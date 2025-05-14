package main

import (
	"context"

	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	"github.com/Layr-Labs/eigensdk-go/challenger"
	challengerprocessor "github.com/Layr-Labs/eigensdk-go/challenger/challenger-processor"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/utils"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	sdktypes "github.com/Layr-Labs/eigensdk-go/types"

	examplecommon "github.com/Layr-Labs/eigensdk-go/examples/awesome-vault-service/common"
	taskmanager "github.com/Layr-Labs/eigensdk-go/examples/awesome-vault-service/contracts/bindings/AwesomeVaultTaskManager"
)

func main() {
	logger, err := logging.NewZapLogger(logging.Production)
	if err != nil {
		println("Failure creating logger")
		return
	}

	taskManagerAbi, err := taskmanager.ContractAwesomeVaultTaskManagerMetaData.GetAbi()
	if err != nil {
		logger.Errorf("Failed to get task manager abi: %w", err)
		return
	}

	ethHttpUrl := "http://localhost:8545"
	ethClient, err := ethclient.Dial(ethHttpUrl)
	if err != nil {
		logger.Errorf("Failed to dial ethclient: %w", err)
		return
	}

	taskManagerAddr := gethcommon.HexToAddress("0x7bc06c482dead17c0e297afbc32f6e63d3846650")

	challengerPrivateKey := "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	ecdsaPrivateKey, err := crypto.HexToECDSA(challengerPrivateKey)
	if err != nil {
		logger.Errorf("Failed to create ecdsa private key: %w", err)
		return
	}

	txMgr, err := txmgr.NewSimpleTxManagerFromPrivateKey(logger, ethClient, ecdsaPrivateKey)
	if err != nil {
		logger.Errorf("Failed to create tx manager from private key: %w", err)
		return
	}

	challengeRaiser, err := NewChallengeRaiser(taskManagerAddr, ethClient, txMgr)
	if err != nil {
		logger.Errorf("Failed to create challenger raiser: %w", err)
		return
	}

	vaultServiceResponseCalc := examplecommon.NewVaultServiceResponseCalculator()

	vaultSetValidation := challengerprocessor.ResponseValidationFunctionFromResponseCalculator(vaultServiceResponseCalc, func(a, b [32]byte) bool { return a == b })

	challengerProcessor, err := challengerprocessor.NewIndexingChallengerProcessor(logger, vaultSetValidation, challengeRaiser)
	if err != nil {
		logger.Errorf("Failed to create challenger verifier: %w", err)
		return
	}

	challengerConfig := challenger.ChallengerConfig{
		Logger:         logger,
		TaskManagerAbi: taskManagerAbi,
		EthClient:      ethClient,
		EthWsUrl:       "ws://localhost:8545",
	}
	challenger, err := challenger.NewChallenger(challengerConfig, challengerProcessor)
	if err != nil {
		logger.Errorf("Failed to create challenger: %w", err)
		return
	}

	err = challenger.Start(context.Background())
	if err != nil {
		logger.Errorf("Failure while running challenger: %w", err)
		return
	}
}

type ChallengeRaiser struct {
	taskManager *taskmanager.ContractAwesomeVaultTaskManager
	txMgr       txmgr.TxManager
}

var _ challengerprocessor.ChallengeRaiser[examplecommon.TaskInput, [32]byte] = (*ChallengeRaiser)(nil)

func NewChallengeRaiser(address gethcommon.Address, ethClient *ethclient.Client, txMgr txmgr.TxManager) (*ChallengeRaiser, error) {
	tm, err := taskmanager.NewContractAwesomeVaultTaskManager(address, ethClient)
	if err != nil {
		return nil, err
	}

	return &ChallengeRaiser{
		taskManager: tm,
		txMgr:       txMgr,
	}, nil
}

func (cr *ChallengeRaiser) RaiseChallenge(task sdktypes.GenericInputTask[examplecommon.TaskInput], taskResponse sdktypes.GenericOutputTaskResponse[[32]byte], taskResponseMetadata sdktypes.GenericTaskResponseMetadata, nonSigningOperatorPubKeys []sdktypes.BN254G1Point) error {
	txOpts, err := cr.txMgr.GetNoSendTxOpts()
	if err != nil {
		return utils.WrapError("Error getting tx opts", err)
	}
	taskInput := taskmanager.IAwesomeVaultTaskManagerTaskInput{
		Key:   task.InputValue.Key,
		Value: task.InputValue.Value,
	}
	contractTask := taskmanager.IAwesomeVaultTaskManagerTask{
		Input:                     taskInput,
		TaskCreatedBlock:          task.TaskCreatedBlock,
		QuorumNumbers:             task.QuorumNumbers,
		QuorumThresholdPercentage: task.QuorumThresholdPercentage,
	}
	contractTaskResponse := taskmanager.IAwesomeVaultTaskManagerTaskResponse{
		ReferenceTaskIndex: taskResponse.ReferenceTaskIndex,
		Result:             taskResponse.OutputValue,
	}
	contractTaskResponseMetadata := taskmanager.IAwesomeVaultTaskManagerTaskResponseMetadata{
		TaskRespondedBlock: taskResponseMetadata.TaskRespondedBlock,
		HashOfNonSigners:   taskResponseMetadata.HashOfNonSigners,
	}
	contractNonSigningOperatorPubKeys := make([]taskmanager.BN254G1Point, len(nonSigningOperatorPubKeys))
	for i, pubKey := range nonSigningOperatorPubKeys {
		contractNonSigningOperatorPubKeys[i] = taskmanager.BN254G1Point{
			X: pubKey.X,
			Y: pubKey.Y,
		}
	}
	tx, err := cr.taskManager.RaiseAndResolveChallenge(txOpts, contractTask, contractTaskResponse, contractTaskResponseMetadata, contractNonSigningOperatorPubKeys)
	if err != nil {
		return utils.WrapError("Error raising challenge", err)
	}
	receipt, err := cr.txMgr.Send(context.Background(), tx, true)
	if err != nil {
		return utils.WrapError("Error submitting RaiseChallenge tx", err)
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		return utils.WrapError("RaiseChallenge tx reverted", nil)
	}
	return nil
}
