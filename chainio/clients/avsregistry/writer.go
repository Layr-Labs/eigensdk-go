package avsregistry

import (
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethcommon "github.com/ethereum/go-ethereum/common"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/elcontracts"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	chainioutils "github.com/Layr-Labs/eigensdk-go/chainio/utils"
	blsapkregistry "github.com/Layr-Labs/eigensdk-go/contracts/bindings/BLSApkRegistry"
	opstateretriever "github.com/Layr-Labs/eigensdk-go/contracts/bindings/OperatorStateRetriever"
	regcoord "github.com/Layr-Labs/eigensdk-go/contracts/bindings/RegistryCoordinator"
	servicemanager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/ServiceManagerBase"
	stakeregistry "github.com/Layr-Labs/eigensdk-go/contracts/bindings/StakeRegistry"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/Layr-Labs/eigensdk-go/utils"
)

type eLReader interface {
	CalculateOperatorAVSRegistrationDigestHash(
		ctx context.Context,
		operatorAddr gethcommon.Address,
		serviceManagerAddr gethcommon.Address,
		operatorToAvsRegistrationSigSalt [32]byte,
		operatorToAvsRegistrationSigExpiry *big.Int,
	) ([32]byte, error)
}

// The ChainWriter provides methods to call the
// AVS registry contract's state-changing functions.
type ChainWriter struct {
	serviceManagerAddr     gethcommon.Address
	registryCoordinator    *regcoord.ContractRegistryCoordinator
	operatorStateRetriever *opstateretriever.ContractOperatorStateRetriever
	stakeRegistry          *stakeregistry.ContractStakeRegistry
	blsApkRegistry         *blsapkregistry.ContractBLSApkRegistry
	elReader               eLReader
	logger                 logging.Logger
	ethClient              eth.HttpBackend
	txMgr                  txmgr.TxManager
}

// Returns a new instance of ChainWriter.
func NewChainWriter(
	serviceManagerAddr gethcommon.Address,
	registryCoordinator *regcoord.ContractRegistryCoordinator,
	operatorStateRetriever *opstateretriever.ContractOperatorStateRetriever,
	stakeRegistry *stakeregistry.ContractStakeRegistry,
	blsApkRegistry *blsapkregistry.ContractBLSApkRegistry,
	elReader eLReader,
	logger logging.Logger,
	ethClient eth.HttpBackend,
	txMgr txmgr.TxManager,
) *ChainWriter {
	logger = logger.With(logging.ComponentKey, "avsregistry/ChainWriter")

	return &ChainWriter{
		serviceManagerAddr:     serviceManagerAddr,
		registryCoordinator:    registryCoordinator,
		operatorStateRetriever: operatorStateRetriever,
		stakeRegistry:          stakeRegistry,
		blsApkRegistry:         blsApkRegistry,
		elReader:               elReader,
		logger:                 logger,
		ethClient:              ethClient,
		txMgr:                  txMgr,
	}
}

// NewWriterFromConfig creates a new ChainWriter from the provided config
func NewWriterFromConfig(
	cfg Config,
	client eth.HttpBackend,
	txMgr txmgr.TxManager,
	logger logging.Logger,
) (*ChainWriter, error) {
	bindings, err := NewBindingsFromConfig(cfg, client, logger)
	if err != nil {
		return nil, err
	}
	elReader, err := elcontracts.NewReaderFromConfig(elcontracts.Config{
		DelegationManagerAddress: bindings.DelegationManagerAddr,
		AvsDirectoryAddress:      bindings.AvsDirectoryAddr,
		DontUseAllocationManager: cfg.DontUseAllocationManager,
	}, client, logger)
	if err != nil {
		return nil, err
	}

	return NewChainWriter(
		bindings.ServiceManagerAddr,
		bindings.RegistryCoordinator,
		bindings.OperatorStateRetriever,
		bindings.StakeRegistry,
		bindings.BlsApkRegistry,
		elReader,
		logger,
		client,
		txMgr,
	), nil
}

// RegisterOperator is used to register a single operator with the AVS's registry coordinator.
func (w *ChainWriter) RegisterOperator(
	ctx context.Context,
	// we need to pass the private key explicitly and can't use the signer because registering requires signing a
	// message which isn't a transaction and the signer can only signs transactions. See operatorSignature in
	// https://github.com/Layr-Labs/eigenlayer-middleware/blob/m2-mainnet/docs/RegistryCoordinator.md#registeroperator
	// TODO(madhur): check to see if we can make the signer and txmgr more flexible so we can use them (and remote
	// signers) to sign non txs
	operatorEcdsaPrivateKey *ecdsa.PrivateKey,
	blsKeyPair *bls.KeyPair,
	quorumNumbers types.QuorumNums,
	socket string,
	waitForReceipt bool,
) (*gethtypes.Receipt, error) {
	operatorAddr := crypto.PubkeyToAddress(operatorEcdsaPrivateKey.PublicKey)
	w.logger.Info(
		"registering operator with the AVS's registry coordinator",
		"avs-service-manager",
		w.serviceManagerAddr,
		"operator",
		operatorAddr,
		"quorumNumbers",
		quorumNumbers,
		"socket",
		socket,
	)
	// params to register bls pubkey with bls apk registry
	g1HashedMsgToSign, err := w.registryCoordinator.PubkeyRegistrationMessageHash(&bind.CallOpts{}, operatorAddr)
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

	// generate a random salt and 1 hour expiry for the signature
	var signatureSalt [32]byte
	_, err = rand.Read(signatureSalt[:])
	if err != nil {
		return nil, err
	}

	curBlockNum, err := w.ethClient.BlockNumber(context.Background())
	if err != nil {
		return nil, err
	}
	curBlock, err := w.ethClient.BlockByNumber(context.Background(), new(big.Int).SetUint64(curBlockNum))
	if err != nil {
		return nil, err
	}
	sigValidForSeconds := int64(60 * 60) // 1 hour

	curTime := new(big.Int).SetUint64(curBlock.Time())
	signatureExpiry := new(big.Int).Add(curTime, big.NewInt(sigValidForSeconds))

	// params to register operator in delegation manager's operator-avs mapping
	msgToSign, err := w.elReader.CalculateOperatorAVSRegistrationDigestHash(
		ctx,
		operatorAddr,
		w.serviceManagerAddr,
		signatureSalt,
		signatureExpiry,
	)
	if err != nil {
		return nil, err
	}
	operatorSignature, err := crypto.Sign(msgToSign[:], operatorEcdsaPrivateKey)
	if err != nil {
		return nil, err
	}
	// the crypto library is low level and deals with 0/1 v values, whereas ethereum expects 27/28, so we add 27
	// see https://github.com/ethereum/go-ethereum/issues/28757#issuecomment-1874525854
	// and https://twitter.com/pcaversaccio/status/1671488928262529031
	operatorSignature[64] += 27
	operatorSignatureWithSaltAndExpiry := regcoord.ISignatureUtilsSignatureWithSaltAndExpiry{
		Signature: operatorSignature,
		Salt:      signatureSalt,
		Expiry:    signatureExpiry,
	}

	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, err
	}
	// TODO: this call will fail if max number of operators are already registered
	// in that case, need to call churner to kick out another operator. See eigenDA's node/operator.go implementation
	tx, err := w.registryCoordinator.RegisterOperator(
		noSendTxOpts,
		quorumNumbers.UnderlyingType(),
		socket,
		pubkeyRegParams,
		operatorSignatureWithSaltAndExpiry,
	)
	if err != nil {
		return nil, err
	}
	receipt, err := w.txMgr.Send(ctx, tx, waitForReceipt)
	if err != nil {
		return nil, utils.WrapError("failed to send tx with err", err)
	}
	w.logger.Info(
		"successfully registered operator with AVS registry coordinator",
		"txHash",
		receipt.TxHash.String(),
		"avs-service-manager",
		w.serviceManagerAddr,
		"operator",
		operatorAddr,
		"quorumNumbers",
		quorumNumbers,
	)
	return receipt, nil
}

// UpdateStakesOfEntireOperatorSetForQuorums is used by avs teams running https://github.com/Layr-Labs/avs-sync
// to updates the stake of their entire operator set.
// Because of high gas costs of this operation, it typically needs to be called for every quorum, or perhaps for a
// small grouping of quorums
// (highly dependent on number of operators per quorum)
func (w *ChainWriter) UpdateStakesOfEntireOperatorSetForQuorums(
	ctx context.Context,
	operatorsPerQuorum [][]gethcommon.Address,
	quorumNumbers types.QuorumNums,
	waitForReceipt bool,
) (*gethtypes.Receipt, error) {
	w.logger.Info("updating stakes for entire operator set", "quorumNumbers", quorumNumbers)
	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, err
	}
	tx, err := w.registryCoordinator.UpdateOperatorsForQuorum(
		noSendTxOpts,
		operatorsPerQuorum,
		quorumNumbers.UnderlyingType(),
	)
	if err != nil {
		return nil, err
	}
	receipt, err := w.txMgr.Send(ctx, tx, waitForReceipt)
	if err != nil {
		return nil, utils.WrapError("failed to send tx with err: ", err)
	}
	w.logger.Info(
		"successfully updated stakes for entire operator set",
		"txHash",
		receipt.TxHash.String(),
		"quorumNumbers",
		quorumNumbers,
	)
	return receipt, nil

}

// Registers an operator while replacing existing operators in full quorums. If any quorum reaches its maximum
// operator capacity, `operatorKickParams` is used to replace an old operator with the new one.
func (w *ChainWriter) RegisterOperatorWithChurn(
	ctx context.Context,
	operatorEcdsaPrivateKey *ecdsa.PrivateKey,
	churnApprovalEcdsaPrivateKey *ecdsa.PrivateKey,
	blsKeyPair *bls.KeyPair,
	quorumNumbers types.QuorumNums,
	quorumNumbersToKick types.QuorumNums,
	operatorsToKick []gethcommon.Address,
	socket string,
	waitForReceipt bool,
) (*gethtypes.Receipt, error) {
	operatorAddr := crypto.PubkeyToAddress(operatorEcdsaPrivateKey.PublicKey)
	g1HashedMsgToSign, err := w.registryCoordinator.PubkeyRegistrationMessageHash(&bind.CallOpts{}, operatorAddr)
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

	// generate a random salt and 1 hour expiry for the signature
	var signatureSalt [32]byte
	_, err = rand.Read(signatureSalt[:])
	if err != nil {
		return nil, err
	}

	curBlockNum, err := w.ethClient.BlockNumber(context.Background())
	if err != nil {
		return nil, err
	}
	curBlock, err := w.ethClient.BlockByNumber(context.Background(), new(big.Int).SetUint64(curBlockNum))
	if err != nil {
		return nil, err
	}
	sigValidForSeconds := int64(60 * 60) // 1 hour

	curTime := new(big.Int).SetUint64(curBlock.Time())
	signatureExpiry := new(big.Int).Add(curTime, big.NewInt(sigValidForSeconds))

	// params to register operator in delegation manager's operator-avs mapping
	msgToSign, err := w.elReader.CalculateOperatorAVSRegistrationDigestHash(
		ctx,
		operatorAddr,
		w.serviceManagerAddr,
		signatureSalt,
		signatureExpiry,
	)
	if err != nil {
		return nil, err
	}
	operatorSignature, err := crypto.Sign(msgToSign[:], operatorEcdsaPrivateKey)
	if err != nil {
		return nil, err
	}
	// the crypto library is low level and deals with 0/1 v values, whereas ethereum expects 27/28, so we add 27
	// see https://github.com/ethereum/go-ethereum/issues/28757#issuecomment-1874525854
	// and https://twitter.com/pcaversaccio/status/1671488928262529031
	operatorSignature[64] += 27
	operatorSignatureWithSaltAndExpiry := regcoord.ISignatureUtilsSignatureWithSaltAndExpiry{
		Signature: operatorSignature,
		Salt:      signatureSalt,
		Expiry:    signatureExpiry,
	}

	var operatorKickParams []regcoord.ISlashingRegistryCoordinatorTypesOperatorKickParam
	for i, operatorToKick := range operatorsToKick {
		operatorKickParams = append(operatorKickParams, regcoord.ISlashingRegistryCoordinatorTypesOperatorKickParam{
			Operator:     operatorToKick,
			QuorumNumber: quorumNumbersToKick[i].UnderlyingType(),
		})
	}
	var churnSignatureSalt [32]byte
	_, err = rand.Read(churnSignatureSalt[:])
	if err != nil {
		return nil, err
	}

	var operatorIdBytes [32]byte

	// `GetOperatorID()` returns the operator ID with `0x` prefix.
	// We need to remove the prefix and decode the hex string to bytes.
	operatorId := blsKeyPair.GetPubKeyG1().GetOperatorID()
	operatorIdNoPrefix := strings.TrimPrefix(operatorId, "0x")
	operatorIdBytesDecoded, err := hex.DecodeString(operatorIdNoPrefix)
	if err != nil {
		return nil, err
	}
	copy(operatorIdBytes[:], operatorIdBytesDecoded)

	churnMsgToSign, err := w.registryCoordinator.CalculateOperatorChurnApprovalDigestHash(
		&bind.CallOpts{Context: ctx},
		operatorAddr,
		operatorIdBytes,
		operatorKickParams,
		churnSignatureSalt,
		signatureExpiry,
	)
	if err != nil {
		return nil, err
	}

	churnApprovalSignature, err := crypto.Sign(churnMsgToSign[:], churnApprovalEcdsaPrivateKey)
	if err != nil {
		return nil, err
	}

	// the crypto library is low level and deals with 0/1 v values, whereas ethereum expects 27/28, so we add 27
	// see https://github.com/ethereum/go-ethereum/issues/28757#issuecomment-1874525854
	// and https://twitter.com/pcaversaccio/status/1671488928262529031
	churnApprovalSignature[64] += 27
	churnApproverSignatureWithSaltAndExpiry := regcoord.ISignatureUtilsSignatureWithSaltAndExpiry{
		Signature: churnApprovalSignature,
		Salt:      churnSignatureSalt,
		Expiry:    signatureExpiry,
	}

	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, err
	}

	tx, err := w.registryCoordinator.RegisterOperatorWithChurn(
		noSendTxOpts,
		quorumNumbers.UnderlyingType(),
		socket,
		pubkeyRegParams,
		operatorKickParams,
		churnApproverSignatureWithSaltAndExpiry,
		operatorSignatureWithSaltAndExpiry,
	)

	if err != nil {
		return nil, err
	}
	receipt, err := w.txMgr.Send(ctx, tx, waitForReceipt)
	if err != nil {
		return nil, utils.WrapError("failed to send tx with err", err.Error())
	}
	w.logger.Info(
		"successfully registered operator with AVS registry coordinator",
		"txHash",
		receipt.TxHash.String(),
		"avs-service-manager",
		w.serviceManagerAddr,
		"operator",
		operatorAddr,
		"quorumNumbers",
		quorumNumbers,
	)
	return receipt, nil
}

// Updates the stakes of a the given `operators` for all the quorums.
// On success, returns the receipt of the transaction.
func (w *ChainWriter) UpdateStakesOfOperatorSubsetForAllQuorums(
	ctx context.Context,
	operators []gethcommon.Address,
	waitForReceipt bool,
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
	receipt, err := w.txMgr.Send(ctx, tx, waitForReceipt)
	if err != nil {
		return nil, utils.WrapError("failed to send tx with err", err)
	}
	w.logger.Info(
		"successfully updated stakes of operator subset for all quorums",
		"txHash",
		receipt.TxHash.String(),
		"operators",
		operators,
	)
	return receipt, nil
}

// Deregisters the caller from the quorums given by `quorumNumbers`.
// On success, returns the receipt of the transaction.
func (w *ChainWriter) DeregisterOperator(
	ctx context.Context,
	quorumNumbers types.QuorumNums,
	pubkey regcoord.BN254G1Point,
	waitForReceipt bool,
) (*gethtypes.Receipt, error) {
	w.logger.Info("deregistering operator with the AVS's registry coordinator")
	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, err
	}
	tx, err := w.registryCoordinator.DeregisterOperator0(noSendTxOpts, quorumNumbers.UnderlyingType())
	if err != nil {
		return nil, err
	}
	receipt, err := w.txMgr.Send(ctx, tx, waitForReceipt)
	if err != nil {
		return nil, utils.WrapError("failed to send tx with err", err)
	}
	w.logger.Info(
		"successfully deregistered operator with the AVS's registry coordinator",
		"txHash",
		receipt.TxHash.String(),
	)
	return receipt, nil
}

// Updates the socket of the sender (if it is a registered operator).
// On success, returns the receipt of the transaction.
func (w *ChainWriter) UpdateSocket(
	ctx context.Context,
	socket types.Socket,
	waitForReceipt bool,
) (*gethtypes.Receipt, error) {
	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, err
	}
	tx, err := w.registryCoordinator.UpdateSocket(noSendTxOpts, socket.String())
	if err != nil {
		return nil, err
	}
	receipt, err := w.txMgr.Send(ctx, tx, waitForReceipt)
	if err != nil {
		return nil, utils.WrapError("failed to send UpdateSocket tx with err", err)
	}
	return receipt, nil
}

// Sets the rewards initiator address as the received as parameter. This address is the only one
// that can initiate rewards. Returns the receipt of the transaction in case of success.
func (w *ChainWriter) SetRewardsInitiator(
	ctx context.Context,
	rewardsInitiatorAddr gethcommon.Address,
	waitForReceipt bool,
) (*gethtypes.Receipt, error) {
	w.logger.Info("setting rewards initiator with addr ", rewardsInitiatorAddr)

	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, err
	}

	// TODO: create ServiceManager binding in `NewChainWriter`
	// (see https://github.com/Layr-Labs/eigensdk-go/issues/552)
	serviceManagerContract, err := servicemanager.NewContractServiceManagerBase(
		w.serviceManagerAddr,
		w.ethClient,
	)
	if err != nil {
		return nil, utils.WrapError("failed to create ServiceManager contract", err)
	}
	tx, err := serviceManagerContract.SetRewardsInitiator(noSendTxOpts, rewardsInitiatorAddr)
	if err != nil {
		return nil, err
	}
	receipt, err := w.txMgr.Send(ctx, tx, waitForReceipt)
	if err != nil {
		return nil, utils.WrapError("failed to send SetRewardsInitiator tx with err", err)
	}
	return receipt, nil
}

// Receives the quorum number to modify and the new look ahead period to set. Sets the look ahead
// time for checking operator shares for a specific quorum, and returns the receipt of the
// transaction in case of success.
func (w *ChainWriter) SetSlashableStakeLookahead(
	ctx context.Context,
	quorumNumber uint8,
	lookAheadPeriod uint32,
	waitForReceipt bool,
) (*gethtypes.Receipt, error) {
	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, err
	}
	tx, err := w.stakeRegistry.SetSlashableStakeLookahead(noSendTxOpts, quorumNumber, lookAheadPeriod)
	if err != nil {
		return nil, err
	}
	receipt, err := w.txMgr.Send(ctx, tx, waitForReceipt)
	if err != nil {
		return nil, utils.WrapError("failed to send SetSlashableStakeLookahead tx with err", err)
	}
	return receipt, nil
}

func (w *ChainWriter) SetMinimumStakeForQuorum(
	ctx context.Context,
	quorumNumber uint8,
	minimumStake *big.Int,
	waitForReceipt bool,
) (*gethtypes.Receipt, error) {
	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, err
	}
	tx, err := w.stakeRegistry.SetMinimumStakeForQuorum(noSendTxOpts, quorumNumber, minimumStake)
	if err != nil {
		return nil, err
	}
	receipt, err := w.txMgr.Send(ctx, tx, waitForReceipt)
	if err != nil {
		return nil, utils.WrapError("failed to send SetMinimumStake tx with err", err)
	}
	return receipt, nil
}

// Creates a new quorum that tracks total delegated stake for operators.
// It receives the operator set parameters for the given quorum and the minimum stake required to register.
// Returns the transaction receipt in case of success.
func (w *ChainWriter) CreateTotalDelegatedStakeQuorum(
	ctx context.Context,
	operatorSetParams regcoord.ISlashingRegistryCoordinatorTypesOperatorSetParam,
	minimumStakeRequired *big.Int,
	strategyParams []regcoord.IStakeRegistryTypesStrategyParams,
	waitForReceipt bool,
) (*gethtypes.Receipt, error) {
	w.logger.Info("Creating total delegated stake quorum")

	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, err
	}

	tx, err := w.registryCoordinator.CreateTotalDelegatedStakeQuorum(
		noSendTxOpts,
		operatorSetParams,
		minimumStakeRequired,
		strategyParams,
	)
	if err != nil {
		return nil, err
	}
	receipt, err := w.txMgr.Send(ctx, tx, waitForReceipt)
	if err != nil {
		return nil, utils.WrapError("failed to send CreateTotalDelegatedStakeQuorum tx with err", err)
	}
	return receipt, nil
}

// Creates a new quorum that tracks slashable stake for operators.
// It receives the operator set parameters for the given quorum, the minimum stake required to register,
// and the number of blocks to look ahead when calculating slashable stake.
// Returns the transaction receipt in case of success.
// Note: This function only works on M2 AVSs.
func (w *ChainWriter) CreateSlashableStakeQuorum(
	ctx context.Context,
	operatorSetParams regcoord.ISlashingRegistryCoordinatorTypesOperatorSetParam,
	minimumStakeRequired *big.Int,
	strategyParams []regcoord.IStakeRegistryTypesStrategyParams,
	lookAheadPeriod uint32,
	waitForReceipt bool,
) (*gethtypes.Receipt, error) {
	w.logger.Info("Creating slashable stake quorum")

	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, err
	}

	tx, err := w.registryCoordinator.CreateSlashableStakeQuorum(
		noSendTxOpts,
		operatorSetParams,
		minimumStakeRequired,
		strategyParams,
		lookAheadPeriod,
	)
	if err != nil {
		return nil, err
	}
	receipt, err := w.txMgr.Send(ctx, tx, waitForReceipt)
	if err != nil {
		return nil, utils.WrapError("failed to send CreateSlashableStakeQuorum tx with err", err)
	}
	return receipt, nil
}

// Receives an operator address and quorum numbers and ejects the operator from the given quorums.
// Note: if the operator is not registered, the call will not fail, but will do nothing.
func (w *ChainWriter) EjectOperator(
	ctx context.Context,
	operatorAddress gethcommon.Address,
	quorumNumbers types.QuorumNums,
	waitForReceipt bool,
) (*gethtypes.Receipt, error) {
	w.logger.Info("ejecting operator with address ", operatorAddress, " from quorum numbers ", quorumNumbers)

	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, err
	}
	tx, err := w.registryCoordinator.EjectOperator(noSendTxOpts, operatorAddress, quorumNumbers.UnderlyingType())
	if err != nil {
		return nil, err
	}
	receipt, err := w.txMgr.Send(ctx, tx, waitForReceipt)
	if err != nil {
		return nil, utils.WrapError("failed to send EjectOperator tx with err", err)
	}
	return receipt, nil
}

// Sets the operator set params for the quorum which id matches the quorum number.
// Params consists in a new max operator count and operator churn parameters
// Returns the transaction receipt in case of success.
func (w *ChainWriter) SetOperatorSetParams(
	ctx context.Context,
	quorumNumber uint8,
	operatorSetParams regcoord.ISlashingRegistryCoordinatorTypesOperatorSetParam,
	waitForReceipt bool,
) (*gethtypes.Receipt, error) {
	w.logger.Info("setting operator set params for quorum ", quorumNumber)

	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, err
	}
	tx, err := w.registryCoordinator.SetOperatorSetParams(noSendTxOpts, quorumNumber, operatorSetParams)
	if err != nil {
		return nil, err
	}
	receipt, err := w.txMgr.Send(ctx, tx, waitForReceipt)
	if err != nil {
		return nil, utils.WrapError("failed to send SetOperatorSetParams tx with err", err)
	}
	return receipt, nil
}

// Sets the churnApprover as the address received as parameter. The churnApprover's signature is required in
// churn related methods (like churn registration). Returns the receipt of the transaction in case of success.
func (w *ChainWriter) SetChurnApprover(
	ctx context.Context,
	churnApproverAddress gethcommon.Address,
	waitForReceipt bool,
) (*gethtypes.Receipt, error) {
	w.logger.Info("setting churn approver with address ", churnApproverAddress)

	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, err
	}
	tx, err := w.registryCoordinator.SetChurnApprover(noSendTxOpts, churnApproverAddress)
	if err != nil {
		return nil, err
	}
	receipt, err := w.txMgr.Send(ctx, tx, waitForReceipt)
	if err != nil {
		return nil, utils.WrapError("failed to send SetChurnApprover tx with err", err)
	}
	return receipt, nil
}

// Sets the ejector address to the one received by parameter. This address is the only one
// that can eject operators. Returns the receipt of the transaction in case of success.
func (w *ChainWriter) SetEjector(
	ctx context.Context,
	ejectorAddress gethcommon.Address,
	waitForReceipt bool,
) (*gethtypes.Receipt, error) {
	w.logger.Info("setting ejector with address ", ejectorAddress)

	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, err
	}
	tx, err := w.registryCoordinator.SetEjector(noSendTxOpts, ejectorAddress)
	if err != nil {
		return nil, err
	}
	receipt, err := w.txMgr.Send(ctx, tx, waitForReceipt)
	if err != nil {
		return nil, utils.WrapError("failed to send SetEjector tx with err", err)
	}
	return receipt, nil
}

// Modifies the multiplier of existing strategies for the given quorum number.
func (w *ChainWriter) ModifyStrategyParams(
	ctx context.Context,
	quorumNumber types.QuorumNum,
	strategyIndices []*big.Int,
	multipliers []*big.Int,
	waitForReceipt bool,
) (*gethtypes.Receipt, error) {
	w.logger.Info("modifying strategy params for quorum ", quorumNumber)

	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, err
	}
	tx, err := w.stakeRegistry.ModifyStrategyParams(
		noSendTxOpts,
		quorumNumber.UnderlyingType(),
		strategyIndices,
		multipliers,
	)
	if err != nil {
		return nil, err
	}
	receipt, err := w.txMgr.Send(ctx, tx, waitForReceipt)
	if err != nil {
		return nil, utils.WrapError("failed to send ModifyStrategyParams tx with err", err)
	}
	return receipt, nil
}

// Sets the avs as the address received as parameter. Identifier should only be set once, since
// changing it could break existing operator sets. Returns the receipt of the transaction in case of success.
func (w *ChainWriter) SetAvs(
	ctx context.Context,
	avsAddress gethcommon.Address,
	waitForReceipt bool,
) (*gethtypes.Receipt, error) {
	w.logger.Info("setting account identifier with address ", avsAddress)

	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, err
	}
	tx, err := w.registryCoordinator.SetAVS(noSendTxOpts, avsAddress)
	if err != nil {
		return nil, err
	}
	receipt, err := w.txMgr.Send(ctx, tx, waitForReceipt)
	if err != nil {
		return nil, utils.WrapError("failed to send SetAvs tx with err", err)
	}
	return receipt, nil
}

// Sets the ejection cooldown with the value received by parameter. The ejection cooldown is the time an operator has to
// wait to join any quorum after being rejected. Returns the receipt of the transaction in case of success.
func (w *ChainWriter) SetEjectionCooldown(
	ctx context.Context,
	ejectionCooldown *big.Int,
	waitForReceipt bool,
) (*gethtypes.Receipt, error) {
	w.logger.Info("setting ejection cooldown with value ", ejectionCooldown)

	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, err
	}
	tx, err := w.registryCoordinator.SetEjectionCooldown(noSendTxOpts, ejectionCooldown)
	if err != nil {
		return nil, err
	}
	receipt, err := w.txMgr.Send(ctx, tx, waitForReceipt)
	if err != nil {
		return nil, utils.WrapError("failed to send SetEjectionCooldown tx with err", err)
	}
	return receipt, nil
}

// Adds new strategies and their associated multipliers to the specified quorum.
func (w *ChainWriter) AddStrategies(
	ctx context.Context,
	quorumNumber types.QuorumNum,
	strategyParams []stakeregistry.IStakeRegistryTypesStrategyParams,
	waitForReceipt bool,
) (*gethtypes.Receipt, error) {
	w.logger.Info("adding strategies for quorum ", quorumNumber.UnderlyingType())

	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, err
	}
	tx, err := w.stakeRegistry.AddStrategies(noSendTxOpts, quorumNumber.UnderlyingType(), strategyParams)
	if err != nil {
		return nil, err
	}
	receipt, err := w.txMgr.Send(ctx, tx, waitForReceipt)
	if err != nil {
		return nil, utils.WrapError("failed to send AddStrategies tx with err", err)
	}
	return receipt, nil
}

// Updates the metadata URI for the AVS. Returns the receipt of the transaction in case of success.
func (w *ChainWriter) UpdateAVSMetadataURI(
	ctx context.Context,
	metadataUri string,
	waitForReceipt bool,
) (*gethtypes.Receipt, error) {
	w.logger.Info("updating AVS metadata URI with value ", metadataUri)

	// TODO: create ServiceManager binding in `NewChainWriter`
	// (see https://github.com/Layr-Labs/eigensdk-go/issues/552)
	serviceManagerContract, err := servicemanager.NewContractServiceManagerBase(
		w.serviceManagerAddr,
		w.ethClient,
	)
	if err != nil {
		return nil, utils.WrapError("failed to create ServiceManager contract", err)
	}

	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, err
	}
	tx, err := serviceManagerContract.UpdateAVSMetadataURI(noSendTxOpts, metadataUri)
	if err != nil {
		return nil, err
	}
	receipt, err := w.txMgr.Send(ctx, tx, waitForReceipt)
	if err != nil {
		return nil, utils.WrapError("failed to send updateAVSMetadataURI tx", err.Error())
	}
	return receipt, nil
}

// Removes strategies and their associated weights from the specified quorum.
func (w *ChainWriter) RemoveStrategies(
	ctx context.Context,
	quorumNumber types.QuorumNum,
	indicesToRemove []*big.Int,
	waitForReceipt bool,
) (*gethtypes.Receipt, error) {
	w.logger.Info("removing strategies from quorum ", quorumNumber)

	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, err
	}

	tx, err := w.stakeRegistry.RemoveStrategies(noSendTxOpts, quorumNumber.UnderlyingType(), indicesToRemove)
	if err != nil {
		return nil, err
	}

	receipt, err := w.txMgr.Send(ctx, tx, waitForReceipt)
	if err != nil {
		return nil, utils.WrapError("failed to remove strategies from quorum", err)
	}
	return receipt, nil
}

// Creates a new rewards submission to the EigenLayer RewardsCoordinator contract,
// to be split amongst the set of stakers delegated to operators who are registered
// to this `avs`. Returns the receipt of the transaction in case of success.
//
// Can fail in some cases:
//   - Only callable by the permissioned rewardsInitiator address
//   - The duration of the `rewardsSubmission` cannot exceed `MAX_REWARDS_DURATION`
//   - The tokens are sent to the `RewardsCoordinator` contract
//   - Strategies must be in ascending order of addresses to check for duplicates
//   - This function may fail to execute with a large number of submissions due to gas limits. Use a
//     smaller array of submissions if necessary.
func (w *ChainWriter) CreateAVSRewardsSubmission(
	ctx context.Context,
	rewardsSubmission []servicemanager.IRewardsCoordinatorTypesRewardsSubmission,
	waitForReceipt bool,
) (*gethtypes.Receipt, error) {
	w.logger.Info("creating AVS rewards submission ", "rewardsSubmission", rewardsSubmission)

	// TODO: store binding in struct
	serviceManagerContract, err := servicemanager.NewContractServiceManagerBase(
		w.serviceManagerAddr,
		w.ethClient,
	)
	if err != nil {
		return nil, utils.WrapError("failed to create ServiceManager contract", err)
	}

	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, err
	}
	tx, err := serviceManagerContract.CreateAVSRewardsSubmission(noSendTxOpts, rewardsSubmission)
	if err != nil {
		return nil, err
	}
	receipt, err := w.txMgr.Send(ctx, tx, waitForReceipt)
	if err != nil {
		return nil, utils.WrapError("failed to send CreateAVSRewardsSubmission tx with err", err)
	}
	return receipt, nil
}

// Creates a new operator-directed rewards submission, to be split amongst the operators and
// set of stakers delegated to operators who are registered to this `avs`.
// Returns the receipt of the transaction in case of success.
//
// Can fail in some cases:
//   - Only callable by the permissioned rewardsInitiator address
//   - The duration of the `rewardsSubmission` cannot exceed `MAX_REWARDS_DURATION`
//   - The tokens are sent to the `RewardsCoordinator` contract
//   - This contract needs a token approval of sum of all `operatorRewards` in the
//     `operatorDirectedRewardsSubmissions`, before calling this function.
//   - Strategies must be in ascending order of addresses to check for duplicates
//   - Operators must be in ascending order of addresses to check for duplicates
//   - This function will revert if the `operatorDirectedRewardsSubmissions` is malformed.
//   - This function may fail to execute with a large number of submissions due to gas limits. Use a
//     smaller array of submissions if necessary.
func (w *ChainWriter) CreateOperatorDirectedAVSRewardsSubmission(
	ctx context.Context,
	operatorDirectedRewardsSubmissions []servicemanager.IRewardsCoordinatorTypesOperatorDirectedRewardsSubmission,
	waitForReceipt bool,
) (*gethtypes.Receipt, error) {
	w.logger.Info("creating operator directed AVS rewards submission ", operatorDirectedRewardsSubmissions)

	// TODO: store binding in struct
	serviceManagerContract, err := servicemanager.NewContractServiceManagerBase(
		w.serviceManagerAddr,
		w.ethClient,
	)
	if err != nil {
		return nil, utils.WrapError("failed to create ServiceManager contract", err)
	}

	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, err
	}
	tx, err := serviceManagerContract.CreateOperatorDirectedAVSRewardsSubmission(
		noSendTxOpts,
		operatorDirectedRewardsSubmissions,
	)
	if err != nil {
		return nil, err
	}
	receipt, err := w.txMgr.Send(ctx, tx, waitForReceipt)
	if err != nil {
		return nil, utils.WrapError("failed to send CreateOperatorDirectedAVSRewardsSubmission tx with err", err)
	}
	return receipt, nil
}
