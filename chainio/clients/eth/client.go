package eth

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

// HttpBackend is the HTTP ETH Client interface
type HttpBackend interface {
	bind.ContractBackend

	BlockNumber(ctx context.Context) (uint64, error)
	BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error)
}

// WsBackend is the Websocket ETH Client interface
type WsBackend interface {
	bind.ContractBackend

	BlockNumber(ctx context.Context) (uint64, error)
	BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error)
}
