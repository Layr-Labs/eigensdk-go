## Signer
This module is used for initializing the signer used to sign smart contract transactions.

There are two types of signer we support
#### Private Key signer
This expects an ECDSA key as argument
```
signer, err := NewPrivateKeySigner(privateKey, chainID)
```

#### Private key from local keystore 
This expects a local path to encrypted ECDSA key json file and password to decrypt it.
```
signer, err := NewPrivateKeyFromKeystoreSigner(path, password, chainID)
```