package txmgr

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
)

type TxManager interface {
	// Send is used to sign and send a transaction to an evm chain
	// It takes an unsigned transaction and then signs it before sending
	// It might also take care of nonce management and gas estimation, depending on the implementation
	Send(ctx context.Context, tx *types.Transaction) (*types.Receipt, error)

	// GetNoSendTxOpts generates a TransactOpts with
	// - NoSend=true: b/c we want to manage the sending ourselves
	// - Signer=NoopSigner: b/c we want the wallet to manage signing
	// - From=sender: unfortunately needed as first parameter to
	// This is needed when using abigen to construct transactions.
	// this to generate the transaction without actually sending it
	GetNoSendTxOpts() (*bind.TransactOpts, error)
}
