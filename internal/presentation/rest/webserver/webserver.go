package webserver

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/pauloRohling/txplorer/internal/presentation/rest/json"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
	"time"
)

type WebServer struct {
	port       int32
	router     chi.Router
	httpServer *http.Server
}

func NewWebServer(port int32, corsOptions *cors.Options) *WebServer {
	router := chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(middleware.SetHeader("Content-Type", "application/json"))
	router.Use(middleware.AllowContentType("application/json"))
	router.Use(middleware.Timeout(30 * time.Second))

	if corsOptions != nil {
		router.Use(cors.Handler(*corsOptions))
	}

	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		json.WriteError(w, nil)
	})

	return &WebServer{
		router: router,
		port:   port,
		httpServer: &http.Server{
			Addr:    fmt.Sprintf(":%d", port),
			Handler: router,
		},
	}
}

func (server *WebServer) AddRoute(routable Routable) {
	server.router.Route(routable.Endpoint(), routable.Route)
}

func (server *WebServer) Start() error {
	err := server.httpServer.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

func (server *WebServer) AddSwaggerRoute() {
	url := fmt.Sprintf("http://localhost:%d/swagger/doc.json", server.port)
	handler := httpSwagger.Handler(httpSwagger.URL(url))
	server.router.Get("/swagger/*", handler)
}

func (server *WebServer) Shutdown(ctx context.Context) error {
	return server.httpServer.Shutdown(ctx)
}
