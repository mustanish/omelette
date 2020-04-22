package authschemas

import "github.com/mustanish/omelette/app/constants"

// Login maps /login route
type Login struct {
	OTP string `json:"otp"`
}

var (
	loginRules = map[string][]string{
		"otp": {"required", "digits:6"},
	}
	loginCheck    = []string{"body"}
	loginMessages = map[string][]string{
		"otp": {"required:" + constants.OTPRequired, "len:" + constants.OTPMaxLen},
	}
	// LoginOpts represents validation options for middleware
	LoginOpts = map[string]interface{}{
		"messages": loginMessages,
		"rules":    loginRules,
		"check":    loginCheck,
	}
)
