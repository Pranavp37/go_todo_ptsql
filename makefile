# Load environment variables
include .env
export

# Goose Commands

goose-status:
	goose -dir $(GOOSE_MIGRATION_DIR) status

goose-up:
	goose -dir $(GOOSE_MIGRATION_DIR) up

goose-down:
	goose -dir $(GOOSE_MIGRATION_DIR) down

goose-reset:
	goose -dir $(GOOSE_MIGRATION_DIR) reset

goose-create:
	goose -dir $(GOOSE_MIGRATION_DIR) create $(name) sql
