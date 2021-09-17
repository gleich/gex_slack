##########
# Building
##########

build-docker-prod:
	docker build -f docker/Dockerfile -t mattgleich/gex_slack:latest .
build-docker-dev:
	docker build -f docker/dev.Dockerfile -t mattgleich/gex_slack:dev .
build-docker-test:
	docker build -f docker/test.Dockerfile -t mattgleich/gex_slack:test .
build-docker-lint:
	docker build -f docker/lint.Dockerfile -t mattgleich/gex_slack:lint .
build-go:
	go get -v -t -d ./...
	go build -v .

#########
# Linting
#########

lint-golangci:
	golangci-lint run
lint-gomod:
	go mod tidy
	git diff --exit-code go.mod
	git diff --exit-code go.sum
lint-goreleaser:
	goreleaser check
lint-hadolint:
	hadolint docker/Dockerfile
	hadolint docker/dev.Dockerfile
	hadolint docker/test.Dockerfile
	hadolint docker/lint.Dockerfile
lint-in-docker: build-docker-lint
	docker run mattgleich/gex_slack:lint

#########
# Testing
#########

test-go:
	go get -v -t -d ./...
	go test -v ./...
test-in-docker: build-docker-test
	docker run mattgleich/gex_slack:test

##########
# Grouping
##########

# Testing
local-test: test-go
docker-test: test-in-docker
# Linting
local-lint: lint-golangci lint-goreleaser lint-hadolint lint-gomod
docker-lint: lint-in-docker
# Build
local-build: build-docker-prod build-docker-dev build-docker-dev-lint
# Development
dev:
	docker compose up
