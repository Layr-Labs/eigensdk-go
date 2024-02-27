package txsender

import (
	"context"

	"github.com/ethereum/go-ethereum/core/types"
)

type TxID = string

// TxSender is an interface for signing and sending transactions to the network
// This interface is used to abstract the process of sending transactions to the Ethereum network
// For example, for an MPC signer, the transaction would be broadcasted via an external API endpoint
// and the status is tracked via another external endpoint instead of being broadcasted
// and retrieved via an Ethereum client.
type TxSender interface {
	SendTransaction(ctx context.Context, tx *types.Transaction) (TxID, error)
	GetTransactionReceipt(ctx context.Context, txID TxID) (*types.Receipt, error)
}
