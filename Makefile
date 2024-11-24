include .envrc
VOLUME_NAME = go-nakama-apps_db-data
MIGRATIONS_PATH = ./cmd/migrate/migrations
IMAGE_NAME = kandlagifari/nakama-app
IMAGE_TAG = latest

.PHONY: compose-up
compose-up:
	@docker compose up --build -d

.PHONY: compose-down
compose-down:
	@docker compose down
	@docker volume rm ${VOLUME_NAME}

.PHONY: migrate-create
migrate-create:
	@migrate create -seq -ext sql -dir $(MIGRATIONS_PATH) $(filter-out $@,$(MAKECMDGOALS))

.PHONY: migrate-up
migrate-up:
	@migrate -path=$(MIGRATIONS_PATH) -database=$(DB_ADDR) up

.PHONY: migrate-down
migrate-down:
	@migrate -path=$(MIGRATIONS_PATH) -database=$(DB_ADDR) down $(filter-out $@,$(MAKECMDGOALS))

.PHONY: generate-docs
generate-docs:
	@swag init -g ./api/main.go -d cmd,internal && swag fmt

.PHONY: build
build:
	@docker build -t $(IMAGE_NAME):$(IMAGE_TAG) .

.PHONY: push
push:
	@docker login -u $(DOCKER_USERNAME) -p $(DOCKER_PASSWORD)
	@docker push $(IMAGE_NAME):$(IMAGE_TAG)

.PHONY: clean
clean:
	@docker rm -f nakama-api

# .PHONY: seed
# seed: 
# 	@go run cmd/migrate/seed/main.go
