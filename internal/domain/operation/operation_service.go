package operation

import (
	"context"
)

type Service struct {
	depositAction  *DepositAction
	transferAction *TransferAction
}

func NewService(depositAction *DepositAction, transferAction *TransferAction) *Service {
	return &Service{
		depositAction:  depositAction,
		transferAction: transferAction,
	}
}

func (service *Service) Deposit(ctx context.Context, input DepositInput) (*DepositOutput, error) {
	return service.depositAction.Execute(ctx, input)
}

func (service *Service) Transfer(ctx context.Context, input TransferInput) (*TransferOutput, error) {
	return service.transferAction.Execute(ctx, input)
}
