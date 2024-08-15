package service

import (
	"context"
	"github.com/pauloRohling/txplorer/internal/domain/account"
)

type AccountService interface {
	Create(ctx context.Context, input account.CreateAccountInput) (*account.CreateAccountOutput, error)
}
