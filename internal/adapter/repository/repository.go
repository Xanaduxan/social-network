package repository

import (
	"gitlab.golang-school.ru/potok-1/okarpova/my-app/internal/adapter/postgres"
	"gitlab.golang-school.ru/potok-1/okarpova/my-app/pkg/redis"
	"time"
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
