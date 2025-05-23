# Database Migration

This project uses [golang-migrate](https://github.com/golang-migrate/migrate) for managing database schema changes.

## Prerequisites

- PostgreSQL
- golang-migrate CLI installed  
  You can install via Homebrew on macOS:

  ```
  brew install golang-migrate
  ```

## Migration Folder

All migration files are located in:

```
db/migrations/
```

## Database URL

The Makefile assumes the following local development database:

```
postgres://kaeruashi:kaeruashipass@localhost:5432/kaeruashi_dev?sslmode=disable
```

## Available Commands

### Create a new migration

```sh
make create name=create_examples
```

- Generates a timestamp-based version using JST (Asia/Tokyo timezone)
- Example output:

  ```
  20250518_154522_create_examples.up.sql
  20250518_154522_create_examples.down.sql
  ```

### Apply all up migrations

```sh
make up
```

### Roll back the last migration

```sh
make down
```

### Force the current version (manual override)

```sh
make force
```

### Show the current migration version

```sh
make version
```

## Notes

- All timestamps are generated based on **Asia/Tokyo timezone**, not UTC.
- Migration format is `YYYYMMDD_HHMMSS_description`.
