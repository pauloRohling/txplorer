package mapper

import (
	"github.com/pauloRohling/txplorer/internal/domain/account"
	"github.com/pauloRohling/txplorer/internal/persistence/store"
)

type AccountMapper interface {
	ToModel(account store.Account) *account.Account
}
