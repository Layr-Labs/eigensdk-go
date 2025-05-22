package main

import (
	"context"
	"math/big"

	"github.com/Layr-Labs/eigensdk-go/aggregator"
	taskprocessor "github.com/Layr-Labs/eigensdk-go/aggregator/task-processor"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	"github.com/Layr-Labs/eigensdk-go/logging"
	taskmanager "github.com/Layr-Labs/eigensdk-go/task-manager"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	cstaskmanager "github.com/Layr-Labs/eigensdk-go/examples/incredible-squaring/bindings/taskManager"
)

func main() {
	// 1. Create the logger where all the loggs will appear
	logger, err := logging.NewZapLogger(logging.Production) // Change here if want to change logging level
	if err != nil {
		println("Failure creating logger")
		return
	}

	// 2. Create the ethereum client that will send the RPC messages to the node
	ethHttpUrl := "http://localhost:8545"
	ethHttpClient, err := ethclient.Dial(ethHttpUrl)
	if err != nil {
		return
	}

	// 3. Create the transaction manager, that will manage the transaction sending
	ecdsaPrivateKey, err := crypto.HexToECDSA("2a871d0798f97d79848a013d4936a73bf4cc922c825d33c1cf7073dff6d409c6")
	if err != nil {
		logger.Errorf("Cannot parse ecdsa private key", "err", err)
		return
	}

	txMgr, err := txmgr.NewSimpleTxManagerFromPrivateKey(logger, ethHttpClient, ecdsaPrivateKey)
	if err != nil {
		logger.Errorf("Failed to create transaction manager", "err", err)
		return
	}

	// 4. Get the ABI of the task manager contract's binding
	taskManagerAbi, err := cstaskmanager.ContractIncredibleSquaringTaskManagerMetaData.GetAbi()
	if err != nil {
		logger.Fatalf(err.Error())
	}

	// 5. Create the task responder, which will send the aggregated responses to the on-chain TaskManager contract.
	// Here we use an SDK implementation that satisfies the TaskResponder interface, but you can create your own wrapper
	// which implements the interface and provide it to the Indexing Task Processor.
	taskManagerAddr := common.HexToAddress("0x2bdcc0de6be1f7d2ee689a0342d76f52e8efaba3")
	taskResponder, err := taskmanager.NewTaskManagerFromAbi[*big.Int, *big.Int](
		taskManagerAddr,
		taskManagerAbi,
		txMgr,
		ethHttpClient,
	)
	if err != nil {
		logger.Fatalf(err.Error())
	}

	// 6. Create the Task Processor, which will manage the processing of the tasks and the aggregated responses.
	// Here we use the IndexingTaskProcessor, a generic implementation provided by the SDK that saves the tasks
	// in a map and delegates the sending of aggregated responses to the on-chain TaskManager contract.
	taskProcessor, err := taskprocessor.NewIndexingTaskProcessor(logger, taskResponder)
	if err != nil {
		logger.Fatalf(err.Error())
	}

	// 7. Create the config passed to the aggregator, including some addresses and ethereum node urls.
	cfg := aggregator.Config{
		RegistryCoordinatorAddress:    common.HexToAddress("0x7bc06c482dead17c0e297afbc32f6e63d3846650"),
		OperatorStateRetrieverAddress: common.HexToAddress("0x4c5859f0f772848b2d91f1d83e2fe57935348029"),
		EthHttpUrl:                    ethHttpUrl,
		EthWsUrl:                      "ws://localhost:8545",
		AggregatorServerIpPortAddr:    "localhost:8090",
	}

	// 8. Build the aggregator, providing aggregator config, logger, task processor and the task manager ABI.
	agg, err := aggregator.NewAggregator(cfg, logger, taskProcessor, taskManagerAbi)
	if err != nil {
		logger.Fatalf(err.Error())
	}

	// 9. Run the created aggregator
	err = agg.Start(context.Background())
	if err != nil {
		logger.Fatalf(err.Error())
	}
}
