package user

import (
	"context"
)

type Repository interface {
	Create(ctx context.Context, name string, email string, password string) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
}
