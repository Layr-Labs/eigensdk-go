# Task Manager

This package provides utilities for interacting with a user-defined `TaskManager` binding.

## Overview

### Task Manager interface

A Task Manager abstracts the on-chain `TaskManager` binding, letting the SDK to create tasks, submit responses and raise challenges through a `TaskManager` interface without worrying to specific implementation details.

Implementing the [`TaskManager`] interface requires that your on-chain contract expose some specific Solidity types and functions and provide some specific parameters. Our SDK will wire up all the rest; you need to supply your generic `Input`/`Output`, the task manager contract and ABI, a transaction manager and an ethereum client.the selectors for the new task event and the task responded event via [`TaskManagerDefs`]. You can refer to [this contract](https://github.com/Layr-Labs/eigensdk-go/blob/v2-dev-2/examples/incredible-squaring/contracts/src/IIncredibleSquaringTaskManager.sol) as an example.

**Required Solidity Structs**:

```solidity
struct Task {
    Input input; // user-defined input type
    uint32 taskCreatedBlock;
    bytes  quorumNumbers;
    uint32 quorumThresholdPercentage;
}

struct TaskResponse {
    uint32 referenceTaskIndex;
    Output response; // user-defined output type
}

struct TaskResponseMetadata {
    uint32 taskRespondedBlock;
    bytes32 hashOfNonSigners;
}
```

**Required Solidity Functions**:

```solidity
function createNewTask(
    Input    calldata input,
    uint32   quorumThresholdPercentage,
    bytes    calldata quorumNumbers
) external;

function respondToTask(
    Task calldata task,
    TaskResponse calldata taskResponse,
    NonSignerStakesAndSignature memory nonSignerStakesAndSignature
) external;

function raiseAndResolveChallenge(
    Task calldata task,
    TaskResponse calldata taskResponse,
    TaskResponseMetadata calldata taskResponseMetadata,
    BN254.G1Point[] calldata pubkeysOfNonSigningOperators
) external;
```

**Required Solidity Events**:

```solidity
event NewTaskCreated(uint32 indexed taskIndex, Task task);
event TaskResponded(TaskResponse taskResponse, TaskResponseMetadata taskResponseMetadata);
```

Once those are in place, you have to define the AVS Input/Output and create the Implement [`TaskManagerDefs`] to choose your `Input`/`Output` and event selectors.

``` go
    taskManagerAbi, err := avtaskmanager.ContractAwesomeVaultTaskManagerMetaData.GetAbi()
    if err != nil {
        logger.Errorf("Failed to get task manager abi: %w", err)
        return
    }

    ethHttpUrl := "http://localhost:8545"
    ethClient, err := ethclient.Dial(ethHttpUrl)
    if err != nil {
        logger.Errorf("Failed to dial ethclient: %w", err)
        return
    }

    taskManagerAddr := gethcommon.HexToAddress("0x7bc06c482dead17c0e297afbc32f6e63d3846650")

    challengerPrivateKey := "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
    ecdsaPrivateKey, err := crypto.HexToECDSA(challengerPrivateKey)
    if err != nil {
        logger.Errorf("Failed to create ecdsa private key: %w", err)
        return
    }

    txMgr, err := txmgr.NewSimpleTxManagerFromPrivateKey(logger, ethClient, ecdsaPrivateKey)
    if err != nil {
        logger.Errorf("Failed to create tx manager from private key: %w", err)
        return
    }

    challengerRaiser, err := taskmanager.NewTaskManagerFromAbi[examplecommon.TaskInput, [32]byte](taskManagerAddr, taskManagerAbi, txMgr, ethClient)
    if err != nil {
        logger.Errorf("Failed to create challenger raiser: %w", err)
        return
    }
```
