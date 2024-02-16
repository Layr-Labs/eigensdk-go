package eth

import (
	"context"
	"math/big"
	"time"

	rpccalls "github.com/Layr-Labs/eigensdk-go/metrics/collectors/rpc_calls"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// InstrumentedClient is a wrapper around the geth ethclient that instruments
// all the calls made to it. It counts each eth_ call made to it, and records the duration of each call,
// and exposes these as prometheus metrics
//
// TODO: ideally this should be done at the rpcclient level, not the ethclient level, which would make
// our life much easier... but geth implemented the gethclient using an rpcClient struct instead of interface...
// see https://github.com/ethereum/go-ethereum/issues/28267
type InstrumentedClient struct {
	client            *ethclient.Client
	rpcCallsCollector *rpccalls.Collector
	// we store both client and version because that's what the web3_clientVersion jsonrpc call returns
	// https://ethereum.org/en/developers/docs/apis/json-rpc/#web3_clientversion
	clientAndVersion string
}

var _ Client = (*InstrumentedClient)(nil)

func NewInstrumentedClient(rpcAddress string, rpcCallsCollector *rpccalls.Collector) (*InstrumentedClient, error) {
	client, err := ethclient.Dial(rpcAddress)
	if err != nil {
		return nil, err
	}
	clientAndVersion := getClientAndVersion(client)
	return &InstrumentedClient{
		client:            client,
		rpcCallsCollector: rpcCallsCollector,
		clientAndVersion:  clientAndVersion,
	}, nil
}

func NewInstrumentedClientFromClient(client *ethclient.Client, rpcCallsCollector *rpccalls.Collector) *InstrumentedClient {
	clientAndVersion := getClientAndVersion(client)
	return &InstrumentedClient{
		client:            client,
		rpcCallsCollector: rpcCallsCollector,
		clientAndVersion:  clientAndVersion,
	}
}

// gethClient interface methods

func (iec *InstrumentedClient) ChainID(ctx context.Context) (*big.Int, error) {
	chainID := func() (*big.Int, error) { return iec.client.ChainID(ctx) }
	id, err := instrumentFunction[*big.Int](chainID, "eth_chainId", iec)
	return id, err
}

func (iec *InstrumentedClient) BalanceAt(
	ctx context.Context,
	account common.Address,
	blockNumber *big.Int,
) (*big.Int, error) {
	balanceAt := func() (*big.Int, error) { return iec.client.BalanceAt(ctx, account, blockNumber) }
	balance, err := instrumentFunction[*big.Int](balanceAt, "eth_getBalance", iec)
	if err != nil {
		return nil, err
	}
	return balance, nil
}

func (iec *InstrumentedClient) BlockByHash(ctx context.Context, hash common.Hash) (*types.Block, error) {
	blockByHash := func() (*types.Block, error) { return iec.client.BlockByHash(ctx, hash) }
	block, err := instrumentFunction[*types.Block](blockByHash, "eth_getBlockByHash", iec)
	if err != nil {
		return nil, err
	}
	return block, nil
}

func (iec *InstrumentedClient) BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error) {
	blockByNumber := func() (*types.Block, error) { return iec.client.BlockByNumber(ctx, number) }
	block, err := instrumentFunction[*types.Block](
		blockByNumber,
		"eth_getBlockByNumber",
		iec,
	)
	if err != nil {
		return nil, err
	}
	return block, nil
}

func (iec *InstrumentedClient) BlockNumber(ctx context.Context) (uint64, error) {
	blockNumber := func() (uint64, error) { return iec.client.BlockNumber(ctx) }
	number, err := instrumentFunction[uint64](blockNumber, "eth_blockNumber", iec)
	if err != nil {
		return 0, err
	}
	return number, nil
}

func (iec *InstrumentedClient) CallContract(
	ctx context.Context,
	call ethereum.CallMsg,
	blockNumber *big.Int,
) ([]byte, error) {
	callContract := func() ([]byte, error) { return iec.client.CallContract(ctx, call, blockNumber) }
	bytes, err := instrumentFunction[[]byte](callContract, "eth_call", iec)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (iec *InstrumentedClient) CallContractAtHash(
	ctx context.Context,
	msg ethereum.CallMsg,
	blockHash common.Hash,
) ([]byte, error) {
	callContractAtHash := func() ([]byte, error) { return iec.client.CallContractAtHash(ctx, msg, blockHash) }
	bytes, err := instrumentFunction[[]byte](callContractAtHash, "eth_call", iec)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (iec *InstrumentedClient) CodeAt(
	ctx context.Context,
	contract common.Address,
	blockNumber *big.Int,
) ([]byte, error) {
	call := func() ([]byte, error) { return iec.client.CodeAt(ctx, contract, blockNumber) }
	bytes, err := instrumentFunction[[]byte](call, "eth_getCode", iec)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (iec *InstrumentedClient) EstimateGas(ctx context.Context, call ethereum.CallMsg) (uint64, error) {
	estimateGas := func() (uint64, error) { return iec.client.EstimateGas(ctx, call) }
	gas, err := instrumentFunction[uint64](estimateGas, "eth_estimateGas", iec)
	if err != nil {
		return 0, err
	}
	return gas, nil
}

func (iec *InstrumentedClient) FeeHistory(
	ctx context.Context,
	blockCount uint64,
	lastBlock *big.Int,
	rewardPercentiles []float64,
) (*ethereum.FeeHistory, error) {
	feeHistory := func() (*ethereum.FeeHistory, error) {
		return iec.client.FeeHistory(ctx, blockCount, lastBlock, rewardPercentiles)
	}
	history, err := instrumentFunction[*ethereum.FeeHistory](
		feeHistory,
		"eth_feeHistory",
		iec,
	)
	if err != nil {
		return nil, err
	}
	return history, nil
}

func (iec *InstrumentedClient) FilterLogs(ctx context.Context, query ethereum.FilterQuery) ([]types.Log, error) {
	filterLogs := func() ([]types.Log, error) { return iec.client.FilterLogs(ctx, query) }
	logs, err := instrumentFunction[[]types.Log](filterLogs, "eth_getLogs", iec)
	if err != nil {
		return nil, err
	}
	return logs, nil
}

func (iec *InstrumentedClient) HeaderByHash(ctx context.Context, hash common.Hash) (*types.Header, error) {
	headerByHash := func() (*types.Header, error) { return iec.client.HeaderByHash(ctx, hash) }
	header, err := instrumentFunction[*types.Header](
		headerByHash,
		"eth_getBlockByHash",
		iec,
	)
	if err != nil {
		return nil, err
	}
	return header, nil
}

func (iec *InstrumentedClient) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	headerByNumber := func() (*types.Header, error) { return iec.client.HeaderByNumber(ctx, number) }
	header, err := instrumentFunction[*types.Header](
		headerByNumber,
		"eth_getBlockByNumber",
		iec,
	)
	if err != nil {
		return nil, err
	}
	return header, nil
}

func (iec *InstrumentedClient) NetworkID(ctx context.Context) (*big.Int, error) {
	networkID := func() (*big.Int, error) { return iec.client.NetworkID(ctx) }
	id, err := instrumentFunction[*big.Int](networkID, "net_version", iec)
	if err != nil {
		return nil, err
	}
	return id, nil
}

func (iec *InstrumentedClient) NonceAt(
	ctx context.Context,
	account common.Address,
	blockNumber *big.Int,
) (uint64, error) {
	nonceAt := func() (uint64, error) { return iec.client.NonceAt(ctx, account, blockNumber) }
	nonce, err := instrumentFunction[uint64](nonceAt, "eth_getTransactionCount", iec)
	if err != nil {
		return 0, err
	}
	return nonce, nil
}

func (iec *InstrumentedClient) PeerCount(ctx context.Context) (uint64, error) {
	peerCount := func() (uint64, error) { return iec.client.PeerCount(ctx) }
	count, err := instrumentFunction[uint64](peerCount, "net_peerCount", iec)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (iec *InstrumentedClient) PendingBalanceAt(ctx context.Context, account common.Address) (*big.Int, error) {
	pendingBalanceAt := func() (*big.Int, error) { return iec.client.PendingBalanceAt(ctx, account) }
	balance, err := instrumentFunction[*big.Int](pendingBalanceAt, "eth_getBalance", iec)
	if err != nil {
		return nil, err
	}
	return balance, nil
}

func (iec *InstrumentedClient) PendingCallContract(ctx context.Context, call ethereum.CallMsg) ([]byte, error) {
	pendingCallContract := func() ([]byte, error) { return iec.client.PendingCallContract(ctx, call) }
	bytes, err := instrumentFunction[[]byte](pendingCallContract, "eth_call", iec)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (iec *InstrumentedClient) PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error) {
	pendingCodeAt := func() ([]byte, error) { return iec.client.PendingCodeAt(ctx, account) }
	bytes, err := instrumentFunction[[]byte](pendingCodeAt, "eth_getCode", iec)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (iec *InstrumentedClient) PendingNonceAt(ctx context.Context, account common.Address) (uint64, error) {
	pendingNonceAt := func() (uint64, error) { return iec.client.PendingNonceAt(ctx, account) }
	nonce, err := instrumentFunction[uint64](
		pendingNonceAt,
		"eth_getTransactionCount",
		iec,
	)
	if err != nil {
		return 0, err
	}
	return nonce, nil
}

func (iec *InstrumentedClient) PendingStorageAt(
	ctx context.Context,
	account common.Address,
	key common.Hash,
) ([]byte, error) {
	pendingStorageAt := func() ([]byte, error) { return iec.client.PendingStorageAt(ctx, account, key) }
	bytes, err := instrumentFunction[[]byte](pendingStorageAt, "eth_getStorageAt", iec)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (iec *InstrumentedClient) PendingTransactionCount(ctx context.Context) (uint, error) {
	pendingTransactionCount := func() (uint, error) { return iec.client.PendingTransactionCount(ctx) }
	count, err := instrumentFunction[uint](
		pendingTransactionCount,
		"eth_getBlockTransactionCountByNumber",
		iec,
	)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (iec *InstrumentedClient) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	// instrumentFunction takes a function that returns a value and an error
	// so we just wrap the SendTransaction method in a function that returns 0 as its value,
	// which we throw out below
	sendTransaction := func() (int, error) { return 0, iec.client.SendTransaction(ctx, tx) }
	_, err := instrumentFunction[int](sendTransaction, "eth_sendRawTransaction", iec)
	return err
}

func (iec *InstrumentedClient) StorageAt(
	ctx context.Context,
	account common.Address,
	key common.Hash,
	blockNumber *big.Int,
) ([]byte, error) {
	storageAt := func() ([]byte, error) { return iec.client.StorageAt(ctx, account, key, blockNumber) }
	bytes, err := instrumentFunction[[]byte](storageAt, "eth_getStorageAt", iec)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (iec *InstrumentedClient) SubscribeFilterLogs(
	ctx context.Context,
	query ethereum.FilterQuery,
	ch chan<- types.Log,
) (ethereum.Subscription, error) {
	subscribeFilterLogs := func() (ethereum.Subscription, error) { return iec.client.SubscribeFilterLogs(ctx, query, ch) }
	subscription, err := instrumentFunction[ethereum.Subscription](
		subscribeFilterLogs,
		"eth_subscribe",
		iec,
	)
	if err != nil {
		return nil, err
	}
	return subscription, nil
}

func (iec *InstrumentedClient) SubscribeNewHead(
	ctx context.Context,
	ch chan<- *types.Header,
) (ethereum.Subscription, error) {
	subscribeNewHead := func() (ethereum.Subscription, error) { return iec.client.SubscribeNewHead(ctx, ch) }
	subscription, err := instrumentFunction[ethereum.Subscription](
		subscribeNewHead,
		"eth_subscribe",
		iec,
	)
	if err != nil {
		return nil, err
	}
	return subscription, nil
}

func (iec *InstrumentedClient) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	suggestGasPrice := func() (*big.Int, error) { return iec.client.SuggestGasPrice(ctx) }
	gasPrice, err := instrumentFunction[*big.Int](suggestGasPrice, "eth_gasPrice", iec)
	if err != nil {
		return nil, err
	}
	return gasPrice, nil
}

func (iec *InstrumentedClient) SuggestGasTipCap(ctx context.Context) (*big.Int, error) {
	suggestGasTipCap := func() (*big.Int, error) { return iec.client.SuggestGasTipCap(ctx) }
	gasTipCap, err := instrumentFunction[*big.Int](
		suggestGasTipCap,
		"eth_maxPriorityFeePerGas",
		iec,
	)
	if err != nil {
		return nil, err
	}
	return gasTipCap, nil
}

func (iec *InstrumentedClient) SyncProgress(ctx context.Context) (*ethereum.SyncProgress, error) {
	syncProgress := func() (*ethereum.SyncProgress, error) { return iec.client.SyncProgress(ctx) }
	progress, err := instrumentFunction[*ethereum.SyncProgress](
		syncProgress,
		"eth_syncing",
		iec,
	)
	if err != nil {
		return nil, err
	}
	return progress, nil
}

// We write the instrumentation of this function directly because instrumentFunction[] generic fct only takes a single
// return value
func (iec *InstrumentedClient) TransactionByHash(
	ctx context.Context,
	hash common.Hash,
) (tx *types.Transaction, isPending bool, err error) {
	start := time.Now()
	tx, isPending, err = iec.client.TransactionByHash(ctx, hash)
	// we count both successful and erroring calls (even though this is not well defined in the spec)
	iec.rpcCallsCollector.AddRPCRequestTotal("eth_getTransactionByHash", iec.clientAndVersion)
	if err != nil {
		return nil, false, err
	}
	rpcRequestDuration := time.Since(start)
	// we only observe the duration of successful calls (even though this is not well defined in the spec)
	iec.rpcCallsCollector.ObserveRPCRequestDurationSeconds(
		float64(rpcRequestDuration),
		"eth_getTransactionByHash",
		iec.clientAndVersion,
	)

	return tx, isPending, nil
}

func (iec *InstrumentedClient) TransactionCount(ctx context.Context, blockHash common.Hash) (uint, error) {
	transactionCount := func() (uint, error) { return iec.client.TransactionCount(ctx, blockHash) }
	count, err := instrumentFunction[uint](
		transactionCount,
		"eth_getBlockTransactionCountByHash",
		iec,
	)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (iec *InstrumentedClient) TransactionInBlock(
	ctx context.Context,
	blockHash common.Hash,
	index uint,
) (*types.Transaction, error) {
	transactionInBlock := func() (*types.Transaction, error) { return iec.client.TransactionInBlock(ctx, blockHash, index) }
	tx, err := instrumentFunction[*types.Transaction](
		transactionInBlock,
		"eth_getTransactionByBlockHashAndIndex",
		iec,
	)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func (iec *InstrumentedClient) TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	transactionReceipt := func() (*types.Receipt, error) { return iec.client.TransactionReceipt(ctx, txHash) }
	receipt, err := instrumentFunction[*types.Receipt](
		transactionReceipt,
		"eth_getTransactionReceipt",
		iec,
	)
	if err != nil {
		return nil, err
	}
	return receipt, nil
}

func (iec *InstrumentedClient) TransactionSender(
	ctx context.Context,
	tx *types.Transaction,
	block common.Hash,
	index uint,
) (common.Address, error) {
	transactionSender := func() (common.Address, error) { return iec.client.TransactionSender(ctx, tx, block, index) }
	address, err := instrumentFunction[common.Address](
		transactionSender,
		"eth_getSender",
		iec,
	)
	if err != nil {
		return common.Address{}, err
	}
	return address, nil
}

// Extra methods

// TODO(samlaf): feels weird that we have to write this function ourselves
//
//	perhaps the gethClient interface should be the instrumented client,
//	and then ethclient can take an instrumentedGethClient?
func (iec *InstrumentedClient) WaitForTransactionReceipt(ctx context.Context, txHash common.Hash) *types.Receipt {
	for {
		// verifying transaction receipt
		receipt, err := iec.TransactionReceipt(ctx, txHash)
		if err != nil {
			time.Sleep(2 * time.Second)
		} else {
			return receipt
		}
	}
}

// Not sure why this method is not exposed in the ethclient itself...
// but it is needed to comply with the rpc metrics defined in avs-node spec
// https://eigen.nethermind.io/docs/metrics/metrics-prom-spec
func getClientAndVersion(client *ethclient.Client) string {
	var clientVersion string
	err := client.Client().Call(&clientVersion, "web3_clientVersion")
	if err != nil {
		return "unavailable"
	}
	return clientVersion
}

// Generic function used to instrument all the eth calls that we make below
func instrumentFunction[T any](
	rpcCall func() (T, error),
	rpcMethodName string,
	iec *InstrumentedClient,
) (value T, err error) {
	start := time.Now()
	result, err := rpcCall()
	// we count both successful and erroring calls (even though this is not well defined in the spec)
	iec.rpcCallsCollector.AddRPCRequestTotal(rpcMethodName, iec.clientAndVersion)
	if err != nil {
		return value, err
	}
	rpcRequestDuration := time.Since(start)
	// we only observe the duration of successful calls (even though this is not well defined in the spec)
	iec.rpcCallsCollector.ObserveRPCRequestDurationSeconds(
		float64(rpcRequestDuration),
		rpcMethodName,
		iec.clientAndVersion,
	)
	return result, nil
}
