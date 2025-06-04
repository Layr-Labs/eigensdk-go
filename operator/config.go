package operator

import (
	"math/big"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/ethereum/go-ethereum/common"
)

// Operator configuration struct
type Config struct {
	// The address for the operator
	OperatorAddress string `toml:"operator_address"`

	// The registry coordinator address is used to create the AVS reader
	RegistryCoordinatorAddress string `toml:"registry_coordinator_address"`

	// Ethereum HTTP RPC URL to use for interacting with on-chain contracts
	EthRpcUrl string `toml:"eth_http_url"`

	// Ethereum WebSocket RPC URL to use for subscribing to on-chain events
	EthWsUrl string `toml:"eth_ws_url"`

	// The config used to create the BLS signer for the operator
	BlsSignerCfg BlsSignerConfig `toml:"bls_signer"`

	// IP address and port where the aggregator will listen to operator task responses
	AggregatorServerIpPortAddress string `toml:"aggregator_server_ip_port"`

	// The config used to register an operator
	Registration RegistrationConfig `toml:"registration"`
}

// For the BLS signing we accept two different options:
// - If the key pair is provided, we use that pair
// - If not, we use the specified keystore path and the password.
// We use the go-ethereum keystore package to encrypt/decrypt the private keys, more specifically the V3 encryption,
// so to see the supported or unsuported keystore formats refer to the geth keystore package documentation.
type BlsSignerConfig struct {
	// Path to the BLS keystore on local storage
	KeystorePath string `toml:"keystore_path"`

	// Password for the keystore
	KeystorePassword string `toml:"keystore_password"`

	// The bls Key Pair, if nil will use the store path and password above
	// Note: the toml tag is private_key because from the private key generates the key pair
	BlsKeyPair *bls.KeyPair `toml:"private_key"`
}

// This config is used to register an operator on startup.
// A TODO of this config is to make some values optional to perform some registration operations instead of all
type RegistrationConfig struct {
	// If set true, should register the operator on startup
	RegisterOnStartup bool `toml:"register_on_startup"`

	// Used for setting the allogation delay to zero and initialize allocations
	AllocationManagerAddr common.Address `toml:"allocation_manager_address"`

	// Used to register operator in operator sets and initialize allocations
	AvsAddress common.Address `toml:"service_manager_address"`

	// Used to deposit into these strategies for operator and initialize allocations on these strategies
	StrategyAddrs []common.Address `toml:"strategy_addresses"`

	// Used to create eigenlayer chain reader and writer
	DelegationManagerAddress    common.Address `toml:"delegation_manager_address"`
	RewardsCoordinatorAddress   common.Address `toml:"rewards_coordinator_address"`
	PermissionControllerAddress common.Address `toml:"permission_controller_address"`

	// The path to the location of the ecdsa private key on local storage
	EcdsaKeyStorePath string `toml:"ecdsa_private_key_store_path"`

	// The ammount to mint to the operator
	AmountToMint *big.Int `toml:"amount_to_mint"`

	// The magnitudes to be allocatable (slashable) in the strategies for the operator
	AllocatableMagnitudes []uint64 `toml:"allocatable_magnitudes"`

	// The IDs of the operator sets to be registered
	OperatorSetIds []uint32 `toml:"operator_set_ids"`

	// The operator metadata url to be set in eigenlayer registration
	MetadataUrl string `toml:"metadata_url"`

	// The socket used in the operator set registration request
	Socket string `toml:"socket"`

	// The allocation delay set for the operator, set as zero for immediate allocation
	AllocationDelay uint32 `toml:"allocation_delay"`
}
