package elcontracts

import (
	"context"
	"errors"

	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethcommon "github.com/ethereum/go-ethereum/common"
	gethtypes "github.com/ethereum/go-ethereum/core/types"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	delegationmanager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/DelegationManager"
	avsdirectory "github.com/Layr-Labs/eigensdk-go/contracts/bindings/IAVSDirectory"
	allocationmanager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/IAllocationManager"
	erc20 "github.com/Layr-Labs/eigensdk-go/contracts/bindings/IERC20"
	rewardscoordinator "github.com/Layr-Labs/eigensdk-go/contracts/bindings/IRewardsCoordinator"
	strategy "github.com/Layr-Labs/eigensdk-go/contracts/bindings/IStrategy"
	strategymanager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/StrategyManager"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/metrics"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/Layr-Labs/eigensdk-go/utils"
)

type Reader interface {
	GetStrategyAndUnderlyingERC20Token(
		opts *bind.CallOpts, strategyAddr gethcommon.Address,
	) (*strategy.ContractIStrategy, erc20.ContractIERC20Methods, gethcommon.Address, error)
}

type ChainWriter struct {
	delegationManager   *delegationmanager.ContractDelegationManager
	strategyManager     *strategymanager.ContractStrategyManager
	rewardsCoordinator  *rewardscoordinator.ContractIRewardsCoordinator
	avsDirectory        *avsdirectory.ContractIAVSDirectory
	allocationManager   *allocationmanager.ContractIAllocationManager
	strategyManagerAddr gethcommon.Address
	elChainReader       Reader
	ethClient           eth.HttpBackend
	logger              logging.Logger
	txMgr               txmgr.TxManager
}

func NewChainWriter(
	delegationManager *delegationmanager.ContractDelegationManager,
	strategyManager *strategymanager.ContractStrategyManager,
	rewardsCoordinator *rewardscoordinator.ContractIRewardsCoordinator,
	avsDirectory *avsdirectory.ContractIAVSDirectory,
	allocationManager *allocationmanager.ContractIAllocationManager,
	strategyManagerAddr gethcommon.Address,
	elChainReader Reader,
	ethClient eth.HttpBackend,
	logger logging.Logger,
	eigenMetrics metrics.Metrics,
	txMgr txmgr.TxManager,
) *ChainWriter {
	logger = logger.With(logging.ComponentKey, "elcontracts/writer")

	return &ChainWriter{
		delegationManager:   delegationManager,
		strategyManager:     strategyManager,
		strategyManagerAddr: strategyManagerAddr,
		rewardsCoordinator:  rewardsCoordinator,
		allocationManager:   allocationManager,
		avsDirectory:        avsDirectory,
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
	ethClient eth.HttpBackend,
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
		elContractBindings.DelegationManager,
		elContractBindings.StrategyManager,
		elContractBindings.AvsDirectory,
		elContractBindings.RewardsCoordinator,
		elContractBindings.AllocationManager,
		logger,
		ethClient,
	)
	return NewChainWriter(
		elContractBindings.DelegationManager,
		elContractBindings.StrategyManager,
		elContractBindings.RewardsCoordinator,
		elContractBindings.AvsDirectory,
		elContractBindings.AllocationManager,
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
	ethClient eth.HttpBackend,
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
		elContractBindings.DelegationManager,
		elContractBindings.StrategyManager,
		elContractBindings.AvsDirectory,
		elContractBindings.RewardsCoordinator,
		elContractBindings.AllocationManager,
		logger,
		ethClient,
	)
	return NewChainWriter(
		elContractBindings.DelegationManager,
		elContractBindings.StrategyManager,
		elContractBindings.RewardsCoordinator,
		elContractBindings.AvsDirectory,
		elContractBindings.AllocationManager,
		elContractBindings.StrategyManagerAddr,
		elChainReader,
		ethClient,
		logger,
		eigenMetrics,
		txMgr,
	), nil
}

func (w *ChainWriter) RegisterAsOperator(
	ctx context.Context,
	operator types.Operator,
	waitForReceipt bool,
) (*gethtypes.Receipt, error) {
	if w.delegationManager == nil {
		return nil, errors.New("DelegationManager contract not provided")
	}

	w.logger.Infof("registering operator %s to EigenLayer", operator.Address)
	opDetails := delegationmanager.IDelegationManagerTypesOperatorDetails{
		// Earning receiver has been deprecated, so we just use the operator address as a dummy value
		// Any reward related setup is via RewardsCoordinator contract
		DeprecatedEarningsReceiver: gethcommon.HexToAddress(operator.Address),
		// DeprecatedStakerOptOutWindowBlocks has been deprecated, so we just use the operator's
		// StakerOptOutWindowBlocks
		DeprecatedStakerOptOutWindowBlocks: operator.StakerOptOutWindowBlocks,
		DelegationApprover:                 gethcommon.HexToAddress(operator.DelegationApproverAddress),
	}

	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, err
	}
	tx, err := w.delegationManager.RegisterAsOperator(
		noSendTxOpts,
		opDetails,
		operator.AllocationDelay,
		operator.MetadataUrl,
	)
	if err != nil {
		return nil, err
	}
	receipt, err := w.txMgr.Send(ctx, tx, waitForReceipt)
	if err != nil {
		return nil, errors.New("failed to send tx with err: " + err.Error())
	}
	w.logger.Info("tx successfully included", "txHash", receipt.TxHash.String())

	return receipt, nil
}

func (w *ChainWriter) UpdateOperatorDetails(
	ctx context.Context,
	operator types.Operator,
	waitForReceipt bool,
) (*gethtypes.Receipt, error) {
	if w.delegationManager == nil {
		return nil, errors.New("DelegationManager contract not provided")
	}

	w.logger.Infof("updating operator details of operator %s to EigenLayer", operator.Address)
	opDetails := delegationmanager.IDelegationManagerTypesOperatorDetails{
		// Earning receiver has been deprecated, so we just use the operator address as a dummy value
		// Any reward related setup is via RewardsCoordinator contract
		DeprecatedEarningsReceiver: gethcommon.HexToAddress(operator.Address),
		DelegationApprover:         gethcommon.HexToAddress(operator.DelegationApproverAddress),
		// DeprecatedStakerOptOutWindowBlocks has been deprecated, so we just use the operator's
		// StakerOptOutWindowBlocks
		DeprecatedStakerOptOutWindowBlocks: operator.StakerOptOutWindowBlocks,
	}

	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, err
	}

	tx, err := w.delegationManager.ModifyOperatorDetails(noSendTxOpts, opDetails)
	if err != nil {
		return nil, err
	}
	receipt, err := w.txMgr.Send(ctx, tx, waitForReceipt)
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

func (w *ChainWriter) UpdateMetadataURI(ctx context.Context, uri string, waitForReceipt bool,
) (*gethtypes.Receipt, error) {
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
	receipt, err := w.txMgr.Send(ctx, tx, waitForReceipt)
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
	waitForReceipt bool,
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
	_, err = w.txMgr.Send(ctx, tx, waitForReceipt)
	if err != nil {
		return nil, errors.New("failed to send tx with err: " + err.Error())
	}

	tx, err = w.strategyManager.DepositIntoStrategy(noSendTxOpts, strategyAddr, underlyingTokenAddr, amount)
	if err != nil {
		return nil, err
	}
	receipt, err := w.txMgr.Send(ctx, tx, waitForReceipt)
	if err != nil {
		return nil, errors.New("failed to send tx with err: " + err.Error())
	}

	w.logger.Infof("deposited %s into strategy %s", amount.String(), strategyAddr)
	return receipt, nil
}

func (w *ChainWriter) SetClaimerFor(
	ctx context.Context,
	claimer gethcommon.Address,
	waitForReceipt bool,
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
	receipt, err := w.txMgr.Send(ctx, tx, waitForReceipt)
	if err != nil {
		return nil, utils.WrapError("failed to send tx", err)
	}

	return receipt, nil
}

func (w *ChainWriter) ProcessClaim(
	ctx context.Context,
	claim rewardscoordinator.IRewardsCoordinatorTypesRewardsMerkleClaim,
	earnerAddress gethcommon.Address,
	waitForReceipt bool,
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
	receipt, err := w.txMgr.Send(ctx, tx, waitForReceipt)
	if err != nil {
		return nil, utils.WrapError("failed to send tx", err)
	}

	return receipt, nil
}

func (w *ChainWriter) ForceDeregisterFromOperatorSets(
	ctx context.Context,
	operator gethcommon.Address,
	avs gethcommon.Address,
	operatorSetIds []uint32,
	operatorSignature avsdirectory.ISignatureUtilsSignatureWithSaltAndExpiry,
	waitForReceipt bool,
) (*gethtypes.Receipt, error) {
	if w.avsDirectory == nil {
		return nil, errors.New("AVSDirectory contract not provided")
	}

	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, utils.WrapError("failed to get no send tx opts", err)
	}

	tx, err := w.avsDirectory.ForceDeregisterFromOperatorSets(
		noSendTxOpts,
		operator,
		avs,
		operatorSetIds,
		operatorSignature,
	)

	if err != nil {
		return nil, utils.WrapError("failed to create ForceDeregisterFromOperatorSets tx", err)
	}

	receipt, err := w.txMgr.Send(ctx, tx, waitForReceipt)
	if err != nil {
		return nil, utils.WrapError("failed to send tx", err)
	}

	return receipt, nil
}

//func (w *ChainWriter) SetOperatorCommissionBips(
//	ctx context.Context,
//	operatorSet rewardscoordinator.OperatorSet,
//	rewardType uint8,
//	commissionBips uint16,
//	waitForReceipt bool,
//) (*gethtypes.Receipt, error) {
//	if w.rewardsCoordinator == nil {
//		return nil, errors.New("RewardsCoordinator contract not provided")
//	}
//
//	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
//	if err != nil {
//		return nil, utils.WrapError("failed to get no send tx opts", err)
//	}
//
//	tx, err := w.rewardsCoordinator.SetOperatorCommissionBips(noSendTxOpts, operatorSet, rewardType, commissionBips)
//	if err != nil {
//		return nil, utils.WrapError("failed to create SetOperatorCommissionBips tx", err)
//	}
//
//	receipt, err := w.txMgr.Send(ctx, tx, waitForReceipt)
//	if err != nil {
//		return nil, utils.WrapError("failed to send tx", err)
//	}
//
//	return receipt, nil
//}

func (w *ChainWriter) ModifyAllocations(
	ctx context.Context,
	allocations []allocationmanager.IAllocationManagerTypesMagnitudeAllocation,
	waitForReceipt bool,
) (*gethtypes.Receipt, error) {
	if w.allocationManager == nil {
		return nil, errors.New("AllocationManager contract not provided")
	}

	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, utils.WrapError("failed to get no send tx opts", err)
	}

	tx, err := w.allocationManager.ModifyAllocations(noSendTxOpts, allocations)
	if err != nil {
		return nil, utils.WrapError("failed to create ModifyAllocations tx", err)
	}

	receipt, err := w.txMgr.Send(ctx, tx, waitForReceipt)
	if err != nil {
		return nil, utils.WrapError("failed to send tx", err)
	}

	return receipt, nil
}

func (w *ChainWriter) InitializeAllocationDelay(
	ctx context.Context,
	operatorAddress gethcommon.Address,
	delay uint32,
	waitForReceipt bool,
) (*gethtypes.Receipt, error) {
	if w.allocationManager == nil {
		return nil, errors.New("AllocationManager contract not provided")
	}

	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, utils.WrapError("failed to get no send tx opts", err)
	}

	tx, err := w.allocationManager.SetAllocationDelay(noSendTxOpts, operatorAddress, delay)
	if err != nil {
		return nil, utils.WrapError("failed to create InitializeAllocationDelay tx", err)
	}
	receipt, err := w.txMgr.Send(ctx, tx, waitForReceipt)
	if err != nil {
		return nil, utils.WrapError("failed to send tx", err)
	}

	return receipt, nil
}
