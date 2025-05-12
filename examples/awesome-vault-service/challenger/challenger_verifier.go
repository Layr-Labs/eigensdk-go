package challenger

import (
	"context"

	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	"github.com/Layr-Labs/eigensdk-go/challenger"
	"github.com/Layr-Labs/eigensdk-go/logging"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
	"github.com/Layr-Labs/eigensdk-go/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	delegationmanager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/DelegationManager"
	examplecommon "github.com/Layr-Labs/eigensdk-go/examples/awesome-vault-service/common"
	taskmanager "github.com/Layr-Labs/eigensdk-go/examples/awesome-vault-service/contracts/bindings/AwesomeVaultTaskManager"
)

type ChallengeVerifier struct {
	logger                    logging.Logger
	taskManagerContract       *taskmanager.ContractAwesomeVaultTaskManager
	txMgr                     txmgr.TxManager
	delegationManagerContract *delegationmanager.ContractDelegationManager
}

func NewChallengeVerifier(
	logger logging.Logger,
	taskManagerAddr common.Address,
	ethclient ethclient.Client,
	txMgr txmgr.TxManager,
	delegationManagerAddr common.Address,
) (ChallengeVerifier, error) {
	// Get the task manager contract with the given address
	taskManagerContract, err := taskmanager.NewContractAwesomeVaultTaskManager(taskManagerAddr, &ethclient)
	if err != nil {
		logger.Errorf("Failed to get Task Manager Contract: %w", err)
		return ChallengeVerifier{}, err
	}

	delegationManagerContract, err := delegationmanager.NewContractDelegationManager(delegationManagerAddr, &ethclient)
	if err != nil {
		logger.Errorf("Failed to get delegation Manager Contract: %w", err)
		return ChallengeVerifier{}, err
	}

	return ChallengeVerifier{
		logger:                    logger,
		taskManagerContract:       taskManagerContract,
		txMgr:                     txMgr,
		delegationManagerContract: delegationManagerContract,
	}, nil
}

var _ challenger.ChallengeVerifier[examplecommon.TaskInput, [32]byte] = (*ChallengeVerifier)(nil)

func (cv ChallengeVerifier) VerifyChallenge(taskIndex uint32, task sdktypes.GenericInputTask[examplecommon.TaskInput], taskResponse sdktypes.TaskResponseData[[32]byte]) error {
	// Calculate response
	result, err := examplecommon.VaultSet(taskIndex, task.InputValue)
	if err != nil {
		cv.logger.Errorf("Failed to calculate task response: %w", err)
		return err
	}

	// Compare submitted response with calculated here
	receivedResponse := taskResponse.TaskResponse.OutputValue
	// TODO: compare
	shouldRaiseChallenge := result == [32]byte{0}

	if shouldRaiseChallenge {
		cv.logger.Infof("Response was not correct, expected %v and got %v", result, receivedResponse)
		// Call the Raise Challenge method of the on-chain contract

		noSendTxOpts, err := cv.txMgr.GetNoSendTxOpts()
		if err != nil {
			cv.logger.Errorf("Failed to get tx opts: %w", err)
			return err
		}

		taskValue := taskmanager.IAwesomeVaultTaskManagerTask{
			Input:                     taskmanager.IAwesomeVaultTaskManagerTaskInput{Key: task.InputValue.Key, Value: task.InputValue.Value},
			TaskCreatedBlock:          task.TaskCreatedBlock,
			QuorumNumbers:             task.QuorumNumbers,
			QuorumThresholdPercentage: task.QuorumThresholdPercentage,
		}

		taskResponseValue := taskmanager.IAwesomeVaultTaskManagerTaskResponse{
			ReferenceTaskIndex: taskResponse.TaskResponse.ReferenceTaskIndex,
			Result:             taskResponse.TaskResponse.OutputValue,
		}

		taskResponseMetadataValue := taskmanager.IAwesomeVaultTaskManagerTaskResponseMetadata{
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

		tx, err := cv.taskManagerContract.RaiseAndResolveChallenge(noSendTxOpts, taskValue, taskResponseValue, taskResponseMetadataValue, pubkeysOfNonSigningOperators)
		if err != nil {
			cv.logger.Errorf("Failed to create raise and resolve challenge tx: %w", err)
			return err
		}

		receipt, err := cv.txMgr.Send(context.Background(), tx, true)
		if err != nil {
			cv.logger.Errorf("Failed to send raise and resolve challenge tx: %w", err)
			return err
		}
		if receipt.Status != types.ReceiptStatusSuccessful {
			cv.logger.Error("receipt status was not success sending raise challenge tx")
			return utils.WrapError(err, "receipt status was not success")
		}

		// Note: this printing logic is not necessary, but left here to show how the operator shares
		// at strategy decreases after raising a challenge (because it's being slashed)
		shares, err := cv.delegationManagerContract.OperatorShares(
			&bind.CallOpts{},
			common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"), // Operator address
			common.HexToAddress("0x2b961e3959b79326a8e7f64ef0d2d825707669b5"), // strategy address
		)
		if err != nil {
			cv.logger.Errorf("Failed to get operator shares. Err: %w", err)
			return err
		}

		cv.logger.Infof("After raising challenge, operator shares are %v", shares)
	}

	return nil
}
