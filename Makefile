BRANCH = $(shell git rev-parse --abbrev-ref HEAD)
DATE = $(shell git show --pretty=format:%cs --no-patch)
HASH = $(shell git show --pretty=format:%h --no-patch)
TAG = $(shell git for-each-ref refs/tags --sort=-taggerdate --format='%(refname:short)' --count=1)
GIT_INFO = $(BRANCH) - $(DATE) - $(HASH)

RED = \x1b[31;49;3;1m
GREEN = \x1b[32;49;3;1m
YELLOW = \x1b[33;49;3;1m
BLUE = \x1b[34;49;3;1m
MAGENTA = \x1b[35;49;3;1m
CYAN = \x1b[36;49;3;1m
WHITE = \x1b[49;3;1m

help: ## Display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
.PHONY: help

build-dev: swag ## Build the dev binary file
	@echo "${GREEN}build dev binary file"
	@go build -tags dev -v -a -ldflags="-X 'among/config/config.configTarget=dev'" -o bin/among cmd/app/main.go
.PHONY: build-dev

build-stage: swag ## Build the dev binary file
	@echo "${GREEN}build dev binary file"
	@go build -tags stage -a -ldflags="-X 'main.GitInfo=$(GIT_INFO)' -X 'main.Tag=$(TAG)'" -o bin/among cmd/app/main.go
.PHONY: build-stage

build-prod: swag ## Build the dev binary file
	@echo "${GREEN}build dev binary file"
	@go build -tags prod -a -ldflags="-X 'main.GitInfo=$(GIT_INFO)' -X 'main.Tag=$(TAG)'" -o bin/among cmd/app/main.go
.PHONY: build-prod

swag: ### swag init
	@echo "${RED}swag init"
	@swag init -g cmd/app/main.go
.PHONY: swag

run: swag ## Run the swag binary file
	@echo "${GREEN}run the main.go file"
	@go mod tidy && go mod download && go run ./cmd/app/main.go
.PHONY: run

mock: ### run mockery
	@mockery --all --keeptree --dir $(PWD)/internal --output $(PWD)/mocks --disable-version-string
.PHONY: mock

compose-up: ### Run docker-compose
	docker-compose up --build -d postgres rabbitmq && docker-compose logs -f
.PHONY: compose-up

compose-up-integration-test: ### Run docker-compose with integration test
	docker-compose up --build --abort-on-container-exit --exit-code-from integration
.PHONY: compose-up-integration-test

compose-down: ### Down docker-compose
	docker-compose down --remove-orphans
.PHONY: compose-down

docker-rm-volume: ### remove docker volume
	docker volume rm go-clean-template_pg-data
.PHONY: docker-rm-volume

linter-golangci: ### check by golangci linter
	golangci-lint run
.PHONY: linter-golangci

linter-hadolint: ### check by hadolint linter
	git ls-files --exclude='Dockerfile*' --ignored | xargs hadolint
.PHONY: linter-hadolint

linter-dotenv: ### check by dotenv linter
	dotenv-linter
.PHONY: linter-dotenv

test: ### run test
	go test -v -cover -race ./internal/...
.PHONY: test

integration-test: ### run integration-test
	go clean -testcache && go test -v ./integration-test/...
.PHONY: integration-test

migrate-create:  ### create new migration
	migrate create -ext sql -dir migrations 'migrate_name'
.PHONY: migrate-create

migrate-up: ### migration up
	migrate -path migrations -database '$(PG_URL)?sslmode=disable' up
.PHONY: migrate-up
