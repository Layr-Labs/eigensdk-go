package testutils

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func NewEcdsaSkAndAddress() (*ecdsa.PrivateKey, common.Address, error) {
	ecdsaSk, err := crypto.GenerateKey()
	if err != nil {
		return nil, common.Address{}, err
	}
	pk := ecdsaSk.Public()
	address := crypto.PubkeyToAddress(*pk.(*ecdsa.PublicKey))
	return ecdsaSk, address, nil
}

func ZeroAddress() *common.Address {
	return &common.Address{}
}
