package userschemas

import "github.com/mustanish/omelette/app/constants"

// UpdateUser maps / route to update profile
type UpdateUser struct {
	Name     string `json:"name"`
	UserName string `json:"userName"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	DOB      string `json:"dob"`
}

var (
	updateUser      UpdateUser
	updateUserRules = map[string][]string{
		"name":     {"alpha", "between:3,50"},
		"userName": {"alpha_dash", "between:3,50"},
		"email":    {"regex:^" + constants.EmailRegex + "$"},
		"phone":    {"regex:^" + constants.MobileRegex + "$"},
		"dob":      {"date"},
	}
	updateUserMessages = map[string][]string{
		"email": {"regex:" + constants.IdentityRegex},
		"phone": {"regex:" + constants.IdentityRegex},
	}
	// UpdateUserOpts represents validation options for middleware
	UpdateUserOpts = map[string]interface{}{
		"data":     &updateUser,
		"messages": updateUserMessages,
		"rules":    updateUserRules,
	}
)
