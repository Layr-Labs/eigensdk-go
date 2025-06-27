package challenger

import (
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/ethclient"
)

type ChallengerConfig struct {
	EthWsUrl       string
	Logger         logging.Logger
	TaskManagerAbi *abi.ABI
	EthClient      *ethclient.Client
}
