package main

import (
	"encoding/hex"
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

	err := run(args)
	assert.NoError(t, err)

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

	err := run(args)
	assert.NoError(t, err)

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

func TestDeriveOperatorId(t *testing.T) {
	privateKey := "13710126902690889134622698668747132666439281256983827313388062967626731803599"

	args := []string{"egnkey", "derive-operator-id"}
	args = append(args, "--private-key", privateKey)
	err := run(args)
	assert.NoError(t, err)
}

func TestStore(t *testing.T) {
	tempDir := t.TempDir()
	tempFile := filepath.Join(tempDir, "tempFile")

	privateKeyHex := "6e3de7104453736e95c62d1436519813505e0120e9404dca749c973a66f9a748"
	password := "A4SzV5NRow0ve9B6QSbh"

	args := []string{"egnkey", "convert"}
	args = append(args, "--private-key", privateKeyHex)
	args = append(args, "--password", password)
	args = append(args, "--file", tempFile)

	err := run(args)
	assert.NoError(t, err)

	// Verify the content of tempFile
	key, err := sdkEcdsa.ReadKey(tempFile, password)
	assert.NoError(t, err)

	hexString := hex.EncodeToString(key.D.Bytes())
	assert.Equal(t, privateKeyHex, hexString)
}
