package userRequest

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type Verify struct {
	Otp   string `json:"otp"`
	Token string `json:"token"`
}

// Validate is used to validate input JSON
func (v Verify) Validate() error {
	return validation.ValidateStruct(&v,
		validation.Field(&v.Otp, validation.Required, validation.Length(6, 6).Error("Sorry that is not a valid OTP")),
	)
}
