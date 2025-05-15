package main

import (
	"context"

	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	"github.com/Layr-Labs/eigensdk-go/challenger"
	challengerprocessor "github.com/Layr-Labs/eigensdk-go/challenger/challenger-processor"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/utils"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	examplecommon "github.com/Layr-Labs/eigensdk-go/examples/awesome-vault-service/common"
	avtaskmanager "github.com/Layr-Labs/eigensdk-go/examples/awesome-vault-service/contracts/bindings/AwesomeVaultTaskManager"
)

func main() {
	logger, err := logging.NewZapLogger(logging.Production)
	if err != nil {
		println("Failure creating logger")
		return
	}

	taskManagerAbi, err := avtaskmanager.ContractAwesomeVaultTaskManagerMetaData.GetAbi()
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

	vaultServiceResponseCalc := examplecommon.NewVaultServiceResponseCalculator()

	challengeRaiser, err := examplecommon.NewAwesomeVaultTaskManager(taskManagerAddr, taskManagerAbi, txMgr, ethClient)
	if err != nil {
		logger.Errorf("Failed to create challenge raiser: %v", err)
		return
	}

	vaultSetValidationWithProof := getProofGeneratingVerifier(vaultServiceResponseCalc)

	challengerProcessor, err := challengerprocessor.NewIndexingChallengerProcessor(logger, vaultSetValidationWithProof, challengeRaiser)
	if err != nil {
		logger.Errorf("Failed to create challenger verifier: %w", err)
		return
	}

	challengerConfig := challenger.Config{
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

func getProofGeneratingVerifier(vaultServiceResponseCalc *examplecommon.VaultServiceResponseCalculator) challengerprocessor.ResponseValidationFunction[examplecommon.TaskInput, [32]byte, []examplecommon.TaskInput] {
	vaultSetValidation := challengerprocessor.ResponseValidationFunctionFromResponseCalculator(vaultServiceResponseCalc, func(a, b [32]byte) bool { return a == b })

	return func(taskIndex uint32, input examplecommon.TaskInput, output [32]byte) (bool, []examplecommon.TaskInput, error) {
		var proof []examplecommon.TaskInput
		shouldRaise, _, err := vaultSetValidation(taskIndex, input, output)
		if err != nil {
			return shouldRaise, proof, err
		}
		if shouldRaise {
			oldLeaves, err := vaultServiceResponseCalc.GetPreviousState(input)
			if err != nil {
				return false, nil, utils.WrapError("Error getting previous state", err)
			}
			proof = oldLeaves
		}
		return shouldRaise, proof, nil
	}
}
