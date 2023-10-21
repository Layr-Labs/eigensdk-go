package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/urfave/cli/v2"

	blspubkeyreg "github.com/Layr-Labs/eigensdk-go/contracts/bindings/BLSPubkeyRegistry"
	blsregistrycoordinator "github.com/Layr-Labs/eigensdk-go/contracts/bindings/BLSRegistryCoordinatorWithIndices"
	slasher "github.com/Layr-Labs/eigensdk-go/contracts/bindings/Slasher"
)

var (
	RegistryCoordinatorAddrFlag = &cli.StringFlag{
		Name:     "registry-coordinator",
		Usage:    "BLSRegistryCoordinatorWithIndices contract address",
		Required: true,
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

	registryCoordinatorAddr := common.HexToAddress(c.String(RegistryCoordinatorAddrFlag.Name))
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
