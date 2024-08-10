package mapper

import (
	"github.com/pauloRohling/txplorer/internal/model"
	"github.com/pauloRohling/txplorer/internal/persistance/store"
)

type AccountMapper struct {
}

func NewAccountMapper() *AccountMapper {
	return &AccountMapper{}
}

func (mapper *AccountMapper) ToModel(account store.Account) *model.Account {
	return &model.Account{
		ID:      account.ID,
		Balance: account.Balance,
	}
}
