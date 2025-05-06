package main

import (
	"context"
	"iter"
	"math/big"
	"time"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/wallet"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/signerv2"
	taskspammer "github.com/Layr-Labs/eigensdk-go/task-spammer"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	taskmanager "github.com/Layr-Labs/eigensdk-go/examples/incredible-dot-product/contracts/bindings/IncredibleDotProductTaskManager"
)

type DotProductInput struct {
	X []*big.Int
	Y []*big.Int
}

func main() {
	logger, err := logging.NewZapLogger(logging.Production)
	if err != nil {
		panic(err)
	}

	taskManagerAbi, err := taskmanager.ContractIncredibleDotProductTaskManagerMetaData.GetAbi()
	if err != nil {
		logger.Errorf("Failed to get task manager abi: %w", err)
	}

	ethHttpUrl := "http://localhost:8545"
	ethClient, err := ethclient.Dial(ethHttpUrl)
	if err != nil {
		logger.Errorf("Failed to dial ethclient: %w", err)
	}

	taskManagerAddr := common.HexToAddress("0x2bdcc0de6be1f7d2ee689a0342d76f52e8efaba3")

	taskSpammerPrivateKey := "4bbbf85ce3377467afe5d46f804f221813b2bb87f24d81f60f1fcdbf7cbf4356"
	ecdsaPrivateKey, err := crypto.HexToECDSA(taskSpammerPrivateKey)
	if err != nil {
		logger.Errorf("Failed to create ecdsa private key: %w", err)
	}

	chainId, err := ethClient.ChainID(context.Background())
	if err != nil {
		logger.Error("Cannot get chainId", "err", err)
	}

	taskSpammerAddr := common.HexToAddress("0x14dC79964da2C08b23698B3D3cc7Ca32193d9955")

	signerV2, _, err := signerv2.SignerFromConfig(signerv2.Config{PrivateKey: ecdsaPrivateKey}, chainId)
	if err != nil {
		panic(err)
	}
	skWallet, err := wallet.NewPrivateKeyWallet(ethClient, signerV2, taskSpammerAddr, logger)
	if err != nil {
		panic(err)
	}
	txMgr := txmgr.NewSimpleTxManager(skWallet, ethClient, logger, taskSpammerAddr)

	taskCreator, err := taskspammer.NewTaskCreatorFromAbi[DotProductInput](taskManagerAddr, *taskManagerAbi, txMgr, ethClient)
	if err != nil {
		logger.Errorf("Failed to create Task Creator: %w", err)
	}

	taskSpammerConfig := taskspammer.Config{
		Logger:                    logger,
		TimeBetweenTasks:          10 * time.Second,
		QuorumThresholdPercentage: 100,
		QuorumNumbers:             []uint8{0},
	}
	taskSpammer, err := taskspammer.NewTaskSpammer(taskCreator, taskSpammerConfig)
	if err != nil {
		logger.Errorf("Failed to create Task Spammer: %w", err)
	}

	seq := NewLinearSequence()

	err = taskSpammer.Start(context.Background(), seq)
	if err != nil {
		logger.Errorf("Failure while running Task Spammer: %w", err)
	}
}

// Returns an iterator for the sequence 1, 2, 3, ...
// Inspired on the one from examples/incredible-squaring/task-spammer/task_spammer_use_example.go
func NewLinearSequence() iter.Seq[DotProductInput] {
	acc := big.NewInt(1)
	delta := big.NewInt(1)
	return func(yield func(DotProductInput) bool) {
		for {
			newInput := DotProductInput{X: []*big.Int{acc}, Y: []*big.Int{acc}}
			if !yield(newInput) {
				break
			}
			acc.Add(acc, delta)
		}
	}
}
