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

* Added field `DontUseAllocationManager` to `BuildAllConfig` by @MegaRedHand in [#580](https://github.com/Layr-Labs/eigensdk-go/pull/580)
* Added `AnvilC` field to `clients.Clients` struct by @maximopalopoli in [#585](https://github.com/Layr-Labs/eigensdk-go/pull/585)
* Added `IsOperatorRegisteredWithAvs`, `GetAVSRegistrar` methods to elcontracts chain reader and and `SetAVSRegistrar` to chain writer by @maximopalopoli in [#585](https://github.com/Layr-Labs/eigensdk-go/pull/585)
  * An example for `IsOperatorRegisteredWithAvs` would be the following:
    ```go
      // Given an operator registered to a M2 Quorum
      isOperator, err := clients.ElChainReader.IsOperatorRegisteredWithAvs(ctx, operatorAddress, avsAddress)
      assert.NoError(t, err)
      assert.Equal(t, isOperator, true) // Assuming is registered
    ```
  * An example for `GetAVSRegistrar` would be the following:
    ```go
      avsRegistrar, err := clients.ElChainReader.GetAVSRegistrar(context.Background(), avsAddress)
      assert.NoError(t, err)
    ```
  * An example for `SetAVSRegistrar` would be the following:
    ```go
      // Usually the AVSRegistrar is the registryCoordinator
      receipt, err := clients.ElChainWriter.SetAVSRegistrar(context.Background(), avsAddress, contractAddrs.RegistryCoordinator, true)
      require.NoError(t, err)
      require.Equal(t, gethtypes.ReceiptStatusSuccessful, receipt.Status)
    ```
* Added `RegisterAsOperatorPreSlashing` to register an operator in M2 workflows by @maximopalopoli in [#595](https://github.com/Layr-Labs/eigensdk-go/pull/595)
  * A use example would be the following:
    ```go
      operator :=
        types.M2Operator{
          Address:                   fundedAccount,
          DelegationApproverAddress: "0xd5e099c71b797516c10ed0f0d895f429c2781142",
          StakerOptOutWindowBlocks:  100,
          MetadataUrl:               "https://madhur-test-public.s3.us-east-2.amazonaws.com/metadata.json",
        }

      receipt, err := clients.ElChainWriter.RegisterAsOperatorPreSlashing(context.Background(), operator, true)
      assert.NoError(t, err)
      assert.True(t, receipt.Status == 1)
    ```
  This PR also introduces the `M2Operator` type

* Added support for registering operators in operator sets with churn approval in [#596](https://github.com/Layr-Labs/eigensdk-go/pull/596)
  * We added the fields `ChurnApprovalEcdsaPrivateKey` and `OperatorKickParams` to `elcontracts.RegistrationRequest`. Specifying the first one makes the `ChainWriter.RegisterForOperatorSets` function sign a churn approval for the operator registration, making the AVS eject operators specified by the other field to make space for the registering operator.

    ```go
    request := elcontracts.RegistrationRequest{
      // ...old fields are required...
      ChurnApprovalEcdsaPrivateKey: /* ECDSA key of the AVS churn approver */,
      OperatorKickParams:  /* which operators to kick for each registering quorum */,
    }
    receipt, err := chainWriter.RegisterForOperatorSets(
      context.Background(),
      registryCoordinatorAddress,
      request,
    )
    ```

* Bump middleware contracts version up to [v1.3.0-rc.0](https://github.com/Layr-Labs/eigenlayer-middleware/releases/tag/v1.3.0-rc.0) by @maximopalopoli in [#605](https://github.com/Layr-Labs/eigensdk-go/pull/605).

* Bump middleware contracts version up to [v1.4.0-testnet-holesky](https://github.com/Layr-Labs/eigenlayer-middleware/releases/tag/v1.4.0-testnet-holesky) by @maximopalopoli in [#606](https://github.com/Layr-Labs/eigensdk-go/pull/606).

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
* In elcontracts, `ChainReader.IsOperatorRegisteredWithOperatorSet` no longer queries the `AVSDirectory`, and so now only works for operator sets by @maximopalopoli in [#585](https://github.com/Layr-Labs/eigensdk-go/pull/585)
  * To query if an operator is registered to an M2 quorum you should now use `chainReader.IsOperatorRegisteredWithAvs`, which queries the `AVSDirectory`.
* `egnaddrs` utility now works with slashing release middleware contracts and, in that case, the returned service manager will be the zero address, unless the `--service-manager` is specified by @maximopalopoli in [#585](https://github.com/Layr-Labs/eigensdk-go/pull/585).
* The `elcontracts.NewChainWriter` function now receives an additional parameter, the delegation manager address by @maximopalopoli in [#595](https://github.com/Layr-Labs/eigensdk-go/pull/595).

* Renamed `SetAccountIdentifier` to `SetAvs` [#597](https://github.com/Layr-Labs/eigensdk-go/pull/597)
  * The underlying call was renamed in [the v1.1.1 eigenlayer-middleware release](https://github.com/Layr-Labs/eigenlayer-middleware/releases/tag/v1.1.1-testnet-slashing).

* Added bindings for M2 contracts by @maximopalopoli in [595](https://github.com/Layr-Labs/eigensdk-go/pull/595)
  * The M2 bindings are available at `github.com/Layr-Labs/eigensdk-goM2-contracts/bindings`. Users that wish to use an old version of a binding should change the `contracts` part of the import for `M2-contracts`:
    ```go
      // slashing bindings import
      // import "github.com/Layr-Labs/eigensdk-go/contracts/bindings/RegistryCoordinator"

      // M2 bindings import
      import "github.com/Layr-Labs/eigensdk-go/M2-contracts/bindings/RegistryCoordinator"
    ```

* Bumped up slashing bindings to [v1.3.0-rc.0](https://github.com/Layr-Labs/eigenlayer-contracts/releases/tag/v1.3.0) by @maximopalopoli in [#602](https://github.com/Layr-Labs/eigensdk-go/pull/602)
  * Introduces two new functions for `elcontracts.chainReader`: `IsOperatorSlashable` and `GetAllocatedStake`. The first can be used this way:
    ``` Go
      isSlashable, err := chainReader.IsOperatorSlashable(
        context.Background(),
        operatorAddress,
        operatorSet,
      )
      require.NoError(t, err)
    ```

    A use example for the second one would be the following:
    ``` Go
      allocatedStakes, err = chainReader.GetAllocatedStake(
        context.Background(),
        operatorSet,
        operatorAddresses,
        strategyAddresses,
      )
      require.NoError(t, err)
    ```


### Removed
* Removed `IsOperatorSetQuorum` method of avsRegistry chain reader by @maximopalopoli in [#585](https://github.com/Layr-Labs/eigensdk-go/pull/585)
  * This function was removed in [the v1.1.1 eigenlayer-middleware release](https://github.com/Layr-Labs/eigenlayer-middleware/releases/tag/v1.1.1-testnet-slashing).

------------

Changes made in the v0.1.X versions weren't tracked by this changelog.
