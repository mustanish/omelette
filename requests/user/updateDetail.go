package userrequest

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/mustanish/omelette/config"
)

// User is used to map input JSON
type User struct {
	Name        string `json:"name"`
	UserName    string `json:"userName"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Age         int64  `json:"age"`
}

// Validate is used to validate input JSON
func (u User) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Name, validation.Length(3, 50).Error("Sorry that is not a valid name must be between 3-50 characters long")),
		validation.Field(&u.UserName, validation.Match(regexp.MustCompile("^[a-z0-9_-]{3,15}$")).Error("Sorry that is not a valid username must be between 3-15 characters long and allowed characters are alphabet, numbers, underscore and hyphen")),
		validation.Field(&u.Email, is.Email.Error("Sorry that is not a valid email address")),
		validation.Field(&u.PhoneNumber, validation.Match(regexp.MustCompile("^"+config.PhoneRegex+"$")).Error("Sorry that is not a valid phone number")),
		validation.Field(&u.Age, validation.Min(13).Error("Sorry that is not a valid age must be greater than 13 years")),
	)
}
