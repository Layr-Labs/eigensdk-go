package signerv2_test

import (
	"context"
	"fmt"
	"log/slog"
	"math/big"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"

	eigenkms "github.com/Layr-Labs/eigensdk-go/aws/kms"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/wallet"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/signerv2"
	"github.com/Layr-Labs/eigensdk-go/testutils"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
	"github.com/ethereum/go-ethereum/common"
	gtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
)

var (
	mappedLocalstackPort string
	keyMetadata          *types.KeyMetadata
	anvilEndpoint        string
	localstack           testcontainers.Container
	anvil                testcontainers.Container
	rpcClient            *rpc.Client
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
	localstack, err = testutils.StartLocalstackContainer("kms_signer_test")
	if err != nil {
		return fmt.Errorf("failed to start Localstack container: %w", err)
	}
	mappedPort, err := localstack.MappedPort(context.Background(), testutils.LocalStackPort)
	if err != nil {
		return fmt.Errorf("failed to get mapped port: %w", err)
	}
	mappedLocalstackPort = string(mappedPort)

	anvil, err = testutils.StartAnvilContainer("")
	if err != nil {
		return fmt.Errorf("failed to start Anvil container: %w", err)
	}
	endpoint, err := anvil.Endpoint(context.Background(), "")
	if err != nil {
		return fmt.Errorf("failed to get Anvil endpoint: %w", err)
	}
	anvilEndpoint = fmt.Sprintf("http://%s", endpoint)
	rpcClient, err = rpc.Dial(anvilEndpoint)
	if err != nil {
		return fmt.Errorf("failed to dial Anvil RPC: %w", err)
	}
	keyMetadata, err = testutils.CreateKMSKey(mappedLocalstackPort)
	if err != nil {
		return fmt.Errorf("failed to create KMS key: %w", err)
	}
	return nil
}

func teardown() {
	_ = localstack.Terminate(context.Background())
	_ = anvil.Terminate(context.Background())
}

func TestSendTransaction(t *testing.T) {
	c, err := testutils.NewKMSClient(mappedLocalstackPort)
	assert.Nil(t, err)
	assert.NotNil(t, keyMetadata.KeyId)
	pk, err := eigenkms.GetECDSAPublicKey(context.Background(), c, *keyMetadata.KeyId)
	assert.Nil(t, err)
	assert.NotNil(t, pk)
	keyAddr := crypto.PubkeyToAddress(*pk)
	t.Logf("Public key address: %s", keyAddr.String())
	assert.NotEqual(t, keyAddr, common.Address{0})
	err = rpcClient.CallContext(context.Background(), nil, "anvil_setBalance", keyAddr, 2_000_000_000_000_000_000)
	assert.Nil(t, err)

	logger := logging.NewTextSLogger(os.Stdout, &logging.SLoggerOptions{Level: slog.LevelDebug})
	ethClient, err := ethclient.Dial(anvilEndpoint)
	assert.Nil(t, err)
	chainID, err := ethClient.ChainID(context.Background())
	assert.Nil(t, err)
	signer := signerv2.NewKMSSigner(context.Background(), c, pk, *keyMetadata.KeyId, chainID)
	assert.Nil(t, err)
	assert.NotNil(t, signer)
	pkWallet, err := wallet.NewPrivateKeyWallet(ethClient, signer, keyAddr, logger)
	assert.Nil(t, err)
	assert.NotNil(t, pkWallet)
	txMgr := txmgr.NewSimpleTxManager(pkWallet, ethClient, logger, keyAddr)
	assert.NotNil(t, txMgr)
	zeroAddr := common.HexToAddress("0x0")
	receipt, err := txMgr.Send(context.Background(), gtypes.NewTx(&gtypes.DynamicFeeTx{
		ChainID: chainID,
		Nonce:   0,
		To:      &zeroAddr,
		Value:   big.NewInt(1_000_000_000_000_000_000),
	}))
	assert.Nil(t, err)
	assert.NotNil(t, receipt)
	balance, err := ethClient.BalanceAt(context.Background(), keyAddr, nil)
	assert.Nil(t, err)
	assert.Equal(t, big.NewInt(999979000000000000), balance)
}
