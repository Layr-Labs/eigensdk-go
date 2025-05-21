![Unit Tests](https://github.com/Layr-Labs/eigensdk-go/actions/workflows/unit-tests.yml/badge.svg)
![Linter](https://github.com/Layr-Labs/eigensdk-go/actions/workflows/golangci-lint.yml/badge.svg)
![Go Coverage](https://github.com/Layr-Labs/eigensdk-go/wiki/coverage.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/Layr-Labs/eigensdk-go)](https://goreportcard.com/report/github.com/Layr-Labs/eigensdk-go)

## EigenSDK

This SDK provides a set of primitive Go modules for developing AVSs on EigenLayer.

## Installation

``` bash
go get github.com/Layr-Labs/eigensdk-go
```

## Modules

We support following modules right now.
> **_NOTE:_** All modules are in active development and interfaces might change.

* [Logging](./logging)
* [ECDSA Signer](./signerv2)
* [BLS Signer](./signer)
* [ChainIO](./chainio)
* [Services](./services)

## AVS use examples

This SDK has three AVS use examples:

* [Incredible Squaring](https://github.com/Layr-Labs/eigensdk-go/tree/v2-dev-2/examples/incredible-squaring)
* [Incredible Dot Product](https://github.com/Layr-Labs/eigensdk-go/tree/v2-dev-2/examples/incredible-dot-product)
* [Awesome Vault Service](https://github.com/Layr-Labs/eigensdk-go/tree/v2-dev-2/examples/awesome-vault-service)

These are examples of the code needed to create an AVS from scratch. All include the contracts needed to execute them, with further instructions in their respective readme files.

### Business logic in entities

* Aggregator: The aggregator does not have much business logic, since the task processor implementation lies in the Indexing Task Processor implementation, which can be seen as the default one. If wanted to create your task processor, you can base it on the ITP implementation, and change what you need.
* Challenger: The challenger business logic lies in Task response validation. To validate the response, the challenger first calculates the response with the same function as the operator and then compares it with the received response, raising a challenge if they differ.
* Operator: The operator responds to tasks using the ResponseCalculator struct corresponding to the task to complete, which has a ComputeResponse method (satisfying the ResponseCalculator interface).
* Task spammer: The task spammer logic lies in the sequence that generates the numbers pulled by the spammer at the SDK level.

## Development

Clone the repo

``` bash
git clone https://github.com/Layr-Labs/eigensdk-go.git
```

Initialize git submodules

``` bash
git submodule update --init --recursive
```

Follow the [contribution guidelines](CONTRIBUTING.md) to contribute to eigensdk-go

## Branches

For consistency with [eigenlayer-middleware](https://github.com/Layr-Labs/eigenlayer-middleware) and [eigenlayer-contracts](https://github.com/Layr-Labs/eigenlayer-contracts) repos, we no longer use the `master` branch and instead use `dev` as the default branch, which will track as closely as possible the `dev` branch of eigenlayer-middleware (which in turn tracks the `dev` branch of eigenlayer-contracts). This convention will also be followed for other important branches. For eg, the m2-mainnet branch of this repo will track the m2-mainnet branch of eigenlayer-middleware (which tracks the unfortunately named mainnet branch of eigenlayer-contracts), and same with the testnet-holesky branch.

## Security Bugs

Please report security vulnerabilities to security@eigenlabs.org. Do NOT report security bugs via Github Issues.

## Disclaimer

🚧 EigenSDK-go is under active development and has not been audited. EigenSDK-go is rapidly being upgraded, features may be added, removed or otherwise improved or modified and interfaces will have breaking changes. EigenSDK-go should be used only for testing purposes and not in production. EigenSDK-go is provided "as is" and Eigen Labs, Inc. does not guarantee its functionality or provide support for its use in production. 🚧
