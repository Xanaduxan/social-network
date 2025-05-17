package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog/log"
	"gitlab.golang-school.ru/potok-1/okarpova/my-app/config"
	"gitlab.golang-school.ru/potok-1/okarpova/my-app/internal/adapter/kafka"
	"gitlab.golang-school.ru/potok-1/okarpova/my-app/internal/adapter/postgres"
	"gitlab.golang-school.ru/potok-1/okarpova/my-app/internal/adapter/redis"
	"gitlab.golang-school.ru/potok-1/okarpova/my-app/internal/adapter/repository"
	"gitlab.golang-school.ru/potok-1/okarpova/my-app/internal/controller/grpc"
	"gitlab.golang-school.ru/potok-1/okarpova/my-app/internal/controller/http"
	"gitlab.golang-school.ru/potok-1/okarpova/my-app/internal/controller/kafka_consumer"
	"gitlab.golang-school.ru/potok-1/okarpova/my-app/internal/controller/worker"
	"gitlab.golang-school.ru/potok-1/okarpova/my-app/internal/usecase"
	"gitlab.golang-school.ru/potok-1/okarpova/my-app/pkg/httpclient"
	"gitlab.golang-school.ru/potok-1/okarpova/my-app/pkg/httpserver"
	"gitlab.golang-school.ru/potok-1/okarpova/my-app/pkg/metrics"
	pgpool "gitlab.golang-school.ru/potok-1/okarpova/my-app/pkg/postgres"
	redislib "gitlab.golang-school.ru/potok-1/okarpova/my-app/pkg/redis"
	"gitlab.golang-school.ru/potok-1/okarpova/my-app/pkg/router"
	"gitlab.golang-school.ru/potok-1/okarpova/my-app/pkg/transaction"
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
