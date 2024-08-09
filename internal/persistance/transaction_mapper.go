package persistance

import (
	"xplorer/internal/model"
	"xplorer/internal/persistance/queries"
)

type TransactionMapper interface {
	ToModel(transaction queries.Transaction) *model.Transaction
}
