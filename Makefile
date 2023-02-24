ENV_DEV_FILE := env.dev
ENV_DEV = $(shell cat $(ENV_DEV_FILE))
ENV_TEST_FILE := env.test
ENV_TEST = $(shell cat $(ENV_TEST_FILE))
ENV_GITLABCI_FILE := env.gitlabci
ENV_GITLABCI = $(shell cat $(ENV_GITLABCI_FILE))

# App Server
.PHONY: run-dev
run-dev:
	$(ENV_DEV) docker-compose up

.PHONY: run-for-test
run-for-test:
	$(ENV_TEST) docker-compose up

# Todo: Rebuild Docker Image when Dockerfile or docker-compose.dev-db.yml is updated
# MySQL
.PHONY: up-db
up-db:
	docker-compose -f docker-compose.dev-db.yml up -d

.PHONY: init-db
init-db:
	docker-compose exec template-backend go run mysql/init_db.go

.PHONY: down-db
down-db:
	docker-compose -f docker-compose.deps.yml down

.PHONY: up-test-db
up-test-db:
	docker-compose -f docker-compose.test.yml up -d

.PHONY: down-test-db
down-test-db:
	docker-compose -f docker-compose.test.yml down

# Tools
.PHONY: tools
tools:
	cat tool/tools.go | grep "_" | awk -F'"' '{print $$2}' | xargs -tI % go install %

# Lint, Format
.PHONY: lint
lint: tools
	golangci-lint run ./... --timeout=5m

.PHONY: format
format: tools
	golangci-lint run ./... --fix

.PHONY: test
test:
	docker-compose exec template-backend go test -v ./...

.PHONY: test-without-db
test-without-db:
	docker-compose exec template-backend go test `go list ./... | grep -v gitlab.com/soy-app/stock-api/integration_tests`

.PHONY: test-gitlab
test-gitlab:
	$(ENV_GITLABCI) go test -v ./...

.PHONY: test-coverage
test-coverage:
	docker-compose exec template-backend go test -v ./... -covermode=count

.PHONY: check
check:
	echo "called"
