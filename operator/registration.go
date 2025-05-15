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

func RegisterOperatorOnStartup(c RegistrationConfig) error {
	ethRpcClient, err := ethclient.Dial(c.EthHttpUrl)
	if err != nil {
		c.Logger.Errorf("Cannot create http ethclient", "err", err)
		return err
	}

	elcontractsConfig := elcontracts.Config{
		DelegationManagerAddress:    c.DelegationManagerAddress,
		RewardsCoordinatorAddress:   c.RewardsCoordinatorAddress,
		PermissionControllerAddress: c.PermissionControllerAddress,
	}

	rpcCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	chainid, err := ethRpcClient.ChainID(rpcCtx)
	if err != nil {
		c.Logger.Error("Cannot get chain id", "err", err)
		return err
	}

	ecdsaKeyPassword, ok := os.LookupEnv("OPERATOR_ECDSA_KEY_PASSWORD")
	if !ok {
		c.Logger.Warnf("OPERATOR_ECDSA_KEY_PASSWORD env var not set. using empty string")
	}

	operatorEcdsaPrivateKey, err := ecdsa.ReadKey(
		c.EcdsaKeyStorePath,
		ecdsaKeyPassword,
	)
	if err != nil {
		return err
	}

	signerV2, senderAddr, err := signerv2.SignerFromConfig(signerv2.Config{
		PrivateKey: operatorEcdsaPrivateKey,
	}, chainid)
	if err != nil {
		c.Logger.Fatalf(err.Error())
	}

	pkWallet, err := wallet.NewPrivateKeyWallet(ethRpcClient, signerV2, senderAddr, c.Logger)
	if err != nil {
		return err
	}

	txMgr := txmgr.NewSimpleTxManager(pkWallet, ethRpcClient, c.Logger, senderAddr)

	err = RegisterOperatorWithEigenlayer(
		c.OperatorAddr,
		elcontractsConfig,
		ethRpcClient,
		c.Logger,
		txMgr,
	)
	if err != nil {
		c.Logger.Fatalf("Failed to register operator with EigenLayer on startup: %v", err.Error())
	}

	err = DepositIntoStrategyForOperator(
		c.Logger,
		elcontractsConfig,
		ethRpcClient,
		c.StrategyAddr,
		txMgr,
		c.OperatorAddr,
		c.AmountToMint,
	)
	if err != nil {
		c.Logger.Fatalf("Failed to deposit into strategy for operator on startup: %v", err.Error())
	}

	blsKeyPassword, ok := os.LookupEnv("OPERATOR_BLS_KEY_PASSWORD")
	if !ok {
		c.Logger.Warnf("OPERATOR_BLS_KEY_PASSWORD env var not set. using empty string")
	}
	blsKeyPair, err := bls.ReadPrivateKeyFromFile(c.BlsKeyStorePath, blsKeyPassword)
	if err != nil {
		c.Logger.Errorf("Cannot parse bls private key", "err", err)
		return err
	}

	err = RegisterForOperatorSets(
		c.OperatorAddr,
		c.Logger,
		elcontractsConfig,
		ethRpcClient,
		txMgr,
		c.RegistryCoordinatorAddr,
		c.AvsAddress,
		c.OperatorSetIds,
		*blsKeyPair,
		"",
	)
	if err != nil {
		c.Logger.Fatalf("Failed to register operator for operator sets on startup: %v", err.Error())
	}

	err = SetAllocationDelay(
		c.Logger,
		c.OperatorAddr,
		ethRpcClient,
		c.AllocationManagerAddr,
		txMgr,
		0,
	)
	if err != nil {
		c.Logger.Fatalf("Failed to set allocation delay: %v", err.Error())
	}

	err = modifyAllocations(
		c.OperatorAddr,
		c.AllocationManagerAddr,
		c.AvsAddress,
		[]common.Address{c.StrategyAddr},
		[]uint64{c.AllocatableMagnitude},
		c.EthHttpUrl,
		txMgr,
		c.OperatorSetIds,
		c.Logger,
	)
	if err != nil {
		c.Logger.Fatalf("Failed to set modify allocations: %v", err.Error())
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
	operatorSetsIds []uint32,
	logger logging.Logger,
) error {
	txOpts, _ := txMgr.GetNoSendTxOpts()

	ethRpcClient, _ := ethclient.Dial(httpUrl)
	waitForReceipt := true
	allocationManagerContract, _ := allocationmanager.NewContractAllocationManager(allocationManagerAddr, ethRpcClient)

	allocations := []allocationmanager.IAllocationManagerTypesAllocateParams{}
	for _, setId := range operatorSetsIds {
		operatorSet := allocationmanager.OperatorSet{Avs: avsAddress, Id: setId}
		newAllocation := allocationmanager.IAllocationManagerTypesAllocateParams{
			OperatorSet:   operatorSet,
			Strategies:    strategies,
			NewMagnitudes: newMagnitudes,
		}
		allocations = append(allocations, newAllocation)
	}

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
