package avsecdsaregistry

import (
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethcommon "github.com/ethereum/go-ethereum/common"

	opstateretriever "github.com/Layr-Labs/eigensdk-go/contracts/bindings/ECDSAOperatorStateRetriever"
	regcoord "github.com/Layr-Labs/eigensdk-go/contracts/bindings/ECDSARegistryCoordinator"
)

type AvsEcdsaRegistryReader interface {
	GetOperatorsStakeInQuorumsAtBlock(
		opts *bind.CallOpts,
		quorumNumbers []byte,
		blockNumber uint32,
	) ([][]opstateretriever.ECDSAOperatorStateRetrieverOperator, error)

	GetCheckSignaturesIndices(
		opts *bind.CallOpts,
		referenceBlockNumber uint32,
		quorumNumbers []byte,
		nonSignerOperatorIds []gethcommon.Address,
	) (opstateretriever.ECDSAOperatorStateRetrieverCheckSignaturesIndices, error)

	IsOperatorRegistered(
		opts *bind.CallOpts,
		operatorAddress gethcommon.Address,
	) (bool, error)
}

type AvsEcdsaRegistryChainReader struct {
	logger                  logging.Logger
	registryCoordinatorAddr gethcommon.Address
	registryCoordinator     *regcoord.ContractECDSARegistryCoordinator
	operatorStateRetriever  *opstateretriever.ContractECDSAOperatorStateRetriever
	ethClient               eth.EthClient
}

// forces AvsReader to implement the clients.ReaderInterface interface
var _ AvsEcdsaRegistryReader = (*AvsEcdsaRegistryChainReader)(nil)

func NewAvsEcdsaRegistryChainReader(
	registryCoordinatorAddr gethcommon.Address,
	registerCoordinator *regcoord.ContractECDSARegistryCoordinator,
	operatorStateRetriever *opstateretriever.ContractECDSAOperatorStateRetriever,
	logger logging.Logger,
	ethClient eth.EthClient,
) *AvsEcdsaRegistryChainReader {
	return &AvsEcdsaRegistryChainReader{
		registryCoordinatorAddr: registryCoordinatorAddr,
		registryCoordinator:     registerCoordinator,
		operatorStateRetriever:  operatorStateRetriever,
		logger:                  logger,
		ethClient:               ethClient,
	}
}

func BuildAvsEcdsaRegistryChainReader(
	registryCoordinatorAddr gethcommon.Address,
	operatorStateRetrieverAddr gethcommon.Address,
	ethClient eth.EthClient,
	logger logging.Logger,
) (*AvsEcdsaRegistryChainReader, error) {
	contractOperatorStateRetriever, err := opstateretriever.NewContractECDSAOperatorStateRetriever(
		operatorStateRetrieverAddr,
		ethClient,
	)
	if err != nil {
		return nil, err
	}
	contractRegistryCoordinator, err := regcoord.NewContractECDSARegistryCoordinator(
		registryCoordinatorAddr,
		ethClient,
	)
	if err != nil {
		return nil, err
	}
	return NewAvsEcdsaRegistryChainReader(
		registryCoordinatorAddr,
		contractRegistryCoordinator,
		contractOperatorStateRetriever,
		logger,
		ethClient,
	), nil
}

func (r *AvsEcdsaRegistryChainReader) GetCheckSignaturesIndices(
	opts *bind.CallOpts,
	referenceBlockNumber uint32,
	quorumNumbers []byte,
	signerOperatorIds []gethcommon.Address,
) (opstateretriever.ECDSAOperatorStateRetrieverCheckSignaturesIndices, error) {
	checkSignatureIndices, err := r.operatorStateRetriever.GetCheckSignaturesIndices(
		opts,
		r.registryCoordinatorAddr,
		referenceBlockNumber,
		quorumNumbers,
		signerOperatorIds,
	)
	if err != nil {
		r.logger.Error("Failed to get check signatures indices", "err", err)
		return opstateretriever.ECDSAOperatorStateRetrieverCheckSignaturesIndices{}, err
	}
	return checkSignatureIndices, nil
}

// the contract stores historical state, so blockNumber should be the block number of the state you want to query
// and the blockNumber in opts should be the block number of the latest block (or set to nil, which is equivalent)
func (r *AvsEcdsaRegistryChainReader) GetOperatorsStakeInQuorumsAtBlock(
	opts *bind.CallOpts,
	quorumNumbers []byte,
	blockNumber uint32,
) ([][]opstateretriever.ECDSAOperatorStateRetrieverOperator, error) {
	operatorStakes, err := r.operatorStateRetriever.GetOperatorState0(
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

func (r *AvsEcdsaRegistryChainReader) IsOperatorRegistered(
	opts *bind.CallOpts,
	operatorAddress gethcommon.Address,
) (bool, error) {
	operatorStatus, err := r.registryCoordinator.GetOperatorStatus(opts, operatorAddress)
	if err != nil {
		r.logger.Error("Cannot get operator status", "err", err)
		return false, err
	}

	// 0 = NEVER_REGISTERED, 1 = REGISTERED, 2 = DEREGISTERED
	registeredWithAvs := operatorStatus == 1
	return registeredWithAvs, nil
}
