package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/urfave/cli/v2"

	dm "github.com/Layr-Labs/eigensdk-go/contracts/bindings/DelegationManager"
	iblssigchecker "github.com/Layr-Labs/eigensdk-go/contracts/bindings/IBLSSignatureChecker"
	regcoord "github.com/Layr-Labs/eigensdk-go/contracts/bindings/RegistryCoordinator"
)

var (
	ServiceManagerAddrFlag = &cli.StringFlag{
		Name:     "service-manager",
		Usage:    "ServiceManager contract address",
		Required: false,
	}
	RegistryCoordinatorAddrFlag = &cli.StringFlag{
		Name:     "registry-coordinator",
		Usage:    "BLSRegistryCoordinatorWithIndices contract address",
		Required: false,
	}
	RpcUrlFlag = &cli.StringFlag{
		Name:        "rpc-url",
		Usage:       "rpc url",
		Value:       "http://localhost:8545",
		DefaultText: "http://localhost:8545",
	}
)

func main() {
	run(os.Args)
}

// We structure run in this way to make it testable.
// see https://github.com/urfave/cli/issues/731
func run(args []string) {
	app := cli.NewApp()
	app.Name = "egnaddrs"
	app.Usage = "Used to help debug and test deployments and contract setups."
	app.Description = "This utility facilitates the debugging and testing of Eigenlayer and AVS contract deployments by retrieving and displaying a comprehensive list of contract addresses. Starting from an initial contract address provided, it recursively identifies and prints addresses for all relevant Eigenlayer and AVS contracts within the network. This includes service managers, registry coordinators, and various registries, thus providing a view of the deployment's structure within the network."

	app.Action = printAddrs
	app.Flags = []cli.Flag{
		ServiceManagerAddrFlag,
		RegistryCoordinatorAddrFlag,
		RpcUrlFlag,
	}

	if err := app.Run(args); err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

}

func printAddrs(c *cli.Context) error {
	client, err := ethclient.Dial(c.String(RpcUrlFlag.Name))
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	chainId, err := client.ChainID(ctx)
	if err != nil {
		return err
	}

	registryCoordinatorAddr, serviceManagerAddr, err := getRegCoordAndServiceMngrAddr(c, client)
	if err != nil {
		return err
	}
	eigenlayerContractAddrs, err := getEigenlayerContractAddrs(client, serviceManagerAddr)
	if err != nil {
		return err
	}
	avsContractAddrs, err := getAvsContractAddrs(client, registryCoordinatorAddr)
	if err != nil {
		return err
	}

	addrsDict := map[string]any{
		"network": map[string]string{
			"rpc-url":  c.String(RpcUrlFlag.Name),
			"chain-id": chainId.String(),
		},
		"eigenlayer": eigenlayerContractAddrs,
		"avs":        avsContractAddrs,
	}
	addrsMarshaled, err := json.MarshalIndent(addrsDict, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(addrsMarshaled))
	return nil
}

func getRegCoordAndServiceMngrAddr(
	c *cli.Context, client *ethclient.Client,
) (regcoordAddr common.Address, servicemanagerAddr common.Address, err error) {
	// if registry coordinator addr was passed as argument
	registryCoordinatorAddrString := c.String(RegistryCoordinatorAddrFlag.Name)
	if registryCoordinatorAddrString != "" {
		registryCoordinatorAddr := common.HexToAddress(registryCoordinatorAddrString)
		registryCoordinatorC, err := regcoord.NewContractRegistryCoordinator(
			registryCoordinatorAddr,
			client,
		)
		if err != nil {
			return common.Address{}, common.Address{}, err
		}
		serviceManagerAddr, err := registryCoordinatorC.ServiceManager(&bind.CallOpts{})
		if err != nil {
			return common.Address{}, common.Address{}, err
		}
		return registryCoordinatorAddr, serviceManagerAddr, nil
	}

	// else service manager addr was passed as argument
	serviceManagerAddrString := c.String(ServiceManagerAddrFlag.Name)
	if serviceManagerAddrString != "" {
		serviceManagerAddr := common.HexToAddress(serviceManagerAddrString)
		// we use the IBLSSignatureChecker interface because the IServiceManager interface doesn't have a getter for the
		// registry coordinator
		// because we don't want to restrict teams to use our registry contracts.
		// However, egnaddrs is targetted at teams using our registry contracts, so we assume that they are using our
		// registries and that their service manager inherits the IBLSSignatureChecker interface (to check signatures
		// against the BLSPubkeyRegistry).
		serviceManagerContract, err := iblssigchecker.NewContractIBLSSignatureChecker(serviceManagerAddr, client)
		if err != nil {
			return common.Address{}, common.Address{}, err
		}
		registryCoordinatorAddr, err := serviceManagerContract.RegistryCoordinator(&bind.CallOpts{})
		if err != nil {
			return common.Address{}, common.Address{}, err
		}
		return registryCoordinatorAddr, serviceManagerAddr, nil
	}
	// else neither registry coordinator addr nor service manager addr was passed as argument
	return common.Address{}, common.Address{}, errors.New(
		"must provide either --registry-coordinator or --service-manager flag",
	)
}

func getAvsContractAddrs(client *ethclient.Client, registryCoordinatorAddr common.Address) (map[string]string, error) {
	blsRegistryCoordinatorWithIndicesC, err := regcoord.NewContractRegistryCoordinator(
		registryCoordinatorAddr,
		client,
	)
	if err != nil {
		return nil, err
	}
	_ = blsRegistryCoordinatorWithIndicesC

	serviceManagerAddr, err := blsRegistryCoordinatorWithIndicesC.ServiceManager(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}

	// 3 registries
	blsPubkeyApkAddr, err := blsRegistryCoordinatorWithIndicesC.BlsApkRegistry(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	indexRegistryAddr, err := blsRegistryCoordinatorWithIndicesC.IndexRegistry(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	stakeRegistryAddr, err := blsRegistryCoordinatorWithIndicesC.StakeRegistry(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}

	addrsDict := map[string]string{
		"service-manager":      serviceManagerAddr.Hex(),
		"registry-coordinator": registryCoordinatorAddr.Hex(),
		"bls-apk-registry":     blsPubkeyApkAddr.Hex(),
		"index-registry":       indexRegistryAddr.Hex(),
		"stake-registry":       stakeRegistryAddr.Hex(),
	}
	return addrsDict, nil
}

func getEigenlayerContractAddrs(
	client *ethclient.Client,
	serviceManagerAddr common.Address,
) (map[string]string, error) {
	serviceManagerC, err := iblssigchecker.NewContractIBLSSignatureChecker(
		serviceManagerAddr,
		client,
	)
	if err != nil {
		return nil, err
	}

	delegationManagerAddr, err := serviceManagerC.Delegation(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	delegationManagerC, err := dm.NewContractDelegationManager(delegationManagerAddr, client)
	if err != nil {
		return nil, err
	}
	slasherAddr, err := delegationManagerC.Slasher(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	strategyManagerAddr, err := delegationManagerC.StrategyManager(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}

	addrsDict := map[string]string{
		"slasher":            slasherAddr.Hex(),
		"delegation-manager": delegationManagerAddr.Hex(),
		"strategy-manager":   strategyManagerAddr.Hex(),
	}
	return addrsDict, nil
}
