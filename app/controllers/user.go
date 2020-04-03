package controllers

import (
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/go-chi/render"
	"github.com/mustanish/omelette/app/constants"
	"github.com/mustanish/omelette/app/helpers"
	"github.com/mustanish/omelette/app/repository"
	"github.com/mustanish/omelette/app/responses"
	userschemas "github.com/mustanish/omelette/app/schemas/user"
	"golang.org/x/crypto/bcrypt"
)

var (
	user  = new(repository.User)
	token = new(repository.Token)
)

// Authenticate validates email/mobile via OTP
func Authenticate(res http.ResponseWriter, req *http.Request) {
	var (
		data     = req.Context().Value("data").(*userschemas.Authenticate)
		now      = time.Now()
		OTP      = helpers.GenerateOTP(constants.OTPLength)
		hash, _  = bcrypt.GenerateFromPassword([]byte(OTP), bcrypt.MinCost)
		isEmail  = regexp.MustCompile("^" + constants.EmailRegex + "$")
		msg      string
		response = make(map[string]interface{})
	)
	log.Println(OTP)
	docKey, err := user.Exist(data.Identity)
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
	token, _, _, _ := helpers.GenerateToken(docKey, constants.OTPValidity)
	response["message"] = msg
	response["accessToken"] = token
	render.Render(res, req, responses.NewHTTPSucess(http.StatusOK, response))
}

// VerifyUser verifies generated OTP
func VerifyUser(res http.ResponseWriter, req *http.Request) {
	ID := req.Context().Value("ID").(string)
	docKey, err := user.Exist(ID)
	if err != nil {
		render.Render(res, req, responses.NewHTTPError(http.StatusServiceUnavailable, constants.Unavailable))
		return
	} else if len(docKey) < 0 {
		render.Render(res, req, responses.NewHTTPError(http.StatusBadRequest, constants.NotFoundResource))
		return
	}
	var (
		data       = req.Context().Value("data").(*userschemas.VerifyUser)
		now        = time.Now().Unix()
		response   = make(map[string]interface{})
		timeDiff   = time.Unix(user.OtpValidity, 0).Sub(time.Now())
		comparePwd = bcrypt.CompareHashAndPassword([]byte(user.OTP), []byte(data.OTP))
	)
	if comparePwd != nil || timeDiff < 0 {
		render.Render(res, req, responses.NewHTTPError(http.StatusBadRequest, constants.InvalidOTP))
		return
	}
	if user.OtpType == constants.OTPType["email"] {
		user.EmailVerify = 1
	} else if user.OtpType == constants.OTPType["phone"] {
		user.PhoneVerify = 1
	}
	user.OTP, user.OtpType, user.OtpValidity, user.LastLogedIn, user.UpdatedAt = "", "", 0, now, now
	_, err = user.UpdateUser(docKey)
	if err != nil {
		render.Render(res, req, responses.NewHTTPError(http.StatusServiceUnavailable, constants.Unavailable))
		return
	}
	accessToken, accessID, accessExpires, _ := helpers.GenerateToken(docKey, constants.AccessTokenValidity)
	token.Key = accessID
	token.Type = "aceessToken"
	token.ExpiresAt = accessExpires
	err = token.AddToken()
	if err != nil {
		render.Render(res, req, responses.NewHTTPError(http.StatusServiceUnavailable, constants.Unavailable))
		return
	}
	refreshToken, refreshID, refreshExpires, _ := helpers.GenerateToken(docKey, constants.RefreshTokenValidity)
	token.Key = refreshID
	token.Type = "refreshToken"
	token.ExpiresAt = refreshExpires
	err = token.AddToken()
	if err != nil {
		render.Render(res, req, responses.NewHTTPError(http.StatusServiceUnavailable, constants.Unavailable))
		return
	}
	response["accessToken"] = accessToken
	response["refreshToken"] = refreshToken
	render.Render(res, req, responses.NewHTTPSucess(http.StatusOK, response))
}

// UpdateUser updates details of an existing user
func UpdateUser(res http.ResponseWriter, req *http.Request) {
}

// GetUser fetches details of an existing user
func GetUser(res http.ResponseWriter, req *http.Request) {
}

// Logout logs out existing user
func Logout(res http.ResponseWriter, req *http.Request) {
}
