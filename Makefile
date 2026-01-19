.PHONY: migrate

migrate:
ifndef name
	$(error name is required. Usage: make migrate name=create_users_table)
endif
	@mkdir -p ./migrations
	@timestamp=$$(date +%Y%m%d%H%M%S); \
	touch ./migrations/$${timestamp}_$(name).up.sql; \
	touch ./migrations/$${timestamp}_$(name).down.sql; \
	echo "Created migrations/$${timestamp}_$(name).up.sql"; \
	echo "Created migrations/$${timestamp}_$(name).down.sql"
