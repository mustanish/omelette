package userschemas

import "github.com/mustanish/omelette/app/constants"

// VerifyUser maps /verify route
type VerifyUser struct {
	OTP string `json:"otp"`
}

var (
	verifyUser VerifyUser

	verifyUserRules = map[string][]string{
		"otp": {"required", "digits:6"},
	}
	verifyUserMessages = map[string][]string{
		"otp": {"required:" + constants.OTPRequired, "len:" + constants.OTPMaxLen},
	}
	// VerifyUserOpts represents validation options for middleware
	VerifyUserOpts = map[string]interface{}{
		"data":     &verifyUser,
		"messages": verifyUserMessages,
		"rules":    verifyUserRules,
	}
)
