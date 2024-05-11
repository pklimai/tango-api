include architect.mk

#======================================#
# VARIABLES
#======================================#

## GENERAL
SERVICE_NAME = "tango-api"
MIGRATION_DIR = $(CURDIR)/migration
PG_DSN="host=localhost port=1331 user=tango-api-user password=1234 dbname=tango-api sslmode=disable"

## BIN
GOOSE_BIN = $(LOCAL_BIN)/goose


#======================================#
# INSTALLATION
#======================================#

.custom-bin-deps: export GOBIN := $(LOCAL_BIN)
.custom-bin-deps:
	$(info Installing custom bins for project...)    
	tmp=$$(mktemp -d) && cd $$tmp && go mod init temp && \
		go install github.com/pressly/goose/v3/cmd/goose@latest
		go install github.com/gojuno/minimock/v3/cmd/minimock@latest
	rm -rf $$tmp

custom-bin-deps: .custom-bin-deps ## install custom necessary bins

#======================================#
# GENERATION
#======================================#

.generate-mocks:
	$(info Generating mocks...) 
	PATH="$(LOCAL_BIN):$$PATH" go generate -x -run minimock ./...

generate-mocks: .generate-mocks ## geenrate mocks for tests

generate: .generate-mocks

#======================================#
# DOCKER
#======================================#

.compose-up: 
	docker compose -p $(SERVICE_NAME) -f ./local/docker/docker-compose.yml up -d

compose-down: 
	docker compose -p $(SERVICE_NAME) -f ./local/docker/docker-compose.yml down
 
.compose-rm:
	docker compose -p $(SERVICE_NAME) -f ./local/docker/docker-compose.yml rm -fvs  #f - force, v - any anonymous volumes, s - stop

.compose-rs:
	make compose-down
	make compose-up

compose-up: .compose-up ## start docker containers

compose-down: .compose-down ## stop docker containers

compose-rm: .compose-rm ## stop and remove docker containers

compose-rs: .compose-rs ## restart docker containers

#======================================#
# MIGRATIONS
#======================================#

.migration-up: 
	$(GOOSE_BIN) -dir $(MIGRATION_DIR) postgres $(PG_DSN) up

.migration-down:
	$(GOOSE_BIN) -dir $(MIGRATION_DIR) postgres $(PG_DSN) down

migration-up: .migration-up ## run up pg migratins

migration-down: .migration-pg-down ## run down pg migratins

migration-create: ## create migration file in pg migration folder
	$(GOOSE_BIN) -dir $(MIGRATION_DIR) create $(name) sql

#======================================#
# BUILD & RUN
#======================================#

run-full: ## run service with db
	@make .compose-up
	@echo "Waiting for DBs to start"
	@sleep 5
	@make .migration-up
	@make .run

run-test:
	go run ./cmd/test

#======================================#
# .PHONY
#======================================#

.PHONY: \
	.custom-bin-deps \
    custom-bin-deps \
	.generate-mocks \
	generate-mocks \
	.compose-up \
	.compose-down \
	.compose-rm \
	.compose-rs \
	compose-up \
	compose-down \
	compose-rm \
	compose-rs \
	migration-create \
	.migration-up \
	.migration-down \
	migration-up \
	migration-down \
	run-full
