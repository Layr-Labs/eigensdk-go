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

### Business logic in the entities

- Aggregator: The aggregator does not have much business logic, since the task processor implementation lies in the Indexing Task Processor one, which can be seen as the default. If wanted to create your task processor, you can base it on the ITP implementation, and change what you need.
- Challenger: The challenger business logic lies in Task response validation. To validate the response, the challenger first calculates the response with the same function as the operator (the ComputeResponse method of the Squaring ResponseCalculator) and then compares it with the received response, raising a challenge if they differ.
- Operator: The operator responds to tasks using the Squaring ResponseCalculator struct, which has a ComputeResponse method (satisfying the ResponseCalculator interface). The Squaring response calculator is made with the builder provided by the SDK for non-state-saving calculators.
- Task spammer: The task spammer logic lies in the sequence that generates the values pulled by the spammer at the SDK level. More specifically, the input passed to the spammer when it makes a pull is passed to the yield function inside the closure.

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
