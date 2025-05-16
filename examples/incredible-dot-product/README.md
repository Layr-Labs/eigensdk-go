# Incredible Dot Product example

This example proposes a more complex AVS than the proposed incredible squaring. The input type handled is a struct that contains two vectors, representing two points, that should be processed making the dot product, being the output type a big int, representing the result of the product. In this sense, the task for the operators to complete is executing the dot product between the two points, and returning the result as the response value submitted to the Task Manager on-chain contract.

## Structure

### Types

The task type of the solidity contract is the following:

``` solidity
    struct DotProductInput {
        uint256[] X;
        uint256[] Y;
    }

    struct Task {
        DotProductInput pointsToMultiply;
        uint32 taskCreatedBlock;
        bytes quorumNumbers;
        uint32 quorumThresholdPercentage;
    }
```

The input is the pair of vectors X and Y, both represented by an `uint256` array.

The task response type is:

``` solidity
    struct TaskResponse {
        uint32 referenceTaskIndex;
        uint256 result;
    }
```

The `result` field represents the result of the dot product operation between the two received points.

### Specific business logic

For computing the response, the challenger and operator use the ComputeResponse method of the DotProduct ResponseCalculator. This DotProduct response calculator is made with the builder provided by the SDK for non-state-saving calculators, that receives a function to calculate the logic.

In the specific case of the dot product the function receives an input with two vectors, and performs the dot product of them, returning the result.

To create the sequence that passes input values to the task spammer, we use the sequence generator in the task manager main (in `examples/incredible-dot-product/task-spammer/main.go`), that creates a sequence that on each iterarion advances on 1 and gives as input a fixed pair of vectors defined from the iteration number.

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

By default, the `start-operator` command will also register the operator. To disable this, set `RegisterOperatorOnStartup` to false in the operator main.

The operator will produce invalid results often because the function passed to it has a failure probability. These failures result in slashing once they're challenged.
To see this in action, start the challenger with:

```bash
make start-challenger
```

To start the cycle, start the task spammer:

``` bash
make start-task-spammer
```
