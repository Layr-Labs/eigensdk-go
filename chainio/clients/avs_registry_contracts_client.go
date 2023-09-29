package clients

import (
	"math/big"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	blsoperatorstateretrievar "github.com/Layr-Labs/eigensdk-go/contracts/bindings/BLSOperatorStateRetriever"
	blspubkeyregistry "github.com/Layr-Labs/eigensdk-go/contracts/bindings/BLSPubkeyRegistry"
	regcoord "github.com/Layr-Labs/eigensdk-go/contracts/bindings/BLSRegistryCoordinatorWithIndices"
	"github.com/Layr-Labs/eigensdk-go/logging"

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
		quorumNumbers []byte,
		blockNumber uint32,
	) ([][]blsoperatorstateretrievar.BLSOperatorStateRetrieverOperator, error)

	GetCheckSignaturesIndices(
		opts *bind.CallOpts,
		referenceBlockNumber uint32,
		quorumNumbers []byte,
		nonSignerOperatorIds [][32]byte,
	) (blsoperatorstateretrievar.BLSOperatorStateRetrieverCheckSignaturesIndices, error)

	UpdateStakes(
		opts *bind.TransactOpts,
		operators []gethcommon.Address,
		operatorIds [][32]byte,
		prevElements []*big.Int,
	) (*gethTypes.Transaction, error)

	DeregisterOperator(
		opts *bind.TransactOpts,
		operator gethcommon.Address,
		quorumNumbers []byte,
		pubkey blspubkeyregistry.BN254G1Point,
	) (*gethTypes.Transaction, error)
}

type AvsregistryContractsChainClient struct {
	avsRegistryBindings     *avsRegistryContractBindings
	ethHttpClient           eth.EthClient
	registryCoordinatorAddr gethcommon.Address
	logger                  logging.Logger
}

var _ AVSRegistryContractsClient = (*AvsregistryContractsChainClient)(nil)

func NewAVSContractsChainClient(
	blsRegistryCoordinatorAddr gethcommon.Address,
	blsOperatorStateRetrieverAddr gethcommon.Address,
	stakeregistryAddr gethcommon.Address,
	blsPubkeyRegistryAddr gethcommon.Address,
	ethClient eth.EthClient,
	logger logging.Logger,
) (*AvsregistryContractsChainClient, error) {
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

	return &AvsregistryContractsChainClient{
		avsRegistryBindings:     avsRegistryBindings,
		ethHttpClient:           ethClient,
		registryCoordinatorAddr: blsRegistryCoordinatorAddr,
		logger:                  logger,
	}, nil
}

// Registration specific functions
func (a *AvsregistryContractsChainClient) RegisterOperatorWithCoordinator(
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

func (a *AvsregistryContractsChainClient) GetOperatorId(
	opts *bind.CallOpts,
	operatorAddress gethcommon.Address,
) ([32]byte, error) {
	return a.avsRegistryBindings.RegistryCoordinator.GetOperatorId(opts, operatorAddress)
}

func (a *AvsregistryContractsChainClient) GetOperatorsStakeInQuorumsAtBlock(
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

func (a *AvsregistryContractsChainClient) GetCheckSignaturesIndices(
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

func (a *AvsregistryContractsChainClient) UpdateStakes(
	opts *bind.TransactOpts,
	operators []gethcommon.Address,
	operatorIds [][32]byte,
	prevElements []*big.Int,
) (*gethTypes.Transaction, error) {
	return a.avsRegistryBindings.StakeRegistry.UpdateStakes(
		opts,
		operators,
		operatorIds,
		prevElements)
}

func (a *AvsregistryContractsChainClient) DeregisterOperator(
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
