package usecase

import (
	"context"
	"fmt"

	"gitlab.golang-school.ru/potok-1/okarpova/my-app/pkg/otel/tracer"

	"github.com/google/uuid"
	"gitlab.golang-school.ru/potok-1/okarpova/my-app/internal/domain"
	"gitlab.golang-school.ru/potok-1/okarpova/my-app/internal/dto"
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
