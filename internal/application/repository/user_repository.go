package repository

import (
	"context"
	"github.com/pauloRohling/txplorer/internal/domain"
)

type UserRepository interface {
	Create(ctx context.Context, name string, email string, password string) (*domain.User, error)
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
}
