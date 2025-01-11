package account

import (
	"context"
	"github.com/pauloRohling/txplorer/internal/application/password"
	"github.com/pauloRohling/txplorer/internal/domain/account"
	"github.com/pauloRohling/txplorer/internal/domain/throw"
	"github.com/pauloRohling/txplorer/internal/domain/user"
	"github.com/pauloRohling/txplorer/pkg/transaction"
)

type CreateAccountInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateAccountOutput struct {
	*account.Account
}

type CreateAccountAction struct {
	accountRepository  account.Repository
	userRepository     user.Repository
	transactionManager transaction.Manager
	passwordEncoder    password.Encoder
}

func NewCreateAccountAction(accountRepository account.Repository, userRepository user.Repository, transactionManager transaction.Manager, passwordEncoder password.Encoder) *CreateAccountAction {
	return &CreateAccountAction{
		accountRepository:  accountRepository,
		userRepository:     userRepository,
		transactionManager: transactionManager,
		passwordEncoder:    passwordEncoder,
	}
}

func (action *CreateAccountAction) Execute(ctx context.Context, input CreateAccountInput) (*CreateAccountOutput, error) {
	if input.Name == "" || len(input.Name) < 2 || len(input.Name) > 255 {
		return nil, throw.ValidationError("Name must be between 2 and 255 characters")
	}

	if input.Email == "" || len(input.Email) < 5 || len(input.Email) > 255 {
		return nil, throw.ValidationError("Email must be between 5 and 255 characters")
	}

	if input.Password == "" || len(input.Password) < 8 || len(input.Password) > 128 {
		return nil, throw.ValidationError("Password must be between 8 and 128 characters")
	}

	var err error
	input.Password, err = action.passwordEncoder.Encode(input.Password)
	if err != nil {
		return nil, throw.InternalError("Failed to encode password", err)
	}

	var newUser *user.User
	var newAccount *account.Account

	err = action.transactionManager.RunTransaction(ctx, func(ctx context.Context) error {
		newUser, err = action.userRepository.Create(ctx, input.Name, input.Email, input.Password)
		if err != nil {
			return throw.InternalError("Failed to create user", err)
		}

		newAccount, err = action.accountRepository.Create(ctx, newUser.ID)
		if err != nil {
			return throw.InternalError("Failed to create account", err)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &CreateAccountOutput{Account: newAccount}, nil
}
