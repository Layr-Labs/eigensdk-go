package challengerexample

import (
	"context"
	"errors"
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

	taskgeneratorexample "github.com/Layr-Labs/eigensdk-go/examples/task-generator"
	cstaskmanager "github.com/Layr-Labs/eigensdk-go/examples/task-generator/bindings/taskManager"
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

	challenger, err := challenger.NewChallenger(cfg, challengerLogicImpl, newTaskEventHash, taskRespondedEventHash)
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

type ChallengerLogicImpl struct {
	logger        logging.Logger
	ethClient     *ethclient.Client
	avsWriter     AvsWriter
	tasks         map[uint32]cstaskmanager.IIncredibleSquaringTaskManagerTask
	taskResponses map[uint32]TaskResponseData
}

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
		logger:        c.Logger,
		ethClient:     c.EthHttpClient,
		avsWriter:     *avsWriter,
		tasks:         make(map[uint32]cstaskmanager.IIncredibleSquaringTaskManagerTask),
		taskResponses: make(map[uint32]TaskResponseData),
	}, nil
}

func (c *ChallengerLogicImpl) verifyChallenge(taskIndex uint32) error {
	numberToBeSquared := c.tasks[taskIndex].NumberToBeSquared
	answerInResponse := c.taskResponses[taskIndex].TaskResponse.NumberSquared
	trueAnswer := numberToBeSquared.Exp(numberToBeSquared, big.NewInt(2), nil)

	// checking if the answer in the response submitted by aggregator is correct
	if trueAnswer.Cmp(answerInResponse) != 0 {
		c.logger.Info("The number squared is not correct", "expectedAnswer", trueAnswer, "gotAnswer", answerInResponse)

		// raise challenge
		c.logger.Info("Challenger raising challenge.", "taskIndex", taskIndex)

		_, err := c.avsWriter.RaiseChallenge(
			context.Background(),
			c.tasks[taskIndex],
			c.taskResponses[taskIndex].TaskResponse,
			c.taskResponses[taskIndex].TaskResponseMetadata,
			c.taskResponses[taskIndex].NonSigningOperatorPubKeys,
		)
		if err != nil {
			c.logger.Error("Challenger failed to raise challenge:", "err", err)
			return fmt.Errorf("challenger failed to raise challenge: %w", err)
		}

		return nil
	} else {
		c.logger.Info("The number squared is correct")
		return errors.New("100. Task response is valid")
	}
}

func (c *ChallengerLogicImpl) ProcessNewTaskCreatedLog(
	newTaskCreatedEvent any,
) error {
	var newTaskCreatedLog cstaskmanager.ContractIncredibleSquaringTaskManagerNewTaskCreated

	log, ok := newTaskCreatedEvent.(types.Log)
	if !ok {
		c.logger.Errorf("Event was not a types.Log. Event: %v", newTaskCreatedEvent)
		return errors.New("invalid type event, expected types.Log")
	}

	taskManagerAbi, err := cstaskmanager.ContractIncredibleSquaringTaskManagerMetaData.GetAbi()
	if err != nil {
		c.logger.Fatalf("Error obtaining task manager ABI: %v", err)
	}

	err = taskManagerAbi.UnpackIntoInterface(&newTaskCreatedLog, "NewTaskCreated", log.Data)
	if err != nil {
		return fmt.Errorf("error unpacking the log: %w", err)
	}

	taskIndex := newTaskCreatedLog.TaskIndex
	c.tasks[taskIndex] = newTaskCreatedLog.Task

	if _, found := c.taskResponses[taskIndex]; found {
		_ = c.verifyChallenge(taskIndex)
	}

	return nil
}

func (c *ChallengerLogicImpl) ProcessTaskResponseLog(
	taskResponseEvent any,
) error {
	var taskRespondedLog cstaskmanager.ContractIncredibleSquaringTaskManagerTaskResponded

	log, ok := taskResponseEvent.(types.Log)
	if !ok {
		c.logger.Errorf("Event was not a types.Log. Event: %v", taskResponseEvent)
		return errors.New("invalid type event, expected types.Log")
	}

	taskManagerAbi, err := cstaskmanager.ContractIncredibleSquaringTaskManagerMetaData.GetAbi()
	if err != nil {
		c.logger.Fatalf("Error obtaining task manager ABI: %v", err)
	}

	err = taskManagerAbi.UnpackIntoInterface(&taskRespondedLog, "TaskResponded", log.Data)
	if err != nil {
		return fmt.Errorf("error unpacking the log: %w", err)
	}

	taskIndex := taskRespondedLog.TaskResponse.ReferenceTaskIndex

	// get the inputs necessary for raising a challenge
	nonSigningOperatorPubKeys := c.getNonSigningOperatorPubKeys(&taskRespondedLog)
	taskResponseData := TaskResponseData{
		TaskResponse:              taskRespondedLog.TaskResponse,
		TaskResponseMetadata:      taskRespondedLog.TaskResponseMetadata,
		NonSigningOperatorPubKeys: nonSigningOperatorPubKeys,
	}

	c.taskResponses[taskIndex] = taskResponseData

	if _, found := c.tasks[taskIndex]; found {
		_ = c.verifyChallenge(taskIndex)
	}

	return nil
}

func (c *ChallengerLogicImpl) getNonSigningOperatorPubKeys(
	vLog *cstaskmanager.ContractIncredibleSquaringTaskManagerTaskResponded,
) []cstaskmanager.BN254G1Point {
	// get the nonSignerStakesAndSignature
	txHash := vLog.Raw.TxHash
	tx, _, err := c.ethClient.TransactionByHash(context.Background(), txHash)
	if err != nil {
		c.logger.Error("Error getting transaction by hash",
			"txHash", txHash,
			"err", err,
		)
	}

	taskManagerAbi, err := cstaskmanager.ContractIncredibleSquaringTaskManagerMetaData.GetAbi()
	if err != nil {
		c.logger.Error("Error getting Abi", "err", err)
	}

	calldata := tx.Data()
	methodSig := calldata[:4]
	method, err := taskManagerAbi.MethodById(methodSig)
	if err != nil {
		c.logger.Error("Error getting method", "err", err)
	}

	inputs, err := method.Inputs.Unpack(calldata[4:])
	if err != nil {
		c.logger.Error("Error unpacking calldata", "err", err)
	}

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
		[]cstaskmanager.BN254G1Point,
		len(nonSignerStakesAndSignatureInput.NonSignerPubkeys),
	)
	for i, pubkey := range nonSignerStakesAndSignatureInput.NonSignerPubkeys {
		nonSigningOperatorPubKeys[i] = cstaskmanager.BN254G1Point{
			X: pubkey.X,
			Y: pubkey.Y,
		}
	}

	return nonSigningOperatorPubKeys
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
