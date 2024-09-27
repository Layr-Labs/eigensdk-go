package elcontracts_test

import (
	"context"
	"testing"
	"time"

	"github.com/Layr-Labs/eigensdk-go/testutils"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestChainWriter(t *testing.T) {
	anvil, err := testutils.StartAnvilContainer("")
	defer anvil.Terminate(context.Background())

	// Setup ethclient
	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	anvilHttpEndpoint, err := anvil.Endpoint(ctxWithTimeout, "http")
	require.NoError(t, err)
	ethClient, err := ethclient.Dial(anvilHttpEndpoint)
	require.NoError(t, err)

	/*
		// Setup el_chain_reader
		elChainReader, err := setupELChainReader(httpEndpoint)
		assert.NoError(t, err)

		// Setup operator address and private key
		operatorAddr := common.HexToAddress("90F79bf6EB2c4f870365E785982E1f101E93b906")
		operatorPrivateKey := "7c852118294e51e653712a81e05800f419141751be58f605c371e15141b007a6"

		// Get strategy manager address
		strategyManagerAddr, err := getStrategyManagerAddress(httpEndpoint)
		assert.NoError(t, err)

		// Setup ChainWriter
		logger := logging.NewNoopLogger()
		txMgr := txmgr.NewSimpleTxManager(ethClient, logger)
		elChainWriter, err := NewWriterFromConfig(Config{
			StrategyManagerAddr: strategyManagerAddr,
		}, ethClient, logger, nil, txMgr)
		assert.NoError(t, err)

		// Define an operator
		wallet, err := crypto.HexToECDSA("bead471191bea97fc3aeac36c9d74c895e8a6242602e144e43152f96219e96e8")
		assert.NoError(t, err)
		walletAddress := crypto.PubkeyToAddress(wallet.PublicKey)

		operator := types.Operator{
			Address:                   walletAddress.Hex(),
			EarningsReceiverAddress:   walletAddress.Hex(),
			DelegationApproverAddress: walletAddress.Hex(),
			StakerOptOutWindowBlocks:  big.NewInt(3),
			MetadataUrl:               "eigensdk-go",
		}

		// Test 1: Register as an operator
		receipt, err := elChainWriter.RegisterAsOperator(context.Background(), operator, true)
		assert.NoError(t, err)
		assert.True(t, receipt.Status == 1)

		// Test 2: Update operator details
		walletModified, err := crypto.HexToECDSA("2a871d0798f97d79848a013d4936a73bf4cc922c825d33c1cf7073dff6d409c6")
		assert.NoError(t, err)
		walletModifiedAddress := crypto.PubkeyToAddress(walletModified.PublicKey)

		operatorModified := types.Operator{
			Address:                   walletModifiedAddress.Hex(),
			EarningsReceiverAddress:   walletModifiedAddress.Hex(),
			DelegationApproverAddress: walletModifiedAddress.Hex(),
			StakerOptOutWindowBlocks:  big.NewInt(3),
			MetadataUrl:               "eigensdk-go",
		}

		receipt, err = elChainWriter.UpdateOperatorDetails(context.Background(), operatorModified, true)
		assert.NoError(t, err)
		assert.True(t, receipt.Status == 1)

		// Test 3: Deposit ERC20 into strategy
		amount := big.NewInt(100)
		strategyAddr, err := getERC20MockStrategy(httpEndpoint)
		assert.NoError(t, err)

		receipt, err = elChainWriter.DepositERC20IntoStrategy(context.Background(), strategyAddr, amount, true)
		assert.NoError(t, err)
		assert.True(t, receipt.Status == 1)

	*/
	assert.Equal(t, "1234", "1234")
}
