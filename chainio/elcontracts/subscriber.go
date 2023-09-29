package elcontracts

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/event"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients"
	pubkeycompendium "github.com/Layr-Labs/eigensdk-go/contracts/bindings/BLSPublicKeyCompendium"
	"github.com/Layr-Labs/eigensdk-go/logging"
)

type ELSubscriber interface {
	SubscribeToNewPubkeyRegistrations() (chan *pubkeycompendium.ContractBLSPublicKeyCompendiumNewPubkeyRegistration, event.Subscription, error)
}

type ELChainSubscriber struct {
	logger            logging.Logger
	elContractsClient clients.ELContractsClient
}

// forces EthSubscriber to implement the chainio.Subscriber interface
var _ ELSubscriber = (*ELChainSubscriber)(nil)

func NewELChainSubscriber(
	elContractsClient clients.ELContractsClient,
	logger logging.Logger,
) (*ELChainSubscriber, error) {

	return &ELChainSubscriber{
		logger:            logger,
		elContractsClient: elContractsClient,
	}, nil
}

func (s *ELChainSubscriber) SubscribeToNewPubkeyRegistrations() (chan *pubkeycompendium.ContractBLSPublicKeyCompendiumNewPubkeyRegistration, event.Subscription, error) {
	newPubkeyRegistrationChan := make(chan *pubkeycompendium.ContractBLSPublicKeyCompendiumNewPubkeyRegistration)
	sub, err := s.elContractsClient.WatchNewPubkeyRegistration(
		&bind.WatchOpts{}, newPubkeyRegistrationChan, nil,
	)
	if err != nil {
		s.logger.Error("Failed to subscribe to NewPubkeyRegistration events", "err", err)
		return nil, nil, err
	}
	return newPubkeyRegistrationChan, sub, nil
}
