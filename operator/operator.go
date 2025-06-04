package operator

import (
	"context"
	"fmt"
	"math/big"

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
	taskmanager "github.com/Layr-Labs/eigensdk-go/task-manager"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
	"github.com/Layr-Labs/eigensdk-go/utils"
)

// The operator responds to tasks created by the task manager, and sends the task responses to
// the aggregator via rpc. To do this, receives a function to calculate the task response and
// another to calculate the task response hash.
type Operator[Input any, Output any] struct {
	logger logging.Logger

	// The operator ID, used to sign the task responses
	operatorId sdktypes.OperatorId

	// The aggregator RPC client is responsible of the communication with the Aggregator
	aggregatorRpcClient AggregatorRpcClienter[Output]

	// The BLS key pair used to sign the task responses
	blsKeypair *bls.KeyPair

	// channel that receives new task created event logs
	newTaskCreatedLogs chan types.Log

	// The ABI of the task manager contract
	taskManagerAbi *abi.ABI

	// The function used to calculate the response for the received tasks
	responseCalculator ResponseCalculator[Input, Output]
	// The function used to hash the task responses
	taskResponseHashFn TaskResponseHashFunction[Output]
}

// The function used to hash the task responses, receiving the generic sdk task response and returning the digest
type TaskResponseHashFunction[Output any] func(taskResponse taskmanager.TaskResponse[Output]) ([32]byte, error)

// NewOperator creates a new Operator with the provided config and the functions to calculate and hashing the response.
func NewOperator[Input any, Output any](
	logger logging.Logger,
	c Config,
	taskManagerAbi *abi.ABI,
	responseCalculator ResponseCalculator[Input, Output],
	taskResponseHashFn TaskResponseHashFunction[Output],
) (*Operator[Input, Output], error) {
	// Note: here we only assign the registry coordinator address because is the only address we use
	// when we use the AVS registry reader. If you want to do more things with avs registry reader,
	// you should add those addresses to the operator config and assign them here.
	avsConfig := avsregistry.Config{
		RegistryCoordinatorAddress: common.HexToAddress(c.RegistryCoordinatorAddress),
	}

	ethHttpClient, err := ethclient.Dial(c.EthRpcUrl)
	if err != nil {
		return nil, utils.WrapError("Failed to create Eth Http client", err)
	}

	avsReader, err := avsregistry.NewReaderFromConfig(avsConfig, ethHttpClient, logger)
	if err != nil {
		logger.Error("Cannot create AvsReader", "err", err)
		return nil, err
	}

	blsKeyPair := c.BlsSignerCfg.BlsKeyPair
	if blsKeyPair == nil {
		logger.Info("Bls Key pair was nil, using the private key store path and password params...")
		blsKeyPair, err = bls.ReadPrivateKeyFromFile(c.BlsSignerCfg.KeystorePath, c.BlsSignerCfg.KeystorePassword)
		if err != nil {
			logger.Errorf("Cannot parse bls private key", "err", err)
			return nil, err
		}
	}

	// Check if operator was registered, if its not registered and register on startup flag is not set, then will fail.
	// If its not registered and should be registered on startup, make the registration.
	operatorIsRegistered, err := avsReader.IsOperatorRegistered(&bind.CallOpts{}, common.HexToAddress(c.OperatorAddress))
	if err != nil {
		logger.Error("Error checking if operator is registered", "err", err)
		return nil, err
	}
	if !operatorIsRegistered {
		if c.Registration.RegisterOnStartup {
			err = registerOperatorOnStartup(
				c.Registration,
				logger,
				avsConfig.RegistryCoordinatorAddress,
				common.HexToAddress(c.OperatorAddress),
				ethHttpClient,
				blsKeyPair,
			)
			if err != nil {
				logger.Errorf("Failure while registering operator on startup: %w", err)
				return nil, err
			}
		} else {
			// We bubble the error all the way up instead of using logger.Fatal because logger.Fatal prints a huge stack
			// trace that hides the actual error message. This error msg is more explicit and doesn't require showing a
			// stack trace to the user.
			return nil, fmt.Errorf(
				"Operator is not registered and register on startup flag is false, try registering operator using the operator-cli before starting operator",
			)
		}
	}

	operatorId, err := avsReader.GetOperatorId(&bind.CallOpts{}, common.HexToAddress(c.OperatorAddress))
	if err != nil {
		logger.Error("Cannot get operator id", "err", err)
		return nil, err
	}

	calculatedOperatorId := sdktypes.OperatorIdFromKeyPair(blsKeyPair)
	if operatorId != calculatedOperatorId {
		return nil, fmt.Errorf("the operator Id calculated from BLS public key does not match the on chain operator Id")
	}

	aggregatorRpcClient, err := NewAggregatorRpcClient[Output](c.AggregatorServerIpPortAddress, logger)
	if err != nil {
		logger.Error("Cannot create AggregatorRpcClient. Is aggregator running?", "err", err)
		return nil, err
	}

	wsClient, err := ethclient.Dial(c.EthWsUrl)
	if err != nil {
		logger.Fatal("error connecting to web socket", "err", err)
	}

	eventHash := taskManagerAbi.Events["NewTaskCreated"].ID
	query := ethereum.FilterQuery{
		Addresses: []common.Address{},
		Topics:    [][]common.Hash{{eventHash}},
	}

	newTaskCreatedLogs := make(chan types.Log)
	_, err = wsClient.SubscribeFilterLogs(context.Background(), query, newTaskCreatedLogs)
	if err != nil {
		logger.Fatal("error subscribing to newTaskCreated events", "err", err)
	}

	if taskResponseHashFn == nil {
		taskResponseType, err := extractTypeFromAbi(taskManagerAbi)
		if err != nil {
			logger.Error("Failed to get task response type in default abi.", "err", err)
			return nil, err
		}

		taskResponseHashFn = getDefaultHashFunction[Output](taskResponseType)
	}

	operator := &Operator[Input, Output]{
		logger:              logger,
		blsKeypair:          blsKeyPair,
		aggregatorRpcClient: *aggregatorRpcClient,
		operatorId:          operatorId,
		newTaskCreatedLogs:  newTaskCreatedLogs,
		taskManagerAbi:      taskManagerAbi,
		responseCalculator:  responseCalculator,
		taskResponseHashFn:  taskResponseHashFn,
	}

	logger.Info("Operator info",
		"operatorId", operatorId,
		"operatorAddr", c.OperatorAddress,
		"operatorG1Pubkey", operator.blsKeypair.GetPubKeyG1(),
		"operatorG2Pubkey", operator.blsKeypair.GetPubKeyG2(),
	)

	return operator, nil
}

// Runs the Operator in a separate goroutine. This should be called only one time per Operator.
// Will return an error if execution fails or nil in case the context is cancelled.
func (o *Operator[Input, Output]) Start(ctx context.Context) <-chan error {
	errChan := make(chan error)

	go func() {
		errChan <- o.run(ctx)
	}()

	return errChan
}

// The run method executes the main loop for the operator, that basically listens to new task created events
// and respond to those tasks, signs them and sends them to the aggregator via aggregator RPC client.
func (o *Operator[Input, Output]) run(ctx context.Context) error {
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
) (*taskmanager.TaskResponse[Output], error) {
	var newTaskCreatedLog taskmanager.NewTaskCreatedEvent[Input]

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
	taskResponse := &taskmanager.TaskResponse[Output]{
		ReferenceTaskIndex: newTaskIndex,
		OutputValue:        output,
	}
	return taskResponse, nil
}

// Receives a task response and signs it with the task response hash function, the BLS signature
// and the operator ID.
func (o *Operator[Input, Output]) signTaskResponse(
	taskResponse *taskmanager.TaskResponse[Output],
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

// Receives an ABI and returns the ABI type for the AVS output value, or an error in case of failure
// The idea is, instead of receiving the type as parameter, read it from the received ABI
func extractTypeFromAbi(taskManagerAbi *abi.ABI) (abi.Type, error) {
	taskResponseType, err := abi.NewType("tuple", "", []abi.ArgumentMarshaling{
		{
			Name: "referenceTaskIndex",
			Type: "uint32",
		},
		{
			Name: "OutputValue", // Left because ABI does not support purely anonymous or underscored fields
			Type: taskManagerAbi.Events["TaskResponded"].Inputs[0].Type.TupleElems[1].String(),
		},
	})
	if err != nil {
		return abi.Type{}, fmt.Errorf("error creating abi task response type: %w", err)
	}

	return taskResponseType, nil
}

// Receives an ABI type and returns a generic task response hash function that uses that type and returns
// the hash of the response encoded on the value
func getDefaultHashFunction[Output any](taskResponseType abi.Type) TaskResponseHashFunction[Output] {
	return func(taskResponse taskmanager.TaskResponse[Output]) ([32]byte, error) {
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
