package operator

import (
	"context"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	sdkaggregator "github.com/Layr-Labs/eigensdk-go/aggregator"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/avsregistry"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	sdklogging "github.com/Layr-Labs/eigensdk-go/logging"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
	"github.com/Layr-Labs/eigensdk-go/utils"
)

const AVS_NAME = "incredible-squaring"
const SEM_VER = "0.0.1"

type OperatorConfig struct {
	Production                       bool   `yaml:"production"`
	OperatorAddress                  string `yaml:"operator_address"`
	OperatorStateRetrieverAddress    string `yaml:"operator_state_retriever_address"`
	IncredibleSquaringServiceManager string `yaml:"service_manager_address"`
	InstantSlasher                   string `yaml:"instant_slasher_address"`
	AVSRegistryCoordinatorAddress    string `yaml:"avs_registry_coordinator_address"`
	RewardsCoordinatorAddress        string `yaml:"rewards_coordinator_address"`
	PermissionControllerAddress      string `yaml:"permission_controller_address"`
	AllocationManagerAddress         string `yaml:"allocation_manager_address"`
	TokenStrategyAddr                string `yaml:"token_strategy_addr"`
	EthRpcUrl                        string `yaml:"eth_rpc_url"`
	EthWsUrl                         string `yaml:"eth_ws_url"`
	BlsPrivateKeyStorePath           string `yaml:"bls_private_key_store_path"`
	EcdsaPrivateKeyStorePath         string `yaml:"ecdsa_private_key_store_path"`
	AggregatorServerIpPortAddress    string `yaml:"aggregator_server_ip_port_address"`
	RegisterOperatorOnStartup        bool   `yaml:"register_operator_on_startup"`
	EigenMetricsIpPortAddress        string `yaml:"eigen_metrics_ip_port_address"`
	EnableMetrics                    bool   `yaml:"enable_metrics"`
	NodeApiIpPortAddress             string `yaml:"node_api_ip_port_address"`
	EnableNodeApi                    bool   `yaml:"enable_node_api"`
	OperatorSetId                    uint32 `yaml:"operator_set_id"`
	Socket                           string `yaml:"socket"`
	MaxOperatorCount                 uint32 `yaml:"max_operator_count"`
	KickBIPsOfOperatorStake          uint16 `yaml:"kick_bips_of_operator_stake"`
	KickBIPsOfTotalStake             uint16 `yaml:"kick_bips_of_total_stake"`
	MinimumStake                     int64  `yaml:"minimum_stake"`
	Multiplier                       int64  `yaml:"multiplier"`
	TimesFailing                     int    `yaml:"times_failing"`
}

type OperatorTaskProcessor interface {
	ProcessNewTaskCreatedLog(newTaskCreated types.Log) (sdkaggregator.TaskResponse)
}

type Operator struct {
//	config    OperatorConfig
	logger    sdklogging.Logger
	operatorId       sdktypes.OperatorId
	aggregatorRpcClient AggregatorRpcClienter
	EthWsUrl                         string
	blsKeypair       *bls.KeyPair
	taskProcessor OperatorTaskProcessor

	//nodeApi          *nodeapi.NodeApi
	//avsWriter        *avsregistry.ChainWriter
	//avsReader        *avsregistry.ChainReader
	//avsSubscriber    *avsregistry.ChainSubscriber
	//eigenlayerReader *elcontracts.ChainReader
	//eigenlayerWriter *elcontracts.ChainWriter
	//operatorAddr     common.Address
	// receive new tasks in this chan (typically from listening to onchain event)
	// newTaskCreatedChan chan *cstaskmanager.ContractIncredibleSquaringTaskManagerNewTaskCreated
	// ip address of aggregator
	//aggregatorServerIpPortAddr string
	// needed when opting in to avs (allow this service manager contract to slash operator)
	//credibleSquaringServiceManagerAddr common.Address

	// If bigger than zero, submits wrong responses that many times every 100
	timesFailing int

	newTaskCreatedLogs    chan types.Log

}

// TODO(samlaf): config is a mess right now, since the chainio client constructors
//
//	take the config in core (which is shared with aggregator and challenger)
func NewOperatorFromConfig(c OperatorConfig, eventHash common.Hash) (*Operator, error) {
	var logLevel sdklogging.LogLevel
	if c.Production {
		logLevel = sdklogging.Production
	} else {
		logLevel = sdklogging.Development
	}
	logger, err := sdklogging.NewZapLogger(logLevel)
	if err != nil {
		return nil, err
	}

	avs_config := avsregistry.Config{
		RegistryCoordinatorAddress: common.HexToAddress(c.AVSRegistryCoordinatorAddress),
		OperatorStateRetrieverAddress: common.HexToAddress(c.OperatorStateRetrieverAddress),
		ServiceManagerAddress: common.HexToAddress(c.IncredibleSquaringServiceManager),
	}

	ethHttpClient, err := ethclient.Dial(c.EthRpcUrl)
	if err != nil {
		return nil, utils.WrapError("Failed to create Eth Http client", err)
	}

	avsReader, err := avsregistry.NewReaderFromConfig(avs_config, ethHttpClient, logger)
	if err != nil {
		logger.Error("Cannot create AvsReader", "err", err)
		return nil, err
	}

	// Check if operator was registered, return error if not
	operatorIsRegistered, err := avsReader.IsOperatorRegistered(&bind.CallOpts{}, common.HexToAddress(c.OperatorAddress))
	if err != nil {
		logger.Error("Error checking if operator is registered", "err", err)
		return nil, err
	}
	if !operatorIsRegistered {
		// We bubble the error all the way up instead of using logger.Fatal because logger.Fatal prints a huge stack
		// trace that hides the actual error message. This error msg is more explicit and doesn't require showing a
		// stack trace to the user.
		return nil, fmt.Errorf(
			"operator is not registered. Registering operator using the operator-cli before starting operator",
		)
	}

	operatorId, err := avsReader.GetOperatorId(&bind.CallOpts{}, common.HexToAddress(c.OperatorAddress))
	if err != nil {
		logger.Error("Cannot get operator id", "err", err)
		return nil, err
	}

	aggregatorRpcClient, err := NewAggregatorRpcClient(c.AggregatorServerIpPortAddress, logger)
	if err != nil {
		logger.Error("Cannot create AggregatorRpcClient. Is aggregator running?", "err", err)
		return nil, err
	}

	blsKeyPassword, ok := os.LookupEnv("OPERATOR_BLS_KEY_PASSWORD")
	if !ok {
		logger.Warnf("OPERATOR_BLS_KEY_PASSWORD env var not set. using empty string")
	}
	blsKeyPair, err := bls.ReadPrivateKeyFromFile(c.BlsPrivateKeyStorePath, blsKeyPassword)
	if err != nil {
		logger.Errorf("Cannot parse bls private key", "err", err)
		return nil, err
	}

	wsClient, err := ethclient.Dial(c.EthWsUrl)
	if err != nil {
		logger.Fatal("error connecting to web socket", "err", err)
	}

	query := ethereum.FilterQuery{
		Addresses: []common.Address{},
		Topics:    [][]common.Hash{{eventHash}},
	}

	newTaskCreatedLogs := make(chan types.Log)
	_, err = wsClient.SubscribeFilterLogs(context.Background(), query, newTaskCreatedLogs)
	if err != nil {
		logger.Fatal("error subscribing to newTaskCreated events", "err", err)
	}

	operator := &Operator{
		logger:                     logger,
		blsKeypair:                 blsKeyPair,
		aggregatorRpcClient:        aggregatorRpcClient,
		operatorId:                         operatorId,
		timesFailing:                       c.TimesFailing,
		newTaskCreatedLogs: 		newTaskCreatedLogs,
	}

	// Operator registration on startup should be deprecated already
	operator.operatorId = operatorId
	logger.Info("Operator info",
		"operatorId", operatorId,
		"operatorAddr", c.OperatorAddress,
		"operatorG1Pubkey", operator.blsKeypair.GetPubKeyG1(),
		"operatorG2Pubkey", operator.blsKeypair.GetPubKeyG2(),
	)

	return operator, nil
}

func (o *Operator) Start(ctx context.Context, taskResponseType interface{}) error {
	o.logger.Info("Starting operator.")

	for {
		select {
		case <-ctx.Done():
			return nil

		case log := <-o.newTaskCreatedLogs:
			// let task_response = self.taskProcessor.process_new_task(data);
            // let signed_task_response =
            //     Self::sign_task_response(&self.key_pair, &self.operator_id, task_response)?;
            // self.client_aggregator
            //     .send_signed_task_response(signed_task_response)
            //     .await?;

			taskResponse := o.taskProcessor.ProcessNewTaskCreatedLog(log)
			signedTaskResponse, err := o.SignTaskResponse(taskResponse)
			if err != nil {
				continue
			}
			go o.aggregatorRpcClient.SendSignedTaskResponseToAggregator(signedTaskResponse, taskResponseType)
		}
	}
}
/* 
// This function should be at incredible squaring lever
func (o *Operator) ProcessNewTaskCreatedLog(
	newTaskCreatedLog types.Log,
) sdkaggregator.TaskResponse {
	// o.logger.Debug("Received new task", "task", newTaskCreatedLog)
	// o.logger.Info("Received new task",
	// 	"numberToBeSquared", newTaskCreatedLog.Task.NumberToBeSquared,
	// 	"taskIndex", newTaskCreatedLog.TaskIndex,
	// 	"taskCreatedBlock", newTaskCreatedLog.Task.TaskCreatedBlock,
	// 	"quorumNumbers", newTaskCreatedLog.Task.QuorumNumbers,
	// 	"QuorumThresholdPercentage", newTaskCreatedLog.Task.QuorumThresholdPercentage,
	// )

	// Cast log to new task created event, and then create the Task response

	// cstaskmanager.ContractIncredibleSquaringTaskManagerNewTaskCreated


	numberSquared := big.NewInt(0).Exp(newTaskCreatedLog.Task.NumberToBeSquared, big.NewInt(2), nil)

	if o.timesFailing > 0 {
		rand.Seed(uint64((time.Now().UnixNano())))
		num := rand.Intn(100)
		if num < o.timesFailing {
			numberSquared = big.NewInt(908243203843)
			o.logger.Info("Operator computed wrong task result")
		}
	}
	taskResponse := &cstaskmanager.IIncredibleSquaringTaskManagerTaskResponse{
		ReferenceTaskIndex: newTaskCreatedLog.TaskIndex,
		NumberSquared:      numberSquared,
	}
	return taskResponse
} 
*/


func (o *Operator) SignTaskResponse(
	taskResponse sdkaggregator.TaskResponse,
) (*sdkaggregator.SignedTaskResponse, error) {
	taskResponseHash := taskResponse.Digest()

	blsSignature := o.blsKeypair.SignMessage(taskResponseHash) // Digest returns 256 bytes but keypair uses 32
	signedTaskResponse := &sdkaggregator.SignedTaskResponse{
		TaskResponse: taskResponse,
		BlsSignature: *blsSignature,
		OperatorId:   o.operatorId,
	}
	o.logger.Debug("Signed task response", "signedTaskResponse", signedTaskResponse)
	return signedTaskResponse, nil
}
