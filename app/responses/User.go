package responses

// User act as user response
type User struct {
	Name        string `json:"name"`
	UserName    string `json:"userName"`
	Email       string `json:"email"`
	EmailVerify int64  `json:"emailVerify"`
	Phone       string `json:"phone"`
	PhoneVerify int64  `json:"phoneVerify"`
	DOB         string `json:"dob"`
}
