package txmgr

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/signerv2"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
)

var (
	FallbackGasTipCap = big.NewInt(15_000_000_000)
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

type EthBackend interface {
	bind.ContractBackend

	// TransactionReceipt queries the backend for a receipt associated with
	// txHash. If lookup does not fail, but the transaction is not found,
	// nil should be returned for both values.
	TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error)

	ChainID(ctx context.Context) (*big.Int, error)
}

type SimpleTxManager struct {
	backend   EthBackend
	signerFn  signerv2.SignerFn
	log       logging.Logger
	sender    common.Address
	contracts map[common.Address]*bind.BoundContract
}

var _ TxManager = (*SimpleTxManager)(nil)

// NewSimpleTxManager creates a new simpleTxManager which can be used
// to send a transaction to smart contracts on the Ethereum node
func NewSimpleTxManager(
	backend EthBackend,
	log logging.Logger,
	signerFn signerv2.SignerFn,
	sender common.Address,
) *SimpleTxManager {
	return &SimpleTxManager{
		backend:   backend,
		log:       log,
		signerFn:  signerFn,
		sender:    sender,
		contracts: map[common.Address]*bind.BoundContract{},
	}
}

// Send is used to send a transaction to the Ethereum node. It takes an unsigned/signed transaction
// and then sends it to the Ethereum node.
// It also takes care of gas estimation and adds a buffer to the gas limit
// If you pass in a signed transaction it will ignore the signature
// and resign the transaction after adding the nonce and gas limit.
// To check out the whole flow on how this works, check out the README.md in this folder
func (m *SimpleTxManager) Send(ctx context.Context, tx *types.Transaction) (*types.Receipt, error) {

	// Estimate gas and nonce
	m.log.Debug("Estimating gas and nonce", "tx", tx.Hash().Hex())
	tx, err := m.estimateGasAndNonce(ctx, tx)
	if err != nil {
		return nil, err
	}

	m.log.Debug("Getting signer for tx", "tx", tx.Hash().Hex())
	signer, err := m.signerFn(ctx, m.sender)
	if err != nil {
		return nil, err
	}

	m.log.Debug("Sending transaction", "tx", tx.Hash().Hex())
	opts := &bind.TransactOpts{
		From:      m.sender,
		Nonce:     new(big.Int).SetUint64(tx.Nonce()),
		Signer:    signer,
		Value:     tx.Value(),
		GasFeeCap: tx.GasFeeCap(),
		GasTipCap: tx.GasTipCap(),
		GasLimit:  tx.Gas(),
		Context:   ctx,
	}

	contract := m.contracts[*tx.To()]
	// if the contract has not been cached
	if contract == nil {
		// create a dummy bound contract tied to the `to` address of the transaction
		contract = bind.NewBoundContract(*tx.To(), abi.ABI{}, m.backend, m.backend, m.backend)
		// cache the contract for later use
		m.contracts[*tx.To()] = contract
	}

	tx, err = contract.RawTransact(opts, tx.Data())
	if err != nil {
		return nil, fmt.Errorf("send: failed to send txn: %w", err)
	}
	if err != nil {
		return nil, err
	}

	receipt := m.waitForTx(ctx, tx)

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

// estimateGasAndNonce we are explicitly implementing this because
// * We want to support legacy transactions (i.e. not dynamic fee)
// * We want to support gas management, i.e. add buffer to gas limit
func (m *SimpleTxManager) estimateGasAndNonce(ctx context.Context, tx *types.Transaction) (*types.Transaction, error) {
	gasTipCap, err := m.backend.SuggestGasTipCap(ctx)
	if err != nil {
		// If the transaction failed because the backend does not support
		// eth_maxPriorityFeePerGas, fallback to using the default constant.
		m.log.Info("eth_maxPriorityFeePerGas is unsupported by current backend, using fallback gasTipCap")
		gasTipCap = FallbackGasTipCap
	}

	header, err := m.backend.HeaderByNumber(ctx, nil)
	if err != nil {
		return nil, err
	}

	gasFeeCap := new(big.Int).Add(header.BaseFee, gasTipCap)

	gasLimit, err := m.backend.EstimateGas(ctx, ethereum.CallMsg{
		From:      m.sender,
		To:        tx.To(),
		GasTipCap: gasTipCap,
		GasFeeCap: gasFeeCap,
		Value:     tx.Value(),
		Data:      tx.Data(),
	})
	if err != nil {
		return nil, err
	}

	rawTx := &types.DynamicFeeTx{
		ChainID:   tx.ChainId(),
		To:        tx.To(),
		GasTipCap: gasTipCap,
		GasFeeCap: gasFeeCap,
		Data:      tx.Data(),
		Value:     tx.Value(),
		Gas:       gasLimit,   // TODO(add buffer)
		Nonce:     tx.Nonce(), // We are not doing any nonce management for now but we probably should later for more robustness
	}

	return types.NewTx(rawTx), nil
}

// waitForTx calls waitMined, and then return the receipt
func (m *SimpleTxManager) waitForTx(ctx context.Context, tx *types.Transaction) *types.Receipt {
	// Poll for the transaction to be ready & then send the result to receiptChan
	receipt, err := m.waitMined(ctx, tx)
	if err != nil {
		log.Info("Transaction receipt not found", "err", err)
		return nil
	}
	return receipt
}

// waitMined waits for the transaction to be mined or for the context to be cancelled.
func (m *SimpleTxManager) waitMined(ctx context.Context, tx *types.Transaction) (*types.Receipt, error) {
	txHash := tx.Hash()
	queryTicker := time.NewTicker(2 * time.Second)
	defer queryTicker.Stop()
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-queryTicker.C:
			if receipt := m.queryReceipt(ctx, txHash); receipt != nil {
				return receipt, nil
			}
		}
	}
}

// queryReceipt queries for the receipt and returns the receipt
func (m *SimpleTxManager) queryReceipt(ctx context.Context, txHash common.Hash) *types.Receipt {
	receipt, err := m.backend.TransactionReceipt(ctx, txHash)
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
