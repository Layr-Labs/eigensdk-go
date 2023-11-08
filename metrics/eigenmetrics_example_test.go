// We make this package metrics_test instead of metrics because the goal of this file is just to
// have an example of how to use the EigenMetrics struct. Given that AVS teams will be importing
// the metrics package to construct their own EigenMetrics struct, we want to mimic that here.
// This way the example shows for eg. `metrics.NewEigenMetrics` instead of just `NewEigenMetrics`.
package metrics_test

import (
	"context"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/chainio/constructor"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/metrics/collectors/economic"
	"github.com/Layr-Labs/eigensdk-go/metrics/collectors/rpc_calls"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/ethereum/go-ethereum/common"
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

	chainioConfig := constructor.Config{
		EcdsaPrivateKeyString:         "0x0",
		EthHttpUrl:                    "http://localhost:8545",
		EthWsUrl:                      "ws://localhost:8545",
		BlsRegistryCoordinatorAddr:    "0x0",
		BlsOperatorStateRetrieverAddr: "0x0",
		AvsName:                       "exampleAvs",
		PromMetricsIpPortAddress:      ":9090",
	}
	clients, err := constructor.BuildClients(chainioConfig, logger)
	if err != nil {
		panic(err)
	}
	reg, eigenMetrics := clients.PrometheusRegistry, clients.Metrics

	operatorAddr := common.HexToAddress("0x0")
	quorumNames := map[types.QuorumNum]string{
		0: "ethQuorum",
		1: "someOtherTokenQuorum",
	}
	// We must register the economic metrics separately because they are exported metrics (from jsonrpc or subgraph calls)
	// and not instrumented metrics: see https://prometheus.io/docs/instrumenting/writing_clientlibs/#overall-structure
	economicMetricsCollector := economic.NewCollector(clients.ElChainReader, clients.AvsRegistryChainReader, "exampleAvs", logger, operatorAddr, quorumNames)
	reg.MustRegister(economicMetricsCollector)

	rpcCallsCollector := rpccalls.NewCollector("exampleAvs", reg)
	instrumentedEthClient, err := eth.NewInstrumentedClient("http://localhost:8545", rpcCallsCollector)
	if err != nil {
		panic(err)
	}

	eigenMetrics.Start(context.Background(), reg)

	// use instrumentedEthClient as you would a normal ethClient
	_ = instrumentedEthClient
}
