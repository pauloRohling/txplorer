package json

import (
	"encoding/json"
	"errors"
	"github.com/pauloRohling/txplorer/internal/model"
	"io"
	"log/slog"
	"net/http"
)

func Parse[T any](r *http.Request) (*T, error) {
	defer func(Body io.ReadCloser) { _ = Body.Close() }(r.Body)

	var input T
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		return nil, err
	}

	return &input, nil
}

func WriteJSON(w http.ResponseWriter, status int, payload any) {
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(payload)

	if err != nil {
		_ = json.NewEncoder(w).Encode(err)
	}
}

func WriteError(w http.ResponseWriter, err error) {
	if err == nil {
		err = model.InternalError("Empty error")
	}

	var customErr model.Error
	if !errors.As(err, &customErr) {
		customErr = model.InternalError(err.Error())
	}

	if customErr.Err == nil {
		slog.Error(customErr.Error())
	} else {
		slog.Error(customErr.Error(), "description", customErr.Err.Error())
	}

	slog.Error(customErr.StackTrace)

	response := NewResponseFromError(customErr)
	WriteJSON(w, response.Status, response)
}
