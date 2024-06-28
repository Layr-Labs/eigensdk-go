package avsregistry

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	blsapkreg "github.com/Layr-Labs/eigensdk-go/contracts/bindings/BLSApkRegistry"
	regcoord "github.com/Layr-Labs/eigensdk-go/contracts/bindings/RegistryCoordinator"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/utils"
)

type Subscriber interface {
	SubscribeToNewPubkeyRegistrations() (chan *blsapkreg.ContractBLSApkRegistryNewPubkeyRegistration, event.Subscription, error)
	SubscribeToOperatorSocketUpdates() (chan *regcoord.ContractRegistryCoordinatorOperatorSocketUpdate, event.Subscription, error)
}

type ChainSubscriber struct {
	logger         logging.Logger
	regCoord       regcoord.ContractRegistryCoordinatorFilters
	blsApkRegistry blsapkreg.ContractBLSApkRegistryFilters
}

// forces EthSubscriber to implement the chainio.Subscriber interface
var _ Subscriber = (*ChainSubscriber)(nil)

func NewChainSubscriber(
	logger logging.Logger,
	regCoord regcoord.ContractRegistryCoordinatorFilters,
	blsApkRegistry blsapkreg.ContractBLSApkRegistryFilters,
) (*ChainSubscriber, error) {
	logger = logger.With("module", "avsregistry/ChainSubscriber")

	return &ChainSubscriber{
		logger:         logger,
		regCoord:       regCoord,
		blsApkRegistry: blsApkRegistry,
	}, nil
}

func BuildAvsRegistryChainSubscriber(
	regCoordAddr common.Address,
	ethWsClient eth.Client,
	logger logging.Logger,
) (*ChainSubscriber, error) {
	regCoord, err := regcoord.NewContractRegistryCoordinator(regCoordAddr, ethWsClient)
	if err != nil {
		return nil, utils.WrapError("Failed to create RegistryCoordinator contract", err)
	}
	blsApkRegAddr, err := regCoord.BlsApkRegistry(&bind.CallOpts{})
	if err != nil {
		return nil, utils.WrapError("Failed to get BLSApkRegistry address from RegistryCoordinator", err)
	}
	blsApkReg, err := blsapkreg.NewContractBLSApkRegistry(blsApkRegAddr, ethWsClient)
	if err != nil {
		return nil, utils.WrapError("Failed to create BLSApkRegistry contract", err)
	}
	return NewChainSubscriber(logger, regCoord, blsApkReg)
}

func NewSubscriberFromConfig(
	cfg Config,
	ethClient eth.Client,
	logger logging.Logger,
) (*ChainSubscriber, error) {
	bindings, err := NewBindingsFromConfig(cfg, ethClient, logger)
	if err != nil {
		return nil, err
	}

	return NewChainSubscriber(logger, bindings.RegistryCoordinator, bindings.BlsApkRegistry)
}

func (s *ChainSubscriber) SubscribeToNewPubkeyRegistrations() (chan *blsapkreg.ContractBLSApkRegistryNewPubkeyRegistration, event.Subscription, error) {
	newPubkeyRegistrationChan := make(chan *blsapkreg.ContractBLSApkRegistryNewPubkeyRegistration)
	sub, err := s.blsApkRegistry.WatchNewPubkeyRegistration(
		&bind.WatchOpts{}, newPubkeyRegistrationChan, nil,
	)
	if err != nil {
		return nil, nil, utils.WrapError("Failed to subscribe to NewPubkeyRegistration events", err)
	}
	return newPubkeyRegistrationChan, sub, nil
}

func (s *ChainSubscriber) SubscribeToOperatorSocketUpdates() (chan *regcoord.ContractRegistryCoordinatorOperatorSocketUpdate, event.Subscription, error) {
	operatorSocketUpdateChan := make(chan *regcoord.ContractRegistryCoordinatorOperatorSocketUpdate)
	sub, err := s.regCoord.WatchOperatorSocketUpdate(
		&bind.WatchOpts{}, operatorSocketUpdateChan, nil,
	)
	if err != nil {
		return nil, nil, utils.WrapError("Failed to subscribe to OperatorSocketUpdate events", err)
	}
	return operatorSocketUpdateChan, sub, nil
}
