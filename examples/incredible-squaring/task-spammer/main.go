package main

import (
	"context"
	"iter"
	"math/big"
	"time"

	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	cstaskmanager "github.com/Layr-Labs/eigensdk-go/examples/incredible-squaring/bindings/taskManager"
	"github.com/Layr-Labs/eigensdk-go/logging"
	taskspammer "github.com/Layr-Labs/eigensdk-go/task-spammer"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	logger, err := logging.NewZapLogger(logging.Development)
	if err != nil {
		return
	}

	// This pk should be related to the address passed to TaskManager as task_spammer_addr when initialized
	taskSpammerPk := "0x4bbbf85ce3377467afe5d46f804f221813b2bb87f24d81f60f1fcdbf7cbf4356"

	ethHttpUrl := "http://localhost:8545"
	ethHttpClient, err := ethclient.Dial(ethHttpUrl)
	if err != nil {
		return
	}

	ecdsaPrivateKey, err := crypto.HexToECDSA(taskSpammerPk)
	if err != nil {
		return
	}

	txMgr, err := txmgr.NewSimpleTxManagerFromPrivateKey(logger, ethHttpClient, ecdsaPrivateKey)
	if err != nil {
		return
	}

	// This value is extracted from the deployment output files
	taskManagerAddress := common.HexToAddress("0x2bdcc0de6be1f7d2ee689a0342d76f52e8efaba3")

	abi, err := cstaskmanager.ContractIncredibleSquaringTaskManagerMetaData.GetAbi()
	if err != nil {
		return
	}

	taskCreator, err := taskspammer.NewTaskCreatorFromAbi[*big.Int](taskManagerAddress, *abi, txMgr, ethHttpClient)
	if err != nil {
		return
	}

	taskSpammerConfig := taskspammer.Config{
		Logger: logger,

		// This means TaskGenerator will send tasks every 10 seconds
		TimeBetweenTasks: 10 * time.Second,

		QuorumThresholdPercentage: 100,
		QuorumNumbers:             []uint8{0},
	}
	taskGen, err := taskspammer.NewTaskSpammer(taskCreator, taskSpammerConfig)
	if err != nil {
		return
	}

	seq := NewNumberToSquareSequence()

	err = taskGen.Start(context.Background(), seq)
	if err != nil {
		return
	}
}

// Returns an iterator for the sequence 1, 2, 3, ...
func NewNumberToSquareSequence() iter.Seq[*big.Int] {
	acc := big.NewInt(1)
	delta := big.NewInt(1)
	return func(yield func(*big.Int) bool) {
		for {
			if !yield(acc) {
				break
			}
			acc.Add(acc, delta)
		}
	}
}
