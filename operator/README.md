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
   - Sends the signed response to an Aggregator service through a RPC request.

## How to Set Up an Operator

1. **Task Manager ABI**: Get the ABI of the task manager from your bindings.
   
   ```go
      taskManagerAbi, err := cstaskmanager.ContractIncredibleSquaringTaskManagerMetaData.GetAbi()
   ```

2. **Create the operator configuration**: Create a `operator.Config` struct
    - By default, the operator attempts to register itself to EigenLayer using the values provided in the `operator.RegistrationConfig` struct. To skip automatic registration, leave the `RegistrationCfg` field unset, but make sure the operator is already registered to EigenLayer.

    - Config fields:
      - `OperatorAddress`: The address of the operator
      - `OperatorStateRetrieverAddress`: The address of the operator state retriever
      - `ServiceManagerAddress`: The address of the service manager
      - `AVSRegistryCoordinatorAddress`: The address of the AVS registry coordinator
      - `EthRpcUrl`: The URL of the Ethereum RPC
      - `EthWsUrl`: The URL of the Ethereum WebSocket
      - `BlsPrivateKeyStorePath`: The path to the BLS private key store
      - `AggregatorServerIpPortAddress`: The IP and port of the aggregator
      - `Logger`: The logger
      - `TaskManagerAbi`: The ABI of the task manager
      - `RegistrationCfg`: The registration configuration

      ```go
          operatorConfig := operator.Config{
            OperatorAddress:               "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266",
            OperatorStateRetrieverAddress: "0x4c5859f0f772848b2d91f1d83e2fe57935348029",
            ServiceManagerAddress:         "0x5f3f1dbd7b74c6b46e8c44f98792a1daf8d69154",
            AVSRegistryCoordinatorAddress: "0x7bc06c482dead17c0e297afbc32f6e63d3846650",
            EthRpcUrl:                     "http://localhost:8545",
            EthWsUrl:                      "ws://localhost:8545",
            BlsPrivateKeyStorePath:        "keys/test.bls.key.json",
            AggregatorServerIpPortAddress: "localhost:8090",
            Logger:                        logger,
            TaskManagerAbi:                taskManagerAbi,
            RegistrationCfg:               registrationConfig,
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

4. **Response Calculator**: To abstract your computation into the operator, we provide a `ResponseCalculator` interface with a standar `functionResponseCalculator` struct. This struct implements the interface and a helper method for turning your function into `functionResponseCalculator`:
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
        operator, err := operator.NewOperatorFromConfig(operatorConfig, logic, nil)
        if err != nil {
          logger.Errorf("Failed to create operator from config: %v", err)
          return
        }

        err = operator.Start(context.Background())
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

