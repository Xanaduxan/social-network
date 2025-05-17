package usecase

import (
	"context"
	"fmt"

	"github.com/okarpova/my-app/pkg/otel"

	"github.com/segmentio/kafka-go"

	"github.com/okarpova/my-app/pkg/otel/tracer"

	"github.com/okarpova/my-app/pkg/transaction"

	"github.com/google/uuid"
)

func (u *UseCase) GenerateMessages(ctx context.Context, msgCount int) error {
	ctx, span := tracer.Start(ctx, "usecase GenerateMessages")
	defer span.End()

	msgs := make([]kafka.Message, 0, msgCount)

	for range msgCount {
		msgs = append(msgs, kafka.Message{
			Topic: "okarpova-my-app-topic",
			Key:   []byte(uuid.New().String()),
			Value: []byte(uuid.New().String()),
		})
	}

	// Трекаем каждое сообщение
	otel.InjectPropagateHeaders(ctx, msgs...)

	// В транзакции
	err := transaction.Wrap(ctx, func(ctx context.Context) error {
		// Пишем в outbox таблицу БД
		err := u.postgres.SaveOutboxKafka(ctx, msgs...)
		if err != nil {
			return fmt.Errorf("u.postgres.SaveOutboxKafka: %w", err)
		}

		// И что-то ещё записываем в БД

		return nil
	})
	if err != nil {
		return fmt.Errorf("transaction.Wrap: %w", err)
	}

	return nil
}
