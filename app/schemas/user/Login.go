package userschemas

import "github.com/mustanish/omelette/app/constants"

// Login maps /login route
type Login struct {
	OTP string `json:"otp"`
}

var (
	login Login

	loginRules = map[string][]string{
		"otp": {"required", "digits:6"},
	}
	loginMessages = map[string][]string{
		"otp": {"required:" + constants.OTPRequired, "len:" + constants.OTPMaxLen},
	}
	// LoginOpts represents validation options for middleware
	LoginOpts = map[string]interface{}{
		"data":     &login,
		"messages": loginMessages,
		"rules":    loginRules,
	}
)
