package elcontracts_test

import (
	"context"
	"log/slog"
	"os"
	"testing"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/testutils"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestChainWriter(t *testing.T) {
	anvilStateFileName := "contracts-deployed-anvil-state.json"
	anvilC, err := testutils.StartAnvilContainer(anvilStateFileName)
	require.NoError(t, err)

	anvilHttpEndpoint, err := anvilC.Endpoint(context.Background(), "http")
	require.NoError(t, err)

	defer anvilC.Terminate(context.Background())
	logger := logging.NewTextSLogger(os.Stdout, &logging.SLoggerOptions{Level: slog.LevelDebug})

	ecdsaPrivKeyHex := "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	ecdsaPrivateKey, err := crypto.HexToECDSA(ecdsaPrivKeyHex)
	require.NoError(t, err)

	contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilHttpEndpoint)
	require.NoError(t, err)

	RegistryCoordinatorAddress := contractAddrs.RegistryCoordinator
	OperatorStateRetrieverAddress := contractAddrs.OperatorStateRetriever

	chainioConfig := clients.BuildAllConfig{
		EthHttpUrl:                 "http://localhost:8545",
		EthWsUrl:                   "ws://localhost:8545",
		RegistryCoordinatorAddr:    RegistryCoordinatorAddress.String(),
		OperatorStateRetrieverAddr: OperatorStateRetrieverAddress.String(),
		AvsName:                    "exampleAvs",
		PromMetricsIpPortAddress:   ":9090",
	}

	//// Setup el_chain_reader
	clients, err := clients.BuildAll(
		chainioConfig,
		ecdsaPrivateKey,
		logger,
	)
	require.NoError(t, err)

	// Define an operator
	operator :=
		types.Operator{
			Address:                   "0xd5e099c71b797516c10ed0f0d895f429c2781142",
			DelegationApproverAddress: "0xd5e099c71b797516c10ed0f0d895f429c2781142",
			StakerOptOutWindowBlocks:  100,
			MetadataUrl:               "https://madhur-test-public.s3.us-east-2.amazonaws.com/metadata.json",
		}

	// Test 1: Register as an operator
	receipt, err := clients.ElChainWriter.RegisterAsOperator(context.Background(), operator, true)
	assert.NoError(t, err)
	assert.True(t, receipt.Status == 1)

	// Test 2: Update operator details
	walletModified, err := crypto.HexToECDSA("2a871d0798f97d79848a013d4936a73bf4cc922c825d33c1cf7073dff6d409c6")
	assert.NoError(t, err)
	walletModifiedAddress := crypto.PubkeyToAddress(walletModified.PublicKey)

	operatorModified := types.Operator{
		Address:                   walletModifiedAddress.Hex(),
		DelegationApproverAddress: walletModifiedAddress.Hex(),
		StakerOptOutWindowBlocks:  101,
		MetadataUrl:               "eigensdk-go",
	}

	receipt, err = clients.ElChainWriter.UpdateOperatorDetails(context.Background(), operatorModified, true)
	assert.NoError(t, err)
	assert.True(t, receipt.Status == 1)

	assert.Equal(t, "1234", "1234")
}
