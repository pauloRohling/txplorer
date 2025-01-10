package mapper

import (
	"github.com/pauloRohling/txplorer/internal/model"
	"github.com/pauloRohling/txplorer/internal/persistence/store"
)

type AccountMapper interface {
	ToModel(account store.Account) *model.Account
}
