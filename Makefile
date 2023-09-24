.DEFAULT_GOAL := all

.PHONY: all
all: install lint test

.PHONY: install
install:
	@go mod vendor

.PHONY: lint
lint:
	@golangci-lint run

.PHONY: format
format:
	@golangci-lint run --fix

.PHONY: test
test:
	@CGO_ENABLED=0 go test ./cmd/... ./internal/...
