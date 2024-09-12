package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/pauloRohling/txplorer/internal/domain/account"
	"github.com/pauloRohling/txplorer/internal/presentation/rest/json"
	"github.com/pauloRohling/txplorer/internal/presentation/rest/types"
	"github.com/pauloRohling/txplorer/internal/presentation/rest/webserver"
	"github.com/pauloRohling/txplorer/internal/presentation/service"
	"net/http"
)

type AccountRouter struct {
	accountService service.AccountService
}

func NewAccountRouter(accountService service.AccountService) *AccountRouter {
	return &AccountRouter{accountService: accountService}
}

func (router *AccountRouter) Endpoint() string {
	return webserver.AccountsApi
}

func (router *AccountRouter) Route(r chi.Router) {
	r.Post("/", webserver.Endpoint(router.Create, http.StatusOK))
}

// Create godoc
//
//	@Summary		Create Account
//	@Description	Creates a new account and a new user
//	@Tags			Account
//	@Accept			json
//	@Produce		json
//	@Param			account	body		types.CreateAccountInput	true	"Account"
//	@Success		200		{object}	account.CreateAccountOutput
//	@Failure		400		{object}	model.Error
//	@Failure		401		{object}	model.Error
//	@Failure		500		{object}	model.Error
//	@Router			/accounts [post]
func (router *AccountRouter) Create(_ http.ResponseWriter, r *http.Request) (*account.CreateAccountOutput, error) {
	jsonInput, err := json.Parse[types.CreateAccountInput](r)
	if err != nil {
		return nil, err
	}

	input := &account.CreateAccountInput{
		Name:     jsonInput.Name,
		Email:    jsonInput.Email,
		Password: jsonInput.Password,
	}

	return router.accountService.Create(r.Context(), *input)
}
