package rest

import (
	"context"
	"github.com/pauloRohling/txplorer/internal/domain/operation"
)

type TransactionService interface {
	Transfer(ctx context.Context, input operation.TransferInput) (*operation.TransferOutput, error)
}
