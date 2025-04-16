package challenger

import (
	"context"
	"fmt"
	"math/big"

	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type ChallengerLogic[Input any] interface {
	GetNonSigningOperatorPubKeys(common.Hash)([]BN254G1Point)
	VerifyChallenge(uint32, GenericInputTask[Input], TaskResponseData[Input])(error)
}

type GenericInputTask[Input any] struct {
	InputValue         Input
	TaskCreatedBlock          uint32
	QuorumNumbers             []byte
	QuorumThresholdPercentage uint32
}

type GenericInputTaskResponse[Input any] struct {
	ReferenceTaskIndex uint32
	InputValue      Input
}

type GenericTaskResponseMetadata struct {
	TaskRespondedBlock uint32
	HashOfNonSigners   [32]byte
}

type BN254G1Point struct {
	X *big.Int
	Y *big.Int
}

type TaskResponseData[Input any] struct {
	TaskResponse              GenericInputTaskResponse[Input]
	TaskResponseMetadata      GenericTaskResponseMetadata
	NonSigningOperatorPubKeys []BN254G1Point
}

type NewTaskCreatedEvent[Input any] interface {
	InnerTask()(GenericInputTask[Input])
}

type TaskRespondedEvent[Input any] interface {
	TaskIndex()(uint32)
	GetTaskResponse()(GenericInputTaskResponse[Input])
	GetTaskResponseMetadata()(GenericTaskResponseMetadata)
}


type ChallengerConfig struct {
	EthWsUrl string
	Logger   logging.Logger
}

type Challenger[Input any, NewTaskCreated NewTaskCreatedEvent[Input], TaskResponded TaskRespondedEvent[Input]] struct {
	logger             logging.Logger
	logic              ChallengerLogic[Input]
	taskResponseChan   chan types.Log
	newTaskCreatedChan chan types.Log

	taskManagerAbi	*abi.ABI
	tasks         map[uint32]GenericInputTask[Input]
	taskResponses map[uint32]TaskResponseData[Input]

}

func NewChallenger[Input any, NewTaskCreated NewTaskCreatedEvent[Input], TaskResponded TaskRespondedEvent[Input]](
	c ChallengerConfig,
	logic ChallengerLogic[Input],
	newTaskEventHash common.Hash,
	taskProcessedEventHash common.Hash,
	taskManagerAbi *abi.ABI,
) (*Challenger[Input, NewTaskCreated, TaskResponded], error) {
	client, err := ethclient.Dial(c.EthWsUrl)
	if err != nil {
		c.Logger.Fatalf("error connecting to web socket: %v", err)
	}

	query := ethereum.FilterQuery{
		Addresses: []common.Address{},
		Topics:    [][]common.Hash{{newTaskEventHash}},
	}

	newTaskCreatedLogs := make(chan types.Log)
	_, err = client.SubscribeFilterLogs(context.Background(), query, newTaskCreatedLogs)
	if err != nil {
		c.Logger.Fatalf("error subscribing to newTaskCreated events: %v", err)
	}

	query.Topics[0][0] = taskProcessedEventHash

	taskRespondedLogs := make(chan types.Log)
	_, err = client.SubscribeFilterLogs(context.Background(), query, taskRespondedLogs)
	if err != nil {
		c.Logger.Fatalf("error subscribing to taskResponded events: %v", err)
	}

	return &Challenger[Input, NewTaskCreated, TaskResponded]{
		logger:             c.Logger,
		logic:              logic,
		newTaskCreatedChan: newTaskCreatedLogs,
		taskResponseChan:   taskRespondedLogs,
		taskManagerAbi: taskManagerAbi,
		tasks:              make(map[uint32]GenericInputTask[Input]),
		taskResponses:      make(map[uint32]TaskResponseData[Input]),
	}, nil
}

func (c *Challenger[Input, NewTaskCreated, TaskResponded]) Start(ctx context.Context) error {
	c.logger.Info("Starting Challenger.")

	for {
		select {
		case newTaskCreatedLog := <-c.newTaskCreatedChan:
			c.logger.Info("New task created log received")
			err := c.ProcessNewTaskCreatedLog(newTaskCreatedLog)
			if err != nil {
				c.logger.Fatalf("Error processing NewTaskCreated log: %v", err)
			}
		case taskResponseLog := <-c.taskResponseChan:
			c.logger.Info("Task response log received")
			err := c.ProcessTaskResponseLog(taskResponseLog)
			if err != nil {
				c.logger.Fatalf("Error processing TaskResponded log: %v", err)
			}
		}
	}

}

func (c *Challenger[Input, NewTaskCreated, TaskResponded]) ProcessNewTaskCreatedLog(log types.Log) error {
	var newTaskCreatedLog NewTaskCreated

	err := c.taskManagerAbi.UnpackIntoInterface(&newTaskCreatedLog, "NewTaskCreated", log.Data)
	if err != nil {
		return fmt.Errorf("error unpacking the log: %w", err)
	}

	newTaskIndex := uint32(new(big.Int).SetBytes(log.Topics[1].Bytes()).Uint64())
	c.tasks[newTaskIndex] = newTaskCreatedLog.InnerTask()

	return nil
}

func (c *Challenger[Input, NewTaskCreated, TaskResponded]) ProcessTaskResponseLog(
	log types.Log,
) error {
	var taskRespondedLog TaskResponded

	err := c.taskManagerAbi.UnpackIntoInterface(&taskRespondedLog, "TaskResponded", log.Data)
	if err != nil {
		return fmt.Errorf("error unpacking the log: %w", err)
	}

	taskIndex := taskRespondedLog.TaskIndex()

	// get the inputs necessary for raising a challenge
	nonSigningOperatorPubKeys := c.logic.GetNonSigningOperatorPubKeys(log.TxHash)
	taskResponseData := TaskResponseData[Input]{
		TaskResponse:              taskRespondedLog.GetTaskResponse(),
		TaskResponseMetadata:      taskRespondedLog.GetTaskResponseMetadata(),
		NonSigningOperatorPubKeys: nonSigningOperatorPubKeys,
	}

	c.taskResponses[taskIndex] = taskResponseData

	if task, found := c.tasks[taskIndex]; found {
		err = c.logic.VerifyChallenge(taskIndex, task, taskResponseData)
		if err != nil {
			return fmt.Errorf("error verifying the challenge: %w", err)
		}
	}

	return nil
}
