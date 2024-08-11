package webserver

import "github.com/go-chi/chi/v5"

type Routable interface {
	Endpoint() string
	Route(r chi.Router)
}
