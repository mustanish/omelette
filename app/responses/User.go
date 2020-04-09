package responses

// User act as user response
type User struct {
	Name        string `json:"name,omitempty"`
	UserName    string `json:"userName,omitempty"`
	Email       string `json:"email,omitempty"`
	EmailVerify int64  `json:"emailVerify,omitempty"`
	Phone       string `json:"phone,omitempty"`
	PhoneVerify int64  `json:"phoneVerify,omitempty"`
	DOB         string `json:"dob,omitempty"`
}
