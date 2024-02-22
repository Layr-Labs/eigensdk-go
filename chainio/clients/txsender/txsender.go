package txsender

import (
	"context"

	"github.com/ethereum/go-ethereum/core/types"
)

type TxID = string

// TxSender is an interface for signing and sending transactions to the network
type TxSender interface {
	SendTransaction(ctx context.Context, tx *types.Transaction) (TxID, error)
	GetTransactionReceipt(ctx context.Context, txID TxID) (*types.Receipt, error)
}
