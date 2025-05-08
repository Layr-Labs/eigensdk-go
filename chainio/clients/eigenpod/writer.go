package eigenpod

import (
	"context"
	"math/big"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eigenpod/bindings"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/utils"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type ChainWriter struct {
	logger    logging.Logger
	ethClient eth.HttpBackend
	eigenPod  *bindings.IEigenPod
	txMgr     txmgr.TxManager
}

type ManagerChainWriter struct {
	logger    logging.Logger
	ethClient eth.HttpBackend
	manager   *bindings.IEigenPodManager
	txMgr     txmgr.TxManager
}

func newChainWriter(
	eigenPod *bindings.IEigenPod,
	ethClient eth.HttpBackend,
	logger logging.Logger,
	txMgr txmgr.TxManager,
) *ChainWriter {
	logger = logger.With(logging.ComponentKey, "eigenpod/writer")

	return &ChainWriter{
		logger:    logger,
		ethClient: ethClient,
		eigenPod:  eigenPod,
		txMgr:     txMgr,
	}
}

func newManagerChainWriter(
	manager *bindings.IEigenPodManager,
	ethClient eth.HttpBackend,
	logger logging.Logger,
	txMgr txmgr.TxManager,
) *ManagerChainWriter {
	logger = logger.With(logging.ComponentKey, "eigenpodmanager/writer")

	return &ManagerChainWriter{
		logger:    logger,
		ethClient: ethClient,
		manager:   manager,
		txMgr:     txMgr,
	}
}

func NewWriter(
	eigenPodAddress common.Address,
	ethClient eth.HttpBackend,
	txMgr txmgr.TxManager,
	logger logging.Logger,
) (*ChainWriter, error) {
	pod, err := bindings.NewIEigenPod(eigenPodAddress, ethClient)
	if err != nil {
		return nil, utils.WrapError("Failed to create EigenPod contract", err)
	}

	return newChainWriter(pod, ethClient, logger, txMgr), nil
}

func NewManagerWriter(
	eigenPodManagerAddress common.Address,
	ethClient eth.HttpBackend,
	txMgr txmgr.TxManager,
	logger logging.Logger,
) (*ManagerChainWriter, error) {
	manager, err := bindings.NewIEigenPodManager(eigenPodManagerAddress, ethClient)
	if err != nil {
		return nil, utils.WrapError("Failed to create EigenPodManager contract", err)
	}

	return newManagerChainWriter(manager, ethClient, logger, txMgr), nil
}

func (w *ChainWriter) VerifyWithdrawalCredentials(
	ctx context.Context,
	beaconTimestamp uint64,
	stateRootProof bindings.BeaconChainProofsStateRootProof,
	validatorIndices []*big.Int,
	validatorFieldsProofs [][]byte,
	validatorFields [][][32]byte,
) (*types.Receipt, error) {
	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, err
	}

	tx, err := w.eigenPod.VerifyWithdrawalCredentials(
		noSendTxOpts,
		beaconTimestamp,
		stateRootProof,
		validatorIndices,
		validatorFieldsProofs,
		validatorFields,
	)
	if err != nil {
		return nil, utils.WrapError("failed to create tx", err)
	}

	receipt, err := w.txMgr.Send(ctx, tx)
	if err != nil {
		return nil, utils.WrapError("failed to send tx", err)
	}

	return receipt, nil
}

func (w *ChainWriter) VerifyCheckpointProofs(
	ctx context.Context,
	balanceContainerProof bindings.BeaconChainProofsBalanceContainerProof,
	balanceProof []bindings.BeaconChainProofsBalanceProof,
) (*types.Receipt, error) {
	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, err
	}

	tx, err := w.eigenPod.VerifyCheckpointProofs(
		noSendTxOpts,
		balanceContainerProof,
		balanceProof,
	)
	if err != nil {
		return nil, utils.WrapError("failed to create tx", err)
	}

	receipt, err := w.txMgr.Send(ctx, tx)
	if err != nil {
		return nil, utils.WrapError("failed to send tx", err)
	}

	return receipt, nil
}

func (w *ChainWriter) StartCheckpoint(
	ctx context.Context,
	revertIfNoBalance bool,
) (*types.Receipt, error) {
	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, err
	}

	tx, err := w.eigenPod.StartCheckpoint(
		noSendTxOpts,
		revertIfNoBalance,
	)
	if err != nil {
		return nil, utils.WrapError("failed to create tx", err)
	}

	receipt, err := w.txMgr.Send(ctx, tx)
	if err != nil {
		return nil, utils.WrapError("failed to send tx", err)
	}

	return receipt, nil
}
