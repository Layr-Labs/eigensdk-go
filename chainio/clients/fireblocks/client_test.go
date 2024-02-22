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

func TestListContracts(t *testing.T) {
	t.Skip("skipping test as it's meant for manual runs only")

	secretKey, err := os.ReadFile(secretKeyPath)
	assert.NoError(t, err)
	logger, err := logging.NewZapLogger(logging.Development)
	assert.NoError(t, err)
	c, err := fireblocks.NewFireblocksClient(apiKey, secretKey, sandboxAPI, 5*time.Second, logger)
	assert.NoError(t, err)
	contracts, err := c.ListContracts(context.Background())
	assert.NoError(t, err)
	for _, contract := range contracts {
		t.Logf("Contract: %+v", contract)
	}
}

func TestContractCall(t *testing.T) {
	t.Skip("skipping test as it's meant for manual runs only")

	secretKey, err := os.ReadFile(secretKeyPath)
	assert.NoError(t, err)

	logger, err := logging.NewZapLogger(logging.Development)
	assert.NoError(t, err)

	c, err := fireblocks.NewFireblocksClient(apiKey, secretKey, sandboxAPI, 5*time.Second, logger)
	assert.NoError(t, err)
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
	)
	resp, err := c.ContractCall(context.Background(), req)
	assert.NoError(t, err)
	t.Logf("txID: %s, status: %s", resp.ID, resp.Status)
}
