package wallet

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/Layr-Labs/eigensdk-go/signerv2"
	"github.com/Layr-Labs/eigensdk-go/testutils"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
)

var (
	chainId = big.NewInt(31337)
)

func TestPrivateKeyWallet(t *testing.T) {
	logger := testutils.NewTestLogger()

	t.Run("SendTransaction + GetTransactionReceipt", func(t *testing.T) {
		anvilC, err := testutils.StartAnvilContainer("")
		require.NoError(t, err)
		ctxWithTimeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		anvilHttpEndpoint, err := anvilC.Endpoint(ctxWithTimeout, "http")
		require.NoError(t, err)
		ethClient, err := ethclient.Dial(anvilHttpEndpoint)
		require.NoError(t, err)

		ecdsaPrivKey, err := crypto.HexToECDSA(testutils.ANVIL_FIRST_PRIVATE_KEY)
		require.NoError(t, err)
		signerV2, signerAddr, err := signerv2.SignerFromConfig(signerv2.Config{PrivateKey: ecdsaPrivKey}, chainId)
		if err != nil {
			panic(err)
		}

		skWallet, err := NewPrivateKeyWallet(ethClient, signerV2, signerAddr, logger)
		require.NoError(t, err)

		tx := types.NewTx(&types.DynamicFeeTx{
			ChainID:   chainId,
			Nonce:     0,
			GasTipCap: big.NewInt(1),
			GasFeeCap: big.NewInt(1_000_000_000),
			Gas:       21000,
			To:        &signerAddr,
			Value:     big.NewInt(1),
		})
		ctxWithTimeout, cancel = context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		txId, err := skWallet.SendTransaction(ctxWithTimeout, tx)
		require.NoError(t, err)

		// need to give some time for anvil to process the tx and mine the block
		// TODO: shall we expose a public WaitForTxReceipt function in the wallet interface, or somewhere else?
		time.Sleep(3 * time.Second)

		ctxWithTimeout, cancel = context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		receipt, err := skWallet.GetTransactionReceipt(ctxWithTimeout, txId)
		require.NoError(t, err)
		// make sure the txHash in the mined tx receipt matches the once we sent
		require.Equal(t, txId, receipt.TxHash.String())
	})
}
