package transaction

import (
	"context"
)

type Service struct {
	transferAction *TransferAction
}

func NewService(transferAction *TransferAction) *Service {
	return &Service{transferAction: transferAction}
}

func (service *Service) Transfer(ctx context.Context, input TransferInput) (*TransferOutput, error) {
	return service.transferAction.Execute(ctx, input)
}
