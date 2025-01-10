-- name: InsertOperation :one
INSERT INTO operations (id, from_account_id, to_account_id, amount, type, created_by)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: UpdateOperationStatus :one
UPDATE operations
SET status = $1
WHERE id = $2
RETURNING *;
