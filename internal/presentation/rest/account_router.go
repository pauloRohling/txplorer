package rest

import (
	"github.com/go-chi/chi/v5"
	"github.com/pauloRohling/txplorer/internal/domain/account"
	"github.com/pauloRohling/txplorer/internal/presentation/rest/json"
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

func (router *AccountRouter) Create(_ http.ResponseWriter, r *http.Request) (*account.CreateAccountOutput, error) {
	input, err := json.Parse[account.CreateAccountInput](r)
	if err != nil {
		return nil, err
	}

	return router.accountService.Create(r.Context(), *input)
}
