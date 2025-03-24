package signerv2

import (
	"crypto/ecdsa"
	sdkEcdsa "github.com/Layr-Labs/eigensdk-go/crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func PrivateKeyMessageSignerFn(privateKey *ecdsa.PrivateKey) func([]byte) ([]byte, error) {
	return func(message []byte) ([]byte, error) {
		return crypto.Sign(message, privateKey)
	}
}

func KeyStoreMessageSignerFn(path, password string) (func([]byte) ([]byte, error), error) {
	privateKey, err := sdkEcdsa.ReadKey(path, password)
	if err != nil {
		return nil, err
	}
	return PrivateKeyMessageSignerFn(privateKey), nil
}

func Web3MessageSignerFn(sender common.Address, remoteSignerUrl string) (func([]byte) ([]byte, error), error) {
	client := NewWeb3SignerClient(remoteSignerUrl)

	return func(message []byte) ([]byte, error) {
		return client.SignMessage(sender, message)
	}, nil
}
