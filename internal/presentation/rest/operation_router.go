package rest

import (
	"github.com/go-chi/chi/v5"
	"github.com/pauloRohling/txplorer/internal/domain/operation"
	"github.com/pauloRohling/txplorer/internal/presentation/rest/json"
	"github.com/pauloRohling/txplorer/internal/presentation/rest/webserver"
	"github.com/pauloRohling/txplorer/internal/presentation/service"
	"net/http"
)

type OperationRouter struct {
	operationService service.OperationService
}

func NewOperationRouter(operationService service.OperationService) *OperationRouter {
	return &OperationRouter{operationService: operationService}
}

func (router *OperationRouter) Endpoint() string {
	return webserver.OperationsApi
}

func (router *OperationRouter) Route(r chi.Router) {
	r.Post("/deposit", webserver.Endpoint(router.Deposit, http.StatusOK))
	r.Post("/transfer", webserver.Endpoint(router.Transfer, http.StatusOK))
}

func (router *OperationRouter) Deposit(_ http.ResponseWriter, r *http.Request) (*operation.DepositOutput, error) {
	input, err := json.Parse[operation.DepositInput](r)
	if err != nil {
		return nil, err
	}

	return router.operationService.Deposit(r.Context(), *input)
}

func (router *OperationRouter) Transfer(_ http.ResponseWriter, r *http.Request) (*operation.TransferOutput, error) {
	input, err := json.Parse[operation.TransferInput](r)
	if err != nil {
		return nil, err
	}

	return router.operationService.Transfer(r.Context(), *input)
}
