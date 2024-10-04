package clients

import (
	"context"
	"crypto/ecdsa"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/avsregistry"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/elcontracts"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/wallet"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/metrics"
	"github.com/Layr-Labs/eigensdk-go/signerv2"
	"github.com/Layr-Labs/eigensdk-go/utils"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/prometheus/client_golang/prometheus"
)

type BuildAllConfig struct {
	EthHttpUrl                 string
	EthWsUrl                   string
	RegistryCoordinatorAddr    string
	OperatorStateRetrieverAddr string
	AvsName                    string
	PromMetricsIpPortAddress   string
}

// ReadClients is a struct that holds only the read clients for interacting with the AVS and EL contracts.
type ReadClients struct {
	AvsRegistryChainReader      *avsregistry.ChainReader
	AvsRegistryChainSubscriber  *avsregistry.ChainSubscriber
	ElChainReader               *elcontracts.ChainReader
	EthHttpClient               eth.HttpBackend
	EthWsClient                 eth.WsBackend
	AvsRegistryContractBindings *avsregistry.ContractBindings
	EigenlayerContractBindings  *elcontracts.ContractBindings
	Metrics                     *metrics.EigenMetrics // exposes main avs node spec metrics that need to be incremented by avs code and used to start the metrics server
	PrometheusRegistry          *prometheus.Registry
}

// Clients is a struct that holds all the clients that are needed to interact with the AVS and EL contracts.
type Clients struct {
	ReadClients
	Wallet                 wallet.Wallet
	TxManager              txmgr.TxManager
	ElChainWriter          *elcontracts.ChainWriter
	AvsRegistryChainWriter *avsregistry.ChainWriter
}

// BuildReadClients creates all the read clients needed to interact with the AVS and EL contracts.
func BuildReadClients(
	config BuildAllConfig,
	logger logging.Logger,
) (*ReadClients, error) {
	config.validate(logger)

	// Create the metrics server
	promReg := prometheus.NewRegistry()
	eigenMetrics := metrics.NewEigenMetrics(config.AvsName, config.PromMetricsIpPortAddress, promReg, logger)

	// creating two types of Eth clients: HTTP and WS
	ethHttpClient, err := ethclient.Dial(config.EthHttpUrl)
	if err != nil {
		return nil, utils.WrapError("Failed to create Eth Http client", err)
	}

	ethWsClient, err := ethclient.Dial(config.EthWsUrl)
	if err != nil {
		return nil, utils.WrapError("Failed to create Eth WS client", err)
	}

	// creating AVS clients: Reader
	avsRegistryChainReader, avsRegistryChainSubscriber, avsRegistryContractBindings, err := avsregistry.BuildReadClients(
		avsregistry.Config{
			RegistryCoordinatorAddress:    gethcommon.HexToAddress(config.RegistryCoordinatorAddr),
			OperatorStateRetrieverAddress: gethcommon.HexToAddress(config.OperatorStateRetrieverAddr),
		},
		ethHttpClient,
		ethWsClient,
		logger,
	)
	if err != nil {
		return nil, utils.WrapError("Failed to create AVS Registry Reader and Writer", err)
	}

	// creating EL clients: Reader and EigenLayer Contract Bindings
	elChainReader, elContractBindings, err := elcontracts.BuildReadClients(
		elcontracts.Config{
			DelegationManagerAddress: avsRegistryContractBindings.DelegationManagerAddr,
			AvsDirectoryAddress:      avsRegistryContractBindings.AvsDirectoryAddr,
		},
		ethHttpClient,
		logger,
		eigenMetrics,
	)
	if err != nil {
		return nil, utils.WrapError("Failed to create EL Reader and Writer", err)
	}

	readClients := ReadClients{
		ElChainReader:               elChainReader,
		AvsRegistryChainReader:      avsRegistryChainReader,
		AvsRegistryChainSubscriber:  avsRegistryChainSubscriber,
		EthHttpClient:               ethHttpClient,
		EthWsClient:                 ethWsClient,
		EigenlayerContractBindings:  elContractBindings,
		AvsRegistryContractBindings: avsRegistryContractBindings,
		Metrics:                     eigenMetrics,
		PrometheusRegistry:          promReg,
	}
	return &readClients, nil
}

// BuildAll creates all the clients needed to interact with the AVS and EL contracts. For both read and write
// operations.
// TODO: this is confusing right now because clients are not instrumented clients, but
// we return metrics and prometheus reg, so user has to build instrumented clients at the call
// site if they need them. We should probably separate into two separate constructors, one
// for non-instrumented clients that doesn't return metrics/reg, and another instrumented-constructor
// that returns instrumented clients and the metrics/reg.
func BuildAll(
	config BuildAllConfig,
	ecdsaPrivateKey *ecdsa.PrivateKey,
	logger logging.Logger,
) (*Clients, error) {
	config.validate(logger)

	// Create the metrics server
	promReg := prometheus.NewRegistry()
	eigenMetrics := metrics.NewEigenMetrics(config.AvsName, config.PromMetricsIpPortAddress, promReg, logger)

	// creating two types of Eth clients: HTTP and WS
	ethHttpClient, err := ethclient.Dial(config.EthHttpUrl)
	if err != nil {
		return nil, utils.WrapError("Failed to create Eth Http client", err)
	}

	ethWsClient, err := ethclient.Dial(config.EthWsUrl)
	if err != nil {
		return nil, utils.WrapError("Failed to create Eth WS client", err)
	}

	rpcCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	chainid, err := ethHttpClient.ChainID(rpcCtx)
	if err != nil {
		logger.Fatal("Cannot get chain id", "err", err)
	}
	signerV2, addr, err := signerv2.SignerFromConfig(signerv2.Config{PrivateKey: ecdsaPrivateKey}, chainid)
	if err != nil {
		return nil, utils.WrapError("Failed to create the signer from the given config", err)
	}

	pkWallet, err := wallet.NewPrivateKeyWallet(ethHttpClient, signerV2, addr, logger)
	if err != nil {
		return nil, utils.WrapError("Failed to create transaction sender", err)
	}
	txMgr := txmgr.NewSimpleTxManager(pkWallet, ethHttpClient, logger, addr)

	// creating AVS clients: Reader and Writer
	avsRegistryChainReader, avsRegistryChainSubscriber, avsRegistryChainWriter, avsRegistryContractBindings, err := avsregistry.BuildClients(
		avsregistry.Config{
			RegistryCoordinatorAddress:    gethcommon.HexToAddress(config.RegistryCoordinatorAddr),
			OperatorStateRetrieverAddress: gethcommon.HexToAddress(config.OperatorStateRetrieverAddr),
		},
		ethHttpClient,
		ethWsClient,
		txMgr,
		logger,
	)
	if err != nil {
		return nil, utils.WrapError("Failed to create AVS Registry Reader and Writer", err)
	}

	// creating EL clients: Reader, Writer and EigenLayer Contract Bindings
	elChainReader, elChainWriter, elContractBindings, err := elcontracts.BuildClients(
		elcontracts.Config{
			DelegationManagerAddress: avsRegistryContractBindings.DelegationManagerAddr,
			AvsDirectoryAddress:      avsRegistryContractBindings.AvsDirectoryAddr,
		},
		ethHttpClient,
		txMgr,
		logger,
		eigenMetrics,
	)
	if err != nil {
		return nil, utils.WrapError("Failed to create EL Reader and Writer", err)
	}

	readClients := ReadClients{
		ElChainReader:               elChainReader,
		AvsRegistryChainReader:      avsRegistryChainReader,
		AvsRegistryChainSubscriber:  avsRegistryChainSubscriber,
		EthHttpClient:               ethHttpClient,
		EthWsClient:                 ethWsClient,
		EigenlayerContractBindings:  elContractBindings,
		AvsRegistryContractBindings: avsRegistryContractBindings,
		Metrics:                     eigenMetrics,
		PrometheusRegistry:          promReg,
	}

	return &Clients{
		ReadClients:            readClients,
		ElChainWriter:          elChainWriter,
		AvsRegistryChainWriter: avsRegistryChainWriter,
		Wallet:                 pkWallet,
		TxManager:              txMgr,
	}, nil
}

// Very basic validation that makes sure all fields are nonempty
// we might eventually want more sophisticated validation, based on regexp,
// or use something like https://json-schema.org/ (?)
func (config *BuildAllConfig) validate(logger logging.Logger) {
	if config.EthHttpUrl == "" {
		logger.Fatalf("BuildAllConfig.validate: Missing eth http url")
	}
	if config.EthWsUrl == "" {
		logger.Fatalf("BuildAllConfig.validate: Missing eth ws url")
	}
	if config.RegistryCoordinatorAddr == "" {
		logger.Fatalf("BuildAllConfig.validate: Missing bls registry coordinator address")
	}
	if config.OperatorStateRetrieverAddr == "" {
		logger.Fatalf("BuildAllConfig.validate: Missing bls operator state retriever address")
	}
	if config.AvsName == "" {
		logger.Fatalf("BuildAllConfig.validate: Missing avs name")
	}
	if config.PromMetricsIpPortAddress == "" {
		logger.Fatalf("BuildAllConfig.validate: Missing prometheus metrics ip port address")
	}
}
