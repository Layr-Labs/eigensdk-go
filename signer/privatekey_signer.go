package signer

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// Deprecated: Use SignerV2 instead
type PrivateKeySigner struct {
	txOpts *bind.TransactOpts
}

var _ Signer = (*PrivateKeySigner)(nil)

func NewPrivateKeySigner(privateKey *ecdsa.PrivateKey, chainID *big.Int) (Signer, error) {
	// TODO (madhur): chainId should not be hardcoded
	txOpts, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, err
	}
	return &PrivateKeySigner{
		txOpts: txOpts,
	}, nil
}

func NewPrivateKeyFromKeystoreSigner(path string, password string, chainID *big.Int) (Signer, error) {
	privateKey, err := getECDSAPrivateKey(path, password)
	if err != nil {
		return nil, err
	}

	// TODO (madhur): chainId should not be hardcoded
	txOpts, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, err
	}
	return &PrivateKeySigner{
		txOpts: txOpts,
	}, nil
}

func (p *PrivateKeySigner) GetTxOpts() *bind.TransactOpts {
	return p.txOpts
}

func (p *PrivateKeySigner) SendToExternal(ctx context.Context, tx *types.Transaction) (common.Hash, error) {
	return common.Hash{}, errors.New("this signer does not support external signing")
}

func getECDSAPrivateKey(keyStoreFile string, password string) (*ecdsa.PrivateKey, error) {
	keyStoreContents, err := os.ReadFile(keyStoreFile)
	if err != nil {
		return nil, err
	}

	sk, err := keystore.DecryptKey(keyStoreContents, password)
	if err != nil {
		return nil, err
	}

	return sk.PrivateKey, nil
}
