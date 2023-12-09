package clients

import (
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/avsregistry"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/elcontracts"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	chainioutils "github.com/Layr-Labs/eigensdk-go/chainio/utils"
	blspubkeycompendium "github.com/Layr-Labs/eigensdk-go/contracts/bindings/BLSPublicKeyCompendium"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/metrics"
	"github.com/Layr-Labs/eigensdk-go/signerv2"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/prometheus/client_golang/prometheus"
)

type BuildAllConfig struct {
	EthHttpUrl                    string `yaml:"eth_http_url"`
	EthWsUrl                      string `yaml:"eth_ws_url"`
	BlsRegistryCoordinatorAddr    string `yaml:"bls_registry_coordinator_address"`
	BlsOperatorStateRetrieverAddr string `yaml:"bls_operator_state_retriever_address"`
	AvsName                       string `yaml:"avs_name"`
	PromMetricsIpPortAddress      string `yaml:"prometheus_metrics_ip_port_address"`
}

// TODO: this is confusing right now because clients are not instrumented clients, but
// we return metrics and prometheus reg, so user has to build instrumented clients at the call
// site if they need them. We should probably separate into two separate constructors, one
// for non-instrumented clients that doesn't return metrics/reg, and another instrumented-constructor
// that returns instrumented clients and the metrics/reg.
type Clients struct {
	AvsRegistryChainReader avsregistry.AvsRegistryReader
	AvsRegistryChainWriter avsregistry.AvsRegistryWriter
	ElChainReader          elcontracts.ELReader
	ElChainSubscriber      elcontracts.ELSubscriber
	ElChainWriter          elcontracts.ELWriter
	EthHttpClient          *eth.Client
	EthWsClient            *eth.Client
	Metrics                *metrics.EigenMetrics // exposes main avs node spec metrics that need to be incremented by avs code and used to start the metrics server
	PrometheusRegistry     *prometheus.Registry  // Used if avs teams need to register avs-specific metrics
}

func BuildAll(config BuildAllConfig, signer signerv2.SignerFn, logger logging.Logger) (*Clients, error) {
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

	txMgr := txmgr.NewSimpleTxManager(ethHttpClient, logger, signer, gethcommon.Address{})
	// creating EL clients: Reader, Writer and Subscriber
	elChainReader, elChainWriter, elChainSubscriber, err := config.buildElClients(
		ethHttpClient,
		ethWsClient,
		txMgr,
		logger,
		eigenMetrics,
	)
	if err != nil {
		logger.Error("Failed to create EL Reader, Writer and Subscriber", "err", err)
		return nil, err
	}

	// creating AVS clients: Reader and Writer
	avsRegistryChainReader, avsRegistryChainWriter, err := config.buildAvsClients(
		ethHttpClient,
		txMgr,
		logger,
	)
	if err != nil {
		logger.Error("Failed to create AVS Registry Reader and Writer", "err", err)
		return nil, err
	}

	return &Clients{
		ElChainReader:          elChainReader,
		ElChainSubscriber:      elChainSubscriber,
		ElChainWriter:          elChainWriter,
		AvsRegistryChainReader: avsRegistryChainReader,
		AvsRegistryChainWriter: avsRegistryChainWriter,
		EthHttpClient:          ethHttpClient,
		EthWsClient:            ethWsClient,
		Metrics:                eigenMetrics,
		PrometheusRegistry:     promReg,
	}, nil

}

func (config *BuildAllConfig) buildElClients(
	ethHttpClient eth.EthClient,
	ethWsClient eth.EthClient,
	txMgr txmgr.TxManager,
	logger logging.Logger,
	eigenMetrics *metrics.EigenMetrics,
) (elcontracts.ELReader, elcontracts.ELWriter, elcontracts.ELSubscriber, error) {

	avsRegistryContractBindings, err := chainioutils.NewAVSRegistryContractBindings(
		gethcommon.HexToAddress(config.BlsRegistryCoordinatorAddr),
		gethcommon.HexToAddress(config.BlsOperatorStateRetrieverAddr),
		ethHttpClient,
		logger,
	)
	if err != nil {
		logger.Error("Failed to create AVSRegistryContractBindings", "err", err)
		return nil, nil, nil, err
	}

	slasherAddr, err := avsRegistryContractBindings.RegistryCoordinator.Slasher(&bind.CallOpts{})
	if err != nil {
		logger.Fatal("Failed to fetch Slasher contract", "err", err)
	}

	elContractBindings, err := chainioutils.NewEigenlayerContractBindings(
		slasherAddr,
		avsRegistryContractBindings.BlsPubkeyCompendiumAddr,
		ethHttpClient,
		logger,
	)
	if err != nil {
		logger.Error("Failed to create EigenlayerContractBindings", "err", err)
		return nil, nil, nil, err
	}

	// get the Reader for the EL contracts
	elChainReader := elcontracts.NewELChainReader(
		elContractBindings.Slasher,
		elContractBindings.DelegationManager,
		elContractBindings.StrategyManager,
		elContractBindings.BlsPubkeyCompendium,
		elContractBindings.BlspubkeyCompendiumAddr,
		logger,
		ethHttpClient,
	)

	// get the Subscriber for the EL contracts
	contractBlsPubkeyCompendiumWs, err := blspubkeycompendium.NewContractBLSPublicKeyCompendium(
		elContractBindings.BlspubkeyCompendiumAddr,
		ethWsClient,
	)
	if err != nil {
		logger.Fatal("Failed to fetch BLSPublicKeyCompendium contract", "err", err)
	}
	elChainSubscriber, err := elcontracts.NewELChainSubscriber(
		contractBlsPubkeyCompendiumWs,
		logger,
	)
	if err != nil {
		logger.Error("Failed to create ELChainSubscriber", "err", err)
		return nil, nil, nil, err
	}

	elChainWriter := elcontracts.NewELChainWriter(
		elContractBindings.Slasher,
		elContractBindings.DelegationManager,
		elContractBindings.StrategyManager,
		elContractBindings.StrategyManagerAddr,
		elContractBindings.BlsPubkeyCompendium,
		elContractBindings.BlspubkeyCompendiumAddr,
		elChainReader,
		ethHttpClient,
		logger,
		eigenMetrics,
		txMgr,
	)
	if err != nil {
		logger.Error("Failed to create ELChainWriter", "err", err)
		return nil, nil, nil, err
	}

	return elChainReader, elChainWriter, elChainSubscriber, nil
}

func (config *BuildAllConfig) buildAvsClients(
	ethHttpClient eth.EthClient,
	txMgr txmgr.TxManager,
	logger logging.Logger,
) (avsregistry.AvsRegistryReader, avsregistry.AvsRegistryWriter, error) {

	avsRegistryContractBindings, err := chainioutils.NewAVSRegistryContractBindings(
		gethcommon.HexToAddress(config.BlsRegistryCoordinatorAddr),
		gethcommon.HexToAddress(config.BlsOperatorStateRetrieverAddr),
		ethHttpClient,
		logger,
	)
	if err != nil {
		logger.Error("Failed to create AVSRegistryContractBindings", "err", err)
		return nil, nil, err
	}

	avsRegistryChainReader, err := avsregistry.NewAvsRegistryReader(
		avsRegistryContractBindings.RegistryCoordinatorAddr,
		avsRegistryContractBindings.RegistryCoordinator,
		avsRegistryContractBindings.BlsOperatorStateRetriever,
		avsRegistryContractBindings.StakeRegistry,
		logger,
		ethHttpClient,
	)
	if err != nil {
		logger.Error("Failed to create AVSRegistryChainReader", "err", err)
		return nil, nil, err
	}

	avsRegistryChainWriter, err := avsregistry.NewAvsRegistryWriter(
		avsRegistryContractBindings.RegistryCoordinator,
		avsRegistryContractBindings.BlsOperatorStateRetriever,
		avsRegistryContractBindings.StakeRegistry,
		avsRegistryContractBindings.BlsPubkeyRegistry,
		logger,
		ethHttpClient,
		txMgr,
	)
	if err != nil {
		logger.Error("Failed to create AVSRegistryChainWriter", "err", err)
		return nil, nil, err
	}

	return avsRegistryChainReader, avsRegistryChainWriter, nil
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
	if config.BlsRegistryCoordinatorAddr == "" {
		logger.Fatalf("BuildAllConfig.validate: Missing bls registry coordinator address")
	}
	if config.BlsOperatorStateRetrieverAddr == "" {
		logger.Fatalf("BuildAllConfig.validate: Missing bls operator state retriever address")
	}
	if config.AvsName == "" {
		logger.Fatalf("BuildAllConfig.validate: Missing avs name")
	}
	if config.PromMetricsIpPortAddress == "" {
		logger.Fatalf("BuildAllConfig.validate: Missing prometheus metrics ip port address")
	}
}
