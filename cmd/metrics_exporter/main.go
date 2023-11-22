package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/urfave/cli/v2"

	"github.com/Layr-Labs/eigensdk-go/chainio/constructor"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/metrics/collectors/economic"
)

var (
	AvsNameFlag = &cli.StringFlag{
		Name:     "avs-name",
		Usage:    "AVS name",
		Required: true,
	}
	QuorumNamesDictFlag = &cli.StringFlag{
		Name:     "quorum-names-dict",
		Usage:    "quorum names dict (eg: '{0:eth,1:mycoin}')",
		Required: true,
	}
	OperatorAddrFlag = &cli.StringFlag{
		Name:     "operator-addr",
		Usage:    "operator address",
		Required: true,
	}
	RegistryCoordinatorAddrFlag = &cli.StringFlag{
		Name:     "registry-coordinator-addr",
		Usage:    "BLSRegistryCoordinatorWithIndices contract address",
		Required: true,
	}
	MetricsListeningAddrFlag = &cli.StringFlag{
		Name:     "metrics-listening-addr",
		Usage:    "prometheus metrics listening address (eg: localhost:9095)",
		Required: true,
	}
	RpcUrlFlag = &cli.StringFlag{
		Name:        "rpc-url",
		Usage:       "rpc url",
		Value:       "http://localhost:8545",
		DefaultText: "http://localhost:8545",
	}
	LogLevelFlag = &cli.StringFlag{
		Name:        "log-level",
		Usage:       "logging level (production or development)",
		Value:       "production",
		DefaultText: "production",
	}
)

func main() {
	run(os.Args)
}

// We structure run in this way to make it testable.
// see https://github.com/urfave/cli/issues/731
func run(args []string) {
	app := cli.NewApp()
	app.Name = "metrics-exporter"
	app.Usage = "Ran as a AVS node sidecar to export metrics that are fetched from onchain contracts."
	app.Description = "Currently only exports economic metrics."
	app.Action = startCollector
	app.Flags = []cli.Flag{
		AvsNameFlag,
		QuorumNamesDictFlag,
		OperatorAddrFlag,
		RegistryCoordinatorAddrFlag,
		MetricsListeningAddrFlag,
		RpcUrlFlag,
		LogLevelFlag,
	}

	if err := app.Run(args); err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

}

func startCollector(c *cli.Context) error {
	logLevelStr := c.String(LogLevelFlag.Name)
	var logLevel logging.LogLevel
	if logLevelStr == "production" {
		logLevel = logging.Production
	} else if logLevelStr == "development" {
		logLevel = logging.Development
	} else {
		return fmt.Errorf("invalid log level: %s", logLevelStr)
	}
	logger, err := logging.NewZapLogger(logLevel)
	if err != nil {
		return err
	}

	// TODO: we hardcode a bunch of things for now, waiting for chainio refactor PR to be merged
	//      https://github.com/Layr-Labs/eigensdk-go/pull/73
	ethHttpUrl := c.String(RpcUrlFlag.Name)
	// we don't even need this but are forced to pass one because constructor builds the websocket clients
	ethWsUrl := "ws" + ethHttpUrl[4:]
	constructorConfig := constructor.Config{
		// we don't need writers for this exporter, but still need to pass this fake privatekey..
		EcdsaPrivateKeyString:         "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80",
		EthHttpUrl:                    ethHttpUrl,
		EthWsUrl:                      ethWsUrl,
		BlsRegistryCoordinatorAddr:    c.String(RegistryCoordinatorAddrFlag.Name),
		BlsOperatorStateRetrieverAddr: "", // not used by this exporter
		AvsName:                       c.String(AvsNameFlag.Name),
		PromMetricsIpPortAddress:      c.String(MetricsListeningAddrFlag.Name),
	}
	clients, err := constructor.BuildClients(constructorConfig, logger)
	if err != nil {
		return err
	}

	operatorAddr := common.HexToAddress(c.String(OperatorAddrFlag.Name))
	var quorumNamesDict map[uint8]string
	if err := json.Unmarshal([]byte(c.String(QuorumNamesDictFlag.Name)), &quorumNamesDict); err != nil {
		return err
	}
	economicCollector := economic.NewCollector(
		clients.ElChainReader,
		clients.AvsRegistryChainReader,
		c.String(AvsNameFlag.Name),
		logger,
		operatorAddr,
		quorumNamesDict,
	)
	reg := prometheus.NewRegistry()
	reg.MustRegister(economicCollector)

	metricsListeningAddr := c.String(MetricsListeningAddrFlag.Name)
	logger.Infof("Starting metrics server at port %v", metricsListeningAddr)
	http.Handle("/metrics", promhttp.HandlerFor(
		reg,
		promhttp.HandlerOpts{},
	))
	err = http.ListenAndServe(metricsListeningAddr, nil)
	return err
}
