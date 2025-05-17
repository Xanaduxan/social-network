package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID          uuid.UUID
	AuthorID    uuid.UUID
	Content     string
	Visibility  string
	Attachments []string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type PostStats struct {
	PostID    uuid.UUID
	Likes     int
	Views     int
	Shares    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewPost(authorID uuid.UUID, content, visibility string, attachments []string) (*Post, error) {
	if authorID == uuid.Nil {
		return nil, errors.New("author ID cannot be empty")
	}
	if content == "" {
		return nil, errors.New("content cannot be empty")
	}
	if !isValidVisibility(visibility) {
		return nil, errors.New("invalid visibility value")
	}

	return &Post{
		ID:          uuid.New(),
		AuthorID:    authorID,
		Content:     content,
		Visibility:  visibility,
		Attachments: attachments,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}, nil
}

func NewPostStats(postID uuid.UUID) *PostStats {
	return &PostStats{
		PostID:    postID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
}

func isValidVisibility(v string) bool {
	switch v {
	case "public", "friends", "private":
		return true
	default:
		return false
	}
}
