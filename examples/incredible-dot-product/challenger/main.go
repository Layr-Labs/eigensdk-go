package main

import (
	"context"
	"math/big"

	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	"github.com/Layr-Labs/eigensdk-go/challenger"
	challengerprocessor "github.com/Layr-Labs/eigensdk-go/challenger/challenger-processor"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/operator"
	taskmanager "github.com/Layr-Labs/eigensdk-go/task-manager"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	common "github.com/Layr-Labs/eigensdk-go/examples/common"
	examplecommon "github.com/Layr-Labs/eigensdk-go/examples/incredible-dot-product/common"
	idptaskmanager "github.com/Layr-Labs/eigensdk-go/examples/incredible-dot-product/contracts/bindings/IncredibleDotProductTaskManager"
)

func main() {
	// 0. Create the logger where all the logs will appear
	logger, err := logging.NewZapLogger(logging.Production) // Change here if want to change logging level
	if err != nil {
		println("Failure creating logger")
		return
	}

	// 1. Create the config passed to the challenger

	// Get the ABI of the task manager contract's binding
	taskManagerAbi, err := idptaskmanager.ContractIncredibleDotProductTaskManagerMetaData.GetAbi()
	if err != nil {
		logger.Fatalf(err.Error())
	}

	// Create the ethereum client that will send the RPC messages to the node
	ethClient, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		logger.Errorf("Failed to dial ethclient: %w", err)
		return
	}

	challengerConfig := challenger.Config{
		EthWsUrl:       "ws://localhost:8545",
		Logger:         logger,
		TaskManagerAbi: taskManagerAbi,
		EthClient:      ethClient,
	}

	// 2. Define a function that verifies the response for a task. Note that here we wrap the logic into a
	// `ResponseCalculator` implementation, and then we use the ResponseValidationFunctionFromResponseCalculator,
	// which creates a validation function that will compute the logic function to validate the response
	dotProductCalculator := operator.NewFunctionResponseCalculator(examplecommon.DotProduct)
	dotProductValidation := challengerprocessor.ResponseValidationFunctionFromResponseCalculator(dotProductCalculator, common.BigIntEqual)

	// 3. Provide a struct implementing the `ChallengerProcessor` interface, first instantiating the values
	// required for creating it

	// Create the transaction manager, that will manage the transaction sending
	// This is the anvil first rich account's private key
	ecdsaPrivateKey, err := crypto.HexToECDSA("ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80")
	if err != nil {
		logger.Errorf("Failed to create ecdsa private key: %w", err)
		return
	}

	txMgr, err := txmgr.NewSimpleTxManagerFromPrivateKey(logger, ethClient, ecdsaPrivateKey)
	if err != nil {
		logger.Errorf("Failed to create transaction manager", "err", err)
		return
	}

	// Provide a struct that implements the ChallengeRaiser interface. Here we provide a SDK implementation that
	// satisfies the ChallengeRaiser interface.
	// Note that in this step we define the input and output types that we are using on our AVS. In this case
	// the DotProductInput struct (a pair of vectors) and a big int.
	challengeRaiser, err := taskmanager.NewTaskManagerFromAbi[examplecommon.DotProductInput, *big.Int](
		gethcommon.HexToAddress("0x2bdcc0de6be1f7d2ee689a0342d76f52e8efaba3"),
		taskManagerAbi,
		txMgr,
		ethClient,
	)
	if err != nil {
		logger.Errorf("Failed to create challenger raiser: %v", err)
		return
	}

	// 4. Create the Challenger Processor, with the challenger raiser created above.
	challengerProcessor, err := challengerprocessor.NewIndexingChallengerProcessor(logger, dotProductValidation, challengeRaiser)
	if err != nil {
		logger.Errorf("Failed to create challenger processor: %v", err)
		return
	}

	// 5. Create a Challenger from the Config and the ChallengerProcessor.
	challenger, err := challenger.NewChallenger(
		challengerConfig,
		challengerProcessor,
	)
	if err != nil {
		logger.Errorf("Failed to create challenger: %w", err)
		return
	}

	// 6. Start the challenger
	err = challenger.Start(context.Background())
	if err != nil {
		logger.Errorf("Failure while running challenger: %w", err)
		return
	}
}
