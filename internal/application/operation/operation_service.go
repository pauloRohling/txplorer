package operation

import (
	"context"
)

type Service struct {
	depositAction  *DepositAction
	transferAction *TransferAction
	withdrawAction *WithdrawAction
}

func NewService(depositAction *DepositAction, transferAction *TransferAction, withdrawAction *WithdrawAction) *Service {
	return &Service{
		depositAction:  depositAction,
		transferAction: transferAction,
		withdrawAction: withdrawAction,
	}
}

func (service *Service) Deposit(ctx context.Context, input DepositInput) (*DepositOutput, error) {
	return service.depositAction.Execute(ctx, input)
}

func (service *Service) Transfer(ctx context.Context, input TransferInput) (*TransferOutput, error) {
	return service.transferAction.Execute(ctx, input)
}

func (service *Service) Withdraw(ctx context.Context, input WithdrawInput) (*WithdrawOutput, error) {
	return service.withdrawAction.Execute(ctx, input)
}
