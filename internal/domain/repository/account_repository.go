package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/pauloRohling/txplorer/internal/model"
)

type AccountRepository interface {
	AddBalanceById(ctx context.Context, id uuid.UUID, balance int64) (*model.Account, error)
	Create(ctx context.Context, userId uuid.UUID) (*model.Account, error)
	GetById(ctx context.Context, id uuid.UUID) (*model.Account, error)
}
