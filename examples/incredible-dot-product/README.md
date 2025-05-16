# Incredible Dot Product example

This example proposes a more complex AVS than the proposed on incredible aquaring. The input type handled is a struct that contains two vectors, representing two points, that shpuld be processed making the dot product, being the output type a big int, representing the result of the product. In this sense, the task for the operators to complete is executing the dot product between the two points, and returning the result as the response value submitted to the Task Manager on-chain contract.

## Structure

## Types

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

We can see the input is the pair of points X and Y, both represented by an uint256 array.

The task response type is:

``` solidity
    struct TaskResponse {
        uint32 referenceTaskIndex;
        uint256 result;
    }
```

The result the result of the dot product operation between the two received points.

### Business logic in the entities

- Aggregator: The aggregator does not have much business logic, since the task processor implementation lies in the Indexing Task Processor one, which can be seen as the default. If wanted to create your task processor, you can base it on the ITP implementation, and change what you need.
- Challenger: The challenger business logic lies in Task response validation. To validate the response, the challenger first calculates the response with the same function as the operator (the ComputeResponse method of the DotProduct ResponseCalculator) and then compares it with the received response, raising a challenge if they differ.
- Operator: The operator responds to tasks using the DotProduct ResponseCalculator struct, which has a ComputeResponse method (satisfying the ResponseCalculator interface). The DotProduct response calculator is made with the builder provided by the SDK for non state-saving calculators.
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

By default, the `start-operator` command will also register the operator. To disable this, set `RegisterOperatorOnStartup` to false in the operator main.

The operator will produce invalid results often, because the function passed to it has a failure probability. These failures result in slashing once they're challenged.
To see this in action, start the challenger with:

```bash
make start-challenger
```

To start the cycle, start the task spammer:

``` bash
make start-task-spammer
```
