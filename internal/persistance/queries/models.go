// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package queries

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID      uuid.UUID `json:"id"`
	Balance int64     `json:"balance"`
}

type Transaction struct {
	ID            uuid.UUID `json:"id"`
	FromAccountID uuid.UUID `json:"fromAccountId"`
	ToAccountID   uuid.UUID `json:"toAccountId"`
	Amount        int64     `json:"amount"`
	Timestamp     time.Time `json:"timestamp"`
	Status        string    `json:"status"`
}
