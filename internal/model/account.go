package model

import (
	"github.com/google/uuid"
	"time"
)

type Account struct {
	ID        uuid.UUID     `json:"id"`
	Balance   int64         `json:"balance"`
	UserID    uuid.UUID     `json:"userId"`
	CreatedAt time.Time     `json:"createdAt"`
	UpdatedAt time.Time     `json:"updatedAt"`
	Status    AccountStatus `json:"status"`
}
