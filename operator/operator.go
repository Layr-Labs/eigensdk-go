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
	"golang.org/x/crypto/sha3"

	sdkaggregator "github.com/Layr-Labs/eigensdk-go/aggregator"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/avsregistry"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/logging"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
	"github.com/Layr-Labs/eigensdk-go/utils"
)

type Operator[Input any, Output any] struct {
	logger              logging.Logger
	operatorId          sdktypes.OperatorId
	aggregatorRpcClient AggregatorRpcClienter[Output]
	EthWsUrl            string
	blsKeypair          *bls.KeyPair
	newTaskCreatedLogs  chan types.Log
	taskManagerAbi      *abi.ABI
	responseCalculator  ResponseCalculator[Input, Output]
	taskResponseHashFn  TaskResponseHashFunction[Output]
}

type TaskResponseHashFunction[Output any] func(taskResponse sdktypes.TaskResponse[Output]) ([32]byte, error)

func NewOperatorFromConfig[Input any, Output any](
	c OperatorConfig,
	responseCalculator ResponseCalculator[Input, Output],
	taskResponseHashFn TaskResponseHashFunction[Output],
) (*Operator[Input, Output], error) {
	avs_config := avsregistry.Config{
		RegistryCoordinatorAddress:    common.HexToAddress(c.AVSRegistryCoordinatorAddress),
		OperatorStateRetrieverAddress: common.HexToAddress(c.OperatorStateRetrieverAddress),
		ServiceManagerAddress:         common.HexToAddress(c.ServiceManagerAddress),
	}

	ethHttpClient, err := ethclient.Dial(c.EthRpcUrl)
	if err != nil {
		return nil, utils.WrapError("Failed to create Eth Http client", err)
	}

	avsReader, err := avsregistry.NewReaderFromConfig(avs_config, ethHttpClient, c.Logger)
	if err != nil {
		c.Logger.Error("Cannot create AvsReader", "err", err)
		return nil, err
	}

	// Check if operator was registered, return error if not
	operatorIsRegistered, err := avsReader.IsOperatorRegistered(&bind.CallOpts{}, common.HexToAddress(c.OperatorAddress))
	if err != nil {
		c.Logger.Error("Error checking if operator is registered", "err", err)
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
		c.Logger.Error("Cannot get operator id", "err", err)
		return nil, err
	}

	aggregatorRpcClient, err := NewAggregatorRpcClient[Output](c.AggregatorServerIpPortAddress, c.Logger)
	if err != nil {
		c.Logger.Error("Cannot create AggregatorRpcClient. Is aggregator running?", "err", err)
		return nil, err
	}

	blsKeyPassword, ok := os.LookupEnv("OPERATOR_BLS_KEY_PASSWORD")
	if !ok {
		c.Logger.Warnf("OPERATOR_BLS_KEY_PASSWORD env var not set. using empty string")
	}
	blsKeyPair, err := bls.ReadPrivateKeyFromFile(c.BlsPrivateKeyStorePath, blsKeyPassword)
	if err != nil {
		c.Logger.Errorf("Cannot parse bls private key", "err", err)
		return nil, err
	}

	wsClient, err := ethclient.Dial(c.EthWsUrl)
	if err != nil {
		c.Logger.Fatal("error connecting to web socket", "err", err)
	}

	eventHash := c.TaskManagerAbi.Events["NewTaskCreated"].ID
	query := ethereum.FilterQuery{
		Addresses: []common.Address{},
		Topics:    [][]common.Hash{{eventHash}},
	}

	newTaskCreatedLogs := make(chan types.Log)
	_, err = wsClient.SubscribeFilterLogs(context.Background(), query, newTaskCreatedLogs)
	if err != nil {
		c.Logger.Fatal("error subscribing to newTaskCreated events", "err", err)
	}

	if taskResponseHashFn == nil {
		taskResponseType, err := extractTypeFromAbi(c.TaskManagerAbi)
		if err != nil {
			c.Logger.Error("Failed to get task response type in default abi.", "err", err)
			return nil, err
		}

		taskResponseHashFn = getDefaultHashFunction[Output](taskResponseType)
	}

	operator := &Operator[Input, Output]{
		logger:              c.Logger,
		blsKeypair:          blsKeyPair,
		aggregatorRpcClient: *aggregatorRpcClient,
		operatorId:          operatorId,
		newTaskCreatedLogs:  newTaskCreatedLogs,
		taskManagerAbi:      c.TaskManagerAbi,
		responseCalculator:  responseCalculator,
		taskResponseHashFn:  taskResponseHashFn,
	}

	c.Logger.Info("Operator info",
		"operatorId", operatorId,
		"operatorAddr", c.OperatorAddress,
		"operatorG1Pubkey", operator.blsKeypair.GetPubKeyG1(),
		"operatorG2Pubkey", operator.blsKeypair.GetPubKeyG2(),
	)

	return operator, nil
}

func (o *Operator[Input, Output]) Start(ctx context.Context) error {
	o.logger.Info("Starting operator.")

	for {
		select {
		case <-ctx.Done():
			return nil

		case log := <-o.newTaskCreatedLogs:
			taskResponse, err := o.processNewTaskCreatedLog(log)
			if err != nil {
				o.logger.Error("Error processing new task created log", "err", err)
				return err
			}
			signedTaskResponse, err := o.signTaskResponse(taskResponse)
			if err != nil {
				o.logger.Info("Error signing task response", "err", err)
				continue
			}
			go o.aggregatorRpcClient.SendSignedTaskResponseToAggregator(signedTaskResponse)
		}
	}
}

// Takes a NewTaskCreatedLog struct as input and returns a TaskResponseHeader struct.
// The TaskResponseHeader struct is the struct that is signed and sent to the contract as a task response.
func (o *Operator[Input, Output]) processNewTaskCreatedLog(
	log types.Log,
) (*sdktypes.TaskResponse[Output], error) {
	var newTaskCreatedLog sdktypes.NewTaskCreatedEvent[Input]

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

	output, err := o.responseCalculator.ComputeResponse(newTaskIndex, newTaskCreatedLog.Task.InputValue)
	if err != nil {
		return nil, fmt.Errorf("error calculating task response: %w", err)
	}
	taskResponse := &sdktypes.TaskResponse[Output]{
		ReferenceTaskIndex: newTaskIndex,
		OutputValue:        output,
	}
	return taskResponse, nil
}

func (o *Operator[Input, Output]) signTaskResponse(
	taskResponse *sdktypes.TaskResponse[Output],
) (*sdkaggregator.SignedTaskResponse[Output], error) {
	taskResponseDigest, err := o.taskResponseHashFn(*taskResponse)
	if err != nil {
		o.logger.Errorf("Failed to get task response digest: %w", err)
		return nil, err
	}

	blsSignature := o.blsKeypair.SignMessage(taskResponseDigest)
	signedTaskResponse := &sdkaggregator.SignedTaskResponse[Output]{
		TaskResponse: *taskResponse,
		BlsSignature: *blsSignature,
		OperatorId:   o.operatorId,
	}
	o.logger.Debug("Signed task response", "signedTaskResponse", signedTaskResponse)
	return signedTaskResponse, nil
}

func extractTypeFromAbi(taskManagerAbi *abi.ABI) (abi.Type, error) {
	taskResponseType, err := abi.NewType("tuple", "", []abi.ArgumentMarshaling{
		{
			Name: "referenceTaskIndex",
			Type: "uint32",
		},
		{
			Name: "OutputValue", // Left because abi does not support purely anonymous or underscored fields
			Type: taskManagerAbi.Events["TaskResponded"].Inputs[0].Type.TupleElems[1].String(),
		},
	})
	if err != nil {
		return abi.Type{}, fmt.Errorf("error creating abi task response type: %w", err)
	}

	return taskResponseType, nil
}

func getDefaultHashFunction[Output any](taskResponseType abi.Type) TaskResponseHashFunction[Output] {
	return func(taskResponse sdktypes.TaskResponse[Output]) ([32]byte, error) {
		arguments := abi.Arguments{
			{
				Type: taskResponseType,
			},
		}

		encodeTaskResponseByte, err := arguments.Pack(taskResponse)
		if err != nil {
			return [32]byte{}, fmt.Errorf("error encoding task response: %w", err)
		}

		var taskResponseDigest [32]byte
		hasher := sha3.NewLegacyKeccak256()
		hasher.Write(encodeTaskResponseByte)
		copy(taskResponseDigest[:], hasher.Sum(nil)[:32])
		return taskResponseDigest, nil
	}
}
