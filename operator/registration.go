package operator

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"os"
	"strings"

	allocationmanager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/AllocationManager"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	sdkecdsa "github.com/Layr-Labs/eigensdk-go/crypto/ecdsa"
	"github.com/Layr-Labs/eigensdk-go/signerv2"

	erc20mock "github.com/Layr-Labs/eigensdk-go/contracts/bindings/MockERC20"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/avsregistry"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/elcontracts"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/wallet"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/metrics"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/Layr-Labs/eigensdk-go/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
)

// This function performs startup operations for the operator, including:
//   - Register the operator in Eigenlayer
//   - Mint tokens for the operator in the received strategy
//   - Register the operator in the received operator sets
//   - Set the allocation delay as zero, to performs allocations immediatly
//   - Initialize allocations for the operator sets
func handleRegistration(
	logger logging.Logger,
	config *RegistrationConfig,
	operatorAddr common.Address,
	registryCoordinatorAddr common.Address,
	avsReader *avsregistry.ChainReader, // TODO: Change this for a bool or an addr
	ethHttpClient *ethclient.Client,
	blsKeyPair *bls.KeyPair,
) error {
	if config == nil {
		// We bubble the error all the way up instead of using logger.Fatal because logger.Fatal prints a huge stack
		// trace that hides the actual error message. This error msg is more explicit and doesn't require showing a
		// stack trace to the user.
		return fmt.Errorf(
			"RegistrationConfig is nil, no startup changes required ",
		)
	}

	// Registration set up
	elcontractsConfig := elcontracts.Config{
		DelegationManagerAddress:    config.DelegationManagerAddress,
		RewardsCoordinatorAddress:   config.RewardsCoordinatorAddress,
		PermissionControllerAddress: config.PermissionControllerAddress,
	}

	chainid, err := ethHttpClient.ChainID(context.Background())
	if err != nil {
		logger.Error("Cannot get chain id", "err", err)
		return err
	}

	var ecdsaPk *ecdsa.PrivateKey
	if config.EcdsaSignerCfg.PrivateKey == "" {
		logger.Info("ECDSA private key was nil, using the private key store path and password params...")
		ecdsaKeystorePassword := ""

		envPassword, ok := os.LookupEnv("OPERATOR_ECDSA_KEY_PASSWORD")
		if !ok {
			logger.Info("ECDSA keystore password was not set at env, reading value from config")
			if config.EcdsaSignerCfg.KeystorePassword != nil {
				ecdsaKeystorePassword = *config.EcdsaSignerCfg.KeystorePassword
			} else {
				logger.Warnf("ECDSA keystore password not found in config, using empty string")
			}
		} else {
			ecdsaKeystorePassword = envPassword
		}

		ecdsaPk, err = sdkecdsa.ReadKey(
			config.EcdsaSignerCfg.KeystorePath,
			ecdsaKeystorePassword,
		)
		if err != nil {
			return utils.WrapError("Failed to read the ECDSA private key from keystore", err)
		}
	} else {
		operatorEcdsaPkString := strings.TrimPrefix(config.EcdsaSignerCfg.PrivateKey, "0x")

		ecdsaPk, err = crypto.HexToECDSA(operatorEcdsaPkString)
		if err != nil {
			return utils.WrapError("Failed to convert hex key to ECDSA", err)
		}
	}

	signerV2, senderAddr, err := signerv2.SignerFromConfig(signerv2.Config{
		PrivateKey: ecdsaPk,
	}, chainid)
	if err != nil {
		logger.Fatalf(err.Error())
	}

	pkWallet, err := wallet.NewPrivateKeyWallet(ethHttpClient, signerV2, senderAddr, logger)
	if err != nil {
		return err
	}

	// TODO: Use NewSimpleTxManagerFromPrivateKey instead, that receives only the priv key
	txMgr := txmgr.NewSimpleTxManager(pkWallet, ethHttpClient, logger, senderAddr)

	// Register operator in EigenLayer

	// Check if operator was registered, if its not registered and register on startup flag is not set, then will fail.
	// If its not registered and should be registered on startup, make the registration.
	operatorIsRegistered, err := avsReader.IsOperatorRegistered(&bind.CallOpts{}, operatorAddr)
	if err != nil {
		logger.Error("Error checking if operator is registered", "err", err)
		return err
	}
	if !operatorIsRegistered {
		err = RegisterOperatorWithEigenlayer(
			operatorAddr,
			elcontractsConfig,
			ethHttpClient,
			logger,
			txMgr,
			config.MetadataUrl,
		)
		if err != nil {
			logger.Fatalf("Failed to register operator with EigenLayer on startup: %v", err.Error())
		}
	} else {
		logger.Info("Operator already registered to EigenLayer, skipped EL regisration")
	}

	operatorsStatusInQuorums, err := avsReader.GetOperatorsStakeInQuorumsAtCurrentBlock(&bind.CallOpts{}, types.QuorumNums{0})
	if err != nil {
		logger.Fatalf("Failed to get operator stake at registration: %v", err.Error())
	}
	operatorStatus := operatorsStatusInQuorums[0][0]

	var differenceToMint *big.Int
	differenceToMint = differenceToMint.Sub(operatorStatus.Stake, config.AmountToMint)
	if differenceToMint.Int64() > 0 {
		err = DepositIntoStrategyForOperator(
			logger,
			elcontractsConfig,
			ethHttpClient,
			config.StrategyAddrs,
			txMgr,
			operatorAddr,
			differenceToMint,
		)
		if err != nil {
			logger.Fatalf("Failed to deposit into strategy for operator on startup: %v", err.Error())
		}
	} else {
		logger.Info("Operator already have the required amount to min, skipped depositing into strategy for operator")
	}

	elReader, err := elcontracts.NewReaderFromConfig(elcontractsConfig, ethHttpClient, logger)
	if err != nil {
		logger.Error("Error creating eigenlayer chain writer", "err", err)
		return err
	}

	isOperatorRegisteredToQuorum, err := elReader.IsOperatorRegisteredWithOperatorSet(
		context.Background(),
		operatorAddr,
		allocationmanager.OperatorSet{
			Avs: config.AvsAddress,
			Id:  config.OperatorSetIds[0],
		},
	)
	if err != nil {
		logger.Fatalf("Failed to check if operator is registered to quorum at registration: %v", err.Error())
	}

	if !isOperatorRegisteredToQuorum {
		err = RegisterForOperatorSets(
			operatorAddr,
			logger,
			elcontractsConfig,
			ethHttpClient,
			txMgr,
			registryCoordinatorAddr,
			config.AvsAddress,
			config.OperatorSetIds,
			*blsKeyPair,
			config.Socket,
		)
		if err != nil {
			logger.Fatalf("Failed to register operator for operator sets on startup: %v", err.Error())
		}
	} else {
		logger.Info("Operator is already registered to the required operator sets, skipped quorum registration")
	}

	err = SetAllocationDelay(
		logger,
		operatorAddr,
		ethHttpClient,
		config.AllocationManagerAddr,
		txMgr,
		config.AllocationDelay,
	)
	if err != nil {
		logger.Fatalf("Failed to set allocation delay: %v", err.Error())
	}

	err = modifyAllocations(
		operatorAddr,
		config.AllocationManagerAddr,
		config.AvsAddress,
		config.StrategyAddrs,
		config.AllocatableMagnitudes,
		ethHttpClient,
		txMgr,
		config.OperatorSetIds,
		logger,
	)
	if err != nil {
		logger.Fatalf("Failed to set modify allocations: %v", err.Error())
	}

	return nil
}

// This function registers the operator with Eigenlayer. To do this needs the delegationManager
// address in the elcontracts config
func RegisterOperatorWithEigenlayer(
	operatorAddr common.Address,
	elcontractsConfig elcontracts.Config,
	ethClient *ethclient.Client,
	logger logging.Logger,
	txMgr txmgr.TxManager,
	metadataUrl string,
) error {
	op := types.Operator{
		Address:                   operatorAddr.String(),
		DelegationApproverAddress: operatorAddr.String(),
		MetadataUrl:               metadataUrl,
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

// This function registers the operator in the operator sets received as parameter. To do this needs the
// allocationManager address in the elcontracts config and the registryCoordinator address.
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

// This function sets the allocation delay for the operator to a value received as parameter (currently zero).
// To do this needs the allocationManager address.
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

// This function deposits into the token of the received strategies in an amount received as parameter.
// To do this needs the strategyManager address and the addresses of the strategies for the operator
// to be deposited.
func DepositIntoStrategyForOperator(
	logger logging.Logger,
	elcontractsConfig elcontracts.Config,
	ethClient *ethclient.Client,
	strategyAddrs []common.Address,
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

	for _, strategyAddr := range strategyAddrs {
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
	}

	return nil
}

// This function initializes the allocations for the operator, setting the allocatable magnitude (the available to slash).
// To do this needs the allocationManager address.
func modifyAllocations(
	operatorAddr common.Address,
	allocationManagerAddr common.Address,
	avsAddress common.Address,
	strategies []common.Address,
	newMagnitudes []uint64,
	ethHttpClient *ethclient.Client,
	txMgr txmgr.TxManager,
	operatorSetsIds []uint32,
	logger logging.Logger,
) error {
	txOpts, _ := txMgr.GetNoSendTxOpts()

	waitForReceipt := true
	allocationManagerContract, _ := allocationmanager.NewContractAllocationManager(allocationManagerAddr, ethHttpClient)

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
