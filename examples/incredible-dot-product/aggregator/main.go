package main

import (
	"context"
	"math/big"

	"github.com/Layr-Labs/eigensdk-go/aggregator"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/wallet"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/signerv2"
	taskprocessor "github.com/Layr-Labs/eigensdk-go/task-processor"
	gethcommon "github.com/ethereum/go-ethereum/common"
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
		println("Failure creating logger")
		return
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

	aggregatorPrivateKey := "2a871d0798f97d79848a013d4936a73bf4cc922c825d33c1cf7073dff6d409c6"
	ecdsaPrivateKey, err := crypto.HexToECDSA(aggregatorPrivateKey)
	if err != nil {
		logger.Errorf("Failed to create ecdsa private key: %w", err)
	}

	chainId, err := ethClient.ChainID(context.Background())
	if err != nil {
		logger.Error("Cannot get chainId", "err", err)
	}

	aggregatorAddr := gethcommon.HexToAddress("0xa0Ee7A142d267C1f36714E4a8F75612F20a79720")

	signerV2, _, err := signerv2.SignerFromConfig(signerv2.Config{PrivateKey: ecdsaPrivateKey}, chainId)
	if err != nil {
		logger.Fatalf(err.Error())
	}
	skWallet, err := wallet.NewPrivateKeyWallet(ethClient, signerV2, aggregatorAddr, logger)
	if err != nil {
		logger.Fatalf(err.Error())
	}
	txMgr := txmgr.NewSimpleTxManager(skWallet, ethClient, logger, aggregatorAddr)

	aggConfig := aggregator.AggregatorConfig{
		Logger:             logger,
		TaskManagerAbi:     taskManagerAbi,
		TaskResponseHashFn: nil,

		EthHttpUrl:                 ethHttpUrl,
		EthWsUrl:                   "ws://localhost:8545",
		AggregatorServerIpPortAddr: "localhost:8090",

		RegistryCoordinatorAddress:    gethcommon.HexToAddress("0x7bc06c482dead17c0e297afbc32f6e63d3846650"),
		OperatorStateRetrieverAddress: gethcommon.HexToAddress("0x4c5859f0f772848b2d91f1d83e2fe57935348029"),
		ServiceManagerAddress:         gethcommon.HexToAddress("0x5f3f1dbd7b74c6b46e8c44f98792a1daf8d69154"),

		EthHttpClient:   ethClient,
		EcdsaPrivateKey: ecdsaPrivateKey,
	}

	taskManagerAddr := gethcommon.HexToAddress("0x2bdcc0de6be1f7d2ee689a0342d76f52e8efaba3")
	taskResponder, err := taskprocessor.NewTaskResponderFromAbi[DotProductInput, *big.Int](taskManagerAddr, taskManagerAbi, txMgr, ethClient)
	if err != nil {
		logger.Errorf("Failed to create Task Responder: %w", err)
	}

	taskProcessor, err := taskprocessor.NewIndexingTaskProcessor(logger, taskResponder)
	if err != nil {
		logger.Errorf("Failed to create Task Processor: %w", err)
	}

	aggregator, err := aggregator.NewAggregator(aggConfig, taskProcessor)
	if err != nil {
		logger.Errorf("Failed to create aggregator: %w", err)
	}

	err = aggregator.Start(context.Background())
	if err != nil {
		logger.Errorf("Failure while running aggregator: %w", err)
	}
}
