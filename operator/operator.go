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
	logger                logging.Logger
	operatorId            sdktypes.OperatorId
	aggregatorRpcClient   AggregatorRpcClienter[Output]
	EthWsUrl              string
	blsKeypair            *bls.KeyPair
	newTaskCreatedLogs    chan types.Log
	taskManagerAbi        *abi.ABI
	responseCalculationFn ResponseCalculationFunction[Input, Output]
}

type ResponseCalculationFunction[Input any, Output any] func(task sdktypes.GenericInputTask[Input], taskIndex uint32) (sdktypes.GenericOutputTaskResponse[Output], error)

func NewOperatorFromConfig[Input any, Output any](
	c OperatorConfig,
	responseCalculationFn ResponseCalculationFunction[Input, Output],
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

	operator := &Operator[Input, Output]{
		logger:                c.Logger,
		blsKeypair:            blsKeyPair,
		aggregatorRpcClient:   *aggregatorRpcClient,
		operatorId:            operatorId,
		newTaskCreatedLogs:    newTaskCreatedLogs,
		taskManagerAbi:        c.TaskManagerAbi,
		responseCalculationFn: responseCalculationFn,
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
) (*sdktypes.GenericOutputTaskResponse[Output], error) {
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

	taskResponse, err := o.responseCalculationFn(newTaskCreatedLog.Task, newTaskIndex)
	if err != nil {
		return nil, fmt.Errorf("error calculating task response: %w", err)
	}

	return &taskResponse, nil
}

// func outputValueType[T any](value T) string {
// 	switch any(value).(type) {
// 	case string:
// 		return "string"
// 	case int:
// 		return "uint32"
// 	case *big.Int:
// 		return "uint256"
// 	case common.Address:
// 		return "address"
// 	case []byte:
// 		return "bytes"
// 	case sdktypes.Bytes32: // Check if this is the same as using [32]byte
// 		return "bytes"
// 	default:
// 		t := reflect.TypeOf(value)
// 		switch t.Kind() {
// 		case reflect.Slice:
// 			return "tuple[]"
// 		case reflect.Struct:
// 			return "tuple"
// 		default:
// 			return "unknown"
// 		}
// 	}
// }

func (o *Operator[Input, Output]) signTaskResponse(
	taskResponse *sdktypes.GenericOutputTaskResponse[Output],
) (*sdkaggregator.SignedTaskResponse[Output], error) {
	// Calculate the abi output type depending on generic Output type
	// taskResponseType, err := abiTypeFromOutputValue(taskResponse.OutputValue)
	// if err != nil {
	// 	return nil, err
	// }

	taskResponseType, err := abi.NewType("tuple", "", []abi.ArgumentMarshaling{
		{
			Name: "referenceTaskIndex",
			Type: "uint32",
		},
		{
			Name: "OutputValue",
			Type: o.taskManagerAbi.Events["TaskResponded"].Inputs[0].Type.TupleElems[1].String(),
		},
	})
	if err != nil {
		return nil, fmt.Errorf("error creating abi task response type: %w", err)
	}
	arguments := abi.Arguments{
		{
			Type: taskResponseType,
		},
	}

	encodeTaskResponseByte, err := arguments.Pack(taskResponse)
	if err != nil {
		return nil, fmt.Errorf("error encoding task response: %w", err)
	}

	var taskResponseDigest [32]byte
	hasher := sha3.NewLegacyKeccak256()
	hasher.Write(encodeTaskResponseByte)
	copy(taskResponseDigest[:], hasher.Sum(nil)[:32])

	blsSignature := o.blsKeypair.SignMessage(taskResponseDigest)
	signedTaskResponse := &sdkaggregator.SignedTaskResponse[Output]{
		TaskResponse: *taskResponse,
		BlsSignature: *blsSignature,
		OperatorId:   o.operatorId,
	}
	o.logger.Debug("Signed task response", "signedTaskResponse", signedTaskResponse)
	return signedTaskResponse, nil
}
