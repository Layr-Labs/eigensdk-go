package challenger

import (
	"context"
	"errors"
	"math/big"

	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	"github.com/Layr-Labs/eigensdk-go/challenger"
	"github.com/Layr-Labs/eigensdk-go/logging"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	taskmanager "github.com/Layr-Labs/eigensdk-go/examples/incredible-dot-product/contracts/bindings/IncredibleDotProductTaskManager"
)

type DotProductInput struct {
	X []*big.Int
	Y []*big.Int
}

type DotProductOutput struct {
	Result *big.Int
}

type ChallengeVerifier struct {
	logger              logging.Logger
	taskManagerContract *taskmanager.ContractIncredibleDotProductTaskManager
	txMgr               txmgr.TxManager
}

func NewChallengeVerifier(
	logger logging.Logger,
	taskManagerAddr common.Address,
	ethclient ethclient.Client,
	txMgr txmgr.TxManager,
) (ChallengeVerifier, error) {
	// Get the task manager contract with the given address
	taskManagerContract, err := taskmanager.NewContractIncredibleDotProductTaskManager(taskManagerAddr, &ethclient)
	if err != nil {
		logger.Errorf("Failed to get Task Manager Contract: %w", err)
		return ChallengeVerifier{}, err
	}
	return ChallengeVerifier{
		logger:              logger,
		taskManagerContract: taskManagerContract,
		txMgr:               txMgr,
	}, nil
}

var _ challenger.ChallengeVerifier[DotProductInput, DotProductOutput] = (*ChallengeVerifier)(nil)

func (cv ChallengeVerifier) VerifyChallenge(taskIndex uint32, task sdktypes.GenericInputTask[DotProductInput], taskResponse sdktypes.TaskResponseData[DotProductOutput]) error {
	// Calculate response
	totalSum := big.NewInt(0)
	for i := range task.InputValue.X {
		currentSum := big.NewInt(0).Mul(task.InputValue.X[i], task.InputValue.Y[i])
		totalSum.Add(totalSum, currentSum)
	}

	// Compare submitted response with calculated here
	receivedResponse := taskResponse.TaskResponse.OutputValue.Result
	shouldRaiseChallenge := totalSum == receivedResponse

	if shouldRaiseChallenge {
		cv.logger.Infof("Response was not correct, expected %v and got %v", totalSum, receivedResponse)
		// Call the Raise Challenge method of the on-chain contract

		noSendTxOpts, err := cv.txMgr.GetNoSendTxOpts()
		if err != nil {
			cv.logger.Errorf("Failed to get tx opts: %w", err)
			return err
		}

		incredibleTask := taskmanager.IIncredibleDotProductTaskManagerTask{
			PointsToMultiply:          taskmanager.IIncredibleDotProductTaskManagerDotProductInput{XAxis: task.InputValue.X, YAxis: task.InputValue.Y},
			TaskCreatedBlock:          task.TaskCreatedBlock,
			QuorumNumbers:             task.QuorumNumbers,
			QuorumThresholdPercentage: task.QuorumThresholdPercentage,
		}

		incredibleTaskResponse := taskmanager.IIncredibleDotProductTaskManagerTaskResponse{
			ReferenceTaskIndex: taskResponse.TaskResponse.ReferenceTaskIndex,
			Result:             taskResponse.TaskResponse.OutputValue.Result,
		}

		incredibleTaskResponseMetadata := taskmanager.IIncredibleDotProductTaskManagerTaskResponseMetadata{
			TaskRespondedBlock: taskResponse.TaskResponseMetadata.TaskRespondedBlock,
			HashOfNonSigners:   taskResponse.TaskResponseMetadata.HashOfNonSigners,
		}

		pubkeysOfNonSigningOperators := []taskmanager.BN254G1Point{}
		for i, pubkey := range taskResponse.NonSigningOperatorPubKeys {
			pubkeysOfNonSigningOperators[i] = taskmanager.BN254G1Point{
				X: pubkey.X,
				Y: pubkey.Y,
			}
		}

		tx, err := cv.taskManagerContract.RaiseAndResolveChallenge(noSendTxOpts, incredibleTask, incredibleTaskResponse, incredibleTaskResponseMetadata, pubkeysOfNonSigningOperators)
		if err != nil {
			cv.logger.Errorf("Failed to create raise and resolve challenge tx: %w", err)
			return err
		}

		receipt, err := cv.txMgr.Send(context.Background(), tx, true)
		if err != nil {
			cv.logger.Errorf("Failed to send raise and resolve challenge tx: %w", err)
			return err
		}
		if receipt.Status != uint64(1) {
			return errors.New("failed to send raise and resolve challenge tx: receipt status was not success")
		}

	}

	return nil
}
