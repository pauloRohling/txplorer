package rest

import (
	"github.com/go-chi/chi/v5"
	"github.com/pauloRohling/txplorer/internal/domain/operation"
	"github.com/pauloRohling/txplorer/internal/presentation/rest/json"
	"github.com/pauloRohling/txplorer/internal/presentation/rest/webserver"
	"net/http"
)

type OperationRouter struct {
	transactionService OperationService
}

func NewOperationRouter(transactionService OperationService) *OperationRouter {
	return &OperationRouter{transactionService: transactionService}
}

func (router *OperationRouter) Endpoint() string {
	return webserver.OperationsApi
}

func (router *OperationRouter) Route(r chi.Router) {
	r.Post("/transfer", webserver.Endpoint(router.Transfer, http.StatusOK))
}

func (router *OperationRouter) Transfer(_ http.ResponseWriter, r *http.Request) (*operation.TransferOutput, error) {
	input, err := json.Parse[operation.TransferInput](r)
	if err != nil {
		return nil, err
	}

	return router.transactionService.Transfer(r.Context(), *input)
}
