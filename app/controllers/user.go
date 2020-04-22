package controllers

import (
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/go-chi/render"
	"github.com/jinzhu/copier"
	"github.com/mustanish/omelette/app/constants"
	"github.com/mustanish/omelette/app/helpers"
	"github.com/mustanish/omelette/app/repository"
	"github.com/mustanish/omelette/app/responses"
	authschemas "github.com/mustanish/omelette/app/schemas/auth"
	userschemas "github.com/mustanish/omelette/app/schemas/user"
	"golang.org/x/crypto/bcrypt"
)

// Authenticate accepts email/mobile and generate OTP
func Authenticate(res http.ResponseWriter, req *http.Request) {
	var (
		msg         string
		user        repository.User
		data        = req.Context().Value("data").(*authschemas.Authenticate)
		now         = time.Now()
		OTP         = helpers.GenerateOTP(constants.OTPLength)
		hash, _     = bcrypt.GenerateFromPassword([]byte(OTP), bcrypt.MinCost)
		isEmail     = regexp.MustCompile("^" + constants.EmailRegex + "$")
		response    = make(map[string]interface{})
		docKey, err = user.Exist(data.Identity)
	)
	if err != nil {
		render.Render(res, req, responses.NewHTTPError(http.StatusServiceUnavailable, constants.Unavailable))
		return
	}
	if isEmail.MatchString(data.Identity) {
		user.Email = data.Identity
		user.OtpType = constants.OTPType["email"]
		msg = strings.Replace(constants.OTPMsg, "identity", "email "+helpers.MaskEmail(data.Identity), -1)
	} else {
		user.Phone = data.Identity
		user.OtpType = constants.OTPType["phone"]
		msg = strings.Replace(constants.OTPMsg, "identity", "mobile "+helpers.MaskNumber(data.Identity), -1)
	}
	user.OTP = string(hash)
	user.OtpValidity, user.UpdatedAt = now.Add(time.Second*constants.OTPValidity).Unix(), now.Unix()
	if len(docKey) > 0 {
		docKey, err = user.UpdateUser(docKey)
	} else {
		user.CreatedAt = now.Unix()
		docKey, err = user.Authenticate()
	}
	if err != nil {
		render.Render(res, req, responses.NewHTTPError(http.StatusServiceUnavailable, constants.Unavailable))
		return
	}
	tempToken, err := generateToken(docKey, "tempToken", constants.OTPValidity)
	if err != nil {
		render.Render(res, req, responses.NewHTTPError(http.StatusServiceUnavailable, constants.Unavailable))
		return
	}
	response["message"] = msg
	response["accessToken"] = tempToken
	render.Render(res, req, responses.NewHTTPSucess(http.StatusOK, response))
}

// Login verifies OTP and generate access & refresh token
func Login(res http.ResponseWriter, req *http.Request) {
	var (
		user        repository.User
		ID          = req.Context().Value("ID").(string)
		data        = req.Context().Value("data").(*authschemas.Login)
		now         = time.Now().Unix()
		response    = make(map[string]interface{})
		docKey, err = user.Exist(ID)
		timeDiff    = time.Unix(user.OtpValidity, 0).Sub(time.Now())
		compareOtp  = bcrypt.CompareHashAndPassword([]byte(user.OTP), []byte(data.OTP))
	)
	if err != nil {
		render.Render(res, req, responses.NewHTTPError(http.StatusServiceUnavailable, constants.Unavailable))
		return
	} else if len(docKey) == 0 {
		render.Render(res, req, responses.NewHTTPError(http.StatusBadRequest, constants.NotFoundResource))
		return
	}
	if compareOtp != nil || timeDiff < 0 {
		render.Render(res, req, responses.NewHTTPError(http.StatusBadRequest, constants.InvalidOTP))
		return
	}
	if user.OtpType == constants.OTPType["email"] {
		user.EmailVerify = 1
	} else if user.OtpType == constants.OTPType["phone"] {
		user.PhoneVerify = 1
	}
	user.OTP, user.OtpType, user.OtpValidity, user.LastLogedIn, user.UpdatedAt = "", "", 0, now, now
	docKey, err = user.UpdateUser(docKey)
	if err != nil {
		render.Render(res, req, responses.NewHTTPError(http.StatusServiceUnavailable, constants.Unavailable))
		return
	}
	accessToken, err := generateToken(docKey, "aceessToken", constants.AccessTokenValidity)
	if err != nil {
		render.Render(res, req, responses.NewHTTPError(http.StatusServiceUnavailable, constants.Unavailable))
		return
	}
	refreshToken, err := generateToken(docKey, "refreshToken", constants.RefreshTokenValidity)
	if err != nil {
		render.Render(res, req, responses.NewHTTPError(http.StatusServiceUnavailable, constants.Unavailable))
		return
	}
	response["accessToken"] = accessToken
	response["refreshToken"] = refreshToken
	response["expiresIn"] = constants.AccessTokenValidity
	render.Render(res, req, responses.NewHTTPSucess(http.StatusOK, response))
}

// UpdateUser updates details of an existing user
func UpdateUser(res http.ResponseWriter, req *http.Request) {
	var (
		user        repository.User
		response    responses.User
		ID          = req.Context().Value("ID").(string)
		data        = req.Context().Value("data").(*userschemas.UpdateUser)
		now         = time.Now()
		docKey, err = user.Exist(ID)
	)
	if err != nil {
		render.Render(res, req, responses.NewHTTPError(http.StatusServiceUnavailable, constants.Unavailable))
		return
	} else if len(docKey) == 0 {
		render.Render(res, req, responses.NewHTTPError(http.StatusBadRequest, constants.NotFoundResource))
		return
	}
	copier.Copy(&user, data)
	user.UpdatedAt = now.Unix()
	docKey, err = user.UpdateUser(docKey)
	if err != nil {
		render.Render(res, req, responses.NewHTTPError(http.StatusServiceUnavailable, constants.Unavailable))
		return
	}
	response.ID = docKey
	copier.Copy(&response, user)
	render.Render(res, req, responses.NewHTTPSucess(http.StatusOK, response))
}

// GetUser fetches details of an existing user
func GetUser(res http.ResponseWriter, req *http.Request) {
	var (
		user        repository.User
		response    responses.User
		ID          = req.Context().Value("ID").(string)
		docKey, err = user.Exist(ID)
	)
	if err != nil {
		render.Render(res, req, responses.NewHTTPError(http.StatusServiceUnavailable, constants.Unavailable))
		return
	} else if len(docKey) == 0 {
		render.Render(res, req, responses.NewHTTPError(http.StatusBadRequest, constants.NotFoundResource))
		return
	}
	response.ID = docKey
	copier.Copy(&response, user)
	render.Render(res, req, responses.NewHTTPSucess(http.StatusOK, response))
}

// Logout deletes existing access token
func Logout(res http.ResponseWriter, req *http.Request) {
	var (
		token    repository.Token
		tokenID  = req.Context().Value("tokenID").(string)
		response = make(map[string]interface{})
		msg      = constants.Logout
		err      = token.RemoveToken(tokenID)
	)
	if err != nil {
		render.Render(res, req, responses.NewHTTPError(http.StatusServiceUnavailable, constants.Unavailable))
		return
	}
	response["message"] = msg
	render.Render(res, req, responses.NewHTTPSucess(http.StatusOK, response))
}

// RefreshToken refreshes access token
func RefreshToken(res http.ResponseWriter, req *http.Request) {
	var (
		token       repository.Token
		user        repository.User
		ID          = req.Context().Value("ID").(string)
		tokenID     = req.Context().Value("tokenID").(string)
		response    = make(map[string]interface{})
		docKey, err = user.Exist(ID)
	)
	if err != nil {
		render.Render(res, req, responses.NewHTTPError(http.StatusServiceUnavailable, constants.Unavailable))
		return
	} else if len(docKey) == 0 {
		render.Render(res, req, responses.NewHTTPError(http.StatusBadRequest, constants.NotFoundResource))
		return
	}
	err = token.RemoveToken(tokenID)
	if err != nil {
		render.Render(res, req, responses.NewHTTPError(http.StatusServiceUnavailable, constants.Unavailable))
		return
	}
	accessToken, err := generateToken(docKey, "aceessToken", constants.AccessTokenValidity)
	if err != nil {
		render.Render(res, req, responses.NewHTTPError(http.StatusServiceUnavailable, constants.Unavailable))
		return
	}
	refreshToken, err := generateToken(docKey, "refreshToken", constants.RefreshTokenValidity)
	if err != nil {
		render.Render(res, req, responses.NewHTTPError(http.StatusServiceUnavailable, constants.Unavailable))
		return
	}
	response["accessToken"] = accessToken
	response["refreshToken"] = refreshToken
	response["expiresIn"] = constants.AccessTokenValidity
	render.Render(res, req, responses.NewHTTPSucess(http.StatusOK, response))
}

func generateToken(docKey string, tokenType string, validity time.Duration) (string, error) {
	var (
		token                                 repository.Token
		tokenString, tokenID, tokenExpires, _ = helpers.GenerateToken(docKey, validity)
	)
	token.Key = tokenID
	token.Type = tokenType
	token.ExpiresAt = tokenExpires
	err := token.AddToken()
	if err != nil {
		return tokenString, err
	}
	return tokenString, nil
}
