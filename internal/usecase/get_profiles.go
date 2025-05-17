package usecase

import (
	"context"
	"fmt"

	"github.com/okarpova/my-app/pkg/otel/tracer"

	"github.com/okarpova/my-app/internal/domain"
	"github.com/okarpova/my-app/internal/dto"
)

func (u *UseCase) GetProfiles(ctx context.Context, input dto.GetProfilesInput) (dto.GetProfilesOutput, error) {
	ctx, span := tracer.Start(ctx, "usecase GetProfiles")
	defer span.End()

	var output dto.GetProfilesOutput

	err := input.Validate()
	if err != nil {
		return output, fmt.Errorf("input.Validate: %w", err)
	}

	if input.Limit == 0 {
		input.Limit = 10
	}

	if input.Order == "" {
		input.Order = "asc"
	}

	profiles, err := u.postgres.GetProfiles(ctx, input)
	if err != nil {
		return output, fmt.Errorf("u.postgres.GetProfiles: %w", err)
	}

	if len(profiles) == 0 {
		return output, domain.ErrNotFound
	}

	output.Profiles = profiles

	return output, nil
}
