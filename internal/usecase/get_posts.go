package usecase

import (
	"context"
	"fmt"

	"github.com/okarpova/my-app/internal/domain"
	"github.com/okarpova/my-app/internal/dto"
	"github.com/okarpova/my-app/pkg/otel/tracer"
)

func (u *UseCase) GetPosts(ctx context.Context, input dto.GetPostsInput) (dto.GetPostsOutput, error) {
	ctx, span := tracer.Start(ctx, "usecase GetPosts")
	defer span.End()

	var output dto.GetPostsOutput

	err := input.Validate()
	if err != nil {
		return output, fmt.Errorf("input.Validate: %w", err)
	}

	if input.Limit == 0 {
		input.Limit = 10
	}

	if input.Order == "" {
		input.Order = "desc"
	}

	// Получаем доменные модели
	domainPosts, total, err := u.postgres.GetPosts(ctx, input)
	if err != nil {
		return output, fmt.Errorf("u.postgres.GetPosts: %w", err)
	}

	if len(domainPosts) == 0 {
		return output, domain.ErrNotFound
	}

	// Конвертируем доменные модели в DTO
	output.Posts = make([]dto.Post, 0, len(domainPosts))
	for _, p := range domainPosts {
		output.Posts = append(output.Posts, toPostDTO(p))
	}

	output.Total = total

	return output, nil
}

// toPostDTO конвертирует доменную модель в DTO
func toPostDTO(p domain.Post) dto.Post {
	return dto.Post{
		ID:        p.ID,
		AuthorID:  p.AuthorID,
		Content:   p.Content,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
		// Дополнительные поля при необходимости
	}
}
