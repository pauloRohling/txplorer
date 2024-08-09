package transaction

import (
	"context"
	"github.com/google/uuid"
	"xplorer/internal/model"
)

type TransactionRepository interface {
	Create(ctx context.Context, entity *model.Transaction) (*model.Transaction, error)
	UpdateStatus(ctx context.Context, id uuid.UUID, status model.TransactionStatus) (*model.Transaction, error)
}
