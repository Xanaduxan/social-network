package dto

import "time"

type CreateSubscribeInput struct {
	SubscriberID     string `json:"subscriber_id" validate:"required"`
	TargetID         string `json:"target_id" validate:"required"`
	SubscriptionType string `json:"subscription_type" validate:"oneof=regular premium"`
}

type CreateSubscribeOutput struct {
	SubscriptionID string    `json:"subscription_id"`
	CreatedAt      time.Time `json:"created_at"`
}
