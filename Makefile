############################# HELP MESSAGE #############################
# Make sure the help command stays first, so that it's printed by default when `make` is called without arguments
.PHONY: help bindings mocks tests tests-cover fmt format-lines lint

GO_LINES_IGNORED_DIRS=contracts
GO_PACKAGES=./chainio/... ./crypto/... ./logging/... \
	./types/... ./utils/... ./signer/... ./cmd/... \
	./signerv2/...
GO_FOLDERS=$(shell echo ${GO_PACKAGES} | sed -e "s/\.\///g" | sed -e "s/\/\.\.\.//g")
help:
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

bindings: ## generates contract bindings
	cd contracts && rm -rf bindings/* && ./generate-bindings.sh 

generate: ## generates mocks and wire files
	go install go.uber.org/mock/mockgen@v0.4.0
	go install github.com/google/wire/cmd/wire@v0.6.0
	go generate ./...

tests: ## runs all tests
	go test -race ./... -timeout=1m

tests-cover: ## run all tests with test coverge
	go test -race ./... -coverprofile=coverage.out -covermode=atomic -v -count=1
	go tool cover -html=coverage.out -o coverage.html
	open coverage.html

godoc-port = 6060
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

fmt: ## formats all go files
	go fmt ./...
	make format-lines

format-lines: ## formats all go files with golines
	go install github.com/segmentio/golines@latest
	golines -w -m 120 --ignore-generated --shorten-comments --ignored-dirs=${GO_LINES_IGNORED_DIRS} ${GO_FOLDERS}

lint: ## runs all linters
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	golangci-lint run ./...

___CONTRACTS___: ## 

deploy-contracts-to-anvil-and-save-state: ## 
	./contracts/anvil/deploy-contracts-save-anvil-state.sh

start-anvil-with-contracts-deployed: ## 
	./contracts/anvil/start-anvil-chain-with-el-and-avs-deployed.sh
