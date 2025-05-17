package http

import (
	"github.com/okarpova/my-app/pkg/metrics"
	"github.com/okarpova/my-app/pkg/otel"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/go-chi/chi/v5"

	ver1 "github.com/okarpova/my-app/internal/controller/http/v1"

	"github.com/okarpova/my-app/internal/usecase"
	"github.com/okarpova/my-app/pkg/logger"
)

func ProfileRouter(r *chi.Mux, uc *usecase.UseCase, m *metrics.HTTPServer) {
	v1 := ver1.New(uc)

	r.Handle("/metrics", promhttp.Handler())

	r.Route("/okarpova/my-app/api", func(r chi.Router) {
		r.Use(logger.Middleware)
		r.Use(metrics.NewMiddleware(m))
		r.Use(otel.Middleware)

		r.Route("/v1", func(r chi.Router) {
			r.Post("/profile", v1.CreateProfile)
			r.Put("/profile", v1.UpdateProfile)
			r.Get("/profile/{id}", v1.GetProfile)
			r.Get("/profiles", v1.GetProfiles)
			r.Delete("/profile/{id}", v1.DeleteProfile)
		})

	})
}
