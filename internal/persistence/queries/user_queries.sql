-- name: InsertUser :one
INSERT INTO users (id, name, email, password)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: FindUserByEmail :one
SELECT *
FROM users
WHERE email = $1;