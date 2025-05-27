package taskprocessor

import (
	"sync"
	"time"

	"github.com/Layr-Labs/eigensdk-go/logging"
	blsagg "github.com/Layr-Labs/eigensdk-go/services/bls_aggregation"
	taskmanager "github.com/Layr-Labs/eigensdk-go/task-manager"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
	"github.com/Layr-Labs/eigensdk-go/utils"
)

// The Indexing Task Processor is a generic implementation provided by the SDK that
// satisfies the `TaskProcessor` interface expected by the `Aggregator`
type IndexingTaskProcessor[Input any, Output any] struct {
	tasks   map[sdktypes.TaskIndex]taskmanager.Task[Input]
	tasksMu sync.RWMutex

	// The task responder is the way the processor has to communicate with the task manager contract
	taskResponder taskmanager.TaskResponder[Input, Output]

	logger logging.Logger
}

const (
	// number of blocks after which a task is considered expired this hardcoded here because it's also
	//  hardcoded in the contracts, but should ideally be fetched from the contracts
	taskChallengeWindowBlock = 100
	blockTimeSeconds         = 12 * time.Second
)

// Creates an Indexing Task Processor from a logger and a task responder.
func NewIndexingTaskProcessor[Input any, Output any](
	logger logging.Logger,
	taskResponder taskmanager.TaskResponder[Input, Output],
) (*IndexingTaskProcessor[Input, Output], error) {
	return &IndexingTaskProcessor[Input, Output]{
		tasks:         make(map[sdktypes.TaskIndex]taskmanager.Task[Input]),
		taskResponder: taskResponder,
		logger:        logger,
	}, nil
}

// Processes a new task, saving it in the tasks map and creating the metadata for the BLS aggregation
// service, which it returns
func (itp *IndexingTaskProcessor[Input, Output]) ProcessNewTask(
	taskIndex sdktypes.TaskIndex,
	task taskmanager.Task[Input],
) (blsagg.TaskMetadata, error) {
	itp.logger.Infof("Indexing task processor received new task: %v: ", task)

	itp.tasksMu.Lock()
	itp.tasks[taskIndex] = task
	itp.tasksMu.Unlock()

	quorumThresholdPercentages := make(sdktypes.QuorumThresholdPercentages, len(task.QuorumNumbers))
	for i := range task.QuorumNumbers {
		quorumThresholdPercentages[i] = sdktypes.QuorumThresholdPercentage(task.QuorumThresholdPercentage)
	}

	// TODO(samlaf): we use seconds for now, but we should ideally pass a blocknumber to the blsAggregationService
	// and it should monitor the chain and only expire the task aggregation once the chain has reached that block
	// number.
	taskTimeToExpiry := taskChallengeWindowBlock * blockTimeSeconds
	var quorumNums sdktypes.QuorumNums
	for _, quorumNum := range task.QuorumNumbers {
		quorumNums = append(quorumNums, sdktypes.QuorumNum(quorumNum))
	}
	metadata := blsagg.NewTaskMetadata(
		taskIndex,
		task.TaskCreatedBlock,
		quorumNums,
		quorumThresholdPercentages,
		taskTimeToExpiry,
	)

	return metadata, nil
}

// Processes a Task response, delegating the hashing of the response to the Task responder.
func (itp *IndexingTaskProcessor[Input, Output]) ProcessTaskResponse(taskResponse taskmanager.TaskResponse[Output]) ([32]byte, error) {
	return itp.taskResponder.HashTaskResponse(taskResponse)
}

// Processes an aggregated response, creating the required types and sending them to the on-chain Task Manager contract. After
// sending the response, deletes the completed task from the tasks map.
func (itp *IndexingTaskProcessor[Input, Output]) ProcessAggregatedResponse(
	taskIndex sdktypes.TaskIndex,
	taskResponse taskmanager.TaskResponse[Output],
	nonSignerStakesAndSignature sdktypes.NonSignerStakesAndSignature,
) error {
	itp.tasksMu.RLock()
	task := itp.tasks[taskIndex]
	itp.tasksMu.RUnlock()

	itp.logger.Info("Threshold reached. Sending aggregated response onchain.", "taskIndex", taskIndex)

	err := itp.taskResponder.RespondToTask(task, taskResponse, nonSignerStakesAndSignature)
	if err != nil {
		return utils.WrapError("Aggregator failed to respond to task", err)
	}

	itp.tasksMu.RLock()
	delete(itp.tasks, taskIndex)
	itp.tasksMu.RUnlock()

	return nil
}
