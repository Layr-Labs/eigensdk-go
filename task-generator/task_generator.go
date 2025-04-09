package taskgenerator

import (
	"context"
	"time"

	"github.com/Layr-Labs/eigensdk-go/logging"
)

type TaskGenerator struct {
	logger          logging.Logger
	logic			TaskGeneratorLogic
	secondsBetweenTasks time.Duration // Maybe this name is confusing, should fine a better way to name it
}

func BuildTaskGenerator(logger logging.Logger, logic TaskGeneratorLogic, timeBetweenTasks time.Duration) (*TaskGenerator, error) {
	return &TaskGenerator{
		logger,
		logic,
		timeBetweenTasks,
	}, nil
}

func (taskGen *TaskGenerator) Start(ctx context.Context) error {
	time.Sleep(time.Duration(2 * time.Second))

	taskGen.logger.Info("Starting Task Generator.")
	taskGen.logger.Info("Starting Task Generator rpc server.")

	ticker := time.NewTicker(taskGen.secondsBetweenTasks * time.Second)
	defer ticker.Stop()
	taskGen.logger.Info("Task Generator set to send new task every %v seconds...", taskGen.secondsBetweenTasks)

	taskNum := int64(0)

	// Send a task before looping
	err := taskGen.logic.sendNewTask(taskNum)
	if err != nil {
		taskGen.logger.Error("Aggregator failed to send number to square", "err", err)
		return err
	}
	taskNum++

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-ticker.C:
			taskGen.logger.Infof("Task Generator sending new task, number to square: %v", taskNum)
			err := taskGen.logic.sendNewTask(taskNum)
			if err != nil {
				taskGen.logger.Error("Aggregator failed to send number to square", "err", err)
				return err
			}
			taskNum++
		}
	}
}

type TaskGeneratorLogic interface{
	sendNewTask(taskNumber int64)(error)
}

/* 
type TaskGenLogic struct{
	avsWriter AvsWriter
	thresholdNumerator uint8
	quorumNumbers uint8
	logger logging.Logger
}

func NewTaskGenLogic(thresholdNumerator uint8, quorumNumbers uint8, logger logging.Logger) (TaskGenLogic, error) {
	avsWriter, err := BuildAvsWriterFromConfig(c)
	if err != nil {
		c.Logger.Errorf("Cannot create avsWriter", "err", err)
		return nil, err
	}

	contractServiceManager, err := csservicemanager.NewContractIncredibleSquaringServiceManager(
		c.IncredibleSquaringServiceManager,
		&c.EthHttpClient,
	)
	if err != nil {
		c.Logger.Error("Failed to fetch IServiceManager contract", "err", err)
		return nil, err
	}
	taskManagerAddr, err := contractServiceManager.IncredibleSquaringTaskManager(&bind.CallOpts{})
	if err != nil {
		c.Logger.Error("Failed to fetch TaskManager address", "err", err)
		return nil, err
	}

	return TaskGenLogic{
		avsWriter, thresholdNumerator, quorumNumbers, logger}, nil
}

func (tgl TaskGenLogic)sendNewTask(taskNumber int64)(error){
	err := tgl.avsWriter.SendNewTaskNumberToSquare(context.Background(), big.NewInt(taskNumber),
	tgl.thresholdNumerator, tgl.quorumNumbers)
	if err != nil {
		tgl.logger.Error("Aggregator failed to send number to square", "err", err)
		return err
	}

	return nil
}

func main(){
	thresholdNumerator := types.QUORUM_THRESHOLD_NUMERATOR
	quorumNumbers := types.QUORUM_NUMBERS

	logic := TaskGenLogic
}

type AvsWriter struct {
	avsregistry.ChainWriter
	logger              logging.Logger
	TxMgr               txmgr.TxManager
}

type AvsConfig struct {
	Logger                    logging.Logger
	EthWsRpcUrl                               string
	EthHttpUrl                    	          string
	OperatorStateRetrieverAddr                common.Address
	IncredibleSquaringRegistryCoordinatorAddr common.Address
	IncredibleSquaringServiceManager          common.Address
	TxMgr                 txmgr.TxManager
}

func BuildAvsWriterFromConfig(c *AvsConfig) (*AvsWriter, error) {
	ethWsClient, err := ethclient.Dial(c.EthWsRpcUrl)
	if err != nil {
		return nil, utils.WrapError("Failed to create Eth WS client", err)
	}
	ethHttpClient, err := ethclient.Dial(c.EthHttpUrl)
	if err != nil {
		return nil, utils.WrapError("Failed to create Eth Http client", err)
	}

	return BuildAvsWriter(
		c.TxMgr,
		c.IncredibleSquaringServiceManager,
		c.IncredibleSquaringRegistryCoordinatorAddr,
		c.OperatorStateRetrieverAddr,
		ethWsClient,
		&ethHttpClient,
		c.Logger,
	)
}

func BuildAvsWriter(
	txMgr txmgr.TxManager,
	serviceManagerAddr common.Address,
	registryCoordinatorAddr, operatorStateRetrieverAddr common.Address,
	wsClient eth.WsBackend,
	ethHttpClient sdkcommon.EthClientInterface,
	logger logging.Logger,
) (*AvsWriter, error) {

	if err != nil {
		logger.Error("Failed to create contract bindings", "err", err)
		return nil, err
	}
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
	return NewAvsWriter(*avsRegistryWriter, logger, txMgr), nil
}

func NewAvsWriter(
	avsRegistryWriter avsregistry.ChainWriter,
	logger logging.Logger,
	txMgr txmgr.TxManager,
) *AvsWriter {
	return &AvsWriter{
		ChainWriter:         avsRegistryWriter,
		logger:              logger,
		TxMgr:               txMgr,
	}
}

// returns the tx receipt, as well as the task index (which it gets from parsing the tx receipt logs)
func (w *AvsWriter) SendNewTaskNumberToSquare(
	ctx context.Context,
	numToSquare *big.Int,
	quorumThresholdPercentage sdktypes.QuorumThresholdPercentage,
	quorumNumbers sdktypes.QuorumNums,
) error {
	txOpts, err := w.TxMgr.GetNoSendTxOpts()
	if err != nil {
		w.logger.Errorf("Error getting tx opts")
		return cstaskmanager.IIncredibleSquaringTaskManagerTask{}, 0, err
	}
	tx, err := w.AvsContractBindings.TaskManager.CreateNewTask(
		txOpts,
		numToSquare,
		uint32(quorumThresholdPercentage),
		quorumNumbers.UnderlyingType(),
	)
	if err != nil {
		w.logger.Errorf("Error assembling CreateNewTask tx")
		return cstaskmanager.IIncredibleSquaringTaskManagerTask{}, 0, err
	}
	receipt, err := w.TxMgr.Send(ctx, tx, true)
	if err != nil {
		w.logger.Errorf("Error submitting CreateNewTask tx")
		return cstaskmanager.IIncredibleSquaringTaskManagerTask{}, 0, err
	}
	newTaskCreatedEvent, err := w.AvsContractBindings.TaskManager.ContractIncredibleSquaringTaskManagerFilterer.ParseNewTaskCreated(
		*receipt.Logs[0],
	)
	if err != nil {
		w.logger.Error("Aggregator failed to parse new task created event", "err", err)
		return cstaskmanager.IIncredibleSquaringTaskManagerTask{}, 0, err
	}
	return newTaskCreatedEvent.Task, newTaskCreatedEvent.TaskIndex, nil
}
 */
