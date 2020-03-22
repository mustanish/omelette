package constants

// EmailRegex is exported to be used globally
const EmailRegex = "[a-z0-9._%+\\-]+@[a-z0-9.\\-]+\\.[a-z]{2,4}"

// MobileRegex is exported to be used globally
const MobileRegex = "[6-9]\\d{9}"

// IdentityRequired is exported to be used globally
const IdentityRequired = "Enter your email/mobile number"

// IdentityRegex is exported to be used globally
const IdentityRegex = "Sorry that is not a valid email/mobile number"

// Failed is exported to be used globally
const Failed = "failed"

// Success is exported to be used globally
const Success = "success"

// InvalidReq is exported to be used globally
const InvalidReq = "Invalid request. Please review your request and try again"

// NotFound is exported to be used globally
const NotFound = "The reuested URL was not found on this server"

// Unavailable is exported to be used globally
const Unavailable = "Unable to service your request. Please try again later"

// MethodNotAllowed is exported to be used globally
const MethodNotAllowed = "The requested method is not allowed for this URL"
