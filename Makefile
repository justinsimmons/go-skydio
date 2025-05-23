# Copyright 2025 The go-skydio AUTHORS. All rights reserved.
#
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

.DEFAULT: make
.PHONY: make fmt lint test diff help

make:
	go generate ./...
	go fmt ./...

fmt: ## Format the entire repo, shorthand for go fmt.
	go fmt ./...

lint: ## Run linting tool(s).
	go vet ./...
	golangci-lint run ./...

test: ## Run all unit tests.
	go test -v ./...

diff:  ## Outputs the status of the repository, and the volume of per file changes
	$(info The status of the repository, and the volume of per file changes:)
	@git status
	@git diff --stat

help: ## Print help
	@printf "\033[36m%-30s\033[0m %s\n" "(service)"         "Create one of the supported services: $(SERVICES)"
	@printf "\033[36m%-30s\033[0m %s\n" "compose-(service)" "Build the docker compose for the services listed, if supported: $(SERVICES)"
	@printf "\033[36m%-30s\033[0m %s\n" "up-(service)"      "Bring up docker compose for the services listed, if supported: $(SERVICES)"
	@printf "\033[36m%-30s\033[0m %s\n" "down-(service)"    "Bring down docker compose for the services listed, if supported: $(SERVICES)"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

