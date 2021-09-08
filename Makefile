SHELL := /bin/bash

.PHONY: run
run:
	@air

.PHONY: lint
lint:
	@golangci-lint run --fix ./...

.PHONY:
test:
	@go test -shuffle=on -coverprofile=./tmp/cover.out ./...
	@go tool cover -html=./tmp/cover.out -o=./tmp/cover.html
