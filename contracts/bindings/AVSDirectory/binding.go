// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractAVSDirectory

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ISignatureUtilsSignatureWithSaltAndExpiry is an auto generated low-level Go binding around an user-defined struct.
type ISignatureUtilsSignatureWithSaltAndExpiry struct {
	Signature []byte
	Salt      [32]byte
	Expiry    *big.Int
}

// OperatorSet is an auto generated low-level Go binding around an user-defined struct.
type OperatorSet struct {
	Avs           common.Address
	OperatorSetId uint32
}

// ContractAVSDirectoryMetaData contains all meta data concerning the ContractAVSDirectory contract.
var ContractAVSDirectoryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_delegation\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"},{\"name\":\"_DEALLOCATION_DELAY\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DEALLOCATION_DELAY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"OPERATOR_AVS_REGISTRATION_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"OPERATOR_SET_FORCE_DEREGISTRATION_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"OPERATOR_SET_REGISTRATION_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addStrategiesToOperatorSet\",\"inputs\":[{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"avsOperatorStatus\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumIAVSDirectoryTypes.OperatorAVSRegistrationStatus\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"becomeOperatorSetAVS\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"calculateOperatorAVSRegistrationDigestHash\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateOperatorSetForceDeregistrationTypehash\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateOperatorSetRegistrationDigestHash\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"cancelSalt\",\"inputs\":[{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createOperatorSets\",\"inputs\":[{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deregisterOperatorFromAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deregisterOperatorFromOperatorSets\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"domainSeparator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"forceDeregisterFromOperatorSets\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getNumOperatorSetsOfOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNumOperatorsInOperatorSet\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorSetsOfOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"operatorSets\",\"type\":\"tuple[]\",\"internalType\":\"structOperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorsInOperatorSet\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStrategiesInOperatorSet\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"inTotalOperatorSets\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"initialPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isMember\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOperatorSet\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOperatorSetAVS\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOperatorSetBatch\",\"inputs\":[{\"name\":\"operatorSets\",\"type\":\"tuple[]\",\"internalType\":\"structOperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOperatorSlashable\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"migrateOperatorsToOperatorSets\",\"inputs\":[{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"operatorSaltIsSpent\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"isSpent\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorSetMemberAtIndex\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorSetStatus\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"registered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"lastDeregisteredTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorSetsMemberOfAtIndex\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerOperatorToAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerOperatorToOperatorSets\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeStrategiesFromOperatorSet\",\"inputs\":[{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPauserRegistry\",\"inputs\":[{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateAVSMetadataURI\",\"inputs\":[{\"name\":\"metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AVSMetadataURIUpdated\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"metadataURI\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AVSMigratedToOperatorSets\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorAVSRegistrationStatusUpdated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"status\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumIAVSDirectoryTypes.OperatorAVSRegistrationStatus\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorAddedToOperatorSet\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorMigratedToOperatorSets\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"indexed\":false,\"internalType\":\"uint32[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorRemovedFromOperatorSet\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSetCreated\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PauserRegistrySet\",\"inputs\":[{\"name\":\"pauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StrategyAddedToOperatorSet\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StrategyRemovedFromOperatorSet\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CurrentlyPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputAddressZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAVS\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidNewPausedStatus\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOperator\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOperatorSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyPauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyUnpauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorAlreadyRegisteredToAVS\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorNotRegisteredToAVS\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorNotRegisteredToEigenLayer\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SaltSpent\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SignatureExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StrategyAlreadyInOperatorSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StrategyNotInOperatorSet\",\"inputs\":[]}]",
	Bin: "0x61010060405234801561001157600080fd5b50604051613b4f380380613b4f833981016040819052610030916101ce565b6001600160a01b03821660805263ffffffff811660a0524660c052610053610065565b60e05261005e61010f565b505061021d565b600060c05146146101085750604080518082018252600a81526922b4b3b2b72630bcb2b960b11b60209182015281517f8cad95687ba82c2ce50e74f7b754645e5117c3a5bec8151c0726d5857980a866818301527f71b625cfad44bac63b13dba07f2e1d6084ee04b6f8752101ece6126d584ee6ea81840152466060820152306080808301919091528351808303909101815260a0909101909252815191012090565b5060e05190565b600054610100900460ff161561017b5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff908116146101cc576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b600080604083850312156101e157600080fd5b82516001600160a01b03811681146101f857600080fd5b602084015190925063ffffffff8116811461021257600080fd5b809150509250929050565b60805160a05160c05160e0516138e461026b60003960006128ae015260006127ee0152600081816103d1015261090a01526000818161071801528181610cf7015261160201526138e46000f3fe608060405234801561001057600080fd5b50600436106102a05760003560e01c80638da5cb5b11610167578063cbdf0e42116100ce578063e88d804911610087578063e88d80491461068c578063ec76f4421461073a578063ef2dfa8d1461076e578063f2fde38b14610781578063f698da2514610794578063fabc1cbc1461079c57600080fd5b8063cbdf0e421461068c578063ce7b5e4b1461069f578063d79aceab146106b2578063da2ff05d146106d9578063dce974b9146106ec578063df5cf7231461071357600080fd5b8063aec205c511610120578063aec205c514610604578063afe02ed51461060c578063b2841d481461061f578063b5a768ca14610632578063c1a8e2c514610652578063c825fe681461066557600080fd5b80638da5cb5b14610594578063955e6696146105a55780639926ee7d146105b8578063a1060c88146105cb578063a364f4da146105de578063a98fb355146105f157600080fd5b8063411d415b1161020b578063715018a6116101c4578063715018a6146105025780637357723b1461050a5780637673e93a1461051d578063769993421461054057806384d76f7b14610553578063886f11951461058157600080fd5b8063411d415b146104495780634177a87c1461047457806349075da314610494578063595c6a67146104cf5780635ac86ab7146104d75780635c975abb146104fa57600080fd5b80631e2199e21161025d5780631e2199e2146103495780631e68134e1461035c57806320c4e236146103b95780632981eb77146103cc578063374823b5146104085780633fee332d1461043657600080fd5b80631023aa35146102a557806310d67a2f146102cb5780631352c3e6146102e0578063136439dd1461030357806316ae76cb146103165780631794bb3c14610336575b600080fd5b6102b86102b3366004612d9b565b6107af565b6040519081526020015b60405180910390f35b6102de6102d9366004612db7565b6107da565b005b6102f36102ee366004612dd4565b61088e565b60405190151581526020016102c2565b6102de610311366004612e0a565b61093f565b610329610324366004612e23565b610a2a565b6040516102c29190612e58565b6102de610344366004612ebf565b610b66565b6102de610357366004612ff9565b610c8a565b61039d61036a366004613078565b609f60209081526000938452604080852082529284528284209052825290205460ff811690610100900463ffffffff1682565b60408051921515835263ffffffff9091166020830152016102c2565b6102f36103c73660046130bf565b610e5b565b6103f37f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff90911681526020016102c2565b6102f3610416366004613134565b609960209081526000928352604080842090915290825290205460ff1681565b6102de610444366004613160565b610f19565b61045c6104573660046131f2565b611046565b6040516001600160a01b0390911681526020016102c2565b610487610482366004612d9b565b61107c565b6040516102c2919061321d565b6104c26104a236600461325e565b609860209081526000928352604080842090915290825290205460ff1681565b6040516102c291906132ad565b6102de611147565b6102f36104e53660046132d5565b606654600160ff9092169190911b9081161490565b6066546102b8565b6102de61120f565b6104876105183660046132f8565b611223565b6102f361052b366004612db7565b609a6020526000908152604090205460ff1681565b6102de61054e36600461332c565b611311565b6102f361056136600461337e565b609b60209081526000928352604080842090915290825290205460ff1681565b60655461045c906001600160a01b031681565b6033546001600160a01b031661045c565b6102b86105b33660046133aa565b61145f565b6102de6105c6366004613410565b6114c4565b6102b86105d936600461345f565b611733565b6102de6105ec366004612db7565b61179d565b6102de6105ff3660046134a5565b6118b6565b6102de6118fd565b6102de61061a366004613507565b611970565b6102b861062d3660046133aa565b611ae0565b610645610640366004613134565b611b20565b6040516102c29190613548565b6102de61066036600461356e565b611b5a565b6102b87f809c5ac049c45b7a7f050a20f00c16cf63797efbf8b1eb8d749fdfa39ff8f92981565b6102b861069a366004612db7565b611b8f565b6102de6106ad36600461332c565b611bb0565b6102b87fda2c89bafdd34776a2b8bb9c83c82f419e20cc8c67207f70edd58249b92661bd81565b6102f36106e7366004612dd4565b611cfe565b6102b87f4ee65f64218c67b68da66fd0db16560040a6b973290b9e71912d661ee53fe49581565b61045c7f000000000000000000000000000000000000000000000000000000000000000081565b6102de610748366004612e0a565b33600090815260996020908152604080832093835292905220805460ff19166001179055565b6102de61077c36600461358e565b611d2a565b6102de61078f366004612db7565b611fd2565b6102b8612048565b6102de6107aa366004612e0a565b612057565b60006107d4609d60006107c18561215f565b81526020019081526020016000206121c4565b92915050565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561082d573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061085191906135fd565b6001600160a01b0316336001600160a01b0316146108825760405163794821ff60e01b815260040160405180910390fd5b61088b816121ce565b50565b600061089a8383611cfe565b156108a7575060016107d4565b81516001600160a01b039081166000908152609f6020908152604080832093871683529281528282208186015163ffffffff90811684529082529183902083518085019094525460ff811615158452610100900490911690820181905261092f907f000000000000000000000000000000000000000000000000000000000000000090613630565b63ffffffff164210949350505050565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa158015610987573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109ab919061364c565b6109c857604051631d77d47760e21b815260040160405180910390fd5b606654818116146109ec5760405163c61dca5d60e01b815260040160405180910390fd5b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d906020015b60405180910390a250565b6001600160a01b0383166000908152609c60205260408120606091908490610a51906121c4565b610a5b919061366e565b905080831115610a69578092505b826001600160401b03811115610a8157610a81612ca0565b604051908082528060200260200182016040528015610ac657816020015b6040805180820190915260008082526020820152815260200190600190039081610a9f5790505b50915060005b83811015610b5d57610b38610b02610ae48388613681565b6001600160a01b0389166000908152609c602052604090209061225e565b60408051808201909152600080825260208201525060408051808201909152606082901c815263ffffffff909116602082015290565b838281518110610b4a57610b4a613694565b6020908102919091010152600101610acc565b50509392505050565b600054610100900460ff1615808015610b865750600054600160ff909116105b80610ba05750303b158015610ba0575060005460ff166001145b610c085760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b6000805460ff191660011790558015610c2b576000805461ff0019166101001790555b610c35838361226a565b610c3e846122ef565b8015610c84576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050565b606654600190600290811603610cb35760405163840a48d560e01b815260040160405180910390fd5b4282604001511015610cd857604051630819bdcd60e01b815260040160405180910390fd5b6040516336b87bd760e11b81526001600160a01b0386811660048301527f00000000000000000000000000000000000000000000000000000000000000001690636d70f7ae90602401602060405180830381865afa158015610d3e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d62919061364c565b610d7f57604051639f88c8af60e01b815260040160405180910390fd5b336000908152609a602052604090205460ff16610daf576040516366e565df60e01b815260040160405180910390fd5b6001600160a01b038516600090815260996020908152604080832085830151845290915290205460ff1615610df757604051630d4c4c9160e21b815260040160405180910390fd5b610e1785610e103387878760200151886040015161145f565b8451612341565b6001600160a01b03851660009081526099602090815260408083208583015184529091529020805460ff19166001179055610e5485338686612372565b5050505050565b6000805b82811015610f0f57609b6000858584818110610e7d57610e7d613694565b610e939260206040909202019081019150612db7565b6001600160a01b03166001600160a01b031681526020019081526020016000206000858584818110610ec757610ec7613694565b9050604002016020016020810190610edf91906136aa565b63ffffffff16815260208101919091526040016000205460ff16610f075760009150506107d4565b600101610e5f565b5060019392505050565b606654600190600290811603610f425760405163840a48d560e01b815260040160405180910390fd5b815151600003610f7a57336001600160a01b03871614610f755760405163ccea9e6f60e01b815260040160405180910390fd5b611032565b4282604001511015610f9f57604051630819bdcd60e01b815260040160405180910390fd5b6001600160a01b038616600090815260996020908152604080832085830151845290915290205460ff1615610fe757604051630d4c4c9160e21b815260040160405180910390fd5b61100086610e1087878787602001518860400151611ae0565b6001600160a01b03861660009081526099602090815260408083208583015184529091529020805460ff191660011790555b61103e85878686612572565b505050505050565b600061107582609d60006110598761215f565b815260200190815260200160002061225e90919063ffffffff16565b9392505050565b606060006110898361215f565b6000818152609e60205260408120919250906110a4906121c4565b9050806001600160401b038111156110be576110be612ca0565b6040519080825280602002602001820160405280156110e7578160200160208202803683370190505b50925060005b8181101561113f576000838152609e6020526040902061110d908261225e565b84828151811061111f5761111f613694565b6001600160a01b03909216602092830291909101909101526001016110ed565b505050919050565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa15801561118f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111b3919061364c565b6111d057604051631d77d47760e21b815260040160405180910390fd5b600019606681905560405190815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2565b611217612707565b61122160006122ef565b565b606060006112308561215f565b6000818152609d6020526040812091925090859061124d906121c4565b611257919061366e565b905080841115611265578093505b836001600160401b0381111561127d5761127d612ca0565b6040519080825280602002602001820160405280156112a6578160200160208202803683370190505b50925060005b84811015611307576112d56112c18288613681565b6000858152609d602052604090209061225e565b8482815181106112e7576112e7613694565b6001600160a01b03909216602092830291909101909101526001016112ac565b5050509392505050565b6040805180820182523380825263ffffffff861660208084018290526000928352609b815284832091835252919091205460ff1661136257604051631fb1705560e21b815260040160405180910390fd5b600061136d8261215f565b905060005b8381101561103e576113b885858381811061138f5761138f613694565b90506020020160208101906113a49190612db7565b6000848152609e6020526040902090612761565b6113d55760405163585cfb2f60e01b815260040160405180910390fd5b7f7ab260fe0af193db5f4986770d831bda4ea46099dc817e8b6716dcae8af8e88b8386868481811061140957611409613694565b905060200201602081019061141e9190612db7565b6040805183516001600160a01b03908116825260209485015163ffffffff1694820194909452929091169082015260600160405180910390a1600101611372565b60006114ba7f809c5ac049c45b7a7f050a20f00c16cf63797efbf8b1eb8d749fdfa39ff8f929878787878760405160200161149f9695949392919061370a565b60405160208183030381529060405280519060200120612776565b9695505050505050565b6066546000906001908116036114ed5760405163840a48d560e01b815260040160405180910390fd5b428260400151101561151257604051630819bdcd60e01b815260040160405180910390fd5b336000908152609a602052604090205460ff1615611543576040516366e565df60e01b815260040160405180910390fd5b60013360009081526098602090815260408083206001600160a01b038816845290915290205460ff16600181111561157d5761157d613297565b0361159b57604051631aa528bb60e11b815260040160405180910390fd5b6001600160a01b038316600090815260996020908152604080832085830151845290915290205460ff16156115e357604051630d4c4c9160e21b815260040160405180910390fd5b6040516336b87bd760e11b81526001600160a01b0384811660048301527f00000000000000000000000000000000000000000000000000000000000000001690636d70f7ae90602401602060405180830381865afa158015611649573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061166d919061364c565b61168a57604051639f88c8af60e01b815260040160405180910390fd5b6116a283610e10853386602001518760400151611733565b6001600160a01b038316600081815260996020908152604080832086830151845282528083208054600160ff19918216811790925533808652609885528386208787529094529382902080549094168117909355519092917ff0952b1c65271d819d39983d2abb044b9cace59bcc4d4dd389f586ebdcb15b419161172691906132ad565b60405180910390a3505050565b604080517fda2c89bafdd34776a2b8bb9c83c82f419e20cc8c67207f70edd58249b92661bd60208201526001600160a01b038087169282019290925290841660608201526080810183905260a081018290526000906117949060c00161149f565b95945050505050565b6066546000906001908116036117c65760405163840a48d560e01b815260040160405180910390fd5b60013360009081526098602090815260408083206001600160a01b038716845290915290205460ff16600181111561180057611800613297565b1461181e576040516352df45c960e01b815260040160405180910390fd5b336000908152609a602052604090205460ff161561184f576040516366e565df60e01b815260040160405180910390fd5b3360008181526098602090815260408083206001600160a01b0387168085529252808320805460ff191690555190917ff0952b1c65271d819d39983d2abb044b9cace59bcc4d4dd389f586ebdcb15b41916118aa91906132ad565b60405180910390a35050565b336001600160a01b03167fa89c1dc243d8908a96dd84944bcc97d6bc6ac00dd78e20621576be6a3c94371383836040516118f192919061374a565b60405180910390a25050565b336000908152609a602052604090205460ff161561192e576040516366e565df60e01b815260040160405180910390fd5b336000818152609a6020526040808220805460ff19166001179055517f702b0c1f6cb1cf511aaa81f72bc05a215bb3497632d72c690c822b044ab494bf9190a2565b60005b81811015611adb57336000908152609b602052604081209084848481811061199d5761199d613694565b90506020020160208101906119b291906136aa565b63ffffffff16815260208101919091526040016000205460ff16156119ea57604051631fb1705560e21b815260040160405180910390fd5b336000908152609b60205260408120600191858585818110611a0e57611a0e613694565b9050602002016020810190611a2391906136aa565b63ffffffff1663ffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055507f31629285ead2335ae0933f86ed2ae63321f7af77b4e6eaabc42c057880977e6c6040518060400160405280336001600160a01b03168152602001858585818110611aa157611aa1613694565b9050602002016020810190611ab691906136aa565b63ffffffff169052604051611acb9190613548565b60405180910390a1600101611973565b505050565b60006114ba7f4ee65f64218c67b68da66fd0db16560040a6b973290b9e71912d661ee53fe495878787878760405160200161149f9695949392919061370a565b60408051808201909152600080825260208201526001600160a01b0383166000908152609c6020526040902061107590610b02908461225e565b606654600190600290811603611b835760405163840a48d560e01b815260040160405180910390fd5b610c8433858585612572565b6001600160a01b0381166000908152609c602052604081206107d4906121c4565b6040805180820182523380825263ffffffff861660208084018290526000928352609b815284832091835252919091205460ff16611c0157604051631fb1705560e21b815260040160405180910390fd5b6000611c0c8261215f565b905060005b8381101561103e57611c57858583818110611c2e57611c2e613694565b9050602002016020810190611c439190612db7565b6000848152609e60205260409020906127bd565b611c74576040516331bc342760e11b815260040160405180910390fd5b7f7b4b073d80dcac55a11177d8459ad9f664ceeb91f71f27167bb14f8152a7eeee83868684818110611ca857611ca8613694565b9050602002016020810190611cbd9190612db7565b6040805183516001600160a01b03908116825260209485015163ffffffff1694820194909452929091169082015260600160405180910390a1600101611c11565b6000611075611d0c8361215f565b6001600160a01b0385166000908152609c60205260409020906127d2565b606654600190600290811603611d535760405163840a48d560e01b815260040160405180910390fd5b336000908152609a602052604090205460ff16611d83576040516366e565df60e01b815260040160405180910390fd5b60005b8481101561103e57600133600090815260986020526040812090888885818110611db257611db2613694565b9050602002016020810190611dc79190612db7565b6001600160a01b0316815260208101919091526040016000205460ff166001811115611df557611df5613297565b14611e135760405163ccea9e6f60e01b815260040160405180910390fd5b611e67868683818110611e2857611e28613694565b9050602002016020810190611e3d9190612db7565b33868685818110611e5057611e50613694565b9050602002810190611e629190613779565b612372565b33600090815260986020526040812081888885818110611e8957611e89613694565b9050602002016020810190611e9e9190612db7565b6001600160a01b031681526020810191909152604001600020805460ff191660018381811115611ed057611ed0613297565b021790555033868683818110611ee857611ee8613694565b9050602002016020810190611efd9190612db7565b6001600160a01b03167ff0952b1c65271d819d39983d2abb044b9cace59bcc4d4dd389f586ebdcb15b416000604051611f3691906132ad565b60405180910390a333868683818110611f5157611f51613694565b9050602002016020810190611f669190612db7565b6001600160a01b03167f54f33cfdd1ca703d795986b986fd47d742eab1904ecd2a5fdb8d6595e5904a01868685818110611fa257611fa2613694565b9050602002810190611fb49190613779565b604051611fc29291906137c2565b60405180910390a3600101611d86565b611fda612707565b6001600160a01b03811661203f5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610bff565b61088b816122ef565b60006120526127ea565b905090565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156120aa573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906120ce91906135fd565b6001600160a01b0316336001600160a01b0316146120ff5760405163794821ff60e01b815260040160405180910390fd5b6066541981196066541916146121285760405163c61dca5d60e01b815260040160405180910390fd5b606681905560405181815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c90602001610a1f565b60008160000151826020015163ffffffff166040516020016121ac92919060609290921b6bffffffffffffffffffffffff1916825260a01b6001600160a01b031916601482015260200190565b6040516020818303038152906040526107d4906137de565b60006107d4825490565b6001600160a01b0381166121f5576040516339b190bb60e11b815260040160405180910390fd5b606554604080516001600160a01b03928316815291831660208301527f6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6910160405180910390a1606580546001600160a01b0319166001600160a01b0392909216919091179055565b600061107583836128d0565b6065546001600160a01b031615801561228b57506001600160a01b03821615155b6122a8576040516339b190bb60e11b815260040160405180910390fd5b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a26122eb826121ce565b5050565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6123556001600160a01b03841683836128fa565b611adb57604051638baa579f60e01b815260040160405180910390fd5b60005b81811015610e545760006040518060400160405280866001600160a01b031681526020018585858181106123ab576123ab613694565b90506020020160208101906123c091906136aa565b63ffffffff1690526001600160a01b0386166000908152609b602052604081209192508585858181106123f5576123f5613694565b905060200201602081019061240a91906136aa565b63ffffffff16815260208101919091526040016000205460ff1661244157604051631fb1705560e21b815260040160405180910390fd5b600061244c8261215f565b6001600160a01b0388166000908152609c602052604090209091506124719082612951565b506000818152609d6020526040902061248a9088612761565b506001600160a01b038087166000908152609f60209081526040808320938b168352929052908120818787878181106124c5576124c5613694565b90506020020160208101906124da91906136aa565b63ffffffff1681526020810191909152604001600020805490915060ff16156125165760405163ccea9e6f60e01b815260040160405180910390fd5b805460ff191660011781556040516001600160a01b038916907f43232edf9071753d2321e5fa7e018363ee248e5f2142e6c08edd3265bfb4895e9061255c908690613548565b60405180910390a2505050806001019050612375565b60005b81811015610e545760006040518060400160405280876001600160a01b031681526020018585858181106125ab576125ab613694565b90506020020160208101906125c091906136aa565b63ffffffff169052905060006125d58261215f565b6001600160a01b0387166000908152609c602052604090209091506125fa908261295d565b506000818152609d6020526040902061261390876127bd565b506001600160a01b038088166000908152609f60209081526040808320938a1683529290529081208187878781811061264e5761264e613694565b905060200201602081019061266391906136aa565b63ffffffff1681526020810191909152604001600020805490915060ff1661269e5760405163ccea9e6f60e01b815260040160405180910390fd5b805464ffffffffff19166101004263ffffffff16021781556040516001600160a01b038816907fad34c3070be1dffbcaa499d000ba2b8d9848aefcac3059df245dd95c4ece14fe906126f1908690613548565b60405180910390a2505050806001019050612575565b6033546001600160a01b031633146112215760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610bff565b6000611075836001600160a01b038416612969565b60006127806127ea565b60405161190160f01b6020820152602281019190915260428101839052606201604051602081830303815290604052805190602001209050919050565b6000611075836001600160a01b0384166129b8565b60008181526001830160205260408120541515611075565b60007f000000000000000000000000000000000000000000000000000000000000000046146128ab5750604080518082018252600a81526922b4b3b2b72630bcb2b960b11b60209182015281517f8cad95687ba82c2ce50e74f7b754645e5117c3a5bec8151c0726d5857980a866818301527f71b625cfad44bac63b13dba07f2e1d6084ee04b6f8752101ece6126d584ee6ea81840152466060820152306080808301919091528351808303909101815260a0909101909252815191012090565b507f000000000000000000000000000000000000000000000000000000000000000090565b60008260000182815481106128e7576128e7613694565b9060005260206000200154905092915050565b60008060006129098585612aab565b9092509050600081600481111561292257612922613297565b1480156129405750856001600160a01b0316826001600160a01b0316145b806114ba57506114ba868686612af0565b60006110758383612969565b600061107583836129b8565b60008181526001830160205260408120546129b0575081546001818101845560008481526020808220909301849055845484825282860190935260409020919091556107d4565b5060006107d4565b60008181526001830160205260408120548015612aa15760006129dc60018361366e565b85549091506000906129f09060019061366e565b9050818114612a55576000866000018281548110612a1057612a10613694565b9060005260206000200154905080876000018481548110612a3357612a33613694565b6000918252602080832090910192909255918252600188019052604090208390555b8554869080612a6657612a66613805565b6001900381819060005260206000200160009055905585600101600086815260200190815260200160002060009055600193505050506107d4565b60009150506107d4565b6000808251604103612ae15760208301516040840151606085015160001a612ad587828585612bdc565b94509450505050612ae9565b506000905060025b9250929050565b6000806000856001600160a01b0316631626ba7e60e01b8686604051602401612b1a92919061383f565b60408051601f198184030181529181526020820180516001600160e01b03166001600160e01b0319909416939093179092529051612b589190613879565b600060405180830381855afa9150503d8060008114612b93576040519150601f19603f3d011682016040523d82523d6000602084013e612b98565b606091505b5091509150818015612bac57506020815110155b80156114ba57508051630b135d3f60e11b90612bd19083016020908101908401613895565b149695505050505050565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0831115612c135750600090506003612c97565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015612c67573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b038116612c9057600060019250925050612c97565b9150600090505b94509492505050565b634e487b7160e01b600052604160045260246000fd5b604051606081016001600160401b0381118282101715612cd857612cd8612ca0565b60405290565b604051601f8201601f191681016001600160401b0381118282101715612d0657612d06612ca0565b604052919050565b6001600160a01b038116811461088b57600080fd5b803563ffffffff81168114612d3757600080fd5b919050565b600060408284031215612d4e57600080fd5b604080519081016001600160401b0381118282101715612d7057612d70612ca0565b6040529050808235612d8181612d0e565b8152612d8f60208401612d23565b60208201525092915050565b600060408284031215612dad57600080fd5b6110758383612d3c565b600060208284031215612dc957600080fd5b813561107581612d0e565b60008060608385031215612de757600080fd5b8235612df281612d0e565b9150612e018460208501612d3c565b90509250929050565b600060208284031215612e1c57600080fd5b5035919050565b600080600060608486031215612e3857600080fd5b8335612e4381612d0e565b95602085013595506040909401359392505050565b602080825282518282018190526000918401906040840190835b81811015612eb457612e9e83855180516001600160a01b0316825260209081015163ffffffff16910152565b6020939093019260409290920191600101612e72565b509095945050505050565b600080600060608486031215612ed457600080fd5b8335612edf81612d0e565b92506020840135612eef81612d0e565b929592945050506040919091013590565b60008083601f840112612f1257600080fd5b5081356001600160401b03811115612f2957600080fd5b6020830191508360208260051b8501011115612ae957600080fd5b600060608284031215612f5657600080fd5b612f5e612cb6565b905081356001600160401b03811115612f7657600080fd5b8201601f81018413612f8757600080fd5b80356001600160401b03811115612fa057612fa0612ca0565b612fb3601f8201601f1916602001612cde565b818152856020838501011115612fc857600080fd5b8160208401602083013760006020928201830152835283810135908301525060409182013591810191909152919050565b6000806000806060858703121561300f57600080fd5b843561301a81612d0e565b935060208501356001600160401b0381111561303557600080fd5b61304187828801612f00565b90945092505060408501356001600160401b0381111561306057600080fd5b61306c87828801612f44565b91505092959194509250565b60008060006060848603121561308d57600080fd5b833561309881612d0e565b925060208401356130a881612d0e565b91506130b660408501612d23565b90509250925092565b600080602083850312156130d257600080fd5b82356001600160401b038111156130e857600080fd5b8301601f810185136130f957600080fd5b80356001600160401b0381111561310f57600080fd5b8560208260061b840101111561312457600080fd5b6020919091019590945092505050565b6000806040838503121561314757600080fd5b823561315281612d0e565b946020939093013593505050565b60008060008060006080868803121561317857600080fd5b853561318381612d0e565b9450602086013561319381612d0e565b935060408601356001600160401b038111156131ae57600080fd5b6131ba88828901612f00565b90945092505060608601356001600160401b038111156131d957600080fd5b6131e588828901612f44565b9150509295509295909350565b6000806060838503121561320557600080fd5b61320f8484612d3c565b946040939093013593505050565b602080825282518282018190526000918401906040840190835b81811015612eb45783516001600160a01b0316835260209384019390920191600101613237565b6000806040838503121561327157600080fd5b823561327c81612d0e565b9150602083013561328c81612d0e565b809150509250929050565b634e487b7160e01b600052602160045260246000fd5b60208101600283106132cf57634e487b7160e01b600052602160045260246000fd5b91905290565b6000602082840312156132e757600080fd5b813560ff8116811461107557600080fd5b60008060006080848603121561330d57600080fd5b6133178585612d3c565b95604085013595506060909401359392505050565b60008060006040848603121561334157600080fd5b61334a84612d23565b925060208401356001600160401b0381111561336557600080fd5b61337186828701612f00565b9497909650939450505050565b6000806040838503121561339157600080fd5b823561339c81612d0e565b9150612e0160208401612d23565b6000806000806000608086880312156133c257600080fd5b85356133cd81612d0e565b945060208601356001600160401b038111156133e857600080fd5b6133f488828901612f00565b9699909850959660408101359660609091013595509350505050565b6000806040838503121561342357600080fd5b823561342e81612d0e565b915060208301356001600160401b0381111561344957600080fd5b61345585828601612f44565b9150509250929050565b6000806000806080858703121561347557600080fd5b843561348081612d0e565b9350602085013561349081612d0e565b93969395505050506040820135916060013590565b600080602083850312156134b857600080fd5b82356001600160401b038111156134ce57600080fd5b8301601f810185136134df57600080fd5b80356001600160401b038111156134f557600080fd5b85602082840101111561312457600080fd5b6000806020838503121561351a57600080fd5b82356001600160401b0381111561353057600080fd5b61353c85828601612f00565b90969095509350505050565b81516001600160a01b0316815260208083015163ffffffff1690820152604081016107d4565b60008060006040848603121561358357600080fd5b833561334a81612d0e565b600080600080604085870312156135a457600080fd5b84356001600160401b038111156135ba57600080fd5b6135c687828801612f00565b90955093505060208501356001600160401b038111156135e557600080fd5b6135f187828801612f00565b95989497509550505050565b60006020828403121561360f57600080fd5b815161107581612d0e565b634e487b7160e01b600052601160045260246000fd5b63ffffffff81811683821601908111156107d4576107d461361a565b60006020828403121561365e57600080fd5b8151801515811461107557600080fd5b818103818111156107d4576107d461361a565b808201808211156107d4576107d461361a565b634e487b7160e01b600052603260045260246000fd5b6000602082840312156136bc57600080fd5b61107582612d23565b81835260208301925060008160005b848110156137005763ffffffff6136ea83612d23565b16865260209586019591909101906001016136d4565b5093949350505050565b8681526001600160a01b038616602082015260a06040820181905260009061373590830186886136c5565b60608301949094525060800152949350505050565b60208152816020820152818360408301376000818301604090810191909152601f909201601f19160101919050565b6000808335601e1984360301811261379057600080fd5b8301803591506001600160401b038211156137aa57600080fd5b6020019150600581901b3603821315612ae957600080fd5b6020815260006137d66020830184866136c5565b949350505050565b805160208083015191908110156137ff576000198160200360031b1b821691505b50919050565b634e487b7160e01b600052603160045260246000fd5b60005b8381101561383657818101518382015260200161381e565b50506000910152565b828152604060208201526000825180604084015261386481606085016020870161381b565b601f01601f1916919091016060019392505050565b6000825161388b81846020870161381b565b9190910192915050565b6000602082840312156138a757600080fd5b505191905056fea26469706673582212207688cbb7f3c517bbd106b6acdfbc909011606bdf7617d1f4d9cef871c330425f64736f6c634300081b0033",
}

// ContractAVSDirectoryABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractAVSDirectoryMetaData.ABI instead.
var ContractAVSDirectoryABI = ContractAVSDirectoryMetaData.ABI

// ContractAVSDirectoryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractAVSDirectoryMetaData.Bin instead.
var ContractAVSDirectoryBin = ContractAVSDirectoryMetaData.Bin

// DeployContractAVSDirectory deploys a new Ethereum contract, binding an instance of ContractAVSDirectory to it.
func DeployContractAVSDirectory(auth *bind.TransactOpts, backend bind.ContractBackend, _delegation common.Address, _DEALLOCATION_DELAY uint32) (common.Address, *types.Transaction, *ContractAVSDirectory, error) {
	parsed, err := ContractAVSDirectoryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractAVSDirectoryBin), backend, _delegation, _DEALLOCATION_DELAY)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractAVSDirectory{ContractAVSDirectoryCaller: ContractAVSDirectoryCaller{contract: contract}, ContractAVSDirectoryTransactor: ContractAVSDirectoryTransactor{contract: contract}, ContractAVSDirectoryFilterer: ContractAVSDirectoryFilterer{contract: contract}}, nil
}

// ContractAVSDirectoryMethods is an auto generated interface around an Ethereum contract.
type ContractAVSDirectoryMethods interface {
	ContractAVSDirectoryCalls
	ContractAVSDirectoryTransacts
	ContractAVSDirectoryFilters
}

// ContractAVSDirectoryCalls is an auto generated interface that defines the call methods available for an Ethereum contract.
type ContractAVSDirectoryCalls interface {
	DEALLOCATIONDELAY(opts *bind.CallOpts) (uint32, error)

	OPERATORAVSREGISTRATIONTYPEHASH(opts *bind.CallOpts) ([32]byte, error)

	OPERATORSETFORCEDEREGISTRATIONTYPEHASH(opts *bind.CallOpts) ([32]byte, error)

	OPERATORSETREGISTRATIONTYPEHASH(opts *bind.CallOpts) ([32]byte, error)

	AvsOperatorStatus(opts *bind.CallOpts, avs common.Address, operator common.Address) (uint8, error)

	CalculateOperatorAVSRegistrationDigestHash(opts *bind.CallOpts, operator common.Address, avs common.Address, salt [32]byte, expiry *big.Int) ([32]byte, error)

	CalculateOperatorSetForceDeregistrationTypehash(opts *bind.CallOpts, avs common.Address, operatorSetIds []uint32, salt [32]byte, expiry *big.Int) ([32]byte, error)

	CalculateOperatorSetRegistrationDigestHash(opts *bind.CallOpts, avs common.Address, operatorSetIds []uint32, salt [32]byte, expiry *big.Int) ([32]byte, error)

	Delegation(opts *bind.CallOpts) (common.Address, error)

	DomainSeparator(opts *bind.CallOpts) ([32]byte, error)

	GetNumOperatorSetsOfOperator(opts *bind.CallOpts, operator common.Address) (*big.Int, error)

	GetNumOperatorsInOperatorSet(opts *bind.CallOpts, operatorSet OperatorSet) (*big.Int, error)

	GetOperatorSetsOfOperator(opts *bind.CallOpts, operator common.Address, start *big.Int, length *big.Int) ([]OperatorSet, error)

	GetOperatorsInOperatorSet(opts *bind.CallOpts, operatorSet OperatorSet, start *big.Int, length *big.Int) ([]common.Address, error)

	GetStrategiesInOperatorSet(opts *bind.CallOpts, operatorSet OperatorSet) ([]common.Address, error)

	InTotalOperatorSets(opts *bind.CallOpts, operator common.Address) (*big.Int, error)

	IsMember(opts *bind.CallOpts, operator common.Address, operatorSet OperatorSet) (bool, error)

	IsOperatorSet(opts *bind.CallOpts, avs common.Address, operatorSetId uint32) (bool, error)

	IsOperatorSetAVS(opts *bind.CallOpts, avs common.Address) (bool, error)

	IsOperatorSetBatch(opts *bind.CallOpts, operatorSets []OperatorSet) (bool, error)

	IsOperatorSlashable(opts *bind.CallOpts, operator common.Address, operatorSet OperatorSet) (bool, error)

	OperatorSaltIsSpent(opts *bind.CallOpts, operator common.Address, salt [32]byte) (bool, error)

	OperatorSetMemberAtIndex(opts *bind.CallOpts, operatorSet OperatorSet, index *big.Int) (common.Address, error)

	OperatorSetStatus(opts *bind.CallOpts, operator common.Address, avs common.Address, operatorSetId uint32) (struct {
		Registered                bool
		LastDeregisteredTimestamp uint32
	}, error)

	OperatorSetsMemberOfAtIndex(opts *bind.CallOpts, operator common.Address, index *big.Int) (OperatorSet, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	Paused(opts *bind.CallOpts, index uint8) (bool, error)

	Paused0(opts *bind.CallOpts) (*big.Int, error)

	PauserRegistry(opts *bind.CallOpts) (common.Address, error)
}

// ContractAVSDirectoryTransacts is an auto generated interface that defines the transact methods available for an Ethereum contract.
type ContractAVSDirectoryTransacts interface {
	AddStrategiesToOperatorSet(opts *bind.TransactOpts, operatorSetId uint32, strategies []common.Address) (*types.Transaction, error)

	BecomeOperatorSetAVS(opts *bind.TransactOpts) (*types.Transaction, error)

	CancelSalt(opts *bind.TransactOpts, salt [32]byte) (*types.Transaction, error)

	CreateOperatorSets(opts *bind.TransactOpts, operatorSetIds []uint32) (*types.Transaction, error)

	DeregisterOperatorFromAVS(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error)

	DeregisterOperatorFromOperatorSets(opts *bind.TransactOpts, operator common.Address, operatorSetIds []uint32) (*types.Transaction, error)

	ForceDeregisterFromOperatorSets(opts *bind.TransactOpts, operator common.Address, avs common.Address, operatorSetIds []uint32, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error)

	Initialize(opts *bind.TransactOpts, initialOwner common.Address, _pauserRegistry common.Address, initialPausedStatus *big.Int) (*types.Transaction, error)

	MigrateOperatorsToOperatorSets(opts *bind.TransactOpts, operators []common.Address, operatorSetIds [][]uint32) (*types.Transaction, error)

	Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error)

	PauseAll(opts *bind.TransactOpts) (*types.Transaction, error)

	RegisterOperatorToAVS(opts *bind.TransactOpts, operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error)

	RegisterOperatorToOperatorSets(opts *bind.TransactOpts, operator common.Address, operatorSetIds []uint32, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error)

	RemoveStrategiesFromOperatorSet(opts *bind.TransactOpts, operatorSetId uint32, strategies []common.Address) (*types.Transaction, error)

	RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	SetPauserRegistry(opts *bind.TransactOpts, newPauserRegistry common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error)

	Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error)

	UpdateAVSMetadataURI(opts *bind.TransactOpts, metadataURI string) (*types.Transaction, error)
}

// ContractAVSDirectoryFilterer is an auto generated interface that defines the log filtering methods available for an Ethereum contract.
type ContractAVSDirectoryFilters interface {
	FilterAVSMetadataURIUpdated(opts *bind.FilterOpts, avs []common.Address) (*ContractAVSDirectoryAVSMetadataURIUpdatedIterator, error)
	WatchAVSMetadataURIUpdated(opts *bind.WatchOpts, sink chan<- *ContractAVSDirectoryAVSMetadataURIUpdated, avs []common.Address) (event.Subscription, error)
	ParseAVSMetadataURIUpdated(log types.Log) (*ContractAVSDirectoryAVSMetadataURIUpdated, error)

	FilterAVSMigratedToOperatorSets(opts *bind.FilterOpts, avs []common.Address) (*ContractAVSDirectoryAVSMigratedToOperatorSetsIterator, error)
	WatchAVSMigratedToOperatorSets(opts *bind.WatchOpts, sink chan<- *ContractAVSDirectoryAVSMigratedToOperatorSets, avs []common.Address) (event.Subscription, error)
	ParseAVSMigratedToOperatorSets(log types.Log) (*ContractAVSDirectoryAVSMigratedToOperatorSets, error)

	FilterInitialized(opts *bind.FilterOpts) (*ContractAVSDirectoryInitializedIterator, error)
	WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractAVSDirectoryInitialized) (event.Subscription, error)
	ParseInitialized(log types.Log) (*ContractAVSDirectoryInitialized, error)

	FilterOperatorAVSRegistrationStatusUpdated(opts *bind.FilterOpts, operator []common.Address, avs []common.Address) (*ContractAVSDirectoryOperatorAVSRegistrationStatusUpdatedIterator, error)
	WatchOperatorAVSRegistrationStatusUpdated(opts *bind.WatchOpts, sink chan<- *ContractAVSDirectoryOperatorAVSRegistrationStatusUpdated, operator []common.Address, avs []common.Address) (event.Subscription, error)
	ParseOperatorAVSRegistrationStatusUpdated(log types.Log) (*ContractAVSDirectoryOperatorAVSRegistrationStatusUpdated, error)

	FilterOperatorAddedToOperatorSet(opts *bind.FilterOpts, operator []common.Address) (*ContractAVSDirectoryOperatorAddedToOperatorSetIterator, error)
	WatchOperatorAddedToOperatorSet(opts *bind.WatchOpts, sink chan<- *ContractAVSDirectoryOperatorAddedToOperatorSet, operator []common.Address) (event.Subscription, error)
	ParseOperatorAddedToOperatorSet(log types.Log) (*ContractAVSDirectoryOperatorAddedToOperatorSet, error)

	FilterOperatorMigratedToOperatorSets(opts *bind.FilterOpts, operator []common.Address, avs []common.Address) (*ContractAVSDirectoryOperatorMigratedToOperatorSetsIterator, error)
	WatchOperatorMigratedToOperatorSets(opts *bind.WatchOpts, sink chan<- *ContractAVSDirectoryOperatorMigratedToOperatorSets, operator []common.Address, avs []common.Address) (event.Subscription, error)
	ParseOperatorMigratedToOperatorSets(log types.Log) (*ContractAVSDirectoryOperatorMigratedToOperatorSets, error)

	FilterOperatorRemovedFromOperatorSet(opts *bind.FilterOpts, operator []common.Address) (*ContractAVSDirectoryOperatorRemovedFromOperatorSetIterator, error)
	WatchOperatorRemovedFromOperatorSet(opts *bind.WatchOpts, sink chan<- *ContractAVSDirectoryOperatorRemovedFromOperatorSet, operator []common.Address) (event.Subscription, error)
	ParseOperatorRemovedFromOperatorSet(log types.Log) (*ContractAVSDirectoryOperatorRemovedFromOperatorSet, error)

	FilterOperatorSetCreated(opts *bind.FilterOpts) (*ContractAVSDirectoryOperatorSetCreatedIterator, error)
	WatchOperatorSetCreated(opts *bind.WatchOpts, sink chan<- *ContractAVSDirectoryOperatorSetCreated) (event.Subscription, error)
	ParseOperatorSetCreated(log types.Log) (*ContractAVSDirectoryOperatorSetCreated, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractAVSDirectoryOwnershipTransferredIterator, error)
	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractAVSDirectoryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error)
	ParseOwnershipTransferred(log types.Log) (*ContractAVSDirectoryOwnershipTransferred, error)

	FilterPaused(opts *bind.FilterOpts, account []common.Address) (*ContractAVSDirectoryPausedIterator, error)
	WatchPaused(opts *bind.WatchOpts, sink chan<- *ContractAVSDirectoryPaused, account []common.Address) (event.Subscription, error)
	ParsePaused(log types.Log) (*ContractAVSDirectoryPaused, error)

	FilterPauserRegistrySet(opts *bind.FilterOpts) (*ContractAVSDirectoryPauserRegistrySetIterator, error)
	WatchPauserRegistrySet(opts *bind.WatchOpts, sink chan<- *ContractAVSDirectoryPauserRegistrySet) (event.Subscription, error)
	ParsePauserRegistrySet(log types.Log) (*ContractAVSDirectoryPauserRegistrySet, error)

	FilterStrategyAddedToOperatorSet(opts *bind.FilterOpts) (*ContractAVSDirectoryStrategyAddedToOperatorSetIterator, error)
	WatchStrategyAddedToOperatorSet(opts *bind.WatchOpts, sink chan<- *ContractAVSDirectoryStrategyAddedToOperatorSet) (event.Subscription, error)
	ParseStrategyAddedToOperatorSet(log types.Log) (*ContractAVSDirectoryStrategyAddedToOperatorSet, error)

	FilterStrategyRemovedFromOperatorSet(opts *bind.FilterOpts) (*ContractAVSDirectoryStrategyRemovedFromOperatorSetIterator, error)
	WatchStrategyRemovedFromOperatorSet(opts *bind.WatchOpts, sink chan<- *ContractAVSDirectoryStrategyRemovedFromOperatorSet) (event.Subscription, error)
	ParseStrategyRemovedFromOperatorSet(log types.Log) (*ContractAVSDirectoryStrategyRemovedFromOperatorSet, error)

	FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*ContractAVSDirectoryUnpausedIterator, error)
	WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ContractAVSDirectoryUnpaused, account []common.Address) (event.Subscription, error)
	ParseUnpaused(log types.Log) (*ContractAVSDirectoryUnpaused, error)
}

// ContractAVSDirectory is an auto generated Go binding around an Ethereum contract.
type ContractAVSDirectory struct {
	ContractAVSDirectoryCaller     // Read-only binding to the contract
	ContractAVSDirectoryTransactor // Write-only binding to the contract
	ContractAVSDirectoryFilterer   // Log filterer for contract events
}

// ContractAVSDirectory implements the ContractAVSDirectoryMethods interface.
var _ ContractAVSDirectoryMethods = (*ContractAVSDirectory)(nil)

// ContractAVSDirectoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractAVSDirectoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractAVSDirectoryCaller implements the ContractAVSDirectoryCalls interface.
var _ ContractAVSDirectoryCalls = (*ContractAVSDirectoryCaller)(nil)

// ContractAVSDirectoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractAVSDirectoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractAVSDirectoryTransactor implements the ContractAVSDirectoryTransacts interface.
var _ ContractAVSDirectoryTransacts = (*ContractAVSDirectoryTransactor)(nil)

// ContractAVSDirectoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractAVSDirectoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractAVSDirectoryFilterer implements the ContractAVSDirectoryFilters interface.
var _ ContractAVSDirectoryFilters = (*ContractAVSDirectoryFilterer)(nil)

// ContractAVSDirectorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractAVSDirectorySession struct {
	Contract     *ContractAVSDirectory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ContractAVSDirectoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractAVSDirectoryCallerSession struct {
	Contract *ContractAVSDirectoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// ContractAVSDirectoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractAVSDirectoryTransactorSession struct {
	Contract     *ContractAVSDirectoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// ContractAVSDirectoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractAVSDirectoryRaw struct {
	Contract *ContractAVSDirectory // Generic contract binding to access the raw methods on
}

// ContractAVSDirectoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractAVSDirectoryCallerRaw struct {
	Contract *ContractAVSDirectoryCaller // Generic read-only contract binding to access the raw methods on
}

// ContractAVSDirectoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractAVSDirectoryTransactorRaw struct {
	Contract *ContractAVSDirectoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractAVSDirectory creates a new instance of ContractAVSDirectory, bound to a specific deployed contract.
func NewContractAVSDirectory(address common.Address, backend bind.ContractBackend) (*ContractAVSDirectory, error) {
	contract, err := bindContractAVSDirectory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractAVSDirectory{ContractAVSDirectoryCaller: ContractAVSDirectoryCaller{contract: contract}, ContractAVSDirectoryTransactor: ContractAVSDirectoryTransactor{contract: contract}, ContractAVSDirectoryFilterer: ContractAVSDirectoryFilterer{contract: contract}}, nil
}

// NewContractAVSDirectoryCaller creates a new read-only instance of ContractAVSDirectory, bound to a specific deployed contract.
func NewContractAVSDirectoryCaller(address common.Address, caller bind.ContractCaller) (*ContractAVSDirectoryCaller, error) {
	contract, err := bindContractAVSDirectory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractAVSDirectoryCaller{contract: contract}, nil
}

// NewContractAVSDirectoryTransactor creates a new write-only instance of ContractAVSDirectory, bound to a specific deployed contract.
func NewContractAVSDirectoryTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractAVSDirectoryTransactor, error) {
	contract, err := bindContractAVSDirectory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractAVSDirectoryTransactor{contract: contract}, nil
}

// NewContractAVSDirectoryFilterer creates a new log filterer instance of ContractAVSDirectory, bound to a specific deployed contract.
func NewContractAVSDirectoryFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractAVSDirectoryFilterer, error) {
	contract, err := bindContractAVSDirectory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractAVSDirectoryFilterer{contract: contract}, nil
}

// bindContractAVSDirectory binds a generic wrapper to an already deployed contract.
func bindContractAVSDirectory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractAVSDirectoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractAVSDirectory *ContractAVSDirectoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractAVSDirectory.Contract.ContractAVSDirectoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractAVSDirectory *ContractAVSDirectoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractAVSDirectory.Contract.ContractAVSDirectoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractAVSDirectory *ContractAVSDirectoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractAVSDirectory.Contract.ContractAVSDirectoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractAVSDirectory *ContractAVSDirectoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractAVSDirectory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractAVSDirectory *ContractAVSDirectoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractAVSDirectory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractAVSDirectory *ContractAVSDirectoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractAVSDirectory.Contract.contract.Transact(opts, method, params...)
}

// DEALLOCATIONDELAY is a free data retrieval call binding the contract method 0x2981eb77.
//
// Solidity: function DEALLOCATION_DELAY() view returns(uint32)
func (_ContractAVSDirectory *ContractAVSDirectoryCaller) DEALLOCATIONDELAY(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractAVSDirectory.contract.Call(opts, &out, "DEALLOCATION_DELAY")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// DEALLOCATIONDELAY is a free data retrieval call binding the contract method 0x2981eb77.
//
// Solidity: function DEALLOCATION_DELAY() view returns(uint32)
func (_ContractAVSDirectory *ContractAVSDirectorySession) DEALLOCATIONDELAY() (uint32, error) {
	return _ContractAVSDirectory.Contract.DEALLOCATIONDELAY(&_ContractAVSDirectory.CallOpts)
}

// DEALLOCATIONDELAY is a free data retrieval call binding the contract method 0x2981eb77.
//
// Solidity: function DEALLOCATION_DELAY() view returns(uint32)
func (_ContractAVSDirectory *ContractAVSDirectoryCallerSession) DEALLOCATIONDELAY() (uint32, error) {
	return _ContractAVSDirectory.Contract.DEALLOCATIONDELAY(&_ContractAVSDirectory.CallOpts)
}

// OPERATORAVSREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0xd79aceab.
//
// Solidity: function OPERATOR_AVS_REGISTRATION_TYPEHASH() view returns(bytes32)
func (_ContractAVSDirectory *ContractAVSDirectoryCaller) OPERATORAVSREGISTRATIONTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ContractAVSDirectory.contract.Call(opts, &out, "OPERATOR_AVS_REGISTRATION_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OPERATORAVSREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0xd79aceab.
//
// Solidity: function OPERATOR_AVS_REGISTRATION_TYPEHASH() view returns(bytes32)
func (_ContractAVSDirectory *ContractAVSDirectorySession) OPERATORAVSREGISTRATIONTYPEHASH() ([32]byte, error) {
	return _ContractAVSDirectory.Contract.OPERATORAVSREGISTRATIONTYPEHASH(&_ContractAVSDirectory.CallOpts)
}

// OPERATORAVSREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0xd79aceab.
//
// Solidity: function OPERATOR_AVS_REGISTRATION_TYPEHASH() view returns(bytes32)
func (_ContractAVSDirectory *ContractAVSDirectoryCallerSession) OPERATORAVSREGISTRATIONTYPEHASH() ([32]byte, error) {
	return _ContractAVSDirectory.Contract.OPERATORAVSREGISTRATIONTYPEHASH(&_ContractAVSDirectory.CallOpts)
}

// OPERATORSETFORCEDEREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0xdce974b9.
//
// Solidity: function OPERATOR_SET_FORCE_DEREGISTRATION_TYPEHASH() view returns(bytes32)
func (_ContractAVSDirectory *ContractAVSDirectoryCaller) OPERATORSETFORCEDEREGISTRATIONTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ContractAVSDirectory.contract.Call(opts, &out, "OPERATOR_SET_FORCE_DEREGISTRATION_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OPERATORSETFORCEDEREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0xdce974b9.
//
// Solidity: function OPERATOR_SET_FORCE_DEREGISTRATION_TYPEHASH() view returns(bytes32)
func (_ContractAVSDirectory *ContractAVSDirectorySession) OPERATORSETFORCEDEREGISTRATIONTYPEHASH() ([32]byte, error) {
	return _ContractAVSDirectory.Contract.OPERATORSETFORCEDEREGISTRATIONTYPEHASH(&_ContractAVSDirectory.CallOpts)
}

// OPERATORSETFORCEDEREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0xdce974b9.
//
// Solidity: function OPERATOR_SET_FORCE_DEREGISTRATION_TYPEHASH() view returns(bytes32)
func (_ContractAVSDirectory *ContractAVSDirectoryCallerSession) OPERATORSETFORCEDEREGISTRATIONTYPEHASH() ([32]byte, error) {
	return _ContractAVSDirectory.Contract.OPERATORSETFORCEDEREGISTRATIONTYPEHASH(&_ContractAVSDirectory.CallOpts)
}

// OPERATORSETREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0xc825fe68.
//
// Solidity: function OPERATOR_SET_REGISTRATION_TYPEHASH() view returns(bytes32)
func (_ContractAVSDirectory *ContractAVSDirectoryCaller) OPERATORSETREGISTRATIONTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ContractAVSDirectory.contract.Call(opts, &out, "OPERATOR_SET_REGISTRATION_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OPERATORSETREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0xc825fe68.
//
// Solidity: function OPERATOR_SET_REGISTRATION_TYPEHASH() view returns(bytes32)
func (_ContractAVSDirectory *ContractAVSDirectorySession) OPERATORSETREGISTRATIONTYPEHASH() ([32]byte, error) {
	return _ContractAVSDirectory.Contract.OPERATORSETREGISTRATIONTYPEHASH(&_ContractAVSDirectory.CallOpts)
}

// OPERATORSETREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0xc825fe68.
//
// Solidity: function OPERATOR_SET_REGISTRATION_TYPEHASH() view returns(bytes32)
func (_ContractAVSDirectory *ContractAVSDirectoryCallerSession) OPERATORSETREGISTRATIONTYPEHASH() ([32]byte, error) {
	return _ContractAVSDirectory.Contract.OPERATORSETREGISTRATIONTYPEHASH(&_ContractAVSDirectory.CallOpts)
}

// AvsOperatorStatus is a free data retrieval call binding the contract method 0x49075da3.
//
// Solidity: function avsOperatorStatus(address avs, address operator) view returns(uint8)
func (_ContractAVSDirectory *ContractAVSDirectoryCaller) AvsOperatorStatus(opts *bind.CallOpts, avs common.Address, operator common.Address) (uint8, error) {
	var out []interface{}
	err := _ContractAVSDirectory.contract.Call(opts, &out, "avsOperatorStatus", avs, operator)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// AvsOperatorStatus is a free data retrieval call binding the contract method 0x49075da3.
//
// Solidity: function avsOperatorStatus(address avs, address operator) view returns(uint8)
func (_ContractAVSDirectory *ContractAVSDirectorySession) AvsOperatorStatus(avs common.Address, operator common.Address) (uint8, error) {
	return _ContractAVSDirectory.Contract.AvsOperatorStatus(&_ContractAVSDirectory.CallOpts, avs, operator)
}

// AvsOperatorStatus is a free data retrieval call binding the contract method 0x49075da3.
//
// Solidity: function avsOperatorStatus(address avs, address operator) view returns(uint8)
func (_ContractAVSDirectory *ContractAVSDirectoryCallerSession) AvsOperatorStatus(avs common.Address, operator common.Address) (uint8, error) {
	return _ContractAVSDirectory.Contract.AvsOperatorStatus(&_ContractAVSDirectory.CallOpts, avs, operator)
}

// CalculateOperatorAVSRegistrationDigestHash is a free data retrieval call binding the contract method 0xa1060c88.
//
// Solidity: function calculateOperatorAVSRegistrationDigestHash(address operator, address avs, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_ContractAVSDirectory *ContractAVSDirectoryCaller) CalculateOperatorAVSRegistrationDigestHash(opts *bind.CallOpts, operator common.Address, avs common.Address, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _ContractAVSDirectory.contract.Call(opts, &out, "calculateOperatorAVSRegistrationDigestHash", operator, avs, salt, expiry)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateOperatorAVSRegistrationDigestHash is a free data retrieval call binding the contract method 0xa1060c88.
//
// Solidity: function calculateOperatorAVSRegistrationDigestHash(address operator, address avs, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_ContractAVSDirectory *ContractAVSDirectorySession) CalculateOperatorAVSRegistrationDigestHash(operator common.Address, avs common.Address, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _ContractAVSDirectory.Contract.CalculateOperatorAVSRegistrationDigestHash(&_ContractAVSDirectory.CallOpts, operator, avs, salt, expiry)
}

// CalculateOperatorAVSRegistrationDigestHash is a free data retrieval call binding the contract method 0xa1060c88.
//
// Solidity: function calculateOperatorAVSRegistrationDigestHash(address operator, address avs, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_ContractAVSDirectory *ContractAVSDirectoryCallerSession) CalculateOperatorAVSRegistrationDigestHash(operator common.Address, avs common.Address, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _ContractAVSDirectory.Contract.CalculateOperatorAVSRegistrationDigestHash(&_ContractAVSDirectory.CallOpts, operator, avs, salt, expiry)
}

// CalculateOperatorSetForceDeregistrationTypehash is a free data retrieval call binding the contract method 0xb2841d48.
//
// Solidity: function calculateOperatorSetForceDeregistrationTypehash(address avs, uint32[] operatorSetIds, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_ContractAVSDirectory *ContractAVSDirectoryCaller) CalculateOperatorSetForceDeregistrationTypehash(opts *bind.CallOpts, avs common.Address, operatorSetIds []uint32, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _ContractAVSDirectory.contract.Call(opts, &out, "calculateOperatorSetForceDeregistrationTypehash", avs, operatorSetIds, salt, expiry)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateOperatorSetForceDeregistrationTypehash is a free data retrieval call binding the contract method 0xb2841d48.
//
// Solidity: function calculateOperatorSetForceDeregistrationTypehash(address avs, uint32[] operatorSetIds, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_ContractAVSDirectory *ContractAVSDirectorySession) CalculateOperatorSetForceDeregistrationTypehash(avs common.Address, operatorSetIds []uint32, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _ContractAVSDirectory.Contract.CalculateOperatorSetForceDeregistrationTypehash(&_ContractAVSDirectory.CallOpts, avs, operatorSetIds, salt, expiry)
}

// CalculateOperatorSetForceDeregistrationTypehash is a free data retrieval call binding the contract method 0xb2841d48.
//
// Solidity: function calculateOperatorSetForceDeregistrationTypehash(address avs, uint32[] operatorSetIds, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_ContractAVSDirectory *ContractAVSDirectoryCallerSession) CalculateOperatorSetForceDeregistrationTypehash(avs common.Address, operatorSetIds []uint32, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _ContractAVSDirectory.Contract.CalculateOperatorSetForceDeregistrationTypehash(&_ContractAVSDirectory.CallOpts, avs, operatorSetIds, salt, expiry)
}

// CalculateOperatorSetRegistrationDigestHash is a free data retrieval call binding the contract method 0x955e6696.
//
// Solidity: function calculateOperatorSetRegistrationDigestHash(address avs, uint32[] operatorSetIds, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_ContractAVSDirectory *ContractAVSDirectoryCaller) CalculateOperatorSetRegistrationDigestHash(opts *bind.CallOpts, avs common.Address, operatorSetIds []uint32, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _ContractAVSDirectory.contract.Call(opts, &out, "calculateOperatorSetRegistrationDigestHash", avs, operatorSetIds, salt, expiry)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateOperatorSetRegistrationDigestHash is a free data retrieval call binding the contract method 0x955e6696.
//
// Solidity: function calculateOperatorSetRegistrationDigestHash(address avs, uint32[] operatorSetIds, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_ContractAVSDirectory *ContractAVSDirectorySession) CalculateOperatorSetRegistrationDigestHash(avs common.Address, operatorSetIds []uint32, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _ContractAVSDirectory.Contract.CalculateOperatorSetRegistrationDigestHash(&_ContractAVSDirectory.CallOpts, avs, operatorSetIds, salt, expiry)
}

// CalculateOperatorSetRegistrationDigestHash is a free data retrieval call binding the contract method 0x955e6696.
//
// Solidity: function calculateOperatorSetRegistrationDigestHash(address avs, uint32[] operatorSetIds, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_ContractAVSDirectory *ContractAVSDirectoryCallerSession) CalculateOperatorSetRegistrationDigestHash(avs common.Address, operatorSetIds []uint32, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _ContractAVSDirectory.Contract.CalculateOperatorSetRegistrationDigestHash(&_ContractAVSDirectory.CallOpts, avs, operatorSetIds, salt, expiry)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractAVSDirectory *ContractAVSDirectoryCaller) Delegation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractAVSDirectory.contract.Call(opts, &out, "delegation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractAVSDirectory *ContractAVSDirectorySession) Delegation() (common.Address, error) {
	return _ContractAVSDirectory.Contract.Delegation(&_ContractAVSDirectory.CallOpts)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractAVSDirectory *ContractAVSDirectoryCallerSession) Delegation() (common.Address, error) {
	return _ContractAVSDirectory.Contract.Delegation(&_ContractAVSDirectory.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_ContractAVSDirectory *ContractAVSDirectoryCaller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ContractAVSDirectory.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_ContractAVSDirectory *ContractAVSDirectorySession) DomainSeparator() ([32]byte, error) {
	return _ContractAVSDirectory.Contract.DomainSeparator(&_ContractAVSDirectory.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_ContractAVSDirectory *ContractAVSDirectoryCallerSession) DomainSeparator() ([32]byte, error) {
	return _ContractAVSDirectory.Contract.DomainSeparator(&_ContractAVSDirectory.CallOpts)
}

// GetNumOperatorSetsOfOperator is a free data retrieval call binding the contract method 0xe88d8049.
//
// Solidity: function getNumOperatorSetsOfOperator(address operator) view returns(uint256)
func (_ContractAVSDirectory *ContractAVSDirectoryCaller) GetNumOperatorSetsOfOperator(opts *bind.CallOpts, operator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ContractAVSDirectory.contract.Call(opts, &out, "getNumOperatorSetsOfOperator", operator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNumOperatorSetsOfOperator is a free data retrieval call binding the contract method 0xe88d8049.
//
// Solidity: function getNumOperatorSetsOfOperator(address operator) view returns(uint256)
func (_ContractAVSDirectory *ContractAVSDirectorySession) GetNumOperatorSetsOfOperator(operator common.Address) (*big.Int, error) {
	return _ContractAVSDirectory.Contract.GetNumOperatorSetsOfOperator(&_ContractAVSDirectory.CallOpts, operator)
}

// GetNumOperatorSetsOfOperator is a free data retrieval call binding the contract method 0xe88d8049.
//
// Solidity: function getNumOperatorSetsOfOperator(address operator) view returns(uint256)
func (_ContractAVSDirectory *ContractAVSDirectoryCallerSession) GetNumOperatorSetsOfOperator(operator common.Address) (*big.Int, error) {
	return _ContractAVSDirectory.Contract.GetNumOperatorSetsOfOperator(&_ContractAVSDirectory.CallOpts, operator)
}

// GetNumOperatorsInOperatorSet is a free data retrieval call binding the contract method 0x1023aa35.
//
// Solidity: function getNumOperatorsInOperatorSet((address,uint32) operatorSet) view returns(uint256)
func (_ContractAVSDirectory *ContractAVSDirectoryCaller) GetNumOperatorsInOperatorSet(opts *bind.CallOpts, operatorSet OperatorSet) (*big.Int, error) {
	var out []interface{}
	err := _ContractAVSDirectory.contract.Call(opts, &out, "getNumOperatorsInOperatorSet", operatorSet)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNumOperatorsInOperatorSet is a free data retrieval call binding the contract method 0x1023aa35.
//
// Solidity: function getNumOperatorsInOperatorSet((address,uint32) operatorSet) view returns(uint256)
func (_ContractAVSDirectory *ContractAVSDirectorySession) GetNumOperatorsInOperatorSet(operatorSet OperatorSet) (*big.Int, error) {
	return _ContractAVSDirectory.Contract.GetNumOperatorsInOperatorSet(&_ContractAVSDirectory.CallOpts, operatorSet)
}

// GetNumOperatorsInOperatorSet is a free data retrieval call binding the contract method 0x1023aa35.
//
// Solidity: function getNumOperatorsInOperatorSet((address,uint32) operatorSet) view returns(uint256)
func (_ContractAVSDirectory *ContractAVSDirectoryCallerSession) GetNumOperatorsInOperatorSet(operatorSet OperatorSet) (*big.Int, error) {
	return _ContractAVSDirectory.Contract.GetNumOperatorsInOperatorSet(&_ContractAVSDirectory.CallOpts, operatorSet)
}

// GetOperatorSetsOfOperator is a free data retrieval call binding the contract method 0x16ae76cb.
//
// Solidity: function getOperatorSetsOfOperator(address operator, uint256 start, uint256 length) view returns((address,uint32)[] operatorSets)
func (_ContractAVSDirectory *ContractAVSDirectoryCaller) GetOperatorSetsOfOperator(opts *bind.CallOpts, operator common.Address, start *big.Int, length *big.Int) ([]OperatorSet, error) {
	var out []interface{}
	err := _ContractAVSDirectory.contract.Call(opts, &out, "getOperatorSetsOfOperator", operator, start, length)

	if err != nil {
		return *new([]OperatorSet), err
	}

	out0 := *abi.ConvertType(out[0], new([]OperatorSet)).(*[]OperatorSet)

	return out0, err

}

// GetOperatorSetsOfOperator is a free data retrieval call binding the contract method 0x16ae76cb.
//
// Solidity: function getOperatorSetsOfOperator(address operator, uint256 start, uint256 length) view returns((address,uint32)[] operatorSets)
func (_ContractAVSDirectory *ContractAVSDirectorySession) GetOperatorSetsOfOperator(operator common.Address, start *big.Int, length *big.Int) ([]OperatorSet, error) {
	return _ContractAVSDirectory.Contract.GetOperatorSetsOfOperator(&_ContractAVSDirectory.CallOpts, operator, start, length)
}

// GetOperatorSetsOfOperator is a free data retrieval call binding the contract method 0x16ae76cb.
//
// Solidity: function getOperatorSetsOfOperator(address operator, uint256 start, uint256 length) view returns((address,uint32)[] operatorSets)
func (_ContractAVSDirectory *ContractAVSDirectoryCallerSession) GetOperatorSetsOfOperator(operator common.Address, start *big.Int, length *big.Int) ([]OperatorSet, error) {
	return _ContractAVSDirectory.Contract.GetOperatorSetsOfOperator(&_ContractAVSDirectory.CallOpts, operator, start, length)
}

// GetOperatorsInOperatorSet is a free data retrieval call binding the contract method 0x7357723b.
//
// Solidity: function getOperatorsInOperatorSet((address,uint32) operatorSet, uint256 start, uint256 length) view returns(address[] operators)
func (_ContractAVSDirectory *ContractAVSDirectoryCaller) GetOperatorsInOperatorSet(opts *bind.CallOpts, operatorSet OperatorSet, start *big.Int, length *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _ContractAVSDirectory.contract.Call(opts, &out, "getOperatorsInOperatorSet", operatorSet, start, length)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOperatorsInOperatorSet is a free data retrieval call binding the contract method 0x7357723b.
//
// Solidity: function getOperatorsInOperatorSet((address,uint32) operatorSet, uint256 start, uint256 length) view returns(address[] operators)
func (_ContractAVSDirectory *ContractAVSDirectorySession) GetOperatorsInOperatorSet(operatorSet OperatorSet, start *big.Int, length *big.Int) ([]common.Address, error) {
	return _ContractAVSDirectory.Contract.GetOperatorsInOperatorSet(&_ContractAVSDirectory.CallOpts, operatorSet, start, length)
}

// GetOperatorsInOperatorSet is a free data retrieval call binding the contract method 0x7357723b.
//
// Solidity: function getOperatorsInOperatorSet((address,uint32) operatorSet, uint256 start, uint256 length) view returns(address[] operators)
func (_ContractAVSDirectory *ContractAVSDirectoryCallerSession) GetOperatorsInOperatorSet(operatorSet OperatorSet, start *big.Int, length *big.Int) ([]common.Address, error) {
	return _ContractAVSDirectory.Contract.GetOperatorsInOperatorSet(&_ContractAVSDirectory.CallOpts, operatorSet, start, length)
}

// GetStrategiesInOperatorSet is a free data retrieval call binding the contract method 0x4177a87c.
//
// Solidity: function getStrategiesInOperatorSet((address,uint32) operatorSet) view returns(address[] strategies)
func (_ContractAVSDirectory *ContractAVSDirectoryCaller) GetStrategiesInOperatorSet(opts *bind.CallOpts, operatorSet OperatorSet) ([]common.Address, error) {
	var out []interface{}
	err := _ContractAVSDirectory.contract.Call(opts, &out, "getStrategiesInOperatorSet", operatorSet)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetStrategiesInOperatorSet is a free data retrieval call binding the contract method 0x4177a87c.
//
// Solidity: function getStrategiesInOperatorSet((address,uint32) operatorSet) view returns(address[] strategies)
func (_ContractAVSDirectory *ContractAVSDirectorySession) GetStrategiesInOperatorSet(operatorSet OperatorSet) ([]common.Address, error) {
	return _ContractAVSDirectory.Contract.GetStrategiesInOperatorSet(&_ContractAVSDirectory.CallOpts, operatorSet)
}

// GetStrategiesInOperatorSet is a free data retrieval call binding the contract method 0x4177a87c.
//
// Solidity: function getStrategiesInOperatorSet((address,uint32) operatorSet) view returns(address[] strategies)
func (_ContractAVSDirectory *ContractAVSDirectoryCallerSession) GetStrategiesInOperatorSet(operatorSet OperatorSet) ([]common.Address, error) {
	return _ContractAVSDirectory.Contract.GetStrategiesInOperatorSet(&_ContractAVSDirectory.CallOpts, operatorSet)
}

// InTotalOperatorSets is a free data retrieval call binding the contract method 0xcbdf0e42.
//
// Solidity: function inTotalOperatorSets(address operator) view returns(uint256)
func (_ContractAVSDirectory *ContractAVSDirectoryCaller) InTotalOperatorSets(opts *bind.CallOpts, operator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ContractAVSDirectory.contract.Call(opts, &out, "inTotalOperatorSets", operator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InTotalOperatorSets is a free data retrieval call binding the contract method 0xcbdf0e42.
//
// Solidity: function inTotalOperatorSets(address operator) view returns(uint256)
func (_ContractAVSDirectory *ContractAVSDirectorySession) InTotalOperatorSets(operator common.Address) (*big.Int, error) {
	return _ContractAVSDirectory.Contract.InTotalOperatorSets(&_ContractAVSDirectory.CallOpts, operator)
}

// InTotalOperatorSets is a free data retrieval call binding the contract method 0xcbdf0e42.
//
// Solidity: function inTotalOperatorSets(address operator) view returns(uint256)
func (_ContractAVSDirectory *ContractAVSDirectoryCallerSession) InTotalOperatorSets(operator common.Address) (*big.Int, error) {
	return _ContractAVSDirectory.Contract.InTotalOperatorSets(&_ContractAVSDirectory.CallOpts, operator)
}

// IsMember is a free data retrieval call binding the contract method 0xda2ff05d.
//
// Solidity: function isMember(address operator, (address,uint32) operatorSet) view returns(bool)
func (_ContractAVSDirectory *ContractAVSDirectoryCaller) IsMember(opts *bind.CallOpts, operator common.Address, operatorSet OperatorSet) (bool, error) {
	var out []interface{}
	err := _ContractAVSDirectory.contract.Call(opts, &out, "isMember", operator, operatorSet)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMember is a free data retrieval call binding the contract method 0xda2ff05d.
//
// Solidity: function isMember(address operator, (address,uint32) operatorSet) view returns(bool)
func (_ContractAVSDirectory *ContractAVSDirectorySession) IsMember(operator common.Address, operatorSet OperatorSet) (bool, error) {
	return _ContractAVSDirectory.Contract.IsMember(&_ContractAVSDirectory.CallOpts, operator, operatorSet)
}

// IsMember is a free data retrieval call binding the contract method 0xda2ff05d.
//
// Solidity: function isMember(address operator, (address,uint32) operatorSet) view returns(bool)
func (_ContractAVSDirectory *ContractAVSDirectoryCallerSession) IsMember(operator common.Address, operatorSet OperatorSet) (bool, error) {
	return _ContractAVSDirectory.Contract.IsMember(&_ContractAVSDirectory.CallOpts, operator, operatorSet)
}

// IsOperatorSet is a free data retrieval call binding the contract method 0x84d76f7b.
//
// Solidity: function isOperatorSet(address avs, uint32 operatorSetId) view returns(bool)
func (_ContractAVSDirectory *ContractAVSDirectoryCaller) IsOperatorSet(opts *bind.CallOpts, avs common.Address, operatorSetId uint32) (bool, error) {
	var out []interface{}
	err := _ContractAVSDirectory.contract.Call(opts, &out, "isOperatorSet", avs, operatorSetId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperatorSet is a free data retrieval call binding the contract method 0x84d76f7b.
//
// Solidity: function isOperatorSet(address avs, uint32 operatorSetId) view returns(bool)
func (_ContractAVSDirectory *ContractAVSDirectorySession) IsOperatorSet(avs common.Address, operatorSetId uint32) (bool, error) {
	return _ContractAVSDirectory.Contract.IsOperatorSet(&_ContractAVSDirectory.CallOpts, avs, operatorSetId)
}

// IsOperatorSet is a free data retrieval call binding the contract method 0x84d76f7b.
//
// Solidity: function isOperatorSet(address avs, uint32 operatorSetId) view returns(bool)
func (_ContractAVSDirectory *ContractAVSDirectoryCallerSession) IsOperatorSet(avs common.Address, operatorSetId uint32) (bool, error) {
	return _ContractAVSDirectory.Contract.IsOperatorSet(&_ContractAVSDirectory.CallOpts, avs, operatorSetId)
}

// IsOperatorSetAVS is a free data retrieval call binding the contract method 0x7673e93a.
//
// Solidity: function isOperatorSetAVS(address avs) view returns(bool)
func (_ContractAVSDirectory *ContractAVSDirectoryCaller) IsOperatorSetAVS(opts *bind.CallOpts, avs common.Address) (bool, error) {
	var out []interface{}
	err := _ContractAVSDirectory.contract.Call(opts, &out, "isOperatorSetAVS", avs)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperatorSetAVS is a free data retrieval call binding the contract method 0x7673e93a.
//
// Solidity: function isOperatorSetAVS(address avs) view returns(bool)
func (_ContractAVSDirectory *ContractAVSDirectorySession) IsOperatorSetAVS(avs common.Address) (bool, error) {
	return _ContractAVSDirectory.Contract.IsOperatorSetAVS(&_ContractAVSDirectory.CallOpts, avs)
}

// IsOperatorSetAVS is a free data retrieval call binding the contract method 0x7673e93a.
//
// Solidity: function isOperatorSetAVS(address avs) view returns(bool)
func (_ContractAVSDirectory *ContractAVSDirectoryCallerSession) IsOperatorSetAVS(avs common.Address) (bool, error) {
	return _ContractAVSDirectory.Contract.IsOperatorSetAVS(&_ContractAVSDirectory.CallOpts, avs)
}

// IsOperatorSetBatch is a free data retrieval call binding the contract method 0x20c4e236.
//
// Solidity: function isOperatorSetBatch((address,uint32)[] operatorSets) view returns(bool)
func (_ContractAVSDirectory *ContractAVSDirectoryCaller) IsOperatorSetBatch(opts *bind.CallOpts, operatorSets []OperatorSet) (bool, error) {
	var out []interface{}
	err := _ContractAVSDirectory.contract.Call(opts, &out, "isOperatorSetBatch", operatorSets)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperatorSetBatch is a free data retrieval call binding the contract method 0x20c4e236.
//
// Solidity: function isOperatorSetBatch((address,uint32)[] operatorSets) view returns(bool)
func (_ContractAVSDirectory *ContractAVSDirectorySession) IsOperatorSetBatch(operatorSets []OperatorSet) (bool, error) {
	return _ContractAVSDirectory.Contract.IsOperatorSetBatch(&_ContractAVSDirectory.CallOpts, operatorSets)
}

// IsOperatorSetBatch is a free data retrieval call binding the contract method 0x20c4e236.
//
// Solidity: function isOperatorSetBatch((address,uint32)[] operatorSets) view returns(bool)
func (_ContractAVSDirectory *ContractAVSDirectoryCallerSession) IsOperatorSetBatch(operatorSets []OperatorSet) (bool, error) {
	return _ContractAVSDirectory.Contract.IsOperatorSetBatch(&_ContractAVSDirectory.CallOpts, operatorSets)
}

// IsOperatorSlashable is a free data retrieval call binding the contract method 0x1352c3e6.
//
// Solidity: function isOperatorSlashable(address operator, (address,uint32) operatorSet) view returns(bool)
func (_ContractAVSDirectory *ContractAVSDirectoryCaller) IsOperatorSlashable(opts *bind.CallOpts, operator common.Address, operatorSet OperatorSet) (bool, error) {
	var out []interface{}
	err := _ContractAVSDirectory.contract.Call(opts, &out, "isOperatorSlashable", operator, operatorSet)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperatorSlashable is a free data retrieval call binding the contract method 0x1352c3e6.
//
// Solidity: function isOperatorSlashable(address operator, (address,uint32) operatorSet) view returns(bool)
func (_ContractAVSDirectory *ContractAVSDirectorySession) IsOperatorSlashable(operator common.Address, operatorSet OperatorSet) (bool, error) {
	return _ContractAVSDirectory.Contract.IsOperatorSlashable(&_ContractAVSDirectory.CallOpts, operator, operatorSet)
}

// IsOperatorSlashable is a free data retrieval call binding the contract method 0x1352c3e6.
//
// Solidity: function isOperatorSlashable(address operator, (address,uint32) operatorSet) view returns(bool)
func (_ContractAVSDirectory *ContractAVSDirectoryCallerSession) IsOperatorSlashable(operator common.Address, operatorSet OperatorSet) (bool, error) {
	return _ContractAVSDirectory.Contract.IsOperatorSlashable(&_ContractAVSDirectory.CallOpts, operator, operatorSet)
}

// OperatorSaltIsSpent is a free data retrieval call binding the contract method 0x374823b5.
//
// Solidity: function operatorSaltIsSpent(address operator, bytes32 salt) view returns(bool isSpent)
func (_ContractAVSDirectory *ContractAVSDirectoryCaller) OperatorSaltIsSpent(opts *bind.CallOpts, operator common.Address, salt [32]byte) (bool, error) {
	var out []interface{}
	err := _ContractAVSDirectory.contract.Call(opts, &out, "operatorSaltIsSpent", operator, salt)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OperatorSaltIsSpent is a free data retrieval call binding the contract method 0x374823b5.
//
// Solidity: function operatorSaltIsSpent(address operator, bytes32 salt) view returns(bool isSpent)
func (_ContractAVSDirectory *ContractAVSDirectorySession) OperatorSaltIsSpent(operator common.Address, salt [32]byte) (bool, error) {
	return _ContractAVSDirectory.Contract.OperatorSaltIsSpent(&_ContractAVSDirectory.CallOpts, operator, salt)
}

// OperatorSaltIsSpent is a free data retrieval call binding the contract method 0x374823b5.
//
// Solidity: function operatorSaltIsSpent(address operator, bytes32 salt) view returns(bool isSpent)
func (_ContractAVSDirectory *ContractAVSDirectoryCallerSession) OperatorSaltIsSpent(operator common.Address, salt [32]byte) (bool, error) {
	return _ContractAVSDirectory.Contract.OperatorSaltIsSpent(&_ContractAVSDirectory.CallOpts, operator, salt)
}

// OperatorSetMemberAtIndex is a free data retrieval call binding the contract method 0x411d415b.
//
// Solidity: function operatorSetMemberAtIndex((address,uint32) operatorSet, uint256 index) view returns(address)
func (_ContractAVSDirectory *ContractAVSDirectoryCaller) OperatorSetMemberAtIndex(opts *bind.CallOpts, operatorSet OperatorSet, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ContractAVSDirectory.contract.Call(opts, &out, "operatorSetMemberAtIndex", operatorSet, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OperatorSetMemberAtIndex is a free data retrieval call binding the contract method 0x411d415b.
//
// Solidity: function operatorSetMemberAtIndex((address,uint32) operatorSet, uint256 index) view returns(address)
func (_ContractAVSDirectory *ContractAVSDirectorySession) OperatorSetMemberAtIndex(operatorSet OperatorSet, index *big.Int) (common.Address, error) {
	return _ContractAVSDirectory.Contract.OperatorSetMemberAtIndex(&_ContractAVSDirectory.CallOpts, operatorSet, index)
}

// OperatorSetMemberAtIndex is a free data retrieval call binding the contract method 0x411d415b.
//
// Solidity: function operatorSetMemberAtIndex((address,uint32) operatorSet, uint256 index) view returns(address)
func (_ContractAVSDirectory *ContractAVSDirectoryCallerSession) OperatorSetMemberAtIndex(operatorSet OperatorSet, index *big.Int) (common.Address, error) {
	return _ContractAVSDirectory.Contract.OperatorSetMemberAtIndex(&_ContractAVSDirectory.CallOpts, operatorSet, index)
}

// OperatorSetStatus is a free data retrieval call binding the contract method 0x1e68134e.
//
// Solidity: function operatorSetStatus(address operator, address avs, uint32 operatorSetId) view returns(bool registered, uint32 lastDeregisteredTimestamp)
func (_ContractAVSDirectory *ContractAVSDirectoryCaller) OperatorSetStatus(opts *bind.CallOpts, operator common.Address, avs common.Address, operatorSetId uint32) (struct {
	Registered                bool
	LastDeregisteredTimestamp uint32
}, error) {
	var out []interface{}
	err := _ContractAVSDirectory.contract.Call(opts, &out, "operatorSetStatus", operator, avs, operatorSetId)

	outstruct := new(struct {
		Registered                bool
		LastDeregisteredTimestamp uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Registered = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.LastDeregisteredTimestamp = *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return *outstruct, err

}

// OperatorSetStatus is a free data retrieval call binding the contract method 0x1e68134e.
//
// Solidity: function operatorSetStatus(address operator, address avs, uint32 operatorSetId) view returns(bool registered, uint32 lastDeregisteredTimestamp)
func (_ContractAVSDirectory *ContractAVSDirectorySession) OperatorSetStatus(operator common.Address, avs common.Address, operatorSetId uint32) (struct {
	Registered                bool
	LastDeregisteredTimestamp uint32
}, error) {
	return _ContractAVSDirectory.Contract.OperatorSetStatus(&_ContractAVSDirectory.CallOpts, operator, avs, operatorSetId)
}

// OperatorSetStatus is a free data retrieval call binding the contract method 0x1e68134e.
//
// Solidity: function operatorSetStatus(address operator, address avs, uint32 operatorSetId) view returns(bool registered, uint32 lastDeregisteredTimestamp)
func (_ContractAVSDirectory *ContractAVSDirectoryCallerSession) OperatorSetStatus(operator common.Address, avs common.Address, operatorSetId uint32) (struct {
	Registered                bool
	LastDeregisteredTimestamp uint32
}, error) {
	return _ContractAVSDirectory.Contract.OperatorSetStatus(&_ContractAVSDirectory.CallOpts, operator, avs, operatorSetId)
}

// OperatorSetsMemberOfAtIndex is a free data retrieval call binding the contract method 0xb5a768ca.
//
// Solidity: function operatorSetsMemberOfAtIndex(address operator, uint256 index) view returns((address,uint32))
func (_ContractAVSDirectory *ContractAVSDirectoryCaller) OperatorSetsMemberOfAtIndex(opts *bind.CallOpts, operator common.Address, index *big.Int) (OperatorSet, error) {
	var out []interface{}
	err := _ContractAVSDirectory.contract.Call(opts, &out, "operatorSetsMemberOfAtIndex", operator, index)

	if err != nil {
		return *new(OperatorSet), err
	}

	out0 := *abi.ConvertType(out[0], new(OperatorSet)).(*OperatorSet)

	return out0, err

}

// OperatorSetsMemberOfAtIndex is a free data retrieval call binding the contract method 0xb5a768ca.
//
// Solidity: function operatorSetsMemberOfAtIndex(address operator, uint256 index) view returns((address,uint32))
func (_ContractAVSDirectory *ContractAVSDirectorySession) OperatorSetsMemberOfAtIndex(operator common.Address, index *big.Int) (OperatorSet, error) {
	return _ContractAVSDirectory.Contract.OperatorSetsMemberOfAtIndex(&_ContractAVSDirectory.CallOpts, operator, index)
}

// OperatorSetsMemberOfAtIndex is a free data retrieval call binding the contract method 0xb5a768ca.
//
// Solidity: function operatorSetsMemberOfAtIndex(address operator, uint256 index) view returns((address,uint32))
func (_ContractAVSDirectory *ContractAVSDirectoryCallerSession) OperatorSetsMemberOfAtIndex(operator common.Address, index *big.Int) (OperatorSet, error) {
	return _ContractAVSDirectory.Contract.OperatorSetsMemberOfAtIndex(&_ContractAVSDirectory.CallOpts, operator, index)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractAVSDirectory *ContractAVSDirectoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractAVSDirectory.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractAVSDirectory *ContractAVSDirectorySession) Owner() (common.Address, error) {
	return _ContractAVSDirectory.Contract.Owner(&_ContractAVSDirectory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractAVSDirectory *ContractAVSDirectoryCallerSession) Owner() (common.Address, error) {
	return _ContractAVSDirectory.Contract.Owner(&_ContractAVSDirectory.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractAVSDirectory *ContractAVSDirectoryCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _ContractAVSDirectory.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractAVSDirectory *ContractAVSDirectorySession) Paused(index uint8) (bool, error) {
	return _ContractAVSDirectory.Contract.Paused(&_ContractAVSDirectory.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractAVSDirectory *ContractAVSDirectoryCallerSession) Paused(index uint8) (bool, error) {
	return _ContractAVSDirectory.Contract.Paused(&_ContractAVSDirectory.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractAVSDirectory *ContractAVSDirectoryCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractAVSDirectory.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractAVSDirectory *ContractAVSDirectorySession) Paused0() (*big.Int, error) {
	return _ContractAVSDirectory.Contract.Paused0(&_ContractAVSDirectory.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractAVSDirectory *ContractAVSDirectoryCallerSession) Paused0() (*big.Int, error) {
	return _ContractAVSDirectory.Contract.Paused0(&_ContractAVSDirectory.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractAVSDirectory *ContractAVSDirectoryCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractAVSDirectory.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractAVSDirectory *ContractAVSDirectorySession) PauserRegistry() (common.Address, error) {
	return _ContractAVSDirectory.Contract.PauserRegistry(&_ContractAVSDirectory.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractAVSDirectory *ContractAVSDirectoryCallerSession) PauserRegistry() (common.Address, error) {
	return _ContractAVSDirectory.Contract.PauserRegistry(&_ContractAVSDirectory.CallOpts)
}

// AddStrategiesToOperatorSet is a paid mutator transaction binding the contract method 0x76999342.
//
// Solidity: function addStrategiesToOperatorSet(uint32 operatorSetId, address[] strategies) returns()
func (_ContractAVSDirectory *ContractAVSDirectoryTransactor) AddStrategiesToOperatorSet(opts *bind.TransactOpts, operatorSetId uint32, strategies []common.Address) (*types.Transaction, error) {
	return _ContractAVSDirectory.contract.Transact(opts, "addStrategiesToOperatorSet", operatorSetId, strategies)
}

// AddStrategiesToOperatorSet is a paid mutator transaction binding the contract method 0x76999342.
//
// Solidity: function addStrategiesToOperatorSet(uint32 operatorSetId, address[] strategies) returns()
func (_ContractAVSDirectory *ContractAVSDirectorySession) AddStrategiesToOperatorSet(operatorSetId uint32, strategies []common.Address) (*types.Transaction, error) {
	return _ContractAVSDirectory.Contract.AddStrategiesToOperatorSet(&_ContractAVSDirectory.TransactOpts, operatorSetId, strategies)
}

// AddStrategiesToOperatorSet is a paid mutator transaction binding the contract method 0x76999342.
//
// Solidity: function addStrategiesToOperatorSet(uint32 operatorSetId, address[] strategies) returns()
func (_ContractAVSDirectory *ContractAVSDirectoryTransactorSession) AddStrategiesToOperatorSet(operatorSetId uint32, strategies []common.Address) (*types.Transaction, error) {
	return _ContractAVSDirectory.Contract.AddStrategiesToOperatorSet(&_ContractAVSDirectory.TransactOpts, operatorSetId, strategies)
}

// BecomeOperatorSetAVS is a paid mutator transaction binding the contract method 0xaec205c5.
//
// Solidity: function becomeOperatorSetAVS() returns()
func (_ContractAVSDirectory *ContractAVSDirectoryTransactor) BecomeOperatorSetAVS(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractAVSDirectory.contract.Transact(opts, "becomeOperatorSetAVS")
}

// BecomeOperatorSetAVS is a paid mutator transaction binding the contract method 0xaec205c5.
//
// Solidity: function becomeOperatorSetAVS() returns()
func (_ContractAVSDirectory *ContractAVSDirectorySession) BecomeOperatorSetAVS() (*types.Transaction, error) {
	return _ContractAVSDirectory.Contract.BecomeOperatorSetAVS(&_ContractAVSDirectory.TransactOpts)
}

// BecomeOperatorSetAVS is a paid mutator transaction binding the contract method 0xaec205c5.
//
// Solidity: function becomeOperatorSetAVS() returns()
func (_ContractAVSDirectory *ContractAVSDirectoryTransactorSession) BecomeOperatorSetAVS() (*types.Transaction, error) {
	return _ContractAVSDirectory.Contract.BecomeOperatorSetAVS(&_ContractAVSDirectory.TransactOpts)
}

// CancelSalt is a paid mutator transaction binding the contract method 0xec76f442.
//
// Solidity: function cancelSalt(bytes32 salt) returns()
func (_ContractAVSDirectory *ContractAVSDirectoryTransactor) CancelSalt(opts *bind.TransactOpts, salt [32]byte) (*types.Transaction, error) {
	return _ContractAVSDirectory.contract.Transact(opts, "cancelSalt", salt)
}

// CancelSalt is a paid mutator transaction binding the contract method 0xec76f442.
//
// Solidity: function cancelSalt(bytes32 salt) returns()
func (_ContractAVSDirectory *ContractAVSDirectorySession) CancelSalt(salt [32]byte) (*types.Transaction, error) {
	return _ContractAVSDirectory.Contract.CancelSalt(&_ContractAVSDirectory.TransactOpts, salt)
}

// CancelSalt is a paid mutator transaction binding the contract method 0xec76f442.
//
// Solidity: function cancelSalt(bytes32 salt) returns()
func (_ContractAVSDirectory *ContractAVSDirectoryTransactorSession) CancelSalt(salt [32]byte) (*types.Transaction, error) {
	return _ContractAVSDirectory.Contract.CancelSalt(&_ContractAVSDirectory.TransactOpts, salt)
}

// CreateOperatorSets is a paid mutator transaction binding the contract method 0xafe02ed5.
//
// Solidity: function createOperatorSets(uint32[] operatorSetIds) returns()
func (_ContractAVSDirectory *ContractAVSDirectoryTransactor) CreateOperatorSets(opts *bind.TransactOpts, operatorSetIds []uint32) (*types.Transaction, error) {
	return _ContractAVSDirectory.contract.Transact(opts, "createOperatorSets", operatorSetIds)
}

// CreateOperatorSets is a paid mutator transaction binding the contract method 0xafe02ed5.
//
// Solidity: function createOperatorSets(uint32[] operatorSetIds) returns()
func (_ContractAVSDirectory *ContractAVSDirectorySession) CreateOperatorSets(operatorSetIds []uint32) (*types.Transaction, error) {
	return _ContractAVSDirectory.Contract.CreateOperatorSets(&_ContractAVSDirectory.TransactOpts, operatorSetIds)
}

// CreateOperatorSets is a paid mutator transaction binding the contract method 0xafe02ed5.
//
// Solidity: function createOperatorSets(uint32[] operatorSetIds) returns()
func (_ContractAVSDirectory *ContractAVSDirectoryTransactorSession) CreateOperatorSets(operatorSetIds []uint32) (*types.Transaction, error) {
	return _ContractAVSDirectory.Contract.CreateOperatorSets(&_ContractAVSDirectory.TransactOpts, operatorSetIds)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_ContractAVSDirectory *ContractAVSDirectoryTransactor) DeregisterOperatorFromAVS(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _ContractAVSDirectory.contract.Transact(opts, "deregisterOperatorFromAVS", operator)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_ContractAVSDirectory *ContractAVSDirectorySession) DeregisterOperatorFromAVS(operator common.Address) (*types.Transaction, error) {
	return _ContractAVSDirectory.Contract.DeregisterOperatorFromAVS(&_ContractAVSDirectory.TransactOpts, operator)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_ContractAVSDirectory *ContractAVSDirectoryTransactorSession) DeregisterOperatorFromAVS(operator common.Address) (*types.Transaction, error) {
	return _ContractAVSDirectory.Contract.DeregisterOperatorFromAVS(&_ContractAVSDirectory.TransactOpts, operator)
}

// DeregisterOperatorFromOperatorSets is a paid mutator transaction binding the contract method 0xc1a8e2c5.
//
// Solidity: function deregisterOperatorFromOperatorSets(address operator, uint32[] operatorSetIds) returns()
func (_ContractAVSDirectory *ContractAVSDirectoryTransactor) DeregisterOperatorFromOperatorSets(opts *bind.TransactOpts, operator common.Address, operatorSetIds []uint32) (*types.Transaction, error) {
	return _ContractAVSDirectory.contract.Transact(opts, "deregisterOperatorFromOperatorSets", operator, operatorSetIds)
}

// DeregisterOperatorFromOperatorSets is a paid mutator transaction binding the contract method 0xc1a8e2c5.
//
// Solidity: function deregisterOperatorFromOperatorSets(address operator, uint32[] operatorSetIds) returns()
func (_ContractAVSDirectory *ContractAVSDirectorySession) DeregisterOperatorFromOperatorSets(operator common.Address, operatorSetIds []uint32) (*types.Transaction, error) {
	return _ContractAVSDirectory.Contract.DeregisterOperatorFromOperatorSets(&_ContractAVSDirectory.TransactOpts, operator, operatorSetIds)
}

// DeregisterOperatorFromOperatorSets is a paid mutator transaction binding the contract method 0xc1a8e2c5.
//
// Solidity: function deregisterOperatorFromOperatorSets(address operator, uint32[] operatorSetIds) returns()
func (_ContractAVSDirectory *ContractAVSDirectoryTransactorSession) DeregisterOperatorFromOperatorSets(operator common.Address, operatorSetIds []uint32) (*types.Transaction, error) {
	return _ContractAVSDirectory.Contract.DeregisterOperatorFromOperatorSets(&_ContractAVSDirectory.TransactOpts, operator, operatorSetIds)
}

// ForceDeregisterFromOperatorSets is a paid mutator transaction binding the contract method 0x3fee332d.
//
// Solidity: function forceDeregisterFromOperatorSets(address operator, address avs, uint32[] operatorSetIds, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractAVSDirectory *ContractAVSDirectoryTransactor) ForceDeregisterFromOperatorSets(opts *bind.TransactOpts, operator common.Address, avs common.Address, operatorSetIds []uint32, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractAVSDirectory.contract.Transact(opts, "forceDeregisterFromOperatorSets", operator, avs, operatorSetIds, operatorSignature)
}

// ForceDeregisterFromOperatorSets is a paid mutator transaction binding the contract method 0x3fee332d.
//
// Solidity: function forceDeregisterFromOperatorSets(address operator, address avs, uint32[] operatorSetIds, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractAVSDirectory *ContractAVSDirectorySession) ForceDeregisterFromOperatorSets(operator common.Address, avs common.Address, operatorSetIds []uint32, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractAVSDirectory.Contract.ForceDeregisterFromOperatorSets(&_ContractAVSDirectory.TransactOpts, operator, avs, operatorSetIds, operatorSignature)
}

// ForceDeregisterFromOperatorSets is a paid mutator transaction binding the contract method 0x3fee332d.
//
// Solidity: function forceDeregisterFromOperatorSets(address operator, address avs, uint32[] operatorSetIds, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractAVSDirectory *ContractAVSDirectoryTransactorSession) ForceDeregisterFromOperatorSets(operator common.Address, avs common.Address, operatorSetIds []uint32, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractAVSDirectory.Contract.ForceDeregisterFromOperatorSets(&_ContractAVSDirectory.TransactOpts, operator, avs, operatorSetIds, operatorSignature)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address initialOwner, address _pauserRegistry, uint256 initialPausedStatus) returns()
func (_ContractAVSDirectory *ContractAVSDirectoryTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address, _pauserRegistry common.Address, initialPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractAVSDirectory.contract.Transact(opts, "initialize", initialOwner, _pauserRegistry, initialPausedStatus)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address initialOwner, address _pauserRegistry, uint256 initialPausedStatus) returns()
func (_ContractAVSDirectory *ContractAVSDirectorySession) Initialize(initialOwner common.Address, _pauserRegistry common.Address, initialPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractAVSDirectory.Contract.Initialize(&_ContractAVSDirectory.TransactOpts, initialOwner, _pauserRegistry, initialPausedStatus)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address initialOwner, address _pauserRegistry, uint256 initialPausedStatus) returns()
func (_ContractAVSDirectory *ContractAVSDirectoryTransactorSession) Initialize(initialOwner common.Address, _pauserRegistry common.Address, initialPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractAVSDirectory.Contract.Initialize(&_ContractAVSDirectory.TransactOpts, initialOwner, _pauserRegistry, initialPausedStatus)
}

// MigrateOperatorsToOperatorSets is a paid mutator transaction binding the contract method 0xef2dfa8d.
//
// Solidity: function migrateOperatorsToOperatorSets(address[] operators, uint32[][] operatorSetIds) returns()
func (_ContractAVSDirectory *ContractAVSDirectoryTransactor) MigrateOperatorsToOperatorSets(opts *bind.TransactOpts, operators []common.Address, operatorSetIds [][]uint32) (*types.Transaction, error) {
	return _ContractAVSDirectory.contract.Transact(opts, "migrateOperatorsToOperatorSets", operators, operatorSetIds)
}

// MigrateOperatorsToOperatorSets is a paid mutator transaction binding the contract method 0xef2dfa8d.
//
// Solidity: function migrateOperatorsToOperatorSets(address[] operators, uint32[][] operatorSetIds) returns()
func (_ContractAVSDirectory *ContractAVSDirectorySession) MigrateOperatorsToOperatorSets(operators []common.Address, operatorSetIds [][]uint32) (*types.Transaction, error) {
	return _ContractAVSDirectory.Contract.MigrateOperatorsToOperatorSets(&_ContractAVSDirectory.TransactOpts, operators, operatorSetIds)
}

// MigrateOperatorsToOperatorSets is a paid mutator transaction binding the contract method 0xef2dfa8d.
//
// Solidity: function migrateOperatorsToOperatorSets(address[] operators, uint32[][] operatorSetIds) returns()
func (_ContractAVSDirectory *ContractAVSDirectoryTransactorSession) MigrateOperatorsToOperatorSets(operators []common.Address, operatorSetIds [][]uint32) (*types.Transaction, error) {
	return _ContractAVSDirectory.Contract.MigrateOperatorsToOperatorSets(&_ContractAVSDirectory.TransactOpts, operators, operatorSetIds)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractAVSDirectory *ContractAVSDirectoryTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractAVSDirectory.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractAVSDirectory *ContractAVSDirectorySession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractAVSDirectory.Contract.Pause(&_ContractAVSDirectory.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractAVSDirectory *ContractAVSDirectoryTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractAVSDirectory.Contract.Pause(&_ContractAVSDirectory.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractAVSDirectory *ContractAVSDirectoryTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractAVSDirectory.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractAVSDirectory *ContractAVSDirectorySession) PauseAll() (*types.Transaction, error) {
	return _ContractAVSDirectory.Contract.PauseAll(&_ContractAVSDirectory.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractAVSDirectory *ContractAVSDirectoryTransactorSession) PauseAll() (*types.Transaction, error) {
	return _ContractAVSDirectory.Contract.PauseAll(&_ContractAVSDirectory.TransactOpts)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractAVSDirectory *ContractAVSDirectoryTransactor) RegisterOperatorToAVS(opts *bind.TransactOpts, operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractAVSDirectory.contract.Transact(opts, "registerOperatorToAVS", operator, operatorSignature)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractAVSDirectory *ContractAVSDirectorySession) RegisterOperatorToAVS(operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractAVSDirectory.Contract.RegisterOperatorToAVS(&_ContractAVSDirectory.TransactOpts, operator, operatorSignature)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractAVSDirectory *ContractAVSDirectoryTransactorSession) RegisterOperatorToAVS(operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractAVSDirectory.Contract.RegisterOperatorToAVS(&_ContractAVSDirectory.TransactOpts, operator, operatorSignature)
}

// RegisterOperatorToOperatorSets is a paid mutator transaction binding the contract method 0x1e2199e2.
//
// Solidity: function registerOperatorToOperatorSets(address operator, uint32[] operatorSetIds, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractAVSDirectory *ContractAVSDirectoryTransactor) RegisterOperatorToOperatorSets(opts *bind.TransactOpts, operator common.Address, operatorSetIds []uint32, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractAVSDirectory.contract.Transact(opts, "registerOperatorToOperatorSets", operator, operatorSetIds, operatorSignature)
}

// RegisterOperatorToOperatorSets is a paid mutator transaction binding the contract method 0x1e2199e2.
//
// Solidity: function registerOperatorToOperatorSets(address operator, uint32[] operatorSetIds, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractAVSDirectory *ContractAVSDirectorySession) RegisterOperatorToOperatorSets(operator common.Address, operatorSetIds []uint32, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractAVSDirectory.Contract.RegisterOperatorToOperatorSets(&_ContractAVSDirectory.TransactOpts, operator, operatorSetIds, operatorSignature)
}

// RegisterOperatorToOperatorSets is a paid mutator transaction binding the contract method 0x1e2199e2.
//
// Solidity: function registerOperatorToOperatorSets(address operator, uint32[] operatorSetIds, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractAVSDirectory *ContractAVSDirectoryTransactorSession) RegisterOperatorToOperatorSets(operator common.Address, operatorSetIds []uint32, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractAVSDirectory.Contract.RegisterOperatorToOperatorSets(&_ContractAVSDirectory.TransactOpts, operator, operatorSetIds, operatorSignature)
}

// RemoveStrategiesFromOperatorSet is a paid mutator transaction binding the contract method 0xce7b5e4b.
//
// Solidity: function removeStrategiesFromOperatorSet(uint32 operatorSetId, address[] strategies) returns()
func (_ContractAVSDirectory *ContractAVSDirectoryTransactor) RemoveStrategiesFromOperatorSet(opts *bind.TransactOpts, operatorSetId uint32, strategies []common.Address) (*types.Transaction, error) {
	return _ContractAVSDirectory.contract.Transact(opts, "removeStrategiesFromOperatorSet", operatorSetId, strategies)
}

// RemoveStrategiesFromOperatorSet is a paid mutator transaction binding the contract method 0xce7b5e4b.
//
// Solidity: function removeStrategiesFromOperatorSet(uint32 operatorSetId, address[] strategies) returns()
func (_ContractAVSDirectory *ContractAVSDirectorySession) RemoveStrategiesFromOperatorSet(operatorSetId uint32, strategies []common.Address) (*types.Transaction, error) {
	return _ContractAVSDirectory.Contract.RemoveStrategiesFromOperatorSet(&_ContractAVSDirectory.TransactOpts, operatorSetId, strategies)
}

// RemoveStrategiesFromOperatorSet is a paid mutator transaction binding the contract method 0xce7b5e4b.
//
// Solidity: function removeStrategiesFromOperatorSet(uint32 operatorSetId, address[] strategies) returns()
func (_ContractAVSDirectory *ContractAVSDirectoryTransactorSession) RemoveStrategiesFromOperatorSet(operatorSetId uint32, strategies []common.Address) (*types.Transaction, error) {
	return _ContractAVSDirectory.Contract.RemoveStrategiesFromOperatorSet(&_ContractAVSDirectory.TransactOpts, operatorSetId, strategies)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractAVSDirectory *ContractAVSDirectoryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractAVSDirectory.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractAVSDirectory *ContractAVSDirectorySession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractAVSDirectory.Contract.RenounceOwnership(&_ContractAVSDirectory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractAVSDirectory *ContractAVSDirectoryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractAVSDirectory.Contract.RenounceOwnership(&_ContractAVSDirectory.TransactOpts)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractAVSDirectory *ContractAVSDirectoryTransactor) SetPauserRegistry(opts *bind.TransactOpts, newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractAVSDirectory.contract.Transact(opts, "setPauserRegistry", newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractAVSDirectory *ContractAVSDirectorySession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractAVSDirectory.Contract.SetPauserRegistry(&_ContractAVSDirectory.TransactOpts, newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractAVSDirectory *ContractAVSDirectoryTransactorSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractAVSDirectory.Contract.SetPauserRegistry(&_ContractAVSDirectory.TransactOpts, newPauserRegistry)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractAVSDirectory *ContractAVSDirectoryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractAVSDirectory.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractAVSDirectory *ContractAVSDirectorySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractAVSDirectory.Contract.TransferOwnership(&_ContractAVSDirectory.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractAVSDirectory *ContractAVSDirectoryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractAVSDirectory.Contract.TransferOwnership(&_ContractAVSDirectory.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractAVSDirectory *ContractAVSDirectoryTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractAVSDirectory.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractAVSDirectory *ContractAVSDirectorySession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractAVSDirectory.Contract.Unpause(&_ContractAVSDirectory.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractAVSDirectory *ContractAVSDirectoryTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractAVSDirectory.Contract.Unpause(&_ContractAVSDirectory.TransactOpts, newPausedStatus)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string metadataURI) returns()
func (_ContractAVSDirectory *ContractAVSDirectoryTransactor) UpdateAVSMetadataURI(opts *bind.TransactOpts, metadataURI string) (*types.Transaction, error) {
	return _ContractAVSDirectory.contract.Transact(opts, "updateAVSMetadataURI", metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string metadataURI) returns()
func (_ContractAVSDirectory *ContractAVSDirectorySession) UpdateAVSMetadataURI(metadataURI string) (*types.Transaction, error) {
	return _ContractAVSDirectory.Contract.UpdateAVSMetadataURI(&_ContractAVSDirectory.TransactOpts, metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string metadataURI) returns()
func (_ContractAVSDirectory *ContractAVSDirectoryTransactorSession) UpdateAVSMetadataURI(metadataURI string) (*types.Transaction, error) {
	return _ContractAVSDirectory.Contract.UpdateAVSMetadataURI(&_ContractAVSDirectory.TransactOpts, metadataURI)
}

// ContractAVSDirectoryAVSMetadataURIUpdatedIterator is returned from FilterAVSMetadataURIUpdated and is used to iterate over the raw logs and unpacked data for AVSMetadataURIUpdated events raised by the ContractAVSDirectory contract.
type ContractAVSDirectoryAVSMetadataURIUpdatedIterator struct {
	Event *ContractAVSDirectoryAVSMetadataURIUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractAVSDirectoryAVSMetadataURIUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAVSDirectoryAVSMetadataURIUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractAVSDirectoryAVSMetadataURIUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractAVSDirectoryAVSMetadataURIUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAVSDirectoryAVSMetadataURIUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAVSDirectoryAVSMetadataURIUpdated represents a AVSMetadataURIUpdated event raised by the ContractAVSDirectory contract.
type ContractAVSDirectoryAVSMetadataURIUpdated struct {
	Avs         common.Address
	MetadataURI string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAVSMetadataURIUpdated is a free log retrieval operation binding the contract event 0xa89c1dc243d8908a96dd84944bcc97d6bc6ac00dd78e20621576be6a3c943713.
//
// Solidity: event AVSMetadataURIUpdated(address indexed avs, string metadataURI)
func (_ContractAVSDirectory *ContractAVSDirectoryFilterer) FilterAVSMetadataURIUpdated(opts *bind.FilterOpts, avs []common.Address) (*ContractAVSDirectoryAVSMetadataURIUpdatedIterator, error) {

	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _ContractAVSDirectory.contract.FilterLogs(opts, "AVSMetadataURIUpdated", avsRule)
	if err != nil {
		return nil, err
	}
	return &ContractAVSDirectoryAVSMetadataURIUpdatedIterator{contract: _ContractAVSDirectory.contract, event: "AVSMetadataURIUpdated", logs: logs, sub: sub}, nil
}

// WatchAVSMetadataURIUpdated is a free log subscription operation binding the contract event 0xa89c1dc243d8908a96dd84944bcc97d6bc6ac00dd78e20621576be6a3c943713.
//
// Solidity: event AVSMetadataURIUpdated(address indexed avs, string metadataURI)
func (_ContractAVSDirectory *ContractAVSDirectoryFilterer) WatchAVSMetadataURIUpdated(opts *bind.WatchOpts, sink chan<- *ContractAVSDirectoryAVSMetadataURIUpdated, avs []common.Address) (event.Subscription, error) {

	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _ContractAVSDirectory.contract.WatchLogs(opts, "AVSMetadataURIUpdated", avsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAVSDirectoryAVSMetadataURIUpdated)
				if err := _ContractAVSDirectory.contract.UnpackLog(event, "AVSMetadataURIUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAVSMetadataURIUpdated is a log parse operation binding the contract event 0xa89c1dc243d8908a96dd84944bcc97d6bc6ac00dd78e20621576be6a3c943713.
//
// Solidity: event AVSMetadataURIUpdated(address indexed avs, string metadataURI)
func (_ContractAVSDirectory *ContractAVSDirectoryFilterer) ParseAVSMetadataURIUpdated(log types.Log) (*ContractAVSDirectoryAVSMetadataURIUpdated, error) {
	event := new(ContractAVSDirectoryAVSMetadataURIUpdated)
	if err := _ContractAVSDirectory.contract.UnpackLog(event, "AVSMetadataURIUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAVSDirectoryAVSMigratedToOperatorSetsIterator is returned from FilterAVSMigratedToOperatorSets and is used to iterate over the raw logs and unpacked data for AVSMigratedToOperatorSets events raised by the ContractAVSDirectory contract.
type ContractAVSDirectoryAVSMigratedToOperatorSetsIterator struct {
	Event *ContractAVSDirectoryAVSMigratedToOperatorSets // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractAVSDirectoryAVSMigratedToOperatorSetsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAVSDirectoryAVSMigratedToOperatorSets)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractAVSDirectoryAVSMigratedToOperatorSets)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractAVSDirectoryAVSMigratedToOperatorSetsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAVSDirectoryAVSMigratedToOperatorSetsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAVSDirectoryAVSMigratedToOperatorSets represents a AVSMigratedToOperatorSets event raised by the ContractAVSDirectory contract.
type ContractAVSDirectoryAVSMigratedToOperatorSets struct {
	Avs common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAVSMigratedToOperatorSets is a free log retrieval operation binding the contract event 0x702b0c1f6cb1cf511aaa81f72bc05a215bb3497632d72c690c822b044ab494bf.
//
// Solidity: event AVSMigratedToOperatorSets(address indexed avs)
func (_ContractAVSDirectory *ContractAVSDirectoryFilterer) FilterAVSMigratedToOperatorSets(opts *bind.FilterOpts, avs []common.Address) (*ContractAVSDirectoryAVSMigratedToOperatorSetsIterator, error) {

	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _ContractAVSDirectory.contract.FilterLogs(opts, "AVSMigratedToOperatorSets", avsRule)
	if err != nil {
		return nil, err
	}
	return &ContractAVSDirectoryAVSMigratedToOperatorSetsIterator{contract: _ContractAVSDirectory.contract, event: "AVSMigratedToOperatorSets", logs: logs, sub: sub}, nil
}

// WatchAVSMigratedToOperatorSets is a free log subscription operation binding the contract event 0x702b0c1f6cb1cf511aaa81f72bc05a215bb3497632d72c690c822b044ab494bf.
//
// Solidity: event AVSMigratedToOperatorSets(address indexed avs)
func (_ContractAVSDirectory *ContractAVSDirectoryFilterer) WatchAVSMigratedToOperatorSets(opts *bind.WatchOpts, sink chan<- *ContractAVSDirectoryAVSMigratedToOperatorSets, avs []common.Address) (event.Subscription, error) {

	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _ContractAVSDirectory.contract.WatchLogs(opts, "AVSMigratedToOperatorSets", avsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAVSDirectoryAVSMigratedToOperatorSets)
				if err := _ContractAVSDirectory.contract.UnpackLog(event, "AVSMigratedToOperatorSets", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAVSMigratedToOperatorSets is a log parse operation binding the contract event 0x702b0c1f6cb1cf511aaa81f72bc05a215bb3497632d72c690c822b044ab494bf.
//
// Solidity: event AVSMigratedToOperatorSets(address indexed avs)
func (_ContractAVSDirectory *ContractAVSDirectoryFilterer) ParseAVSMigratedToOperatorSets(log types.Log) (*ContractAVSDirectoryAVSMigratedToOperatorSets, error) {
	event := new(ContractAVSDirectoryAVSMigratedToOperatorSets)
	if err := _ContractAVSDirectory.contract.UnpackLog(event, "AVSMigratedToOperatorSets", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAVSDirectoryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractAVSDirectory contract.
type ContractAVSDirectoryInitializedIterator struct {
	Event *ContractAVSDirectoryInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractAVSDirectoryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAVSDirectoryInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractAVSDirectoryInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractAVSDirectoryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAVSDirectoryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAVSDirectoryInitialized represents a Initialized event raised by the ContractAVSDirectory contract.
type ContractAVSDirectoryInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractAVSDirectory *ContractAVSDirectoryFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractAVSDirectoryInitializedIterator, error) {

	logs, sub, err := _ContractAVSDirectory.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractAVSDirectoryInitializedIterator{contract: _ContractAVSDirectory.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractAVSDirectory *ContractAVSDirectoryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractAVSDirectoryInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractAVSDirectory.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAVSDirectoryInitialized)
				if err := _ContractAVSDirectory.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractAVSDirectory *ContractAVSDirectoryFilterer) ParseInitialized(log types.Log) (*ContractAVSDirectoryInitialized, error) {
	event := new(ContractAVSDirectoryInitialized)
	if err := _ContractAVSDirectory.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAVSDirectoryOperatorAVSRegistrationStatusUpdatedIterator is returned from FilterOperatorAVSRegistrationStatusUpdated and is used to iterate over the raw logs and unpacked data for OperatorAVSRegistrationStatusUpdated events raised by the ContractAVSDirectory contract.
type ContractAVSDirectoryOperatorAVSRegistrationStatusUpdatedIterator struct {
	Event *ContractAVSDirectoryOperatorAVSRegistrationStatusUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractAVSDirectoryOperatorAVSRegistrationStatusUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAVSDirectoryOperatorAVSRegistrationStatusUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractAVSDirectoryOperatorAVSRegistrationStatusUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractAVSDirectoryOperatorAVSRegistrationStatusUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAVSDirectoryOperatorAVSRegistrationStatusUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAVSDirectoryOperatorAVSRegistrationStatusUpdated represents a OperatorAVSRegistrationStatusUpdated event raised by the ContractAVSDirectory contract.
type ContractAVSDirectoryOperatorAVSRegistrationStatusUpdated struct {
	Operator common.Address
	Avs      common.Address
	Status   uint8
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorAVSRegistrationStatusUpdated is a free log retrieval operation binding the contract event 0xf0952b1c65271d819d39983d2abb044b9cace59bcc4d4dd389f586ebdcb15b41.
//
// Solidity: event OperatorAVSRegistrationStatusUpdated(address indexed operator, address indexed avs, uint8 status)
func (_ContractAVSDirectory *ContractAVSDirectoryFilterer) FilterOperatorAVSRegistrationStatusUpdated(opts *bind.FilterOpts, operator []common.Address, avs []common.Address) (*ContractAVSDirectoryOperatorAVSRegistrationStatusUpdatedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _ContractAVSDirectory.contract.FilterLogs(opts, "OperatorAVSRegistrationStatusUpdated", operatorRule, avsRule)
	if err != nil {
		return nil, err
	}
	return &ContractAVSDirectoryOperatorAVSRegistrationStatusUpdatedIterator{contract: _ContractAVSDirectory.contract, event: "OperatorAVSRegistrationStatusUpdated", logs: logs, sub: sub}, nil
}

// WatchOperatorAVSRegistrationStatusUpdated is a free log subscription operation binding the contract event 0xf0952b1c65271d819d39983d2abb044b9cace59bcc4d4dd389f586ebdcb15b41.
//
// Solidity: event OperatorAVSRegistrationStatusUpdated(address indexed operator, address indexed avs, uint8 status)
func (_ContractAVSDirectory *ContractAVSDirectoryFilterer) WatchOperatorAVSRegistrationStatusUpdated(opts *bind.WatchOpts, sink chan<- *ContractAVSDirectoryOperatorAVSRegistrationStatusUpdated, operator []common.Address, avs []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _ContractAVSDirectory.contract.WatchLogs(opts, "OperatorAVSRegistrationStatusUpdated", operatorRule, avsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAVSDirectoryOperatorAVSRegistrationStatusUpdated)
				if err := _ContractAVSDirectory.contract.UnpackLog(event, "OperatorAVSRegistrationStatusUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOperatorAVSRegistrationStatusUpdated is a log parse operation binding the contract event 0xf0952b1c65271d819d39983d2abb044b9cace59bcc4d4dd389f586ebdcb15b41.
//
// Solidity: event OperatorAVSRegistrationStatusUpdated(address indexed operator, address indexed avs, uint8 status)
func (_ContractAVSDirectory *ContractAVSDirectoryFilterer) ParseOperatorAVSRegistrationStatusUpdated(log types.Log) (*ContractAVSDirectoryOperatorAVSRegistrationStatusUpdated, error) {
	event := new(ContractAVSDirectoryOperatorAVSRegistrationStatusUpdated)
	if err := _ContractAVSDirectory.contract.UnpackLog(event, "OperatorAVSRegistrationStatusUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAVSDirectoryOperatorAddedToOperatorSetIterator is returned from FilterOperatorAddedToOperatorSet and is used to iterate over the raw logs and unpacked data for OperatorAddedToOperatorSet events raised by the ContractAVSDirectory contract.
type ContractAVSDirectoryOperatorAddedToOperatorSetIterator struct {
	Event *ContractAVSDirectoryOperatorAddedToOperatorSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractAVSDirectoryOperatorAddedToOperatorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAVSDirectoryOperatorAddedToOperatorSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractAVSDirectoryOperatorAddedToOperatorSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractAVSDirectoryOperatorAddedToOperatorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAVSDirectoryOperatorAddedToOperatorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAVSDirectoryOperatorAddedToOperatorSet represents a OperatorAddedToOperatorSet event raised by the ContractAVSDirectory contract.
type ContractAVSDirectoryOperatorAddedToOperatorSet struct {
	Operator    common.Address
	OperatorSet OperatorSet
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorAddedToOperatorSet is a free log retrieval operation binding the contract event 0x43232edf9071753d2321e5fa7e018363ee248e5f2142e6c08edd3265bfb4895e.
//
// Solidity: event OperatorAddedToOperatorSet(address indexed operator, (address,uint32) operatorSet)
func (_ContractAVSDirectory *ContractAVSDirectoryFilterer) FilterOperatorAddedToOperatorSet(opts *bind.FilterOpts, operator []common.Address) (*ContractAVSDirectoryOperatorAddedToOperatorSetIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractAVSDirectory.contract.FilterLogs(opts, "OperatorAddedToOperatorSet", operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractAVSDirectoryOperatorAddedToOperatorSetIterator{contract: _ContractAVSDirectory.contract, event: "OperatorAddedToOperatorSet", logs: logs, sub: sub}, nil
}

// WatchOperatorAddedToOperatorSet is a free log subscription operation binding the contract event 0x43232edf9071753d2321e5fa7e018363ee248e5f2142e6c08edd3265bfb4895e.
//
// Solidity: event OperatorAddedToOperatorSet(address indexed operator, (address,uint32) operatorSet)
func (_ContractAVSDirectory *ContractAVSDirectoryFilterer) WatchOperatorAddedToOperatorSet(opts *bind.WatchOpts, sink chan<- *ContractAVSDirectoryOperatorAddedToOperatorSet, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractAVSDirectory.contract.WatchLogs(opts, "OperatorAddedToOperatorSet", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAVSDirectoryOperatorAddedToOperatorSet)
				if err := _ContractAVSDirectory.contract.UnpackLog(event, "OperatorAddedToOperatorSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOperatorAddedToOperatorSet is a log parse operation binding the contract event 0x43232edf9071753d2321e5fa7e018363ee248e5f2142e6c08edd3265bfb4895e.
//
// Solidity: event OperatorAddedToOperatorSet(address indexed operator, (address,uint32) operatorSet)
func (_ContractAVSDirectory *ContractAVSDirectoryFilterer) ParseOperatorAddedToOperatorSet(log types.Log) (*ContractAVSDirectoryOperatorAddedToOperatorSet, error) {
	event := new(ContractAVSDirectoryOperatorAddedToOperatorSet)
	if err := _ContractAVSDirectory.contract.UnpackLog(event, "OperatorAddedToOperatorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAVSDirectoryOperatorMigratedToOperatorSetsIterator is returned from FilterOperatorMigratedToOperatorSets and is used to iterate over the raw logs and unpacked data for OperatorMigratedToOperatorSets events raised by the ContractAVSDirectory contract.
type ContractAVSDirectoryOperatorMigratedToOperatorSetsIterator struct {
	Event *ContractAVSDirectoryOperatorMigratedToOperatorSets // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractAVSDirectoryOperatorMigratedToOperatorSetsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAVSDirectoryOperatorMigratedToOperatorSets)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractAVSDirectoryOperatorMigratedToOperatorSets)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractAVSDirectoryOperatorMigratedToOperatorSetsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAVSDirectoryOperatorMigratedToOperatorSetsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAVSDirectoryOperatorMigratedToOperatorSets represents a OperatorMigratedToOperatorSets event raised by the ContractAVSDirectory contract.
type ContractAVSDirectoryOperatorMigratedToOperatorSets struct {
	Operator       common.Address
	Avs            common.Address
	OperatorSetIds []uint32
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterOperatorMigratedToOperatorSets is a free log retrieval operation binding the contract event 0x54f33cfdd1ca703d795986b986fd47d742eab1904ecd2a5fdb8d6595e5904a01.
//
// Solidity: event OperatorMigratedToOperatorSets(address indexed operator, address indexed avs, uint32[] operatorSetIds)
func (_ContractAVSDirectory *ContractAVSDirectoryFilterer) FilterOperatorMigratedToOperatorSets(opts *bind.FilterOpts, operator []common.Address, avs []common.Address) (*ContractAVSDirectoryOperatorMigratedToOperatorSetsIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _ContractAVSDirectory.contract.FilterLogs(opts, "OperatorMigratedToOperatorSets", operatorRule, avsRule)
	if err != nil {
		return nil, err
	}
	return &ContractAVSDirectoryOperatorMigratedToOperatorSetsIterator{contract: _ContractAVSDirectory.contract, event: "OperatorMigratedToOperatorSets", logs: logs, sub: sub}, nil
}

// WatchOperatorMigratedToOperatorSets is a free log subscription operation binding the contract event 0x54f33cfdd1ca703d795986b986fd47d742eab1904ecd2a5fdb8d6595e5904a01.
//
// Solidity: event OperatorMigratedToOperatorSets(address indexed operator, address indexed avs, uint32[] operatorSetIds)
func (_ContractAVSDirectory *ContractAVSDirectoryFilterer) WatchOperatorMigratedToOperatorSets(opts *bind.WatchOpts, sink chan<- *ContractAVSDirectoryOperatorMigratedToOperatorSets, operator []common.Address, avs []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _ContractAVSDirectory.contract.WatchLogs(opts, "OperatorMigratedToOperatorSets", operatorRule, avsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAVSDirectoryOperatorMigratedToOperatorSets)
				if err := _ContractAVSDirectory.contract.UnpackLog(event, "OperatorMigratedToOperatorSets", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOperatorMigratedToOperatorSets is a log parse operation binding the contract event 0x54f33cfdd1ca703d795986b986fd47d742eab1904ecd2a5fdb8d6595e5904a01.
//
// Solidity: event OperatorMigratedToOperatorSets(address indexed operator, address indexed avs, uint32[] operatorSetIds)
func (_ContractAVSDirectory *ContractAVSDirectoryFilterer) ParseOperatorMigratedToOperatorSets(log types.Log) (*ContractAVSDirectoryOperatorMigratedToOperatorSets, error) {
	event := new(ContractAVSDirectoryOperatorMigratedToOperatorSets)
	if err := _ContractAVSDirectory.contract.UnpackLog(event, "OperatorMigratedToOperatorSets", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAVSDirectoryOperatorRemovedFromOperatorSetIterator is returned from FilterOperatorRemovedFromOperatorSet and is used to iterate over the raw logs and unpacked data for OperatorRemovedFromOperatorSet events raised by the ContractAVSDirectory contract.
type ContractAVSDirectoryOperatorRemovedFromOperatorSetIterator struct {
	Event *ContractAVSDirectoryOperatorRemovedFromOperatorSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractAVSDirectoryOperatorRemovedFromOperatorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAVSDirectoryOperatorRemovedFromOperatorSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractAVSDirectoryOperatorRemovedFromOperatorSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractAVSDirectoryOperatorRemovedFromOperatorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAVSDirectoryOperatorRemovedFromOperatorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAVSDirectoryOperatorRemovedFromOperatorSet represents a OperatorRemovedFromOperatorSet event raised by the ContractAVSDirectory contract.
type ContractAVSDirectoryOperatorRemovedFromOperatorSet struct {
	Operator    common.Address
	OperatorSet OperatorSet
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorRemovedFromOperatorSet is a free log retrieval operation binding the contract event 0xad34c3070be1dffbcaa499d000ba2b8d9848aefcac3059df245dd95c4ece14fe.
//
// Solidity: event OperatorRemovedFromOperatorSet(address indexed operator, (address,uint32) operatorSet)
func (_ContractAVSDirectory *ContractAVSDirectoryFilterer) FilterOperatorRemovedFromOperatorSet(opts *bind.FilterOpts, operator []common.Address) (*ContractAVSDirectoryOperatorRemovedFromOperatorSetIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractAVSDirectory.contract.FilterLogs(opts, "OperatorRemovedFromOperatorSet", operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractAVSDirectoryOperatorRemovedFromOperatorSetIterator{contract: _ContractAVSDirectory.contract, event: "OperatorRemovedFromOperatorSet", logs: logs, sub: sub}, nil
}

// WatchOperatorRemovedFromOperatorSet is a free log subscription operation binding the contract event 0xad34c3070be1dffbcaa499d000ba2b8d9848aefcac3059df245dd95c4ece14fe.
//
// Solidity: event OperatorRemovedFromOperatorSet(address indexed operator, (address,uint32) operatorSet)
func (_ContractAVSDirectory *ContractAVSDirectoryFilterer) WatchOperatorRemovedFromOperatorSet(opts *bind.WatchOpts, sink chan<- *ContractAVSDirectoryOperatorRemovedFromOperatorSet, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractAVSDirectory.contract.WatchLogs(opts, "OperatorRemovedFromOperatorSet", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAVSDirectoryOperatorRemovedFromOperatorSet)
				if err := _ContractAVSDirectory.contract.UnpackLog(event, "OperatorRemovedFromOperatorSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOperatorRemovedFromOperatorSet is a log parse operation binding the contract event 0xad34c3070be1dffbcaa499d000ba2b8d9848aefcac3059df245dd95c4ece14fe.
//
// Solidity: event OperatorRemovedFromOperatorSet(address indexed operator, (address,uint32) operatorSet)
func (_ContractAVSDirectory *ContractAVSDirectoryFilterer) ParseOperatorRemovedFromOperatorSet(log types.Log) (*ContractAVSDirectoryOperatorRemovedFromOperatorSet, error) {
	event := new(ContractAVSDirectoryOperatorRemovedFromOperatorSet)
	if err := _ContractAVSDirectory.contract.UnpackLog(event, "OperatorRemovedFromOperatorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAVSDirectoryOperatorSetCreatedIterator is returned from FilterOperatorSetCreated and is used to iterate over the raw logs and unpacked data for OperatorSetCreated events raised by the ContractAVSDirectory contract.
type ContractAVSDirectoryOperatorSetCreatedIterator struct {
	Event *ContractAVSDirectoryOperatorSetCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractAVSDirectoryOperatorSetCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAVSDirectoryOperatorSetCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractAVSDirectoryOperatorSetCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractAVSDirectoryOperatorSetCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAVSDirectoryOperatorSetCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAVSDirectoryOperatorSetCreated represents a OperatorSetCreated event raised by the ContractAVSDirectory contract.
type ContractAVSDirectoryOperatorSetCreated struct {
	OperatorSet OperatorSet
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorSetCreated is a free log retrieval operation binding the contract event 0x31629285ead2335ae0933f86ed2ae63321f7af77b4e6eaabc42c057880977e6c.
//
// Solidity: event OperatorSetCreated((address,uint32) operatorSet)
func (_ContractAVSDirectory *ContractAVSDirectoryFilterer) FilterOperatorSetCreated(opts *bind.FilterOpts) (*ContractAVSDirectoryOperatorSetCreatedIterator, error) {

	logs, sub, err := _ContractAVSDirectory.contract.FilterLogs(opts, "OperatorSetCreated")
	if err != nil {
		return nil, err
	}
	return &ContractAVSDirectoryOperatorSetCreatedIterator{contract: _ContractAVSDirectory.contract, event: "OperatorSetCreated", logs: logs, sub: sub}, nil
}

// WatchOperatorSetCreated is a free log subscription operation binding the contract event 0x31629285ead2335ae0933f86ed2ae63321f7af77b4e6eaabc42c057880977e6c.
//
// Solidity: event OperatorSetCreated((address,uint32) operatorSet)
func (_ContractAVSDirectory *ContractAVSDirectoryFilterer) WatchOperatorSetCreated(opts *bind.WatchOpts, sink chan<- *ContractAVSDirectoryOperatorSetCreated) (event.Subscription, error) {

	logs, sub, err := _ContractAVSDirectory.contract.WatchLogs(opts, "OperatorSetCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAVSDirectoryOperatorSetCreated)
				if err := _ContractAVSDirectory.contract.UnpackLog(event, "OperatorSetCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOperatorSetCreated is a log parse operation binding the contract event 0x31629285ead2335ae0933f86ed2ae63321f7af77b4e6eaabc42c057880977e6c.
//
// Solidity: event OperatorSetCreated((address,uint32) operatorSet)
func (_ContractAVSDirectory *ContractAVSDirectoryFilterer) ParseOperatorSetCreated(log types.Log) (*ContractAVSDirectoryOperatorSetCreated, error) {
	event := new(ContractAVSDirectoryOperatorSetCreated)
	if err := _ContractAVSDirectory.contract.UnpackLog(event, "OperatorSetCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAVSDirectoryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContractAVSDirectory contract.
type ContractAVSDirectoryOwnershipTransferredIterator struct {
	Event *ContractAVSDirectoryOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractAVSDirectoryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAVSDirectoryOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractAVSDirectoryOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractAVSDirectoryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAVSDirectoryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAVSDirectoryOwnershipTransferred represents a OwnershipTransferred event raised by the ContractAVSDirectory contract.
type ContractAVSDirectoryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractAVSDirectory *ContractAVSDirectoryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractAVSDirectoryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractAVSDirectory.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractAVSDirectoryOwnershipTransferredIterator{contract: _ContractAVSDirectory.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractAVSDirectory *ContractAVSDirectoryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractAVSDirectoryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractAVSDirectory.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAVSDirectoryOwnershipTransferred)
				if err := _ContractAVSDirectory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractAVSDirectory *ContractAVSDirectoryFilterer) ParseOwnershipTransferred(log types.Log) (*ContractAVSDirectoryOwnershipTransferred, error) {
	event := new(ContractAVSDirectoryOwnershipTransferred)
	if err := _ContractAVSDirectory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAVSDirectoryPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ContractAVSDirectory contract.
type ContractAVSDirectoryPausedIterator struct {
	Event *ContractAVSDirectoryPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractAVSDirectoryPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAVSDirectoryPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractAVSDirectoryPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractAVSDirectoryPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAVSDirectoryPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAVSDirectoryPaused represents a Paused event raised by the ContractAVSDirectory contract.
type ContractAVSDirectoryPaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractAVSDirectory *ContractAVSDirectoryFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*ContractAVSDirectoryPausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractAVSDirectory.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractAVSDirectoryPausedIterator{contract: _ContractAVSDirectory.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractAVSDirectory *ContractAVSDirectoryFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ContractAVSDirectoryPaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractAVSDirectory.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAVSDirectoryPaused)
				if err := _ContractAVSDirectory.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractAVSDirectory *ContractAVSDirectoryFilterer) ParsePaused(log types.Log) (*ContractAVSDirectoryPaused, error) {
	event := new(ContractAVSDirectoryPaused)
	if err := _ContractAVSDirectory.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAVSDirectoryPauserRegistrySetIterator is returned from FilterPauserRegistrySet and is used to iterate over the raw logs and unpacked data for PauserRegistrySet events raised by the ContractAVSDirectory contract.
type ContractAVSDirectoryPauserRegistrySetIterator struct {
	Event *ContractAVSDirectoryPauserRegistrySet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractAVSDirectoryPauserRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAVSDirectoryPauserRegistrySet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractAVSDirectoryPauserRegistrySet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractAVSDirectoryPauserRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAVSDirectoryPauserRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAVSDirectoryPauserRegistrySet represents a PauserRegistrySet event raised by the ContractAVSDirectory contract.
type ContractAVSDirectoryPauserRegistrySet struct {
	PauserRegistry    common.Address
	NewPauserRegistry common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterPauserRegistrySet is a free log retrieval operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractAVSDirectory *ContractAVSDirectoryFilterer) FilterPauserRegistrySet(opts *bind.FilterOpts) (*ContractAVSDirectoryPauserRegistrySetIterator, error) {

	logs, sub, err := _ContractAVSDirectory.contract.FilterLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return &ContractAVSDirectoryPauserRegistrySetIterator{contract: _ContractAVSDirectory.contract, event: "PauserRegistrySet", logs: logs, sub: sub}, nil
}

// WatchPauserRegistrySet is a free log subscription operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractAVSDirectory *ContractAVSDirectoryFilterer) WatchPauserRegistrySet(opts *bind.WatchOpts, sink chan<- *ContractAVSDirectoryPauserRegistrySet) (event.Subscription, error) {

	logs, sub, err := _ContractAVSDirectory.contract.WatchLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAVSDirectoryPauserRegistrySet)
				if err := _ContractAVSDirectory.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePauserRegistrySet is a log parse operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractAVSDirectory *ContractAVSDirectoryFilterer) ParsePauserRegistrySet(log types.Log) (*ContractAVSDirectoryPauserRegistrySet, error) {
	event := new(ContractAVSDirectoryPauserRegistrySet)
	if err := _ContractAVSDirectory.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAVSDirectoryStrategyAddedToOperatorSetIterator is returned from FilterStrategyAddedToOperatorSet and is used to iterate over the raw logs and unpacked data for StrategyAddedToOperatorSet events raised by the ContractAVSDirectory contract.
type ContractAVSDirectoryStrategyAddedToOperatorSetIterator struct {
	Event *ContractAVSDirectoryStrategyAddedToOperatorSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractAVSDirectoryStrategyAddedToOperatorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAVSDirectoryStrategyAddedToOperatorSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractAVSDirectoryStrategyAddedToOperatorSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractAVSDirectoryStrategyAddedToOperatorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAVSDirectoryStrategyAddedToOperatorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAVSDirectoryStrategyAddedToOperatorSet represents a StrategyAddedToOperatorSet event raised by the ContractAVSDirectory contract.
type ContractAVSDirectoryStrategyAddedToOperatorSet struct {
	OperatorSet OperatorSet
	Strategy    common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterStrategyAddedToOperatorSet is a free log retrieval operation binding the contract event 0x7ab260fe0af193db5f4986770d831bda4ea46099dc817e8b6716dcae8af8e88b.
//
// Solidity: event StrategyAddedToOperatorSet((address,uint32) operatorSet, address strategy)
func (_ContractAVSDirectory *ContractAVSDirectoryFilterer) FilterStrategyAddedToOperatorSet(opts *bind.FilterOpts) (*ContractAVSDirectoryStrategyAddedToOperatorSetIterator, error) {

	logs, sub, err := _ContractAVSDirectory.contract.FilterLogs(opts, "StrategyAddedToOperatorSet")
	if err != nil {
		return nil, err
	}
	return &ContractAVSDirectoryStrategyAddedToOperatorSetIterator{contract: _ContractAVSDirectory.contract, event: "StrategyAddedToOperatorSet", logs: logs, sub: sub}, nil
}

// WatchStrategyAddedToOperatorSet is a free log subscription operation binding the contract event 0x7ab260fe0af193db5f4986770d831bda4ea46099dc817e8b6716dcae8af8e88b.
//
// Solidity: event StrategyAddedToOperatorSet((address,uint32) operatorSet, address strategy)
func (_ContractAVSDirectory *ContractAVSDirectoryFilterer) WatchStrategyAddedToOperatorSet(opts *bind.WatchOpts, sink chan<- *ContractAVSDirectoryStrategyAddedToOperatorSet) (event.Subscription, error) {

	logs, sub, err := _ContractAVSDirectory.contract.WatchLogs(opts, "StrategyAddedToOperatorSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAVSDirectoryStrategyAddedToOperatorSet)
				if err := _ContractAVSDirectory.contract.UnpackLog(event, "StrategyAddedToOperatorSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStrategyAddedToOperatorSet is a log parse operation binding the contract event 0x7ab260fe0af193db5f4986770d831bda4ea46099dc817e8b6716dcae8af8e88b.
//
// Solidity: event StrategyAddedToOperatorSet((address,uint32) operatorSet, address strategy)
func (_ContractAVSDirectory *ContractAVSDirectoryFilterer) ParseStrategyAddedToOperatorSet(log types.Log) (*ContractAVSDirectoryStrategyAddedToOperatorSet, error) {
	event := new(ContractAVSDirectoryStrategyAddedToOperatorSet)
	if err := _ContractAVSDirectory.contract.UnpackLog(event, "StrategyAddedToOperatorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAVSDirectoryStrategyRemovedFromOperatorSetIterator is returned from FilterStrategyRemovedFromOperatorSet and is used to iterate over the raw logs and unpacked data for StrategyRemovedFromOperatorSet events raised by the ContractAVSDirectory contract.
type ContractAVSDirectoryStrategyRemovedFromOperatorSetIterator struct {
	Event *ContractAVSDirectoryStrategyRemovedFromOperatorSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractAVSDirectoryStrategyRemovedFromOperatorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAVSDirectoryStrategyRemovedFromOperatorSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractAVSDirectoryStrategyRemovedFromOperatorSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractAVSDirectoryStrategyRemovedFromOperatorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAVSDirectoryStrategyRemovedFromOperatorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAVSDirectoryStrategyRemovedFromOperatorSet represents a StrategyRemovedFromOperatorSet event raised by the ContractAVSDirectory contract.
type ContractAVSDirectoryStrategyRemovedFromOperatorSet struct {
	OperatorSet OperatorSet
	Strategy    common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterStrategyRemovedFromOperatorSet is a free log retrieval operation binding the contract event 0x7b4b073d80dcac55a11177d8459ad9f664ceeb91f71f27167bb14f8152a7eeee.
//
// Solidity: event StrategyRemovedFromOperatorSet((address,uint32) operatorSet, address strategy)
func (_ContractAVSDirectory *ContractAVSDirectoryFilterer) FilterStrategyRemovedFromOperatorSet(opts *bind.FilterOpts) (*ContractAVSDirectoryStrategyRemovedFromOperatorSetIterator, error) {

	logs, sub, err := _ContractAVSDirectory.contract.FilterLogs(opts, "StrategyRemovedFromOperatorSet")
	if err != nil {
		return nil, err
	}
	return &ContractAVSDirectoryStrategyRemovedFromOperatorSetIterator{contract: _ContractAVSDirectory.contract, event: "StrategyRemovedFromOperatorSet", logs: logs, sub: sub}, nil
}

// WatchStrategyRemovedFromOperatorSet is a free log subscription operation binding the contract event 0x7b4b073d80dcac55a11177d8459ad9f664ceeb91f71f27167bb14f8152a7eeee.
//
// Solidity: event StrategyRemovedFromOperatorSet((address,uint32) operatorSet, address strategy)
func (_ContractAVSDirectory *ContractAVSDirectoryFilterer) WatchStrategyRemovedFromOperatorSet(opts *bind.WatchOpts, sink chan<- *ContractAVSDirectoryStrategyRemovedFromOperatorSet) (event.Subscription, error) {

	logs, sub, err := _ContractAVSDirectory.contract.WatchLogs(opts, "StrategyRemovedFromOperatorSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAVSDirectoryStrategyRemovedFromOperatorSet)
				if err := _ContractAVSDirectory.contract.UnpackLog(event, "StrategyRemovedFromOperatorSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStrategyRemovedFromOperatorSet is a log parse operation binding the contract event 0x7b4b073d80dcac55a11177d8459ad9f664ceeb91f71f27167bb14f8152a7eeee.
//
// Solidity: event StrategyRemovedFromOperatorSet((address,uint32) operatorSet, address strategy)
func (_ContractAVSDirectory *ContractAVSDirectoryFilterer) ParseStrategyRemovedFromOperatorSet(log types.Log) (*ContractAVSDirectoryStrategyRemovedFromOperatorSet, error) {
	event := new(ContractAVSDirectoryStrategyRemovedFromOperatorSet)
	if err := _ContractAVSDirectory.contract.UnpackLog(event, "StrategyRemovedFromOperatorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAVSDirectoryUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ContractAVSDirectory contract.
type ContractAVSDirectoryUnpausedIterator struct {
	Event *ContractAVSDirectoryUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractAVSDirectoryUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAVSDirectoryUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractAVSDirectoryUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractAVSDirectoryUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAVSDirectoryUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAVSDirectoryUnpaused represents a Unpaused event raised by the ContractAVSDirectory contract.
type ContractAVSDirectoryUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractAVSDirectory *ContractAVSDirectoryFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*ContractAVSDirectoryUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractAVSDirectory.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractAVSDirectoryUnpausedIterator{contract: _ContractAVSDirectory.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractAVSDirectory *ContractAVSDirectoryFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ContractAVSDirectoryUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractAVSDirectory.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAVSDirectoryUnpaused)
				if err := _ContractAVSDirectory.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractAVSDirectory *ContractAVSDirectoryFilterer) ParseUnpaused(log types.Log) (*ContractAVSDirectoryUnpaused, error) {
	event := new(ContractAVSDirectoryUnpaused)
	if err := _ContractAVSDirectory.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
