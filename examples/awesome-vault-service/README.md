
# Awesome Vault Service example

This example shows how to create a more complex AVS based on the SDK structure. The AVS shown here is a local Redis version that allows storing key-value pairs in the Task Manager. In this sense, the task for the operators to complete is storing the pair into a local Merkle tree of pairs and recomputing the root hash of the tree, submitting it as the response value to the Task Manager on-chain contract.

## Structure

### Business logic in the entities

- Aggregator: The aggregator does not have much business logic, since the task processor implementation lies in the Indexing Task Processor implementation, which can be seen as the default one. If wanted to create your task processor, you can base it on the ITP implementation, and change what you need.
- Challenger: The challenger business logic lies in Task response validation. To validate the response, the challenger first calculates the response with the same function as the operator (the ComputeResponse method of the VaultServiceResponseCalculator) and then compares it with the received response, raising a challenge if they differ.
- Operator: The operator responds to tasks using the VaultServiceResponseCalculator struct, which has a ComputeResponse method (satisfying the ResponseCalculator interface). We decided to create a response calculator struct because in cases like this, where the operator should keep a state to respond to tasks, the struct allows us to save that state in the struct attributes.
- Task spammer: The task spammer logic lies in the sequence that generates the numbers pulled by the spammer at the SDK level. More specifically, the input passed to the spammer when it makes a pull is passed to the yield function inside the closure.
