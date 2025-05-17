package usecase

import (
	"context"
	"fmt"

	"github.com/okarpova/my-app/pkg/otel/tracer"

	"github.com/google/uuid"
	"github.com/okarpova/my-app/internal/domain"
	"github.com/okarpova/my-app/internal/dto"
)

func (u *UseCase) GetProfile(ctx context.Context, input dto.GetProfileInput) (dto.GetProfileOutput, error) {
	ctx, span := tracer.Start(ctx, "usecase GetProfile")
	defer span.End()

	var output dto.GetProfileOutput

	id, err := uuid.Parse(input.ID)
	if err != nil {
		return output, domain.ErrUUIDInvalid
	}

	profile, err := u.postgres.GetProfile(ctx, id)
	if err != nil {
		return output, fmt.Errorf("u.postgres.GetProfile: %w", err)
	}

	if profile.IsDeleted() {
		return output, domain.ErrNotFound
	}

	return dto.GetProfileOutput{
		Profile: profile,
	}, nil
}
