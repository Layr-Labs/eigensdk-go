package operator

import (
	"context"
	"errors"
	"math/big"
	"os"
	"time"

	allocationmanager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/AllocationManager"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/crypto/ecdsa"
	"github.com/Layr-Labs/eigensdk-go/signerv2"

	erc20mock "github.com/Layr-Labs/eigensdk-go/contracts/bindings/MockERC20"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/elcontracts"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/wallet"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/metrics"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/Layr-Labs/eigensdk-go/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
)

func RegisterOperatorOnStartup(logger logging.Logger) error {
	operatorAddr := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	allocationManagerAddr := common.HexToAddress("0x2279b7a0a67db372996a5fab50d91eaa73d2ebe6")
	avsAddress := common.HexToAddress("0xcd8a1c3ba11cf5ecfa6267617243239504a98d90")
	registryCoordinatorAddr := common.HexToAddress("0xfd471836031dc5108809d173a067e8486b9047a3")
	strategyAddr := common.HexToAddress("0x2b961e3959b79326a8e7f64ef0d2d825707669b5")
	ethHttpUrl := "http://localhost:8545"

	delegationManagerAddress := common.HexToAddress("0x9fe46736679d2d9a65f0992f2272de9f3c7fa6e0")
	rewardsCoordinatorAddress := common.HexToAddress("0xa51c1fc2f0d1a1b8494ed1fe312d7c3a78ed91c0")
	permissionControllerAddress := common.HexToAddress("0x59b670e9fa9d0a427751af201d676719a970857b")

	ecdsaKeyStorePath := "keys/test.ecdsa.key.json"
	blsKeyStorePath := "keys/test.bls.key.json"
	stringMintAmount := "1000000000000000000000"

	allocatableMagnitude := uint64(1000000000000000)
	operatorSetId := uint32(0)

	ethRpcClient, err := ethclient.Dial(ethHttpUrl)
	if err != nil {
		logger.Errorf("Cannot create http ethclient", "err", err)
		return err
	}

	elcontractsConfig := elcontracts.Config{
		DelegationManagerAddress:    delegationManagerAddress,
		RewardsCoordinatorAddress:   rewardsCoordinatorAddress,
		PermissionControllerAddress: permissionControllerAddress,
	}

	rpcCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	chainid, err := ethRpcClient.ChainID(rpcCtx)
	if err != nil {
		logger.Error("Cannot get chain id", "err", err)
		return err
	}

	ecdsaKeyPassword, ok := os.LookupEnv("OPERATOR_ECDSA_KEY_PASSWORD")
	if !ok {
		logger.Warnf("OPERATOR_ECDSA_KEY_PASSWORD env var not set. using empty string")
	}

	operatorEcdsaPrivateKey, err := ecdsa.ReadKey(
		ecdsaKeyStorePath,
		ecdsaKeyPassword,
	)
	if err != nil {
		return err
	}

	signerV2, senderAddr, err := signerv2.SignerFromConfig(signerv2.Config{
		PrivateKey: operatorEcdsaPrivateKey,
	}, chainid)
	if err != nil {
		logger.Fatalf(err.Error())
	}

	pkWallet, err := wallet.NewPrivateKeyWallet(ethRpcClient, signerV2, senderAddr, logger)
	if err != nil {
		return err
	}

	txMgr := txmgr.NewSimpleTxManager(pkWallet, ethRpcClient, logger, senderAddr)

	err = RegisterOperatorWithEigenlayer(
		operatorAddr,
		elcontractsConfig,
		ethRpcClient,
		logger,
		txMgr,
	)
	if err != nil {
		logger.Fatalf("Failed to register operator with EigenLayer on startup: %v", err.Error())
	}

	amount := new(big.Int)
	amount.SetString(stringMintAmount, 10)
	err = DepositIntoStrategyForOperator(
		logger,
		elcontractsConfig,
		ethRpcClient,
		strategyAddr,
		txMgr,
		operatorAddr,
		amount,
	)
	if err != nil {
		logger.Fatalf("Failed to deposit into strategy for operator on startup: %v", err.Error())
	}

	blsKeyPassword, ok := os.LookupEnv("OPERATOR_BLS_KEY_PASSWORD")
	if !ok {
		logger.Warnf("OPERATOR_BLS_KEY_PASSWORD env var not set. using empty string")
	}
	blsKeyPair, err := bls.ReadPrivateKeyFromFile(blsKeyStorePath, blsKeyPassword)
	if err != nil {
		logger.Errorf("Cannot parse bls private key", "err", err)
		return err
	}

	err = RegisterForOperatorSets(
		operatorAddr,
		logger,
		elcontractsConfig,
		ethRpcClient,
		txMgr,
		registryCoordinatorAddr,
		avsAddress,
		[]uint32{operatorSetId},
		*blsKeyPair,
		"",
	)
	if err != nil {
		logger.Fatalf("Failed to register operator for operator sets on startup: %v", err.Error())
	}

	err = SetAllocationDelay(
		logger,
		operatorAddr,
		ethRpcClient,
		allocationManagerAddr,
		txMgr,
		0,
	)
	if err != nil {
		logger.Fatalf("Failed to set allocation delay: %v", err.Error())
	}

	err = modifyAllocations(
		operatorAddr,
		allocationManagerAddr,
		avsAddress,
		[]common.Address{strategyAddr},
		[]uint64{allocatableMagnitude},
		ethHttpUrl,
		txMgr,
		operatorSetId,
		logger,
	)
	if err != nil {
		logger.Fatalf("Failed to set modify allocations: %v", err.Error())
	}

	return nil
}

// The idea is to have a util function that just registers an operator for an operator set if the avs needs it
func RegisterOperatorWithEigenlayer(
	operatorAddr common.Address,
	elcontractsConfig elcontracts.Config,
	ethClient *ethclient.Client,
	logger logging.Logger,
	txMgr txmgr.TxManager,
) error {

	// Register operator with EigenLayer
	op := types.Operator{
		Address:                   operatorAddr.String(),
		DelegationApproverAddress: operatorAddr.String(),
	}

	elWriter, err := elcontracts.NewWriterFromConfig(elcontractsConfig, ethClient, logger, &metrics.EigenMetrics{}, txMgr)
	if err != nil {
		logger.Error("Error creating eigenlayer chain writer", "err", err)
		return err
	}

	_, err = elWriter.RegisterAsOperator(context.Background(), op, true)
	if err != nil {
		logger.Error("Error registering operator with eigenlayer", "err", err)
		return err
	}

	return nil
}

func RegisterForOperatorSets(
	operatorAddr common.Address,
	logger logging.Logger,
	elcontractsConfig elcontracts.Config,
	ethClient *ethclient.Client,
	txMgr txmgr.TxManager,
	registryCoordinatorAddr common.Address,
	avsAddress common.Address,
	operatorSetsIds []uint32,
	blsKeyPair bls.KeyPair,
	socket string,
) error {
	elWriter, err := elcontracts.NewWriterFromConfig(elcontractsConfig, ethClient, logger, &metrics.EigenMetrics{}, txMgr)
	if err != nil {
		logger.Error("Error creating eigenlayer chain writer", "err", err)
		return err
	}
	// Register operator for operator sets
	registrationRequest := elcontracts.RegistrationRequest{
		OperatorAddress: operatorAddr,
		AVSAddress:      avsAddress,
		OperatorSetIds:  operatorSetsIds,
		WaitForReceipt:  true,
		BlsKeyPair:      &blsKeyPair,
		Socket:          socket,
	}

	_, err = elWriter.RegisterForOperatorSets(
		context.Background(),
		registryCoordinatorAddr,
		registrationRequest,
	)

	if err != nil {
		logger.Errorf("Unable to register operator with the operator set")
		return err
	}
	logger.Info("Registered operator with operator set")

	return nil
}

func SetAllocationDelay(
	logger logging.Logger,
	operatorAddr common.Address,
	ethClient *ethclient.Client,
	allocationManagerAddr common.Address,
	txMgr txmgr.TxManager,
	delay uint32,
) error {
	txOpts, _ := txMgr.GetNoSendTxOpts()

	waitForReceipt := true
	allocationManagerContract, _ := allocationmanager.NewContractAllocationManager(allocationManagerAddr, ethClient)

	tx, err := allocationManagerContract.SetAllocationDelay(txOpts, operatorAddr, delay)
	if err != nil {
		return err
	}
	receipt, err := txMgr.Send(context.Background(), tx, waitForReceipt)
	if err != nil {
		return utils.WrapError("failed to send setAllocationDelay tx with err", err)
	}
	logger.Info(
		"tx successfully included for SetAllocationDelay  ",
		"txHash",
		receipt.TxHash.String(),
	)

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

func modifyAllocations(
	operatorAddr common.Address,
	allocationManagerAddr common.Address,
	avsAddress common.Address,
	strategies []common.Address,
	newMagnitudes []uint64,
	httpUrl string,
	txMgr txmgr.TxManager,
	id uint32,
	logger logging.Logger,
) error {
	txOpts, _ := txMgr.GetNoSendTxOpts()

	ethRpcClient, _ := ethclient.Dial(httpUrl)
	waitForReceipt := true
	allocationManagerContract, _ := allocationmanager.NewContractAllocationManager(allocationManagerAddr, ethRpcClient)
	operatorSet := allocationmanager.OperatorSet{Avs: avsAddress, Id: id}
	var allocations []allocationmanager.IAllocationManagerTypesAllocateParams
	allocations_1 := allocationmanager.IAllocationManagerTypesAllocateParams{
		OperatorSet:   operatorSet,
		Strategies:    strategies,
		NewMagnitudes: newMagnitudes,
	}
	allocations = append(allocations, allocations_1)
	tx, err := allocationManagerContract.ModifyAllocations(txOpts, operatorAddr, allocations)
	if err != nil {
		return err
	}
	receipt, err := txMgr.Send(context.Background(), tx, waitForReceipt)
	if err != nil {
		return utils.WrapError("failed to send modifyAllocations tx with err", err)
	}
	logger.Infof("tx successfully included for modifyAllocations. txHash: %v", receipt.TxHash.String())

	return nil
}

// TODO: move to "types" module
type OperatorSet allocationmanager.OperatorSet

type WatchOperatorActivatedOpts struct {
	Context context.Context
	// Operators to watch activation for.
	// If empty, function will return an error.
	Operators []common.Address
	// Operator sets to trigger on.
	// If empty, function will return an error.
	OperatorSets []OperatorSet
}

type OperatorActivated struct {
	Operator    common.Address
	OperatorSet OperatorSet
}

// TODO: handle both HTTP polling and WS subscribing
// Watches for operator activation requests.
// Returns a channel that will receive notice of operator activations. This includes
// some immediately after the subscription is created, for operators that are already active.
// Also returns a subscription that can be used to unsubscribe from the event.
func WatchOperatorActivated(opts *WatchOperatorActivatedOpts, ethClient bind.ContractBackend, allocationManagerAddr common.Address) (<-chan *OperatorActivated, event.Subscription, error) {
	// TODO: should we default to "all operators/opsets" if none are provided?
	if len(opts.OperatorSets) == 0 {
		return nil, nil, errors.New("no operator sets provided")
	}
	if len(opts.Operators) == 0 {
		return nil, nil, errors.New("no operators provided")
	}
	ctx := opts.Context

	allocationManagerContract, err := allocationmanager.NewContractAllocationManager(allocationManagerAddr, ethClient)
	if err != nil {
		return nil, nil, utils.WrapError("failed to create allocation manager contract", err)
	}

	watchOpts := &bind.WatchOpts{Context: ctx}
	eventsC := make(chan *allocationmanager.ContractAllocationManagerOperatorAddedToOperatorSet)
	sub, err := allocationManagerContract.WatchOperatorAddedToOperatorSet(watchOpts, eventsC, opts.Operators)
	if err != nil {
		return nil, nil, utils.WrapError("failed to watch operator added event", err)
	}

	shouldIncludeOpset := make(map[OperatorSet]bool)

	for _, operatorSet := range opts.OperatorSets {
		shouldIncludeOpset[operatorSet] = true
	}

	sink := make(chan *OperatorActivated)
	// Inspired by the code from WatchOperatorAddedToOperatorSet
	ourSub := event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		callOpts := &bind.CallOpts{Context: ctx}
		for _, operator := range opts.Operators {
			for _, operatorSet := range opts.OperatorSets {
				isMember, err := allocationManagerContract.IsMemberOfOperatorSet(callOpts, operator, allocationmanager.OperatorSet(operatorSet))
				if err != nil {
					return utils.WrapError("failed to call IsMemberOfOperatorSet", err)
				}
				if isMember {
					sink <- &OperatorActivated{Operator: operator, OperatorSet: operatorSet}
				}
			}
		}
		for {
			select {
			case operatorActivatedEvent := <-eventsC:
				operatorSet := OperatorSet(operatorActivatedEvent.OperatorSet)

				if !shouldIncludeOpset[operatorSet] {
					continue
				}

				event := new(OperatorActivated)
				event.Operator = operatorActivatedEvent.Operator
				event.OperatorSet = operatorSet

				select {
				case sink <- event:
				// this part was copied from the bindings
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	})
	return sink, ourSub, nil
}
