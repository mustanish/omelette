package responses

// Book act as book response
type Book struct {
	ID         string  `json:"id,omitempty"`
	Name       string  `json:"name,omitempty"`
	Genre      string  `json:"genre,omitempty"`
	AuthorName string  `json:"authorName,omitempty"`
	AddedBy    string  `json:"addedBy,omitempty"`
	Price      float64 `json:"price,omitempty"`
	CoverImage string  `json:"coverImage,omitempty"`
	TotalCount float64 `json:"totalCount,omitempty"`
}
