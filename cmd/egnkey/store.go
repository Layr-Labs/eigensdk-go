package main

import (
	"github.com/Layr-Labs/eigensdk-go/crypto/ecdsa"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/urfave/cli/v2"
)

const defaultFile = "key.json"

var (
	PrivateKey = &cli.StringFlag{
		Name:     "private-key",
		Usage:    "private key to store (in hex)",
		Required: true,
	}
	OutputFile = &cli.StringFlag{
		Name:     "file",
		Usage:    "file to store key",
		Required: false,
	}
	Password = &cli.StringFlag{
		Name:     "password",
		Usage:    "password to encrypt key",
		Required: false,
	}
)

var commandStore = &cli.Command{
	Name:        "convert",
	Aliases:     []string{"c"},
	Description: `Stores an ecdsa key to a file, in web3 secret storage format.`,
	Action:      store,
	Flags: []cli.Flag{
		PrivateKey,
		OutputFile,
		Password,
	},
}

func store(c *cli.Context) error {
	outputFile := c.String(OutputFile.Name)
	if outputFile == "" {
		outputFile = defaultFile
	}
	// use geth crypto lib to create an ecdsa private key
	privateKey, err := crypto.HexToECDSA(c.String(PrivateKey.Name))
	if err != nil {
		return err
	}
	password := c.String(Password.Name)
	err = ecdsa.WriteKey(outputFile, privateKey, password)
	if err != nil {
		return err
	}
	return nil
}
