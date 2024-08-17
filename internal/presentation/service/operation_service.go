package service

import (
	"context"
	"github.com/pauloRohling/txplorer/internal/domain/operation"
)

type OperationService interface {
	Deposit(ctx context.Context, input operation.DepositInput) (*operation.DepositOutput, error)
	Transfer(ctx context.Context, input operation.TransferInput) (*operation.TransferOutput, error)
	Withdraw(ctx context.Context, input operation.WithdrawInput) (*operation.WithdrawOutput, error)
}
