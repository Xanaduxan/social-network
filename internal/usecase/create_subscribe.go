package usecase

import (
	"context"
	"fmt"

	"github.com/okarpova/my-app/internal/domain"
	"github.com/okarpova/my-app/internal/dto"
	"github.com/okarpova/my-app/pkg/otel/tracer"
	"github.com/okarpova/my-app/pkg/transaction"
)

func (u *UseCase) CreateSubscribe(ctx context.Context, input dto.CreateSubscribeInput) (dto.CreateSubscribeOutput, error) {
	ctx, span := tracer.Start(ctx, "usecase CreateSubscribe")
	defer span.End()

	var output dto.CreateSubscribeOutput

	// Создаем доменную модель подписки
	subscription, err := domain.NewSubscription(
		input.SubscriberID,
		input.TargetID,
		input.SubscriptionType,
	)
	if err != nil {
		return output, fmt.Errorf("domain.NewSubscription: %w", err)
	}

	// Создаем связанную сущность (например, уведомление о подписке)
	notification := domain.NewSubscriptionNotification(
		subscription.ID,
		"new_subscription",
		input.TargetID,
	)

	// Начинаем транзакцию
	ctx, err = transaction.Begin(ctx)
	if err != nil {
		return output, fmt.Errorf("transaction.Begin: %w", err)
	}
	defer transaction.Rollback(ctx)

	// Сохраняем подписку в БД
	err = u.postgres.CreateSubscription(ctx, subscription)
	if err != nil {
		return output, fmt.Errorf("u.postgres.CreateSubscription: %w", err)
	}

	// Сохраняем уведомление
	err = u.postgres.CreateSubscriptionNotification(ctx, notification)
	if err != nil {
		return output, fmt.Errorf("u.postgres.CreateSubscriptionNotification: %w", err)
	}

	// Фиксируем транзакцию
	err = transaction.Commit(ctx)
	if err != nil {
		return output, fmt.Errorf("transaction.Commit: %w", err)
	}

	return dto.CreateSubscribeOutput{
		SubscriptionID: subscription.ID,
		CreatedAt:      subscription.CreatedAt,
	}, nil
}
