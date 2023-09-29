package avsregistry

import (
	"context"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	blsoperatorstateretrievar "github.com/Layr-Labs/eigensdk-go/contracts/bindings/BLSOperatorStateRetriever"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethcommon "github.com/ethereum/go-ethereum/common"
)

type AvsRegistryReader interface {
	GetOperatorsStakeInQuorumsAtBlock(
		ctx context.Context,
		quorumNumbers []byte,
		blockNumber uint32,
	) ([][]blsoperatorstateretrievar.BLSOperatorStateRetrieverOperator, error)

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
		r.logger.Error("Failed to get operator state", "err", err)
		return nil, err
	}

	return operatorStakes, nil
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
