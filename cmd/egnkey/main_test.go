package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	sdkEcdsa "github.com/Layr-Labs/eigensdk-go/crypto/ecdsa"
	"github.com/stretchr/testify/assert"
)

func TestGenerateBlsKeys(t *testing.T) {
	tempDir := t.TempDir()

	args := []string{"egnkey", "generate"}
	args = append(args, "--output-dir", tempDir)
	args = append(args, "--key-type", "bls")
	args = append(args, "--num-keys", "1")
	run(args)

	// Decrypt private key from json file and compare it with private_key_hex.txt
	privateKeyPath := filepath.Join(tempDir, "/private_key_hex.txt")
	privateKeyHex, err := os.ReadFile(privateKeyPath)
	assert.NoError(t, err)

	passwordPath := filepath.Join(tempDir, "password.txt")
	password, err := os.ReadFile(passwordPath)
	assert.NoError(t, err)

	blsKeyPath := filepath.Join(tempDir, "/keys/1.bls.key.json")
	// TODO: fix decryption
	decryptedKey, err := bls.ReadPrivateKeyFromFile(blsKeyPath, string(password))
	assert.NoError(t, err, "error decrypting bls key")

	assert.Equal(t, privateKeyHex, decryptedKey.PrivKey.Bytes())
}

func TestGenerateEcdsaKeys(t *testing.T) {
	tempDir := t.TempDir()

	args := []string{"egnkey", "generate"}
	args = append(args, "--output-dir", tempDir)
	args = append(args, "--key-type", "ecdsa")
	args = append(args, "--num-keys", "1")
	run(args)

	// Decrypt private key from json file and compare it with private_key_hex.txt
	privateKeyPath := filepath.Join(tempDir, "/private_key_hex.txt")
	privateKeyHex, err := os.ReadFile(privateKeyPath)
	assert.NoError(t, err)

	passwordPath := filepath.Join(tempDir, "password.txt")
	password, err := os.ReadFile(passwordPath)
	assert.NoError(t, err)

	ecdsaKeyPath := filepath.Join(tempDir, "/keys/1.ecdsa.key.json")
	// TODO: fix decryption
	decryptedKey, err := sdkEcdsa.ReadKey(ecdsaKeyPath, string(password))
	assert.NoError(t, err, "error decrypting bls key")

	assert.Equal(t, privateKeyHex, decryptedKey.D.Bytes())
}
