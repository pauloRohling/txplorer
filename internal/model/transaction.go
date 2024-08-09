package model

import (
	"github.com/google/uuid"
	"time"
)

type Transaction struct {
	ID            uuid.UUID  `json:"id"`
	FromAccountID uuid.UUID  `json:"fromAccountId"`
	ToAccountID   uuid.UUID  `json:"toAccountId"`
	Amount        int64      `json:"amount"`
	Timestamp     *time.Time `json:"timestamp"`
	Status        string     `json:"status"`
}

func NewTransaction(fromAccountID uuid.UUID, toAccountID uuid.UUID, amount int64) *Transaction {
	timestamp := time.Now().UTC()
	return &Transaction{
		ID:            uuid.New(),
		FromAccountID: fromAccountID,
		ToAccountID:   toAccountID,
		Amount:        amount,
		Timestamp:     &timestamp,
		Status:        TransactionStatusPending.String(),
	}
}
