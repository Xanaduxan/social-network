//go:build integration

package test

import (
	"context"
	"testing"
	"time"

	"github.com/okarpova/my-app/pkg/otel"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/okarpova/my-app/config"
	kafka_producer "github.com/okarpova/my-app/internal/adapter/kafka"
	"github.com/okarpova/my-app/internal/app"
	"github.com/okarpova/my-app/internal/controller/grpc"
	"github.com/okarpova/my-app/internal/controller/kafka_consumer"
	"github.com/okarpova/my-app/internal/controller/worker"
	"github.com/okarpova/my-app/pkg/httpserver"
	"github.com/okarpova/my-app/pkg/logger"
	"github.com/okarpova/my-app/pkg/postgres"
	"github.com/okarpova/my-app/pkg/redis"
	"github.com/segmentio/kafka-go"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

// Prepare:  make up
// Run test: make integration-test

var ctx = context.Background()

func Test_Integration(t *testing.T) {
	suite.Run(t, &Suite{})
}

type Suite struct {
	suite.Suite
	*require.Assertions

	profile     *ProfileClient
	kafkaWriter *kafka.Writer
	kafkaReader *kafka.Reader
}

func (s *Suite) SetupSuite() {
	s.Assertions = s.Require()

	s.ResetMigrations()

	// Config
	c := config.Config{
		App: config.App{
			Name:    "my-app",
			Version: "test",
		},
		HTTP: httpserver.Config{
			Port: "8080",
		},
		GRPC: grpc.Config{
			Port: "50051",
		},
		Logger: logger.Config{
			AppName:       "my-app",
			AppVersion:    "test",
			Level:         "debug",
			PrettyConsole: true,
		},
		Postgres: postgres.Config{
			Host:     "localhost",
			Port:     "5432",
			User:     "login",
			Password: "pass",
			DBName:   "postgres",
		},
		Redis: redis.Config{
			Addr: "localhost:6379",
		},
		ProduceWorker: worker.ProduceConfig{
			Timeout:      time.Second,
			MessageCount: 1,
		},
		KafkaProducer: kafka_producer.Config{
			Addr:  []string{"localhost:9094"},
			Topic: "",
		},
		KafkaConsumer: kafka_consumer.Config{
			Addr:     []string{"localhost:9094"},
			Topic:    "okarpova-my-app-topic",
			Group:    "okarpova-my-app-group",
			Disabled: true, // Disable consumer in test!
		},
		OutboxKafkaWorker: worker.OutboxKafkaConfig{
			Limit: 10,
		},
	}

	// logger.Init(c.Logger)
	log.Logger = zerolog.Nop()
	otel.SilentModeInit()

	// Kafka writer for direct produce messages
	s.kafkaWriter = &kafka.Writer{
		Addr:  kafka.TCP(c.KafkaProducer.Addr...),
		Topic: c.KafkaProducer.Topic,
	}

	// Kafka reader for direct consume messages
	s.kafkaReader = kafka.NewReader(kafka.ReaderConfig{
		Brokers: c.KafkaConsumer.Addr,
		Topic:   c.KafkaConsumer.Topic,
		GroupID: c.KafkaConsumer.Group,
	})

	// Server
	go func() {
		err := app.Run(context.Background(), c)
		s.NoError(err)
	}()

	BuildProfile(s)

	time.Sleep(time.Second)
}

func (s *Suite) TearDownSuite() {}

func (s *Suite) SetupTest() {}

func (s *Suite) TearDownTest() {}
