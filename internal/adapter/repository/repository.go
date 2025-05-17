package repository

import (
	"time"

	"github.com/okarpova/my-app/internal/adapter/postgres"
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

func New(client *redis.Client, postgres *postgres.Postgres) *Repository {
	return &Repository{
		redis:    client,
		postgres: postgres,
	}
}
