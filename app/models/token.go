package models

// Token act as model for database
type Token struct {
	Key       string `json:"_key"`
	Type      string `json:"type"`
	ExpiresAt int64  `json:"expiresAt"`
}
