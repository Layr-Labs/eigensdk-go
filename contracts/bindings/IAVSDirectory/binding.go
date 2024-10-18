// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractIAVSDirectory

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

// ContractIAVSDirectoryMetaData contains all meta data concerning the ContractIAVSDirectory contract.
var ContractIAVSDirectoryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"OPERATOR_AVS_REGISTRATION_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"OPERATOR_SET_REGISTRATION_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addStrategiesToOperatorSet\",\"inputs\":[{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"becomeOperatorSetAVS\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"calculateOperatorAVSRegistrationDigestHash\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateOperatorSetForceDeregistrationTypehash\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateOperatorSetRegistrationDigestHash\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"cancelSalt\",\"inputs\":[{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createOperatorSets\",\"inputs\":[{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deregisterOperatorFromAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deregisterOperatorFromOperatorSets\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forceDeregisterFromOperatorSets\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getNumOperatorSetsOfOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNumOperatorsInOperatorSet\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorSetsOfOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"operatorSets\",\"type\":\"tuple[]\",\"internalType\":\"structOperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorsInOperatorSet\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStrategiesInOperatorSet\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"inTotalOperatorSets\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"initialPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isMember\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOperatorSet\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOperatorSetAVS\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOperatorSetBatch\",\"inputs\":[{\"name\":\"operatorSets\",\"type\":\"tuple[]\",\"internalType\":\"structOperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOperatorSlashable\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"migrateOperatorsToOperatorSets\",\"inputs\":[{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"operatorSaltIsSpent\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorSetMemberAtIndex\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorSetStatus\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"registered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"lastDeregisteredTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorSetsMemberOfAtIndex\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerOperatorToAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerOperatorToOperatorSets\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeStrategiesFromOperatorSet\",\"inputs\":[{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateAVSMetadataURI\",\"inputs\":[{\"name\":\"metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AVSMetadataURIUpdated\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"metadataURI\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AVSMigratedToOperatorSets\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorAVSRegistrationStatusUpdated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"status\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumIAVSDirectoryTypes.OperatorAVSRegistrationStatus\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorAddedToOperatorSet\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorMigratedToOperatorSets\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"indexed\":false,\"internalType\":\"uint32[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorRemovedFromOperatorSet\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSetCreated\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StrategyAddedToOperatorSet\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StrategyRemovedFromOperatorSet\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidAVS\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOperator\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOperatorSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorAlreadyRegisteredToAVS\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorNotRegisteredToAVS\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorNotRegisteredToEigenLayer\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SaltSpent\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SignatureExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StrategyAlreadyInOperatorSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StrategyNotInOperatorSet\",\"inputs\":[]}]",
}

// ContractIAVSDirectoryABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractIAVSDirectoryMetaData.ABI instead.
var ContractIAVSDirectoryABI = ContractIAVSDirectoryMetaData.ABI

// ContractIAVSDirectoryMethods is an auto generated interface around an Ethereum contract.
type ContractIAVSDirectoryMethods interface {
	ContractIAVSDirectoryCalls
	ContractIAVSDirectoryTransacts
	ContractIAVSDirectoryFilters
}

// ContractIAVSDirectoryCalls is an auto generated interface that defines the call methods available for an Ethereum contract.
type ContractIAVSDirectoryCalls interface {
	OPERATORAVSREGISTRATIONTYPEHASH(opts *bind.CallOpts) ([32]byte, error)

	OPERATORSETREGISTRATIONTYPEHASH(opts *bind.CallOpts) ([32]byte, error)

	CalculateOperatorAVSRegistrationDigestHash(opts *bind.CallOpts, operator common.Address, avs common.Address, salt [32]byte, expiry *big.Int) ([32]byte, error)

	CalculateOperatorSetForceDeregistrationTypehash(opts *bind.CallOpts, avs common.Address, operatorSetIds []uint32, salt [32]byte, expiry *big.Int) ([32]byte, error)

	CalculateOperatorSetRegistrationDigestHash(opts *bind.CallOpts, avs common.Address, operatorSetIds []uint32, salt [32]byte, expiry *big.Int) ([32]byte, error)

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

	OperatorSetStatus(opts *bind.CallOpts, avs common.Address, operator common.Address, operatorSetId uint32) (struct {
		Registered                bool
		LastDeregisteredTimestamp uint32
	}, error)

	OperatorSetsMemberOfAtIndex(opts *bind.CallOpts, operator common.Address, index *big.Int) (OperatorSet, error)
}

// ContractIAVSDirectoryTransacts is an auto generated interface that defines the transact methods available for an Ethereum contract.
type ContractIAVSDirectoryTransacts interface {
	AddStrategiesToOperatorSet(opts *bind.TransactOpts, operatorSetId uint32, strategies []common.Address) (*types.Transaction, error)

	BecomeOperatorSetAVS(opts *bind.TransactOpts) (*types.Transaction, error)

	CancelSalt(opts *bind.TransactOpts, salt [32]byte) (*types.Transaction, error)

	CreateOperatorSets(opts *bind.TransactOpts, operatorSetIds []uint32) (*types.Transaction, error)

	DeregisterOperatorFromAVS(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error)

	DeregisterOperatorFromOperatorSets(opts *bind.TransactOpts, operator common.Address, operatorSetIds []uint32) (*types.Transaction, error)

	ForceDeregisterFromOperatorSets(opts *bind.TransactOpts, operator common.Address, avs common.Address, operatorSetIds []uint32, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error)

	Initialize(opts *bind.TransactOpts, initialOwner common.Address, _pauserRegistry common.Address, initialPausedStatus *big.Int) (*types.Transaction, error)

	MigrateOperatorsToOperatorSets(opts *bind.TransactOpts, operators []common.Address, operatorSetIds [][]uint32) (*types.Transaction, error)

	RegisterOperatorToAVS(opts *bind.TransactOpts, operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error)

	RegisterOperatorToOperatorSets(opts *bind.TransactOpts, operator common.Address, operatorSetIds []uint32, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error)

	RemoveStrategiesFromOperatorSet(opts *bind.TransactOpts, operatorSetId uint32, strategies []common.Address) (*types.Transaction, error)

	UpdateAVSMetadataURI(opts *bind.TransactOpts, metadataURI string) (*types.Transaction, error)
}

// ContractIAVSDirectoryFilterer is an auto generated interface that defines the log filtering methods available for an Ethereum contract.
type ContractIAVSDirectoryFilters interface {
	FilterAVSMetadataURIUpdated(opts *bind.FilterOpts, avs []common.Address) (*ContractIAVSDirectoryAVSMetadataURIUpdatedIterator, error)
	WatchAVSMetadataURIUpdated(opts *bind.WatchOpts, sink chan<- *ContractIAVSDirectoryAVSMetadataURIUpdated, avs []common.Address) (event.Subscription, error)
	ParseAVSMetadataURIUpdated(log types.Log) (*ContractIAVSDirectoryAVSMetadataURIUpdated, error)

	FilterAVSMigratedToOperatorSets(opts *bind.FilterOpts, avs []common.Address) (*ContractIAVSDirectoryAVSMigratedToOperatorSetsIterator, error)
	WatchAVSMigratedToOperatorSets(opts *bind.WatchOpts, sink chan<- *ContractIAVSDirectoryAVSMigratedToOperatorSets, avs []common.Address) (event.Subscription, error)
	ParseAVSMigratedToOperatorSets(log types.Log) (*ContractIAVSDirectoryAVSMigratedToOperatorSets, error)

	FilterOperatorAVSRegistrationStatusUpdated(opts *bind.FilterOpts, operator []common.Address, avs []common.Address) (*ContractIAVSDirectoryOperatorAVSRegistrationStatusUpdatedIterator, error)
	WatchOperatorAVSRegistrationStatusUpdated(opts *bind.WatchOpts, sink chan<- *ContractIAVSDirectoryOperatorAVSRegistrationStatusUpdated, operator []common.Address, avs []common.Address) (event.Subscription, error)
	ParseOperatorAVSRegistrationStatusUpdated(log types.Log) (*ContractIAVSDirectoryOperatorAVSRegistrationStatusUpdated, error)

	FilterOperatorAddedToOperatorSet(opts *bind.FilterOpts, operator []common.Address) (*ContractIAVSDirectoryOperatorAddedToOperatorSetIterator, error)
	WatchOperatorAddedToOperatorSet(opts *bind.WatchOpts, sink chan<- *ContractIAVSDirectoryOperatorAddedToOperatorSet, operator []common.Address) (event.Subscription, error)
	ParseOperatorAddedToOperatorSet(log types.Log) (*ContractIAVSDirectoryOperatorAddedToOperatorSet, error)

	FilterOperatorMigratedToOperatorSets(opts *bind.FilterOpts, operator []common.Address, avs []common.Address) (*ContractIAVSDirectoryOperatorMigratedToOperatorSetsIterator, error)
	WatchOperatorMigratedToOperatorSets(opts *bind.WatchOpts, sink chan<- *ContractIAVSDirectoryOperatorMigratedToOperatorSets, operator []common.Address, avs []common.Address) (event.Subscription, error)
	ParseOperatorMigratedToOperatorSets(log types.Log) (*ContractIAVSDirectoryOperatorMigratedToOperatorSets, error)

	FilterOperatorRemovedFromOperatorSet(opts *bind.FilterOpts, operator []common.Address) (*ContractIAVSDirectoryOperatorRemovedFromOperatorSetIterator, error)
	WatchOperatorRemovedFromOperatorSet(opts *bind.WatchOpts, sink chan<- *ContractIAVSDirectoryOperatorRemovedFromOperatorSet, operator []common.Address) (event.Subscription, error)
	ParseOperatorRemovedFromOperatorSet(log types.Log) (*ContractIAVSDirectoryOperatorRemovedFromOperatorSet, error)

	FilterOperatorSetCreated(opts *bind.FilterOpts) (*ContractIAVSDirectoryOperatorSetCreatedIterator, error)
	WatchOperatorSetCreated(opts *bind.WatchOpts, sink chan<- *ContractIAVSDirectoryOperatorSetCreated) (event.Subscription, error)
	ParseOperatorSetCreated(log types.Log) (*ContractIAVSDirectoryOperatorSetCreated, error)

	FilterStrategyAddedToOperatorSet(opts *bind.FilterOpts) (*ContractIAVSDirectoryStrategyAddedToOperatorSetIterator, error)
	WatchStrategyAddedToOperatorSet(opts *bind.WatchOpts, sink chan<- *ContractIAVSDirectoryStrategyAddedToOperatorSet) (event.Subscription, error)
	ParseStrategyAddedToOperatorSet(log types.Log) (*ContractIAVSDirectoryStrategyAddedToOperatorSet, error)

	FilterStrategyRemovedFromOperatorSet(opts *bind.FilterOpts) (*ContractIAVSDirectoryStrategyRemovedFromOperatorSetIterator, error)
	WatchStrategyRemovedFromOperatorSet(opts *bind.WatchOpts, sink chan<- *ContractIAVSDirectoryStrategyRemovedFromOperatorSet) (event.Subscription, error)
	ParseStrategyRemovedFromOperatorSet(log types.Log) (*ContractIAVSDirectoryStrategyRemovedFromOperatorSet, error)
}

// ContractIAVSDirectory is an auto generated Go binding around an Ethereum contract.
type ContractIAVSDirectory struct {
	ContractIAVSDirectoryCaller     // Read-only binding to the contract
	ContractIAVSDirectoryTransactor // Write-only binding to the contract
	ContractIAVSDirectoryFilterer   // Log filterer for contract events
}

// ContractIAVSDirectory implements the ContractIAVSDirectoryMethods interface.
var _ ContractIAVSDirectoryMethods = (*ContractIAVSDirectory)(nil)

// ContractIAVSDirectoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractIAVSDirectoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractIAVSDirectoryCaller implements the ContractIAVSDirectoryCalls interface.
var _ ContractIAVSDirectoryCalls = (*ContractIAVSDirectoryCaller)(nil)

// ContractIAVSDirectoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractIAVSDirectoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractIAVSDirectoryTransactor implements the ContractIAVSDirectoryTransacts interface.
var _ ContractIAVSDirectoryTransacts = (*ContractIAVSDirectoryTransactor)(nil)

// ContractIAVSDirectoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractIAVSDirectoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractIAVSDirectoryFilterer implements the ContractIAVSDirectoryFilters interface.
var _ ContractIAVSDirectoryFilters = (*ContractIAVSDirectoryFilterer)(nil)

// ContractIAVSDirectorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractIAVSDirectorySession struct {
	Contract     *ContractIAVSDirectory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ContractIAVSDirectoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractIAVSDirectoryCallerSession struct {
	Contract *ContractIAVSDirectoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// ContractIAVSDirectoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractIAVSDirectoryTransactorSession struct {
	Contract     *ContractIAVSDirectoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// ContractIAVSDirectoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractIAVSDirectoryRaw struct {
	Contract *ContractIAVSDirectory // Generic contract binding to access the raw methods on
}

// ContractIAVSDirectoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractIAVSDirectoryCallerRaw struct {
	Contract *ContractIAVSDirectoryCaller // Generic read-only contract binding to access the raw methods on
}

// ContractIAVSDirectoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractIAVSDirectoryTransactorRaw struct {
	Contract *ContractIAVSDirectoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractIAVSDirectory creates a new instance of ContractIAVSDirectory, bound to a specific deployed contract.
func NewContractIAVSDirectory(address common.Address, backend bind.ContractBackend) (*ContractIAVSDirectory, error) {
	contract, err := bindContractIAVSDirectory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractIAVSDirectory{ContractIAVSDirectoryCaller: ContractIAVSDirectoryCaller{contract: contract}, ContractIAVSDirectoryTransactor: ContractIAVSDirectoryTransactor{contract: contract}, ContractIAVSDirectoryFilterer: ContractIAVSDirectoryFilterer{contract: contract}}, nil
}

// NewContractIAVSDirectoryCaller creates a new read-only instance of ContractIAVSDirectory, bound to a specific deployed contract.
func NewContractIAVSDirectoryCaller(address common.Address, caller bind.ContractCaller) (*ContractIAVSDirectoryCaller, error) {
	contract, err := bindContractIAVSDirectory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractIAVSDirectoryCaller{contract: contract}, nil
}

// NewContractIAVSDirectoryTransactor creates a new write-only instance of ContractIAVSDirectory, bound to a specific deployed contract.
func NewContractIAVSDirectoryTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractIAVSDirectoryTransactor, error) {
	contract, err := bindContractIAVSDirectory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractIAVSDirectoryTransactor{contract: contract}, nil
}

// NewContractIAVSDirectoryFilterer creates a new log filterer instance of ContractIAVSDirectory, bound to a specific deployed contract.
func NewContractIAVSDirectoryFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractIAVSDirectoryFilterer, error) {
	contract, err := bindContractIAVSDirectory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractIAVSDirectoryFilterer{contract: contract}, nil
}

// bindContractIAVSDirectory binds a generic wrapper to an already deployed contract.
func bindContractIAVSDirectory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractIAVSDirectoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractIAVSDirectory *ContractIAVSDirectoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractIAVSDirectory.Contract.ContractIAVSDirectoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractIAVSDirectory *ContractIAVSDirectoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractIAVSDirectory.Contract.ContractIAVSDirectoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractIAVSDirectory *ContractIAVSDirectoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractIAVSDirectory.Contract.ContractIAVSDirectoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractIAVSDirectory *ContractIAVSDirectoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractIAVSDirectory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractIAVSDirectory *ContractIAVSDirectoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractIAVSDirectory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractIAVSDirectory *ContractIAVSDirectoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractIAVSDirectory.Contract.contract.Transact(opts, method, params...)
}

// OPERATORAVSREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0xd79aceab.
//
// Solidity: function OPERATOR_AVS_REGISTRATION_TYPEHASH() view returns(bytes32)
func (_ContractIAVSDirectory *ContractIAVSDirectoryCaller) OPERATORAVSREGISTRATIONTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ContractIAVSDirectory.contract.Call(opts, &out, "OPERATOR_AVS_REGISTRATION_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OPERATORAVSREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0xd79aceab.
//
// Solidity: function OPERATOR_AVS_REGISTRATION_TYPEHASH() view returns(bytes32)
func (_ContractIAVSDirectory *ContractIAVSDirectorySession) OPERATORAVSREGISTRATIONTYPEHASH() ([32]byte, error) {
	return _ContractIAVSDirectory.Contract.OPERATORAVSREGISTRATIONTYPEHASH(&_ContractIAVSDirectory.CallOpts)
}

// OPERATORAVSREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0xd79aceab.
//
// Solidity: function OPERATOR_AVS_REGISTRATION_TYPEHASH() view returns(bytes32)
func (_ContractIAVSDirectory *ContractIAVSDirectoryCallerSession) OPERATORAVSREGISTRATIONTYPEHASH() ([32]byte, error) {
	return _ContractIAVSDirectory.Contract.OPERATORAVSREGISTRATIONTYPEHASH(&_ContractIAVSDirectory.CallOpts)
}

// OPERATORSETREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0xc825fe68.
//
// Solidity: function OPERATOR_SET_REGISTRATION_TYPEHASH() view returns(bytes32)
func (_ContractIAVSDirectory *ContractIAVSDirectoryCaller) OPERATORSETREGISTRATIONTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ContractIAVSDirectory.contract.Call(opts, &out, "OPERATOR_SET_REGISTRATION_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OPERATORSETREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0xc825fe68.
//
// Solidity: function OPERATOR_SET_REGISTRATION_TYPEHASH() view returns(bytes32)
func (_ContractIAVSDirectory *ContractIAVSDirectorySession) OPERATORSETREGISTRATIONTYPEHASH() ([32]byte, error) {
	return _ContractIAVSDirectory.Contract.OPERATORSETREGISTRATIONTYPEHASH(&_ContractIAVSDirectory.CallOpts)
}

// OPERATORSETREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0xc825fe68.
//
// Solidity: function OPERATOR_SET_REGISTRATION_TYPEHASH() view returns(bytes32)
func (_ContractIAVSDirectory *ContractIAVSDirectoryCallerSession) OPERATORSETREGISTRATIONTYPEHASH() ([32]byte, error) {
	return _ContractIAVSDirectory.Contract.OPERATORSETREGISTRATIONTYPEHASH(&_ContractIAVSDirectory.CallOpts)
}

// CalculateOperatorAVSRegistrationDigestHash is a free data retrieval call binding the contract method 0xa1060c88.
//
// Solidity: function calculateOperatorAVSRegistrationDigestHash(address operator, address avs, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_ContractIAVSDirectory *ContractIAVSDirectoryCaller) CalculateOperatorAVSRegistrationDigestHash(opts *bind.CallOpts, operator common.Address, avs common.Address, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _ContractIAVSDirectory.contract.Call(opts, &out, "calculateOperatorAVSRegistrationDigestHash", operator, avs, salt, expiry)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateOperatorAVSRegistrationDigestHash is a free data retrieval call binding the contract method 0xa1060c88.
//
// Solidity: function calculateOperatorAVSRegistrationDigestHash(address operator, address avs, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_ContractIAVSDirectory *ContractIAVSDirectorySession) CalculateOperatorAVSRegistrationDigestHash(operator common.Address, avs common.Address, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _ContractIAVSDirectory.Contract.CalculateOperatorAVSRegistrationDigestHash(&_ContractIAVSDirectory.CallOpts, operator, avs, salt, expiry)
}

// CalculateOperatorAVSRegistrationDigestHash is a free data retrieval call binding the contract method 0xa1060c88.
//
// Solidity: function calculateOperatorAVSRegistrationDigestHash(address operator, address avs, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_ContractIAVSDirectory *ContractIAVSDirectoryCallerSession) CalculateOperatorAVSRegistrationDigestHash(operator common.Address, avs common.Address, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _ContractIAVSDirectory.Contract.CalculateOperatorAVSRegistrationDigestHash(&_ContractIAVSDirectory.CallOpts, operator, avs, salt, expiry)
}

// CalculateOperatorSetForceDeregistrationTypehash is a free data retrieval call binding the contract method 0xb2841d48.
//
// Solidity: function calculateOperatorSetForceDeregistrationTypehash(address avs, uint32[] operatorSetIds, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_ContractIAVSDirectory *ContractIAVSDirectoryCaller) CalculateOperatorSetForceDeregistrationTypehash(opts *bind.CallOpts, avs common.Address, operatorSetIds []uint32, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _ContractIAVSDirectory.contract.Call(opts, &out, "calculateOperatorSetForceDeregistrationTypehash", avs, operatorSetIds, salt, expiry)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateOperatorSetForceDeregistrationTypehash is a free data retrieval call binding the contract method 0xb2841d48.
//
// Solidity: function calculateOperatorSetForceDeregistrationTypehash(address avs, uint32[] operatorSetIds, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_ContractIAVSDirectory *ContractIAVSDirectorySession) CalculateOperatorSetForceDeregistrationTypehash(avs common.Address, operatorSetIds []uint32, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _ContractIAVSDirectory.Contract.CalculateOperatorSetForceDeregistrationTypehash(&_ContractIAVSDirectory.CallOpts, avs, operatorSetIds, salt, expiry)
}

// CalculateOperatorSetForceDeregistrationTypehash is a free data retrieval call binding the contract method 0xb2841d48.
//
// Solidity: function calculateOperatorSetForceDeregistrationTypehash(address avs, uint32[] operatorSetIds, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_ContractIAVSDirectory *ContractIAVSDirectoryCallerSession) CalculateOperatorSetForceDeregistrationTypehash(avs common.Address, operatorSetIds []uint32, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _ContractIAVSDirectory.Contract.CalculateOperatorSetForceDeregistrationTypehash(&_ContractIAVSDirectory.CallOpts, avs, operatorSetIds, salt, expiry)
}

// CalculateOperatorSetRegistrationDigestHash is a free data retrieval call binding the contract method 0x955e6696.
//
// Solidity: function calculateOperatorSetRegistrationDigestHash(address avs, uint32[] operatorSetIds, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_ContractIAVSDirectory *ContractIAVSDirectoryCaller) CalculateOperatorSetRegistrationDigestHash(opts *bind.CallOpts, avs common.Address, operatorSetIds []uint32, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _ContractIAVSDirectory.contract.Call(opts, &out, "calculateOperatorSetRegistrationDigestHash", avs, operatorSetIds, salt, expiry)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateOperatorSetRegistrationDigestHash is a free data retrieval call binding the contract method 0x955e6696.
//
// Solidity: function calculateOperatorSetRegistrationDigestHash(address avs, uint32[] operatorSetIds, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_ContractIAVSDirectory *ContractIAVSDirectorySession) CalculateOperatorSetRegistrationDigestHash(avs common.Address, operatorSetIds []uint32, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _ContractIAVSDirectory.Contract.CalculateOperatorSetRegistrationDigestHash(&_ContractIAVSDirectory.CallOpts, avs, operatorSetIds, salt, expiry)
}

// CalculateOperatorSetRegistrationDigestHash is a free data retrieval call binding the contract method 0x955e6696.
//
// Solidity: function calculateOperatorSetRegistrationDigestHash(address avs, uint32[] operatorSetIds, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_ContractIAVSDirectory *ContractIAVSDirectoryCallerSession) CalculateOperatorSetRegistrationDigestHash(avs common.Address, operatorSetIds []uint32, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _ContractIAVSDirectory.Contract.CalculateOperatorSetRegistrationDigestHash(&_ContractIAVSDirectory.CallOpts, avs, operatorSetIds, salt, expiry)
}

// GetNumOperatorSetsOfOperator is a free data retrieval call binding the contract method 0xe88d8049.
//
// Solidity: function getNumOperatorSetsOfOperator(address operator) view returns(uint256)
func (_ContractIAVSDirectory *ContractIAVSDirectoryCaller) GetNumOperatorSetsOfOperator(opts *bind.CallOpts, operator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ContractIAVSDirectory.contract.Call(opts, &out, "getNumOperatorSetsOfOperator", operator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNumOperatorSetsOfOperator is a free data retrieval call binding the contract method 0xe88d8049.
//
// Solidity: function getNumOperatorSetsOfOperator(address operator) view returns(uint256)
func (_ContractIAVSDirectory *ContractIAVSDirectorySession) GetNumOperatorSetsOfOperator(operator common.Address) (*big.Int, error) {
	return _ContractIAVSDirectory.Contract.GetNumOperatorSetsOfOperator(&_ContractIAVSDirectory.CallOpts, operator)
}

// GetNumOperatorSetsOfOperator is a free data retrieval call binding the contract method 0xe88d8049.
//
// Solidity: function getNumOperatorSetsOfOperator(address operator) view returns(uint256)
func (_ContractIAVSDirectory *ContractIAVSDirectoryCallerSession) GetNumOperatorSetsOfOperator(operator common.Address) (*big.Int, error) {
	return _ContractIAVSDirectory.Contract.GetNumOperatorSetsOfOperator(&_ContractIAVSDirectory.CallOpts, operator)
}

// GetNumOperatorsInOperatorSet is a free data retrieval call binding the contract method 0x1023aa35.
//
// Solidity: function getNumOperatorsInOperatorSet((address,uint32) operatorSet) view returns(uint256)
func (_ContractIAVSDirectory *ContractIAVSDirectoryCaller) GetNumOperatorsInOperatorSet(opts *bind.CallOpts, operatorSet OperatorSet) (*big.Int, error) {
	var out []interface{}
	err := _ContractIAVSDirectory.contract.Call(opts, &out, "getNumOperatorsInOperatorSet", operatorSet)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNumOperatorsInOperatorSet is a free data retrieval call binding the contract method 0x1023aa35.
//
// Solidity: function getNumOperatorsInOperatorSet((address,uint32) operatorSet) view returns(uint256)
func (_ContractIAVSDirectory *ContractIAVSDirectorySession) GetNumOperatorsInOperatorSet(operatorSet OperatorSet) (*big.Int, error) {
	return _ContractIAVSDirectory.Contract.GetNumOperatorsInOperatorSet(&_ContractIAVSDirectory.CallOpts, operatorSet)
}

// GetNumOperatorsInOperatorSet is a free data retrieval call binding the contract method 0x1023aa35.
//
// Solidity: function getNumOperatorsInOperatorSet((address,uint32) operatorSet) view returns(uint256)
func (_ContractIAVSDirectory *ContractIAVSDirectoryCallerSession) GetNumOperatorsInOperatorSet(operatorSet OperatorSet) (*big.Int, error) {
	return _ContractIAVSDirectory.Contract.GetNumOperatorsInOperatorSet(&_ContractIAVSDirectory.CallOpts, operatorSet)
}

// GetOperatorSetsOfOperator is a free data retrieval call binding the contract method 0x16ae76cb.
//
// Solidity: function getOperatorSetsOfOperator(address operator, uint256 start, uint256 length) view returns((address,uint32)[] operatorSets)
func (_ContractIAVSDirectory *ContractIAVSDirectoryCaller) GetOperatorSetsOfOperator(opts *bind.CallOpts, operator common.Address, start *big.Int, length *big.Int) ([]OperatorSet, error) {
	var out []interface{}
	err := _ContractIAVSDirectory.contract.Call(opts, &out, "getOperatorSetsOfOperator", operator, start, length)

	if err != nil {
		return *new([]OperatorSet), err
	}

	out0 := *abi.ConvertType(out[0], new([]OperatorSet)).(*[]OperatorSet)

	return out0, err

}

// GetOperatorSetsOfOperator is a free data retrieval call binding the contract method 0x16ae76cb.
//
// Solidity: function getOperatorSetsOfOperator(address operator, uint256 start, uint256 length) view returns((address,uint32)[] operatorSets)
func (_ContractIAVSDirectory *ContractIAVSDirectorySession) GetOperatorSetsOfOperator(operator common.Address, start *big.Int, length *big.Int) ([]OperatorSet, error) {
	return _ContractIAVSDirectory.Contract.GetOperatorSetsOfOperator(&_ContractIAVSDirectory.CallOpts, operator, start, length)
}

// GetOperatorSetsOfOperator is a free data retrieval call binding the contract method 0x16ae76cb.
//
// Solidity: function getOperatorSetsOfOperator(address operator, uint256 start, uint256 length) view returns((address,uint32)[] operatorSets)
func (_ContractIAVSDirectory *ContractIAVSDirectoryCallerSession) GetOperatorSetsOfOperator(operator common.Address, start *big.Int, length *big.Int) ([]OperatorSet, error) {
	return _ContractIAVSDirectory.Contract.GetOperatorSetsOfOperator(&_ContractIAVSDirectory.CallOpts, operator, start, length)
}

// GetOperatorsInOperatorSet is a free data retrieval call binding the contract method 0x7357723b.
//
// Solidity: function getOperatorsInOperatorSet((address,uint32) operatorSet, uint256 start, uint256 length) view returns(address[] operators)
func (_ContractIAVSDirectory *ContractIAVSDirectoryCaller) GetOperatorsInOperatorSet(opts *bind.CallOpts, operatorSet OperatorSet, start *big.Int, length *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _ContractIAVSDirectory.contract.Call(opts, &out, "getOperatorsInOperatorSet", operatorSet, start, length)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOperatorsInOperatorSet is a free data retrieval call binding the contract method 0x7357723b.
//
// Solidity: function getOperatorsInOperatorSet((address,uint32) operatorSet, uint256 start, uint256 length) view returns(address[] operators)
func (_ContractIAVSDirectory *ContractIAVSDirectorySession) GetOperatorsInOperatorSet(operatorSet OperatorSet, start *big.Int, length *big.Int) ([]common.Address, error) {
	return _ContractIAVSDirectory.Contract.GetOperatorsInOperatorSet(&_ContractIAVSDirectory.CallOpts, operatorSet, start, length)
}

// GetOperatorsInOperatorSet is a free data retrieval call binding the contract method 0x7357723b.
//
// Solidity: function getOperatorsInOperatorSet((address,uint32) operatorSet, uint256 start, uint256 length) view returns(address[] operators)
func (_ContractIAVSDirectory *ContractIAVSDirectoryCallerSession) GetOperatorsInOperatorSet(operatorSet OperatorSet, start *big.Int, length *big.Int) ([]common.Address, error) {
	return _ContractIAVSDirectory.Contract.GetOperatorsInOperatorSet(&_ContractIAVSDirectory.CallOpts, operatorSet, start, length)
}

// GetStrategiesInOperatorSet is a free data retrieval call binding the contract method 0x4177a87c.
//
// Solidity: function getStrategiesInOperatorSet((address,uint32) operatorSet) view returns(address[] strategies)
func (_ContractIAVSDirectory *ContractIAVSDirectoryCaller) GetStrategiesInOperatorSet(opts *bind.CallOpts, operatorSet OperatorSet) ([]common.Address, error) {
	var out []interface{}
	err := _ContractIAVSDirectory.contract.Call(opts, &out, "getStrategiesInOperatorSet", operatorSet)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetStrategiesInOperatorSet is a free data retrieval call binding the contract method 0x4177a87c.
//
// Solidity: function getStrategiesInOperatorSet((address,uint32) operatorSet) view returns(address[] strategies)
func (_ContractIAVSDirectory *ContractIAVSDirectorySession) GetStrategiesInOperatorSet(operatorSet OperatorSet) ([]common.Address, error) {
	return _ContractIAVSDirectory.Contract.GetStrategiesInOperatorSet(&_ContractIAVSDirectory.CallOpts, operatorSet)
}

// GetStrategiesInOperatorSet is a free data retrieval call binding the contract method 0x4177a87c.
//
// Solidity: function getStrategiesInOperatorSet((address,uint32) operatorSet) view returns(address[] strategies)
func (_ContractIAVSDirectory *ContractIAVSDirectoryCallerSession) GetStrategiesInOperatorSet(operatorSet OperatorSet) ([]common.Address, error) {
	return _ContractIAVSDirectory.Contract.GetStrategiesInOperatorSet(&_ContractIAVSDirectory.CallOpts, operatorSet)
}

// InTotalOperatorSets is a free data retrieval call binding the contract method 0xcbdf0e42.
//
// Solidity: function inTotalOperatorSets(address operator) view returns(uint256)
func (_ContractIAVSDirectory *ContractIAVSDirectoryCaller) InTotalOperatorSets(opts *bind.CallOpts, operator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ContractIAVSDirectory.contract.Call(opts, &out, "inTotalOperatorSets", operator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InTotalOperatorSets is a free data retrieval call binding the contract method 0xcbdf0e42.
//
// Solidity: function inTotalOperatorSets(address operator) view returns(uint256)
func (_ContractIAVSDirectory *ContractIAVSDirectorySession) InTotalOperatorSets(operator common.Address) (*big.Int, error) {
	return _ContractIAVSDirectory.Contract.InTotalOperatorSets(&_ContractIAVSDirectory.CallOpts, operator)
}

// InTotalOperatorSets is a free data retrieval call binding the contract method 0xcbdf0e42.
//
// Solidity: function inTotalOperatorSets(address operator) view returns(uint256)
func (_ContractIAVSDirectory *ContractIAVSDirectoryCallerSession) InTotalOperatorSets(operator common.Address) (*big.Int, error) {
	return _ContractIAVSDirectory.Contract.InTotalOperatorSets(&_ContractIAVSDirectory.CallOpts, operator)
}

// IsMember is a free data retrieval call binding the contract method 0xda2ff05d.
//
// Solidity: function isMember(address operator, (address,uint32) operatorSet) view returns(bool)
func (_ContractIAVSDirectory *ContractIAVSDirectoryCaller) IsMember(opts *bind.CallOpts, operator common.Address, operatorSet OperatorSet) (bool, error) {
	var out []interface{}
	err := _ContractIAVSDirectory.contract.Call(opts, &out, "isMember", operator, operatorSet)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMember is a free data retrieval call binding the contract method 0xda2ff05d.
//
// Solidity: function isMember(address operator, (address,uint32) operatorSet) view returns(bool)
func (_ContractIAVSDirectory *ContractIAVSDirectorySession) IsMember(operator common.Address, operatorSet OperatorSet) (bool, error) {
	return _ContractIAVSDirectory.Contract.IsMember(&_ContractIAVSDirectory.CallOpts, operator, operatorSet)
}

// IsMember is a free data retrieval call binding the contract method 0xda2ff05d.
//
// Solidity: function isMember(address operator, (address,uint32) operatorSet) view returns(bool)
func (_ContractIAVSDirectory *ContractIAVSDirectoryCallerSession) IsMember(operator common.Address, operatorSet OperatorSet) (bool, error) {
	return _ContractIAVSDirectory.Contract.IsMember(&_ContractIAVSDirectory.CallOpts, operator, operatorSet)
}

// IsOperatorSet is a free data retrieval call binding the contract method 0x84d76f7b.
//
// Solidity: function isOperatorSet(address avs, uint32 operatorSetId) view returns(bool)
func (_ContractIAVSDirectory *ContractIAVSDirectoryCaller) IsOperatorSet(opts *bind.CallOpts, avs common.Address, operatorSetId uint32) (bool, error) {
	var out []interface{}
	err := _ContractIAVSDirectory.contract.Call(opts, &out, "isOperatorSet", avs, operatorSetId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperatorSet is a free data retrieval call binding the contract method 0x84d76f7b.
//
// Solidity: function isOperatorSet(address avs, uint32 operatorSetId) view returns(bool)
func (_ContractIAVSDirectory *ContractIAVSDirectorySession) IsOperatorSet(avs common.Address, operatorSetId uint32) (bool, error) {
	return _ContractIAVSDirectory.Contract.IsOperatorSet(&_ContractIAVSDirectory.CallOpts, avs, operatorSetId)
}

// IsOperatorSet is a free data retrieval call binding the contract method 0x84d76f7b.
//
// Solidity: function isOperatorSet(address avs, uint32 operatorSetId) view returns(bool)
func (_ContractIAVSDirectory *ContractIAVSDirectoryCallerSession) IsOperatorSet(avs common.Address, operatorSetId uint32) (bool, error) {
	return _ContractIAVSDirectory.Contract.IsOperatorSet(&_ContractIAVSDirectory.CallOpts, avs, operatorSetId)
}

// IsOperatorSetAVS is a free data retrieval call binding the contract method 0x7673e93a.
//
// Solidity: function isOperatorSetAVS(address avs) view returns(bool)
func (_ContractIAVSDirectory *ContractIAVSDirectoryCaller) IsOperatorSetAVS(opts *bind.CallOpts, avs common.Address) (bool, error) {
	var out []interface{}
	err := _ContractIAVSDirectory.contract.Call(opts, &out, "isOperatorSetAVS", avs)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperatorSetAVS is a free data retrieval call binding the contract method 0x7673e93a.
//
// Solidity: function isOperatorSetAVS(address avs) view returns(bool)
func (_ContractIAVSDirectory *ContractIAVSDirectorySession) IsOperatorSetAVS(avs common.Address) (bool, error) {
	return _ContractIAVSDirectory.Contract.IsOperatorSetAVS(&_ContractIAVSDirectory.CallOpts, avs)
}

// IsOperatorSetAVS is a free data retrieval call binding the contract method 0x7673e93a.
//
// Solidity: function isOperatorSetAVS(address avs) view returns(bool)
func (_ContractIAVSDirectory *ContractIAVSDirectoryCallerSession) IsOperatorSetAVS(avs common.Address) (bool, error) {
	return _ContractIAVSDirectory.Contract.IsOperatorSetAVS(&_ContractIAVSDirectory.CallOpts, avs)
}

// IsOperatorSetBatch is a free data retrieval call binding the contract method 0x20c4e236.
//
// Solidity: function isOperatorSetBatch((address,uint32)[] operatorSets) view returns(bool)
func (_ContractIAVSDirectory *ContractIAVSDirectoryCaller) IsOperatorSetBatch(opts *bind.CallOpts, operatorSets []OperatorSet) (bool, error) {
	var out []interface{}
	err := _ContractIAVSDirectory.contract.Call(opts, &out, "isOperatorSetBatch", operatorSets)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperatorSetBatch is a free data retrieval call binding the contract method 0x20c4e236.
//
// Solidity: function isOperatorSetBatch((address,uint32)[] operatorSets) view returns(bool)
func (_ContractIAVSDirectory *ContractIAVSDirectorySession) IsOperatorSetBatch(operatorSets []OperatorSet) (bool, error) {
	return _ContractIAVSDirectory.Contract.IsOperatorSetBatch(&_ContractIAVSDirectory.CallOpts, operatorSets)
}

// IsOperatorSetBatch is a free data retrieval call binding the contract method 0x20c4e236.
//
// Solidity: function isOperatorSetBatch((address,uint32)[] operatorSets) view returns(bool)
func (_ContractIAVSDirectory *ContractIAVSDirectoryCallerSession) IsOperatorSetBatch(operatorSets []OperatorSet) (bool, error) {
	return _ContractIAVSDirectory.Contract.IsOperatorSetBatch(&_ContractIAVSDirectory.CallOpts, operatorSets)
}

// IsOperatorSlashable is a free data retrieval call binding the contract method 0x1352c3e6.
//
// Solidity: function isOperatorSlashable(address operator, (address,uint32) operatorSet) view returns(bool)
func (_ContractIAVSDirectory *ContractIAVSDirectoryCaller) IsOperatorSlashable(opts *bind.CallOpts, operator common.Address, operatorSet OperatorSet) (bool, error) {
	var out []interface{}
	err := _ContractIAVSDirectory.contract.Call(opts, &out, "isOperatorSlashable", operator, operatorSet)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperatorSlashable is a free data retrieval call binding the contract method 0x1352c3e6.
//
// Solidity: function isOperatorSlashable(address operator, (address,uint32) operatorSet) view returns(bool)
func (_ContractIAVSDirectory *ContractIAVSDirectorySession) IsOperatorSlashable(operator common.Address, operatorSet OperatorSet) (bool, error) {
	return _ContractIAVSDirectory.Contract.IsOperatorSlashable(&_ContractIAVSDirectory.CallOpts, operator, operatorSet)
}

// IsOperatorSlashable is a free data retrieval call binding the contract method 0x1352c3e6.
//
// Solidity: function isOperatorSlashable(address operator, (address,uint32) operatorSet) view returns(bool)
func (_ContractIAVSDirectory *ContractIAVSDirectoryCallerSession) IsOperatorSlashable(operator common.Address, operatorSet OperatorSet) (bool, error) {
	return _ContractIAVSDirectory.Contract.IsOperatorSlashable(&_ContractIAVSDirectory.CallOpts, operator, operatorSet)
}

// OperatorSaltIsSpent is a free data retrieval call binding the contract method 0x374823b5.
//
// Solidity: function operatorSaltIsSpent(address operator, bytes32 salt) view returns(bool)
func (_ContractIAVSDirectory *ContractIAVSDirectoryCaller) OperatorSaltIsSpent(opts *bind.CallOpts, operator common.Address, salt [32]byte) (bool, error) {
	var out []interface{}
	err := _ContractIAVSDirectory.contract.Call(opts, &out, "operatorSaltIsSpent", operator, salt)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OperatorSaltIsSpent is a free data retrieval call binding the contract method 0x374823b5.
//
// Solidity: function operatorSaltIsSpent(address operator, bytes32 salt) view returns(bool)
func (_ContractIAVSDirectory *ContractIAVSDirectorySession) OperatorSaltIsSpent(operator common.Address, salt [32]byte) (bool, error) {
	return _ContractIAVSDirectory.Contract.OperatorSaltIsSpent(&_ContractIAVSDirectory.CallOpts, operator, salt)
}

// OperatorSaltIsSpent is a free data retrieval call binding the contract method 0x374823b5.
//
// Solidity: function operatorSaltIsSpent(address operator, bytes32 salt) view returns(bool)
func (_ContractIAVSDirectory *ContractIAVSDirectoryCallerSession) OperatorSaltIsSpent(operator common.Address, salt [32]byte) (bool, error) {
	return _ContractIAVSDirectory.Contract.OperatorSaltIsSpent(&_ContractIAVSDirectory.CallOpts, operator, salt)
}

// OperatorSetMemberAtIndex is a free data retrieval call binding the contract method 0x411d415b.
//
// Solidity: function operatorSetMemberAtIndex((address,uint32) operatorSet, uint256 index) view returns(address)
func (_ContractIAVSDirectory *ContractIAVSDirectoryCaller) OperatorSetMemberAtIndex(opts *bind.CallOpts, operatorSet OperatorSet, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ContractIAVSDirectory.contract.Call(opts, &out, "operatorSetMemberAtIndex", operatorSet, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OperatorSetMemberAtIndex is a free data retrieval call binding the contract method 0x411d415b.
//
// Solidity: function operatorSetMemberAtIndex((address,uint32) operatorSet, uint256 index) view returns(address)
func (_ContractIAVSDirectory *ContractIAVSDirectorySession) OperatorSetMemberAtIndex(operatorSet OperatorSet, index *big.Int) (common.Address, error) {
	return _ContractIAVSDirectory.Contract.OperatorSetMemberAtIndex(&_ContractIAVSDirectory.CallOpts, operatorSet, index)
}

// OperatorSetMemberAtIndex is a free data retrieval call binding the contract method 0x411d415b.
//
// Solidity: function operatorSetMemberAtIndex((address,uint32) operatorSet, uint256 index) view returns(address)
func (_ContractIAVSDirectory *ContractIAVSDirectoryCallerSession) OperatorSetMemberAtIndex(operatorSet OperatorSet, index *big.Int) (common.Address, error) {
	return _ContractIAVSDirectory.Contract.OperatorSetMemberAtIndex(&_ContractIAVSDirectory.CallOpts, operatorSet, index)
}

// OperatorSetStatus is a free data retrieval call binding the contract method 0x1e68134e.
//
// Solidity: function operatorSetStatus(address avs, address operator, uint32 operatorSetId) view returns(bool registered, uint32 lastDeregisteredTimestamp)
func (_ContractIAVSDirectory *ContractIAVSDirectoryCaller) OperatorSetStatus(opts *bind.CallOpts, avs common.Address, operator common.Address, operatorSetId uint32) (struct {
	Registered                bool
	LastDeregisteredTimestamp uint32
}, error) {
	var out []interface{}
	err := _ContractIAVSDirectory.contract.Call(opts, &out, "operatorSetStatus", avs, operator, operatorSetId)

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
// Solidity: function operatorSetStatus(address avs, address operator, uint32 operatorSetId) view returns(bool registered, uint32 lastDeregisteredTimestamp)
func (_ContractIAVSDirectory *ContractIAVSDirectorySession) OperatorSetStatus(avs common.Address, operator common.Address, operatorSetId uint32) (struct {
	Registered                bool
	LastDeregisteredTimestamp uint32
}, error) {
	return _ContractIAVSDirectory.Contract.OperatorSetStatus(&_ContractIAVSDirectory.CallOpts, avs, operator, operatorSetId)
}

// OperatorSetStatus is a free data retrieval call binding the contract method 0x1e68134e.
//
// Solidity: function operatorSetStatus(address avs, address operator, uint32 operatorSetId) view returns(bool registered, uint32 lastDeregisteredTimestamp)
func (_ContractIAVSDirectory *ContractIAVSDirectoryCallerSession) OperatorSetStatus(avs common.Address, operator common.Address, operatorSetId uint32) (struct {
	Registered                bool
	LastDeregisteredTimestamp uint32
}, error) {
	return _ContractIAVSDirectory.Contract.OperatorSetStatus(&_ContractIAVSDirectory.CallOpts, avs, operator, operatorSetId)
}

// OperatorSetsMemberOfAtIndex is a free data retrieval call binding the contract method 0xb5a768ca.
//
// Solidity: function operatorSetsMemberOfAtIndex(address operator, uint256 index) view returns((address,uint32))
func (_ContractIAVSDirectory *ContractIAVSDirectoryCaller) OperatorSetsMemberOfAtIndex(opts *bind.CallOpts, operator common.Address, index *big.Int) (OperatorSet, error) {
	var out []interface{}
	err := _ContractIAVSDirectory.contract.Call(opts, &out, "operatorSetsMemberOfAtIndex", operator, index)

	if err != nil {
		return *new(OperatorSet), err
	}

	out0 := *abi.ConvertType(out[0], new(OperatorSet)).(*OperatorSet)

	return out0, err

}

// OperatorSetsMemberOfAtIndex is a free data retrieval call binding the contract method 0xb5a768ca.
//
// Solidity: function operatorSetsMemberOfAtIndex(address operator, uint256 index) view returns((address,uint32))
func (_ContractIAVSDirectory *ContractIAVSDirectorySession) OperatorSetsMemberOfAtIndex(operator common.Address, index *big.Int) (OperatorSet, error) {
	return _ContractIAVSDirectory.Contract.OperatorSetsMemberOfAtIndex(&_ContractIAVSDirectory.CallOpts, operator, index)
}

// OperatorSetsMemberOfAtIndex is a free data retrieval call binding the contract method 0xb5a768ca.
//
// Solidity: function operatorSetsMemberOfAtIndex(address operator, uint256 index) view returns((address,uint32))
func (_ContractIAVSDirectory *ContractIAVSDirectoryCallerSession) OperatorSetsMemberOfAtIndex(operator common.Address, index *big.Int) (OperatorSet, error) {
	return _ContractIAVSDirectory.Contract.OperatorSetsMemberOfAtIndex(&_ContractIAVSDirectory.CallOpts, operator, index)
}

// AddStrategiesToOperatorSet is a paid mutator transaction binding the contract method 0x76999342.
//
// Solidity: function addStrategiesToOperatorSet(uint32 operatorSetId, address[] strategies) returns()
func (_ContractIAVSDirectory *ContractIAVSDirectoryTransactor) AddStrategiesToOperatorSet(opts *bind.TransactOpts, operatorSetId uint32, strategies []common.Address) (*types.Transaction, error) {
	return _ContractIAVSDirectory.contract.Transact(opts, "addStrategiesToOperatorSet", operatorSetId, strategies)
}

// AddStrategiesToOperatorSet is a paid mutator transaction binding the contract method 0x76999342.
//
// Solidity: function addStrategiesToOperatorSet(uint32 operatorSetId, address[] strategies) returns()
func (_ContractIAVSDirectory *ContractIAVSDirectorySession) AddStrategiesToOperatorSet(operatorSetId uint32, strategies []common.Address) (*types.Transaction, error) {
	return _ContractIAVSDirectory.Contract.AddStrategiesToOperatorSet(&_ContractIAVSDirectory.TransactOpts, operatorSetId, strategies)
}

// AddStrategiesToOperatorSet is a paid mutator transaction binding the contract method 0x76999342.
//
// Solidity: function addStrategiesToOperatorSet(uint32 operatorSetId, address[] strategies) returns()
func (_ContractIAVSDirectory *ContractIAVSDirectoryTransactorSession) AddStrategiesToOperatorSet(operatorSetId uint32, strategies []common.Address) (*types.Transaction, error) {
	return _ContractIAVSDirectory.Contract.AddStrategiesToOperatorSet(&_ContractIAVSDirectory.TransactOpts, operatorSetId, strategies)
}

// BecomeOperatorSetAVS is a paid mutator transaction binding the contract method 0xaec205c5.
//
// Solidity: function becomeOperatorSetAVS() returns()
func (_ContractIAVSDirectory *ContractIAVSDirectoryTransactor) BecomeOperatorSetAVS(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractIAVSDirectory.contract.Transact(opts, "becomeOperatorSetAVS")
}

// BecomeOperatorSetAVS is a paid mutator transaction binding the contract method 0xaec205c5.
//
// Solidity: function becomeOperatorSetAVS() returns()
func (_ContractIAVSDirectory *ContractIAVSDirectorySession) BecomeOperatorSetAVS() (*types.Transaction, error) {
	return _ContractIAVSDirectory.Contract.BecomeOperatorSetAVS(&_ContractIAVSDirectory.TransactOpts)
}

// BecomeOperatorSetAVS is a paid mutator transaction binding the contract method 0xaec205c5.
//
// Solidity: function becomeOperatorSetAVS() returns()
func (_ContractIAVSDirectory *ContractIAVSDirectoryTransactorSession) BecomeOperatorSetAVS() (*types.Transaction, error) {
	return _ContractIAVSDirectory.Contract.BecomeOperatorSetAVS(&_ContractIAVSDirectory.TransactOpts)
}

// CancelSalt is a paid mutator transaction binding the contract method 0xec76f442.
//
// Solidity: function cancelSalt(bytes32 salt) returns()
func (_ContractIAVSDirectory *ContractIAVSDirectoryTransactor) CancelSalt(opts *bind.TransactOpts, salt [32]byte) (*types.Transaction, error) {
	return _ContractIAVSDirectory.contract.Transact(opts, "cancelSalt", salt)
}

// CancelSalt is a paid mutator transaction binding the contract method 0xec76f442.
//
// Solidity: function cancelSalt(bytes32 salt) returns()
func (_ContractIAVSDirectory *ContractIAVSDirectorySession) CancelSalt(salt [32]byte) (*types.Transaction, error) {
	return _ContractIAVSDirectory.Contract.CancelSalt(&_ContractIAVSDirectory.TransactOpts, salt)
}

// CancelSalt is a paid mutator transaction binding the contract method 0xec76f442.
//
// Solidity: function cancelSalt(bytes32 salt) returns()
func (_ContractIAVSDirectory *ContractIAVSDirectoryTransactorSession) CancelSalt(salt [32]byte) (*types.Transaction, error) {
	return _ContractIAVSDirectory.Contract.CancelSalt(&_ContractIAVSDirectory.TransactOpts, salt)
}

// CreateOperatorSets is a paid mutator transaction binding the contract method 0xafe02ed5.
//
// Solidity: function createOperatorSets(uint32[] operatorSetIds) returns()
func (_ContractIAVSDirectory *ContractIAVSDirectoryTransactor) CreateOperatorSets(opts *bind.TransactOpts, operatorSetIds []uint32) (*types.Transaction, error) {
	return _ContractIAVSDirectory.contract.Transact(opts, "createOperatorSets", operatorSetIds)
}

// CreateOperatorSets is a paid mutator transaction binding the contract method 0xafe02ed5.
//
// Solidity: function createOperatorSets(uint32[] operatorSetIds) returns()
func (_ContractIAVSDirectory *ContractIAVSDirectorySession) CreateOperatorSets(operatorSetIds []uint32) (*types.Transaction, error) {
	return _ContractIAVSDirectory.Contract.CreateOperatorSets(&_ContractIAVSDirectory.TransactOpts, operatorSetIds)
}

// CreateOperatorSets is a paid mutator transaction binding the contract method 0xafe02ed5.
//
// Solidity: function createOperatorSets(uint32[] operatorSetIds) returns()
func (_ContractIAVSDirectory *ContractIAVSDirectoryTransactorSession) CreateOperatorSets(operatorSetIds []uint32) (*types.Transaction, error) {
	return _ContractIAVSDirectory.Contract.CreateOperatorSets(&_ContractIAVSDirectory.TransactOpts, operatorSetIds)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_ContractIAVSDirectory *ContractIAVSDirectoryTransactor) DeregisterOperatorFromAVS(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _ContractIAVSDirectory.contract.Transact(opts, "deregisterOperatorFromAVS", operator)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_ContractIAVSDirectory *ContractIAVSDirectorySession) DeregisterOperatorFromAVS(operator common.Address) (*types.Transaction, error) {
	return _ContractIAVSDirectory.Contract.DeregisterOperatorFromAVS(&_ContractIAVSDirectory.TransactOpts, operator)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_ContractIAVSDirectory *ContractIAVSDirectoryTransactorSession) DeregisterOperatorFromAVS(operator common.Address) (*types.Transaction, error) {
	return _ContractIAVSDirectory.Contract.DeregisterOperatorFromAVS(&_ContractIAVSDirectory.TransactOpts, operator)
}

// DeregisterOperatorFromOperatorSets is a paid mutator transaction binding the contract method 0xc1a8e2c5.
//
// Solidity: function deregisterOperatorFromOperatorSets(address operator, uint32[] operatorSetIds) returns()
func (_ContractIAVSDirectory *ContractIAVSDirectoryTransactor) DeregisterOperatorFromOperatorSets(opts *bind.TransactOpts, operator common.Address, operatorSetIds []uint32) (*types.Transaction, error) {
	return _ContractIAVSDirectory.contract.Transact(opts, "deregisterOperatorFromOperatorSets", operator, operatorSetIds)
}

// DeregisterOperatorFromOperatorSets is a paid mutator transaction binding the contract method 0xc1a8e2c5.
//
// Solidity: function deregisterOperatorFromOperatorSets(address operator, uint32[] operatorSetIds) returns()
func (_ContractIAVSDirectory *ContractIAVSDirectorySession) DeregisterOperatorFromOperatorSets(operator common.Address, operatorSetIds []uint32) (*types.Transaction, error) {
	return _ContractIAVSDirectory.Contract.DeregisterOperatorFromOperatorSets(&_ContractIAVSDirectory.TransactOpts, operator, operatorSetIds)
}

// DeregisterOperatorFromOperatorSets is a paid mutator transaction binding the contract method 0xc1a8e2c5.
//
// Solidity: function deregisterOperatorFromOperatorSets(address operator, uint32[] operatorSetIds) returns()
func (_ContractIAVSDirectory *ContractIAVSDirectoryTransactorSession) DeregisterOperatorFromOperatorSets(operator common.Address, operatorSetIds []uint32) (*types.Transaction, error) {
	return _ContractIAVSDirectory.Contract.DeregisterOperatorFromOperatorSets(&_ContractIAVSDirectory.TransactOpts, operator, operatorSetIds)
}

// ForceDeregisterFromOperatorSets is a paid mutator transaction binding the contract method 0x3fee332d.
//
// Solidity: function forceDeregisterFromOperatorSets(address operator, address avs, uint32[] operatorSetIds, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractIAVSDirectory *ContractIAVSDirectoryTransactor) ForceDeregisterFromOperatorSets(opts *bind.TransactOpts, operator common.Address, avs common.Address, operatorSetIds []uint32, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractIAVSDirectory.contract.Transact(opts, "forceDeregisterFromOperatorSets", operator, avs, operatorSetIds, operatorSignature)
}

// ForceDeregisterFromOperatorSets is a paid mutator transaction binding the contract method 0x3fee332d.
//
// Solidity: function forceDeregisterFromOperatorSets(address operator, address avs, uint32[] operatorSetIds, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractIAVSDirectory *ContractIAVSDirectorySession) ForceDeregisterFromOperatorSets(operator common.Address, avs common.Address, operatorSetIds []uint32, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractIAVSDirectory.Contract.ForceDeregisterFromOperatorSets(&_ContractIAVSDirectory.TransactOpts, operator, avs, operatorSetIds, operatorSignature)
}

// ForceDeregisterFromOperatorSets is a paid mutator transaction binding the contract method 0x3fee332d.
//
// Solidity: function forceDeregisterFromOperatorSets(address operator, address avs, uint32[] operatorSetIds, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractIAVSDirectory *ContractIAVSDirectoryTransactorSession) ForceDeregisterFromOperatorSets(operator common.Address, avs common.Address, operatorSetIds []uint32, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractIAVSDirectory.Contract.ForceDeregisterFromOperatorSets(&_ContractIAVSDirectory.TransactOpts, operator, avs, operatorSetIds, operatorSignature)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address initialOwner, address _pauserRegistry, uint256 initialPausedStatus) returns()
func (_ContractIAVSDirectory *ContractIAVSDirectoryTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address, _pauserRegistry common.Address, initialPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractIAVSDirectory.contract.Transact(opts, "initialize", initialOwner, _pauserRegistry, initialPausedStatus)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address initialOwner, address _pauserRegistry, uint256 initialPausedStatus) returns()
func (_ContractIAVSDirectory *ContractIAVSDirectorySession) Initialize(initialOwner common.Address, _pauserRegistry common.Address, initialPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractIAVSDirectory.Contract.Initialize(&_ContractIAVSDirectory.TransactOpts, initialOwner, _pauserRegistry, initialPausedStatus)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address initialOwner, address _pauserRegistry, uint256 initialPausedStatus) returns()
func (_ContractIAVSDirectory *ContractIAVSDirectoryTransactorSession) Initialize(initialOwner common.Address, _pauserRegistry common.Address, initialPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractIAVSDirectory.Contract.Initialize(&_ContractIAVSDirectory.TransactOpts, initialOwner, _pauserRegistry, initialPausedStatus)
}

// MigrateOperatorsToOperatorSets is a paid mutator transaction binding the contract method 0xef2dfa8d.
//
// Solidity: function migrateOperatorsToOperatorSets(address[] operators, uint32[][] operatorSetIds) returns()
func (_ContractIAVSDirectory *ContractIAVSDirectoryTransactor) MigrateOperatorsToOperatorSets(opts *bind.TransactOpts, operators []common.Address, operatorSetIds [][]uint32) (*types.Transaction, error) {
	return _ContractIAVSDirectory.contract.Transact(opts, "migrateOperatorsToOperatorSets", operators, operatorSetIds)
}

// MigrateOperatorsToOperatorSets is a paid mutator transaction binding the contract method 0xef2dfa8d.
//
// Solidity: function migrateOperatorsToOperatorSets(address[] operators, uint32[][] operatorSetIds) returns()
func (_ContractIAVSDirectory *ContractIAVSDirectorySession) MigrateOperatorsToOperatorSets(operators []common.Address, operatorSetIds [][]uint32) (*types.Transaction, error) {
	return _ContractIAVSDirectory.Contract.MigrateOperatorsToOperatorSets(&_ContractIAVSDirectory.TransactOpts, operators, operatorSetIds)
}

// MigrateOperatorsToOperatorSets is a paid mutator transaction binding the contract method 0xef2dfa8d.
//
// Solidity: function migrateOperatorsToOperatorSets(address[] operators, uint32[][] operatorSetIds) returns()
func (_ContractIAVSDirectory *ContractIAVSDirectoryTransactorSession) MigrateOperatorsToOperatorSets(operators []common.Address, operatorSetIds [][]uint32) (*types.Transaction, error) {
	return _ContractIAVSDirectory.Contract.MigrateOperatorsToOperatorSets(&_ContractIAVSDirectory.TransactOpts, operators, operatorSetIds)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractIAVSDirectory *ContractIAVSDirectoryTransactor) RegisterOperatorToAVS(opts *bind.TransactOpts, operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractIAVSDirectory.contract.Transact(opts, "registerOperatorToAVS", operator, operatorSignature)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractIAVSDirectory *ContractIAVSDirectorySession) RegisterOperatorToAVS(operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractIAVSDirectory.Contract.RegisterOperatorToAVS(&_ContractIAVSDirectory.TransactOpts, operator, operatorSignature)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractIAVSDirectory *ContractIAVSDirectoryTransactorSession) RegisterOperatorToAVS(operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractIAVSDirectory.Contract.RegisterOperatorToAVS(&_ContractIAVSDirectory.TransactOpts, operator, operatorSignature)
}

// RegisterOperatorToOperatorSets is a paid mutator transaction binding the contract method 0x1e2199e2.
//
// Solidity: function registerOperatorToOperatorSets(address operator, uint32[] operatorSetIds, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractIAVSDirectory *ContractIAVSDirectoryTransactor) RegisterOperatorToOperatorSets(opts *bind.TransactOpts, operator common.Address, operatorSetIds []uint32, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractIAVSDirectory.contract.Transact(opts, "registerOperatorToOperatorSets", operator, operatorSetIds, operatorSignature)
}

// RegisterOperatorToOperatorSets is a paid mutator transaction binding the contract method 0x1e2199e2.
//
// Solidity: function registerOperatorToOperatorSets(address operator, uint32[] operatorSetIds, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractIAVSDirectory *ContractIAVSDirectorySession) RegisterOperatorToOperatorSets(operator common.Address, operatorSetIds []uint32, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractIAVSDirectory.Contract.RegisterOperatorToOperatorSets(&_ContractIAVSDirectory.TransactOpts, operator, operatorSetIds, operatorSignature)
}

// RegisterOperatorToOperatorSets is a paid mutator transaction binding the contract method 0x1e2199e2.
//
// Solidity: function registerOperatorToOperatorSets(address operator, uint32[] operatorSetIds, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractIAVSDirectory *ContractIAVSDirectoryTransactorSession) RegisterOperatorToOperatorSets(operator common.Address, operatorSetIds []uint32, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractIAVSDirectory.Contract.RegisterOperatorToOperatorSets(&_ContractIAVSDirectory.TransactOpts, operator, operatorSetIds, operatorSignature)
}

// RemoveStrategiesFromOperatorSet is a paid mutator transaction binding the contract method 0xce7b5e4b.
//
// Solidity: function removeStrategiesFromOperatorSet(uint32 operatorSetId, address[] strategies) returns()
func (_ContractIAVSDirectory *ContractIAVSDirectoryTransactor) RemoveStrategiesFromOperatorSet(opts *bind.TransactOpts, operatorSetId uint32, strategies []common.Address) (*types.Transaction, error) {
	return _ContractIAVSDirectory.contract.Transact(opts, "removeStrategiesFromOperatorSet", operatorSetId, strategies)
}

// RemoveStrategiesFromOperatorSet is a paid mutator transaction binding the contract method 0xce7b5e4b.
//
// Solidity: function removeStrategiesFromOperatorSet(uint32 operatorSetId, address[] strategies) returns()
func (_ContractIAVSDirectory *ContractIAVSDirectorySession) RemoveStrategiesFromOperatorSet(operatorSetId uint32, strategies []common.Address) (*types.Transaction, error) {
	return _ContractIAVSDirectory.Contract.RemoveStrategiesFromOperatorSet(&_ContractIAVSDirectory.TransactOpts, operatorSetId, strategies)
}

// RemoveStrategiesFromOperatorSet is a paid mutator transaction binding the contract method 0xce7b5e4b.
//
// Solidity: function removeStrategiesFromOperatorSet(uint32 operatorSetId, address[] strategies) returns()
func (_ContractIAVSDirectory *ContractIAVSDirectoryTransactorSession) RemoveStrategiesFromOperatorSet(operatorSetId uint32, strategies []common.Address) (*types.Transaction, error) {
	return _ContractIAVSDirectory.Contract.RemoveStrategiesFromOperatorSet(&_ContractIAVSDirectory.TransactOpts, operatorSetId, strategies)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string metadataURI) returns()
func (_ContractIAVSDirectory *ContractIAVSDirectoryTransactor) UpdateAVSMetadataURI(opts *bind.TransactOpts, metadataURI string) (*types.Transaction, error) {
	return _ContractIAVSDirectory.contract.Transact(opts, "updateAVSMetadataURI", metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string metadataURI) returns()
func (_ContractIAVSDirectory *ContractIAVSDirectorySession) UpdateAVSMetadataURI(metadataURI string) (*types.Transaction, error) {
	return _ContractIAVSDirectory.Contract.UpdateAVSMetadataURI(&_ContractIAVSDirectory.TransactOpts, metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string metadataURI) returns()
func (_ContractIAVSDirectory *ContractIAVSDirectoryTransactorSession) UpdateAVSMetadataURI(metadataURI string) (*types.Transaction, error) {
	return _ContractIAVSDirectory.Contract.UpdateAVSMetadataURI(&_ContractIAVSDirectory.TransactOpts, metadataURI)
}

// ContractIAVSDirectoryAVSMetadataURIUpdatedIterator is returned from FilterAVSMetadataURIUpdated and is used to iterate over the raw logs and unpacked data for AVSMetadataURIUpdated events raised by the ContractIAVSDirectory contract.
type ContractIAVSDirectoryAVSMetadataURIUpdatedIterator struct {
	Event *ContractIAVSDirectoryAVSMetadataURIUpdated // Event containing the contract specifics and raw log

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
func (it *ContractIAVSDirectoryAVSMetadataURIUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIAVSDirectoryAVSMetadataURIUpdated)
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
		it.Event = new(ContractIAVSDirectoryAVSMetadataURIUpdated)
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
func (it *ContractIAVSDirectoryAVSMetadataURIUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIAVSDirectoryAVSMetadataURIUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIAVSDirectoryAVSMetadataURIUpdated represents a AVSMetadataURIUpdated event raised by the ContractIAVSDirectory contract.
type ContractIAVSDirectoryAVSMetadataURIUpdated struct {
	Avs         common.Address
	MetadataURI string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAVSMetadataURIUpdated is a free log retrieval operation binding the contract event 0xa89c1dc243d8908a96dd84944bcc97d6bc6ac00dd78e20621576be6a3c943713.
//
// Solidity: event AVSMetadataURIUpdated(address indexed avs, string metadataURI)
func (_ContractIAVSDirectory *ContractIAVSDirectoryFilterer) FilterAVSMetadataURIUpdated(opts *bind.FilterOpts, avs []common.Address) (*ContractIAVSDirectoryAVSMetadataURIUpdatedIterator, error) {

	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _ContractIAVSDirectory.contract.FilterLogs(opts, "AVSMetadataURIUpdated", avsRule)
	if err != nil {
		return nil, err
	}
	return &ContractIAVSDirectoryAVSMetadataURIUpdatedIterator{contract: _ContractIAVSDirectory.contract, event: "AVSMetadataURIUpdated", logs: logs, sub: sub}, nil
}

// WatchAVSMetadataURIUpdated is a free log subscription operation binding the contract event 0xa89c1dc243d8908a96dd84944bcc97d6bc6ac00dd78e20621576be6a3c943713.
//
// Solidity: event AVSMetadataURIUpdated(address indexed avs, string metadataURI)
func (_ContractIAVSDirectory *ContractIAVSDirectoryFilterer) WatchAVSMetadataURIUpdated(opts *bind.WatchOpts, sink chan<- *ContractIAVSDirectoryAVSMetadataURIUpdated, avs []common.Address) (event.Subscription, error) {

	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _ContractIAVSDirectory.contract.WatchLogs(opts, "AVSMetadataURIUpdated", avsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIAVSDirectoryAVSMetadataURIUpdated)
				if err := _ContractIAVSDirectory.contract.UnpackLog(event, "AVSMetadataURIUpdated", log); err != nil {
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
func (_ContractIAVSDirectory *ContractIAVSDirectoryFilterer) ParseAVSMetadataURIUpdated(log types.Log) (*ContractIAVSDirectoryAVSMetadataURIUpdated, error) {
	event := new(ContractIAVSDirectoryAVSMetadataURIUpdated)
	if err := _ContractIAVSDirectory.contract.UnpackLog(event, "AVSMetadataURIUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIAVSDirectoryAVSMigratedToOperatorSetsIterator is returned from FilterAVSMigratedToOperatorSets and is used to iterate over the raw logs and unpacked data for AVSMigratedToOperatorSets events raised by the ContractIAVSDirectory contract.
type ContractIAVSDirectoryAVSMigratedToOperatorSetsIterator struct {
	Event *ContractIAVSDirectoryAVSMigratedToOperatorSets // Event containing the contract specifics and raw log

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
func (it *ContractIAVSDirectoryAVSMigratedToOperatorSetsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIAVSDirectoryAVSMigratedToOperatorSets)
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
		it.Event = new(ContractIAVSDirectoryAVSMigratedToOperatorSets)
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
func (it *ContractIAVSDirectoryAVSMigratedToOperatorSetsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIAVSDirectoryAVSMigratedToOperatorSetsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIAVSDirectoryAVSMigratedToOperatorSets represents a AVSMigratedToOperatorSets event raised by the ContractIAVSDirectory contract.
type ContractIAVSDirectoryAVSMigratedToOperatorSets struct {
	Avs common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAVSMigratedToOperatorSets is a free log retrieval operation binding the contract event 0x702b0c1f6cb1cf511aaa81f72bc05a215bb3497632d72c690c822b044ab494bf.
//
// Solidity: event AVSMigratedToOperatorSets(address indexed avs)
func (_ContractIAVSDirectory *ContractIAVSDirectoryFilterer) FilterAVSMigratedToOperatorSets(opts *bind.FilterOpts, avs []common.Address) (*ContractIAVSDirectoryAVSMigratedToOperatorSetsIterator, error) {

	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _ContractIAVSDirectory.contract.FilterLogs(opts, "AVSMigratedToOperatorSets", avsRule)
	if err != nil {
		return nil, err
	}
	return &ContractIAVSDirectoryAVSMigratedToOperatorSetsIterator{contract: _ContractIAVSDirectory.contract, event: "AVSMigratedToOperatorSets", logs: logs, sub: sub}, nil
}

// WatchAVSMigratedToOperatorSets is a free log subscription operation binding the contract event 0x702b0c1f6cb1cf511aaa81f72bc05a215bb3497632d72c690c822b044ab494bf.
//
// Solidity: event AVSMigratedToOperatorSets(address indexed avs)
func (_ContractIAVSDirectory *ContractIAVSDirectoryFilterer) WatchAVSMigratedToOperatorSets(opts *bind.WatchOpts, sink chan<- *ContractIAVSDirectoryAVSMigratedToOperatorSets, avs []common.Address) (event.Subscription, error) {

	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _ContractIAVSDirectory.contract.WatchLogs(opts, "AVSMigratedToOperatorSets", avsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIAVSDirectoryAVSMigratedToOperatorSets)
				if err := _ContractIAVSDirectory.contract.UnpackLog(event, "AVSMigratedToOperatorSets", log); err != nil {
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
func (_ContractIAVSDirectory *ContractIAVSDirectoryFilterer) ParseAVSMigratedToOperatorSets(log types.Log) (*ContractIAVSDirectoryAVSMigratedToOperatorSets, error) {
	event := new(ContractIAVSDirectoryAVSMigratedToOperatorSets)
	if err := _ContractIAVSDirectory.contract.UnpackLog(event, "AVSMigratedToOperatorSets", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIAVSDirectoryOperatorAVSRegistrationStatusUpdatedIterator is returned from FilterOperatorAVSRegistrationStatusUpdated and is used to iterate over the raw logs and unpacked data for OperatorAVSRegistrationStatusUpdated events raised by the ContractIAVSDirectory contract.
type ContractIAVSDirectoryOperatorAVSRegistrationStatusUpdatedIterator struct {
	Event *ContractIAVSDirectoryOperatorAVSRegistrationStatusUpdated // Event containing the contract specifics and raw log

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
func (it *ContractIAVSDirectoryOperatorAVSRegistrationStatusUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIAVSDirectoryOperatorAVSRegistrationStatusUpdated)
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
		it.Event = new(ContractIAVSDirectoryOperatorAVSRegistrationStatusUpdated)
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
func (it *ContractIAVSDirectoryOperatorAVSRegistrationStatusUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIAVSDirectoryOperatorAVSRegistrationStatusUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIAVSDirectoryOperatorAVSRegistrationStatusUpdated represents a OperatorAVSRegistrationStatusUpdated event raised by the ContractIAVSDirectory contract.
type ContractIAVSDirectoryOperatorAVSRegistrationStatusUpdated struct {
	Operator common.Address
	Avs      common.Address
	Status   uint8
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorAVSRegistrationStatusUpdated is a free log retrieval operation binding the contract event 0xf0952b1c65271d819d39983d2abb044b9cace59bcc4d4dd389f586ebdcb15b41.
//
// Solidity: event OperatorAVSRegistrationStatusUpdated(address indexed operator, address indexed avs, uint8 status)
func (_ContractIAVSDirectory *ContractIAVSDirectoryFilterer) FilterOperatorAVSRegistrationStatusUpdated(opts *bind.FilterOpts, operator []common.Address, avs []common.Address) (*ContractIAVSDirectoryOperatorAVSRegistrationStatusUpdatedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _ContractIAVSDirectory.contract.FilterLogs(opts, "OperatorAVSRegistrationStatusUpdated", operatorRule, avsRule)
	if err != nil {
		return nil, err
	}
	return &ContractIAVSDirectoryOperatorAVSRegistrationStatusUpdatedIterator{contract: _ContractIAVSDirectory.contract, event: "OperatorAVSRegistrationStatusUpdated", logs: logs, sub: sub}, nil
}

// WatchOperatorAVSRegistrationStatusUpdated is a free log subscription operation binding the contract event 0xf0952b1c65271d819d39983d2abb044b9cace59bcc4d4dd389f586ebdcb15b41.
//
// Solidity: event OperatorAVSRegistrationStatusUpdated(address indexed operator, address indexed avs, uint8 status)
func (_ContractIAVSDirectory *ContractIAVSDirectoryFilterer) WatchOperatorAVSRegistrationStatusUpdated(opts *bind.WatchOpts, sink chan<- *ContractIAVSDirectoryOperatorAVSRegistrationStatusUpdated, operator []common.Address, avs []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _ContractIAVSDirectory.contract.WatchLogs(opts, "OperatorAVSRegistrationStatusUpdated", operatorRule, avsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIAVSDirectoryOperatorAVSRegistrationStatusUpdated)
				if err := _ContractIAVSDirectory.contract.UnpackLog(event, "OperatorAVSRegistrationStatusUpdated", log); err != nil {
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
func (_ContractIAVSDirectory *ContractIAVSDirectoryFilterer) ParseOperatorAVSRegistrationStatusUpdated(log types.Log) (*ContractIAVSDirectoryOperatorAVSRegistrationStatusUpdated, error) {
	event := new(ContractIAVSDirectoryOperatorAVSRegistrationStatusUpdated)
	if err := _ContractIAVSDirectory.contract.UnpackLog(event, "OperatorAVSRegistrationStatusUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIAVSDirectoryOperatorAddedToOperatorSetIterator is returned from FilterOperatorAddedToOperatorSet and is used to iterate over the raw logs and unpacked data for OperatorAddedToOperatorSet events raised by the ContractIAVSDirectory contract.
type ContractIAVSDirectoryOperatorAddedToOperatorSetIterator struct {
	Event *ContractIAVSDirectoryOperatorAddedToOperatorSet // Event containing the contract specifics and raw log

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
func (it *ContractIAVSDirectoryOperatorAddedToOperatorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIAVSDirectoryOperatorAddedToOperatorSet)
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
		it.Event = new(ContractIAVSDirectoryOperatorAddedToOperatorSet)
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
func (it *ContractIAVSDirectoryOperatorAddedToOperatorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIAVSDirectoryOperatorAddedToOperatorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIAVSDirectoryOperatorAddedToOperatorSet represents a OperatorAddedToOperatorSet event raised by the ContractIAVSDirectory contract.
type ContractIAVSDirectoryOperatorAddedToOperatorSet struct {
	Operator    common.Address
	OperatorSet OperatorSet
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorAddedToOperatorSet is a free log retrieval operation binding the contract event 0x43232edf9071753d2321e5fa7e018363ee248e5f2142e6c08edd3265bfb4895e.
//
// Solidity: event OperatorAddedToOperatorSet(address indexed operator, (address,uint32) operatorSet)
func (_ContractIAVSDirectory *ContractIAVSDirectoryFilterer) FilterOperatorAddedToOperatorSet(opts *bind.FilterOpts, operator []common.Address) (*ContractIAVSDirectoryOperatorAddedToOperatorSetIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractIAVSDirectory.contract.FilterLogs(opts, "OperatorAddedToOperatorSet", operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractIAVSDirectoryOperatorAddedToOperatorSetIterator{contract: _ContractIAVSDirectory.contract, event: "OperatorAddedToOperatorSet", logs: logs, sub: sub}, nil
}

// WatchOperatorAddedToOperatorSet is a free log subscription operation binding the contract event 0x43232edf9071753d2321e5fa7e018363ee248e5f2142e6c08edd3265bfb4895e.
//
// Solidity: event OperatorAddedToOperatorSet(address indexed operator, (address,uint32) operatorSet)
func (_ContractIAVSDirectory *ContractIAVSDirectoryFilterer) WatchOperatorAddedToOperatorSet(opts *bind.WatchOpts, sink chan<- *ContractIAVSDirectoryOperatorAddedToOperatorSet, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractIAVSDirectory.contract.WatchLogs(opts, "OperatorAddedToOperatorSet", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIAVSDirectoryOperatorAddedToOperatorSet)
				if err := _ContractIAVSDirectory.contract.UnpackLog(event, "OperatorAddedToOperatorSet", log); err != nil {
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
func (_ContractIAVSDirectory *ContractIAVSDirectoryFilterer) ParseOperatorAddedToOperatorSet(log types.Log) (*ContractIAVSDirectoryOperatorAddedToOperatorSet, error) {
	event := new(ContractIAVSDirectoryOperatorAddedToOperatorSet)
	if err := _ContractIAVSDirectory.contract.UnpackLog(event, "OperatorAddedToOperatorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIAVSDirectoryOperatorMigratedToOperatorSetsIterator is returned from FilterOperatorMigratedToOperatorSets and is used to iterate over the raw logs and unpacked data for OperatorMigratedToOperatorSets events raised by the ContractIAVSDirectory contract.
type ContractIAVSDirectoryOperatorMigratedToOperatorSetsIterator struct {
	Event *ContractIAVSDirectoryOperatorMigratedToOperatorSets // Event containing the contract specifics and raw log

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
func (it *ContractIAVSDirectoryOperatorMigratedToOperatorSetsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIAVSDirectoryOperatorMigratedToOperatorSets)
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
		it.Event = new(ContractIAVSDirectoryOperatorMigratedToOperatorSets)
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
func (it *ContractIAVSDirectoryOperatorMigratedToOperatorSetsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIAVSDirectoryOperatorMigratedToOperatorSetsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIAVSDirectoryOperatorMigratedToOperatorSets represents a OperatorMigratedToOperatorSets event raised by the ContractIAVSDirectory contract.
type ContractIAVSDirectoryOperatorMigratedToOperatorSets struct {
	Operator       common.Address
	Avs            common.Address
	OperatorSetIds []uint32
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterOperatorMigratedToOperatorSets is a free log retrieval operation binding the contract event 0x54f33cfdd1ca703d795986b986fd47d742eab1904ecd2a5fdb8d6595e5904a01.
//
// Solidity: event OperatorMigratedToOperatorSets(address indexed operator, address indexed avs, uint32[] operatorSetIds)
func (_ContractIAVSDirectory *ContractIAVSDirectoryFilterer) FilterOperatorMigratedToOperatorSets(opts *bind.FilterOpts, operator []common.Address, avs []common.Address) (*ContractIAVSDirectoryOperatorMigratedToOperatorSetsIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _ContractIAVSDirectory.contract.FilterLogs(opts, "OperatorMigratedToOperatorSets", operatorRule, avsRule)
	if err != nil {
		return nil, err
	}
	return &ContractIAVSDirectoryOperatorMigratedToOperatorSetsIterator{contract: _ContractIAVSDirectory.contract, event: "OperatorMigratedToOperatorSets", logs: logs, sub: sub}, nil
}

// WatchOperatorMigratedToOperatorSets is a free log subscription operation binding the contract event 0x54f33cfdd1ca703d795986b986fd47d742eab1904ecd2a5fdb8d6595e5904a01.
//
// Solidity: event OperatorMigratedToOperatorSets(address indexed operator, address indexed avs, uint32[] operatorSetIds)
func (_ContractIAVSDirectory *ContractIAVSDirectoryFilterer) WatchOperatorMigratedToOperatorSets(opts *bind.WatchOpts, sink chan<- *ContractIAVSDirectoryOperatorMigratedToOperatorSets, operator []common.Address, avs []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _ContractIAVSDirectory.contract.WatchLogs(opts, "OperatorMigratedToOperatorSets", operatorRule, avsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIAVSDirectoryOperatorMigratedToOperatorSets)
				if err := _ContractIAVSDirectory.contract.UnpackLog(event, "OperatorMigratedToOperatorSets", log); err != nil {
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
func (_ContractIAVSDirectory *ContractIAVSDirectoryFilterer) ParseOperatorMigratedToOperatorSets(log types.Log) (*ContractIAVSDirectoryOperatorMigratedToOperatorSets, error) {
	event := new(ContractIAVSDirectoryOperatorMigratedToOperatorSets)
	if err := _ContractIAVSDirectory.contract.UnpackLog(event, "OperatorMigratedToOperatorSets", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIAVSDirectoryOperatorRemovedFromOperatorSetIterator is returned from FilterOperatorRemovedFromOperatorSet and is used to iterate over the raw logs and unpacked data for OperatorRemovedFromOperatorSet events raised by the ContractIAVSDirectory contract.
type ContractIAVSDirectoryOperatorRemovedFromOperatorSetIterator struct {
	Event *ContractIAVSDirectoryOperatorRemovedFromOperatorSet // Event containing the contract specifics and raw log

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
func (it *ContractIAVSDirectoryOperatorRemovedFromOperatorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIAVSDirectoryOperatorRemovedFromOperatorSet)
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
		it.Event = new(ContractIAVSDirectoryOperatorRemovedFromOperatorSet)
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
func (it *ContractIAVSDirectoryOperatorRemovedFromOperatorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIAVSDirectoryOperatorRemovedFromOperatorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIAVSDirectoryOperatorRemovedFromOperatorSet represents a OperatorRemovedFromOperatorSet event raised by the ContractIAVSDirectory contract.
type ContractIAVSDirectoryOperatorRemovedFromOperatorSet struct {
	Operator    common.Address
	OperatorSet OperatorSet
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorRemovedFromOperatorSet is a free log retrieval operation binding the contract event 0xad34c3070be1dffbcaa499d000ba2b8d9848aefcac3059df245dd95c4ece14fe.
//
// Solidity: event OperatorRemovedFromOperatorSet(address indexed operator, (address,uint32) operatorSet)
func (_ContractIAVSDirectory *ContractIAVSDirectoryFilterer) FilterOperatorRemovedFromOperatorSet(opts *bind.FilterOpts, operator []common.Address) (*ContractIAVSDirectoryOperatorRemovedFromOperatorSetIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractIAVSDirectory.contract.FilterLogs(opts, "OperatorRemovedFromOperatorSet", operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractIAVSDirectoryOperatorRemovedFromOperatorSetIterator{contract: _ContractIAVSDirectory.contract, event: "OperatorRemovedFromOperatorSet", logs: logs, sub: sub}, nil
}

// WatchOperatorRemovedFromOperatorSet is a free log subscription operation binding the contract event 0xad34c3070be1dffbcaa499d000ba2b8d9848aefcac3059df245dd95c4ece14fe.
//
// Solidity: event OperatorRemovedFromOperatorSet(address indexed operator, (address,uint32) operatorSet)
func (_ContractIAVSDirectory *ContractIAVSDirectoryFilterer) WatchOperatorRemovedFromOperatorSet(opts *bind.WatchOpts, sink chan<- *ContractIAVSDirectoryOperatorRemovedFromOperatorSet, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractIAVSDirectory.contract.WatchLogs(opts, "OperatorRemovedFromOperatorSet", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIAVSDirectoryOperatorRemovedFromOperatorSet)
				if err := _ContractIAVSDirectory.contract.UnpackLog(event, "OperatorRemovedFromOperatorSet", log); err != nil {
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
func (_ContractIAVSDirectory *ContractIAVSDirectoryFilterer) ParseOperatorRemovedFromOperatorSet(log types.Log) (*ContractIAVSDirectoryOperatorRemovedFromOperatorSet, error) {
	event := new(ContractIAVSDirectoryOperatorRemovedFromOperatorSet)
	if err := _ContractIAVSDirectory.contract.UnpackLog(event, "OperatorRemovedFromOperatorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIAVSDirectoryOperatorSetCreatedIterator is returned from FilterOperatorSetCreated and is used to iterate over the raw logs and unpacked data for OperatorSetCreated events raised by the ContractIAVSDirectory contract.
type ContractIAVSDirectoryOperatorSetCreatedIterator struct {
	Event *ContractIAVSDirectoryOperatorSetCreated // Event containing the contract specifics and raw log

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
func (it *ContractIAVSDirectoryOperatorSetCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIAVSDirectoryOperatorSetCreated)
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
		it.Event = new(ContractIAVSDirectoryOperatorSetCreated)
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
func (it *ContractIAVSDirectoryOperatorSetCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIAVSDirectoryOperatorSetCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIAVSDirectoryOperatorSetCreated represents a OperatorSetCreated event raised by the ContractIAVSDirectory contract.
type ContractIAVSDirectoryOperatorSetCreated struct {
	OperatorSet OperatorSet
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorSetCreated is a free log retrieval operation binding the contract event 0x31629285ead2335ae0933f86ed2ae63321f7af77b4e6eaabc42c057880977e6c.
//
// Solidity: event OperatorSetCreated((address,uint32) operatorSet)
func (_ContractIAVSDirectory *ContractIAVSDirectoryFilterer) FilterOperatorSetCreated(opts *bind.FilterOpts) (*ContractIAVSDirectoryOperatorSetCreatedIterator, error) {

	logs, sub, err := _ContractIAVSDirectory.contract.FilterLogs(opts, "OperatorSetCreated")
	if err != nil {
		return nil, err
	}
	return &ContractIAVSDirectoryOperatorSetCreatedIterator{contract: _ContractIAVSDirectory.contract, event: "OperatorSetCreated", logs: logs, sub: sub}, nil
}

// WatchOperatorSetCreated is a free log subscription operation binding the contract event 0x31629285ead2335ae0933f86ed2ae63321f7af77b4e6eaabc42c057880977e6c.
//
// Solidity: event OperatorSetCreated((address,uint32) operatorSet)
func (_ContractIAVSDirectory *ContractIAVSDirectoryFilterer) WatchOperatorSetCreated(opts *bind.WatchOpts, sink chan<- *ContractIAVSDirectoryOperatorSetCreated) (event.Subscription, error) {

	logs, sub, err := _ContractIAVSDirectory.contract.WatchLogs(opts, "OperatorSetCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIAVSDirectoryOperatorSetCreated)
				if err := _ContractIAVSDirectory.contract.UnpackLog(event, "OperatorSetCreated", log); err != nil {
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
func (_ContractIAVSDirectory *ContractIAVSDirectoryFilterer) ParseOperatorSetCreated(log types.Log) (*ContractIAVSDirectoryOperatorSetCreated, error) {
	event := new(ContractIAVSDirectoryOperatorSetCreated)
	if err := _ContractIAVSDirectory.contract.UnpackLog(event, "OperatorSetCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIAVSDirectoryStrategyAddedToOperatorSetIterator is returned from FilterStrategyAddedToOperatorSet and is used to iterate over the raw logs and unpacked data for StrategyAddedToOperatorSet events raised by the ContractIAVSDirectory contract.
type ContractIAVSDirectoryStrategyAddedToOperatorSetIterator struct {
	Event *ContractIAVSDirectoryStrategyAddedToOperatorSet // Event containing the contract specifics and raw log

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
func (it *ContractIAVSDirectoryStrategyAddedToOperatorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIAVSDirectoryStrategyAddedToOperatorSet)
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
		it.Event = new(ContractIAVSDirectoryStrategyAddedToOperatorSet)
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
func (it *ContractIAVSDirectoryStrategyAddedToOperatorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIAVSDirectoryStrategyAddedToOperatorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIAVSDirectoryStrategyAddedToOperatorSet represents a StrategyAddedToOperatorSet event raised by the ContractIAVSDirectory contract.
type ContractIAVSDirectoryStrategyAddedToOperatorSet struct {
	OperatorSet OperatorSet
	Strategy    common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterStrategyAddedToOperatorSet is a free log retrieval operation binding the contract event 0x7ab260fe0af193db5f4986770d831bda4ea46099dc817e8b6716dcae8af8e88b.
//
// Solidity: event StrategyAddedToOperatorSet((address,uint32) operatorSet, address strategy)
func (_ContractIAVSDirectory *ContractIAVSDirectoryFilterer) FilterStrategyAddedToOperatorSet(opts *bind.FilterOpts) (*ContractIAVSDirectoryStrategyAddedToOperatorSetIterator, error) {

	logs, sub, err := _ContractIAVSDirectory.contract.FilterLogs(opts, "StrategyAddedToOperatorSet")
	if err != nil {
		return nil, err
	}
	return &ContractIAVSDirectoryStrategyAddedToOperatorSetIterator{contract: _ContractIAVSDirectory.contract, event: "StrategyAddedToOperatorSet", logs: logs, sub: sub}, nil
}

// WatchStrategyAddedToOperatorSet is a free log subscription operation binding the contract event 0x7ab260fe0af193db5f4986770d831bda4ea46099dc817e8b6716dcae8af8e88b.
//
// Solidity: event StrategyAddedToOperatorSet((address,uint32) operatorSet, address strategy)
func (_ContractIAVSDirectory *ContractIAVSDirectoryFilterer) WatchStrategyAddedToOperatorSet(opts *bind.WatchOpts, sink chan<- *ContractIAVSDirectoryStrategyAddedToOperatorSet) (event.Subscription, error) {

	logs, sub, err := _ContractIAVSDirectory.contract.WatchLogs(opts, "StrategyAddedToOperatorSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIAVSDirectoryStrategyAddedToOperatorSet)
				if err := _ContractIAVSDirectory.contract.UnpackLog(event, "StrategyAddedToOperatorSet", log); err != nil {
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
func (_ContractIAVSDirectory *ContractIAVSDirectoryFilterer) ParseStrategyAddedToOperatorSet(log types.Log) (*ContractIAVSDirectoryStrategyAddedToOperatorSet, error) {
	event := new(ContractIAVSDirectoryStrategyAddedToOperatorSet)
	if err := _ContractIAVSDirectory.contract.UnpackLog(event, "StrategyAddedToOperatorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIAVSDirectoryStrategyRemovedFromOperatorSetIterator is returned from FilterStrategyRemovedFromOperatorSet and is used to iterate over the raw logs and unpacked data for StrategyRemovedFromOperatorSet events raised by the ContractIAVSDirectory contract.
type ContractIAVSDirectoryStrategyRemovedFromOperatorSetIterator struct {
	Event *ContractIAVSDirectoryStrategyRemovedFromOperatorSet // Event containing the contract specifics and raw log

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
func (it *ContractIAVSDirectoryStrategyRemovedFromOperatorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIAVSDirectoryStrategyRemovedFromOperatorSet)
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
		it.Event = new(ContractIAVSDirectoryStrategyRemovedFromOperatorSet)
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
func (it *ContractIAVSDirectoryStrategyRemovedFromOperatorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIAVSDirectoryStrategyRemovedFromOperatorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIAVSDirectoryStrategyRemovedFromOperatorSet represents a StrategyRemovedFromOperatorSet event raised by the ContractIAVSDirectory contract.
type ContractIAVSDirectoryStrategyRemovedFromOperatorSet struct {
	OperatorSet OperatorSet
	Strategy    common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterStrategyRemovedFromOperatorSet is a free log retrieval operation binding the contract event 0x7b4b073d80dcac55a11177d8459ad9f664ceeb91f71f27167bb14f8152a7eeee.
//
// Solidity: event StrategyRemovedFromOperatorSet((address,uint32) operatorSet, address strategy)
func (_ContractIAVSDirectory *ContractIAVSDirectoryFilterer) FilterStrategyRemovedFromOperatorSet(opts *bind.FilterOpts) (*ContractIAVSDirectoryStrategyRemovedFromOperatorSetIterator, error) {

	logs, sub, err := _ContractIAVSDirectory.contract.FilterLogs(opts, "StrategyRemovedFromOperatorSet")
	if err != nil {
		return nil, err
	}
	return &ContractIAVSDirectoryStrategyRemovedFromOperatorSetIterator{contract: _ContractIAVSDirectory.contract, event: "StrategyRemovedFromOperatorSet", logs: logs, sub: sub}, nil
}

// WatchStrategyRemovedFromOperatorSet is a free log subscription operation binding the contract event 0x7b4b073d80dcac55a11177d8459ad9f664ceeb91f71f27167bb14f8152a7eeee.
//
// Solidity: event StrategyRemovedFromOperatorSet((address,uint32) operatorSet, address strategy)
func (_ContractIAVSDirectory *ContractIAVSDirectoryFilterer) WatchStrategyRemovedFromOperatorSet(opts *bind.WatchOpts, sink chan<- *ContractIAVSDirectoryStrategyRemovedFromOperatorSet) (event.Subscription, error) {

	logs, sub, err := _ContractIAVSDirectory.contract.WatchLogs(opts, "StrategyRemovedFromOperatorSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIAVSDirectoryStrategyRemovedFromOperatorSet)
				if err := _ContractIAVSDirectory.contract.UnpackLog(event, "StrategyRemovedFromOperatorSet", log); err != nil {
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
func (_ContractIAVSDirectory *ContractIAVSDirectoryFilterer) ParseStrategyRemovedFromOperatorSet(log types.Log) (*ContractIAVSDirectoryStrategyRemovedFromOperatorSet, error) {
	event := new(ContractIAVSDirectoryStrategyRemovedFromOperatorSet)
	if err := _ContractIAVSDirectory.contract.UnpackLog(event, "StrategyRemovedFromOperatorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
