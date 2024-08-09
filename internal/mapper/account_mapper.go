package mapper

import (
	"xplorer/internal/model"
	"xplorer/internal/persistance/queries"
)

type AccountMapper struct {
}

func NewAccountMapper() *AccountMapper {
	return &AccountMapper{}
}

func (mapper *AccountMapper) ToModel(account queries.Account) *model.Account {
	return &model.Account{
		ID:      account.ID,
		Balance: account.Balance,
	}
}
