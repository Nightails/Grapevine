-- name: AddTransaction :one
INSERT INTO transactions (id, created_at, updated_at, user_id, account_id, amount, currency, description, status, date)
VALUES (?1, ?2, ?3, ?4, ?5, ?6, ?7, ?8, ?9, ?10)
RETURNING *;

-- name: GetAllTransactions :many
SELECT * FROM transactions;

-- name: GetTransactionById :one
SELECT * FROM transactions WHERE id = ?1;

-- name: UpdateTransactionById :exec
UPDATE transactions
SET
    updated_at = ?2,
    amount = ?3,
    description = ?4,
    status = ?5,
    date = ?6
WHERE id = ?1;

-- name: RemoveTransactionById :exec
DELETE FROM transactions WHERE id = ?1;

-- name: RemoveAllTransactions :exec
DELETE FROM transactions;