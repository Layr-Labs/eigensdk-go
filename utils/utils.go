package utils

import (
	"crypto/ecdsa"
	"errors"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func EcdsaPrivateKeyToAddress(privateKey *ecdsa.PrivateKey) (common.Address, error) {
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return common.Address{}, errors.New("ErrCannotGetECDSAPubKey")
	}
	return crypto.PubkeyToAddress(*publicKeyECDSA), nil
}

func Add0x(address string) string {
	if strings.HasPrefix(address, "0x") {
		return address
	}
	return "0x" + address
}

func Trim0x(address string) string {
	return strings.TrimPrefix(address, "0x")
}
