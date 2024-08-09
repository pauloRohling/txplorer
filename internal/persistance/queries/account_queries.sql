-- name: InsertAccount :one
INSERT INTO accounts (id)
VALUES ($1)
RETURNING *;

-- name: GetAccountById :one
SELECT *
FROM accounts
WHERE id = $1;

-- name: AddBalanceById :one
UPDATE accounts
SET balance = balance + $1
WHERE id = $2
RETURNING *;