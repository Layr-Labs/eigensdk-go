![Unit Tests](https://github.com/Layr-Labs/eigensdk-go/actions/workflows/unit-tests.yml/badge.svg)
![Linter](https://github.com/Layr-Labs/eigensdk-go/actions/workflows/golangci-lint.yml/badge.svg)
![Go Coverage](https://github.com/Layr-Labs/eigensdk-go/wiki/coverage.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/Layr-Labs/eigensdk-go)](https://goreportcard.com/report/github.com/Layr-Labs/eigensdk-go)

<p align="center"><b>
🚧 Under active development. EIGENSDK-GO is rapidly being upgraded, features are being added, interfaces will have breaking changes 🚧
</b><p>

**Do not use it in Production, testnet only.**

## EigenSDK
This SDK provides a set of primitive Go modules for developing AVSs used in EigenLayer

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

## Security Bugs
Please report security vulnerabilities to security@eigenlabs.org. Do NOT report security bugs via Github Issues.