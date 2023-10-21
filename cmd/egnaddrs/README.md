## egnaddrs
This tool is used to help debug and test eigenlayer/avs deployments and contract setups.

### How to install
#### Install from repository
```bash
   go install github.com/Layr-Labs/eigensdk-go/cmd/egnaddrs@latest
```

#### Install from local source
If you have git cloned eigensdk-go to your machine, navigate to [cmd/egnaddrs](.) directory and run
```bash
   go install
```

### Usage

Currently egnaddrs only supports deriving contract addresses starting from the registry-coordinator address. It then prints the following datastructure:

```bash
$$$ egnaddrs --registry-coordinator 0x9E545E3C0baAB3E08CdfD552C960A1050f373042
{
  "avs": {
    "bls-pubkey-compendium (shared)": "0x322813Fd9A801c5507c9de605d63CEA4f2CE6c44",
    "bls-pubkey-registry": "0xa82fF9aFd8f496c3d6ac40E2a0F282E47488CFc9",
    "index-registry": "0x1613beB3B2C4f22Ee086B2b38C1476A3cE7f78E8",
    "registry-coordinator": "0x9E545E3C0baAB3E08CdfD552C960A1050f373042",
    "service-manager": "0xc3e53F4d16Ae77Db1c982e75a937B9f60FE63690",
    "stake-registry": "0x851356ae760d987E095750cCeb3bC6014560891C"
  },
  "eigenlayer": {
    "delegation-manager": "0xCf7Ed3AccA5a467e9e704C703E8D87F634fB0Fc9",
    "slasher": "0x5FC8d32690cc91D4c39d9d3abcBD16989F875707",
    "strategy-manager": "0xDc64a140Aa3E981100a9becA4E685f962f0cF6C9"
  },
  "network": {
    "chain-id": "31337",
    "rpc-url": "http://localhost:8545"
  }
}
```