package account

import (
	"github.com/pauloRohling/txplorer/internal/domain/account"
	"github.com/pauloRohling/txplorer/internal/persistence/store"
)

type Mapper interface {
	ToModel(account store.Account) *account.Account
}

type StoreMapper struct {
}

func NewStoreMapper() *StoreMapper {
	return &StoreMapper{}
}

func (mapper *StoreMapper) ToModel(savedAccount store.Account) *account.Account {
	return &account.Account{
		ID:        savedAccount.ID,
		Balance:   savedAccount.Balance,
		UserID:    savedAccount.UserID,
		CreatedAt: savedAccount.CreatedAt,
		UpdatedAt: savedAccount.UpdatedAt,
		Status:    account.Status(savedAccount.Status),
	}
}

var _ Mapper = (*StoreMapper)(nil)
