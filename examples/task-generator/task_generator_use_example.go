package taskgeneratorexample

import (
	"context"
	"math/big"
	"time"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/avsregistry"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/wallet"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	cstaskmanager "github.com/Layr-Labs/eigensdk-go/examples/task-generator/bindings/taskManager"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/signerv2"
	taskgenerator "github.com/Layr-Labs/eigensdk-go/task-generator"
	"github.com/Layr-Labs/eigensdk-go/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type TaskGenLogic struct {
	avsWriter          *AvsWriter
	thresholdNumerator uint8
	quorumNumbers      []uint8
	logger             logging.Logger
}

func NewTaskGenLogic(c *AvsConfig, thresholdNumerator uint8, quorumNumbers []uint8) (*TaskGenLogic, error) {
	avsWriter, err := BuildAvsWriterFromConfig(c)
	if err != nil {
		c.Logger.Errorf("Cannot create avsWriter", "err", err)
		return nil, err
	}

	return &TaskGenLogic{
		avsWriter, thresholdNumerator, quorumNumbers, c.Logger}, nil
}

func (tgl *TaskGenLogic) SendNewTask(taskNumber int64) error {
	err := tgl.avsWriter.SendNewTaskNumberToSquare(context.Background(), big.NewInt(taskNumber),
		tgl.thresholdNumerator, tgl.quorumNumbers)
	if err != nil {
		tgl.logger.Error("Aggregator failed to send number to square", "err", err)
		return err
	}

	return nil
}

func main() {
	logger, err := logging.NewZapLogger(logging.Development)
	if err != nil {
		return
	}

	thresholdNumerator := uint8(100)
	quorumNumbers := []uint8{0}

	// This pk should be related to the address passed to TaskManager as task_generator_addr when initialized
	taskgeneratorPk := "0x4bbbf85ce3377467afe5d46f804f221813b2bb87f24d81f60f1fcdbf7cbf4356"
	ecdsaPrivateKey, err := crypto.HexToECDSA(taskgeneratorPk)
	if err != nil {
		return
	}

	ethHttpUrl := "http://localhost:8545"
	ethHttpClient, err := ethclient.Dial(ethHttpUrl)
	if err != nil {
		return
	}

	rpcCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	chainid, err := ethHttpClient.ChainID(rpcCtx)
	if err != nil {
		logger.Error("Cannot get chain id", "err", err)
		return
	}

	signerV2, senderAddr, err := signerv2.SignerFromConfig(signerv2.Config{PrivateKey: ecdsaPrivateKey}, chainid)
	if err != nil {
		return
	}

	pkWallet, err := wallet.NewPrivateKeyWallet(ethHttpClient, signerV2, senderAddr, logger)
	if err != nil {
		return
	}

	txMgr := txmgr.NewSimpleTxManager(pkWallet, ethHttpClient, logger, senderAddr)

	// The values from this config are extracted from an incredible squaring config file and also the deployment output files
	avsConfig := AvsConfig{
		Logger:                     logger,
		EthHttpUrl:                 ethHttpUrl,
		EthWsRpcUrl:                "ws://localhost:8545",
		OperatorStateRetrieverAddr: common.HexToAddress("0x4c5859f0f772848b2d91f1d83e2fe57935348029"),
		IncredibleSquaringRegistryCoordinatorAddr: common.HexToAddress("0x7bc06c482dead17c0e297afbc32f6e63d3846650"),
		IncredibleSquaringServiceManager:          common.HexToAddress("0x5f3f1dbd7b74c6b46e8c44f98792a1daf8d69154"),
		IncredibleSquaringTaskManager:             common.HexToAddress("0x2bdcc0de6be1f7d2ee689a0342d76f52e8efaba3"),
		TxMgr:                                     txMgr,
		EthHttpClient:                             ethHttpClient,
	}

	logic, err := NewTaskGenLogic(&avsConfig, thresholdNumerator, quorumNumbers)
	if err != nil {
		return
	}

	secondsInterval := 10 // This means TaskGenerator will send tasks every 10 seconds
	taskGen, err := taskgenerator.BuildTaskGenerator(logger, logic, secondsInterval)
	if err != nil {
		return
	}

	err = taskGen.Start(context.Background())
	if err != nil {
		return
	}
}

type AvsWriter struct {
	avsregistry.ChainWriter
	logger              logging.Logger
	TxMgr               txmgr.TxManager
	taskManagerContract *cstaskmanager.ContractIncredibleSquaringTaskManager
}

type AvsConfig struct {
	Logger                                    logging.Logger
	EthWsRpcUrl                               string
	EthHttpUrl                                string
	OperatorStateRetrieverAddr                common.Address
	IncredibleSquaringRegistryCoordinatorAddr common.Address
	IncredibleSquaringServiceManager          common.Address
	IncredibleSquaringTaskManager             common.Address
	TxMgr                                     txmgr.TxManager
	EthHttpClient                             *ethclient.Client
}

func BuildAvsWriterFromConfig(c *AvsConfig) (*AvsWriter, error) {
	ethWsClient, err := ethclient.Dial(c.EthWsRpcUrl)
	if err != nil {
		return nil, utils.WrapError("Failed to create Eth WS client", err)
	}

	return BuildAvsWriter(
		c.TxMgr,
		c.IncredibleSquaringServiceManager,
		c.IncredibleSquaringTaskManager,
		c.IncredibleSquaringRegistryCoordinatorAddr,
		c.OperatorStateRetrieverAddr,
		ethWsClient,
		c.EthHttpClient,
		c.Logger,
	)
}

func BuildAvsWriter(
	txMgr txmgr.TxManager,
	serviceManagerAddr, taskManagerAddr common.Address,
	registryCoordinatorAddr, operatorStateRetrieverAddr common.Address,
	wsClient eth.WsBackend,
	ethHttpClient *ethclient.Client,
	logger logging.Logger,
) (*AvsWriter, error) {
	config := avsregistry.Config{
		RegistryCoordinatorAddress:    registryCoordinatorAddr,
		OperatorStateRetrieverAddress: operatorStateRetrieverAddr,
		DontUseAllocationManager:      false,
		ServiceManagerAddress:         serviceManagerAddr,
	}

	_, _, avsRegistryWriter, _, err := avsregistry.BuildClients(config, ethHttpClient, wsClient, txMgr, logger)
	if err != nil {
		return nil, err
	}

	contractTaskManager, err := cstaskmanager.NewContractIncredibleSquaringTaskManager(
		taskManagerAddr,
		ethHttpClient,
	)
	if err != nil {
		return nil, utils.WrapError("Failed to fetch IServiceManager contract", err)
	}

	return NewAvsWriter(*avsRegistryWriter, logger, txMgr, contractTaskManager), nil
}

func NewAvsWriter(
	avsRegistryWriter avsregistry.ChainWriter,
	logger logging.Logger,
	txMgr txmgr.TxManager,
	contractTaskManager *cstaskmanager.ContractIncredibleSquaringTaskManager,
) *AvsWriter {
	return &AvsWriter{
		ChainWriter:         avsRegistryWriter,
		logger:              logger,
		TxMgr:               txMgr,
		taskManagerContract: contractTaskManager,
	}
}

// returns the tx receipt, as well as the task index (which it gets from parsing the tx receipt logs)
func (w *AvsWriter) SendNewTaskNumberToSquare(
	ctx context.Context,
	numToSquare *big.Int,
	quorumThresholdPercentage uint8,
	quorumNumbers []uint8,
) error {
	txOpts, err := w.TxMgr.GetNoSendTxOpts()
	if err != nil {
		w.logger.Errorf("Error getting tx opts")
		return err
	}

	tx, err := w.taskManagerContract.CreateNewTask(
		txOpts,
		numToSquare,
		uint32(quorumThresholdPercentage),
		quorumNumbers,
	)
	if err != nil {
		w.logger.Errorf("Error assembling CreateNewTask tx")
		return err
	}
	receipt, err := w.TxMgr.Send(ctx, tx, true)
	if err != nil {
		w.logger.Errorf("Error submitting CreateNewTask tx")
		return err
	}
	_, err = w.taskManagerContract.ContractIncredibleSquaringTaskManagerFilterer.ParseNewTaskCreated(
		*receipt.Logs[0],
	)
	if err != nil {
		w.logger.Error("Aggregator failed to parse new task created event", "err", err)
		return err
	}
	return nil
}
