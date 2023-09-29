package signer

import (
	"crypto/ecdsa"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
)

func TestPrivateKeySigner(t *testing.T) {
	privateKey, err := crypto.GenerateKey()
	assert.NoError(t, err)

	tests := map[string]struct {
		privateKey *ecdsa.PrivateKey
		chainID    *big.Int
		success    bool
	}{
		"Successful initialization": {
			privateKey: privateKey,
			chainID:    big.NewInt(5),
			success:    true,
		},
		"Failure with Invalid network": {
			privateKey: privateKey,
			success:    false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			_, err := NewPrivateKeySigner(test.privateKey, test.chainID)
			if test.success {
				assert.NoError(t, err)
				return
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestPrivateKeyFromKeystoreSigner(t *testing.T) {
	tests := map[string]struct {
		path     string
		password string
		chainID  *big.Int
		success  bool
	}{
		"Successful initialization": {
			path:     "./mockdata/dummy.key.json",
			password: "test12345",
			chainID:  big.NewInt(5),
			success:  true,
		},
		"Failure with Invalid network": {
			path:     "./mockdata/dummy.key.json",
			password: "test12345",
			success:  false,
		},
		"Failure with Invalid password": {
			path:     "./mockdata/dummy.key.json",
			password: "wrong_password",
			chainID:  big.NewInt(5),
			success:  false,
		},
		"Failure with unknown path": {
			path:     "./mockdata/wrongfile.key.json",
			password: "test12345",
			chainID:  big.NewInt(5),
			success:  false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			_, err := NewPrivateKeyFromKeystoreSigner(test.path, test.password, test.chainID)
			if test.success {
				assert.NoError(t, err)
				return
			} else {
				assert.Error(t, err)
			}
		})
	}
}
