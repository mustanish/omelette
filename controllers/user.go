package controllers

import (
	"encoding/json"
	"math"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/mustanish/jwtAPI/config"
	"github.com/mustanish/jwtAPI/helpers"
	"github.com/mustanish/jwtAPI/models"
	userRequest "github.com/mustanish/jwtAPI/request/user"
	"github.com/mustanish/jwtAPI/response"
	"github.com/myesui/uuid"
)

// Authorize is used to authorize a user to use our application Implement Diffie-Hellman
func Authorize(res http.ResponseWriter, req *http.Request) {

	var (
		userModel       models.User
		errorResponse   response.Error
		successResponse response.Success
		identity        = regexp.MustCompile("^" + config.EmailRegex + "$")
		authorize       userRequest.Authorize
		error           error
	)
	error = json.NewDecoder(req.Body).Decode(&authorize)
	if error != nil {
		errorResponse.Code = http.StatusBadRequest
		errorResponse.Error = config.InvalidJSON
		helpers.SetResponse(res, http.StatusBadRequest, errorResponse)
		return
	}
	error = authorize.Validate()
	if error != nil {
		errorResponse.Code = http.StatusBadRequest
		errorResponse.Error = helpers.FormatError(error.Error())
		helpers.SetResponse(res, http.StatusBadRequest, errorResponse)
		return
	}
	user := userModel.Check(authorize.Identity)
	user.OTP, _ = strconv.ParseInt(helpers.GenerateOtp(6), 10, 64)
	user.Token = uuid.NewV4().String()
	user.TokenValidity = time.Now().Unix()
	if identity.MatchString(authorize.Identity) {
		user.Email = &authorize.Identity
		successResponse.Msg = strings.Replace(config.AuthorizeMsg, "identity", "email "+helpers.MaskEmail(authorize.Identity), -1)
	} else {
		user.PhoneNumber = &authorize.Identity
		successResponse.Msg = strings.Replace(config.AuthorizeMsg, "identity", "mobile "+helpers.MaskNumber(authorize.Identity), -1)
	}
	if user.ID > 0 {
		error = userModel.UpdateDetail()
	} else {
		error = userModel.Authorize()
	}
	if error != nil {
		errorResponse.Code = http.StatusServiceUnavailable
		errorResponse.Error = config.ServiceUnavailable
		helpers.SetResponse(res, http.StatusServiceUnavailable, errorResponse)
		return
	}
	response := make(map[string]interface{})
	response["authorizationToken"] = user.Token
	successResponse.Code = http.StatusOK
	successResponse.Data = response
	helpers.SetResponse(res, http.StatusOK, successResponse)
}

// Authenticate is used to register/login to our application
func Authenticate(res http.ResponseWriter, req *http.Request) {
}

// Verify is used to verify OTP
func Verify(res http.ResponseWriter, req *http.Request) {
	var (
		userModel       models.User
		errorResponse   response.Error
		successResponse response.Success
		verfiy          userRequest.Verify
		error           error
	)
	error = json.NewDecoder(req.Body).Decode(&verfiy)
	if error != nil {
		errorResponse.Code = http.StatusBadRequest
		errorResponse.Error = config.InvalidJSON
		helpers.SetResponse(res, http.StatusBadRequest, errorResponse)
		return
	}
	error = verfiy.Validate()
	if error != nil {
		errorResponse.Code = http.StatusBadRequest
		errorResponse.Error = helpers.FormatError(error.Error())
		helpers.SetResponse(res, http.StatusBadRequest, errorResponse)
		return
	}
	user := userModel.Check(verfiy.Token)
	elapsed := time.Since(time.Unix(user.TokenValidity, 0))
	otp, _ := strconv.ParseInt(verfiy.Otp, 10, 64)
	if user.ID < 0 || math.Round(elapsed.Seconds()) > 120 || user.OTP != otp {
		errorResponse.Code = http.StatusBadRequest
		errorResponse.Error = config.InvalidToken
		helpers.SetResponse(res, http.StatusBadRequest, errorResponse)
		return
	}
	switch verfiy.Event {
	case "authorize":
		user.LastLogedIn = time.Now()
		successResponse.Msg = config.LoggedInMsg
	case "email":
		user.EmailVerify = 1
		successResponse.Msg = config.EmailVerifyMsg
	case "phone":
		user.PhoneVerify = 1
		successResponse.Msg = config.PhoneVerifyMsg
	}
	user.OTP, user.Token, user.TokenValidity = 0, "", 0
	error = userModel.UpdateDetail()
	if error != nil {
		errorResponse.Code = http.StatusServiceUnavailable
		errorResponse.Error = config.ServiceUnavailable
		helpers.SetResponse(res, http.StatusServiceUnavailable, errorResponse)
		return
	}
	successResponse.Code = http.StatusOK
	helpers.SetResponse(res, http.StatusOK, successResponse)
}

// Resend is used to send OTP via email/sms
func Resend(res http.ResponseWriter, req *http.Request) {
}

// ReadDetail is used to fetch detail of requested user
func ReadDetail(res http.ResponseWriter, req *http.Request) {
}

// UpdateDetail is used to update detail of requested user
func UpdateDetail(res http.ResponseWriter, req *http.Request) {
}
