package avsregistry

import (
	"context"
	"errors"
	"math"
	"math/big"
	"slices"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	apkreg "github.com/Layr-Labs/eigensdk-go/contracts/bindings/BLSApkRegistry"
	opstateretriever "github.com/Layr-Labs/eigensdk-go/contracts/bindings/OperatorStateRetriever"
	regcoord "github.com/Layr-Labs/eigensdk-go/contracts/bindings/RegistryCoordinator"
	servicemanager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/ServiceManagerBase"
	stakeregistry "github.com/Layr-Labs/eigensdk-go/contracts/bindings/StakeRegistry"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/Layr-Labs/eigensdk-go/utils"
)

// DefaultQueryBlockRange different node providers have different eth_getLogs range limits.
// 10k is an arbitrary choice that should work for most
var DefaultQueryBlockRange = big.NewInt(10_000)

type Config struct {
	RegistryCoordinatorAddress    common.Address
	OperatorStateRetrieverAddress common.Address

	/// Setting this to true will disable the fetching of the AllocationManager address.
	/// This is useful for older deployments, which don't have the contract deployed.
	// TODO: remove this once mainnet is updated with the new contracts
	DontUseAllocationManager bool
	/// The address of the ServiceManager contract.
	ServiceManagerAddress common.Address
}

// The ChainReader provides methods to call the
// AVS registry contract's view functions.
type ChainReader struct {
	logger                  logging.Logger
	blsApkRegistryAddr      common.Address
	blsApkRegistry          *apkreg.ContractBLSApkRegistry
	registryCoordinatorAddr common.Address
	serviceManager          *servicemanager.ContractServiceManagerBase
	registryCoordinator     *regcoord.ContractRegistryCoordinator
	operatorStateRetriever  *opstateretriever.ContractOperatorStateRetriever
	stakeRegistry           *stakeregistry.ContractStakeRegistry
	ethClient               eth.HttpBackend
}

// Creates a new instance  of the ChainReader.
func NewChainReader(
	serviceManager *servicemanager.ContractServiceManagerBase,
	registryCoordinatorAddr common.Address,
	blsApkRegistryAddr common.Address,
	registryCoordinator *regcoord.ContractRegistryCoordinator,
	blsApkRegistry *apkreg.ContractBLSApkRegistry,
	operatorStateRetriever *opstateretriever.ContractOperatorStateRetriever,
	stakeRegistry *stakeregistry.ContractStakeRegistry,
	logger logging.Logger,
	ethClient eth.HttpBackend,
) *ChainReader {
	logger = logger.With(logging.ComponentKey, "avsregistry/ChainReader")

	return &ChainReader{
		blsApkRegistryAddr:      blsApkRegistryAddr,
		serviceManager:          serviceManager,
		blsApkRegistry:          blsApkRegistry,
		registryCoordinatorAddr: registryCoordinatorAddr,
		registryCoordinator:     registryCoordinator,
		operatorStateRetriever:  operatorStateRetriever,
		stakeRegistry:           stakeRegistry,
		logger:                  logger,
		ethClient:               ethClient,
	}
}

// NewReaderFromConfig creates a new ChainReader
func NewReaderFromConfig(
	cfg Config,
	client eth.HttpBackend,
	logger logging.Logger,
) (*ChainReader, error) {
	bindings, err := NewBindingsFromConfig(cfg, client, logger)
	if err != nil {
		return nil, err
	}

	return NewChainReader(
		bindings.ServiceManager,
		bindings.RegistryCoordinatorAddr,
		bindings.BlsApkRegistryAddr,
		bindings.RegistryCoordinator,
		bindings.BlsApkRegistry,
		bindings.OperatorStateRetriever,
		bindings.StakeRegistry,
		logger,
		client,
	), nil
}

// Returns the total quorum count read from the RegistryCoordinator
func (r *ChainReader) GetQuorumCount(opts *bind.CallOpts) (uint8, error) {
	if r.registryCoordinator == nil {
		return 0, errors.New("RegistryCoordinator contract not provided")
	}
	return r.registryCoordinator.QuorumCount(opts)
}

// Returns, for each quorum in `quorumNumbers`, a vector of the operators registered for
// that quorum at the current block, containing each operator's `operatorId` and `stake`.
func (r *ChainReader) GetOperatorsStakeInQuorumsAtCurrentBlock(
	opts *bind.CallOpts,
	quorumNumbers types.QuorumNums,
) ([][]opstateretriever.OperatorStateRetrieverOperator, error) {
	if opts.Context == nil {
		opts.Context = context.Background()
	}
	curBlock, err := r.ethClient.BlockNumber(opts.Context)
	if err != nil {
		return nil, utils.WrapError("Cannot get current block number", err)
	}
	if curBlock > math.MaxUint32 {
		return nil, utils.WrapError("Current block number is too large to be converted to uint32", err)
	}
	return r.GetOperatorsStakeInQuorumsAtBlock(opts, quorumNumbers, uint32(curBlock))
}

// the contract stores historical state, so blockNumber should be the block number of the state you want to query
// and the blockNumber in opts should be the block number of the latest block (or set to nil, which is equivalent)
func (r *ChainReader) GetOperatorsStakeInQuorumsAtBlock(
	opts *bind.CallOpts,
	quorumNumbers types.QuorumNums,
	blockNumber uint32,
) ([][]opstateretriever.OperatorStateRetrieverOperator, error) {
	if r.operatorStateRetriever == nil {
		return nil, errors.New("OperatorStateRetriever contract not provided")
	}

	operatorStakes, err := r.operatorStateRetriever.GetOperatorState(
		opts,
		r.registryCoordinatorAddr,
		quorumNumbers.UnderlyingType(),
		blockNumber)
	if err != nil {
		return nil, utils.WrapError("Failed to get operators state", err)
	}
	return operatorStakes, nil
}

// Returns, for each quorum in `quorumNumbers`, a vector of the addresses of the
// operators registered for that quorum at the current block.
func (r *ChainReader) GetOperatorAddrsInQuorumsAtCurrentBlock(
	opts *bind.CallOpts,
	quorumNumbers types.QuorumNums,
) ([][]common.Address, error) {
	if r.operatorStateRetriever == nil {
		return nil, errors.New("OperatorStateRetriever contract not provided")
	}
	if opts.Context == nil {
		opts.Context = context.Background()
	}
	curBlock, err := r.ethClient.BlockNumber(opts.Context)
	if err != nil {
		return nil, utils.WrapError("Failed to get current block number", err)
	}
	if curBlock > math.MaxUint32 {
		return nil, utils.WrapError("Current block number is too large to be converted to uint32", err)
	}
	operatorStakes, err := r.operatorStateRetriever.GetOperatorState(
		opts,
		r.registryCoordinatorAddr,
		quorumNumbers.UnderlyingType(),
		uint32(curBlock),
	)
	if err != nil {
		return nil, utils.WrapError("Failed to get operators state", err)
	}
	var quorumOperatorAddrs [][]common.Address
	for _, quorum := range operatorStakes {
		var operatorAddrs []common.Address
		for _, operator := range quorum {
			operatorAddrs = append(operatorAddrs, operator.Operator)
		}
		quorumOperatorAddrs = append(quorumOperatorAddrs, operatorAddrs)
	}
	return quorumOperatorAddrs, nil

}

// Returns a tuple containing
//   - An array with the quorum IDs in which the given operator is registered at the given block
//   - An array that contains, for each quorum, an array with the address, id and stake
//     of each operator registered in that quorum.
func (r *ChainReader) GetOperatorsStakeInQuorumsOfOperatorAtBlock(
	opts *bind.CallOpts,
	operatorId types.OperatorId,
	blockNumber uint32,
) (types.QuorumNums, [][]opstateretriever.OperatorStateRetrieverOperator, error) {
	if r.operatorStateRetriever == nil {
		return nil, nil, errors.New("OperatorStateRetriever contract not provided")
	}

	quorumBitmap, operatorStakes, err := r.operatorStateRetriever.GetOperatorState0(
		opts,
		r.registryCoordinatorAddr,
		operatorId,
		blockNumber)
	if err != nil {
		return nil, nil, utils.WrapError("Failed to get operators state", err)
	}
	quorums := types.BitmapToQuorumIds(quorumBitmap)
	return quorums, operatorStakes, nil
}

// opts will be modified to have the latest blockNumber, so make sure not to reuse it
// blockNumber in opts will be ignored, and the chain will be queried to get the latest blockNumber
func (r *ChainReader) GetOperatorsStakeInQuorumsOfOperatorAtCurrentBlock(
	opts *bind.CallOpts,
	operatorId types.OperatorId,
) (types.QuorumNums, [][]opstateretriever.OperatorStateRetrieverOperator, error) {
	if opts.Context == nil {
		opts.Context = context.Background()
	}
	curBlock, err := r.ethClient.BlockNumber(opts.Context)
	if err != nil {
		return nil, nil, utils.WrapError("Failed to get current block number", err)
	}
	if curBlock > math.MaxUint32 {
		return nil, nil, utils.WrapError("Current block number is too large to be converted to uint32", err)
	}
	opts.BlockNumber = big.NewInt(int64(curBlock))
	return r.GetOperatorsStakeInQuorumsOfOperatorAtBlock(opts, operatorId, uint32(curBlock))
}

// To avoid a possible race condition, this method must assure that all the calls
// are made with the same blockNumber.
// So, if the blockNumber and blockHash are not set in opts, blockNumber will be set
// to the latest block.
// All calls to the chain use `opts` parameter.
func (r *ChainReader) GetOperatorStakeInQuorumsOfOperatorAtCurrentBlock(
	opts *bind.CallOpts,
	operatorId types.OperatorId,
) (map[types.QuorumNum]types.StakeAmount, error) {
	if r.registryCoordinator == nil {
		return nil, errors.New("RegistryCoordinator contract not provided")
	}

	if r.stakeRegistry == nil {
		return nil, errors.New("StakeRegistry contract not provided")
	}

	// check if opts parameter has not a block number set (BlockNumber)
	var defaultHash common.Hash
	if opts.BlockNumber == nil && opts.BlockHash == defaultHash {
		// if not, set the block number to the latest block
		if opts.Context == nil {
			opts.Context = context.Background()
		}
		latestBlock, err := r.ethClient.BlockNumber(opts.Context)
		if err != nil {
			return nil, utils.WrapError("Failed to get latest block number", err)
		}
		opts.BlockNumber = big.NewInt(int64(latestBlock))
	}

	quorumBitmap, err := r.registryCoordinator.GetCurrentQuorumBitmap(opts, operatorId)
	if err != nil {
		return nil, utils.WrapError("Failed to get operator quorums", err)
	}
	quorums := types.BitmapToQuorumIds(quorumBitmap)
	quorumStakes := make(map[types.QuorumNum]types.StakeAmount)
	for _, quorum := range quorums {
		stake, err := r.stakeRegistry.GetCurrentStake(
			opts,
			operatorId,
			uint8(quorum),
		)
		if err != nil {
			return nil, utils.WrapError("Failed to get operator stake", err)
		}
		quorumStakes[quorum] = stake
	}
	return quorumStakes, nil
}

// This function computes the total weight of the operator in the quorum quorumNumber.
func (r *ChainReader) WeightOfOperatorForQuorum(
	opts *bind.CallOpts,
	quorumNumber uint8,
	operatorAddr common.Address,
) (types.StakeAmount, error) {
	if r.stakeRegistry == nil {
		return nil, errors.New("StakeRegistry contract not provided")
	}

	stake, err := r.stakeRegistry.WeightOfOperatorForQuorum(opts, quorumNumber, operatorAddr)
	if err != nil {
		return nil, utils.WrapError("Failed to get operator stake", err)
	}
	return stake, nil
}

// Returns the length of the dynamic array stored in strategyParams[quorumNumber] in the StakeRegistry contract.
func (r *ChainReader) StrategyParamsLength(
	opts *bind.CallOpts,
	quorumNumber uint8,
) (*big.Int, error) {
	if r.stakeRegistry == nil {
		return nil, errors.New("StakeRegistry contract not provided")
	}

	length, err := r.stakeRegistry.StrategyParamsLength(opts, quorumNumber)
	if err != nil {
		return nil, utils.WrapError("Failed to get strategy params length", err)
	}
	return length, nil
}

// Returns the strategy and weight multiplier for the index'th strategy in the specified quorum.
func (r *ChainReader) StrategyParamsByIndex(
	opts *bind.CallOpts,
	quorumNumber uint8,
	index *big.Int,
) (stakeregistry.IStakeRegistryTypesStrategyParams, error) {
	if r.stakeRegistry == nil {
		return stakeregistry.IStakeRegistryTypesStrategyParams{}, errors.New("StakeRegistry contract not provided")
	}

	param, err := r.stakeRegistry.StrategyParamsByIndex(opts, quorumNumber, index)
	if err != nil {
		return stakeregistry.IStakeRegistryTypesStrategyParams{}, utils.WrapError(
			"Failed to get strategy params by index",
			err,
		)
	}
	return param, nil
}

// Returns the length of an operator's stake history for the given quorum
func (r *ChainReader) GetStakeHistoryLength(
	opts *bind.CallOpts,
	operatorId types.OperatorId,
	quorumNumber uint8,
) (*big.Int, error) {
	if r.stakeRegistry == nil {
		return nil, errors.New("StakeRegistry contract not provided")
	}

	length, err := r.stakeRegistry.GetStakeHistoryLength(opts, operatorId, quorumNumber)
	if err != nil {
		return nil, utils.WrapError("Failed to get stake history length", err)
	}
	return length, nil
}

// Returns the entire operatorStakeHistory[operatorId][quorumNumber] array which contains the operator's
// stake update history for a given quorum.
func (r *ChainReader) GetStakeHistory(
	opts *bind.CallOpts,
	operatorId types.OperatorId,
	quorumNumber uint8,
) ([]stakeregistry.IStakeRegistryTypesStakeUpdate, error) {
	if r.stakeRegistry == nil {
		return []stakeregistry.IStakeRegistryTypesStakeUpdate{}, errors.New("StakeRegistry contract not provided")
	}

	stakeHistory, err := r.stakeRegistry.GetStakeHistory(opts, operatorId, quorumNumber)
	if err != nil {
		return []stakeregistry.IStakeRegistryTypesStakeUpdate{}, utils.WrapError("Failed to get stake history", err)
	}
	return stakeHistory, nil
}

// Returns the most recent stake weight for the `operatorId` for a certain quorum
func (r *ChainReader) GetLatestStakeUpdate(
	opts *bind.CallOpts,
	operatorId types.OperatorId,
	quorumNumber uint8,
) (stakeregistry.IStakeRegistryTypesStakeUpdate, error) {
	if r.stakeRegistry == nil {
		return stakeregistry.IStakeRegistryTypesStakeUpdate{}, errors.New("StakeRegistry contract not provided")
	}

	stakeUpdate, err := r.stakeRegistry.GetLatestStakeUpdate(opts, operatorId, quorumNumber)
	if err != nil {
		return stakeregistry.IStakeRegistryTypesStakeUpdate{}, utils.WrapError("Failed to get latest stake update", err)
	}
	return stakeUpdate, nil
}

// Returns the `index`-th entry in the `operatorStakeHistory[operatorId][quorumNumber]` array, i.e.,
// returns the `index`-th stake update for the operator.
func (r *ChainReader) GetStakeUpdateAtIndex(
	opts *bind.CallOpts,
	operatorId types.OperatorId,
	quorumNumber uint8,
	index *big.Int,
) (stakeregistry.IStakeRegistryTypesStakeUpdate, error) {
	if r.stakeRegistry == nil {
		return stakeregistry.IStakeRegistryTypesStakeUpdate{}, errors.New("StakeRegistry contract not provided")
	}

	stakeUpdate, err := r.stakeRegistry.GetStakeUpdateAtIndex(opts, quorumNumber, operatorId, index)
	if err != nil {
		return stakeregistry.IStakeRegistryTypesStakeUpdate{}, utils.WrapError(
			"Failed to get stake update at index",
			err,
		)
	}
	return stakeUpdate, nil
}

// Returns the stake of the operator for the provided `quorumNumber` at the given `blockNumber`
func (r *ChainReader) GetStakeAtBlockNumber(
	opts *bind.CallOpts,
	operatorId types.OperatorId,
	quorumNumber uint8,
	blockNumber uint32,
) (types.StakeAmount, error) {
	if r.stakeRegistry == nil {
		return nil, errors.New("StakeRegistry contract not provided")
	}

	stake, err := r.stakeRegistry.GetStakeAtBlockNumber(opts, operatorId, quorumNumber, blockNumber)
	if err != nil {
		return nil, utils.WrapError("Failed to get stake at block number", err)
	}
	return stake, nil
}

// Returns the indices of the operator stakes for the provided `quorumNumber` at the given `blockNumber`
func (r *ChainReader) GetStakeUpdateIndexAtBlockNumber(
	opts *bind.CallOpts,
	operatorId types.OperatorId,
	quorumNumber uint8,
	blockNumber uint32,
) (uint32, error) {
	if r.stakeRegistry == nil {
		return 0, errors.New("StakeRegistry contract not provided")
	}

	index, err := r.stakeRegistry.GetStakeUpdateIndexAtBlockNumber(
		opts,
		operatorId,
		quorumNumber,
		blockNumber,
	)
	if err != nil {
		return 0, utils.WrapError("Failed to get stake update index at block number", err)
	}
	return index, nil
}

// Returns the stake weight corresponding to `operatorId` for quorum `quorumNumber`, at the
// `index`-th entry in the operator's stake history array if it was the operator's
// stake at `blockNumber`.
func (r *ChainReader) GetStakeAtBlockNumberAndIndex(
	opts *bind.CallOpts,
	operatorId types.OperatorId,
	quorumNumber uint8,
	blockNumber uint32,
	index *big.Int,
) (types.StakeAmount, error) {
	if r.stakeRegistry == nil {
		return nil, errors.New("StakeRegistry contract not provided")
	}

	stake, err := r.stakeRegistry.GetStakeAtBlockNumberAndIndex(
		opts,
		quorumNumber,
		blockNumber,
		operatorId,
		index,
	)
	if err != nil {
		return nil, utils.WrapError("Failed to get stake at block number and index", err)
	}
	return stake, nil
}

// Returns the length of the total stake history for the given quorum
func (r *ChainReader) GetTotalStakeHistoryLength(
	opts *bind.CallOpts,
	quorumNumber uint8,
) (*big.Int, error) {
	if r.stakeRegistry == nil {
		return nil, errors.New("StakeRegistry contract not provided")
	}

	length, err := r.stakeRegistry.GetTotalStakeHistoryLength(opts, quorumNumber)
	if err != nil {
		return nil, utils.WrapError("Failed to get total stake history length", err)
	}

	return length, nil
}

// Returns a struct containing the indices of the quorum members that signed,
// and the ones that didn't.
func (r *ChainReader) GetCheckSignaturesIndices(
	opts *bind.CallOpts,
	referenceBlockNumber uint32,
	quorumNumbers types.QuorumNums,
	nonSignerOperatorIds []types.OperatorId,
) (opstateretriever.OperatorStateRetrieverCheckSignaturesIndices, error) {
	if r.operatorStateRetriever == nil {
		return opstateretriever.OperatorStateRetrieverCheckSignaturesIndices{}, errors.New(
			"OperatorStateRetriever contract not provided",
		)
	}

	nonSignerOperatorIdsBytes := make([][32]byte, len(nonSignerOperatorIds))
	for i, id := range nonSignerOperatorIds {
		nonSignerOperatorIdsBytes[i] = id
	}
	checkSignatureIndices, err := r.operatorStateRetriever.GetCheckSignaturesIndices(
		opts,
		r.registryCoordinatorAddr,
		referenceBlockNumber,
		quorumNumbers.UnderlyingType(),
		nonSignerOperatorIdsBytes,
	)
	if err != nil {
		return opstateretriever.OperatorStateRetrieverCheckSignaturesIndices{}, utils.WrapError(
			"Failed to get check signatures indices",
			err,
		)
	}
	return checkSignatureIndices, nil
}

// Returns the stake weight from the latest entry in the quorum's stake history
func (r *ChainReader) GetCurrentTotalStake(
	opts *bind.CallOpts,
	quorumNumber uint8,
) (types.StakeAmount, error) {
	if r.stakeRegistry == nil {
		return nil, errors.New("StakeRegistry contract not provided")
	}

	stake, err := r.stakeRegistry.GetCurrentTotalStake(opts, quorumNumber)
	if err != nil {
		return nil, utils.WrapError("Failed to get current total stake", err)
	}
	return stake, nil
}

func (r *ChainReader) GetTotalStakeUpdateAtIndex(
	opts *bind.CallOpts,
	quorumNumber uint8,
	index *big.Int,
) (stakeregistry.IStakeRegistryTypesStakeUpdate, error) {
	if r.stakeRegistry == nil {
		return stakeregistry.IStakeRegistryTypesStakeUpdate{}, errors.New("StakeRegistry contract not provided")
	}

	stakeUpdate, err := r.stakeRegistry.GetTotalStakeUpdateAtIndex(opts, quorumNumber, index)
	if err != nil {
		return stakeregistry.IStakeRegistryTypesStakeUpdate{}, utils.WrapError(
			"Failed to get total stake update at index",
			err,
		)
	}
	return stakeUpdate, nil
}

// Returns the total stake weight for the specified quorum at the `index`-th entry in the
// stake history array if it was the stake at the specified blockNumber.
func (r *ChainReader) GetTotalStakeAtBlockNumberFromIndex(
	opts *bind.CallOpts,
	quorumNumber uint8,
	blockNumber uint32,
	index *big.Int,
) (types.StakeAmount, error) {
	if r.stakeRegistry == nil {
		return nil, errors.New("StakeRegistry contract not provided")
	}

	stake, err := r.stakeRegistry.GetTotalStakeAtBlockNumberFromIndex(
		opts,
		quorumNumber,
		blockNumber,
		index,
	)
	if err != nil {
		return nil, utils.WrapError("Failed to get total stake at block number from index", err)
	}
	return stake, nil
}

// Returns the indices of the total stakes for the provided quorumNumbers at the given blockNumber
func (r *ChainReader) GetTotalStakeIndicesAtBlockNumber(
	opts *bind.CallOpts,
	quorumNumbers types.QuorumNums,
	blockNumber uint32,
) ([]uint32, error) {
	if r.stakeRegistry == nil {
		return []uint32{}, errors.New("StakeRegistry contract not provided")
	}

	indices, err := r.stakeRegistry.GetTotalStakeIndicesAtBlockNumber(opts, blockNumber, quorumNumbers.UnderlyingType())
	if err != nil {
		return []uint32{}, utils.WrapError("Failed to get total stake indices at block number", err)
	}
	return indices, nil
}

func (r *ChainReader) GetMinimumStakeForQuorum(
	opts *bind.CallOpts,
	quorumNumber uint8,
) (types.StakeAmount, error) {
	if r.stakeRegistry == nil {
		return nil, errors.New("StakeRegistry contract not provided")
	}

	stake, err := r.stakeRegistry.MinimumStakeForQuorum(opts, quorumNumber)
	if err != nil {
		return nil, utils.WrapError("Failed to get minimum stake for quorum", err)
	}
	return stake, nil
}

func (r *ChainReader) GetStrategyParamsAtIndex(
	opts *bind.CallOpts,
	quorumNumber uint8,
	index *big.Int,
) (stakeregistry.IStakeRegistryTypesStrategyParams, error) {
	if r.stakeRegistry == nil {
		return stakeregistry.IStakeRegistryTypesStrategyParams{}, errors.New("StakeRegistry contract not provided")
	}
	params, err := r.stakeRegistry.StrategyParams(opts, quorumNumber, index)
	if err != nil {
		return stakeregistry.IStakeRegistryTypesStrategyParams{}, utils.WrapError("Failed to get strategy params", err)
	}
	return params, nil
}

func (r *ChainReader) GetStrategyPerQuorumAtIndex(
	opts *bind.CallOpts,
	quorumNumber uint8,
	index *big.Int,
) (common.Address, error) {
	if r.stakeRegistry == nil {
		return common.Address{}, errors.New("StakeRegistry contract not provided")
	}

	strategy, err := r.stakeRegistry.StrategiesPerQuorum(opts, quorumNumber, index)
	if err != nil {
		return common.Address{}, utils.WrapError("Failed to get strategies per quorum", err)
	}
	return strategy, nil
}

// Returns the list of strategies that the AVS supports for restaking.
// The list returned contains no duplicates.
func (r *ChainReader) GetRestakeableStrategies(opts *bind.CallOpts) ([]common.Address, error) {
	if r.serviceManager == nil {
		return nil, errors.New("ServiceManager contract not provided")
	}

	strategies, err := r.serviceManager.GetRestakeableStrategies(opts)
	if err != nil {
		return nil, utils.WrapError("Failed to get restakeable strategies", err)
	}
	if len(strategies) == 0 {
		return strategies, nil
	}
	return removeDuplicateStrategies(strategies), nil
}

// Returns the list of strategies that the operator has potentially restaked on the AVS.
// No guarantee is made on whether the operator has shares on each of the returned strategies.
func (r *ChainReader) GetOperatorRestakedStrategies(
	opts *bind.CallOpts,
	operator common.Address,
) ([]common.Address, error) {
	if r.serviceManager == nil {
		return nil, errors.New("ServiceManager contract not provided")
	}

	strategies, err := r.serviceManager.GetOperatorRestakedStrategies(opts, operator)
	if err != nil {
		return nil, utils.WrapError("Failed to get operator restaked strategies", err)
	}
	return removeDuplicateStrategies(strategies), nil
}

func (r *ChainReader) GetStakeTypePerQuorum(
	opts *bind.CallOpts,
	quorumNumber uint8,
) (uint8, error) {
	if r.stakeRegistry == nil {
		return 0, errors.New("StakeRegistry contract not provided")
	}

	stakeType, err := r.stakeRegistry.StakeTypePerQuorum(opts, quorumNumber)
	if err != nil {
		return 0, utils.WrapError("Failed to get stake type per quorum", err)
	}
	return stakeType, nil
}

func (r *ChainReader) GetSlashableStakeLookAheadPerQuorum(
	opts *bind.CallOpts,
	quorumNumber uint8,
) (uint32, error) {
	if r.stakeRegistry == nil {
		return 0, errors.New("StakeRegistry contract not provided")
	}

	lookAhead, err := r.stakeRegistry.SlashableStakeLookAheadPerQuorum(opts, quorumNumber)
	if err != nil {
		return 0, utils.WrapError("Failed to get slashable stake look ahead per quorum", err)
	}
	return lookAhead, nil
}

// Given an operator address, returns its ID.
func (r *ChainReader) GetOperatorId(
	opts *bind.CallOpts,
	operatorAddress common.Address,
) ([32]byte, error) {
	if r.registryCoordinator == nil {
		return [32]byte{}, errors.New("RegistryCoordinator contract not provided")
	}

	operatorId, err := r.registryCoordinator.GetOperatorId(
		opts,
		operatorAddress,
	)
	if err != nil {
		return [32]byte{}, utils.WrapError("Failed to get operator id", err)
	}
	return operatorId, nil
}

// Given an operator ID, returns its address.
func (r *ChainReader) GetOperatorFromId(
	opts *bind.CallOpts,
	operatorId types.OperatorId,
) (common.Address, error) {
	if r.registryCoordinator == nil {
		return common.Address{}, errors.New("RegistryCoordinator contract not provided")
	}

	operatorAddress, err := r.registryCoordinator.GetOperatorFromId(
		opts,
		operatorId,
	)
	if err != nil {
		return common.Address{}, utils.WrapError("Failed to get operator address", err)
	}
	return operatorAddress, nil
}

// Returns an array of booleans, where the boolean at index i represents
// whether the operator is registered for the quorum i.
func (r *ChainReader) QueryRegistrationDetail(
	opts *bind.CallOpts,
	operatorAddress common.Address,
) ([]bool, error) {
	operatorId, err := r.GetOperatorId(opts, operatorAddress)
	if err != nil {
		return nil, utils.WrapError("Failed to get operator id", err)
	}
	value, err := r.registryCoordinator.GetCurrentQuorumBitmap(opts, operatorId)
	if err != nil {
		return nil, utils.WrapError("Failed to get operator quorums", err)
	}
	numBits := value.BitLen()
	var quorums []bool
	for i := 0; i < numBits; i++ {
		quorums = append(quorums, value.Int64()&(1<<i) != 0)
	}
	if len(quorums) == 0 {
		numQuorums, err := r.GetQuorumCount(opts)
		if err != nil {
			return nil, utils.WrapError("Failed to get quorum count", err)
		}
		for i := uint8(0); i < numQuorums; i++ {
			quorums = append(quorums, false)
		}
	}
	return quorums, nil
}

// Returns true if the operator is registered, false otherwise.
func (r *ChainReader) IsOperatorRegistered(
	opts *bind.CallOpts,
	operatorAddress common.Address,
) (bool, error) {
	if r.registryCoordinator == nil {
		return false, errors.New("RegistryCoordinator contract not provided")
	}

	operatorStatus, err := r.registryCoordinator.GetOperatorStatus(opts, operatorAddress)
	if err != nil {
		return false, utils.WrapError("Failed to get operator status", err)
	}

	// 0 = NEVER_REGISTERED, 1 = REGISTERED, 2 = DEREGISTERED
	registeredWithAvs := operatorStatus == 1
	return registeredWithAvs, nil
}

// Returns the operator's operatorId (pubkeyhash) given their address
func (r *ChainReader) GetOperatorIdFromOperatorAddress(
	opts *bind.CallOpts,
	operatorAddress common.Address,
) ([32]byte, error) {
	if r.blsApkRegistry == nil {
		return [32]byte{}, errors.New("BLSApkRegistry contract not provided")
	}

	operatorPubkeyHash, err := r.blsApkRegistry.OperatorToPubkeyHash(opts, operatorAddress)
	if err != nil {
		return [32]byte{}, utils.WrapError("Failed to get operator pubkey hash", err)
	}
	return operatorPubkeyHash, nil
}

// Returns the operator address given their operator ID (their pubkeyhash)
func (r *ChainReader) GetOperatorAddressFromOperatorId(
	opts *bind.CallOpts,
	operatorPubkeyHash [32]byte,
) (common.Address, error) {
	if r.blsApkRegistry == nil {
		return common.Address{}, errors.New("BLSApkRegistry contract not provided")
	}

	operatorAddress, err := r.blsApkRegistry.PubkeyHashToOperator(opts, operatorPubkeyHash)
	if err != nil {
		return common.Address{}, utils.WrapError("Failed to get operator address", err)
	}
	return operatorAddress, nil
}

// Returns operator's BLS public key in G1 given their address
func (r *ChainReader) GetPubkeyFromOperatorAddress(
	opts *bind.CallOpts,
	operatorAddress common.Address,
) (bls.G1Point, error) {
	if r.blsApkRegistry == nil {
		return bls.G1Point{}, errors.New("BLSApkRegistry contract not provided")
	}

	operatorPubkey, err := r.blsApkRegistry.OperatorToPubkey(opts, operatorAddress)
	if err != nil {
		return bls.G1Point{}, utils.WrapError("Failed to get operator pubkey", err)
	}

	operatorPubkeyG1 := bls.NewG1Point(operatorPubkey.X, operatorPubkey.Y)

	return *operatorPubkeyG1, nil
}

// Stores the history of aggregate public key updates for `quorumNumber` at `index`
func (r *ChainReader) GetApkUpdate(
	opts *bind.CallOpts,
	quorumNumber uint8,
	index *big.Int,
) (apkreg.IBLSApkRegistryTypesApkUpdate, error) {
	if r.blsApkRegistry == nil {
		return apkreg.IBLSApkRegistryTypesApkUpdate{}, errors.New("StakeRegistry contract not provided")
	}

	update, err := r.blsApkRegistry.ApkHistory(opts, quorumNumber, index)
	if err != nil {
		return apkreg.IBLSApkRegistryTypesApkUpdate{}, utils.WrapError("Failed to get apk history", err)
	}

	apkUpdate := apkreg.IBLSApkRegistryTypesApkUpdate{
		ApkHash:               update.ApkHash,
		UpdateBlockNumber:     update.UpdateBlockNumber,
		NextUpdateBlockNumber: update.NextUpdateBlockNumber,
	}

	return apkUpdate, nil
}

// Gets the current aggregate bls pubkey for a given quorum
func (r *ChainReader) GetCurrentApk(
	opts *bind.CallOpts,
	quorumNumber uint8,
) (bls.G1Point, error) {
	if r.blsApkRegistry == nil {
		return bls.G1Point{}, errors.New("StakeRegistry contract not provided")
	}

	apk, err := r.blsApkRegistry.CurrentApk(opts, quorumNumber)
	if err != nil {
		return bls.G1Point{}, utils.WrapError("Failed to get current apk", err)
	}

	apkG1 := bls.NewG1Point(apk.X, apk.Y)
	return *apkG1, nil
}

// Queries existing operators for a particular block range.
// Returns two arrays. The first one contains the addresses
// of the operators, and the second contains their corresponding public keys.
func (r *ChainReader) QueryExistingRegisteredOperatorPubKeys(
	ctx context.Context,
	startBlock *big.Int,
	stopBlock *big.Int,
	blockRange *big.Int,
) ([]types.OperatorAddr, []types.OperatorPubkeys, error) {
	blsApkRegistryAbi, err := apkreg.ContractBLSApkRegistryMetaData.GetAbi()
	if err != nil {
		return nil, nil, utils.WrapError("Cannot get Abi", err)
	}

	if startBlock == nil {
		startBlock = big.NewInt(0)
	}
	if stopBlock == nil {
		curBlockNum, err := r.ethClient.BlockNumber(ctx)
		if err != nil {
			return nil, nil, utils.WrapError("Cannot get current block number", err)
		}
		stopBlock = new(big.Int).SetUint64(curBlockNum)
	}
	if blockRange == nil {
		blockRange = DefaultQueryBlockRange
	}

	operatorAddresses := make([]types.OperatorAddr, 0)
	operatorPubkeys := make([]types.OperatorPubkeys, 0)
	// QueryExistingRegisteredOperatorPubKeys and QueryExistingRegisteredOperatorSockets
	// both run in parallel and they read and mutate the same variable startBlock,
	// so we clone it to prevent the race condition.
	// TODO: we might want to eventually change the function signature to pass a uint,
	// but that would be a breaking change
	for i := new(big.Int).Set(startBlock); i.Cmp(stopBlock) <= 0; i.Add(i, blockRange) {
		// Subtract 1 since FilterQuery is inclusive
		toBlock := big.NewInt(0).Add(i, big.NewInt(0).Sub(blockRange, big.NewInt(1)))
		if toBlock.Cmp(stopBlock) > 0 {
			toBlock = stopBlock
		}
		query := ethereum.FilterQuery{
			FromBlock: i,
			ToBlock:   toBlock,
			Addresses: []common.Address{
				r.blsApkRegistryAddr,
			},
			Topics: [][]common.Hash{{blsApkRegistryAbi.Events["NewPubkeyRegistration"].ID}},
		}

		logs, err := r.ethClient.FilterLogs(ctx, query)
		if err != nil {
			return nil, nil, utils.WrapError("Cannot filter logs", err)
		}
		r.logger.Debug(
			"avsRegistryChainReader.QueryExistingRegisteredOperatorPubKeys",
			"numTransactionLogs",
			len(logs),
			"fromBlock",
			i,
			"toBlock",
			toBlock,
		)

		for _, vLog := range logs {
			// get the operator address
			operatorAddr := common.HexToAddress(vLog.Topics[1].Hex())
			operatorAddresses = append(operatorAddresses, operatorAddr)

			event, err := blsApkRegistryAbi.Unpack("NewPubkeyRegistration", vLog.Data)
			if err != nil {
				return nil, nil, utils.WrapError("Cannot unpack event data", err)
			}

			G1Pubkey := event[0].(struct {
				X *big.Int "json:\"X\""
				Y *big.Int "json:\"Y\""
			})

			G2Pubkey := event[1].(struct {
				X [2]*big.Int "json:\"X\""
				Y [2]*big.Int "json:\"Y\""
			})

			operatorPubkey := types.OperatorPubkeys{
				G1Pubkey: bls.NewG1Point(
					G1Pubkey.X,
					G1Pubkey.Y,
				),
				G2Pubkey: bls.NewG2Point(
					G2Pubkey.X,
					G2Pubkey.Y,
				),
			}

			operatorPubkeys = append(operatorPubkeys, operatorPubkey)
		}
	}

	return operatorAddresses, operatorPubkeys, nil
}

// Queries existing operator sockets for a particular block range.
// Returns a mapping containing operator IDs as keys and their
// corresponding sockets as values.
func (r *ChainReader) QueryExistingRegisteredOperatorSockets(
	ctx context.Context,
	startBlock *big.Int,
	stopBlock *big.Int,
	blockRange *big.Int,
) (map[types.OperatorId]types.Socket, error) {
	if r.registryCoordinator == nil {
		return nil, errors.New("RegistryCoordinator contract not provided")
	}

	if startBlock == nil {
		startBlock = big.NewInt(0)
	}
	if stopBlock == nil {
		curBlockNum, err := r.ethClient.BlockNumber(ctx)
		if err != nil {
			return nil, utils.WrapError("Cannot get current block number", err)
		}
		stopBlock = new(big.Int).SetUint64(curBlockNum)
	}
	if blockRange == nil {
		blockRange = DefaultQueryBlockRange
	}

	operatorIdToSocketMap := make(map[types.OperatorId]types.Socket)
	// QueryExistingRegisteredOperatorPubKeys and QueryExistingRegisteredOperatorSockets
	// both run in parallel and they read and mutate the same variable startBlock,
	// so we clone it to prevent the race condition.
	// TODO: we might want to eventually change the function signature to pass a uint,
	// but that would be a breaking change
	for i := new(big.Int).Set(startBlock); i.Cmp(stopBlock) <= 0; i.Add(i, blockRange) {
		// Subtract 1 since FilterQuery is inclusive
		toBlock := big.NewInt(0).Add(i, big.NewInt(0).Sub(blockRange, big.NewInt(1)))
		if toBlock.Cmp(stopBlock) > 0 {
			toBlock = stopBlock
		}

		end := toBlock.Uint64()

		filterOpts := &bind.FilterOpts{
			Start: i.Uint64(),
			End:   &end,
		}
		socketUpdates, err := r.registryCoordinator.FilterOperatorSocketUpdate(filterOpts, nil)
		if err != nil {
			return nil, utils.WrapError("Cannot filter operator socket updates", err)
		}

		numSocketUpdates := 0
		for socketUpdates.Next() {
			operatorIdToSocketMap[socketUpdates.Event.OperatorId] = types.Socket(socketUpdates.Event.Socket)
			numSocketUpdates++
		}
		r.logger.Debug(
			"avsRegistryChainReader.QueryExistingRegisteredOperatorSockets",
			"numTransactionLogs",
			numSocketUpdates,
			"fromBlock",
			i,
			"toBlock",
			toBlock,
		)
	}
	return operatorIdToSocketMap, nil
}

// Removes duplicates from the given list of strategies.
func removeDuplicateStrategies(strategies []common.Address) []common.Address {
	slices.SortFunc(strategies, common.Address.Cmp)
	uniqueStrategies := make([]common.Address, 0, len(strategies))
	lastElement := strategies[0]
	uniqueStrategies = append(uniqueStrategies, lastElement)
	for i := range uniqueStrategies[1:] {
		if strategies[i] == lastElement {
			continue
		}
		lastElement = strategies[i]
		uniqueStrategies = append(uniqueStrategies, lastElement)
	}
	return uniqueStrategies
}
