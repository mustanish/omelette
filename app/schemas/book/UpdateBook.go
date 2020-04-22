package bookschemas

// UpdateBook maps /book/{id} route to update an existing book
type UpdateBook struct {
	Name       string  `json:"name"`
	Genre      string  `json:"genre"`
	AuthorName string  `json:"authorName"`
	Price      float64 `json:"price"`
	CoverImage string  `json:"coverImage"`
}

var (
	updateBookRules = map[string][]string{
		"name":       {"alpha_space", "between:3,50"},
		"genre":      {"alpha_space", "between:3,50"},
		"authorName": {"alpha_space", "between:3,50"},
		"price":      {"float"},
		"coverImage": {"url"},
	}
	updateBookCheck = []string{"body"}
	// UpdateBookOpts represents validation options for middleware
	UpdateBookOpts = map[string]interface{}{
		"rules": updateBookRules,
		"check": updateBookCheck,
	}
)
