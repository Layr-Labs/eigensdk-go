package challenger

import (
	"context"
	"fmt"
	"math/big"

	"github.com/Layr-Labs/eigensdk-go/logging"
	taskmanager "github.com/Layr-Labs/eigensdk-go/task-manager"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// The Challenger processor is responsible for processing the challenges
type ChallengerProcessor[Input any, Output any] interface {
	// Processes new tasks, returns an error in case of failure
	ProcessNewTaskCreated(taskIndex uint32, task taskmanager.Task[Input]) error
	// Processes task responses, returns an error in case of failure
	ProcessTaskResponded(taskIndex uint32, taskResponse taskmanager.TaskResponse[Output], taskResponseMetadata sdktypes.TaskResponseMetadata, nonSigningOperatorPubKeys []sdktypes.BN254G1Point) error
}

// The challenger is the entity responsible of validating the task responsed submitted by the BLS aggregation
// service and emited by the Task Manager on-chain contract in the TaskResponded event. To do that will need
// To listen to new task created events (to register the new tasks) and task responded events (to validate
// those task responses).
// Most of these things are delegated to the Challenger Processor interface, that process challenges and
// communicates with the on-chain task manager contract when raising a challenge.
type Challenger[Input any, Output any] struct {
	logger logging.Logger

	// The responsible for processing the challenges
	challengerProcessor ChallengerProcessor[Input, Output]

	// channel that receives task responded event logs
	taskResponseChan chan types.Log
	// channel that receives new task created event logs
	newTaskCreatedChan chan types.Log

	// The ABI of the task manager contract
	taskManagerAbi *abi.ABI

	// The client used to communicate with the anvil node
	ethClient *ethclient.Client
}

// NewChallenger creates a new Aggregator with the provided config and a challenger processor.
func NewChallenger[Input any, Output any](
	logger logging.Logger,
	c Config,
	taskManagerAbi *abi.ABI,
	challengerProcessor ChallengerProcessor[Input, Output],
) (*Challenger[Input, Output], error) {
	client, err := ethclient.Dial(c.EthWsUrl)
	if err != nil {
		logger.Fatalf("error connecting to web socket: %v", err)
	}

	newTaskEventHash := taskManagerAbi.Events["NewTaskCreated"].ID
	query := ethereum.FilterQuery{
		Addresses: []common.Address{},
		Topics:    [][]common.Hash{{newTaskEventHash}},
	}

	newTaskCreatedLogs := make(chan types.Log)
	_, err = client.SubscribeFilterLogs(context.Background(), query, newTaskCreatedLogs)
	if err != nil {
		logger.Fatalf("error subscribing to newTaskCreated events: %v", err)
	}

	taskRespondedEventHash := taskManagerAbi.Events["TaskResponded"].ID
	query.Topics[0][0] = taskRespondedEventHash

	taskRespondedLogs := make(chan types.Log)
	_, err = client.SubscribeFilterLogs(context.Background(), query, taskRespondedLogs)
	if err != nil {
		logger.Fatalf("error subscribing to taskResponded events: %v", err)
	}

	ethClient, err := ethclient.Dial(c.EthHttpUrl)
	if err != nil {
		logger.Fatalf("Failed to dial ethclient: %v", err)
	}

	return &Challenger[Input, Output]{
		logger:              logger,
		challengerProcessor: challengerProcessor,
		newTaskCreatedChan:  newTaskCreatedLogs,
		taskResponseChan:    taskRespondedLogs,
		taskManagerAbi:      taskManagerAbi,
		ethClient:           ethClient,
	}, nil
}

// Runs the Challenger in a separate goroutine. This should be called only one time per Challenger.
// Returns an error channel, that in case of an error in the run method will contain the received error.
func (c *Challenger[Input, Output]) Start(ctx context.Context) <-chan error {
	errChan := make(chan error)

	go func() {
		errChan <- c.run(ctx)
	}()

	return errChan
}

// The run method executes the main loop of the Challenger. This loop has 2 main events:
//   - Receive a task responded event: In this case the challenger will process that event, saving the
//     task index
//   - Receive a new task created event log: In this case the challenger processes that event, and in case
//     the response is wrong, a challenge will be raised
func (c *Challenger[Input, Output]) run(ctx context.Context) error {
	c.logger.Info("Starting Challenger.")

	for {
		select {
		case <-ctx.Done():
			return nil
		case newTaskCreatedLog := <-c.newTaskCreatedChan:
			c.logger.Info("New task created log received")
			err := c.processNewTaskCreatedLog(newTaskCreatedLog)
			if err != nil {
				c.logger.Fatalf("Error processing NewTaskCreated log: %v", err)
			}
		case taskResponseLog := <-c.taskResponseChan:
			c.logger.Info("Task response log received")
			err := c.processTaskRespondedLog(taskResponseLog)
			if err != nil {
				c.logger.Fatalf("Error processing TaskResponded log: %v", err)
			}
		}
	}

}

// When processing a new task created log, the aggregator unpacks the log data into the new task created event and
// sends it to the challenger processor
func (c *Challenger[Input, Output]) processNewTaskCreatedLog(log types.Log) error {
	var newTaskCreatedLog taskmanager.NewTaskCreatedEvent[Input]

	err := c.taskManagerAbi.UnpackIntoInterface(&newTaskCreatedLog, "NewTaskCreated", log.Data)
	if err != nil {
		return fmt.Errorf("error unpacking the log: %w", err)
	}

	newTaskIndex := uint32(new(big.Int).SetBytes(log.Topics[1].Bytes()).Uint64())

	err = c.challengerProcessor.ProcessNewTaskCreated(newTaskIndex, newTaskCreatedLog.Task)
	if err != nil {
		return fmt.Errorf("error processing new task created: %w", err)
	}

	return nil
}

// When processing a task responded log, the challenger unpacks the log data into the task responded event and
// checks if it has to raise a challenge delegating it to the challenger processor
func (c *Challenger[Input, Output]) processTaskRespondedLog(
	log types.Log,
) error {
	var taskRespondedLog taskmanager.TaskRespondedEvent[Output]

	err := c.taskManagerAbi.UnpackIntoInterface(&taskRespondedLog, "TaskResponded", log.Data)
	if err != nil {
		return fmt.Errorf("error unpacking the log: %w", err)
	}

	taskIndex := taskRespondedLog.TaskResponse.ReferenceTaskIndex

	// get the inputs necessary for raising a challenge
	nonSigningOperatorPubKeys := c.getNonSigningOperatorPubKeys(log.TxHash)

	err = c.challengerProcessor.ProcessTaskResponded(
		taskIndex,
		taskRespondedLog.TaskResponse,
		taskRespondedLog.TaskResponseMetadata,
		nonSigningOperatorPubKeys,
	)
	if err != nil {
		return fmt.Errorf("error verifying the challenge: %w", err)
	}

	return nil
}

// Gets the non signing operator public keys from the transaction hash received from the log
func (c *Challenger[Input, Output]) getNonSigningOperatorPubKeys(
	transactionHash common.Hash,
) []sdktypes.BN254G1Point {
	// get the nonSignerStakesAndSignature
	tx, _, err := c.ethClient.TransactionByHash(context.Background(), transactionHash)
	if err != nil {
		c.logger.Error("Error getting transaction by hash",
			"txHash", transactionHash,
			"err", err,
		)
	}

	calldata := tx.Data()
	methodSig := calldata[:4]
	method, err := c.taskManagerAbi.MethodById(methodSig)
	if err != nil {
		c.logger.Error("Error getting method", "err", err)
	}

	inputs, err := method.Inputs.Unpack(calldata[4:])
	if err != nil {
		c.logger.Error("Error unpacking calldata", "err", err)
	}

	// Note: this implies the ABI of the Task Manager contract implemented by the AVS should respect this values, or it
	// wont work. Other solution is to receive this as parameter, but it looks more difficult than only replace this.
	nonSignerStakesAndSignatureInput := inputs[2].(struct {
		NonSignerQuorumBitmapIndices []uint32 "json:\"nonSignerQuorumBitmapIndices\""
		NonSignerPubkeys             []struct {
			X *big.Int "json:\"X\""
			Y *big.Int "json:\"Y\""
		} "json:\"nonSignerPubkeys\""
		QuorumApks []struct {
			X *big.Int "json:\"X\""
			Y *big.Int "json:\"Y\""
		} "json:\"quorumApks\""
		ApkG2 struct {
			X [2]*big.Int "json:\"X\""
			Y [2]*big.Int "json:\"Y\""
		} "json:\"apkG2\""
		Sigma struct {
			X *big.Int "json:\"X\""
			Y *big.Int "json:\"Y\""
		} "json:\"sigma\""
		QuorumApkIndices      []uint32   "json:\"quorumApkIndices\""
		TotalStakeIndices     []uint32   "json:\"totalStakeIndices\""
		NonSignerStakeIndices [][]uint32 "json:\"nonSignerStakeIndices\""
	})

	// get pubkeys of non-signing operators and submit them to the contract
	nonSigningOperatorPubKeys := make(
		[]sdktypes.BN254G1Point,
		len(nonSignerStakesAndSignatureInput.NonSignerPubkeys),
	)
	for i, pubkey := range nonSignerStakesAndSignatureInput.NonSignerPubkeys {
		nonSigningOperatorPubKeys[i] = sdktypes.BN254G1Point{
			X: pubkey.X,
			Y: pubkey.Y,
		}
	}

	return nonSigningOperatorPubKeys
}
