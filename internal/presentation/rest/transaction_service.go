package rest

import (
	"context"
	"github.com/pauloRohling/txplorer/internal/domain/transaction"
)

type TransactionService interface {
	Transfer(ctx context.Context, input transaction.TransferInput) (*transaction.TransferOutput, error)
}
