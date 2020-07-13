package validation

import "omelette/app/constants"

// UpdateUser maps /user route to update profile
type UpdateUser struct {
	Name     string `json:"name"`
	UserName string `json:"userName"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	DOB      string `json:"dob"`
}

var (
	updateUserRules = map[string][]string{
		"name":     {"alpha_space", "between:3,50"},
		"userName": {"alpha_dash", "between:3,50"},
		"email":    {"regex:^" + constants.EmailRegex + "$"},
		"phone":    {"regex:^" + constants.MobileRegex + "*$"},
		"dob":      {"date"},
	}
	updateUserCheck    = []string{"body"}
	updateUserMessages = map[string][]string{
		"email": {"regex:" + constants.IdentityRegex},
		"phone": {"regex:" + constants.IdentityRegex},
	}
	// UpdateUserOpts represents validation options for middleware
	UpdateUserOpts = map[string]interface{}{
		"messages": updateUserMessages,
		"rules":    updateUserRules,
		"check":    updateUserCheck,
	}
)
