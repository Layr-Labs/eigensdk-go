// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.27;

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import "eigenlayer-contracts/src/contracts/permissions/PauserRegistry.sol";
import {IDelegationManager} from "eigenlayer-contracts/src/contracts/interfaces/IDelegationManager.sol";
import {IStrategyManager, IStrategy} from "eigenlayer-contracts/src/contracts/interfaces/IStrategyManager.sol";
import "eigenlayer-contracts/src/test/mocks/EmptyContract.sol";
import {IServiceManager} from "eigenlayer-middleware/src/interfaces/IServiceManager.sol";
import {
    ISlashingRegistryCoordinator,
    ISlashingRegistryCoordinatorTypes
} from "eigenlayer-middleware/src/interfaces/ISlashingRegistryCoordinator.sol";
import {SlashingRegistryCoordinator} from "eigenlayer-middleware/src/SlashingRegistryCoordinator.sol";
import {BLSApkRegistry, IBLSApkRegistry} from "eigenlayer-middleware/src/BLSApkRegistry.sol";
import {IndexRegistry, IIndexRegistry} from "eigenlayer-middleware/src/IndexRegistry.sol";
import {StakeRegistry, IStakeRegistry, IStakeRegistryTypes} from "eigenlayer-middleware/src/StakeRegistry.sol";
import {SocketRegistry, ISocketRegistry} from "eigenlayer-middleware/src/SocketRegistry.sol";
import {OperatorStateRetriever} from "eigenlayer-middleware/src/OperatorStateRetriever.sol";
import {MockAvsContracts} from "./parsers/MockAvsContractsParser.sol";
import {EigenlayerContracts, EigenlayerContractsParser} from "./parsers/EigenlayerContractsParser.sol";
import {ConfigsReadWriter} from "./parsers/ConfigsReadWriter.sol";
import {MockAvsServiceManager} from "../src/MockAvsServiceManager.sol";
import {ContractsRegistry} from "../src/ContractsRegistry.sol";
import {LegacyRegistryCoordinator} from "../src/LegacyRegistryCoordinator.sol";

import "forge-std/Script.sol";
import "forge-std/console.sol";
import "forge-std/StdJson.sol";

contract DeployMockAvsRegistries is Script, ConfigsReadWriter, EigenlayerContractsParser {
    struct Registries {
        IBLSApkRegistry blsApkRegistry;
        IBLSApkRegistry blsApkRegistryImplementation;
        IIndexRegistry indexRegistry;
        IIndexRegistry indexRegistryImplementation;
        IStakeRegistry stakeRegistry;
        IStakeRegistry stakeRegistryImplementation;
        ISocketRegistry socketRegistry;
        ISocketRegistry socketRegistryImplementation;
    }

    struct DeployedContracts {
        ProxyAdmin proxyAdmin;
        PauserRegistry pauserReg;
        SlashingRegistryCoordinator coordinator;
        ISlashingRegistryCoordinator coordinatorImplementation;
        OperatorStateRetriever stateRetriever;
        EmptyContract emptyContract;
    }

    struct MockAvsOpsAddresses {
        address communityMultisig;
        address pauser;
        address churner;
        address ejector;
    }

    Registries private registries;
    DeployedContracts private deployed;

    function _loadAvsOpsAddresses(string memory fileName) internal view returns (MockAvsOpsAddresses memory config) {
        string memory addresses = readInput(fileName);
        config.communityMultisig = stdJson.readAddress(addresses, ".communityMultisig");
        config.pauser = stdJson.readAddress(addresses, ".pauser");
        config.churner = stdJson.readAddress(addresses, ".churner");
        config.ejector = stdJson.readAddress(addresses, ".ejector");
    }

    function _deploymockAvsRegistryContracts(
        EigenlayerContracts memory eigenlayerContracts,
        MockAvsOpsAddresses memory addressConfig,
        MockAvsServiceManager manager,
        MockAvsServiceManager managerImpl
    ) internal returns (MockAvsContracts memory) {
        _deployPauserRegistry(addressConfig);
        _deployProxies();
        deployed.stateRetriever = new OperatorStateRetriever();
        _deployAndUpgradeImplementations(eigenlayerContracts, manager);
        _initializeRegistryCoordinator(addressConfig, manager);

        _setupPermissions(addressConfig.communityMultisig, eigenlayerContracts);

        require(Ownable(address(deployed.coordinator)).owner() != address(0), "Owner uninitialized");
        _writeDeploymentOutput(manager, managerImpl);

        return MockAvsContracts(manager, deployed.coordinator, deployed.stateRetriever);
    }

    function _deployPauserRegistry(MockAvsOpsAddresses memory config) internal {
        address[] memory pausers = new address[](2);
        pausers[0] = config.pauser;
        pausers[1] = config.communityMultisig;
        deployed.pauserReg = new PauserRegistry(pausers, config.communityMultisig);
    }

    function _deployProxies() internal {
        if (address(deployed.emptyContract) == address(0)) {
            deployed.emptyContract = new EmptyContract();
        }
        if (address(deployed.proxyAdmin) == address(0)) {
            deployed.proxyAdmin = new ProxyAdmin();
        }
        deployed.coordinator = SlashingRegistryCoordinator(_deployProxy());
        registries.blsApkRegistry = IBLSApkRegistry(_deployProxy());
        registries.indexRegistry = IIndexRegistry(_deployProxy());
        registries.stakeRegistry = IStakeRegistry(_deployProxy());
        registries.socketRegistry = ISocketRegistry(_deployProxy());
    }

    function _deployProxy() internal returns (address) {
        return
            address(new TransparentUpgradeableProxy(address(deployed.emptyContract), address(deployed.proxyAdmin), ""));
    }

    function _deployAndUpgradeImplementations(EigenlayerContracts memory eigen, MockAvsServiceManager manager)
        internal
    {
        registries.blsApkRegistryImplementation = new BLSApkRegistry(deployed.coordinator);
        _upgradeProxy(address(registries.blsApkRegistry), address(registries.blsApkRegistryImplementation));

        registries.indexRegistryImplementation = new IndexRegistry(deployed.coordinator);
        _upgradeProxy(address(registries.indexRegistry), address(registries.indexRegistryImplementation));

        registries.stakeRegistryImplementation = new StakeRegistry(
            deployed.coordinator, eigen.delegationManager, eigen.avsDirectory, eigen.allocationManager
        );
        _upgradeProxy(address(registries.stakeRegistry), address(registries.stakeRegistryImplementation));

        registries.socketRegistryImplementation =
            new SocketRegistry(ISlashingRegistryCoordinator(address(deployed.coordinator)));
        _upgradeProxy(address(registries.socketRegistry), address(registries.socketRegistryImplementation));

        deployed.coordinatorImplementation = new SlashingRegistryCoordinator(
            registries.stakeRegistry,
            registries.blsApkRegistry,
            registries.indexRegistry,
            registries.socketRegistry,
            eigen.allocationManager,
            deployed.pauserReg
        );
    }

    function _upgradeProxy(address proxy, address implementation) internal {
        deployed.proxyAdmin.upgrade(TransparentUpgradeableProxy(payable(proxy)), implementation);
    }

    function _initializeRegistryCoordinator(MockAvsOpsAddresses memory config, MockAvsServiceManager manager)
        internal
    {
        deployed.proxyAdmin.upgradeAndCall(
            TransparentUpgradeableProxy(payable(address(deployed.coordinator))),
            address(deployed.coordinatorImplementation),
            abi.encodeCall(
                deployed.coordinator.initialize,
                (config.communityMultisig, config.churner, config.ejector, 0, address(manager))
            )
        );
    }

    function _setupPermissions(address avs, EigenlayerContracts memory elContracts) internal {
        address coordinator = address(deployed.coordinator);
        address allocationManager = address(elContracts.allocationManager);
        elContracts.permissionController.setAppointee(
            avs, coordinator, allocationManager, elContracts.allocationManager.createOperatorSets.selector
        );
        elContracts.permissionController.setAppointee(
            avs, coordinator, allocationManager, elContracts.allocationManager.deregisterFromOperatorSets.selector
        );
        address stakeRegistry = address(deployed.coordinator.stakeRegistry());
        elContracts.permissionController.setAppointee(
            avs, stakeRegistry, allocationManager, elContracts.allocationManager.addStrategiesToOperatorSet.selector
        );
        elContracts.permissionController.setAppointee(
            avs,
            stakeRegistry,
            allocationManager,
            elContracts.allocationManager.removeStrategiesFromOperatorSet.selector
        );
        // NOTE: we don't set permissions for the slasher, because there isn't one
        // elContracts.permissionController.setAppointee(
        //     avs, slasher, allocationManager, elContracts.allocationManager.slashOperator.selector
        // );
    }

    function _writeDeploymentOutput(MockAvsServiceManager manager, MockAvsServiceManager managerImpl) internal {
        string memory parent = "parent object";
        string memory addresses = "addresses";

        vm.serializeAddress(addresses, "proxyAdmin", address(deployed.proxyAdmin));
        vm.serializeAddress(addresses, "mockAvsServiceManager", address(manager));
        vm.serializeAddress(addresses, "mockAvsServiceManagerImplementation", address(managerImpl));
        vm.serializeAddress(addresses, "registryCoordinator", address(deployed.coordinator));
        vm.serializeAddress(addresses, "registryCoordinatorImplementation", address(deployed.coordinatorImplementation));
        string memory output =
            vm.serializeAddress(addresses, "operatorStateRetriever", address(deployed.stateRetriever));

        writeOutput(vm.serializeString(parent, addresses, output), "mockavs_deployment_output");
    }

    function _writeContractsToRegistry(
        ContractsRegistry registry,
        EigenlayerContracts memory eigen,
        MockAvsContracts memory avs
    ) internal {
        registry.registerContract("mockAvsServiceManager", address(avs.mockAvsServiceManager));
        registry.registerContract("mockAvsRegistryCoordinator", address(avs.registryCoordinator));
        registry.registerContract("mockAvsOperatorStateRetriever", address(avs.operatorStateRetriever));
        registry.registerContract("delegationManager", address(eigen.delegationManager));
        registry.registerContract("strategyManager", address(eigen.strategyManager));
        registry.registerContract("rewardsCoordinator", address(eigen.rewardsCoordinator));
        registry.registerContract("permissionController", address(eigen.permissionController));
    }
}
