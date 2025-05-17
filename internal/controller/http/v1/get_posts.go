package v1

import (
	"encoding/json"
	"net/http"

	"github.com/okarpova/my-app/internal/dto"
	"github.com/okarpova/my-app/pkg/render"
)

func (h *Handlers) GetPosts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Прямое использование строковых параметров без парсинга UUID
	input := dto.GetPostsInput{
		AuthorID: r.URL.Query().Get("author_id"), // Оставляем как строку
		Order:    r.URL.Query().Get("order"),
		Offset:   atoi(r.URL.Query().Get("offset")),
		Limit:    atoi(r.URL.Query().Get("limit")),
	}

	output, err := h.usecase.GetPosts(ctx, input)
	if err != nil {
		render.Error(ctx, w, err, http.StatusBadRequest, "failed to get posts")
		return
	}

	render.JSON(w, output, http.StatusOK)
}

func atoi(s string) int {
	if s == "" {
		return 0
	}
	n, _ := strconv.Atoi(s)
	return n
}
