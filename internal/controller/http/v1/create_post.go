package v1

import (
	"encoding/json"
	"net/http"

	"github.com/okarpova/my-app/internal/dto"
	"github.com/okarpova/my-app/pkg/render"
)

func (h *Handlers) CreatePost(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	input := dto.CreatePostInput{}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		render.Error(ctx, w, err, http.StatusBadRequest, "json decode error")
		return
	}

	output, err := h.usecase.CreatePost(ctx, input)
	if err != nil {
		render.Error(ctx, w, err, http.StatusInternalServerError, "failed to create post")
		return
	}

	render.JSON(w, output, http.StatusCreated)
}
