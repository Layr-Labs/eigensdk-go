
# Awesome Vault Service example

This example shows how to create a more complex AVS based on the SDK structure. The AVS shown here is a local Redis version that allows storing key-value pairs in the Task Manager. In this sense, the task for the operators to complete is storing the pair into a local Merkle tree of pairs and recomputing the root hash of the tree, submitting it as the response value to the Task Manager on-chain contract.

## Structure

### Types

The task type of the solidity contract is the following:

``` solidity
    struct TaskInput {
        string key;
        string value;
    }

    struct Task {
        TaskInput input;
        uint32 taskCreatedBlock;
        bytes quorumNumbers;
        uint32 quorumThresholdPercentage;
    }
```

We can see the input is the key-value pair mentioned at first.

The task response type is:

``` solidity
    struct TaskResponse {
        uint32 referenceTaskIndex;
        bytes32 result;
    }
```

The result field represents the root hash of the stored merkle tree.

### Business logic in entities

- Aggregator: The aggregator does not have much business logic, since the task processor implementation lies in the Indexing Task Processor implementation, which can be seen as the default one. If wanted to create your task processor, you can base it on the ITP implementation, and change what you need.
- Challenger: The challenger business logic lies in Task response validation. To validate the response, the challenger first calculates the response with the same function as the operator (the ComputeResponse method of the VaultServiceResponseCalculator) and then compares it with the received response, raising a challenge if they differ.
  - For response validation, there should be an additional check to verify that the operator has uploaded the key-value pair, but for that a proof should be added to the challenge cycle.
- Operator: The operator responds to tasks using the VaultServiceResponseCalculator struct, which has a ComputeResponse method (satisfying the ResponseCalculator interface). We decided to create a response calculator struct because in cases like this, where the operator should keep a state to respond to tasks, the struct allows us to save that state in the struct attributes.
- Task spammer: The task spammer logic lies in the sequence that generates the numbers pulled by the spammer at the SDK level. More specifically, the input passed to the spammer when it makes a pull is passed to the yield function inside the closure.

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
