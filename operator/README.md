# Operator

## What is an Operator

Operators are off-chain nodes that perform, sign, and submit verifiable computations for Autonomous Verifiable Services (AVSs) using Ethereum restaking for security. They first register on EigenLayer’s core contracts, then opt-in to provide a range of services to AVSs. Operators listen for new task events, execute the supplied computation logic, cryptographically sign the results with `BLS/ECDSA` keys, and finally send the proofs to an aggregator for final consolidation.

## How the Logic Works

The Operator functions through the following flow:

1. **Task Subscription**:
   - The operator subscribes to specific event signatures emitted by task processors
   - Uses WebSocket connection to listen to blockchain events
   - Filters only for the specific task type it's designed to handle

2. **Task Processing**:
   - When a new task is detected, it extracts the task index and input data
   - Applies a computation function to the input data. This computation function is provided when creating the operator.

3. **Response Signing**:
   - Signs the computed result using the operator's BLS Key Pair.
   - Creates a `SignedTaskResponse` containing the result, signature, and operator ID

4. **Response Submission**:
   - Sends the signed response to an Aggregator service through an RPC request.

## How to Set Up an Operator

1. **Task Manager ABI**: Get the ABI of the task manager from your bindings.

   ```go
      taskManagerAbi, err := cstaskmanager.ContractIncredibleSquaringTaskManagerMetaData.GetAbi()
   ```

2. **Create the operator configuration**: Create an `operator.Config` struct. An alternative way to populate it is to load it from a config file like we do in the example.
    - Config fields:
      - `OperatorAddress`: The address of the operator
      - `RegistryCoordinatorAddress`: The address of the AVS registry coordinator
      - `EthRpcUrl`: The URL of the Ethereum RPC
      - `EthWsUrl`: The URL of the Ethereum WebSocket
      - `BlsPrivateKeyStorePath`: The path to the BLS private key store
      - `AggregatorServerIpPortAddress`: The IP and port of the aggregator
      - `Registration`: The registration configuration. If `Registration.RegisterOnStartup` is `true`, the operator attempts to register itself to EigenLayer using the values provided here. To skip automatic registration, leave this field unset (or set `Registration.RegisterOnStartup` to `false`), but note that this requires the operator to be already registered to EigenLayer.

      ```go
          operatorConfig := operator.Config{
            OperatorAddress:               "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266",
            RegistryCoordinatorAddress:    "0x7bc06c482dead17c0e297afbc32f6e63d3846650",
            EthRpcUrl:                     "http://localhost:8545",
            EthWsUrl:                      "ws://localhost:8545",
            BlsPrivateKeyStorePath:        "keys/test.bls.key.json",
            AggregatorServerIpPortAddress: "localhost:8090",
            Registration:               registrationConfig,
          }
      ```

3. **Processing Logic**: Implement the computation function that processes task inputs and produces outputs
   - This function will be called when the operator receives a `"NewTaskCreated"` event.

      ```go
        func square(taskIndex uint32, numberToSquare *big.Int) (*big.Int, error) {
          numberSquared := big.NewInt(0).Exp(numberToSquare, big.NewInt(2), nil)

          return numberSquared, nil
        }
      ```

4. **Response Calculator**: To abstract your computation into the operator, we provide a `ResponseCalculator` interface with a standard `functionResponseCalculator` struct. This struct implements the interface and a helper method for turning your function into `functionResponseCalculator`:
   - `NewFunctionResponseCalculator`: Create a response calculator from your computation function.

      ```go
          calculator := operator.NewFunctionResponseCalculator(square)
      ```

   - In case you need to save state in the operator, you can use your own struct implementing the `ResponseCalculator` interface.

5. **Failing Response Calculator**: If you want to test what happens when the operator responds incorrectly to a task and see how slashing works, you can wrap your logic with `NewFailingResponseCalculator` method to inject failures and a given failure rate. **Use this for testing purposes only.**

   ```go
      logic, err := operator.NewFailingResponseCalculator(calculator, 50, big.NewInt(0))
   ```

6. **Run the operator**: Initialize the `Operator` with the configuration and the processing logic. Then start the operator.

   ```go
      operator, err := operator.NewOperator(logger, operatorConfig, taskManagerAbi, logic, nil)
      if err != nil {
         logger.Errorf("Failed to create operator from config: %v", err)
         return
      }

      err = <-operator.Start(context.Background())
      if err != nil {
         logger.Errorf("Error while running operator: %v", err)
         return
      }
   ```

## Examples

Here are some examples of operators that are already implemented:

- [Incredible Squaring](https://github.com/Layr-Labs/eigensdk-go/blob/v2-dev-1/examples/incredible-squaring/operator/main.go)
- [Incredible Dot Product](https://github.com/Layr-Labs/eigensdk-go/blob/v2-dev-1/examples/incredible-dot-product/operator/main.go)
- [Awesome Vault Service](https://github.com/Layr-Labs/eigensdk-go/blob/v2-dev-1/examples/awesome-vault-service/operator/main.go)

## How to create your Response Calculator

To create your Response Calculator you have to declare a struct that satisfies the `ResponseCalculator` interface:

``` go
   type ResponseCalculator[Input any, Output any] interface {
      ComputeResponse(taskIndex uint32, input Input) (Output, error)
   }
```

The `ResponseCalculator` interface has the `ComputeResponse` method, which computes the response of a task given its index and input, returning the output, or an error if the computation fails.

We recommend implementing your own `ReponseCalculator` if you need to save state between responses, for that you should implement your Response calculator following the above interface. The response calculation logic should go in the `ComputeResponse` method and you can set your initial state in a constructor.

If you don't need to save a state between the operator responses you can use the `NewFunctionResponseCalculator` function, provided by the SDK in the operator package. You can find it on `operator/response_calculator.go`.
