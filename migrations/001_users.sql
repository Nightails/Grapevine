-- +goose Up
CREATE TABLE users (
    id TEXT PRIMARY KEY,
    created_at TEXT NOT NULL,
    updated_at TEXT NOT NULL,
    password TEXT NOT NULL DEFAULT 'empty'
);

-- +goose Down
DROP TABLE users;