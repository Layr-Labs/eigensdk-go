package avsregistry

import (
	"context"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	blspubkeyregistry "github.com/Layr-Labs/eigensdk-go/contracts/bindings/BLSPubkeyRegistry"
	regcoord "github.com/Layr-Labs/eigensdk-go/contracts/bindings/BLSRegistryCoordinatorWithIndices"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/signer"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type AvsRegistryWriter interface {
	RegisterOperatorWithAVSRegistryCoordinator(
		ctx context.Context,
		quorumNumbers []byte,
		pubkey regcoord.BN254G1Point,
		socket string,
	) (*types.Receipt, error)

	UpdateStakes(
		ctx context.Context,
		operators []gethcommon.Address,
	) (*types.Receipt, error)

	DeregisterOperator(
		ctx context.Context,
		operator gethcommon.Address,
		quorumNumbers []byte,
		pubkey blspubkeyregistry.BN254G1Point,
	) (*types.Receipt, error)
}

type AvsRegistryChainWriter struct {
	avsRegistryContractsClient clients.AVSRegistryContractsClient
	signer                     signer.Signer
	logger                     logging.Logger
	ethClient                  eth.EthClient
}

var _ AvsRegistryWriter = (*AvsRegistryChainWriter)(nil)

func NewAvsRegistryWriter(
	avsRegistryContractsClient clients.AVSRegistryContractsClient,
	logger logging.Logger,
	signer signer.Signer,
	ethClient eth.EthClient,
) (*AvsRegistryChainWriter, error) {
	return &AvsRegistryChainWriter{
		avsRegistryContractsClient: avsRegistryContractsClient,
		signer:                     signer,
		logger:                     logger,
		ethClient:                  ethClient,
	}, nil
}

func (w *AvsRegistryChainWriter) RegisterOperatorWithAVSRegistryCoordinator(
	ctx context.Context,
	quorumNumbers []byte,
	pubkey regcoord.BN254G1Point,
	socket string,
) (*types.Receipt, error) {
	w.logger.Info("registering operator with the AVS's registry coordinator")
	txOpts := w.signer.GetTxOpts()
	tx, err := w.avsRegistryContractsClient.RegisterOperatorWithCoordinator(txOpts, quorumNumbers, pubkey, socket)
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
	tx, err := w.avsRegistryContractsClient.UpdateStakes(txOpts, operators)
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
	operator gethcommon.Address,
	quorumNumbers []byte,
	pubkey blspubkeyregistry.BN254G1Point,
) (*types.Receipt, error) {
	w.logger.Info("deregistering operator with the AVS's registry coordinator")
	txOpts := w.signer.GetTxOpts()
	tx, err := w.avsRegistryContractsClient.DeregisterOperator(txOpts, operator, quorumNumbers, pubkey)
	if err != nil {
		return nil, err
	}
	w.logger.Infof("tx hash: %s", tx.Hash().String())
	receipt := w.ethClient.WaitForTransactionReceipt(ctx, tx.Hash())
	w.logger.Info("deregistered operator with the AVS's registry coordinator")
	return receipt, nil
}
