package avsregistry

import (
	"errors"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	blsapkreg "github.com/Layr-Labs/eigensdk-go/contracts/bindings/BLSApkRegistry"
	"github.com/Layr-Labs/eigensdk-go/logging"
)

type AvsRegistrySubscriber interface {
	SubscribeToNewPubkeyRegistrations() (chan *blsapkreg.ContractBLSApkRegistryNewPubkeyRegistration, event.Subscription, error)
}

type AvsRegistryChainSubscriber struct {
	logger         logging.Logger
	blsApkRegistry blsapkreg.ContractBLSApkRegistryFilters
}

// forces EthSubscriber to implement the chainio.Subscriber interface
var _ AvsRegistrySubscriber = (*AvsRegistryChainSubscriber)(nil)

func NewAvsRegistryChainSubscriber(
	blsApkRegistry blsapkreg.ContractBLSApkRegistryFilters,
	logger logging.Logger,
) (*AvsRegistryChainSubscriber, error) {
	return &AvsRegistryChainSubscriber{
		logger:         logger,
		blsApkRegistry: blsApkRegistry,
	}, nil
}

func BuildAvsRegistryChainSubscriber(
	blsApkRegistryAddr common.Address,
	ethWsClient eth.EthClient,
	logger logging.Logger,
) (*AvsRegistryChainSubscriber, error) {
	blsapkreg, err := blsapkreg.NewContractBLSApkRegistry(blsApkRegistryAddr, ethWsClient)
	if err != nil {
		return nil, errors.Join(errors.New("Failed to create BLSApkRegistry contract"), err)
	}
	return NewAvsRegistryChainSubscriber(blsapkreg, logger)
}

func (s *AvsRegistryChainSubscriber) SubscribeToNewPubkeyRegistrations() (chan *blsapkreg.ContractBLSApkRegistryNewPubkeyRegistration, event.Subscription, error) {
	newPubkeyRegistrationChan := make(chan *blsapkreg.ContractBLSApkRegistryNewPubkeyRegistration)
	sub, err := s.blsApkRegistry.WatchNewPubkeyRegistration(
		&bind.WatchOpts{}, newPubkeyRegistrationChan, nil,
	)
	if err != nil {
		return nil, nil, errors.Join(errors.New("Failed to subscribe to NewPubkeyRegistration events"), err)
	}
	return newPubkeyRegistrationChan, sub, nil
}
