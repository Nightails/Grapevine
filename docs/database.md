# Database Documentation

This document describes the database schema for the BudgetCLI application. The database uses SQLite.

## Tables

### `users`

Stores information about the users of the application.

| Column | Type | Constraints | Description |
| :--- | :--- | :--- | :--- |
| `id` | TEXT | PRIMARY KEY | Unique identifier for the user. |
| `created_at` | TEXT | NOT NULL | Timestamp when the user was created. |
| `updated_at` | TEXT | NOT NULL | Timestamp when the user was last updated. |
| `username` | TEXT | NOT NULL, UNIQUE | User's unique login name. |
| `password` | TEXT | NOT NULL, DEFAULT 'empty' | User's hashed password. |

### `accounts`

Stores financial accounts (e.g., checking, savings) belonging to users.

| Column | Type | Constraints | Description |
| :--- | :--- | :--- | :--- |
| `id` | TEXT | PRIMARY KEY | Unique identifier for the account. |
| `created_at` | TEXT | NOT NULL | Timestamp when the account was created. |
| `updated_at` | TEXT | NOT NULL | Timestamp when the account was last updated. |
| `name` | TEXT | NOT NULL | The name of the account. |
| `type` | TEXT | NOT NULL | The type of account (e.g., credit, debit). |
| `balance` | INTEGER | NOT NULL | Current balance in the smallest currency unit. |
| `currency` | TEXT | NOT NULL | Currency code (e.g., USD). |
| `user_id` | TEXT | NOT NULL, FOREIGN KEY | Reference to the `users` table. |

### `transactions`

Stores financial transactions for users.

| Column | Type | Constraints | Description |
| :--- | :--- | :--- | :--- |
| `id` | TEXT | PRIMARY KEY | Unique identifier for the transaction. |
| `created_at` | TEXT | NOT NULL | Timestamp when the transaction record was created. |
| `updated_at` | TEXT | NOT NULL | Timestamp when the transaction record was last updated. |
| `user_id` | TEXT | NOT NULL, FOREIGN KEY | Reference to the `users` table. |
| `account_id` | TEXT | NOT NULL, FOREIGN KEY | Reference to the `accounts` table. |
| `amount` | INTEGER | NOT NULL, DEFAULT 0 | The transaction amount in the smallest currency unit. |
| `currency` | TEXT | NOT NULL, DEFAULT 'USD' | Currency of the transaction. |
| `description` | TEXT | NOT NULL, DEFAULT '' | A brief description of the transaction. |
| `status` | TEXT | NOT NULL, DEFAULT 'pending' | Status of the transaction (e.g., pending, completed). |
| `date` | TEXT | NOT NULL, DEFAULT CURRENT_DATE | The date of the transaction. |

### `merchants`

Stores information about merchants involved in transactions.

| Column | Type | Constraints | Description |
| :--- | :--- | :--- | :--- |
| `id` | INTEGER | PRIMARY KEY | Unique identifier for the merchant. |
| `created_at` | TEXT | NOT NULL | Timestamp when the merchant was created. |
| `updated_at` | TEXT | NOT NULL | Timestamp when the merchant was last updated. |
| `name` | TEXT | NOT NULL | The name of the merchant. |

## Relationships

- `accounts.user_id` references `users.id` with `ON DELETE CASCADE`.
- `transactions.user_id` references `users.id` with `ON DELETE CASCADE`.
- `transactions.account_id` references `accounts.id` with `ON DELETE CASCADE`.
