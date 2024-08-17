package user

import "context"

type Service struct {
	loginAction *LoginAction
}

func NewService(loginAction *LoginAction) *Service {
	return &Service{loginAction: loginAction}
}

func (service *Service) Login(ctx context.Context, input LoginInput) (*LoginOutput, error) {
	return service.loginAction.Execute(ctx, input)
}
