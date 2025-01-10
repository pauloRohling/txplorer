package account

import (
	"context"
	"github.com/google/uuid"
	"github.com/pauloRohling/txplorer/internal/application/repository"
	"github.com/pauloRohling/txplorer/internal/domain"
)

type GetAccountInput struct {
	UserID uuid.UUID `json:"userId"`
}

type GetAccountOutput struct {
	*domain.Account
}

type GetAccountAction struct {
	accountRepository repository.AccountRepository
}

func NewGetAccountAction(accountRepository repository.AccountRepository) *GetAccountAction {
	return &GetAccountAction{accountRepository: accountRepository}
}

func (action *GetAccountAction) Execute(ctx context.Context, input GetAccountInput) (*GetAccountOutput, error) {
	account, err := action.accountRepository.GetByUserId(ctx, input.UserID)
	if err != nil {
		return nil, err
	}

	return &GetAccountOutput{Account: account}, nil
}
