package avsregistry

import (
	"context"
	"log/slog"
	"os"
	"testing"
	"time"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/wallet"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/signerv2"
	"github.com/Layr-Labs/eigensdk-go/testutils"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
)

func TestWriterMethods(t *testing.T) {
	anvilStateFileName := "contracts-deployed-anvil-state.json"
	anvilC, err := testutils.StartAnvilContainer(anvilStateFileName)
	require.NoError(t, err)
	anvilHttpEndpoint, err := anvilC.Endpoint(context.Background(), "http")
	require.NoError(t, err)

	contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilHttpEndpoint)
	ethHttpClient, err := ethclient.Dial(anvilHttpEndpoint)
	require.NoError(t, err)
	logger := logging.NewTextSLogger(os.Stdout, &logging.SLoggerOptions{Level: slog.LevelDebug})

	ecdsaPrivateKey, err := crypto.HexToECDSA("0x0")
	require.NoError(t, err)

	rpcCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	chainid, err := ethHttpClient.ChainID(rpcCtx)
	require.NoError(t, err)

	signerV2, addr, err := signerv2.SignerFromConfig(signerv2.Config{PrivateKey: ecdsaPrivateKey}, chainid)
	require.NoError(t, err)

	pkWallet, err := wallet.NewPrivateKeyWallet(ethHttpClient, signerV2, addr, logger)
	require.NoError(t, err)
	txMgr := txmgr.NewSimpleTxManager(pkWallet, ethHttpClient, logger, addr)

	chainWriter, err := NewWriterFromConfig(Config{
		RegistryCoordinatorAddress:    contractAddrs.RegistryCoordinator,
		OperatorStateRetrieverAddress: contractAddrs.OperatorStateRetriever,
	}, ethHttpClient, txMgr, logger)
	require.NoError(t, err)
}
