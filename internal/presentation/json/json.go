package json

import (
	"encoding/json"
	"io"
	"net/http"
)

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func Endpoint[T any](method func(http.ResponseWriter, *http.Request) (T, error), status int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		result, err := method(w, r)
		if err != nil {
			WriteError(w, err)
			return
		}

		WriteJSON(w, status, result)
	}
}

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
	payload := Error{Code: http.StatusInternalServerError, Message: err.Error()}
	WriteJSON(w, http.StatusInternalServerError, payload)
}
