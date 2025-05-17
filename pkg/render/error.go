package render

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/okarpova/my-app/internal/dto/baggage"
)

type Err struct {
	Error string `json:"error"`
}

func Error(ctx context.Context, w http.ResponseWriter, err error, status int, message string) {
	baggage.PutError(ctx, err)

	err = unpack(err)
	err = fmt.Errorf("%s: %s", message, err)

	JSON(w, Err{Error: err.Error()}, status)
}

func unpack(err error) error {
	for {
		e := errors.Unwrap(err)
		if e == nil {
			break
		}

		err = e
	}

	return err
}
