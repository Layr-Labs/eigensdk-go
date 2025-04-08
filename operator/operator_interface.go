package operator

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/prometheus/client_golang/prometheus"
	"golang.org/x/exp/rand"

	sdkaggregator "github.com/Layr-Labs/eigensdk-go/aggregator"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/avsregistry"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/elcontracts"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/wallet"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	sdkecdsa "github.com/Layr-Labs/eigensdk-go/crypto/ecdsa"
	sdklogging "github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/metrics/collectors/economic"
	"github.com/Layr-Labs/eigensdk-go/nodeapi"
	"github.com/Layr-Labs/eigensdk-go/signerv2"
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

type Operator struct {
	config    OperatorConfig
	logger    sdklogging.Logger
//	ethClient sdkcommon.EthClientInterface
//	metricsReg       *prometheus.Registry
//	metrics          metrics.Metrics
	nodeApi          *nodeapi.NodeApi
	avsWriter        *avsregistry.ChainWriter
	avsReader        *avsregistry.ChainReader
	avsSubscriber    *avsregistry.ChainSubscriber
	eigenlayerReader *elcontracts.ChainReader
	eigenlayerWriter *elcontracts.ChainWriter
	blsKeypair       *bls.KeyPair
	operatorId       sdktypes.OperatorId
	operatorAddr     common.Address
	// receive new tasks in this chan (typically from listening to onchain event)
	// newTaskCreatedChan chan *cstaskmanager.ContractIncredibleSquaringTaskManagerNewTaskCreated
	// ip address of aggregator
	aggregatorServerIpPortAddr string
	// rpc client to send signed task responses to aggregator
	aggregatorRpcClient AggregatorRpcClienter
	// needed when opting in to avs (allow this service manager contract to slash operator)
	credibleSquaringServiceManagerAddr common.Address
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
	reg := prometheus.NewRegistry()
	//eigenMetrics := sdkmetrics.NewEigenMetrics(AVS_NAME, c.EigenMetricsIpPortAddress, reg, logger)
	//avsAndEigenMetrics := metrics.NewAvsAndEigenMetrics(AVS_NAME, eigenMetrics, reg)

	// Setup Node Api
	nodeApi := nodeapi.NewNodeApi(AVS_NAME, SEM_VER, c.NodeApiIpPortAddress, logger)

	// var ethRpcClient, ethWsClient sdkcommon.EthClientInterface
	// if c.EnableMetrics {
	// 	rpcCallsCollector := rpccalls.NewCollector(AVS_NAME, reg)
	// 	ethRpcClient, err = eth.NewInstrumentedClient(c.EthRpcUrl, rpcCallsCollector)
	// 	if err != nil {
	// 		logger.Errorf("Cannot create http ethclient", "err", err)
	// 		return nil, err
	// 	}
	// 	ethWsClient, err = eth.NewInstrumentedClient(c.EthWsUrl, rpcCallsCollector)
	// 	if err != nil {
	// 		logger.Errorf("Cannot create ws ethclient", "err", err)
	// 		return nil, err
	// 	}
	// } else {
	// 	ethRpcClient, err = ethclient.Dial(c.EthRpcUrl)
	// 	if err != nil {
	// 		logger.Errorf("Cannot create http ethclient", "err", err)
	// 		return nil, err
	// 	}
	// 	ethWsClient, err = ethclient.Dial(c.EthWsUrl)
	// 	if err != nil {
	// 		logger.Errorf("Cannot create ws ethclient", "err", err)
	// 		return nil, err
	// 	}
	// }

	ethHttpClient, err := ethclient.Dial(c.EthRpcUrl)
	if err != nil {
		return nil, utils.WrapError("Failed to create Eth Http client", err)
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
	// TODO(samlaf): should we add the chainId to the config instead?
	// this way we can prevent creating a signer that signs on mainnet by mistake
	// if the config says chainId=5, then we can only create a goerli signer
	chainId, err := ethHttpClient.ChainID(context.Background())
	if err != nil {
		logger.Error("Cannot get chainId", "err", err)
		return nil, err
	}

	ecdsaKeyPassword, ok := os.LookupEnv("OPERATOR_ECDSA_KEY_PASSWORD")
	if !ok {
		logger.Warnf("OPERATOR_ECDSA_KEY_PASSWORD env var not set. using empty string")
	}

	signerV2, _, err := signerv2.SignerFromConfig(signerv2.Config{
		KeystorePath: c.EcdsaPrivateKeyStorePath,
		Password:     ecdsaKeyPassword,
	}, chainId)
	if err != nil {
		panic(err)
	}

	chainioConfig := clients.BuildAllConfig{
		EthHttpUrl:                  c.EthRpcUrl,
		EthWsUrl:                    c.EthWsUrl,
		RegistryCoordinatorAddr:     c.AVSRegistryCoordinatorAddress,
		OperatorStateRetrieverAddr:  c.OperatorStateRetrieverAddress,
		ServiceManagerAddress:       c.IncredibleSquaringServiceManager,
		RewardsCoordinatorAddress:   c.RewardsCoordinatorAddress,
		PermissionControllerAddress: c.PermissionControllerAddress,
		AvsName:                     AVS_NAME,
		PromMetricsIpPortAddress:    c.EigenMetricsIpPortAddress,
	}
	operatorEcdsaPrivateKey, err := sdkecdsa.ReadKey(
		c.EcdsaPrivateKeyStorePath,
		ecdsaKeyPassword,
	)
	if err != nil {
		return nil, err
	}
	sdkClients, err := clients.BuildAll(chainioConfig, operatorEcdsaPrivateKey, logger)
	if err != nil {
		panic(err)
	}
	skWallet, err := wallet.NewPrivateKeyWallet(ethHttpClient, signerV2, common.HexToAddress(c.OperatorAddress), logger)
	if err != nil {
		panic(err)
	}
	txMgr := txmgr.NewSimpleTxManager(skWallet, ethHttpClient, logger, common.HexToAddress(c.OperatorAddress))

	avs_config := avsregistry.Config{
		RegistryCoordinatorAddress: common.HexToAddress(c.AVSRegistryCoordinatorAddress),
		OperatorStateRetrieverAddress: common.HexToAddress(c.OperatorStateRetrieverAddress),
		ServiceManagerAddress: common.HexToAddress(c.IncredibleSquaringServiceManager),
	}

	avsWriter, err := avsregistry.NewWriterFromConfig(avs_config, ethHttpClient, txMgr, logger)
	if err != nil {
		logger.Error("Cannot create AvsWriter", "err", err)
		return nil, err
	}

	avsReader, err := avsregistry.NewReaderFromConfig(avs_config, ethHttpClient, logger)
	if err != nil {
		logger.Error("Cannot create AvsReader", "err", err)
		return nil, err
	}

	avsSubscriber, err := avsregistry.NewSubscriberFromConfig(avs_config, sdkClients.EthWsClient, logger)
	if err != nil {
		logger.Error("Cannot create AvsSubscriber", "err", err)
		return nil, err
	}

	// We must register the economic metrics separately because they are exported metrics (from jsonrpc or subgraph
	// calls)
	// and not instrumented metrics: see https://prometheus.io/docs/instrumenting/writing_clientlibs/#overall-structure
	quorumNames := map[sdktypes.QuorumNum]string{
		0: "quorum0",
	}
	economicMetricsCollector := economic.NewCollector(
		sdkClients.ElChainReader, sdkClients.AvsRegistryChainReader,
		AVS_NAME, logger, common.HexToAddress(c.OperatorAddress), quorumNames)
	reg.MustRegister(economicMetricsCollector)

	aggregatorRpcClient, err := NewAggregatorRpcClient(c.AggregatorServerIpPortAddress, logger)
	if err != nil {
		logger.Error("Cannot create AggregatorRpcClient. Is aggregator running?", "err", err)
		return nil, err
	}

	client, err := ethclient.Dial(c.EthWsUrl)
	if err != nil {
		logger.Fatal("error connecting to web socket", "err", err)
	}

	query := ethereum.FilterQuery{
		Addresses: []common.Address{},
		Topics:    [][]common.Hash{{eventHash}},
	}

	newTaskCreatedLogs := make(chan types.Log)
	_, err = client.SubscribeFilterLogs(context.Background(), query, newTaskCreatedLogs)
	if err != nil {
		logger.Fatal("error subscribing to newTaskCreated events", "err", err)
	}

	operator := &Operator{
		config:                     c,
		logger:                     logger,
//		metricsReg:                 reg,
//		metrics:                    avsAndEigenMetrics,
		nodeApi:                    nodeApi,
//		ethClient:                  ethRpcClient,
		avsWriter:                  avsWriter,
		avsReader:                  avsReader,
		avsSubscriber:              avsSubscriber,
		eigenlayerReader:           sdkClients.ElChainReader,
		eigenlayerWriter:           sdkClients.ElChainWriter,
		blsKeypair:                 blsKeyPair,
		operatorAddr:               common.HexToAddress(c.OperatorAddress),
		aggregatorServerIpPortAddr: c.AggregatorServerIpPortAddress,
		aggregatorRpcClient:        aggregatorRpcClient,
		// newTaskCreatedChan: make(
		// 	chan *cstaskmanager.ContractIncredibleSquaringTaskManagerNewTaskCreated,
		// ),
		credibleSquaringServiceManagerAddr: common.HexToAddress(c.IncredibleSquaringServiceManager),
		operatorId:                         [32]byte{0}, // this is set below
		timesFailing:                       c.TimesFailing,
		newTaskCreatedLogs: 		newTaskCreatedLogs,
	}

	// Operator registration on startup should be deprecated already

	// OperatorId is set in contract during registration so we get it after registering operator.
	operatorId, err := sdkClients.AvsRegistryChainReader.GetOperatorId(&bind.CallOpts{}, operator.operatorAddr)
	if err != nil {
		logger.Error("Cannot get operator id", "err", err)
		return nil, err
	}
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
	operatorIsRegistered, err := o.avsReader.IsOperatorRegistered(&bind.CallOpts{}, o.operatorAddr)
	if err != nil {
		o.logger.Error("Error checking if operator is registered", "err", err)
		return err
	}
	if !operatorIsRegistered {
		// We bubble the error all the way up instead of using logger.Fatal because logger.Fatal prints a huge stack
		// trace that hides the actual error message. This error msg is more explicit and doesn't require showing a
		// stack trace to the user.
		return fmt.Errorf(
			"operator is not registered. Registering operator using the operator-cli before starting operator",
		)
	}

	o.logger.Info("Starting operator.")

	if o.config.EnableNodeApi {
		o.nodeApi.Start()
	}
	// var metricsErrChan <-chan error
	// if o.config.EnableMetrics {
	// 	metricsErrChan = o.metrics.Start(ctx, o.metricsReg)
	// } else {
	// 	metricsErrChan = make(chan error, 1)
	// }

	for {
		select {
		case <-ctx.Done():
			return nil
		// case err := <-metricsErrChan:
		// 	// TODO(samlaf); we should also register the service as unhealthy in the node api
		// 	// https://eigen.nethermind.io/docs/spec/api/
		// 	o.logger.Fatal("Error in metrics server", "err", err)
		// case err := <-sub.Err():
		// 	o.logger.Error("Error in websocket subscription", "err", err)
		// 	// TODO(samlaf): write unit tests to check if this fixed the issues we were seeing
		// 	sub.Unsubscribe()
		// 	// TODO(samlaf): wrap this call with increase in avs-node-spec metric
		// 	sub = o.avsSubscriber.SubscribeToNewTasks(o.newTaskCreatedChan)
		case log := <-o.newTaskCreatedLogs:
			// o.metrics.IncNumTasksReceived()
			taskResponse := o.ProcessNewTaskCreatedLog(log)
			signedTaskResponse, err := o.SignTaskResponse(taskResponse)
			if err != nil {
				continue
			}
			go o.aggregatorRpcClient.SendSignedTaskResponseToAggregator(signedTaskResponse, taskResponseType) // &aggregator.IncredibleSquaringTaskResponse{}
		}
	}
}

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


/*

/// Operator methods
type Operator interface {
    /// type TaskResponse
    /// type NewTaskEvent

    processNewTask(newTaskCreated: NewTaskEvent)(TaskResponse)

    func start_operator(
        avs_registry_reader: &AvsRegistryChainReader,
        key_pair: &BlsKeyPair,
        operator_id: &OperatorId,
        operator_address: Address,
        operator_name: &str,
        client_aggregator: &ClientAggregator,
        ws_rpc_url: &str,
    ) -> impl std::future::Future<Output = Result<(), OperatorError>> + Send {
        async move {
            let is_registered = avs_registry_reader
                .is_operator_registered(operator_address)
                .await
                .map_err(|_| OperatorError::RegistrationError)?;
            info!("is {} registered {}", operator_name, is_registered);
            let arc_client = Arc::new(client_aggregator);
            if !is_registered {
                return Err(OperatorError::RegistrationError);
            }

            info!("Starting operator");

            let ws = WsConnect::new(ws_rpc_url);
            let provider = ProviderBuilder::new()
                .on_ws(ws)
                .await
                .map_err(|_| OperatorError::TransportError)?;

            let filter = Filter::new().event_signature(Self::NewTaskEvent::SIGNATURE_HASH);
            let sub = provider
                .subscribe_logs(&filter)
                .await
                .map_err(|_| OperatorError::SubscribeLogsError)?;
            let mut stream = sub.into_stream();

            // TODO: this loop could be a separate function
            while let Some(log) = stream.next().await {
                // TODO: check error handling (maybe use continue)
                let data: Self::NewTaskEvent = log
                    .log_decode()
                    .map_err(|_| OperatorError::SubscribeLogsError)?
                    .inner
                    .data;
                info!("{} picked up a new task", operator_name);
                let task_response = Self::process_new_task(data);
                let signed_task_response =
                    Self::sign_task_response(key_pair, operator_id, task_response)?;
                let _ = arc_client
                    .send_signed_task_response(signed_task_response)
                    .await;
            }
            Ok(())
        }
    }

    /// Sign the task response
    fn sign_task_response(
        key_pair: &BlsKeyPair,
        operator_id: &OperatorId,
        task_response: Self::TaskResponse,
    ) -> Result<SignedTaskResponse<Self::TaskResponse>, OperatorError> {
        let hash_msg = task_response.digest();
        let signed_msg = key_pair.sign_message(&hash_msg);
        let signed_task_response = SignedTaskResponse::new(task_response, signed_msg, *operator_id);
        Ok(signed_task_response)
    }
}

*/
