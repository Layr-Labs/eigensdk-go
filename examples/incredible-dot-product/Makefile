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

deploy-all: deploy-eigenlayer deploy-avs

bindings: ## generates contract bindings
	cd contracts && ./generate-go-bindings.sh
