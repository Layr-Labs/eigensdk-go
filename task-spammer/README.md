# Task Spammer

## What is a Task Spammer

A Task Spammer is a testing utility designed to generate tasks at configurable intervals. It serves as a simulation tool that allows developers to test how operators, aggregators and challengers respond to a continuous stream of new tasks.

## How the Logic Works

The Task Spammer uses a builder pattern and follows this workflow:

1. **Task Generation Logic**:
    - Uses an iterator to produce a sequence of task inputs
    - Each iterator value becomes the input for a new task
    
2. **Task Submission**:
    - Connects to a TaskManager contract
    - Calls `CreateNewTask` with the generated input

## How to Set Up a Task Spammer

1. **Task Manager Definition**: Create a `taskManagerContractWrapper` struct that implements the `TaskManager` interface.
    - `Input` and `Output` types for your tasks
    - Use the `taskmanager.NewTaskManagerFromAbi` method to build your `taskManagerContractWrapper`.

      ```go
        txMgr, _ := txmgr.NewSimpleTxManagerFromPrivateKey(logger, ethHttpClient, ecdsaPrivateKey)
        abi, _ := cstaskmanager.ContractIncredibleSquaringTaskManagerMetaData.GetAbi()

        taskCreator, _ := taskmanager.NewTaskManagerFromAbi[*big.Int, *big.Int](taskManagerAddress, abi, txMgr, ethHttpClient)
      ```

2. **Task Spammer Config**: Create a `taskspammer.Config` struct with the following fields:
    - `Logger`: The logger
    - `TimeBetweenTasks`: The time between task creations
    - `QuorumThresholdPercentage`: The percentage of the total quorum stake needed by the signers to make the aggregated response valid
    - `QuorumNumbers`: The numbers of the quorums required to respond to tasks for the response to be valid

      ```go
        taskSpammerConfig := taskspammer.Config{
          Logger:                     logger,
          TimeBetweenTasks:           10 * time.Second,
          QuorumThresholdPercentage:  100,
          QuorumNumbers:              []uint8{0},
        }
      ```

3. **Build the Task Spammer**: Use the `NewTaskSpammer` function to build the task spammer

    ```go
      taskSpammer, _ := taskspammer.NewTaskSpammer(taskCreator, taskSpammerConfig)
    ```

4. **Create the iterator**: Define an iterator that creates appropriate input values for your specific AVS
   - The input values will be passed to the `CreateNewTask` function on the `TaskSpammer` struct

      ```go
        func NewNumberToSquareSequence() iter.Seq[*big.Int] {
          acc := big.NewInt(1)
          delta := big.NewInt(1)
          return func(yield func(*big.Int) bool) {
            for {
              if !yield(acc) {
                break
              }
              acc.Add(acc, delta)
            }
          }
        }
        inputGen := NewNumberToSquareSequence()
      ```

5. **Create the Task Spammer**: Create the task spammer using the `NewTaskSpammer` function

    ```go
      taskSpammer, _ := taskspammer.NewTaskSpammer(taskCreator, taskSpammerConfig)
    ```

6. **Start the Task Spammer**: Call the `Start` method to start the task spammer

    ```go
      taskSpammer.Start(context.Background(), inputGen)
    ```

## Examples

Here are some examples of task spammer implementations:

- [Incredible Squaring](https://github.com/Layr-Labs/eigensdk-go/blob/v2-dev-1/examples/incredible-squaring/task-spammer/main.go)
- [Incredible Dot Product](https://github.com/Layr-Labs/eigensdk-go/blob/v2-dev-1/examples/incredible-dot-product/taskc-spammer/main.go)
- [Awesome Vault Service](https://github.com/Layr-Labs/eigensdk-go/blob/v2-dev-1/examples/awesome-vault-service/task-spammer/main.go)

