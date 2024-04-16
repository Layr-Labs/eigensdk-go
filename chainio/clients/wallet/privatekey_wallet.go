package wallet

import (
	"context"
	"fmt"
	"math/big"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/signerv2"
	"github.com/Layr-Labs/eigensdk-go/utils"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
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
		return "", utils.WrapError(fmt.Errorf("send: tx %v failed.", tx.Hash().String()), err)
	}

	return sendingTx.Hash().Hex(), nil
}

func (t *privateKeyWallet) GetTransactionReceipt(ctx context.Context, txID TxID) (*types.Receipt, error) {
	txHash := common.HexToHash(txID)
	return t.ethClient.TransactionReceipt(ctx, txHash)
}

func (t *privateKeyWallet) SenderAddress(ctx context.Context) (common.Address, error) {
	return t.address, nil
}
