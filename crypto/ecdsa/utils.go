package ecdsa

import (
	"bufio"
	"crypto/ecdsa"
	"fmt"
	"os"
	"path/filepath"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
)

func WriteKeyFromHex(path, privateKeyHex, password string) error {
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return err
	}
	return WriteKey(path, privateKey, password)
}

// WriteKey writes the private key to the given path
// The key is encrypted using the given password
// This function will create the directory if it doesn't exist
// If there's an existing file at the given path, it will be overwritten
func WriteKey(path string, privateKey *ecdsa.PrivateKey, password string) error {
	UUID, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	// We are using https://github.com/ethereum/go-ethereum/blob/master/accounts/keystore/key.go#L41
	// to store the keys which requires us to have random UUID for encryption
	key := &keystore.Key{
		Id:         UUID,
		Address:    crypto.PubkeyToAddress(privateKey.PublicKey),
		PrivateKey: privateKey,
	}

	encryptedBytes, err := keystore.EncryptKey(key, password, keystore.StandardScryptN, keystore.StandardScryptP)
	if err != nil {
		return err
	}

	return writeBytesToFile(path, encryptedBytes)
}

func writeBytesToFile(path string, data []byte) error {
	dir := filepath.Dir(path)

	// create the directory if it doesn't exist. If exists, it does nothing
	if err := os.MkdirAll(dir, 0755); err != nil {
		fmt.Println("Error creating directories:", err)
		return err
	}

	file, err := os.Create(filepath.Clean(path))
	if err != nil {
		fmt.Println("file create error")
		return err
	}
	// remember to close the file
	defer func() {
		cerr := file.Close()
		if err == nil {
			err = cerr
		}
	}()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	_, err = file.Write(data)

	return err
}

func ReadKey(keyStoreFile string, password string) (*ecdsa.PrivateKey, error) {
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
