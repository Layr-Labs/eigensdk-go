package challengerexample

import (
	"context"
	"fmt"
	"math/big"

	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	"github.com/Layr-Labs/eigensdk-go/challenger"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/testutils"
	"github.com/Layr-Labs/eigensdk-go/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	cstaskmanager "github.com/Layr-Labs/eigensdk-go/examples/bindings/taskManager"
	taskgeneratorexample "github.com/Layr-Labs/eigensdk-go/examples/task-generator"
)

func main() {
	logger, err := logging.NewZapLogger(logging.Production) // Change here if want to change logging level
	if err != nil {
		println("Failure creating logger")
		return
	}

	taskManagerAbi, err := cstaskmanager.ContractIncredibleSquaringTaskManagerMetaData.GetAbi()
	if err != nil {
		logger.Fatalf(err.Error())
	}

	newTaskEventHash := taskManagerAbi.Events["NewTaskCreated"].ID
	taskRespondedEventHash := taskManagerAbi.Events["TaskResponded"].ID

	cfg := challenger.ChallengerConfig{
		EthWsUrl: "ws://localhost:8545",
		Logger:   logger,
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

	avsConfig := AvsConfig{
		Logger:                        logger,
		IncredibleSquaringTaskManager: common.HexToAddress("0x2bdcc0de6be1f7d2ee689a0342d76f52e8efaba3"),
		TxMgr:                         txMgr,
		EthHttpClient:                 ethHttpClient,
	}
	challengerLogicImpl, err := NewChallengerLogicImpl(&avsConfig)
	if err != nil {
		logger.Errorf("Failed to create challenger logic from config: %v", err)
		return
	}

	challenger, err := challenger.NewChallenger[*big.Int, NewTaskCreatedEvent, TaskRespondedEvent](
		cfg, 
		challengerLogicImpl, 
		newTaskEventHash, 
		taskRespondedEventHash, 
		taskManagerAbi, 
		ethHttpClient,
	)
	if err != nil {
		logger.Errorf("Failed to create challenger from config: %v", err)
		return
	}

	err = challenger.Start(context.Background())
	if err != nil {
		logger.Errorf("Error while running operator: %v", err)
		return
	}
}

// Required structs

type NewTaskCreatedEvent struct {
	TaskIndex uint32
	Task      challenger.GenericInputTask[*big.Int]
	Raw       types.Log
}

func (newTaskEvent NewTaskCreatedEvent) InnerTask() challenger.GenericInputTask[*big.Int] {
	return newTaskEvent.Task
}

type TaskRespondedEvent struct {
	TaskResponse              challenger.GenericInputTaskResponse[*big.Int]
	TaskResponseMetadata      challenger.GenericTaskResponseMetadata
	NonSigningOperatorPubKeys []challenger.BN254G1Point
}

func (taskRespEvent TaskRespondedEvent) TaskIndex() uint32 {
	return taskRespEvent.TaskResponse.ReferenceTaskIndex
}

func (taskRespEvent TaskRespondedEvent) GetTaskResponse() challenger.GenericInputTaskResponse[*big.Int] {
	return taskRespEvent.TaskResponse
}

func (taskRespEvent TaskRespondedEvent) GetTaskResponseMetadata() challenger.GenericTaskResponseMetadata {
	return taskRespEvent.TaskResponseMetadata
}


// Challenger Logic

type ChallengerLogicImpl struct {
	logger    logging.Logger
	ethClient *ethclient.Client
	avsWriter *AvsWriter
}

var _ challenger.ChallengerLogic[*big.Int] = (*ChallengerLogicImpl)(nil)

type TaskResponseData struct {
	TaskResponse              cstaskmanager.IIncredibleSquaringTaskManagerTaskResponse
	TaskResponseMetadata      cstaskmanager.IIncredibleSquaringTaskManagerTaskResponseMetadata
	NonSigningOperatorPubKeys []cstaskmanager.BN254G1Point
}

func NewChallengerLogicImpl(c *AvsConfig) (*ChallengerLogicImpl, error) {
	avsWriter, err := BuildAvsWriterFromConfig(c)
	if err != nil {
		c.Logger.Errorf("Cannot create avsWriter", "err", err)
		return nil, err
	}

	return &ChallengerLogicImpl{
		logger:    c.Logger,
		ethClient: c.EthHttpClient,
		avsWriter: avsWriter,
	}, nil
}

func (c *ChallengerLogicImpl) VerifyChallenge(
	taskIndex uint32,
	task challenger.GenericInputTask[*big.Int],
	responseData challenger.TaskResponseData[*big.Int],
) error {
	nonSignerPubkeys := []cstaskmanager.BN254G1Point{}
	for i, pubkey := range responseData.NonSigningOperatorPubKeys {
		nonSignerPubkeys[i] = cstaskmanager.BN254G1Point{
			X: pubkey.X,
			Y: pubkey.Y,
		}
	}

	numberToBeSquared := task.InputValue
	answerInResponse := responseData.TaskResponse.InputValue
	trueAnswer := numberToBeSquared.Exp(numberToBeSquared, big.NewInt(2), nil)

	// checking if the answer in the response submitted by aggregator is correct
	if trueAnswer.Cmp(answerInResponse) != 0 {
		c.logger.Info("The number squared is not correct", "expectedAnswer", trueAnswer, "gotAnswer", answerInResponse)

		// raise challenge
		c.logger.Info("Challenger raising challenge.", "taskIndex", taskIndex)

		// This conversions are not optimal, but are necessary to send the challenge to the contract
		incredibleSquaringTask := cstaskmanager.IIncredibleSquaringTaskManagerTask{
			NumberToBeSquared:         task.InputValue,
			TaskCreatedBlock:          task.TaskCreatedBlock,
			QuorumNumbers:             task.QuorumNumbers,
			QuorumThresholdPercentage: task.QuorumThresholdPercentage,
		}

		incredibleSquaringTaskResponse := cstaskmanager.IIncredibleSquaringTaskManagerTaskResponse{
			ReferenceTaskIndex: responseData.TaskResponse.ReferenceTaskIndex,
			NumberSquared:      responseData.TaskResponse.InputValue,
		}

		incredibleSquaringTaskResponseMetadata := cstaskmanager.IIncredibleSquaringTaskManagerTaskResponseMetadata{
			TaskRespondedBlock: responseData.TaskResponseMetadata.TaskRespondedBlock,
			HashOfNonSigners:   responseData.TaskResponseMetadata.HashOfNonSigners,
		}

		_, err := c.avsWriter.RaiseChallenge(
			context.Background(),
			incredibleSquaringTask,
			incredibleSquaringTaskResponse,
			incredibleSquaringTaskResponseMetadata,
			nonSignerPubkeys,
		)
		if err != nil {
			c.logger.Error("Challenger failed to raise challenge:", "err", err)
			return fmt.Errorf("challenger failed to raise challenge: %w", err)
		}

		return nil
	} else {
		c.logger.Info("The number squared is correct")
		return nil
	}
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

func (w *AvsWriter) RaiseChallenge(
	ctx context.Context,
	task cstaskmanager.IIncredibleSquaringTaskManagerTask,
	taskResponse cstaskmanager.IIncredibleSquaringTaskManagerTaskResponse,
	taskResponseMetadata cstaskmanager.IIncredibleSquaringTaskManagerTaskResponseMetadata,
	pubkeysOfNonSigningOperators []cstaskmanager.BN254G1Point,
) (*types.Receipt, error) {
	txOpts, err := w.TxMgr.GetNoSendTxOpts()
	if err != nil {
		w.logger.Errorf("Error getting tx opts")
		return nil, fmt.Errorf("error getting tx opts: %w", err)
	}
	tx, err := w.taskManagerContract.RaiseAndResolveChallenge(
		txOpts,
		task,
		taskResponse,
		taskResponseMetadata,
		pubkeysOfNonSigningOperators,
	)
	if err != nil {
		w.logger.Errorf("Error assembling RaiseChallenge tx")
		return nil, fmt.Errorf("error assembling RaiseChallenge tx: %w", err)
	}
	receipt, err := w.TxMgr.Send(ctx, tx, true)
	if err != nil {
		w.logger.Errorf("Error submitting RaiseChallenge tx")
		return nil, fmt.Errorf("error submitting RaiseChallenge tx: %w", err)
	}
	return receipt, nil
}

func (w *AvsWriter) ParseTaskResponded(
	rawLog types.Log,
) (*cstaskmanager.ContractIncredibleSquaringTaskManagerTaskResponded, error) {
	return w.taskManagerContract.ContractIncredibleSquaringTaskManagerFilterer.ParseTaskResponded(rawLog)
}
