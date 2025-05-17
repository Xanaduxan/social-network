package logger

import (
	"context"

	"google.golang.org/grpc/status"

	"github.com/rs/zerolog/log"
	"gitlab.golang-school.ru/potok-1/okarpova/my-app/internal/dto/baggage"
	"google.golang.org/grpc"
)

func Interceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	bag := &baggage.Baggage{}
	ctx = baggage.WithContext(ctx, bag)

	event := log.Info()

	resp, err := handler(ctx, req)
	if err != nil {
		event = log.Error().Err(bag.Err)
	}

	event.
		Str("profile_id", bag.ProfileID).
		Str("proto", "grpc").
		Str("code", status.Code(err).String()).
		Str("method", info.FullMethod).
		Send()

	return resp, err
}
