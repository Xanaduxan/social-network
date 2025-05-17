package main

import (
	"context"

	"gitlab.golang-school.ru/potok-1/okarpova/my-app/pkg/otel"

	"github.com/rs/zerolog/log"
	"gitlab.golang-school.ru/potok-1/okarpova/my-app/config"
	"gitlab.golang-school.ru/potok-1/okarpova/my-app/internal/app"
	"gitlab.golang-school.ru/potok-1/okarpova/my-app/pkg/logger"
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
