package repository

import (
	"context"
	"time"

	"github.com/okarpova/my-app/internal/adapter/postgres"
	"github.com/okarpova/my-app/internal/domain"
	"github.com/okarpova/my-app/pkg/redis"
)

const (
	prefix = "okarpova:my-app:"
	ttl    = time.Minute
)

type Repository struct {
	redis    *redis.Client
	postgres *postgres.Postgres
}

// CreatePost implements usecase.Postgres.
func (r *Repository) CreatePost(ctx context.Context, post domain.Post) error {
	panic("unimplemented")
}

func New(client *redis.Client, postgres *postgres.Postgres) *Repository {
	return &Repository{
		redis:    client,
		postgres: postgres,
	}
}
