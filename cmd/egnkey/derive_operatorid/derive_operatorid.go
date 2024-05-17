package derive_operatorid

import (
	"encoding/hex"
	"fmt"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/urfave/cli/v2"
)

var (
	PrivateKey = &cli.StringFlag{
		Name:     "private-key",
		Usage:    "(bn254) private key from which to derive operatorId from",
		Required: true,
	}
)

var Command = &cli.Command{
	Name:        "derive-operator-id",
	Aliases:     []string{"d"},
	Description: `Given a private key, output its associated operatorId (hash of bn254 G1 pubkey).`,
	Action:      derive,
	Flags: []cli.Flag{
		PrivateKey,
	},
}

func derive(c *cli.Context) error {
	sk := c.String(PrivateKey.Name)
	keypair, err := bls.NewKeyPairFromString(sk)
	if err != nil {
		return err
	}
	operatorId := types.OperatorIdFromKeyPair(keypair)
	fmt.Println("0x" + hex.EncodeToString(operatorId[:]))
	return nil
}
