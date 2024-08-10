package model

import (
	"github.com/google/uuid"
	"time"
)

type Operation struct {
	ID            uuid.UUID `json:"id"`
	FromAccountID uuid.UUID `json:"fromAccountId"`
	ToAccountID   uuid.UUID `json:"toAccountId"`
	Amount        int64     `json:"amount"`
	Type          string    `json:"type"`
	CreatedAt     time.Time `json:"createdAt"`
	CreatedBy     uuid.UUID `json:"createdBy"`
	Status        string    `json:"status"`
}
