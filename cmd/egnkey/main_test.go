package main

import (
	"encoding/hex"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	sdkEcdsa "github.com/Layr-Labs/eigensdk-go/crypto/ecdsa"
	"github.com/stretchr/testify/assert"
)

func TestGenerateBlsKeys(t *testing.T) {
	tempDir := t.TempDir()
	outputDir := filepath.Join(tempDir, "test")

	args := []string{"egnkey", "generate"}
	args = append(args, "--output-dir", outputDir)
	args = append(args, "--key-type", "bls")
	args = append(args, "--num-keys", "1")

	err := run(args)
	assert.NoError(t, err)

	// Decrypt private key from json file and compare it with private_key_hex.txt
	privateKeyPath := filepath.Join(outputDir, "private_key_hex.txt")
	privateKeyBytes, err := os.ReadFile(privateKeyPath)
	assert.NoError(t, err)
	privateKey := strings.TrimSuffix(string(privateKeyBytes), "\n")

	passwordPath := filepath.Join(outputDir, "password.txt")
	passwordBytes, err := os.ReadFile(passwordPath)
	assert.NoError(t, err)
	password := strings.TrimSuffix(string(passwordBytes), "\n")

	blsKeyPath := filepath.Join(outputDir, "keys", "1.bls.key.json")

	decryptedKey, err := bls.ReadPrivateKeyFromFile(blsKeyPath, password)
	assert.NoError(t, err, "error decrypting bls key")

	assert.Equal(t, privateKey, decryptedKey.PrivKey.String())
}

func TestGenerateEcdsaKeys(t *testing.T) {
	tempDir := t.TempDir()
	outputDir := filepath.Join(tempDir, "test")

	args := []string{"egnkey", "generate"}
	args = append(args, "--output-dir", outputDir)
	args = append(args, "--key-type", "ecdsa")
	args = append(args, "--num-keys", "1")

	err := run(args)
	assert.NoError(t, err)

	// Decrypt private key from json file and compare it with private_key_hex.txt
	privateKeyPath := filepath.Join(outputDir, "private_key_hex.txt")
	privateKeyBytes, err := os.ReadFile(privateKeyPath)
	assert.NoError(t, err)
	privateKeyHex := strings.TrimSuffix(string(privateKeyBytes), "\n")
	privateKeyHex = strings.TrimPrefix(privateKeyHex, "0x")

	passwordPath := filepath.Join(outputDir, "password.txt")
	passwordBytes, err := os.ReadFile(passwordPath)
	password := strings.TrimSuffix(string(passwordBytes), "\n")

	assert.NoError(t, err)

	ecdsaKeyPath := filepath.Join(outputDir, "keys", "1.ecdsa.key.json")

	decryptedKey, err := sdkEcdsa.ReadKey(ecdsaKeyPath, password)
	assert.NoError(t, err, "error decrypting ecdsa key")

	privateKeyDecoded, err := hex.DecodeString(privateKeyHex)
	assert.NoError(t, err)
	assert.Equal(t, privateKeyDecoded, decryptedKey.D.Bytes())
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
