package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/pauloRohling/txplorer/internal/domain"
)

type AccountRepository interface {
	AddBalanceById(ctx context.Context, id uuid.UUID, balance int64) (*domain.Account, error)
	Create(ctx context.Context, userId uuid.UUID) (*domain.Account, error)
	GetById(ctx context.Context, id uuid.UUID) (*domain.Account, error)
	GetByUserId(ctx context.Context, userId uuid.UUID) (*domain.Account, error)
}
