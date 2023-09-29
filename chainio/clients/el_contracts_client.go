package clients

import (
	"math/big"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	blspubkeycompendium "github.com/Layr-Labs/eigensdk-go/contracts/bindings/BLSPublicKeyCompendium"
	contractDelegationManager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/DelegationManager"
	strategy "github.com/Layr-Labs/eigensdk-go/contracts/bindings/IStrategy"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	gethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

type ELContractsClient interface {

	// Write methods

	RegisterAsOperator(
		opts *bind.TransactOpts,
		registeringOperatorDetails contractDelegationManager.IDelegationManagerOperatorDetails,
		metadataURI string,
	) (*gethTypes.Transaction, error)

	ModifyOperatorDetails(
		opts *bind.TransactOpts,
		registeringOperatorDetails contractDelegationManager.IDelegationManagerOperatorDetails,
	) (*gethTypes.Transaction, error)

	UpdateOperatorMetadataURI(
		opts *bind.TransactOpts,
		metadataURI string,
	) (*gethTypes.Transaction, error)

	OptIntoSlashing(
		opts *bind.TransactOpts,
		avsServiceManagerAddr common.Address,
	) (*gethTypes.Transaction, error)

	RegisterBLSPublicKey(
		opts *bind.TransactOpts,
		signedMessageHash blspubkeycompendium.BN254G1Point,
		pubkeyG1 blspubkeycompendium.BN254G1Point,
		pubkeyG2 blspubkeycompendium.BN254G2Point,
	) (*gethTypes.Transaction, error)

	DepositIntoStrategy(
		opts *bind.TransactOpts,
		strategy common.Address,
		token common.Address,
		amount *big.Int,
	) (*gethTypes.Transaction, error)

	// Subscribe methods

	WatchNewPubkeyRegistration(
		opts *bind.WatchOpts,
		sink chan<- *blspubkeycompendium.ContractBLSPublicKeyCompendiumNewPubkeyRegistration,
		operator []common.Address,
	) (event.Subscription, error)

	// Read methods

	GetStrategyManagerContractAddress() (common.Address, error)

	GetBLSPublicKeyCompendiumContractAddress() common.Address

	IsOperator(opts *bind.CallOpts, operator common.Address) (bool, error)

	OperatorDetails(
		opts *bind.CallOpts,
		operator common.Address,
	) (contractDelegationManager.IDelegationManagerOperatorDetails, error)

	ContractCanSlashOperatorUntilBlock(
		opts *bind.CallOpts,
		operator common.Address,
		serviceContract common.Address,
	) (uint32, error)

	IsFrozen(opts *bind.CallOpts, staker common.Address) (bool, error)

	GetStrategyAndUnderlyingToken(
		strategyAddr common.Address,
	) (*strategy.ContractIStrategy, common.Address, error)

	GetStrategyAndUnderlyingERC20Token(
		strategyAddr common.Address,
	) (*strategy.ContractIStrategy, ERC20ContractClient, common.Address, error)

	GetOperatorPubkeyHash(
		opts *bind.CallOpts,
		operatorAddr common.Address,
	) ([32]byte, error)

	GetOperatorAddressFromPubkeyHash(opts *bind.CallOpts, pubkeyHash [32]byte) (common.Address, error)

	OperatorShares(opts *bind.CallOpts, operatorAddr common.Address, strategyAddr common.Address) (*big.Int, error)
}

// ElContractsChainClient is really just a wrapper around the go bindings that has a proper interface, so that we can
// mock it in tests.
// TODO(samlaf): should we instead just make the
type ElContractsChainClient struct {
	elHttpBindings *eigenlayerContractBindings
	// TODO(samlaf): currently we're creating a second set of bindings backed by a websocket client, so that we can
	// subscribe to events. perhaps a better way would be to make our EthClient implementation have both an http and
	// websocket client, and use the websocket client for
	// the SubscribeFilterLogs() method, and the http client for everything else...
	// Yet another option is to start using an indexer (maybe
	// https://github.com/ethereum-optimism/optimism/tree/develop/indexer) and drop websocket subscriptions altogether.
	elWsBindings            *eigenlayerContractBindings
	ethHttpClient           eth.EthClient
	logger                  logging.Logger
	blsPubKeyCompendiumAddr common.Address
}

var _ ELContractsClient = (*ElContractsChainClient)(nil)

func NewELContractsChainClient(
	slasherAddr common.Address,
	blsPubKeyCompendiumAddr common.Address,
	ethHttpClient eth.EthClient,
	ethWsClient eth.EthClient,
	logger logging.Logger,
) (*ElContractsChainClient, error) {
	elHttpBindings, err := newEigenlayerContractBindings(slasherAddr, blsPubKeyCompendiumAddr, ethHttpClient, logger)
	if err != nil {
		return nil, err
	}
	elWsBindings, err := newEigenlayerContractBindings(slasherAddr, blsPubKeyCompendiumAddr, ethWsClient, logger)
	if err != nil {
		return nil, err
	}

	return &ElContractsChainClient{
		elHttpBindings:          elHttpBindings,
		elWsBindings:            elWsBindings,
		ethHttpClient:           ethHttpClient,
		logger:                  logger,
		blsPubKeyCompendiumAddr: blsPubKeyCompendiumAddr,
	}, nil
}

func (e *ElContractsChainClient) RegisterAsOperator(
	opts *bind.TransactOpts,
	registeringOperatorDetails contractDelegationManager.IDelegationManagerOperatorDetails,
	metadataURI string,
) (*gethTypes.Transaction, error) {
	return e.elHttpBindings.DelegationManager.RegisterAsOperator(opts, registeringOperatorDetails, metadataURI)
}

func (e *ElContractsChainClient) ModifyOperatorDetails(
	opts *bind.TransactOpts,
	registeringOperatorDetails contractDelegationManager.IDelegationManagerOperatorDetails,
) (*gethTypes.Transaction, error) {
	return e.elHttpBindings.DelegationManager.ModifyOperatorDetails(opts, registeringOperatorDetails)
}

func (e *ElContractsChainClient) UpdateOperatorMetadataURI(
	opts *bind.TransactOpts,
	metadataURI string,
) (*gethTypes.Transaction, error) {
	return e.elHttpBindings.DelegationManager.UpdateOperatorMetadataURI(opts, metadataURI)
}

func (e *ElContractsChainClient) OptIntoSlashing(
	opts *bind.TransactOpts,
	avsServiceManagerAddr common.Address,
) (*gethTypes.Transaction, error) {
	return e.elHttpBindings.Slasher.OptIntoSlashing(opts, avsServiceManagerAddr)
}

func (e *ElContractsChainClient) RegisterBLSPublicKey(
	opts *bind.TransactOpts,
	signedMessageHash blspubkeycompendium.BN254G1Point,
	pubkeyG1 blspubkeycompendium.BN254G1Point,
	pubkeyG2 blspubkeycompendium.BN254G2Point,
) (*gethTypes.Transaction, error) {
	return e.elHttpBindings.BlsPubKeyCompendium.RegisterBLSPublicKey(opts, signedMessageHash, pubkeyG1, pubkeyG2)
}

func (e *ElContractsChainClient) GetStrategyAndUnderlyingToken(
	strategyAddr common.Address,
) (*strategy.ContractIStrategy, common.Address, error) {
	contractStrategy, err := strategy.NewContractIStrategy(strategyAddr, e.ethHttpClient)
	if err != nil {
		e.logger.Error("Failed to fetch strategy contract", "err", err)
		return nil, common.Address{}, err
	}
	underlyingTokenAddr, err := contractStrategy.UnderlyingToken(&bind.CallOpts{})
	if err != nil {
		e.logger.Error("Failed to fetch token contract", "err", err)
		return nil, common.Address{}, err
	}
	return contractStrategy, underlyingTokenAddr, nil
}

func (e *ElContractsChainClient) GetStrategyAndUnderlyingERC20Token(
	strategyAddr common.Address,
) (*strategy.ContractIStrategy, ERC20ContractClient, common.Address, error) {
	contractStrategy, err := strategy.NewContractIStrategy(strategyAddr, e.ethHttpClient)
	if err != nil {
		e.logger.Error("Failed to fetch strategy contract", "err", err)
		return nil, nil, common.Address{}, err
	}
	underlyingTokenAddr, err := contractStrategy.UnderlyingToken(&bind.CallOpts{})
	if err != nil {
		e.logger.Error("Failed to fetch token contract", "err", err)
		return nil, nil, common.Address{}, err
	}
	contractUnderlyingToken, err := NewERC20ContractChainClient(underlyingTokenAddr, e.ethHttpClient)
	if err != nil {
		e.logger.Error("Failed to fetch token contract", "err", err)
		return nil, nil, common.Address{}, err
	}
	return contractStrategy, contractUnderlyingToken, underlyingTokenAddr, nil
}

func (e *ElContractsChainClient) DepositIntoStrategy(
	opts *bind.TransactOpts,
	strategy common.Address,
	token common.Address,
	amount *big.Int,
) (*gethTypes.Transaction, error) {
	return e.elHttpBindings.StrategyManager.DepositIntoStrategy(opts, strategy, token, amount)
}

func (e *ElContractsChainClient) GetStrategyManagerContractAddress() (common.Address, error) {
	return e.elHttpBindings.Slasher.StrategyManager(&bind.CallOpts{})
}

func (e *ElContractsChainClient) WatchNewPubkeyRegistration(
	opts *bind.WatchOpts,
	sink chan<- *blspubkeycompendium.ContractBLSPublicKeyCompendiumNewPubkeyRegistration,
	operator []common.Address,
) (event.Subscription, error) {
	return e.elWsBindings.BlsPubKeyCompendium.WatchNewPubkeyRegistration(opts, sink, operator)
}

func (e *ElContractsChainClient) IsOperator(opts *bind.CallOpts, operator common.Address) (bool, error) {
	return e.elHttpBindings.DelegationManager.IsOperator(opts, operator)
}

func (e *ElContractsChainClient) OperatorDetails(
	opts *bind.CallOpts,
	operator common.Address,
) (contractDelegationManager.IDelegationManagerOperatorDetails, error) {
	return e.elHttpBindings.DelegationManager.OperatorDetails(opts, operator)
}

func (e *ElContractsChainClient) ContractCanSlashOperatorUntilBlock(
	opts *bind.CallOpts,
	operator common.Address,
	serviceContract common.Address,
) (uint32, error) {
	return e.elHttpBindings.Slasher.ContractCanSlashOperatorUntilBlock(opts, operator, serviceContract)
}

func (e *ElContractsChainClient) IsFrozen(opts *bind.CallOpts, staker common.Address) (bool, error) {
	return e.elHttpBindings.Slasher.IsFrozen(opts, staker)
}

func (e *ElContractsChainClient) GetBLSPublicKeyCompendiumContractAddress() common.Address {
	return e.blsPubKeyCompendiumAddr
}

func (e *ElContractsChainClient) GetOperatorPubkeyHash(
	opts *bind.CallOpts,
	operatorAddr common.Address,
) ([32]byte, error) {
	return e.elHttpBindings.BlsPubKeyCompendium.OperatorToPubkeyHash(opts, operatorAddr)
}

func (e *ElContractsChainClient) GetOperatorAddressFromPubkeyHash(
	opts *bind.CallOpts,
	pubkeyHash [32]byte,
) (common.Address, error) {
	return e.elHttpBindings.BlsPubKeyCompendium.PubkeyHashToOperator(opts, pubkeyHash)
}

func (e *ElContractsChainClient) OperatorShares(
	opts *bind.CallOpts,
	operatorAddr common.Address,
	strategyAddr common.Address,
) (*big.Int, error) {
	operatorSharesInStrategy, err := e.elHttpBindings.DelegationManager.OperatorShares(opts, operatorAddr, strategyAddr)
	if err != nil {
		return nil, err
	}
	return operatorSharesInStrategy, nil
}
