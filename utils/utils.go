package utils

import (
	"crypto/ecdsa"
	"errors"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
)

func EcdsaPrivateKeyToAddress(privateKey *ecdsa.PrivateKey) (gethcommon.Address, error) {
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return gethcommon.Address{}, errors.New("ErrCannotGetECDSAPubKey")
	}
	return crypto.PubkeyToAddress(*publicKeyECDSA), nil
}

func RoundUpDivideBig(a, b *big.Int) *big.Int {
	one := new(big.Int).SetUint64(1)
	res := new(big.Int)
	a.Add(a, b)
	a.Sub(a, one)
	res.Div(a, b)
	return res
}
