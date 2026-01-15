-- +goose Up
CREATE TABLE users (
    id TEXT PRIMARY KEY,
    created_at TEXT NOT NULL,
    updated_at TEXT NOT NULL,
    username TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL DEFAULT 'empty'
);

-- +goose Down
DROP TABLE users;