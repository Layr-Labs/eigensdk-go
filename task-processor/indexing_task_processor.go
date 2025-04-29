package taskprocessor

import (
	"sync"
	"time"

	"github.com/Layr-Labs/eigensdk-go/logging"
	blsagg "github.com/Layr-Labs/eigensdk-go/services/bls_aggregation"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
)

type IndexingTaskProcessor[Input any, Output any] struct{
	tasks   map[sdktypes.TaskIndex]sdktypes.GenericInputTask[Input]
	tasksMu sync.RWMutex

	taskResponses   map[sdktypes.TaskIndex]sdktypes.GenericOutputTaskResponse[Output]
	taskResponsesMu sync.RWMutex
	
	taskManagerContract TaskManagerContract[Input, Output]

	logger logging.Logger
}

const (
	// number of blocks after which a task is considered expired this hardcoded here because it's also
	//  hardcoded in the contracts, but should ideally be fetched from the contracts
	taskChallengeWindowBlock = 100
	blockTimeSeconds         = 12 * time.Second
)

func NewIndexingTaskProcessor[Input any, Output any](taskManagerContract TaskManagerContract[Input, Output]) (IndexingTaskProcessor[Input, Output], error) {
	return IndexingTaskProcessor[Input, Output]{
		tasks:	make(map[sdktypes.TaskIndex]sdktypes.GenericInputTask[Input]),
		taskResponses:	make(map[sdktypes.TaskIndex]sdktypes.GenericOutputTaskResponse[Output]),
		taskManagerContract: taskManagerContract,
	}, nil
}

func (itp *IndexingTaskProcessor[Input, Output]) ProcessNewTask(
	taskIndex sdktypes.TaskIndex, 
	task sdktypes.GenericInputTask[Input],
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

func (itp *IndexingTaskProcessor[Input, Output]) ProcessTaskResponse(taskResponse sdktypes.GenericOutputTaskResponse[Output]) ([32]byte, error) {

	return [32]byte{}, nil
}

func (itp *IndexingTaskProcessor[Input, Output]) ProcessAggregatedResponse(response blsagg.BlsAggregationServiceResponse) error {
	return nil
}
