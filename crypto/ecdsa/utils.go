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

func WriteKey(path string, privateKey *ecdsa.PrivateKey, password string) error {
	UUID, err := uuid.NewRandom()
	if err != nil {
		return err
	}
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
