package elcontracts

import (
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"errors"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"

	"math/big"

	gethcommon "github.com/ethereum/go-ethereum/common"
	gethtypes "github.com/ethereum/go-ethereum/core/types"

	m2delegationmanager "github.com/Layr-Labs/eigensdk-go/M2-contracts/bindings/DelegationManager"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	chainioutils "github.com/Layr-Labs/eigensdk-go/chainio/utils"
	avsdirectory "github.com/Layr-Labs/eigensdk-go/contracts/bindings/AVSDirectory"
	allocationmanager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/AllocationManager"
	delegationmanager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/DelegationManager"
	erc20 "github.com/Layr-Labs/eigensdk-go/contracts/bindings/IERC20"
	strategy "github.com/Layr-Labs/eigensdk-go/contracts/bindings/IStrategy"
	permissioncontroller "github.com/Layr-Labs/eigensdk-go/contracts/bindings/PermissionController"
	regcoord "github.com/Layr-Labs/eigensdk-go/contracts/bindings/RegistryCoordinator"
	rewardscoordinator "github.com/Layr-Labs/eigensdk-go/contracts/bindings/RewardsCoordinator"
	strategymanager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/StrategyManager"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/metrics"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/Layr-Labs/eigensdk-go/utils"
)

type Reader interface {
	GetStrategyAndUnderlyingERC20Token(
		ctx context.Context, strategyAddr gethcommon.Address,
	) (*strategy.ContractIStrategy, erc20.ContractIERC20Methods, gethcommon.Address, error)
}

// The ChainWriter provides methods to call the
// EigenLayer core contract's state-changing functions.
type ChainWriter struct {
	delegationManager    *delegationmanager.ContractDelegationManager
	strategyManager      *strategymanager.ContractStrategyManager
	rewardsCoordinator   *rewardscoordinator.ContractRewardsCoordinator
	avsDirectory         *avsdirectory.ContractAVSDirectory
	allocationManager    *allocationmanager.ContractAllocationManager
	permissionController *permissioncontroller.ContractPermissionController
	// This field exists to handle M2 contracts, so the address is used to create
	// the M2 delegationManager binding in registerAsOperatorPreSlashing.
	delegationManagerAddr gethcommon.Address
	strategyManagerAddr   gethcommon.Address
	elChainReader         Reader
	ethClient             eth.HttpBackend
	logger                logging.Logger
	txMgr                 txmgr.TxManager
}

// Returns a new instance of ChainWriter.
func NewChainWriter(
	delegationManager *delegationmanager.ContractDelegationManager,
	strategyManager *strategymanager.ContractStrategyManager,
	rewardsCoordinator *rewardscoordinator.ContractRewardsCoordinator,
	avsDirectory *avsdirectory.ContractAVSDirectory,
	allocationManager *allocationmanager.ContractAllocationManager,
	permissionController *permissioncontroller.ContractPermissionController,
	strategyManagerAddr gethcommon.Address,
	delegationManagerAddr gethcommon.Address,
	elChainReader Reader,
	ethClient eth.HttpBackend,
	logger logging.Logger,
	eigenMetrics metrics.Metrics,
	txMgr txmgr.TxManager,
) *ChainWriter {
	logger = logger.With(logging.ComponentKey, "elcontracts/writer")

	return &ChainWriter{
		delegationManager:     delegationManager,
		delegationManagerAddr: delegationManagerAddr,
		strategyManager:       strategyManager,
		strategyManagerAddr:   strategyManagerAddr,
		rewardsCoordinator:    rewardsCoordinator,
		allocationManager:     allocationManager,
		permissionController:  permissionController,
		avsDirectory:          avsDirectory,
		elChainReader:         elChainReader,
		logger:                logger,
		ethClient:             ethClient,
		txMgr:                 txMgr,
	}
}

// Returns a new instance of ChainWriter from a given config.
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
		elContractBindings.PermissionController,
		logger,
		ethClient,
	)
	return NewChainWriter(
		elContractBindings.DelegationManager,
		elContractBindings.StrategyManager,
		elContractBindings.RewardsCoordinator,
		elContractBindings.AvsDirectory,
		elContractBindings.AllocationManager,
		elContractBindings.PermissionController,
		elContractBindings.StrategyManagerAddr,
		elContractBindings.DelegationManagerAddr,
		elChainReader,
		ethClient,
		logger,
		eigenMetrics,
		txMgr,
	), nil
}

// Registers the caller as an operator in EigenLayer through the
// DelegationManager contract.
func (w *ChainWriter) RegisterAsOperator(
	ctx context.Context,
	operator types.Operator,
	waitForReceipt bool,
) (*gethtypes.Receipt, error) {
	if w.delegationManager == nil {
		return nil, errors.New("DelegationManager contract not provided")
	}

	w.logger.Infof("registering operator %s to EigenLayer", operator.Address)

	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, err
	}
	tx, err := w.delegationManager.RegisterAsOperator(
		noSendTxOpts,
		gethcommon.HexToAddress(operator.DelegationApproverAddress),
		operator.AllocationDelay,
		operator.MetadataUrl,
	)
	if err != nil {
		return nil, err
	}
	receipt, err := w.txMgr.Send(ctx, tx, waitForReceipt)
	if err != nil {
		return nil, utils.WrapError("failed to send tx", err)
	}
	w.logger.Info("tx successfully included", "txHash", receipt.TxHash.String())

	return receipt, nil
}

// Registers the caller as an operator in EigenLayer through the M2 DelegationManager contract.
// Note: This method is only works on the pre-slashing (M2) version of the contracts.
func (w *ChainWriter) RegisterAsOperatorPreSlashing(
	ctx context.Context,
	operator types.M2Operator,
	waitForReceipt bool,
) (*gethtypes.Receipt, error) {
	w.logger.Infof("registering operator %s to EigenLayer", operator.Address)

	m2DelegationManager, err := m2delegationmanager.NewContractDelegationManager(w.delegationManagerAddr, w.ethClient)
	if err != nil {
		return nil, utils.WrapError("Failed to fetch m2DelegationManager contract", err)
	}

	opDetails := m2delegationmanager.IDelegationManagerOperatorDetails{
		// Earning receiver has been deprecated, so we just use the operator address as a dummy value
		// Any reward related setup is via RewardsCoordinator contract
		DeprecatedEarningsReceiver: gethcommon.HexToAddress(operator.Address),
		StakerOptOutWindowBlocks:   operator.StakerOptOutWindowBlocks,
		DelegationApprover:         gethcommon.HexToAddress(operator.DelegationApproverAddress),
	}

	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, err
	}
	tx, err := m2DelegationManager.RegisterAsOperator(noSendTxOpts, opDetails, operator.MetadataUrl)
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

// Updates an operator's stored `delegationApprover` with
// the given `operator.DelegationApproverAddress` by calling
// the `modifyOperatorDetails` function in the DelegationManager contract.
func (w *ChainWriter) UpdateOperatorDetails(
	ctx context.Context,
	operator types.Operator,
	waitForReceipt bool,
) (*gethtypes.Receipt, error) {
	if w.delegationManager == nil {
		return nil, errors.New("DelegationManager contract not provided")
	}

	w.logger.Infof("updating operator details of operator %s to EigenLayer", operator.Address)

	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, err
	}

	tx, err := w.delegationManager.ModifyOperatorDetails(
		noSendTxOpts,
		gethcommon.HexToAddress(operator.Address),
		gethcommon.HexToAddress(operator.DelegationApproverAddress),
	)
	if err != nil {
		return nil, err
	}
	receipt, err := w.txMgr.Send(ctx, tx, waitForReceipt)
	if err != nil {
		return nil, utils.WrapError("failed to send tx", err)
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

// Updates the metadata URI for the given operator.
func (w *ChainWriter) UpdateMetadataURI(
	ctx context.Context,
	operatorAddress gethcommon.Address,
	uri string,
	waitForReceipt bool,
) (*gethtypes.Receipt, error) {
	if w.delegationManager == nil {
		return nil, errors.New("DelegationManager contract not provided")
	}

	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, err
	}

	tx, err := w.delegationManager.UpdateOperatorMetadataURI(noSendTxOpts, operatorAddress, uri)
	if err != nil {
		return nil, err
	}
	receipt, err := w.txMgr.Send(ctx, tx, waitForReceipt)
	if err != nil {
		return nil, utils.WrapError("failed to send tx", err)
	}
	w.logger.Info(
		"successfully updated operator metadata uri",
		"txHash",
		receipt.TxHash.String(),
	)

	return receipt, nil
}

// Deposits `amount` of the `strategyAddr` underlying token
// into the strategy given by `strategyAddr`.
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
		ctx,
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
		return nil, utils.WrapError("failed to send tx", err)
	}

	tx, err = w.strategyManager.DepositIntoStrategy(noSendTxOpts, strategyAddr, underlyingTokenAddr, amount)
	if err != nil {
		return nil, err
	}
	receipt, err := w.txMgr.Send(ctx, tx, waitForReceipt)
	if err != nil {
		return nil, utils.WrapError("failed to send tx", err)
	}

	w.logger.Infof("deposited %s into strategy %s", amount.String(), strategyAddr)
	return receipt, nil
}

// Sets `claimer` as the claimer for the earner (in this case the
// earner is the caller). That means that `claimer` can call `processClaim`
// on behalf of the earner.
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

// Processes the given `claim` for rewards.
// The rewards are transferred to the given `recipientAddress`.
func (w *ChainWriter) ProcessClaim(
	ctx context.Context,
	claim rewardscoordinator.IRewardsCoordinatorTypesRewardsMerkleClaim,
	recipientAddress gethcommon.Address,
	waitForReceipt bool,
) (*gethtypes.Receipt, error) {
	if w.rewardsCoordinator == nil {
		return nil, errors.New("RewardsCoordinator contract not provided")
	}

	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, utils.WrapError("failed to get no send tx opts", err)
	}

	tx, err := w.rewardsCoordinator.ProcessClaim(noSendTxOpts, claim, recipientAddress)
	if err != nil {
		return nil, utils.WrapError("failed to create ProcessClaim tx", err)
	}
	receipt, err := w.txMgr.Send(ctx, tx, waitForReceipt)
	if err != nil {
		return nil, utils.WrapError("failed to send tx", err)
	}

	return receipt, nil
}

// Sets the split for a specific operator for a specific AVS.
// The caller must be a registered operator.
// The split has to be between 0 and 10000 bips (inclusive).
// The split will be activated after activation delay.
func (w *ChainWriter) SetOperatorAVSSplit(
	ctx context.Context,
	operator gethcommon.Address,
	avs gethcommon.Address,
	split uint16,
	waitForReceipt bool,
) (*gethtypes.Receipt, error) {
	if w.rewardsCoordinator == nil {
		return nil, errors.New("RewardsCoordinator contract not provided")
	}

	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, utils.WrapError("failed to get no send tx opts", err)
	}

	tx, err := w.rewardsCoordinator.SetOperatorAVSSplit(noSendTxOpts, operator, avs, split)
	if err != nil {
		return nil, utils.WrapError("failed to create SetOperatorAVSSplit tx", err)
	}
	receipt, err := w.txMgr.Send(ctx, tx, waitForReceipt)
	if err != nil {
		return nil, utils.WrapError("failed to send tx", err)
	}

	return receipt, nil
}

// Sets the split for a specific operator for Programmatic Incentives.
// The caller must be a registered operator.
// The split has to be between 0 and 10000 bips (inclusive).
// The split will be activated after activation delay.
func (w *ChainWriter) SetOperatorPISplit(
	ctx context.Context,
	operator gethcommon.Address,
	split uint16,
	waitForReceipt bool,
) (*gethtypes.Receipt, error) {
	if w.rewardsCoordinator == nil {
		return nil, errors.New("RewardsCoordinator contract not provided")
	}

	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, utils.WrapError("failed to get no send tx opts", err)
	}

	tx, err := w.rewardsCoordinator.SetOperatorPISplit(noSendTxOpts, operator, split)
	if err != nil {
		return nil, utils.WrapError("failed to create SetOperatorAVSSplit tx", err)
	}
	receipt, err := w.txMgr.Send(ctx, tx, waitForReceipt)
	if err != nil {
		return nil, utils.WrapError("failed to send tx", err)
	}

	return receipt, nil
}

func (w *ChainWriter) SetOperatorSetSplit(
	ctx context.Context,
	operator gethcommon.Address,
	operatorSet rewardscoordinator.OperatorSet,
	split uint16,
	waitForReceipt bool,
) (*gethtypes.Receipt, error) {
	if w.rewardsCoordinator == nil {
		return nil, errors.New("RewardsCoordinator contract not provided")
	}

	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, utils.WrapError("failed to get no send tx opts", err)
	}

	tx, err := w.rewardsCoordinator.SetOperatorSetSplit(noSendTxOpts, operator, operatorSet, split)
	if err != nil {
		return nil, utils.WrapError("failed to create SetOperatorSetSplit tx", err)
	}
	receipt, err := w.txMgr.Send(ctx, tx, waitForReceipt)
	if err != nil {
		return nil, utils.WrapError("failed to send tx", err)
	}

	return receipt, nil
}

// Processes the claims given by `claims`.
// The rewards are transferred to the given `recipientAddress`.
func (w *ChainWriter) ProcessClaims(
	ctx context.Context,
	claims []rewardscoordinator.IRewardsCoordinatorTypesRewardsMerkleClaim,
	recipientAddress gethcommon.Address,
	waitForReceipt bool,
) (*gethtypes.Receipt, error) {
	if w.rewardsCoordinator == nil {
		return nil, errors.New("RewardsCoordinator contract not provided")
	}

	if len(claims) == 0 {
		return nil, errors.New("claims is empty, at least one claim must be provided")
	}

	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, utils.WrapError("failed to get no send tx opts", err)
	}

	tx, err := w.rewardsCoordinator.ProcessClaims(noSendTxOpts, claims, recipientAddress)
	if err != nil {
		return nil, utils.WrapError("failed to create ProcessClaims tx", err)
	}
	receipt, err := w.txMgr.Send(ctx, tx, waitForReceipt)
	if err != nil {
		return nil, utils.WrapError("failed to send tx", err)
	}

	return receipt, nil
}

// Modifies the proportions of slashable stake allocated to an operator set
// from a list of strategies, for the given `operatorAddress`.
func (w *ChainWriter) ModifyAllocations(
	ctx context.Context,
	operatorAddress gethcommon.Address,
	allocations []allocationmanager.IAllocationManagerTypesAllocateParams,
	waitForReceipt bool,
) (*gethtypes.Receipt, error) {
	if w.allocationManager == nil {
		return nil, errors.New("AllocationManager contract not provided")
	}

	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, utils.WrapError("failed to get no send tx opts", err)
	}

	tx, err := w.allocationManager.ModifyAllocations(noSendTxOpts, operatorAddress, allocations)
	if err != nil {
		return nil, utils.WrapError("failed to create ModifyAllocations tx", err)
	}

	receipt, err := w.txMgr.Send(ctx, tx, waitForReceipt)
	if err != nil {
		return nil, utils.WrapError("failed to send tx", err)
	}

	return receipt, nil
}

// Receives an operator address, and a list of strategies and numsToClear (number of elements to clear from queue),
// and clears the operators deallocation queue in numbers to clear for the given strategies, by completing the
// pending deallocations if their effect timestamps have passed. Note that strategies and numsToClear should have
// equal length, since there should be a number of elements to clear from queue for each strategy queue.
func (w *ChainWriter) ClearDeallocationQueue(
	ctx context.Context,
	operatorAddress gethcommon.Address,
	strategies []gethcommon.Address,
	numsToClear []uint16,
	waitForReceipt bool,
) (*gethtypes.Receipt, error) {
	if w.allocationManager == nil {
		return nil, errors.New("AllocationManager contract not provided")
	}

	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, utils.WrapError("failed to get no send tx opts", err)
	}

	tx, err := w.allocationManager.ClearDeallocationQueue(noSendTxOpts, operatorAddress, strategies, numsToClear)
	if err != nil {
		return nil, utils.WrapError("failed to create ClearDeallocationQueue tx", err)
	}

	receipt, err := w.txMgr.Send(ctx, tx, waitForReceipt)
	if err != nil {
		return nil, utils.WrapError("failed to send tx", err)
	}

	return receipt, nil
}

// Sets the allocation delay for an operator.
// The allocation delay is the number of blocks between the operator
// allocating a magnitude to an operator set, and the magnitude becoming
// slashable.
func (w *ChainWriter) SetAllocationDelay(
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

// Sets the avsRegistrar for the received avs.
// The avsRegistrar is the contract called to complete registration,
// usually a RegistryCoordinator
func (w *ChainWriter) SetAVSRegistrar(
	ctx context.Context,
	avs gethcommon.Address,
	registrar gethcommon.Address,
	waitForReceipt bool,
) (*gethtypes.Receipt, error) {
	if w.allocationManager == nil {
		return nil, errors.New("AllocationManager contract not provided")
	}

	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, utils.WrapError("failed to get no send tx opts", err)
	}

	tx, err := w.allocationManager.SetAVSRegistrar(noSendTxOpts, avs, registrar)
	if err != nil {
		return nil, utils.WrapError("failed to create SetAVSRegistrar tx", err)
	}
	receipt, err := w.txMgr.Send(ctx, tx, waitForReceipt)
	if err != nil {
		return nil, utils.WrapError("failed to send tx", err)
	}

	return receipt, nil
}

// Deregister an operator from one or more of the AVS's operator sets.
// If the operator has any slashable stake allocated to the AVS,
// it remains slashable until the deallocation delay has passed.
func (w *ChainWriter) DeregisterFromOperatorSets(
	ctx context.Context,
	operator gethcommon.Address,
	request DeregistrationRequest,
) (*gethtypes.Receipt, error) {
	if w.allocationManager == nil {
		return nil, errors.New("AllocationManager contract not provided")
	}

	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, utils.WrapError("failed to get no send tx opts", err)
	}

	tx, err := w.allocationManager.DeregisterFromOperatorSets(
		noSendTxOpts,
		allocationmanager.IAllocationManagerTypesDeregisterParams{
			Operator:       operator,
			Avs:            request.AVSAddress,
			OperatorSetIds: request.OperatorSetIds,
		})
	if err != nil {
		return nil, utils.WrapError("failed to create DeregisterFromOperatorSets tx", err)
	}

	receipt, err := w.txMgr.Send(ctx, tx, request.WaitForReceipt)
	if err != nil {
		return nil, utils.WrapError("failed to send tx", err)
	}

	return receipt, nil
}

// Register an operator for one or more operator sets for an AVS.
// If `churnApprovalEcdsaPrivateKey` is provided, the churn approver
// parameters will be used for replacing existing operators in full quorums.
// If the operator has any stake allocated to these operator sets,
// it immediately becomes slashable.
func (w *ChainWriter) RegisterForOperatorSets(
	ctx context.Context,
	registryCoordinatorAddr gethcommon.Address,
	request RegistrationRequest,
) (*gethtypes.Receipt, error) {
	if w.allocationManager == nil {
		return nil, errors.New("AllocationManager contract not provided")
	}

	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, utils.WrapError("failed to get no send tx opts", err)
	}

	pubkeyRegParams, err := getPubkeyRegistrationParams(
		w.ethClient,
		registryCoordinatorAddr,
		request.OperatorAddress,
		request.BlsKeyPair,
	)
	if err != nil {
		return nil, utils.WrapError("failed to get public key registration params", err)
	}

	var data []byte

	if request.ChurnApprovalEcdsaPrivateKey != nil {
		// it's a registration with churn approval
		signatureWithSaltAndExpiry, err := signChurnRegistration(
			ctx,
			w.ethClient,
			registryCoordinatorAddr,
			request.OperatorAddress,
			request.ChurnApprovalEcdsaPrivateKey,
			request.BlsKeyPair.GetPubKeyG1(),
			request.OperatorKickParams,
		)
		if err != nil {
			return nil, utils.WrapError("failed to compute churn registration signature", err)
		}
		data, err = AbiEncodeRegistrationWithChurnParams(
			request.Socket,
			*pubkeyRegParams,
			request.OperatorKickParams,
			signatureWithSaltAndExpiry,
		)
		if err != nil {
			return nil, utils.WrapError("failed to encode churn registration params", err)
		}
	} else {
		// it's a normal registration
		data, err = AbiEncodeNormalRegistrationParams(request.Socket, *pubkeyRegParams)
		if err != nil {
			return nil, utils.WrapError("failed to encode normal registration params", err)
		}
	}
	tx, err := w.allocationManager.RegisterForOperatorSets(
		noSendTxOpts,
		request.OperatorAddress,
		allocationmanager.IAllocationManagerTypesRegisterParams{
			Avs:            request.AVSAddress,
			OperatorSetIds: request.OperatorSetIds,
			Data:           data,
		})
	if err != nil {
		return nil, utils.WrapError("failed to create RegisterForOperatorSets tx", err)
	}

	receipt, err := w.txMgr.Send(ctx, tx, request.WaitForReceipt)
	if err != nil {
		return nil, utils.WrapError("failed to send tx", err)
	}

	return receipt, nil
}

// Removes permission of an appointee for a specific function
// (given by request.selector) on a target contract, given an account address.
func (w *ChainWriter) RemovePermission(
	ctx context.Context,
	request RemovePermissionRequest,
) (*gethtypes.Receipt, error) {
	txOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, utils.WrapError("failed to get no-send tx opts", err)
	}
	tx, err := w.NewRemovePermissionTx(txOpts, request)
	if err != nil {
		return nil, utils.WrapError("failed to create NewRemovePermissionTx", err)
	}
	return w.txMgr.Send(ctx, tx, request.WaitForReceipt)
}

// Builds a transaction for the PermissionController's removeAppointee function.
func (w *ChainWriter) NewRemovePermissionTx(
	txOpts *bind.TransactOpts,
	request RemovePermissionRequest,
) (*gethtypes.Transaction, error) {
	if w.permissionController == nil {
		return nil, errors.New("permission contract not provided")
	}

	return w.permissionController.RemoveAppointee(
		txOpts,
		request.AccountAddress,
		request.AppointeeAddress,
		request.Target,
		request.Selector,
	)
}

// Builds a transaction for the PermissionController's setAppointee function.
func (w *ChainWriter) NewSetPermissionTx(
	txOpts *bind.TransactOpts,
	request SetPermissionRequest,
) (*gethtypes.Transaction, error) {
	if w.permissionController == nil {
		return nil, errors.New("permission contract not provided")
	}
	return w.permissionController.SetAppointee(
		txOpts,
		request.AccountAddress,
		request.AppointeeAddress,
		request.Target,
		request.Selector,
	)
}

// Set an appointee for a given account.
// Only the admin of the account can set an appointee.
// The appointee will be able to call the function given
// by `request.Selector` on the contract given by `request.Target`.
func (w *ChainWriter) SetPermission(
	ctx context.Context,
	request SetPermissionRequest,
) (*gethtypes.Receipt, error) {
	txOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, utils.WrapError("failed to get no-send tx opts", err)
	}

	tx, err := w.NewSetPermissionTx(txOpts, request)
	if err != nil {
		return nil, utils.WrapError("failed to create NewSetPermissionTx", err)
	}

	return w.txMgr.Send(ctx, tx, request.WaitForReceipt)
}

// Builds a transaction for the PermissionController's acceptAdmin function.
func (w *ChainWriter) NewAcceptAdminTx(
	txOpts *bind.TransactOpts,
	request AcceptAdminRequest,
) (*gethtypes.Transaction, error) {
	if w.permissionController == nil {
		return nil, errors.New("permission contract not provided")
	}
	return w.permissionController.AcceptAdmin(txOpts, request.AccountAddress)
}

// Accept a pending admin for the account given by `request.AccountAddress`.
// The sender of the transaction must be the pending admin.
func (w *ChainWriter) AcceptAdmin(
	ctx context.Context,
	request AcceptAdminRequest,
) (*gethtypes.Receipt, error) {
	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, utils.WrapError("failed to get no send tx opts", err)
	}

	tx, err := w.NewAcceptAdminTx(noSendTxOpts, request)
	if err != nil {
		return nil, utils.WrapError("failed to create AcceptAdmin transaction", err)
	}
	return w.txMgr.Send(ctx, tx, request.WaitForReceipt)
}

// Builds a transaction for the PermissionController's addPendingAdmin function.
func (w *ChainWriter) NewAddPendingAdminTx(
	txOpts *bind.TransactOpts,
	request AddPendingAdminRequest,
) (*gethtypes.Transaction, error) {
	if w.permissionController == nil {
		return nil, errors.New("permission contract not provided")
	}
	return w.permissionController.AddPendingAdmin(txOpts, request.AccountAddress, request.AdminAddress)
}

// Set a pending admin. Multiple admins can be set for an account.
// The caller must be an admin. If the account does not have an admin,
// the caller must be the account.
func (w *ChainWriter) AddPendingAdmin(ctx context.Context, request AddPendingAdminRequest) (*gethtypes.Receipt, error) {
	txOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, utils.WrapError("failed to get no send tx opts", err)
	}
	tx, err := w.NewAddPendingAdminTx(txOpts, request)
	if err != nil {
		return nil, utils.WrapError("failed to create AddPendingAdminTx", err)
	}
	return w.txMgr.Send(ctx, tx, request.WaitForReceipt)
}

// Builds a transaction for the PermissionController's removeAdmin function.
func (w *ChainWriter) NewRemoveAdminTx(
	txOpts *bind.TransactOpts,
	request RemoveAdminRequest,
) (*gethtypes.Transaction, error) {
	if w.permissionController == nil {
		return nil, errors.New("permission contract not provided")
	}
	return w.permissionController.RemoveAdmin(txOpts, request.AccountAddress, request.AdminAddress)
}

// Removes the admin given by `request.AdminAddress` from the account given
// by `request.AccountAddress`. The sender of the transaction must be an admin.
func (w *ChainWriter) RemoveAdmin(
	ctx context.Context,
	request RemoveAdminRequest,
) (*gethtypes.Receipt, error) {
	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, utils.WrapError("failed to get no send tx opts", err)
	}

	tx, err := w.NewRemoveAdminTx(noSendTxOpts, request)
	if err != nil {
		return nil, utils.WrapError("failed to create RemoveAdmin transaction", err)
	}
	return w.txMgr.Send(ctx, tx, request.WaitForReceipt)
}

// Builds a transaction for the PermissionController's removePendingAdmin function.
func (w *ChainWriter) NewRemovePendingAdminTx(
	txOpts *bind.TransactOpts,
	request RemovePendingAdminRequest,
) (*gethtypes.Transaction, error) {
	if w.permissionController == nil {
		return nil, errors.New("permission contract not provided")
	}
	return w.permissionController.RemovePendingAdmin(txOpts, request.AccountAddress, request.AdminAddress)
}

// Remove pending admin given by `request.AdminAddress` from the account given
// by `request.AccountAddress`. Only the admin of the account can remove a pending admin.
func (w *ChainWriter) RemovePendingAdmin(
	ctx context.Context,
	request RemovePendingAdminRequest,
) (*gethtypes.Receipt, error) {
	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, utils.WrapError("failed to get no send tx opts", err)
	}

	tx, err := w.NewRemovePendingAdminTx(noSendTxOpts, request)
	if err != nil {
		return nil, utils.WrapError("failed to create RemovePendingAdmin transaction", err)
	}

	return w.txMgr.Send(ctx, tx, request.WaitForReceipt)
}

// Returns the pubkey registration params for the operator given by `operatorAddress`.
func getPubkeyRegistrationParams(
	ethClient bind.ContractBackend,
	registryCoordinatorAddr, operatorAddress gethcommon.Address,
	blsKeyPair *bls.KeyPair,
) (*regcoord.IBLSApkRegistryTypesPubkeyRegistrationParams, error) {
	registryCoordinator, err := regcoord.NewContractRegistryCoordinator(registryCoordinatorAddr, ethClient)
	if err != nil {
		return nil, utils.WrapError("failed to create registry coordinator", err)
	}
	// params to register bls pubkey with bls apk registry
	g1HashedMsgToSign, err := registryCoordinator.PubkeyRegistrationMessageHash(
		&bind.CallOpts{},
		operatorAddress,
	)
	if err != nil {
		return nil, err
	}
	signedMsg := chainioutils.ConvertToBN254G1Point(
		blsKeyPair.SignHashedToCurveMessage(chainioutils.ConvertBn254GethToGnark(g1HashedMsgToSign)).G1Point,
	)
	G1pubkeyBN254 := chainioutils.ConvertToBN254G1Point(blsKeyPair.GetPubKeyG1())
	G2pubkeyBN254 := chainioutils.ConvertToBN254G2Point(blsKeyPair.GetPubKeyG2())
	pubkeyRegParams := regcoord.IBLSApkRegistryTypesPubkeyRegistrationParams{
		PubkeyRegistrationSignature: signedMsg,
		PubkeyG1:                    G1pubkeyBN254,
		PubkeyG2:                    G2pubkeyBN254,
	}
	return &pubkeyRegParams, nil
}

// Returns the pubkey registration params for the operator given by `operatorAddress`.
func signChurnRegistration(
	ctx context.Context,
	ethClient eth.HttpBackend,
	registryCoordinatorAddr, operatorAddress gethcommon.Address,
	churnApprovalEcdsaPrivateKey *ecdsa.PrivateKey,
	pubKeyG1 *bls.G1Point,
	operatorKickParams []OperatorKickParam,
) (SignatureWithSaltAndExpiry, error) {
	curBlockNum, err := ethClient.BlockNumber(context.Background())
	if err != nil {
		return SignatureWithSaltAndExpiry{}, err
	}
	curBlock, err := ethClient.BlockByNumber(context.Background(), new(big.Int).SetUint64(curBlockNum))
	if err != nil {
		return SignatureWithSaltAndExpiry{}, err
	}

	sigValidForSeconds := int64(60 * 60) // 1 hour

	curTime := new(big.Int).SetUint64(curBlock.Time())
	expiry := new(big.Int).Add(curTime, big.NewInt(sigValidForSeconds))

	var salt [32]byte
	_, err = rand.Read(salt[:])
	if err != nil {
		return SignatureWithSaltAndExpiry{}, err
	}
	registryCoordinator, err := regcoord.NewContractRegistryCoordinator(registryCoordinatorAddr, ethClient)
	if err != nil {
		return SignatureWithSaltAndExpiry{}, utils.WrapError("failed to create registry coordinator", err)
	}
	operatorId := types.OperatorIdFromG1Pubkey(pubKeyG1)
	kickParams := make([]regcoord.ISlashingRegistryCoordinatorTypesOperatorKickParam, len(operatorKickParams))
	for i, kickParam := range operatorKickParams {
		kickParams[i] = regcoord.ISlashingRegistryCoordinatorTypesOperatorKickParam{
			QuorumNumber: kickParam.QuorumNumber,
			Operator:     kickParam.Operator,
		}
	}
	msgToSign, err := registryCoordinator.CalculateOperatorChurnApprovalDigestHash(
		&bind.CallOpts{Context: ctx},
		operatorAddress,
		operatorId,
		kickParams,
		salt,
		expiry,
	)
	if err != nil {
		return SignatureWithSaltAndExpiry{}, err
	}
	signature, err := crypto.Sign(msgToSign[:], churnApprovalEcdsaPrivateKey)
	if err != nil {
		return SignatureWithSaltAndExpiry{}, err
	}
	// the crypto library is low level and deals with 0/1 v values, whereas ethereum expects 27/28, so we add 27
	// see https://github.com/ethereum/go-ethereum/issues/28757#issuecomment-1874525854
	// and https://twitter.com/pcaversaccio/status/1671488928262529031
	signature[64] += 27

	signatureWithSaltAndExpiry := SignatureWithSaltAndExpiry{
		Signature: signature,
		Salt:      salt,
		Expiry:    expiry,
	}
	return signatureWithSaltAndExpiry, nil
}

func getNormalRegistrationAbi() []abi.ArgumentMarshaling {
	return []abi.ArgumentMarshaling{
		{Name: "RegistrationType", Type: "uint8"},
		{Name: "Socket", Type: "string"},
		{Name: "PubkeyRegParams", Type: "tuple", Components: []abi.ArgumentMarshaling{
			{Name: "PubkeyRegistrationSignature", Type: "tuple", Components: []abi.ArgumentMarshaling{
				{Name: "X", Type: "uint256"},
				{Name: "Y", Type: "uint256"},
			}},
			{Name: "PubkeyG1", Type: "tuple", Components: []abi.ArgumentMarshaling{
				{Name: "X", Type: "uint256"},
				{Name: "Y", Type: "uint256"},
			}},
			{Name: "PubkeyG2", Type: "tuple", Components: []abi.ArgumentMarshaling{
				{Name: "X", Type: "uint256[2]"},
				{Name: "Y", Type: "uint256[2]"},
			}},
		}},
	}
}

func getRegistrationWithChurnAbi() []abi.ArgumentMarshaling {
	return append(getNormalRegistrationAbi(), []abi.ArgumentMarshaling{
		{Name: "OperatorKickParams", Type: "tuple[]", Components: []abi.ArgumentMarshaling{
			{Name: "QuorumNumber", Type: "uint8"},
			{Name: "Operator", Type: "address"},
		}},
		{Name: "SignatureWithSaltAndExpiry", Type: "tuple", Components: []abi.ArgumentMarshaling{
			{Name: "Signature", Type: "bytes"},
			{Name: "Salt", Type: "bytes32"},
			{Name: "Expiry", Type: "uint256"},
		}},
	}...)
}

// Returns the ABI encoding of the given normal registration params.
func AbiEncodeNormalRegistrationParams(
	socket string,
	pubkeyRegistrationParams regcoord.IBLSApkRegistryTypesPubkeyRegistrationParams,
) ([]byte, error) {
	registrationParamsType, err := abi.NewType("tuple", "", getNormalRegistrationAbi())
	if err != nil {
		return nil, err
	}

	registrationParams := struct {
		RegistrationType RegistrationType
		Socket           string
		PubkeyRegParams  regcoord.IBLSApkRegistryTypesPubkeyRegistrationParams
	}{
		RegistrationTypeNormal,
		socket,
		pubkeyRegistrationParams,
	}

	args := abi.Arguments{
		{Type: registrationParamsType, Name: "registrationParams"},
	}

	data, err := args.Pack(&registrationParams)
	if err != nil {
		return nil, err
	}
	// The encoder is prepending 32 bytes to the data as if it was used in a dynamic function parameter.
	// This is not used when decoding the bytes directly, so we need to remove it.
	return data[32:], nil
}

// Returns the ABI encoding of the given registration with churn params.
func AbiEncodeRegistrationWithChurnParams(
	socket string,
	pubkeyRegistrationParams regcoord.IBLSApkRegistryTypesPubkeyRegistrationParams,
	operatorKickParams []OperatorKickParam,
	churnApproverSignature SignatureWithSaltAndExpiry,
) ([]byte, error) {
	registrationParamsType, err := abi.NewType("tuple", "", getRegistrationWithChurnAbi())
	if err != nil {
		return nil, err
	}

	registrationParams := struct {
		RegistrationType           RegistrationType
		Socket                     string
		PubkeyRegParams            regcoord.IBLSApkRegistryTypesPubkeyRegistrationParams
		OperatorKickParams         []OperatorKickParam
		SignatureWithSaltAndExpiry SignatureWithSaltAndExpiry
	}{
		RegistrationTypeChurn,
		socket,
		pubkeyRegistrationParams,
		operatorKickParams,
		churnApproverSignature,
	}

	args := abi.Arguments{
		{Type: registrationParamsType, Name: "registrationParams"},
	}

	data, err := args.Pack(&registrationParams)
	if err != nil {
		return nil, err
	}
	// The encoder is prepending 32 bytes to the data as if it was used in a dynamic function parameter.
	// This is not used when decoding the bytes directly, so we need to remove it.
	return data[32:], nil
}
