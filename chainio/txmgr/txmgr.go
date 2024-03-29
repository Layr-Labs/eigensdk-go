package txmgr

import (
	"context"
	"errors"
	"math/big"
	"time"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/wallet"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
)

var (
	// 5 gwei in case the backend does not support eth_maxPriorityFeePerGas (no idea if/when this ever happens..)
	FallbackGasTipCap = big.NewInt(5_000_000_000)
	// 1.20x gas limit multiplier. This is arbitrary but should be safe for most cases
	FallbackGasLimitMultiplier = 1.20
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
	wallet             wallet.Wallet
	client             eth.Client
	log                logging.Logger
	sender             common.Address
	gasLimitMultiplier float64
}

var _ TxManager = (*SimpleTxManager)(nil)

// NewSimpleTxManager creates a new simpleTxManager which can be used
// to send a transaction to smart contracts on the Ethereum node
func NewSimpleTxManager(
	wallet wallet.Wallet,
	client eth.Client,
	log logging.Logger,
	sender common.Address,
) *SimpleTxManager {
	return &SimpleTxManager{
		wallet:             wallet,
		client:             client,
		log:                log,
		sender:             sender,
		gasLimitMultiplier: FallbackGasLimitMultiplier,
	}
}

func (m *SimpleTxManager) WithGasLimitMultiplier(multiplier float64) *SimpleTxManager {
	m.gasLimitMultiplier = multiplier
	return m
}

// Send is used to send a transaction to the Ethereum node. It takes an unsigned/signed transaction
// and then sends it to the Ethereum node.
// It also takes care of gas estimation and adds a buffer to the gas limit
// If you pass in a signed transaction it will ignore the signature
// and resign the transaction after adding the nonce and gas limit.
// To check out the whole flow on how this works, check out the README.md in this folder
func (m *SimpleTxManager) Send(ctx context.Context, tx *types.Transaction) (*types.Receipt, error) {
	// Estimate gas and nonce
	// can't print tx hash in logs because the tx changes below when we complete and sign it
	// so the txHash is meaningless at this point
	m.log.Debug("Estimating gas and nonce")
	tx, err := m.estimateGasAndNonce(ctx, tx)
	if err != nil {
		return nil, err
	}
	bumpedGasTx := &types.DynamicFeeTx{
		To:        tx.To(),
		Nonce:     tx.Nonce(),
		GasFeeCap: tx.GasFeeCap(),
		GasTipCap: tx.GasTipCap(),
		Gas:       uint64(float64(tx.Gas()) * m.gasLimitMultiplier),
		Value:     tx.Value(),
		Data:      tx.Data(),
	}
	txID, err := m.wallet.SendTransaction(ctx, types.NewTx(bumpedGasTx))
	if err != nil {
		return nil, errors.Join(errors.New("send: failed to estimate gas and nonce"), err)
	}

	receipt, err := m.waitForReceipt(ctx, txID)
	if err != nil {
		log.Info("Transaction receipt not found", "err", err)
		return nil, err
	}

	return receipt, nil
}

func NoopSigner(addr common.Address, tx *types.Transaction) (*types.Transaction, error) {
	return tx, nil
}

// GetNoSendTxOpts This generates a noSend TransactOpts so that we can use
// this to generate the transaction without actually sending it
func (m *SimpleTxManager) GetNoSendTxOpts() (*bind.TransactOpts, error) {
	return &bind.TransactOpts{
		From:   m.sender,
		NoSend: true,
		Signer: NoopSigner,
	}, nil
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
	receipt, err := m.wallet.GetTransactionReceipt(ctx, txID)
	if errors.Is(err, ethereum.NotFound) {
		m.log.Info("Transaction not yet mined", "txID", txID)
		return nil
	} else if err != nil {
		m.log.Info("Receipt retrieval failed", "txID", txID, "err", err)
		return nil
	} else if receipt == nil {
		m.log.Warn("Receipt and error are both nil", "txID", txID)
		return nil
	}

	return receipt
}

// estimateGasAndNonce we are explicitly implementing this because
// * We want to support legacy transactions (i.e. not dynamic fee)
// * We want to support gas management, i.e. add buffer to gas limit
func (m *SimpleTxManager) estimateGasAndNonce(ctx context.Context, tx *types.Transaction) (*types.Transaction, error) {
	gasTipCap, err := m.client.SuggestGasTipCap(ctx)
	if err != nil {
		// If the transaction failed because the backend does not support
		// eth_maxPriorityFeePerGas, fallback to using the default constant.
		m.log.Info("eth_maxPriorityFeePerGas is unsupported by current backend, using fallback gasTipCap")
		gasTipCap = FallbackGasTipCap
	}

	header, err := m.client.HeaderByNumber(ctx, nil)
	if err != nil {
		return nil, err
	}

	// 2*baseFee + gasTipCap makes sure that the tx remains includeable for 6 consecutive 100% full blocks.
	// see https://www.blocknative.com/blog/eip-1559-fees
	gasFeeCap := new(big.Int).Add(header.BaseFee.Mul(header.BaseFee, big.NewInt(2)), gasTipCap)

	gasLimit := tx.Gas()
	// we only estimate if gasLimit is not already set
	if gasLimit == 0 {
		from, err := m.wallet.SenderAddress(ctx)
		if err != nil {
			return nil, errors.Join(errors.New("send: failed to get sender address"), err)
		}
		gasLimit, err = m.client.EstimateGas(ctx, ethereum.CallMsg{
			From:      from,
			To:        tx.To(),
			GasTipCap: gasTipCap,
			GasFeeCap: gasFeeCap,
			Value:     tx.Value(),
			Data:      tx.Data(),
		})
		if err != nil {
			return nil, errors.Join(errors.New("send: failed to estimate gas"), err)
		}
	}

	rawTx := &types.DynamicFeeTx{
		ChainID:   tx.ChainId(),
		To:        tx.To(),
		GasTipCap: gasTipCap,
		GasFeeCap: gasFeeCap,
		Data:      tx.Data(),
		Value:     tx.Value(),
		Gas:       uint64(float64(gasLimit) * m.gasLimitMultiplier),
		Nonce:     tx.Nonce(), // We are not doing any nonce management for now but we probably should later for more robustness
	}

	return types.NewTx(rawTx), nil
}
