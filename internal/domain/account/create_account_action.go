package account

import (
	"context"
	"github.com/pauloRohling/txplorer/internal/domain/password"
	"github.com/pauloRohling/txplorer/internal/domain/repository"
	"github.com/pauloRohling/txplorer/internal/model"
	"github.com/pauloRohling/txplorer/pkg/transaction"
)

type CreateAccountInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateAccountOutput struct {
	*model.Account
}

type CreateAccountAction struct {
	accountRepository  repository.AccountRepository
	userRepository     repository.UserRepository
	transactionManager transaction.Manager
	passwordEncoder    password.Encoder
}

func NewCreateAccountAction(accountRepository repository.AccountRepository, userRepository repository.UserRepository, transactionManager transaction.Manager, passwordEncoder password.Encoder) *CreateAccountAction {
	return &CreateAccountAction{
		accountRepository:  accountRepository,
		userRepository:     userRepository,
		transactionManager: transactionManager,
		passwordEncoder:    passwordEncoder,
	}
}

func (action *CreateAccountAction) Execute(ctx context.Context, input CreateAccountInput) (*CreateAccountOutput, error) {
	if input.Name == "" || len(input.Name) < 2 || len(input.Name) > 255 {
		return nil, model.ValidationError("Name must be between 2 and 255 characters")
	}

	if input.Email == "" || len(input.Email) < 5 || len(input.Email) > 255 {
		return nil, model.ValidationError("Email must be between 5 and 255 characters")
	}

	if input.Password == "" || len(input.Password) < 8 || len(input.Password) > 128 {
		return nil, model.ValidationError("Password must be between 8 and 128 characters")
	}

	var err error
	input.Password, err = action.passwordEncoder.Encode(input.Password)
	if err != nil {
		return nil, model.InternalError("Failed to encode password", err)
	}

	var user *model.User
	var account *model.Account

	err = action.transactionManager.RunTransaction(ctx, func(ctx context.Context) error {
		user, err = action.userRepository.Create(ctx, input.Name, input.Email, input.Password)
		if err != nil {
			return model.InternalError("Failed to create user", err)
		}

		account, err = action.accountRepository.Create(ctx, user.ID)
		if err != nil {
			return model.InternalError("Failed to create account", err)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &CreateAccountOutput{Account: account}, nil
}
