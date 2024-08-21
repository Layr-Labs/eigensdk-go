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

// IAVSDirectoryMagnitudeAllocation is an auto generated low-level Go binding around an user-defined struct.
type IAVSDirectoryMagnitudeAllocation struct {
	Strategy               common.Address
	ExpectedTotalMagnitude uint64
	OperatorSets           []IAVSDirectoryOperatorSet
	Magnitudes             []uint64
}

// IAVSDirectoryOperatorSet is an auto generated low-level Go binding around an user-defined struct.
type IAVSDirectoryOperatorSet struct {
	Avs           common.Address
	OperatorSetId uint32
}

// ISignatureUtilsSignatureWithSaltAndExpiry is an auto generated low-level Go binding around an user-defined struct.
type ISignatureUtilsSignatureWithSaltAndExpiry struct {
	Signature []byte
	Salt      [32]byte
	Expiry    *big.Int
}

// ContractIAVSDirectoryMetaData contains all meta data concerning the ContractIAVSDirectory contract.
var ContractIAVSDirectoryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"INITIAL_TOTAL_MAGNITUDE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"OPERATOR_AVS_REGISTRATION_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"OPERATOR_SET_REGISTRATION_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"becomeOperatorSetAVS\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"calculateMagnitudeAllocationDigestHash\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allocations\",\"type\":\"tuple[]\",\"internalType\":\"structIAVSDirectory.MagnitudeAllocation[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"expectedTotalMagnitude\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"operatorSets\",\"type\":\"tuple[]\",\"internalType\":\"structIAVSDirectory.OperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"magnitudes\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}]},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateOperatorAVSRegistrationDigestHash\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateOperatorSetForceDeregistrationTypehash\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateOperatorSetRegistrationDigestHash\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"cancelSalt\",\"inputs\":[{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createOperatorSets\",\"inputs\":[{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deregisterOperatorFromOperatorSets\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"domainSeparator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"forceDeregisterFromOperatorSets\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllocatableMagnitude\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllocationDelay\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLatestTotalMagnitude\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSlashablePPM\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structIAVSDirectory.OperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"timestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"linear\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint24[]\",\"internalType\":\"uint24[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initializeAllocationDelay\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delay\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isMember\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structIAVSDirectory.OperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOperatorSet\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOperatorSetAVS\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOperatorSlashable\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structIAVSDirectory.OperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"migrateOperatorsToOperatorSets\",\"inputs\":[{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"modifyAllocations\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allocations\",\"type\":\"tuple[]\",\"internalType\":\"structIAVSDirectory.MagnitudeAllocation[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"expectedTotalMagnitude\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"operatorSets\",\"type\":\"tuple[]\",\"internalType\":\"structIAVSDirectory.OperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"magnitudes\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}]},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"operatorSaltIsSpent\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorSetMemberCount\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerOperatorToOperatorSets\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"slashOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"bipsToSlash\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateAVSMetadataURI\",\"inputs\":[{\"name\":\"metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateFreeMagnitude\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"numToComplete\",\"type\":\"uint8[]\",\"internalType\":\"uint8[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AVSMetadataURIUpdated\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"metadataURI\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AVSMigratedToOperatorSets\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MagnitudeAllocated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIAVSDirectory.OperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"magnitudeToAllocate\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"effectTimestamp\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MagnitudeDeallocationCompleted\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIAVSDirectory.OperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"freeMagnitudeAdded\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MagnitudeQueueDeallocated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIAVSDirectory.OperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"magnitudeToDeallocate\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"completableTimestamp\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorAVSRegistrationStatusUpdated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"status\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumIAVSDirectory.OperatorAVSRegistrationStatus\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorAddedToOperatorSet\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIAVSDirectory.OperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorMigratedToOperatorSets\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"indexed\":false,\"internalType\":\"uint32[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorRemovedFromOperatorSet\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIAVSDirectory.OperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSetCreated\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIAVSDirectory.OperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSlashed\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"bipsToSlash\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false}]",
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
	INITIALTOTALMAGNITUDE(opts *bind.CallOpts) (uint64, error)

	OPERATORAVSREGISTRATIONTYPEHASH(opts *bind.CallOpts) ([32]byte, error)

	OPERATORSETREGISTRATIONTYPEHASH(opts *bind.CallOpts) ([32]byte, error)

	CalculateMagnitudeAllocationDigestHash(opts *bind.CallOpts, operator common.Address, allocations []IAVSDirectoryMagnitudeAllocation, salt [32]byte, expiry *big.Int) ([32]byte, error)

	CalculateOperatorAVSRegistrationDigestHash(opts *bind.CallOpts, operator common.Address, avs common.Address, salt [32]byte, expiry *big.Int) ([32]byte, error)

	CalculateOperatorSetForceDeregistrationTypehash(opts *bind.CallOpts, avs common.Address, operatorSetIds []uint32, salt [32]byte, expiry *big.Int) ([32]byte, error)

	CalculateOperatorSetRegistrationDigestHash(opts *bind.CallOpts, avs common.Address, operatorSetIds []uint32, salt [32]byte, expiry *big.Int) ([32]byte, error)

	DomainSeparator(opts *bind.CallOpts) ([32]byte, error)

	GetAllocatableMagnitude(opts *bind.CallOpts, operator common.Address, strategy common.Address) (uint64, error)

	GetAllocationDelay(opts *bind.CallOpts, operator common.Address) (uint32, error)

	GetLatestTotalMagnitude(opts *bind.CallOpts, operator common.Address, strategy common.Address) (uint64, error)

	GetSlashablePPM(opts *bind.CallOpts, operator common.Address, operatorSet IAVSDirectoryOperatorSet, strategies []common.Address, timestamp uint32, linear bool) ([]*big.Int, error)

	IsMember(opts *bind.CallOpts, operator common.Address, operatorSet IAVSDirectoryOperatorSet) (bool, error)

	IsOperatorSet(opts *bind.CallOpts, avs common.Address, operatorSetId uint32) (bool, error)

	IsOperatorSetAVS(opts *bind.CallOpts, avs common.Address) (bool, error)

	IsOperatorSlashable(opts *bind.CallOpts, operator common.Address, operatorSet IAVSDirectoryOperatorSet) (bool, error)

	OperatorSaltIsSpent(opts *bind.CallOpts, operator common.Address, salt [32]byte) (bool, error)

	OperatorSetMemberCount(opts *bind.CallOpts, avs common.Address, operatorSetId uint32) (*big.Int, error)
}

// ContractIAVSDirectoryTransacts is an auto generated interface that defines the transact methods available for an Ethereum contract.
type ContractIAVSDirectoryTransacts interface {
	BecomeOperatorSetAVS(opts *bind.TransactOpts) (*types.Transaction, error)

	CancelSalt(opts *bind.TransactOpts, salt [32]byte) (*types.Transaction, error)

	CreateOperatorSets(opts *bind.TransactOpts, operatorSetIds []uint32) (*types.Transaction, error)

	DeregisterOperatorFromOperatorSets(opts *bind.TransactOpts, operator common.Address, operatorSetIds []uint32) (*types.Transaction, error)

	ForceDeregisterFromOperatorSets(opts *bind.TransactOpts, operator common.Address, avs common.Address, operatorSetIds []uint32, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error)

	InitializeAllocationDelay(opts *bind.TransactOpts, operator common.Address, delay uint32) (*types.Transaction, error)

	MigrateOperatorsToOperatorSets(opts *bind.TransactOpts, operators []common.Address, operatorSetIds [][]uint32) (*types.Transaction, error)

	ModifyAllocations(opts *bind.TransactOpts, operator common.Address, allocations []IAVSDirectoryMagnitudeAllocation, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error)

	RegisterOperatorToOperatorSets(opts *bind.TransactOpts, operator common.Address, operatorSetIds []uint32, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error)

	SlashOperator(opts *bind.TransactOpts, operator common.Address, operatorSetId uint32, strategies []common.Address, bipsToSlash uint16) (*types.Transaction, error)

	UpdateAVSMetadataURI(opts *bind.TransactOpts, metadataURI string) (*types.Transaction, error)

	UpdateFreeMagnitude(opts *bind.TransactOpts, operator common.Address, strategies []common.Address, numToComplete []uint8) (*types.Transaction, error)
}

// ContractIAVSDirectoryFilterer is an auto generated interface that defines the log filtering methods available for an Ethereum contract.
type ContractIAVSDirectoryFilters interface {
	FilterAVSMetadataURIUpdated(opts *bind.FilterOpts, avs []common.Address) (*ContractIAVSDirectoryAVSMetadataURIUpdatedIterator, error)
	WatchAVSMetadataURIUpdated(opts *bind.WatchOpts, sink chan<- *ContractIAVSDirectoryAVSMetadataURIUpdated, avs []common.Address) (event.Subscription, error)
	ParseAVSMetadataURIUpdated(log types.Log) (*ContractIAVSDirectoryAVSMetadataURIUpdated, error)

	FilterAVSMigratedToOperatorSets(opts *bind.FilterOpts, avs []common.Address) (*ContractIAVSDirectoryAVSMigratedToOperatorSetsIterator, error)
	WatchAVSMigratedToOperatorSets(opts *bind.WatchOpts, sink chan<- *ContractIAVSDirectoryAVSMigratedToOperatorSets, avs []common.Address) (event.Subscription, error)
	ParseAVSMigratedToOperatorSets(log types.Log) (*ContractIAVSDirectoryAVSMigratedToOperatorSets, error)

	FilterMagnitudeAllocated(opts *bind.FilterOpts) (*ContractIAVSDirectoryMagnitudeAllocatedIterator, error)
	WatchMagnitudeAllocated(opts *bind.WatchOpts, sink chan<- *ContractIAVSDirectoryMagnitudeAllocated) (event.Subscription, error)
	ParseMagnitudeAllocated(log types.Log) (*ContractIAVSDirectoryMagnitudeAllocated, error)

	FilterMagnitudeDeallocationCompleted(opts *bind.FilterOpts) (*ContractIAVSDirectoryMagnitudeDeallocationCompletedIterator, error)
	WatchMagnitudeDeallocationCompleted(opts *bind.WatchOpts, sink chan<- *ContractIAVSDirectoryMagnitudeDeallocationCompleted) (event.Subscription, error)
	ParseMagnitudeDeallocationCompleted(log types.Log) (*ContractIAVSDirectoryMagnitudeDeallocationCompleted, error)

	FilterMagnitudeQueueDeallocated(opts *bind.FilterOpts) (*ContractIAVSDirectoryMagnitudeQueueDeallocatedIterator, error)
	WatchMagnitudeQueueDeallocated(opts *bind.WatchOpts, sink chan<- *ContractIAVSDirectoryMagnitudeQueueDeallocated) (event.Subscription, error)
	ParseMagnitudeQueueDeallocated(log types.Log) (*ContractIAVSDirectoryMagnitudeQueueDeallocated, error)

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

	FilterOperatorSlashed(opts *bind.FilterOpts) (*ContractIAVSDirectoryOperatorSlashedIterator, error)
	WatchOperatorSlashed(opts *bind.WatchOpts, sink chan<- *ContractIAVSDirectoryOperatorSlashed) (event.Subscription, error)
	ParseOperatorSlashed(log types.Log) (*ContractIAVSDirectoryOperatorSlashed, error)
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

// INITIALTOTALMAGNITUDE is a free data retrieval call binding the contract method 0x9a543ca4.
//
// Solidity: function INITIAL_TOTAL_MAGNITUDE() view returns(uint64)
func (_ContractIAVSDirectory *ContractIAVSDirectoryCaller) INITIALTOTALMAGNITUDE(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _ContractIAVSDirectory.contract.Call(opts, &out, "INITIAL_TOTAL_MAGNITUDE")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// INITIALTOTALMAGNITUDE is a free data retrieval call binding the contract method 0x9a543ca4.
//
// Solidity: function INITIAL_TOTAL_MAGNITUDE() view returns(uint64)
func (_ContractIAVSDirectory *ContractIAVSDirectorySession) INITIALTOTALMAGNITUDE() (uint64, error) {
	return _ContractIAVSDirectory.Contract.INITIALTOTALMAGNITUDE(&_ContractIAVSDirectory.CallOpts)
}

// INITIALTOTALMAGNITUDE is a free data retrieval call binding the contract method 0x9a543ca4.
//
// Solidity: function INITIAL_TOTAL_MAGNITUDE() view returns(uint64)
func (_ContractIAVSDirectory *ContractIAVSDirectoryCallerSession) INITIALTOTALMAGNITUDE() (uint64, error) {
	return _ContractIAVSDirectory.Contract.INITIALTOTALMAGNITUDE(&_ContractIAVSDirectory.CallOpts)
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

// CalculateMagnitudeAllocationDigestHash is a free data retrieval call binding the contract method 0x686b686e.
//
// Solidity: function calculateMagnitudeAllocationDigestHash(address operator, (address,uint64,(address,uint32)[],uint64[])[] allocations, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_ContractIAVSDirectory *ContractIAVSDirectoryCaller) CalculateMagnitudeAllocationDigestHash(opts *bind.CallOpts, operator common.Address, allocations []IAVSDirectoryMagnitudeAllocation, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _ContractIAVSDirectory.contract.Call(opts, &out, "calculateMagnitudeAllocationDigestHash", operator, allocations, salt, expiry)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateMagnitudeAllocationDigestHash is a free data retrieval call binding the contract method 0x686b686e.
//
// Solidity: function calculateMagnitudeAllocationDigestHash(address operator, (address,uint64,(address,uint32)[],uint64[])[] allocations, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_ContractIAVSDirectory *ContractIAVSDirectorySession) CalculateMagnitudeAllocationDigestHash(operator common.Address, allocations []IAVSDirectoryMagnitudeAllocation, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _ContractIAVSDirectory.Contract.CalculateMagnitudeAllocationDigestHash(&_ContractIAVSDirectory.CallOpts, operator, allocations, salt, expiry)
}

// CalculateMagnitudeAllocationDigestHash is a free data retrieval call binding the contract method 0x686b686e.
//
// Solidity: function calculateMagnitudeAllocationDigestHash(address operator, (address,uint64,(address,uint32)[],uint64[])[] allocations, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_ContractIAVSDirectory *ContractIAVSDirectoryCallerSession) CalculateMagnitudeAllocationDigestHash(operator common.Address, allocations []IAVSDirectoryMagnitudeAllocation, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _ContractIAVSDirectory.Contract.CalculateMagnitudeAllocationDigestHash(&_ContractIAVSDirectory.CallOpts, operator, allocations, salt, expiry)
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

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_ContractIAVSDirectory *ContractIAVSDirectoryCaller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ContractIAVSDirectory.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_ContractIAVSDirectory *ContractIAVSDirectorySession) DomainSeparator() ([32]byte, error) {
	return _ContractIAVSDirectory.Contract.DomainSeparator(&_ContractIAVSDirectory.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_ContractIAVSDirectory *ContractIAVSDirectoryCallerSession) DomainSeparator() ([32]byte, error) {
	return _ContractIAVSDirectory.Contract.DomainSeparator(&_ContractIAVSDirectory.CallOpts)
}

// GetAllocatableMagnitude is a free data retrieval call binding the contract method 0x6cfb4481.
//
// Solidity: function getAllocatableMagnitude(address operator, address strategy) view returns(uint64)
func (_ContractIAVSDirectory *ContractIAVSDirectoryCaller) GetAllocatableMagnitude(opts *bind.CallOpts, operator common.Address, strategy common.Address) (uint64, error) {
	var out []interface{}
	err := _ContractIAVSDirectory.contract.Call(opts, &out, "getAllocatableMagnitude", operator, strategy)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetAllocatableMagnitude is a free data retrieval call binding the contract method 0x6cfb4481.
//
// Solidity: function getAllocatableMagnitude(address operator, address strategy) view returns(uint64)
func (_ContractIAVSDirectory *ContractIAVSDirectorySession) GetAllocatableMagnitude(operator common.Address, strategy common.Address) (uint64, error) {
	return _ContractIAVSDirectory.Contract.GetAllocatableMagnitude(&_ContractIAVSDirectory.CallOpts, operator, strategy)
}

// GetAllocatableMagnitude is a free data retrieval call binding the contract method 0x6cfb4481.
//
// Solidity: function getAllocatableMagnitude(address operator, address strategy) view returns(uint64)
func (_ContractIAVSDirectory *ContractIAVSDirectoryCallerSession) GetAllocatableMagnitude(operator common.Address, strategy common.Address) (uint64, error) {
	return _ContractIAVSDirectory.Contract.GetAllocatableMagnitude(&_ContractIAVSDirectory.CallOpts, operator, strategy)
}

// GetAllocationDelay is a free data retrieval call binding the contract method 0xb9fbaed1.
//
// Solidity: function getAllocationDelay(address operator) view returns(uint32)
func (_ContractIAVSDirectory *ContractIAVSDirectoryCaller) GetAllocationDelay(opts *bind.CallOpts, operator common.Address) (uint32, error) {
	var out []interface{}
	err := _ContractIAVSDirectory.contract.Call(opts, &out, "getAllocationDelay", operator)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetAllocationDelay is a free data retrieval call binding the contract method 0xb9fbaed1.
//
// Solidity: function getAllocationDelay(address operator) view returns(uint32)
func (_ContractIAVSDirectory *ContractIAVSDirectorySession) GetAllocationDelay(operator common.Address) (uint32, error) {
	return _ContractIAVSDirectory.Contract.GetAllocationDelay(&_ContractIAVSDirectory.CallOpts, operator)
}

// GetAllocationDelay is a free data retrieval call binding the contract method 0xb9fbaed1.
//
// Solidity: function getAllocationDelay(address operator) view returns(uint32)
func (_ContractIAVSDirectory *ContractIAVSDirectoryCallerSession) GetAllocationDelay(operator common.Address) (uint32, error) {
	return _ContractIAVSDirectory.Contract.GetAllocationDelay(&_ContractIAVSDirectory.CallOpts, operator)
}

// GetLatestTotalMagnitude is a free data retrieval call binding the contract method 0x7ed0687d.
//
// Solidity: function getLatestTotalMagnitude(address operator, address strategy) view returns(uint64)
func (_ContractIAVSDirectory *ContractIAVSDirectoryCaller) GetLatestTotalMagnitude(opts *bind.CallOpts, operator common.Address, strategy common.Address) (uint64, error) {
	var out []interface{}
	err := _ContractIAVSDirectory.contract.Call(opts, &out, "getLatestTotalMagnitude", operator, strategy)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetLatestTotalMagnitude is a free data retrieval call binding the contract method 0x7ed0687d.
//
// Solidity: function getLatestTotalMagnitude(address operator, address strategy) view returns(uint64)
func (_ContractIAVSDirectory *ContractIAVSDirectorySession) GetLatestTotalMagnitude(operator common.Address, strategy common.Address) (uint64, error) {
	return _ContractIAVSDirectory.Contract.GetLatestTotalMagnitude(&_ContractIAVSDirectory.CallOpts, operator, strategy)
}

// GetLatestTotalMagnitude is a free data retrieval call binding the contract method 0x7ed0687d.
//
// Solidity: function getLatestTotalMagnitude(address operator, address strategy) view returns(uint64)
func (_ContractIAVSDirectory *ContractIAVSDirectoryCallerSession) GetLatestTotalMagnitude(operator common.Address, strategy common.Address) (uint64, error) {
	return _ContractIAVSDirectory.Contract.GetLatestTotalMagnitude(&_ContractIAVSDirectory.CallOpts, operator, strategy)
}

// GetSlashablePPM is a free data retrieval call binding the contract method 0x1c699d65.
//
// Solidity: function getSlashablePPM(address operator, (address,uint32) operatorSet, address[] strategies, uint32 timestamp, bool linear) view returns(uint24[])
func (_ContractIAVSDirectory *ContractIAVSDirectoryCaller) GetSlashablePPM(opts *bind.CallOpts, operator common.Address, operatorSet IAVSDirectoryOperatorSet, strategies []common.Address, timestamp uint32, linear bool) ([]*big.Int, error) {
	var out []interface{}
	err := _ContractIAVSDirectory.contract.Call(opts, &out, "getSlashablePPM", operator, operatorSet, strategies, timestamp, linear)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetSlashablePPM is a free data retrieval call binding the contract method 0x1c699d65.
//
// Solidity: function getSlashablePPM(address operator, (address,uint32) operatorSet, address[] strategies, uint32 timestamp, bool linear) view returns(uint24[])
func (_ContractIAVSDirectory *ContractIAVSDirectorySession) GetSlashablePPM(operator common.Address, operatorSet IAVSDirectoryOperatorSet, strategies []common.Address, timestamp uint32, linear bool) ([]*big.Int, error) {
	return _ContractIAVSDirectory.Contract.GetSlashablePPM(&_ContractIAVSDirectory.CallOpts, operator, operatorSet, strategies, timestamp, linear)
}

// GetSlashablePPM is a free data retrieval call binding the contract method 0x1c699d65.
//
// Solidity: function getSlashablePPM(address operator, (address,uint32) operatorSet, address[] strategies, uint32 timestamp, bool linear) view returns(uint24[])
func (_ContractIAVSDirectory *ContractIAVSDirectoryCallerSession) GetSlashablePPM(operator common.Address, operatorSet IAVSDirectoryOperatorSet, strategies []common.Address, timestamp uint32, linear bool) ([]*big.Int, error) {
	return _ContractIAVSDirectory.Contract.GetSlashablePPM(&_ContractIAVSDirectory.CallOpts, operator, operatorSet, strategies, timestamp, linear)
}

// IsMember is a free data retrieval call binding the contract method 0xda2ff05d.
//
// Solidity: function isMember(address operator, (address,uint32) operatorSet) view returns(bool)
func (_ContractIAVSDirectory *ContractIAVSDirectoryCaller) IsMember(opts *bind.CallOpts, operator common.Address, operatorSet IAVSDirectoryOperatorSet) (bool, error) {
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
func (_ContractIAVSDirectory *ContractIAVSDirectorySession) IsMember(operator common.Address, operatorSet IAVSDirectoryOperatorSet) (bool, error) {
	return _ContractIAVSDirectory.Contract.IsMember(&_ContractIAVSDirectory.CallOpts, operator, operatorSet)
}

// IsMember is a free data retrieval call binding the contract method 0xda2ff05d.
//
// Solidity: function isMember(address operator, (address,uint32) operatorSet) view returns(bool)
func (_ContractIAVSDirectory *ContractIAVSDirectoryCallerSession) IsMember(operator common.Address, operatorSet IAVSDirectoryOperatorSet) (bool, error) {
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

// IsOperatorSlashable is a free data retrieval call binding the contract method 0x1352c3e6.
//
// Solidity: function isOperatorSlashable(address operator, (address,uint32) operatorSet) view returns(bool)
func (_ContractIAVSDirectory *ContractIAVSDirectoryCaller) IsOperatorSlashable(opts *bind.CallOpts, operator common.Address, operatorSet IAVSDirectoryOperatorSet) (bool, error) {
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
func (_ContractIAVSDirectory *ContractIAVSDirectorySession) IsOperatorSlashable(operator common.Address, operatorSet IAVSDirectoryOperatorSet) (bool, error) {
	return _ContractIAVSDirectory.Contract.IsOperatorSlashable(&_ContractIAVSDirectory.CallOpts, operator, operatorSet)
}

// IsOperatorSlashable is a free data retrieval call binding the contract method 0x1352c3e6.
//
// Solidity: function isOperatorSlashable(address operator, (address,uint32) operatorSet) view returns(bool)
func (_ContractIAVSDirectory *ContractIAVSDirectoryCallerSession) IsOperatorSlashable(operator common.Address, operatorSet IAVSDirectoryOperatorSet) (bool, error) {
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

// OperatorSetMemberCount is a free data retrieval call binding the contract method 0xdae226b6.
//
// Solidity: function operatorSetMemberCount(address avs, uint32 operatorSetId) view returns(uint256)
func (_ContractIAVSDirectory *ContractIAVSDirectoryCaller) OperatorSetMemberCount(opts *bind.CallOpts, avs common.Address, operatorSetId uint32) (*big.Int, error) {
	var out []interface{}
	err := _ContractIAVSDirectory.contract.Call(opts, &out, "operatorSetMemberCount", avs, operatorSetId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OperatorSetMemberCount is a free data retrieval call binding the contract method 0xdae226b6.
//
// Solidity: function operatorSetMemberCount(address avs, uint32 operatorSetId) view returns(uint256)
func (_ContractIAVSDirectory *ContractIAVSDirectorySession) OperatorSetMemberCount(avs common.Address, operatorSetId uint32) (*big.Int, error) {
	return _ContractIAVSDirectory.Contract.OperatorSetMemberCount(&_ContractIAVSDirectory.CallOpts, avs, operatorSetId)
}

// OperatorSetMemberCount is a free data retrieval call binding the contract method 0xdae226b6.
//
// Solidity: function operatorSetMemberCount(address avs, uint32 operatorSetId) view returns(uint256)
func (_ContractIAVSDirectory *ContractIAVSDirectoryCallerSession) OperatorSetMemberCount(avs common.Address, operatorSetId uint32) (*big.Int, error) {
	return _ContractIAVSDirectory.Contract.OperatorSetMemberCount(&_ContractIAVSDirectory.CallOpts, avs, operatorSetId)
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

// InitializeAllocationDelay is a paid mutator transaction binding the contract method 0x80fe6da4.
//
// Solidity: function initializeAllocationDelay(address operator, uint32 delay) returns()
func (_ContractIAVSDirectory *ContractIAVSDirectoryTransactor) InitializeAllocationDelay(opts *bind.TransactOpts, operator common.Address, delay uint32) (*types.Transaction, error) {
	return _ContractIAVSDirectory.contract.Transact(opts, "initializeAllocationDelay", operator, delay)
}

// InitializeAllocationDelay is a paid mutator transaction binding the contract method 0x80fe6da4.
//
// Solidity: function initializeAllocationDelay(address operator, uint32 delay) returns()
func (_ContractIAVSDirectory *ContractIAVSDirectorySession) InitializeAllocationDelay(operator common.Address, delay uint32) (*types.Transaction, error) {
	return _ContractIAVSDirectory.Contract.InitializeAllocationDelay(&_ContractIAVSDirectory.TransactOpts, operator, delay)
}

// InitializeAllocationDelay is a paid mutator transaction binding the contract method 0x80fe6da4.
//
// Solidity: function initializeAllocationDelay(address operator, uint32 delay) returns()
func (_ContractIAVSDirectory *ContractIAVSDirectoryTransactorSession) InitializeAllocationDelay(operator common.Address, delay uint32) (*types.Transaction, error) {
	return _ContractIAVSDirectory.Contract.InitializeAllocationDelay(&_ContractIAVSDirectory.TransactOpts, operator, delay)
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

// ModifyAllocations is a paid mutator transaction binding the contract method 0xf8e91d16.
//
// Solidity: function modifyAllocations(address operator, (address,uint64,(address,uint32)[],uint64[])[] allocations, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractIAVSDirectory *ContractIAVSDirectoryTransactor) ModifyAllocations(opts *bind.TransactOpts, operator common.Address, allocations []IAVSDirectoryMagnitudeAllocation, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractIAVSDirectory.contract.Transact(opts, "modifyAllocations", operator, allocations, operatorSignature)
}

// ModifyAllocations is a paid mutator transaction binding the contract method 0xf8e91d16.
//
// Solidity: function modifyAllocations(address operator, (address,uint64,(address,uint32)[],uint64[])[] allocations, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractIAVSDirectory *ContractIAVSDirectorySession) ModifyAllocations(operator common.Address, allocations []IAVSDirectoryMagnitudeAllocation, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractIAVSDirectory.Contract.ModifyAllocations(&_ContractIAVSDirectory.TransactOpts, operator, allocations, operatorSignature)
}

// ModifyAllocations is a paid mutator transaction binding the contract method 0xf8e91d16.
//
// Solidity: function modifyAllocations(address operator, (address,uint64,(address,uint32)[],uint64[])[] allocations, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractIAVSDirectory *ContractIAVSDirectoryTransactorSession) ModifyAllocations(operator common.Address, allocations []IAVSDirectoryMagnitudeAllocation, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractIAVSDirectory.Contract.ModifyAllocations(&_ContractIAVSDirectory.TransactOpts, operator, allocations, operatorSignature)
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

// SlashOperator is a paid mutator transaction binding the contract method 0xbd74a06c.
//
// Solidity: function slashOperator(address operator, uint32 operatorSetId, address[] strategies, uint16 bipsToSlash) returns()
func (_ContractIAVSDirectory *ContractIAVSDirectoryTransactor) SlashOperator(opts *bind.TransactOpts, operator common.Address, operatorSetId uint32, strategies []common.Address, bipsToSlash uint16) (*types.Transaction, error) {
	return _ContractIAVSDirectory.contract.Transact(opts, "slashOperator", operator, operatorSetId, strategies, bipsToSlash)
}

// SlashOperator is a paid mutator transaction binding the contract method 0xbd74a06c.
//
// Solidity: function slashOperator(address operator, uint32 operatorSetId, address[] strategies, uint16 bipsToSlash) returns()
func (_ContractIAVSDirectory *ContractIAVSDirectorySession) SlashOperator(operator common.Address, operatorSetId uint32, strategies []common.Address, bipsToSlash uint16) (*types.Transaction, error) {
	return _ContractIAVSDirectory.Contract.SlashOperator(&_ContractIAVSDirectory.TransactOpts, operator, operatorSetId, strategies, bipsToSlash)
}

// SlashOperator is a paid mutator transaction binding the contract method 0xbd74a06c.
//
// Solidity: function slashOperator(address operator, uint32 operatorSetId, address[] strategies, uint16 bipsToSlash) returns()
func (_ContractIAVSDirectory *ContractIAVSDirectoryTransactorSession) SlashOperator(operator common.Address, operatorSetId uint32, strategies []common.Address, bipsToSlash uint16) (*types.Transaction, error) {
	return _ContractIAVSDirectory.Contract.SlashOperator(&_ContractIAVSDirectory.TransactOpts, operator, operatorSetId, strategies, bipsToSlash)
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

// UpdateFreeMagnitude is a paid mutator transaction binding the contract method 0x84b78fc2.
//
// Solidity: function updateFreeMagnitude(address operator, address[] strategies, uint8[] numToComplete) returns()
func (_ContractIAVSDirectory *ContractIAVSDirectoryTransactor) UpdateFreeMagnitude(opts *bind.TransactOpts, operator common.Address, strategies []common.Address, numToComplete []uint8) (*types.Transaction, error) {
	return _ContractIAVSDirectory.contract.Transact(opts, "updateFreeMagnitude", operator, strategies, numToComplete)
}

// UpdateFreeMagnitude is a paid mutator transaction binding the contract method 0x84b78fc2.
//
// Solidity: function updateFreeMagnitude(address operator, address[] strategies, uint8[] numToComplete) returns()
func (_ContractIAVSDirectory *ContractIAVSDirectorySession) UpdateFreeMagnitude(operator common.Address, strategies []common.Address, numToComplete []uint8) (*types.Transaction, error) {
	return _ContractIAVSDirectory.Contract.UpdateFreeMagnitude(&_ContractIAVSDirectory.TransactOpts, operator, strategies, numToComplete)
}

// UpdateFreeMagnitude is a paid mutator transaction binding the contract method 0x84b78fc2.
//
// Solidity: function updateFreeMagnitude(address operator, address[] strategies, uint8[] numToComplete) returns()
func (_ContractIAVSDirectory *ContractIAVSDirectoryTransactorSession) UpdateFreeMagnitude(operator common.Address, strategies []common.Address, numToComplete []uint8) (*types.Transaction, error) {
	return _ContractIAVSDirectory.Contract.UpdateFreeMagnitude(&_ContractIAVSDirectory.TransactOpts, operator, strategies, numToComplete)
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

// ContractIAVSDirectoryMagnitudeAllocatedIterator is returned from FilterMagnitudeAllocated and is used to iterate over the raw logs and unpacked data for MagnitudeAllocated events raised by the ContractIAVSDirectory contract.
type ContractIAVSDirectoryMagnitudeAllocatedIterator struct {
	Event *ContractIAVSDirectoryMagnitudeAllocated // Event containing the contract specifics and raw log

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
func (it *ContractIAVSDirectoryMagnitudeAllocatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIAVSDirectoryMagnitudeAllocated)
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
		it.Event = new(ContractIAVSDirectoryMagnitudeAllocated)
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
func (it *ContractIAVSDirectoryMagnitudeAllocatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIAVSDirectoryMagnitudeAllocatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIAVSDirectoryMagnitudeAllocated represents a MagnitudeAllocated event raised by the ContractIAVSDirectory contract.
type ContractIAVSDirectoryMagnitudeAllocated struct {
	Operator            common.Address
	Strategy            common.Address
	OperatorSet         IAVSDirectoryOperatorSet
	MagnitudeToAllocate uint64
	EffectTimestamp     uint32
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterMagnitudeAllocated is a free log retrieval operation binding the contract event 0x6d7d0079582cb2c5e70d4135b37f36711415ee6c260778b716bd65e026eb4f1a.
//
// Solidity: event MagnitudeAllocated(address operator, address strategy, (address,uint32) operatorSet, uint64 magnitudeToAllocate, uint32 effectTimestamp)
func (_ContractIAVSDirectory *ContractIAVSDirectoryFilterer) FilterMagnitudeAllocated(opts *bind.FilterOpts) (*ContractIAVSDirectoryMagnitudeAllocatedIterator, error) {

	logs, sub, err := _ContractIAVSDirectory.contract.FilterLogs(opts, "MagnitudeAllocated")
	if err != nil {
		return nil, err
	}
	return &ContractIAVSDirectoryMagnitudeAllocatedIterator{contract: _ContractIAVSDirectory.contract, event: "MagnitudeAllocated", logs: logs, sub: sub}, nil
}

// WatchMagnitudeAllocated is a free log subscription operation binding the contract event 0x6d7d0079582cb2c5e70d4135b37f36711415ee6c260778b716bd65e026eb4f1a.
//
// Solidity: event MagnitudeAllocated(address operator, address strategy, (address,uint32) operatorSet, uint64 magnitudeToAllocate, uint32 effectTimestamp)
func (_ContractIAVSDirectory *ContractIAVSDirectoryFilterer) WatchMagnitudeAllocated(opts *bind.WatchOpts, sink chan<- *ContractIAVSDirectoryMagnitudeAllocated) (event.Subscription, error) {

	logs, sub, err := _ContractIAVSDirectory.contract.WatchLogs(opts, "MagnitudeAllocated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIAVSDirectoryMagnitudeAllocated)
				if err := _ContractIAVSDirectory.contract.UnpackLog(event, "MagnitudeAllocated", log); err != nil {
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

// ParseMagnitudeAllocated is a log parse operation binding the contract event 0x6d7d0079582cb2c5e70d4135b37f36711415ee6c260778b716bd65e026eb4f1a.
//
// Solidity: event MagnitudeAllocated(address operator, address strategy, (address,uint32) operatorSet, uint64 magnitudeToAllocate, uint32 effectTimestamp)
func (_ContractIAVSDirectory *ContractIAVSDirectoryFilterer) ParseMagnitudeAllocated(log types.Log) (*ContractIAVSDirectoryMagnitudeAllocated, error) {
	event := new(ContractIAVSDirectoryMagnitudeAllocated)
	if err := _ContractIAVSDirectory.contract.UnpackLog(event, "MagnitudeAllocated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIAVSDirectoryMagnitudeDeallocationCompletedIterator is returned from FilterMagnitudeDeallocationCompleted and is used to iterate over the raw logs and unpacked data for MagnitudeDeallocationCompleted events raised by the ContractIAVSDirectory contract.
type ContractIAVSDirectoryMagnitudeDeallocationCompletedIterator struct {
	Event *ContractIAVSDirectoryMagnitudeDeallocationCompleted // Event containing the contract specifics and raw log

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
func (it *ContractIAVSDirectoryMagnitudeDeallocationCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIAVSDirectoryMagnitudeDeallocationCompleted)
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
		it.Event = new(ContractIAVSDirectoryMagnitudeDeallocationCompleted)
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
func (it *ContractIAVSDirectoryMagnitudeDeallocationCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIAVSDirectoryMagnitudeDeallocationCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIAVSDirectoryMagnitudeDeallocationCompleted represents a MagnitudeDeallocationCompleted event raised by the ContractIAVSDirectory contract.
type ContractIAVSDirectoryMagnitudeDeallocationCompleted struct {
	Operator           common.Address
	Strategy           common.Address
	OperatorSet        IAVSDirectoryOperatorSet
	FreeMagnitudeAdded uint64
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterMagnitudeDeallocationCompleted is a free log retrieval operation binding the contract event 0x1e5c8e13c62c31d6252ac205e592477d643c7e95831d5b46d99a3c60c2fad8db.
//
// Solidity: event MagnitudeDeallocationCompleted(address operator, address strategy, (address,uint32) operatorSet, uint64 freeMagnitudeAdded)
func (_ContractIAVSDirectory *ContractIAVSDirectoryFilterer) FilterMagnitudeDeallocationCompleted(opts *bind.FilterOpts) (*ContractIAVSDirectoryMagnitudeDeallocationCompletedIterator, error) {

	logs, sub, err := _ContractIAVSDirectory.contract.FilterLogs(opts, "MagnitudeDeallocationCompleted")
	if err != nil {
		return nil, err
	}
	return &ContractIAVSDirectoryMagnitudeDeallocationCompletedIterator{contract: _ContractIAVSDirectory.contract, event: "MagnitudeDeallocationCompleted", logs: logs, sub: sub}, nil
}

// WatchMagnitudeDeallocationCompleted is a free log subscription operation binding the contract event 0x1e5c8e13c62c31d6252ac205e592477d643c7e95831d5b46d99a3c60c2fad8db.
//
// Solidity: event MagnitudeDeallocationCompleted(address operator, address strategy, (address,uint32) operatorSet, uint64 freeMagnitudeAdded)
func (_ContractIAVSDirectory *ContractIAVSDirectoryFilterer) WatchMagnitudeDeallocationCompleted(opts *bind.WatchOpts, sink chan<- *ContractIAVSDirectoryMagnitudeDeallocationCompleted) (event.Subscription, error) {

	logs, sub, err := _ContractIAVSDirectory.contract.WatchLogs(opts, "MagnitudeDeallocationCompleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIAVSDirectoryMagnitudeDeallocationCompleted)
				if err := _ContractIAVSDirectory.contract.UnpackLog(event, "MagnitudeDeallocationCompleted", log); err != nil {
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

// ParseMagnitudeDeallocationCompleted is a log parse operation binding the contract event 0x1e5c8e13c62c31d6252ac205e592477d643c7e95831d5b46d99a3c60c2fad8db.
//
// Solidity: event MagnitudeDeallocationCompleted(address operator, address strategy, (address,uint32) operatorSet, uint64 freeMagnitudeAdded)
func (_ContractIAVSDirectory *ContractIAVSDirectoryFilterer) ParseMagnitudeDeallocationCompleted(log types.Log) (*ContractIAVSDirectoryMagnitudeDeallocationCompleted, error) {
	event := new(ContractIAVSDirectoryMagnitudeDeallocationCompleted)
	if err := _ContractIAVSDirectory.contract.UnpackLog(event, "MagnitudeDeallocationCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIAVSDirectoryMagnitudeQueueDeallocatedIterator is returned from FilterMagnitudeQueueDeallocated and is used to iterate over the raw logs and unpacked data for MagnitudeQueueDeallocated events raised by the ContractIAVSDirectory contract.
type ContractIAVSDirectoryMagnitudeQueueDeallocatedIterator struct {
	Event *ContractIAVSDirectoryMagnitudeQueueDeallocated // Event containing the contract specifics and raw log

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
func (it *ContractIAVSDirectoryMagnitudeQueueDeallocatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIAVSDirectoryMagnitudeQueueDeallocated)
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
		it.Event = new(ContractIAVSDirectoryMagnitudeQueueDeallocated)
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
func (it *ContractIAVSDirectoryMagnitudeQueueDeallocatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIAVSDirectoryMagnitudeQueueDeallocatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIAVSDirectoryMagnitudeQueueDeallocated represents a MagnitudeQueueDeallocated event raised by the ContractIAVSDirectory contract.
type ContractIAVSDirectoryMagnitudeQueueDeallocated struct {
	Operator              common.Address
	Strategy              common.Address
	OperatorSet           IAVSDirectoryOperatorSet
	MagnitudeToDeallocate uint64
	CompletableTimestamp  uint32
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterMagnitudeQueueDeallocated is a free log retrieval operation binding the contract event 0x2e68db1fe51107d7e451ae268d1631796989ab9d7925054e9b247854cb5be950.
//
// Solidity: event MagnitudeQueueDeallocated(address operator, address strategy, (address,uint32) operatorSet, uint64 magnitudeToDeallocate, uint32 completableTimestamp)
func (_ContractIAVSDirectory *ContractIAVSDirectoryFilterer) FilterMagnitudeQueueDeallocated(opts *bind.FilterOpts) (*ContractIAVSDirectoryMagnitudeQueueDeallocatedIterator, error) {

	logs, sub, err := _ContractIAVSDirectory.contract.FilterLogs(opts, "MagnitudeQueueDeallocated")
	if err != nil {
		return nil, err
	}
	return &ContractIAVSDirectoryMagnitudeQueueDeallocatedIterator{contract: _ContractIAVSDirectory.contract, event: "MagnitudeQueueDeallocated", logs: logs, sub: sub}, nil
}

// WatchMagnitudeQueueDeallocated is a free log subscription operation binding the contract event 0x2e68db1fe51107d7e451ae268d1631796989ab9d7925054e9b247854cb5be950.
//
// Solidity: event MagnitudeQueueDeallocated(address operator, address strategy, (address,uint32) operatorSet, uint64 magnitudeToDeallocate, uint32 completableTimestamp)
func (_ContractIAVSDirectory *ContractIAVSDirectoryFilterer) WatchMagnitudeQueueDeallocated(opts *bind.WatchOpts, sink chan<- *ContractIAVSDirectoryMagnitudeQueueDeallocated) (event.Subscription, error) {

	logs, sub, err := _ContractIAVSDirectory.contract.WatchLogs(opts, "MagnitudeQueueDeallocated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIAVSDirectoryMagnitudeQueueDeallocated)
				if err := _ContractIAVSDirectory.contract.UnpackLog(event, "MagnitudeQueueDeallocated", log); err != nil {
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

// ParseMagnitudeQueueDeallocated is a log parse operation binding the contract event 0x2e68db1fe51107d7e451ae268d1631796989ab9d7925054e9b247854cb5be950.
//
// Solidity: event MagnitudeQueueDeallocated(address operator, address strategy, (address,uint32) operatorSet, uint64 magnitudeToDeallocate, uint32 completableTimestamp)
func (_ContractIAVSDirectory *ContractIAVSDirectoryFilterer) ParseMagnitudeQueueDeallocated(log types.Log) (*ContractIAVSDirectoryMagnitudeQueueDeallocated, error) {
	event := new(ContractIAVSDirectoryMagnitudeQueueDeallocated)
	if err := _ContractIAVSDirectory.contract.UnpackLog(event, "MagnitudeQueueDeallocated", log); err != nil {
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
	OperatorSet IAVSDirectoryOperatorSet
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
	OperatorSet IAVSDirectoryOperatorSet
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
	OperatorSet IAVSDirectoryOperatorSet
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

// ContractIAVSDirectoryOperatorSlashedIterator is returned from FilterOperatorSlashed and is used to iterate over the raw logs and unpacked data for OperatorSlashed events raised by the ContractIAVSDirectory contract.
type ContractIAVSDirectoryOperatorSlashedIterator struct {
	Event *ContractIAVSDirectoryOperatorSlashed // Event containing the contract specifics and raw log

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
func (it *ContractIAVSDirectoryOperatorSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIAVSDirectoryOperatorSlashed)
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
		it.Event = new(ContractIAVSDirectoryOperatorSlashed)
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
func (it *ContractIAVSDirectoryOperatorSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIAVSDirectoryOperatorSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIAVSDirectoryOperatorSlashed represents a OperatorSlashed event raised by the ContractIAVSDirectory contract.
type ContractIAVSDirectoryOperatorSlashed struct {
	Operator      common.Address
	OperatorSetId uint32
	Strategy      common.Address
	BipsToSlash   uint16
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOperatorSlashed is a free log retrieval operation binding the contract event 0xe672839d3c371691acdb52de9fefc94b3dbf407dc0920ef566c7c059ad575b1c.
//
// Solidity: event OperatorSlashed(address operator, uint32 operatorSetId, address strategy, uint16 bipsToSlash)
func (_ContractIAVSDirectory *ContractIAVSDirectoryFilterer) FilterOperatorSlashed(opts *bind.FilterOpts) (*ContractIAVSDirectoryOperatorSlashedIterator, error) {

	logs, sub, err := _ContractIAVSDirectory.contract.FilterLogs(opts, "OperatorSlashed")
	if err != nil {
		return nil, err
	}
	return &ContractIAVSDirectoryOperatorSlashedIterator{contract: _ContractIAVSDirectory.contract, event: "OperatorSlashed", logs: logs, sub: sub}, nil
}

// WatchOperatorSlashed is a free log subscription operation binding the contract event 0xe672839d3c371691acdb52de9fefc94b3dbf407dc0920ef566c7c059ad575b1c.
//
// Solidity: event OperatorSlashed(address operator, uint32 operatorSetId, address strategy, uint16 bipsToSlash)
func (_ContractIAVSDirectory *ContractIAVSDirectoryFilterer) WatchOperatorSlashed(opts *bind.WatchOpts, sink chan<- *ContractIAVSDirectoryOperatorSlashed) (event.Subscription, error) {

	logs, sub, err := _ContractIAVSDirectory.contract.WatchLogs(opts, "OperatorSlashed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIAVSDirectoryOperatorSlashed)
				if err := _ContractIAVSDirectory.contract.UnpackLog(event, "OperatorSlashed", log); err != nil {
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

// ParseOperatorSlashed is a log parse operation binding the contract event 0xe672839d3c371691acdb52de9fefc94b3dbf407dc0920ef566c7c059ad575b1c.
//
// Solidity: event OperatorSlashed(address operator, uint32 operatorSetId, address strategy, uint16 bipsToSlash)
func (_ContractIAVSDirectory *ContractIAVSDirectoryFilterer) ParseOperatorSlashed(log types.Log) (*ContractIAVSDirectoryOperatorSlashed, error) {
	event := new(ContractIAVSDirectoryOperatorSlashed)
	if err := _ContractIAVSDirectory.contract.UnpackLog(event, "OperatorSlashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
