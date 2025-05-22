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
	// 1. Create the logger where all the loggs will appear
	logger, err := logging.NewZapLogger(logging.Production) // Change here if want to change logging level
	if err != nil {
		println("Failure creating logger")
		return
	}

	// 2. Get the ABI of the task manager contract's binding
	taskManagerAbi, err := idptaskmanager.ContractIncredibleDotProductTaskManagerMetaData.GetAbi()
	if err != nil {
		logger.Fatalf(err.Error())
	}

	// 3. Create the ethereum client that will send the RPC messages to the node
	ethClient, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		logger.Errorf("Failed to dial ethclient: %w", err)
		return
	}

	// 4. Create the transaction manager, that will manage the transaction sending
	challengerPrivateKey := "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	ecdsaPrivateKey, err := crypto.HexToECDSA(challengerPrivateKey)
	if err != nil {
		logger.Errorf("Failed to create ecdsa private key: %w", err)
		return
	}

	txMgr, err := txmgr.NewSimpleTxManagerFromPrivateKey(logger, ethClient, ecdsaPrivateKey)
	if err != nil {
		logger.Errorf("Failed to create transaction manager", "err", err)
		return
	}

	// 5. Create the challenger raiser, which will raise the challenges to the on-chain TaskManager contract.
	// Here we use an SDK implementation that satisfies the ChallengeRaiser interface, but you can create your wrapper
	// (which implements the interface) and provide it to the Indexing Task Processor.
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

	// 6. Create the calculator and validation function with the AVS calculation logic. Note that here we create a
	// Response calculator with the NewFunctionResponseCalculator from the operator package, and with that
	// calculator, we create the validator with the ResponseValidationFunctionFromResponseCalculator builder from
	// the challengerprocessor package.
	dotProductCalculator := operator.NewFunctionResponseCalculator(examplecommon.DotProduct)
	dotProductValidation := challengerprocessor.ResponseValidationFunctionFromResponseCalculator(dotProductCalculator, common.BigIntEqual)

	// 7. Create the Challenger Processor, which will manage the challenge raising in case the response is
	// different. Here we use the IndexingChallengerProcessor, a generic implementation provided by the SDK that
	// saves the tasks in a map and in case of receiving a wrong response delegates the raising of the
	// challenges to the on-chain TaskManager contract.
	challengerProcessor, err := challengerprocessor.NewIndexingChallengerProcessor(logger, dotProductValidation, challengeRaiser)
	if err != nil {
		logger.Errorf("Failed to create challenger processor: %v", err)
		return
	}

	// 8. Create the config passed to the challenger.
	challengerConfig := challenger.Config{
		Logger:         logger,
		TaskManagerAbi: taskManagerAbi,
		EthClient:      ethClient,
		EthWsUrl:       "ws://localhost:8545",
	}

	// 9. Build the challenger, providing challenger config, and the challenger processor.
	challenger, err := challenger.NewChallenger(
		challengerConfig,
		challengerProcessor,
	)
	if err != nil {
		logger.Errorf("Failed to create challenger: %w", err)
		return
	}

	// 10. Run the created challenger
	err = challenger.Start(context.Background())
	if err != nil {
		logger.Errorf("Failure while running challenger: %w", err)
		return
	}
}
