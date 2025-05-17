package redis

import (
	"time"

	"github.com/okarpova/my-app/pkg/redis"
)

const (
	idempotencyPrefix = "okarpova:my-app:idempotency:"
	ttl               = time.Hour
)

type Redis struct {
	redis *redis.Client
}

func New(client *redis.Client) *Redis {
	return &Redis{
		redis: client,
	}
}
