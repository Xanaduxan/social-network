package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/okarpova/my-app/pkg/otel/tracer"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/okarpova/my-app/internal/domain"
	"github.com/okarpova/my-app/pkg/transaction"
)

func (p *Postgres) DeleteProfile(ctx context.Context, id uuid.UUID) error {
	ctx, span := tracer.Start(ctx, "adapter postgres DeleteProfile")
	defer span.End()

	const sql = `UPDATE profile SET deleted_at = NOW()
                     WHERE id = $1`

	txOrPool := transaction.TryExtractTX(ctx)

	_, err := txOrPool.Exec(ctx, sql, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.ErrNotFound
		}

		return fmt.Errorf("txOrPool.Exec: %w", err)
	}

	return nil
}
