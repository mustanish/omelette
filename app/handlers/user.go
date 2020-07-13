package handlers

import (
	"log"
	"net/http"
	"omelette/app/connectors"
	"omelette/app/constants"
	"omelette/app/responses"
	"omelette/app/schemas/model"
	"omelette/app/schemas/validation"
	"omelette/helpers"
	"regexp"
	"strings"
	"time"

	"github.com/go-chi/render"
	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

// Authenticate accepts email/mobile and generate OTP
func Authenticate(res http.ResponseWriter, req *http.Request) {
	var (
		query strings.Builder
		msg   string
	)
	user := new(model.User)
	data := req.Context().Value("data").(*validation.Authenticate)
	otpHash, _ := bcrypt.GenerateFromPassword([]byte(helpers.GenerateOTP(constants.OTPLength)), bcrypt.MinCost)
	isEmail := regexp.MustCompile("^" + constants.EmailRegex + "$")
	bindVars := make(map[string]interface{})
	response := make(map[string]interface{})
	if isEmail.MatchString(data.Identity) {
		query.WriteString("UPSERT { email:@identity } INSERT { email:@identity, ")
		bindVars["type"] = constants.OTPType["email"]
		msg = strings.Replace(constants.OTPMsg, "identity", "email "+helpers.MaskEmail(data.Identity), -1)
	} else {
		query.WriteString("UPSERT { phone:@identity } INSERT { phone:@identity, ")
		bindVars["type"] = constants.OTPType["phone"]
		msg = strings.Replace(constants.OTPMsg, "identity", "mobile "+helpers.MaskNumber(data.Identity), -1)
	}
	query.WriteString("otp:@otp, otpType:@type, otpValidity:ROUND(DATE_TIMESTAMP(DATE_ADD(DATE_NOW(),@validity,'seconds'))/1000), createdAt:ROUND(DATE_NOW()/1000) }")
	query.WriteString("UPDATE { otp:@otp, otpType:@type, otpValidity:ROUND(DATE_TIMESTAMP(DATE_ADD(DATE_NOW(),@validity,'seconds'))/1000), updatedAt: ROUND(DATE_NOW()/1000) } IN users RETURN NEW")
	bindVars["identity"] = data.Identity
	bindVars["otp"] = string(otpHash)
	bindVars["validity"] = constants.OTPValidity
	cursor, err := connectors.QueryDocument(query.String(), bindVars)
	if err != nil {
		render.Render(res, req, responses.NewHTTPError(http.StatusServiceUnavailable, constants.Unavailable))
		return
	}
	meta, err := cursor.ReadDocument(nil, user)
	if err != nil {
		log.Println("FAILED::could not read user document because of", err.Error())
	}
	tempToken, err := generateToken(meta.Key, "tempToken", constants.OTPValidity)
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
	user := new(model.User)
	ID := req.Context().Value("ID").(string)
	data := req.Context().Value("data").(*validation.Login)
	now := time.Now().Unix()
	response := make(map[string]interface{})
	docKey, err := connectors.ReadDocument("users", ID, user)
	if err != nil {
		render.Render(res, req, responses.NewHTTPError(http.StatusServiceUnavailable, constants.Unavailable))
		return
	} else if len(docKey) == 0 {
		render.Render(res, req, responses.NewHTTPError(http.StatusBadRequest, constants.NotFoundResource))
		return
	}
	timeDiff := time.Unix(user.OtpValidity, 0).Sub(time.Now())
	compareOtp := bcrypt.CompareHashAndPassword([]byte(user.OTP), []byte(data.OTP))
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
	docKey, err = connectors.UpdateDocument("users", docKey, user)
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
	render.Render(res, req, responses.NewHTTPSucess(http.StatusOK, response))
}

// GetUser fetches details of an existing user
func GetUser(res http.ResponseWriter, req *http.Request) {
	var response responses.User
	user := new(model.User)
	ID := req.Context().Value("ID").(string)
	docKey, err := connectors.ReadDocument("users", ID, user)
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

// UpdateUser updates details of an existing user
func UpdateUser(res http.ResponseWriter, req *http.Request) {
	var response responses.User
	user := new(model.User)
	ID := req.Context().Value("ID").(string)
	data := req.Context().Value("data").(*validation.UpdateUser)
	now := time.Now()
	docKey, err := connectors.ReadDocument("users", ID, user)
	if err != nil {
		render.Render(res, req, responses.NewHTTPError(http.StatusServiceUnavailable, constants.Unavailable))
		return
	} else if len(docKey) == 0 {
		render.Render(res, req, responses.NewHTTPError(http.StatusBadRequest, constants.NotFoundResource))
		return
	}
	copier.Copy(&user, data)
	user.UpdatedAt = now.Unix()
	docKey, err = connectors.UpdateDocument("users", docKey, user)
	if err != nil {
		render.Render(res, req, responses.NewHTTPError(http.StatusServiceUnavailable, constants.Unavailable))
		return
	}
	response.ID = docKey
	copier.Copy(&response, user)
	render.Render(res, req, responses.NewHTTPSucess(http.StatusOK, response))
}

// Logout deletes existing access token
func Logout(res http.ResponseWriter, req *http.Request) {
	tokenID := req.Context().Value("tokenID").(string)
	response := make(map[string]interface{})
	msg := constants.Logout
	_, err := connectors.RemoveDocument("tokens", tokenID)
	if err != nil {
		render.Render(res, req, responses.NewHTTPError(http.StatusServiceUnavailable, constants.Unavailable))
		return
	}
	response["message"] = msg
	render.Render(res, req, responses.NewHTTPSucess(http.StatusOK, response))
}

// RefreshToken refreshes access token
func RefreshToken(res http.ResponseWriter, req *http.Request) {
	ID := req.Context().Value("ID").(string)
	tokenID := req.Context().Value("tokenID").(string)
	response := make(map[string]interface{})
	docKey, err := connectors.ReadDocument("users", ID, nil)
	if err != nil {
		render.Render(res, req, responses.NewHTTPError(http.StatusServiceUnavailable, constants.Unavailable))
		return
	} else if len(docKey) == 0 {
		render.Render(res, req, responses.NewHTTPError(http.StatusBadRequest, constants.NotFoundResource))
		return
	}
	_, err = connectors.RemoveDocument("tokens", tokenID)
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
	render.Render(res, req, responses.NewHTTPSucess(http.StatusOK, response))
}

func generateToken(docKey string, tokenType string, validity time.Duration) (string, error) {
	var token model.Token
	tokenString, tokenID, tokenExpires, _ := helpers.GenerateToken(docKey, validity)
	token.Key = tokenID
	token.Type = tokenType
	token.ExpiresAt = tokenExpires
	_, err := connectors.CreateDocument("tokens", token)
	if err != nil {
		return tokenString, err
	}
	return tokenString, nil
}
