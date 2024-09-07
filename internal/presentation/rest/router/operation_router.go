package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/pauloRohling/txplorer/internal/domain/operation"
	presentation "github.com/pauloRohling/txplorer/internal/presentation/rest/auth"
	"github.com/pauloRohling/txplorer/internal/presentation/rest/json"
	"github.com/pauloRohling/txplorer/internal/presentation/rest/middleware"
	"github.com/pauloRohling/txplorer/internal/presentation/rest/types"
	"github.com/pauloRohling/txplorer/internal/presentation/rest/webserver"
	"github.com/pauloRohling/txplorer/internal/presentation/service"
	"net/http"
)

type OperationRouter struct {
	operationService service.OperationService
	secretHolder     presentation.SecretHolder
}

func NewOperationRouter(operationService service.OperationService, secretHolder presentation.SecretHolder) *OperationRouter {
	return &OperationRouter{operationService: operationService, secretHolder: secretHolder}
}

func (router *OperationRouter) Endpoint() string {
	return webserver.OperationsApi
}

func (router *OperationRouter) Route(r chi.Router) {
	secret := router.secretHolder.Get()
	r.Use(jwtauth.Verifier(secret))
	r.Use(middleware.Authenticator(secret))
	r.Post("/deposit", webserver.Endpoint(router.Deposit, http.StatusOK))
	r.Post("/transfer", webserver.Endpoint(router.Transfer, http.StatusOK))
	r.Post("/withdraw", webserver.Endpoint(router.Withdraw, http.StatusOK))
}

func (router *OperationRouter) Deposit(_ http.ResponseWriter, r *http.Request) (*operation.DepositOutput, error) {
	userId, err := middleware.GetUserId(r.Context())
	if err != nil {
		return nil, err
	}

	jsonInput, err := json.Parse[types.DepositInput](r)
	if err != nil {
		return nil, err
	}

	input := &operation.DepositInput{
		AccountID:   jsonInput.AccountID,
		RequesterID: userId,
		Amount:      jsonInput.Amount,
	}

	return router.operationService.Deposit(r.Context(), *input)
}

func (router *OperationRouter) Transfer(_ http.ResponseWriter, r *http.Request) (*operation.TransferOutput, error) {
	userId, err := middleware.GetUserId(r.Context())
	if err != nil {
		return nil, err
	}

	jsonInput, err := json.Parse[types.TransferInput](r)
	if err != nil {
		return nil, err
	}

	input := &operation.TransferInput{
		FromAccountID: jsonInput.FromAccountID,
		ToAccountID:   jsonInput.ToAccountID,
		RequesterID:   userId,
		Amount:        jsonInput.Amount,
	}

	return router.operationService.Transfer(r.Context(), *input)
}

func (router *OperationRouter) Withdraw(_ http.ResponseWriter, r *http.Request) (*operation.WithdrawOutput, error) {
	userId, err := middleware.GetUserId(r.Context())
	if err != nil {
		return nil, err
	}

	jsonInput, err := json.Parse[types.WithdrawInput](r)
	if err != nil {
		return nil, err
	}

	input := &operation.WithdrawInput{
		AccountID:   jsonInput.AccountID,
		RequesterID: userId,
		Amount:      jsonInput.Amount,
	}

	return router.operationService.Withdraw(r.Context(), *input)
}
