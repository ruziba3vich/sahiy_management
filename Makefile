.PHONY: migrate-create swagger

swagger:
	@swag init -g cmd/api/main.go -o docs
	@echo "Swagger docs generated in ./docs"

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
