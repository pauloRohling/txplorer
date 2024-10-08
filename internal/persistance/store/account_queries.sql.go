// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: account_queries.sql

package store

import (
	"context"

	"github.com/google/uuid"
)

const addBalanceById = `-- name: AddBalanceById :one
UPDATE accounts
SET balance = balance + $1
WHERE id = $2
RETURNING id, balance, user_id, created_at, updated_at, status
`

type AddBalanceByIdParams struct {
	Balance int64     `json:"balance"`
	ID      uuid.UUID `json:"id"`
}

func (q *Queries) AddBalanceById(ctx context.Context, arg AddBalanceByIdParams) (Account, error) {
	row := q.db.QueryRowContext(ctx, addBalanceById, arg.Balance, arg.ID)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Balance,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Status,
	)
	return i, err
}

const getAccountById = `-- name: GetAccountById :one
SELECT id, balance, user_id, created_at, updated_at, status
FROM accounts
WHERE id = $1
`

func (q *Queries) GetAccountById(ctx context.Context, id uuid.UUID) (Account, error) {
	row := q.db.QueryRowContext(ctx, getAccountById, id)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Balance,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Status,
	)
	return i, err
}

const getAccountByUserId = `-- name: GetAccountByUserId :one
SELECT id, balance, user_id, created_at, updated_at, status
FROM accounts
WHERE user_id = $1
`

func (q *Queries) GetAccountByUserId(ctx context.Context, userID uuid.UUID) (Account, error) {
	row := q.db.QueryRowContext(ctx, getAccountByUserId, userID)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Balance,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Status,
	)
	return i, err
}

const insertAccount = `-- name: InsertAccount :one
INSERT INTO accounts (id, user_id)
VALUES ($1, $2)
RETURNING id, balance, user_id, created_at, updated_at, status
`

type InsertAccountParams struct {
	ID     uuid.UUID `json:"id"`
	UserID uuid.UUID `json:"userId"`
}

func (q *Queries) InsertAccount(ctx context.Context, arg InsertAccountParams) (Account, error) {
	row := q.db.QueryRowContext(ctx, insertAccount, arg.ID, arg.UserID)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Balance,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Status,
	)
	return i, err
}
