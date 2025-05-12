.PHONY: help
help:
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

AGGREGATOR_ECDSA_PRIV_KEY=0x2a871d0798f97d79848a013d4936a73bf4cc922c825d33c1cf7073dff6d409c6
CHALLENGER_ECDSA_PRIV_KEY=0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80
TASK_SPAMMER_PRIVATE_KEY=0x4bbbf85ce3377467afe5d46f804f221813b2bb87f24d81f60f1fcdbf7cbf4356
CHAINID=31337
DEPLOYMENT_FILE=contracts/script/deployments/incredible-squaring/${CHAINID}.json
CORE_DEPLOYMENT_FILE=contracts/script/deployments/core/${CHAINID}.json

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
	go run aggregator/main.go --config config-files/aggregator.yaml \
		--credible-squaring-deployment ${DEPLOYMENT_FILE} \
		--core-deployment ${CORE_DEPLOYMENT_FILE} \
		--ecdsa-private-key ${AGGREGATOR_ECDSA_PRIV_KEY} \
		2>&1 | zap-pretty

start-operator: ## 
	go run ./operator --config config-files/operator.anvil.yaml \
		2>&1 | zap-pretty

start-challenger: ## 
	go run challenger/main.go --config config-files/challenger.yaml \
		--credible-squaring-deployment ${DEPLOYMENT_FILE} \
		--core-deployment ${CORE_DEPLOYMENT_FILE} \
		--ecdsa-private-key ${CHALLENGER_ECDSA_PRIV_KEY} \
		2>&1 | zap-pretty

start-task-spammer: ## 
	go run task-spammer/main.go --config config-files/task-spammer.yaml \
		--credible-squaring-deployment ${DEPLOYMENT_FILE} \
		--core-deployment ${CORE_DEPLOYMENT_FILE} \
		--ecdsa-private-key ${TASK_SPAMMER_PRIVATE_KEY} \
		2>&1 | zap-pretty

.PHONY: lint
lint: ## runs all linters
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	golangci-lint run ./...
