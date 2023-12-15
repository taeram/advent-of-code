.PHONY: lint ready run test tidy update workflow
.SILENT: lint ready run test tidy update workflow

SHELL=/bin/bash

lint:
	golangci-lint run --fix ./$(day)/ ./internal/...

ready: tidy lint test run

run:
	go run $(day)/main.go $(day)/input.txt

test:
	gotest ./$(day)/

tidy:
	go mod tidy

update:
	rm -f go.sum
	go get -u ./...
	go mod tidy

workflow:
	act --job test
