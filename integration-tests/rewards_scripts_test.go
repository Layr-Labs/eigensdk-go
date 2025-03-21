package integration_test

import (
	"context"
	"log"
	"math/big"
	"testing"

	strategy "github.com/Layr-Labs/eigensdk-go/contracts/bindings/IStrategy"
	mockerc20 "github.com/Layr-Labs/eigensdk-go/contracts/bindings/MockERC20"
	"github.com/Layr-Labs/eigensdk-go/testutils"
	"github.com/Layr-Labs/eigensdk-go/testutils/testclients"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestIntegrationRewards(t *testing.T) {
	log.Println("This test takes ~50 seconds to run...")

	clients, anvilHttpEndpoint := testclients.BuildTestClients(t)
	contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilHttpEndpoint)

	contractStrategy, err := strategy.NewContractIStrategy(contractAddrs.Erc20MockStrategy, clients.EthHttpClient)
	require.NoError(t, err)

	tokenAddr, err := contractStrategy.UnderlyingToken(&bind.CallOpts{Context: context.Background()})
	require.NoError(t, err)

	mockToken, err := mockerc20.NewContractMockERC20(tokenAddr, clients.EthHttpClient)
	require.NoError(t, err)

	txMgr := clients.TxManager
	noSendTxOpts, err := txMgr.GetNoSendTxOpts()
	require.NoError(t, err)

	// This is the avsRegistry writer address (anvil nineth address)
	tx, err := mockToken.Mint(noSendTxOpts, common.HexToAddress("0xa0Ee7A142d267C1f36714E4a8F75612F20a79720"), big.NewInt(1000))
	require.NoError(t, err)

	receipt, err := txMgr.Send(context.Background(), tx, true)
	require.NoError(t, err)
	require.Equal(t, receipt.Status, uint64(1))
}
