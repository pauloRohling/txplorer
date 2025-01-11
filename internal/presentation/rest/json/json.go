package json

import (
	"encoding/json"
	"errors"
	"github.com/pauloRohling/throw"
	"io"
	"log/slog"
	"net/http"
)

func Parse[T any](r *http.Request) (*T, error) {
	defer func(Body io.ReadCloser) { _ = Body.Close() }(r.Body)

	var input T
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		return nil, throw.Validation().Err(err).Msg("Failed to parse request body")
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
	response := ErrorHandler(err)
	WriteJSON(w, response.Status, response)
}

func ErrorHandler(err error) *HttpErrorResponse {
	if err == nil {
		err = throw.Internal().Msg("Unexpected error")
	}

	var customError *throw.Error
	if !errors.As(err, &customError) {
		customError = throw.Internal().Err(err).Msg("Unexpected error")
	}

	statusCode := throw.ErrorType(customError.Type()).StatusCode()
	errResponse := &HttpErrorResponse{
		Err:    customError.Unwrap(),
		Title:  http.StatusText(statusCode),
		Detail: customError.Error(),
		Status: statusCode,
	}

	slog.Error(
		errResponse.Title,
		"status", errResponse.Status,
		"detail", errResponse.Detail,
		"error", errResponse.Err,
	)

	return errResponse
}
