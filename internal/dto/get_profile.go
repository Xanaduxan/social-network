package dto

import (
	"gitlab.golang-school.ru/potok-1/okarpova/my-app/internal/domain"
)

type GetProfileOutput struct {
	domain.Profile
}

type GetProfileInput struct {
	ID string
}
