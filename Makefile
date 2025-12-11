DAYS?=9
DAY?=$(shell printf "%02g" $(DAYS))

.PHONY: help build run scaffold clean lint

help:
	@echo 'Usage: make [target] ...'
	@echo
	@echo 'Targets:'
	@grep -F -h "## " $(MAKEFILE_LIST) | grep -v grep  \
	| sed 's/^\(.*\):[^#]*##*\(.*\)/\x1b[36m\1\x1b[0m:\2/' \
	| column -t -s ':'

build: ## Build all days and place binaries under ./bin
	@mkdir -p ./bin
	@bash -c 'for d in $$(seq -f "%02g" 1 $(DAYS)); do go build -o bin/$$d $$d/main.go; done'

run: ## Timed run of all days
	@bash -c 'for d in $$(seq -f "%02g" 1 9); do echo -e "+--------+\n| Day $$d |\n+--------+"; time ./bin/$$d; echo; done'

scaffold: ## Scaffold a day, prefix it with DAYS=xx
	@mkdir -p ./$(DAY)
	@touch ./inputs/$(DAY).ex
	@bash -c 'test -f ./$(DAY)/main.go || DAY=$(DAY) envsubst <./day.tpl.go >./$(DAY)/main.go'

clean: ## Nuke the ./bin dir from orbit
	@/sbin/rm -rf ./bin

lint: ## Run linters and checkers
	go fmt ./...
	golangci-lint run ./...
