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
	"github.com/Layr-Labs/eigensdk-go/testutils"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/Layr-Labs/eigensdk-go/examples/common"
	examplecommon "github.com/Layr-Labs/eigensdk-go/examples/incredible-squaring/common"
	cstaskmanager "github.com/Layr-Labs/eigensdk-go/examples/incredible-squaring/contracts/bindings/IncredibleSquaringTaskManager"
)

// This config contains the values needed to create and run the Challenger
type Config struct {
	TaskManagerAddress string `toml:"task_manager_address"`

	challenger.Config
}

func main() {
	// 0. Create the logger where all the logs will appear
	logger, err := logging.NewZapLogger(logging.Production) // Change here if want to change logging level
	if err != nil {
		println("Failure creating logger")
		return
	}

	// 1. Create the config passed to the challenger

	challengerConfig := &Config{}
	err = examplecommon.ReadTomlConfig("config/config.toml", challengerConfig)
	if err != nil {
		logger.Fatalf(err.Error())
	}

	// Get the ABI of the task manager contract's binding
	taskManagerAbi, err := cstaskmanager.ContractIncredibleSquaringTaskManagerMetaData.GetAbi()
	if err != nil {
		logger.Fatalf(err.Error())
	}

	// Create the ethereum client that will send the RPC messages to the node
	ethHttpClient, err := ethclient.Dial(challengerConfig.EthHttpUrl)
	if err != nil {
		logger.Errorf("Failed to dial ethclient: %w", err)
		return
	}

	cfg := challengerConfig.Config

	// 2. Define a function that verifies the response for a task. Note that here we wrap the logic into a
	// `ResponseCalculator` implementation, and then we use the ResponseValidationFunctionFromResponseCalculator,
	// which creates a validation function that will compute the logic function to validate the response
	squareCalculator := operator.NewFunctionResponseCalculator(examplecommon.Square)
	squareValidation := challengerprocessor.ResponseValidationFunctionFromResponseCalculator(squareCalculator, common.BigIntEqual)

	// 3. Provide a struct implementing the `ChallengerProcessor` interface, first instantiating the values
	// required for creating it

	// Create the transaction manager, that will manage the transaction sending
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

	// Provide a struct that implements the ChallengeRaiser interface. Here we provide a SDK implementation that
	// satisfies the ChallengeRaiser interface.
	// Note that in this step we define the input and output types that we are using on our AVS. In this case
	// the input and output are both big int numbers.
	challengeRaiser, err := taskmanager.NewTaskManagerFromAbi[*big.Int, *big.Int](
		gethcommon.HexToAddress(challengerConfig.TaskManagerAddress),
		taskManagerAbi,
		txMgr,
		ethHttpClient,
	)
	if err != nil {
		logger.Errorf("Failed to create challenger raiser: %v", err)
		return
	}

	// Create the Challenger Processor, with the challenger raiser created above.
	indexingChallengerProcessor, err := challengerprocessor.NewIndexingChallengerProcessor(logger, squareValidation, challengeRaiser)
	if err != nil {
		logger.Errorf("Failed to create challenger logic from config: %v", err)
		return
	}

	// 4. Create a Challenger from the Config and the ChallengerProcessor.
	challenger, err := challenger.NewChallenger(
		logger,
		cfg,
		taskManagerAbi,
		indexingChallengerProcessor,
	)
	if err != nil {
		logger.Errorf("Failed to create challenger from config: %v", err)
		return
	}

	// 5. Start the challenger
	err = <-challenger.Start(context.Background())
	if err != nil {
		logger.Errorf("Error while running challenger: %v", err)
		return
	}
}
