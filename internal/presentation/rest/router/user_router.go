package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/pauloRohling/txplorer/internal/domain/user"
	"github.com/pauloRohling/txplorer/internal/presentation/rest/json"
	"github.com/pauloRohling/txplorer/internal/presentation/rest/types"
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
	r.Post("/login", webserver.Endpoint(router.Login, http.StatusOK))
}

func (router *UserRouter) Login(_ http.ResponseWriter, r *http.Request) (*user.LoginOutput, error) {
	jsonInput, err := json.Parse[types.LoginInput](r)
	if err != nil {
		return nil, err
	}

	input := &user.LoginInput{
		Email:    jsonInput.Email,
		Password: jsonInput.Password,
	}

	return router.userService.Login(r.Context(), *input)
}
