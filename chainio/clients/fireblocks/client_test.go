package fireblocks_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/fireblocks"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/stretchr/testify/assert"
)

const (
	sandboxAPI    = "https://sandbox-api.fireblocks.io"
	secretKeyPath = "FILL_ME_IN"
	apiKey        = "FILL_ME_IN"
)

func newFireblocksClient(t *testing.T) fireblocks.Client {
	secretKey, err := os.ReadFile(secretKeyPath)
	assert.NoError(t, err)
	logger, err := logging.NewZapLogger(logging.Development)
	assert.NoError(t, err)
	c, err := fireblocks.NewClient(apiKey, secretKey, sandboxAPI, 5*time.Second, logger)
	assert.NoError(t, err)
	return c

}

func TestListContracts(t *testing.T) {
	t.Skip("skipping test as it's meant for manual runs only")

	c := newFireblocksClient(t)
	contracts, err := c.ListContracts(context.Background())
	assert.NoError(t, err)
	for _, contract := range contracts {
		t.Logf("Contract: %+v", contract)
	}
}

func TestContractCall(t *testing.T) {
	t.Skip("skipping test as it's meant for manual runs only")

	c := newFireblocksClient(t)
	destinationAccountID := "FILL_ME_IN"
	idempotenceKey := "FILL_ME_IN"
	// Tests the contract call against this contract:
	// https://holesky.etherscan.io/address/0x2c61EA360D6500b58E7f481541A36B443Bc858c6
	req := fireblocks.NewContractCallRequest(
		idempotenceKey,
		"ETH_TEST6",
		"2",
		destinationAccountID,
		"0",
		"0x5140a548000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000030000000000000000000000002177dee1f66d6dbfbf517d9c4f316024c6a21aeb000000000000000000000000ad9770d6b5514724c7b766f087bea8a784038cbe000000000000000000000000cb14cfaac122e52024232583e7354589aede74ff00000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000",
		"", // replaceTxByHash
		"", // gasPrice
		"", // gasLimit
		"", // maxFee
		"", // priorityFee
		fireblocks.FeeLevelHigh,
	)
	resp, err := c.ContractCall(context.Background(), req)
	assert.NoError(t, err)
	t.Logf("txID: %s, status: %s", resp.ID, resp.Status)
}

func TestListVaultAccounts(t *testing.T) {
	t.Skip("skipping test as it's meant for manual runs only")

	c := newFireblocksClient(t)
	accounts, err := c.ListVaultAccounts(context.Background())
	assert.NoError(t, err)
	for _, account := range accounts {
		t.Logf("Account: %+v", account)
	}
}

func TestGetTransaction(t *testing.T) {
	t.Skip("skipping test as it's meant for manual runs only")

	c := newFireblocksClient(t)
	txID := "FILL_ME_IN"
	tx, err := c.GetTransaction(context.Background(), txID)
	assert.NoError(t, err)
	t.Logf("Transaction: %+v", tx)
}

func TestGetAssetAddresses(t *testing.T) {
	t.Skip("skipping test as it's meant for manual runs only")

	c := newFireblocksClient(t)
	vaultID := "FILL_ME_IN"
	assetID := fireblocks.AssetID("FILL_ME_IN")
	addresses, err := c.GetAssetAddresses(context.Background(), vaultID, assetID)
	assert.NoError(t, err)
	for _, address := range addresses {
		t.Logf("Address: %+v", address)
	}
}
