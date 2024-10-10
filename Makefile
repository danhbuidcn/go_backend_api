GOOSE_DBSTRING= "root:root@tcp(127.0.0.1:3308)/shopdevgo"
GOOSE_MIGRATION_DIR ?= sql/schema

# name app
APP_NAME = server

run:
	go run ./cmd/$(APP_NAME)/

build:
	go build -o bin/migration_app ./cmd/$(APP_NAME)/

dev:
	docker-compose up && go run ./cmd/$(APP_NAME)

kill:
	docker-compose kill

up:
	docker-compose up -d

down:
	docker-compose down

upse:
	@GOOSE_DRIVER=mysql GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) up

downse:
	@GOOSE_DRIVER=mysql GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE MIGRATION_DIR) down

resetse:
	@GOOSE_DRIVER=mysql GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE MIGRATION_DIR) reset

.PHONY: run build upse downse resetse
