package webserver

import (
	"github.com/pauloRohling/txplorer/internal/presentation/rest/json"
	"net/http"
)

type RestHandler[T any] func(http.ResponseWriter, *http.Request) (T, error)

func Endpoint[T any](method RestHandler[T], status int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		result, err := method(w, r)
		if err != nil {
			json.WriteError(w, err)
			return
		}

		json.WriteJSON(w, status, result)
	}
}
