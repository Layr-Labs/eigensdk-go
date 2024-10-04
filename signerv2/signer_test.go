package signerv2_test

import (
	"math/big"
	"testing"

	"github.com/Layr-Labs/eigensdk-go/signerv2"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
)

func TestPrivateKeySignerFn(t *testing.T) {
	privateKeyHex := "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	require.NoError(t, err)
	chainID := big.NewInt(1)

	_, err = signerv2.PrivateKeySignerFn(privateKey, chainID)
	require.NoError(t, err)
}
