package account

import (
	"context"
	"github.com/google/uuid"
)

type Repository interface {
	AddBalanceById(ctx context.Context, id uuid.UUID, balance int64) (*Account, error)
	Create(ctx context.Context, userId uuid.UUID) (*Account, error)
	GetById(ctx context.Context, id uuid.UUID) (*Account, error)
	GetByUserId(ctx context.Context, userId uuid.UUID) (*Account, error)
}
