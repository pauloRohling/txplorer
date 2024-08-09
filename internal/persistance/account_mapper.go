package persistance

import (
	"xplorer/internal/model"
	"xplorer/internal/persistance/queries"
)

type AccountMapper interface {
	ToModel(account queries.Account) *model.Account
}
