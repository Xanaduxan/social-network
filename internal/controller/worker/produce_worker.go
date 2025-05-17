package worker

import (
	"context"
	"time"

	"github.com/okarpova/my-app/pkg/otel/tracer"
	"go.opentelemetry.io/otel/trace"

	"github.com/okarpova/my-app/internal/usecase"
	"github.com/rs/zerolog/log"
)

type ProduceConfig struct {
	Timeout      time.Duration `envconfig:"PRODUCE_WORKER_TIMEOUT" default:"10s"`
	MessageCount int           `envconfig:"PRODUCE_WORKER_MESSAGE_COUNT" default:"1"`
	Disabled     bool          `envconfig:"PRODUCE_WORKER_DISABLED"`
}

type ProduceWorker struct {
	config  ProduceConfig
	usecase *usecase.UseCase
	stop    chan struct{}
	done    chan struct{}
}

func NewProduceWorker(c ProduceConfig, uc *usecase.UseCase) *ProduceWorker {
	w := &ProduceWorker{
		config:  c,
		usecase: uc,
		stop:    make(chan struct{}),
		done:    make(chan struct{}),
	}

	if w.config.Disabled {
		log.Info().Msg("produce worker: disabled")

		return w
	}

	go w.run()

	return w
}

func (w *ProduceWorker) run() {
	log.Info().Msg("produce worker: started")

FOR:
	for {
		ctx := context.Background()
		ctx, span := tracer.Start(ctx, "worker produce", trace.WithSpanKind(trace.SpanKindInternal))

		err := w.usecase.GenerateMessages(ctx, w.config.MessageCount)
		if err != nil {
			log.Error().Err(err).Msg("produce worker: GenerateMessages failed")
		}

		span.End() // Закрываем span

		select {
		case <-w.stop:
			break FOR
		case <-time.After(w.config.Timeout):
		}
	}

	log.Info().Msg("produce worker: stopped")

	close(w.done)
}

func (w *ProduceWorker) Stop() {
	if w.config.Disabled {
		return
	}

	close(w.stop)

	<-w.done
}
