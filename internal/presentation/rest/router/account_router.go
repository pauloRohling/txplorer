package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/pauloRohling/txplorer/internal/domain/account"
	presentation "github.com/pauloRohling/txplorer/internal/presentation/rest/auth"
	"github.com/pauloRohling/txplorer/internal/presentation/rest/json"
	"github.com/pauloRohling/txplorer/internal/presentation/rest/middleware"
	"github.com/pauloRohling/txplorer/internal/presentation/rest/types"
	"github.com/pauloRohling/txplorer/internal/presentation/rest/webserver"
	"github.com/pauloRohling/txplorer/internal/presentation/service"
	"net/http"
)

type AccountRouter struct {
	accountService service.AccountService
	secretHolder   presentation.SecretHolder
}

func NewAccountRouter(accountService service.AccountService, secretHolder presentation.SecretHolder) *AccountRouter {
	return &AccountRouter{
		accountService: accountService,
		secretHolder:   secretHolder,
	}
}

func (router *AccountRouter) Endpoint() string {
	return webserver.AccountsApi
}

func (router *AccountRouter) Route(r chi.Router) {
	r.Post("/", webserver.Endpoint(router.Create, http.StatusOK))

	r.Route("/", func(r chi.Router) {
		secret := router.secretHolder.Get()
		r.Use(jwtauth.Verifier(secret))
		r.Use(middleware.Authenticator(secret))
		r.Get("/", webserver.Endpoint(router.Get, http.StatusOK))
	})
}

// Create godoc
//
//	@Summary		Create Account
//	@Description	Creates a new account and a new user
//	@Security		BearerAuth
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

// Get godoc
//
//	@Summary		Get Account
//	@Description	Gets an account by User ID from token
//	@Security		BearerAuth
//	@Tags			Account
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	account.GetAccountOutput
//	@Failure		400	{object}	model.Error
//	@Failure		401	{object}	model.Error
//	@Failure		500	{object}	model.Error
//	@Router			/accounts [get]
func (router *AccountRouter) Get(_ http.ResponseWriter, r *http.Request) (*account.GetAccountOutput, error) {
	userId, err := middleware.GetUserId(r.Context())
	if err != nil {
		return nil, err
	}

	input := &account.GetAccountInput{UserID: userId}
	return router.accountService.Get(r.Context(), *input)
}
