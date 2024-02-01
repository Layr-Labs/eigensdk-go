package avsecdsaregistry

import (
	"context"
	"errors"
	"math/big"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/elcontracts"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	sdkecdsa "github.com/Layr-Labs/eigensdk-go/crypto/ecdsa"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethcommon "github.com/ethereum/go-ethereum/common"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	regcoord "github.com/Layr-Labs/eigensdk-go/contracts/bindings/ECDSARegistryCoordinator"
	stakereg "github.com/Layr-Labs/eigensdk-go/contracts/bindings/ECDSAStakeRegistry"
)

type AvsEcdsaRegistryWriter interface {
	// TODO(samlaf): an operator that is already registered in a quorum can register with another quorum without passing signatures
	//               perhaps we should add another sdk function for this purpose, that just takes in a quorumNumber?
	// RegisterOperatorInQuorumWithAVSRegistryCoordinator is used to register a single operator with the AVS's registry coordinator.
	//  - operatorEcdsaPrivateKey is the operator's ecdsa private key (used to sign a message to register operator in eigenlayer's delegation manager)
	//  - operatorToAvsRegistrationSigSalt is a random salt used to prevent replay attacks
	//  - operatorToAvsRegistrationSigExpiry is the expiry time of the signature
	RegisterOperatorInQuorumWithAVSRegistryCoordinator(
		ctx context.Context,
		operatorEcdsaPrivateKey *sdkecdsa.PrivateKey,
		operatorToAvsRegistrationSigSalt [32]byte,
		operatorToAvsRegistrationSigExpiry *big.Int,
		avsEcdsaKeyPair *sdkecdsa.PrivateKey,
		quorumNumbers []byte,
	) (*gethtypes.Receipt, error)

	// UpdateStakesOfEntireOperatorSetForQuorums is used by avs teams running https://github.com/Layr-Labs/avs-sync
	// to updates the stake of their entire operator set.
	// Because of high gas costs of this operation, it typically needs to be called for every quorum, or perhaps for a small grouping of quorums
	// (highly dependent on number of operators per quorum)
	UpdateStakesOfEntireOperatorSetForQuorums(
		ctx context.Context,
		operatorsPerQuorum [][]gethcommon.Address,
		quorumNumbers []byte,
	) (*gethtypes.Receipt, error)

	// UpdateStakesOfOperatorSubsetForAllQuorums is meant to be used by single operators (or teams of operators, possibly running https://github.com/Layr-Labs/avs-sync)
	// to update the stake of their own operator(s). This might be needed in the case that they received a lot of new stake delegations, and want this to be reflected
	// in the AVS's registry coordinator.
	UpdateStakesOfOperatorSubsetForAllQuorums(
		ctx context.Context,
		operators []gethcommon.Address,
	) (*gethtypes.Receipt, error)

	DeregisterOperator(
		ctx context.Context,
		quorumNumbers []byte,
	) (*gethtypes.Receipt, error)
}

type AvsEcdsaRegistryChainWriter struct {
	serviceManagerAddr  gethcommon.Address
	registryCoordinator *regcoord.ContractECDSARegistryCoordinator
	elReader            elcontracts.ELReader
	logger              logging.Logger
	ethClient           eth.EthClient
	txMgr               txmgr.TxManager
}

var _ AvsEcdsaRegistryWriter = (*AvsEcdsaRegistryChainWriter)(nil)

func NewAvsRegistryChainWriter(
	serviceManagerAddr gethcommon.Address,
	registryCoordinator *regcoord.ContractECDSARegistryCoordinator,
	elReader elcontracts.ELReader,
	logger logging.Logger,
	ethClient eth.EthClient,
	txMgr txmgr.TxManager,
) (*AvsEcdsaRegistryChainWriter, error) {
	return &AvsEcdsaRegistryChainWriter{
		serviceManagerAddr:  serviceManagerAddr,
		registryCoordinator: registryCoordinator,
		elReader:            elReader,
		logger:              logger,
		ethClient:           ethClient,
		txMgr:               txMgr,
	}, nil
}

func BuildAvsRegistryChainWriter(
	registryCoordinatorAddr gethcommon.Address,
	operatorStateRetrieverAddr gethcommon.Address,
	logger logging.Logger,
	ethClient eth.EthClient,
	txMgr txmgr.TxManager,
) (*AvsEcdsaRegistryChainWriter, error) {
	registryCoordinator, err := regcoord.NewContractECDSARegistryCoordinator(registryCoordinatorAddr, ethClient)
	if err != nil {
		return nil, err
	}
	serviceManagerAddr, err := registryCoordinator.ServiceManager(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	stakeRegistryAddr, err := registryCoordinator.StakeRegistry(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	stakeRegistry, err := stakereg.NewContractECDSAStakeRegistry(stakeRegistryAddr, ethClient)
	if err != nil {
		return nil, err
	}
	delegationManagerAddr, err := stakeRegistry.Delegation(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	elReader, err := elcontracts.BuildELChainReader(delegationManagerAddr, ethClient, logger)
	if err != nil {
		return nil, err
	}
	return NewAvsRegistryChainWriter(
		serviceManagerAddr,
		registryCoordinator,
		elReader,
		logger,
		ethClient,
		txMgr,
	)
}

// RegisterOperatorInQuorumWithAVSRegistryCoordinator is used to register a single operator with the AVS's registry coordinator,
// in a given set of quorum numbers.
// We reuse operatorToAvsRegistrationSigSalt and operatorToAvsRegistrationSigExpiry for both the operator's registration with the AVS's registry coordinator
// and the operator's registration with the eigenlayer's delegation manager.
func (w *AvsEcdsaRegistryChainWriter) RegisterOperatorInQuorumWithAVSRegistryCoordinator(
	ctx context.Context,
	// we need to pass the private key explicitly and can't use the signer because registering requires signing a message which isn't a transaction
	// and the signer can only signs transactions
	// see operatorSignature in https://github.com/Layr-Labs/eigenlayer-middleware/blob/m2-mainnet/docs/RegistryCoordinator.md#registeroperator
	// TODO(madhur): check to see if we can make the signer and txmgr more flexible so we can use them (and remote signers) to sign non txs
	operatorEcdsaPrivateKey *sdkecdsa.PrivateKey,
	operatorToAvsRegistrationSigSalt [32]byte,
	operatorToAvsRegistrationSigExpiry *big.Int,
	avsEcdsaPrivateKey *sdkecdsa.PrivateKey,
	quorumNumbers []byte,
) (*gethtypes.Receipt, error) {
	w.logger.Info("registering operator with the AVS's registry coordinator")
	// params to register avs ecdsa pubkey with ecdsaRegistryCoordinator
	operatorAddr := crypto.PubkeyToAddress(operatorEcdsaPrivateKey.PublicKey)
	// Note that we sign a message containing the operatorAddr, which basically tells the contract
	// "I am binding this avsEcdsa key that is signing the message to my operatorAddr"
	msgHashToSign, err := w.registryCoordinator.PubkeyRegistrationMessageHash(&bind.CallOpts{}, operatorAddr)
	if err != nil {
		return nil, err
	}
	signedMsg, err := sdkecdsa.SignMsg(msgHashToSign[:], avsEcdsaPrivateKey)
	if err != nil {
		return nil, err
	}
	avsEcdsaAddr := crypto.PubkeyToAddress(avsEcdsaPrivateKey.PublicKey)
	pubkeyRegParams := regcoord.ECDSARegistryCoordinatorECDSAPubkeyRegistrationParams{
		SigningAddress: avsEcdsaAddr,
		SignatureAndExpiry: regcoord.ISignatureUtilsSignatureWithSaltAndExpiry{
			Signature: signedMsg,
			Salt:      operatorToAvsRegistrationSigSalt,
			Expiry:    operatorToAvsRegistrationSigExpiry,
		},
	}

	// params to register operator in delegation manager's operator-avs mapping
	msgToSign, err := w.elReader.CalculateDelegationApprovalDigestHash(
		&bind.CallOpts{}, operatorAddr, w.serviceManagerAddr, operatorToAvsRegistrationSigSalt, operatorToAvsRegistrationSigExpiry)
	if err != nil {
		return nil, err
	}
	operatorSignature, err := sdkecdsa.SignMsg(msgToSign[:], operatorEcdsaPrivateKey)
	if err != nil {
		return nil, err
	}
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

func (w *AvsEcdsaRegistryChainWriter) UpdateStakesOfEntireOperatorSetForQuorums(
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

func (w *AvsEcdsaRegistryChainWriter) UpdateStakesOfOperatorSubsetForAllQuorums(
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

func (w *AvsEcdsaRegistryChainWriter) DeregisterOperator(
	ctx context.Context,
	quorumNumbers []byte,
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
