package constants

// EmailRegex represents regex to validate email
const EmailRegex = "[a-z0-9._%+\\-]+@[a-z0-9.\\-]+\\.[a-z]{2,4}"

// MobileRegex represents regex to validate mobile number
const MobileRegex = "[6-9]\\d{9}"

// OTPLength represents max length of OTP
const OTPLength = 6

// OTPValidity represents validity of OTP in seconds
const OTPValidity = 300

// OTPType represents types of OTP
var OTPType = map[string]string{"email": "verifyEmail", "phone": "verifyPhone"}

// AccessTokenValidity represents access token validity in seconds
const AccessTokenValidity = 3600

// RefreshTokenValidity represents refresh token validity in seconds
const RefreshTokenValidity = 604800

// Jwtsecret represents secret to create token
const Jwtsecret = "^hBlHu3pSCX@_KQ'JSX*6I*CX^brqM=@2nPIU*LSc~;LLwFG-Fk1-3F6WDT][5U"
