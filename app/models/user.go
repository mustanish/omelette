package models

// User act as
type User struct {
	Name        string `json:"name"`
	UserName    string `json:"userName"`
	Email       string `json:"email"`
	EmailVerify int64  `json:"emailVerify"`
	Phone       string `json:"phone"`
	PhoneVerify int64  `json:"phoneVerify"`
	Age         int64  `json:"age"`
	OTP         string `json:"otp"`
	OtpType     string `json:"otpType"`
	OtpValidity int64  `json:"otpValidity"`
	LastLogedIn int64  `json:"lastLoggedIn"`
	CreatedAt   int64  `json:"createdAt"`
	UpdatedAt   int64  `json:"updatedAt"`
	DeletedAt   int64  `json:"deletedAt"`
}
