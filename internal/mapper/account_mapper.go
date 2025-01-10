package mapper

import (
	"github.com/pauloRohling/txplorer/internal/domain"
	"github.com/pauloRohling/txplorer/internal/persistence/store"
)

type AccountMapper struct {
}

func NewAccountMapper() *AccountMapper {
	return &AccountMapper{}
}

func (mapper *AccountMapper) ToModel(account store.Account) *domain.Account {
	return &domain.Account{
		ID:        account.ID,
		Balance:   account.Balance,
		UserID:    account.UserID,
		CreatedAt: account.CreatedAt,
		UpdatedAt: account.UpdatedAt,
		Status:    domain.AccountStatus(account.Status),
	}
}
