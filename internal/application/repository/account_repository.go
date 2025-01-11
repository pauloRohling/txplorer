package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/pauloRohling/txplorer/internal/domain/account"
)

type AccountRepository interface {
	AddBalanceById(ctx context.Context, id uuid.UUID, balance int64) (*account.Account, error)
	Create(ctx context.Context, userId uuid.UUID) (*account.Account, error)
	GetById(ctx context.Context, id uuid.UUID) (*account.Account, error)
	GetByUserId(ctx context.Context, userId uuid.UUID) (*account.Account, error)
}
