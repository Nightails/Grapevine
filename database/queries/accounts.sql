-- name: AddAccount :one
INSERT INTO accounts (id, created_at, updated_at, name, type, balance, currency, user_id)
VALUES (?1, ?2, ?3, ?4, ?5, ?6, ?7, ?8)
RETURNING *;

-- name: RemoveAccountById :exec
DELETE FROM accounts WHERE id = ?1;