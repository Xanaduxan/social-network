package v1

import (
	"encoding/json"
	"net/http"

	"gitlab.golang-school.ru/potok-1/okarpova/my-app/internal/dto/baggage"

	"gitlab.golang-school.ru/potok-1/okarpova/my-app/internal/dto"
	"gitlab.golang-school.ru/potok-1/okarpova/my-app/pkg/render"
)

// UpdateProfile обновляет существующий или создаёт новый профиль
func (h *Handlers) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	input := dto.UpdateProfileInput{}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		render.Error(ctx, w, err, http.StatusBadRequest, "json decode error")

		return
	}

	baggage.PutProfileID(ctx, input.ID)

	err = h.usecase.UpdateProfile(ctx, input)
	if err != nil {
		render.Error(ctx, w, err, http.StatusBadRequest, "request failed")

		return
	}

	w.WriteHeader(http.StatusNoContent)
}
