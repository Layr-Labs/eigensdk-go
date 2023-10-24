package clients

import (
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	blsoperatorstateretrievar "github.com/Layr-Labs/eigensdk-go/contracts/bindings/BLSOperatorStateRetriever"
	blspubkeyregistry "github.com/Layr-Labs/eigensdk-go/contracts/bindings/BLSPubkeyRegistry"
	regcoord "github.com/Layr-Labs/eigensdk-go/contracts/bindings/BLSRegistryCoordinatorWithIndices"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/types"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethcommon "github.com/ethereum/go-ethereum/common"
	gethTypes "github.com/ethereum/go-ethereum/core/types"
)

type AVSRegistryContractsClient interface {
	RegisterOperatorWithCoordinator(
		opts *bind.TransactOpts,
		quorumNumbers []byte,
		pubkey regcoord.BN254G1Point,
		socket string,
	) (*gethTypes.Transaction, error)

	GetOperatorId(
		opts *bind.CallOpts,
		operatorAddress gethcommon.Address,
	) ([32]byte, error)

	GetOperatorsStakeInQuorumsAtBlock(
		opts *bind.CallOpts,
		quorumNumbers []types.QuorumNum,
		blockNumber uint32,
	) ([][]blsoperatorstateretrievar.BLSOperatorStateRetrieverOperator, error)

	GetOperatorsStakeInQuorumsOfOperatorAtBlock(
		opts *bind.CallOpts,
		operatorId types.OperatorId,
		blockNumber uint32,
	) ([]types.QuorumNum, [][]blsoperatorstateretrievar.BLSOperatorStateRetrieverOperator, error)

	GetCheckSignaturesIndices(
		opts *bind.CallOpts,
		referenceBlockNumber uint32,
		quorumNumbers []byte,
		nonSignerOperatorIds [][32]byte,
	) (blsoperatorstateretrievar.BLSOperatorStateRetrieverCheckSignaturesIndices, error)

	UpdateStakes(
		opts *bind.TransactOpts,
		operators []gethcommon.Address,
	) (*gethTypes.Transaction, error)

	DeregisterOperator(
		opts *bind.TransactOpts,
		operator gethcommon.Address,
		quorumNumbers []byte,
		pubkey blspubkeyregistry.BN254G1Point,
	) (*gethTypes.Transaction, error)
}

type AvsRegistryContractsChainClient struct {
	avsRegistryBindings     *avsRegistryContractBindings
	ethHttpClient           eth.EthClient
	registryCoordinatorAddr gethcommon.Address
	logger                  logging.Logger
}

var _ AVSRegistryContractsClient = (*AvsRegistryContractsChainClient)(nil)

func NewAvsRegistryContractsChainClient(
	blsRegistryCoordinatorAddr gethcommon.Address,
	blsOperatorStateRetrieverAddr gethcommon.Address,
	stakeregistryAddr gethcommon.Address,
	blsPubkeyRegistryAddr gethcommon.Address,
	ethClient eth.EthClient,
	logger logging.Logger,
) (*AvsRegistryContractsChainClient, error) {
	avsRegistryBindings, err := newAVSRegistryContractBindings(
		blsRegistryCoordinatorAddr,
		blsOperatorStateRetrieverAddr,
		stakeregistryAddr,
		blsPubkeyRegistryAddr,
		ethClient,
		logger)
	if err != nil {
		return nil, err
	}

	return &AvsRegistryContractsChainClient{
		avsRegistryBindings:     avsRegistryBindings,
		ethHttpClient:           ethClient,
		registryCoordinatorAddr: blsRegistryCoordinatorAddr,
		logger:                  logger,
	}, nil
}

// Registration specific functions
func (a *AvsRegistryContractsChainClient) RegisterOperatorWithCoordinator(
	opts *bind.TransactOpts,
	quorumNumbers []byte,
	pubkey regcoord.BN254G1Point,
	socket string,
) (*gethTypes.Transaction, error) {
	return a.avsRegistryBindings.RegistryCoordinator.RegisterOperatorWithCoordinator1(
		opts,
		quorumNumbers,
		pubkey,
		socket,
	)
}

func (a *AvsRegistryContractsChainClient) GetOperatorId(
	opts *bind.CallOpts,
	operatorAddress gethcommon.Address,
) ([32]byte, error) {
	return a.avsRegistryBindings.RegistryCoordinator.GetOperatorId(opts, operatorAddress)
}

func (a *AvsRegistryContractsChainClient) GetOperatorsStakeInQuorumsAtBlock(
	opts *bind.CallOpts,
	quorumNumbers []byte,
	blockNumber uint32,
) ([][]blsoperatorstateretrievar.BLSOperatorStateRetrieverOperator, error) {
	return a.avsRegistryBindings.BlsOperatorStateRetriever.GetOperatorState(
		opts,
		a.registryCoordinatorAddr,
		quorumNumbers,
		blockNumber)
}

func (a *AvsRegistryContractsChainClient) GetOperatorsStakeInQuorumsOfOperatorAtBlock(
	opts *bind.CallOpts,
	operatorId types.OperatorId,
	blockNumber uint32,
) ([]types.QuorumNum, [][]blsoperatorstateretrievar.BLSOperatorStateRetrieverOperator, error) {
	quorumBitmap, blsOperatorStateRetrieverOperator, err := a.avsRegistryBindings.BlsOperatorStateRetriever.GetOperatorState0(
		opts,
		a.registryCoordinatorAddr,
		operatorId,
		blockNumber,
	)
	if err != nil {
		return nil, nil, err
	}
	quorums := types.BitmapToQuorumIds(quorumBitmap)
	return quorums, blsOperatorStateRetrieverOperator, nil
}

func (a *AvsRegistryContractsChainClient) GetCheckSignaturesIndices(
	opts *bind.CallOpts,
	referenceBlockNumber uint32,
	quorumNumbers []byte,
	nonSignerOperatorIds [][32]byte,
) (blsoperatorstateretrievar.BLSOperatorStateRetrieverCheckSignaturesIndices, error) {
	return a.avsRegistryBindings.BlsOperatorStateRetriever.GetCheckSignaturesIndices(
		opts,
		a.registryCoordinatorAddr,
		referenceBlockNumber,
		quorumNumbers,
		nonSignerOperatorIds)
}

func (a *AvsRegistryContractsChainClient) UpdateStakes(
	opts *bind.TransactOpts,
	operators []gethcommon.Address,
) (*gethTypes.Transaction, error) {
	return a.avsRegistryBindings.StakeRegistry.UpdateStakes(
		opts,
		operators,
	)
}

func (a *AvsRegistryContractsChainClient) DeregisterOperator(
	opts *bind.TransactOpts,
	operator gethcommon.Address,
	quorumNumbers []byte,
	pubkey blspubkeyregistry.BN254G1Point,
) (*gethTypes.Transaction, error) {
	return a.avsRegistryBindings.BlsPubkeyRegistry.DeregisterOperator(
		opts,
		operator,
		quorumNumbers,
		pubkey)
}
