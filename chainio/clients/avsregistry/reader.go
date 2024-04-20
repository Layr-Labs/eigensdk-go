package avsregistry

import (
	"context"
	"errors"
	"math"
	"math/big"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	gethcommon "github.com/ethereum/go-ethereum/common"

	apkreg "github.com/Layr-Labs/eigensdk-go/contracts/bindings/BLSApkRegistry"
	opstateretriever "github.com/Layr-Labs/eigensdk-go/contracts/bindings/OperatorStateRetriever"
	regcoord "github.com/Layr-Labs/eigensdk-go/contracts/bindings/RegistryCoordinator"
	stakeregistry "github.com/Layr-Labs/eigensdk-go/contracts/bindings/StakeRegistry"
)

type AvsRegistryReader interface {
	GetQuorumCount(opts *bind.CallOpts) (uint8, error)

	GetOperatorsStakeInQuorumsAtCurrentBlock(
		opts *bind.CallOpts,
		quorumNumbers types.QuorumNums,
	) ([][]opstateretriever.OperatorStateRetrieverOperator, error)

	GetOperatorsStakeInQuorumsAtBlock(
		opts *bind.CallOpts,
		quorumNumbers types.QuorumNums,
		blockNumber uint32,
	) ([][]opstateretriever.OperatorStateRetrieverOperator, error)

	GetOperatorAddrsInQuorumsAtCurrentBlock(
		opts *bind.CallOpts,
		quorumNumbers types.QuorumNums,
	) ([][]common.Address, error)

	GetOperatorsStakeInQuorumsOfOperatorAtBlock(
		opts *bind.CallOpts,
		operatorId types.OperatorId,
		blockNumber uint32,
	) (types.QuorumNums, [][]opstateretriever.OperatorStateRetrieverOperator, error)

	GetOperatorsStakeInQuorumsOfOperatorAtCurrentBlock(
		opts *bind.CallOpts,
		operatorId types.OperatorId,
	) (types.QuorumNums, [][]opstateretriever.OperatorStateRetrieverOperator, error)

	GetOperatorStakeInQuorumsOfOperatorAtCurrentBlock(
		opts *bind.CallOpts,
		operatorId types.OperatorId,
	) (map[types.QuorumNum]types.StakeAmount, error)

	GetCheckSignaturesIndices(
		opts *bind.CallOpts,
		referenceBlockNumber uint32,
		quorumNumbers types.QuorumNums,
		nonSignerOperatorIds []types.OperatorId,
	) (opstateretriever.OperatorStateRetrieverCheckSignaturesIndices, error)

	GetOperatorId(opts *bind.CallOpts, operatorAddress gethcommon.Address) ([32]byte, error)

	GetOperatorFromId(opts *bind.CallOpts, operatorId types.OperatorId) (gethcommon.Address, error)

	IsOperatorRegistered(opts *bind.CallOpts, operatorAddress gethcommon.Address) (bool, error)

	QueryExistingRegisteredOperatorPubKeys(
		ctx context.Context,
		startBlock *big.Int,
		stopBlock *big.Int,
	) ([]types.OperatorAddr, []types.OperatorPubkeys, error)

	QueryExistingRegisteredOperatorSockets(
		ctx context.Context,
		startBlock *big.Int,
		stopBlock *big.Int,
	) (map[types.OperatorId]types.Socket, error)
}

type AvsRegistryChainReader struct {
	logger                  logging.Logger
	blsApkRegistryAddr      gethcommon.Address
	registryCoordinatorAddr gethcommon.Address
	registryCoordinator     *regcoord.ContractRegistryCoordinator
	operatorStateRetriever  *opstateretriever.ContractOperatorStateRetriever
	stakeRegistry           *stakeregistry.ContractStakeRegistry
	ethClient               eth.Client
}

// forces AvsReader to implement the clients.ReaderInterface interface
var _ AvsRegistryReader = (*AvsRegistryChainReader)(nil)

func NewAvsRegistryChainReader(
	registryCoordinatorAddr gethcommon.Address,
	blsApkRegistryAddr gethcommon.Address,
	registryCoordinator *regcoord.ContractRegistryCoordinator,
	operatorStateRetriever *opstateretriever.ContractOperatorStateRetriever,
	stakeRegistry *stakeregistry.ContractStakeRegistry,
	logger logging.Logger,
	ethClient eth.Client,
) *AvsRegistryChainReader {
	return &AvsRegistryChainReader{
		blsApkRegistryAddr:      blsApkRegistryAddr,
		registryCoordinatorAddr: registryCoordinatorAddr,
		registryCoordinator:     registryCoordinator,
		operatorStateRetriever:  operatorStateRetriever,
		stakeRegistry:           stakeRegistry,
		logger:                  logger,
		ethClient:               ethClient,
	}
}

func BuildAvsRegistryChainReader(
	registryCoordinatorAddr gethcommon.Address,
	operatorStateRetrieverAddr gethcommon.Address,
	ethClient eth.Client,
	logger logging.Logger,
) (*AvsRegistryChainReader, error) {
	contractRegistryCoordinator, err := regcoord.NewContractRegistryCoordinator(registryCoordinatorAddr, ethClient)
	if err != nil {
		return nil, types.WrapError(errors.New("Failed to create contractRegistryCoordinator"), err)
	}
	blsApkRegistryAddr, err := contractRegistryCoordinator.BlsApkRegistry(&bind.CallOpts{})
	if err != nil {
		return nil, types.WrapError(errors.New("Failed to get blsApkRegistryAddr"), err)
	}
	stakeRegistryAddr, err := contractRegistryCoordinator.StakeRegistry(&bind.CallOpts{})
	if err != nil {
		return nil, types.WrapError(errors.New("Failed to get stakeRegistryAddr"), err)
	}
	contractStakeRegistry, err := stakeregistry.NewContractStakeRegistry(stakeRegistryAddr, ethClient)
	if err != nil {
		return nil, types.WrapError(errors.New("Failed to create contractStakeRegistry"), err)
	}
	contractOperatorStateRetriever, err := opstateretriever.NewContractOperatorStateRetriever(
		operatorStateRetrieverAddr,
		ethClient,
	)
	if err != nil {
		return nil, types.WrapError(errors.New("Failed to create contractOperatorStateRetriever"), err)
	}
	return NewAvsRegistryChainReader(
		registryCoordinatorAddr,
		blsApkRegistryAddr,
		contractRegistryCoordinator,
		contractOperatorStateRetriever,
		contractStakeRegistry,
		logger,
		ethClient,
	), nil
}

func (r *AvsRegistryChainReader) GetQuorumCount(opts *bind.CallOpts) (uint8, error) {
	return r.registryCoordinator.QuorumCount(opts)
}

func (r *AvsRegistryChainReader) GetOperatorsStakeInQuorumsAtCurrentBlock(
	opts *bind.CallOpts,
	quorumNumbers types.QuorumNums,
) ([][]opstateretriever.OperatorStateRetrieverOperator, error) {
	if opts.Context == nil {
		opts.Context = context.Background()
	}
	curBlock, err := r.ethClient.BlockNumber(opts.Context)
	if err != nil {
		return nil, types.WrapError(errors.New("Cannot get current block number"), err)
	}
	if curBlock > math.MaxUint32 {
		return nil, types.WrapError(errors.New("Current block number is too large to be converted to uint32"), err)
	}
	return r.GetOperatorsStakeInQuorumsAtBlock(opts, quorumNumbers, uint32(curBlock))
}

// the contract stores historical state, so blockNumber should be the block number of the state you want to query
// and the blockNumber in opts should be the block number of the latest block (or set to nil, which is equivalent)
func (r *AvsRegistryChainReader) GetOperatorsStakeInQuorumsAtBlock(
	opts *bind.CallOpts,
	quorumNumbers types.QuorumNums,
	blockNumber uint32,
) ([][]opstateretriever.OperatorStateRetrieverOperator, error) {
	operatorStakes, err := r.operatorStateRetriever.GetOperatorState(
		opts,
		r.registryCoordinatorAddr,
		quorumNumbers.UnderlyingType(),
		blockNumber)
	if err != nil {
		return nil, types.WrapError(errors.New("Failed to get operators state"), err)
	}
	return operatorStakes, nil
}

func (r *AvsRegistryChainReader) GetOperatorAddrsInQuorumsAtCurrentBlock(
	opts *bind.CallOpts,
	quorumNumbers types.QuorumNums,
) ([][]common.Address, error) {
	if opts.Context == nil {
		opts.Context = context.Background()
	}
	curBlock, err := r.ethClient.BlockNumber(opts.Context)
	if err != nil {
		return nil, types.WrapError(errors.New("Failed to get current block number"), err)
	}
	if curBlock > math.MaxUint32 {
		return nil, types.WrapError(errors.New("Current block number is too large to be converted to uint32"), err)
	}
	operatorStakes, err := r.operatorStateRetriever.GetOperatorState(
		opts,
		r.registryCoordinatorAddr,
		quorumNumbers.UnderlyingType(),
		uint32(curBlock),
	)
	if err != nil {
		return nil, types.WrapError(errors.New("Failed to get operators state"), err)
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

func (r *AvsRegistryChainReader) GetOperatorsStakeInQuorumsOfOperatorAtBlock(
	opts *bind.CallOpts,
	operatorId types.OperatorId,
	blockNumber uint32,
) (types.QuorumNums, [][]opstateretriever.OperatorStateRetrieverOperator, error) {
	quorumBitmap, operatorStakes, err := r.operatorStateRetriever.GetOperatorState0(
		opts,
		r.registryCoordinatorAddr,
		operatorId,
		blockNumber)
	if err != nil {
		return nil, nil, types.WrapError(errors.New("Failed to get operators state"), err)
	}
	quorums := types.BitmapToQuorumIds(quorumBitmap)
	return quorums, operatorStakes, nil
}

// opts will be modified to have the latest blockNumber, so make sure not to reuse it
// blockNumber in opts will be ignored, and the chain will be queried to get the latest blockNumber
func (r *AvsRegistryChainReader) GetOperatorsStakeInQuorumsOfOperatorAtCurrentBlock(
	opts *bind.CallOpts,
	operatorId types.OperatorId,
) (types.QuorumNums, [][]opstateretriever.OperatorStateRetrieverOperator, error) {
	if opts.Context == nil {
		opts.Context = context.Background()
	}
	curBlock, err := r.ethClient.BlockNumber(opts.Context)
	if err != nil {
		return nil, nil, types.WrapError(errors.New("Failed to get current block number"), err)
	}
	if curBlock > math.MaxUint32 {
		return nil, nil, types.WrapError(errors.New("Current block number is too large to be converted to uint32"), err)
	}
	opts.BlockNumber = big.NewInt(int64(curBlock))
	return r.GetOperatorsStakeInQuorumsOfOperatorAtBlock(opts, operatorId, uint32(curBlock))
}

// GetOperatorStakeInQuorumsOfOperatorAtCurrentBlock could have race conditions
// it currently makes a bunch of calls to fetch "current block" information,
// so some of them could actually return information from different blocks
func (r *AvsRegistryChainReader) GetOperatorStakeInQuorumsOfOperatorAtCurrentBlock(
	opts *bind.CallOpts,
	operatorId types.OperatorId,
) (map[types.QuorumNum]types.StakeAmount, error) {
	quorumBitmap, err := r.registryCoordinator.GetCurrentQuorumBitmap(opts, operatorId)
	if err != nil {
		return nil, types.WrapError(errors.New("Failed to get operator quorums"), err)
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
			return nil, types.WrapError(errors.New("Failed to get operator stake"), err)
		}
		quorumStakes[quorum] = stake
	}
	return quorumStakes, nil
}

func (r *AvsRegistryChainReader) GetCheckSignaturesIndices(
	opts *bind.CallOpts,
	referenceBlockNumber uint32,
	quorumNumbers types.QuorumNums,
	nonSignerOperatorIds []types.OperatorId,
) (opstateretriever.OperatorStateRetrieverCheckSignaturesIndices, error) {
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
		return opstateretriever.OperatorStateRetrieverCheckSignaturesIndices{}, types.WrapError(errors.New("Failed to get check signatures indices"), err)
	}
	return checkSignatureIndices, nil
}

func (r *AvsRegistryChainReader) GetOperatorId(
	opts *bind.CallOpts,
	operatorAddress gethcommon.Address,
) ([32]byte, error) {
	operatorId, err := r.registryCoordinator.GetOperatorId(
		opts,
		operatorAddress,
	)
	if err != nil {
		return [32]byte{}, types.WrapError(errors.New("Failed to get operator id"), err)
	}
	return operatorId, nil
}

func (r *AvsRegistryChainReader) GetOperatorFromId(
	opts *bind.CallOpts,
	operatorId types.OperatorId,
) (gethcommon.Address, error) {
	operatorAddress, err := r.registryCoordinator.GetOperatorFromId(
		opts,
		operatorId,
	)
	if err != nil {
		return gethcommon.Address{}, types.WrapError(errors.New("Failed to get operator address"), err)
	}
	return operatorAddress, nil
}

func (r *AvsRegistryChainReader) IsOperatorRegistered(
	opts *bind.CallOpts,
	operatorAddress gethcommon.Address,
) (bool, error) {
	operatorStatus, err := r.registryCoordinator.GetOperatorStatus(opts, operatorAddress)
	if err != nil {
		return false, types.WrapError(errors.New("Failed to get operator status"), err)
	}

	// 0 = NEVER_REGISTERED, 1 = REGISTERED, 2 = DEREGISTERED
	registeredWithAvs := operatorStatus == 1
	return registeredWithAvs, nil
}

func (r *AvsRegistryChainReader) QueryExistingRegisteredOperatorPubKeys(
	ctx context.Context,
	startBlock *big.Int,
	stopBlock *big.Int,
) ([]types.OperatorAddr, []types.OperatorPubkeys, error) {

	blsApkRegistryAbi, err := apkreg.ContractBLSApkRegistryMetaData.GetAbi()
	if err != nil {
		return nil, nil, types.WrapError(errors.New("Cannot get Abi"), err)
	}

	if startBlock == nil {
		startBlock = big.NewInt(0)
	}
	if stopBlock == nil {
		curBlockNum, err := r.ethClient.BlockNumber(ctx)
		if err != nil {
			return nil, nil, types.WrapError(errors.New("Cannot get current block number"), err)
		}
		stopBlock = big.NewInt(int64(curBlockNum))
	}

	operatorAddresses := make([]types.OperatorAddr, 0)
	operatorPubkeys := make([]types.OperatorPubkeys, 0)

	// eth_getLogs is limited to a 10,000 range, so we need to iterate over the range
	for i := startBlock; i.Cmp(stopBlock) <= 0; i.Add(i, big.NewInt(10_000)) {
		toBlock := big.NewInt(0).Add(i, big.NewInt(10_000))
		if toBlock.Cmp(stopBlock) > 0 {
			toBlock = stopBlock
		}
		query := ethereum.FilterQuery{
			FromBlock: i,
			ToBlock:   toBlock,
			Addresses: []gethcommon.Address{
				r.blsApkRegistryAddr,
			},
			Topics: [][]gethcommon.Hash{{blsApkRegistryAbi.Events["NewPubkeyRegistration"].ID}},
		}

		logs, err := r.ethClient.FilterLogs(ctx, query)
		if err != nil {
			return nil, nil, types.WrapError(errors.New("Cannot filter logs"), err)
		}
		r.logger.Debug("avsRegistryChainReader.QueryExistingRegisteredOperatorPubKeys", "numTransactionLogs", len(logs), "fromBlock", i, "toBlock", toBlock)

		for _, vLog := range logs {

			// get the operator address
			operatorAddr := gethcommon.HexToAddress(vLog.Topics[1].Hex())
			operatorAddresses = append(operatorAddresses, operatorAddr)

			event, err := blsApkRegistryAbi.Unpack("NewPubkeyRegistration", vLog.Data)
			if err != nil {
				return nil, nil, types.WrapError(errors.New("Cannot unpack event data"), err)
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

func (r *AvsRegistryChainReader) QueryExistingRegisteredOperatorSockets(
	ctx context.Context,
	startBlock *big.Int,
	stopBlock *big.Int,
) (map[types.OperatorId]types.Socket, error) {

	if startBlock == nil {
		startBlock = big.NewInt(0)
	}
	if stopBlock == nil {
		curBlockNum, err := r.ethClient.BlockNumber(ctx)
		if err != nil {
			return nil, types.WrapError(errors.New("Cannot get current block number"), err)
		}
		stopBlock = big.NewInt(int64(curBlockNum))
	}

	operatorIdToSocketMap := make(map[types.OperatorId]types.Socket)

	// eth_getLogs is limited to a 10,000 range, so we need to iterate over the range
	for i := startBlock; i.Cmp(stopBlock) <= 0; i.Add(i, big.NewInt(10_000)) {
		toBlock := big.NewInt(0).Add(i, big.NewInt(10_000))
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
			return nil, types.WrapError(errors.New("Cannot filter operator socket updates"), err)
		}
		r.logger.Debug("avsRegistryChainReader.QueryExistingRegisteredOperatorSockets", "fromBlock", i, "toBlock", toBlock)

		for update := socketUpdates; update.Event != nil; update.Next() {
			operatorIdToSocketMap[update.Event.OperatorId] = types.Socket(update.Event.Socket)
		}
	}
	return operatorIdToSocketMap, nil
}
