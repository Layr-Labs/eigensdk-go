package main

import (
	"context"

	"github.com/Layr-Labs/eigensdk-go/challenger"
	examplechallenger "github.com/Layr-Labs/eigensdk-go/examples/incredible-dot-product/challenger"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	taskmanager "github.com/Layr-Labs/eigensdk-go/examples/incredible-dot-product/contracts/bindings/IncredibleDotProductTaskManager"
)

func main() {
	logger, err := logging.NewZapLogger(logging.Development)
	if err != nil {
		panic(err)
	}

	taskManagerAbi, err := taskmanager.ContractIncredibleDotProductTaskManagerMetaData.GetAbi()
	if err != nil {
		logger.Errorf("Failed to get task manager abi: %w", err)
	}

	ethHttpUrl := "127.0.0.1:8545"
	ethClient, err := ethclient.Dial(ethHttpUrl)
	if err != nil {
		logger.Errorf("Failed to dial ethclient: %w", err)
	}

	taskManagerAddr := common.HexToAddress("0x")

	challengerVerifier, err := examplechallenger.NewChallengeVerifier(logger, taskManagerAddr)
	if err != nil {
		logger.Errorf("Failed to create challenger verifier: %w", err)
	}

	challengerConfig := challenger.ChallengerConfig{
		Logger:         logger,
		TaskManagerAbi: taskManagerAbi,
		EthClient:      ethClient,
		// EthWsUrl: "ws:8080",
	}
	challenger, err := challenger.NewChallenger(challengerConfig, challengerVerifier)
	if err != nil {
		logger.Errorf("Failed to create challenger: %w", err)
	}

	err = challenger.Start(context.Background())
	if err != nil {
		logger.Errorf("Failure while running challenger: %w", err)
	}
}
