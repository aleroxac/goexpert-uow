ifeq ($(wildcard .env),)
    $(shell cp .env.example .env)
endif

include .env
export


## ---------- UTILS
.PHONY: help
help: ## Show this menu
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: clean
clean: ## Clean all temp files
	@rm -rf tmp/



## ---------- INSTALL
.PHONY: install-sqlc
install-sqlc: ## install sqlc
	@if [ ! -f ~/.go/bin/sqlc ]; then \
		go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest; \
	else \
		echo "Great, you already have [sqlc] installed"; \
	fi

.PHONY: install
install: install-sqlc ## install all requirements



## ---------- COMPOSE
.PHONY: compose_up
compose_up: ## run docker-compose up
	@docker-compose up -d

.PHONY: compose.down
compose_down: ## run docker-compose down
	@docker-compose down



## ---------- MAIN
.PHONY: generate
generate: ## run sqls generate
	@sqlc generate

.PHONY: test
test: ## run unit tests
	@go test -v ./...
