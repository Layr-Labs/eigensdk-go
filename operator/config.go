package operator

import (
	"math/big"

	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

type Config struct {
	OperatorAddress string

	// Avs Reader addresses
	OperatorStateRetrieverAddress string
	ServiceManagerAddress         string
	AVSRegistryCoordinatorAddress string

	EthRpcUrl string
	EthWsUrl  string

	BlsPrivateKeyStorePath        string
	AggregatorServerIpPortAddress string

	RegisterOnStartup bool

	Logger         logging.Logger
	TaskManagerAbi *abi.ABI

	RegistrationCfg RegistrationConfig
}

type RegistrationConfig struct {
	Logger logging.Logger

	OperatorAddr            common.Address
	AllocationManagerAddr   common.Address
	AvsAddress              common.Address
	RegistryCoordinatorAddr common.Address
	StrategyAddr            common.Address

	DelegationManagerAddress    common.Address
	RewardsCoordinatorAddress   common.Address
	PermissionControllerAddress common.Address

	EthHttpUrl string

	EcdsaKeyStorePath string
	BlsKeyStorePath   string

	AmountToMint         *big.Int
	AllocatableMagnitude uint64

	OperatorSetId uint32
}
