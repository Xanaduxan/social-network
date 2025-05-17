package usecase

import (
	"context"
	"fmt"

	"gitlab.golang-school.ru/potok-1/okarpova/my-app/pkg/otel/tracer"

	"github.com/google/uuid"
	"gitlab.golang-school.ru/potok-1/okarpova/my-app/internal/domain"
	"gitlab.golang-school.ru/potok-1/okarpova/my-app/internal/dto"
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
