package config

import (
	"fmt"

	"gitlab.golang-school.ru/potok-1/okarpova/my-app/pkg/otel"

	"gitlab.golang-school.ru/potok-1/okarpova/my-app/internal/adapter/kafka"
	"gitlab.golang-school.ru/potok-1/okarpova/my-app/internal/controller/kafka_consumer"

	"gitlab.golang-school.ru/potok-1/okarpova/my-app/internal/controller/worker"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"gitlab.golang-school.ru/potok-1/okarpova/my-app/internal/controller/grpc"
	"gitlab.golang-school.ru/potok-1/okarpova/my-app/pkg/httpclient"
	"gitlab.golang-school.ru/potok-1/okarpova/my-app/pkg/httpserver"
	"gitlab.golang-school.ru/potok-1/okarpova/my-app/pkg/logger"
	"gitlab.golang-school.ru/potok-1/okarpova/my-app/pkg/postgres"
	"gitlab.golang-school.ru/potok-1/okarpova/my-app/pkg/redis"
)

type App struct {
	Name    string `envconfig:"APP_NAME"    required:"true"`
	Version string `envconfig:"APP_VERSION" required:"true"`
}

type Config struct {
	App               App
	HTTP              httpserver.Config
	GRPC              grpc.Config
	Logger            logger.Config
	OTEL              otel.Config
	Postgres          postgres.Config
	Redis             redis.Config
	Client            httpclient.Config
	KafkaConsumer     kafka_consumer.Config
	KafkaProducer     kafka.Config
	ProduceWorker     worker.ProduceConfig
	OutboxKafkaWorker worker.OutboxKafkaConfig
}

func New() (Config, error) {
	var config Config

	err := godotenv.Load(".env")
	if err != nil {
		return config, fmt.Errorf("godotenv.Load: %w", err)
	}

	err = envconfig.Process("", &config)
	if err != nil {
		return config, fmt.Errorf("envconfig.Process: %w", err)
	}

	return config, nil
}
