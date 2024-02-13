package avsregistry

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"math/big"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/elcontracts"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	"github.com/Layr-Labs/eigensdk-go/chainio/utils"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethcommon "github.com/ethereum/go-ethereum/common"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	blsapkregistry "github.com/Layr-Labs/eigensdk-go/contracts/bindings/BLSApkRegistry"
	opstateretriever "github.com/Layr-Labs/eigensdk-go/contracts/bindings/OperatorStateRetriever"
	regcoord "github.com/Layr-Labs/eigensdk-go/contracts/bindings/RegistryCoordinator"
	smbase "github.com/Layr-Labs/eigensdk-go/contracts/bindings/ServiceManagerBase"
	stakeregistry "github.com/Layr-Labs/eigensdk-go/contracts/bindings/StakeRegistry"
)

type AvsRegistryWriter interface {
	// TODO(samlaf): an operator that is already registered in a quorum can register with another quorum without passing
	// signatures perhaps we should add another sdk function for this purpose, that just takes in a quorumNumber and
	// socket? RegisterOperatorInQuorumWithAVSRegistryCoordinator is used to register a single operator with the AVS's
	// registry coordinator. - operatorEcdsaPrivateKey is the operator's ecdsa private key (used to sign a message to
	// register operator in eigenlayer's delegation manager)
	//  - operatorToAvsRegistrationSigSalt is a random salt used to prevent replay attacks
	//  - operatorToAvsRegistrationSigExpiry is the expiry time of the signature
	RegisterOperatorInQuorumWithAVSRegistryCoordinator(
		ctx context.Context,
		operatorEcdsaPrivateKey *ecdsa.PrivateKey,
		operatorToAvsRegistrationSigSalt [32]byte,
		operatorToAvsRegistrationSigExpiry *big.Int,
		blsKeyPair *bls.KeyPair,
		quorumNumbers []byte,
		socket string,
	) (*gethtypes.Receipt, error)

	// UpdateStakesOfEntireOperatorSetForQuorums is used by avs teams running https://github.com/Layr-Labs/avs-sync
	// to updates the stake of their entire operator set.
	// Because of high gas costs of this operation, it typically needs to be called for every quorum, or perhaps for a
	// small grouping of quorums
	// (highly dependent on number of operators per quorum)
	UpdateStakesOfEntireOperatorSetForQuorums(
		ctx context.Context,
		operatorsPerQuorum [][]gethcommon.Address,
		quorumNumbers []byte,
	) (*gethtypes.Receipt, error)

	// UpdateStakesOfOperatorSubsetForAllQuorums is meant to be used by single operators (or teams of operators,
	// possibly running https://github.com/Layr-Labs/avs-sync) to update the stake of their own operator(s). This might
	// be needed in the case that they received a lot of new stake delegations, and want this to be reflected
	// in the AVS's registry coordinator.
	UpdateStakesOfOperatorSubsetForAllQuorums(
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
	serviceManagerAddr     gethcommon.Address
	registryCoordinator    *regcoord.ContractRegistryCoordinator
	operatorStateRetriever *opstateretriever.ContractOperatorStateRetriever
	stakeRegistry          *stakeregistry.ContractStakeRegistry
	blsApkRegistry         *blsapkregistry.ContractBLSApkRegistry
	elReader               elcontracts.ELReader
	logger                 logging.Logger
	ethClient              eth.EthClient
	txMgr                  txmgr.TxManager
}

var _ AvsRegistryWriter = (*AvsRegistryChainWriter)(nil)

func NewAvsRegistryChainWriter(
	serviceManagerAddr gethcommon.Address,
	registryCoordinator *regcoord.ContractRegistryCoordinator,
	operatorStateRetriever *opstateretriever.ContractOperatorStateRetriever,
	stakeRegistry *stakeregistry.ContractStakeRegistry,
	blsApkRegistry *blsapkregistry.ContractBLSApkRegistry,
	elReader elcontracts.ELReader,
	logger logging.Logger,
	ethClient eth.EthClient,
	txMgr txmgr.TxManager,
) (*AvsRegistryChainWriter, error) {
	return &AvsRegistryChainWriter{
		serviceManagerAddr:     serviceManagerAddr,
		registryCoordinator:    registryCoordinator,
		operatorStateRetriever: operatorStateRetriever,
		stakeRegistry:          stakeRegistry,
		blsApkRegistry:         blsApkRegistry,
		elReader:               elReader,
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
		return nil, errors.Join(errors.New("Failed to create RegistryCoordinator contract"), err)
	}
	operatorStateRetriever, err := opstateretriever.NewContractOperatorStateRetriever(
		operatorStateRetrieverAddr,
		ethClient,
	)
	if err != nil {
		return nil, errors.Join(errors.New("Failed to create OperatorStateRetriever contract"), err)
	}
	serviceManagerAddr, err := registryCoordinator.ServiceManager(&bind.CallOpts{})
	if err != nil {
		return nil, errors.Join(errors.New("Failed to get ServiceManager address"), err)
	}
	serviceManager, err := smbase.NewContractServiceManagerBase(serviceManagerAddr, ethClient)
	if err != nil {
		return nil, errors.Join(errors.New("Failed to create ServiceManager contract"), err)
	}
	blsApkRegistryAddr, err := registryCoordinator.BlsApkRegistry(&bind.CallOpts{})
	if err != nil {
		return nil, errors.Join(errors.New("Failed to get BLSApkRegistry address"), err)
	}
	blsApkRegistry, err := blsapkregistry.NewContractBLSApkRegistry(blsApkRegistryAddr, ethClient)
	if err != nil {
		return nil, errors.Join(errors.New("Failed to create BLSApkRegistry contract"), err)
	}
	stakeRegistryAddr, err := registryCoordinator.StakeRegistry(&bind.CallOpts{})
	if err != nil {
		return nil, errors.Join(errors.New("Failed to get StakeRegistry address"), err)
	}
	stakeRegistry, err := stakeregistry.NewContractStakeRegistry(stakeRegistryAddr, ethClient)
	if err != nil {
		return nil, errors.Join(errors.New("Failed to create StakeRegistry contract"), err)
	}
	delegationManagerAddr, err := stakeRegistry.Delegation(&bind.CallOpts{})
	if err != nil {
		return nil, errors.Join(errors.New("Failed to get DelegationManager address"), err)
	}
	avsDirectoryAddr, err := serviceManager.AvsDirectory(&bind.CallOpts{})
	if err != nil {
		return nil, errors.Join(errors.New("Failed to get AvsDirectory address"), err)
	}
	elReader, err := elcontracts.BuildELChainReader(delegationManagerAddr, avsDirectoryAddr, ethClient, logger)
	if err != nil {
		return nil, errors.Join(errors.New("Failed to create ELChainReader"), err)
	}
	return NewAvsRegistryChainWriter(
		serviceManagerAddr,
		registryCoordinator,
		operatorStateRetriever,
		stakeRegistry,
		blsApkRegistry,
		elReader,
		logger,
		ethClient,
		txMgr,
	)
}

// TODO(samlaf): clean up this function
func (w *AvsRegistryChainWriter) RegisterOperatorInQuorumWithAVSRegistryCoordinator(
	ctx context.Context,
	// we need to pass the private key explicitly and can't use the signer because registering requires signing a
	// message which isn't a transaction
	// and the signer can only signs transactions
	// see operatorSignature in
	// https://github.com/Layr-Labs/eigenlayer-middleware/blob/m2-mainnet/docs/RegistryCoordinator.md#registeroperator
	// TODO(madhur): check to see if we can make the signer and txmgr more flexible so we can use them (and remote
	// signers) to sign non txs
	operatorEcdsaPrivateKey *ecdsa.PrivateKey,
	operatorToAvsRegistrationSigSalt [32]byte,
	operatorToAvsRegistrationSigExpiry *big.Int,
	blsKeyPair *bls.KeyPair,
	quorumNumbers []byte,
	socket string,
) (*gethtypes.Receipt, error) {
	w.logger.Info("registering operator with the AVS's registry coordinator")
	// params to register bls pubkey with bls apk registry
	operatorAddr := crypto.PubkeyToAddress(operatorEcdsaPrivateKey.PublicKey)
	g1HashedMsgToSign, err := w.registryCoordinator.PubkeyRegistrationMessageHash(&bind.CallOpts{}, operatorAddr)
	if err != nil {
		return nil, err
	}
	signedMsg := utils.ConvertToBN254G1Point(
		blsKeyPair.SignHashedToCurveMessage(utils.ConvertBn254GethToGnark(g1HashedMsgToSign)).G1Point,
	)
	G1pubkeyBN254 := utils.ConvertToBN254G1Point(blsKeyPair.GetPubKeyG1())
	G2pubkeyBN254 := utils.ConvertToBN254G2Point(blsKeyPair.GetPubKeyG2())
	pubkeyRegParams := regcoord.IBLSApkRegistryPubkeyRegistrationParams{
		PubkeyRegistrationSignature: signedMsg,
		PubkeyG1:                    G1pubkeyBN254,
		PubkeyG2:                    G2pubkeyBN254,
	}

	// params to register operator in delegation manager's operator-avs mapping
	msgToSign, err := w.elReader.CalculateOperatorAVSRegistrationDigestHash(
		&bind.CallOpts{},
		operatorAddr,
		w.serviceManagerAddr,
		operatorToAvsRegistrationSigSalt,
		operatorToAvsRegistrationSigExpiry,
	)
	if err != nil {
		return nil, err
	}
	operatorSignature, err := crypto.Sign(msgToSign[:], operatorEcdsaPrivateKey)
	if err != nil {
		return nil, err
	}
	// this is annoying, and not sure why its needed, but seems like some historical baggage
	// see https://github.com/ethereum/go-ethereum/issues/28757#issuecomment-1874525854
	// and https://twitter.com/pcaversaccio/status/1671488928262529031
	operatorSignature[64] += 27
	operatorSignatureWithSaltAndExpiry := regcoord.ISignatureUtilsSignatureWithSaltAndExpiry{
		Signature: operatorSignature,
		Salt:      operatorToAvsRegistrationSigSalt,
		Expiry:    operatorToAvsRegistrationSigExpiry,
	}

	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, err
	}
	// TODO: this call will fail if max number of operators are already registered
	// in that case, need to call churner to kick out another operator. See eigenDA's node/operator.go implementation
	tx, err := w.registryCoordinator.RegisterOperator(
		noSendTxOpts,
		quorumNumbers,
		socket,
		pubkeyRegParams,
		operatorSignatureWithSaltAndExpiry,
	)
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

func (w *AvsRegistryChainWriter) UpdateStakesOfEntireOperatorSetForQuorums(
	ctx context.Context,
	operatorsPerQuorum [][]gethcommon.Address,
	quorumNumbers []byte,
) (*gethtypes.Receipt, error) {
	w.logger.Info("updating stakes for entire operator set", "quorumNumbers", quorumNumbers)
	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, err
	}
	tx, err := w.registryCoordinator.UpdateOperatorsForQuorum(noSendTxOpts, operatorsPerQuorum, quorumNumbers)
	if err != nil {
		return nil, err
	}
	receipt, err := w.txMgr.Send(ctx, tx)
	if err != nil {
		return nil, errors.New("failed to send tx with err: " + err.Error())
	}
	w.logger.Infof("tx hash: %s", tx.Hash().String())
	w.logger.Info("updated stakes for entire operator set", "quorumNumbers", quorumNumbers)
	return receipt, nil

}

func (w *AvsRegistryChainWriter) UpdateStakesOfOperatorSubsetForAllQuorums(
	ctx context.Context,
	operators []gethcommon.Address,
) (*gethtypes.Receipt, error) {
	w.logger.Info("updating stakes of operator subset for all quorums", "operators", operators)
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
	w.logger.Info("updated stakes of operator subset for all quorums", "operators", operators)
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
