package models

// User act as model for database
type User struct {
	Name        string `json:"name,omitempty"`
	UserName    string `json:"userName,omitempty"`
	Email       string `json:"email,omitempty"`
	EmailVerify int64  `json:"emailVerify,omitempty"`
	Phone       string `json:"phone,omitempty"`
	PhoneVerify int64  `json:"phoneVerify,omitempty"`
	DOB         string `json:"dob,omitempty"`
	OTP         string `json:"otp,omitempty"`
	OtpType     string `json:"otpType,omitempty"`
	OtpValidity int64  `json:"otpValidity,omitempty"`
	LastLogedIn int64  `json:"lastLoggedIn,omitempty"`
	CreatedAt   int64  `json:"createdAt,omitempty"`
	UpdatedAt   int64  `json:"updatedAt,omitempty"`
	DeletedAt   int64  `json:"deletedAt,omitempty"`
}
