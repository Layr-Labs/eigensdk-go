package elcontracts

import (
	"errors"

	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	gethcommon "github.com/ethereum/go-ethereum/common"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	delegationmanager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/DelegationManager"
	avsdirectory "github.com/Layr-Labs/eigensdk-go/contracts/bindings/IAVSDirectory"
	erc20 "github.com/Layr-Labs/eigensdk-go/contracts/bindings/IERC20"
	rewardscoordinator "github.com/Layr-Labs/eigensdk-go/contracts/bindings/IRewardsCoordinator"
	slasher "github.com/Layr-Labs/eigensdk-go/contracts/bindings/ISlasher"
	strategy "github.com/Layr-Labs/eigensdk-go/contracts/bindings/IStrategy"
	strategymanager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/StrategyManager"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/Layr-Labs/eigensdk-go/utils"
)

type Config struct {
	DelegationManagerAddress  common.Address
	AvsDirectoryAddress       common.Address
	RewardsCoordinatorAddress common.Address
}

type ChainReader struct {
	logger             logging.Logger
	slasher            slasher.ContractISlasherCalls
	delegationManager  *delegationmanager.ContractDelegationManager
	strategyManager    *strategymanager.ContractStrategyManager
	avsDirectory       *avsdirectory.ContractIAVSDirectory
	rewardsCoordinator *rewardscoordinator.ContractIRewardsCoordinator
	ethClient          eth.HttpBackend
}

func NewChainReader(
	slasher slasher.ContractISlasherCalls,
	delegationManager *delegationmanager.ContractDelegationManager,
	strategyManager *strategymanager.ContractStrategyManager,
	avsDirectory *avsdirectory.ContractIAVSDirectory,
	rewardsCoordinator *rewardscoordinator.ContractIRewardsCoordinator,
	logger logging.Logger,
	ethClient eth.HttpBackend,
) *ChainReader {
	logger = logger.With(logging.ComponentKey, "elcontracts/reader")

	return &ChainReader{
		slasher:            slasher,
		delegationManager:  delegationManager,
		strategyManager:    strategyManager,
		avsDirectory:       avsDirectory,
		rewardsCoordinator: rewardsCoordinator,
		logger:             logger,
		ethClient:          ethClient,
	}
}

// BuildELChainReader creates a new ELChainReader
// Deprecated: Use BuildFromConfig instead
func BuildELChainReader(
	delegationManagerAddr gethcommon.Address,
	avsDirectoryAddr gethcommon.Address,
	ethClient eth.HttpBackend,
	logger logging.Logger,
) (*ChainReader, error) {
	elContractBindings, err := NewEigenlayerContractBindings(
		delegationManagerAddr,
		avsDirectoryAddr,
		ethClient,
		logger,
	)
	if err != nil {
		return nil, err
	}
	return NewChainReader(
		elContractBindings.Slasher,
		elContractBindings.DelegationManager,
		elContractBindings.StrategyManager,
		elContractBindings.AvsDirectory,
		elContractBindings.RewardsCoordinator,
		logger,
		ethClient,
	), nil
}

func NewReaderFromConfig(
	cfg Config,
	ethClient eth.HttpBackend,
	logger logging.Logger,
) (*ChainReader, error) {
	elContractBindings, err := NewBindingsFromConfig(
		cfg,
		ethClient,
		logger,
	)
	if err != nil {
		return nil, err
	}
	return NewChainReader(
		elContractBindings.Slasher,
		elContractBindings.DelegationManager,
		elContractBindings.StrategyManager,
		elContractBindings.AvsDirectory,
		elContractBindings.RewardsCoordinator,
		logger,
		ethClient,
	), nil
}

func (r *ChainReader) IsOperatorRegistered(opts *bind.CallOpts, operator types.Operator) (bool, error) {
	if r.delegationManager == nil {
		return false, errors.New("DelegationManager contract not provided")
	}

	isOperator, err := r.delegationManager.IsOperator(
		opts,
		gethcommon.HexToAddress(operator.Address),
	)
	if err != nil {
		return false, err
	}

	return isOperator, nil
}

func (r *ChainReader) GetOperatorDetails(opts *bind.CallOpts, operator types.Operator) (types.Operator, error) {
	if r.delegationManager == nil {
		return types.Operator{}, errors.New("DelegationManager contract not provided")
	}

	operatorDetails, err := r.delegationManager.OperatorDetails(
		opts,
		gethcommon.HexToAddress(operator.Address),
	)
	if err != nil {
		return types.Operator{}, err
	}

	return types.Operator{
		Address:                   operator.Address,
		StakerOptOutWindowBlocks:  operatorDetails.StakerOptOutWindowBlocks,
		DelegationApproverAddress: operatorDetails.DelegationApprover.Hex(),
	}, nil
}

// GetStrategyAndUnderlyingToken returns the strategy contract and the underlying token address
func (r *ChainReader) GetStrategyAndUnderlyingToken(
	opts *bind.CallOpts, strategyAddr gethcommon.Address,
) (*strategy.ContractIStrategy, gethcommon.Address, error) {
	contractStrategy, err := strategy.NewContractIStrategy(strategyAddr, r.ethClient)
	if err != nil {
		return nil, common.Address{}, utils.WrapError("Failed to fetch strategy contract", err)
	}
	underlyingTokenAddr, err := contractStrategy.UnderlyingToken(opts)
	if err != nil {
		return nil, common.Address{}, utils.WrapError("Failed to fetch token contract", err)
	}
	return contractStrategy, underlyingTokenAddr, nil
}

// GetStrategyAndUnderlyingERC20Token returns the strategy contract, the erc20 bindings for the underlying token
// and the underlying token address
func (r *ChainReader) GetStrategyAndUnderlyingERC20Token(
	opts *bind.CallOpts, strategyAddr gethcommon.Address,
) (*strategy.ContractIStrategy, erc20.ContractIERC20Methods, gethcommon.Address, error) {
	contractStrategy, err := strategy.NewContractIStrategy(strategyAddr, r.ethClient)
	if err != nil {
		return nil, nil, common.Address{}, utils.WrapError("Failed to fetch strategy contract", err)
	}
	underlyingTokenAddr, err := contractStrategy.UnderlyingToken(opts)
	if err != nil {
		return nil, nil, common.Address{}, utils.WrapError("Failed to fetch token contract", err)
	}
	contractUnderlyingToken, err := erc20.NewContractIERC20(underlyingTokenAddr, r.ethClient)
	if err != nil {
		return nil, nil, common.Address{}, utils.WrapError("Failed to fetch token contract", err)
	}
	return contractStrategy, contractUnderlyingToken, underlyingTokenAddr, nil
}

func (r *ChainReader) ServiceManagerCanSlashOperatorUntilBlock(
	opts *bind.CallOpts,
	operatorAddr gethcommon.Address,
	serviceManagerAddr gethcommon.Address,
) (uint32, error) {
	if r.slasher == nil {
		return uint32(0), errors.New("slasher contract not provided")
	}

	return r.slasher.ContractCanSlashOperatorUntilBlock(
		opts, operatorAddr, serviceManagerAddr,
	)
}

func (r *ChainReader) OperatorIsFrozen(opts *bind.CallOpts, operatorAddr gethcommon.Address) (bool, error) {
	if r.slasher == nil {
		return false, errors.New("slasher contract not provided")
	}

	return r.slasher.IsFrozen(opts, operatorAddr)
}

func (r *ChainReader) GetOperatorSharesInStrategy(
	opts *bind.CallOpts,
	operatorAddr gethcommon.Address,
	strategyAddr gethcommon.Address,
) (*big.Int, error) {
	if r.delegationManager == nil {
		return &big.Int{}, errors.New("DelegationManager contract not provided")
	}

	return r.delegationManager.OperatorShares(
		opts,
		operatorAddr,
		strategyAddr,
	)
}

func (r *ChainReader) CalculateDelegationApprovalDigestHash(
	opts *bind.CallOpts, staker gethcommon.Address, operator gethcommon.Address,
	delegationApprover gethcommon.Address, approverSalt [32]byte, expiry *big.Int,
) ([32]byte, error) {
	if r.delegationManager == nil {
		return [32]byte{}, errors.New("DelegationManager contract not provided")
	}

	return r.delegationManager.CalculateDelegationApprovalDigestHash(
		opts, staker, operator, delegationApprover, approverSalt, expiry,
	)
}

func (r *ChainReader) CalculateOperatorAVSRegistrationDigestHash(
	opts *bind.CallOpts, operator gethcommon.Address, avs gethcommon.Address, salt [32]byte, expiry *big.Int,
) ([32]byte, error) {
	if r.avsDirectory == nil {
		return [32]byte{}, errors.New("AVSDirectory contract not provided")
	}

	return r.avsDirectory.CalculateOperatorAVSRegistrationDigestHash(
		opts, operator, avs, salt, expiry,
	)
}

func (r *ChainReader) GetDistributionRootsLength(opts *bind.CallOpts) (*big.Int, error) {
	if r.rewardsCoordinator == nil {
		return nil, errors.New("RewardsCoordinator contract not provided")
	}

	return r.rewardsCoordinator.GetDistributionRootsLength(opts)
}

func (r *ChainReader) CurrRewardsCalculationEndTimestamp(opts *bind.CallOpts) (uint32, error) {
	if r.rewardsCoordinator == nil {
		return 0, errors.New("RewardsCoordinator contract not provided")
	}

	return r.rewardsCoordinator.CurrRewardsCalculationEndTimestamp(opts)
}

func (r *ChainReader) GetCurrentClaimableDistributionRoot(
	opts *bind.CallOpts,
) (rewardscoordinator.IRewardsCoordinatorDistributionRoot, error) {
	if r.rewardsCoordinator == nil {
		return rewardscoordinator.IRewardsCoordinatorDistributionRoot{}, errors.New(
			"RewardsCoordinator contract not provided",
		)
	}

	return r.rewardsCoordinator.GetCurrentClaimableDistributionRoot(opts)
}

func (r *ChainReader) GetRootIndexFromHash(opts *bind.CallOpts, rootHash [32]byte) (uint32, error) {
	if r.rewardsCoordinator == nil {
		return 0, errors.New("RewardsCoordinator contract not provided")
	}

	return r.rewardsCoordinator.GetRootIndexFromHash(opts, rootHash)
}
