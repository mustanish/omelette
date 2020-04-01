package constants

// EmailRegex represents regex to validate email
const EmailRegex = "[a-z0-9._%+\\-]+@[a-z0-9.\\-]+\\.[a-z]{2,4}"

// MobileRegex represents regex to validate mobile number
const MobileRegex = "[6-9]\\d{9}"

// OTPLength represents max length of OTP
const OTPLength = 6

// OTPValidity represents validity of OTP in seconds
const OTPValidity = 150

// OTPType represents types of OTP
var OTPType = map[string]string{"email": "verifyEmail", "phone": "verifyPhone"}

// AccessTokenValidity represents access token validity in seconds
const AccessTokenValidity = 3600

// RefreshTokenValidity represents refresh token validity in seconds
const RefreshTokenValidity = 604800

// Jwtsecret represents secret to create token
const Jwtsecret = "mYq3t6w9z$C&F)J@NcRfUjXn2r4u7x!A%D*G-KaPdSgVkYp3s6v8y/B?E(H+MbQeThWmZq4t7w!z$C&F)J@NcRfUjXn2r5u8x/A?D*G-KaPdSgVkYp3s6v9y$B&E)H+M"
