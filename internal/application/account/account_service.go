package account

import "context"

type Service struct {
	createAccountAction *CreateAccountAction
	getAccountAction    *GetAccountAction
}

func NewService(createAccountAction *CreateAccountAction, getAccountAction *GetAccountAction) *Service {
	return &Service{
		createAccountAction: createAccountAction,
		getAccountAction:    getAccountAction,
	}
}

func (service *Service) Create(ctx context.Context, input CreateAccountInput) (*CreateAccountOutput, error) {
	return service.createAccountAction.Execute(ctx, input)
}

func (service *Service) Get(ctx context.Context, input GetAccountInput) (*GetAccountOutput, error) {
	return service.getAccountAction.Execute(ctx, input)
}
