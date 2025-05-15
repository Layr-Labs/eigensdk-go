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

	cfg := challenger.Config{
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

	challengerRaiser, err := taskmanager.NewTaskManagerFromAbi[*big.Int, *big.Int, any](
		gethcommon.HexToAddress("0x2bdcc0de6be1f7d2ee689a0342d76f52e8efaba3"),
		taskManagerAbi,
		txMgr,
		ethHttpClient,
	)
	if err != nil {
		logger.Errorf("Failed to create challenger raiser: %v", err)
		return
	}

	squareCalculator := operator.NewFunctionResponseCalculator(examplecommon.Square)
	squareValidation := challengerprocessor.ResponseValidationFunctionFromResponseCalculator(squareCalculator, common.BigIntEqual)
	indexingTaskProcessor, err := challengerprocessor.NewIndexingChallengerProcessor(logger, squareValidation, challengerRaiser)
	if err != nil {
		logger.Errorf("Failed to create challenger logic from config: %v", err)
		return
	}

	challenger, err := challenger.NewChallenger(
		cfg,
		indexingTaskProcessor,
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
