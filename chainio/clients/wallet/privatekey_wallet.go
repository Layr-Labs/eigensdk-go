package wallet

import (
	"context"
	"fmt"

	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/signerv2"
	"github.com/Layr-Labs/eigensdk-go/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type EthBackend interface {
	SendTransaction(ctx context.Context, tx *types.Transaction) error
	TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error)
}

type privateKeyWallet struct {
	ethClient EthBackend
	address   common.Address
	signerFn  signerv2.SignerFn
	logger    logging.Logger
}

var _ Wallet = (*privateKeyWallet)(nil)

func NewPrivateKeyWallet(
	ethClient EthBackend,
	signer signerv2.SignerFn,
	signerAddress common.Address,
	logger logging.Logger,
) (Wallet, error) {
	return &privateKeyWallet{
		ethClient: ethClient,
		address:   signerAddress,
		signerFn:  signer,
		logger:    logger,
	}, nil
}

func (t *privateKeyWallet) SendTransaction(ctx context.Context, tx *types.Transaction) (TxID, error) {

	t.logger.Debug("Getting signer for tx")
	signer, err := t.signerFn(ctx, t.address)
	if err != nil {
		return "", err
	}

	t.logger.Debug("Sending transaction")
	signedTx, err := signer(t.address, tx)
	if err != nil {
		return "", utils.WrapError(fmt.Errorf("sign: tx %v failed.", tx.Hash().String()), err)
	}

	err = t.ethClient.SendTransaction(ctx, signedTx)
	if err != nil {
		return "", utils.WrapError(fmt.Errorf("send: tx %v failed.", tx.Hash().String()), err)
	}

	return signedTx.Hash().Hex(), nil
}

func (t *privateKeyWallet) GetTransactionReceipt(ctx context.Context, txID TxID) (*types.Receipt, error) {
	txHash := common.HexToHash(txID)
	return t.ethClient.TransactionReceipt(ctx, txHash)
}

func (t *privateKeyWallet) SenderAddress(ctx context.Context) (common.Address, error) {
	return t.address, nil
}
