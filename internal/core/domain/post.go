package domain

import "time"

type Post struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	AuthorID uint   `json:"author_id"`
	// TODO: Update by User type
	Author string `json:"author"`
	// TODO: Update by Tag type
	Tags      []string  `json:"tags"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PostCreate struct {
	Title    string `json:"title" validate:"required,min=3"`
	Content  string `json:"content" validate:"required,min=10"`
	AuthorID uint   `json:"author_id" validate:"required"`
	// TODO: Update by Tag type
	Tags []string `json:"tags"`
}
