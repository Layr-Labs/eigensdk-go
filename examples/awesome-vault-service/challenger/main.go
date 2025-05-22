package main

import (
	"context"

	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	"github.com/Layr-Labs/eigensdk-go/challenger"
	challengerprocessor "github.com/Layr-Labs/eigensdk-go/challenger/challenger-processor"
	"github.com/Layr-Labs/eigensdk-go/logging"
	taskmanager "github.com/Layr-Labs/eigensdk-go/task-manager"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	examplecommon "github.com/Layr-Labs/eigensdk-go/examples/awesome-vault-service/common"
	avtaskmanager "github.com/Layr-Labs/eigensdk-go/examples/awesome-vault-service/contracts/bindings/AwesomeVaultTaskManager"
)

func main() {
	// 1. Create the logger where all the loggs will appear
	logger, err := logging.NewZapLogger(logging.Production) // Change here if want to change logging level
	if err != nil {
		println("Failure creating logger")
		return
	}

	// 2. Get the ABI of the task manager contract's binding
	taskManagerAbi, err := avtaskmanager.ContractAwesomeVaultTaskManagerMetaData.GetAbi()
	if err != nil {
		logger.Errorf("Failed to get task manager abi: %w", err)
		return
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
	// Here we use an SDK implementation that satisfies the ChallengeRaiser interface, but you can create your own wrapper
	// which implements the interface and provide it to the Indexing Task Processor.
	// Note that in this step we define the input and output types that we are using on our AVS. In this case
	// the TaskInput struct (a key-value pair) and 32 bytes.
	challengeRaiser, err := taskmanager.NewTaskManagerFromAbi[examplecommon.TaskInput, [32]byte](
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
	// custom Response calculator, declared on examples/awesome-vault-service/common/response_calculator.go. We
	// decided to create a custom Response Calculator because we have to save state between task responses, so
	// the NewFunctionResponseCalculator from the operator package wont be useful. You can see the implementation
	// to view how simple is to build one for your AVS. With that calculator we create the validator with the
	// ResponseValidationFunctionFromResponseCalculator builder from the challengerprocessor package.
	vaultServiceResponseCalc := examplecommon.NewVaultServiceResponseCalculator()
	vaultSetValidation := challengerprocessor.ResponseValidationFunctionFromResponseCalculator(vaultServiceResponseCalc, func(a, b [32]byte) bool { return a == b })

	// 7. Create the Challenger Processor, which will manage the challenge raising in case the response is not
	// correct. Here we use the IndexingChallengerProcessor, a generic implementation provided by the SDK that
	// saves the tasks in a map and in case of receiving a wrong response delegates the raising of the
	// challenges to the on-chain TaskManager contract.
	challengerProcessor, err := challengerprocessor.NewIndexingChallengerProcessor(logger, vaultSetValidation, challengeRaiser)
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
