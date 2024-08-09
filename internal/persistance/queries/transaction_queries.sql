-- name: InsertTransaction :one
INSERT INTO transactions (id, from_account_id, to_account_id, amount, timestamp, status)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: UpdateTransactionStatus :one
UPDATE transactions
SET status = $1
WHERE id = $2
RETURNING *;
