ENV_DEV_FILE := env.dev
ENV_DEV = $(shell cat $(ENV_DEV_FILE))
ENV_TEST_FILE := env.test
ENV_TEST = $(shell cat $(ENV_TEST_FILE))

# App Server
.PHONY: run-dev
run-dev:
	$(ENV_DEV) go run main.go

# Todo: Rebuild Docker Image when Dockerfile or docker-compose.deps.yml is updated
# MySQL
# .PHONY: up-db
# up-db:
# 	docker-compose -f docker-compose.yml up -d

# .PHONY: down-db
# down-db:
# 	docker-compose -f docker-compose.yml down

# .PHONY: up-test-db
# up-test-db:
# 	docker-compose -f docker-compose.test.yml up -d

# .PHONY: down-test-db
# down-test-db:
# 	docker-compose -f docker-compose.test.yml down

# Tools
.PHONY: tools
tools:
	cat tools/tools.go | grep "_" | awk -F'"' '{print $$2}' | xargs -tI % go install %

# Lint, Format
.PHONY: lint
lint: tools
	golangci-lint run ./...

.PHONY: format
format: tools
	golangci-lint run ./... --fix

.PHONY: test
test:
	$(ENV_TEST) go test -v ./...

.PHONY: test-coverage
test-coverage:
	$(ENV_TEST) go test -v ./... -covermode=count

.PHONY: check
check:
	echo "called"
