package elcontracts

import (
	"context"
	"math/big"

	gethcommon "github.com/ethereum/go-ethereum/common"
	gethtypes "github.com/ethereum/go-ethereum/core/types"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/chainio/utils"
	chainioutils "github.com/Layr-Labs/eigensdk-go/chainio/utils"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/metrics"
	"github.com/Layr-Labs/eigensdk-go/signer"
	"github.com/Layr-Labs/eigensdk-go/types"

	blspubkeycompendium "github.com/Layr-Labs/eigensdk-go/contracts/bindings/BLSPublicKeyCompendium"
	delegationmanager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/DelegationManager"
	slasher "github.com/Layr-Labs/eigensdk-go/contracts/bindings/Slasher"
	strategymanager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/StrategyManager"
)

type ELWriter interface {
	RegisterAsOperator(ctx context.Context, operator types.Operator) (*gethtypes.Receipt, error)

	UpdateOperatorDetails(ctx context.Context, operator types.Operator) (*gethtypes.Receipt, error)

	RegisterBLSPublicKey(
		ctx context.Context,
		blsKeyPair *bls.KeyPair,
		operator types.Operator,
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
	slasher                 slasher.ContractSlasherTransacts
	delegationManager       delegationmanager.ContractDelegationManagerTransacts
	strategyManager         strategymanager.ContractStrategyManagerTransacts
	strategyManagerAddr     gethcommon.Address
	blsPubkeyCompendium     blspubkeycompendium.ContractBLSPublicKeyCompendiumTransacts
	blsPubkeyCompendiumAddr gethcommon.Address
	elChainReader           ELReader
	ethClient               eth.EthClient
	signer                  signer.Signer
	logger                  logging.Logger
}

var _ ELWriter = (*ELChainWriter)(nil)

func NewELChainWriter(
	slasher slasher.ContractSlasherTransacts,
	delegationManager delegationmanager.ContractDelegationManagerTransacts,
	strategyManager strategymanager.ContractStrategyManagerTransacts,
	strategyManagerAddr gethcommon.Address,
	blsPubkeyCompendium blspubkeycompendium.ContractBLSPublicKeyCompendiumTransacts,
	blsPubkeyCompendiumAddr gethcommon.Address,
	elChainReader ELReader,
	ethClient eth.EthClient,
	signer signer.Signer,
	logger logging.Logger,
	eigenMetrics metrics.Metrics,
) *ELChainWriter {
	return &ELChainWriter{
		slasher:                 slasher,
		delegationManager:       delegationManager,
		strategyManager:         strategyManager,
		strategyManagerAddr:     strategyManagerAddr,
		blsPubkeyCompendium:     blsPubkeyCompendium,
		blsPubkeyCompendiumAddr: blsPubkeyCompendiumAddr,
		elChainReader:           elChainReader,
		signer:                  signer,
		logger:                  logger,
		ethClient:               ethClient,
	}
}

func BuildELChainWriter(
	slasherAddr gethcommon.Address,
	blsPubKeyCompendiumAddr gethcommon.Address,
	ethClient eth.EthClient,
	signer signer.Signer,
	logger logging.Logger,
	eigenMetrics metrics.Metrics,
) (*ELChainWriter, error) {
	elContractBindings, err := chainioutils.NewEigenlayerContractBindings(slasherAddr, blsPubKeyCompendiumAddr, ethClient, logger)
	if err != nil {
		return nil, err
	}
	elChainReader := NewELChainReader(
		elContractBindings.Slasher,
		elContractBindings.DelegationManager,
		elContractBindings.StrategyManager,
		elContractBindings.BlsPubkeyCompendium,
		blsPubKeyCompendiumAddr,
		logger,
		ethClient,
	)
	return NewELChainWriter(
		elContractBindings.Slasher,
		elContractBindings.DelegationManager,
		elContractBindings.StrategyManager,
		elContractBindings.StrategyManagerAddr,
		elContractBindings.BlsPubkeyCompendium,
		blsPubKeyCompendiumAddr,
		elChainReader,
		ethClient,
		signer,
		logger,
		eigenMetrics,
	), nil
}

// TODO(madhur): do we really want to wait for txreceipts like this in these functions?
// feels like this should be something decided by callers
func (w *ELChainWriter) RegisterAsOperator(ctx context.Context, operator types.Operator) (*gethtypes.Receipt, error) {
	w.logger.Infof("registering operator %s to EigenLayer", operator.Address)
	opDetails := delegationmanager.IDelegationManagerOperatorDetails{
		EarningsReceiver:         gethcommon.HexToAddress(operator.EarningsReceiverAddress),
		StakerOptOutWindowBlocks: operator.StakerOptOutWindowBlocks,
	}
	txOpts := w.signer.GetTxOpts()
	tx, err := w.delegationManager.RegisterAsOperator(txOpts, opDetails, operator.MetadataUrl)
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
	opDetails := delegationmanager.IDelegationManagerOperatorDetails{
		EarningsReceiver:         gethcommon.HexToAddress(operator.EarningsReceiverAddress),
		DelegationApprover:       gethcommon.HexToAddress(operator.DelegationApproverAddress),
		StakerOptOutWindowBlocks: operator.StakerOptOutWindowBlocks,
	}

	tx, err := w.delegationManager.ModifyOperatorDetails(txOpts, opDetails)
	if err != nil {
		return nil, err
	}
	w.logger.Infof("tx hash: %s", tx.Hash().String())
	w.ethClient.WaitForTransactionReceipt(ctx, tx.Hash())

	w.logger.Infof("updated operator metadata URI for operator %s to EigenLayer", operator.Address)
	tx, err = w.delegationManager.UpdateOperatorMetadataURI(txOpts, operator.MetadataUrl)
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
	_, underlyingTokenContract, underlyingTokenAddr, err := w.elChainReader.GetStrategyAndUnderlyingERC20Token(
		ctx,
		strategyAddr,
	)
	if err != nil {
		return nil, err
	}

	tx, err := underlyingTokenContract.Approve(txOpts, w.strategyManagerAddr, amount)
	if err != nil {
		return nil, err
	}
	// not sure if this is necessary or if nonce management will be smart enough to queue them one after the other,
	// but playing it safe by waiting for approve tx to be mined before sending deposit tx
	w.ethClient.WaitForTransactionReceipt(ctx, tx.Hash())

	tx, err = w.strategyManager.DepositIntoStrategy(txOpts, strategyAddr, underlyingTokenAddr, amount)
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
	tx, err := w.slasher.OptIntoSlashing(txOpts, avsServiceManagerAddr)
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
) (*gethtypes.Receipt, error) {
	w.logger.Infof("Registering BLS Public key to eigenlayer for operator %s", operator.Address)
	txOpts := w.signer.GetTxOpts()
	chainID, err := w.ethClient.ChainID(ctx)
	if err != nil {
		return nil, err
	}
	signedMsgHash := blsKeyPair.MakePubkeyRegistrationData(
		gethcommon.HexToAddress(operator.Address),
		w.blsPubkeyCompendiumAddr,
		chainID,
	)
	signedMsgHashBN254 := blspubkeycompendium.BN254G1Point(utils.ConvertToBN254G1Point(signedMsgHash))
	G1pubkeyBN254 := blspubkeycompendium.BN254G1Point(utils.ConvertToBN254G1Point(blsKeyPair.GetPubKeyG1()))
	G2pubkeyBN254 := blspubkeycompendium.BN254G2Point(utils.ConvertToBN254G2Point(blsKeyPair.GetPubKeyG2()))
	tx, err := w.blsPubkeyCompendium.RegisterBLSPublicKey(
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
