package aggregator_example

import (
	"context"
	"math/big"

	"github.com/Layr-Labs/eigensdk-go/aggregator"
	"github.com/Layr-Labs/eigensdk-go/logging"
	taskprocessor "github.com/Layr-Labs/eigensdk-go/task-processor"
	"github.com/Layr-Labs/eigensdk-go/testutils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	cstaskmanager "github.com/Layr-Labs/eigensdk-go/examples/incredible-squaring/bindings/taskManager"
	taskspammerexample "github.com/Layr-Labs/eigensdk-go/examples/incredible-squaring/task-spammer"
)

func main() {
	logger, err := logging.NewZapLogger(logging.Production) // Change here if want to change logging level
	if err != nil {
		println("Failure creating logger")
		return
	}

	ethHttpUrl := "http://localhost:8545"
	ethHttpClient, err := ethclient.Dial(ethHttpUrl)
	if err != nil {
		return
	}

	txMgr, err := taskspammerexample.GetTxManager(logger, ethHttpClient, testutils.ANVIL_FIRST_PRIVATE_KEY)
	if err != nil {
		return
	}

	ecdsaPrivateKey, err := crypto.HexToECDSA(testutils.ANVIL_FIRST_PRIVATE_KEY)
	if err != nil {
		logger.Errorf("Cannot parse ecdsa private key", "err", err)
		return
	}

	taskManagerAbi, err := cstaskmanager.ContractIncredibleSquaringTaskManagerMetaData.GetAbi()
	if err != nil {
		logger.Fatalf(err.Error())
	}

	cfg := aggregator.AggregatorConfig{
		RegistryCoordinatorAddress:    common.HexToAddress("0x7bc06c482dead17c0e297afbc32f6e63d3846650"),
		OperatorStateRetrieverAddress: common.HexToAddress("0x4c5859f0f772848b2d91f1d83e2fe57935348029"),
		ServiceManagerAddress:         common.HexToAddress("0x5f3f1dbd7b74c6b46e8c44f98792a1daf8d69154"),
		EthHttpClient:                 ethHttpClient,
		Logger:                        logger,
		EthHttpUrl:                    ethHttpUrl,
		EthWsUrl:                      "ws://localhost:8545",
		EcdsaPrivateKey:               ecdsaPrivateKey,
		AggregatorServerIpPortAddr:    "localhost:8090",
		TaskManagerAbi:                taskManagerAbi,
	}

	taskManagerAddr := common.HexToAddress("0x2bdcc0de6be1f7d2ee689a0342d76f52e8efaba3")

	taskResponder, err := taskprocessor.NewTaskResponderFromAbi[*big.Int, *big.Int](
		taskManagerAddr,
		taskManagerAbi,
		txMgr,
		ethHttpClient,
	)
	if err != nil {
		logger.Fatalf(err.Error())
	}

	taskProcessor, err := taskprocessor.NewIndexingTaskProcessor(taskManagerAbi, logger, taskResponder)
	if err != nil {
		logger.Fatalf(err.Error())
	}

	agg, err := aggregator.NewAggregator(cfg, taskProcessor)
	if err != nil {
		logger.Fatalf(err.Error())
	}

	err = agg.Start(context.Background())
	if err != nil {
		logger.Fatalf(err.Error())
	}
}
