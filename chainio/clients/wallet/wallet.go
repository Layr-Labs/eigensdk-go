package wallet

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type TxID = string

// Wallet is an interface for signing and sending transactions to the txpool.
// For a higher-level interface that includes nonce management and gas bumping, use the TxManager interface.
// This interface is used to abstract the process of sending transactions to the Ethereum network
// For example, for an MPC signer, the transaction would be broadcasted via an external API endpoint
// and the status is tracked via another external endpoint instead of being broadcasted
// and retrieved via an Ethereum client.
type Wallet interface {
	SendTransaction(ctx context.Context, tx *types.Transaction) (TxID, error)
	GetTransactionReceipt(ctx context.Context, txID TxID) (*types.Receipt, error)
	// SenderAddress returns the address of the wallet
	SenderAddress(ctx context.Context) (common.Address, error)
}
