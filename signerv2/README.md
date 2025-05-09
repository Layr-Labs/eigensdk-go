# Signer v2

Signerv2 is a module for signing messages. It provides a simple and unified way to produce cryptographic signatures.
Signers instantiated from this module is required to create some SDK transaction managers (see [`NewPrivateKeyWallet`](../chainio/clients/wallet/privatekey_wallet.go) and [`NewSimpleTxManager`](../chainio/txmgr/simple.go)/[`NewGeometricTxnManager`](../chainio/txmgr/geometric/geometric.go)).

## Features

- Sign messages using raw private keys
- Sign messages using encrypted keystores
- Sign messages using a remote signer (web3 or KMS)

### Comparison to Old Signer

In comparison to the old signer, Signerv2 offers:

- New signing mechanisms
- A simplified API for easier extension

### Using SignerFromConfig

SignerV2 introduces `SignerFromConfig`

The `SignerFromConfig` function allows you to create a signer function based on a configuration.
This configuration specifies whether to use a private key signer, a local keystore signer, or a remote web3 signer.

```go
package main

import (
    "github.com/Layr-Labs/eigensdk-go/signerv2"
)

func main() {
    config := signerv2.Config{
        // ...initialize your configuration...
    }
    chainID := // Set your chain ID
    signerFn, signerAddr, err := signerv2.SignerFromConfig(config, chainID)
    if err != nil {
        // Handle error
        return
    }
    // Use signerFn and signerAddr as needed
}
```

Internally, `SignerFromConfig` calls different signer functions depending on the config it receives: `PrivateKeySignerFn`, `KeyStoreSignerFn`, or `Web3SignerFn`.
Those functions are also available to users.

### KMSSignerFn

This module includes support for signing messages using a Key Management Service (KMS) key.  Use `KMSSignerFn` to create a signer for KMS-managed keys.

## Upgrade from Signer (v1)

`NewPrivateKeySigner` and `NewPrivateKeyFromKeystoreSigner` functions should be upgraded to use the new `SignerFromConfig` and a `Config` with a `PrivateKey`, or `KeystorePath` and `Password`, respectively.

The functionality given by the `Signer` interface and `BasicSigner` type was redesigned into the [`wallet`](../chainio/clients/wallet) and [`txmgr`](../chainio/txmgr) modules.
After generating a `SignerFn` as specified in ["UsingSignerFromConfig"](#using-signerfromconfig), you can generate a transaction manager via `NewPrivateKeyWallet` and `NewSimpleTxManager` (or `NewGeometricTxnManager` for geometric gas pricing)
