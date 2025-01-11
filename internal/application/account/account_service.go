package account

import "context"

type Service interface {
	Create(ctx context.Context, input CreateAccountInput) (*CreateAccountOutput, error)
	Get(ctx context.Context, input GetAccountInput) (*GetAccountOutput, error)
}

type FacadeService struct {
	createAccountAction *CreateAccountAction
	getAccountAction    *GetAccountAction
}

func NewFacadeService(createAccountAction *CreateAccountAction, getAccountAction *GetAccountAction) *FacadeService {
	return &FacadeService{
		createAccountAction: createAccountAction,
		getAccountAction:    getAccountAction,
	}
}

func (service *FacadeService) Create(ctx context.Context, input CreateAccountInput) (*CreateAccountOutput, error) {
	return service.createAccountAction.Execute(ctx, input)
}

func (service *FacadeService) Get(ctx context.Context, input GetAccountInput) (*GetAccountOutput, error) {
	return service.getAccountAction.Execute(ctx, input)
}

var _ Service = (*FacadeService)(nil)
