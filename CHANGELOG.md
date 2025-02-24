# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

Each version will have a separate `Breaking Changes` section as well. To describe in how to upgrade from one version to another if needed

## [Unreleased]

### Added 🎉

* feat: new BLS aggregation service interface by @maximopalopoli in <https://github.com/Layr-Labs/eigensdk-go/pull/578>
  * The new interface implies starting the service before using it, interact with it using a handler and receiving the aggregated responses in a separate channel.
  * An example using the interface is:

    ```Go
    // initialize service
    blsAgg := NewBlsAggregatorBuilder(fakeAvsRegistryService, hashFunction, logger)
    handler, aggResponsesC := blsAgg.Start()

    // Initialize task
    metadata := NewTaskMetadata(taskIndex, blockNum, quorumNumbers, quorumThresholdPercentages, tasksTimeToExpiry)
    err := handler.InitializeNewTask(metadata)

    // Process signature
    taskSignature := NewTaskSignature(taskIndex, taskResponse, blsSig, testOperator1.OperatorId)
    err = handler.ProcessNewSignature(
      context.Background(),
      taskSignature,
    )

    // Receive responses
    aggregationServiceResponse := <-aggResponsesC
    ```

* Added field `DontUseAllocationManager` to `BuildAllConfig` in [#580](https://github.com/Layr-Labs/eigensdk-go/pull/580)

### Changed

* Fixed BLS aggregation for multiple quorums by @TomasArrachea in [#394](https://github.com/Layr-Labs/eigensdk-go/pull/394)
* fix: change requested pr url in changelog's workflow by @maximopalopoli in [#575](https://github.com/Layr-Labs/eigensdk-go/pull/575)
* chore: use utils WrapError function instead of fmt's Errorf by @pablodeymo and @maximopalopoli in <https://github.com/Layr-Labs/eigensdk-go/pull/579>
* fix: propagate DontUseAllocationManager flag between builders configs by @maximopalopoli in [#581](https://github.com/Layr-Labs/eigensdk-go/pull/581)

### Breaking changes

* refactor: encapsulate parameters into `TaskSignature` in [#487](https://github.com/Layr-Labs/eigensdk-go/pull/487)

  * Introduced `TaskSignature` struct to encapsulate parameters related to task signatures:
  * Updated `ProcessNewSignature` to accept a `TaskSignature` struct instead of multiple parameters.

    ```go
    // BEFORE
    blsAggServ.ProcessNewSignature(
        context.Background(),
        taskIndex,
        taskResponse,
        blsSigOp1,
        testOperator1.OperatorId,
    )
    
    // AFTER
    taskSignature := NewTaskSignature(taskIndex, taskResponse, blsSig, testOperator1.OperatorId)

    blsAggServ.ProcessNewSignature(
        context.Background(),
        taskSignature,
    )
    ```
  
* refactor: update interface on `bls aggregation` in [#485](https://github.com/Layr-Labs/eigensdk-go/pull/485).
  * Introduces a new struct `TaskMetadata` with a constructor `NewTaskMetadata` to initialize a new task and a method `WithWindowDuration` to set the window duration.
  * Refactors `InitializeNewTask` and `singleTaskAggregatorGoroutineFunc` to accept a `TaskMetadata` struct instead of multiple parameters.

    ```go
    // BEFORE
    blsAggServ := NewBlsAggregatorService(fakeAvsRegistryService, hashFunction, logger)

    blsAggServ.InitializeNewTask(
        taskIndex,
        blockNum,
        quorumNumbers,
        quorumThresholdPercentages,
        tasksTimeToExpiry,
    )
    
    // AFTER
    blsAggServ := NewBlsAggregatorService(fakeAvsRegistryService, hashFunction, logger)

    metadata := NewTaskMetadata(taskIndex, blockNum, quorumNumbers, quorumThresholdPercentages, tasksTimeToExpiry)
    blsAggServ.InitializeNewTask(metadata)
    ```

  * Removes `InitializeNewTaskWithWindow` since `windowDuration` can now be set in `TaskMetadata`.

    ```go
    // BEFORE
    blsAggServ := NewBlsAggregatorService(fakeAvsRegistryService, hashFunction, logger)
    err = blsAggServ.InitializeNewTaskWithWindow(
        taskIndex,
        blockNum,
        quorumNumbers,
        quorumThresholdPercentages,
        timeToExpiry,
        windowDuration,
    )

    // AFTER
    blsAggServ := NewBlsAggregatorService(fakeAvsRegistryService, hashFunction, logger)

    metadata := NewTaskMetadata(
        taskIndex,
        blockNum,
        quorumNumbers,
        quorumThresholdPercentages,
        tasksTimeToExpiry,
    ).WithWindowDuration(windowDuration)
    blsAggServ.InitializeNewTask(metadata)
    ```

### Removed

------------

Changes made in the v0.1.X versions weren't tracked by this changelog.
