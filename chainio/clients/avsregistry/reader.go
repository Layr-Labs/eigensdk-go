package avsregistry

import (
	"context"
	"errors"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	apkreg "github.com/Layr-Labs/eigensdk-go/contracts/bindings/BLSApkRegistry"
	opstateretriever "github.com/Layr-Labs/eigensdk-go/contracts/bindings/OperatorStateRetriever"
	regcoord "github.com/Layr-Labs/eigensdk-go/contracts/bindings/RegistryCoordinator"
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
}

type ChainReader struct {
	logger                  logging.Logger
	blsApkRegistryAddr      common.Address
	registryCoordinatorAddr common.Address
	registryCoordinator     *regcoord.ContractRegistryCoordinator
	operatorStateRetriever  *opstateretriever.ContractOperatorStateRetriever
	stakeRegistry           *stakeregistry.ContractStakeRegistry
	ethClient               eth.Client
}

func NewChainReader(
	registryCoordinatorAddr common.Address,
	blsApkRegistryAddr common.Address,
	registryCoordinator *regcoord.ContractRegistryCoordinator,
	operatorStateRetriever *opstateretriever.ContractOperatorStateRetriever,
	stakeRegistry *stakeregistry.ContractStakeRegistry,
	logger logging.Logger,
	ethClient eth.Client,
) *ChainReader {
	logger = logger.With(logging.ComponentKey, "avsregistry/ChainReader")

	return &ChainReader{
		blsApkRegistryAddr:      blsApkRegistryAddr,
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
	client eth.Client,
	logger logging.Logger,
) (*ChainReader, error) {
	bindings, err := NewBindingsFromConfig(cfg, client, logger)
	if err != nil {
		return nil, err
	}

	return NewChainReader(
		bindings.RegistryCoordinatorAddr,
		bindings.BlsApkRegistryAddr,
		bindings.RegistryCoordinator,
		bindings.OperatorStateRetriever,
		bindings.StakeRegistry,
		logger,
		client,
	), nil
}

// BuildAvsRegistryChainReader creates a new ChainReader
// Deprecated: Use NewReaderFromConfig instead
func BuildAvsRegistryChainReader(
	registryCoordinatorAddr common.Address,
	operatorStateRetrieverAddr common.Address,
	ethClient eth.Client,
	logger logging.Logger,
) (*ChainReader, error) {
	contractRegistryCoordinator, err := regcoord.NewContractRegistryCoordinator(registryCoordinatorAddr, ethClient)
	if err != nil {
		return nil, utils.WrapError("Failed to create contractRegistryCoordinator", err)
	}
	blsApkRegistryAddr, err := contractRegistryCoordinator.BlsApkRegistry(&bind.CallOpts{})
	if err != nil {
		return nil, utils.WrapError("Failed to get blsApkRegistryAddr", err)
	}
	stakeRegistryAddr, err := contractRegistryCoordinator.StakeRegistry(&bind.CallOpts{})
	if err != nil {
		return nil, utils.WrapError("Failed to get stakeRegistryAddr", err)
	}
	contractStakeRegistry, err := stakeregistry.NewContractStakeRegistry(stakeRegistryAddr, ethClient)
	if err != nil {
		return nil, utils.WrapError("Failed to create contractStakeRegistry", err)
	}
	contractOperatorStateRetriever, err := opstateretriever.NewContractOperatorStateRetriever(
		operatorStateRetrieverAddr,
		ethClient,
	)
	if err != nil {
		return nil, utils.WrapError("Failed to create contractOperatorStateRetriever", err)
	}
	return NewChainReader(
		registryCoordinatorAddr,
		blsApkRegistryAddr,
		contractRegistryCoordinator,
		contractOperatorStateRetriever,
		contractStakeRegistry,
		logger,
		ethClient,
	), nil
}

func (r *ChainReader) GetQuorumCount(opts *bind.CallOpts) (uint8, error) {
	if r.registryCoordinator == nil {
		return 0, errors.New("RegistryCoordinator contract not provided")
	}
	return r.registryCoordinator.QuorumCount(opts)
}

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

// GetOperatorStakeInQuorumsOfOperatorAtCurrentBlock could have race conditions
// it currently makes a bunch of calls to fetch "current block" information,
// so some of them could actually return information from different blocks
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

	quorumBitmap, err := r.registryCoordinator.GetCurrentQuorumBitmap(opts, operatorId)
	if err != nil {
		return nil, utils.WrapError("Failed to get operator quorums", err)
	}
	quorums := types.BitmapToQuorumIds(quorumBitmap)
	quorumStakes := make(map[types.QuorumNum]types.StakeAmount)
	for _, quorum := range quorums {
		stake, err := r.stakeRegistry.GetCurrentStake(
			&bind.CallOpts{},
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
		stopBlock = big.NewInt(int64(curBlockNum))
	}
	if blockRange == nil {
		blockRange = DefaultQueryBlockRange
	}

	operatorAddresses := make([]types.OperatorAddr, 0)
	operatorPubkeys := make([]types.OperatorPubkeys, 0)
	for i := startBlock; i.Cmp(stopBlock) <= 0; i.Add(i, blockRange) {
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
		stopBlock = big.NewInt(int64(curBlockNum))
	}
	if blockRange == nil {
		blockRange = DefaultQueryBlockRange
	}

	operatorIdToSocketMap := make(map[types.OperatorId]types.Socket)
	for i := startBlock; i.Cmp(stopBlock) <= 0; i.Add(i, blockRange) {
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
