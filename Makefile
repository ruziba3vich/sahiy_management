include .env
export

.PHONY: migrate-create migrate-up migrate-down migrate-force swagger

swagger:
	@swag init -g cmd/api/main.go -o docs
	@echo "Swagger docs generated in ./docs"

migrate-up:
	@migrate -path ./migrations -database "$(DB_URL)" up
	@echo "Migrations applied successfully"

migrate-down:
	@migrate -path ./migrations -database "$(DB_URL)" down 1
	@echo "Rolled back 1 migration"

migrate-force:
ifndef version
	$(error version is required. Usage: make migrate-force version=20260119120010)
endif
	@migrate -path ./migrations -database "$(DB_URL)" force $(version)
	@echo "Forced migration version to $(version)"

migrate-create:
ifndef name
	$(error name is required. Usage: make migrate-create name=create_users_table)
endif
	@mkdir -p ./migrations
	@timestamp=$$(date +%Y%m%d%H%M%S); \
	touch ./migrations/$${timestamp}_$(name).up.sql; \
	touch ./migrations/$${timestamp}_$(name).down.sql; \
	echo "Created migrations/$${timestamp}_$(name).up.sql"; \
	echo "Created migrations/$${timestamp}_$(name).down.sql"
