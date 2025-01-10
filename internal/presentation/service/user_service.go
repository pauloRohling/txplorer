package service

import (
	"context"
	"github.com/pauloRohling/txplorer/internal/application/user"
)

type UserService interface {
	Login(ctx context.Context, input user.LoginInput) (*user.LoginOutput, error)
}
