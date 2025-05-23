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

The input is the key-value pair mentioned at first.

The task response type is:

``` solidity
    struct TaskResponse {
        uint32 referenceTaskIndex;
        bytes32 result;
    }
```

The `result` field represents the root hash of the stored Merkle tree.

### Specific business logic

The challenger and operator use the `VaultServiceResponseCalculator.ComputeResponse()` method for computing responses. You can see the specific implementation in `examples/awesome-vault-service/common/response_calculator.go`.

The method inserts the key-value pair into the vaults array representing the Merkle tree. Then, computes the tree root hash and returns it as the task response value.

We decided to create a response calculator struct because in cases like this, where the operator should keep a state to respond to tasks, the struct allows us to save that state in the struct attributes.

For response validation, there should be an additional check to verify that the operator has uploaded the key-value pair, but for that, proof telling the operator has set the value should be added to the challenge cycle.

To create the sequence that passes input values to the task spammer, we use the sequence generator in the task manager main (in `examples/awesome-vault-service/task-spammer/main.go`), that creates a sequence that on each iteration advances on 1 and gives as input a fixed key-pair defined from the iteration number.

### Moving parts

The system flow is composed of 4 binaries:

- Aggregator: Aggregates responses from operators and sends the aggregated responses to the task manager contract. [Link to use example](https://github.com/Layr-Labs/eigensdk-go/blob/v2-dev-2/examples/incredible-dot-product/aggregator/main.go)
- Challenger: If the aggregated responses are incorrect, raises a challenge to the task manager contract. [Link to use example](https://github.com/Layr-Labs/eigensdk-go/blob/v2-dev-2/examples/incredible-dot-product/challenger/main.go)
- Operator: Responds to tasks and sends them to the aggregator. [Link to use example](https://github.com/Layr-Labs/eigensdk-go/blob/v2-dev-2/examples/incredible-dot-product/operator/main.go)
- Task Spammer: Creates new tasks and sends them to the task manager contract. [Link to use example](https://github.com/Layr-Labs/eigensdk-go/blob/v2-dev-2/examples/incredible-dot-product/task-spammer/main.go)

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
