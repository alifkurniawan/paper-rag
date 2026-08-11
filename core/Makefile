include .env
export

MIGRATION_DIR=internal/db/migrations

.PHONY: migrate-up migrate-down migrate-status sqlc-generate dev

migrate-up:
	goose -dir $(MIGRATION_DIR) postgres "$(DATABASE_URL)" up

migrate-down:
	goose -dir $(MIGRATION_DIR) postgres "$(DATABASE_URL)" down

migrate-status:
	goose -dir $(MIGRATION_DIR) postgres "$(DATABASE_URL)" status

sqlc-generate:
	sqlc generate

dev:
	air