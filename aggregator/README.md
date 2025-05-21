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

1. **Create the aggregator configuration**: Create a `aggregator.Config` struct
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
                EthHttpUrl:                    ethHttpUrl,
                EthWsUrl:                      "ws://localhost:8545",
                AggregatorServerIpPortAddr:    "localhost:8090",
            }
        ```

2. **Provide a Task Processor**: Provide a struct implementing the `TaskProcessor` interface. This interface contains user-defined logic to handle new tasks, signed responses, and the final aggregated result. We provide a standard `IndexingTaskProcessor` implementation that can be used as is for most cases, you can create it usig the `NewIndexingTaskProcessor` builder from `taskprocessor` package:

    1. Instantiate an ethereum client and a transaction manager for the task responder:

        ``` go
            ethHttpClient, err := ethclient.Dial("http://localhost:8545")

            ecdsaPrivateKey, err := crypto.HexToECDSA("2a871d0798f97d79848a013d4936a73bf4cc922c825d33c1cf7073dff6d409c6")
            txMgr, err := txmgr.NewSimpleTxManagerFromPrivateKey(logger, ethHttpClient, ecdsaPrivateKey)
        ```

    2. Get the `TaskManager` ABI from the contract binding:

        ``` go
            taskManagerAbi, err := cstaskmanager.ContractIncredibleSquaringTaskManagerMetaData.GetAbi()
        ```

    3. Provide a struct that implements the `TaskResponder` interface. That interface requires a method to send the aggregated task responses and another for hashing task responses. Here we provide an SDK implementation that satisfies the Task Responder interface, receiving the TaskManager address and ABI, a transaction manager and an ethereum client:

        ``` go
            taskResponder, err := taskmanager.NewTaskManagerFromAbi[*big.Int, *big.Int](
                taskManagerAddr,
                taskManagerAbi,
                txMgr,
                ethHttpClient,
            )
        ```

    4. Create the `IndexingTaskProcessor` with the `TaskResponder` created below:

        ```go
            taskProcessor, err := taskprocessor.NewIndexingTaskProcessor(logger, taskResponder)
        ```

3. **Run the aggregator**: Instantiate an `Aggregator` with the config, a logger, the task processor created below and the task manager contract ABI, and then start it:

    ``` go
        agg, err := aggregator.NewAggregator(cfg, logger, taskProcessor, taskManagerAbi)

        err = agg.Start(context.Background())
    ```

## Examples

Here are some examples of aggregator implementations:

- [Incredible Squaring](https://github.com/Layr-Labs/eigensdk-go/blob/v2-dev-1/examples/incredible-squaring/aggregator/main.go)
- [Incredible Dot Product](https://github.com/Layr-Labs/eigensdk-go/blob/v2-dev-1/examples/incredible-dot-product/aggregator/main.go)
- [Awesome Vault Service](https://github.com/Layr-Labs/eigensdk-go/blob/v2-dev-1/examples/awesome-vault-service/aggregator/main.go)
