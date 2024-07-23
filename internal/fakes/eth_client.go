package fakes

import (
	"context"
	"errors"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

const (
	TransactionHash          = "0x0000000000000000000000000000000000000000000000000000000000001234"
	TransactionNashNotInFake = "0xabcd"
	BlockNumber              = 1234
)

type EthClient struct {
	chainId       *big.Int
	successfulTxs map[common.Hash]bool
}

// NewEthClient returns a new instance of the EthClient
// Right now this client is hardcoded with some values to satisfy current
// testing requirements, but it can be extended to support more features and
// can be made more generic over time when we add more tests.
// Currently used in
// - chainio/clients/wallet/fireblocks_wallet_test.go
func NewEthClient() *EthClient {
	return &EthClient{
		chainId: big.NewInt(5),
		successfulTxs: map[common.Hash]bool{
			common.HexToHash(TransactionHash): true,
		},
	}
}

func (f *EthClient) ChainID(ctx context.Context) (*big.Int, error) {
	return f.chainId, nil
}

func (f *EthClient) TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	if _, ok := f.successfulTxs[txHash]; !ok {
		return nil, errors.New("tx not found")
	}

	return &types.Receipt{
		TxHash:      txHash,
		BlockNumber: big.NewInt(BlockNumber),
	}, nil
}

func (f *EthClient) CodeAt(ctx context.Context, contract common.Address, blockNumber *big.Int) ([]byte, error) {
	return []byte{}, nil
}

func (f *EthClient) CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	return []byte{}, nil
}

func (f *EthClient) EstimateGas(ctx context.Context, call ethereum.CallMsg) (uint64, error) {
	return 0, nil
}

func (f *EthClient) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	return big.NewInt(0), nil
}

func (f *EthClient) SuggestGasTipCap(ctx context.Context) (*big.Int, error) {
	return big.NewInt(0), nil
}

func (f *EthClient) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	return nil
}

func (f *EthClient) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	return &types.Header{}, nil
}

func (f *EthClient) PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error) {
	return []byte{}, nil
}

func (f *EthClient) PendingNonceAt(ctx context.Context, account common.Address) (uint64, error) {
	return 0, nil
}

func (f *EthClient) FilterLogs(ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error) {
	return []types.Log{}, nil
}

func (f *EthClient) SubscribeFilterLogs(ctx context.Context, q ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error) {
	return nil, nil
}
