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

// Deposit godoc
//
//	@Summary		Deposit
//	@Description	Deposits funds to an account
//	@Tags			Operation
//	@Accept			json
//	@Produce		json
//	@Param			account	body		types.DepositInput	true	"Account"
//	@Success		200		{object}	operation.DepositOutput
//	@Failure		400		{object}	model.Error
//	@Failure		401		{object}	model.Error
//	@Failure		500		{object}	model.Error
//	@Router			/operations/deposit [post]
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

// Transfer godoc
//
//	@Summary		Transfer
//	@Description	Transfers funds from one account to another
//	@Tags			Operation
//	@Accept			json
//	@Produce		json
//	@Param			account	body		types.TransferInput	true	"Account"
//	@Success		200		{object}	operation.TransferOutput
//	@Failure		400		{object}	model.Error
//	@Failure		401		{object}	model.Error
//	@Failure		500		{object}	model.Error
//	@Router			/operations/transfer [post]
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

// Withdraw godoc
//
//	@Summary		Withdraw
//	@Description	Withdraws funds from an account
//	@Tags			Operation
//	@Accept			json
//	@Produce		json
//	@Param			account	body		types.WithdrawInput	true	"Account"
//	@Success		200		{object}	operation.WithdrawOutput
//	@Failure		400		{object}	model.Error
//	@Failure		401		{object}	model.Error
//	@Failure		500		{object}	model.Error
//	@Router			/operations/withdraw [post]
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
