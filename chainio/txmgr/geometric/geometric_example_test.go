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
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	chainid = big.NewInt(31337)
)

func ExampleGeometricTxManager() {
	anvilC, err := testutils.StartAnvilContainer("")
	if err != nil {
		panic(err)
	}
	anvilUrl, err := anvilC.Endpoint(context.TODO(), "http")
	if err != nil {
		panic(err)
	}

	ecdsaPrivateKey, err := crypto.HexToECDSA("ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80")
	if err != nil {
		panic(err)
	}
	pk := ecdsaPrivateKey.PublicKey
	address := crypto.PubkeyToAddress(pk)

	client, txmgr := createTxMgr(anvilUrl, ecdsaPrivateKey)

	tx := createTx(client, address)
	_, err = txmgr.Send(context.TODO(), tx)
	if err != nil {
		panic(err)
	}

	// we just add this to make sure the example runs
	fmt.Println("Tx sent")
	// Output: Tx sent
}

func createTx(client eth.Client, address common.Address) *types.Transaction {
	zeroAddr := common.HexToAddress("0x0")
	nonce, err := client.PendingNonceAt(context.TODO(), address)
	if err != nil {
		panic(err)
	}
	return types.NewTx(&types.DynamicFeeTx{
		To:    &zeroAddr,
		Nonce: nonce,
	})
}

func createTxMgr(rpcUrl string, ecdsaPrivateKey *ecdsa.PrivateKey) (eth.Client, *GeometricTxManager) {
	logger := logging.NewTextSLogger(os.Stdout, &logging.SLoggerOptions{})
	client, err := eth.NewClient(rpcUrl)
	if err != nil {
		panic(err)
	}
	signerV2, signerAddr, err := signerv2.SignerFromConfig(signerv2.Config{PrivateKey: ecdsaPrivateKey}, chainid)
	if err != nil {
		panic(err)
	}
	wallet, err := wallet.NewPrivateKeyWallet(client, signerV2, signerAddr, logger)
	if err != nil {
		panic(err)
	}
	reg := prometheus.NewRegistry()
	metrics := NewMetrics(reg, "example", logger)
	return client, NewGeometricTxnManager(client, wallet, logger, metrics, GeometricTxnManagerParams{})
}
