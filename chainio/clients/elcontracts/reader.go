package elcontracts

import (
	"errors"
	"math"

	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethcommon "github.com/ethereum/go-ethereum/common"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	avsdirectory "github.com/Layr-Labs/eigensdk-go/contracts/bindings/AVSDirectory"
	allocationmanager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/AllocationManager"
	delegationmanager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/DelegationManager"
	erc20 "github.com/Layr-Labs/eigensdk-go/contracts/bindings/IERC20"
	rewardscoordinator "github.com/Layr-Labs/eigensdk-go/contracts/bindings/IRewardsCoordinator"
	strategy "github.com/Layr-Labs/eigensdk-go/contracts/bindings/IStrategy"
	strategymanager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/StrategyManager"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/Layr-Labs/eigensdk-go/utils"
)

type Config struct {
	DelegationManagerAddress  gethcommon.Address
	AvsDirectoryAddress       gethcommon.Address
	RewardsCoordinatorAddress gethcommon.Address
}

type ChainReader struct {
	logger             logging.Logger
	delegationManager  *delegationmanager.ContractDelegationManager
	strategyManager    *strategymanager.ContractStrategyManager
	avsDirectory       *avsdirectory.ContractAVSDirectory
	rewardsCoordinator *rewardscoordinator.ContractIRewardsCoordinator
	allocationManager  *allocationmanager.ContractAllocationManager
	ethClient          eth.HttpBackend
}

var errLegacyAVSsNotSupported = errors.New("method not supported for legacy AVSs")

func NewChainReader(
	delegationManager *delegationmanager.ContractDelegationManager,
	strategyManager *strategymanager.ContractStrategyManager,
	avsDirectory *avsdirectory.ContractAVSDirectory,
	rewardsCoordinator *rewardscoordinator.ContractIRewardsCoordinator,
	allocationManager *allocationmanager.ContractAllocationManager,
	logger logging.Logger,
	ethClient eth.HttpBackend,
) *ChainReader {
	logger = logger.With(logging.ComponentKey, "elcontracts/reader")

	return &ChainReader{
		delegationManager:  delegationManager,
		strategyManager:    strategyManager,
		avsDirectory:       avsDirectory,
		rewardsCoordinator: rewardsCoordinator,
		allocationManager:  allocationManager,
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
		elContractBindings.DelegationManager,
		elContractBindings.StrategyManager,
		elContractBindings.AvsDirectory,
		elContractBindings.RewardsCoordinator,
		elContractBindings.AllocationManager,
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
		elContractBindings.DelegationManager,
		elContractBindings.StrategyManager,
		elContractBindings.AvsDirectory,
		elContractBindings.RewardsCoordinator,
		elContractBindings.AllocationManager,
		logger,
		ethClient,
	), nil
}

// GetStakerShares returns the amount of shares that a staker has in all of the strategies in which they have nonzero
// shares
func (r *ChainReader) GetStakerShares(
	opts *bind.CallOpts,
	stakerAddress gethcommon.Address,
	blockNumber *big.Int,
) ([]gethcommon.Address, []*big.Int, error) {
	return r.delegationManager.GetDepositedShares(opts, stakerAddress)
}

// GetDelegatedOperator returns the operator that a staker has delegated to
func (r *ChainReader) GetDelegatedOperator(
	opts *bind.CallOpts,
	stakerAddress gethcommon.Address,
	blockNumber *big.Int,
) (gethcommon.Address, error) {
	return r.delegationManager.DelegatedTo(opts, stakerAddress)
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

	allocationDelayDetails, err := r.allocationManager.GetAllocationDelay(
		opts,
		gethcommon.HexToAddress(operator.Address),
	)
	if err == nil {
		return types.Operator{}, err
	}

	var allocationDelay uint32
	if allocationDelayDetails.IsSet {
		allocationDelay = allocationDelayDetails.Delay
	} else {
		allocationDelay = 0
	}

	return types.Operator{
		Address:                   operator.Address,
		StakerOptOutWindowBlocks:  operatorDetails.DeprecatedStakerOptOutWindowBlocks,
		DelegationApproverAddress: operatorDetails.DelegationApprover.Hex(),
		AllocationDelay:           allocationDelay,
	}, nil
}

// GetStrategyAndUnderlyingToken returns the strategy contract and the underlying token address
func (r *ChainReader) GetStrategyAndUnderlyingToken(
	opts *bind.CallOpts, strategyAddr gethcommon.Address,
) (*strategy.ContractIStrategy, gethcommon.Address, error) {
	contractStrategy, err := strategy.NewContractIStrategy(strategyAddr, r.ethClient)
	if err != nil {
		return nil, gethcommon.Address{}, utils.WrapError("Failed to fetch strategy contract", err)
	}
	underlyingTokenAddr, err := contractStrategy.UnderlyingToken(opts)
	if err != nil {
		return nil, gethcommon.Address{}, utils.WrapError("Failed to fetch token contract", err)
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
		return nil, nil, gethcommon.Address{}, utils.WrapError("Failed to fetch strategy contract", err)
	}
	underlyingTokenAddr, err := contractStrategy.UnderlyingToken(opts)
	if err != nil {
		return nil, nil, gethcommon.Address{}, utils.WrapError("Failed to fetch token contract", err)
	}
	contractUnderlyingToken, err := erc20.NewContractIERC20(underlyingTokenAddr, r.ethClient)
	if err != nil {
		return nil, nil, gethcommon.Address{}, utils.WrapError("Failed to fetch token contract", err)
	}
	return contractStrategy, contractUnderlyingToken, underlyingTokenAddr, nil
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
) (rewardscoordinator.IRewardsCoordinatorTypesDistributionRoot, error) {
	if r.rewardsCoordinator == nil {
		return rewardscoordinator.IRewardsCoordinatorTypesDistributionRoot{}, errors.New(
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

func (r *ChainReader) GetCumulativeClaimed(
	opts *bind.CallOpts,
	earner gethcommon.Address,
	token gethcommon.Address,
) (*big.Int, error) {
	if r.rewardsCoordinator == nil {
		return nil, errors.New("RewardsCoordinator contract not provided")
	}

	return r.rewardsCoordinator.CumulativeClaimed(opts, earner, token)
}

func (r *ChainReader) CheckClaim(
	opts *bind.CallOpts,
	claim rewardscoordinator.IRewardsCoordinatorTypesRewardsMerkleClaim,
) (bool, error) {
	if r.rewardsCoordinator == nil {
		return false, errors.New("RewardsCoordinator contract not provided")
	}

	return r.rewardsCoordinator.CheckClaim(opts, claim)
}

func (r *ChainReader) GetAllocatableMagnitude(
	opts *bind.CallOpts,
	operatorAddress gethcommon.Address,
	strategyAddress gethcommon.Address,
) (uint64, error) {
	if r.allocationManager == nil {
		return 0, errors.New("AllocationManager contract not provided")
	}

	return r.allocationManager.GetAllocatableMagnitude(opts, operatorAddress, strategyAddress)
}

func (r *ChainReader) GetMaxMagnitudes(
	opts *bind.CallOpts,
	operatorAddress gethcommon.Address,
	strategyAddresses []gethcommon.Address,
) ([]uint64, error) {
	if r.allocationManager == nil {
		return []uint64{}, errors.New("AllocationManager contract not provided")
	}

	return r.allocationManager.GetMaxMagnitudes(opts, operatorAddress, strategyAddresses)
}

func (r *ChainReader) GetAllocationInfo(
	opts *bind.CallOpts,
	operatorAddress gethcommon.Address,
	strategyAddress gethcommon.Address,
) ([]AllocationInfo, error) {
	if r.allocationManager == nil {
		return nil, errors.New("AllocationManager contract not provided")
	}

	opSets, allocationInfo, err := r.allocationManager.GetAllocationInfo1(opts, operatorAddress, strategyAddress)
	if err != nil {
		return nil, err
	}

	allocationsInfo := make([]AllocationInfo, len(opSets))
	for i, opSet := range opSets {
		allocationsInfo[i] = AllocationInfo{
			OperatorSetId:        opSet.OperatorSetId,
			AvsAddress:           opSet.Avs,
			CurrentMagnitude:     big.NewInt(int64(allocationInfo[i].CurrentMagnitude)),
			PendingDiff:          allocationInfo[i].PendingDiff,
			CompletableTimestamp: allocationInfo[i].EffectTimestamp,
		}
	}

	return allocationsInfo, nil
}

func (r *ChainReader) GetOperatorShares(
	opts *bind.CallOpts,
	operatorAddress gethcommon.Address,
	strategyAddresses []gethcommon.Address,
) ([]*big.Int, error) {
	if r.delegationManager == nil {
		return nil, errors.New("DelegationManager contract not provided")
	}

	return r.delegationManager.GetOperatorShares(opts, operatorAddress, strategyAddresses)
}

func (r *ChainReader) GetOperatorsShares(
	opts *bind.CallOpts,
	operatorAddress []gethcommon.Address,
	strategyAddresses []gethcommon.Address,
) ([][]*big.Int, error) {
	if r.delegationManager == nil {
		return nil, errors.New("DelegationManager contract not provided")
	}

	return r.delegationManager.GetOperatorsShares(opts, operatorAddress, strategyAddresses)
}

// GetNumberOfOperatorSetsForOperator returns the number of operator sets that an operator is part of
// Doesn't include M2 AVSs
func (r *ChainReader) GetNumOperatorSetsForOperator(
	opts *bind.CallOpts,
	operatorAddress gethcommon.Address,
) (*big.Int, error) {
	return r.avsDirectory.GetNumOperatorSetsOfOperator(opts, operatorAddress)
}

// GetOperatorSetsForOperator returns the list of operator sets that an operator is part of
// Doesn't include M2 AVSs
func (r *ChainReader) GetOperatorSetsForOperator(
	opts *bind.CallOpts,
	operatorAddress gethcommon.Address,
) ([]avsdirectory.OperatorSet, error) {
	// TODO: we're fetching max int64 operatorSets here. What's the practical limit for timeout by RPC? do we need to
	// paginate?
	return r.avsDirectory.GetOperatorSetsOfOperator(opts, operatorAddress, gethcommon.Big0, big.NewInt(math.MaxInt64))
}

// IsOperatorRegisteredWithOperatorSet returns if an operator is registered with a specific operator set
func (r *ChainReader) IsOperatorRegisteredWithOperatorSet(
	opts *bind.CallOpts,
	operatorAddress gethcommon.Address,
	operatorSet avsdirectory.OperatorSet,
) (bool, error) {
	if operatorSet.OperatorSetId == 0 {
		// this is an M2 AVS
		status, err := r.avsDirectory.AvsOperatorStatus(opts, operatorSet.Avs, operatorAddress)
		if err != nil {
			return false, err
		}

		return status == 1, nil
	} else {
		registered, err := r.avsDirectory.IsMember(opts, operatorAddress, operatorSet)
		if err != nil {
			return false, err
		}

		return registered, nil
	}
}

// GetOperatorsForOperatorSet returns the list of operators in a specific operator set
// Not supported for M2 AVSs
func (r *ChainReader) GetOperatorsForOperatorSet(
	opts *bind.CallOpts,
	operatorSet avsdirectory.OperatorSet,
) ([]gethcommon.Address, error) {
	if operatorSet.OperatorSetId == 0 {
		return nil, errLegacyAVSsNotSupported
	} else {
		return r.avsDirectory.GetOperatorsInOperatorSet(opts, operatorSet, gethcommon.Big0, big.NewInt(math.MaxInt64))
	}
}

// GetNumOperatorsForOperatorSet returns the number of operators in a specific operator set
func (r *ChainReader) GetNumOperatorsForOperatorSet(
	opts *bind.CallOpts,
	operatorSet avsdirectory.OperatorSet,
) (*big.Int, error) {
	if operatorSet.OperatorSetId == 0 {
		return nil, errLegacyAVSsNotSupported
	} else {
		return r.avsDirectory.GetNumOperatorsInOperatorSet(opts, operatorSet)
	}
}

// GetStrategiesForOperatorSet returns the list of strategies that an operator set takes into account
// Not supported for M2 AVSs
func (r *ChainReader) GetStrategiesForOperatorSet(
	opts *bind.CallOpts,
	operatorSet avsdirectory.OperatorSet,
) ([]gethcommon.Address, error) {
	if operatorSet.OperatorSetId == 0 {
		return nil, errLegacyAVSsNotSupported
	} else {
		return r.avsDirectory.GetStrategiesInOperatorSet(opts, operatorSet)
	}
}

type OperatorSetStakes struct {
	OperatorSet     avsdirectory.OperatorSet
	Strategies      []gethcommon.Address
	Operators       []gethcommon.Address
	DelegatedStakes [][]*big.Int
	SlashableStakes [][]*big.Int
}

// GetDelegatedAndSlashableSharesForOperatorSets returns the strategies the operatorSets take into account, their
// operators,
// and the minimum amount of shares that multiple operators delegated to them and slashable by the operatorSets.
// Not supported for M2 AVSs
func (r *ChainReader) GetDelegatedAndSlashableSharesForOperatorSets(
	opts *bind.CallOpts,
	operatorSets []avsdirectory.OperatorSet,
) ([]OperatorSetStakes, error) {
	operatorSetStakes := make([]OperatorSetStakes, len(operatorSets))
	for i, operatorSet := range operatorSets {
		operators, err := r.GetOperatorsForOperatorSet(opts, operatorSet)
		if err != nil {
			return nil, err
		}

		strategies, err := r.GetStrategiesForOperatorSet(opts, operatorSet)
		if err != nil {
			return nil, err
		}

		delegatedShares, slashableShares, err := r.allocationManager.GetCurrentDelegatedAndSlashableOperatorShares(
			opts,
			allocationmanager.OperatorSet{
				OperatorSetId: operatorSet.OperatorSetId,
				Avs:           operatorSet.Avs,
			},
			operators,
			strategies,
		)
		if err != nil {
			return nil, err
		}

		operatorSetStakes[i] = OperatorSetStakes{
			OperatorSet:     operatorSet,
			Strategies:      strategies,
			Operators:       operators,
			DelegatedStakes: delegatedShares,
			SlashableStakes: slashableShares,
		}
	}

	return operatorSetStakes, nil
}

// GetDelegatedAndSlashableSharesForOperatorSetsBefore returns the strategies the operatorSets take into account, their
// operators, and the minimum amount of shares that multiple operators delegated to them and slashable by the
// operatorSets before a given timestamp.
// Timestamp must be in the future. Used to underestimate future slashable stake.
// Not supported for M2 AVSs
func (r *ChainReader) GetDelegatedAndSlashableSharesForOperatorSetsBefore(
	opts *bind.CallOpts,
	operatorSets []avsdirectory.OperatorSet,
	beforeTimestamp uint32,
) ([]OperatorSetStakes, error) {
	operatorSetStakes := make([]OperatorSetStakes, len(operatorSets))
	for i, operatorSet := range operatorSets {
		operators, err := r.GetOperatorsForOperatorSet(opts, operatorSet)
		if err != nil {
			return nil, err
		}

		strategies, err := r.GetStrategiesForOperatorSet(opts, operatorSet)
		if err != nil {
			return nil, err
		}

		delegatedShares, slashableShares, err := r.allocationManager.GetMinDelegatedAndSlashableOperatorSharesBefore(
			opts,
			allocationmanager.OperatorSet{
				OperatorSetId: operatorSet.OperatorSetId,
				Avs:           operatorSet.Avs,
			},
			operators,
			strategies,
			beforeTimestamp,
		)
		if err != nil {
			return nil, err
		}

		operatorSetStakes[i] = OperatorSetStakes{
			OperatorSet:     operatorSet,
			Strategies:      strategies,
			Operators:       operators,
			DelegatedStakes: delegatedShares,
			SlashableStakes: slashableShares,
		}
	}

	return operatorSetStakes, nil
}
