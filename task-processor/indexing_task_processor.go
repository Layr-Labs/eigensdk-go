package taskprocessor

import (
	"math/big"
	"sync"
	"time"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/logging"
	blsagg "github.com/Layr-Labs/eigensdk-go/services/bls_aggregation"
	taskmanager "github.com/Layr-Labs/eigensdk-go/task-processor/task-manager"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
	"github.com/Layr-Labs/eigensdk-go/utils"
)

type IndexingTaskProcessor[Input any, Output any] struct {
	tasks   map[sdktypes.TaskIndex]taskmanager.Task[Input]
	tasksMu sync.RWMutex

	taskResponder taskmanager.TaskResponder[Input, Output]

	logger logging.Logger
}

const (
	// number of blocks after which a task is considered expired this hardcoded here because it's also
	//  hardcoded in the contracts, but should ideally be fetched from the contracts
	taskChallengeWindowBlock = 100
	blockTimeSeconds         = 12 * time.Second
)

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

func (itp *IndexingTaskProcessor[Input, Output]) ProcessTaskResponse(taskResponse taskmanager.TaskResponse[Output]) ([32]byte, error) {
	return itp.taskResponder.HashTaskResponse(taskResponse)
}

func (itp *IndexingTaskProcessor[Input, Output]) ProcessAggregatedResponse(response blsagg.BlsAggregationServiceResponse) error {
	itp.tasksMu.RLock()
	task := itp.tasks[response.TaskIndex]
	itp.tasksMu.RUnlock()

	if response.Err != nil {
		return utils.WrapError("BlsAggregationServiceResponse contains an error", response.Err)
	}
	nonSignerPubkeys := []sdktypes.BN254G1Point{}
	for _, nonSignerPubkey := range response.NonSignersPubkeysG1 {
		nonSignerPubkeys = append(nonSignerPubkeys, ConvertToBN254G1Point(nonSignerPubkey))
	}
	quorumApks := []sdktypes.BN254G1Point{}
	for _, quorumApk := range response.QuorumApksG1 {
		quorumApks = append(quorumApks, ConvertToBN254G1Point(quorumApk))
	}
	nonSignerStakesAndSignature := sdktypes.NonSignerStakesAndSignature{
		NonSignerPubkeys:             nonSignerPubkeys,
		QuorumApks:                   quorumApks,
		ApkG2:                        ConvertToBN254G2Point(response.SignersApkG2),
		Sigma:                        ConvertToBN254G1Point(response.SignersAggSigG1.G1Point),
		NonSignerQuorumBitmapIndices: response.NonSignerQuorumBitmapIndices,
		QuorumApkIndices:             response.QuorumApkIndices,
		TotalStakeIndices:            response.TotalStakeIndices,
		NonSignerStakeIndices:        response.NonSignerStakeIndices,
	}

	itp.logger.Info("Threshold reached. Sending aggregated response onchain.", "taskIndex", response.TaskIndex)

	taskResponseAgg, ok := response.TaskResponse.(taskmanager.TaskResponse[Output])
	if !ok {
		itp.logger.Error("task Response could not be converted to sdk aggregator's Task Response type")
	}

	err := itp.taskResponder.RespondToTask(task, taskResponseAgg, nonSignerStakesAndSignature)
	if err != nil {
		return utils.WrapError("Aggregator failed to respond to task", err)
	}

	return nil
}

// Utils
func ConvertToBN254G1Point(input *bls.G1Point) sdktypes.BN254G1Point {
	output := sdktypes.BN254G1Point{
		X: input.X.BigInt(big.NewInt(0)),
		Y: input.Y.BigInt(big.NewInt(0)),
	}
	return output
}

func ConvertToBN254G2Point(input *bls.G2Point) sdktypes.BN254G2Point {
	output := sdktypes.BN254G2Point{
		X: [2]*big.Int{input.X.A1.BigInt(big.NewInt(0)), input.X.A0.BigInt(big.NewInt(0))},
		Y: [2]*big.Int{input.Y.A1.BigInt(big.NewInt(0)), input.Y.A0.BigInt(big.NewInt(0))},
	}
	return output
}
