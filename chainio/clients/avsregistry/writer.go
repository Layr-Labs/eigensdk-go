package avsregistry

import (
	"context"
	"errors"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	"github.com/Layr-Labs/eigensdk-go/chainio/utils"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	gethcommon "github.com/ethereum/go-ethereum/common"
	gethtypes "github.com/ethereum/go-ethereum/core/types"

	blsapkregistry "github.com/Layr-Labs/eigensdk-go/contracts/bindings/BLSApkRegistry"
	opstateretriever "github.com/Layr-Labs/eigensdk-go/contracts/bindings/OperatorStateRetriever"
	regcoord "github.com/Layr-Labs/eigensdk-go/contracts/bindings/RegistryCoordinator"
	stakeregistry "github.com/Layr-Labs/eigensdk-go/contracts/bindings/StakeRegistry"
)

type AvsRegistryWriter interface {
	RegisterOperatorWithAVSRegistryCoordinator(
		ctx context.Context,
		blsKeyPair *bls.KeyPair,
		operatorAddr common.Address,
		quorumNumbers []byte,
		socket string,
	) (*gethtypes.Receipt, error)

	UpdateOperatorStakes(
		ctx context.Context,
		operators []gethcommon.Address,
	) (*gethtypes.Receipt, error)

	DeregisterOperator(
		ctx context.Context,
		quorumNumbers []byte,
		pubkey regcoord.BN254G1Point,
	) (*gethtypes.Receipt, error)
}

type AvsRegistryChainWriter struct {
	registryCoordinator    *regcoord.ContractRegistryCoordinator
	operatorStateRetriever *opstateretriever.ContractOperatorStateRetriever
	stakeRegistry          *stakeregistry.ContractStakeRegistry
	blsApkRegistry         *blsapkregistry.ContractBLSApkRegistry
	logger                 logging.Logger
	ethClient              eth.EthClient
	txMgr                  txmgr.TxManager
}

var _ AvsRegistryWriter = (*AvsRegistryChainWriter)(nil)

func NewAvsRegistryChainWriter(
	registryCoordinator *regcoord.ContractRegistryCoordinator,
	operatorStateRetriever *opstateretriever.ContractOperatorStateRetriever,
	stakeRegistry *stakeregistry.ContractStakeRegistry,
	blsApkRegistry *blsapkregistry.ContractBLSApkRegistry,
	logger logging.Logger,
	ethClient eth.EthClient,
	txMgr txmgr.TxManager,
) (*AvsRegistryChainWriter, error) {
	return &AvsRegistryChainWriter{
		registryCoordinator:    registryCoordinator,
		operatorStateRetriever: operatorStateRetriever,
		stakeRegistry:          stakeRegistry,
		blsApkRegistry:         blsApkRegistry,
		logger:                 logger,
		ethClient:              ethClient,
		txMgr:                  txMgr,
	}, nil
}

func BuildAvsRegistryChainWriter(
	registryCoordinatorAddr gethcommon.Address,
	operatorStateRetrieverAddr gethcommon.Address,
	logger logging.Logger,
	ethClient eth.EthClient,
	txMgr txmgr.TxManager,
) (*AvsRegistryChainWriter, error) {
	registryCoordinator, err := regcoord.NewContractRegistryCoordinator(registryCoordinatorAddr, ethClient)
	if err != nil {
		return nil, err
	}
	operatorStateRetriever, err := opstateretriever.NewContractOperatorStateRetriever(operatorStateRetrieverAddr, ethClient)
	if err != nil {
		return nil, err
	}
	blsApkRegistryAddr, err := registryCoordinator.BlsApkRegistry(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	blsApkRegistry, err := blsapkregistry.NewContractBLSApkRegistry(blsApkRegistryAddr, ethClient)
	if err != nil {
		return nil, err
	}
	stakeRegistryAddr, err := registryCoordinator.StakeRegistry(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	stakeRegistry, err := stakeregistry.NewContractStakeRegistry(stakeRegistryAddr, ethClient)
	if err != nil {
		return nil, err
	}
	return NewAvsRegistryChainWriter(
		registryCoordinator,
		operatorStateRetriever,
		stakeRegistry,
		blsApkRegistry,
		logger,
		ethClient,
		txMgr,
	)
}

// TODO(samlaf): clean up this function
func (w *AvsRegistryChainWriter) RegisterOperatorWithAVSRegistryCoordinator(
	ctx context.Context,
	blsKeyPair *bls.KeyPair,
	operatorAddr common.Address,
	quorumNumbers []byte,
	socket string,
) (*gethtypes.Receipt, error) {
	w.logger.Info("registering operator with the AVS's registry coordinator")
	g1HashedMsgToSign, err := w.registryCoordinator.PubkeyRegistrationMessageHash(&bind.CallOpts{}, operatorAddr)
	if err != nil {
		return nil, err
	}
	signedMsg := utils.ConvertToBN254G1Point(blsKeyPair.SignHashedToCurveMessage(utils.ConvertBn254GethToGnark(g1HashedMsgToSign)).G1Point)
	G1pubkeyBN254 := utils.ConvertToBN254G1Point(blsKeyPair.GetPubKeyG1())
	G2pubkeyBN254 := utils.ConvertToBN254G2Point(blsKeyPair.GetPubKeyG2())
	pubkeyRegParams := regcoord.IBLSApkRegistryPubkeyRegistrationParams{
		PubkeyRegistrationSignature: signedMsg,
		PubkeyG1:                    G1pubkeyBN254,
		PubkeyG2:                    G2pubkeyBN254,
	}
	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, err
	}
	operatorSignature := regcoord.ISignatureUtilsSignatureWithSaltAndExpiry{}
	// TODO: this call will fail if max number of operators are already registered
	// in that case, need to call churner to kick out another operator. See eigenDA's node/operator.go implementation
	tx, err := w.registryCoordinator.RegisterOperator(noSendTxOpts, quorumNumbers, socket, pubkeyRegParams, operatorSignature)
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

func (w *AvsRegistryChainWriter) UpdateOperatorStakes(
	ctx context.Context,
	operators []gethcommon.Address,
) (*gethtypes.Receipt, error) {
	w.logger.Info("updating stakes")
	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, err
	}
	tx, err := w.registryCoordinator.UpdateOperators(noSendTxOpts, operators)
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
	pubkey regcoord.BN254G1Point,
) (*gethtypes.Receipt, error) {
	w.logger.Info("deregistering operator with the AVS's registry coordinator")
	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, err
	}
	tx, err := w.registryCoordinator.DeregisterOperator(noSendTxOpts, quorumNumbers)
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
