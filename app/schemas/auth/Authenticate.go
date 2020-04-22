package authschemas

import "github.com/mustanish/omelette/app/constants"

// Authenticate maps /auth route
type Authenticate struct {
	Identity string `json:"identity"`
}

var (
	authenticateRules = map[string][]string{
		"identity": {"required", "regex:^" + constants.EmailRegex + "|" + constants.MobileRegex + "*$"},
	}
	authenticateCheck    = []string{"body"}
	authenticateMessages = map[string][]string{
		"identity": {"required:" + constants.IdentityRequired, "regex:" + constants.IdentityRegex},
	}
	// AuthenticateOpts represents validation options for middleware
	AuthenticateOpts = map[string]interface{}{
		"messages": authenticateMessages,
		"rules":    authenticateRules,
		"check":    authenticateCheck,
	}
)
