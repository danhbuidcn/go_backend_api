# name app
APP_NAME = server

run:
	go run ./cmd/$(APP_NAME)/

dev:
	docker-compose up && go run ./cmd/$(APP_NAME)

kill:
	docker-compose kill

up:
	docker-compose up -d

down:
	docker-compose down

.PHONY: run
