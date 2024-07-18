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
	erc20 "github.com/Layr-Labs/eigensdk-go/contracts/bindings/IERC20"
	rewardscoordinator "github.com/Layr-Labs/eigensdk-go/contracts/bindings/IRewardsCoordinator"
	slasher "github.com/Layr-Labs/eigensdk-go/contracts/bindings/ISlasher"
	strategy "github.com/Layr-Labs/eigensdk-go/contracts/bindings/IStrategy"
	strategymanager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/StrategyManager"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/metrics"
	"github.com/Layr-Labs/eigensdk-go/types"
)

type Reader interface {
	GetStrategyAndUnderlyingERC20Token(
		opts *bind.CallOpts, strategyAddr gethcommon.Address,
	) (*strategy.ContractIStrategy, erc20.ContractIERC20Methods, gethcommon.Address, error)
}

type ChainWriter struct {
	slasher             *slasher.ContractISlasher
	delegationManager   *delegationmanager.ContractDelegationManager
	strategyManager     *strategymanager.ContractStrategyManager
	rewardsCoordinator  *rewardscoordinator.ContractIRewardsCoordinator
	strategyManagerAddr gethcommon.Address
	elChainReader       Reader
	ethClient           eth.Client
	logger              logging.Logger
	txMgr               txmgr.TxManager
}

func NewChainWriter(
	slasher *slasher.ContractISlasher,
	delegationManager *delegationmanager.ContractDelegationManager,
	strategyManager *strategymanager.ContractStrategyManager,
	rewardsCoordinator *rewardscoordinator.ContractIRewardsCoordinator,
	strategyManagerAddr gethcommon.Address,
	elChainReader Reader,
	ethClient eth.Client,
	logger logging.Logger,
	eigenMetrics metrics.Metrics,
	txMgr txmgr.TxManager,
) *ChainWriter {
	logger = logger.With(logging.ComponentKey, "elcontracts/writer")

	return &ChainWriter{
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

// BuildELChainWriter builds an ChainWriter instance.
// Deprecated: Use NewWriterFromConfig instead.
func BuildELChainWriter(
	delegationManagerAddr gethcommon.Address,
	avsDirectoryAddr gethcommon.Address,
	ethClient eth.Client,
	logger logging.Logger,
	eigenMetrics metrics.Metrics,
	txMgr txmgr.TxManager,
) (*ChainWriter, error) {
	elContractBindings, err := NewEigenlayerContractBindings(
		delegationManagerAddr,
		avsDirectoryAddr,
		ethClient,
		logger,
	)
	if err != nil {
		return nil, err
	}
	elChainReader := NewChainReader(
		elContractBindings.Slasher,
		elContractBindings.DelegationManager,
		elContractBindings.StrategyManager,
		elContractBindings.AvsDirectory,
		elContractBindings.RewardsCoordinator,
		logger,
		ethClient,
	)
	return NewChainWriter(
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

func NewWriterFromConfig(
	cfg Config,
	ethClient eth.Client,
	logger logging.Logger,
	eigenMetrics metrics.Metrics,
	txMgr txmgr.TxManager,
) (*ChainWriter, error) {
	elContractBindings, err := NewBindingsFromConfig(
		cfg,
		ethClient,
		logger,
	)
	if err != nil {
		return nil, err
	}
	elChainReader := NewChainReader(
		elContractBindings.Slasher,
		elContractBindings.DelegationManager,
		elContractBindings.StrategyManager,
		elContractBindings.AvsDirectory,
		elContractBindings.RewardsCoordinator,
		logger,
		ethClient,
	)
	return NewChainWriter(
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

func (w *ChainWriter) RegisterAsOperator(ctx context.Context, operator types.Operator) (*gethtypes.Receipt, error) {
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

func (w *ChainWriter) UpdateOperatorDetails(
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

func (w *ChainWriter) UpdateMetadataURI(ctx context.Context, uri string) (*gethtypes.Receipt, error) {
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

func (w *ChainWriter) DepositERC20IntoStrategy(
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

func (w *ChainWriter) SetClaimerFor(
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

func (w *ChainWriter) ProcessClaim(
	ctx context.Context,
	claim rewardscoordinator.IRewardsCoordinatorRewardsMerkleClaim,
	earnerAddress gethcommon.Address,
) (*gethtypes.Receipt, error) {
	if w.rewardsCoordinator == nil {
		return nil, errors.New("RewardsCoordinator contract not provided")
	}

	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, utils.WrapError("failed to get no send tx opts", err)
	}

	tx, err := w.rewardsCoordinator.ProcessClaim(noSendTxOpts, claim, earnerAddress)
	if err != nil {
		return nil, utils.WrapError("failed to create ProcessClaim tx", err)
	}
	receipt, err := w.txMgr.Send(ctx, tx)
	if err != nil {
		return nil, utils.WrapError("failed to send tx", err)
	}

	return receipt, nil
}
