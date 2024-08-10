package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/pauloRohling/txplorer/internal/model"
)

type AccountRepository interface {
	AddBalanceById(ctx context.Context, id uuid.UUID, balance int64) (*model.Account, error)
	Create(ctx context.Context, id uuid.UUID, userId uuid.UUID) (*model.Account, error)
}
