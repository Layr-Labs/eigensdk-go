package types

import "github.com/ethereum/go-ethereum/common"

type ElChainReaderConfig struct {
	DelegationManagerAddress common.Address
	AvsDirectoryAddress      common.Address
}
