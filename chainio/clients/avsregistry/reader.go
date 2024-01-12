package avsregistry

import (
	"bytes"
	"context"
	"math"
	"math/big"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	gethcommon "github.com/ethereum/go-ethereum/common"

	contractOperatorStateRetriever "github.com/Layr-Labs/eigensdk-go/contracts/bindings/OperatorStateRetriever"
	opstateretriever "github.com/Layr-Labs/eigensdk-go/contracts/bindings/OperatorStateRetriever"
	regcoord "github.com/Layr-Labs/eigensdk-go/contracts/bindings/RegistryCoordinator"
	stakeregistry "github.com/Layr-Labs/eigensdk-go/contracts/bindings/StakeRegistry"
	eigenabi "github.com/Layr-Labs/eigensdk-go/types/abi"
)

type AvsRegistryReader interface {
	GetQuorumCount(opts *bind.CallOpts) (uint8, error)

	GetOperatorsStakeInQuorumsAtBlock(
		opts *bind.CallOpts,
		quorumNumbers []byte,
		blockNumber uint32,
	) ([][]opstateretriever.OperatorStateRetrieverOperator, error)

	GetOperatorAddrsInQuorumsAtCurrentBlock(
		opts *bind.CallOpts,
		quorumNumbers []byte,
	) ([][]common.Address, error)

	GetOperatorStakeInQuorumsOfOperatorAtCurrentBlock(
		opts *bind.CallOpts,
		operatorId types.OperatorId,
	) (map[types.QuorumNum]types.StakeAmount, error)

	GetOperatorsStakeInQuorumsOfOperatorAtBlock(
		opts *bind.CallOpts,
		operatorId types.OperatorId,
		blockNumber uint32,
	) ([]types.QuorumNum, [][]opstateretriever.OperatorStateRetrieverOperator, error)

	GetOperatorsStakeInQuorumsOfOperatorAtCurrentBlock(
		opts *bind.CallOpts,
		operatorId types.OperatorId,
	) ([]types.QuorumNum, [][]opstateretriever.OperatorStateRetrieverOperator, error)

	GetCheckSignaturesIndices(
		opts *bind.CallOpts,
		referenceBlockNumber uint32,
		quorumNumbers []byte,
		nonSignerOperatorIds [][32]byte,
	) (opstateretriever.OperatorStateRetrieverCheckSignaturesIndices, error)

	GetOperatorId(opts *bind.CallOpts, operatorAddress gethcommon.Address) ([32]byte, error)

	GetOperatorFromId(opts *bind.CallOpts, operatorId types.OperatorId) (gethcommon.Address, error)

	IsOperatorRegistered(opts *bind.CallOpts, operatorAddress gethcommon.Address) (bool, error)

	QueryExistingRegisteredOperatorPubKeys(
		ctx context.Context,
		startBlock *big.Int,
		stopBlock *big.Int,
	) ([]types.OperatorAddr, []types.OperatorPubkeys, error)
}

type AvsRegistryChainReader struct {
	logger                  logging.Logger
	blsApkRegistryAddr      gethcommon.Address
	registryCoordinatorAddr gethcommon.Address
	registryCoordinator     *regcoord.ContractRegistryCoordinator
	operatorStateRetriever  *opstateretriever.ContractOperatorStateRetriever
	stakeRegistry           *stakeregistry.ContractStakeRegistry
	ethClient               eth.EthClient
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
	ethClient eth.EthClient,
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
	ethClient eth.EthClient,
	logger logging.Logger,
) (*AvsRegistryChainReader, error) {
	contractRegistryCoordinator, err := regcoord.NewContractRegistryCoordinator(registryCoordinatorAddr, ethClient)
	if err != nil {
		return nil, err
	}
	blsApkRegistryAddr, err := contractRegistryCoordinator.BlsApkRegistry(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	stakeRegistryAddr, err := contractRegistryCoordinator.StakeRegistry(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	contractStakeRegistry, err := stakeregistry.NewContractStakeRegistry(stakeRegistryAddr, ethClient)
	if err != nil {
		return nil, err
	}
	contractOperatorStateRetriever, err := contractOperatorStateRetriever.NewContractOperatorStateRetriever(
		operatorStateRetrieverAddr,
		ethClient,
	)
	if err != nil {
		return nil, err
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

// the contract stores historical state, so blockNumber should be the block number of the state you want to query
// and the blockNumber in opts should be the block number of the latest block (or set to nil, which is equivalent)
func (r *AvsRegistryChainReader) GetOperatorsStakeInQuorumsAtBlock(
	opts *bind.CallOpts,
	quorumNumbers []byte,
	blockNumber uint32,
) ([][]opstateretriever.OperatorStateRetrieverOperator, error) {
	operatorStakes, err := r.operatorStateRetriever.GetOperatorState(
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

func (r *AvsRegistryChainReader) GetOperatorAddrsInQuorumsAtCurrentBlock(
	opts *bind.CallOpts,
	quorumNumbers []byte,
) ([][]common.Address, error) {
	if opts.Context == nil {
		opts.Context = context.Background()
	}
	curBlock, err := r.ethClient.BlockNumber(opts.Context)
	if err != nil {
		r.logger.Error("Failed to get current block number", "err", err)
		return nil, err
	}
	if curBlock > math.MaxUint32 {
		r.logger.Error("Current block number is too large to be converted to uint32")
		return nil, err
	}
	operatorStakes, err := r.operatorStateRetriever.GetOperatorState(
		opts,
		r.registryCoordinatorAddr,
		quorumNumbers,
		uint32(curBlock),
	)
	if err != nil {
		r.logger.Error("Failed to get operators state", "err", err)
		return nil, err
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
) ([]types.QuorumNum, [][]opstateretriever.OperatorStateRetrieverOperator, error) {
	quorumBitmap, operatorStakes, err := r.operatorStateRetriever.GetOperatorState0(
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
) ([]types.QuorumNum, [][]opstateretriever.OperatorStateRetrieverOperator, error) {
	if opts.Context == nil {
		opts.Context = context.Background()
	}
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
	quorumBitmap, err := r.registryCoordinator.GetCurrentQuorumBitmap(opts, operatorId)
	if err != nil {
		r.logger.Error("Failed to get operator quorums", "err", err)
		return nil, err
	}
	quorums := types.BitmapToQuorumIds(quorumBitmap)
	quorumStakes := make(map[types.QuorumNum]types.StakeAmount)
	for _, quorum := range quorums {
		stake, err := r.stakeRegistry.GetCurrentStake(
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
) (opstateretriever.OperatorStateRetrieverCheckSignaturesIndices, error) {
	checkSignatureIndices, err := r.operatorStateRetriever.GetCheckSignaturesIndices(
		opts,
		r.registryCoordinatorAddr,
		referenceBlockNumber,
		quorumNumbers,
		nonSignerOperatorIds,
	)
	if err != nil {
		r.logger.Error("Failed to get check signatures indices", "err", err)
		return opstateretriever.OperatorStateRetrieverCheckSignaturesIndices{}, err
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

func (r *AvsRegistryChainReader) GetOperatorFromId(
	opts *bind.CallOpts,
	operatorId types.OperatorId,
) (gethcommon.Address, error) {
	operatorAddress, err := r.registryCoordinator.GetOperatorFromId(
		opts,
		operatorId,
	)
	if err != nil {
		r.logger.Error("Failed to get operator address", "err", err)
		return gethcommon.Address{}, err
	}
	return operatorAddress, nil
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

func (r *AvsRegistryChainReader) QueryExistingRegisteredOperatorPubKeys(
	ctx context.Context,
	startBlock *big.Int,
	stopBlock *big.Int,
) ([]types.OperatorAddr, []types.OperatorPubkeys, error) {

	blsApkRegistryAbi, err := abi.JSON(bytes.NewReader(eigenabi.BLSApkRegistryAbi))
	if err != nil {
		r.logger.Error("Error getting Abi", "err", err)
		return nil, nil, err
	}

	query := ethereum.FilterQuery{
		FromBlock: startBlock,
		ToBlock:   stopBlock,
		Addresses: []gethcommon.Address{
			r.blsApkRegistryAddr,
		},
		Topics: [][]gethcommon.Hash{{blsApkRegistryAbi.Events["NewPubkeyRegistration"].ID}},
	}

	logs, err := r.ethClient.FilterLogs(ctx, query)
	if err != nil {
		r.logger.Error("Error filtering logs", "err", err)
		return nil, nil, err
	}
	r.logger.Info("logs:", "logs", logs)

	operatorAddresses := make([]types.OperatorAddr, 0)
	operatorPubkeys := make([]types.OperatorPubkeys, 0)

	for _, vLog := range logs {

		// get the operator address
		r.logger.Infof("deleteme", "vLog", vLog.Topics)
		operatorAddr := gethcommon.HexToAddress(vLog.Topics[1].Hex())
		operatorAddresses = append(operatorAddresses, operatorAddr)

		event, err := blsApkRegistryAbi.Unpack("NewPubkeyRegistration", vLog.Data)
		if err != nil {
			r.logger.Error("Error unpacking event data", "err", err)
			return nil, nil, err
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

	return operatorAddresses, operatorPubkeys, nil
}
