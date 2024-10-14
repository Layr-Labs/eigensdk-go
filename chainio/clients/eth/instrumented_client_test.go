package eth_test

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	rpccalls "github.com/Layr-Labs/eigensdk-go/metrics/collectors/rpc_calls"
	"github.com/Layr-Labs/eigensdk-go/testutils"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
)

var (
	anvil             testcontainers.Container
	anvilHttpEndpoint string
	anvilWsEndpoint   string
	rpcCallsCollector *rpccalls.Collector
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
	anvilWsEndpoint, err = anvil.Endpoint(context.Background(), "ws")
	if err != nil {
		return fmt.Errorf("failed to get Anvil endpoint: %w", err)
	}
	reg := prometheus.NewRegistry()
	rpcCallsCollector = rpccalls.NewCollector("exampleAvs", reg)
	return nil
}

func teardown() {
	_ = anvil.Terminate(context.Background())
}

func TestNewInstrumentedClient(t *testing.T) {
	client, err := eth.NewInstrumentedClient(anvilHttpEndpoint, rpcCallsCollector)
	assert.NoError(t, err)
	assert.NotNil(t, client)
}

func TestChainID(t *testing.T) {
	client, err := eth.NewInstrumentedClient(anvilHttpEndpoint, rpcCallsCollector)
	assert.NoError(t, err)

	chainID, err := client.ChainID(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, chainID, big.NewInt(31337))
}

func TestBalanceAt(t *testing.T) {
	client, err := eth.NewInstrumentedClient(anvilHttpEndpoint, rpcCallsCollector)
	assert.NoError(t, err)

	account := common.HexToAddress("0x0000000000000000000000000000000000000000")
	balance, err := client.BalanceAt(context.Background(), account, nil)
	assert.NoError(t, err)
	assert.Equal(t, balance.Uint64(), uint64(0))
}

func TestBlockByHash(t *testing.T) {
	client, err := eth.NewInstrumentedClient(anvilHttpEndpoint, rpcCallsCollector)
	assert.NoError(t, err)

	curBlock, err := client.BlockByNumber(context.Background(), nil)
	assert.NoError(t, err)

	block, err := client.BlockByHash(context.Background(), curBlock.Hash())
	assert.NoError(t, err)
	assert.NotNil(t, block)
}

func TestBlockByNumber(t *testing.T) {
	client, err := eth.NewInstrumentedClient(anvilHttpEndpoint, rpcCallsCollector)
	assert.NoError(t, err)

	number := big.NewInt(0)
	block, err := client.BlockByNumber(context.Background(), number)
	assert.NoError(t, err)
	assert.Equal(t, block.NumberU64(), uint64(0))
}

func TestBlockNumber(t *testing.T) {
	client, err := eth.NewInstrumentedClient(anvilHttpEndpoint, rpcCallsCollector)
	assert.NoError(t, err)

	number, err := client.BlockNumber(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, number, uint64(0))
}

func TestCallContract(t *testing.T) {
	client, err := eth.NewInstrumentedClient(anvilHttpEndpoint, rpcCallsCollector)
	assert.NoError(t, err)

	to := common.HexToAddress("0x0")
	callMsg := ethereum.CallMsg{
		To:    &to,
		Value: big.NewInt(0),
	}
	result, err := client.CallContract(context.Background(), callMsg, nil)
	assert.NoError(t, err)
	assert.Equal(t, result, []byte{})
}

func TestCodeAt(t *testing.T) {
	client, err := eth.NewInstrumentedClient(anvilHttpEndpoint, rpcCallsCollector)
	assert.NoError(t, err)

	contract := common.HexToAddress("0x0")
	result, err := client.CodeAt(context.Background(), contract, nil)
	assert.NoError(t, err)
	assert.Equal(t, result, []byte{})
}

func TestEstimateGas(t *testing.T) {
	client, err := eth.NewInstrumentedClient(anvilHttpEndpoint, rpcCallsCollector)
	assert.NoError(t, err)

	to := common.HexToAddress("0x0")
	callMsg := ethereum.CallMsg{
		To: &to,
	}
	result, err := client.EstimateGas(context.Background(), callMsg)
	assert.NoError(t, err)
	assert.NotEqual(t, result, uint64(0))
}

func TestFeeHistory(t *testing.T) {
	client, err := eth.NewInstrumentedClient(anvilHttpEndpoint, rpcCallsCollector)
	assert.NoError(t, err)

	rewardsPercentiles := []float64{0.2}
	feeHistory, err := client.FeeHistory(context.Background(), 1, nil, rewardsPercentiles)
	assert.NoError(t, err)
	assert.Equal(t, feeHistory.OldestBlock.Uint64(), uint64(0))
	assert.Equal(t, feeHistory.Reward[0][0].Uint64(), uint64(0))
	assert.Equal(t, feeHistory.BaseFee[0].Uint64(), uint64(0))
	assert.Equal(t, feeHistory.GasUsedRatio, []float64{0})
}

func TestFilterLogs(t *testing.T) {
	client, err := eth.NewInstrumentedClient(anvilHttpEndpoint, rpcCallsCollector)
	assert.NoError(t, err)

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(0),
		ToBlock:   big.NewInt(10000),
		Addresses: []common.Address{common.HexToAddress("0x0")},
	}
	logs, err := client.FilterLogs(context.Background(), query)
	assert.NoError(t, err)
	assert.Empty(t, logs)
}

func TestHeaderByHash(t *testing.T) {
	client, err := eth.NewInstrumentedClient(anvilHttpEndpoint, rpcCallsCollector)
	assert.NoError(t, err)

	curBlock, err := client.BlockByNumber(context.Background(), nil)
	assert.NoError(t, err)

	header, err := client.HeaderByHash(context.Background(), curBlock.Hash())
	assert.NoError(t, err)

	headerBytes, err := header.MarshalJSON()
	assert.NoError(t, err)
	expectedHeaderBytes, err := curBlock.Header().MarshalJSON()
	assert.NoError(t, err)

	assert.Equal(t, headerBytes, expectedHeaderBytes)
}

func TestHeaderByNumber(t *testing.T) {
	client, err := eth.NewInstrumentedClient(anvilHttpEndpoint, rpcCallsCollector)
	assert.NoError(t, err)

	curBlockNumber, err := client.BlockNumber(context.Background())
	assert.NoError(t, err)

	header, err := client.HeaderByNumber(context.Background(), nil)
	assert.NoError(t, err)

	assert.Equal(t, header.Number.Uint64(), curBlockNumber)
}

func TestNonceAt(t *testing.T) {
	client, err := eth.NewInstrumentedClient(anvilHttpEndpoint, rpcCallsCollector)
	assert.NoError(t, err)

	account := common.HexToAddress("0x0000000000000000000000000000000000000000")
	nonce, err := client.NonceAt(context.Background(), account, nil)
	assert.NoError(t, err)

	assert.Equal(t, nonce, uint64(0))
}

func TestPendingBalanceAt(t *testing.T) {
	client, err := eth.NewInstrumentedClient(anvilHttpEndpoint, rpcCallsCollector)
	assert.NoError(t, err)

	account := common.HexToAddress("0x0")
	balance, err := client.PendingBalanceAt(context.Background(), account)
	assert.NoError(t, err)

	assert.Equal(t, balance.Uint64(), uint64(0))
}

func TestPendingCallContract(t *testing.T) {
	client, err := eth.NewInstrumentedClient(anvilHttpEndpoint, rpcCallsCollector)
	assert.NoError(t, err)

	to := common.HexToAddress("0x0")
	callMsg := ethereum.CallMsg{
		To: &to,
	}
	result, err := client.PendingCallContract(context.Background(), callMsg)
	assert.NoError(t, err)
	assert.Equal(t, result, []byte{})
}

func TestPendingCodeAt(t *testing.T) {
	client, err := eth.NewInstrumentedClient(anvilHttpEndpoint, rpcCallsCollector)
	assert.NoError(t, err)

	account := common.HexToAddress("0x0")
	code, err := client.PendingCodeAt(context.Background(), account)
	assert.NoError(t, err)
	assert.Equal(t, code, []byte{})
}

func TestPendingNonceAt(t *testing.T) {
	client, err := eth.NewInstrumentedClient(anvilHttpEndpoint, rpcCallsCollector)
	assert.NoError(t, err)

	account := common.HexToAddress("0x0")
	nonce, err := client.PendingNonceAt(context.Background(), account)
	assert.NoError(t, err)
	assert.Equal(t, nonce, uint64(0))
}

func TestPendingStorageAt(t *testing.T) {
	client, err := eth.NewInstrumentedClient(anvilHttpEndpoint, rpcCallsCollector)
	assert.NoError(t, err)

	account := common.HexToAddress("0x0")
	key := common.HexToHash("0x123")
	storage, err := client.PendingStorageAt(context.Background(), account, key)
	assert.NoError(t, err)
	assert.Equal(t, storage, make([]byte, 32))
}

func TestPendingTransactionCount(t *testing.T) {
	client, err := eth.NewInstrumentedClient(anvilHttpEndpoint, rpcCallsCollector)
	assert.NoError(t, err)

	count, err := client.PendingTransactionCount(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, count, uint(0))
}

func TestStorageAt(t *testing.T) {
	client, err := eth.NewInstrumentedClient(anvilHttpEndpoint, rpcCallsCollector)
	assert.NoError(t, err)

	account := common.HexToAddress("0x0")
	key := common.HexToHash("0x123")
	storage, err := client.StorageAt(context.Background(), account, key, nil)
	assert.NoError(t, err)
	assert.Equal(t, storage, make([]byte, 32))
}

func TestSubscribeFilterLogs(t *testing.T) {
	client, err := eth.NewInstrumentedClient(anvilWsEndpoint, rpcCallsCollector)
	assert.NoError(t, err)

	query := ethereum.FilterQuery{
		Addresses: []common.Address{common.HexToAddress("0x0")},
	}
	ch := make(chan types.Log)

	subscription, err := client.SubscribeFilterLogs(context.Background(), query, ch)
	assert.NoError(t, err)
	assert.NotNil(t, subscription)
}

func TestSubscribeNewHead(t *testing.T) {
	client, err := eth.NewInstrumentedClient(anvilWsEndpoint, rpcCallsCollector)
	assert.NoError(t, err)

	ch := make(chan *types.Header)

	subscription, err := client.SubscribeNewHead(context.Background(), ch)
	assert.NoError(t, err)
	assert.NotNil(t, subscription)
}

func TestSuggestGasPrice(t *testing.T) {
	client, err := eth.NewInstrumentedClient(anvilHttpEndpoint, rpcCallsCollector)
	assert.NoError(t, err)

	gasPrice, err := client.SuggestGasPrice(context.Background())
	assert.NoError(t, err)
	assert.NotEqual(t, gasPrice.Uint64(), uint64(0))
}

func TestSuggestGasTipCap(t *testing.T) {
	client, err := eth.NewInstrumentedClient(anvilHttpEndpoint, rpcCallsCollector)
	assert.NoError(t, err)

	gasTipCap, err := client.SuggestGasTipCap(context.Background())
	assert.NoError(t, err)
	assert.NotEqual(t, gasTipCap.Uint64(), uint64(0))
}

func TestSyncProgress(t *testing.T) {
	client, err := eth.NewInstrumentedClient(anvilHttpEndpoint, rpcCallsCollector)
	assert.NoError(t, err)

	syncProgress, err := client.SyncProgress(context.Background())
	assert.NoError(t, err)
	assert.Nil(t, syncProgress) // is nil since there is no current syncing in place
}

func TestTransactionMethods(t *testing.T) {
	client, err := eth.NewInstrumentedClient(anvilHttpEndpoint, rpcCallsCollector)
	assert.NoError(t, err)

	to := common.HexToAddress("0x123")
	tx := types.NewTx(&types.DynamicFeeTx{
		Value: big.NewInt(1),
		To:    &to,
		Gas:   21000,
	})

	ecdsaPrivKeyHex := "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	ecdsaPrivKey, err := crypto.HexToECDSA(ecdsaPrivKeyHex)
	assert.NoError(t, err)
	signer := types.LatestSignerForChainID(big.NewInt(31337))
	signature, err := crypto.Sign(signer.Hash(tx).Bytes(), ecdsaPrivKey)
	assert.NoError(t, err)
	signedTx, err := tx.WithSignature(signer, signature)
	assert.NoError(t, err)

	err = client.SendTransaction(context.Background(), signedTx)
	assert.NoError(t, err)

	// this sleep is needed because we have to wait since the transaction is not ready yet
	time.Sleep(2 * time.Second)

	t.Run("transaction by hash", func(t *testing.T) {
		_, _, err = client.TransactionByHash(context.Background(), signedTx.Hash())
		assert.NoError(t, err)
	})

	curBlock, err := client.BlockByNumber(context.Background(), nil)
	assert.NoError(t, err)
	blockHash := curBlock.Hash()

	t.Run("transaction count", func(t *testing.T) {
		count, err := client.TransactionCount(context.Background(), blockHash)
		assert.NoError(t, err)
		assert.Equal(t, count, uint(1))
	})

	t.Run("transaction in block", func(t *testing.T) {
		txIndex := uint(0)

		tx, err := client.TransactionInBlock(context.Background(), blockHash, txIndex)
		assert.NoError(t, err)
		assert.Equal(t, tx.Hash(), signedTx.Hash())
	})
}

func TestTransactionCount(t *testing.T) {
	client, err := eth.NewInstrumentedClient(anvilHttpEndpoint, rpcCallsCollector)
	assert.NoError(t, err)

	curBlock, err := client.BlockByNumber(context.Background(), nil)
	assert.NoError(t, err)
	blockHash := curBlock.Hash()

	count, err := client.TransactionCount(context.Background(), blockHash)
	assert.NoError(t, err)
	assert.Equal(t, count, uint(0))
}
