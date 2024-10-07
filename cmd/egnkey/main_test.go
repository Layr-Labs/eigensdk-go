package main

import (
	"bytes"
	"os"
	"testing"

	sdkEcdsa "github.com/Layr-Labs/eigensdk-go/crypto/ecdsa"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
)

func TestGenerateBlsKeys(t *testing.T) {
	tempDir := t.TempDir()

	args := []string{"egnkey", "generate"}
	args = append(args, "--output-dir", tempDir)
	args = append(args, "--key-type", "bls")
	args = append(args, "--num-keys", "1")
	run(args)

	// Descrypt private key from json file and compare it with private_key_hex.txt
	privateKeyHex, err := os.ReadFile(tempDir + "/private_key_hex.txt")
	assert.NoError(t, err)

	password, err := os.ReadFile(tempDir + "/password.txt")
	assert.NoError(t, err)

	decryptedKey, err := sdkEcdsa.ReadKey(tempDir+"/keys/1.bls.key.json", string(password))
	assert.NoError(t, err)
	decryptedKeyBytes := crypto.FromECDSA(decryptedKey)

	if !bytes.Equal(privateKeyHex, decryptedKeyBytes) {
		t.Errorf("Decrypted private key does not match the original private key")
	}
}
