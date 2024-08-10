package mapper

import (
	"github.com/pauloRohling/txplorer/internal/model"
	"github.com/pauloRohling/txplorer/internal/persistance/store"
)

type TransactionMapper struct {
}

func NewTransactionMapper() *TransactionMapper {
	return &TransactionMapper{}
}

func (mapper *TransactionMapper) ToModel(transaction store.Transaction) *model.Transaction {
	return &model.Transaction{
		ID:            transaction.ID,
		FromAccountID: transaction.FromAccountID,
		ToAccountID:   transaction.ToAccountID,
		Amount:        transaction.Amount,
		Timestamp:     &transaction.Timestamp,
		Status:        transaction.Status,
	}
}
