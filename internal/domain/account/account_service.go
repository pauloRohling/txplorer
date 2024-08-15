package account

import "context"

type Service struct {
	createAccountAction *CreateAccountAction
}

func NewService(createAccountAction *CreateAccountAction) *Service {
	return &Service{createAccountAction: createAccountAction}
}

func (service *Service) Create(ctx context.Context, input CreateAccountInput) (*CreateAccountOutput, error) {
	return service.createAccountAction.Execute(ctx, input)
}
