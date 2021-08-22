BIN := "./bin/calendar"
DOCKER_IMG="calendar:develop"

GIT_HASH := $(shell git log --format="%h" -n 1)
LDFLAGS := -X main.release="develop" -X main.buildDate=$(shell date -u +%Y-%m-%dT%H:%M:%S) -X main.gitHash=$(GIT_HASH)

run: build
	$(BIN) -config ./configs/config.yaml

up:
	docker-compose -f deployments/docker-compose.yaml up -d

version: build
	$(BIN) version

test:
	go test -race ./internal/...

install-lint-deps:
	(which golangci-lint > /dev/null) || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.33.0

lint: install-lint-deps
	golangci-lint run ./...

migrate:
	docker-compose -f deployments/docker-compose.yaml exec -it

swagger-init:
	go get -u github.com/swaggo/swag/cmd/swag

swagger: swagger-init
 	swag init -g ./internal/server/http/router.go -o api