package response

// Error is used set error
type Error struct {
	Code  int         `json:"code"`
	Error interface{} `json:"error"`
}
