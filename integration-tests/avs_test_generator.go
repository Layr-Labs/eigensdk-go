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
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/operator"
	taskmanager "github.com/Layr-Labs/eigensdk-go/task-manager"
	taskspammer "github.com/Layr-Labs/eigensdk-go/task-spammer"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
)

type AvsConfig[Input any, Output any] struct {
	// Task manager related
	TaskManagerAddr common.Address
	TaskManagerAbi  *abi.ABI

	// Eth Urls
	EthHttpUrl string
	EthWsUrl   string

	// A function that creates a calculator to respond the tasks
	ResponseCalculatorBuilder func() operator.ResponseCalculator[Input, Output]

	// The function to compare the calculated and the received output in the challenger
	EqualFn func(a, b Output) bool

	// The sequence to generate the inputs sent to the task manager
	InputSequence iter.Seq[Input]

	// Address from which the Aggregator will listen to signed task responses from operators
	AggregatorServerIpPortAddr string

	// Avs deployment Addresses
	RegistryCoordinatorAddress    common.Address
	OperatorStateRetrieverAddress common.Address
	AvsAddress                    common.Address

	// Operator private keys
	OperatorBlsPrivateKey string
	OperatorPrivateKey    string

	// Entities private keys
	AggregatorPrivateKey string
	ChallengerPrivateKey string
	// This one must match the task_generator_addr passed to the Task manager in deployment
	TaskSpammerPrivateKey string

	// Core deployment addresses
	AllocationManagerAddr       common.Address
	StrategyAddr                common.Address
	DelegationManagerAddress    common.Address
	RewardsCoordinatorAddress   common.Address
	PermissionControllerAddress common.Address

	// Operator config values
	OperatorAddr string

	// Registration config values
	AmountToMint         string
	AllocatableMagnitude uint64
	OperatorSetId        uint32
	MetadataUrl          string
	Socket               string
	AllocationDelay      uint32

	// Task spammer config values
	TimeBetweenTasks          time.Duration
	QuorumThresholdPercentage uint32
	QuorumNumbers             []uint8
}

type Avs struct {
	Aggregator  <-chan error
	Challenger  <-chan error
	Operator    <-chan error
	TaskSpammer <-chan error
}

func StartAvs[Input any, Output any](t *testing.T, ctx context.Context, config AvsConfig[Input, Output]) Avs {
	// Aggregator
	aggregator := createAvsAggregator(t, config)

	aggErrC := aggregator.Start(ctx)

	// Challenger
	challenger := createAvsChallenger(t, config)

	chErrC := challenger.Start(ctx)

	// Operator
	operator := createAvsOperator(t, config)

	opErrC := operator.Start(ctx)

	// Task Spammer
	taskSpammer := createAvsTaskSpammer(t, config)

	tsErrC := taskSpammer.Start(ctx)

	return Avs{
		Aggregator:  aggErrC,
		Challenger:  chErrC,
		Operator:    opErrC,
		TaskSpammer: tsErrC,
	}
}

func createAvsAggregator[Input any, Output any](t *testing.T, config AvsConfig[Input, Output]) *aggregator.Aggregator[Input, Output] {
	t.Helper()

	logger, err := logging.NewZapLogger(logging.Production)
	require.NoError(t, err, "Failure creating logger")

	cfg := aggregator.Config{
		RegistryCoordinatorAddress:    config.RegistryCoordinatorAddress,
		OperatorStateRetrieverAddress: config.OperatorStateRetrieverAddress,
		EthHttpUrl:                    config.EthHttpUrl,
		EthWsUrl:                      config.EthWsUrl,
		AggregatorServerIpPortAddr:    config.AggregatorServerIpPortAddr,
	}

	ethClient, err := ethclient.Dial(config.EthHttpUrl)
	require.NoError(t, err, "Failure creating ethclient")

	ecdsaPrivateKey, err := crypto.HexToECDSA(config.AggregatorPrivateKey)
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
	config AvsConfig[Input, Output],
) *challenger.Challenger[Input, Output] {
	t.Helper()

	logger, err := logging.NewZapLogger(logging.Production)
	require.NoError(t, err, "Failure creating logger")

	ethHttpClient, err := ethclient.Dial(config.EthHttpUrl)
	require.NoError(t, err, "Failed to create eth client")

	challengerCfg := challenger.Config{
		EthWsUrl:   config.EthWsUrl,
		EthHttpUrl: config.EthHttpUrl,
	}

	responseCalculator := config.ResponseCalculatorBuilder()
	logicValidation := challenger.ResponseValidationFunctionFromResponseCalculator(responseCalculator, config.EqualFn)

	ecdsaPrivateKey, err := crypto.HexToECDSA(config.ChallengerPrivateKey)
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
	config AvsConfig[Input, Output],
) *operator.Operator[Input, Output] {
	t.Helper()

	logger, err := logging.NewZapLogger(logging.Production)
	require.NoError(t, err, "Failure creating logger")

	ecdsaCfg := operator.EcdsaSignerConfig{
		PrivateKey: config.OperatorPrivateKey,
	}

	keyPair, err := bls.NewKeyPairFromString(config.OperatorBlsPrivateKey)
	require.NoError(t, err)

	blsCfg := operator.BlsSignerConfig{
		BlsKeyPair: keyPair,
	}

	amount := new(big.Int)
	amount.SetString(config.AmountToMint, 10)

	depositConfig := operator.DepositConfig{
		StrategyAddrs:         config.StrategyAddr,
		AmountToMint:          amount,
		AllocatableMagnitudes: config.AllocatableMagnitude,
	}

	registrationConfig := operator.RegistrationConfig{
		AvsAddress: config.AvsAddress,

		DelegationManagerAddress: config.DelegationManagerAddress,

		EcdsaSignerCfg: ecdsaCfg,

		OperatorSetConfigs: []operator.OperatorSetConfig{
			{
				ID:       config.OperatorSetId,
				Deposits: []operator.DepositConfig{depositConfig},
			},
		},

		MetadataUrl:     config.MetadataUrl,
		Socket:          config.Socket,
		AllocationDelay: config.AllocationDelay,
	}

	operatorConfig := operator.Config{
		OperatorAddress:               config.OperatorAddr,
		RegistryCoordinatorAddress:    config.RegistryCoordinatorAddress,
		EthRpcUrl:                     config.EthHttpUrl,
		EthWsUrl:                      config.EthWsUrl,
		BlsSignerCfg:                  blsCfg,
		AggregatorServerIpPortAddress: config.AggregatorServerIpPortAddr,
		Registration:                  &registrationConfig,
	}

	responseCalculator := config.ResponseCalculatorBuilder()
	operator, err := operator.NewOperator(logger, operatorConfig, config.TaskManagerAbi, responseCalculator, nil)
	require.NoError(t, err, "Failed to create operator from config")

	return operator
}

func createAvsTaskSpammer[Input any, Output any](
	t *testing.T,
	config AvsConfig[Input, Output],
) *taskspammer.TaskSpammer[Input] {
	t.Helper()

	logger, err := logging.NewZapLogger(logging.Production)
	require.NoError(t, err, "Failure creating logger")

	ethHttpClient, err := ethclient.Dial(config.EthHttpUrl)
	require.NoError(t, err, "Failed to dial ethclient")

	ecdsaPrivateKey, err := crypto.HexToECDSA(config.TaskSpammerPrivateKey)
	require.NoError(t, err, "Failed to create ecdsa private key")

	txMgr, err := txmgr.NewSimpleTxManagerFromPrivateKey(logger, ethHttpClient, ecdsaPrivateKey)
	require.NoError(t, err, "Failed to create transaction manager from private key")

	taskCreator, err := taskmanager.NewTaskManagerFromAbi[Input, Output](config.TaskManagerAddr, config.TaskManagerAbi, txMgr, ethHttpClient)
	require.NoError(t, err, "Failed to create Task Creator")

	taskSpammerConfig := taskspammer.Config{
		// This means TaskSpammer will send tasks every 10 seconds
		TimeBetweenTasks: config.TimeBetweenTasks,

		QuorumThresholdPercentage: config.QuorumThresholdPercentage,
		QuorumNumbers:             config.QuorumNumbers,
	}

	taskSpammer, err := taskspammer.NewTaskSpammer(logger, taskSpammerConfig, taskCreator, config.InputSequence)
	require.NoError(t, err, "Failed to create Task Spammer")

	return taskSpammer
}
