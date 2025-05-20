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

1. **Task Manager ABI**: Get the ABI of the task manager from your bindings.

   ```go
      taskManagerAbi, err := cstaskmanager.ContractIncredibleSquaringTaskManagerMetaData.GetAbi()
   ```


2. **Task Verification Logic**: Define a function that computes the expected result for a task, which will be used to verify operator responses
   - This would be the logic to compute a new task.

      ```go
        func square(taskIndex uint32, numberToSquare *big.Int) (*big.Int, error) {
          numberSquared := big.NewInt(0).Exp(numberToSquare, big.NewInt(2), nil)
          return numberSquared, nil
        }
      ```

4. **Response Calculator**: To abstract your computation into the challenger, we provide a `ResponseCalculator` interface with a standar `functionResponseCalculator` struct. This struct implements the interface and a helper method for turning your function into `functionResponseCalculator`:
   - `NewFunctionResponseCalculator`: Create a response calculator from your computation function.

      ```go
          calculator := operator.NewFunctionResponseCalculator(square)
      ```

5. **Verifier**: Create a verifier from the response calculator.
   - This will be in charge of computing the response of a task and will use a user-defined function to compare the computed response with the operator's response.

      ```go
          // BigIntEqual receives two *big.Int and returns whether they are equal
          func BigIntEqual(a, b *big.Int) bool {
            return a.Cmp(b) == 0
          }

          validation := challengerprocessor.ResponseValidationFunctionFromResponseCalculator(calculator, BigIntEqual)
      ```

6. **Challenger Task Processor**: Create a [`ChallengerTaskProcessor`] interface implementation.
    - This will be in charge of processing the task and the response.
    - We provide a standard [`IndexingChallengerProcessor`] struct that can be used as a starting point.

      ```go
          indexingTaskProcessor, err := challengerprocessor.NewIndexingChallengerProcessor(logger, validation, challengerRaiser)
      ```

7. **Challenger**: Create a [`Challenger`] from the [`Config`] and the [`ChallengerTaskProcessor`].

    ```go
        challenger, _ := challenger.NewChallenger(
          cfg,
          indexingTaskProcessor,
        )
    ```

8. **Start the Challenger**: Start the challenger.

    ```go
        challenger.Start(context.Background())
    ```


## Examples

Here are some examples of challenger implementations:

- [Incredible Squaring](https://github.com/Layr-Labs/eigensdk-go/blob/v2-dev-1/examples/incredible-squaring/challenger/main.go)
- [Incredible Dot Product](https://github.com/Layr-Labs/eigensdk-go/blob/v2-dev-1/examples/incredible-dot-product/challenger/main.go)
- [Awesome Vault Service](https://github.com/Layr-Labs/eigensdk-go/blob/v2-dev-1/examples/awesome-vault-service/challenger/main.go)
