package main

import (
	"context"
	"fmt"
	"math/big"

	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	"github.com/Layr-Labs/eigensdk-go/challenger"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/testutils"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
	"github.com/Layr-Labs/eigensdk-go/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	cstaskmanager "github.com/Layr-Labs/eigensdk-go/examples/incredible-squaring/bindings/taskManager"
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

	ethHttpUrl := "http://localhost:8545"
	ethHttpClient, err := ethclient.Dial(ethHttpUrl)
	if err != nil {
		return
	}

	cfg := challenger.ChallengerConfig{
		EthWsUrl:       "ws://localhost:8545",
		Logger:         logger,
		TaskManagerAbi: taskManagerAbi,
		EthClient:      ethHttpClient,
	}

	ecdsaPrivateKey, err := crypto.HexToECDSA(testutils.ANVIL_FIRST_PRIVATE_KEY)
	if err != nil {
		logger.Errorf("Cannot parse ecdsa private key", "err", err)
		return
	}

	txMgr, err := txmgr.NewSimpleTxManagerFromPrivateKey(logger, ethHttpClient, ecdsaPrivateKey)
	if err != nil {
		logger.Errorf("Failed to create transaction manager", "err", err)
		return
	}

	avsConfig := AvsConfig{
		Logger:                        logger,
		IncredibleSquaringTaskManager: common.HexToAddress("0x2bdcc0de6be1f7d2ee689a0342d76f52e8efaba3"),
		TxMgr:                         txMgr,
		EthHttpClient:                 ethHttpClient,
	}
	challengeVerifierImpl, err := NewChallengeVerifierImpl(&avsConfig)
	if err != nil {
		logger.Errorf("Failed to create challenger logic from config: %v", err)
		return
	}

	challenger, err := challenger.NewChallenger(
		cfg,
		challengeVerifierImpl,
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

// Challenger Logic

type ChallengeVerifierImpl struct {
	logger    logging.Logger
	ethClient *ethclient.Client
	avsWriter *AvsWriter
}

var _ challenger.ChallengeVerifier[*big.Int, *big.Int] = (*ChallengeVerifierImpl)(nil)

type TaskResponseData struct {
	TaskResponse              cstaskmanager.IIncredibleSquaringTaskManagerTaskResponse
	TaskResponseMetadata      cstaskmanager.IIncredibleSquaringTaskManagerTaskResponseMetadata
	NonSigningOperatorPubKeys []cstaskmanager.BN254G1Point
}

func NewChallengeVerifierImpl(c *AvsConfig) (*ChallengeVerifierImpl, error) {
	avsWriter, err := BuildAvsWriterFromConfig(c)
	if err != nil {
		c.Logger.Errorf("Cannot create avsWriter", "err", err)
		return nil, err
	}

	return &ChallengeVerifierImpl{
		logger:    c.Logger,
		ethClient: c.EthHttpClient,
		avsWriter: avsWriter,
	}, nil
}

func (c *ChallengeVerifierImpl) VerifyChallenge(
	taskIndex uint32,
	task sdktypes.GenericInputTask[*big.Int],
	responseData sdktypes.TaskResponseData[*big.Int],
) error {
	nonSignerPubkeys := []cstaskmanager.BN254G1Point{}
	for i, pubkey := range responseData.NonSigningOperatorPubKeys {
		nonSignerPubkeys[i] = cstaskmanager.BN254G1Point{
			X: pubkey.X,
			Y: pubkey.Y,
		}
	}

	numberToBeSquared := task.InputValue
	answerInResponse := responseData.TaskResponse.OutputValue
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
			NumberSquared:      responseData.TaskResponse.OutputValue,
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
