# SQLBoiler Usage

This project uses [SQLBoiler](https://github.com/volatiletech/sqlboiler) to generate type-safe Go models from the postgres database schema.

## Configuration

Settings are defined in the `sqlboiler.toml` file.

## Code Generation

To generate Go models with minimal CRUD logic (no joins, no hooks), run:

```bash
sqlboiler psql
```

## Notes

- Make sure you're inside the internal/db directory when running the command,
  as the sqlboiler.toml file is expected to be located there
- You must run database migrations first (`make up`) so the tables exist
- Should regenerate after any schema change
