package eth_test

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"testing"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	rpccalls "github.com/Layr-Labs/eigensdk-go/metrics/collectors/rpc_calls"
	"github.com/Layr-Labs/eigensdk-go/testutils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
)

var (
	anvil             testcontainers.Container
	anvilHttpEndpoint string
)

func TestMain(m *testing.M) {
	err := setup()
	if err != nil {
		fmt.Println("Error setting up test environment:", err)
		teardown()
		os.Exit(1)
	}
	exitCode := m.Run()
	teardown()
	os.Exit(exitCode)
}

func setup() error {
	var err error
	anvil, err = testutils.StartAnvilContainer("")
	if err != nil {
		return fmt.Errorf("failed to start Anvil container: %w", err)
	}
	anvilHttpEndpoint, err = anvil.Endpoint(context.Background(), "http")
	if err != nil {
		return fmt.Errorf("failed to get Anvil endpoint: %w", err)
	}
	return nil
}

func teardown() {
	_ = anvil.Terminate(context.Background())
}

func TestNewInstrumentedClient(t *testing.T) {
	reg := prometheus.NewRegistry()
	rpcCallsCollector := rpccalls.NewCollector("exampleAvs", reg)

	client, err := eth.NewInstrumentedClient(anvilHttpEndpoint, rpcCallsCollector)
	require.NoError(t, err)
	assert.NotNil(t, client)
}

func TestChainID(t *testing.T) {
	reg := prometheus.NewRegistry()
	rpcCallsCollector := rpccalls.NewCollector("exampleAvs", reg)

	client, err := eth.NewInstrumentedClient(anvilHttpEndpoint, rpcCallsCollector)
	require.NoError(t, err)

	chainID, err := client.ChainID(context.Background())
	require.NoError(t, err)
	assert.Equal(t, chainID, big.NewInt(31337))
}

func TestBalanceAt(t *testing.T) {
	reg := prometheus.NewRegistry()
	rpcCallsCollector := rpccalls.NewCollector("exampleAvs", reg)

	client, err := eth.NewInstrumentedClient(anvilHttpEndpoint, rpcCallsCollector)
	require.NoError(t, err)

	account := common.HexToAddress("0x0000000000000000000000000000000000000000")
	balance, err := client.BalanceAt(context.Background(), account, nil)
	require.NoError(t, err)
	assert.NotNil(t, balance)
}

func TestBlockByHash(t *testing.T) {
	reg := prometheus.NewRegistry()
	rpcCallsCollector := rpccalls.NewCollector("exampleAvs", reg)

	client, err := eth.NewInstrumentedClient(anvilHttpEndpoint, rpcCallsCollector)
	require.NoError(t, err)

	curBlock, err := client.BlockByNumber(context.Background(), big.NewInt(0))
	require.NoError(t, err)

	block, err := client.BlockByHash(context.Background(), curBlock.Hash())
	require.NoError(t, err)
	assert.NotNil(t, block)
}

func TestBlockByNumber(t *testing.T) {
	reg := prometheus.NewRegistry()
	rpcCallsCollector := rpccalls.NewCollector("exampleAvs", reg)

	client, err := eth.NewInstrumentedClient(anvilHttpEndpoint, rpcCallsCollector)
	require.NoError(t, err)

	number := big.NewInt(0)
	block, err := client.BlockByNumber(context.Background(), number)
	require.NoError(t, err)
	assert.NotNil(t, block)
}

/*
func TestSendTransaction(t *testing.T) {
	reg := prometheus.NewRegistry()
	rpcCallsCollector := rpccalls.NewCollector("exampleAvs", reg)

	client, err := eth.NewInstrumentedClient(anvilHttpEndpoint, rpcCallsCollector)
	require.NoError(t, err)

	tx := types.NewTransaction(
		1,
		common.HexToAddress("0x0000000000000000000000000000000000000000"),
		big.NewInt(1),
		21000,
		big.NewInt(1),
		nil,
	)
	err = client.SendTransaction(context.Background(), tx)
	require.NoError(t, err)
}
*/
