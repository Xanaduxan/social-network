package v1

import (
	"net/http"

	"github.com/okarpova/my-app/internal/dto/baggage"

	"github.com/go-chi/chi/v5"
	"github.com/okarpova/my-app/internal/dto"
	"github.com/okarpova/my-app/pkg/render"
)

func (h *Handlers) DeleteProfile(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	input := dto.DeleteProfileInput{
		ID: chi.URLParam(r, "id"),
	}

	baggage.PutProfileID(ctx, input.ID)

	err := h.usecase.DeleteProfile(ctx, input)
	if err != nil {
		render.Error(ctx, w, err, http.StatusBadRequest, "request failed")

		return
	}

	w.WriteHeader(http.StatusNoContent)
}
