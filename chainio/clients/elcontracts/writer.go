package elcontracts

import (
	"context"
	"errors"
	"math/big"

	"github.com/Layr-Labs/eigensdk-go-master/chainio/txmgr"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethcommon "github.com/ethereum/go-ethereum/common"
	gethtypes "github.com/ethereum/go-ethereum/core/types"

	"github.com/Layr-Labs/eigensdk-go-master/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go-master/chainio/utils"
	chainioutils "github.com/Layr-Labs/eigensdk-go-master/chainio/utils"
	"github.com/Layr-Labs/eigensdk-go-master/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go-master/logging"
	"github.com/Layr-Labs/eigensdk-go-master/metrics"
	"github.com/Layr-Labs/eigensdk-go-master/types"

	blspubkeycompendium "github.com/Layr-Labs/eigensdk-go-master/contracts/bindings/BLSPublicKeyCompendium"
	delegationmanager "github.com/Layr-Labs/eigensdk-go-master/contracts/bindings/DelegationManager"
	slasher "github.com/Layr-Labs/eigensdk-go-master/contracts/bindings/Slasher"
	strategymanager "github.com/Layr-Labs/eigensdk-go-master/contracts/bindings/StrategyManager"
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
	logger                  logging.Logger
	txMgr                   txmgr.TxManager
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
	logger logging.Logger,
	eigenMetrics metrics.Metrics,
	txMgr txmgr.TxManager,
) *ELChainWriter {
	return &ELChainWriter{
		slasher:                 slasher,
		delegationManager:       delegationManager,
		strategyManager:         strategyManager,
		strategyManagerAddr:     strategyManagerAddr,
		blsPubkeyCompendium:     blsPubkeyCompendium,
		blsPubkeyCompendiumAddr: blsPubkeyCompendiumAddr,
		elChainReader:           elChainReader,
		logger:                  logger,
		ethClient:               ethClient,
		txMgr:                   txMgr,
	}
}

func BuildELChainWriter(
	slasherAddr gethcommon.Address,
	blsPubKeyCompendiumAddr gethcommon.Address,
	ethClient eth.EthClient,
	logger logging.Logger,
	eigenMetrics metrics.Metrics,
	txMgr txmgr.TxManager,
) (*ELChainWriter, error) {
	elContractBindings, err := chainioutils.NewEigenlayerContractBindings(
		slasherAddr,
		blsPubKeyCompendiumAddr,
		ethClient,
		logger,
	)
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
		logger,
		eigenMetrics,
		txMgr,
	), nil
}

// TODO(madhur): we wait for txreceipts in these functions right now, but
// this will be changed once we have a better tx manager design implemented
// see https://github.com/Layr-Labs/eigensdk-go-master/pull/75
func (w *ELChainWriter) RegisterAsOperator(ctx context.Context, operator types.Operator) (*gethtypes.Receipt, error) {
	w.logger.Infof("registering operator %s to EigenLayer", operator.Address)
	opDetails := delegationmanager.IDelegationManagerOperatorDetails{
		EarningsReceiver:         gethcommon.HexToAddress(operator.EarningsReceiverAddress),
		StakerOptOutWindowBlocks: operator.StakerOptOutWindowBlocks,
	}

	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, err
	}
	tx, err := w.delegationManager.RegisterAsOperator(noSendTxOpts, opDetails, operator.MetadataUrl)
	if err != nil {
		return nil, err
	}
	receipt, err := w.txMgr.Send(ctx, tx)
	if err != nil {
		return nil, errors.New("failed to send tx with err: " + err.Error())
	}
	w.logger.Infof("tx hash: %s", tx.Hash().String())

	return receipt, nil
}

func (w *ELChainWriter) UpdateOperatorDetails(
	ctx context.Context,
	operator types.Operator,
) (*gethtypes.Receipt, error) {

	w.logger.Infof("updating operator details of operator %s to EigenLayer", operator.Address)
	opDetails := delegationmanager.IDelegationManagerOperatorDetails{
		EarningsReceiver:         gethcommon.HexToAddress(operator.EarningsReceiverAddress),
		DelegationApprover:       gethcommon.HexToAddress(operator.DelegationApproverAddress),
		StakerOptOutWindowBlocks: operator.StakerOptOutWindowBlocks,
	}

	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, err
	}

	tx, err := w.delegationManager.ModifyOperatorDetails(noSendTxOpts, opDetails)
	if err != nil {
		return nil, err
	}
	_, err = w.txMgr.Send(ctx, tx)
	if err != nil {
		return nil, errors.New("failed to send tx with err: " + err.Error())
	}
	w.logger.Infof("tx hash: %s", tx.Hash().String())
	w.logger.Infof("updated operator metadata URI for operator %s to EigenLayer", operator.Address)

	tx, err = w.delegationManager.UpdateOperatorMetadataURI(noSendTxOpts, operator.MetadataUrl)
	if err != nil {
		return nil, err
	}
	receipt, err := w.txMgr.Send(ctx, tx)
	if err != nil {
		return nil, errors.New("failed to send tx with err: " + err.Error())
	}
	w.logger.Infof("tx hash: %s", tx.Hash().String())

	w.logger.Infof("updated operator details of operator %s to EigenLayer", operator.Address)
	return receipt, nil
}

func (w *ELChainWriter) DepositERC20IntoStrategy(
	ctx context.Context,
	strategyAddr gethcommon.Address,
	amount *big.Int,
) (*gethtypes.Receipt, error) {
	w.logger.Infof("depositing %s tokens into strategy %s", amount.String(), strategyAddr)
	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, err
	}
	_, underlyingTokenContract, underlyingTokenAddr, err := w.elChainReader.GetStrategyAndUnderlyingERC20Token(
		&bind.CallOpts{Context: ctx},
		strategyAddr,
	)
	if err != nil {
		return nil, err
	}

	tx, err := underlyingTokenContract.Approve(noSendTxOpts, w.strategyManagerAddr, amount)
	if err != nil {
		return nil, err
	}
	_, err = w.txMgr.Send(ctx, tx)
	if err != nil {
		return nil, errors.New("failed to send tx with err: " + err.Error())
	}

	tx, err = w.strategyManager.DepositIntoStrategy(noSendTxOpts, strategyAddr, underlyingTokenAddr, amount)
	if err != nil {
		return nil, err
	}
	receipt, err := w.txMgr.Send(ctx, tx)
	if err != nil {
		return nil, errors.New("failed to send tx with err: " + err.Error())
	}

	w.logger.Infof("deposited %s into strategy %s", amount.String(), strategyAddr)
	return receipt, nil
}

// OptOperatorIntoSlashing operator opting into slashing is the w.signer wallet
// this is meant to be called by the operator CLI
func (w *ELChainWriter) OptOperatorIntoSlashing(
	ctx context.Context,
	avsServiceManagerAddr gethcommon.Address,
) (*gethtypes.Receipt, error) {
	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, err
	}
	tx, err := w.slasher.OptIntoSlashing(noSendTxOpts, avsServiceManagerAddr)
	if err != nil {
		return nil, err
	}
	receipt, err := w.txMgr.Send(ctx, tx)
	if err != nil {
		return nil, errors.New("failed to send tx with err: " + err.Error())
	}

	w.logger.Infof(
		"Operator %s opted into slashing by service manager contract %s \n",
		noSendTxOpts.From,
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
	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, err
	}
	chainID, err := w.ethClient.ChainID(ctx)
	if err != nil {
		return nil, err
	}
	signedMsgHash := blsKeyPair.MakePubkeyRegistrationData(
		gethcommon.HexToAddress(operator.Address),
		w.blsPubkeyCompendiumAddr,
		chainID,
	)
	signedMsgHashBN254 := utils.ConvertToBN254G1Point(signedMsgHash)
	G1pubkeyBN254 := utils.ConvertToBN254G1Point(blsKeyPair.GetPubKeyG1())
	G2pubkeyBN254 := utils.ConvertToBN254G2Point(blsKeyPair.GetPubKeyG2())
	tx, err := w.blsPubkeyCompendium.RegisterBLSPublicKey(
		noSendTxOpts,
		signedMsgHashBN254,
		G1pubkeyBN254,
		G2pubkeyBN254,
	)
	if err != nil {
		return nil, err
	}
	receipt, err := w.txMgr.Send(ctx, tx)
	if err != nil {
		return nil, errors.New("failed to send tx with err: " + err.Error())
	}

	w.logger.Infof("Operator %s has registered BLS public key to EigenLayer \n", operator.Address)
	return receipt, nil
}
