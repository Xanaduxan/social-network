package main

import (
	"context"

	"github.com/okarpova/my-app/pkg/otel"

	"github.com/okarpova/my-app/config"
	"github.com/okarpova/my-app/internal/app"
	"github.com/okarpova/my-app/pkg/logger"
	"github.com/rs/zerolog/log"
	_ "go.uber.org/automaxprocs"
)

func main() {
	c, err := config.New()
	if err != nil {
		log.Fatal().Err(err).Msg("config.New")
	}

	logger.Init(c.Logger)

	ctx := context.Background()

	err = otel.Init(ctx, c.OTEL)
	if err != nil {
		log.Fatal().Err(err).Msg("otel.Init")
	}
	defer otel.Close()

	err = app.Run(ctx, c)
	if err != nil {
		log.Fatal().Err(err).Msg("app.Run")
	}
}
