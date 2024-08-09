package transaction

import (
	"context"
	"github.com/google/uuid"
	"xplorer/internal/model"
)

type AccountRepository interface {
	AddBalanceById(ctx context.Context, id uuid.UUID, balance int64) (*model.Account, error)
}
