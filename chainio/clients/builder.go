package clients

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"time"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/avsregistry"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/elcontracts"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/wallet"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	chainioutils "github.com/Layr-Labs/eigensdk-go/chainio/utils"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/metrics"
	"github.com/Layr-Labs/eigensdk-go/signerv2"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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
	EthHttpClient              eth.Client
	EthWsClient                eth.Client
	Metrics                    *metrics.EigenMetrics // exposes main avs node spec metrics that need to be incremented by avs code and used to start the metrics server
	PrometheusRegistry         *prometheus.Registry  // Used if avs teams need to register avs-specific metrics
}

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
	ethHttpClient, err := eth.NewClient(config.EthHttpUrl)
	if err != nil {
		return nil, types.WrapError(errors.New("Failed to create Eth Http client"), err)
	}

	ethWsClient, err := eth.NewClient(config.EthWsUrl)
	if err != nil {
		return nil, types.WrapError(errors.New("Failed to create Eth WS client"), err)
	}

	rpcCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	chainid, err := ethHttpClient.ChainID(rpcCtx)
	if err != nil {
		logger.Fatal("Cannot get chain id", "err", err)
	}
	signerV2, addr, err := signerv2.SignerFromConfig(signerv2.Config{PrivateKey: ecdsaPrivateKey}, chainid)
	if err != nil {
		panic(err)
	}

	pkWallet, err := wallet.NewPrivateKeyWallet(ethHttpClient, signerV2, addr, logger)
	if err != nil {
		return nil, types.WrapError(errors.New("Failed to create transaction sender"), err)
	}
	txMgr := txmgr.NewSimpleTxManager(pkWallet, ethHttpClient, logger, addr)
	// creating EL clients: Reader, Writer and Subscriber
	elChainReader, elChainWriter, err := config.buildElClients(
		ethHttpClient,
		txMgr,
		logger,
		eigenMetrics,
	)
	if err != nil {
		return nil, types.WrapError(errors.New("Failed to create EL Reader, Writer and Subscriber"), err)
	}

	// creating AVS clients: Reader and Writer
	avsRegistryChainReader, avsRegistryChainSubscriber, avsRegistryChainWriter, err := config.buildAvsClients(
		elChainReader,
		ethHttpClient,
		ethWsClient,
		txMgr,
		logger,
	)
	if err != nil {
		return nil, types.WrapError(errors.New("Failed to create AVS Registry Reader and Writer"), err)
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
	ethHttpClient eth.Client,
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
		return nil, nil, types.WrapError(errors.New("Failed to create AVSRegistryContractBindings"), err)
	}

	delegationManagerAddr, err := avsRegistryContractBindings.StakeRegistry.Delegation(&bind.CallOpts{})
	if err != nil {
		logger.Fatal("Failed to fetch Slasher contract", "err", err)
	}
	avsDirectoryAddr, err := avsRegistryContractBindings.ServiceManager.AvsDirectory(&bind.CallOpts{})
	if err != nil {
		logger.Fatal("Failed to fetch Slasher contract", "err", err)
	}

	elContractBindings, err := chainioutils.NewEigenlayerContractBindings(
		delegationManagerAddr,
		avsDirectoryAddr,
		ethHttpClient,
		logger,
	)
	if err != nil {
		return nil, nil, types.WrapError(errors.New("Failed to create EigenlayerContractBindings"), err)
	}

	// get the Reader for the EL contracts
	elChainReader := elcontracts.NewELChainReader(
		elContractBindings.Slasher,
		elContractBindings.DelegationManager,
		elContractBindings.StrategyManager,
		elContractBindings.AvsDirectory,
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

	return elChainReader, elChainWriter, nil
}

func (config *BuildAllConfig) buildAvsClients(
	elReader elcontracts.ELReader,
	ethHttpClient eth.Client,
	ethWsClient eth.Client,
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
		return nil, nil, nil, types.WrapError(errors.New("Failed to create AVSRegistryContractBindings"), err)
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
		avsRegistryContractBindings.ServiceManagerAddr,
		avsRegistryContractBindings.RegistryCoordinator,
		avsRegistryContractBindings.OperatorStateRetriever,
		avsRegistryContractBindings.StakeRegistry,
		avsRegistryContractBindings.BlsApkRegistry,
		elReader,
		logger,
		ethHttpClient,
		txMgr,
	)
	if err != nil {
		return nil, nil, nil, types.WrapError(errors.New("Failed to create AVSRegistryChainWriter"), err)
	}

	// get the Subscriber for Avs Registry contracts
	// note that the subscriber needs a ws connection instead of http
	avsRegistrySubscriber, err := avsregistry.BuildAvsRegistryChainSubscriber(
		avsRegistryContractBindings.RegistryCoordinatorAddr,
		ethWsClient,
		logger,
	)
	if err != nil {
		return nil, nil, nil, types.WrapError(errors.New("Failed to create ELChainSubscriber"), err)
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
