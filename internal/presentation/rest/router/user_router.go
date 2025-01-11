package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/pauloRohling/txplorer/internal/application/user"
	"github.com/pauloRohling/txplorer/internal/presentation/rest/json"
	"github.com/pauloRohling/txplorer/internal/presentation/rest/types"
	"github.com/pauloRohling/txplorer/internal/presentation/rest/webserver"
	"net/http"
)

type UserRouter struct {
	userService user.Service
}

func NewUserRouter(userService user.Service) *UserRouter {
	return &UserRouter{userService: userService}
}

func (router *UserRouter) Endpoint() string {
	return webserver.UsersApi
}

func (router *UserRouter) Route(r chi.Router) {
	r.Post("/login", webserver.Endpoint(router.Login, http.StatusOK))
}

// Login godoc
//
//	@Summary		Login
//	@Description	Generates a JWT token to authenticate the user
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			credentials	body		types.LoginInput	true	"Email"
//	@Success		200			{object}	user.LoginOutput
//	@Failure		400			{object}	model.Error
//	@Failure		401			{object}	model.Error
//	@Failure		500			{object}	model.Error
//	@Router			/users/login [post]
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
