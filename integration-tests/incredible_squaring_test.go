package integration_test

import (
	"context"
	"iter"
	"math/big"
	"testing"
	"time"

	"github.com/Layr-Labs/eigensdk-go/aggregator"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	"github.com/Layr-Labs/eigensdk-go/challenger"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/operator"
	taskmanager "github.com/Layr-Labs/eigensdk-go/task-manager"
	taskspammer "github.com/Layr-Labs/eigensdk-go/task-spammer"
	"github.com/Layr-Labs/eigensdk-go/testutils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"

	cstaskmanager "github.com/Layr-Labs/eigensdk-go/examples/incredible-squaring/contracts/bindings/IncredibleSquaringTaskManager"
)

var taskManagerAddress = common.HexToAddress("0x2bdcc0de6be1f7d2ee689a0342d76f52e8efaba3")

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

	taskManager, err := cstaskmanager.NewContractIncredibleSquaringTaskManager(taskManagerAddress, ethClient)
	require.NoError(t, err, "Failed to create task manager contract")

	initialTaskIndex, err := taskManager.TaskNumber(&bind.CallOpts{})
	require.NoError(t, err, "Failed to get initial task index")
	require.Equal(t, uint32(0), initialTaskIndex, "Initial task index should be 0")

	// Aggregator
	aggregator := createIncredibleSquaringAggregator(t, ethHttpUrl, ethWsUrl)

	aggErrC := aggregator.Start(ctx)

	// Challenger
	challenger := createIncredibleSquaringChallenger(t, ethHttpUrl, ethWsUrl)

	chErrC := challenger.Start(ctx)

	// Operator
	operator := createIncredibleSquaringOperator(t, ethHttpUrl, ethWsUrl)

	opErrC := operator.Start(ctx)

	// Task Spammer
	taskSpammer := createIncredibleSquaringTaskSpammer(t, ethHttpUrl)

	tsErrC := taskSpammer.Start(ctx)

	timer := time.NewTimer(32 * time.Second)

	select {
	case err := <-aggErrC:
		t.Fatal("Aggregator error:", err)
	case err := <-chErrC:
		t.Fatal("Challenger error:", err)
	case err := <-opErrC:
		t.Fatal("Operator error:", err)
	case err := <-tsErrC:
		// Here we handle the input generation termination (in that case returns nil but does not imply an error)
		if err == nil {
			cancel()
		} else {
			t.Fatal("Task Spammer error:", err)
		}
	case <-timer.C:
		// Time passed, so we can stop the test
		cancel()
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

func createIncredibleSquaringAggregator(t *testing.T, ethHttpUrl, ethWsUrl string) *aggregator.Aggregator[*big.Int, *big.Int] {
	t.Helper()

	logger, err := logging.NewZapLogger(logging.Production)
	require.NoError(t, err, "Failure creating logger")

	cfg := aggregator.Config{
		RegistryCoordinatorAddress:    common.HexToAddress("0x7bc06c482dead17c0e297afbc32f6e63d3846650"),
		OperatorStateRetrieverAddress: common.HexToAddress("0x4c5859f0f772848b2d91f1d83e2fe57935348029"),
		EthHttpUrl:                    ethHttpUrl,
		EthWsUrl:                      ethWsUrl,
		AggregatorServerIpPortAddr:    "localhost:8090",
	}

	ethClient, err := ethclient.Dial(ethHttpUrl)
	require.NoError(t, err, "Failure creating ethclient")

	aggregatorPrivateKey := "2a871d0798f97d79848a013d4936a73bf4cc922c825d33c1cf7073dff6d409c6"
	ecdsaPrivateKey, err := crypto.HexToECDSA(aggregatorPrivateKey)
	require.NoError(t, err, "Failed to create ecdsa private key")

	txMgr, err := txmgr.NewSimpleTxManagerFromPrivateKey(logger, ethClient, ecdsaPrivateKey)
	require.NoError(t, err, "Failed to create tx manager from private key")

	taskManagerAbi, err := cstaskmanager.ContractIncredibleSquaringTaskManagerMetaData.GetAbi()
	require.NoError(t, err, "Failed to get task manager abi")

	taskResponder, err := taskmanager.NewTaskManagerFromAbi[*big.Int, *big.Int](
		taskManagerAddress,
		taskManagerAbi,
		txMgr,
		ethClient,
	)
	require.NoError(t, err, "Failed to create Task Responder")

	aggregatorProcessor, err := aggregator.NewIndexingProcessor(logger, taskResponder)
	require.NoError(t, err, "Failed to create Processor")

	aggregator, err := aggregator.NewAggregator(logger, cfg, taskManagerAbi, aggregatorProcessor)
	require.NoError(t, err, "Failed to create aggregator")

	return aggregator
}

func createIncredibleSquaringChallenger(t *testing.T, ethHttpUrl, ethWsUrl string) *challenger.Challenger[*big.Int, *big.Int] {
	t.Helper()

	logger, err := logging.NewZapLogger(logging.Production)
	require.NoError(t, err, "Failure creating logger")

	taskManagerAbi, err := cstaskmanager.ContractIncredibleSquaringTaskManagerMetaData.GetAbi()
	require.NoError(t, err, "Failed to get task manager abi")

	ethHttpClient, err := ethclient.Dial(ethHttpUrl)
	require.NoError(t, err, "Failed to create eth client")

	challengerCfg := challenger.Config{
		EthWsUrl:   ethWsUrl,
		EthHttpUrl: ethHttpUrl,
	}

	squareCalculator := operator.NewFunctionResponseCalculator(square)
	squareValidation := challenger.ResponseValidationFunctionFromResponseCalculator(squareCalculator, func(a, b *big.Int) bool {
		return a.Cmp(b) == 0
	})

	ecdsaPrivateKey, err := crypto.HexToECDSA(testutils.ANVIL_FIRST_PRIVATE_KEY)
	require.NoError(t, err, "Failed to parse ecdsa private key")

	txMgr, err := txmgr.NewSimpleTxManagerFromPrivateKey(logger, ethHttpClient, ecdsaPrivateKey)
	require.NoError(t, err, "Failed to create transaction manager")

	challengeRaiser, err := taskmanager.NewTaskManagerFromAbi[*big.Int, *big.Int](
		taskManagerAddress,
		taskManagerAbi,
		txMgr,
		ethHttpClient,
	)
	require.NoError(t, err, "Failed to create challenge raiser")

	challengerProcessor, err := challenger.NewIndexingProcessor(logger, squareValidation, challengeRaiser)
	require.NoError(t, err, "Failed to create challenger processor")

	challenger, err := challenger.NewChallenger(
		logger,
		challengerCfg,
		taskManagerAbi,
		challengerProcessor,
	)
	require.NoError(t, err, "Failed to create challenger from config")

	return challenger
}

func createIncredibleSquaringOperator(t *testing.T, ethHttpUrl, ethWsUrl string) *operator.Operator[*big.Int, *big.Int] {
	t.Helper()

	logger, err := logging.NewZapLogger(logging.Production)
	require.NoError(t, err, "Failure creating logger")

	taskManagerAbi, err := cstaskmanager.ContractIncredibleSquaringTaskManagerMetaData.GetAbi()
	require.NoError(t, err, "Failed to get task manager abi")

	amount := new(big.Int)
	amount.SetString("1000000000000000000000", 10)
	registrationConfig := operator.RegistrationConfig{
		RegisterOnStartup: true,

		AllocationManagerAddr: common.HexToAddress("0x2279b7a0a67db372996a5fab50d91eaa73d2ebe6"),
		AvsAddress:            common.HexToAddress("0x5f3f1dbd7b74c6b46e8c44f98792a1daf8d69154"),
		StrategyAddrs:         []common.Address{common.HexToAddress("0x2b961e3959b79326a8e7f64ef0d2d825707669b5")},

		DelegationManagerAddress:    common.HexToAddress("0x9fe46736679d2d9a65f0992f2272de9f3c7fa6e0"),
		RewardsCoordinatorAddress:   common.HexToAddress("0xa51c1fc2f0d1a1b8494ed1fe312d7c3a78ed91c0"),
		PermissionControllerAddress: common.HexToAddress("0x59b670e9fa9d0a427751af201d676719a970857b"),

		// Current dir is integration-tests
		EcdsaKeyStorePath: "../examples/incredible-squaring/keys/test.ecdsa.key.json",

		AmountToMint:          amount,
		AllocatableMagnitudes: []uint64{1000000000000000},

		OperatorSetIds: []uint32{0},

		MetadataUrl:     "",
		Socket:          "",
		AllocationDelay: 0,
	}

	blsSignerConfig := operator.BlsSignerConfig{
		BlsPrivateKeyStorePath: "../examples/incredible-squaring/keys/test.bls.key.json",
		BlsPrivateKeyPassword:  "",
	}
	operatorConfig := operator.Config{
		OperatorAddress:            "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266",
		RegistryCoordinatorAddress: "0x7bc06c482dead17c0e297afbc32f6e63d3846650",
		EthRpcUrl:                  ethHttpUrl,
		EthWsUrl:                   ethWsUrl,
		// Current dir is integration-tests
		BlsSignerCfg:                  blsSignerConfig,
		AggregatorServerIpPortAddress: "localhost:8090",
		Registration:                  registrationConfig,
	}

	calculator := operator.NewFunctionResponseCalculator(square)

	operator, err := operator.NewOperator(logger, operatorConfig, taskManagerAbi, calculator, nil)
	require.NoError(t, err, "Failed to create operator from config")

	return operator
}

func createIncredibleSquaringTaskSpammer(t *testing.T, ethHttpUrl string) *taskspammer.TaskSpammer[*big.Int] {
	t.Helper()

	logger, err := logging.NewZapLogger(logging.Production)
	require.NoError(t, err, "Failure creating logger")

	ethHttpClient, err := ethclient.Dial(ethHttpUrl)
	require.NoError(t, err, "Failed to dial ethclient")

	ecdsaPrivateKey, err := crypto.HexToECDSA("2a871d0798f97d79848a013d4936a73bf4cc922c825d33c1cf7073dff6d409c6")
	require.NoError(t, err, "Failed to create ecdsa private key")

	txMgr, err := txmgr.NewSimpleTxManagerFromPrivateKey(logger, ethHttpClient, ecdsaPrivateKey)
	require.NoError(t, err, "Failed to create transaction manager from private key")

	abi, err := cstaskmanager.ContractIncredibleSquaringTaskManagerMetaData.GetAbi()
	require.NoError(t, err, "Failed to get task manager abi")

	taskManagerAddress := taskManagerAddress
	taskCreator, err := taskmanager.NewTaskManagerFromAbi[*big.Int, *big.Int](taskManagerAddress, abi, txMgr, ethHttpClient)
	require.NoError(t, err, "Failed to create Task Creator")

	taskSpammerConfig := taskspammer.Config{
		// This means TaskGenerator will send tasks every 10 seconds
		TimeBetweenTasks: 10 * time.Second,

		QuorumThresholdPercentage: 100,
		QuorumNumbers:             []uint8{0},
	}

	squaringSequence := newNumberToSquareSequence()

	taskSpammer, err := taskspammer.NewTaskSpammer(logger, taskSpammerConfig, taskCreator, squaringSequence)
	require.NoError(t, err, "Failed to create Task Spammer")

	return taskSpammer
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
