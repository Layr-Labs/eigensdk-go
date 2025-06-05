package integration_test

import (
	"context"
	"iter"
	"slices"
	"strconv"
	"strings"
	"testing"
	"time"

	cstaskmanager "github.com/Layr-Labs/eigensdk-go/examples/awesome-vault-service/contracts/bindings/AwesomeVaultTaskManager"
	"github.com/Layr-Labs/eigensdk-go/testutils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
)

func TestAwesomeVaultService(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	anvilC, err := testutils.StartAnvilContainer("../../examples/awesome-vault-service/tests/anvil/awesome-vault-service-anvil-state.json")
	require.NoError(t, err)
	ethHttpUrl, err := anvilC.Endpoint(context.Background(), "http")
	require.NoError(t, err)
	ethWsUrl, err := anvilC.Endpoint(context.Background(), "ws")
	require.NoError(t, err)

	ethClient, err := ethclient.Dial(ethHttpUrl)
	require.NoError(t, err, "Failed to create eth client")

	taskManagerAddress := common.HexToAddress("0x7bc06c482dead17c0e297afbc32f6e63d3846650")

	taskManager, err := cstaskmanager.NewContractAwesomeVaultTaskManager(taskManagerAddress, ethClient)
	require.NoError(t, err, "Failed to create task manager contract")

	initialTaskIndex, err := taskManager.TaskNumber(&bind.CallOpts{})
	require.NoError(t, err, "Failed to get initial task index")
	require.Equal(t, uint32(0), initialTaskIndex, "Initial task index should be 0")

	taskManagerAbi, err := cstaskmanager.ContractAwesomeVaultTaskManagerMetaData.GetAbi()
	require.NoError(t, err, "Failed to get task manager abi")

	equalFn := func(a, b [32]byte) bool { return a == b }

	testConfig := AvsConfig[TaskInput, [32]byte]{
		TaskManagerAddr: taskManagerAddress,
		TaskManagerAbi:  taskManagerAbi,

		EthHttpUrl: ethHttpUrl,
		EthWsUrl:   ethWsUrl,

		ResponseCalculator: NewVaultServiceResponseCalculator(),
		EqualFn:            equalFn,
		InputSequence:      LinearRangeSequence(),

		RegistryCoordinatorAddress:    "0xfd471836031dc5108809d173a067e8486b9047a3",
		OperatorStateRetrieverAddress: common.HexToAddress("0x5f3f1dbd7b74c6b46e8c44f98792a1daf8d69154"),
		AvsAddress:                    common.HexToAddress("0xcd8a1c3ba11cf5ecfa6267617243239504a98d90"),

		EcdsaPrivateKey: "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80",
		BlsPrivateKey:   "0x2518600ef40ef39cb4ab8b828ce303b3e02ac01ec6ba6bd0d0cf0663e1252ff0",

		TaskSpammerPrivateKey: "4bbbf85ce3377467afe5d46f804f221813b2bb87f24d81f60f1fcdbf7cbf4356",

		AggregatorServerIpPortAddr:  "localhost:8092",
		AggregatorPrivateKey:        "2a871d0798f97d79848a013d4936a73bf4cc922c825d33c1cf7073dff6d409c6",
		ChallengerPrivateKey:        testutils.ANVIL_FIRST_PRIVATE_KEY,
		AmountToMint:                "1000000000000000000000",
		AllocationManagerAddr:       common.HexToAddress("0x2279b7a0a67db372996a5fab50d91eaa73d2ebe6"),
		StrategyAddr:                common.HexToAddress("0x2b961e3959b79326a8e7f64ef0d2d825707669b5"),
		DelegationManagerAddress:    common.HexToAddress("0x9fe46736679d2d9a65f0992f2272de9f3c7fa6e0"),
		RewardsCoordinatorAddress:   common.HexToAddress("0xa51c1fc2f0d1a1b8494ed1fe312d7c3a78ed91c0"),
		PermissionControllerAddress: common.HexToAddress("0x59b670e9fa9d0a427751af201d676719a970857b"),

		AllocatableMagnitude: 1000000000000000,
		OperatorSetId:        0,

		MetadataUrl:     "",
		Socket:          "",
		AllocationDelay: 0,

		BlsKeystorePassword: "",

		OperatorAddr: "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266",

		TimeBetweenTasks:          10 * time.Second,
		QuorumThresholdPercentage: 100,
		QuorumNumbers:             []uint8{0},
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

type TaskInput struct {
	Key   string
	Value string
}

type VaultServiceResponseCalculator struct {
	vaults []TaskInput
}

func NewVaultServiceResponseCalculator() *VaultServiceResponseCalculator {
	vaults := make([]TaskInput, 0)

	return &VaultServiceResponseCalculator{
		vaults: vaults,
	}
}

func (vsrc *VaultServiceResponseCalculator) ComputeResponse(taskIndex uint32, input TaskInput) ([32]byte, error) {
	cmpFn := func(vault TaskInput, key string) int {
		return strings.Compare(vault.Key, key)
	}
	index, wasFound := slices.BinarySearchFunc(vsrc.vaults, input.Key, cmpFn)
	if wasFound {
		vsrc.vaults[index].Value = input.Value
	} else {
		vsrc.vaults = slices.Insert(vsrc.vaults, index, input)
	}

	return ComputeVaultsRoot(vsrc.vaults), nil
}

func ComputeVaultsRoot(vaults []TaskInput) [32]byte {
	leaves := make([][32]byte, len(vaults))
	for i, vault := range vaults {
		leaves[i] = hashVault(vault)
	}
	for len(leaves) > 1 {
		halfLength := (len(leaves) + 1) / 2
		for i := range halfLength {
			rightIdx := i*2 + 1
			if rightIdx >= len(leaves) {
				rightIdx = i * 2
			}
			leaves[i] = hashNodes(leaves[i*2], leaves[rightIdx])
		}
		leaves = leaves[:halfLength]
	}
	return leaves[0]
}

func hashVault(input TaskInput) [32]byte {
	return crypto.Keccak256Hash([]byte(input.Key + input.Value))
}

func hashNodes(leftNode [32]byte, rightNode [32]byte) [32]byte {
	if slices.Compare(leftNode[:], rightNode[:]) > 0 {
		leftNode, rightNode = rightNode, leftNode
	}
	return crypto.Keccak256Hash(leftNode[:], rightNode[:])
}

func LinearRangeSequence() iter.Seq[TaskInput] {
	keys := []string{"foo", "bar", "baz"}
	n := 1
	count := 0
	return func(yield func(TaskInput) bool) {
		for count < 3 {
			input := TaskInput{Key: keys[n%len(keys)], Value: strconv.Itoa(n)}
			if !yield(input) {
				break
			}
			n++
			count++
		}
	}
}
