// We make this package metrics_test instead of metrics because the goal of this file is just to
// have an example of how to use the EigenMetrics struct. Given that AVS teams will be importing
// the metrics package to construct their own EigenMetrics struct, we want to mimic that here.
// This way the example shows for eg. `metrics.NewEigenMetrics` instead of just `NewEigenMetrics`.
package metrics_test

import (
	"context"

	"github.com/Layr-Labs/eigensdk-go/chainio/avsregistry"
	sdkclients "github.com/Layr-Labs/eigensdk-go/chainio/clients"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/chainio/elcontracts"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/metrics"
	"github.com/Layr-Labs/eigensdk-go/metrics/collectors/economic"
	rpccalls "github.com/Layr-Labs/eigensdk-go/metrics/collectors/rpc_calls"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/prometheus/client_golang/prometheus"
)

// ExampleEigenMetrics is a testable example (https://go.dev/blog/examples), so tests skip it.
// It's purpose is to show up on godocs to give example of how to construct the and work with the EigenMetrics struct.
// It needs to be in a _test so that it's hidden from normal compilation, but still tested in tests so that
// we are forced to update it and keep it in sync with the code when we make refactors or code changes.
func ExampleEigenMetrics() {

	logger, err := logging.NewZapLogger("development")
	if err != nil {
		panic(err)
	}

	reg := prometheus.NewRegistry()
	eigenMetrics := metrics.NewEigenMetrics("exampleAvs", ":9090", reg, logger)

	slasherAddr := common.HexToAddress("0x0")
	blsPubKeyCompendiumAddr := common.HexToAddress("0x0")
	ethHttpClient, err := eth.NewClient("http://localhost:8545")
	if err != nil {
		panic(err)
	}
	ethWsClient, err := eth.NewClient("ws://localhost:8545")
	if err != nil {
		panic(err)
	}
	elContractsClient, err := sdkclients.NewELContractsChainClient(slasherAddr, blsPubKeyCompendiumAddr, ethHttpClient, ethWsClient, logger)
	if err != nil {
		panic(err)
	}
	eigenlayerReader, err := elcontracts.NewELChainReader(elContractsClient, logger, ethHttpClient)
	if err != nil {
		panic(err)
	}

	blsRegistryCoordAddr := common.HexToAddress("0x0")
	blsOperatorStateRetrieverAddr := common.HexToAddress("0x0")
	stakeRegistry := common.HexToAddress("0x0")
	blsPubkeyRegistryAddr := common.HexToAddress("0x0")
	avsRegistryClients, err := sdkclients.NewAvsRegistryContractsChainClient(
		blsRegistryCoordAddr, blsOperatorStateRetrieverAddr, stakeRegistry, blsPubkeyRegistryAddr, ethHttpClient, logger,
	)
	if err != nil {
		panic(err)
	}
	avsRegistryReader, err := avsregistry.NewAvsRegistryReader(avsRegistryClients, logger, ethHttpClient)
	if err != nil {
		panic(err)
	}

	operatorAddr := common.HexToAddress("0x0")
	quorumNames := map[types.QuorumNum]string{
		0: "ethQuorum",
		1: "someOtherTokenQuorum",
	}
	// We must register the economic metrics separately because they are exported metrics (from jsonrpc or subgraph calls)
	// and not instrumented metrics: see https://prometheus.io/docs/instrumenting/writing_clientlibs/#overall-structure
	economicMetricsCollector := economic.NewCollector(eigenlayerReader, avsRegistryReader, "exampleAvs", logger, operatorAddr, quorumNames)
	reg.MustRegister(economicMetricsCollector)

	rpcCallsCollector := rpccalls.NewCollector("eigen", "exampleAvs", reg)
	eth.NewInstrumentedClient("http://localhost:8545", rpcCallsCollector)

	eigenMetrics.Start(context.Background(), reg)
}
