package signerv2_test

import (
	"math/big"
	"testing"

	"github.com/Layr-Labs/eigensdk-go/signerv2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
)

func TestPrivateKeySignerFn(t *testing.T) {
	privateKeyHex := "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	require.NoError(t, err)
	chainID := big.NewInt(1)

	signer, err := signerv2.PrivateKeySignerFn(privateKey, chainID)
	require.NoError(t, err)

	address := crypto.PubkeyToAddress(privateKey.PublicKey)
	tx := types.NewTx(&types.DynamicFeeTx{
		Nonce:   0,
		Value:   big.NewInt(0),
		ChainID: chainID,
		Data:    common.Hex2Bytes("6057361d00000000000000000000000000000000000000000000000000000000000f4240"),
	})
	signedTx, err := signer(address, tx)
	require.NoError(t, err)

	// Verify the sender address of the signed transaction
	from, err := types.Sender(types.LatestSignerForChainID(chainID), signedTx)
	require.NoError(t, err)
	require.Equal(t, address, from)
}
