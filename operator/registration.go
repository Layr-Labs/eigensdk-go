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

	erc20mock "github.com/Layr-Labs/eigensdk-go/contracts/bindings/MockERC20"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/elcontracts"
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
//   - Register the operator in Eigenlayer (if it's not already registered)
//   - Mint tokens for the operator in the received strategy (if it has not the required ammount)
//   - Register the operator in the received operator sets (if it's not already registered to those operator sets)
//   - Set the allocation delay as to the value received in config (if the allocation delay value is set)
//   - Initialize allocations for the operator sets (if the registration config has defined the pameters needed)
func handleRegistration(
	logger logging.Logger,
	config *RegistrationConfig,
	operatorAddr common.Address,
	registryCoordinatorAddr common.Address,
	ethHttpClient *ethclient.Client,
	blsKeyPair *bls.KeyPair,
) error {
	// Required:
	// - DelegationManager (to check if operator is registered to EigenLayer)
	// - AllocationManager (to get the allocated stake for an operator and check if the operator is registered to an operator set)

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
		DelegationManagerAddress: config.DelegationManagerAddress,
	}

	elReader, err := elcontracts.NewReaderFromConfig(elcontractsConfig, ethHttpClient, logger)
	if err != nil {
		logger.Error("Error creating eigenlayer chain reader", "err", err)
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

	txMgr, err := txmgr.NewSimpleTxManagerFromPrivateKey(logger, ethHttpClient, ecdsaPk)
	if err != nil {
		logger.Errorf("Error creating tx managerfor registration: %v", err)
		return err
	}

	elWriter, err := elcontracts.NewWriterFromConfig(elcontractsConfig, ethHttpClient, logger, &metrics.EigenMetrics{}, txMgr)
	if err != nil {
		logger.Error("Error creating eigenlayer chain writer", "err", err)
		return err
	}

	// Register operator in EigenLayer
	// Check if operator was registered, if its not registered and register on startup flag is not set, then will fail.
	// If its not registered and should be registered on startup, make the registration.
	operatorIsRegistered, err := elReader.IsOperatorRegistered(
		context.Background(),
		types.Operator{Address: operatorAddr.Hex()}, // TODO: We are turning this into string to
	)
	if err != nil {
		logger.Error("Error checking if operator is registered", "err", err)
		return err
	}
	if !operatorIsRegistered {
		err = registerOperatorWithEigenlayer(
			logger,
			elWriter,
			operatorAddr,
			config.MetadataUrl,
		)
		if err != nil {
			logger.Fatalf("Failed to register operator with EigenLayer on startup: %v", err.Error())
		}
	} else {
		logger.Info("Operator already registered to EigenLayer, skipped EL regisration")
	}

	// Deposit tokens into strategies for the operator
	err = handleDepositTokenAmount(
		txMgr,
		logger,
		ethHttpClient,
		elReader,
		elWriter,
		operatorAddr,
		config.DepositConfig,
	)
	if err != nil {
		logger.Fatalf("Failed to deposit tokens into strategies for operator on startup: %v", err.Error())
	}

	// operatorSet := allocationmanager.OperatorSet{
	// 	Avs: config.AvsAddress,
	// 	Id:  config.OperatorSetIds[0],
	// }
	// allocatedStakes, err := elReader.GetAllocatedStake(context.Background(), operatorSet, []common.Address{operatorAddr}, config.StrategyAddrs)
	// if err != nil {
	// 	logger.Fatalf("Failed to get operator stake at registration: %v", err.Error())
	// }
	// allocatedStake := allocatedStakes[0][0]

	// differenceToMint := big.NewInt(0)
	// differenceToMint = differenceToMint.Sub(config.AmountToMint, allocatedStake)
	// if differenceToMint.Int64() > 0 {
	// 	err = depositIntoStrategyForOperator(
	// 		logger,
	// 		elReader,
	// 		elWriter,
	// 		ethHttpClient,
	// 		config.StrategyAddrs,
	// 		txMgr,
	// 		operatorAddr,
	// 		differenceToMint,
	// 	)
	// 	if err != nil {
	// 		logger.Fatalf("Failed to deposit into strategy for operator on startup: %v", err.Error())
	// 	}
	// } else {
	// 	logger.Info("Operator already have the required amount to min, skipped depositing into strategy for operator")
	// }

	// isOperatorRegisteredToQuorum, err := elReader.IsOperatorRegisteredWithOperatorSet(
	// 	context.Background(),
	// 	operatorAddr,
	// 	allocationmanager.OperatorSet{
	// 		Avs: config.AvsAddress,
	// 		Id:  config.OperatorSetIds[0],
	// 	},
	// )
	// if err != nil {
	// 	logger.Fatalf("Failed to check if operator is registered to quorum at registration: %v", err.Error())
	// }

	// if !isOperatorRegisteredToQuorum {
	// 	err = registerForOperatorSets(
	// 		logger,
	// 		operatorAddr,
	// 		elWriter,
	// 		registryCoordinatorAddr,
	// 		config.AvsAddress,
	// 		config.OperatorSetIds,
	// 		*blsKeyPair,
	// 		config.Socket,
	// 	)
	// 	if err != nil {
	// 		logger.Fatalf("Failed to register operator for operator sets on startup: %v", err.Error())
	// 	}
	// } else {
	// 	logger.Info("Operator is already registered to the required operator sets, skipped quorum registration")
	// }

	// err = setAllocationDelay(
	// 	logger,
	// 	elWriter,
	// 	operatorAddr,
	// 	config.AllocationDelay,
	// )
	// if err != nil {
	// 	logger.Fatalf("Failed to set allocation delay: %v", err.Error())
	// }

	// err = modifyAllocations(
	// 	logger,
	// 	elWriter,
	// 	operatorAddr,
	// 	config.AvsAddress,
	// 	config.StrategyAddrs,
	// 	config.AllocatableMagnitudes,
	// 	config.OperatorSetIds,
	// )
	// if err != nil {
	// 	logger.Fatalf("Failed to set modify allocations: %v", err.Error())
	// }

	return nil
}

// This function registers the operator with Eigenlayer. To do this needs the delegationManager
// address in the elcontracts config
func registerOperatorWithEigenlayer(
	logger logging.Logger,
	elWriter *elcontracts.ChainWriter,
	operatorAddr common.Address,
	metadataUrl string,
) error {
	op := types.Operator{
		Address:                   operatorAddr.String(),
		DelegationApproverAddress: operatorAddr.String(),
		MetadataUrl:               metadataUrl,
	}

	_, err := elWriter.RegisterAsOperator(context.Background(), op, true)
	if err != nil {
		logger.Error("Error registering operator with eigenlayer", "err", err)
		return err
	}

	return nil
}

func handleDepositTokenAmount(
	txMgr txmgr.TxManager,
	logger logging.Logger,
	ethClient *ethclient.Client,
	elReader *elcontracts.ChainReader,
	elWriter *elcontracts.ChainWriter,
	operatorAddr common.Address,
	depositConfig []DepositConfig,
) error {
	for _, deposit := range depositConfig {
		depositAmount, err := elReader.GetOperatorSharesInStrategy(context.Background(), operatorAddr, deposit.StrategyAddrs)
		if err != nil {
			logger.Errorf("Error getting operator shares in strategy: %s. Error: %v", deposit.StrategyAddrs.String(), err)
			return err
		}
		logger.Infof("Operator has %v shares in strategy %x", depositAmount, deposit.StrategyAddrs.String())

		if depositAmount.Cmp(deposit.AmountToMint) == -1 {
			amountToDeposit := big.NewInt(0).Sub(deposit.AmountToMint, depositAmount)
			logger.Infof("Expected deposit amount: %v. Difference between expected and deposited amount: %v", deposit.AmountToMint, amountToDeposit)
			logger.Infof("Depositing %v tokens into strategy %x", amountToDeposit, deposit.StrategyAddrs.String())

			_, tokenAddr, err := elReader.GetStrategyAndUnderlyingToken(context.Background(), deposit.StrategyAddrs)
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

			tx, err := contractErc20Mock.Mint(txOpts, operatorAddr, deposit.AmountToMint)
			if err != nil {
				logger.Errorf("Error assembling Mint tx")
				return err
			}
			_, err = txMgr.Send(context.Background(), tx, true)
			if err != nil {
				logger.Errorf("Error submitting Mint tx")
				return err
			}

			_, err = elWriter.DepositERC20IntoStrategy(context.Background(), deposit.StrategyAddrs, deposit.AmountToMint, true)
			if err != nil {
				logger.Errorf("Error depositing into strategy: %s. Error: %v", deposit.StrategyAddrs.String(), err)
				return err
			}

		} else {
			logger.Infof("Operator has enough shares in strategy %x, skipping deposit", deposit.StrategyAddrs.String())
		}
	}

	return nil
}

// This function registers the operator in the operator sets received as parameter. To do this needs the
// allocationManager address in the elcontracts config and the registryCoordinator address.
func registerForOperatorSets(
	logger logging.Logger,
	operatorAddr common.Address,
	elWriter *elcontracts.ChainWriter,
	registryCoordinatorAddr common.Address,
	avsAddress common.Address,
	operatorSetsIds []uint32,
	blsKeyPair bls.KeyPair,
	socket string,
) error {
	// Register operator for operator sets
	registrationRequest := elcontracts.RegistrationRequest{
		OperatorAddress: operatorAddr,
		AVSAddress:      avsAddress,
		OperatorSetIds:  operatorSetsIds,
		WaitForReceipt:  true,
		BlsKeyPair:      &blsKeyPair,
		Socket:          socket,
	}

	_, err := elWriter.RegisterForOperatorSets(
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
func setAllocationDelay(
	logger logging.Logger,
	elWriter *elcontracts.ChainWriter,
	operatorAddr common.Address,
	delay uint32,
) error {
	receipt, err := elWriter.SetAllocationDelay(context.Background(), operatorAddr, delay, true)
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
// func depositIntoStrategyForOperator(
// 	logger logging.Logger,
// 	elReader *elcontracts.ChainReader,
// 	elWriter *elcontracts.ChainWriter,
// 	ethClient *ethclient.Client,
// 	operatorAddr common.Address,
// 	amount *big.Int,
// ) error {
// 	for _, strategyAddr := range strategyAddrs {
// 		_, tokenAddr, err := elReader.GetStrategyAndUnderlyingToken(context.Background(), strategyAddr)
// 		if err != nil {
// 			logger.Error("Failed to fetch strategy contract", "err", err)
// 			return err
// 		}
// 		logger.Info(tokenAddr.String())

// 		contractErc20Mock, err := erc20mock.NewContractMockERC20(tokenAddr, ethClient)
// 		if err != nil {
// 			logger.Error("Failed to fetch ERC20Mock contract", "err", err)
// 			return err
// 		}
// 		txOpts, err := txMgr.GetNoSendTxOpts()
// 		if err != nil {
// 			logger.Errorf("Error in GetNoSendTxOpts")
// 			return err
// 		}

// 		tx, err := contractErc20Mock.Mint(txOpts, operatorAddr, deposit.AmountToMint)
// 		if err != nil {
// 			logger.Errorf("Error assembling Mint tx")
// 			return err
// 		}
// 		_, err = txMgr.Send(context.Background(), tx, true)
// 		if err != nil {
// 			logger.Errorf("Error submitting Mint tx")
// 			return err
// 		}

// 		_, err = elWriter.DepositERC20IntoStrategy(context.Background(), deposit.StrategyAddrs, deposit.AmountToMint, true)
// 		if err != nil {
// 			logger.Errorf("Error depositing into strategy", "err", err)
// 			return err
// 		}
// 	}

// 	return nil
// }

// This function initializes the allocations for the operator, setting the allocatable magnitude (the available to slash).
// To do this needs the allocationManager address.
func modifyAllocations(
	logger logging.Logger,
	elWriter *elcontracts.ChainWriter,
	operatorAddr common.Address,
	avsAddress common.Address,
	strategies []common.Address,
	newMagnitudes []uint64,
	operatorSetsIds []uint32,
) error {
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

	receipt, err := elWriter.ModifyAllocations(context.Background(), operatorAddr, allocations, true)
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
