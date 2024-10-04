package eth_test

import (
	"context"
	"testing"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	rpccalls "github.com/Layr-Labs/eigensdk-go/metrics/collectors/rpc_calls"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewInstrumentedClient(t *testing.T) {
	reg := prometheus.NewRegistry()
	rpcCallsCollector := rpccalls.NewCollector("exampleAvs", reg)
	rpcAddress := "http://localhost:8545"

	client, err := eth.NewInstrumentedClient(rpcAddress, rpcCallsCollector)
	require.NoError(t, err)
	assert.NotNil(t, client)
}

func TestChainID(t *testing.T) {
	reg := prometheus.NewRegistry()
	rpcCallsCollector := rpccalls.NewCollector("exampleAvs", reg)
	rpcAddress := "http://localhost:8545"

	client, err := eth.NewInstrumentedClient(rpcAddress, rpcCallsCollector)
	require.NoError(t, err)

	chainID, err := client.ChainID(context.Background())
	require.NoError(t, err)
	assert.NotNil(t, chainID)
}

/*
func TestBalanceAt(t *testing.T) {
	rpcAddress := "http://localhost:8545"
	rpcCallsCollector := rpccalls.NewCollector()

	client, err := eth.NewInstrumentedClient(rpcAddress, rpcCallsCollector)
	require.NoError(t, err)

	account := common.HexToAddress("0x0000000000000000000000000000000000000000")
	balance, err := client.BalanceAt(context.Background(), account, nil)
	require.NoError(t, err)
	assert.NotNil(t, balance)
}

func TestBlockByHash(t *testing.T) {
	rpcAddress := "http://localhost:8545"
	rpcCallsCollector := rpccalls.NewCollector()

	client, err := eth.NewInstrumentedClient(rpcAddress, rpcCallsCollector)
	require.NoError(t, err)

	hash := common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000")
	block, err := client.BlockByHash(context.Background(), hash)
	require.NoError(t, err)
	assert.NotNil(t, block)
}

func TestBlockByNumber(t *testing.T) {
	rpcAddress := "http://localhost:8545"
	rpcCallsCollector := rpccalls.NewCollector()

	client, err := eth.NewInstrumentedClient(rpcAddress, rpcCallsCollector)
	require.NoError(t, err)

	number := big.NewInt(1)
	block, err := client.BlockByNumber(context.Background(), number)
	require.NoError(t, err)
	assert.NotNil(t, block)
}

func TestSendTransaction(t *testing.T) {
	rpcAddress := "http://localhost:8545"
	rpcCallsCollector := rpccalls.NewCollector()

	client, err := eth.NewInstrumentedClient(rpcAddress, rpcCallsCollector)
	require.NoError(t, err)

	tx := types.NewTransaction(1, common.HexToAddress("0x0000000000000000000000000000000000000000"), big.NewInt(1), 21000, big.NewInt(1), nil)
	err = client.SendTransaction(context.Background(), tx)
	require.NoError(t, err)
}
*/
