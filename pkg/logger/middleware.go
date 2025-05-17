package logger

import (
	"fmt"
	"net/http"

	"gitlab.golang-school.ru/potok-1/okarpova/my-app/pkg/router"

	"github.com/rs/zerolog/log"
	"gitlab.golang-school.ru/potok-1/okarpova/my-app/internal/dto/baggage"
)

func Middleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		bag := &baggage.Baggage{}
		ctx := baggage.WithContext(r.Context(), bag)

		ww := router.WriterWrapper(w)
		next.ServeHTTP(ww, r.WithContext(ctx))

		event := log.Info()

		if bag.Err != nil {
			event = log.Error().Err(bag.Err)
		}

		event.
			Str("profile_id", bag.ProfileID).
			Str("proto", "http").
			Int("code", ww.Code()).
			Str("method", fmt.Sprintf("%s %s", r.Method, router.ExtractPath(ctx))).
			Send()
	}

	return http.HandlerFunc(fn)
}
