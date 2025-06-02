package txmgr_test

import (
	"context"
	"errors"
	"math/big"
	"testing"
	"time"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/wallet"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	"github.com/Layr-Labs/eigensdk-go/signerv2"
	"github.com/Layr-Labs/eigensdk-go/testutils"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newTx() *types.Transaction {
	return types.NewTx(&types.DynamicFeeTx{
		ChainID:   big.NewInt(31337),
		Nonce:     0,
		GasTipCap: big.NewInt(1),
		GasFeeCap: big.NewInt(1_000_000_000),
		Gas:       21000,
		To:        testutils.ZeroAddress(),
		Value:     big.NewInt(1),
	})
}

func TestSendWithRetryWithNoError(t *testing.T) {
	// Test SendWithRetry with a non-failing transaction to verify normal behavior
	ecdsaPrivateKey, err := crypto.HexToECDSA("ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80")
	require.NoError(t, err)
	anvilC, err := testutils.StartAnvilContainer("")
	require.NoError(t, err)
	anvilHttpEndpoint, err := anvilC.Endpoint(context.Background(), "http")
	require.NoError(t, err)
	logger := testutils.NewTestLogger()

	ethHttpClient, err := ethclient.Dial(anvilHttpEndpoint)
	require.NoError(t, err)

	chainid, err := ethHttpClient.ChainID(context.Background())
	require.NoError(t, err)

	signerV2, addr, err := signerv2.SignerFromConfig(signerv2.Config{PrivateKey: ecdsaPrivateKey}, chainid)
	require.NoError(t, err)

	pkWallet, err := wallet.NewPrivateKeyWallet(ethHttpClient, signerV2, addr, logger)
	require.NoError(t, err)

	txMgr := txmgr.NewSimpleTxManager(pkWallet, ethHttpClient, logger, addr)

	tx := newTx()
	retryTimeout := 200 * time.Millisecond
	maxElapsedTime := 2 * time.Second
	multiplier := 1.5

	_, err = txMgr.SendWithRetry(context.Background(), tx, retryTimeout, maxElapsedTime, multiplier)
	require.NoError(t, err)
}

func TestSendWithRetryDoesBackoff(t *testing.T) {
	// Test SendWithRetry using a FailingEthBackend to simulate errors when sending transactions
	logger := testutils.NewTestLogger()
	ethBackend := NewFailingEthBackend(3)

	chainid := big.NewInt(31337)
	ecdsaSk, _, err := testutils.NewEcdsaSkAndAddress()
	require.NoError(t, err)

	signerV2, addr, err := signerv2.SignerFromConfig(signerv2.Config{PrivateKey: ecdsaSk}, chainid)
	require.NoError(t, err)

	pkWallet, err := wallet.NewPrivateKeyWallet(ethBackend, signerV2, addr, logger)
	require.NoError(t, err)

	txMgr := txmgr.NewSimpleTxManager(pkWallet, ethBackend, logger, addr)

	tx := newTx()
	retryTimeout := 200 * time.Millisecond
	maxElapsedTime := 3 * time.Second
	multiplier := 1.5

	_, err = txMgr.SendWithRetry(context.Background(), tx, retryTimeout, maxElapsedTime, multiplier)
	require.NoError(t, err)
	assert.Equal(t, ethBackend.pendingFailures, uint32(0))
}

// Mock of the EthBackend that returns an error when sending transactions.
// Once pendingFailures reaches zero, SendTransaction will no longer fail
type FailingEthBackend struct {
	pendingFailures uint32
}

func NewFailingEthBackend(pendingFailures uint32) *FailingEthBackend {
	backend := &FailingEthBackend{pendingFailures: pendingFailures}
	return backend
}

func (s *FailingEthBackend) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	if s.pendingFailures == 0 {
		return nil
	}
	s.pendingFailures--
	return errors.New("did not send tx")
}

func (s *FailingEthBackend) TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	return &types.Receipt{}, nil
}

func (s *FailingEthBackend) EstimateGas(ctx context.Context, msg ethereum.CallMsg) (uint64, error) {
	return 0, nil
}

func (s *FailingEthBackend) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	return &types.Header{
		BaseFee: big.NewInt(0),
	}, nil
}

func (s *FailingEthBackend) SuggestGasTipCap(ctx context.Context) (*big.Int, error) {
	return big.NewInt(0), nil
}
