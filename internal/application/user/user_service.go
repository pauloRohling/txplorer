package user

import "context"

type Service interface {
	Login(ctx context.Context, input LoginInput) (*LoginOutput, error)
}

type FacadeService struct {
	loginAction *LoginAction
}

func NewFacadeService(loginAction *LoginAction) *FacadeService {
	return &FacadeService{loginAction: loginAction}
}

func (service *FacadeService) Login(ctx context.Context, input LoginInput) (*LoginOutput, error) {
	return service.loginAction.Execute(ctx, input)
}
