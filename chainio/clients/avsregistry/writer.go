package avsregistry

import (
	"context"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/signer"
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
	signer                    signer.Signer
	logger                    logging.Logger
	ethClient                 eth.EthClient
}

var _ AvsRegistryWriter = (*AvsRegistryChainWriter)(nil)

func NewAvsRegistryWriter(
	registryCoordinator *blsregistrycoordinator.ContractBLSRegistryCoordinatorWithIndices,
	blsOperatorStateRetriever *blsoperatorstateretriever.ContractBLSOperatorStateRetriever,
	stakeRegistry *stakeregistry.ContractStakeRegistry,
	blsPubkeyRegistry *blspubkeyregistry.ContractBLSPubkeyRegistry,
	logger logging.Logger,
	signer signer.Signer,
	ethClient eth.EthClient,
) (*AvsRegistryChainWriter, error) {
	return &AvsRegistryChainWriter{
		registryCoordinator:       registryCoordinator,
		blsOperatorStateRetriever: blsOperatorStateRetriever,
		stakeRegistry:             stakeRegistry,
		blsPubkeyRegistry:         blsPubkeyRegistry,
		signer:                    signer,
		logger:                    logger,
		ethClient:                 ethClient,
	}, nil
}

func (w *AvsRegistryChainWriter) RegisterOperatorWithAVSRegistryCoordinator(
	ctx context.Context,
	quorumNumbers []byte,
	pubkey blsregistrycoordinator.BN254G1Point,
	socket string,
) (*types.Receipt, error) {
	w.logger.Info("registering operator with the AVS's registry coordinator")
	txOpts := w.signer.GetTxOpts()
	// TODO: this call will fail if max number of operators are already registered
	// in that case, need to call churner to kick out another operator. See eigenDA's node/operator.go implementation
	tx, err := w.registryCoordinator.RegisterOperatorWithCoordinator1(txOpts, quorumNumbers, pubkey, socket)
	if err != nil {
		return nil, err
	}
	w.logger.Infof("tx hash: %s", tx.Hash().String())
	receipt := w.ethClient.WaitForTransactionReceipt(ctx, tx.Hash())
	w.logger.Info("registered operator with the AVS's registry coordinator")
	return receipt, nil
}

func (w *AvsRegistryChainWriter) UpdateStakes(
	ctx context.Context,
	operators []gethcommon.Address,
) (*types.Receipt, error) {
	w.logger.Info("updating stakes")
	txOpts := w.signer.GetTxOpts()
	tx, err := w.stakeRegistry.UpdateStakes(txOpts, operators)
	if err != nil {
		return nil, err
	}
	w.logger.Infof("tx hash: %s", tx.Hash().String())
	receipt := w.ethClient.WaitForTransactionReceipt(ctx, tx.Hash())
	w.logger.Info("updated stakes")
	return receipt, nil

}

func (w *AvsRegistryChainWriter) DeregisterOperator(
	ctx context.Context,
	quorumNumbers []byte,
	pubkey blsregistrycoordinator.BN254G1Point,
) (*types.Receipt, error) {
	w.logger.Info("deregistering operator with the AVS's registry coordinator")
	txOpts := w.signer.GetTxOpts()
	tx, err := w.registryCoordinator.DeregisterOperatorWithCoordinator(txOpts, quorumNumbers, pubkey)
	if err != nil {
		return nil, err
	}
	w.logger.Infof("tx hash: %s", tx.Hash().String())
	receipt := w.ethClient.WaitForTransactionReceipt(ctx, tx.Hash())
	w.logger.Info("deregistered operator with the AVS's registry coordinator")
	return receipt, nil
}
