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
//   - Initialize allocations for the operator sets (if the registration config has defined the pameters needed)
//   - Set the allocation delay as to the value received in config (if the allocation delay value is set)
//   - Register the operator in the received operator sets (if it's not already registered to those operator sets)
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
	err = handleRegistrationWithEigenlayer(
		logger,
		elReader,
		elWriter,
		operatorAddr,
		config.MetadataUrl,
		config.AllocationDelay,
	)
	if err != nil {
		logger.Fatalf("Failed to register operator with EigenLayer on startup: %v", err.Error())
	}

	// Extract operator sets and deposits from config
	operatorSets := []allocationmanager.OperatorSet{}
	allDeposits := []DepositConfig{}
	for _, opSetConfig := range config.OperatorSetConfigs {
		operatorSets = append(operatorSets, allocationmanager.OperatorSet{
			Avs: config.AvsAddress,
			Id:  opSetConfig.ID,
		})
		allDeposits = append(allDeposits, opSetConfig.Deposits...)
	}

	// Deposit tokens into strategies for the operator
	err = handleDepositTokenAmount(
		txMgr,
		logger,
		ethHttpClient,
		elReader,
		elWriter,
		operatorAddr,
		allDeposits,
	)
	if err != nil {
		logger.Fatalf("Failed to deposit tokens into strategies on startup: %v", err.Error())
	}

	// Handle allocated stake for operator
	err = handleAllocatedStake(
		logger,
		elReader,
		elWriter,
		operatorAddr,
		config,
	)
	if err != nil {
		logger.Fatalf("Failed to allocate stake on startup: %v", err.Error())
	}

	err = handleAllocationDelay(
		logger,
		elReader,
		elWriter,
		operatorAddr,
	)
	if err != nil {
		logger.Fatalf("Failed to set allocation delay on startup: %v", err.Error())
	}

	err = handleRegistrationToOperatorSets(
		logger,
		elWriter,
		elReader,
		operatorAddr,
		registryCoordinatorAddr,
		*blsKeyPair,
		operatorSets,
		config.Socket,
	)
	if err != nil {
		logger.Fatalf("Failed to register to operetorSets on startup: %v", err.Error())
	}

	return nil
}

// Register operator in EigenLayer
// Check if operator was registered, if its not registered, try to register it
// If the operator is registered, skip the registration
func handleRegistrationWithEigenlayer(
	logger logging.Logger,
	elReader *elcontracts.ChainReader,
	elWriter *elcontracts.ChainWriter,
	operatorAddr common.Address,
	metadataUrl string,
	allocationDelay uint32,
) error {
	operatorIsRegistered, err := elReader.IsOperatorRegistered(
		context.Background(),
		types.Operator{Address: operatorAddr.Hex()},
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
			metadataUrl,
			allocationDelay,
		)
		if err != nil {
			logger.Fatalf("Failed to register operator with EigenLayer on startup: %v", err.Error())
		}
	} else {
		logger.Info("Operator already registered to EigenLayer, skipped EL regisration")
	}

	return nil
}

// This function registers the operator with Eigenlayer.
func registerOperatorWithEigenlayer(
	logger logging.Logger,
	elWriter *elcontracts.ChainWriter,
	operatorAddr common.Address,
	metadataUrl string,
	allocationDelay uint32,
) error {
	op := types.Operator{
		Address:                   operatorAddr.String(),
		DelegationApproverAddress: operatorAddr.String(),
		MetadataUrl:               metadataUrl,
		AllocationDelay:           allocationDelay,
	}

	_, err := elWriter.RegisterAsOperator(context.Background(), op, true)
	if err != nil {
		logger.Error("Error registering operator with eigenlayer", "err", err)
		return err
	}

	return nil
}

// Handles token deposits into EigenLayer strategies.
// This function ensures that the operator has deposited the required amounts of tokens
// into each specified strategy. If the operator has already deposited the required amount,
// skip the deposit. If not, deposit the difference between the required amount and the deposited amount.
func handleDepositTokenAmount(
	txMgr txmgr.TxManager,
	logger logging.Logger,
	ethClient *ethclient.Client,
	elReader *elcontracts.ChainReader,
	elWriter *elcontracts.ChainWriter,
	operatorAddr common.Address,
	depositConfig []DepositConfig,
) error {

	if len(depositConfig) == 0 {
		logger.Info("No deposits to make, skipping deposit token amount")
		return nil
	}

	// For each deposit, we get the operator shares in the strategy and check if the operator has deposited the required amount
	for _, deposit := range depositConfig {
		depositAmount, err := elReader.GetOperatorSharesInStrategy(context.Background(), operatorAddr, deposit.StrategyAddrs)
		if err != nil {
			logger.Errorf("Error getting operator shares in strategy: %x. Error: %v", deposit.StrategyAddrs, err)
			return err
		}
		logger.Infof("Operator has %v shares in strategy %x", depositAmount, deposit.StrategyAddrs)

		// If the operator has less shares than the amount to mint, we need to deposit the difference
		if depositAmount.Cmp(deposit.AmountToMint) == -1 {
			amountToDeposit := big.NewInt(0).Sub(deposit.AmountToMint, depositAmount)
			logger.Infof("Expected deposit amount: %v. Difference between expected and deposited amount: %v", deposit.AmountToMint, amountToDeposit)
			logger.Infof("Depositing %v tokens into strategy %x", amountToDeposit, deposit.StrategyAddrs)

			// TODO: Should we use GetStrategyAndUnderlyingToken or GetStrategyAndUnderlyingERC20Token?
			_, tokenAddr, err := elReader.GetStrategyAndUnderlyingToken(context.Background(), deposit.StrategyAddrs)
			if err != nil {
				logger.Error("Failed to fetch strategy contract", "err", err)
				return err
			}
			logger.Infof("Token address: %x", tokenAddr)

			// TODO: This is a mock contract, we need to use the real contract
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

			// Deposit the tokens into the strategy
			_, err = elWriter.DepositERC20IntoStrategy(context.Background(), deposit.StrategyAddrs, deposit.AmountToMint, true)
			if err != nil {
				logger.Errorf("Error depositing into strategy: %s. Error: %v", deposit.StrategyAddrs.String(), err)
				return err
			}

		} else {
			logger.Infof("Operator has enough shares in strategy %x, skipping deposit", deposit.StrategyAddrs)
		}
	}

	return nil
}

// Handles the allocated stake for the operator.
// This function ensures that the operator has the desired allocation in each strategy.
// If the operator has less allocation than the desired allocation, it will modify the allocations.
// In other case, it will skip the allocation modification.
func handleAllocatedStake(
	logger logging.Logger,
	elReader *elcontracts.ChainReader,
	elWriter *elcontracts.ChainWriter,
	operatorAddr common.Address,
	config *RegistrationConfig,
) error {

	if len(config.OperatorSetConfigs) == 0 {
		logger.Info("No operator set configs to handle, skipping allocated stake")
		return nil
	}

	allocateParams := []allocationmanager.IAllocationManagerTypesAllocateParams{}

	// For each operator set config, we get the allocated stake for the operator and the desired allocation
	for _, opSetConfig := range config.OperatorSetConfigs {
		operatorSet := allocationmanager.OperatorSet{
			Avs: config.AvsAddress,
			Id:  opSetConfig.ID,
		}

		// Strategies and expected magnitudes to batch the calls to the chain
		strategies := make([]common.Address, len(opSetConfig.Deposits))

		for i, deposit := range opSetConfig.Deposits {
			strategies[i] = deposit.StrategyAddrs
		}

		// Get the allocated stake for the strategies in the operator set
		allocated, err := elReader.GetAllocatedStake(context.Background(), operatorSet, []common.Address{operatorAddr}, strategies)
		if err != nil {
			logger.Errorf("Error getting allocated stake for operator set %v: %v", operatorSet.Id, err)
			return err
		}

		// Collect strategies and magnitudes that need updates
		strategiesToAllocate := []common.Address{}
		magnitudesToAllocate := []uint64{}

		// Check if the allocated stake is less than the desired allocation for each strategy
		// If it is, we add the strategy and magnitude to the allocateParams
		for i, deposit := range opSetConfig.Deposits {
			currentAllocatedStake := allocated[0][i]
			allocateMagnitude := big.NewInt(int64(deposit.AllocatableMagnitudes))

			logger.Infof(
				"Current allocated stake: %v, desired: %v for strategy %x in operator set %v",
				currentAllocatedStake, allocateMagnitude, deposit.StrategyAddrs, operatorSet.Id,
			)

			if currentAllocatedStake.Cmp(allocateMagnitude) == -1 {
				strategiesToAllocate = append(strategiesToAllocate, deposit.StrategyAddrs)
				magnitudesToAllocate = append(magnitudesToAllocate, deposit.AllocatableMagnitudes)
			}
		}

		if len(strategiesToAllocate) > 0 {
			allocateParams = append(allocateParams, allocationmanager.IAllocationManagerTypesAllocateParams{
				OperatorSet:   operatorSet,
				Strategies:    strategiesToAllocate,
				NewMagnitudes: magnitudesToAllocate,
			})
		}
	}

	// Execute batch allocation changes if any are needed
	if len(allocateParams) > 0 {
		logger.Infof("Modifying %v allocations", len(allocateParams))
		_, err := elWriter.ModifyAllocations(context.Background(), operatorAddr, allocateParams, true)
		if err != nil {
			logger.Errorf("Error modifying the allocations")
			return err
		}
	} else {
		logger.Info("All allocations are correct")
	}

	return nil
}

func handleAllocationDelay(
	logger logging.Logger,
	elReader *elcontracts.ChainReader,
	elWriter *elcontracts.ChainWriter,
	operatorAddr common.Address,
) error {
	allocationDelay, err := elReader.GetAllocationDelay(context.Background(), operatorAddr)
	if err != nil {
		logger.Errorf("Error getting allocation delay: %v", err)
		return err
	}

	if allocationDelay != 0 {
		logger.Infof("Allocation delay is %v, skipping allocation delay modification", allocationDelay)
		return nil
	}

	_, err = elWriter.SetAllocationDelay(context.Background(), operatorAddr, allocationDelay, true)
	if err != nil {
		logger.Errorf("Error setting allocation delay: %v", err)
		return err
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

func handleRegistrationToOperatorSets(
	logger logging.Logger,
	elWriter *elcontracts.ChainWriter,
	elReader *elcontracts.ChainReader,
	operatorAddr common.Address,
	registryCoordinatorAddr common.Address,
	blsKeyPair bls.KeyPair,
	operatorSets []allocationmanager.OperatorSet,
	socket string,
) error {
	if len(operatorSets) == 0 {
		logger.Info("No operator sets to register, skipping registration to operator sets")
		return nil
	}

	operatorSetsByAvs := map[common.Address][]uint32{}

	for _, operatorSet := range operatorSets {
		isOperatorRegisteredToQuorum, err := elReader.IsOperatorRegisteredWithOperatorSet(
			context.Background(),
			operatorAddr,
			allocationmanager.OperatorSet{
				Avs: operatorSet.Avs,
				Id:  operatorSet.Id,
			},
		)
		if err != nil {
			logger.Fatalf("Failed to check if operator is registered to quorum at registration: %v", err.Error())
		}

		if !isOperatorRegisteredToQuorum {
			logger.Infof("Operator set %x: %v ID requires registration", operatorSet.Avs, operatorSet.Id)
			operatorSetsByAvs[operatorSet.Avs] = append(operatorSetsByAvs[operatorSet.Avs], operatorSet.Id)
		} else {
			logger.Infof("Operator set %x: %v already registered", operatorSet.Avs, operatorSet.Id)
		}
	}

	for avsAddress, operatorSetIds := range operatorSetsByAvs {
		err := registerForOperatorSets(
			logger,
			operatorAddr,
			elWriter,
			registryCoordinatorAddr,
			avsAddress,
			operatorSetIds,
			blsKeyPair,
			socket,
		)
		if err != nil {
			logger.Fatalf("Failed to register operator for operator sets on startup: %v", err.Error())
		}
	}
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
