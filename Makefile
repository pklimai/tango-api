include architect.mk

#======================================#
# VARIABLES
#======================================#

## GENERAL
SERVICE_NAME = "tango-api"
MIGRATION_DIR = $(CURDIR)/migration

## BIN
GOOSE_BIN = $(LOCAL_BIN)/goose

## PG
PG_MIGRATION_DIR = $(MIGRATION_DIR)/postgresql
PG_DSN="host=localhost port=1331 user=tango-api-user password=1234 dbname=tango-api sslmode=disable"

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
	PATH="$(LOCAL_BIN):$$PATH" go generate -x ./...

generate-mocks: .generate-mocks ## geenrate mocks for tests

# Generating mocks...
# PATH="/home/zigal0/Dev/Go/gitlab/zigal0/go-reference-project/bin:$PATH" go generate -x run=minimock ./...
# sh -c rm -rf mocks && mkdir -p mocks
# minimock -i * -o ./mocks -s _mock.go -g
# minimock: ./mocks/employee_repo_mock.go
# minimock: ./mocks/task_repo_mock.go

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
	make compose-rm
	make compose-up

compose-up: .compose-up ## start docker containers

compose-down: .compose-down ## stop docker containers

compose-rm: .compose-rm ## stop and remove docker containers

compose-rs: .compose-rs ## restart docker containers

#======================================#
# MIGRATIONS
#======================================#

## PG

.migration-pg-up: 
	$(GOOSE_BIN) -dir $(PG_MIGRATION_DIR) postgres $(PG_DSN) up

.migration-pg-down:
	$(GOOSE_BIN) -dir $(PG_MIGRATION_DIR) postgres $(PG_DSN) down

migration-pg-up: .migration-pg-up ## run up pg migratins

migration-pg-down: .migration-pg-down ## run down pg migratins

migration-pg-create: ## create migration file in pg migration folder
	$(GOOSE_BIN) -dir $(PG_MIGRATION_DIR) create $(name) sql

## General

migration-up: ## run up all migratins 
	@make .migration-pg-up 

migration-down: ## run down all migratins
	@make .migration-pg-down 

#======================================#
# BUILD & RUN
#======================================#

run-full: ## run service with db
	@make compose-up
	@echo "Waiting for DBs to start"
	@sleep 5
	@make migration-pg-up
	@make .run

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
	migration-pg-create \
	.migration-pg-up \
	.migration-pg-down \
	migration-up \
	migration-down \
	run-full
