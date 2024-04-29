package signer

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	sdkethclient "github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/utils"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethcommon "github.com/ethereum/go-ethereum/common"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
)

// exported as the default so that users can call NewBasicSigner with it if they don't know any better
var FallbackGasTipCap = big.NewInt(15000000000)

type BasicSigner struct {
	logger            logging.Logger
	ethClient         sdkethclient.Client
	privateKey        *ecdsa.PrivateKey
	accountAddress    gethcommon.Address
	contracts         map[gethcommon.Address]*bind.BoundContract
	fallbackGasTipCap *big.Int
}

func NewBasicSigner(
	privateKey *ecdsa.PrivateKey,
	ethClient sdkethclient.Client,
	logger logging.Logger,
	fallbackGasTipCap *big.Int,
) (*BasicSigner, error) {

	accountAddress, err := utils.EcdsaPrivateKeyToAddress(privateKey)
	if err != nil {
		return nil, utils.WrapError("Cannot get account address", err)
	}

	return &BasicSigner{
		logger:            logger,
		ethClient:         ethClient,
		privateKey:        privateKey,
		accountAddress:    accountAddress,
		contracts:         make(map[gethcommon.Address]*bind.BoundContract),
		fallbackGasTipCap: fallbackGasTipCap,
	}, nil
}

// TODO: is this the right place to put this?
//
//	it's needed when making calls calls using the bindings, eg
//	AvsContractBindings.TaskManager.CreateNewTask(w.NoSendTransactOpts, numToSquare)
func (s *BasicSigner) GetNoSendTransactOpts() (*bind.TransactOpts, error) {
	chainIDBigInt, err := s.ethClient.ChainID(context.Background())
	if err != nil {
		return nil, utils.WrapError("Cannot get chainId", err)
	}
	opts, err := bind.NewKeyedTransactorWithChainID(s.privateKey, chainIDBigInt)
	if err != nil {
		return nil, utils.WrapError("Cannot create NoSendTransactOpts", err)
	}
	opts.NoSend = true
	return opts, nil
}

// EstimateGasPriceAndLimitAndSendTx sends and returns an otherwise identical txn
// to the one provided but with updated gas prices sampled from the existing network
// conditions and an accurate gasLimit
//
// Note: tx must be a to a contract, not an EOA
//
// Slightly modified from:
// https://github.com/ethereum-optimism/optimism/blob/ec266098641820c50c39c31048aa4e953bece464/batch-submitter/drivers/sequencer/driver.go#L314
func (s *BasicSigner) EstimateGasPriceAndLimitAndSendTx(
	ctx context.Context,
	tx *gethtypes.Transaction,
	tag string,
) (*gethtypes.Receipt, error) {
	s.logger.Debug("Entering EstimateGasPriceAndLimitAndSendTx function...")
	defer s.logger.Debug("Exiting EstimateGasPriceAndLimitAndSendTx function...")

	gasTipCap, err := s.ethClient.SuggestGasTipCap(ctx)
	if err != nil {
		// If the transaction failed because the backend does not support
		// eth_maxPriorityFeePerGas, fallback to using the default constant.
		// Currently Alchemy is the only backend provider that exposes this
		// method, so in the event their API is unreachable we can fallback to a
		// degraded mode of operation. This also applies to our test
		// environments, as hardhat doesn't support the query either.
		s.logger.Info("eth_maxPriorityFeePerGas is unsupported by current backend, using fallback gasTipCap")
		gasTipCap = s.fallbackGasTipCap
	}

	header, err := s.ethClient.HeaderByNumber(ctx, nil)
	if err != nil {
		return nil, err
	}
	gasFeeCap := new(big.Int).Add(header.BaseFee, gasTipCap)

	// The estimated gas limits performed by RawTransact fail semi-regularly
	// with out of gas exceptions. To remedy this we extract the internal calls
	// to perform gas price/gas limit estimation here and add a buffer to
	// account for any network variability.
	gasLimit, err := s.ethClient.EstimateGas(ctx, ethereum.CallMsg{
		From:      s.accountAddress,
		To:        tx.To(),
		GasTipCap: gasTipCap,
		GasFeeCap: gasFeeCap,
		Value:     nil,
		Data:      tx.Data(),
	})

	if err != nil {
		return nil, err
	}

	opts, err := bind.NewKeyedTransactorWithChainID(s.privateKey, tx.ChainId())
	if err != nil {
		return nil, utils.WrapError("Cannot create transactOpts", err)
	}
	opts.Context = ctx
	opts.Nonce = new(big.Int).SetUint64(tx.Nonce())
	opts.GasTipCap = gasTipCap
	opts.GasFeeCap = gasFeeCap
	opts.GasLimit = addGasBuffer(gasLimit)

	contract := s.contracts[*tx.To()]
	// if the contract has not been cached
	if contract == nil {
		// create a dummy bound contract tied to the `to` address of the transaction
		contract = bind.NewBoundContract(*tx.To(), abi.ABI{}, s.ethClient, s.ethClient, s.ethClient)
		// cache the contract for later use
		s.contracts[*tx.To()] = contract
	}

	tx, err = contract.RawTransact(opts, tx.Data())
	if err != nil {
		return nil, utils.WrapError(fmt.Errorf("Failed to send transaction with tag: %v", tag), err)
	}

	receipt, err := s.EnsureTransactionEvaled(
		tx,
		tag,
	)
	if err != nil {
		return nil, err
	}

	return receipt, err
}

func (s *BasicSigner) EnsureTransactionEvaled(tx *gethtypes.Transaction, tag string) (*gethtypes.Receipt, error) {
	receipt, err := bind.WaitMined(context.Background(), s.ethClient, tx)
	if err != nil {
		return nil, utils.WrapError(fmt.Errorf("Failed to wait for transaction to mine with tag: %v", tag), err)
	}
	if receipt.Status != 1 {
		return nil, fmt.Errorf(
			"Transaction failed (tag: %v, txHash: %v, status: %v, gasUsed: %v)",
			tag,
			tx.Hash().Hex(),
			receipt.Status,
			receipt.GasUsed,
		)
	}
	s.logger.Debug("successfully submitted transaction",
		"txHash", receipt.TxHash.Hex(),
		"tag", tag,
		"gasUsed", receipt.GasUsed,
	)
	return receipt, nil
}

func addGasBuffer(gasLimit uint64) uint64 {
	return 6 * gasLimit / 5 // add 20% buffer to gas limit
}
