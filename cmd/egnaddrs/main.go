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

	blspubkeyreg "github.com/Layr-Labs/eigensdk-go/contracts/bindings/BLSPubkeyRegistry"
	blsregistrycoordinator "github.com/Layr-Labs/eigensdk-go/contracts/bindings/BLSRegistryCoordinatorWithIndices"
	iblssigchecker "github.com/Layr-Labs/eigensdk-go/contracts/bindings/IBlsSignatureChecker"
	slasher "github.com/Layr-Labs/eigensdk-go/contracts/bindings/Slasher"
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
	app := cli.NewApp()
	app.Name = "egnaddrs"
	app.Usage = "Used to help debug and test deployments and contract setups."
	// TODO: add a much more descriptive description
	app.Description = "Prints all reachable Eigenlayer and AVS contract addresses starting from any one contract."
	app.Action = printAddrs
	app.Flags = []cli.Flag{
		ServiceManagerAddrFlag,
		RegistryCoordinatorAddrFlag,
		RpcUrlFlag,
	}

	if err := app.Run(os.Args); err != nil {
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

	registryCoordinatorAddr, err := getRegistryCoordinatorAddr(c, client)
	if err != nil {
		return err
	}
	eigenlayerContractAddrs, err := getEigenlayerContractAddrs(client, registryCoordinatorAddr)
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

func getRegistryCoordinatorAddr(c *cli.Context, client *ethclient.Client) (common.Address, error) {
	registryCoordinatorAddrString := c.String(RegistryCoordinatorAddrFlag.Name)
	if registryCoordinatorAddrString != "" {
		registryCoordinatorAddr := common.HexToAddress(registryCoordinatorAddrString)
		return registryCoordinatorAddr, nil
	}
	serviceManagerAddrString := c.String(ServiceManagerAddrFlag.Name)
	if serviceManagerAddrString != "" {
		serviceManagerAddr := common.HexToAddress(serviceManagerAddrString)
		// we use the IBLSSignatureChecker interface because the IServiceManager interface doesn't have a getter for the registry coordinator
		// because we don't want to restrict teams to use our registry contracts.
		// However, egnaddrs is targetted at teams using our registry contracts, so we assume that they are using our registries
		// and that their service manager inherits the IBLSSignatureChecker interface (to check signatures against the BLSPubkeyRegistry).
		serviceManagerContract, err := iblssigchecker.NewContractIBLSSignatureChecker(serviceManagerAddr, client)
		if err != nil {
			return common.Address{}, err
		}
		registryCoordinatorAddr, err := serviceManagerContract.RegistryCoordinator(&bind.CallOpts{})
		if err != nil {
			return common.Address{}, err
		}
		return registryCoordinatorAddr, nil
	}
	return common.Address{}, errors.New("must provide either --registry-coordinator or --service-manager flag")
}

func getAvsContractAddrs(client *ethclient.Client, registryCoordinatorAddr common.Address) (map[string]string, error) {
	blsRegistryCoordinatorWithIndicesC, err := blsregistrycoordinator.NewContractBLSRegistryCoordinatorWithIndices(registryCoordinatorAddr, client)
	if err != nil {
		return nil, err
	}
	_ = blsRegistryCoordinatorWithIndicesC

	serviceManagerAddr, err := blsRegistryCoordinatorWithIndicesC.ServiceManager(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}

	// 3 registries
	blsPubkeyRegistryAddr, err := blsRegistryCoordinatorWithIndicesC.BlsPubkeyRegistry(&bind.CallOpts{})
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

	// pubkey compendium (shared avs contract)
	contractBlsPubkeyRegistry, err := blspubkeyreg.NewContractBLSPubkeyRegistry(blsPubkeyRegistryAddr, client)
	if err != nil {
		return nil, err
	}
	blsPubKeyCompendiumAddr, err := contractBlsPubkeyRegistry.PubkeyCompendium(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}

	addrsDict := map[string]string{
		"service-manager":                serviceManagerAddr.Hex(),
		"registry-coordinator":           registryCoordinatorAddr.Hex(),
		"bls-pubkey-registry":            blsPubkeyRegistryAddr.Hex(),
		"index-registry":                 indexRegistryAddr.Hex(),
		"stake-registry":                 stakeRegistryAddr.Hex(),
		"bls-pubkey-compendium (shared)": blsPubKeyCompendiumAddr.Hex(),
	}
	return addrsDict, nil
}

func getEigenlayerContractAddrs(client *ethclient.Client, registryCoordinatorAddr common.Address) (map[string]string, error) {
	blsRegistryCoordinatorWithIndicesC, err := blsregistrycoordinator.NewContractBLSRegistryCoordinatorWithIndices(registryCoordinatorAddr, client)
	if err != nil {
		return nil, err
	}

	slasherAddr, err := blsRegistryCoordinatorWithIndicesC.Slasher(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	slasherC, err := slasher.NewContractSlasher(slasherAddr, client)
	if err != nil {
		return nil, err
	}
	delegationManagerAddr, err := slasherC.Delegation(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	strategyManagerAddr, err := slasherC.StrategyManager(&bind.CallOpts{})
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
