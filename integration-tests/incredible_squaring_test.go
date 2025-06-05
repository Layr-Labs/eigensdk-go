package integration_test

import (
	"context"
	"iter"
	"math/big"
	"testing"
	"time"

	"github.com/Layr-Labs/eigensdk-go/testutils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"

	cstaskmanager "github.com/Layr-Labs/eigensdk-go/examples/incredible-squaring/contracts/bindings/IncredibleSquaringTaskManager"
)

func TestIncredibleSquaring(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	anvilC, err := testutils.StartAnvilContainer("../../examples/incredible-squaring/tests/anvil/incredible-squaring-anvil-state.json")
	require.NoError(t, err)
	ethHttpUrl, err := anvilC.Endpoint(context.Background(), "http")
	require.NoError(t, err)
	ethWsUrl, err := anvilC.Endpoint(context.Background(), "ws")
	require.NoError(t, err)

	ethClient, err := ethclient.Dial(ethHttpUrl)
	require.NoError(t, err, "Failed to create eth client")

	taskManagerAddress := common.HexToAddress("0x2bdcc0de6be1f7d2ee689a0342d76f52e8efaba3")

	taskManager, err := cstaskmanager.NewContractIncredibleSquaringTaskManager(taskManagerAddress, ethClient)
	require.NoError(t, err, "Failed to create task manager contract")

	initialTaskIndex, err := taskManager.TaskNumber(&bind.CallOpts{})
	require.NoError(t, err, "Failed to get initial task index")
	require.Equal(t, uint32(0), initialTaskIndex, "Initial task index should be 0")

	taskManagerAbi, err := cstaskmanager.ContractIncredibleSquaringTaskManagerMetaData.GetAbi()
	require.NoError(t, err, "Failed to get task manager abi")

	equalFn := func(a, b *big.Int) bool {
		return a.Cmp(b) == 0
	}

	testConfig := AvsConfig[*big.Int, *big.Int]{
		TaskManagerAddr: taskManagerAddress,
		TaskManagerAbi:  taskManagerAbi,

		EthHttpUrl: ethHttpUrl,
		EthWsUrl:   ethWsUrl,

		LogicFn:       square,
		EqualFn:       equalFn,
		InputSequence: newNumberToSquareSequence(),

		RegistryCoordinatorAddress:    "0x7bc06c482dead17c0e297afbc32f6e63d3846650",
		OperatorStateRetrieverAddress: common.HexToAddress("0x4c5859f0f772848b2d91f1d83e2fe57935348029"),
		AvsAddress:                    common.HexToAddress("0x5f3f1dbd7b74c6b46e8c44f98792a1daf8d69154"),

		EcdsaKeyStorePath: "../examples/incredible-squaring/keys/test.ecdsa.key.json",
		BlsKeyStorePath:   "../examples/incredible-squaring/keys/test.bls.key.json",

		TaskSpammerPrivateKey: "2a871d0798f97d79848a013d4936a73bf4cc922c825d33c1cf7073dff6d409c6",

		AggregatorServerIpPortAddr: "localhost:8090",
	}

	avs := StartAvs(t, ctx, testConfig)

	timer := time.NewTimer(32 * time.Second)

	select {
	case err := <-avs.Aggregator:
		t.Fatal("Aggregator error:", err)
	case err := <-avs.Challenger:
		t.Fatal("Challenger error:", err)
	case err := <-avs.Operator:
		t.Fatal("Operator error:", err)
	case err := <-avs.TaskSpammer:
		// Here we handle the input generation termination (in that case returns nil but does not imply an error)
		if err == nil {
			cancel()
		} else {
			t.Fatal("Task Spammer error:", err)
		}
	case <-timer.C:
		// If reached this point, there must be a problem with the tasks generation.
		t.Fatal("Timer expired before the task spammer finished sending tasks")
	}

	taskIndex, err := taskManager.TaskNumber(&bind.CallOpts{})
	require.NoError(t, err, "Failed to get final task index")
	require.Equal(t, taskIndex, uint32(3), "Final task index should be equal to the amount of tasks generated")

	for i := range uint32(3) {
		taskResponse, err := taskManager.AllTaskResponses(&bind.CallOpts{}, i)
		require.NoError(t, err, "Failed to get task response number %v", i)
		require.NotZero(t, taskResponse, "Task response number %v should not be nil", i)
	}
}

func square(taskIndex uint32, input *big.Int) (*big.Int, error) {
	result := new(big.Int).Mul(input, input)
	return result, nil
}

// Returns an iterator for the sequence 1, 2, 3, ...
func newNumberToSquareSequence() iter.Seq[*big.Int] {
	acc := big.NewInt(1)
	delta := big.NewInt(1)
	count := 0
	return func(yield func(*big.Int) bool) {
		for count < 3 {
			if !yield(acc) {
				break
			}
			acc.Add(acc, delta)
			count++
		}
	}
}
