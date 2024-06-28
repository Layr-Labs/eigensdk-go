package txmgr

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
)

// We are taking inspiration from the optimism TxManager interface
// https://github.com/ethereum-optimism/optimism/blob/develop/op-service/txmgr/txmgr.go

type TxManager interface {
	// Send is used to send a transaction and
	// It takes an unsigned transaction and then signs it before sending
	// It also takes care of nonce management and gas estimation
	Send(ctx context.Context, tx *types.Transaction) (*types.Receipt, error)

	// GetNoSendTxOpts generates a TransactOpts with
	// - NoSend=true: b/c we want to manage the sending ourselves
	// - Signer=NoopSigner: b/c we want the wallet to manage signing
	// - From=sender: unfortunately needed as first parameter to
	// This is needed when using abigen to construct transactions.
	// this to generate the transaction without actually sending it
	GetNoSendTxOpts() (*bind.TransactOpts, error)
}

// TxnManager receives transactions from the caller, sends them to the chain, and monitors their status.
// It also handles the case where a transaction is not mined within a certain time. In this case, it will
// resend the transaction with a higher gas price. It is assumed that all transactions originate from the
// same account.
type AsyncTxManager interface {
	Start(ctx context.Context)
	ProcessTransaction(ctx context.Context, req *TxnRequest) error
	ReceiptChan() chan *ReceiptOrErr
}
