# Aggregator

## What is an Aggregator

The Aggregator is a service component that coordinates BLS signature aggregation from multiple operators for each task. It acts as the bridge between the off-chain operator network and the on-chain smart contracts.

## How the Logic Works

The Aggregator operates through three main asynchronous processes:

1. **RPC Server Process**:

    - Runs a RPC-based server that listens for incoming operator responses
    - When an operator submits a signed task response, the Aggregator forwards the signature to the BLS aggregation service.

2. **Task Monitoring Process**:

    - Subscribes to blockchain events for new tasks
    - When a new task is detected, it creates a task metadata record
    - Sends the task metadata to the BLS Aggregation Service to begin signature collection

3. **Aggregation Process**:

    - Listens for aggregated results from the BLS aggregation service
    - When enough signatures are collected (meeting the quorum threshold), processes the result
    - Submits the aggregated signature along with information about non-signing operators to the blockchain

This flow ensures tasks are initialized, signatures collected, and the final response confirmed and forwarded to the AVS logic.

## How to Set Up an Aggregator

1. **Task Manager ABI**: Get the ABI of the task manager from your bindings.

    ```go
        taskManagerAbi, err := cstaskmanager.ContractIncredibleSquaringTaskManagerMetaData.GetAbi()
    ```

2. **Task Manager Definition**: Create a `taskManagerContractWrapper` struct that implements the `TaskManager` interface, which already implements `TaskResponder` one.
    - In the AVS examples we use a default task manager contract wrapper, that implements the interface, that can be created with the `NewTaskManagerFromAbi` function from `taskmanager` package, that creates a task manager contract wrapper from the task manager address and abi, a tx manager and an eth HTTP client.

    ``` go
        ethHttpUrl := "http://localhost:8545"
        ethHttpClient, err := ethclient.Dial(ethHttpUrl)
        if err != nil {
            return
        }

        ecdsaPrivateKey, err := crypto.HexToECDSA("2a871d0798f97d79848a013d4936a73bf4cc922c825d33c1cf7073dff6d409c6")
        if err != nil {
            logger.Errorf("Cannot parse ecdsa private key", "err", err)
            return
        }

        txMgr, err := txmgr.NewSimpleTxManagerFromPrivateKey(logger, ethHttpClient, ecdsaPrivateKey)
        if err != nil {
            logger.Errorf("Failed to create transaction manager", "err", err)
            return
        }

        taskManagerAbi, err := cstaskmanager.ContractIncredibleSquaringTaskManagerMetaData.GetAbi()
        if err != nil {
            logger.Fatalf(err.Error())
        }

        taskResponder, err := taskmanager.NewTaskManagerFromAbi[*big.Int, *big.Int](
            taskManagerAddr,
            taskManagerAbi,
            txMgr,
            ethHttpClient,
        )

    ```

3. **Create the aggregator configuration**: Create a `aggregator.Config` struct
    - Config fields:
        - `RegistryCoordinatorAddress`: The address of the AVS registry coordinator
        - `OperatorStateRetrieverAddress`: The address of the operator state retriever
        - `ServiceManagerAddress`: The address of the service manager
        - `EthHttpUrl`: The URL of the Ethereum HTTP RPC
        - `EthWsUrl`: The URL of the Ethereum WebSocket
        - `EcdsaPrivateKey`: The Ecdsa private key used by the aggregator to build the avs registry reader and subscriber
        - `AggregatorServerIpPortAddr`: The IP and port of the aggregator for listening for RPC calls

        ```go
            cfg := aggregator.Config{
                RegistryCoordinatorAddress:    common.HexToAddress("0x7bc06c482dead17c0e297afbc32f6e63d3846650"),
                OperatorStateRetrieverAddress: common.HexToAddress("0x4c5859f0f772848b2d91f1d83e2fe57935348029"),
                ServiceManagerAddress:         common.HexToAddress("0x5f3f1dbd7b74c6b46e8c44f98792a1daf8d69154"),
                EthHttpUrl:                    ethHttpUrl,
                EthWsUrl:                      "ws://localhost:8545",
                EcdsaPrivateKey:               ecdsaPrivateKey,
                AggregatorServerIpPortAddr:    "localhost:8090",
            }
        ```

4. **Task Processor Definition**: Create a struct implementing the `TaskProcessor` interface:

    - This is a interface that contains user-defined logic to handle new tasks, signed responses, and the final aggregated result.
    - We provide a standard `IndexingTaskProcessor` implementation that can be used as is for most cases. Its builder, `NewIndexingTaskProcessor` from `taskprocessor` package, receives a logger to put the logs and the task responder created below.

    ``` go
        taskProcessor, err := taskprocessor.NewIndexingTaskProcessor(logger, taskResponder)
    ```

5. **Run the aggregator**: Create an `Aggregator` with the config, a logger, the task processor created below and the task manager contract ABI, and then start it:

    ``` go
        agg, err := aggregator.NewAggregator(cfg, logger, taskProcessor, taskManagerAbi)

        err = agg.Start(context.Background())
    ```

## Examples

Here are some examples of aggregator implementations:

- [Incredible Squaring](https://github.com/Layr-Labs/eigensdk-go/blob/v2-dev-1/examples/incredible-squaring/aggregator/main.go)
- [Incredible Dot Product](https://github.com/Layr-Labs/eigensdk-go/blob/v2-dev-1/examples/incredible-dot-product/aggregator/main.go)
- [Awesome Vault Service](https://github.com/Layr-Labs/eigensdk-go/blob/v2-dev-1/examples/awesome-vault-service/aggregator/main.go)
