package eth

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

// HttpBackend is the HTTP ETH Client interface
// It is exactly the same as the WsBackend and there is no difference between them to the compiler,
// but we keep them separate as a signal to the programmer that an eth.Client with an underlying http/ws connection is expected
type HttpBackend interface {
	bind.ContractBackend

	BlockNumber(ctx context.Context) (uint64, error)
	BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error)
}

// WsBackend is the Websocket ETH Client interface
// It is exactly the same as the HttpBackend and there is no difference between them to the compiler,
// but we keep them separate as a signal to the programmer that an eth.Client with an underlying http/ws connection is expected
type WsBackend interface {
	bind.ContractBackend

	BlockNumber(ctx context.Context) (uint64, error)
	BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error)
}
