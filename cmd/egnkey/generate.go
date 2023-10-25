package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"

	"github.com/Layr-Labs/eigensdk-go/crypto/ecdsa"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
	"github.com/urfave/cli/v2"
)

const (
	DefaultKeyFolder = "keys"

	PasswordFile      = "password.txt"
	PrivateKeyHexFile = "private_key_hex.txt"
)

var (
	KeyTypeFlag = &cli.StringFlag{
		Name:     "key-type",
		Usage:    "key type to create (ecdsa or bls)",
		Required: true,
	}
	NumKeysFlag = &cli.IntFlag{
		Name:        "num-keys",
		Usage:       "number of keys to create",
		Required:    true,
		DefaultText: "1",
	}
	OutputDirFlag = &cli.StringFlag{
		Name:     "output-dir",
		Usage:    "folder to store keys",
		Required: false,
	}
)

var commandGenerate = &cli.Command{
	Name:    "generate",
	Aliases: []string{"g"},
	Description: `Generate keys for testing purpose.
This command creates ecdsa or bls key pair for testing purposes. It generates
all the relevant files for reading and keys and decrypted it and also gets 
you the private keys in plaintext.

It creates the following artifacts based on arguments
- passwords.txt - contains all passwords to decrypt keys
- private_key_hex.txt - will create plaintext private keys
- keys/* - create all the encrypted json files in this folder
`,
	Action: generate,
	Flags: []cli.Flag{
		KeyTypeFlag,
		NumKeysFlag,
		OutputDirFlag,
	},
}

func generate(c *cli.Context) error {
	keyType := c.String(KeyTypeFlag.Name)
	if keyType != "ecdsa" && keyType != "bls" {
		return cli.Exit("Invalid key type", 1)
	}
	numKeys := c.Int(NumKeysFlag.Name)
	if numKeys < 1 {
		return cli.Exit("Invalid number of keys", 1)
	}

	folder := c.String(OutputDirFlag.Name)
	if folder == "" {
		id, err := uuid.NewUUID()
		if err != nil {
			return cli.Exit("Failed to generate UUID", 1)
		}
		folder = id.String()
	}

	_, err := os.Stat(folder)
	if !os.IsNotExist(err) {
		return cli.Exit("Folder already exists, choose a different folder or delete the existing folder", 1)
	}

	err = os.MkdirAll(folder+"/"+DefaultKeyFolder, 0755)
	if err != nil {
		return err
	}

	passwordFile, err := os.Create(filepath.Clean(folder + "/" + PasswordFile))
	if err != nil {
		return err
	}
	privateKeyFile, err := os.Create(filepath.Clean(folder + "/" + PrivateKeyHexFile))
	if err != nil {
		return err
	}

	if keyType == "ecdsa" {
		err := generateECDSAKeys(numKeys, folder, passwordFile, privateKeyFile)
		if err != nil {
			return err
		}
	} else if keyType == "bls" {
		err := generateBlsKeys(numKeys, folder, passwordFile, privateKeyFile)
		if err != nil {
			return err
		}
	} else {
		return cli.Exit("Invalid key type", 1)
	}
	log.Println("Generated all keys")

	return nil
}

func generateBlsKeys(numKeys int, path string, passwordFile, privateKeyFile *os.File) error {
	for i := 0; i < numKeys; i++ {
		key, err := bls.GenRandomBlsKeys()
		if err != nil {
			return err
		}

		password := generateRandomPassword()
		if err != nil {
			return err
		}

		privateKeyHex := key.PrivKey.String()
		fileName := fmt.Sprintf("%d.bls.key.json", i+1)
		err = key.SaveToFile(filepath.Clean(path+"/"+DefaultKeyFolder+"/"+fileName), password)
		if err != nil {
			return err
		}

		_, err = passwordFile.WriteString(password + "\n")
		if err != nil {
			return err
		}

		_, err = privateKeyFile.WriteString(privateKeyHex + "\n")
		if err != nil {
			return err
		}
		if (i+1)%50 == 0 {
			log.Printf("Generated %d keys\n", i+1)
		}
	}
	return nil
}

func generateECDSAKeys(numKeys int, path string, passwordFile, privateKeyFile *os.File) error {
	for i := 0; i < numKeys; i++ {
		key, err := crypto.GenerateKey()
		if err != nil {
			return err
		}

		password := generateRandomPassword()
		if err != nil {
			return err
		}

		privateKeyBytes := key.D.Bytes()

		// We don't want to save the private key if it is not 32 bytes
		if len(privateKeyBytes) != 32 {
			log.Println("Private key length is not 32 bytes, skipping this iteration")
			i--
			continue
		}
		privateKeyHex := hex.EncodeToString(privateKeyBytes)
		fileName := fmt.Sprintf("%d.ecdsa.key.json", i+1)
		err = ecdsa.WriteKey(filepath.Clean(path+"/"+DefaultKeyFolder+"/"+fileName), key, password)
		if err != nil {
			return err
		}

		_, err = passwordFile.WriteString(password + "\n")
		if err != nil {
			return err
		}

		_, err = privateKeyFile.WriteString("0x" + privateKeyHex + "\n")
		if err != nil {
			return err
		}
		if (i+1)%50 == 0 {
			log.Printf("Generated %d keys\n", i+1)
		}
	}
	return nil
}

func generateRandomPassword() string {
	// Seed the random number generator
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Define character sets for the password
	uppercaseLetters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowercaseLetters := "abcdefghijklmnopqrstuvwxyz"
	digits := "0123456789"
	//specialSymbols := "!@#$%^&*()-_+=[]{}|;:,.<>?/\\"

	// Combine character sets into one
	//allCharacters := uppercaseLetters + lowercaseLetters + digits + specialSymbols
	allCharacters := uppercaseLetters + lowercaseLetters + digits

	// Length of the password you want
	passwordLength := 20

	// Generate the password
	password := make([]byte, passwordLength)
	for i := range password {
		password[i] = allCharacters[random.Intn(len(allCharacters))]
	}
	return string(password)
}
