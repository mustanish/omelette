package userSchemas

import "github.com/mustanish/omelette/app/constants"

type authenticate struct {
	Identity string `json:"identity"`
}

var (
	data authenticate

	rules = map[string][]string{
		"identity": {"required", "regex:^" + constants.EmailRegex + "|" + constants.MobileRegex + "$"},
	}

	messages = map[string][]string{
		"identity": {"required:" + constants.IdentityRequired, "regex:" + constants.IdentityRegex},
	}

	// Authenticate is exported to be in validation middleware
	Authenticate = map[string]interface{}{
		"data":     &data,
		"rules":    rules,
		"messages": messages,
	}
)
