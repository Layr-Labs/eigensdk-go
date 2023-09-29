package clients

import (
	"math/big"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	erc20 "github.com/Layr-Labs/eigensdk-go/contracts/bindings/IERC20"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type ERC20ContractClient interface {
	Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error)
}

type eRC20ContractChainClient struct {
	erc20Contract *erc20.ContractIERC20
}

var _ ERC20ContractClient = (*eRC20ContractChainClient)(nil)

func NewERC20ContractChainClient(
	address common.Address,
	client eth.EthClient,
) (*eRC20ContractChainClient, error) {
	contract, err := erc20.NewContractIERC20(address, client)
	if err != nil {
		panic(err)
	}
	return &eRC20ContractChainClient{erc20Contract: contract}, nil
}

// Approve implements ERC20ContractService.
func (e *eRC20ContractChainClient) Approve(
	opts *bind.TransactOpts,
	spender common.Address,
	amount *big.Int,
) (*types.Transaction, error) {
	return e.erc20Contract.Approve(opts, spender, amount)
}
