package controllers

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/mustanish/omelette/config"
	"github.com/mustanish/omelette/helpers"
	"github.com/mustanish/omelette/models"
	userRequest "github.com/mustanish/omelette/request/user"
	"github.com/mustanish/omelette/response"
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
		errorResponse.Error = config.InvalidRequest
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
	user.OtpValidity = time.Now().Unix()
	if identity.MatchString(authorize.Identity) {
		user.Email = &authorize.Identity
		user.OtpType = "authorizeEmail"
		successResponse.Msg = strings.Replace(config.AuthorizeMsg, "identity", "email "+helpers.MaskEmail(authorize.Identity), -1)
	} else {
		user.PhoneNumber = &authorize.Identity
		user.OtpType = "authorizePhone"
		successResponse.Msg = strings.Replace(config.AuthorizeMsg, "identity", "mobile "+helpers.MaskNumber(authorize.Identity), -1)
	}
	if user.ID > 0 {
		error = userModel.UpdateDetail()
	} else {
		user, error = userModel.Authorize()
	}
	token, error := helpers.GenerateToken(user.ID)
	if error != nil {
		errorResponse.Code = http.StatusServiceUnavailable
		errorResponse.Error = config.ServiceUnavailable
		helpers.SetResponse(res, http.StatusServiceUnavailable, errorResponse)
		return
	}
	response := make(map[string]interface{})
	response["authorizationToken"] = token
	successResponse.Code = http.StatusOK
	successResponse.Data = response
	helpers.SetResponse(res, http.StatusOK, successResponse)
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
		errorResponse.Error = config.InvalidRequest
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
	user := userModel.Check(req.Context().Value("identity"))
	//log.Println(user)
	otp, _ := strconv.ParseInt(verfiy.Otp, 10, 64)
	if user.ID < 0 || user.OTP != otp || !helpers.InArray(user.OtpType, []string{"authorizeEmail", "authorizePhone", "verifyEmail", "verifyPhone"}) {
		errorResponse.Code = http.StatusBadRequest
		errorResponse.Error = config.InvalidToken
		helpers.SetResponse(res, http.StatusBadRequest, errorResponse)
		return
	}
	switch user.OtpType {
	case "authorizeEmail":
		user.EmailVerify = 1
		user.LastLogedIn = time.Now()
		successResponse.Msg = config.LoggedInMsg
	case "authorizePhone":
		user.PhoneVerify = 1
		user.LastLogedIn = time.Now()
		successResponse.Msg = config.LoggedInMsg
	case "verifyEmail":
		user.EmailVerify = 1
		successResponse.Msg = config.EmailVerifyMsg
	case "verifyPhone":
		user.PhoneVerify = 1
		successResponse.Msg = config.PhoneVerifyMsg
	}
	user.OTP, user.OtpType, user.OtpValidity = 0, "", 0
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

// ReadDetail is used to fetch detail of requested user
func ReadDetail(res http.ResponseWriter, req *http.Request) {
}

// UpdateDetail is used to update detail of requested user
func UpdateDetail(res http.ResponseWriter, req *http.Request) {
}
