package postgres

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/okarpova/my-app/internal/domain"
	"github.com/okarpova/my-app/internal/dto"
	"github.com/okarpova/my-app/pkg/otel/tracer"
	"github.com/okarpova/my-app/pkg/transaction"
)

type GetPostsDTO struct {
	ID          pgtype.UUID
	AuthorID    pgtype.UUID
	Content     pgtype.Text
	Visibility  pgtype.Text
	CreatedAt   pgtype.Timestamptz
	UpdatedAt   pgtype.Timestamptz
	DeletedAt   pgtype.Timestamptz
	Attachments []byte
}

func (d *GetPostsDTO) ToDomain() (domain.Post, error) {
	var attachments []string
	if len(d.Attachments) > 0 {
		err := json.Unmarshal(d.Attachments, &attachments)
		if err != nil {
			return domain.Post{}, fmt.Errorf("cannot unmarshal attachments: %w", err)
		}
	}

	return domain.Post{
		ID:          d.ID.Bytes,
		AuthorID:    d.AuthorID.Bytes,
		Content:     d.Content.String,
		Visibility:  d.Visibility.String,
		CreatedAt:   d.CreatedAt.Time,
		UpdatedAt:   d.UpdatedAt.Time,
		Attachments: attachments,
	}, nil
}

func (d *GetPostsDTO) Dest() []any {
	return []any{
		&d.ID,
		&d.AuthorID,
		&d.Content,
		&d.Visibility,
		&d.CreatedAt,
		&d.UpdatedAt,
		&d.DeletedAt,
		&d.Attachments,
	}
}

func (p *Postgres) GetPosts(ctx context.Context, input dto.GetPostsInput) ([]domain.Post, error) {
	ctx, span := tracer.Start(ctx, "adapter postgres GetPosts")
	defer span.End()

	sql := `SELECT id, author_id, content, visibility, created_at, updated_at, deleted_at, attachments
           FROM posts
           WHERE deleted_at IS NULL
           ORDER BY %s %s
           OFFSET %d
           LIMIT %d`

	sql = fmt.Sprintf(sql, input.Sort, input.Order, input.Offset, input.Limit)

	txOrPool := transaction.TryExtractTX(ctx)

	rows, err := txOrPool.Query(ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("txOrPool.Query: %w", err)
	}

	defer rows.Close()

	posts := make([]domain.Post, 0, input.Limit)

	for rows.Next() {
		var d GetPostsDTO

		err = rows.Scan(d.Dest()...)
		if err != nil {
			return nil, fmt.Errorf("rows.Scan: %w", err)
		}

		post, err := d.ToDomain()
		if err != nil {
			return nil, fmt.Errorf("dto.ToDomain: %w", err)
		}

		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows.Err: %w", err)
	}

	return posts, nil
}
