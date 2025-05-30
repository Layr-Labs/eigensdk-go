# Challenger

## What is a Challenger

A Challenger is a validator component that monitors the network for the creation of new tasks and task responses submitted by operators, verifies their correctness, and raises challenges when incorrect responses are detected. If the challenge is successful, the operator will be slashed.

## How the Logic Works

The Challenger operates through a well-defined workflow:

1. **Event Subscription**:

    - Subscribes to blockchain events for new tasks and task responses
    - Monitors for `NewTaskCreated` to track new tasks created in the system
    - Watches for `TaskResponded` when operators submit responses to tasks

2. **Verification Process**:

    - When a new task is detected, user defined logic is used to process the new task
    - When a task response is received, user defined logic is used to process the task response and verify it
    - Uses a user-defined verification function to determine if the response is correct
    - The verification logic can be customized based on the specific AVS requirements

3. **Challenge Mechanism**:

    - If a response is verified as correct, the challenger logs the result and takes no action
    - If a response is determined to be incorrect, the challenger raises a challenge
    - Includes identifying the non-signing operators who might have abstained from the incorrect response

## How to Set Up a Challenger

1. **Challenger Configuration**: Create a `challenger.Config` struct with the following fields:
   - `EthWsUrl`: The URL of the Ethereum websocket
   - `Logger`: The logger
   - `TaskManagerAbi`: The ABI of the task manager
   - `EthClient`: The Ethereum client

      ```go
        ethHttpUrl := "http://localhost:8545"
        ethHttpClient, _ := ethclient.Dial(ethHttpUrl)
      
        cfg := challenger.Config{
          EthWsUrl:       "ws://localhost:8545",
          Logger:         logger,
          TaskManagerAbi: taskManagerAbi,
          EthClient:      ethHttpClient,
        }
      ```

2. **Task Verification Logic**: Define a function that verifies the response for a task, returning `true` for valid tasks.

    ```go
      func isValidSquare(taskIndex uint32, numberToSquare, numberSquared *big.Int) (bool, error) {
          expectedNumberSquared := big.NewInt(0).Exp(numberToSquare, big.NewInt(2), nil)
          return expectedNumberSquared.Cmp(numberSquared) == 0, nil
      }
    ```

    - There is another approach where you can use the logic from the operator to compute the task. To do this, you can wrap the logic into a `ResponseCalculator` implementation. Then you can use `ResponseValidationFunctionFromResponseCalculator`, which will be in charge of computing the response of a task and will use a user-defined function to compare the computed response with the operator's response.

3. **Provide a Challenger Processor**: Provide a struct implementing the `ChallengerProcessor` interface. This interface contains user-defined logic to process new tasks and task responses. We provide a standard `IndexingChallengerProcessor` implementation that can be used as-is for most cases, you can create it using the `NewIndexingChallengerProcessor` constructor from `challengerprocessor` package:

    1. Instantiate an Ethereum client and a transaction manager for the `ChallengerProcessor`:

        ``` go
            ethHttpClient, err := ethclient.Dial("http://localhost:8545")

            ecdsaPrivateKey, err := crypto.HexToECDSA("2a871d0798f97d79848a013d4936a73bf4cc922c825d33c1cf7073dff6d409c6")
            txMgr, err := txmgr.NewSimpleTxManagerFromPrivateKey(logger, ethHttpClient, ecdsaPrivateKey)
        ```

    2. Get the `TaskManager` ABI from the contract binding:

        ``` go
            taskManagerAbi, err := cstaskmanager.ContractIncredibleSquaringTaskManagerMetaData.GetAbi()
        ```

    3. Provide a struct that implements the `ChallengeRaiser` interface. That interface requires a method to raise an on-chain challenge when operator responses are incorrect. Here we provide a SDK implementation that satisfies the `ChallengeRaiser` interface, receiving the TaskManager address and ABI, a transaction manager and an Ethereum client:

        ``` go
            challengerRaiser, err := taskmanager.NewTaskManagerFromAbi[*big.Int, *big.Int](
                taskManagerAddr,
                taskManagerAbi,
                txMgr,
                ethHttpClient,
            )
        ```

    4. Create the `NewIndexingChallengerProcessor` with the `ChallengeRaiser` created above:

        ```go
            challengerProcessor, err := challengerprocessor.NewIndexingChallengerProcessor(logger, isValidSquare, challengerRaiser)
        ```

4. **Challenger**: Create a `Challenger` from the `Config` and the `ChallengerProcessor`.

    ```go
        challenger, _ := challenger.NewChallenger(logger, cfg, taskManagerAbi, challengerProcessor)
    ```

5. **Start the Challenger**: Start the challenger.

    ```go
        err = <-challenger.Start(context.Background())
    ```

## Examples

Here are some examples of challenger implementations:

- [Incredible Squaring](https://github.com/Layr-Labs/eigensdk-go/blob/v2-dev-1/examples/incredible-squaring/challenger/main.go)
- [Incredible Dot Product](https://github.com/Layr-Labs/eigensdk-go/blob/v2-dev-1/examples/incredible-dot-product/challenger/main.go)
- [Awesome Vault Service](https://github.com/Layr-Labs/eigensdk-go/blob/v2-dev-1/examples/awesome-vault-service/challenger/main.go)

## How to implement a custom Challenger Processor

To implement a custom Challenger Processor, you must implement the `ChallengerProcessor` interface:

```go
  type ChallengerProcessor[Input any, Output any] interface {
    ProcessNewTaskCreated(taskIndex uint32, task taskmanager.Task[Input]) error
    ProcessTaskResponded(taskIndex uint32, taskResponse taskmanager.TaskResponse[Output], taskResponseMetadata sdktypes.TaskResponseMetadata, nonSigningOperatorPubKeys []sdktypes.BN254G1Point) error
  }
```

This interface has two methods, which we explain in the next sections. Refer to the [`IndexingChallengerProcessor`](https://github.com/Layr-Labs/eigensdk-go/blob/v2-dev-2/challenger/challenger-processor/challenger_task_processor.go) implementation for an example of how to implement a custom Challenger Processor.

### `ProcessNewTaskCreated`

Invoked when a new task is emitted by the contract. `ProcessNewTaskCreated` is commonly used to store the task (e.g., in a map, in a database, etc.), so that when a response arrives, you can retrieve the corresponding input.

### `ProcessTaskResponded`

Invoked when a task response is received. This method should verify the operator’s response and raise a challenge for invalid responses. Step by step, this method:

1. Retrieves the original task using the index.
2. Verifies the operator’s response against the task's input.
3. Raises a challenge through the `TaskManager` if the responses differ.
