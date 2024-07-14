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

// NewChainSubscriber creates a new instance of ChainSubscriber
// The bindings must be created using websocket ETH Client
func NewChainSubscriber(
	regCoord regcoord.ContractRegistryCoordinatorFilters,
	blsApkRegistry blsapkreg.ContractBLSApkRegistryFilters,
	logger logging.Logger,
) *ChainSubscriber {
	logger = logger.With(logging.ComponentKey, "avsregistry/ChainSubscriber")

	return &ChainSubscriber{
		regCoord:       regCoord,
		blsApkRegistry: blsApkRegistry,
		logger:         logger,
	}
}

func NewChainSubscriberFromBindings(
	bindings *ContractBindings,
	logger logging.Logger,
) *ChainSubscriber {
	return NewChainSubscriber(bindings.RegistryCoordinator, bindings.BlsApkRegistry, logger)
}

// BuildAvsRegistryChainSubscriber creates a new instance of ChainSubscriber
// Deprecated: Use NewSubscriberFromConfig instead
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
	return NewChainSubscriber(regCoord, blsApkReg, logger), nil
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
