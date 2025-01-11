package operation

import (
	"context"
)

type Service interface {
	Deposit(ctx context.Context, input DepositInput) (*DepositOutput, error)
	Transfer(ctx context.Context, input TransferInput) (*TransferOutput, error)
	Withdraw(ctx context.Context, input WithdrawInput) (*WithdrawOutput, error)
}

type FacadeService struct {
	depositAction  *DepositAction
	transferAction *TransferAction
	withdrawAction *WithdrawAction
}

func NewFacadeService(depositAction *DepositAction, transferAction *TransferAction, withdrawAction *WithdrawAction) *FacadeService {
	return &FacadeService{
		depositAction:  depositAction,
		transferAction: transferAction,
		withdrawAction: withdrawAction,
	}
}

func (service *FacadeService) Deposit(ctx context.Context, input DepositInput) (*DepositOutput, error) {
	return service.depositAction.Execute(ctx, input)
}

func (service *FacadeService) Transfer(ctx context.Context, input TransferInput) (*TransferOutput, error) {
	return service.transferAction.Execute(ctx, input)
}

func (service *FacadeService) Withdraw(ctx context.Context, input WithdrawInput) (*WithdrawOutput, error) {
	return service.withdrawAction.Execute(ctx, input)
}
