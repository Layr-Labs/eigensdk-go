# Incredible Squaring example

This example is a basic proposal of AVS, being input and output type a big int, representing the number to be squared and the number squared. In this sense, the task for the operators to complete is squaring the received number, and returning the result of the operation as the response value submitted to the Task Manager on-chain contract.

## Structure

### Business logic in the entities

- Aggregator: The aggregator does not have much business logic, since the task processor implementation lies in the Indexing Task Processor one, which can be seen as the default. If wanted to create your task processor, you can base it on the ITP implementation, and change what you need.
- Challenger: The challenger business logic lies in Task response validation. To validate the response, the challenger first calculates the response with the same function as the operator (the ComputeResponse method of the Squaring ResponseCalculator) and then compares it with the received response, raising a challenge if they differ.
- Operator: The operator responds to tasks using the Squaring ResponseCalculator struct, which has a ComputeResponse method (satisfying the ResponseCalculator interface). The Squaring response calculator is made with the builder provided by the SDK for non state-saving calculators.
- Task spammer: The task spammer logic lies in the sequence that generates the values pulled by the spammer at the SDK level. More specifically, the input passed to the spammer when it makes a pull is passed to the yield function inside the closure.

## How to run

This example cannot be run since does not have the contracts included, but you can see this example running on the updated-v2-dev-1 branch of [this fork](https://github.com/Layr-Labs/incredible-squaring-avs), which is up to date with the v2-dev-1 branch.
