package clients

import (
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/avsregistry"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/elcontracts"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	chainioutils "github.com/Layr-Labs/eigensdk-go/chainio/utils"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/metrics"
	"github.com/Layr-Labs/eigensdk-go/signerv2"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/prometheus/client_golang/prometheus"
)

type BuildAllConfig struct {
	EthHttpUrl                 string `yaml:"eth_http_url"`
	EthWsUrl                   string `yaml:"eth_ws_url"`
	RegistryCoordinatorAddr    string `yaml:"bls_registry_coordinator_address"`
	OperatorStateRetrieverAddr string `yaml:"bls_operator_state_retriever_address"`
	AvsName                    string `yaml:"avs_name"`
	PromMetricsIpPortAddress   string `yaml:"prometheus_metrics_ip_port_address"`
}

// TODO: this is confusing right now because clients are not instrumented clients, but
// we return metrics and prometheus reg, so user has to build instrumented clients at the call
// site if they need them. We should probably separate into two separate constructors, one
// for non-instrumented clients that doesn't return metrics/reg, and another instrumented-constructor
// that returns instrumented clients and the metrics/reg.
type Clients struct {
	AvsRegistryChainReader     *avsregistry.AvsRegistryChainReader
	AvsRegistryChainSubscriber *avsregistry.AvsRegistryChainSubscriber
	AvsRegistryChainWriter     *avsregistry.AvsRegistryChainWriter
	ElChainReader              *elcontracts.ELChainReader
	ElChainWriter              *elcontracts.ELChainWriter
	EthHttpClient              *eth.Client
	EthWsClient                *eth.Client
	Metrics                    *metrics.EigenMetrics // exposes main avs node spec metrics that need to be incremented by avs code and used to start the metrics server
	PrometheusRegistry         *prometheus.Registry  // Used if avs teams need to register avs-specific metrics
}

func BuildAll(config BuildAllConfig, signerAddr gethcommon.Address, signerFn signerv2.SignerFn, logger logging.Logger) (*Clients, error) {
	config.validate(logger)

	// Create the metrics server
	promReg := prometheus.NewRegistry()
	eigenMetrics := metrics.NewEigenMetrics(config.AvsName, config.PromMetricsIpPortAddress, promReg, logger)

	// creating two types of Eth clients: HTTP and WS
	ethHttpClient, err := eth.NewClient(config.EthHttpUrl)
	if err != nil {
		logger.Error("Failed to create Eth Http client", "err", err)
		return nil, err
	}

	ethWsClient, err := eth.NewClient(config.EthWsUrl)
	if err != nil {
		logger.Error("Failed to create Eth WS client", "err", err)
		return nil, err
	}

	// TODO(madhur): is it fine to not set the sender address?
	txMgr := txmgr.NewSimpleTxManager(ethHttpClient, logger, signerFn, signerAddr)
	// creating EL clients: Reader, Writer and Subscriber
	elChainReader, elChainWriter, err := config.buildElClients(
		ethHttpClient,
		txMgr,
		logger,
		eigenMetrics,
	)
	if err != nil {
		logger.Error("Failed to create EL Reader, Writer and Subscriber", "err", err)
		return nil, err
	}

	// creating AVS clients: Reader and Writer
	avsRegistryChainReader, avsRegistryChainSubscriber, avsRegistryChainWriter, err := config.buildAvsClients(
		ethHttpClient,
		ethWsClient,
		txMgr,
		logger,
	)
	if err != nil {
		logger.Error("Failed to create AVS Registry Reader and Writer", "err", err)
		return nil, err
	}

	return &Clients{
		ElChainReader:              elChainReader,
		ElChainWriter:              elChainWriter,
		AvsRegistryChainReader:     avsRegistryChainReader,
		AvsRegistryChainSubscriber: avsRegistryChainSubscriber,
		AvsRegistryChainWriter:     avsRegistryChainWriter,
		EthHttpClient:              ethHttpClient,
		EthWsClient:                ethWsClient,
		Metrics:                    eigenMetrics,
		PrometheusRegistry:         promReg,
	}, nil

}

func (config *BuildAllConfig) buildElClients(
	ethHttpClient eth.EthClient,
	txMgr txmgr.TxManager,
	logger logging.Logger,
	eigenMetrics *metrics.EigenMetrics,
) (*elcontracts.ELChainReader, *elcontracts.ELChainWriter, error) {

	avsRegistryContractBindings, err := chainioutils.NewAVSRegistryContractBindings(
		gethcommon.HexToAddress(config.RegistryCoordinatorAddr),
		gethcommon.HexToAddress(config.OperatorStateRetrieverAddr),
		ethHttpClient,
		logger,
	)
	if err != nil {
		logger.Error("Failed to create AVSRegistryContractBindings", "err", err)
		return nil, nil, err
	}

	delegationManagerAddr, err := avsRegistryContractBindings.StakeRegistry.Delegation(&bind.CallOpts{})
	if err != nil {
		logger.Fatal("Failed to fetch Slasher contract", "err", err)
	}

	elContractBindings, err := chainioutils.NewEigenlayerContractBindings(
		delegationManagerAddr,
		ethHttpClient,
		logger,
	)
	if err != nil {
		logger.Error("Failed to create EigenlayerContractBindings", "err", err)
		return nil, nil, err
	}

	// get the Reader for the EL contracts
	elChainReader := elcontracts.NewELChainReader(
		elContractBindings.Slasher,
		elContractBindings.DelegationManager,
		elContractBindings.StrategyManager,
		logger,
		ethHttpClient,
	)

	elChainWriter := elcontracts.NewELChainWriter(
		elContractBindings.Slasher,
		elContractBindings.DelegationManager,
		elContractBindings.StrategyManager,
		elContractBindings.StrategyManagerAddr,
		elChainReader,
		ethHttpClient,
		logger,
		eigenMetrics,
		txMgr,
	)
	if err != nil {
		logger.Error("Failed to create ELChainWriter", "err", err)
		return nil, nil, err
	}

	return elChainReader, elChainWriter, nil
}

func (config *BuildAllConfig) buildAvsClients(
	ethHttpClient eth.EthClient,
	ethWsClient eth.EthClient,
	txMgr txmgr.TxManager,
	logger logging.Logger,
) (*avsregistry.AvsRegistryChainReader, *avsregistry.AvsRegistryChainSubscriber, *avsregistry.AvsRegistryChainWriter, error) {

	avsRegistryContractBindings, err := chainioutils.NewAVSRegistryContractBindings(
		gethcommon.HexToAddress(config.RegistryCoordinatorAddr),
		gethcommon.HexToAddress(config.OperatorStateRetrieverAddr),
		ethHttpClient,
		logger,
	)
	if err != nil {
		logger.Error("Failed to create AVSRegistryContractBindings", "err", err)
		return nil, nil, nil, err
	}

	avsRegistryChainReader := avsregistry.NewAvsRegistryChainReader(
		avsRegistryContractBindings.RegistryCoordinatorAddr,
		avsRegistryContractBindings.BlsApkRegistryAddr,
		avsRegistryContractBindings.RegistryCoordinator,
		avsRegistryContractBindings.OperatorStateRetriever,
		avsRegistryContractBindings.StakeRegistry,
		logger,
		ethHttpClient,
	)

	avsRegistryChainWriter, err := avsregistry.NewAvsRegistryChainWriter(
		avsRegistryContractBindings.RegistryCoordinator,
		avsRegistryContractBindings.OperatorStateRetriever,
		avsRegistryContractBindings.StakeRegistry,
		avsRegistryContractBindings.BlsApkRegistry,
		logger,
		ethHttpClient,
		txMgr,
	)
	if err != nil {
		logger.Error("Failed to create AVSRegistryChainWriter", "err", err)
		return nil, nil, nil, err
	}

	// get the Subscriber for Avs Registry contracts
	// note that the subscriber needs a ws connection instead of http
	avsRegistrySubscriber, err := avsregistry.BuildAvsRegistryChainSubscriber(
		gethcommon.HexToAddress(config.RegistryCoordinatorAddr),
		ethWsClient,
		logger,
	)
	if err != nil {
		logger.Error("Failed to create ELChainSubscriber", "err", err)
		return nil, nil, nil, err
	}

	return avsRegistryChainReader, avsRegistrySubscriber, avsRegistryChainWriter, nil
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
