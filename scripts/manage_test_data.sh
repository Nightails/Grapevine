#!/usr/bin/env bash

DB_PATH="./database/app.db"
TEST_DATA="./database/tests/test_data.sql"

function check_db() {
    if [ ! -f "$DB_PATH" ]; then
        echo "❌ Error: Database file not found at $DB_PATH"
        echo "Please run migrations first: goose -dir database/migrations sqlite3 $DB_PATH up"
        exit 1
    fi
}

function empty() {
    check_db
    # Delete records to allow fresh insertion without constraint errors
    # Updated: Added accounts and merchants tables
    sqlite3 "$DB_PATH" "DELETE FROM transactions; DELETE FROM accounts; DELETE FROM merchants; DELETE FROM users;"
    echo "🧹 Database cleared."
}

function populate() {
    check_db
    if [ ! -f "$TEST_DATA" ]; then
        echo "❌ Error: $TEST_DATA not found."
        exit 1
    fi
    
    # Always empty before populating to prevent duplicate key errors
    empty
    
    sqlite3 "$DB_PATH" < "$TEST_DATA"
    echo "✅ Database populated with fresh test data."
}

case "$1" in
    up)
        populate
        ;;
    down)
        empty
        ;;
    *)
        echo "Usage: $0 {up|down}"
        exit 1
esac