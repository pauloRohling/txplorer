package mapper

import (
	"xplorer/internal/model"
	"xplorer/internal/persistance/queries"
)

type TransactionMapper struct {
}

func NewTransactionMapper() *TransactionMapper {
	return &TransactionMapper{}
}

func (mapper *TransactionMapper) ToModel(transaction queries.Transaction) *model.Transaction {
	return &model.Transaction{
		ID:            transaction.ID,
		FromAccountID: transaction.FromAccountID,
		ToAccountID:   transaction.ToAccountID,
		Amount:        transaction.Amount,
		Timestamp:     &transaction.Timestamp,
		Status:        transaction.Status,
	}
}
