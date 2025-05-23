# Incredible Squaring example

This example is a basic proposal of AVS, where the input and output type are `uint256` values, representing the number to be squared and the number squared. In this sense, the task for the operators to complete is squaring the received number, and returning the result of the operation as the response value submitted to the Task Manager on-chain contract.

## Structure

### Types

The task type of the solidity contract is the following:

``` solidity
    struct Task {
        uint256 numberToBeSquared;
        uint32 taskCreatedBlock;
        bytes quorumNumbers;
        uint32 quorumThresholdPercentage;
 }
```

The input is an `uint256` representing the number to be squared.

The task response type is:

``` solidity
    struct TaskResponse {
        uint32 referenceTaskIndex;
        uint256 numberSquared;
 }
```

The `numberSquared` field represents the result of the squaring operation with the received number to square.

### Specific business logic

The challenger and operator use the `SquaringResponseCalculator.ComputeResponse()` method for computing responses.  This Squaring response calculator is made with the SDK-provided builder for non-state-saving calculators, which receives a function for calculating the logic.

In the specific case of squaring, the function receives an input of an uint256, and performs the squaring of it, returning the result.

To create the sequence that passes input values to the task spammer, we use the sequence generator in the task manager main (in `examples/incredible-squaring/task-spammer/main.go`) which creates a sequence that advances on 1 and gives the iteration number as input on each iteration.

## How to run

This simple session illustrates the basic flow of the AVS:

Start anvil in a separate terminal:

```bash
anvil
```

Deploy contracts, set UAM permissions, and create a quorum in a single command:

```bash
make deploy-all
```

Start the aggregator:

```bash
make start-aggregator
```

Start the operator:

```bash
make start-operator
```

By default, the `start-operator` command will register the operator. To disable this, set `RegisterOperatorOnStartup` to false in the operator main.

The operator will produce invalid results often because the function passed to it has a failure probability. These failures result in slashing once they're challenged.
To see this in action, start the challenger with:

```bash
make start-challenger
```

To start the cycle, start the task spammer:

``` bash
make start-task-spammer
```

## Link to implementations

To see the specific implementation on each member of this example, you can see:

- [Aggregator implementation](https://github.com/Layr-Labs/eigensdk-go/blob/v2-dev-2/examples/incredible-squaring/aggregator/main.go)
- [Challenger implementation](https://github.com/Layr-Labs/eigensdk-go/blob/v2-dev-2/examples/incredible-squaring/challenger/main.go)
- [Operator implementation](https://github.com/Layr-Labs/eigensdk-go/blob/v2-dev-2/examples/incredible-squaring/operator/main.go)
- [Task Spammer implementation](https://github.com/Layr-Labs/eigensdk-go/blob/v2-dev-2/examples/incredible-squaring/task-spammer/main.go)
