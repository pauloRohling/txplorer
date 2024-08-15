package account

import (
	"context"
	"fmt"
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
		return nil, fmt.Errorf("[CreateAccountAction] invalid name")
	}

	if input.Email == "" || len(input.Email) < 5 || len(input.Email) > 255 {
		return nil, fmt.Errorf("[CreateAccountAction] invalid email")
	}

	if input.Password == "" || len(input.Password) < 8 || len(input.Password) > 128 {
		return nil, fmt.Errorf("[CreateAccountAction] invalid password")
	}

	var err error
	input.Password, err = action.passwordEncoder.Encode(input.Password)
	if err != nil {
		return nil, fmt.Errorf("[CreateAccountAction] failed to encode password: %w", err)
	}

	var user *model.User
	var account *model.Account

	err = action.transactionManager.RunTransaction(ctx, func(ctx context.Context) error {
		user, err = action.userRepository.Create(ctx, input.Name, input.Email, input.Password)
		if err != nil {
			return fmt.Errorf("[CreateAccountAction] failed to create user: %w", err)
		}

		account, err = action.accountRepository.Create(ctx, user.ID)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &CreateAccountOutput{Account: account}, nil
}
