package service

import (
	"context"
	"github.com/pauloRohling/txplorer/internal/domain/account"
)

type AccountService interface {
	Create(ctx context.Context, input account.CreateAccountInput) (*account.CreateAccountOutput, error)
	Get(ctx context.Context, input account.GetAccountInput) (*account.GetAccountOutput, error)
}
