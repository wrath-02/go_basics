.DEFAULT_GOAL := help
.PHONY: help run all api test test-topic bench fmt vet check

TOPICS := $(sort $(dir $(wildcard [0-9][0-9]-*/main.go)))

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## ' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-12s\033[0m %s\n", $$1, $$2}'

run: ## Run one topic, e.g. make run T=05
	@test -n "$(T)" || (echo "usage: make run T=05" && exit 1)
	@dir=$$(ls -d $(T)-*/ 2>/dev/null | head -1); \
	if [ -z "$$dir" ]; then echo "no topic $(T)"; exit 1; \
	elif [ -f $$dir/main.go ]; then go run ./$$dir; \
	elif [ -d $$dir/cmd/api ]; then go run ./$${dir}cmd/api; \
	else go test -v ./$$dir...; fi

all: ## Run every runnable topic one after another
	@for d in $(TOPICS); do printf "\n==================== %s\n" "$$d"; go run ./$$d </dev/null || exit 1; done

api: ## Start the API server from topic 11 on :8000
	go run ./11-building-an-api/cmd/api

test: ## Run all tests with the race detector
	go test -race ./...

test-topic: ## Run tests for one topic, e.g. make test-topic T=14
	@go test -v ./$$(ls -d $(T)-*/ | head -1)...

bench: ## Run all benchmarks
	go test -run=^$$ -bench=. -benchmem ./...

fmt: ## Format all code
	gofmt -w .

vet: ## Run go vet
	go vet ./...

check: ## fmt check + vet + build + test (same as CI)
	@test -z "$$(gofmt -l .)" || (echo "unformatted files:" && gofmt -l . && exit 1)
	go vet ./...
	go build ./...
	go test -race ./...
