package model

import (
	"github.com/google/uuid"
)

type Account struct {
	ID      uuid.UUID `json:"id"`
	Balance int64     `json:"balance"`
}

func NewAccount(id uuid.UUID) *Account {
	return &Account{
		ID:      id,
		Balance: 0,
	}
}
