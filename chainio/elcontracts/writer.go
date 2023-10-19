package elcontracts

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	gethcommon "github.com/ethereum/go-ethereum/common"
	gethtypes "github.com/ethereum/go-ethereum/core/types"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/chainio/utils"
	blspubkeycompendium "github.com/Layr-Labs/eigensdk-go/contracts/bindings/BLSPublicKeyCompendium"
	contractDelegationManager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/DelegationManager"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/metrics"
	"github.com/Layr-Labs/eigensdk-go/signer"
	"github.com/Layr-Labs/eigensdk-go/types"
)

type ELWriter interface {
	RegisterAsOperator(ctx context.Context, operator types.Operator) (*gethtypes.Receipt, error)

	UpdateOperatorDetails(ctx context.Context, operator types.Operator) (*gethtypes.Receipt, error)

	RegisterBLSPublicKey(
		ctx context.Context,
		blsKeyPair *bls.KeyPair,
		operator types.Operator,
		blsPubkeyCompendiumAddr common.Address,
	) (*gethtypes.Receipt, error)

	// DepositERC20IntoStrategy deposits ERC20 tokens into a strategy contract.
	DepositERC20IntoStrategy(
		ctx context.Context,
		strategyAddr gethcommon.Address,
		amount *big.Int,
	) (*gethtypes.Receipt, error)

	OptOperatorIntoSlashing(ctx context.Context, avsServiceManagerAddr gethcommon.Address) (*gethtypes.Receipt, error)
}

type ELChainWriter struct {
	elContractsClient clients.ELContractsClient
	signer            signer.Signer
	logger            logging.Logger
	ethClient         eth.EthClient
}

var _ ELWriter = (*ELChainWriter)(nil)

func NewELChainWriter(
	elContractsClient clients.ELContractsClient,
	ethClient eth.EthClient,
	signer signer.Signer,
	logger logging.Logger,
	eigenMetrics metrics.Metrics,
) *ELChainWriter {
	return &ELChainWriter{
		elContractsClient: elContractsClient,
		signer:            signer,
		logger:            logger,
		ethClient:         ethClient,
	}
}

func (w *ELChainWriter) RegisterAsOperator(ctx context.Context, operator types.Operator) (*gethtypes.Receipt, error) {
	w.logger.Infof("registering operator %s to EigenLayer", operator.Address)
	opDetails := contractDelegationManager.IDelegationManagerOperatorDetails{
		EarningsReceiver:         gethcommon.HexToAddress(operator.EarningsReceiverAddress),
		StakerOptOutWindowBlocks: operator.StakerOptOutWindowBlocks,
	}
	txOpts := w.signer.GetTxOpts()
	tx, err := w.elContractsClient.RegisterAsOperator(txOpts, opDetails, operator.MetadataUrl)
	if err != nil {
		return nil, err
	}
	w.logger.Infof("tx hash: %s", tx.Hash().String())
	receipt := w.ethClient.WaitForTransactionReceipt(ctx, tx.Hash())

	w.logger.Infof("registered operator %s with EigenLayer", operator.Address)
	return receipt, nil
}

func (w *ELChainWriter) UpdateOperatorDetails(
	ctx context.Context,
	operator types.Operator,
) (*gethtypes.Receipt, error) {
	txOpts := w.signer.GetTxOpts()

	w.logger.Infof("updating operator details of operator %s to EigenLayer", operator.Address)
	opDetails := contractDelegationManager.IDelegationManagerOperatorDetails{
		EarningsReceiver:         gethcommon.HexToAddress(operator.EarningsReceiverAddress),
		DelegationApprover:       gethcommon.HexToAddress(operator.DelegationApproverAddress),
		StakerOptOutWindowBlocks: operator.StakerOptOutWindowBlocks,
	}

	tx, err := w.elContractsClient.ModifyOperatorDetails(txOpts, opDetails)
	if err != nil {
		return nil, err
	}
	w.logger.Infof("tx hash: %s", tx.Hash().String())
	w.ethClient.WaitForTransactionReceipt(ctx, tx.Hash())

	w.logger.Infof("updated operator metadata URI for operator %s to EigenLayer", operator.Address)
	tx, err = w.elContractsClient.UpdateOperatorMetadataURI(txOpts, operator.MetadataUrl)
	if err != nil {
		return nil, err
	}
	w.logger.Infof("tx hash: %s", tx.Hash().String())
	receipt := w.ethClient.WaitForTransactionReceipt(ctx, tx.Hash())

	w.logger.Infof("updated operator details of operator %s to EigenLayer", operator.Address)
	return receipt, nil
}

func (w *ELChainWriter) DepositERC20IntoStrategy(
	ctx context.Context,
	strategyAddr gethcommon.Address,
	amount *big.Int,
) (*gethtypes.Receipt, error) {
	w.logger.Infof("depositing %s tokens into strategy %s", amount.String(), strategyAddr)
	txOpts := w.signer.GetTxOpts()
	_, underlyingTokenContract, underlyingTokenAddr, err := w.elContractsClient.GetStrategyAndUnderlyingERC20Token(
		strategyAddr,
	)
	if err != nil {
		return nil, err
	}

	strategyManagerAddr, err := w.elContractsClient.GetStrategyManagerContractAddress()
	if err != nil {
		return nil, err
	}

	tx, err := underlyingTokenContract.Approve(txOpts, strategyManagerAddr, amount)
	if err != nil {
		return nil, err
	}
	// not sure if this is necessary or if nonce management will be smart enough to queue them one after the other,
	// but playing it safe by waiting for approve tx to be mined before sending deposit tx
	w.ethClient.WaitForTransactionReceipt(ctx, tx.Hash())

	tx, err = w.elContractsClient.DepositIntoStrategy(txOpts, strategyAddr, underlyingTokenAddr, amount)
	if err != nil {
		return nil, err
	}
	receipt := w.ethClient.WaitForTransactionReceipt(ctx, tx.Hash())

	w.logger.Infof("deposited %s into strategy %s", amount.String(), strategyAddr)
	return receipt, nil
}

// operator opting into slashing is the w.signer wallet
// this is meant to be called by the operator CLI
func (w *ELChainWriter) OptOperatorIntoSlashing(
	ctx context.Context,
	avsServiceManagerAddr gethcommon.Address,
) (*gethtypes.Receipt, error) {
	txOpts := w.signer.GetTxOpts()
	tx, err := w.elContractsClient.OptIntoSlashing(txOpts, avsServiceManagerAddr)
	if err != nil {
		return nil, err
	}
	receipt := w.ethClient.WaitForTransactionReceipt(ctx, tx.Hash())

	w.logger.Infof(
		"Operator %s opted into slashing by service manager contract %s \n",
		txOpts.From,
		avsServiceManagerAddr,
	)
	return receipt, nil
}

func (w *ELChainWriter) RegisterBLSPublicKey(
	ctx context.Context,
	blsKeyPair *bls.KeyPair,
	operator types.Operator,
	blsPubkeyCompendiumAddr common.Address,
) (*gethtypes.Receipt, error) {
	w.logger.Infof("Registering BLS Public key to eigenlayer for operator %s", operator.Address)
	txOpts := w.signer.GetTxOpts()
	chainID, err := w.ethClient.ChainID(ctx)
	if err != nil {
		return nil, err
	}
	signedMsgHash := blsKeyPair.MakePubkeyRegistrationData(gethcommon.HexToAddress(operator.Address), blsPubkeyCompendiumAddr, chainID)
	signedMsgHashBN254 := blspubkeycompendium.BN254G1Point(utils.ConvertToBN254G1Point(signedMsgHash))
	G1pubkeyBN254 := blspubkeycompendium.BN254G1Point(utils.ConvertToBN254G1Point(blsKeyPair.GetPubKeyG1()))
	G2pubkeyBN254 := blspubkeycompendium.BN254G2Point(utils.ConvertToBN254G2Point(blsKeyPair.GetPubKeyG2()))
	tx, err := w.elContractsClient.RegisterBLSPublicKey(
		txOpts,
		signedMsgHashBN254,
		G1pubkeyBN254,
		G2pubkeyBN254,
	)
	if err != nil {
		return nil, err
	}
	receipt := w.ethClient.WaitForTransactionReceipt(ctx, tx.Hash())

	w.logger.Infof("Operator %s has registered BLS public key to EigenLayer \n", operator.Address)
	return receipt, nil
}
