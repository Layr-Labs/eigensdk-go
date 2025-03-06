// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.27;

import {RegistryCoordinator} from "eigenlayer-middleware/src/RegistryCoordinator.sol";
import {SlashingRegistryCoordinator} from "eigenlayer-middleware/src/SlashingRegistryCoordinator.sol";

import {IPauserRegistry} from "eigenlayer-contracts/src/contracts/interfaces/IPauserRegistry.sol";
import {IAllocationManager} from "eigenlayer-contracts/src/contracts/interfaces/IAllocationManager.sol";
import {IBLSApkRegistry} from "eigenlayer-middleware/src/interfaces/IBLSApkRegistry.sol";
import {IStakeRegistry} from "eigenlayer-middleware/src/interfaces/IStakeRegistry.sol";
import {IIndexRegistry} from "eigenlayer-middleware/src/interfaces/IIndexRegistry.sol";
import {IServiceManager} from "eigenlayer-middleware/src/interfaces/IServiceManager.sol";
import {ISocketRegistry} from "eigenlayer-middleware/src/interfaces/ISocketRegistry.sol";

/// Extension of RegistryCoordinator that allows disabling operator sets and enabling M2 quorums.
/// Used for testing M2 compatibility.
contract LegacyRegistryCoordinator is RegistryCoordinator {
    constructor(
        IServiceManager _serviceManager,
        IStakeRegistry _stakeRegistry,
        IBLSApkRegistry _blsApkRegistry,
        IIndexRegistry _indexRegistry,
        ISocketRegistry _socketRegistry,
        IAllocationManager _allocationManager,
        IPauserRegistry _pauserRegistry
    )
        RegistryCoordinator(
            _serviceManager,
            _stakeRegistry,
            _blsApkRegistry,
            _indexRegistry,
            _socketRegistry,
            _allocationManager,
            _pauserRegistry
        )
    {}

    /// Disables operator sets mode
    /// @dev USE ONLY FOR TESTING
    function disableOperatorSets() external onlyOwner {
        operatorSetsEnabled = false;
    }

    /// Enables M2 quorums mode
    /// @dev USE ONLY FOR TESTING
    function enableM2QuorumRegistration() external onlyOwner {
        isM2QuorumRegistrationDisabled = false;
    }
}
