package avsregistry

import (
	"context"
	"errors"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	"github.com/Layr-Labs/eigensdk-go/signerv2"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/logging"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	blsoperatorstateretriever "github.com/Layr-Labs/eigensdk-go/contracts/bindings/BLSOperatorStateRetriever"
	blspubkeyregistry "github.com/Layr-Labs/eigensdk-go/contracts/bindings/BLSPubkeyRegistry"
	blsregistrycoordinator "github.com/Layr-Labs/eigensdk-go/contracts/bindings/BLSRegistryCoordinatorWithIndices"
	stakeregistry "github.com/Layr-Labs/eigensdk-go/contracts/bindings/StakeRegistry"
)

type AvsRegistryWriter interface {
	RegisterOperatorWithAVSRegistryCoordinator(
		ctx context.Context,
		quorumNumbers []byte,
		pubkey blsregistrycoordinator.BN254G1Point,
		socket string,
	) (*types.Receipt, error)

	UpdateStakes(
		ctx context.Context,
		operators []gethcommon.Address,
	) (*types.Receipt, error)

	DeregisterOperator(
		ctx context.Context,
		quorumNumbers []byte,
		pubkey blsregistrycoordinator.BN254G1Point,
	) (*types.Receipt, error)
}

type AvsRegistryChainWriter struct {
	registryCoordinator       *blsregistrycoordinator.ContractBLSRegistryCoordinatorWithIndices
	blsOperatorStateRetriever *blsoperatorstateretriever.ContractBLSOperatorStateRetriever
	stakeRegistry             *stakeregistry.ContractStakeRegistry
	blsPubkeyRegistry         *blspubkeyregistry.ContractBLSPubkeyRegistry
	signer                    signerv2.SignerFn
	logger                    logging.Logger
	ethClient                 eth.EthClient
	txMgr                     txmgr.TxManager
}

var _ AvsRegistryWriter = (*AvsRegistryChainWriter)(nil)

func NewAvsRegistryWriter(
	registryCoordinator *blsregistrycoordinator.ContractBLSRegistryCoordinatorWithIndices,
	blsOperatorStateRetriever *blsoperatorstateretriever.ContractBLSOperatorStateRetriever,
	stakeRegistry *stakeregistry.ContractStakeRegistry,
	blsPubkeyRegistry *blspubkeyregistry.ContractBLSPubkeyRegistry,
	logger logging.Logger,
	ethClient eth.EthClient,
	txMgr txmgr.TxManager,
) (*AvsRegistryChainWriter, error) {
	return &AvsRegistryChainWriter{
		registryCoordinator:       registryCoordinator,
		blsOperatorStateRetriever: blsOperatorStateRetriever,
		stakeRegistry:             stakeRegistry,
		blsPubkeyRegistry:         blsPubkeyRegistry,
		logger:                    logger,
		ethClient:                 ethClient,
		txMgr:                     txMgr,
	}, nil
}

func (w *AvsRegistryChainWriter) RegisterOperatorWithAVSRegistryCoordinator(
	ctx context.Context,
	quorumNumbers []byte,
	pubkey blsregistrycoordinator.BN254G1Point,
	socket string,
) (*types.Receipt, error) {
	w.logger.Info("registering operator with the AVS's registry coordinator")
	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	// TODO: this call will fail if max number of operators are already registered
	// in that case, need to call churner to kick out another operator. See eigenDA's node/operator.go implementation
	tx, err := w.registryCoordinator.RegisterOperatorWithCoordinator1(noSendTxOpts, quorumNumbers, pubkey, socket)
	if err != nil {
		return nil, err
	}
	receipt, err := w.txMgr.Send(ctx, tx)
	if err != nil {
		return nil, errors.New("failed to send tx with err: " + err.Error())
	}
	w.logger.Infof("tx hash: %s", tx.Hash().String())
	w.logger.Info("registered operator with the AVS's registry coordinator")
	return receipt, nil
}

func (w *AvsRegistryChainWriter) UpdateStakes(
	ctx context.Context,
	operators []gethcommon.Address,
) (*types.Receipt, error) {
	w.logger.Info("updating stakes")
	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	tx, err := w.stakeRegistry.UpdateStakes(noSendTxOpts, operators)
	if err != nil {
		return nil, err
	}
	receipt, err := w.txMgr.Send(ctx, tx)
	if err != nil {
		return nil, errors.New("failed to send tx with err: " + err.Error())
	}
	w.logger.Infof("tx hash: %s", tx.Hash().String())
	w.logger.Info("updated stakes")
	return receipt, nil

}

func (w *AvsRegistryChainWriter) DeregisterOperator(
	ctx context.Context,
	quorumNumbers []byte,
	pubkey blsregistrycoordinator.BN254G1Point,
) (*types.Receipt, error) {
	w.logger.Info("deregistering operator with the AVS's registry coordinator")
	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	tx, err := w.registryCoordinator.DeregisterOperatorWithCoordinator(noSendTxOpts, quorumNumbers, pubkey)
	if err != nil {
		return nil, err
	}
	receipt, err := w.txMgr.Send(ctx, tx)
	if err != nil {
		return nil, errors.New("failed to send tx with err: " + err.Error())
	}
	w.logger.Infof("tx hash: %s", tx.Hash().String())
	w.logger.Info("deregistered operator with the AVS's registry coordinator")
	return receipt, nil
}
