package mapper

import (
	"github.com/pauloRohling/txplorer/internal/domain/account"
	"github.com/pauloRohling/txplorer/internal/persistence/store"
)

type AccountMapper struct {
}

func NewAccountMapper() *AccountMapper {
	return &AccountMapper{}
}

func (mapper *AccountMapper) ToModel(savedAccount store.Account) *account.Account {
	return &account.Account{
		ID:        savedAccount.ID,
		Balance:   savedAccount.Balance,
		UserID:    savedAccount.UserID,
		CreatedAt: savedAccount.CreatedAt,
		UpdatedAt: savedAccount.UpdatedAt,
		Status:    account.Status(savedAccount.Status),
	}
}
