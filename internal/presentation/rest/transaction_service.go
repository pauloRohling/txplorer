package rest

import (
	"context"
	"xplorer/internal/domain/transaction"
)

type TransactionService interface {
	Transfer(ctx context.Context, input transaction.TransferInput) (*transaction.TransferOutput, error)
}
