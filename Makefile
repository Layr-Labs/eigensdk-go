############################# HELP MESSAGE #############################
# Make sure the help command stays first, so that it's printed by default when `make` is called without arguments

GO_LINES_IGNORED_DIRS=contracts
GO_PACKAGES=./chainio/... ./crypto/... ./logging/... \
	./types/... ./utils/... ./signer/... ./cmd/... \
	./signerv2/... ./aws/... ./internal/... ./metrics/... \
	./nodeapi/... ./cmd/... ./services/... ./testutils/...
GO_FOLDERS=$(shell echo ${GO_PACKAGES} | sed -e "s/\.\///g" | sed -e "s/\/\.\.\.//g")

.PHONY: help
help:
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: mocks
mocks: ## generates mocks
	go install go.uber.org/mock/mockgen@v0.4.0
	go generate ./...

.PHONY: tests
tests: ## runs all tests
	go test -race ./... -timeout=1m

.PHONY: tests-cover
tests-cover: ## run all tests with test coverge
	go test -race ./... -coverprofile=coverage.out -covermode=atomic -v -count=1
	go tool cover -html=coverage.out -o coverage.html
	open coverage.html

godoc-port = 6060
.PHONY: godoc
godoc: ## runs godoc server and opens in browser
	@echo "Starting godoc server on port $(godoc-port)..."
	@-godoc -http=:$(godoc-port) & echo $$! > godoc.pid
	@echo "Godoc server started!"
	@echo "Opening localhost page..."
	@sleep 1 # wait for godoc to start
	@open http://localhost:$(godoc-port)/pkg/github.com/Layr-Labs/eigensdk-go/ || xdg-open http://localhost:$(godoc-port)/pkg/github.com/Layr-Labs/eigensdk-go/
	@echo "Press CTRL+C to stop the godoc server..."
	@# The trap command ensures that, when the script exits (e.g., due to CTRL+C), it will kill the godoc process using the PID saved in the godoc.pid file.
	@# The read varname command will keep the make command running (waiting for input) until you press CTRL+C.
	@trap 'kill `cat godoc.pid` && rm -f godoc.pid' EXIT; read varname

.PHONY: fmt
fmt: ## formats all go files
	go fmt ./...
	make format-lines

.PHONY: format-lines
format-lines: ## formats all go files with golines
	go install github.com/segmentio/golines@latest
	golines -w -m 120 --ignore-generated --shorten-comments --ignored-dirs=${GO_LINES_IGNORED_DIRS} ${GO_FOLDERS}

.PHONY: lint
lint: ## runs all linters
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	golangci-lint run ./...


___BINDINGS___: ## 

core_default := "DelegationManager IRewardsCoordinator StrategyManager EigenPod EigenPodManager IStrategy AVSDirectory AllocationManager"
core_location := "./lib/eigenlayer-middleware/lib/eigenlayer-contracts"
core_bindings_location := "../../../../bindings"

middleware_default := "RegistryCoordinator IndexRegistry OperatorStateRetriever StakeRegistry BLSApkRegistry IBLSSignatureChecker ServiceManagerBase IERC20"
middleware_location := "./lib/eigenlayer-middleware"
middleware_bindings_location := "../../bindings"

sdk_default := "MockAvsServiceManager ContractsRegistry"
sdk_location := "."
sdk_bindings_location := "./bindings"

# To generate bindings for specific contracts, run `make core-bindings contracts="DelegationManager IRewardsCoordinator"`
.PHONY: core-bindings ## generates core contracts bindings
core-bindings: ## generates core bindings
	@echo "Starting core bindings generation"
ifneq ($(contracts),)
	@echo "Contracts: $(contracts)"
	cd contracts && ./generate-bindings.sh $(core_location) $(contracts) $(core_bindings_location)
else
	@echo "Contracts: $(core_default)"
	cd contracts && ./generate-bindings.sh $(core_location) $(core_default) $(core_bindings_location)
endif

# To generate bindings for specific contracts, run `make middleware-bindings contracts="RegistryCoordinator"`
.PHONY: middleware-bindings ## generates middleware contracts bindings
middleware-bindings: ## generates middleware bindings
	@echo "Starting middleware bindings generation"
ifneq ($(contracts),)
	@echo "Contracts: $(contracts)"
	cd contracts && ./generate-bindings.sh $(middleware_location) $(contracts) $(middleware_bindings_location)
else
	@echo "Contracts: $(middleware_default)"
	cd contracts && ./generate-bindings.sh $(middleware_location) $(middleware_default) $(middleware_bindings_location)
endif

# To generate bindings for specific contracts, run `make sdk-bindings contracts="MockAvsServiceManager"`
.PHONY: sdk-bindings ## generates sdk contracts bindings
sdk-bindings: ## generates sdk bindings
	@echo "Starting sdk bindings generation"
ifneq ($(contracts),)
	@echo "Contracts: $(contracts)"
	cd contracts && ./generate-bindings.sh $(sdk_location) $(contracts) $(sdk_bindings_location)
else
	@echo "Contracts: $(middleware_default)"
	cd contracts && ./generate-bindings.sh $(sdk_location) $(sdk_default) $(sdk_bindings_location)
endif

.PHONY: eigenpod-bindings
eigenpod-bindings: ## generates contract bindings for eigenpod
	cd chainio/clients/eigenpod && ./generate.sh

.PHONY: bindings
bindings: ## generates all contract bindings
	rm -rf bindings/* && make core-bindings middleware-bindings sdk-bindings eigenpod-bindings


___CONTRACTS___: ## 

.PHONY: deploy-contracts-to-anvil-and-save-state
deploy-contracts-to-anvil-and-save-state: ##
	./contracts/anvil/deploy-contracts-save-anvil-state.sh

.PHONY: start-anvil-with-contracts-deployed
start-anvil-with-contracts-deployed: ## 
	./contracts/anvil/start-anvil-chain-with-el-and-avs-deployed.sh
