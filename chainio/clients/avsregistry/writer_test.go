package avsregistry_test

import (
	"context"
	"testing"

	chainioutils "github.com/Layr-Labs/eigensdk-go/chainio/utils"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/types"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
)

func TestWriterMethods(t *testing.T) {
	clients := BuildTestClients(t)
	chainWriter := clients.AvsRegistryChainWriter

	keypair, err := bls.NewKeyPairFromString("0x01")
	require.NoError(t, err)

	addr := gethcommon.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	ecdsaPrivKeyHex := "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	ecdsaPrivateKey, err := crypto.HexToECDSA(ecdsaPrivKeyHex)
	require.NoError(t, err)

	quorumNumbers := types.QuorumNums{0}

	t.Run("register operator", func(t *testing.T) {
		receipt, err := chainWriter.RegisterOperator(
			context.Background(),
			ecdsaPrivateKey,
			keypair,
			quorumNumbers,
			"",
			true,
		)
		require.NoError(t, err)
		require.NotNil(t, receipt)
	})

	t.Run("update stake of operator subset", func(t *testing.T) {
		receipt, err := chainWriter.UpdateStakesOfOperatorSubsetForAllQuorums(
			context.Background(),
			[]gethcommon.Address{addr},
			true,
		)
		require.NoError(t, err)
		require.NotNil(t, receipt)
	})

	t.Run("update stake of entire operator set", func(t *testing.T) {
		receipt, err := chainWriter.UpdateStakesOfEntireOperatorSetForQuorums(
			context.Background(),
			[][]gethcommon.Address{{addr}},
			quorumNumbers,
			true,
		)
		require.NoError(t, err)
		require.NotNil(t, receipt)
	})

	t.Run("deregister operator", func(t *testing.T) {
		receipt, err := chainWriter.DeregisterOperator(
			context.Background(),
			quorumNumbers,
			chainioutils.ConvertToBN254G1Point(keypair.PubKey),
			true,
		)
		require.NoError(t, err)
		require.NotNil(t, receipt)
	})
}
