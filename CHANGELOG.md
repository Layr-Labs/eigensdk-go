# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

Each version will have a separate `Breaking Changes` section as well. To describe in how to upgrade from one version to another if needed

## [Unreleased]

### Security 🔒

### Added 🎉

### Breaking Changes 🛠

### Deprecated ⚠️

* Deprecated `RoundUpDivideBig` from the `utils` module in [#340](https://github.com/Layr-Labs/eigensdk-go/pull/340)

### Removed 🗑

### Documentation 📚

* Fixed broken link in documentation and updated "Modules" section in [#641](https://github.com/Layr-Labs/eigensdk-go/pull/641)

### Other Changes

* chore: bump dependencies in [#739](https://github.com/Layr-Labs/eigensdk-go/pull/739)
  * `go.mod`
    * github.com/consensys/gnark-crypto from `v0.16.0` to `v0.18.0`
    * github.com/lmittmann/tint from `v1.0.4` to `v1.1.2`
    * github.com/testcontainers/testcontainers-go from `v0.35.0` to `v0.37.0`
    * github.com/urfave/cli/v2 from `v2.27.5` to `v2.27.7`
    * go.uber.org/mock from `v0.4.0` to `v0.5.2`
    * golang.org/x/crypto from `v0.35.0` to `v0.39.0`
    * golang.org/x/sync from `v0.11.0` to `v0.15.0`

  * `signer/go.mod`
    * go from `1.21.13` to `1.23.0`
    * toolchain from `go1.22.5` to `go1.23.4`
    * github.com/consensys/gnark-crypto from `v0.12.1` to `v0.16.0`
    * github.com/stretchr/testify from `v1.9.0` to `v1.10.0`

* added rewards utilities integration test by @maximopalopoli in [#608](https://github.com/Layr-Labs/eigensdk-go/pull/608)
* fixed expired timer handling in bls aggregation service in [#616](https://github.com/Layr-Labs/eigensdk-go/pull/616)
* use the eigenlayer contracts repository bindings instead of creating them by @maximopalopoli in [#668](https://github.com/Layr-Labs/eigensdk-go/pull/668)
* Added unit tests to the `utils` module by @estensen in [#340](https://github.com/Layr-Labs/eigensdk-go/pull/340)

## [0.3.0] - 2025-03-19

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

* Added field `DontUseAllocationManager` to bindings configuration structs
  * feat: make SDK compatible with mainnet contracts by @MegaRedHand in <https://github.com/Layr-Labs/eigensdk-go/pull/464>
  * Added field `DontUseAllocationManager` to `BuildAllConfig` by @MegaRedHand in [#580](https://github.com/Layr-Labs/eigensdk-go/pull/580)
  * fix: propagate `DontUseAllocationManager` flag between builders configs by @maximopalopoli in [#581](https://github.com/Layr-Labs/eigensdk-go/pull/581)
* feat: add missing contract addresses on `BuildAllConfig` by @MegaRedHand in <https://github.com/Layr-Labs/eigensdk-go/pull/558>
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

* Introduces two new functions for `elcontracts.chainReader` from core contracts [v1.3.0-rc.0 release](https://github.com/Layr-Labs/eigenlayer-contracts/releases/tag/v1.3.0) by @maximopalopoli in [#602](https://github.com/Layr-Labs/eigensdk-go/pull/602)
  * `IsOperatorSlashable`

    ```go
      isSlashable, err := chainReader.IsOperatorSlashable(
        context.Background(),
        operatorAddress,
        operatorSet,
      )
      require.NoError(t, err)
    ```

  * `GetAllocatedStake`

    ```go
      allocatedStakes, err = chainReader.GetAllocatedStake(
        context.Background(),
        operatorSet,
        operatorAddresses,
        strategyAddresses,
      )
      require.NoError(t, err)
    ```

* feat: added `setMinimumStakeForQuorum` by @Sidu28 in <https://github.com/Layr-Labs/eigensdk-go/pull/554>
* feat: added `StakeRegistryStorage` functions by @Sidu28 in <https://github.com/Layr-Labs/eigensdk-go/pull/555>
* feat: added `BlsApkRegistryStorage` functions by @Sidu28 in <https://github.com/Layr-Labs/eigensdk-go/pull/556>
* feat: added `AllocationManagerStorage` functions by @Sidu28 in <https://github.com/Layr-Labs/eigensdk-go/pull/563>
* feat: added `DelegationManagerStorage` functions by @Sidu28 in <https://github.com/Layr-Labs/eigensdk-go/pull/564>
* feat: added `ServiceManagerBase.getRestakeableStrategies` wrapper by @MegaRedHand in <https://github.com/Layr-Labs/eigensdk-go/pull/559>
* feat: added `ServiceManagerBase.createOperatorDirectedAVSRewardsSubmission` wrapper by @MegaRedHand in <https://github.com/Layr-Labs/eigensdk-go/pull/>561
* feat: added new method `registerOperatorWithChurn` by @damiramirez in <https://github.com/Layr-Labs/eigensdk-go/pull/566>
* feat: added method `ModifyStrategyParams` by @damiramirez in <https://github.com/Layr-Labs/eigensdk-go/pull/570>
* feat: added new method `AddStrategies` by @damiramirez in <https://github.com/Layr-Labs/eigensdk-go/pull/568>
* feat: added `ServiceManagerBase.getOperatorRestakedStrategies` wrapper by @MegaRedHand in <https://github.com/Layr-Labs/eigensdk-go/pull/560>
* feat: added method `RemoveStrategies` by @damiramirez in <https://github.com/Layr-Labs/eigensdk-go/pull/569>
* feat: added `updateAVSMetadataURI` wrapper by @MegaRedHand in <https://github.com/Layr-Labs/eigensdk-go/pull/553>
* feat: added `ServiceManagerBase.createAVSRewardsSubmission` wrapper by @MegaRedHand in <https://github.com/Layr-Labs/eigensdk-go/pull/557>
* feat: Add `RewardsCoordinatorStorage` functions by @Sidu28 in <https://github.com/Layr-Labs/eigensdk-go/pull/565>

### Changed

* Fixed BLS aggregation for multiple quorums by @TomasArrachea in [#394](https://github.com/Layr-Labs/eigensdk-go/pull/394)
* fix: change PR url in Changelog workflow by @maximopalopoli in [#575](https://github.com/Layr-Labs/eigensdk-go/pull/575)
* chore: use utils WrapError function instead of fmt's Errorf by @pablodeymo and @maximopalopoli in <https://github.com/Layr-Labs/eigensdk-go/pull/579>

### Breaking Changes 🛠

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

* Updated bindings to the [v1.4.0-testnet-holesky](https://github.com/Layr-Labs/eigenlayer-middleware/releases/tag/v1.4.0-testnet-holesky) release of middleware.
  * Bump middleware contracts version up to [v1.3.0-rc.0](https://github.com/Layr-Labs/eigenlayer-middleware/releases/tag/v1.3.0-rc.0) by @maximopalopoli in [#605](https://github.com/Layr-Labs/eigensdk-go/pull/605).
  * Bump middleware contracts version up to [v1.4.0-testnet-holesky](https://github.com/Layr-Labs/eigenlayer-middleware/releases/tag/v1.4.0-testnet-holesky) by @maximopalopoli in [#606](https://github.com/Layr-Labs/eigensdk-go/pull/606).

### Removed 🗑

* Removed `IsOperatorSetQuorum` method of avsRegistry chain reader by @maximopalopoli in [#585](https://github.com/Layr-Labs/eigensdk-go/pull/585)
  * This function was removed in [the v1.1.1 eigenlayer-middleware release](https://github.com/Layr-Labs/eigenlayer-middleware/releases/tag/v1.1.1-testnet-slashing).

### Other Changes

* chore: remove duplicated function by @MegaRedHand in <https://github.com/Layr-Labs/eigensdk-go/pull/426>
* fix: respect context in new signature call by @MegaRedHand in <https://github.com/Layr-Labs/eigensdk-go/pull/502>
* ci: add job to enforce updates to the changelog by @ricomateo in <https://github.com/Layr-Labs/eigensdk-go/pull/483>
* test: `avsregistry` add new tests cases for reader methods by @damiramirez in <https://github.com/Layr-Labs/eigensdk-go/pull/458>
* chore(deps): bump github.com/ethereum/go-ethereum from 1.14.0 to 1.15.0 by @dependabot in <https://github.com/Layr-Labs/eigensdk-go/pull/543>
* chore(deps): bump github.com/ethereum/go-ethereum from 1.14.0 to 1.14.13 in /signer by @dependabot in <https://github.com/Layr-Labs/eigensdk-go/pull/490>
* chore(deps): bump github.com/testcontainers/testcontainers-go from 0.30.0 to 0.35.0 by @dependabot in <https://github.com/Layr-Labs/eigensdk-go/pull/590>
* chore(deps): bump github.com/urfave/cli/v2 from 2.27.1 to 2.27.5 by @dependabot in <https://github.com/Layr-Labs/eigensdk-go/pull/431>
* chore(deps): bump github.com/prometheus/client_golang from 1.19.0 to 1.20.5 by @dependabot in <https://github.com/Layr-Labs/eigensdk-go/pull/432>
* chore(deps): bump github.com/consensys/gnark-crypto from 0.14.0 to 0.16.0 by @dependabot in <https://github.com/Layr-Labs/eigensdk-go/pull/573>

------------

Changes made previous to v0.3.0 weren't tracked by this changelog.
