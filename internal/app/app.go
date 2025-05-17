package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/okarpova/my-app/config"
	"github.com/okarpova/my-app/internal/adapter/kafka"
	"github.com/okarpova/my-app/internal/adapter/postgres"
	"github.com/okarpova/my-app/internal/adapter/redis"
	"github.com/okarpova/my-app/internal/adapter/repository"
	"github.com/okarpova/my-app/internal/controller/grpc"
	"github.com/okarpova/my-app/internal/controller/http"
	"github.com/okarpova/my-app/internal/controller/kafka_consumer"
	"github.com/okarpova/my-app/internal/controller/worker"
	"github.com/okarpova/my-app/internal/usecase"
	"github.com/okarpova/my-app/pkg/httpclient"
	"github.com/okarpova/my-app/pkg/httpserver"
	"github.com/okarpova/my-app/pkg/metrics"
	pgpool "github.com/okarpova/my-app/pkg/postgres"
	redislib "github.com/okarpova/my-app/pkg/redis"
	"github.com/okarpova/my-app/pkg/router"
	"github.com/okarpova/my-app/pkg/transaction"
	"github.com/rs/zerolog/log"
)

func Run(ctx context.Context, c config.Config) (err error) {
	// Postgres
	pgPool, err := pgpool.New(ctx, c.Postgres)
	if err != nil {
		return fmt.Errorf("postgres.New: %w", err)
	}

	transaction.Init(pgPool)

	// Redis
	redisClient, err := redislib.New(c.Redis)
	if err != nil {
		return fmt.Errorf("redislib.New: %w", err)
	}

	entityMetrics := metrics.NewEntity()
	httpMetrics := metrics.NewHTTPServer()

	// Kafka producer
	kafkaProducer := kafka.NewProducer(c.KafkaProducer, entityMetrics)

	// UseCase
	uc := usecase.New(
		repository.New(redisClient, postgres.New()),
		httpclient.New(c.Client),
		kafkaProducer,
		redis.New(redisClient),
	)

	// Kafka consumer
	kafkaConsumer := kafka_consumer.New(c.KafkaConsumer, entityMetrics, uc)

	// Produce worker
	produceWorker := worker.NewProduceWorker(c.ProduceWorker, uc)

	// Outbox Kafka worker
	outboxKafkaWorker := worker.NewOutboxKafkaWorker(uc, c.OutboxKafkaWorker)

	// Worker
	someWorker, err := worker.NewSomeWorker(uc)
	if err != nil {
		return fmt.Errorf("worker.NewSomeWorker: %w", err)
	}

	// GRPC
	grpcServer, err := grpc.New(c.GRPC, uc)
	if err != nil {
		return fmt.Errorf("grpc.New: %w", err)
	}

	// HTTP
	r := router.New()
	http.ProfileRouter(r, uc, httpMetrics)
	httpServer := httpserver.New(r, c.HTTP)

	log.Info().Msg("App started!")

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	<-sig // wait signal

	log.Info().Msg("App got signal to stop")

	// Contollers
	outboxKafkaWorker.Stop()
	kafkaConsumer.Close()
	someWorker.Stop()
	grpcServer.Close()
	httpServer.Close()
	produceWorker.Stop()

	// Adapters
	redisClient.Close()
	kafkaProducer.Close()
	pgPool.Close()

	log.Info().Msg("App stopped!")

	return nil
}
