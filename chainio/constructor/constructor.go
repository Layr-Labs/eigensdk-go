package constructor

import (
	"context"

	avsregistry "github.com/Layr-Labs/eigensdk-go/chainio/avsregistry"
	clients "github.com/Layr-Labs/eigensdk-go/chainio/clients"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	elcontracts "github.com/Layr-Labs/eigensdk-go/chainio/elcontracts"
	blspubkeyreg "github.com/Layr-Labs/eigensdk-go/contracts/bindings/BLSPubkeyRegistry"
	blsregcoord "github.com/Layr-Labs/eigensdk-go/contracts/bindings/BLSRegistryCoordinatorWithIndices"
	logging "github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/metrics"
	"github.com/Layr-Labs/eigensdk-go/signer"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethcommon "github.com/ethereum/go-ethereum/common"
	crypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/prometheus/client_golang/prometheus"
)

type Config struct {
	EcdsaPrivateKeyString         string `yaml:"ecdsa_private_key_string"`
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

func BuildClients(config Config, logger logging.Logger) (*Clients, error) {

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

	// creating EL contracts client (simple wrapper around abigen bindings)
	elContractsClient, err := config.buildElContractsChainClient(logger, ethHttpClient, ethWsClient)
	if err != nil {
		logger.Error("Failed to create ELHttpContractsChainClient", "err", err)
		return nil, err
	}

	// creating EL clients: Reader, Writer and Subscriber
	elChainReader, elChainWriter, elChainSubscriber, err := config.buildElClients(
		elContractsClient,
		ethHttpClient,
		logger,
		eigenMetrics,
	)
	if err != nil {
		logger.Error("Failed to create EL Reader, Writer and Subscriber", "err", err)
		return nil, err
	}

	// creating AVS contracts client
	avsRegistryContractsClient, err := config.buildAvsRegistryContractsChainClient(logger, ethHttpClient)
	if err != nil {
		logger.Error("Failed to create AvsRegistryContractsChainClient", "err", err)
		return nil, err
	}

	// creating AVS clients: Reader and Writer
	avsRegistryChainReader, avsRegistryChainWriter, err := config.buildAvsClients(
		avsRegistryContractsClient,
		logger,
		ethHttpClient,
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
		PrometheusRegistry:     promReg,
	}, nil

}

func (config *Config) buildElContractsChainClient(
	logger logging.Logger,
	ethHttpClient eth.EthClient,
	ethWsClient eth.EthClient,
) (clients.ELContractsClient, error) {
	blsRegistryCoordinatorAddr := gethcommon.HexToAddress(config.BlsRegistryCoordinatorAddr)
	contractBLSRegistryCoordWithIndices, err := blsregcoord.NewContractBLSRegistryCoordinatorWithIndices(
		blsRegistryCoordinatorAddr,
		ethHttpClient,
	)
	if err != nil {
		logger.Fatal("Failed to fetch BLSRegistryCoordinatorWithIndices contract", "err", err)
	}

	blsPubkeyRegistryAddr, err := contractBLSRegistryCoordWithIndices.BlsPubkeyRegistry(&bind.CallOpts{})
	if err != nil {
		logger.Fatal("Failed to fetch BlsPubkeyRegistry contract", "err", err)
	}
	contractBlsPubkeyRegistry, err := blspubkeyreg.NewContractBLSPubkeyRegistry(blsPubkeyRegistryAddr, ethHttpClient)
	if err != nil {
		logger.Fatal("Failed to construct BlsPubkeyRegistry contract", "err", err)
	}
	blsPubKeyCompendiumAddr, err := contractBlsPubkeyRegistry.PubkeyCompendium(&bind.CallOpts{})
	if err != nil {
		logger.Fatal("Failed to fetch PubkeyCompendium contract", "err", err)
	}

	slasherAddr, err := contractBLSRegistryCoordWithIndices.Slasher(&bind.CallOpts{})
	if err != nil {
		logger.Error("Failed to fetch Slasher contract", "err", err)
		return nil, err
	}
	elContractsChainClient, err := clients.NewELContractsChainClient(
		slasherAddr,
		blsPubKeyCompendiumAddr,
		ethHttpClient,
		ethWsClient,
		logger,
	)
	if err != nil {
		logger.Error("Failed to create ELContractsChainClient", "err", err)
		return nil, err
	}

	return elContractsChainClient, nil
}

func (config *Config) buildElClients(
	elContractsClient clients.ELContractsClient,
	ethHttpClient eth.EthClient,
	logger logging.Logger,
	eigenMetrics *metrics.EigenMetrics,
) (elcontracts.ELReader, elcontracts.ELWriter, elcontracts.ELSubscriber, error) {

	// get the Reader for the EL contracts
	elChainReader, err := elcontracts.NewELChainReader(
		elContractsClient,
		logger,
		ethHttpClient,
	)
	if err != nil {
		logger.Error("Failed to create ELChainReader", "err", err)
		return nil, nil, nil, err
	}

	// get the Subscriber for the EL contracts
	elChainSubscriber, err := elcontracts.NewELChainSubscriber(
		elContractsClient,
		logger,
	)
	if err != nil {
		logger.Error("Failed to create ELChainSubscriber", "err", err)
		return nil, nil, nil, err
	}

	// get the Writer for the EL contracts
	ecdsaPrivateKey, err := crypto.HexToECDSA(config.EcdsaPrivateKeyString)
	if err != nil {
		logger.Errorf("Cannot parse ecdsa private key", "err", err)
		return nil, nil, nil, err
	}

	chainId, err := ethHttpClient.ChainID(context.Background())
	if err != nil {
		logger.Error("Cannot get chainId", "err", err)
		return nil, nil, nil, err
	}

	privateKeySigner, err := signer.NewPrivateKeySigner(ecdsaPrivateKey, chainId)
	if err != nil {
		logger.Error("Cannot create signer", "err", err)
		return nil, nil, nil, err
	}

	elChainWriter := elcontracts.NewELChainWriter(
		elContractsClient,
		ethHttpClient,
		privateKeySigner,
		logger,
		eigenMetrics,
	)
	if err != nil {
		logger.Error("Failed to create ELChainWriter", "err", err)
		return nil, nil, nil, err
	}

	return elChainReader, elChainWriter, elChainSubscriber, nil
}

func (config *Config) buildAvsRegistryContractsChainClient(
	logger logging.Logger,
	ethHttpClient eth.EthClient,
) (clients.AVSRegistryContractsClient, error) {
	blsRegistryCoordinatorAddr := gethcommon.HexToAddress(config.BlsRegistryCoordinatorAddr)
	contractBLSRegistryCoordWithIndices, err := blsregcoord.NewContractBLSRegistryCoordinatorWithIndices(
		blsRegistryCoordinatorAddr,
		ethHttpClient,
	)
	if err != nil {
		logger.Error("Failed to fetch BLSRegistryCoordinatorWithIndices contract", "err", err)
		return nil, err
	}

	stakeregistryAddr, err := contractBLSRegistryCoordWithIndices.StakeRegistry(&bind.CallOpts{})
	if err != nil {
		logger.Error("Failed to fetch StakeRegistry contract", "err", err)
		return nil, err
	}

	blsPubkeyRegistryAddr, err := contractBLSRegistryCoordWithIndices.BlsPubkeyRegistry(&bind.CallOpts{})
	if err != nil {
		logger.Error("Failed to fetch BlsPubkeyRegistry contract", "err", err)
		return nil, err
	}

	avsregistryContractsChainClient, err := clients.NewAvsRegistryContractsChainClient(
		blsRegistryCoordinatorAddr,
		gethcommon.HexToAddress(config.BlsOperatorStateRetrieverAddr),
		stakeregistryAddr,
		blsPubkeyRegistryAddr,
		ethHttpClient,
		logger,
	)
	if err != nil {
		logger.Error("Failed to create AVSRegistryContractsChainClient", "err", err)
		return nil, err
	}

	return avsregistryContractsChainClient, nil
}

func (config *Config) buildAvsClients(
	avsContractsClient clients.AVSRegistryContractsClient,
	logger logging.Logger,
	ethHttpClient eth.EthClient,
) (avsregistry.AvsRegistryReader, avsregistry.AvsRegistryWriter, error) {
	avsRegistryChainReader, err := avsregistry.NewAvsRegistryReader(
		avsContractsClient,
		logger,
		ethHttpClient,
	)
	if err != nil {
		logger.Error("Failed to create AVSRegistryChainReader", "err", err)
		return nil, nil, err
	}

	ecdsaPrivateKey, err := crypto.HexToECDSA(config.EcdsaPrivateKeyString)
	if err != nil {
		logger.Errorf("Cannot parse ecdsa private key", "err", err)
		return nil, nil, err
	}

	chainId, err := ethHttpClient.ChainID(context.Background())
	if err != nil {
		logger.Error("Cannot get chainId", "err", err)
		return nil, nil, err
	}

	privateKeySigner, err := signer.NewPrivateKeySigner(ecdsaPrivateKey, chainId)
	if err != nil {
		logger.Error("Cannot create signer", "err", err)
		return nil, nil, err
	}

	avsRegistryChainWriter, err := avsregistry.NewAvsRegistryWriter(
		avsContractsClient,
		logger,
		privateKeySigner,
		ethHttpClient,
	)
	if err != nil {
		logger.Error("Failed to create AVSRegistryChainWriter", "err", err)
		return nil, nil, err
	}

	return avsRegistryChainReader, avsRegistryChainWriter, nil
}
