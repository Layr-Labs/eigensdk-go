package main

import (
	"context"

	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	"github.com/Layr-Labs/eigensdk-go/challenger"
	examplechallenger "github.com/Layr-Labs/eigensdk-go/examples/incredible-dot-product/challenger"
	"github.com/Layr-Labs/eigensdk-go/logging"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	taskmanager "github.com/Layr-Labs/eigensdk-go/examples/incredible-dot-product/contracts/bindings/IncredibleDotProductTaskManager"
)

func main() {
	logger, err := logging.NewZapLogger(logging.Production)
	if err != nil {
		println("Failure creating logger")
		return
	}

	taskManagerAbi, err := taskmanager.ContractIncredibleDotProductTaskManagerMetaData.GetAbi()
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

	taskManagerAddr := gethcommon.HexToAddress("0x2bdcc0de6be1f7d2ee689a0342d76f52e8efaba3")

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

	challengerVerifier, err := examplechallenger.NewChallengeVerifier(logger, taskManagerAddr, *ethClient, txMgr)
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
	challenger, err := challenger.NewChallenger(challengerConfig, challengerVerifier)
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
