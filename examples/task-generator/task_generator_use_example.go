package taskgeneratorexample

import (
	"context"
	"math/big"
	"time"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/avsregistry"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	cstaskmanager "github.com/Layr-Labs/eigensdk-go/examples/task-generator/bindings/taskManager"
	"github.com/Layr-Labs/eigensdk-go/logging"
	taskgenerator "github.com/Layr-Labs/eigensdk-go/task-generator"
	"github.com/Layr-Labs/eigensdk-go/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type TaskGenLogic struct{
	avsWriter *AvsWriter
	thresholdNumerator uint8
	quorumNumbers []uint8
	logger logging.Logger
}

func NewTaskGenLogic(c *AvsConfig, thresholdNumerator uint8, quorumNumbers []uint8, logger logging.Logger) (*TaskGenLogic, error) {
	avsWriter, err := BuildAvsWriterFromConfig(c)
	if err != nil {
		c.Logger.Errorf("Cannot create avsWriter", "err", err)
		return nil, err
	}

	return &TaskGenLogic{
		avsWriter, thresholdNumerator, quorumNumbers, logger}, nil
}

func (tgl *TaskGenLogic) SendNewTask(taskNumber int64)(error){
	err := tgl.avsWriter.SendNewTaskNumberToSquare(context.Background(), big.NewInt(taskNumber),
	tgl.thresholdNumerator, tgl.quorumNumbers)
	if err != nil {
		tgl.logger.Error("Aggregator failed to send number to square", "err", err)
		return err
	}

	return nil
}

func main(){
	thresholdNumerator := uint8(100)
	quorumNumbers := []uint8{0}

	logger, err := logging.NewZapLogger(logging.Development)
	if err != nil {
		return
	}

	avsConfig := AvsConfig{

	}
	logic, err := NewTaskGenLogic(&avsConfig, thresholdNumerator, quorumNumbers, logger)
	if err != nil {
		return
	}

	durationInterval := time.Duration(10) // This means TaskGenerator will send tasks every 10 seconds
	taskGen, err := taskgenerator.BuildTaskGenerator(logger, logic, durationInterval)
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
	Logger                    logging.Logger
	EthWsRpcUrl                               string
	EthHttpUrl                    	          string
	OperatorStateRetrieverAddr                common.Address
	IncredibleSquaringRegistryCoordinatorAddr common.Address
	IncredibleSquaringServiceManager          common.Address
	IncredibleSquaringTaskManager          common.Address
	TxMgr                 txmgr.TxManager
	EthHttpClient *ethclient.Client
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
