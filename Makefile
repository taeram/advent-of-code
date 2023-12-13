.PHONY: run build update docker docker-build docker-environment generate generate-docs lint ready test test-bench test-integration test-generate tidy chore
.SILENT: run build update docker docker-build docker-environment generate generate-docs lint ready test test-bench test-integration test-generate tidy chore

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
