package repository

import (
	"context"
	"github.com/pauloRohling/txplorer/internal/model"
)

type UserRepository interface {
	Create(ctx context.Context, name string, email string, password string) (*model.User, error)
	FindByEmail(ctx context.Context, email string) (*model.User, error)
}
