# Database Documentation

This document describes the database schema for the BudgetCLI application. The database uses SQLite.

## Tables

### `users`

Stores information about the users of the application.

| Column | Type | Constraints | Description |
| :--- | :--- | :--- | :--- |
| `id` | INTEGER | PRIMARY KEY | Unique identifier for the user. |
| `created_at` | TEXT | NOT NULL | Timestamp when the user was created. |
| `updated_at` | TEXT | NOT NULL | Timestamp when the user was last updated. |
| `password` | TEXT | NOT NULL, DEFAULT 'empty' | User's hashed password. |

### `transactions`

Stores financial transactions for users.

| Column | Type | Constraints | Description |
| :--- | :--- | :--- | :--- |
| `id` | INTEGER | PRIMARY KEY | Unique identifier for the transaction. |
| `created_at` | TEXT | NOT NULL | Timestamp when the transaction record was created. |
| `updated_at` | TEXT | NOT NULL | Timestamp when the transaction record was last updated. |
| `user_id` | INTEGER | NOT NULL, FOREIGN KEY | Reference to the `users` table. |
| `amount` | INTEGER | NOT NULL, DEFAULT 0 | The transaction amount (likely in cents or smallest currency unit). |
| `description` | TEXT | NOT NULL, DEFAULT '' | A brief description of the transaction. |
| `status` | TEXT | NOT NULL, DEFAULT 'pending' | Status of the transaction (e.g., pending, completed). |
| `date` | TEXT | NOT NULL, DEFAULT CURRENT_DATE | The date of the transaction. |

## Relationships

- `transactions.user_id` references `users.id` with `ON DELETE CASCADE`.
