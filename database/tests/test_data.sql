-- Test Data for BudgetCLI

-- Insert Sample Users
-- Note: Dates are in ISO8601 format as expected by SQLite TEXT fields
INSERT INTO users (id, created_at, updated_at, username, password) VALUES
('1', '2025-12-01 10:00:00', '2025-12-01 10:00:00', 'jdoe',   'hashed_pass_1'),
('2', '2025-12-05 14:30:00', '2025-12-05 14:30:00', 'asmith', 'hashed_pass_2');

-- Insert Sample Merchants
INSERT INTO merchants (id, created_at, updated_at, name) VALUES
('m1', '2025-12-01 09:00:00', '2025-12-01 09:00:00', 'SuperMart'),
('m2', '2025-12-01 09:00:00', '2025-12-01 09:00:00', 'GasAndGo'),
('m3', '2025-12-01 09:00:00', '2025-12-01 09:00:00', 'InternetProvider');

-- Insert Sample Accounts
INSERT INTO accounts (id, user_id, name, type, balance, currency, created_at, updated_at) VALUES
('a1', '1', 'Main Checking', 'checking',  500000,  'USD', '2025-12-01 10:01:00', '2025-12-01 10:01:00'),
('a2', '1', 'Savings',       'savings',  1200000,  'USD', '2025-12-01 10:01:00', '2025-12-01 10:01:00'),
('a3', '2', 'Checking',      'checking',  350000,  'USD', '2025-12-06 11:01:00', '2025-12-06 11:01:00');

-- Insert Sample Transactions
-- Note: `merchant` exists on `transactions` (TEXT, nullable) in migration 003.
-- User 1 Transactions
INSERT INTO transactions (id, user_id, account_id, merchant, amount, currency, description, status, date, created_at, updated_at) VALUES
('t1', '1', 'a1', 'SuperMart',        500000, 'USD', 'Monthly Salary',      'completed', '2025-12-01', '2025-12-01 10:05:00', '2025-12-01 10:05:00'),
('t2', '1', 'a1', 'SuperMart',        -12050, 'USD', 'Grocery Store',       'completed', '2025-12-02', '2025-12-02 18:00:00', '2025-12-02 18:00:00'),
('t3', '1', 'a1', 'GasAndGo',          -4500, 'USD', 'Gas Station',         'pending',   '2025-12-03', '2025-12-03 09:15:00', '2025-12-03 09:15:00'),
('t4', '1', 'a2', 'InternetProvider',  -2500, 'EUR', 'Online Subscription', 'completed', '2025-12-04', '2025-12-04 14:20:00', '2025-12-04 14:20:00');

-- User 2 Transactions
INSERT INTO transactions (id, user_id, account_id, merchant, amount, currency, description, status, date, created_at, updated_at) VALUES
('t5', '2', 'a3', NULL,           450000, 'USD', 'Freelance Project', 'completed', '2025-12-06', '2025-12-06 11:00:00', '2025-12-06 11:00:00'),
('t6', '2', 'a3', 'InternetProvider', -8000, 'USD', 'Internet Bill',     'completed', '2025-12-07', '2025-12-07 10:00:00', '2025-12-07 10:00:00'),
('t7', '2', 'a3', NULL,            -1500, 'USD', 'Coffee Shop',       'completed', '2025-12-08', '2025-12-08 08:30:00', '2025-12-08 08:30:00');