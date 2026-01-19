-include .env
export

.PHONY: migrate-create migrate-up migrate-down migrate-force swagger build run docker-build docker-up docker-down docker-logs

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
	@version=$$(ls -1 ./migrations/*.up.sql 2>/dev/null | sort -r | head -1 | sed 's/.*\/\([0-9]*\)_.*/\1/'); \
	if [ -z "$$version" ]; then \
		echo "No migrations found"; \
		exit 1; \
	fi; \
	migrate -path ./migrations -database "$(DB_URL)" force $$version; \
	echo "Forced migration version to $$version"

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

# Build commands
build:
	@go build -o bin/api ./cmd/api
	@echo "Binary built at ./bin/api"

run:
	@go run ./cmd/api

# Docker commands
docker-build:
	@docker compose build
	@echo "Docker image built successfully"

docker-up:
	@docker compose up -d
	@echo "Services started. API available at http://localhost:$${PORT:-8080}"
	@echo "Swagger UI: http://localhost:$${PORT:-8080}/swagger/index.html"

docker-down:
	@docker compose down
	@echo "Services stopped"

docker-logs:
	@docker compose logs -f api

docker-clean:
	@docker compose down -v
	@echo "Services stopped and volumes removed"
