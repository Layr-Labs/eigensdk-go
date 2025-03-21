package integration_test

import (
	"context"
	"log"
	"math/big"
	"testing"

	strategy "github.com/Layr-Labs/eigensdk-go/contracts/bindings/IStrategy"
	mockerc20 "github.com/Layr-Labs/eigensdk-go/contracts/bindings/MockERC20"
	servicemanager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/ServiceManagerBase"
	"github.com/Layr-Labs/eigensdk-go/testutils"
	"github.com/Layr-Labs/eigensdk-go/testutils/testclients"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIntegrationRewards(t *testing.T) {
	log.Println("This test takes ~50 seconds to run...")

	clients, anvilHttpEndpoint := testclients.BuildTestClients(t)
	contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilHttpEndpoint)

	contractStrategy, err := strategy.NewContractIStrategy(contractAddrs.Erc20MockStrategy, clients.EthHttpClient)
	require.NoError(t, err)

	tokenAddr, err := contractStrategy.UnderlyingToken(&bind.CallOpts{Context: context.Background()})
	require.NoError(t, err)

	mockToken, err := mockerc20.NewContractMockERC20(tokenAddr, clients.EthHttpClient)
	require.NoError(t, err)

	txMgr := clients.TxManager
	noSendTxOpts, err := txMgr.GetNoSendTxOpts()
	require.NoError(t, err)

	// This is the avsRegistry writer address (anvil nineth address)
	tx, err := mockToken.Mint(noSendTxOpts, common.HexToAddress("0xa0Ee7A142d267C1f36714E4a8F75612F20a79720"), big.NewInt(1000))
	require.NoError(t, err)

	receipt, err := txMgr.Send(context.Background(), tx, true)
	require.NoError(t, err)
	require.Equal(t, receipt.Status, uint64(1))

	amountPerPayment := int64(100)
	numPayments := int64(8)
	tx, err = mockToken.IncreaseAllowance(noSendTxOpts, contractAddrs.ServiceManager,
		big.NewInt(amountPerPayment*numPayments*100000),
	)
	require.NoError(t, err)

	receipt, err = txMgr.Send(context.Background(), tx, true)
	require.NoError(t, err)
	require.Equal(t, receipt.Status, uint64(1))

	// Initially, claimer balance in strategy is zero
	initialBalance, err := mockToken.BalanceOf(&bind.CallOpts{}, common.HexToAddress("0x0000000000000000000000000000000000000001"))
	require.NoError(t, err)
	assert.Zero(t, initialBalance.Int64())

	// Now call the rewards scripts or related and assert the balance has changed.
	stratAndMul := []servicemanager.IRewardsCoordinatorTypesStrategyAndMultiplier{
		{
			Strategy:   contractAddrs.Erc20MockStrategy,
			Multiplier: big.NewInt(1_000_000),
		},
	}
	// These values were taken from Go Incredible Squaring AVS's rewards scripts
	rewardsSubmission := []servicemanager.IRewardsCoordinatorTypesRewardsSubmission{
		{
			StrategiesAndMultipliers: stratAndMul,
			Token:                    tokenAddr,
			Amount:                   big.NewInt(100),
			StartTimestamp:           1743033600,
			Duration:                 6048000,
		},
	}

	receipt, err = clients.AvsRegistryChainWriter.CreateAVSRewardsSubmission(context.Background(), rewardsSubmission, true)
	require.NoError(t, err)
	require.Equal(t, receipt.Status, uint64(1))
}
