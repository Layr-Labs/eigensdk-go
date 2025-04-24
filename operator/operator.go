package operator

import (
	"context"
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	sdkaggregator "github.com/Layr-Labs/eigensdk-go/aggregator"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/avsregistry"
	"github.com/Layr-Labs/eigensdk-go/challenger"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/logging"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
	"github.com/Layr-Labs/eigensdk-go/utils"
)

type OperatorConfig struct {
	OperatorAddress string

	// Avs Reader addresses
	OperatorStateRetrieverAddress string
	ServiceManagerAddress         string
	AVSRegistryCoordinatorAddress string

	EthRpcUrl string
	EthWsUrl  string

	BlsPrivateKeyStorePath        string
	AggregatorServerIpPortAddress string

	RegisterOnStartup bool
}

type OperatorTaskProcessor[Input any] interface {
	DigestResponse(response *challenger.GenericInputTaskResponse[Input]) [32]byte
}

type Operator[Input any] struct {
	logger                logging.Logger
	operatorId            sdktypes.OperatorId
	aggregatorRpcClient   AggregatorRpcClienter[Input]
	EthWsUrl              string
	blsKeypair            *bls.KeyPair
	taskProcessor         OperatorTaskProcessor[Input]
	newTaskCreatedLogs    chan types.Log
	taskManagerAbi        *abi.ABI
	responseCalculationFn ResponseCalculationFunction[Input]
}

type ResponseCalculationFunction[Input any] func(task challenger.GenericInputTask[Input], taskIndex uint32) (challenger.GenericInputTaskResponse[Input], error)

func NewOperatorFromConfig[Input any](
	c OperatorConfig,
	eventHash common.Hash,
	taskProcessor OperatorTaskProcessor[Input],
	logger logging.Logger,
	taskManagerAbi *abi.ABI,
	responseCalculationFn ResponseCalculationFunction[Input],
) (*Operator[Input], error) {
	avs_config := avsregistry.Config{
		RegistryCoordinatorAddress:    common.HexToAddress(c.AVSRegistryCoordinatorAddress),
		OperatorStateRetrieverAddress: common.HexToAddress(c.OperatorStateRetrieverAddress),
		ServiceManagerAddress:         common.HexToAddress(c.ServiceManagerAddress),
	}

	ethHttpClient, err := ethclient.Dial(c.EthRpcUrl)
	if err != nil {
		return nil, utils.WrapError("Failed to create Eth Http client", err)
	}

	avsReader, err := avsregistry.NewReaderFromConfig(avs_config, ethHttpClient, logger)
	if err != nil {
		logger.Error("Cannot create AvsReader", "err", err)
		return nil, err
	}

	// Check if operator was registered, return error if not
	operatorIsRegistered, err := avsReader.IsOperatorRegistered(&bind.CallOpts{}, common.HexToAddress(c.OperatorAddress))
	if err != nil {
		logger.Error("Error checking if operator is registered", "err", err)
		return nil, err
	}
	if !operatorIsRegistered {
		// We bubble the error all the way up instead of using logger.Fatal because logger.Fatal prints a huge stack
		// trace that hides the actual error message. This error msg is more explicit and doesn't require showing a
		// stack trace to the user.
		return nil, fmt.Errorf(
			"operator is not registered. Registering operator using the operator-cli before starting operator",
		)
	}

	operatorId, err := avsReader.GetOperatorId(&bind.CallOpts{}, common.HexToAddress(c.OperatorAddress))
	if err != nil {
		logger.Error("Cannot get operator id", "err", err)
		return nil, err
	}

	aggregatorRpcClient, err := NewAggregatorRpcClient[Input](c.AggregatorServerIpPortAddress, logger)
	if err != nil {
		logger.Error("Cannot create AggregatorRpcClient. Is aggregator running?", "err", err)
		return nil, err
	}

	blsKeyPassword, ok := os.LookupEnv("OPERATOR_BLS_KEY_PASSWORD")
	if !ok {
		logger.Warnf("OPERATOR_BLS_KEY_PASSWORD env var not set. using empty string")
	}
	blsKeyPair, err := bls.ReadPrivateKeyFromFile(c.BlsPrivateKeyStorePath, blsKeyPassword)
	if err != nil {
		logger.Errorf("Cannot parse bls private key", "err", err)
		return nil, err
	}

	wsClient, err := ethclient.Dial(c.EthWsUrl)
	if err != nil {
		logger.Fatal("error connecting to web socket", "err", err)
	}

	query := ethereum.FilterQuery{
		Addresses: []common.Address{},
		Topics:    [][]common.Hash{{eventHash}},
	}

	newTaskCreatedLogs := make(chan types.Log)
	_, err = wsClient.SubscribeFilterLogs(context.Background(), query, newTaskCreatedLogs)
	if err != nil {
		logger.Fatal("error subscribing to newTaskCreated events", "err", err)
	}

	operator := &Operator[Input]{
		logger:                logger,
		blsKeypair:            blsKeyPair,
		aggregatorRpcClient:   *aggregatorRpcClient,
		operatorId:            operatorId,
		newTaskCreatedLogs:    newTaskCreatedLogs,
		taskProcessor:         taskProcessor,
		taskManagerAbi:        taskManagerAbi,
		responseCalculationFn: responseCalculationFn,
	}

	logger.Info("Operator info",
		"operatorId", operatorId,
		"operatorAddr", c.OperatorAddress,
		"operatorG1Pubkey", operator.blsKeypair.GetPubKeyG1(),
		"operatorG2Pubkey", operator.blsKeypair.GetPubKeyG2(),
	)

	return operator, nil
}

func (o *Operator[Input]) Start(ctx context.Context) error {
	o.logger.Info("Starting operator.")

	for {
		select {
		case <-ctx.Done():
			return nil

		case log := <-o.newTaskCreatedLogs:
			taskResponse, err := o.processNewTaskCreatedLog(log)
			if err != nil {
				o.logger.Error("Error checking if operator is registered", "err", err)
				return err
			}
			signedTaskResponse, err := o.SignTaskResponse(taskResponse)
			if err != nil {
				continue
			}
			go o.aggregatorRpcClient.SendSignedTaskResponseToAggregator(signedTaskResponse)
		}
	}
}

// Takes a NewTaskCreatedLog struct as input and returns a TaskResponseHeader struct.
// The TaskResponseHeader struct is the struct that is signed and sent to the contract as a task response.
func (o *Operator[Input]) processNewTaskCreatedLog(
	log types.Log,
) (*challenger.GenericInputTaskResponse[Input], error) {
	var newTaskCreatedLog challenger.NewTaskCreatedEvent[Input]

	err := o.taskManagerAbi.UnpackIntoInterface(&newTaskCreatedLog, "NewTaskCreated", log.Data)
	if err != nil {
		return nil, fmt.Errorf("error unpacking the log: %w", err)
	}

	newTaskIndex := uint32(new(big.Int).SetBytes(log.Topics[1].Bytes()).Uint64())

	o.logger.Debug("Received new task", "task", newTaskCreatedLog)
	o.logger.Info("Received new task",
		"inputValue", newTaskCreatedLog.Task.InputValue,
		"taskIndex", newTaskIndex,
		"taskCreatedBlock", newTaskCreatedLog.Task.TaskCreatedBlock,
		"quorumNumbers", newTaskCreatedLog.Task.QuorumNumbers,
		"QuorumThresholdPercentage", newTaskCreatedLog.Task.QuorumThresholdPercentage,
	)

	taskResponse, err := o.responseCalculationFn(newTaskCreatedLog.Task, newTaskIndex)
	if err != nil {
		return nil, fmt.Errorf("error calculating task response: %w", err)
	}

	return &taskResponse, nil
}

func (o *Operator[Input]) SignTaskResponse(
	taskResponse *challenger.GenericInputTaskResponse[Input],
) (*sdkaggregator.SignedTaskResponse[Input], error) {
	taskResponseHash := o.taskProcessor.DigestResponse(taskResponse)

	blsSignature := o.blsKeypair.SignMessage(taskResponseHash)
	signedTaskResponse := &sdkaggregator.SignedTaskResponse[Input]{
		TaskResponse: *taskResponse,
		BlsSignature: *blsSignature,
		OperatorId:   o.operatorId,
	}
	o.logger.Debug("Signed task response", "signedTaskResponse", signedTaskResponse)
	return signedTaskResponse, nil
}
