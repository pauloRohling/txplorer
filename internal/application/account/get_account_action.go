package account

import (
	"context"
	"github.com/google/uuid"
	"github.com/pauloRohling/txplorer/internal/domain/account"
)

type GetAccountInput struct {
	UserID uuid.UUID `json:"userId"`
}

type GetAccountOutput struct {
	*account.Account
}

type GetAccountAction struct {
	accountRepository account.Repository
}

func NewGetAccountAction(accountRepository account.Repository) *GetAccountAction {
	return &GetAccountAction{accountRepository: accountRepository}
}

func (action *GetAccountAction) Execute(ctx context.Context, input GetAccountInput) (*GetAccountOutput, error) {
	savedAccount, err := action.accountRepository.GetByUserId(ctx, input.UserID)
	if err != nil {
		return nil, err
	}

	return &GetAccountOutput{Account: savedAccount}, nil
}
