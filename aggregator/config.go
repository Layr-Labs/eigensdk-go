package aggregator

import (
	"crypto/ecdsa"

	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type AggregatorConfig struct {
	serverAddress string

	EthHttpUrl string
	EthWsUrl   string

	RegistryCoordinatorAddress    common.Address
	OperatorStateRetrieverAddress common.Address
	ServiceManagerAddress         common.Address

	EthHttpClient *ethclient.Client
	TxMgr         txmgr.TxManager
	Logger        logging.Logger

	EcdsaPrivateKey *ecdsa.PrivateKey
	BlsPrivateKey   *bls.PrivateKey
}
