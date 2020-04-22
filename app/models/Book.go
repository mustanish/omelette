package models

// Book act as model for database
type Book struct {
	Name       string  `json:"name,omitempty"`
	Genre      string  `json:"genre,omitempty"`
	AuthorName string  `json:"authorName,omitempty"`
	CoverImage string  `json:"coverImage,omitempty"`
	Price      float64 `json:"price,omitempty"`
	AddedBy    string  `json:"addedBy,omitempty"`
	Deleted    int64   `json:"deleted,omitempty"`
	CreatedAt  int64   `json:"createdAt,omitempty"`
	UpdatedAt  int64   `json:"updatedAt,omitempty"`
	DeletedAt  int64   `json:"deletedAt,omitempty"`
}
