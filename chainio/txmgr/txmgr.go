package txmgr

import (
	"context"
	"errors"
	"time"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/wallet"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/signerv2"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
)

// We are taking inspiration from the optimism TxManager interface
// https://github.com/ethereum-optimism/optimism/blob/develop/op-service/txmgr/txmgr.go

type TxManager interface {
	// Send is used to send a transaction
	// It takes an unsigned transaction and then signs it before sending
	// It also takes care of nonce management and gas estimation
	Send(ctx context.Context, tx *types.Transaction) (*types.Receipt, error)

	// GetNoSendTxOpts This generates a noSend TransactOpts so that we can use
	// this to generate the transaction without actually sending it
	GetNoSendTxOpts() (*bind.TransactOpts, error)
}

type SimpleTxManager struct {
	wallet   wallet.Wallet
	client   eth.Client
	signerFn signerv2.SignerFn
	log      logging.Logger
	sender   common.Address
}

var _ TxManager = (*SimpleTxManager)(nil)

// NewSimpleTxManager creates a new simpleTxManager which can be used
// to send a transaction to smart contracts on the Ethereum node
func NewSimpleTxManager(
	wallet wallet.Wallet,
	client eth.Client,
	log logging.Logger,
	signerFn signerv2.SignerFn,
	sender common.Address,
) *SimpleTxManager {
	return &SimpleTxManager{
		wallet:   wallet,
		client:   client,
		log:      log,
		signerFn: signerFn,
		sender:   sender,
	}
}

// Send is used to send a transaction to the Ethereum node. It takes an unsigned/signed transaction
// and then sends it to the Ethereum node.
// It also takes care of gas estimation and adds a buffer to the gas limit
// If you pass in a signed transaction it will ignore the signature
// and resign the transaction after adding the nonce and gas limit.
// To check out the whole flow on how this works, check out the README.md in this folder
func (m *SimpleTxManager) Send(ctx context.Context, tx *types.Transaction) (*types.Receipt, error) {
	txID, err := m.wallet.SendTransaction(ctx, tx)
	if err != nil {
		return nil, err
	}

	receipt, err := m.waitForReceipt(ctx, txID)
	if err != nil {
		log.Info("Transaction receipt not found", "err", err)
		return nil, err
	}

	return receipt, nil
}

// GetNoSendTxOpts This generates a noSend TransactOpts so that we can use
// this to generate the transaction without actually sending it
func (m *SimpleTxManager) GetNoSendTxOpts() (*bind.TransactOpts, error) {
	signer, err := m.signerFn(context.Background(), m.sender)
	if err != nil {
		return nil, err
	}
	noSendTxOpts := &bind.TransactOpts{
		From:   m.sender,
		Signer: signer,
		NoSend: true,
	}
	return noSendTxOpts, nil
}

func (m *SimpleTxManager) waitForReceipt(ctx context.Context, txID wallet.TxID) (*types.Receipt, error) {
	// TODO: make this ticker adjustable
	queryTicker := time.NewTicker(2 * time.Second)
	defer queryTicker.Stop()
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-queryTicker.C:
			if receipt := m.queryReceipt(ctx, txID); receipt != nil {
				return receipt, nil
			}
		}
	}
}

func (m *SimpleTxManager) queryReceipt(ctx context.Context, txID wallet.TxID) *types.Receipt {
	txHash := common.HexToHash(txID)
	receipt, err := m.wallet.GetTransactionReceipt(ctx, txHash.Hex())
	if errors.Is(err, ethereum.NotFound) {
		m.log.Info("Transaction not yet mined", "hash", txHash)
		return nil
	} else if err != nil {
		m.log.Info("Receipt retrieval failed", "hash", txHash, "err", err)
		return nil
	} else if receipt == nil {
		m.log.Warn("Receipt and error are both nil", "hash", txHash)
		return nil
	}

	return receipt
}
