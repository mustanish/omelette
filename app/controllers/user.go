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

var model repository.User

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
	user, docKey, err := model.Exist(data.Identity)
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
		_, docKey, err = user.UpdateUser(docKey)
	} else {
		user.CreatedAt = now.Unix()
		docKey, err = user.Authenticate()
	}
	if err != nil {
		render.Render(res, req, responses.NewHTTPError(http.StatusServiceUnavailable, constants.Unavailable))
	} else {
		token, _ := helpers.GenerateToken(docKey, constants.OTPValidity)
		response["message"] = msg
		response["accessToken"] = token
		render.Render(res, req, responses.NewHTTPSucess(http.StatusOK, response))
	}
}

// VerifyUser verifies generated OTP
func VerifyUser(res http.ResponseWriter, req *http.Request) {
	ID := req.Context().Value("ID").(string)
	user, docKey, err := model.Exist(ID)
	if err != nil {
		render.Render(res, req, responses.NewHTTPError(http.StatusServiceUnavailable, constants.Unavailable))
	} else if len(docKey) < 0 {
		render.Render(res, req, responses.NewHTTPError(http.StatusBadRequest, constants.NotFoundResource))
	} else {
		var (
			data       = req.Context().Value("data").(*userschemas.VerifyUser)
			now        = time.Now().Unix()
			response   = make(map[string]interface{})
			timeDiff   = time.Unix(user.OtpValidity, 0).Sub(time.Now())
			comparePwd = bcrypt.CompareHashAndPassword([]byte(user.OTP), []byte(data.OTP))
		)
		if comparePwd != nil || timeDiff < 0 {
			render.Render(res, req, responses.NewHTTPError(http.StatusBadRequest, constants.InvalidOTP))
		} else {
			if user.OtpType == constants.OTPType["email"] {
				user.EmailVerify = 1
			} else if user.OtpType == constants.OTPType["phone"] {
				user.PhoneVerify = 1
			}
			user.OTP, user.OtpType, user.OtpValidity, user.LastLogedIn, user.UpdatedAt = "", "", 0, now, now
			_, _, err = user.UpdateUser(docKey)
			if err != nil {
				render.Render(res, req, responses.NewHTTPError(http.StatusServiceUnavailable, constants.Unavailable))
				return
			}
			accessToken, _ := helpers.GenerateToken(docKey, constants.AccessTokenValidity)
			refreshToken, _ := helpers.GenerateToken(docKey, constants.RefreshTokenValidity)
			response["accessToken"] = accessToken
			response["refreshToken"] = refreshToken
			render.Render(res, req, responses.NewHTTPSucess(http.StatusOK, response))
		}
	}
}

// UpdateUser updates details of an existing user
func UpdateUser(res http.ResponseWriter, req *http.Request) {
}

// DeleteUser blocks an existing user
func DeleteUser(res http.ResponseWriter, req *http.Request) {
}

// GetUser fetches details of an existing user
func GetUser(res http.ResponseWriter, req *http.Request) {
}
