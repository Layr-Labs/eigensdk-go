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

1. **Create the aggregator configuration**: Create an `aggregator.Config`. An alternative way to populate it is to load it from a config file like we do in the example.
    - Config fields:
        - `RegistryCoordinatorAddress`: The address of the AVS registry coordinator
        - `OperatorStateRetrieverAddress`: The address of the operator state retriever
        - `EthHttpUrl`: The URL of the Ethereum HTTP RPC
        - `EthWsUrl`: The URL of the Ethereum WebSocket
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

2. **Provide an Aggregator Processor**: Provide a struct implementing the `AggregatorProcessor` interface. This interface contains user-defined logic to handle new tasks, signed responses, and the final aggregated result. We provide a standard `IndexingAggregatorProcessor` implementation that can be used as is for most cases, you can create it using the `NewIndexingAggregatorProcessor` constructor from `aggregatorprocessor` package:

    1. Instantiate an Ethereum client and a transaction manager for the task responder:

        ``` go
            ethHttpClient, err := ethclient.Dial("http://localhost:8545")

            ecdsaPrivateKey, err := crypto.HexToECDSA("2a871d0798f97d79848a013d4936a73bf4cc922c825d33c1cf7073dff6d409c6")
            txMgr, err := txmgr.NewSimpleTxManagerFromPrivateKey(logger, ethHttpClient, ecdsaPrivateKey)
        ```

    2. Get the `TaskManager` ABI from the contract binding:

        ``` go
            taskManagerAbi, err := cstaskmanager.ContractIncredibleSquaringTaskManagerMetaData.GetAbi()
        ```

    3. Provide a struct that implements the `TaskResponder` interface. That interface requires a method to send the aggregated task responses and another for hashing task responses. Here we provide an SDK implementation that satisfies the Task Responder interface, receiving the TaskManager address and ABI, a transaction manager and an Ethereum client:

        ``` go
            taskResponder, err := taskmanager.NewTaskManagerFromAbi[*big.Int, *big.Int](
                taskManagerAddr,
                taskManagerAbi,
                txMgr,
                ethHttpClient,
            )
        ```

    4. Create the `IndexingAggregatorProcessor` with the `TaskResponder` created above:

        ```go
            aggregatorProcessor, err := aggregatorprocessor.NewIndexingAggregatorProcessor(logger, taskResponder)
        ```

3. **Run the aggregator**: Instantiate an `Aggregator` with the config, a logger, the aggregator processor created above and the task manager contract ABI, and then start it:

    ``` go
        agg, err := aggregator.NewAggregator(logger, cfg, taskManagerAbi, aggregatorProcessor)

        err = <-agg.Start(context.Background())
    ```

## Examples

Here are some examples of aggregator implementations:

- [Incredible Squaring](https://github.com/Layr-Labs/eigensdk-go/blob/v2-dev-1/examples/incredible-squaring/aggregator/main.go)
- [Incredible Dot Product](https://github.com/Layr-Labs/eigensdk-go/blob/v2-dev-1/examples/incredible-dot-product/aggregator/main.go)
- [Awesome Vault Service](https://github.com/Layr-Labs/eigensdk-go/blob/v2-dev-1/examples/awesome-vault-service/aggregator/main.go)

## How to create your Aggregator Processor

To create your Aggregator Processor you have to declare a struct that satisfies the `AggregatorProcessor` interface:

``` go
    type AggregatorProcessor[Input any, Output any] interface {
        ProcessNewTask(taskIndex sdktypes.TaskIndex, task taskmanager.Task[Input]) (blsagg.TaskMetadata, error)
        ProcessTaskResponse(taskResponse taskmanager.TaskResponse[Output]) ([32]byte, error)
        ProcessAggregatedResponse(taskIndex sdktypes.TaskIndex, taskResponse taskmanager.TaskResponse[Output], nonSignerStakesAndSignature sdktypes.NonSignerStakesAndSignature) error
    }
```

If you want to see an example of `AggregatorProcessor` you can watch our `IndexingAggregatorProcessor` on `aggregator/aggregator-processor/indexing_aggregator_processor.go`.

The `AggregatorProcessor` interface has the following methods:

1. `ProcessNewTask`, that should:

    - Saves the task with the associated task index.
    - Creates the task metadata that will be sent to the BLS aggregation service to be processed.
    - Returns the BLS metadata

2. `ProcessTaskResponse`, that should:

    - Encode the task response in the `TaskManager` ABI
    - Hash the encoded task response
    - Return the digest

3. `ProcessAggregatedResponse` receives a task index, a task response and the non signer stakes and signature, and sends the aggregated response information to the on-chain TaskManager contract, returning an error if the processing fails.

    - Obtain the task with the received task index.
    - Process the BLS aggregated response, obtaining the nonsigner stakes and signature
    - Send the task, task response, and nonsigner stakes and signature to the on-chain `TaskManager` contract
