.PHONY: test* run build

PACKAGE_NAME := github.com/akhilesharora/rolling-hash

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build: clean ## Compile app
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/app cmd/main.go

run: build ## Run app
	go run cmd/main.go

test: ## Run all tests
	go test ./... -v count=1

test-race:
	go test -race -short ./...

test-clean-testcache:
	go clean -testcache && go test -v ./...

coverage: clean ## Run tests and generate coverage files per package
	mkdir .coverage 2> /dev/null || true
	rm -rf .coverage/*.out || true
	go test -race ./... -coverprofile=coverage.out -covermode=atomic

clean:
	rm -rf .coverage/ build/
