-- name: InsertAccount :one
INSERT INTO accounts (id, user_id)
VALUES ($1, $2)
RETURNING *;

-- name: AddBalanceById :one
UPDATE accounts
SET balance = balance + $1
WHERE id = $2
RETURNING *;

-- name: GetAccountById :one
SELECT *
FROM accounts
WHERE id = $1;

-- name: GetAccountByUserId :one
SELECT *
FROM accounts
WHERE user_id = $1;