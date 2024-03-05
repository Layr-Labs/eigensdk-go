package wallet

import (
	"context"
	"fmt"
	"math/big"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/signerv2"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

var (
	FallbackGasTipCap = big.NewInt(15_000_000_000)
)

var _ Wallet = (*privateKeyWallet)(nil)

type privateKeyWallet struct {
	ethClient eth.Client
	address   common.Address
	signerFn  signerv2.SignerFn
	logger    logging.Logger

	// cache
	contracts map[common.Address]*bind.BoundContract
}

func NewPrivateKeyWallet(ethClient eth.Client, signer signerv2.SignerFn, signerAddress common.Address, logger logging.Logger) (Wallet, error) {
	return &privateKeyWallet{
		ethClient: ethClient,
		address:   signerAddress,
		signerFn:  signer,
		logger:    logger,
		contracts: make(map[common.Address]*bind.BoundContract, 0),
	}, nil
}

func (t *privateKeyWallet) SendTransaction(ctx context.Context, tx *types.Transaction) (TxID, error) {
	// Estimate gas and nonce
	// can't print tx hash in logs because the tx changes below when we complete and sign it
	// so the txHash is meaningless at this point
	t.logger.Debug("Estimating gas and nonce")
	tx, err := t.estimateGasAndNonce(ctx, tx)
	if err != nil {
		return "", err
	}

	t.logger.Debug("Getting signer for tx")
	signer, err := t.signerFn(ctx, t.address)
	if err != nil {
		return "", err
	}

	t.logger.Debug("Sending transaction")
	opts := &bind.TransactOpts{
		From:      t.address,
		Nonce:     new(big.Int).SetUint64(tx.Nonce()),
		Signer:    signer,
		Value:     tx.Value(),
		GasFeeCap: tx.GasFeeCap(),
		GasTipCap: tx.GasTipCap(),
		GasLimit:  tx.Gas(),
		Context:   ctx,
	}

	contract := t.contracts[*tx.To()]
	// if the contract has not been cached
	if contract == nil {
		// create a dummy bound contract tied to the `to` address of the transaction
		contract = bind.NewBoundContract(*tx.To(), abi.ABI{}, t.ethClient, t.ethClient, t.ethClient)
		// cache the contract for later use
		t.contracts[*tx.To()] = contract
	}

	sendingTx, err := contract.RawTransact(opts, tx.Data())
	if err != nil {
		return "", sdktypes.WrapError(fmt.Errorf("send: tx %v failed.", tx.Hash().String()), err)
	}

	return sendingTx.Hash().Hex(), nil
}

func (t *privateKeyWallet) GetTransactionReceipt(ctx context.Context, txID TxID) (*types.Receipt, error) {
	txHash := common.HexToHash(txID)
	return t.ethClient.TransactionReceipt(ctx, txHash)
}

// estimateGasAndNonce we are explicitly implementing this because
// * We want to support legacy transactions (i.e. not dynamic fee)
// * We want to support gas management, i.e. add buffer to gas limit
func (t *privateKeyWallet) estimateGasAndNonce(ctx context.Context, tx *types.Transaction) (*types.Transaction, error) {
	gasTipCap, err := t.ethClient.SuggestGasTipCap(ctx)
	if err != nil {
		// If the transaction failed because the backend does not support
		// eth_maxPriorityFeePerGas, fallback to using the default constant.
		t.logger.Info("eth_maxPriorityFeePerGas is unsupported by current backend, using fallback gasTipCap")
		gasTipCap = FallbackGasTipCap
	}

	header, err := t.ethClient.HeaderByNumber(ctx, nil)
	if err != nil {
		return nil, err
	}

	gasFeeCap := new(big.Int).Add(header.BaseFee, gasTipCap)

	gasLimit, err := t.ethClient.EstimateGas(ctx, ethereum.CallMsg{
		From:      t.address,
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
