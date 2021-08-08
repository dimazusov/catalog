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

generate:
	protoc --proto_path=internal/server/grpc --go_out=internal/server/grpc/pb --go-grpc_out=internal/server/grpc/pb internal/server/grpc/*.proto

install-lint-deps:
	(which golangci-lint > /dev/null) || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.33.0

lint: install-lint-deps
	golangci-lint run ./...

install-migrate-deps:
	go get -u github.com/pressly/goose/cmd/goose

migrate: install-migrate-deps
	cd migrations && goose postgres "user=db password=db dbname=db sslmode=disable" up

.PHONY: build test lint
.PHONY: build run build-img run-img version test lint
