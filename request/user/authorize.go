package userRequest

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/mustanish/jwtAPI/config"
)

// Authorize is used to map input JSON
type Authorize struct {
	Identity string `json:"identity"`
}

// Validate is used to validate input JSON
func (a Authorize) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Identity, validation.Required, validation.Match(regexp.MustCompile("^"+config.EmailRegex+"|"+config.PhoneRegex+"$")).Error("Sorry that is not a valid email ID/phone number")),
	)
}
