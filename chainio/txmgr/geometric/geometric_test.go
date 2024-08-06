package geometric

import (
	"context"
	"fmt"
	"math/big"
	"sync"
	"testing"
	"time"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/wallet"
	"github.com/Layr-Labs/eigensdk-go/crypto/ecdsa"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/signerv2"
	"github.com/Layr-Labs/eigensdk-go/testutils"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
	"golang.org/x/sync/errgroup"
)

var chainId = big.NewInt(31337)

type testHarness struct {
	fakeEthBackend *fakeEthBackend
	txmgr          *GeometricTxManager
}

func (h *testHarness) validateTxReceipt(t *testing.T, txReceipt *types.Receipt) {
	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	receiptFromEthBackend, err := h.fakeEthBackend.TransactionReceipt(ctxWithTimeout, txReceipt.TxHash)
	require.NoError(t, err)
	require.Equal(t, txReceipt, receiptFromEthBackend)
}

func newTestHarness(t *testing.T) *testHarness {
	logger := testutils.NewTestLogger()
	ethBackend := NewFakeEthBackend()

	ecdsaSk, ecdsaAddr, err := testutils.NewEcdsaSkAndAddress()
	require.NoError(t, err)

	signerFn, _, err := signerv2.SignerFromConfig(signerv2.Config{PrivateKey: ecdsaSk}, chainId)
	require.NoError(t, err)

	skWallet, err := wallet.NewPrivateKeyWallet(ethBackend, signerFn, ecdsaAddr, logger)
	require.NoError(t, err)

	txmgr := NewGeometricTxnManager(ethBackend, skWallet, logger, NewNoopMetrics(), GeometricTxnManagerParams{
		GetTxReceiptTickerDuration: 100 * time.Millisecond,
		// set to 100 so that no buffer is added to the gasTipCap
		// this way we can test that the txmgr will bump the gasTipCap to a working value
		// and also simulate a congested network (with fakeEthBackend.congestedBlocks) where txs won't be mined
		GasTipMultiplierPercentage: 100,
		// set to 1 second (instead of default 2min) so that we can test that the txmgr will bump the gasTipCap to a
		// working value
		TxnBroadcastTimeout: 1 * time.Second,
	})

	return &testHarness{
		fakeEthBackend: ethBackend,
		txmgr:          txmgr,
	}
}

func TestGeometricTxManager(t *testing.T) {
	t.Run("Send 1 tx", func(t *testing.T) {
		h := newTestHarness(t)

		unsignedTx := newUnsignedEthTransferTx(0, nil)
		ctxWithTimeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		txReceipt, err := h.txmgr.Send(ctxWithTimeout, unsignedTx)
		require.NoError(t, err)

		h.validateTxReceipt(t, txReceipt)
	})

	t.Run("Send 1 tx to congested network", func(t *testing.T) {
		h := newTestHarness(t)
		h.fakeEthBackend.setCongestedBlocks(3)
		h.txmgr.params.TxnConfirmationTimeout = 1 * time.Second

		unsignedTx := newUnsignedEthTransferTx(0, nil)
		ctxWithTimeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		txReceipt, err := h.txmgr.Send(ctxWithTimeout, unsignedTx)
		require.NoError(t, err)

		h.validateTxReceipt(t, txReceipt)
	})

	t.Run("gasFeeCap gets overwritten when sending tx", func(t *testing.T) {
		h := newTestHarness(t)

		unsignedTx := newUnsignedEthTransferTx(0, big.NewInt(1))
		ctxWithTimeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		txReceipt, err := h.txmgr.Send(ctxWithTimeout, unsignedTx)
		// ethBackend returns an error if the tx's gasFeeCap is less than the baseFeePerGas
		// this test makes sure that even setting a gasFeeCap less than the baseFeePerGas in the tx (1 above) still
		// works,
		// because the txmgr will overwrite it and set the gasFeeCap to a working value
		require.NoError(t, err)

		h.validateTxReceipt(t, txReceipt)
	})

	t.Run("Send n=3 txs in parallel", func(t *testing.T) {
		n := 3
		h := newTestHarness(t)

		g := new(errgroup.Group)

		txs := make([]*types.Transaction, n)
		txReceipts := make([]*types.Receipt, n)

		for nonce := 0; nonce < n; nonce++ {
			tx := newUnsignedEthTransferTx(uint64(nonce), nil)
			txs[nonce] = tx
			nonce := nonce
			g.Go(func() error {
				ctxWithTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
				defer cancel()
				txReceipt, err := h.txmgr.Send(ctxWithTimeout, tx)
				if err == nil {
					txReceipts[nonce] = txReceipt
				}
				return err
			})
		}

		err := g.Wait()
		require.NoError(t, err)

	})

	t.Run("Send n=3 txs sequentially", func(t *testing.T) {
		n := uint64(3)
		h := newTestHarness(t)

		for nonce := uint64(0); nonce < n; nonce++ {
			tx := newUnsignedEthTransferTx(nonce, nil)
			ctxWithTimeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			_, err := h.txmgr.Send(ctxWithTimeout, tx)
			require.NoError(t, err)
		}

	})

	t.Run("Send 2 txs with same nonce", func(t *testing.T) {
		h := newTestHarness(t)

		unsignedTx := newUnsignedEthTransferTx(0, nil)
		ctxWithTimeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		_, err := h.txmgr.Send(ctxWithTimeout, unsignedTx)
		require.NoError(t, err)

		unsignedTx = newUnsignedEthTransferTx(0, nil)
		_, err = h.txmgr.Send(ctxWithTimeout, unsignedTx)
		require.Error(t, err)
	})
}

func newUnsignedEthTransferTx(nonce uint64, gasFeeCap *big.Int) *types.Transaction {
	if gasFeeCap == nil {
		// 1 gwei is anvil's default starting baseFeePerGas on its genesis block
		gasFeeCap = big.NewInt(1_000_000_000)
	}
	return types.NewTx(&types.DynamicFeeTx{
		ChainID:   chainId,
		Nonce:     nonce,
		GasTipCap: big.NewInt(1),
		GasFeeCap: gasFeeCap,
		Gas:       21000,
		To:        testutils.ZeroAddress(),
		Value:     big.NewInt(1),
	})
}

type fakeEthBackend struct {
	// congestedBlocks can be set to a positive number to start a congestion period.
	// while congestedBlocks > 0, gasTipCap will increase everytime it is read, simulating a congested network
	// this way whenever the txmgr reads the gasTipCap to bump a tx and resend it, the gasTipCap will increase
	// so as to prevent the tx from being mined
	// everytime a block is mined, congestedBlocks will decrease by 1
	congestedBlocks uint64
	gasTipCap       *big.Int
	baseFeePerGas   *big.Int
	// mu protects all the below fields which are updated in "mining" goroutines (see Send)
	mu          sync.Mutex
	blockNumber uint64
	// we use a single nonce for now because the txmgr only supports a single sender
	nonce    uint64
	mempool  map[uint64]*types.Transaction // nonce -> tx
	minedTxs map[common.Hash]*types.Receipt
	logger   logging.Logger
}

func NewFakeEthBackend() *fakeEthBackend {
	logger := testutils.NewTestLogger().With("component", "fakeEthBackend")
	backend := &fakeEthBackend{
		congestedBlocks: 0,
		gasTipCap:       big.NewInt(1),             // 1 wei, same default as anvil
		baseFeePerGas:   big.NewInt(1_000_000_000), // 1 gwei, same default as anvil
		mu:              sync.Mutex{},
		blockNumber:     0,
		nonce:           0,
		mempool:         make(map[uint64]*types.Transaction),
		minedTxs:        make(map[common.Hash]*types.Receipt),
		logger:          logger,
	}
	backend.startMining()
	return backend
}

func (s *fakeEthBackend) WithBaseFeePerGas(baseFeePerGas *big.Int) *fakeEthBackend {
	s.baseFeePerGas = baseFeePerGas
	return s
}

func (s *fakeEthBackend) startMining() {
	go func() {
		for {
			s.mu.Lock()
			// if there's a tx in the mempool with the current nonce and its gasTipCap is >= baseFeePerGas, mine it
			if tx, ok := s.mempool[s.nonce]; ok {
				if tx.GasTipCapIntCmp(s.gasTipCap) >= 0 {
					delete(s.mempool, s.nonce)
					s.minedTxs[tx.Hash()] = &types.Receipt{
						BlockNumber: big.NewInt(int64(s.blockNumber)),
						TxHash:      tx.Hash(),
					}
					s.blockNumber++
					s.nonce++
					s.logger.Debug("mined tx", "txHash", tx.Hash(), "nonce", tx.Nonce())
				} else {
					s.logger.Info("tx.gasTipCap < fakeEthBackend.gasTipCap, not mining", "txHash", tx.Hash(), "nonce", tx.Nonce(), "tx.gasTipCap", tx.GasTipCap(), "fakeEthBackend.gasTipCap", s.gasTipCap)
				}
			}
			if s.congestedBlocks > 0 {
				s.congestedBlocks--
			}
			s.mu.Unlock()
			// mine a block every 100 ms
			time.Sleep(100 * time.Millisecond)
		}
	}()
}

func (s *fakeEthBackend) setCongestedBlocks(n uint64) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.congestedBlocks = n

}

func (s *fakeEthBackend) BlockNumber(context.Context) (uint64, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.blockNumber, nil
}

func (s *fakeEthBackend) SuggestGasTipCap(ctx context.Context) (*big.Int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	gasTipCap := big.NewInt(0).Set(s.gasTipCap)
	if s.congestedBlocks > 0 {
		s.gasTipCap = s.gasTipCap.Add(s.gasTipCap, big.NewInt(1))
	}
	return gasTipCap, nil
}

func (s *fakeEthBackend) EstimateGas(ctx context.Context, msg ethereum.CallMsg) (uint64, error) {
	return 0, nil

}

func (s *fakeEthBackend) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	return &types.Header{
		BaseFee: big.NewInt(0).Set(s.baseFeePerGas),
	}, nil
}

func (s *fakeEthBackend) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	if tx.GasFeeCapIntCmp(s.baseFeePerGas) < 0 {
		return fmt.Errorf("tx.gasFeeCap (%d) < baseFeeCap (%d)", tx.GasFeeCap(), s.baseFeePerGas)
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	if tx.Nonce() < s.nonce {
		return fmt.Errorf("tx.nonce (%d) < current nonce (%d)", tx.Nonce(), s.nonce)
	}
	s.mempool[tx.Nonce()] = tx
	return nil
}

func (s *fakeEthBackend) TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if receipt, ok := s.minedTxs[txHash]; ok {
		return receipt, nil
	}
	return nil, ethereum.NotFound
}

// ========================== INTEGRATION TESTS ==========================

type integrationEthBackend interface {
	wallet.EthBackend
	EthBackend
}
type integrationTestHarness struct {
	ethBackend integrationEthBackend
	txmgr      *GeometricTxManager
}

func newIntegrationTestHarness(t *testing.T) *integrationTestHarness {
	logger := testutils.NewTestLogger()
	anvilC, err := testutils.StartAnvilContainer("")
	require.NoError(t, err)
	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	anvilHttpEndpoint, err := anvilC.Endpoint(ctxWithTimeout, "http")
	require.NoError(t, err)
	anvilHttpClient, err := ethclient.Dial(anvilHttpEndpoint)
	require.NoError(t, err)

	ecdsaSk, ecdsaAddr, err := ecdsa.KeyAndAddressFromHexKey(
		"0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80",
	)
	require.NoError(t, err)

	signerFn, _, err := signerv2.SignerFromConfig(signerv2.Config{PrivateKey: ecdsaSk}, chainId)
	require.NoError(t, err)

	skWallet, err := wallet.NewPrivateKeyWallet(anvilHttpClient, signerFn, ecdsaAddr, logger)
	require.NoError(t, err)

	txmgr := NewGeometricTxnManager(anvilHttpClient, skWallet, logger, NewNoopMetrics(), GeometricTxnManagerParams{})
	return &integrationTestHarness{
		ethBackend: anvilHttpClient,
		txmgr:      txmgr,
	}
}

func (h *integrationTestHarness) validateTxReceipt(t *testing.T, txReceipt *types.Receipt) {
	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	receiptFromEthBackend, err := h.ethBackend.TransactionReceipt(ctxWithTimeout, txReceipt.TxHash)
	require.NoError(t, err)
	require.Equal(t, txReceipt, receiptFromEthBackend)
}

func TestGeometricTxManagerIntegration(t *testing.T) {
	t.Run("Send 1 tx", func(t *testing.T) {
		h := newIntegrationTestHarness(t)

		unsignedTx := newUnsignedEthTransferTx(0, nil)
		ctxWithTimeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		txReceipt, err := h.txmgr.Send(ctxWithTimeout, unsignedTx)
		require.NoError(t, err)

		h.validateTxReceipt(t, txReceipt)
	})
}
