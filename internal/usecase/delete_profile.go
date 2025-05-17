package usecase

import (
	"context"
	"fmt"

	"github.com/okarpova/my-app/pkg/otel/tracer"

	"github.com/google/uuid"
	"github.com/okarpova/my-app/internal/domain"
	"github.com/okarpova/my-app/internal/dto"
)

func (u *UseCase) DeleteProfile(ctx context.Context, input dto.DeleteProfileInput) error {
	ctx, span := tracer.Start(ctx, "usecase DeleteProfile")
	defer span.End()

	id, err := uuid.Parse(input.ID)
	if err != nil {
		return domain.ErrUUIDInvalid
	}

	err = u.postgres.DeleteProfile(ctx, id)
	if err != nil {
		return fmt.Errorf("u.postgres.DeleteProfile: %w", err)
	}

	return nil
}
