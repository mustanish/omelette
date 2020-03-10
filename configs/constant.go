package config

// EmailRegex is exported to be used globally
const EmailRegex = "[a-z0-9._%+\\-]+@[a-z0-9.\\-]+\\.[a-z]{2,4}"

// PhoneRegex is exported to be used globally
const PhoneRegex = "[6-9]\\d{9}"

// InvalidRequest is exported to be used globally
const InvalidRequest = "Sorry that is an invalid request"

// ServiceUnavailable is exported to be used globally
const ServiceUnavailable = "Unable to service your request. Please try again later"

// AlreadyTaken is exported to be used globally
const AlreadyTaken = "Sorry identity already taken, please chosse another"

// AuthorizeMsg is exported to be used globally
const AuthorizeMsg = "One Time Password (OTP) has been sent to your identity, please enter the same here to login"

// InvalidToken is exported to be used globally
const InvalidToken = "Sorry either the token has expired or it is invalid"

// NotFound is exported to be used globally
const NotFound = "The requested url was not found on this server"

// LoggedInMsg is exported to be used globally
const LoggedInMsg = "You have successfully logged in"

// EmailVerifyMsg is exported to be used globally
const EmailVerifyMsg = "You have successfully verified your email"

// PhoneVerifyMsg is exported to be used globally
const PhoneVerifyMsg = "You have successfully verified your phone number"

// Jwtsecret is exported to be used globally
const Jwtsecret = "mYq3t6w9z$C&F)J@NcRfUjXn2r4u7x!A%D*G-KaPdSgVkYp3s6v8y/B?E(H+MbQeThWmZq4t7w!z$C&F)J@NcRfUjXn2r5u8x/A?D*G-KaPdSgVkYp3s6v9y$B&E)H+M"
