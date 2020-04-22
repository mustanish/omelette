package bookschemas

// AddBook maps /book route to add new book
type AddBook struct {
	Name       string  `json:"name"`
	Genre      string  `json:"genre"`
	AuthorName string  `json:"authorName"`
	Price      float64 `json:"price"`
	CoverImage string  `json:"coverImage"`
}

var (
	addBookRules = map[string][]string{
		"name":       {"required", "alpha_space", "between:3,50"},
		"genre":      {"required", "alpha_space", "between:3,50"},
		"authorName": {"required", "alpha_space", "between:3,50"},
		"price":      {"required", "float"},
		"coverImage": {"url"},
	}
	addBookCheck = []string{"body"}
	// AddBookOpts represents validation options for middleware
	AddBookOpts = map[string]interface{}{
		"rules": addBookRules,
		"check": addBookCheck,
	}
)
