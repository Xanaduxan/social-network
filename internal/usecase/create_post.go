package usecase

import (
	"context"
	"fmt"

	"github.com/okarpova/my-app/internal/domain"
	"github.com/okarpova/my-app/internal/dto"
	"github.com/okarpova/my-app/pkg/otel/tracer"
	"github.com/okarpova/my-app/pkg/transaction"
)

func (u *UseCase) CreatePost(ctx context.Context, input dto.CreatePostInput) (dto.CreatePostOutput, error) {
	ctx, span := tracer.Start(ctx, "usecase CreatePost")
	defer span.End()

	var output dto.CreatePostOutput

	post, err := domain.NewPost(
		input.AuthorID,
		input.Content,
		input.Visibility,
		input.Attachments,
	)
	if err != nil {
		return output, fmt.Errorf("domain.NewPost: %w", err)
	}

	ctx, err = transaction.Begin(ctx)
	if err != nil {
		return output, fmt.Errorf("transaction.Begin: %w", err)
	}
	defer transaction.Rollback(ctx)

	err = u.postgres.CreatePost(ctx, *post)
	if err != nil {
		return output, fmt.Errorf("u.postgres.CreatePost: %w", err)
	}

	err = transaction.Commit(ctx)
	if err != nil {
		return output, fmt.Errorf("transaction.Commit: %w", err)
	}

	return dto.CreatePostOutput{
		ID:        post.ID,
		CreatedAt: post.CreatedAt,
	}, nil
}
