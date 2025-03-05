package elcontracts

import (
	"context"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethcommon "github.com/ethereum/go-ethereum/common"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	avsdirectory "github.com/Layr-Labs/eigensdk-go/contracts/bindings/AVSDirectory"
	allocationmanager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/AllocationManager"
	delegationmanager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/DelegationManager"
	erc20 "github.com/Layr-Labs/eigensdk-go/contracts/bindings/IERC20"
	strategy "github.com/Layr-Labs/eigensdk-go/contracts/bindings/IStrategy"
	permissioncontroller "github.com/Layr-Labs/eigensdk-go/contracts/bindings/PermissionController"
	rewardscoordinator "github.com/Layr-Labs/eigensdk-go/contracts/bindings/RewardsCoordinator"
	strategymanager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/StrategyManager"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/Layr-Labs/eigensdk-go/utils"
)

type Config struct {
	DelegationManagerAddress    gethcommon.Address
	AvsDirectoryAddress         gethcommon.Address
	RewardsCoordinatorAddress   gethcommon.Address
	PermissionControllerAddress gethcommon.Address

	/// Setting this to true will disable the fetching of the AllocationManager address.
	/// This is useful for older deployments, which don't have the contract deployed.
	// TODO: remove this once mainnet is updated with the new contracts
	DontUseAllocationManager bool
}

type ChainReader struct {
	logger               logging.Logger
	delegationManager    *delegationmanager.ContractDelegationManager
	strategyManager      *strategymanager.ContractStrategyManager
	avsDirectory         *avsdirectory.ContractAVSDirectory
	rewardsCoordinator   *rewardscoordinator.ContractRewardsCoordinator
	allocationManager    *allocationmanager.ContractAllocationManager
	permissionController *permissioncontroller.ContractPermissionController
	ethClient            eth.HttpBackend
}

var errLegacyAVSsNotSupported = errors.New("method not supported for legacy AVSs")

// Returns a new instance of ChainReader from a given set of bindings.
func NewChainReader(
	delegationManager *delegationmanager.ContractDelegationManager,
	strategyManager *strategymanager.ContractStrategyManager,
	avsDirectory *avsdirectory.ContractAVSDirectory,
	rewardsCoordinator *rewardscoordinator.ContractRewardsCoordinator,
	allocationManager *allocationmanager.ContractAllocationManager,
	permissionController *permissioncontroller.ContractPermissionController,
	logger logging.Logger,
	ethClient eth.HttpBackend,
) *ChainReader {
	logger = logger.With(logging.ComponentKey, "elcontracts/reader")

	return &ChainReader{
		delegationManager:    delegationManager,
		strategyManager:      strategyManager,
		avsDirectory:         avsDirectory,
		rewardsCoordinator:   rewardsCoordinator,
		allocationManager:    allocationManager,
		permissionController: permissionController,
		logger:               logger,
		ethClient:            ethClient,
	}
}

// Returns a new instance of ChainReader from a given config.
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
		elContractBindings.PermissionController,
		logger,
		ethClient,
	), nil
}

// Returns `true` if the operator is registered to the EigenLayer protocol, `false` otherwise.
// Can return an error if the `DelegationManager` contract address was not provided, or due to
// errors in the underlying contract call.
func (r *ChainReader) IsOperatorRegistered(
	ctx context.Context,
	operator types.Operator,
) (bool, error) {
	if r.delegationManager == nil {
		return false, errors.New("DelegationManager contract not provided")
	}

	return r.delegationManager.IsOperator(&bind.CallOpts{Context: ctx}, gethcommon.HexToAddress(operator.Address))
}

// Returns the amount of shares that a staker has in all of the strategies they have shares in.
// Can return an error if the `DelegationManager` contract address was not provided, or due to
// errors in the underlying contract call.
func (r *ChainReader) GetStakerShares(
	ctx context.Context,
	stakerAddress gethcommon.Address,
) ([]gethcommon.Address, []*big.Int, error) {
	if r.delegationManager == nil {
		return nil, nil, errors.New("DelegationManager contract not provided")
	}
	return r.delegationManager.GetDepositedShares(&bind.CallOpts{Context: ctx}, stakerAddress)
}

// Returns the AVSRegistrar of the avs received as parameter.
// Can return an error if the `DelegationManager` contract address was not provided, or due to
// errors in the underlying contract call.
func (r *ChainReader) GetAVSRegistrar(
	ctx context.Context,
	avs gethcommon.Address,
) (gethcommon.Address, error) {
	if r.allocationManager == nil {
		return gethcommon.Address{}, errors.New("AllocationManager contract not provided")
	}
	return r.allocationManager.GetAVSRegistrar(&bind.CallOpts{Context: ctx}, avs)
}

// Returns the operator that a staker has delegated to.
// Can return an error if the `DelegationManager` contract address was not provided, or due to
// errors in the underlying contract call.
func (r *ChainReader) GetDelegatedOperator(
	ctx context.Context,
	stakerAddress gethcommon.Address,
	blockNumber *big.Int,
) (gethcommon.Address, error) {
	if r.delegationManager == nil {
		return gethcommon.Address{}, errors.New("DelegationManager contract not provided")
	}
	return r.delegationManager.DelegatedTo(&bind.CallOpts{Context: ctx}, stakerAddress)
}

// Returns detailed information on an operator.
// Can return an error if the `DelegationManager` contract address was not provided, or due to
// errors in the underlying contract call.
func (r *ChainReader) GetOperatorDetails(
	ctx context.Context,
	operator types.Operator,
) (types.Operator, error) {
	if r.delegationManager == nil {
		return types.Operator{}, errors.New("DelegationManager contract not provided")
	}

	delegationManagerAddress, err := r.delegationManager.DelegationApprover(
		&bind.CallOpts{Context: ctx},
		gethcommon.HexToAddress(operator.Address),
	)
	// This call should not fail since it's a getter
	if err != nil {
		return types.Operator{}, err
	}

	isSet, delay, err := r.allocationManager.GetAllocationDelay(
		&bind.CallOpts{
			Context: ctx,
		},
		gethcommon.HexToAddress(operator.Address),
	)
	// This call should not fail
	if err != nil {
		return types.Operator{}, err
	}

	var allocationDelay uint32
	if isSet {
		allocationDelay = delay
	} else {
		allocationDelay = 0
	}

	return types.Operator{
		Address:                   operator.Address,
		DelegationApproverAddress: delegationManagerAddress.Hex(),
		AllocationDelay:           allocationDelay,
	}, nil
}

// Returns the bindings of a given strategy and the address of its underlying token.
// Can return an error due to errors in the underlying contract call.
func (r *ChainReader) GetStrategyAndUnderlyingToken(
	ctx context.Context,
	strategyAddr gethcommon.Address,
) (*strategy.ContractIStrategy, gethcommon.Address, error) {
	contractStrategy, err := strategy.NewContractIStrategy(strategyAddr, r.ethClient)
	// This call should not fail since it's an init
	if err != nil {
		return nil, gethcommon.Address{}, utils.WrapError("Failed to fetch strategy contract", err)
	}
	underlyingTokenAddr, err := contractStrategy.UnderlyingToken(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, gethcommon.Address{}, utils.WrapError("Failed to fetch token contract", err)
	}
	return contractStrategy, underlyingTokenAddr, nil
}

// Returns the bindings of a given strategy and the bindings and address of its underlying token.
// Can return an error due to errors in the underlying contract call.
func (r *ChainReader) GetStrategyAndUnderlyingERC20Token(
	ctx context.Context,
	strategyAddr gethcommon.Address,
) (*strategy.ContractIStrategy, erc20.ContractIERC20Methods, gethcommon.Address, error) {
	contractStrategy, err := strategy.NewContractIStrategy(strategyAddr, r.ethClient)
	// This call should not fail since it's an init
	if err != nil {
		return nil, nil, gethcommon.Address{}, utils.WrapError("Failed to fetch strategy contract", err)
	}
	underlyingTokenAddr, err := contractStrategy.UnderlyingToken(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, nil, gethcommon.Address{}, utils.WrapError("Failed to fetch token contract", err)
	}
	contractUnderlyingToken, err := erc20.NewContractIERC20(underlyingTokenAddr, r.ethClient)
	// This call should not fail, if the strategy does not have an underlying token then it would enter the if above
	if err != nil {
		return nil, nil, gethcommon.Address{}, utils.WrapError("Failed to fetch token contract", err)
	}
	return contractStrategy, contractUnderlyingToken, underlyingTokenAddr, nil
}

// Returns the shares an operator has in a given strategy.
// Can return an error if the `DelegationManager` contract address was not provided, or due to
// errors in the underlying contract call.
func (r *ChainReader) GetOperatorSharesInStrategy(
	ctx context.Context,
	operatorAddr gethcommon.Address,
	strategyAddr gethcommon.Address,
) (*big.Int, error) {
	if r.delegationManager == nil {
		return &big.Int{}, errors.New("DelegationManager contract not provided")
	}

	return r.delegationManager.OperatorShares(
		&bind.CallOpts{Context: ctx},
		operatorAddr,
		strategyAddr,
	)
}

// Returns the digest hash to be signed by the operator's delegation approver to be used
// when delegating to an operator.
// Can return an error if the `DelegationManager` contract address was not provided, or due to
// errors in the underlying contract call.
func (r *ChainReader) CalculateDelegationApprovalDigestHash(
	ctx context.Context,
	staker gethcommon.Address,
	operator gethcommon.Address,
	delegationApprover gethcommon.Address,
	approverSalt [32]byte,
	expiry *big.Int,
) ([32]byte, error) {
	if r.delegationManager == nil {
		return [32]byte{}, errors.New("DelegationManager contract not provided")
	}

	return r.delegationManager.CalculateDelegationApprovalDigestHash(
		&bind.CallOpts{Context: ctx},
		staker,
		operator,
		delegationApprover,
		approverSalt,
		expiry,
	)
}

// Returns the digest hash to be signed by an operator to register with an AVS.
// Can return an error if the `AVSDirectory` contract address was not provided, or due to
// errors in the underlying contract call.
func (r *ChainReader) CalculateOperatorAVSRegistrationDigestHash(
	ctx context.Context,
	operator gethcommon.Address,
	avs gethcommon.Address,
	salt [32]byte,
	expiry *big.Int,
) ([32]byte, error) {
	if r.avsDirectory == nil {
		return [32]byte{}, errors.New("AVSDirectory contract not provided")
	}

	return r.avsDirectory.CalculateOperatorAVSRegistrationDigestHash(
		&bind.CallOpts{Context: ctx},
		operator,
		avs,
		salt,
		expiry,
	)
}

// Returns the number of distribution roots published.
// Can return an error if the `RewardsCoordinator` contract address was not provided, or due to
// errors in the underlying contract call.
func (r *ChainReader) GetDistributionRootsLength(ctx context.Context) (*big.Int, error) {
	if r.rewardsCoordinator == nil {
		return nil, errors.New("RewardsCoordinator contract not provided")
	}

	return r.rewardsCoordinator.GetDistributionRootsLength(&bind.CallOpts{Context: ctx})
}

// Returns the timestamp until which rewards submissions have been calculated.
// Can return an error if the `RewardsCoordinator` contract address was not provided, or due to
// errors in the underlying contract call.
func (r *ChainReader) CurrRewardsCalculationEndTimestamp(ctx context.Context) (uint32, error) {
	if r.rewardsCoordinator == nil {
		return 0, errors.New("RewardsCoordinator contract not provided")
	}

	return r.rewardsCoordinator.CurrRewardsCalculationEndTimestamp(&bind.CallOpts{Context: ctx})
}

// Returns the latest root that can be claimed against.
// Can return an error if the `RewardsCoordinator` contract address was not provided, or due to
// errors in the underlying contract call.
func (r *ChainReader) GetCurrentClaimableDistributionRoot(
	ctx context.Context,
) (rewardscoordinator.IRewardsCoordinatorTypesDistributionRoot, error) {
	if r.rewardsCoordinator == nil {
		return rewardscoordinator.IRewardsCoordinatorTypesDistributionRoot{}, errors.New(
			"RewardsCoordinator contract not provided",
		)
	}

	return r.rewardsCoordinator.GetCurrentClaimableDistributionRoot(&bind.CallOpts{Context: ctx})
}

// Returns the index of the latest root that can be claimed against.
// Can return an error if the `RewardsCoordinator` contract address was not provided, or due to
// errors in the underlying contract call.
func (r *ChainReader) GetRootIndexFromHash(
	ctx context.Context,
	rootHash [32]byte,
) (uint32, error) {
	if r.rewardsCoordinator == nil {
		return 0, errors.New("RewardsCoordinator contract not provided")
	}

	return r.rewardsCoordinator.GetRootIndexFromHash(&bind.CallOpts{Context: ctx}, rootHash)
}

// Returns the number of `token` tokens the `earner` has claimed.
// Can return an error if the `RewardsCoordinator` contract address was not provided, or due to
// errors in the underlying contract call.
func (r *ChainReader) GetCumulativeClaimed(
	ctx context.Context,
	earner gethcommon.Address,
	token gethcommon.Address,
) (*big.Int, error) {
	if r.rewardsCoordinator == nil {
		return nil, errors.New("RewardsCoordinator contract not provided")
	}

	return r.rewardsCoordinator.CumulativeClaimed(&bind.CallOpts{Context: ctx}, earner, token)
}

// Returns `true` if the claim would currently pass the check in
// `ChainWriter.ProcessClaims` or return an error if invalid.
// Can return an error if the `RewardsCoordinator` contract address was not provided, or due to
// errors in the underlying contract call.
func (r *ChainReader) CheckClaim(
	ctx context.Context,
	claim rewardscoordinator.IRewardsCoordinatorTypesRewardsMerkleClaim,
) (bool, error) {
	if r.rewardsCoordinator == nil {
		return false, errors.New("RewardsCoordinator contract not provided")
	}

	// TODO: this returns an error if the claim is invalid. We map this error to false instead
	return r.rewardsCoordinator.CheckClaim(&bind.CallOpts{Context: ctx}, claim)
}

// Returns the split configured by the `operator` for the `avs`.
// Can return an error if the `RewardsCoordinator` contract address was not provided, or due to
// errors in the underlying contract call.
func (r *ChainReader) GetOperatorAVSSplit(
	ctx context.Context,
	operator gethcommon.Address,
	avs gethcommon.Address,
) (uint16, error) {
	if r.rewardsCoordinator == nil {
		return 0, errors.New("RewardsCoordinator contract not provided")
	}

	return r.rewardsCoordinator.GetOperatorAVSSplit(&bind.CallOpts{Context: ctx}, operator, avs)
}

// Returns the split configured by the `operator` for Programmatic Incentives.
// Can return an error if the `RewardsCoordinator` contract address was not provided, or due to
// errors in the underlying contract call.
func (r *ChainReader) GetOperatorPISplit(
	ctx context.Context,
	operator gethcommon.Address,
) (uint16, error) {
	if r.rewardsCoordinator == nil {
		return 0, errors.New("RewardsCoordinator contract not provided")
	}

	return r.rewardsCoordinator.GetOperatorPISplit(&bind.CallOpts{Context: ctx}, operator)
}

// Returns the split for an operator in an operator set.
// Can return an error if the `RewardsCoordinator` contract address was not provided, or due to
// errors in the underlying contract call.
func (r *ChainReader) GetOperatorSetSplit(
	ctx context.Context,
	operator gethcommon.Address,
	operatorSet rewardscoordinator.OperatorSet,
) (uint16, error) {
	if r.rewardsCoordinator == nil {
		return 0, errors.New("RewardsCoordinator contract not provided")
	}

	return r.rewardsCoordinator.GetOperatorSetSplit(&bind.CallOpts{Context: ctx}, operator, operatorSet)
}

// Gets the interval in seconds at which the calculation for rewards distribution is done.
func (r *ChainReader) GetCalculationIntervalSeconds(
	ctx context.Context,
) (uint32, error) {
	if r.rewardsCoordinator == nil {
		return 0, errors.New("RewardsCoordinator contract not provided")
	}
	return r.rewardsCoordinator.CALCULATIONINTERVALSECONDS(&bind.CallOpts{Context: ctx})
}

// Gets the maximum amount of time (seconds) that a rewards submission can span over
func (r *ChainReader) GetMaxRewardsDuration(
	ctx context.Context,
) (uint32, error) {
	if r.rewardsCoordinator == nil {
		return 0, errors.New("RewardsCoordinator contract not provided")
	}
	return r.rewardsCoordinator.MAXREWARDSDURATION(&bind.CallOpts{Context: ctx})
}

// Get the max amount of time (seconds) that a rewards submission can start in the past
func (r *ChainReader) GetMaxRetroactiveLength(
	ctx context.Context,
) (uint32, error) {
	if r.rewardsCoordinator == nil {
		return 0, errors.New("RewardsCoordinator contract not provided")
	}
	return r.rewardsCoordinator.MAXRETROACTIVELENGTH(&bind.CallOpts{Context: ctx})
}

// Get the max amount of time (seconds) that a rewards submission can start in the future
func (r *ChainReader) GetMaxFutureLength(
	ctx context.Context,
) (uint32, error) {
	if r.rewardsCoordinator == nil {
		return 0, errors.New("RewardsCoordinator contract not provided")
	}
	return r.rewardsCoordinator.MAXFUTURELENGTH(&bind.CallOpts{Context: ctx})
}

// Get absolute min timestamp (seconds) that a rewards submission can start at
func (r *ChainReader) GetGenesisRewardsTimestamp(
	ctx context.Context,
) (uint32, error) {
	if r.rewardsCoordinator == nil {
		return 0, errors.New("RewardsCoordinator contract not provided")
	}
	return r.rewardsCoordinator.GENESISREWARDSTIMESTAMP(&bind.CallOpts{Context: ctx})
}

// Get the address of the entity that can update the contract with new merkle roots
func (r *ChainReader) GetRewardsUpdater(
	ctx context.Context,
) (gethcommon.Address, error) {
	if r.rewardsCoordinator == nil {
		return gethcommon.Address{}, errors.New("RewardsCoordinator contract not provided")
	}
	return r.rewardsCoordinator.RewardsUpdater(&bind.CallOpts{Context: ctx})
}

// Get delay in timestamp (seconds) before a posted root can be claimed against
func (r *ChainReader) GetActivationDelay(
	ctx context.Context,
) (uint32, error) {
	if r.rewardsCoordinator == nil {
		return 0, errors.New("RewardsCoordinator contract not provided")
	}
	return r.rewardsCoordinator.ActivationDelay(&bind.CallOpts{Context: ctx})
}

// Get timestamp for last submitted DistributionRoot
func (r *ChainReader) GetCurrRewardsCalculationEndTimestamp(
	ctx context.Context,
) (uint32, error) {
	if r.rewardsCoordinator == nil {
		return 0, errors.New("RewardsCoordinator contract not provided")
	}
	return r.rewardsCoordinator.CurrRewardsCalculationEndTimestamp(&bind.CallOpts{Context: ctx})
}

// Get the default split for all operators across all avss in bips.
func (r *ChainReader) GetDefaultOperatorSplitBips(
	ctx context.Context,
) (uint16, error) {
	if r.rewardsCoordinator == nil {
		return 0, errors.New("RewardsCoordinator contract not provided")
	}
	return r.rewardsCoordinator.DefaultOperatorSplitBips(&bind.CallOpts{Context: ctx})
}

func (r *ChainReader) GetClaimerFor(
	ctx context.Context,
	earner gethcommon.Address,
) (gethcommon.Address, error) {
	if r.rewardsCoordinator == nil {
		return gethcommon.Address{}, errors.New("RewardsCoordinator contract not provided")
	}
	return r.rewardsCoordinator.ClaimerFor(&bind.CallOpts{Context: ctx}, earner)
}

// Returns the submission nonce for an avs
func (r *ChainReader) GetSubmissionNonce(
	ctx context.Context,
	avs gethcommon.Address,
) (*big.Int, error) {
	if r.rewardsCoordinator == nil {
		return nil, errors.New("RewardsCoordinator contract not provided")
	}
	return r.rewardsCoordinator.SubmissionNonce(&bind.CallOpts{Context: ctx}, avs)
}

// Returns whether a hash is a valid rewards submission hash for a given avs
func (r *ChainReader) GetIsAVSRewardsSubmissionHash(
	ctx context.Context,
	avs gethcommon.Address,
	hash [32]byte,
) (bool, error) {
	if r.rewardsCoordinator == nil {
		return false, errors.New("RewardsCoordinator contract not provided")
	}
	return r.rewardsCoordinator.IsAVSRewardsSubmissionHash(&bind.CallOpts{Context: ctx}, avs, hash)
}

// Returns whether a hash is a valid rewards submission for all hash for a given avs
func (r *ChainReader) GetIsRewardsSubmissionForAllHash(
	ctx context.Context,
	avs gethcommon.Address,
	hash [32]byte,
) (bool, error) {
	if r.rewardsCoordinator == nil {
		return false, errors.New("RewardsCoordinator contract not provided")
	}
	return r.rewardsCoordinator.IsRewardsSubmissionForAllHash(&bind.CallOpts{Context: ctx}, avs, hash)
}

// Returns whether a submitter is a valid rewards for all submitter
func (r *ChainReader) GetIsRewardsForAllSubmitter(
	ctx context.Context,
	submitter gethcommon.Address,
) (bool, error) {
	if r.rewardsCoordinator == nil {
		return false, errors.New("RewardsCoordinator contract not provided")
	}
	return r.rewardsCoordinator.IsRewardsForAllSubmitter(&bind.CallOpts{Context: ctx}, submitter)
}

// Returns whether a hash is a valid rewards submission for all earners hash for a given avs
func (r *ChainReader) GetIsRewardsSubmissionForAllEarnersHash(
	ctx context.Context,
	avs gethcommon.Address,
	hash [32]byte,
) (bool, error) {
	if r.rewardsCoordinator == nil {
		return false, errors.New("RewardsCoordinator contract not provided")
	}
	return r.rewardsCoordinator.IsRewardsSubmissionForAllEarnersHash(&bind.CallOpts{Context: ctx}, avs, hash)
}

// Returns whether a hash is a valid operator set performance rewards submission hash for a given avs
func (r *ChainReader) GetIsOperatorDirectedAVSRewardsSubmissionHash(
	ctx context.Context,
	avs gethcommon.Address,
	hash [32]byte,
) (bool, error) {
	if r.rewardsCoordinator == nil {
		return false, errors.New("RewardsCoordinator contract not provided")
	}
	return r.rewardsCoordinator.IsOperatorDirectedAVSRewardsSubmissionHash(&bind.CallOpts{Context: ctx}, avs, hash)
}

// Returns whether a hash is a valid operator set performance rewards submission hash for a given avs
func (r *ChainReader) GetIsOperatorDirectedOperatorSetRewardsSubmissionHash(
	ctx context.Context,
	avs gethcommon.Address,
	hash [32]byte,
) (bool, error) {
	if r.rewardsCoordinator == nil {
		return false, errors.New("RewardsCoordinator contract not provided")
	}
	return r.rewardsCoordinator.IsOperatorDirectedOperatorSetRewardsSubmissionHash(
		&bind.CallOpts{Context: ctx},
		avs,
		hash,
	)
}

// Returns the amount of magnitude on a strategy not currently allocated to any operator set,
// by an operator.
// Can return an error if the `AllocationManager` contract address was not provided, or due to
// errors in the underlying contract call.
func (r *ChainReader) GetAllocatableMagnitude(
	ctx context.Context,
	operatorAddress gethcommon.Address,
	strategyAddress gethcommon.Address,
) (uint64, error) {
	if r.allocationManager == nil {
		return 0, errors.New("AllocationManager contract not provided")
	}

	return r.allocationManager.GetAllocatableMagnitude(&bind.CallOpts{Context: ctx}, operatorAddress, strategyAddress)
}

// Returns the amount of magnitude an operator has allocated to operator sets for a given strategy
func (r *ChainReader) GetEncumberedMagnitude(
	ctx context.Context,
	operatorAddress gethcommon.Address,
	strategyAddress gethcommon.Address,
) (uint64, error) {
	if r.allocationManager == nil {
		return 0, errors.New("AllocationManager contract not provided")
	}

	return r.allocationManager.EncumberedMagnitude(&bind.CallOpts{Context: ctx}, operatorAddress, strategyAddress)
}

// Returns the delay within which deallocations are slashable.
func (r *ChainReader) GetDeallocationDelay(
	ctx context.Context,
) (uint32, error) {
	if r.allocationManager == nil {
		return 0, errors.New("AllocationManager contract not provided")
	}
	return r.allocationManager.DEALLOCATIONDELAY(&bind.CallOpts{Context: ctx})
}

// Returns the delay before allocation delay modifications take effect.
func (r *ChainReader) GetAllocationConfigurationDelay(
	ctx context.Context,
) (uint32, error) {
	if r.allocationManager == nil {
		return 0, errors.New("AllocationManager contract not provided")
	}
	return r.allocationManager.ALLOCATIONCONFIGURATIONDELAY(&bind.CallOpts{Context: ctx})
}

// Returns the maximum magnitude an operator can allocate for the given strategies.
// Can return an error if the `AllocationManager` contract address was not provided, or due to
// errors in the underlying contract call.
func (r *ChainReader) GetMaxMagnitudes(
	ctx context.Context,
	operatorAddress gethcommon.Address,
	strategyAddresses []gethcommon.Address,
) ([]uint64, error) {
	if r.allocationManager == nil {
		return []uint64{}, errors.New("AllocationManager contract not provided")
	}

	return r.allocationManager.GetMaxMagnitudes0(&bind.CallOpts{Context: ctx}, operatorAddress, strategyAddresses)
}

// Returns the allocation info of a given operator and strategy.
// Can return an error if the `AllocationManager` contract address was not provided, or due to
// errors in the underlying contract call.
func (r *ChainReader) GetAllocationInfo(
	ctx context.Context,
	operatorAddress gethcommon.Address,
	strategyAddress gethcommon.Address,
) ([]AllocationInfo, error) {
	if r.allocationManager == nil {
		return nil, errors.New("AllocationManager contract not provided")
	}

	opSets, allocationInfo, err := r.allocationManager.GetStrategyAllocations(
		&bind.CallOpts{Context: ctx},
		operatorAddress,
		strategyAddress,
	)
	// This call should not fail since it's a getter
	if err != nil {
		return nil, err
	}

	allocationsInfo := make([]AllocationInfo, len(opSets))
	for i, opSet := range opSets {
		allocationsInfo[i] = AllocationInfo{
			OperatorSetId:    opSet.Id,
			AvsAddress:       opSet.Avs,
			CurrentMagnitude: big.NewInt(int64(allocationInfo[i].CurrentMagnitude)),
			PendingDiff:      allocationInfo[i].PendingDiff,
			EffectBlock:      allocationInfo[i].EffectBlock,
		}
	}

	return allocationsInfo, nil
}

// Returns the shares an operator has delegated to them on a strategy.
// Can return an error if the `DelegationManager` contract address was not provided, or due to
// errors in the underlying contract call.
func (r *ChainReader) GetOperatorShares(
	ctx context.Context,
	operatorAddress gethcommon.Address,
	strategyAddresses []gethcommon.Address,
) ([]*big.Int, error) {
	if r.delegationManager == nil {
		return nil, errors.New("DelegationManager contract not provided")
	}

	return r.delegationManager.GetOperatorShares(&bind.CallOpts{
		Context: ctx,
	}, operatorAddress, strategyAddresses)
}

// Returns the shares that a set of operators have delegated to them in a set of strategies.
// Can return an error if the `DelegationManager` contract address was not provided, or due to
// errors in the underlying contract call.
func (r *ChainReader) GetOperatorsShares(
	ctx context.Context,
	operatorAddresses []gethcommon.Address,
	strategyAddresses []gethcommon.Address,
) ([][]*big.Int, error) {
	if r.delegationManager == nil {
		return nil, errors.New("DelegationManager contract not provided")
	}
	return r.delegationManager.GetOperatorsShares(&bind.CallOpts{Context: ctx}, operatorAddresses, strategyAddresses)
}

// Returns whether `delegationApprover` has already used the given `salt`.
func (r *ChainReader) GetDelegationApproverSaltIsSpent(
	ctx context.Context,
	delegationApprover gethcommon.Address,
	approverSalt [32]byte,
) (bool, error) {
	if r.delegationManager == nil {
		return false, errors.New("DelegationManager contract not provided")
	}

	return r.delegationManager.DelegationApproverSaltIsSpent(
		&bind.CallOpts{Context: ctx},
		delegationApprover,
		approverSalt,
	)
}

// Returns whether a withdrawal is pending for a given `withdrawalRoot`.
func (r *ChainReader) GetPendingWithdrawalStatus(
	ctx context.Context,
	withdrawalRoot [32]byte,
) (bool, error) {
	if r.delegationManager == nil {
		return false, errors.New("DelegationManager contract not provided")
	}

	return r.delegationManager.PendingWithdrawals(&bind.CallOpts{Context: ctx}, withdrawalRoot)
}

// Returns the total number of withdrawals that have been queued for a given `staker`
func (r *ChainReader) GetCumulativeWithdrawalsQueued(
	ctx context.Context,
	staker gethcommon.Address,
) (*big.Int, error) {
	if r.delegationManager == nil {
		return big.NewInt(0), errors.New("DelegationManager contract not provided")
	}
	return r.delegationManager.CumulativeWithdrawalsQueued(&bind.CallOpts{Context: ctx}, staker)
}

// Returns the number of operator sets that an operator is part of.
// This doesn't include M2 quorums.
// Can return an error if the `AllocationManager` contract address was not provided, or due to
// errors in the underlying contract call.
func (r *ChainReader) GetNumOperatorSetsForOperator(
	ctx context.Context,
	operatorAddress gethcommon.Address,
) (*big.Int, error) {
	if r.allocationManager == nil {
		return nil, errors.New("AllocationManager contract not provided")
	}
	opSets, err := r.allocationManager.GetAllocatedSets(&bind.CallOpts{Context: ctx}, operatorAddress)
	if err != nil {
		return nil, err
	}
	return big.NewInt(int64(len(opSets))), nil
}

// Returns the list of operator sets the operator has current or pending allocations/deallocations in.
// This doesn't include M2 quorums.
// Can return an error if the `AllocationManager` contract address was not provided, or due to
// errors in the underlying contract call.
func (r *ChainReader) GetOperatorSetsForOperator(
	ctx context.Context,
	operatorAddress gethcommon.Address,
) ([]allocationmanager.OperatorSet, error) {
	if r.allocationManager == nil {
		return nil, errors.New("AllocationManager contract not provided")
	}
	// TODO: we're fetching max int64 operatorSets here. What's the practical limit for timeout by RPC? do we need to
	// paginate?
	return r.allocationManager.GetAllocatedSets(&bind.CallOpts{Context: ctx}, operatorAddress)
}

// Returns `true` if an operator is registered with a specific operator set or M2 quorum.
// Can return an error if the `AVSDirectory` or `AllocationManager` contract addresses were
// not provided, or due to errors in the underlying contract call.
func (r *ChainReader) IsOperatorRegisteredWithOperatorSet(
	ctx context.Context,
	operatorAddress gethcommon.Address,
	operatorSet allocationmanager.OperatorSet,
) (bool, error) {
	if r.allocationManager == nil {
		return false, errors.New("AllocationManager contract not provided")
	}
	registeredOperatorSets, err := r.allocationManager.GetRegisteredSets(&bind.CallOpts{Context: ctx}, operatorAddress)
	// This call should not fail since it's a getter
	if err != nil {
		return false, err
	}
	for _, registeredOperatorSet := range registeredOperatorSets {
		if registeredOperatorSet.Id == operatorSet.Id && registeredOperatorSet.Avs == operatorSet.Avs {
			return true, nil
		}
	}

	return false, nil
}

// Returns the list of operators in a specific operator set.
// Not supported for M2 AVSs.
// Can return an error if the `AllocationManager` contract address was not provided, or due to
// errors in the underlying contract call.
func (r *ChainReader) GetOperatorsForOperatorSet(
	ctx context.Context,
	operatorSet allocationmanager.OperatorSet,
) ([]gethcommon.Address, error) {
	if operatorSet.Id == 0 {
		return nil, errLegacyAVSsNotSupported
	} else {
		if r.allocationManager == nil {
			return nil, errors.New("AllocationManager contract not provided")
		}

		return r.allocationManager.GetMembers(&bind.CallOpts{Context: ctx}, operatorSet)
	}
}

// Returns the number of operators in a specific operator set.
// Not supported for M2 AVSs.
// Can return an error if the `AllocationManager` contract address was not provided, or due to
// errors in the underlying contract call.
func (r *ChainReader) GetNumOperatorsForOperatorSet(
	ctx context.Context,
	operatorSet allocationmanager.OperatorSet,
) (*big.Int, error) {
	if operatorSet.Id == 0 {
		return nil, errLegacyAVSsNotSupported
	} else {
		if r.allocationManager == nil {
			return nil, errors.New("AllocationManager contract not provided")
		}

		return r.allocationManager.GetMemberCount(&bind.CallOpts{Context: ctx}, operatorSet)
	}
}

// Returns the list of strategies that an operator set takes into account.
// Not supported for M2 AVSs.
// Can return an error if the `AllocationManager` contract address was not provided, or due to
// errors in the underlying contract call.
func (r *ChainReader) GetStrategiesForOperatorSet(
	ctx context.Context,
	operatorSet allocationmanager.OperatorSet,
) ([]gethcommon.Address, error) {
	if operatorSet.Id == 0 {
		return nil, errLegacyAVSsNotSupported
	} else {
		if r.allocationManager == nil {
			return nil, errors.New("AllocationManager contract not provided")
		}

		return r.allocationManager.GetStrategiesInOperatorSet(&bind.CallOpts{Context: ctx}, operatorSet)
	}
}

// Returns a list of the number of shares slashable by the operator set
// for each of the given strategies.
// Can return an error if the `AllocationManager` contract address was not provided, or due to
// errors in the underlying contract call.
func (r *ChainReader) GetSlashableShares(
	ctx context.Context,
	operatorAddress gethcommon.Address,
	operatorSet allocationmanager.OperatorSet,
	strategies []gethcommon.Address,
) (map[gethcommon.Address]*big.Int, error) {
	if r.allocationManager == nil {
		return nil, errors.New("AllocationManager contract not provided")
	}

	currentBlock, err := r.ethClient.BlockNumber(ctx)
	// This call should not fail since it's a getter
	if err != nil {
		return nil, err
	}

	slashableShares, err := r.allocationManager.GetMinimumSlashableStake(
		&bind.CallOpts{Context: ctx},
		operatorSet,
		[]gethcommon.Address{operatorAddress},
		strategies,
		uint32(currentBlock),
	)
	// This call should not fail since it's a getter
	if err != nil {
		return nil, err
	}
	if len(slashableShares) == 0 {
		return nil, errors.New("no slashable shares found for operator")
	}

	slashableShareStrategyMap := make(map[gethcommon.Address]*big.Int)
	for i, strat := range strategies {
		// The reason we use 0 here is because we only have one operator in the list
		slashableShareStrategyMap[strat] = slashableShares[0][i]
	}

	return slashableShareStrategyMap, nil
}

// Returns the strategies the `operatorSets` take into account, their
// operators, and the minimum amount of shares that are slashable by the operatorSets.
// Not supported for M2 AVSs.
// Can return an error if the `AllocationManager` contract address was not provided, or due to
// errors in the underlying contract call.
func (r *ChainReader) GetSlashableSharesForOperatorSets(
	ctx context.Context,
	operatorSets []allocationmanager.OperatorSet,
) ([]OperatorSetStakes, error) {
	currentBlock, err := r.ethClient.BlockNumber(ctx)
	// This call should not fail since it's a getter
	if err != nil {
		return nil, err
	}
	return r.GetSlashableSharesForOperatorSetsBefore(ctx, operatorSets, uint32(currentBlock))
}

// Returns the strategies the `operatorSets` take into account, their
// operators, and the minimum amount of shares slashable by the
// `operatorSets` before a given timestamp.
// Timestamp must be in the future. Used to estimate future slashable stake.
// Not supported for M2 AVSs.
// Can return an error if the `AllocationManager` contract address was not provided, or due to
// errors in the underlying contract call.
func (r *ChainReader) GetSlashableSharesForOperatorSetsBefore(
	ctx context.Context,
	operatorSets []allocationmanager.OperatorSet,
	futureBlock uint32,
) ([]OperatorSetStakes, error) {
	operatorSetStakes := make([]OperatorSetStakes, len(operatorSets))
	for i, operatorSet := range operatorSets {
		operators, err := r.GetOperatorsForOperatorSet(ctx, operatorSet)
		if err != nil {
			return nil, err
		}

		strategies, err := r.GetStrategiesForOperatorSet(ctx, operatorSet)
		// If operator setId is 0 will fail on if above
		if err != nil {
			return nil, err
		}

		slashableShares, err := r.allocationManager.GetMinimumSlashableStake(
			&bind.CallOpts{Context: ctx},
			allocationmanager.OperatorSet{
				Id:  operatorSet.Id,
				Avs: operatorSet.Avs,
			},
			operators,
			strategies,
			futureBlock,
		)
		// This call should not fail since it's a getter
		if err != nil {
			return nil, err
		}

		operatorSetStakes[i] = OperatorSetStakes{
			OperatorSet:     operatorSet,
			Strategies:      strategies,
			Operators:       operators,
			SlashableStakes: slashableShares,
		}
	}

	return operatorSetStakes, nil
}

// Returns the time in blocks between an operator allocating slashable magnitude
// and the magnitude becoming slashable.
// Returns an error if the operator has not set an allocation delay.
// Can return an error if the `AllocationManager` contract address was not provided, or due to
// errors in the underlying contract call.
func (r *ChainReader) GetAllocationDelay(
	ctx context.Context,
	operatorAddress gethcommon.Address,
) (uint32, error) {
	if r.allocationManager == nil {
		return 0, errors.New("AllocationManager contract not provided")
	}
	isSet, delay, err := r.allocationManager.GetAllocationDelay(&bind.CallOpts{Context: ctx}, operatorAddress)
	// This call should not fail since it's a getter
	if err != nil {
		return 0, err
	}
	if !isSet {
		return 0, errors.New("allocation delay not set")
	}
	return delay, nil
}

// Returns a list of all operator sets the operator is registered for.
// Can return an error if the `AllocationManager` contract address was not provided, or due to
// errors in the underlying contract call.
func (r *ChainReader) GetRegisteredSets(
	ctx context.Context,
	operatorAddress gethcommon.Address,
) ([]allocationmanager.OperatorSet, error) {
	if r.allocationManager == nil {
		return nil, errors.New("AllocationManager contract not provided")
	}
	return r.allocationManager.GetRegisteredSets(&bind.CallOpts{Context: ctx}, operatorAddress)
}

// Returns `true` if `appointeeAddress` has permission to call the function with the given
// `selector` on the `target` contract, on behalf of `accountAddress`, and `false` otherwise.
// Can return an error if the `PermissionController` contract address was not provided, or due to
// errors in the underlying contract call.
func (r *ChainReader) CanCall(
	ctx context.Context,
	accountAddress gethcommon.Address,
	appointeeAddress gethcommon.Address,
	target gethcommon.Address,
	selector [4]byte,
) (bool, error) {
	canCall, err := r.permissionController.CanCall(
		&bind.CallOpts{Context: ctx},
		accountAddress,
		appointeeAddress,
		target,
		selector,
	)
	// This call should not fail since it's a getter
	if err != nil {
		return false, utils.WrapError("call to permission controller failed", err)
	}
	return canCall, nil
}

// Returns the list of appointees for a given account, target and function selector.
// Note that this doesn't include any of the appointed admins.
// Can return an error if the `PermissionController` contract address was not provided, or due to
// errors in the underlying contract call.
func (r *ChainReader) ListAppointees(
	ctx context.Context,
	accountAddress gethcommon.Address,
	target gethcommon.Address,
	selector [4]byte,
) ([]gethcommon.Address, error) {
	appointees, err := r.permissionController.GetAppointees(
		&bind.CallOpts{Context: ctx},
		accountAddress,
		target,
		selector,
	)
	// This call should not fail since it's a getter
	if err != nil {
		return nil, utils.WrapError("call to permission controller failed", err)
	}
	return appointees, nil
}

// Returns the list of permissions of an appointee for a given account.
// Can return an error if the `PermissionController` contract address was not provided, or due to
// errors in the underlying contract call.
func (r *ChainReader) ListAppointeePermissions(
	ctx context.Context,
	accountAddress gethcommon.Address,
	appointeeAddress gethcommon.Address,
) ([]gethcommon.Address, [][4]byte, error) {
	targets, selectors, err := r.permissionController.GetAppointeePermissions(
		&bind.CallOpts{Context: ctx},
		accountAddress,
		appointeeAddress,
	)
	// This call should not fail since it's a getter
	if err != nil {
		return nil, nil, utils.WrapError("call to permission controller failed", err)
	}
	return targets, selectors, nil
}

// Returns the pending admins of an account.
// Can return an error if the `PermissionController` contract address was not provided, or due to
// errors in the underlying contract call.
func (r *ChainReader) ListPendingAdmins(
	ctx context.Context,
	accountAddress gethcommon.Address,
) ([]gethcommon.Address, error) {
	pendingAdmins, err := r.permissionController.GetPendingAdmins(&bind.CallOpts{Context: ctx}, accountAddress)
	// This call should not fail since it's a getter
	if err != nil {
		return nil, utils.WrapError("call to permission controller failed", err)
	}
	return pendingAdmins, nil
}

// Returns the admins of an account.
// Can return an error if the `PermissionController` contract address was not provided, or due to
// errors in the underlying contract call.
func (r *ChainReader) ListAdmins(
	ctx context.Context,
	accountAddress gethcommon.Address,
) ([]gethcommon.Address, error) {
	pendingAdmins, err := r.permissionController.GetAdmins(&bind.CallOpts{Context: ctx}, accountAddress)
	// This call should not fail since it's a getter
	if err != nil {
		return nil, utils.WrapError("call to permission controller failed", err)
	}
	return pendingAdmins, nil
}

// Returns `true` if `pendingAdminAddress` is a pending admin for `accountAddress`, `false` otherwise.
// Can return an error if the `PermissionController` contract address was not provided, or due to
// errors in the underlying contract call.
func (r *ChainReader) IsPendingAdmin(
	ctx context.Context,
	accountAddress gethcommon.Address,
	pendingAdminAddress gethcommon.Address,
) (bool, error) {
	isPendingAdmin, err := r.permissionController.IsPendingAdmin(
		&bind.CallOpts{Context: ctx},
		accountAddress,
		pendingAdminAddress,
	)
	// This call should not fail since it's a getter
	if err != nil {
		return false, utils.WrapError("call to permission controller failed", err)
	}
	return isPendingAdmin, nil
}

// Returns `true` if `adminAddress` is an admin of `accountAddress`.
// Can return an error if the `PermissionController` contract address was not provided, or due to
// errors in the underlying contract call.
func (r *ChainReader) IsAdmin(
	ctx context.Context,
	accountAddress gethcommon.Address,
	adminAddress gethcommon.Address,
) (bool, error) {
	isAdmin, err := r.permissionController.IsAdmin(&bind.CallOpts{Context: ctx}, accountAddress, adminAddress)
	// This call should not fail since it's a getter
	if err != nil {
		return false, utils.WrapError("call to permission controller failed", err)
	}
	return isAdmin, nil
}
