package userRequest

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation"
)

type Verify struct {
	Otp   string `json:"otp"`
	Event string `json:"event"`
	Token string `json:"token"`
}

// Validate is used to validate input JSON
func (v Verify) Validate() error {
	return validation.ValidateStruct(&v,
		validation.Field(&v.Otp, validation.Required, validation.Length(6, 6).Error("Sorry that is not a valid OTP")),
		validation.Field(&v.Event, validation.Required, validation.Match(regexp.MustCompile("^(?:email|phone|authorize)$")).Error("Sorry that is not a valid event")),
		validation.Field(&v.Token, validation.Required.Error("Sorry that is not a valid token")),
	)
}
