package elcontracts_test

import (
	"context"
	"log/slog"
	"os"
	"testing"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients"
	erc20 "github.com/Layr-Labs/eigensdk-go/contracts/bindings/IERC20"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/testutils"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestChainReader(t *testing.T) {
	anvilStateFileName := "contracts-deployed-anvil-state.json"
	anvilC, err := testutils.StartAnvilContainer(anvilStateFileName)
	require.NoError(t, err)

	anvilHttpEndpoint, err := anvilC.Endpoint(context.Background(), "http")
	require.NoError(t, err)

	anvilWsEndpoint, err := anvilC.Endpoint(context.Background(), "ws")
	require.NoError(t, err)
	logger := logging.NewTextSLogger(os.Stdout, &logging.SLoggerOptions{Level: slog.LevelDebug})

	privateKeyHex := "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	ecdsaPrivateKey, err := crypto.HexToECDSA(privateKeyHex)
	require.NoError(t, err)

	contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilHttpEndpoint)
	require.NoError(t, err)

	chainioConfig := clients.BuildAllConfig{
		EthHttpUrl:                 anvilHttpEndpoint,
		EthWsUrl:                   anvilWsEndpoint,
		RegistryCoordinatorAddr:    contractAddrs.RegistryCoordinator.String(),
		OperatorStateRetrieverAddr: contractAddrs.OperatorStateRetriever.String(),
		AvsName:                    "exampleAvs",
		PromMetricsIpPortAddress:   ":9090",
	}

	clients, err := clients.BuildAll(
		chainioConfig,
		ecdsaPrivateKey,
		logger,
	)
	require.NoError(t, err)

	operatorAddress := "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"
	operator := types.Operator{
		Address:                   operatorAddress,
		DelegationApproverAddress: operatorAddress,
		StakerOptOutWindowBlocks:  101,
		MetadataUrl:               "test",
	}

	t.Run("is operator registered", func(t *testing.T) {
		isOperator, err := clients.ElChainReader.IsOperatorRegistered(&bind.CallOpts{}, operator)
		assert.NoError(t, err)
		assert.Equal(t, isOperator, true)
	})

	t.Run("get strategy and underlying token", func(t *testing.T) {
		strategyAddr := contractAddrs.Erc20MockStrategy
		strategy, underlyingTokenAddr, err := clients.ElChainReader.GetStrategyAndUnderlyingToken(&bind.CallOpts{}, strategyAddr)
		assert.NoError(t, err)
		assert.NotNil(t, strategy)
		assert.NotEqual(t, common.Address{}, underlyingTokenAddr)

		// Verify that it's a valid ERC20 token
		erc20Token, err := erc20.NewContractIERC20(underlyingTokenAddr, clients.EthHttpClient)
		assert.NoError(t, err)

		tokenName, err := erc20Token.Name(&bind.CallOpts{})
		assert.NoError(t, err)
		assert.NotEmpty(t, tokenName)
	})
}
