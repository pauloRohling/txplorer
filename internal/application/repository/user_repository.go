package repository

import (
	"context"
	"github.com/pauloRohling/txplorer/internal/domain/user"
)

type UserRepository interface {
	Create(ctx context.Context, name string, email string, password string) (*user.User, error)
	FindByEmail(ctx context.Context, email string) (*user.User, error)
}
