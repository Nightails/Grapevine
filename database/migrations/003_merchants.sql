-- +goose Up
CREATE TABLE merchants (
    id INTEGER PRIMARY KEY,
    created_at TEXT NOT NULL,
    updated_at TEXT NOT NULL,
    name TEXT NOT NULL
);

-- +goose Down
DROP TABLE merchants;