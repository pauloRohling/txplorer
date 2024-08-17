package rest

import (
	"github.com/go-chi/chi/v5"
	"github.com/pauloRohling/txplorer/internal/domain/user"
	"github.com/pauloRohling/txplorer/internal/presentation/rest/json"
	"github.com/pauloRohling/txplorer/internal/presentation/rest/webserver"
	"github.com/pauloRohling/txplorer/internal/presentation/service"
	"net/http"
)

type UserRouter struct {
	userService service.UserService
}

func NewUserRouter(userService service.UserService) *UserRouter {
	return &UserRouter{userService: userService}
}

func (router *UserRouter) Endpoint() string {
	return webserver.UsersApi
}

func (router *UserRouter) Route(r chi.Router) {
	r.Post("/", webserver.Endpoint(router.Login, http.StatusOK))
}

func (router *UserRouter) Login(_ http.ResponseWriter, r *http.Request) (*user.LoginOutput, error) {
	input, err := json.Parse[user.LoginInput](r)
	if err != nil {
		return nil, err
	}

	return router.userService.Login(r.Context(), *input)
}
