![Unit Tests](https://github.com/Layr-Labs/eigensdk-go/actions/workflows/unit-tests.yml/badge.svg)
![Linter](https://github.com/Layr-Labs/eigensdk-go/actions/workflows/golangci-lint.yml/badge.svg)
![Go Coverage](https://github.com/Layr-Labs/eigensdk-go/wiki/coverage.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/Layr-Labs/eigensdk-go)](https://goreportcard.com/report/github.com/Layr-Labs/eigensdk-go)




## EigenSDK
This SDK provides a set of primitive Go modules for developing AVSs on EigenLayer.

## Installation
```
go get github.com/Layr-Labs/eigensdk-go
```

## Modules
We support following modules right now. 
> **_NOTE:_** All modules are in active development and interfaces might change. 
* [Logging](./logging/README.md)
* [Signer](./signer/README.md)
* [ChainIO](./chainio/README.md)
* [Services](./services/README.md)

## Development
Clone the repo
```
git clone https://github.com/Layr-Labs/eigensdk-go.git
```
Initialize git submodules
```
git submodule update --init
```

Follow the [contribution guidelines](CONTRIBUTING.md) to contribute to eigensdk-go

## Branches

For consistency with [eigenlayer-middleware](https://github.com/Layr-Labs/eigenlayer-middleware) and [eigenlayer-contracts](https://github.com/Layr-Labs/eigenlayer-contracts) repos, we no longer use the `master` branch and instead use `dev` as the default branch, which will track as closely as possible the `dev` branch of eigenlayer-middleware (which in turn tracks the `dev` branch of eigenlayer-contracts). This convention will also be followed for other important branches. For eg, the m2-mainnet branch of this repo will track the m2-mainnet branch of eigenlayer-middleware (which tracks the unfortunately named mainnet branch of eigenlayer-contracts), and same with the testnet-holesky branch.

## Security Bugs
Please report security vulnerabilities to security@eigenlabs.org. Do NOT report security bugs via Github Issues.

## Disclaimer
🚧 EigenSDK-go is under active development and has not been audited. EigenSDK-go is rapidly being upgraded, features may be added, removed or otherwise improved or modified and interfaces will have breaking changes. EigenSDK-go should be used only for testing purposes and not in production. EigenSDK-go is provided "as is" and Eigen Labs, Inc. does not guarantee its functionality or provide support for its use in production. 🚧
