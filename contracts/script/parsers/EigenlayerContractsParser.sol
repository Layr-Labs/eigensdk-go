// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.27;

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "eigenlayer-contracts/src/contracts/permissions/PauserRegistry.sol";

import {IAVSDirectory} from "eigenlayer-contracts/src/contracts/interfaces/IAVSDirectory.sol";
import {IDelegationManager} from "eigenlayer-contracts/src/contracts/interfaces/IDelegationManager.sol";
import {IStrategyManager, IStrategy} from "eigenlayer-contracts/src/contracts/interfaces/IStrategyManager.sol";
import {StrategyBaseTVLLimits} from "eigenlayer-contracts/src/contracts/strategies/StrategyBaseTVLLimits.sol";
import {IRewardsCoordinator} from "eigenlayer-contracts/src/contracts/interfaces/IRewardsCoordinator.sol";
import {IPermissionController} from "eigenlayer-contracts/src/contracts/interfaces/IPermissionController.sol";
import {IAllocationManager} from "eigenlayer-contracts/src/contracts/interfaces/IAllocationManager.sol";

import {ConfigsReadWriter} from "./ConfigsReadWriter.sol";
import "forge-std/StdJson.sol";
import {IAllocationManager} from
    "../../lib/eigenlayer-middleware/lib/eigenlayer-contracts/src/contracts/interfaces/IAllocationManager.sol";

struct EigenlayerContracts {
    ProxyAdmin eigenlayerProxyAdmin;
    PauserRegistry eigenlayerPauserReg;
    IStrategyManager strategyManager;
    IDelegationManager delegationManager;
    IAVSDirectory avsDirectory;
    IRewardsCoordinator rewardsCoordinator;
    StrategyBaseTVLLimits baseStrategyImplementation;
    IPermissionController permissionController;
    IAllocationManager allocationManager;
}

contract EigenlayerContractsParser is ConfigsReadWriter {
    function _loadEigenlayerDeployedContracts() internal view returns (EigenlayerContracts memory) {
        // Eigenlayer contracts
        string memory eigenlayerDeployedContracts = readOutput("eigenlayer_deployment_output");
        ProxyAdmin eigenlayerProxyAdmin =
            ProxyAdmin(stdJson.readAddress(eigenlayerDeployedContracts, ".addresses.eigenLayerProxyAdmin"));
        PauserRegistry eigenlayerPauserReg =
            PauserRegistry(stdJson.readAddress(eigenlayerDeployedContracts, ".addresses.eigenLayerPauserReg"));
        IStrategyManager strategyManager =
            IStrategyManager(stdJson.readAddress(eigenlayerDeployedContracts, ".addresses.strategyManager"));
        IDelegationManager delegationManager =
            IDelegationManager(stdJson.readAddress(eigenlayerDeployedContracts, ".addresses.delegationManager"));
        IAVSDirectory avsDirectory =
            IAVSDirectory(stdJson.readAddress(eigenlayerDeployedContracts, ".addresses.avsDirectory"));
        StrategyBaseTVLLimits baseStrategyImplementation = StrategyBaseTVLLimits(
            stdJson.readAddress(eigenlayerDeployedContracts, ".addresses.baseStrategyImplementation")
        );

        IRewardsCoordinator rewardsCoordinator =
            IRewardsCoordinator(stdJson.readAddress(eigenlayerDeployedContracts, ".addresses.rewardsCoordinator"));

        IPermissionController permissionController =
            IPermissionController(stdJson.readAddress(eigenlayerDeployedContracts, ".addresses.permissionController"));

        IAllocationManager allocationManager =
            IAllocationManager(stdJson.readAddress(eigenlayerDeployedContracts, ".addresses.allocationManager"));
        return EigenlayerContracts(
            eigenlayerProxyAdmin,
            eigenlayerPauserReg,
            strategyManager,
            delegationManager,
            avsDirectory,
            rewardsCoordinator,
            baseStrategyImplementation,
            permissionController,
            allocationManager
        );
    }
}
