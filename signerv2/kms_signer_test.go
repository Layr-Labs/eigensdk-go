package signerv2_test

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"testing"

	eigenkms "github.com/Layr-Labs/eigensdk-go/aws/kms"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/wallet"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	"github.com/Layr-Labs/eigensdk-go/signerv2"
	"github.com/Layr-Labs/eigensdk-go/testutils"
	"github.com/Layr-Labs/eigensdk-go/utils"

	"github.com/aws/aws-sdk-go-v2/service/kms/types"

	"github.com/ethereum/go-ethereum/common"
	gtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
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
		return utils.WrapError("failed to start Localstack container", err)
	}
	mappedPort, err := localstack.MappedPort(context.Background(), testutils.LocalStackPort)
	if err != nil {
		return utils.WrapError("failed to get mapped port", err)
	}
	mappedLocalstackPort = string(mappedPort)

	anvil, err = testutils.StartAnvilContainer("")
	if err != nil {
		return utils.WrapError("failed to start Anvil container", err)
	}
	endpoint, err := anvil.Endpoint(context.Background(), "")
	if err != nil {
		return utils.WrapError("failed to get Anvil endpoint", err)
	}
	anvilEndpoint = fmt.Sprintf("http://%s", endpoint)
	rpcClient, err = rpc.Dial(anvilEndpoint)
	if err != nil {
		return utils.WrapError("failed to dial Anvil RPC", err)
	}
	keyMetadata, err = testutils.CreateKMSKey(mappedLocalstackPort)
	if err != nil {
		return utils.WrapError("failed to create KMS key", err)
	}
	return nil
}

func teardown() {
	_ = localstack.Terminate(context.Background())
	_ = anvil.Terminate(context.Background())
}

func TestSignTransactionWithKmsSigner(t *testing.T) {
	logger := testutils.GetTestLogger()
	ethClient, err := ethclient.Dial(anvilEndpoint)
	assert.Nil(t, err)
	chainID, err := ethClient.ChainID(context.Background())
	assert.Nil(t, err)
	zeroAddr := common.HexToAddress("0x0")

	// read input from JSON if available, otherwise use default values
	var defaultInput = struct {
		ChainID big.Int        `json:"chain_id"`
		Nonce   uint64         `json:"nonce"`
		To      common.Address `json:"to"`
		Value   big.Int        `json:"value"`
	}{
		ChainID: *big.NewInt(0),
		Nonce:   0,
		To:      zeroAddr,
		Value:   *big.NewInt(1_000_000_000_000_000_000),
	}
	testData := testutils.NewTestData(defaultInput)

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

	signer := signerv2.NewKMSSigner(context.Background(), c, pk, *keyMetadata.KeyId, chainID)
	assert.Nil(t, err)
	assert.NotNil(t, signer)
	pkWallet, err := wallet.NewPrivateKeyWallet(ethClient, signer, keyAddr, logger)
	assert.Nil(t, err)
	assert.NotNil(t, pkWallet)
	txMgr := txmgr.NewSimpleTxManager(pkWallet, ethClient, logger, keyAddr)
	assert.NotNil(t, txMgr)
	receipt, err := txMgr.Send(context.Background(), gtypes.NewTx(&gtypes.DynamicFeeTx{
		ChainID: &testData.Input.ChainID,
		Nonce:   testData.Input.Nonce,
		To:      &zeroAddr,
		Value:   &testData.Input.Value,
	}), true)
	assert.Nil(t, err)
	assert.NotNil(t, receipt)
	balance, err := ethClient.BalanceAt(context.Background(), keyAddr, nil)
	assert.Nil(t, err)
	assert.Equal(t, big.NewInt(999979000000000000), balance)
}
