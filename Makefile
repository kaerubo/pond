MIGRATION_DIR = db/migrations
DB_URL=postgres://kaeruashi:pass@localhost:5432/kaeruashi-dev?sslmode=disable
TIME_FORMAT = 20060102_150405
MIGRATE=migrate -path $(MIGRATION_DIR) -database $(DB_URL)

.PHONY: create up down force version

create:
	migrate create -ext sql -dir $(MIGRATION_DIR) -format $(TIME_FORMAT) -tz "Asia/Tokyo" $(name)

up:
	$(MIGRATE) up

down:
	$(MIGRATE) down 1

force:
	$(MIGRATE) force

version:
	$(MIGRATE) version
