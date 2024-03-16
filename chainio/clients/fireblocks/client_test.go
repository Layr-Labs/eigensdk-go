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
	// https://goerli.etherscan.io/address/0x5f9ef6e1bb2acb8f592a483052b732ceb78e58ca
	req := fireblocks.NewContractCallRequest(
		idempotenceKey,
		"ETH_TEST3",
		"1",
		destinationAccountID,
		"0",
		"0x6057361d00000000000000000000000000000000000000000000000000000000000f4240",
		"",
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
