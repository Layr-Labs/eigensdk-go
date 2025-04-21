package aggregator_example

import (
	"context"
	"fmt"
	"math/big"
	"sync"

	"github.com/Layr-Labs/eigensdk-go/aggregator"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/testutils"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/Layr-Labs/eigensdk-go/utils"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/crypto/sha3"

	sdkaggregator "github.com/Layr-Labs/eigensdk-go/aggregator"
	cstaskmanager "github.com/Layr-Labs/eigensdk-go/examples/bindings/taskManager"
	taskgeneratorexample "github.com/Layr-Labs/eigensdk-go/examples/task-generator"
	blsagg "github.com/Layr-Labs/eigensdk-go/services/bls_aggregation"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
)

type TaskResponseData struct {
	TaskResponse              cstaskmanager.IIncredibleSquaringTaskManagerTaskResponse
	TaskResponseMetadata      cstaskmanager.IIncredibleSquaringTaskManagerTaskResponseMetadata
	NonSigningOperatorPubKeys []cstaskmanager.BN254G1Point
}

func main() {
	logger, err := logging.NewZapLogger(logging.Production) // Change here if want to change logging level
	if err != nil {
		println("Failure creating logger")
		return
	}

	ethHttpUrl := "http://localhost:8545"
	ethHttpClient, err := ethclient.Dial(ethHttpUrl)
	if err != nil {
		return
	}

	txMgr, err := taskgeneratorexample.GetTxManager(logger, ethHttpClient, testutils.ANVIL_FIRST_PRIVATE_KEY)
	if err != nil {
		return
	}

	ecdsaPrivateKey, err := crypto.HexToECDSA(testutils.ANVIL_FIRST_PRIVATE_KEY)
	if err != nil {
		logger.Errorf("Cannot parse ecdsa private key", "err", err)
		return
	}

	cfg := aggregator.AggregatorConfig{
		RegistryCoordinatorAddress:    common.HexToAddress("0x7bc06c482dead17c0e297afbc32f6e63d3846650"),
		OperatorStateRetrieverAddress: common.HexToAddress("0x4c5859f0f772848b2d91f1d83e2fe57935348029"),
		ServiceManagerAddress:         common.HexToAddress("0x5f3f1dbd7b74c6b46e8c44f98792a1daf8d69154"),
		EthHttpClient:                 ethHttpClient,
		TxMgr:                         txMgr,
		Logger:                        logger,
		EthHttpUrl:                    ethHttpUrl,
		EthWsUrl:                      "ws://localhost:8545",
		EcdsaPrivateKey:               ecdsaPrivateKey,
		AggregatorServerIpPortAddr:    "localhost:8090",
	}

	taskProcessor, err := NewTaskProcessor(&cfg)
	if err != nil {
		logger.Fatalf(err.Error())
	}

	taskManagerAbi, err := cstaskmanager.ContractIncredibleSquaringTaskManagerMetaData.GetAbi()
	if err != nil {
		logger.Fatalf(err.Error())
	}

	// This is the same hash function used by the operator to hash the task response before signing it.
	hashFunction := func(taskResponse types.TaskResponse) (types.TaskResponseDigest, error) {
		// The order here has to match the field ordering of cstaskmanager.IIncredibleSquaringTaskManagerTaskResponse
		taskResponseType, err := abi.NewType("tuple", "", []abi.ArgumentMarshaling{
			{
				Name: "referenceTaskIndex",
				Type: "uint32",
			},
			{
				Name: "numberSquared",
				Type: "uint256",
			},
		})
		if err != nil {
			return types.TaskResponseDigest{}, utils.WrapError("Error creating taskResponseType", err)
		}
		arguments := abi.Arguments{
			{
				Type: taskResponseType,
			},
		}

		encodeTaskResponseByte, err := arguments.Pack(taskResponse)
		if err != nil {
			return types.TaskResponseDigest{}, utils.WrapError("Error Packing taskResponse", err)
		}

		var taskResponseDigest [32]byte
		hasher := sha3.NewLegacyKeccak256()
		hasher.Write(encodeTaskResponseByte)
		copy(taskResponseDigest[:], hasher.Sum(nil)[:32])

		return taskResponseDigest, nil
	}
	cfg.TaskResponseHashFn = hashFunction

	blockHash := taskManagerAbi.Events["NewTaskCreated"].ID
	agg, err := aggregator.NewAggregator[IncredibleSquaringTaskResponse](cfg, taskProcessor, blockHash)
	if err != nil {
		logger.Fatalf(err.Error())
	}

	err = agg.Start(context.Background())
	if err != nil {
		logger.Fatalf(err.Error())
	}

	return

}

type IncredibleTaskProcessor struct {
	logger        logging.Logger
	avsWriter     AvsWriter
	tasks         map[types.TaskIndex]cstaskmanager.IIncredibleSquaringTaskManagerTask
	tasksMu       sync.RWMutex
	taskResponses map[uint32]TaskResponseData
}

var _ sdkaggregator.TaskProcessor = (*IncredibleTaskProcessor)(nil)

func NewTaskProcessor(c *aggregator.AggregatorConfig) (*IncredibleTaskProcessor, error) {
	avsConfig := AvsConfig{
		Logger: c.Logger,
		//IncredibleSquaringTaskManager: taskMana,
		TxMgr:         c.TxMgr,
		EthHttpClient: c.EthHttpClient,
	}
	avsWriter, err := BuildAvsWriterFromConfig(&avsConfig)
	if err != nil {
		c.Logger.Errorf("Cannot create avsWriter", "err", err)
		return nil, err
	}

	return &IncredibleTaskProcessor{
		logger:        c.Logger,
		avsWriter:     *avsWriter,
		tasks:         make(map[types.TaskIndex]cstaskmanager.IIncredibleSquaringTaskManagerTask),
		taskResponses: make(map[uint32]TaskResponseData),
	}, nil
}

func (tp *IncredibleTaskProcessor) ProcessNewTask(ctx context.Context, log gethtypes.Log) (blsagg.TaskMetadata, error) {
	var newTaskCreatedLog cstaskmanager.ContractIncredibleSquaringTaskManagerNewTaskCreated

	taskManagerAbi, err := cstaskmanager.ContractIncredibleSquaringTaskManagerMetaData.GetAbi()
	if err != nil {
		tp.logger.Fatalf("Error obtaining task manager ABI: %v", err)
	}

	err = taskManagerAbi.UnpackIntoInterface(&newTaskCreatedLog, "NewTaskCreated", log.Data)
	if err != nil {
		return blsagg.TaskMetadata{}, fmt.Errorf("error unpacking the log: %w", err)
	}

	// This is done this way because the taskIndex value in this event is indexed, so we take it from the log
	newTaskIndex := uint32(new(big.Int).SetBytes(log.Topics[1].Bytes()).Uint64())

	tp.logger.Infof("Aggregator received new task: %v: ", newTaskCreatedLog)

	newTask := newTaskCreatedLog.Task
	tp.tasksMu.Lock()
	tp.tasks[newTaskIndex] = newTask
	tp.tasksMu.Unlock()

	quorumThresholdPercentages := make(types.QuorumThresholdPercentages, len(newTask.QuorumNumbers))
	for i := range newTask.QuorumNumbers {
		quorumThresholdPercentages[i] = types.QuorumThresholdPercentage(newTask.QuorumThresholdPercentage)
	}
	// TODO(samlaf): we use seconds for now, but we should ideally pass a blocknumber to the blsAggregationService
	// and it should monitor the chain and only expire the task aggregation once the chain has reached that block
	// number.
	taskTimeToExpiry := taskChallengeWindowBlock * blockTimeSeconds
	var quorumNums types.QuorumNums
	for _, quorumNum := range newTask.QuorumNumbers {
		quorumNums = append(quorumNums, types.QuorumNum(quorumNum))
	}
	metadata := blsagg.NewTaskMetadata(
		newTaskIndex,
		newTask.TaskCreatedBlock,
		quorumNums,
		quorumThresholdPercentages,
		taskTimeToExpiry,
	)

	return metadata, nil
}

func (tp *IncredibleTaskProcessor) ProcessTaskResponse(
	ctx context.Context,
	event aggregator.TaskResponse,
) ([32]byte, error) {
	return event.Digest(), nil
}

func (tp *IncredibleTaskProcessor) ProcessAggregatedResponse(
	ctx context.Context,
	response blsagg.BlsAggregationServiceResponse,
) error {
	if response.Err != nil {
		return utils.WrapError("BlsAggregationServiceResponse contains an error", response.Err)
	}
	nonSignerPubkeys := []cstaskmanager.BN254G1Point{}
	for _, nonSignerPubkey := range response.NonSignersPubkeysG1 {
		nonSignerPubkeys = append(nonSignerPubkeys, core.ConvertToBN254G1Point(nonSignerPubkey))
	}
	quorumApks := []cstaskmanager.BN254G1Point{}
	for _, quorumApk := range response.QuorumApksG1 {
		quorumApks = append(quorumApks, core.ConvertToBN254G1Point(quorumApk))
	}
	nonSignerStakesAndSignature := cstaskmanager.IBLSSignatureCheckerTypesNonSignerStakesAndSignature{
		NonSignerPubkeys:             nonSignerPubkeys,
		QuorumApks:                   quorumApks,
		ApkG2:                        core.ConvertToBN254G2Point(response.SignersApkG2),
		Sigma:                        core.ConvertToBN254G1Point(response.SignersAggSigG1.G1Point),
		NonSignerQuorumBitmapIndices: response.NonSignerQuorumBitmapIndices,
		QuorumApkIndices:             response.QuorumApkIndices,
		TotalStakeIndices:            response.TotalStakeIndices,
		NonSignerStakeIndices:        response.NonSignerStakeIndices,
	}

	tp.logger.Info("Threshold reached. Sending aggregated response onchain.", "taskIndex", response.TaskIndex)

	tp.tasksMu.RLock()
	task := tp.tasks[response.TaskIndex]
	tp.tasksMu.RUnlock()

	taskResponseAgg, ok := response.TaskResponse.(*IncredibleSquaringTaskResponse)
	if !ok {
		tp.logger.Error("task Response could not be converted to sdk aggregator's Task Response type")
	}

	taskResponse := cstaskmanager.IIncredibleSquaringTaskManagerTaskResponse(*taskResponseAgg)

	_, err := tp.avsWriter.SendAggregatedResponse(
		context.Background(),
		task,
		taskResponse,
		nonSignerStakesAndSignature,
	)
	if err != nil {
		return utils.WrapError("Aggregator failed to respond to task", err)
	}
	return nil
}

type IncredibleSquaringTaskResponse struct {
	ReferenceTaskIndex uint32
	NumberSquared      *big.Int
}

func (tr IncredibleSquaringTaskResponse) TaskIndex() types.TaskIndex {
	return tr.ReferenceTaskIndex
}

func (tr IncredibleSquaringTaskResponse) Digest() [32]byte {
	tmresponse := cstaskmanager.IIncredibleSquaringTaskManagerTaskResponse(tr)
	taskResponseHash, err := core.GetTaskResponseDigest(&tmresponse)
	if err != nil {
		return [32]byte{}
	}
	return taskResponseHash
}

// Avs Writer
type AvsConfig struct {
	Logger                        logging.Logger
	IncredibleSquaringTaskManager common.Address
	TxMgr                         txmgr.TxManager
	EthHttpClient                 *ethclient.Client
}

type AvsWriter struct {
	logger              logging.Logger
	TxMgr               txmgr.TxManager
	taskManagerContract *cstaskmanager.ContractIncredibleSquaringTaskManager
}

func BuildAvsWriterFromConfig(c *AvsConfig) (*AvsWriter, error) {
	contractTaskManager, err := cstaskmanager.NewContractIncredibleSquaringTaskManager(
		c.IncredibleSquaringTaskManager,
		c.EthHttpClient,
	)
	if err != nil {
		return nil, utils.WrapError("Failed to fetch task manager contract", err)
	}

	return &AvsWriter{
		logger:              c.Logger,
		TxMgr:               c.TxMgr,
		taskManagerContract: contractTaskManager,
	}, nil
}

func (w *AvsWriter) SendAggregatedResponse(
	ctx context.Context, task cstaskmanager.IIncredibleSquaringTaskManagerTask,
	taskResponse cstaskmanager.IIncredibleSquaringTaskManagerTaskResponse,
	nonSignerStakesAndSignature cstaskmanager.IBLSSignatureCheckerTypesNonSignerStakesAndSignature,
) (*gethtypes.Receipt, error) {
	txOpts, err := w.TxMgr.GetNoSendTxOpts()
	if err != nil {
		w.logger.Errorf("Error getting tx opts")
		return nil, err
	}
	tx, err := w.taskManagerContract.RespondToTask(txOpts, task, taskResponse, nonSignerStakesAndSignature)
	if err != nil {
		w.logger.Error("Error submitting SubmitTaskResponse tx while calling respondToTask", "err", err)
		return nil, err
	}
	receipt, err := w.TxMgr.Send(ctx, tx, true)
	if err != nil {
		w.logger.Errorf("Error submitting respondToTask tx")
		return nil, err
	}
	w.logger.Info("tx hash :respond to task")
	return receipt, nil
}
