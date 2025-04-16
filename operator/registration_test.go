package operator_test

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/elcontracts"
	allocationmanager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/AllocationManager"
	regcoord "github.com/Layr-Labs/eigensdk-go/contracts/bindings/RegistryCoordinator"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/operator"
	"github.com/Layr-Labs/eigensdk-go/testutils"
	"github.com/Layr-Labs/eigensdk-go/testutils/testclients"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	gethtypes "github.com/ethereum/go-ethereum/core/types"
)

func TestWatchOperatorActivated(t *testing.T) {
	clients, anvilHttpEndpoint := testclients.BuildTestClients(t)
	contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilHttpEndpoint)

	// Create operator clients
	chainWriter := clients.ElChainWriter

	chainReader := clients.ElChainReader

	avsAddress := contractAddrs.ServiceManager
	operatorSetId := uint32(1)
	erc20MockStrategyAddr := contractAddrs.Erc20MockStrategy

	// Create an operator set
	err := createTotalStakeOperatorSet(
		clients,
		erc20MockStrategyAddr,
	)
	require.NoError(t, err)

	// Start watching for operator activation
	operatorAddress := common.HexToAddress(testutils.ANVIL_FIRST_ADDRESS)
	keypair, err := bls.NewKeyPairFromString("0x01")
	require.NoError(t, err)

	watchOpts := &operator.WatchOperatorActivatedOpts{
		Context:      context.Background(),
		OperatorSets: []operator.OperatorSet{{Id: operatorSetId, Avs: avsAddress}},
	}
	operatorActivatedC, sub, err := operator.WatchOperatorActivated(watchOpts, clients.EthWsClient, contractAddrs.AllocationManager)
	require.NoError(t, err)
	defer sub.Unsubscribe()

	// Register the operator
	request := elcontracts.RegistrationRequest{
		OperatorAddress: operatorAddress,
		AVSAddress:      avsAddress,
		OperatorSetIds:  []uint32{operatorSetId},
		WaitForReceipt:  true,
		Socket:          "socket",
		BlsKeyPair:      keypair,
	}

	operatorSet := allocationmanager.OperatorSet{
		Avs: avsAddress,
		Id:  operatorSetId,
	}
	registryCoordinatorAddress := contractAddrs.RegistryCoordinator
	receipt, err := chainWriter.RegisterForOperatorSets(
		context.Background(),
		registryCoordinatorAddress,
		request,
	)

	require.NoError(t, err)
	require.Equal(t, gethtypes.ReceiptStatusSuccessful, receipt.Status)

	isRegistered, err := chainReader.IsOperatorRegisteredWithOperatorSet(
		context.Background(),
		operatorAddress,
		operatorSet,
	)
	require.NoError(t, err)
	require.Equal(t, true, isRegistered)

	// Wait for the operator to be activated

	expectedEvent := operator.OperatorActivated{
		Operator:    operatorAddress,
		OperatorSet: operator.OperatorSet{Id: operatorSetId, Avs: avsAddress},
	}

	assertReceives(t, &expectedEvent, operatorActivatedC, 10*time.Second)
}

// TODO: taken from the writer_test package from chainio. This should be moved to the testing utils
// Creates an operator set with a single strategy. Note that operator set Id will be
// defined sequentially (as the new amount of operator sets minus one)
func createTotalStakeOperatorSet(
	clients *clients.Clients,
	erc20MockStrategyAddr common.Address,
) error {
	waitForReceipt := true

	operatorSetParam := regcoord.ISlashingRegistryCoordinatorTypesOperatorSetParam{
		MaxOperatorCount:        10,
		KickBIPsOfOperatorStake: 100,
		KickBIPsOfTotalStake:    1000,
	}
	minimumStake := big.NewInt(1)

	strategyParams := regcoord.IStakeRegistryTypesStrategyParams{
		Strategy:   erc20MockStrategyAddr,
		Multiplier: big.NewInt(1),
	}
	strategyParamsArray := []regcoord.IStakeRegistryTypesStrategyParams{strategyParams}
	_, err := clients.AvsRegistryChainWriter.CreateTotalDelegatedStakeQuorum(
		context.Background(),
		operatorSetParam,
		minimumStake,
		strategyParamsArray,
		waitForReceipt,
	)
	return err
}

// Asserts that the channel receives the expected value within the timeout duration.
func assertReceives[T any](t *testing.T, expected T, ch <-chan T, timeout time.Duration) {
	t.Helper()
	select {
	case received := <-ch:
		require.Equal(t, expected, received)
	case <-time.After(timeout):
		t.Fatal("did not receive")
	}
}
