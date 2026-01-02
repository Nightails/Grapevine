-- Test Data for BudgetCLI

-- Insert Sample Users
-- Note: Dates are in ISO8601 format as expected by SQLite TEXT fields
INSERT INTO users (id, created_at, updated_at, password) VALUES
(1, '2025-12-01 10:00:00', '2025-12-01 10:00:00', 'hashed_pass_1'),
(2, '2025-12-05 14:30:00', '2025-12-05 14:30:00', 'hashed_pass_2');

-- Insert Sample Transactions
-- User 1 Transactions
INSERT INTO transactions (user_id, amount, description, status, date, created_at, updated_at) VALUES
(1, 5000, 'Monthly Salary', 'completed', '2025-12-01', '2025-12-01 10:05:00', '2025-12-01 10:05:00'),
(1, -120, 'Grocery Store', 'completed', '2025-12-02', '2025-12-02 18:00:00', '2025-12-02 18:00:00'),
(1, -45, 'Gas Station', 'pending', '2025-12-03', '2025-12-03 09:15:00', '2025-12-03 09:15:00');

-- User 2 Transactions
INSERT INTO transactions (user_id, amount, description, status, date, created_at, updated_at) VALUES
(2, 4500, 'Freelance Project', 'completed', '2025-12-06', '2025-12-06 11:00:00', '2025-12-06 11:00:00'),
(2, -80, 'Internet Bill', 'completed', '2025-12-07', '2025-12-07 10:00:00', '2025-12-07 10:00:00');