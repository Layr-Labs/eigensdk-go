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
	cstaskmanager "github.com/Layr-Labs/eigensdk-go/examples/incredible-squaring/bindings/taskManager"
	examplecommon "github.com/Layr-Labs/eigensdk-go/examples/incredible-squaring/common"
)

func main() {
	// 1. Create the logger where all the loggs will appear
	logger, err := logging.NewZapLogger(logging.Production) // Change here if want to change logging level
	if err != nil {
		println("Failure creating logger")
		return
	}

	// 2. Get the ABI of the task manager contract's binding
	taskManagerAbi, err := cstaskmanager.ContractIncredibleSquaringTaskManagerMetaData.GetAbi()
	if err != nil {
		logger.Fatalf(err.Error())
	}

	// 3. Create the ethereum client that will send the RPC messages to the node
	ethHttpClient, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		return
	}

	// 4. Create the transaction manager, that will manage the transaction sending
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

	// 5. Create the challenger raiser, which will raise the challenges to the on-chain TaskManager contract.
	// Here we use an SDK implementation that satisfies the ChallengeRaiser interface, but you can create your own wrapper
	// which implements the interface and provide it to the Indexing Task Processor.
	// Note that in this step we define the input and output types that we are using on our AVS. In this case
	// the input and output are both big int numbers.
	challengerRaiser, err := taskmanager.NewTaskManagerFromAbi[*big.Int, *big.Int](
		gethcommon.HexToAddress("0x2bdcc0de6be1f7d2ee689a0342d76f52e8efaba3"),
		taskManagerAbi,
		txMgr,
		ethHttpClient,
	)
	if err != nil {
		logger.Errorf("Failed to create challenger raiser: %v", err)
		return
	}

	// 6. Create the calculator and validation function with the AVS calculation logic. Note that here we create a
	// Response calculator with the NewFunctionResponseCalculator from the operator package, and with that
	// calculator we create the validator with the ResponseValidationFunctionFromResponseCalculator builder from
	// the challengerprocessor package.
	squareCalculator := operator.NewFunctionResponseCalculator(examplecommon.Square)
	squareValidation := challengerprocessor.ResponseValidationFunctionFromResponseCalculator(squareCalculator, common.BigIntEqual)

	// 7. Create the Challenger Processor, which will manage the challenge raising in case the response is not
	// correct. Here we use the IndexingChallengerProcessor, a generic implementation provided by the SDK that
	// saves the tasks in a map and in case of receiving a wrong response delegates the raising of the
	// challenges to the on-chain TaskManager contract.
	indexingTaskProcessor, err := challengerprocessor.NewIndexingChallengerProcessor(logger, squareValidation, challengerRaiser)
	if err != nil {
		logger.Errorf("Failed to create challenger logic from config: %v", err)
		return
	}

	// 8. Create the config passed to the challenger.
	cfg := challenger.Config{
		EthWsUrl:       "ws://localhost:8545",
		Logger:         logger,
		TaskManagerAbi: taskManagerAbi,
		EthClient:      ethHttpClient,
	}

	// 9. Build the challenger, providing challenger config, and the challenger processor.
	challenger, err := challenger.NewChallenger(
		cfg,
		indexingTaskProcessor,
	)
	if err != nil {
		logger.Errorf("Failed to create challenger from config: %v", err)
		return
	}

	// 10. Run the created challenger
	err = challenger.Start(context.Background())
	if err != nil {
		logger.Errorf("Error while running operator: %v", err)
		return
	}
}
