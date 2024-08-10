package rest

import (
	"github.com/pauloRohling/txplorer/internal/domain/transaction"
	"github.com/pauloRohling/txplorer/internal/presentation/json"
	"net/http"
)

type TransactionRouter struct {
	transactionService TransactionService
}

func NewTransactionRouter(transactionService TransactionService) *TransactionRouter {
	return &TransactionRouter{transactionService: transactionService}
}

func (router *TransactionRouter) Transfer(_ http.ResponseWriter, r *http.Request) (*transaction.TransferOutput, error) {
	input, err := json.Parse[transaction.TransferInput](r)
	if err != nil {
		return nil, err
	}

	return router.transactionService.Transfer(r.Context(), *input)
}
