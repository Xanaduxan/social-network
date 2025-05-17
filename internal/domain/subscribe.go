package domain

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type SubscriptionType string

const (
	Regular SubscriptionType = "regular"
	Premium SubscriptionType = "premium"
)

type Subscription struct {
	ID               uuid.UUID        `json:"id"`
	CreatedAt        time.Time        `json:"created_at"`
	UpdatedAt        time.Time        `json:"updated_at"`
	DeletedAt        time.Time        `json:"deleted_at"`
	SubscriberID     uuid.UUID        `json:"subscriber_id" validate:"required"`
	TargetID         uuid.UUID        `json:"target_id" validate:"required"`
	SubscriptionType SubscriptionType `json:"subscription_type" validate:"oneof=regular premium"`
}

type SubscriptionNotification struct {
	ID             uuid.UUID `json:"id"`
	SubscriptionID uuid.UUID `json:"subscription_id" validate:"required"`
	Type           string    `json:"type" validate:"required"`
	TargetID       uuid.UUID `json:"target_id" validate:"required"`
	CreatedAt      time.Time `json:"created_at"`
}

func NewSubscription(subscriberID, targetID uuid.UUID, subType SubscriptionType) (Subscription, error) {
	sub := Subscription{
		ID:               uuid.New(),
		CreatedAt:        time.Now().UTC(),
		UpdatedAt:        time.Now().UTC(),
		SubscriberID:     subscriberID,
		TargetID:         targetID,
		SubscriptionType: subType,
	}

	if err := sub.Validate(); err != nil {
		return Subscription{}, fmt.Errorf("sub.Validate: %w", err)
	}

	return sub, nil
}

func NewSubscriptionNotification(subscriptionID, targetID uuid.UUID) SubscriptionNotification {
	return SubscriptionNotification{
		ID:             uuid.New(),
		SubscriptionID: subscriptionID,
		Type:           "new_subscription",
		TargetID:       targetID,
		CreatedAt:      time.Now().UTC(),
	}
}

func (s Subscription) Validate() error {
	if s.SubscriberID == s.TargetID {
		return fmt.Errorf("subscriber cannot subscribe to themselves")
	}

	err := validate.Struct(s)
	if err != nil {
		return fmt.Errorf("validate.Struct: %w", err)
	}

	return nil
}

func (s Subscription) IsDeleted() bool {
	return !s.DeletedAt.IsZero()
}

func (s Subscription) IsPremium() bool {
	return s.SubscriptionType == Premium
}
