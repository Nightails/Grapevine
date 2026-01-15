# Environment Variables

The project uses environment variables for database configuration. However, these are now **optional** as default values are provided or can be overridden via command-line flags.

| Variable              | Description                                      | Default Value         |
|-----------------------|--------------------------------------------------|-----------------------|
| `GOOSE_DRIVER`        | Database driver for Goose migrations             | `sqlite3`             |
| `GOOSE_DBSTRING`      | Path to the SQLite database file                 | `database/app.db`     |
| `GOOSE_MIGRATION_DIR` | Directory containing Goose migration files       | `database/migrations` |

## Usage with Goose

If you choose not to use environment variables, you can run migrations by providing the driver and connection string directly:

```bash
goose -dir database/migrations sqlite3 database/app.db up
```
