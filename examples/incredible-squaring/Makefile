.PHONY: help
help:
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

___CONTRACTS___: ## 

build-contracts: ## builds all contracts
	cd contracts && forge build

deploy-eigenlayer: ## Deploy eigenlayer
	./tests/anvil/deploy-eigenlayer.sh

deploy-avs: ## Deploy avs
	./tests/anvil/deploy-avs.sh

uam-permissions:
	./tests/anvil/uam-permissions.sh

create-quorum:
	./tests/anvil/create-quorum.sh

deploy-all: deploy-eigenlayer deploy-avs uam-permissions create-quorum

bindings: ## generates contract bindings
	cd contracts && ./generate-go-bindings.sh

____OFFCHAIN_SOFTWARE___: ## 
start-aggregator: ## 
	go run aggregator/main.go 2>&1 | zap-pretty

start-operator: ## 
	go run ./operator 2>&1 | zap-pretty

start-challenger: ## 
	go run challenger/main.go 2>&1 | zap-pretty

start-task-spammer: ## 
	go run task-spammer/main.go 2>&1 | zap-pretty

.PHONY: lint
lint: ## runs all linters
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	golangci-lint run ./...
