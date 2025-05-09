package geometric

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"os"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/wallet"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/signerv2"
	"github.com/Layr-Labs/eigensdk-go/testutils"
	"github.com/Layr-Labs/eigensdk-go/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	chainid = big.NewInt(31337)
)

func createTx(client eth.HttpBackend, address common.Address) (*types.Transaction, error) {
	zeroAddr := common.HexToAddress("0x0")
	nonce, err := client.PendingNonceAt(context.TODO(), address)
	if err != nil {
		return nil, utils.WrapError("Failed to get PendingNonceAt", err)
	}
	return types.NewTx(&types.DynamicFeeTx{
		To:    &zeroAddr,
		Nonce: nonce,
	}), nil
}

func createTxMgr(rpcUrl string, ecdsaPrivateKey *ecdsa.PrivateKey) (eth.HttpBackend, *GeometricTxManager, error) {
	logger := logging.NewTextSLogger(os.Stdout, &logging.SLoggerOptions{})
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		return nil, nil, err
	}
	signerV2, signerAddr, err := signerv2.SignerFromConfig(signerv2.Config{PrivateKey: ecdsaPrivateKey}, chainid)
	if err != nil {
		return nil, nil, err
	}
	wallet, err := wallet.NewPrivateKeyWallet(client, signerV2, signerAddr, logger)
	if err != nil {
		return nil, nil, err
	}
	reg := prometheus.NewRegistry()
	metrics := NewMetrics(reg, "example", logger)
	return client, NewGeometricTxnManager(client, wallet, logger, metrics, GeometricTxnManagerParams{}), nil
}

func ExampleGeometricTxManager() {
	anvilC, err := testutils.StartAnvilContainer("")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	anvilUrl, err := anvilC.Endpoint(context.TODO(), "http")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	ecdsaPrivateKey, err := crypto.HexToECDSA(testutils.ANVIL_FIRST_PRIVATE_KEY)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	pk := ecdsaPrivateKey.PublicKey
	address := crypto.PubkeyToAddress(pk)

	client, txmgr, err := createTxMgr(anvilUrl, ecdsaPrivateKey)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	tx, err := createTx(client, address)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	_, err = txmgr.Send(context.TODO(), tx, true)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// we just add this to make sure the example runs
	fmt.Println("Tx sent")
	// Output: Tx sent
}
