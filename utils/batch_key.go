package utils

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

const (
	// TODO: move to a common constatns file which
	// can be used by utils and cmd
	DefaultKeyFolder  = "keys"
	PasswordFile      = "password.txt"
	PrivateKeyHexFile = "private_key_hex.txt"
)

// ReadBatchKeys reads the batch keys from the given folder
// and returns the list of BatchKey
// folder: folder where the keys are stored, relative to the current directory
func ReadBatchKeys(folder string, isECDSA bool) ([]BatchKey, error) {
	var batchKey []BatchKey
	absFolder, err := filepath.Abs(folder)
	if err != nil {
		return nil, err
	}

	// read the private key file
	privateKeyFile, err := os.Open(absFolder + "/" + PrivateKeyHexFile)
	if err != nil {
		fmt.Println("Error opening the file:", err)
		return nil, err
	}
	defer func(privateKeyFile *os.File) {
		err := privateKeyFile.Close()
		if err != nil {

		}
	}(privateKeyFile)

	// read the password file
	passwordFile, err := os.Open(absFolder + "/" + PasswordFile)
	if err != nil {
		fmt.Println("Error opening the file:", err)
		return nil, err
	}
	defer func(passwordFile *os.File) {
		err := passwordFile.Close()
		if err != nil {

		}
	}(passwordFile)

	privateKeyScanner := bufio.NewScanner(privateKeyFile)
	passwordScanner := bufio.NewScanner(passwordFile)

	// To create file names
	fileCtr := 1
	keyType := "ecdsa"
	if !isECDSA {
		keyType = "bls"
	}
	for privateKeyScanner.Scan() && passwordScanner.Scan() {
		privateKey := privateKeyScanner.Text()
		password := passwordScanner.Text()
		fileName := fmt.Sprintf("%d.%s.key.json", fileCtr, keyType)
		filePath := fmt.Sprintf("%s/%s/%s", absFolder, DefaultKeyFolder, fileName)
		// Since a last line with empty string is also read
		// I need to break here
		// TODO(madhur): remove empty string when it gets created
		if privateKey == "" {
			break
		}

		bK := BatchKey{
			FilePath:   filePath,
			Password:   password,
			PrivateKey: privateKey,
		}
		batchKey = append(batchKey, bK)
		fileCtr++
	}

	if err := privateKeyScanner.Err(); err != nil {
		fmt.Println("Error reading privateKey file: ", err)
	}

	if err := passwordScanner.Err(); err != nil {
		fmt.Println("Error reading password file: ", err)
	}

	return batchKey, nil
}
