package signerv2

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"math/big"

	sdkEcdsa "github.com/Layr-Labs/eigensdk-go/crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

type SignerFn func(ctx context.Context, address common.Address) (bind.SignerFn, error)

func PrivateKeySignerFn(privateKey *ecdsa.PrivateKey, chainID *big.Int) (bind.SignerFn, error) {
	from := crypto.PubkeyToAddress(privateKey.PublicKey)
	signer := types.LatestSignerForChainID(chainID)

	return func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
		if address != from {
			return nil, bind.ErrNotAuthorized
		}
		signature, err := crypto.Sign(signer.Hash(tx).Bytes(), privateKey)
		if err != nil {
			return nil, err
		}
		return tx.WithSignature(signer, signature)
	}, nil
}

func KeyStoreSignerFn(path string, password string, chainID *big.Int) (bind.SignerFn, error) {
	privateKey, err := sdkEcdsa.ReadKey(path, password)
	if err != nil {
		return nil, err
	}
	return PrivateKeySignerFn(privateKey, chainID)
}

// RemoteSignerFn creates a signer function that uses a remote signer
// It should expose `eth_SignTransaction` endpoint which return rlp
// encoded signed tx
func RemoteSignerFn(remoteSignerUrl string) (bind.SignerFn, error) {
	client := NewRemoteSignerClient(remoteSignerUrl)

	return func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
		return client.SignTransaction(address, tx)
	}, nil
}

func SignerFromConfig(c Config, chainID *big.Int) (SignerFn, common.Address, error) {
	var signer SignerFn
	var senderAddress common.Address
	var err error
	if c.IsPrivateKeySigner() {
		senderAddress = crypto.PubkeyToAddress(c.PrivateKey.PublicKey)
		signer = func(ctx context.Context, address common.Address) (bind.SignerFn, error) {
			return PrivateKeySignerFn(c.PrivateKey, chainID)
		}
	} else if c.IsLocalKeystoreSigner() {
		senderAddress, err = sdkEcdsa.GetAddressFromKeyStoreFile(c.KeystorePath)
		if err != nil {
			return nil, common.Address{}, err
		}
		signer = func(ctx context.Context, address common.Address) (bind.SignerFn, error) {
			return KeyStoreSignerFn(c.KeystorePath, c.Password, chainID)
		}
	} else if c.IsRemoteSigner() {
		senderAddress = common.HexToAddress(c.Address)
		signer = func(ctx context.Context, address common.Address) (bind.SignerFn, error) {
			return RemoteSignerFn(c.Endpoint)
		}
	} else {
		return nil, common.Address{}, errors.New("no signer found")
	}
	return signer, senderAddress, nil
}
