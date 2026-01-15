-- Test Data for BudgetCLI

-- Insert Sample Users
-- Note: Dates are in ISO8601 format as expected by SQLite TEXT fields
-- Updated: Added 'username' field as required by migration 001
INSERT INTO users (id, username, created_at, updated_at, password) VALUES
(1, 'jdoe', '2025-12-01 10:00:00', '2025-12-01 10:00:00', 'hashed_pass_1'),
(2, 'asmith', '2025-12-05 14:30:00', '2025-12-05 14:30:00', 'hashed_pass_2');

-- Insert Sample Merchants (New)
INSERT INTO merchants (id, created_at, updated_at, name) VALUES
('1', '2025-12-01 09:00:00', '2025-12-01 09:00:00', 'SuperMart'),
('2', '2025-12-01 09:00:00', '2025-12-01 09:00:00', 'GasAndGo'),
('3', '2025-12-01 09:00:00', '2025-12-01 09:00:00', 'InternetProvider');

-- Insert Sample Accounts (New)
INSERT INTO accounts (id, user_id, name, type, balance, currency, created_at, updated_at) VALUES
('1', '1', 'Main Checking', 'checking', 500000, 'USD', '2025-12-01 10:01:00', '2025-12-01 10:01:00'),
('2', '1', 'Savings', 'savings', 1200000, 'USD', '2025-12-01 10:01:00', '2025-12-01 10:01:00'),
('3', '2', 'Checking', 'checking', 350000, 'USD', '2025-12-06 11:01:00', '2025-12-06 11:01:00');

-- Insert Sample Transactions
-- User 1 Transactions
-- Updated: Added 'id', 'account_id', and 'currency' columns
INSERT INTO transactions (id, user_id, account_id, amount, currency, description, status, date, created_at, updated_at) VALUES
('1', '1', '1', 500000, 'USD', 'Monthly Salary', 'completed', '2025-12-01', '2025-12-01 10:05:00', '2025-12-01 10:05:00'),
('2', '1', '1', -12050, 'USD', 'Grocery Store', 'completed', '2025-12-02', '2025-12-02 18:00:00', '2025-12-02 18:00:00'),
('3', '1', '1', -4500, 'USD', 'Gas Station', 'pending', '2025-12-03', '2025-12-03 09:15:00', '2025-12-03 09:15:00'),
('4', '1', '2', -2500, 'EUR', 'Online Subscription', 'completed', '2025-12-04', '2025-12-04 14:20:00', '2025-12-04 14:20:00');

-- User 2 Transactions
INSERT INTO transactions (id, user_id, account_id, amount, currency, description, status, date, created_at, updated_at) VALUES
('5', '2', '3', 450000, 'USD', 'Freelance Project', 'completed', '2025-12-06', '2025-12-06 11:00:00', '2025-12-06 11:00:00'),
('6', '2', '3', -8000, 'USD', 'Internet Bill', 'completed', '2025-12-07', '2025-12-07 10:00:00', '2025-12-07 10:00:00'),
('7', '2', '3', -1500, 'USD', 'Coffee Shop', 'completed', '2025-12-08', '2025-12-08 08:30:00', '2025-12-08 08:30:00');