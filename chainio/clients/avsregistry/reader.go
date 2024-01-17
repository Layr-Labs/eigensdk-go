package avsregistry

import (
	"math"
	"math/big"

	"github.com/Layr-Labs/eigensdk-go-master/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go-master/logging"
	"github.com/Layr-Labs/eigensdk-go-master/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethcommon "github.com/ethereum/go-ethereum/common"

	blsoperatorstateretriever "github.com/Layr-Labs/eigensdk-go-master/contracts/bindings/BLSOperatorStateRetriever"
	blsregistrycoordinator "github.com/Layr-Labs/eigensdk-go-master/contracts/bindings/BLSRegistryCoordinatorWithIndices"
	stakeregistry "github.com/Layr-Labs/eigensdk-go-master/contracts/bindings/StakeRegistry"
)

type AvsRegistryReader interface {
	GetOperatorsStakeInQuorumsAtBlock(
		opts *bind.CallOpts,
		quorumNumbers []byte,
		blockNumber uint32,
	) ([][]blsoperatorstateretriever.BLSOperatorStateRetrieverOperator, error)

	GetOperatorStakeInQuorumsOfOperatorAtCurrentBlock(
		opts *bind.CallOpts,
		operatorId types.OperatorId,
	) (map[types.QuorumNum]types.StakeAmount, error)

	GetOperatorsStakeInQuorumsOfOperatorAtBlock(
		opts *bind.CallOpts,
		operatorId types.OperatorId,
		blockNumber uint32,
	) ([]types.QuorumNum, [][]blsoperatorstateretriever.BLSOperatorStateRetrieverOperator, error)

	GetOperatorsStakeInQuorumsOfOperatorAtCurrentBlock(
		opts *bind.CallOpts,
		operatorId types.OperatorId,
	) ([]types.QuorumNum, [][]blsoperatorstateretriever.BLSOperatorStateRetrieverOperator, error)

	GetCheckSignaturesIndices(
		opts *bind.CallOpts,
		referenceBlockNumber uint32,
		quorumNumbers []byte,
		nonSignerOperatorIds [][32]byte,
	) (blsoperatorstateretriever.BLSOperatorStateRetrieverCheckSignaturesIndices, error)

	GetOperatorId(opts *bind.CallOpts, operatorAddress gethcommon.Address) ([32]byte, error)

	IsOperatorRegistered(opts *bind.CallOpts, operatorAddress gethcommon.Address) (bool, error)
}

type AvsRegistryChainReader struct {
	logger                    logging.Logger
	registryCoordinatorAddr   gethcommon.Address
	registryCoordinator       *blsregistrycoordinator.ContractBLSRegistryCoordinatorWithIndices
	blsOperatorStateRetriever *blsoperatorstateretriever.ContractBLSOperatorStateRetriever
	stakeRegistry             *stakeregistry.ContractStakeRegistry
	ethClient                 eth.EthClient
}

// forces AvsReader to implement the clients.ReaderInterface interface
var _ AvsRegistryReader = (*AvsRegistryChainReader)(nil)

func NewAvsRegistryReader(
	registryCoordinatorAddr gethcommon.Address,
	registryCoordinator *blsregistrycoordinator.ContractBLSRegistryCoordinatorWithIndices,
	blsOperatorStateRetriever *blsoperatorstateretriever.ContractBLSOperatorStateRetriever,
	stakeRegistry *stakeregistry.ContractStakeRegistry,
	logger logging.Logger,
	ethClient eth.EthClient,
) (*AvsRegistryChainReader, error) {
	return &AvsRegistryChainReader{
		registryCoordinatorAddr:   registryCoordinatorAddr,
		registryCoordinator:       registryCoordinator,
		blsOperatorStateRetriever: blsOperatorStateRetriever,
		stakeRegistry:             stakeRegistry,
		logger:                    logger,
		ethClient:                 ethClient,
	}, nil
}

// the contract stores historical state, so blockNumber should be the block number of the state you want to query
// and the blockNumber in opts should be the block number of the latest block (or set to nil, which is equivalent)
func (r *AvsRegistryChainReader) GetOperatorsStakeInQuorumsAtBlock(
	opts *bind.CallOpts,
	quorumNumbers []byte,
	blockNumber uint32,
) ([][]blsoperatorstateretriever.BLSOperatorStateRetrieverOperator, error) {
	operatorStakes, err := r.blsOperatorStateRetriever.GetOperatorState(
		opts,
		r.registryCoordinatorAddr,
		quorumNumbers,
		blockNumber)
	if err != nil {
		r.logger.Error("Failed to get operators state", "err", err)
		return nil, err
	}

	return operatorStakes, nil
}

func (r *AvsRegistryChainReader) GetOperatorsStakeInQuorumsOfOperatorAtBlock(
	opts *bind.CallOpts,
	operatorId types.OperatorId,
	blockNumber uint32,
) ([]types.QuorumNum, [][]blsoperatorstateretriever.BLSOperatorStateRetrieverOperator, error) {
	quorumBitmap, operatorStakes, err := r.blsOperatorStateRetriever.GetOperatorState0(
		opts,
		r.registryCoordinatorAddr,
		operatorId,
		blockNumber)
	if err != nil {
		r.logger.Error(
			"Failed to get operators state",
			"err",
			err,
			"fn",
			"AvsRegistryChainReader.GetOperatorsStakeInQuorumsOfOperatorAtBlock",
		)
		return nil, nil, err
	}
	quorums := types.BitmapToQuorumIds(quorumBitmap)
	return quorums, operatorStakes, nil
}

// opts will be modified to have the latest blockNumber, so make sure not to reuse it
// blockNumber in opts will be ignored, and the chain will be queried to get the latest blockNumber
func (r *AvsRegistryChainReader) GetOperatorsStakeInQuorumsOfOperatorAtCurrentBlock(
	opts *bind.CallOpts,
	operatorId types.OperatorId,
) ([]types.QuorumNum, [][]blsoperatorstateretriever.BLSOperatorStateRetrieverOperator, error) {
	curBlock, err := r.ethClient.BlockNumber(opts.Context)
	if err != nil {
		r.logger.Error("Failed to get current block number", "err", err)
		return nil, nil, err
	}
	if curBlock > math.MaxUint32 {
		r.logger.Error("Current block number is too large to be converted to uint32")
		return nil, nil, err
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
	quorumBitmap, err := r.registryCoordinator.GetCurrentQuorumBitmapByOperatorId(opts, operatorId)
	if err != nil {
		r.logger.Error("Failed to get operator quorums", "err", err)
		return nil, err
	}
	quorums := types.BitmapToQuorumIds(quorumBitmap)
	quorumStakes := make(map[types.QuorumNum]types.StakeAmount)
	for _, quorum := range quorums {
		stake, err := r.stakeRegistry.GetCurrentOperatorStakeForQuorum(
			&bind.CallOpts{},
			operatorId,
			quorum,
		)
		if err != nil {
			r.logger.Error("Failed to get operator stake", "err", err)
			return nil, err
		}
		quorumStakes[quorum] = stake
	}
	return quorumStakes, nil
}

func (r *AvsRegistryChainReader) GetCheckSignaturesIndices(
	opts *bind.CallOpts,
	referenceBlockNumber uint32,
	quorumNumbers []byte,
	nonSignerOperatorIds [][32]byte,
) (blsoperatorstateretriever.BLSOperatorStateRetrieverCheckSignaturesIndices, error) {
	checkSignatureIndices, err := r.blsOperatorStateRetriever.GetCheckSignaturesIndices(
		opts,
		r.registryCoordinatorAddr,
		referenceBlockNumber,
		quorumNumbers,
		nonSignerOperatorIds,
	)
	if err != nil {
		r.logger.Error("Failed to get check signatures indices", "err", err)
		return blsoperatorstateretriever.BLSOperatorStateRetrieverCheckSignaturesIndices{}, err
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
		r.logger.Error("Failed to get operator id", "err", err)
		return [32]byte{}, err
	}
	return operatorId, nil
}

func (r *AvsRegistryChainReader) IsOperatorRegistered(
	opts *bind.CallOpts,
	operatorAddress gethcommon.Address,
) (bool, error) {
	operatorId, err := r.registryCoordinator.GetOperatorId(opts, operatorAddress)
	if err != nil {
		r.logger.Error("Cannot get operator id", "err", err)
		return false, err
	}
	// OperatorId is set in contract during registration, so if it is not set, the operator is not registered
	registeredWithAvs := operatorId != [32]byte{}
	return registeredWithAvs, nil
}
