package challenger

import (
	"context"

	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type ChallengerClient interface {
	TransactionByHash(ctx context.Context, hash common.Hash) (tx *types.Transaction, isPending bool, err error)
}

type ChallengerLogic interface {
	ProcessNewTaskCreatedLog(newTaskEvent any) error
	ProcessTaskResponseLog(processTaskEvent any) error
}

type ChallengerConfig struct {
	EthWsUrl string
	Logger   logging.Logger
}

type Challenger struct {
	logger             logging.Logger
	logic              ChallengerLogic
	taskResponseChan   chan types.Log
	newTaskCreatedChan chan types.Log
}

func NewChallenger(c ChallengerConfig, logic ChallengerLogic, newTaskEventHash common.Hash, taskProcessedEventHash common.Hash) (*Challenger, error) {
	client, err := ethclient.Dial(c.EthWsUrl)
	if err != nil {
		c.Logger.Fatalf("error connecting to web socket: %v",err)
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

	query = ethereum.FilterQuery{
		Addresses: []common.Address{},
		Topics:    [][]common.Hash{{taskProcessedEventHash}},
	}

	taskRespondedLogs := make(chan types.Log)
	_, err = client.SubscribeFilterLogs(context.Background(), query, taskRespondedLogs)
	if err != nil {
		c.Logger.Fatalf("error subscribing to newTaskCreated events: %v", err)
	}

	return &Challenger{
		logger:    c.Logger,
		logic:     logic,
		newTaskCreatedChan: newTaskCreatedLogs,
		taskResponseChan: taskRespondedLogs,
	}, nil
}

func (c *Challenger) Start(ctx context.Context) error {
	c.logger.Info("Starting Challenger.")

	for {
		select {
		case newTaskCreatedLog := <-c.newTaskCreatedChan:
			c.logger.Info("New task created log received")
			err := c.logic.ProcessNewTaskCreatedLog(newTaskCreatedLog)
			if err != nil {
				c.logger.Fatalf("Error processing NewTaskCreated log: %v", err)
			}
		case taskResponseLog := <-c.taskResponseChan:
			c.logger.Info("Task response log received")
			err := c.logic.ProcessTaskResponseLog(taskResponseLog)
			if err != nil {
				c.logger.Fatalf("Error processing TaskResponded log: %v", err)
			}
		}
	}

}
