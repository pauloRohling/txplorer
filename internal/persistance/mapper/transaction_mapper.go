package mapper

import (
	"github.com/pauloRohling/txplorer/internal/model"
	"github.com/pauloRohling/txplorer/internal/persistance/store"
)

type TransactionMapper interface {
	ToModel(transaction store.Transaction) *model.Transaction
}
