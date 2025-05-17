package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreatePostInput struct {
	AuthorID    uuid.UUID `json:"author_id" validate:"required"`
	Content     string    `json:"content" validate:"required,min=1,max=5000"`
	Visibility  string    `json:"visibility" validate:"oneof=public friends private"`
	Attachments []string  `json:"attachments"`
}

type CreatePostOutput struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
}
