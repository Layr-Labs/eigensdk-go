package elcontracts

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/event"

	pubkeycompendium "github.com/Layr-Labs/eigensdk-go-master/contracts/bindings/BLSPublicKeyCompendium"
	"github.com/Layr-Labs/eigensdk-go-master/logging"
)

type ELSubscriber interface {
	SubscribeToNewPubkeyRegistrations() (chan *pubkeycompendium.ContractBLSPublicKeyCompendiumNewPubkeyRegistration, event.Subscription, error)
}

type ELChainSubscriber struct {
	logger              logging.Logger
	blsPubkeyCompendium pubkeycompendium.ContractBLSPublicKeyCompendiumFilters
}

// forces EthSubscriber to implement the chainio.Subscriber interface
var _ ELSubscriber = (*ELChainSubscriber)(nil)

func NewELChainSubscriber(
	blsPubkeyCompendium pubkeycompendium.ContractBLSPublicKeyCompendiumFilters,
	logger logging.Logger,
) (*ELChainSubscriber, error) {

	return &ELChainSubscriber{
		logger:              logger,
		blsPubkeyCompendium: blsPubkeyCompendium,
	}, nil
}

func (s *ELChainSubscriber) SubscribeToNewPubkeyRegistrations() (chan *pubkeycompendium.ContractBLSPublicKeyCompendiumNewPubkeyRegistration, event.Subscription, error) {
	newPubkeyRegistrationChan := make(chan *pubkeycompendium.ContractBLSPublicKeyCompendiumNewPubkeyRegistration)
	sub, err := s.blsPubkeyCompendium.WatchNewPubkeyRegistration(
		&bind.WatchOpts{}, newPubkeyRegistrationChan, nil,
	)
	if err != nil {
		s.logger.Error("Failed to subscribe to NewPubkeyRegistration events", "err", err)
		return nil, nil, err
	}
	return newPubkeyRegistrationChan, sub, nil
}
