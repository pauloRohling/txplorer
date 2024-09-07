package rest

import "github.com/google/uuid"

type DepositInput struct {
	AccountID uuid.UUID `json:"accountId"`
	Amount    int64     `json:"amount"`
}

type TransferInput struct {
	FromAccountID uuid.UUID `json:"fromAccountId"`
	ToAccountID   uuid.UUID `json:"toAccountId"`
	Amount        int64     `json:"amount"`
}

type WithdrawInput struct {
	AccountID uuid.UUID `json:"accountId"`
	Amount    int64     `json:"amount"`
}
