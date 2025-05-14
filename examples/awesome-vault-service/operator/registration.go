package main

import (
	"context"
	"math/big"
	"os"
	"time"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/elcontracts"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/wallet"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	erc20mock "github.com/Layr-Labs/eigensdk-go/contracts/bindings/MockERC20"
	"github.com/Layr-Labs/eigensdk-go/crypto/ecdsa"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/metrics"
	sdkoperator "github.com/Layr-Labs/eigensdk-go/operator"
	"github.com/Layr-Labs/eigensdk-go/signerv2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func RegisterOperatorOnStartup(logger logging.Logger) error {
	amount := new(big.Int)
	amount.SetString("1000000000000000000000", 10)

	registrationConfig := sdkoperator.RegistrationConfig{
		Logger: logger,

		OperatorAddr:            common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"),
		AllocationManagerAddr:   common.HexToAddress("0x2279b7a0a67db372996a5fab50d91eaa73d2ebe6"),
		AvsAddress:              common.HexToAddress("0xcd8a1c3ba11cf5ecfa6267617243239504a98d90"),
		RegistryCoordinatorAddr: common.HexToAddress("0xfd471836031dc5108809d173a067e8486b9047a3"),
		StrategyAddr:            common.HexToAddress("0x2b961e3959b79326a8e7f64ef0d2d825707669b5"),

		DelegationManagerAddress:    common.HexToAddress("0x9fe46736679d2d9a65f0992f2272de9f3c7fa6e0"),
		RewardsCoordinatorAddress:   common.HexToAddress("0xa51c1fc2f0d1a1b8494ed1fe312d7c3a78ed91c0"),
		PermissionControllerAddress: common.HexToAddress("0x59b670e9fa9d0a427751af201d676719a970857b"),

		EthHttpUrl: "http://localhost:8545",

		EcdsaKeyStorePath: "keys/test.ecdsa.key.json",
		BlsKeyStorePath:   "keys/test.bls.key.json",

		AmountToMint:         amount,
		AllocatableMagnitude: 1000000000000000,

		OperatorSetId: 0,
	}

	err := sdkoperator.RegisterOperatorOnStartup(registrationConfig)
	if err != nil {
		logger.Errorf("Failed to register operator on startup: %w", err)
		return err
	}

	elcontractsConfig := elcontracts.Config{
		DelegationManagerAddress:    registrationConfig.DelegationManagerAddress,
		RewardsCoordinatorAddress:   registrationConfig.RewardsCoordinatorAddress,
		PermissionControllerAddress: registrationConfig.PermissionControllerAddress,
	}

	ethRpcClient, err := ethclient.Dial(registrationConfig.EthHttpUrl)
	if err != nil {
		registrationConfig.Logger.Errorf("Cannot create http ethclient", "err", err)
		return err
	}

	rpcCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	chainid, err := ethRpcClient.ChainID(rpcCtx)
	if err != nil {
		registrationConfig.Logger.Error("Cannot get chain id", "err", err)
		return err
	}

	ecdsaKeyPassword, ok := os.LookupEnv("OPERATOR_ECDSA_KEY_PASSWORD")
	if !ok {
		registrationConfig.Logger.Warnf("OPERATOR_ECDSA_KEY_PASSWORD env var not set. using empty string")
	}

	operatorEcdsaPrivateKey, err := ecdsa.ReadKey(
		registrationConfig.EcdsaKeyStorePath,
		ecdsaKeyPassword,
	)
	if err != nil {
		return err
	}

	signerV2, senderAddr, err := signerv2.SignerFromConfig(signerv2.Config{
		PrivateKey: operatorEcdsaPrivateKey,
	}, chainid)
	if err != nil {
		registrationConfig.Logger.Fatalf(err.Error())
	}

	pkWallet, err := wallet.NewPrivateKeyWallet(ethRpcClient, signerV2, senderAddr, registrationConfig.Logger)
	if err != nil {
		return err
	}

	txMgr := txmgr.NewSimpleTxManager(pkWallet, ethRpcClient, registrationConfig.Logger, senderAddr)

	err = DepositIntoStrategyForOperator(
		registrationConfig.Logger,
		elcontractsConfig,
		ethRpcClient,
		registrationConfig.StrategyAddr,
		txMgr,
		registrationConfig.OperatorAddr,
		registrationConfig.AmountToMint,
	)
	if err != nil {
		registrationConfig.Logger.Fatalf("Failed to deposit into strategy for operator on startup: %v", err.Error())
	}

	return nil
}

func DepositIntoStrategyForOperator(
	logger logging.Logger,
	elcontractsConfig elcontracts.Config,
	ethClient *ethclient.Client,
	strategyAddr common.Address,
	txMgr txmgr.TxManager,
	operatorAddr common.Address,
	amount *big.Int,
) error {
	elReader, err := elcontracts.NewReaderFromConfig(elcontractsConfig, ethClient, logger)
	if err != nil {
		logger.Error("Error creating eigenlayer chain writer", "err", err)
		return err
	}

	elWriter, err := elcontracts.NewWriterFromConfig(
		elcontractsConfig,
		ethClient,
		logger,
		&metrics.EigenMetrics{},
		txMgr,
	)
	if err != nil {
		logger.Error("Error creating eigenlayer chain writer", "err", err)
		return err
	}

	_, tokenAddr, err := elReader.GetStrategyAndUnderlyingToken(context.Background(), strategyAddr)
	if err != nil {
		logger.Error("Failed to fetch strategy contract", "err", err)
		return err
	}
	logger.Info(tokenAddr.String())

	contractErc20Mock, err := erc20mock.NewContractMockERC20(tokenAddr, ethClient)
	if err != nil {
		logger.Error("Failed to fetch ERC20Mock contract", "err", err)
		return err
	}
	txOpts, err := txMgr.GetNoSendTxOpts()
	if err != nil {
		logger.Errorf("Error in GetNoSendTxOpts")
		return err
	}

	tx, err := contractErc20Mock.Mint(txOpts, operatorAddr, amount)
	if err != nil {
		logger.Errorf("Error assembling Mint tx")
		return err
	}
	_, err = txMgr.Send(context.Background(), tx, true)
	if err != nil {
		logger.Errorf("Error submitting Mint tx")
		return err
	}

	_, err = elWriter.DepositERC20IntoStrategy(context.Background(), strategyAddr, amount, true)
	if err != nil {
		logger.Errorf("Error depositing into strategy", "err", err)
		return err
	}

	return nil
}
