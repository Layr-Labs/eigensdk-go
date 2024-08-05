package signer

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// Signer is an interface that defines the methods that a signer must implement.
// There are two kinds of signer.
//  1. A signer that can implement a Signer function and return the signed transaction.
//     They need to implement GetSigner() function and leave SendToExternal() unimplemented.
//
// 2. A signer (remote signer) that we will send to an external rpc which takes care of signing and broadcasting the
// signed transaction. They need to implement SendToExternal() function and leave GetSigner() unimplemented.
// Deprecated: Use SignerV2 instead
type Signer interface {
	GetTxOpts() *bind.TransactOpts
	SendToExternal(ctx context.Context, tx *types.Transaction) (common.Hash, error)
}
