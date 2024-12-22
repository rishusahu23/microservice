MIGRATE_CMD=go run db/migrate.go
MIGRATIONS_DIR=./migrations

help:
	@echo "Usage:"
	@echo "  make create-migration NAME=<name>  # Create a new migration file"
	@echo "  make upgradeDB DB_NAME=<db_name>  # Apply migrations to a specific database"
	@echo "  make downgradeDB DB_NAME=<db_name>  # Rollback the last migration for a specific database"
	@echo "  make help                         # Show this help message"

create-migration:
ifndef DB_NAME
	$(error DB_NAME is required for creating a migration)
endif
	migrate create -ext sql -dir db/${DB_NAME}/migrations -seq ''


upgradeDB:
ifndef DB_NAME
	$(error DB_NAME is required to specify which database to migrate)
endif
	$(MIGRATE_CMD) up $(DB_NAME)

downgradeDB:
ifndef DB_NAME
	$(error DB_NAME is required to specify which database to rollback)
endif
	$(MIGRATE_CMD) down $(DB_NAME)

snapshotDB:
ifndef DB_NAME
	$(error DB_NAME is required to specify which database to snapshot)
endif
	@echo "> Taking snapshot of database schema"
	$(MIGRATE_CMD) snapshot $(DB_NAME)
	@echo "> Snapshot completed!"
