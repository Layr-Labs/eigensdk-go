package aggregator

import (
	"crypto/ecdsa"

	"github.com/Layr-Labs/eigensdk-go/logging"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type AggregatorConfig struct {
	EthHttpUrl string
	EthWsUrl   string
	AggregatorServerIpPortAddr string

	RegistryCoordinatorAddress    common.Address
	OperatorStateRetrieverAddress common.Address
	ServiceManagerAddress         common.Address

	EthHttpClient *ethclient.Client
	Logger        logging.Logger
	EcdsaPrivateKey *ecdsa.PrivateKey

	TaskResponseHashFn sdktypes.TaskResponseHashFunction

	TaskManagerAbi *abi.ABI
}
