package elcontracts

import (
	"context"
	"errors"
	"github.com/Layr-Labs/eigensdk-go/utils"

	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethcommon "github.com/ethereum/go-ethereum/common"
	gethtypes "github.com/ethereum/go-ethereum/core/types"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	delegationmanager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/DelegationManager"
	rewardscoordinator "github.com/Layr-Labs/eigensdk-go/contracts/bindings/IRewardsCoordinator"
	slasher "github.com/Layr-Labs/eigensdk-go/contracts/bindings/ISlasher"
	strategymanager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/StrategyManager"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/metrics"
	"github.com/Layr-Labs/eigensdk-go/types"
)

type ELWriter interface {
	// RegisterAsOperator registers an operator onchain.
	RegisterAsOperator(ctx context.Context, operator types.Operator) (*gethtypes.Receipt, error)

	// UpdateOperatorDetails updates the operator details onchain.
	// This doesn't update the metadata URI. Use UpdateMetadataURI for that.
	UpdateOperatorDetails(ctx context.Context, operator types.Operator) (*gethtypes.Receipt, error)

	// UpdateMetadataURI updates the operator metadata URI onchain
	UpdateMetadataURI(ctx context.Context, uri string) (*gethtypes.Receipt, error)

	// DepositERC20IntoStrategy deposits ERC20 tokens into a strategy contract.
	DepositERC20IntoStrategy(
		ctx context.Context,
		strategyAddr gethcommon.Address,
		amount *big.Int,
	) (*gethtypes.Receipt, error)

	SetClaimerFor(
		ctx context.Context,
		claimer gethcommon.Address,
	) (*gethtypes.Receipt, error)
}

type ELChainWriter struct {
	slasher             *slasher.ContractISlasher
	delegationManager   *delegationmanager.ContractDelegationManager
	strategyManager     *strategymanager.ContractStrategyManager
	rewardsCoordinator  *rewardscoordinator.ContractIRewardsCoordinator
	strategyManagerAddr gethcommon.Address
	elChainReader       ELReader
	ethClient           eth.Client
	logger              logging.Logger
	txMgr               txmgr.TxManager
}

var _ ELWriter = (*ELChainWriter)(nil)

func NewELChainWriter(
	slasher *slasher.ContractISlasher,
	delegationManager *delegationmanager.ContractDelegationManager,
	strategyManager *strategymanager.ContractStrategyManager,
	rewardsCoordinator *rewardscoordinator.ContractIRewardsCoordinator,
	strategyManagerAddr gethcommon.Address,
	elChainReader ELReader,
	ethClient eth.Client,
	logger logging.Logger,
	eigenMetrics metrics.Metrics,
	txMgr txmgr.TxManager,
) *ELChainWriter {
	logger = logger.With("module", "elcontracts/writer")

	return &ELChainWriter{
		slasher:             slasher,
		delegationManager:   delegationManager,
		strategyManager:     strategyManager,
		strategyManagerAddr: strategyManagerAddr,
		rewardsCoordinator:  rewardsCoordinator,
		elChainReader:       elChainReader,
		logger:              logger,
		ethClient:           ethClient,
		txMgr:               txMgr,
	}
}

// BuildELChainWriter builds an ELChainWriter instance.
// Deprecated: Use NewWriterFromConfig instead.
func BuildELChainWriter(
	delegationManagerAddr gethcommon.Address,
	avsDirectoryAddr gethcommon.Address,
	ethClient eth.Client,
	logger logging.Logger,
	eigenMetrics metrics.Metrics,
	txMgr txmgr.TxManager,
) (*ELChainWriter, error) {
	elContractBindings, err := NewEigenlayerContractBindings(
		delegationManagerAddr,
		avsDirectoryAddr,
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
		elContractBindings.AvsDirectory,
		logger,
		ethClient,
	)
	return NewELChainWriter(
		elContractBindings.Slasher,
		elContractBindings.DelegationManager,
		elContractBindings.StrategyManager,
		nil,
		elContractBindings.StrategyManagerAddr,
		elChainReader,
		ethClient,
		logger,
		eigenMetrics,
		txMgr,
	), nil
}

func NewWriterFromConfig(
	cfg Config,
	ethClient eth.Client,
	logger logging.Logger,
	eigenMetrics metrics.Metrics,
	txMgr txmgr.TxManager,
) (*ELChainWriter, error) {
	elContractBindings, err := NewBindingsFromConfig(
		cfg,
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
		elContractBindings.AvsDirectory,
		logger,
		ethClient,
	)
	return NewELChainWriter(
		elContractBindings.Slasher,
		elContractBindings.DelegationManager,
		elContractBindings.StrategyManager,
		elContractBindings.RewardsCoordinator,
		elContractBindings.StrategyManagerAddr,
		elChainReader,
		ethClient,
		logger,
		eigenMetrics,
		txMgr,
	), nil
}

func (w *ELChainWriter) RegisterAsOperator(ctx context.Context, operator types.Operator) (*gethtypes.Receipt, error) {
	if w.delegationManager == nil {
		return nil, errors.New("DelegationManager contract not provided")
	}

	w.logger.Infof("registering operator %s to EigenLayer", operator.Address)
	opDetails := delegationmanager.IDelegationManagerOperatorDetails{
		DeprecatedEarningsReceiver: gethcommon.HexToAddress(operator.EarningsReceiverAddress),
		StakerOptOutWindowBlocks:   operator.StakerOptOutWindowBlocks,
		DelegationApprover:         gethcommon.HexToAddress(operator.DelegationApproverAddress),
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
	w.logger.Info("tx successfully included", "txHash", receipt.TxHash.String())

	return receipt, nil
}

func (w *ELChainWriter) UpdateOperatorDetails(
	ctx context.Context,
	operator types.Operator,
) (*gethtypes.Receipt, error) {
	if w.delegationManager == nil {
		return nil, errors.New("DelegationManager contract not provided")
	}

	w.logger.Infof("updating operator details of operator %s to EigenLayer", operator.Address)
	opDetails := delegationmanager.IDelegationManagerOperatorDetails{
		DeprecatedEarningsReceiver: gethcommon.HexToAddress(operator.EarningsReceiverAddress),
		DelegationApprover:         gethcommon.HexToAddress(operator.DelegationApproverAddress),
		StakerOptOutWindowBlocks:   operator.StakerOptOutWindowBlocks,
	}

	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, err
	}

	tx, err := w.delegationManager.ModifyOperatorDetails(noSendTxOpts, opDetails)
	if err != nil {
		return nil, err
	}
	receipt, err := w.txMgr.Send(ctx, tx)
	if err != nil {
		return nil, errors.New("failed to send tx with err: " + err.Error())
	}
	w.logger.Info(
		"successfully updated operator details",
		"txHash",
		receipt.TxHash.String(),
		"operator",
		operator.Address,
	)

	return receipt, nil
}

func (w *ELChainWriter) UpdateMetadataURI(ctx context.Context, uri string) (*gethtypes.Receipt, error) {
	if w.delegationManager == nil {
		return nil, errors.New("DelegationManager contract not provided")
	}

	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, err
	}

	tx, err := w.delegationManager.UpdateOperatorMetadataURI(noSendTxOpts, uri)
	if err != nil {
		return nil, err
	}
	receipt, err := w.txMgr.Send(ctx, tx)
	if err != nil {
		return nil, errors.New("failed to send tx with err: " + err.Error())
	}
	w.logger.Info(
		"successfully updated operator metadata uri",
		"txHash",
		receipt.TxHash.String(),
	)

	return receipt, nil
}

func (w *ELChainWriter) DepositERC20IntoStrategy(
	ctx context.Context,
	strategyAddr gethcommon.Address,
	amount *big.Int,
) (*gethtypes.Receipt, error) {
	if w.strategyManager == nil {
		return nil, errors.New("StrategyManager contract not provided")
	}

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
		return nil, errors.Join(errors.New("failed to approve token transfer"), err)
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

func (w *ELChainWriter) SetClaimerFor(
	ctx context.Context,
	claimer gethcommon.Address,
) (*gethtypes.Receipt, error) {
	if w.rewardsCoordinator == nil {
		return nil, errors.New("RewardsCoordinator contract not provided")
	}

	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, err
	}

	tx, err := w.rewardsCoordinator.SetClaimerFor(noSendTxOpts, claimer)
	if err != nil {
		return nil, err
	}
	receipt, err := w.txMgr.Send(ctx, tx)
	if err != nil {
		return nil, utils.WrapError("failed to send tx", err)
	}

	return receipt, nil
}
