-- name: AddMerchant :one
INSERT INTO merchants (id, created_at, updated_at, name)
VALUES (?1, ?2, ?3, ?4)
RETURNING *;

-- name: RemoveMerchantById :exec
DELETE FROM merchants WHERE id = ?1;