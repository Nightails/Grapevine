-- name: AddUser :one
INSERT INTO users (id, created_at, updated_at, password)
VALUES (?1, ?2, ?3, ?4)
RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users WHERE id = ?1;

-- name: UpdateUserByID :exec
UPDATE users
SET
    updated_at = ?2,
    password = ?3
WHERE id = ?1;

-- name: RemoveUserByID :exec
DELETE FROM users WHERE id = ?1;

-- name: RemoveAllUsers :exec
DELETE FROM users;