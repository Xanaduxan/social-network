package usecase

import (
	"context"

	"github.com/segmentio/kafka-go"

	"github.com/google/uuid"

	"gitlab.golang-school.ru/potok-1/okarpova/my-app/internal/domain"
	"gitlab.golang-school.ru/potok-1/okarpova/my-app/internal/dto"
	"gitlab.golang-school.ru/potok-1/okarpova/my-app/pkg/httpclient"
)

type Redis interface {
	IsExists(ctx context.Context, idempotencyKey string) bool
}

type Kafka interface {
	Produce(ctx context.Context, msgs ...kafka.Message) error
}

type Profile interface {
	Create(ctx context.Context, name string, age int, email, phone string) (uuid.UUID, error)
	Delete(ctx context.Context, id string) error
	GetProfile(ctx context.Context, id string) (httpclient.Profile, error)
	Update(ctx context.Context, id string, name *string, age *int, email, phone *string) error
}

type Postgres interface {
	CreateProfile(ctx context.Context, profile domain.Profile) error
	CreateProperty(ctx context.Context, property domain.Property) error
	GetProfile(ctx context.Context, profileID uuid.UUID) (domain.Profile, error)
	GetProfiles(ctx context.Context, input dto.GetProfilesInput) ([]domain.Profile, error)
	UpdateProfile(ctx context.Context, profile domain.Profile) error
	DeleteProfile(ctx context.Context, id uuid.UUID) error

	ReadOutboxKafka(ctx context.Context, limit int) ([]kafka.Message, error)
	SaveOutboxKafka(ctx context.Context, msgs ...kafka.Message) error
}

type UseCase struct {
	profile  Profile
	postgres Postgres
	kafka    Kafka
	redis    Redis
}

func New(postgres Postgres, profile Profile, kafka Kafka, redis Redis) *UseCase {
	return &UseCase{
		postgres: postgres,
		profile:  profile,
		kafka:    kafka,
		redis:    redis,
	}
}
