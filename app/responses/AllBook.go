package responses

// AllBook act as Allbook response
type AllBook struct {
	Books      []Book         `json:"books,omitempty"`
	Pagination map[string]int `json:"pagination,omitempty"`
}
