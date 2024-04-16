package elcontracts

import (
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	gethcommon "github.com/ethereum/go-ethereum/common"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	chainioutils "github.com/Layr-Labs/eigensdk-go/chainio/utils"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/Layr-Labs/eigensdk-go/utils"

	avsdirectory "github.com/Layr-Labs/eigensdk-go/contracts/bindings/AVSDirectory"
	delegationmanager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/DelegationManager"
	erc20 "github.com/Layr-Labs/eigensdk-go/contracts/bindings/IERC20"
	slasher "github.com/Layr-Labs/eigensdk-go/contracts/bindings/ISlasher"
	strategy "github.com/Layr-Labs/eigensdk-go/contracts/bindings/IStrategy"
	strategymanager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/StrategyManager"
)

type ELReader interface {
	IsOperatorRegistered(opts *bind.CallOpts, operator types.Operator) (bool, error)

	GetOperatorDetails(opts *bind.CallOpts, operator types.Operator) (types.Operator, error)

	// GetStrategyAndUnderlyingToken returns the strategy contract and the underlying token address
	// use GetStrategyAndUnderlyingERC20Token if the contract address confirms with ERC20 standard
	GetStrategyAndUnderlyingToken(
		opts *bind.CallOpts, strategyAddr gethcommon.Address,
	) (*strategy.ContractIStrategy, gethcommon.Address, error)

	// GetStrategyAndUnderlyingERC20Token returns the strategy contract and the underlying ERC20 token address
	GetStrategyAndUnderlyingERC20Token(
		opts *bind.CallOpts, strategyAddr gethcommon.Address,
	) (*strategy.ContractIStrategy, erc20.ContractIERC20Methods, gethcommon.Address, error)

	ServiceManagerCanSlashOperatorUntilBlock(
		opts *bind.CallOpts,
		operatorAddr gethcommon.Address,
		serviceManagerAddr gethcommon.Address,
	) (uint32, error)

	OperatorIsFrozen(opts *bind.CallOpts, operatorAddr gethcommon.Address) (bool, error)

	GetOperatorSharesInStrategy(
		opts *bind.CallOpts,
		operatorAddr gethcommon.Address,
		strategyAddr gethcommon.Address,
	) (*big.Int, error)

	CalculateDelegationApprovalDigestHash(
		opts *bind.CallOpts, staker gethcommon.Address, operator gethcommon.Address,
		delegationApprover gethcommon.Address, approverSalt [32]byte, expiry *big.Int,
	) ([32]byte, error)

	CalculateOperatorAVSRegistrationDigestHash(
		opts *bind.CallOpts, operator gethcommon.Address, avs gethcommon.Address, salt [32]byte, expiry *big.Int,
	) ([32]byte, error)
}

type ELChainReader struct {
	logger            logging.Logger
	slasher           slasher.ContractISlasherCalls
	delegationManager delegationmanager.ContractDelegationManagerCalls
	strategyManager   strategymanager.ContractStrategyManagerCalls
	avsDirectory      avsdirectory.ContractAVSDirectoryCalls
	ethClient         eth.Client
}

// forces EthReader to implement the chainio.Reader interface
var _ ELReader = (*ELChainReader)(nil)

func NewELChainReader(
	slasher slasher.ContractISlasherCalls,
	delegationManager delegationmanager.ContractDelegationManagerCalls,
	strategyManager strategymanager.ContractStrategyManagerCalls,
	avsDirectory avsdirectory.ContractAVSDirectoryCalls,
	logger logging.Logger,
	ethClient eth.Client,
) *ELChainReader {
	return &ELChainReader{
		slasher:           slasher,
		delegationManager: delegationManager,
		strategyManager:   strategyManager,
		avsDirectory:      avsDirectory,
		logger:            logger,
		ethClient:         ethClient,
	}
}

func BuildELChainReader(
	delegationManagerAddr gethcommon.Address,
	avsDirectoryAddr gethcommon.Address,
	ethClient eth.Client,
	logger logging.Logger,
) (*ELChainReader, error) {
	elContractBindings, err := chainioutils.NewEigenlayerContractBindings(
		delegationManagerAddr,
		avsDirectoryAddr,
		ethClient,
		logger,
	)
	if err != nil {
		return nil, err
	}
	return NewELChainReader(
		elContractBindings.Slasher,
		elContractBindings.DelegationManager,
		elContractBindings.StrategyManager,
		elContractBindings.AvsDirectory,
		logger,
		ethClient,
	), nil
}

func (r *ELChainReader) IsOperatorRegistered(opts *bind.CallOpts, operator types.Operator) (bool, error) {
	isOperator, err := r.delegationManager.IsOperator(
		opts,
		gethcommon.HexToAddress(operator.Address),
	)
	if err != nil {
		return false, err
	}

	return isOperator, nil
}

func (r *ELChainReader) GetOperatorDetails(opts *bind.CallOpts, operator types.Operator) (types.Operator, error) {
	operatorDetails, err := r.delegationManager.OperatorDetails(
		opts,
		gethcommon.HexToAddress(operator.Address),
	)
	if err != nil {
		return types.Operator{}, err
	}

	return types.Operator{
		Address:                   operator.Address,
		EarningsReceiverAddress:   operatorDetails.EarningsReceiver.Hex(),
		StakerOptOutWindowBlocks:  operatorDetails.StakerOptOutWindowBlocks,
		DelegationApproverAddress: operatorDetails.DelegationApprover.Hex(),
	}, nil
}

// GetStrategyAndUnderlyingToken returns the strategy contract and the underlying token address
func (r *ELChainReader) GetStrategyAndUnderlyingToken(
	opts *bind.CallOpts, strategyAddr gethcommon.Address,
) (*strategy.ContractIStrategy, gethcommon.Address, error) {
	contractStrategy, err := strategy.NewContractIStrategy(strategyAddr, r.ethClient)
	if err != nil {
		return nil, common.Address{}, utils.WrapError(errors.New("Failed to fetch strategy contract"), err)
	}
	underlyingTokenAddr, err := contractStrategy.UnderlyingToken(opts)
	if err != nil {
		return nil, common.Address{}, utils.WrapError(errors.New("Failed to fetch token contract"), err)
	}
	return contractStrategy, underlyingTokenAddr, nil
}

// GetStrategyAndUnderlyingERC20Token returns the strategy contract, the erc20 bindings for the underlying token
// and the underlying token address
func (r *ELChainReader) GetStrategyAndUnderlyingERC20Token(
	opts *bind.CallOpts, strategyAddr gethcommon.Address,
) (*strategy.ContractIStrategy, erc20.ContractIERC20Methods, gethcommon.Address, error) {
	contractStrategy, err := strategy.NewContractIStrategy(strategyAddr, r.ethClient)
	if err != nil {
		return nil, nil, common.Address{}, utils.WrapError(errors.New("Failed to fetch strategy contract"), err)
	}
	underlyingTokenAddr, err := contractStrategy.UnderlyingToken(opts)
	if err != nil {
		return nil, nil, common.Address{}, utils.WrapError(errors.New("Failed to fetch token contract"), err)
	}
	contractUnderlyingToken, err := erc20.NewContractIERC20(underlyingTokenAddr, r.ethClient)
	if err != nil {
		return nil, nil, common.Address{}, utils.WrapError(errors.New("Failed to fetch token contract"), err)
	}
	return contractStrategy, contractUnderlyingToken, underlyingTokenAddr, nil
}

func (r *ELChainReader) ServiceManagerCanSlashOperatorUntilBlock(
	opts *bind.CallOpts,
	operatorAddr gethcommon.Address,
	serviceManagerAddr gethcommon.Address,
) (uint32, error) {
	serviceManagerCanSlashOperatorUntilBlock, err := r.slasher.ContractCanSlashOperatorUntilBlock(
		opts, operatorAddr, serviceManagerAddr,
	)
	if err != nil {
		return 0, err
	}
	return serviceManagerCanSlashOperatorUntilBlock, nil
}

func (r *ELChainReader) OperatorIsFrozen(opts *bind.CallOpts, operatorAddr gethcommon.Address) (bool, error) {
	operatorIsFrozen, err := r.slasher.IsFrozen(opts, operatorAddr)
	if err != nil {
		return false, err
	}
	return operatorIsFrozen, nil
}

func (r *ELChainReader) GetOperatorSharesInStrategy(
	opts *bind.CallOpts,
	operatorAddr gethcommon.Address,
	strategyAddr gethcommon.Address,
) (*big.Int, error) {
	operatorSharesInStrategy, err := r.delegationManager.OperatorShares(
		opts,
		operatorAddr,
		strategyAddr,
	)
	if err != nil {
		return nil, err
	}
	return operatorSharesInStrategy, nil
}

func (r *ELChainReader) CalculateDelegationApprovalDigestHash(
	opts *bind.CallOpts, staker gethcommon.Address, operator gethcommon.Address,
	delegationApprover gethcommon.Address, approverSalt [32]byte, expiry *big.Int,
) ([32]byte, error) {
	return r.delegationManager.CalculateDelegationApprovalDigestHash(
		opts, staker, operator, delegationApprover, approverSalt, expiry,
	)
}

func (r *ELChainReader) CalculateOperatorAVSRegistrationDigestHash(
	opts *bind.CallOpts, operator gethcommon.Address, avs gethcommon.Address, salt [32]byte, expiry *big.Int,
) ([32]byte, error) {
	return r.avsDirectory.CalculateOperatorAVSRegistrationDigestHash(
		opts, operator, avs, salt, expiry,
	)
}
