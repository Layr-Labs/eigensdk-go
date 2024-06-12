// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import "eigenlayer-contracts/src/contracts/permissions/PauserRegistry.sol";

import {IDelegationManager} from "eigenlayer-contracts/src/contracts/interfaces/IDelegationManager.sol";
import {IStrategyManager, IStrategy} from "eigenlayer-contracts/src/contracts/interfaces/IStrategyManager.sol";
import {ISlasher} from "eigenlayer-contracts/src/contracts/interfaces/ISlasher.sol";
import "eigenlayer-contracts/src/test/mocks/EmptyContract.sol";

import "eigenlayer-middleware/src/RegistryCoordinator.sol" as blsregcoord;
import {IServiceManager} from "eigenlayer-middleware/src/interfaces/IServiceManager.sol";
import {IBLSApkRegistry, IIndexRegistry, IStakeRegistry} from "eigenlayer-middleware/src/RegistryCoordinator.sol";
import {BLSApkRegistry} from "eigenlayer-middleware/src/BLSApkRegistry.sol";
import {IndexRegistry} from "eigenlayer-middleware/src/IndexRegistry.sol";
import {StakeRegistry} from "eigenlayer-middleware/src/StakeRegistry.sol";
import {OperatorStateRetriever} from "eigenlayer-middleware/src/OperatorStateRetriever.sol";

import {MockAvsContracts} from "./parsers/MockAvsContractsParser.sol";
import {EigenlayerContracts, EigenlayerContractsParser} from "./parsers/EigenlayerContractsParser.sol";
import {ConfigsReadWriter} from "./parsers/ConfigsReadWriter.sol";
import {MockAvsServiceManager} from "../src/MockAvsServiceManager.sol";
import {ContractsRegistry} from "../src/ContractsRegistry.sol";

import "forge-std/Script.sol";
import "forge-std/StdJson.sol";

// forge script script/DeployMockAvs.s.sol --rpc-url $RPC_URL --private-key $PRIVATE_KEY --etherscan-api-key $ETHERSCAN_API_KEY --broadcast --verify
contract Deploy is Script, ConfigsReadWriter, EigenlayerContractsParser {
    // MockAvs contracts
    ProxyAdmin public mockAvsSquaringProxyAdmin;
    PauserRegistry public mockAvsPauserReg;
    blsregcoord.RegistryCoordinator public registryCoordinator;
    blsregcoord.IRegistryCoordinator public registryCoordinatorImplementation;
    IBLSApkRegistry public blsApkRegistry;
    IBLSApkRegistry public blsApkRegistryImplementation;
    IIndexRegistry public indexRegistry;
    IIndexRegistry public indexRegistryImplementation;
    IStakeRegistry public stakeRegistry;
    IStakeRegistry public stakeRegistryImplementation;
    OperatorStateRetriever public operatorStateRetriever;
    MockAvsServiceManager public mockAvsServiceManager;
    MockAvsServiceManager public mockAvsServiceManagerImplementation;

    struct MockAvsOpsAddresses {
        address communityMultisig;
        address pauser;
        address churner;
        address ejector;
    }

    function run() public virtual {
        ContractsRegistry contractsRegistry = ContractsRegistry(
            0x5FbDB2315678afecb367f032d93F642f64180aa3
        );
        EigenlayerContracts
            memory eigenlayerContracts = _loadEigenlayerDeployedContracts();
        string memory opsAddresses = readInput("ops_addresses");
        MockAvsOpsAddresses memory addressConfig;
        addressConfig.communityMultisig = stdJson.readAddress(
            opsAddresses,
            ".mockAvs.communityMultisig"
        );
        addressConfig.pauser = stdJson.readAddress(
            opsAddresses,
            ".mockAvs.pauser"
        );
        addressConfig.churner = stdJson.readAddress(
            opsAddresses,
            ".mockAvs.churner"
        );
        addressConfig.ejector = stdJson.readAddress(
            opsAddresses,
            ".mockAvs.ejector"
        );

        vm.startBroadcast();
        MockAvsContracts memory mockAvsContracts = _deploymockAvsContracts(
            eigenlayerContracts,
            addressConfig
        );

        if (block.chainid == 31337 || block.chainid == 1337) {
            _writeContractsToRegistry(
                contractsRegistry,
                eigenlayerContracts,
                mockAvsContracts
            );
        }
        vm.stopBroadcast();
    }

    function _deploymockAvsContracts(
        EigenlayerContracts memory eigenlayerContracts,
        MockAvsOpsAddresses memory addressConfig
    ) internal returns (MockAvsContracts memory) {
        // deploy proxy admin for ability to upgrade proxy contracts
        ProxyAdmin mockAvsProxyAdmin = new ProxyAdmin();

        // deploy pauser registry
        {
            address[] memory pausers = new address[](2);
            pausers[0] = addressConfig.pauser;
            pausers[1] = addressConfig.communityMultisig;
            mockAvsPauserReg = new PauserRegistry(
                pausers,
                addressConfig.communityMultisig
            );
        }

        EmptyContract emptyContract = new EmptyContract();
        /**
         * First, deploy upgradeable proxy contracts that **will point** to the implementations. Since the implementation contracts are
         * not yet deployed, we give these proxies an empty contract as the initial implementation, to act as if they have no code.
         */
        mockAvsServiceManager = MockAvsServiceManager(
            address(
                new TransparentUpgradeableProxy(
                    address(emptyContract),
                    address(mockAvsProxyAdmin),
                    ""
                )
            )
        );
        registryCoordinator = blsregcoord.RegistryCoordinator(
            address(
                new TransparentUpgradeableProxy(
                    address(emptyContract),
                    address(mockAvsProxyAdmin),
                    ""
                )
            )
        );
        blsApkRegistry = IBLSApkRegistry(
            address(
                new TransparentUpgradeableProxy(
                    address(emptyContract),
                    address(mockAvsProxyAdmin),
                    ""
                )
            )
        );
        indexRegistry = IIndexRegistry(
            address(
                new TransparentUpgradeableProxy(
                    address(emptyContract),
                    address(mockAvsProxyAdmin),
                    ""
                )
            )
        );
        stakeRegistry = IStakeRegistry(
            address(
                new TransparentUpgradeableProxy(
                    address(emptyContract),
                    address(mockAvsProxyAdmin),
                    ""
                )
            )
        );

        operatorStateRetriever = new OperatorStateRetriever();

        // Second, deploy the *implementation* contracts, using the *proxy contracts* as inputs
        blsApkRegistryImplementation = new BLSApkRegistry(registryCoordinator);

        mockAvsProxyAdmin.upgrade(
            TransparentUpgradeableProxy(payable(address(blsApkRegistry))),
            address(blsApkRegistryImplementation)
        );

        indexRegistryImplementation = new IndexRegistry(registryCoordinator);

        mockAvsProxyAdmin.upgrade(
            TransparentUpgradeableProxy(payable(address(indexRegistry))),
            address(indexRegistryImplementation)
        );

        {
            stakeRegistryImplementation = new StakeRegistry(
                registryCoordinator,
                eigenlayerContracts.delegationManager
            );

            mockAvsProxyAdmin.upgrade(
                TransparentUpgradeableProxy(payable(address(stakeRegistry))),
                address(stakeRegistryImplementation)
            );
        }

        registryCoordinatorImplementation = new blsregcoord.RegistryCoordinator(
            blsregcoord.IServiceManager(address(mockAvsServiceManager)),
            blsregcoord.IStakeRegistry(address(stakeRegistry)),
            blsregcoord.IBLSApkRegistry(address(blsApkRegistry)),
            blsregcoord.IIndexRegistry(address(indexRegistry))
        );

        {
            uint numQuorums = 0;
            // for each quorum to setup, we need to define
            // quorumsOperatorSetParams, quorumsMinimumStake, and quorumsStrategyParams
            blsregcoord.RegistryCoordinator.OperatorSetParam[]
                memory quorumsOperatorSetParams = new blsregcoord.RegistryCoordinator.OperatorSetParam[](
                    numQuorums
                );
            for (uint i = 0; i < numQuorums; i++) {
                // hard code these for now
                quorumsOperatorSetParams[i] = blsregcoord
                    .IRegistryCoordinator
                    .OperatorSetParam({
                        maxOperatorCount: 10000,
                        kickBIPsOfOperatorStake: 15000,
                        kickBIPsOfTotalStake: 100
                    });
            }
            // set to 0 for every quorum
            uint96[] memory quorumsMinimumStake = new uint96[](numQuorums);
            IStakeRegistry.StrategyParams[][]
                memory quorumsStrategyParams = new IStakeRegistry.StrategyParams[][](
                    numQuorums
                );
            // We don't setup up any quorums so this is commented out for now
            // (since deployedStrategyArray doesn't exist)
            // for (uint i = 0; i < numQuorums; i++) {
            //     quorumsStrategyParams[i] = new IStakeRegistry.StrategyParams[](numStrategies);
            //     for (uint j = 0; j < numStrategies; j++) {
            //         quorumsStrategyParams[i][j] = IStakeRegistry.StrategyParams({
            //             strategy: deployedStrategyArray[j],
            //             // setting this to 1 ether since the divisor is also 1 ether
            //             // therefore this allows an operator to register with even just 1 token
            //             // see https://github.com/Layr-Labs/eigenlayer-middleware/blob/m2-mainnet/src/StakeRegistry.sol#L484
            //             //    weight += uint96(sharesAmount * strategyAndMultiplier.multiplier / WEIGHTING_DIVISOR);
            //             multiplier: 1 ether
            //         });
            //     }
            // }
            mockAvsProxyAdmin.upgradeAndCall(
                TransparentUpgradeableProxy(
                    payable(address(registryCoordinator))
                ),
                address(registryCoordinatorImplementation),
                abi.encodeWithSelector(
                    blsregcoord.RegistryCoordinator.initialize.selector,
                    addressConfig.communityMultisig,
                    addressConfig.churner,
                    addressConfig.ejector,
                    addressConfig.pauser,
                    0, // 0 initialPausedStatus means everything unpaused
                    quorumsOperatorSetParams,
                    quorumsMinimumStake,
                    quorumsStrategyParams
                )
            );
        }

        mockAvsServiceManagerImplementation = new MockAvsServiceManager(
            registryCoordinator,
            eigenlayerContracts.avsDirectory,
            eigenlayerContracts.rewardsCoordinator
        );

        mockAvsProxyAdmin.upgradeAndCall(
            TransparentUpgradeableProxy(
                payable(address(mockAvsServiceManager))
            ),
            address(mockAvsServiceManagerImplementation),
            abi.encodeWithSelector(
                mockAvsServiceManager.initialize.selector,
                addressConfig.communityMultisig
            )
        );

        require(
            Ownable(address(registryCoordinator)).owner() != address(0),
            "Owner uninitialized"
        );
        require(
            Ownable(address(mockAvsServiceManager)).owner() != address(0),
            "Owner uninitialized"
        );

        // WRITE JSON DATA
        {
            string memory parent_object = "parent object";
            string memory deployed_addresses = "addresses";
            vm.serializeAddress(
                deployed_addresses,
                "proxyAdmin",
                address(mockAvsProxyAdmin)
            );
            vm.serializeAddress(
                deployed_addresses,
                "mockAvsServiceManager",
                address(mockAvsServiceManager)
            );
            vm.serializeAddress(
                deployed_addresses,
                "mockAvsServiceManagerImplementation",
                address(mockAvsServiceManagerImplementation)
            );
            vm.serializeAddress(
                deployed_addresses,
                "registryCoordinator",
                address(registryCoordinator)
            );
            vm.serializeAddress(
                deployed_addresses,
                "registryCoordinatorImplementation",
                address(registryCoordinatorImplementation)
            );
            string memory deployed_addresses_output = vm.serializeAddress(
                deployed_addresses,
                "operatorStateRetriever",
                address(operatorStateRetriever)
            );

            // serialize all the data
            string memory finalJson = vm.serializeString(
                parent_object,
                deployed_addresses,
                deployed_addresses_output
            );

            writeOutput(finalJson, "mockavs_deployment_output");
        }
        return
            MockAvsContracts(
                mockAvsServiceManager,
                registryCoordinator,
                operatorStateRetriever
            );
    }

    function _writeContractsToRegistry(
        ContractsRegistry contractsRegistry,
        EigenlayerContracts memory eigenlayerContracts,
        MockAvsContracts memory mockAvsContracts
    ) internal {
        contractsRegistry.registerContract(
            "mockAvsServiceManager",
            address(mockAvsContracts.mockAvsServiceManager)
        );
        contractsRegistry.registerContract(
            "mockAvsRegistryCoordinator",
            address(mockAvsContracts.registryCoordinator)
        );
        contractsRegistry.registerContract(
            "mockAvsOperatorStateRetriever",
            address(mockAvsContracts.operatorStateRetriever)
        );
        contractsRegistry.registerContract(
            "delegationManager",
            address(eigenlayerContracts.delegationManager)
        );
        contractsRegistry.registerContract(
            "strategyManager",
            address(eigenlayerContracts.strategyManager)
        );
    }
}
