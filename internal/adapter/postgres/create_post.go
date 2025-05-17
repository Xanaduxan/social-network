package postgres

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/okarpova/my-app/internal/domain"
	"github.com/okarpova/my-app/pkg/otel/tracer"
	"github.com/okarpova/my-app/pkg/transaction"
)

func (p *Postgres) CreatePost(ctx context.Context, post domain.Post) error {
	ctx, span := tracer.Start(ctx, "adapter postgres CreatePost")
	defer span.End()

	const sql = `INSERT INTO posts (id, author_id, content, visibility, created_at, updated_at, attachments)
                 VALUES ($1, $2, $3, $4, $5, $6, $7)`

	attachments, err := json.Marshal(post.Attachments)
	if err != nil {
		return fmt.Errorf("json.Marshal attachments: %w", err)
	}

	args := []any{
		post.ID,
		post.AuthorID,
		post.Content,
		post.Visibility,
		post.CreatedAt,
		post.UpdatedAt,
		attachments,
	}

	txOrPool := transaction.TryExtractTX(ctx)

	_, err = txOrPool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("txOrPool.Exec: %w", err)
	}

	return nil
}

func (p *Postgres) CreatePostStats(ctx context.Context, stats domain.PostStats) error {
	ctx, span := tracer.Start(ctx, "adapter postgres CreatePostStats")
	defer span.End()

	const sql = `INSERT INTO post_stats (post_id, likes, views, shares, created_at, updated_at)
                 VALUES ($1, $2, $3, $4, $5, $6)`

	args := []any{
		stats.PostID,
		stats.Likes,
		stats.Views,
		stats.Shares,
		stats.CreatedAt,
		stats.UpdatedAt,
	}

	txOrPool := transaction.TryExtractTX(ctx)

	_, err := txOrPool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("txOrPool.Exec: %w", err)
	}

	return nil
}
