package rest

import (
	"net/http"
	"xplorer/internal/domain/transaction"
	"xplorer/internal/presentation/json"
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
