package dto

import (
	"time"

	"github.com/google/uuid"
)

type GetPostsInput struct {
	AuthorID uuid.UUID `json:"-"`
	Sort     string    `json:"-" validate:"oneof=created_at updated_at"`
	Order    string    `json:"-" validate:"oneof=asc desc"`
	Offset   int       `json:"-" validate:"min=0"`
	Limit    int       `json:"-" validate:"min=1,max=100"`
}

func (i *GetPostsInput) Validate() error {

	return validate.Struct(i)
}

type GetPostsOutput struct {
	Posts []Post `json:"posts"`
	Total int    `json:"total"`
}

type Post struct {
	ID        uuid.UUID `json:"id"`
	AuthorID  uuid.UUID `json:"author_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
