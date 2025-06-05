package integration_test

import (
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
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
)

type AvsConfig[Input any, Output any] struct {
	TaskManagerAddr common.Address

	TaskManagerAbi *abi.ABI

	// The path to the file with the anvil state
	//AnvilStateFileName string

	// A type that implements a task manager inteface would be quite much overhead,
	// maybe a bound contract like with the task manager wrapper

	// The function to calculate the logic
	LogicFn func(taskIndex uint32, input Input) (Output, error)

	// The function to compare the calculated and the received output in the challenger
	EqualFn func(a, b Output) bool

	// The sequence to generate the inputs sent to the task manager
	InputSequence iter.Seq[Input]

	// Avs Addresses
	RegistryCoordinatorAddress    string
	OperatorStateRetrieverAddress common.Address
	AvsAddress                    common.Address

	BlsKeyStorePath   string
	EcdsaKeyStorePath string

	TaskSpammerPrivateKey string

	// Check if really needed
	AggregatorServerIpPortAddr string
}

func createAvsAggregator[Input any, Output any](t *testing.T, ethHttpUrl, ethWsUrl string, config AvsConfig[Input, Output]) *aggregator.Aggregator[Input, Output] {
	t.Helper()

	logger, err := logging.NewZapLogger(logging.Production)
	require.NoError(t, err, "Failure creating logger")

	cfg := aggregator.Config{
		RegistryCoordinatorAddress:    common.HexToAddress(config.RegistryCoordinatorAddress),
		OperatorStateRetrieverAddress: config.OperatorStateRetrieverAddress,
		EthHttpUrl:                    ethHttpUrl,
		EthWsUrl:                      ethWsUrl,
		AggregatorServerIpPortAddr:    config.AggregatorServerIpPortAddr,
	}

	ethClient, err := ethclient.Dial(ethHttpUrl)
	require.NoError(t, err, "Failure creating ethclient")

	aggregatorPrivateKey := "2a871d0798f97d79848a013d4936a73bf4cc922c825d33c1cf7073dff6d409c6"
	ecdsaPrivateKey, err := crypto.HexToECDSA(aggregatorPrivateKey)
	require.NoError(t, err, "Failed to create ecdsa private key")

	txMgr, err := txmgr.NewSimpleTxManagerFromPrivateKey(logger, ethClient, ecdsaPrivateKey)
	require.NoError(t, err, "Failed to create tx manager from private key")

	taskResponder, err := taskmanager.NewTaskManagerFromAbi[Input, Output](
		config.TaskManagerAddr,
		config.TaskManagerAbi,
		txMgr,
		ethClient,
	)
	require.NoError(t, err, "Failed to create Task Responder")

	aggregatorProcessor, err := aggregator.NewIndexingProcessor(logger, taskResponder)
	require.NoError(t, err, "Failed to create Processor")

	aggregator, err := aggregator.NewAggregator(logger, cfg, config.TaskManagerAbi, aggregatorProcessor)
	require.NoError(t, err, "Failed to create aggregator")

	return aggregator
}

func createAvsChallenger[Input any, Output any](
	t *testing.T,
	ethHttpUrl, ethWsUrl string,
	config AvsConfig[Input, Output],
) *challenger.Challenger[Input, Output] {
	t.Helper()

	logger, err := logging.NewZapLogger(logging.Production)
	require.NoError(t, err, "Failure creating logger")

	ethHttpClient, err := ethclient.Dial(ethHttpUrl)
	require.NoError(t, err, "Failed to create eth client")

	challengerCfg := challenger.Config{
		EthWsUrl:   ethWsUrl,
		EthHttpUrl: ethHttpUrl,
	}

	logicCalculator := operator.NewFunctionResponseCalculator(config.LogicFn)
	logicValidation := challenger.ResponseValidationFunctionFromResponseCalculator(logicCalculator, config.EqualFn)

	ecdsaPrivateKey, err := crypto.HexToECDSA(testutils.ANVIL_FIRST_PRIVATE_KEY)
	require.NoError(t, err, "Failed to parse ecdsa private key")

	txMgr, err := txmgr.NewSimpleTxManagerFromPrivateKey(logger, ethHttpClient, ecdsaPrivateKey)
	require.NoError(t, err, "Failed to create transaction manager")

	challengeRaiser, err := taskmanager.NewTaskManagerFromAbi[Input, Output](
		config.TaskManagerAddr,
		config.TaskManagerAbi,
		txMgr,
		ethHttpClient,
	)
	require.NoError(t, err, "Failed to create challenge raiser")

	challengerProcessor, err := challenger.NewIndexingProcessor(logger, logicValidation, challengeRaiser)
	require.NoError(t, err, "Failed to create challenger processor")

	challenger, err := challenger.NewChallenger(
		logger,
		challengerCfg,
		config.TaskManagerAbi,
		challengerProcessor,
	)
	require.NoError(t, err, "Failed to create challenger from config")

	return challenger
}

func createAvsOperator[Input any, Output any](
	t *testing.T,
	ethHttpUrl, ethWsUrl string,
	config AvsConfig[Input, Output],
) *operator.Operator[Input, Output] {
	t.Helper()

	logger, err := logging.NewZapLogger(logging.Production)
	require.NoError(t, err, "Failure creating logger")

	amount := new(big.Int)
	amount.SetString("1000000000000000000000", 10)
	registrationConfig := operator.RegistrationConfig{
		RegisterOnStartup: true,

		AllocationManagerAddr: common.HexToAddress("0x2279b7a0a67db372996a5fab50d91eaa73d2ebe6"),
		AvsAddress:            config.AvsAddress,
		StrategyAddrs:         []common.Address{common.HexToAddress("0x2b961e3959b79326a8e7f64ef0d2d825707669b5")},

		DelegationManagerAddress:    common.HexToAddress("0x9fe46736679d2d9a65f0992f2272de9f3c7fa6e0"),
		RewardsCoordinatorAddress:   common.HexToAddress("0xa51c1fc2f0d1a1b8494ed1fe312d7c3a78ed91c0"),
		PermissionControllerAddress: common.HexToAddress("0x59b670e9fa9d0a427751af201d676719a970857b"),

		EcdsaKeyStorePath: config.EcdsaKeyStorePath,

		AmountToMint:          amount,
		AllocatableMagnitudes: []uint64{1000000000000000},

		OperatorSetIds: []uint32{0},

		MetadataUrl:     "",
		Socket:          "",
		AllocationDelay: 0,
	}

	blsSignerConfig := operator.BlsSignerConfig{
		KeystorePath:     config.BlsKeyStorePath,
		KeystorePassword: "",
	}
	operatorConfig := operator.Config{
		OperatorAddress:               "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266",
		RegistryCoordinatorAddress:    config.RegistryCoordinatorAddress,
		EthRpcUrl:                     ethHttpUrl,
		EthWsUrl:                      ethWsUrl,
		BlsSignerCfg:                  blsSignerConfig,
		AggregatorServerIpPortAddress: config.AggregatorServerIpPortAddr,
		Registration:                  registrationConfig,
	}

	calculator := operator.NewFunctionResponseCalculator(config.LogicFn)

	operator, err := operator.NewOperator(logger, operatorConfig, config.TaskManagerAbi, calculator, nil)
	require.NoError(t, err, "Failed to create operator from config")

	return operator
}

func createAvsTaskSpammer[Input any, Output any](
	t *testing.T,
	ethHttpUrl string,
	config AvsConfig[Input, Output],
) *taskspammer.TaskSpammer[Input] {
	t.Helper()

	logger, err := logging.NewZapLogger(logging.Production)
	require.NoError(t, err, "Failure creating logger")

	ethHttpClient, err := ethclient.Dial(ethHttpUrl)
	require.NoError(t, err, "Failed to dial ethclient")

	ecdsaPrivateKey, err := crypto.HexToECDSA(config.TaskSpammerPrivateKey)
	require.NoError(t, err, "Failed to create ecdsa private key")

	txMgr, err := txmgr.NewSimpleTxManagerFromPrivateKey(logger, ethHttpClient, ecdsaPrivateKey)
	require.NoError(t, err, "Failed to create transaction manager from private key")

	taskCreator, err := taskmanager.NewTaskManagerFromAbi[Input, Output](config.TaskManagerAddr, config.TaskManagerAbi, txMgr, ethHttpClient)
	require.NoError(t, err, "Failed to create Task Creator")

	taskSpammerConfig := taskspammer.Config{
		// This means TaskGenerator will send tasks every 10 seconds
		TimeBetweenTasks: 10 * time.Second,

		QuorumThresholdPercentage: 100,
		QuorumNumbers:             []uint8{0},
	}

	taskSpammer, err := taskspammer.NewTaskSpammer(logger, taskSpammerConfig, taskCreator, config.InputSequence)
	require.NoError(t, err, "Failed to create Task Spammer")

	return taskSpammer
}
