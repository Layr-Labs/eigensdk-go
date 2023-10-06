package avsregistry

import (
	"context"
	"math"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	blsoperatorstateretrievar "github.com/Layr-Labs/eigensdk-go/contracts/bindings/BLSOperatorStateRetriever"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethcommon "github.com/ethereum/go-ethereum/common"
)

type AvsRegistryReader interface {
	GetOperatorsStakeInQuorumsAtBlock(
		ctx context.Context,
		quorumNumbers []byte,
		blockNumber uint32,
	) ([][]blsoperatorstateretrievar.BLSOperatorStateRetrieverOperator, error)

	GetOperatorsStakeInQuorumsOfOperatorAtBlock(
		ctx context.Context,
		operatorId types.OperatorId,
		blockNumber uint32,
	) ([]types.QuorumNum, [][]blsoperatorstateretrievar.BLSOperatorStateRetrieverOperator, error)

	GetOperatorsStakeInQuorumsOfOperatorAtCurrentBlock(
		ctx context.Context,
		operatorId types.OperatorId,
	) ([]types.QuorumNum, [][]blsoperatorstateretrievar.BLSOperatorStateRetrieverOperator, error)

	GetCheckSignaturesIndices(
		ctx context.Context,
		referenceBlockNumber uint32,
		quorumNumbers []byte,
		nonSignerOperatorIds [][32]byte,
	) (blsoperatorstateretrievar.BLSOperatorStateRetrieverCheckSignaturesIndices, error)

	GetOperatorId(ctx context.Context, operatorAddress gethcommon.Address) ([32]byte, error)

	IsOperatorRegistered(ctx context.Context, operatorAddress gethcommon.Address) (bool, error)
}

type AvsRegistryChainReader struct {
	logger                     logging.Logger
	avsRegistryContractsClient clients.AVSRegistryContractsClient
	ethClient                  eth.EthClient
}

// forces AvsReader to implement the clients.ReaderInterface interface
var _ AvsRegistryReader = (*AvsRegistryChainReader)(nil)

func NewAvsRegistryReader(
	avsRegistryContractsClient clients.AVSRegistryContractsClient,
	logger logging.Logger,
	ethClient eth.EthClient,
) (*AvsRegistryChainReader, error) {
	return &AvsRegistryChainReader{
		avsRegistryContractsClient: avsRegistryContractsClient,
		logger:                     logger,
		ethClient:                  ethClient,
	}, nil
}

func (r *AvsRegistryChainReader) GetOperatorsStakeInQuorumsAtBlock(
	ctx context.Context,
	quorumNumbers []byte,
	blockNumber uint32,
) ([][]blsoperatorstateretrievar.BLSOperatorStateRetrieverOperator, error) {
	operatorStakes, err := r.avsRegistryContractsClient.GetOperatorsStakeInQuorumsAtBlock(
		&bind.CallOpts{},
		quorumNumbers,
		blockNumber)
	if err != nil {
		r.logger.Error("Failed to get operators state", "err", err)
		return nil, err
	}

	return operatorStakes, nil
}

func (r *AvsRegistryChainReader) GetOperatorsStakeInQuorumsOfOperatorAtBlock(
	ctx context.Context,
	operatorId types.OperatorId,
	blockNumber uint32,
) ([]types.QuorumNum, [][]blsoperatorstateretrievar.BLSOperatorStateRetrieverOperator, error) {
	quorums, operatorStakes, err := r.avsRegistryContractsClient.GetOperatorsStakeInQuorumsOfOperatorAtBlock(
		&bind.CallOpts{},
		operatorId,
		blockNumber)
	if err != nil {
		r.logger.Error("Failed to get operators state", "err", err)
		return nil, nil, err
	}

	return quorums, operatorStakes, nil
}

func (r *AvsRegistryChainReader) GetOperatorsStakeInQuorumsOfOperatorAtCurrentBlock(
	ctx context.Context,
	operatorId types.OperatorId,
) ([]types.QuorumNum, [][]blsoperatorstateretrievar.BLSOperatorStateRetrieverOperator, error) {
	curBlock, err := r.ethClient.BlockNumber(ctx)
	if err != nil {
		r.logger.Error("Failed to get current block number", "err", err)
		return nil, nil, err
	}
	if curBlock > math.MaxUint32 {
		r.logger.Error("Current block number is too large to be converted to uint32")
		return nil, nil, err
	}
	return r.GetOperatorsStakeInQuorumsOfOperatorAtBlock(ctx, operatorId, uint32(curBlock))
}

func (r *AvsRegistryChainReader) GetCheckSignaturesIndices(
	ctx context.Context,
	referenceBlockNumber uint32,
	quorumNumbers []byte,
	nonSignerOperatorIds [][32]byte,
) (blsoperatorstateretrievar.BLSOperatorStateRetrieverCheckSignaturesIndices, error) {
	checkSignatureIndices, err := r.avsRegistryContractsClient.GetCheckSignaturesIndices(
		&bind.CallOpts{},
		referenceBlockNumber,
		quorumNumbers,
		nonSignerOperatorIds,
	)
	if err != nil {
		r.logger.Error("Failed to get check signatures indices", "err", err)
		return blsoperatorstateretrievar.BLSOperatorStateRetrieverCheckSignaturesIndices{}, err
	}
	return checkSignatureIndices, nil
}

func (r *AvsRegistryChainReader) GetOperatorId(
	ctx context.Context,
	operatorAddress gethcommon.Address,
) ([32]byte, error) {
	operatorId, err := r.avsRegistryContractsClient.GetOperatorId(
		&bind.CallOpts{},
		operatorAddress,
	)
	if err != nil {
		r.logger.Error("Failed to get operator id", "err", err)
		return [32]byte{}, err
	}
	return operatorId, nil
}

func (r *AvsRegistryChainReader) IsOperatorRegistered(
	ctx context.Context,
	operatorAddress gethcommon.Address,
) (bool, error) {
	operatorId, err := r.avsRegistryContractsClient.GetOperatorId(&bind.CallOpts{}, operatorAddress)
	if err != nil {
		r.logger.Error("Cannot get operator id", "err", err)
		return false, err
	}
	// OperatorId is set in contract during registration, so if it is not set, the operator is not registered
	registeredWithAvs := operatorId != [32]byte{}
	return registeredWithAvs, nil
}
