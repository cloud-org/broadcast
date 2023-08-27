#/bin/bash
PWD ?= $(shell pwd)
export PATH := $(PWD)/bin:$(PATH)

default: help
help:  ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z0-9_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
.PHONY: help

build: build-main build-notify ## Build main and notify
.PHONY: build

build-main:
	@go build -o bin/main main.go
.PHONY: build-main

build-notify:
	@go build -o bin/notify notify/main.go
.PHONY: build-notify

e2e: ## Run e2e test
	@cd test/e2e && ginkgo -r
.PHONY: e2e

.PHONY: fmt
fmt: ## Format all go codes
	./scripts/goimports-reviser.sh
