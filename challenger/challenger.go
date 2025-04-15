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

type ChallengerLogic interface {
	// ProcessNewTaskCreatedLog(log types.Log) error
	// ProcessTaskResponseLog(log types.Log) error
	GetNonSigningOperatorPubKeys(common.Hash)([]BN254G1Point)
	VerifyChallenge(uint32, GenericTask, TaskResponseData)(error)
}

type GenericTask interface{

}

type GenericTaskResponse interface{

}

type GenericTaskResponseMetadata interface{

}

type BN254G1Point struct {
	X *big.Int
	Y *big.Int
}

type TaskResponseData struct {
	TaskResponse              GenericTaskResponse
	TaskResponseMetadata      GenericTaskResponseMetadata
	NonSigningOperatorPubKeys []BN254G1Point
}

type NewTaskCreatedEvent interface {
	Task()(GenericTask)
}

type TaskRespondedEvent interface {
	TaskIndex()(uint32)
	TaskResponse()(GenericTaskResponse)
	TaskResponseMetadata()(GenericTaskResponseMetadata)
}


type ChallengerConfig struct {
	EthWsUrl string
	Logger   logging.Logger
}

type Challenger[NewTaskCreated NewTaskCreatedEvent, TaskResponded TaskRespondedEvent, Task GenericTask, TaskResponse TaskResponseData] struct {
	logger             logging.Logger
	logic              ChallengerLogic
	taskResponseChan   chan types.Log
	newTaskCreatedChan chan types.Log

	taskManagerAbi	*abi.ABI
	tasks         map[uint32]GenericTask
	taskResponses map[uint32]TaskResponseData

}

func NewChallenger[NewTaskCreated NewTaskCreatedEvent, TaskResponded TaskRespondedEvent, Task GenericTask, TaskResponse TaskResponseData](
	c ChallengerConfig,
	logic ChallengerLogic,
	newTaskEventHash common.Hash,
	taskProcessedEventHash common.Hash,
	taskManagerAbi *abi.ABI,
) (*Challenger[NewTaskCreated, TaskResponded, Task, TaskResponse], error) {
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

	return &Challenger[NewTaskCreated, TaskResponded, Task, TaskResponse]{
		logger:             c.Logger,
		logic:              logic,
		newTaskCreatedChan: newTaskCreatedLogs,
		taskResponseChan:   taskRespondedLogs,
		taskManagerAbi: taskManagerAbi,
	}, nil
}

func (c *Challenger[NewTaskCreated, TaskResponded, Task, TaskResponse]) Start(ctx context.Context) error {
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

func (c *Challenger[NewTaskCreated, TaskResponded, Task, TaskResponse]) ProcessNewTaskCreatedLog(log types.Log) error {
	var newTaskCreatedLog NewTaskCreated

	err := c.taskManagerAbi.UnpackIntoInterface(&newTaskCreatedLog, "NewTaskCreated", log.Data)
	if err != nil {
		return fmt.Errorf("error unpacking the log: %w", err)
	}

	newTaskIndex := uint32(new(big.Int).SetBytes(log.Topics[1].Bytes()).Uint64())
	c.tasks[newTaskIndex] = newTaskCreatedLog.Task()

	return nil
}

func (c *Challenger[NewTaskCreated, TaskResponded, Task, TaskResponse]) ProcessTaskResponseLog(
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
	taskResponseData := TaskResponseData{
		TaskResponse:              taskRespondedLog.TaskResponse,
		TaskResponseMetadata:      taskRespondedLog.TaskResponseMetadata,
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
