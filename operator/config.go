package operator

import (
	"math/big"

	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

// Operator configuration struct
type Config struct {
	// The address for the operator
	OperatorAddress string

	// The registry coordinator address is used to create the AVS reader
	AVSRegistryCoordinatorAddress string

	// Ethereum HTTP RPC URL to use for interacting with on-chain contracts
	EthRpcUrl string

	// Ethereum WebSocket RPC URL to use for subscribing to on-chain events
	EthWsUrl string

	// The path to the location of the bls private key on local storage
	BlsPrivateKeyStorePath string

	// IP address and port where the aggregator will listen to operator task responses
	AggregatorServerIpPortAddress string

	// The logger to use when logging
	Logger logging.Logger

	// The ABI of the task manager contract
	TaskManagerAbi *abi.ABI

	// The config used to register an operator
	RegistrationCfg RegistrationConfig
}

// This config is used to register an operator on startup.
// A TODO of this config is to make some values optional to perform some registration operations instead of all
type RegistrationConfig struct {
	// If set true, should register the operator on startup
	RegisterOnStartup bool

	// The address for the operator
	OperatorAddr common.Address

	// Used for setting the allogation delay to zero and initialize allocations
	AllocationManagerAddr common.Address

	// Used to register operator in operator sets and initialize allocations
	AvsAddress              common.Address
	RegistryCoordinatorAddr common.Address

	// Used to deposit into these strategies for operator and initialize allocations on these strategies
	StrategyAddrs []common.Address

	// Used to create eigenlayer chain reader and writer
	DelegationManagerAddress    common.Address
	RewardsCoordinatorAddress   common.Address
	PermissionControllerAddress common.Address

	// The url exposed by the anvil node to call the contract methods
	EthRpcUrl string

	// The path to the location of the bls and ecdsa private keys on local storage
	EcdsaKeyStorePath string
	BlsKeyStorePath   string

	// The ammount to mint to the operator
	AmountToMint *big.Int

	// The magnitudes to be allocatable (slashable) in the strategies for the operator
	AllocatableMagnitudes []uint64

	// The IDs of the operator sets to be registered
	OperatorSetIds []uint32
}
