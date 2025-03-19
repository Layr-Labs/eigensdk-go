// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractSlashingRegistryCoordinator

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

// BN254G1Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G1Point struct {
	X *big.Int
	Y *big.Int
}

// ISlashingRegistryCoordinatorTypesOperatorInfo is an auto generated low-level Go binding around an user-defined struct.
type ISlashingRegistryCoordinatorTypesOperatorInfo struct {
	OperatorId [32]byte
	Status     uint8
}

// ISlashingRegistryCoordinatorTypesOperatorKickParam is an auto generated low-level Go binding around an user-defined struct.
type ISlashingRegistryCoordinatorTypesOperatorKickParam struct {
	QuorumNumber uint8
	Operator     common.Address
}

// ISlashingRegistryCoordinatorTypesOperatorSetParam is an auto generated low-level Go binding around an user-defined struct.
type ISlashingRegistryCoordinatorTypesOperatorSetParam struct {
	MaxOperatorCount        uint32
	KickBIPsOfOperatorStake uint16
	KickBIPsOfTotalStake    uint16
}

// ISlashingRegistryCoordinatorTypesQuorumBitmapUpdate is an auto generated low-level Go binding around an user-defined struct.
type ISlashingRegistryCoordinatorTypesQuorumBitmapUpdate struct {
	UpdateBlockNumber     uint32
	NextUpdateBlockNumber uint32
	QuorumBitmap          *big.Int
}

// IStakeRegistryTypesStrategyParams is an auto generated low-level Go binding around an user-defined struct.
type IStakeRegistryTypesStrategyParams struct {
	Strategy   common.Address
	Multiplier *big.Int
}

// ContractSlashingRegistryCoordinatorMetaData contains all meta data concerning the ContractSlashingRegistryCoordinator contract.
var ContractSlashingRegistryCoordinatorMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_stakeRegistry\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"},{\"name\":\"_blsApkRegistry\",\"type\":\"address\",\"internalType\":\"contractIBLSApkRegistry\"},{\"name\":\"_indexRegistry\",\"type\":\"address\",\"internalType\":\"contractIIndexRegistry\"},{\"name\":\"_socketRegistry\",\"type\":\"address\",\"internalType\":\"contractISocketRegistry\"},{\"name\":\"_allocationManager\",\"type\":\"address\",\"internalType\":\"contractIAllocationManager\"},{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"_version\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"OPERATOR_CHURN_APPROVAL_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PUBKEY_REGISTRATION_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allocationManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIAllocationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"avs\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"blsApkRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBLSApkRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateOperatorChurnApprovalDigestHash\",\"inputs\":[{\"name\":\"registeringOperator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"registeringOperatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"operatorKickParams\",\"type\":\"tuple[]\",\"internalType\":\"structISlashingRegistryCoordinatorTypes.OperatorKickParam[]\",\"components\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculatePubkeyRegistrationMessageHash\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"churnApprover\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createSlashableStakeQuorum\",\"inputs\":[{\"name\":\"operatorSetParams\",\"type\":\"tuple\",\"internalType\":\"structISlashingRegistryCoordinatorTypes.OperatorSetParam\",\"components\":[{\"name\":\"maxOperatorCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"kickBIPsOfOperatorStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"kickBIPsOfTotalStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"name\":\"minimumStake\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"strategyParams\",\"type\":\"tuple[]\",\"internalType\":\"structIStakeRegistryTypes.StrategyParams[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"lookAheadPeriod\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createTotalDelegatedStakeQuorum\",\"inputs\":[{\"name\":\"operatorSetParams\",\"type\":\"tuple\",\"internalType\":\"structISlashingRegistryCoordinatorTypes.OperatorSetParam\",\"components\":[{\"name\":\"maxOperatorCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"kickBIPsOfOperatorStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"kickBIPsOfTotalStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"name\":\"minimumStake\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"strategyParams\",\"type\":\"tuple[]\",\"internalType\":\"structIStakeRegistryTypes.StrategyParams[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deregisterOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"domainSeparator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"eip712Domain\",\"inputs\":[],\"outputs\":[{\"name\":\"fields\",\"type\":\"bytes1\",\"internalType\":\"bytes1\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"verifyingContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"extensions\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ejectOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ejectionCooldown\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ejector\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentQuorumBitmap\",\"inputs\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint192\",\"internalType\":\"uint192\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structISlashingRegistryCoordinatorTypes.OperatorInfo\",\"components\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumISlashingRegistryCoordinatorTypes.OperatorStatus\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorFromId\",\"inputs\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorId\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorSetParams\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structISlashingRegistryCoordinatorTypes.OperatorSetParam\",\"components\":[{\"name\":\"maxOperatorCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"kickBIPsOfOperatorStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"kickBIPsOfTotalStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorStatus\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumISlashingRegistryCoordinatorTypes.OperatorStatus\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumBitmapAtBlockNumberByIndex\",\"inputs\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint192\",\"internalType\":\"uint192\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumBitmapHistoryLength\",\"inputs\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumBitmapIndicesAtBlockNumber\",\"inputs\":[{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumBitmapUpdateByIndex\",\"inputs\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structISlashingRegistryCoordinatorTypes.QuorumBitmapUpdate\",\"components\":[{\"name\":\"updateBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"nextUpdateBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumBitmap\",\"type\":\"uint192\",\"internalType\":\"uint192\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"indexRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIIndexRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"churnApprover\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"ejector\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isChurnApproverSaltUsed\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lastEjectionTimestamp\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pubkeyRegistrationMessageHash\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"quorumCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"quorumUpdateBlockNumber\",\"inputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setAVS\",\"inputs\":[{\"name\":\"_avs\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChurnApprover\",\"inputs\":[{\"name\":\"_churnApprover\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setEjectionCooldown\",\"inputs\":[{\"name\":\"_ejectionCooldown\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setEjector\",\"inputs\":[{\"name\":\"_ejector\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setOperatorSetParams\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"operatorSetParams\",\"type\":\"tuple\",\"internalType\":\"structISlashingRegistryCoordinatorTypes.OperatorSetParam\",\"components\":[{\"name\":\"maxOperatorCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"kickBIPsOfOperatorStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"kickBIPsOfTotalStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"socketRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractISocketRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stakeRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"supportsAVS\",\"inputs\":[{\"name\":\"_avs\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateOperators\",\"inputs\":[{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateOperatorsForQuorum\",\"inputs\":[{\"name\":\"operatorsPerQuorum\",\"type\":\"address[][]\",\"internalType\":\"address[][]\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateSocket\",\"inputs\":[{\"name\":\"socket\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"AVSUpdated\",\"inputs\":[{\"name\":\"prevAVS\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newAVS\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChurnApproverUpdated\",\"inputs\":[{\"name\":\"prevChurnApprover\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newChurnApprover\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EIP712DomainChanged\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EjectionCooldownUpdated\",\"inputs\":[{\"name\":\"prevEjectionCooldown\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newEjectionCooldown\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EjectorUpdated\",\"inputs\":[{\"name\":\"prevEjector\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newEjector\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorDeregistered\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorRegistered\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSetParamsUpdated\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"uint8\"},{\"name\":\"operatorSetParams\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structISlashingRegistryCoordinatorTypes.OperatorSetParam\",\"components\":[{\"name\":\"maxOperatorCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"kickBIPsOfOperatorStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"kickBIPsOfTotalStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSocketUpdate\",\"inputs\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"socket\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"QuorumBlockNumberUpdated\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"uint8\"},{\"name\":\"blocknumber\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"QuorumCreated\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"uint8\"},{\"name\":\"operatorSetParams\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structISlashingRegistryCoordinatorTypes.OperatorSetParam\",\"components\":[{\"name\":\"maxOperatorCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"kickBIPsOfOperatorStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"kickBIPsOfTotalStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"name\":\"minimumStake\",\"type\":\"uint96\",\"indexed\":false,\"internalType\":\"uint96\"},{\"name\":\"strategyParams\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"structIStakeRegistryTypes.StrategyParams[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"stakeType\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumIStakeRegistryTypes.StakeType\"},{\"name\":\"lookAheadPeriod\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyRegisteredForQuorums\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BitmapCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BitmapEmpty\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BitmapUpdateIsAfterBlockNumber\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BitmapValueTooLarge\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BytesArrayLengthTooLong\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BytesArrayNotOrdered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotChurnSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotKickOperatorAboveThreshold\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotReregisterYet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChurnApproverSaltUsed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CurrentlyPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpModFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputAddressZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputLengthMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientStakeForChurn\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAVS\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidNewPausedStatus\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRegistrationType\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidShortString\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LookAheadPeriodTooLong\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MaxOperatorCountReached\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MaxQuorumsReached\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NextBitmapUpdateIsBeforeBlockNumber\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotRegisteredForQuorum\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotSorted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyAllocationManager\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyEjector\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyPauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyUnpauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorNotRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"QuorumDoesNotExist\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"QuorumOperatorCountMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SignatureExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StringTooLong\",\"inputs\":[{\"name\":\"str\",\"type\":\"string\",\"internalType\":\"string\"}]}]",
	Bin: "0x610160604052348015610010575f5ffd5b50604051615b88380380615b8883398101604081905261002f916101d9565b6001600160a01b0380881660c05280871660a05280861660e0528085166080528316610100528181610060816100ad565b61012052506001600160a01b03811661008c576040516339b190bb60e11b815260040160405180910390fd5b6001600160a01b0316610140526100a16100f3565b50505050505050610357565b5f5f829050601f815111156100e0578260405163305a27a960e01b81526004016100d791906102fc565b60405180910390fd5b80516100eb82610331565b179392505050565b603254610100900460ff161561015b5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b60648201526084016100d7565b60325460ff908116146101ac576032805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b03811681146101c2575f5ffd5b50565b634e487b7160e01b5f52604160045260245ffd5b5f5f5f5f5f5f5f60e0888a0312156101ef575f5ffd5b87516101fa816101ae565b602089015190975061020b816101ae565b604089015190965061021c816101ae565b606089015190955061022d816101ae565b608089015190945061023e816101ae565b60a089015190935061024f816101ae565b60c08901519092506001600160401b0381111561026a575f5ffd5b8801601f81018a1361027a575f5ffd5b80516001600160401b03811115610293576102936101c5565b604051601f8201601f19908116603f011681016001600160401b03811182821017156102c1576102c16101c5565b6040528181528282016020018c10156102d8575f5ffd5b8160208401602083015e5f6020838301015280935050505092959891949750929550565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b80516020808301519190811015610351575f198160200360031b1b821691505b50919050565b60805160a05160c05160e05161010051610120516101405161573f6104495f395f818161069b01528181610b87015281816113aa015261185b01525f6112f801525f81816107c801528181611e7f015281816124310152818161255c015261355a01525f81816106f101528181610f2e015281816121b1015281816126960152612f9d01525f81816105f201528181611aaa01528181612136015281816124cf0152818161261501528181612efb015261382501525f81816105cb01528181610cf0015281816120bd0152818161270b01528181612c100152612e8201525f818161087d0152611c77015261573f5ff3fe608060405234801561000f575f5ffd5b50600436106102e3575f3560e01c80635df459461161018d578063a65497c6116100e0578063de1164bb1161008f578063de1164bb146107ea578063e65797ad146107fd578063ea32afae14610878578063f2fde38b1461089f578063f698da25146108b2578063fabc1cbc146108ba578063fd39105a146108cd575f5ffd5b8063a65497c61461073a578063a96f783e1461074d578063b526578714610756578063c391425e14610769578063c63fd50214610789578063ca0de8821461079c578063ca8aa7c7146107c3575f5ffd5b806384ca52131161013c57806384ca521314610670578063871ef04914610683578063886f1195146106965780638da5cb5b146106bd5780639aa1653d146106ce5780639e9923c2146106ec5780639feab85914610713575f5ffd5b80635df45946146105c657806368304835146105ed5780636e3b17db14610614578063715018a614610627578063734479921461062f5780638281ab751461064257806384b0196e14610655575f5ffd5b8063296bb06411610245578063530b97a4116101f4578063530b97a41461053857806354fd4d501461054b5780635865c60c14610560578063595c6a67146105805780635ac86ab7146105885780635b0b829f146105ab5780635c975abb146105be575f5ffd5b8063296bb064146104a657806329d1e0c3146104b95780632cdd1e86146104cc578063303ca956146104df5780633c2a7f4c146104f25780633eef3a51146105125780635140a54814610525575f5ffd5b8063125e0584116102a1578063125e05841461039f57806313542a4e146103be578063136439dd146103e65780631478851f146103f95780631eb812da1461042b578063249a0c421461047457806328f61b3114610493575f5ffd5b8062cf2ab5146102e757806303fd3492146102fc57806304ec63511461032e578063054310e6146103595780630cf4b767146103795780630d3f21341461038c575b5f5ffd5b6102fa6102f53660046140f2565b610908565b005b61031b61030a366004614123565b5f9081526002602052604090205490565b6040519081526020015b60405180910390f35b61034161033c36600461414b565b610aa9565b6040516001600160c01b039091168152602001610325565b60075461036c906001600160a01b031681565b6040516103259190614180565b6102fa610387366004614205565b610ac1565b6102fa61039a366004614123565b610b23565b61031b6103ad366004614236565b60096020525f908152604090205481565b61031b6103cc366004614236565b6001600160a01b03165f9081526003602052604090205490565b6102fa6103f4366004614123565b610b70565b61041b610407366004614123565b60046020525f908152604090205460ff1681565b6040519015158152602001610325565b61043e610439366004614251565b610c4a565b60408051825163ffffffff908116825260208085015190911690820152918101516001600160c01b031690820152606001610325565b61031b610482366004614286565b60056020525f908152604090205481565b60085461036c906001600160a01b031681565b61036c6104b4366004614123565b610cd8565b6102fa6104c7366004614236565b610d61565b6102fa6104da366004614236565b610d72565b6102fa6104ed366004614303565b610d83565b610505610500366004614236565b610df7565b6040516103259190614360565b6102fa610520366004614481565b610e1b565b6102fa61053336600461452d565b610e37565b6102fa61054636600461460a565b61115b565b6105536112f1565b604051610325919061469c565b61057361056e366004614236565b611321565b60405161032591906146d6565b6102fa611393565b61041b610596366004614286565b603354600160ff9092169190911b9081161490565b6102fa6105b93660046146f1565b611447565b60335461031b565b61036c7f000000000000000000000000000000000000000000000000000000000000000081565b61036c7f000000000000000000000000000000000000000000000000000000000000000081565b6102fa610622366004614723565b611463565b6102fa61148f565b61031b61063d366004614236565b6114a0565b6102fa61065036600461476f565b611504565b61065d611519565b60405161032597969594939291906147c1565b61031b61067e3660046148eb565b6115b2565b610341610691366004614123565b6115fb565b61036c7f000000000000000000000000000000000000000000000000000000000000000081565b6096546001600160a01b031661036c565b5f546106da9060ff1681565b60405160ff9091168152602001610325565b61036c7f000000000000000000000000000000000000000000000000000000000000000081565b61031b7f2bd82124057f0913bc3b772ce7b83e8057c1ad1f3510fc83778be20f10ec5de681565b6102fa610748366004614236565b611605565b61031b600a5481565b61041b610764366004614236565b611616565b61077c610777366004614950565b61162a565b6040516103259190614a35565b6102fa610797366004614a47565b611638565b61031b7f4d404e3276e7ac2163d8ee476afa6a41d1f68fb71f2d8b6546b24e55ce01b72a81565b61036c7f000000000000000000000000000000000000000000000000000000000000000081565b600b5461036c906001600160a01b031681565b61086b61080b366004614286565b60408051606080820183525f808352602080840182905292840181905260ff9490941684526001825292829020825193840183525463ffffffff8116845261ffff600160201b8204811692850192909252600160301b9004169082015290565b6040516103259190614ad4565b61036c7f000000000000000000000000000000000000000000000000000000000000000081565b6102fa6108ad366004614236565b6117da565b61031b611850565b6102fa6108c8366004614123565b611859565b6108fb6108db366004614236565b6001600160a01b03165f9081526003602052604090206001015460ff1690565b6040516103259190614b07565b6033546002906004908116036109315760405163840a48d560e01b815260040160405180910390fd5b5f5b8251811015610aa4576040805160018082528183019092525f916020808301908036833701905050905083828151811061096f5761096f614b15565b6020026020010151815f8151811061098957610989614b15565b6001600160a01b0392909216602092830291909101909101526040805160018082528183019092525f9181602001602082028036833701905050905060035f8685815181106109da576109da614b15565b60200260200101516001600160a01b03166001600160a01b031681526020019081526020015f205f0154815f81518110610a1657610a16614b15565b6020026020010181815250505f610a45825f81518110610a3857610a38614b15565b6020026020010151611970565b90505f610a5a826001600160c01b031661197c565b90505f5b8151811015610a9357610a8b8585848481518110610a7e57610a7e614b15565b016020015160f81c611a45565b600101610a5e565b505060019093019250610933915050565b505050565b5f610ab76002858585611b7c565b90505b9392505050565b6001335f9081526003602052604090206001015460ff166002811115610ae957610ae96146ae565b14610b075760405163aba4733960e01b815260040160405180910390fd5b335f90815260036020526040902054610b209082611c60565b50565b610b2b611d0b565b600a80549082905560408051828152602081018490527fa77a91bea7b6d95a8eb5a54878a1d9e3c875e26c86a9b70e3420c5c5db193b62910160405180910390a15050565b60405163237dfb4760e11b81526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906346fbf68e90610bbc903390600401614180565b602060405180830381865afa158015610bd7573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610bfb9190614b38565b610c1857604051631d77d47760e21b815260040160405180910390fd5b6033548181168114610c3d5760405163c61dca5d60e01b815260040160405180910390fd5b610c4682611d65565b5050565b604080516060810182525f80825260208201819052918101919091525f838152600260205260409020805483908110610c8557610c85614b15565b5f91825260209182902060408051606081018252919092015463ffffffff8082168352600160201b820416938201939093526001600160c01b03600160401b909304929092169082015290505b92915050565b6040516308f6629d60e31b8152600481018290525f907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906347b314e890602401602060405180830381865afa158015610d3d573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610cd29190614b51565b610d69611d0b565b610b2081611da2565b610d7a611d0b565b610b2081611e0b565b610d8b611e74565b603354600190600290811603610db45760405163840a48d560e01b815260040160405180910390fd5b610dbd83611616565b610dda576040516366e565df60e01b815260040160405180910390fd5b5f610de483611ebd565b9050610df08582611f65565b5050505050565b604080518082019091525f8082526020820152610cd2610e16836114a0565b612213565b610e23611d0b565b610e3184848460018561229d565b50505050565b603354600290600490811603610e605760405163840a48d560e01b815260040160405180910390fd5b610ea283838080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201829052505460ff1692506127ba915050565b5083518214610ec45760405163aaad13f760e01b815260040160405180910390fd5b5f5b82811015610df0575f848483818110610ee157610ee1614b15565b885192013560f81c92505f9188915084908110610f0057610f00614b15565b60209081029190910101516040516379a0849160e11b815260ff841660048201529091506001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063f341092290602401602060405180830381865afa158015610f73573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610f979190614b6c565b63ffffffff16815114610fbd57604051638e5aeee760e01b815260040160405180910390fd5b5f81516001600160401b03811115610fd757610fd7613fbb565b604051908082528060200260200182016040528015611000578160200160208202803683370190505b5090505f805b83518110156110f6575f84828151811061102257611022614b15565b6020026020010151905060035f826001600160a01b03166001600160a01b031681526020019081526020015f205f015484838151811061106457611064614b15565b6020026020010181815250505f611086858481518110610a3857610a38614b15565b905060016001600160c01b03821660ff89161c8116146110b95760405163d053aa2160e01b815260040160405180910390fd5b836001600160a01b0316826001600160a01b0316116110eb5760405163ba50f91160e01b815260040160405180910390fd5b509150600101611006565b50611102838386611a45565b60ff84165f81815260056020908152604091829020439081905591519182527f46077d55330763f16269fd75e5761663f4192d2791747c0189b16ad31db07db4910160405180910390a250505050806001019050610ec6565b603254610100900460ff161580801561117b5750603254600160ff909116105b806111955750303b158015611195575060325460ff166001145b6111fd5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b6032805460ff191660011790558015611220576032805461ff0019166101001790555b6112766040518060400160405280601681526020017520ab29a932b3b4b9ba393ca1b7b7b93234b730ba37b960511b8152506040518060400160405280600681526020016576302e302e3160d01b8152506127ee565b61127f8661281f565b61128885611da2565b61129183611d65565b61129a84611e0b565b6112a382612870565b80156112e9576032805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050505050565b606061131c7f00000000000000000000000000000000000000000000000000000000000000006128dc565b905090565b604080518082019091525f80825260208201526001600160a01b0382165f908152600360209081526040918290208251808401909352805483526001810154909183019060ff166002811115611379576113796146ae565b600281111561138a5761138a6146ae565b90525092915050565b60405163237dfb4760e11b81526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906346fbf68e906113df903390600401614180565b602060405180830381865afa1580156113fa573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061141e9190614b38565b61143b57604051631d77d47760e21b815260040160405180910390fd5b6114455f19611d65565b565b61144f611d0b565b8161145981612919565b610aa48383612941565b61146b6129d2565b6001600160a01b0382165f908152600960205260409020429055610c4682826129fd565b611497611d0b565b6114455f61281f565b5f610cd27f2bd82124057f0913bc3b772ce7b83e8057c1ad1f3510fc83778be20f10ec5de6836040516020016114e99291909182526001600160a01b0316602082015260400190565b60405160208183030381529060405280519060200120612a93565b61150c611d0b565b610aa48383835f5f61229d565b5f6060805f5f5f606060c8545f5f1b148015611535575060c954155b6115795760405162461bcd60e51b81526020600482015260156024820152741152540dcc4c8e88155b9a5b9a5d1a585b1a5e9959605a1b60448201526064016111f4565b611581612abf565b611589612b4f565b604080515f80825260208201909252600f60f81b9b939a50919850469750309650945092509050565b5f6115f17f4d404e3276e7ac2163d8ee476afa6a41d1f68fb71f2d8b6546b24e55ce01b72a87878787876040516020016114e996959493929190614b87565b9695505050505050565b5f610cd282611970565b61160d611d0b565b610b2081612870565b600b546001600160a01b0391821691161490565b6060610aba60028484612b5e565b611640611e74565b6033545f906001908116036116685760405163840a48d560e01b815260040160405180910390fd5b61167185611616565b61168e576040516366e565df60e01b815260040160405180910390fd5b5f61169885611ebd565b90505f80806116a986880188614d16565b9250925092505f6116ba8b83612c0d565b90505f8460018111156116cf576116cf6146ae565b03611778575f6116e38c8388876001612ca6565b5190505f5b8651811015611771575f87828151811061170457611704614b15565b0160209081015160f81c5f8181526001909252604090912054845191925063ffffffff169084908490811061173b5761173b614b15565b602002602001015163ffffffff1611156117685760405163c6b9e76760e01b815260040160405180910390fd5b506001016116e8565b50506117cd565b600184600181111561178c5761178c6146ae565b036117b4575f8061179f898b018b614d71565b945094505050506117718d84898886866130ed565b60405163354bb8ab60e01b815260040160405180910390fd5b5050505050505050505050565b6117e2611d0b565b6001600160a01b0381166118475760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016111f4565b610b208161281f565b5f61131c6132e2565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156118b5573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906118d99190614b51565b6001600160a01b0316336001600160a01b03161461190a5760405163794821ff60e01b815260040160405180910390fd5b603354801982198116146119315760405163c61dca5d60e01b815260040160405180910390fd5b603382905560405182815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c906020015b60405180910390a25050565b5f610cd26002836132eb565b60605f5f61198984613355565b61ffff166001600160401b038111156119a4576119a4613fbb565b6040519080825280601f01601f1916602001820160405280156119ce576020820181803683370190505b5090505f805b8251821080156119e5575061010081105b15611a3b576001811b935085841615611a2b578060f81b838381518110611a0e57611a0e614b15565b60200101906001600160f81b03191690815f1a9053508160010191505b611a3481614e7e565b90506119d4565b5090949350505050565b6040805160018082528183019092525f916020820181803683370190505090508160f81b815f81518110611a7b57611a7b614b15565b60200101906001600160f81b03191690815f1a905350604051636c3fb4bf60e01b81525f906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690636c3fb4bf90611ae390889088908890600401614e96565b5f604051808303815f875af1158015611afe573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052611b259190810190614f26565b90505f5b85518110156112e957818181518110611b4457611b44614b15565b602002602001015115611b7457611b74868281518110611b6657611b66614b15565b6020026020010151846129fd565b600101611b29565b5f838152602085905260408120805482919084908110611b9e57611b9e614b15565b5f91825260209182902060408051606081018252929091015463ffffffff808216808552600160201b8304821695850195909552600160401b9091046001600160c01b03169183019190915290925085161015611c0e57604051636cb19aff60e01b815260040160405180910390fd5b602081015163ffffffff161580611c345750806020015163ffffffff168463ffffffff16105b611c515760405163bbba60cb60e01b815260040160405180910390fd5b6040015190505b949350505050565b6040516378219b3f60e11b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063f043367e90611cae9085908590600401614fb3565b5f604051808303815f87803b158015611cc5575f5ffd5b505af1158015611cd7573d5f5f3e3d5ffd5b50505050817fec2963ab21c1e50e1e582aa542af2e4bf7bf38e6e1403c27b42e1c5d6e621eaa82604051611964919061469c565b6096546001600160a01b031633146114455760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016111f4565b603381905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a250565b6007546040517f315457d8a8fe60f04af17c16e2f5a5e1db612b31648e58030360759ef8f3528c91611de1916001600160a01b03909116908490614fcb565b60405180910390a1600780546001600160a01b0319166001600160a01b0392909216919091179055565b6008546040517f8f30ab09f43a6c157d7fce7e0a13c003042c1c95e8a72e7a146a21c0caa24dc991611e4a916001600160a01b03909116908490614fcb565b60405180910390a1600880546001600160a01b0319166001600160a01b0392909216919091179055565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614611445576040516323d871a560e01b815260040160405180910390fd5b60605f82516001600160401b03811115611ed957611ed9613fbb565b6040519080825280601f01601f191660200182016040528015611f03576020820181803683370190505b5090505f5b8351811015611f5e57838181518110611f2357611f23614b15565b602002602001015160f81b828281518110611f4057611f40614b15565b60200101906001600160f81b03191690815f1a905350600101611f08565b5092915050565b6001600160a01b0382165f90815260036020526040812080549091611f8982611970565b905060018084015460ff166002811115611fa557611fa56146ae565b14611fc35760405163aba4733960e01b815260040160405180910390fd5b5f8054611fd490869060ff166127ba565b90506001600160c01b038116611ffd576040516368b6a87560e11b815260040160405180910390fd5b6120146001600160c01b0382811690841681161490565b6120315760405163d053aa2160e01b815260040160405180910390fd5b6001600160c01b038181161983161661204a848261337f565b6001600160c01b0381166120a6576001600160a01b0387165f81815260036020526040808220600101805460ff19166002179055518692917f396fdcb180cb0fea26928113fb0fd1c3549863f9cd563e6a184f1d578116c8e491a35b60405163f4e24fe560e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063f4e24fe5906120f4908a908a90600401614fe5565b5f604051808303815f87803b15801561210b575f5ffd5b505af115801561211d573d5f5f3e3d5ffd5b505060405163bd29b8cd60e01b81526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016925063bd29b8cd915061216f9087908a90600401614fb3565b5f604051808303815f87803b158015612186575f5ffd5b505af1158015612198573d5f5f3e3d5ffd5b505060405163bd29b8cd60e01b81526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016925063bd29b8cd91506121ea9087908a90600401614fb3565b5f604051808303815f87803b158015612201575f5ffd5b505af11580156117cd573d5f5f3e3d5ffd5b604080518082019091525f80825260208201525f80806122405f5160206156ea5f395f51905f528661501c565b90505b61224c8161338b565b90935091505f5160206156ea5f395f51905f528283098303612284576040805180820190915290815260208101919091529392505050565b5f5160206156ea5f395f51905f52600182089050612243565b5f5460ff1660c081106122c357604051633cb89c9760e01b815260040160405180910390fd5b5f80546001919081906122da90849060ff1661502f565b92506101000a81548160ff021916908360ff1602179055506122fc8187612941565b6040805160018082528183019092525f91816020015b604080518082019091525f8152606060208201528152602001906001900390816123125790505090505f85516001600160401b0381111561235557612355613fbb565b60405190808252806020026020018201604052801561237e578160200160208202803683370190505b5090505f5b86518110156123db5786818151811061239e5761239e614b15565b60200260200101515f01518282815181106123bb576123bb614b15565b6001600160a01b0390921660209283029190910190910152600101612383565b5060405180604001604052808460ff1663ffffffff16815260200182815250825f8151811061240c5761240c614b15565b6020908102919091010152600b54604051630130fc2760e51b81526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000081169263261f84e09261246c9291909116908690600401615048565b5f604051808303815f87803b158015612483575f5ffd5b505af1158015612495573d5f5f3e3d5ffd5b505f92506124a1915050565b8560018111156124b3576124b36146ae565b0361253a57604051633aea0b9d60e11b81526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906375d4173a906125089086908b908b90600401615158565b5f604051808303815f87803b15801561251f575f5ffd5b505af1158015612531573d5f5f3e3d5ffd5b5050505061267e565b600185600181111561254e5761254e6146ae565b0361267e578363ffffffff167f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316632981eb776040518163ffffffff1660e01b8152600401602060405180830381865afa1580156125b6573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906125da9190614b6c565b63ffffffff16116125fe57604051630bd441b960e21b815260040160405180910390fd5b604051630662d3e160e51b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063cc5a7c20906126509086908b9089908c9060040161518b565b5f604051808303815f87803b158015612667575f5ffd5b505af1158015612679573d5f5f3e3d5ffd5b505050505b60405163136ca0f960e11b815260ff841660048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906326d941f2906024015f604051808303815f87803b1580156126df575f5ffd5b505af11580156126f1573d5f5f3e3d5ffd5b505060405163136ca0f960e11b815260ff861660048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031692506326d941f291506024015f604051808303815f87803b158015612756575f5ffd5b505af1158015612768573d5f5f3e3d5ffd5b505050508260ff167fa34835bc2b673ec37fcf1591a91295b163fc2e181e4ea4e733beb27de1ceac4c89898989896040516127a79594939291906151c1565b60405180910390a25b5050505050505050565b5f5f6127c584613407565b9050808360ff166001901b11610aba5760405163ca95733360e01b815260040160405180910390fd5b603254610100900460ff166128155760405162461bcd60e51b81526004016111f490615242565b610c4682826134c2565b609680546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b600b546040516001600160a01b03909116907f9770f3cadfdcbb6f93af935e86047111590c3768271d237e4a2bc0b874bed693906128b19083908590614fcb565b60405180910390a150600b80546001600160a01b0319166001600160a01b0392909216919091179055565b60605f6128e883613510565b6040805160208082528183019092529192505f91906020820181803683375050509182525060208101929092525090565b5f5460ff90811690821610610b2057604051637310cff560e11b815260040160405180910390fd5b60ff82165f8181526001602090815260409182902084518154928601518487015161ffff908116600160301b0267ffff0000000000001991909216600160201b0265ffffffffffff1990951663ffffffff909316929092179390931716919091179055517f3ee6fe8d54610244c3e9d3c066ae4aee997884aa28f10616ae821925401318ac90611964908490614ad4565b6008546001600160a01b03163314611445576040516376d8ab1760e11b815260040160405180910390fd5b6001600160a01b0382165f90815260036020526040902060018082015460ff166002811115612a2e57612a2e6146ae565b14612a4c576040516325ec6c1f60e01b815260040160405180910390fd5b80545f8054612a5f90859060ff166127ba565b90505f612a6b83611970565b9050612a846001600160c01b0383811690831681161490565b156112e9576112e98686613537565b5f610cd2612a9f6132e2565b8360405161190160f01b8152600281019290925260228201526042902090565b606060ca8054612ace9061528d565b80601f0160208091040260200160405190810160405280929190818152602001828054612afa9061528d565b8015612b455780601f10612b1c57610100808354040283529160200191612b45565b820191905f5260205f20905b815481529060010190602001808311612b2857829003601f168201915b5050505050905090565b606060cb8054612ace9061528d565b60605f82516001600160401b03811115612b7a57612b7a613fbb565b604051908082528060200260200182016040528015612ba3578160200160208202803683370190505b5090505f5b8351811015612c0457612bd58686868481518110612bc857612bc8614b15565b60200260200101516135d5565b828281518110612be757612be7614b15565b63ffffffff90921660209283029190910190910152600101612ba8565b50949350505050565b5f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166303c5a6b68484612c4887610df7565b6040518463ffffffff1660e01b8152600401612c66939291906152e7565b6020604051808303815f875af1158015612c82573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610aba9190615360565b612cca60405180606001604052806060815260200160608152602001606081525090565b5f8054612cdb90869060ff166127ba565b90505f612ce787611970565b90506001600160c01b038216612d10576040516313ca465760e01b815260040160405180910390fd5b8082166001600160c01b031615612d3a57604051630c6816cd60e01b815260040160405180910390fd5b600a546001600160a01b0389165f908152600960205260409020546001600160c01b0383811690851617914291612d719190615377565b10612d8f57604051631968677d60e11b815260040160405180910390fd5b612d99888261337f565b612da38887611c60565b60016001600160a01b038a165f9081526003602052604090206001015460ff166002811115612dd457612dd46146ae565b14612e6b57604080518082018252898152600160208083018281526001600160a01b038e165f908152600390925293902082518155925183820180549394939192909160ff191690836002811115612e2e57612e2e6146ae565b0217905550506040518991506001600160a01b038b16907fe8e68cef1c3a761ed7be7e8463a375f27f7bc335e51824223cacce636ec5c3fe905f90a35b604051631fd93ca960e11b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690633fb2795290612eb9908c908b90600401614fe5565b5f604051808303815f87803b158015612ed0575f5ffd5b505af1158015612ee2573d5f5f3e3d5ffd5b5050604051632550477760e01b81526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016925063255047779150612f36908c908c908c9060040161538a565b5f604051808303815f875af1158015612f51573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052612f789190810190615414565b60408087019190915260208601919091525162bff04d60e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169062bff04d90612fd3908b908b90600401614fb3565b5f604051808303815f875af1158015612fee573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052613015919081019061546d565b845284156130e1575f5b87518110156130df575f60015f8a848151811061303e5761303e614b15565b0160209081015160f81c82528181019290925260409081015f208151606081018352905463ffffffff811680835261ffff600160201b8304811695840195909552600160301b909104909316918101919091528751805191935090849081106130a9576130a9614b15565b602002602001015163ffffffff1611156130d65760405163c6b9e76760e01b815260040160405180910390fd5b5060010161301f565b505b50505095945050505050565b835182511461310f5760405163aaad13f760e01b815260040160405180910390fd5b61311b868684846136ed565b5f613129878787875f612ca6565b90505f5b85518110156127b0575f60015f88848151811061314c5761314c614b15565b0160209081015160f81c82528181019290925260409081015f208151606081018352905463ffffffff811680835261ffff600160201b8304811695840195909552600160301b909104909316918101919091528451805191935090849081106131b7576131b7614b15565b602002602001015163ffffffff1611156132d95761324b8783815181106131e0576131e0614b15565b602001015160f81c60f81b60f81c8460400151848151811061320457613204614b15565b60200260200101518b8660200151868151811061322357613223614b15565b602002602001015189878151811061323d5761323d614b15565b602002602001015186613798565b6040805160018082528183019092525f9160208201818036833701905050905087838151811061327d5761327d614b15565b602001015160f81c60f81b815f8151811061329a5761329a614b15565b60200101906001600160f81b03191690815f1a9053506132d78684815181106132c5576132c5614b15565b602002602001015160200151826129fd565b505b5060010161312d565b5f61131c613919565b5f81815260208390526040812054808203613309575f915050610cd2565b5f8381526020859052604090206133216001836154fc565b8154811061333157613331614b15565b5f91825260209091200154600160401b90046001600160c01b03169150610cd29050565b5f805b8215610cd2576133696001846154fc565b90921691806133778161550f565b915050613358565b610c466002838361398c565b5f80805f5160206156ea5f395f51905f5260035f5160206156ea5f395f51905f52865f5160206156ea5f395f51905f52888909090890505f6133fb827f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f525f5160206156ea5f395f51905f52613b45565b91959194509092505050565b5f6101008251111561342c57604051637da54e4760e11b815260040160405180910390fd5b81515f0361343b57505f919050565b5f5f835f8151811061344f5761344f614b15565b0160200151600160f89190911c81901b92505b84518110156134b95784818151811061347d5761347d614b15565b0160200151600160f89190911c1b91508282116134ad57604051631019106960e31b815260040160405180910390fd5b91811791600101613462565b50909392505050565b603254610100900460ff166134e95760405162461bcd60e51b81526004016111f490615242565b60ca6134f58382615573565b5060cb6135028282615573565b50505f60c881905560c95550565b5f60ff8216601f811115610cd257604051632cd44ac360e21b815260040160405180910390fd5b604080516060810182526001600160a01b038481168252600b54811660208301527f00000000000000000000000000000000000000000000000000000000000000001691636e3492b59190810161358d85613bbe565b8152506040518263ffffffff1660e01b81526004016135ac919061562d565b5f604051808303815f87803b1580156135c3575f5ffd5b505af11580156112e9573d5f5f3e3d5ffd5b5f81815260208490526040812054815b818110156136585760016135f982846154fc565b61360391906154fc565b92508463ffffffff16865f8681526020019081526020015f208463ffffffff168154811061363357613633614b15565b5f9182526020909120015463ffffffff1611613650575050610aba565b6001016135e5565b5060405162461bcd60e51b815260206004820152605c60248201527f5265676973747279436f6f7264696e61746f722e67657451756f72756d42697460448201527f6d6170496e6465784174426c6f636b4e756d6265723a206e6f206269746d617060648201527f2075706461746520666f756e6420666f72206f70657261746f72496400000000608482015260a4016111f4565b6020808201515f9081526004909152604090205460ff161561372257604051636fbefec360e11b815260040160405180910390fd5b428160400151101561374757604051630819bdcd60e01b815260040160405180910390fd5b602080820180515f9081526004909252604091829020805460ff19166001179055600754905191830151610e31926001600160a01b039092169161379191889188918891906115b2565b8351613c63565b6020808301516001600160a01b038082165f81815260039094526040909320549192908716036137db576040516356168b4160e11b815260040160405180910390fd5b8760ff16845f015160ff161461380457604051638e5aeee760e01b815260040160405180910390fd5b604051635401ed2760e01b81526004810182905260ff891660248201525f907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690635401ed2790604401602060405180830381865afa158015613872573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906138969190615669565b90506138a28185613c8b565b6001600160601b0316866001600160601b0316116138d357604051634c44995d60e01b815260040160405180910390fd5b6138dd8885613cae565b6001600160601b0316816001600160601b03161061390e5760405163b187e86960e01b815260040160405180910390fd5b505050505050505050565b5f7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f613943613cc7565b61394b613d1f565b60408051602081019490945283019190915260608201524660808201523060a082015260c00160405160208183030381529060405280519060200120905090565b5f8281526020849052604081205490819003613a30575f83815260208581526040808320815160608101835263ffffffff43811682528185018681526001600160c01b03808a16958401958652845460018101865594885295909620915191909201805495519351909416600160401b026001600160401b03938316600160201b0267ffffffffffffffff1990961691909216179390931716919091179055610e31565b5f838152602085905260408120613a486001846154fc565b81548110613a5857613a58614b15565b5f918252602090912001805490915063ffffffff438116911603613a995780546001600160401b0316600160401b6001600160c01b03851602178155610df0565b805463ffffffff438116600160201b81810267ffffffff00000000199094169390931784555f8781526020898152604080832081516060810183529485528483018481526001600160c01b03808c1693870193845282546001810184559286529390942094519401805493519151909216600160401b026001600160401b0391861690960267ffffffffffffffff19909316939094169290921717919091169190911790555050505050565b5f5f613b4f613f7f565b613b57613f9d565b602080825281810181905260408201819052606082018890526080820187905260a082018690528260c08360056107d05a03fa92508280613b9457fe5b5082613bb35760405163d51edae360e01b815260040160405180910390fd5b505195945050505050565b60605f82516001600160401b03811115613bda57613bda613fbb565b604051908082528060200260200182016040528015613c03578160200160208202803683370190505b5090505f5b8351811015611f5e57838181518110613c2357613c23614b15565b602001015160f81c60f81b60f81c60ff16828281518110613c4657613c46614b15565b63ffffffff90921660209283029190910190910152600101613c08565b613c6e838383613d4f565b610aa457604051638baa579f60e01b815260040160405180910390fd5b60208101515f9061271090613ca49061ffff1685615684565b610aba91906156a6565b60408101515f9061271090613ca49061ffff1685615684565b5f5f613cd1612abf565b805190915015613ce8578051602090910120919050565b60c8548015613cf75792915050565b7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a4709250505090565b5f5f613d29612b4f565b805190915015613d40578051602090910120919050565b60c9548015613cf75792915050565b5f5f5f613d5c8585613da3565b90925090505f816004811115613d7457613d746146ae565b148015613d925750856001600160a01b0316826001600160a01b0316145b806115f157506115f1868686613de5565b5f5f8251604103613dd7576020830151604084015160608501515f1a613dcb87828585613ecc565b94509450505050613dde565b505f905060025b9250929050565b5f5f5f856001600160a01b0316631626ba7e60e01b8686604051602401613e0d929190614fb3565b60408051601f198184030181529181526020820180516001600160e01b03166001600160e01b0319909416939093179092529051613e4b91906156d3565b5f60405180830381855afa9150503d805f8114613e83576040519150601f19603f3d011682016040523d82523d5f602084013e613e88565b606091505b5091509150818015613e9c57506020815110155b80156115f157508051630b135d3f60e11b90613ec19083016020908101908401615360565b149695505050505050565b5f806fa2a8918ca85bafe22016d0b997e4df60600160ff1b03831115613ef757505f90506003613f76565b604080515f8082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015613f48573d5f5f3e3d5ffd5b5050604051601f1901519150506001600160a01b038116613f70575f60019250925050613f76565b91505f90505b94509492505050565b60405180602001604052806001906020820280368337509192915050565b6040518060c001604052806006906020820280368337509192915050565b634e487b7160e01b5f52604160045260245ffd5b604051606081016001600160401b0381118282101715613ff157613ff1613fbb565b60405290565b604080519081016001600160401b0381118282101715613ff157613ff1613fbb565b604051601f8201601f191681016001600160401b038111828210171561404157614041613fbb565b604052919050565b5f6001600160401b0382111561406157614061613fbb565b5060051b60200190565b6001600160a01b0381168114610b20575f5ffd5b5f82601f83011261408e575f5ffd5b81356140a161409c82614049565b614019565b8082825260208201915060208360051b8601019250858311156140c2575f5ffd5b602085015b838110156140e85780356140da8161406b565b8352602092830192016140c7565b5095945050505050565b5f60208284031215614102575f5ffd5b81356001600160401b03811115614117575f5ffd5b611c588482850161407f565b5f60208284031215614133575f5ffd5b5035919050565b63ffffffff81168114610b20575f5ffd5b5f5f5f6060848603121561415d575f5ffd5b83359250602084013561416f8161413a565b929592945050506040919091013590565b6001600160a01b0391909116815260200190565b5f82601f8301126141a3575f5ffd5b8135602083015f5f6001600160401b038411156141c2576141c2613fbb565b50601f8301601f19166020016141d781614019565b9150508281528583830111156141eb575f5ffd5b828260208301375f92810160200192909252509392505050565b5f60208284031215614215575f5ffd5b81356001600160401b0381111561422a575f5ffd5b611c5884828501614194565b5f60208284031215614246575f5ffd5b8135610aba8161406b565b5f5f60408385031215614262575f5ffd5b50508035926020909101359150565b803560ff81168114614281575f5ffd5b919050565b5f60208284031215614296575f5ffd5b610aba82614271565b5f82601f8301126142ae575f5ffd5b81356142bc61409c82614049565b8082825260208201915060208360051b8601019250858311156142dd575f5ffd5b602085015b838110156140e85780356142f58161413a565b8352602092830192016142e2565b5f5f5f60608486031215614315575f5ffd5b83356143208161406b565b925060208401356143308161406b565b915060408401356001600160401b0381111561434a575f5ffd5b6143568682870161429f565b9150509250925092565b815181526020808301519082015260408101610cd2565b803561ffff81168114614281575f5ffd5b5f60608284031215614398575f5ffd5b6143a0613fcf565b905081356143ad8161413a565b81526143bb60208301614377565b60208201526143cc60408301614377565b604082015292915050565b6001600160601b0381168114610b20575f5ffd5b5f82601f8301126143fa575f5ffd5b813561440861409c82614049565b8082825260208201915060208360061b860101925085831115614429575f5ffd5b602085015b838110156140e85760408188031215614445575f5ffd5b61444d613ff7565b81356144588161406b565b81526020820135614468816143d7565b602082810191909152908452929092019160400161442e565b5f5f5f5f60c08587031215614494575f5ffd5b61449e8686614388565b935060608501356144ae816143d7565b925060808501356001600160401b038111156144c8575f5ffd5b6144d4878288016143eb565b92505060a08501356144e58161413a565b939692955090935050565b5f5f83601f840112614500575f5ffd5b5081356001600160401b03811115614516575f5ffd5b602083019150836020828501011115613dde575f5ffd5b5f5f5f6040848603121561453f575f5ffd5b83356001600160401b03811115614554575f5ffd5b8401601f81018613614564575f5ffd5b803561457261409c82614049565b8082825260208201915060208360051b850101925088831115614593575f5ffd5b602084015b838110156145d35780356001600160401b038111156145b5575f5ffd5b6145c48b60208389010161407f565b84525060209283019201614598565b50955050505060208401356001600160401b038111156145f1575f5ffd5b6145fd868287016144f0565b9497909650939450505050565b5f5f5f5f5f60a0868803121561461e575f5ffd5b85356146298161406b565b945060208601356146398161406b565b935060408601356146498161406b565b92506060860135915060808601356146608161406b565b809150509295509295909350565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b602081525f610aba602083018461466e565b634e487b7160e01b5f52602160045260245ffd5b600381106146d2576146d26146ae565b9052565b815181526020808301516040830191611f5e908401826146c2565b5f5f60808385031215614702575f5ffd5b61470b83614271565b915061471a8460208501614388565b90509250929050565b5f5f60408385031215614734575f5ffd5b823561473f8161406b565b915060208301356001600160401b03811115614759575f5ffd5b61476585828601614194565b9150509250929050565b5f5f5f60a08486031215614781575f5ffd5b61478b8585614388565b9250606084013561479b816143d7565b915060808401356001600160401b038111156147b5575f5ffd5b614356868287016143eb565b60ff60f81b8816815260e060208201525f6147df60e083018961466e565b82810360408401526147f1818961466e565b606084018890526001600160a01b038716608085015260a0840186905283810360c0850152845180825260208087019350909101905f5b81811015614846578351835260209384019390920191600101614828565b50909b9a5050505050505050505050565b5f82601f830112614866575f5ffd5b813561487461409c82614049565b8082825260208201915060208360061b860101925085831115614895575f5ffd5b602085015b838110156140e857604081880312156148b1575f5ffd5b6148b9613ff7565b6148c282614271565b815260208201356148d28161406b565b602082810191909152908452929092019160400161489a565b5f5f5f5f5f60a086880312156148ff575f5ffd5b853561490a8161406b565b94506020860135935060408601356001600160401b0381111561492b575f5ffd5b61493788828901614857565b9598949750949560608101359550608001359392505050565b5f5f60408385031215614961575f5ffd5b823561496c8161413a565b915060208301356001600160401b03811115614986575f5ffd5b8301601f81018513614996575f5ffd5b80356149a461409c82614049565b8082825260208201915060208360051b8501019250878311156149c5575f5ffd5b6020840193505b828410156149e75783358252602093840193909101906149cc565b809450505050509250929050565b5f8151808452602084019350602083015f5b82811015614a2b57815163ffffffff16865260209586019590910190600101614a07565b5093949350505050565b602081525f610aba60208301846149f5565b5f5f5f5f5f60808688031215614a5b575f5ffd5b8535614a668161406b565b94506020860135614a768161406b565b935060408601356001600160401b03811115614a90575f5ffd5b614a9c8882890161429f565b93505060608601356001600160401b03811115614ab7575f5ffd5b614ac3888289016144f0565b969995985093965092949392505050565b60608101610cd2828463ffffffff815116825261ffff602082015116602083015261ffff60408201511660408301525050565b60208101610cd282846146c2565b634e487b7160e01b5f52603260045260245ffd5b80518015158114614281575f5ffd5b5f60208284031215614b48575f5ffd5b610aba82614b29565b5f60208284031215614b61575f5ffd5b8151610aba8161406b565b5f60208284031215614b7c575f5ffd5b8151610aba8161413a565b5f60c0820188835260018060a01b038816602084015286604084015260c0606084015280865180835260e0850191506020880192505f5b81811015614bf4578351805160ff1684526020908101516001600160a01b03168185015290930192604090920191600101614bbe565b50506080840195909552505060a00152949350505050565b803560028110614281575f5ffd5b5f60408284031215614c2a575f5ffd5b614c32613ff7565b823581526020928301359281019290925250919050565b5f82601f830112614c58575f5ffd5b614c60613ff7565b806040840185811115614c71575f5ffd5b845b81811015614c8b578035845260209384019301614c73565b509095945050505050565b5f818303610100811215614ca8575f5ffd5b614cb0613fcf565b9150614cbc8484614c1a565b8252614ccb8460408501614c1a565b60208301526080607f1982011215614ce1575f5ffd5b50614cea613ff7565b614cf78460808501614c49565b8152614d068460c08501614c49565b6020820152604082015292915050565b5f5f5f6101408486031215614d29575f5ffd5b614d3284614c0c565b925060208401356001600160401b03811115614d4c575f5ffd5b614d5886828701614194565b925050614d688560408601614c96565b90509250925092565b5f5f5f5f5f6101808688031215614d86575f5ffd5b614d8f86614c0c565b945060208601356001600160401b03811115614da9575f5ffd5b614db588828901614194565b945050614dc58760408801614c96565b92506101408601356001600160401b03811115614de0575f5ffd5b614dec88828901614857565b9250506101608601356001600160401b03811115614e08575f5ffd5b860160608189031215614e19575f5ffd5b614e21613fcf565b81356001600160401b03811115614e36575f5ffd5b614e428a828501614194565b8252506020828101359082015260409182013591810191909152949793965091945092919050565b634e487b7160e01b5f52601160045260245ffd5b5f60018201614e8f57614e8f614e6a565b5060010190565b606080825284519082018190525f9060208601906080840190835b81811015614ed85783516001600160a01b0316835260209384019390920191600101614eb1565b5050838103602080860191909152865180835291810192508601905f5b81811015614f13578251845260209384019390920191600101614ef5565b50505060ff841660408401529050611c58565b5f60208284031215614f36575f5ffd5b81516001600160401b03811115614f4b575f5ffd5b8201601f81018413614f5b575f5ffd5b8051614f6961409c82614049565b8082825260208201915060208360051b850101925086831115614f8a575f5ffd5b6020840193505b828410156115f157614fa284614b29565b825260209384019390910190614f91565b828152604060208201525f610ab7604083018461466e565b6001600160a01b0392831681529116602082015260400190565b6001600160a01b03831681526040602082018190525f90610ab79083018461466e565b634e487b7160e01b5f52601260045260245ffd5b5f8261502a5761502a615008565b500690565b60ff8181168382160190811115610cd257610cd2614e6a565b5f6040820160018060a01b03851683526040602084015280845180835260608501915060608160051b8601019250602086015f5b828110156150fd57868503605f190184528151805163ffffffff168652602090810151604082880181905281519088018190529101905f9060608801905b808310156150e55783516001600160a01b0316825260209384019360019390930192909101906150ba565b5096505050602093840193919091019060010161507c565b5092979650505050505050565b5f8151808452602084019350602083015f5b82811015614a2b57815180516001600160a01b031687526020908101516001600160601b0316818801526040909601959091019060010161511c565b60ff841681526001600160601b0383166020820152606060408201525f615182606083018461510a565b95945050505050565b60ff851681526001600160601b038416602082015263ffffffff83166040820152608060608201525f6115f1608083018461510a565b6151f0818763ffffffff815116825261ffff602082015116602083015261ffff60408201511660408301525050565b6001600160601b038516606082015260e060808201525f61521460e083018661510a565b905060028410615226576152266146ae565b8360a083015263ffffffff831660c08301529695505050505050565b6020808252602b908201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960408201526a6e697469616c697a696e6760a81b606082015260800190565b600181811c908216806152a157607f821691505b6020821081036152bf57634e487b7160e01b5f52602260045260245ffd5b50919050565b805f5b6002811015610e315781518452602093840193909101906001016152c8565b6001600160a01b0384168152825180516020808401919091520151604082015261016081016020848101518051606085015290810151608084015250604084015161533660a0840182516152c5565b6020015161534760e08401826152c5565b5082516101208301526020830151610140830152611c58565b5f60208284031215615370575f5ffd5b5051919050565b80820180821115610cd257610cd2614e6a565b60018060a01b0384168152826020820152606060408201525f615182606083018461466e565b5f82601f8301126153bf575f5ffd5b81516153cd61409c82614049565b8082825260208201915060208360051b8601019250858311156153ee575f5ffd5b602085015b838110156140e8578051615406816143d7565b8352602092830192016153f3565b5f5f60408385031215615425575f5ffd5b82516001600160401b0381111561543a575f5ffd5b615446858286016153b0565b92505060208301516001600160401b03811115615461575f5ffd5b614765858286016153b0565b5f6020828403121561547d575f5ffd5b81516001600160401b03811115615492575f5ffd5b8201601f810184136154a2575f5ffd5b80516154b061409c82614049565b8082825260208201915060208360051b8501019250868311156154d1575f5ffd5b6020840193505b828410156115f15783516154eb8161413a565b8252602093840193909101906154d8565b81810381811115610cd257610cd2614e6a565b5f61ffff821661ffff810361552657615526614e6a565b60010192915050565b601f821115610aa457805f5260205f20601f840160051c810160208510156155545750805b601f840160051c820191505b81811015610df0575f8155600101615560565b81516001600160401b0381111561558c5761558c613fbb565b6155a08161559a845461528d565b8461552f565b6020601f8211600181146155d2575f83156155bb5750848201515b5f19600385901b1c1916600184901b178455610df0565b5f84815260208120601f198516915b8281101561560157878501518255602094850194600190920191016155e1565b508482101561561e57868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b602080825282516001600160a01b039081168383015290830151166040808301919091528201516060808301525f90611c5860808401826149f5565b5f60208284031215615679575f5ffd5b8151610aba816143d7565b6001600160601b038181168382160290811690818114611f5e57611f5e614e6a565b5f6001600160601b038316806156be576156be615008565b806001600160601b0384160491505092915050565b5f82518060208501845e5f92019182525091905056fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47a2646970667358221220e111f031b43faf621631c53278218777b5efa0519fc6cf85e62d243ecd3b2a4e64736f6c634300081b0033",
}

// ContractSlashingRegistryCoordinatorABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractSlashingRegistryCoordinatorMetaData.ABI instead.
var ContractSlashingRegistryCoordinatorABI = ContractSlashingRegistryCoordinatorMetaData.ABI

// ContractSlashingRegistryCoordinatorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractSlashingRegistryCoordinatorMetaData.Bin instead.
var ContractSlashingRegistryCoordinatorBin = ContractSlashingRegistryCoordinatorMetaData.Bin

// DeployContractSlashingRegistryCoordinator deploys a new Ethereum contract, binding an instance of ContractSlashingRegistryCoordinator to it.
func DeployContractSlashingRegistryCoordinator(auth *bind.TransactOpts, backend bind.ContractBackend, _stakeRegistry common.Address, _blsApkRegistry common.Address, _indexRegistry common.Address, _socketRegistry common.Address, _allocationManager common.Address, _pauserRegistry common.Address, _version string) (common.Address, *types.Transaction, *ContractSlashingRegistryCoordinator, error) {
	parsed, err := ContractSlashingRegistryCoordinatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractSlashingRegistryCoordinatorBin), backend, _stakeRegistry, _blsApkRegistry, _indexRegistry, _socketRegistry, _allocationManager, _pauserRegistry, _version)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractSlashingRegistryCoordinator{ContractSlashingRegistryCoordinatorCaller: ContractSlashingRegistryCoordinatorCaller{contract: contract}, ContractSlashingRegistryCoordinatorTransactor: ContractSlashingRegistryCoordinatorTransactor{contract: contract}, ContractSlashingRegistryCoordinatorFilterer: ContractSlashingRegistryCoordinatorFilterer{contract: contract}}, nil
}

// ContractSlashingRegistryCoordinatorMethods is an auto generated interface around an Ethereum contract.
type ContractSlashingRegistryCoordinatorMethods interface {
	ContractSlashingRegistryCoordinatorCalls
	ContractSlashingRegistryCoordinatorTransacts
	ContractSlashingRegistryCoordinatorFilters
}

// ContractSlashingRegistryCoordinatorCalls is an auto generated interface that defines the call methods available for an Ethereum contract.
type ContractSlashingRegistryCoordinatorCalls interface {
	OPERATORCHURNAPPROVALTYPEHASH(opts *bind.CallOpts) ([32]byte, error)

	PUBKEYREGISTRATIONTYPEHASH(opts *bind.CallOpts) ([32]byte, error)

	AllocationManager(opts *bind.CallOpts) (common.Address, error)

	Avs(opts *bind.CallOpts) (common.Address, error)

	BlsApkRegistry(opts *bind.CallOpts) (common.Address, error)

	CalculateOperatorChurnApprovalDigestHash(opts *bind.CallOpts, registeringOperator common.Address, registeringOperatorId [32]byte, operatorKickParams []ISlashingRegistryCoordinatorTypesOperatorKickParam, salt [32]byte, expiry *big.Int) ([32]byte, error)

	CalculatePubkeyRegistrationMessageHash(opts *bind.CallOpts, operator common.Address) ([32]byte, error)

	ChurnApprover(opts *bind.CallOpts) (common.Address, error)

	DomainSeparator(opts *bind.CallOpts) ([32]byte, error)

	Eip712Domain(opts *bind.CallOpts) (struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	}, error)

	EjectionCooldown(opts *bind.CallOpts) (*big.Int, error)

	Ejector(opts *bind.CallOpts) (common.Address, error)

	GetCurrentQuorumBitmap(opts *bind.CallOpts, operatorId [32]byte) (*big.Int, error)

	GetOperator(opts *bind.CallOpts, operator common.Address) (ISlashingRegistryCoordinatorTypesOperatorInfo, error)

	GetOperatorFromId(opts *bind.CallOpts, operatorId [32]byte) (common.Address, error)

	GetOperatorId(opts *bind.CallOpts, operator common.Address) ([32]byte, error)

	GetOperatorSetParams(opts *bind.CallOpts, quorumNumber uint8) (ISlashingRegistryCoordinatorTypesOperatorSetParam, error)

	GetOperatorStatus(opts *bind.CallOpts, operator common.Address) (uint8, error)

	GetQuorumBitmapAtBlockNumberByIndex(opts *bind.CallOpts, operatorId [32]byte, blockNumber uint32, index *big.Int) (*big.Int, error)

	GetQuorumBitmapHistoryLength(opts *bind.CallOpts, operatorId [32]byte) (*big.Int, error)

	GetQuorumBitmapIndicesAtBlockNumber(opts *bind.CallOpts, blockNumber uint32, operatorIds [][32]byte) ([]uint32, error)

	GetQuorumBitmapUpdateByIndex(opts *bind.CallOpts, operatorId [32]byte, index *big.Int) (ISlashingRegistryCoordinatorTypesQuorumBitmapUpdate, error)

	IndexRegistry(opts *bind.CallOpts) (common.Address, error)

	IsChurnApproverSaltUsed(opts *bind.CallOpts, arg0 [32]byte) (bool, error)

	LastEjectionTimestamp(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	Paused(opts *bind.CallOpts, index uint8) (bool, error)

	Paused0(opts *bind.CallOpts) (*big.Int, error)

	PauserRegistry(opts *bind.CallOpts) (common.Address, error)

	PubkeyRegistrationMessageHash(opts *bind.CallOpts, operator common.Address) (BN254G1Point, error)

	QuorumCount(opts *bind.CallOpts) (uint8, error)

	QuorumUpdateBlockNumber(opts *bind.CallOpts, arg0 uint8) (*big.Int, error)

	SocketRegistry(opts *bind.CallOpts) (common.Address, error)

	StakeRegistry(opts *bind.CallOpts) (common.Address, error)

	SupportsAVS(opts *bind.CallOpts, _avs common.Address) (bool, error)

	Version(opts *bind.CallOpts) (string, error)
}

// ContractSlashingRegistryCoordinatorTransacts is an auto generated interface that defines the transact methods available for an Ethereum contract.
type ContractSlashingRegistryCoordinatorTransacts interface {
	CreateSlashableStakeQuorum(opts *bind.TransactOpts, operatorSetParams ISlashingRegistryCoordinatorTypesOperatorSetParam, minimumStake *big.Int, strategyParams []IStakeRegistryTypesStrategyParams, lookAheadPeriod uint32) (*types.Transaction, error)

	CreateTotalDelegatedStakeQuorum(opts *bind.TransactOpts, operatorSetParams ISlashingRegistryCoordinatorTypesOperatorSetParam, minimumStake *big.Int, strategyParams []IStakeRegistryTypesStrategyParams) (*types.Transaction, error)

	DeregisterOperator(opts *bind.TransactOpts, operator common.Address, avs common.Address, operatorSetIds []uint32) (*types.Transaction, error)

	EjectOperator(opts *bind.TransactOpts, operator common.Address, quorumNumbers []byte) (*types.Transaction, error)

	Initialize(opts *bind.TransactOpts, initialOwner common.Address, churnApprover common.Address, ejector common.Address, initialPausedStatus *big.Int, avs common.Address) (*types.Transaction, error)

	Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error)

	PauseAll(opts *bind.TransactOpts) (*types.Transaction, error)

	RegisterOperator(opts *bind.TransactOpts, operator common.Address, avs common.Address, operatorSetIds []uint32, data []byte) (*types.Transaction, error)

	RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	SetAVS(opts *bind.TransactOpts, _avs common.Address) (*types.Transaction, error)

	SetChurnApprover(opts *bind.TransactOpts, _churnApprover common.Address) (*types.Transaction, error)

	SetEjectionCooldown(opts *bind.TransactOpts, _ejectionCooldown *big.Int) (*types.Transaction, error)

	SetEjector(opts *bind.TransactOpts, _ejector common.Address) (*types.Transaction, error)

	SetOperatorSetParams(opts *bind.TransactOpts, quorumNumber uint8, operatorSetParams ISlashingRegistryCoordinatorTypesOperatorSetParam) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error)

	Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error)

	UpdateOperators(opts *bind.TransactOpts, operators []common.Address) (*types.Transaction, error)

	UpdateOperatorsForQuorum(opts *bind.TransactOpts, operatorsPerQuorum [][]common.Address, quorumNumbers []byte) (*types.Transaction, error)

	UpdateSocket(opts *bind.TransactOpts, socket string) (*types.Transaction, error)
}

// ContractSlashingRegistryCoordinatorFilterer is an auto generated interface that defines the log filtering methods available for an Ethereum contract.
type ContractSlashingRegistryCoordinatorFilters interface {
	FilterAVSUpdated(opts *bind.FilterOpts) (*ContractSlashingRegistryCoordinatorAVSUpdatedIterator, error)
	WatchAVSUpdated(opts *bind.WatchOpts, sink chan<- *ContractSlashingRegistryCoordinatorAVSUpdated) (event.Subscription, error)
	ParseAVSUpdated(log types.Log) (*ContractSlashingRegistryCoordinatorAVSUpdated, error)

	FilterChurnApproverUpdated(opts *bind.FilterOpts) (*ContractSlashingRegistryCoordinatorChurnApproverUpdatedIterator, error)
	WatchChurnApproverUpdated(opts *bind.WatchOpts, sink chan<- *ContractSlashingRegistryCoordinatorChurnApproverUpdated) (event.Subscription, error)
	ParseChurnApproverUpdated(log types.Log) (*ContractSlashingRegistryCoordinatorChurnApproverUpdated, error)

	FilterEIP712DomainChanged(opts *bind.FilterOpts) (*ContractSlashingRegistryCoordinatorEIP712DomainChangedIterator, error)
	WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *ContractSlashingRegistryCoordinatorEIP712DomainChanged) (event.Subscription, error)
	ParseEIP712DomainChanged(log types.Log) (*ContractSlashingRegistryCoordinatorEIP712DomainChanged, error)

	FilterEjectionCooldownUpdated(opts *bind.FilterOpts) (*ContractSlashingRegistryCoordinatorEjectionCooldownUpdatedIterator, error)
	WatchEjectionCooldownUpdated(opts *bind.WatchOpts, sink chan<- *ContractSlashingRegistryCoordinatorEjectionCooldownUpdated) (event.Subscription, error)
	ParseEjectionCooldownUpdated(log types.Log) (*ContractSlashingRegistryCoordinatorEjectionCooldownUpdated, error)

	FilterEjectorUpdated(opts *bind.FilterOpts) (*ContractSlashingRegistryCoordinatorEjectorUpdatedIterator, error)
	WatchEjectorUpdated(opts *bind.WatchOpts, sink chan<- *ContractSlashingRegistryCoordinatorEjectorUpdated) (event.Subscription, error)
	ParseEjectorUpdated(log types.Log) (*ContractSlashingRegistryCoordinatorEjectorUpdated, error)

	FilterInitialized(opts *bind.FilterOpts) (*ContractSlashingRegistryCoordinatorInitializedIterator, error)
	WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractSlashingRegistryCoordinatorInitialized) (event.Subscription, error)
	ParseInitialized(log types.Log) (*ContractSlashingRegistryCoordinatorInitialized, error)

	FilterOperatorDeregistered(opts *bind.FilterOpts, operator []common.Address, operatorId [][32]byte) (*ContractSlashingRegistryCoordinatorOperatorDeregisteredIterator, error)
	WatchOperatorDeregistered(opts *bind.WatchOpts, sink chan<- *ContractSlashingRegistryCoordinatorOperatorDeregistered, operator []common.Address, operatorId [][32]byte) (event.Subscription, error)
	ParseOperatorDeregistered(log types.Log) (*ContractSlashingRegistryCoordinatorOperatorDeregistered, error)

	FilterOperatorRegistered(opts *bind.FilterOpts, operator []common.Address, operatorId [][32]byte) (*ContractSlashingRegistryCoordinatorOperatorRegisteredIterator, error)
	WatchOperatorRegistered(opts *bind.WatchOpts, sink chan<- *ContractSlashingRegistryCoordinatorOperatorRegistered, operator []common.Address, operatorId [][32]byte) (event.Subscription, error)
	ParseOperatorRegistered(log types.Log) (*ContractSlashingRegistryCoordinatorOperatorRegistered, error)

	FilterOperatorSetParamsUpdated(opts *bind.FilterOpts, quorumNumber []uint8) (*ContractSlashingRegistryCoordinatorOperatorSetParamsUpdatedIterator, error)
	WatchOperatorSetParamsUpdated(opts *bind.WatchOpts, sink chan<- *ContractSlashingRegistryCoordinatorOperatorSetParamsUpdated, quorumNumber []uint8) (event.Subscription, error)
	ParseOperatorSetParamsUpdated(log types.Log) (*ContractSlashingRegistryCoordinatorOperatorSetParamsUpdated, error)

	FilterOperatorSocketUpdate(opts *bind.FilterOpts, operatorId [][32]byte) (*ContractSlashingRegistryCoordinatorOperatorSocketUpdateIterator, error)
	WatchOperatorSocketUpdate(opts *bind.WatchOpts, sink chan<- *ContractSlashingRegistryCoordinatorOperatorSocketUpdate, operatorId [][32]byte) (event.Subscription, error)
	ParseOperatorSocketUpdate(log types.Log) (*ContractSlashingRegistryCoordinatorOperatorSocketUpdate, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractSlashingRegistryCoordinatorOwnershipTransferredIterator, error)
	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractSlashingRegistryCoordinatorOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error)
	ParseOwnershipTransferred(log types.Log) (*ContractSlashingRegistryCoordinatorOwnershipTransferred, error)

	FilterPaused(opts *bind.FilterOpts, account []common.Address) (*ContractSlashingRegistryCoordinatorPausedIterator, error)
	WatchPaused(opts *bind.WatchOpts, sink chan<- *ContractSlashingRegistryCoordinatorPaused, account []common.Address) (event.Subscription, error)
	ParsePaused(log types.Log) (*ContractSlashingRegistryCoordinatorPaused, error)

	FilterQuorumBlockNumberUpdated(opts *bind.FilterOpts, quorumNumber []uint8) (*ContractSlashingRegistryCoordinatorQuorumBlockNumberUpdatedIterator, error)
	WatchQuorumBlockNumberUpdated(opts *bind.WatchOpts, sink chan<- *ContractSlashingRegistryCoordinatorQuorumBlockNumberUpdated, quorumNumber []uint8) (event.Subscription, error)
	ParseQuorumBlockNumberUpdated(log types.Log) (*ContractSlashingRegistryCoordinatorQuorumBlockNumberUpdated, error)

	FilterQuorumCreated(opts *bind.FilterOpts, quorumNumber []uint8) (*ContractSlashingRegistryCoordinatorQuorumCreatedIterator, error)
	WatchQuorumCreated(opts *bind.WatchOpts, sink chan<- *ContractSlashingRegistryCoordinatorQuorumCreated, quorumNumber []uint8) (event.Subscription, error)
	ParseQuorumCreated(log types.Log) (*ContractSlashingRegistryCoordinatorQuorumCreated, error)

	FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*ContractSlashingRegistryCoordinatorUnpausedIterator, error)
	WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ContractSlashingRegistryCoordinatorUnpaused, account []common.Address) (event.Subscription, error)
	ParseUnpaused(log types.Log) (*ContractSlashingRegistryCoordinatorUnpaused, error)
}

// ContractSlashingRegistryCoordinator is an auto generated Go binding around an Ethereum contract.
type ContractSlashingRegistryCoordinator struct {
	ContractSlashingRegistryCoordinatorCaller     // Read-only binding to the contract
	ContractSlashingRegistryCoordinatorTransactor // Write-only binding to the contract
	ContractSlashingRegistryCoordinatorFilterer   // Log filterer for contract events
}

// ContractSlashingRegistryCoordinator implements the ContractSlashingRegistryCoordinatorMethods interface.
var _ ContractSlashingRegistryCoordinatorMethods = (*ContractSlashingRegistryCoordinator)(nil)

// ContractSlashingRegistryCoordinatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractSlashingRegistryCoordinatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSlashingRegistryCoordinatorCaller implements the ContractSlashingRegistryCoordinatorCalls interface.
var _ ContractSlashingRegistryCoordinatorCalls = (*ContractSlashingRegistryCoordinatorCaller)(nil)

// ContractSlashingRegistryCoordinatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractSlashingRegistryCoordinatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSlashingRegistryCoordinatorTransactor implements the ContractSlashingRegistryCoordinatorTransacts interface.
var _ ContractSlashingRegistryCoordinatorTransacts = (*ContractSlashingRegistryCoordinatorTransactor)(nil)

// ContractSlashingRegistryCoordinatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractSlashingRegistryCoordinatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSlashingRegistryCoordinatorFilterer implements the ContractSlashingRegistryCoordinatorFilters interface.
var _ ContractSlashingRegistryCoordinatorFilters = (*ContractSlashingRegistryCoordinatorFilterer)(nil)

// ContractSlashingRegistryCoordinatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSlashingRegistryCoordinatorSession struct {
	Contract     *ContractSlashingRegistryCoordinator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                        // Call options to use throughout this session
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// ContractSlashingRegistryCoordinatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractSlashingRegistryCoordinatorCallerSession struct {
	Contract *ContractSlashingRegistryCoordinatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                              // Call options to use throughout this session
}

// ContractSlashingRegistryCoordinatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractSlashingRegistryCoordinatorTransactorSession struct {
	Contract     *ContractSlashingRegistryCoordinatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                              // Transaction auth options to use throughout this session
}

// ContractSlashingRegistryCoordinatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractSlashingRegistryCoordinatorRaw struct {
	Contract *ContractSlashingRegistryCoordinator // Generic contract binding to access the raw methods on
}

// ContractSlashingRegistryCoordinatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractSlashingRegistryCoordinatorCallerRaw struct {
	Contract *ContractSlashingRegistryCoordinatorCaller // Generic read-only contract binding to access the raw methods on
}

// ContractSlashingRegistryCoordinatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractSlashingRegistryCoordinatorTransactorRaw struct {
	Contract *ContractSlashingRegistryCoordinatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractSlashingRegistryCoordinator creates a new instance of ContractSlashingRegistryCoordinator, bound to a specific deployed contract.
func NewContractSlashingRegistryCoordinator(address common.Address, backend bind.ContractBackend) (*ContractSlashingRegistryCoordinator, error) {
	contract, err := bindContractSlashingRegistryCoordinator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractSlashingRegistryCoordinator{ContractSlashingRegistryCoordinatorCaller: ContractSlashingRegistryCoordinatorCaller{contract: contract}, ContractSlashingRegistryCoordinatorTransactor: ContractSlashingRegistryCoordinatorTransactor{contract: contract}, ContractSlashingRegistryCoordinatorFilterer: ContractSlashingRegistryCoordinatorFilterer{contract: contract}}, nil
}

// NewContractSlashingRegistryCoordinatorCaller creates a new read-only instance of ContractSlashingRegistryCoordinator, bound to a specific deployed contract.
func NewContractSlashingRegistryCoordinatorCaller(address common.Address, caller bind.ContractCaller) (*ContractSlashingRegistryCoordinatorCaller, error) {
	contract, err := bindContractSlashingRegistryCoordinator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractSlashingRegistryCoordinatorCaller{contract: contract}, nil
}

// NewContractSlashingRegistryCoordinatorTransactor creates a new write-only instance of ContractSlashingRegistryCoordinator, bound to a specific deployed contract.
func NewContractSlashingRegistryCoordinatorTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractSlashingRegistryCoordinatorTransactor, error) {
	contract, err := bindContractSlashingRegistryCoordinator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractSlashingRegistryCoordinatorTransactor{contract: contract}, nil
}

// NewContractSlashingRegistryCoordinatorFilterer creates a new log filterer instance of ContractSlashingRegistryCoordinator, bound to a specific deployed contract.
func NewContractSlashingRegistryCoordinatorFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractSlashingRegistryCoordinatorFilterer, error) {
	contract, err := bindContractSlashingRegistryCoordinator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractSlashingRegistryCoordinatorFilterer{contract: contract}, nil
}

// bindContractSlashingRegistryCoordinator binds a generic wrapper to an already deployed contract.
func bindContractSlashingRegistryCoordinator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractSlashingRegistryCoordinatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractSlashingRegistryCoordinator.Contract.ContractSlashingRegistryCoordinatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.Contract.ContractSlashingRegistryCoordinatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.Contract.ContractSlashingRegistryCoordinatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractSlashingRegistryCoordinator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.Contract.contract.Transact(opts, method, params...)
}

// OPERATORCHURNAPPROVALTYPEHASH is a free data retrieval call binding the contract method 0xca0de882.
//
// Solidity: function OPERATOR_CHURN_APPROVAL_TYPEHASH() view returns(bytes32)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCaller) OPERATORCHURNAPPROVALTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ContractSlashingRegistryCoordinator.contract.Call(opts, &out, "OPERATOR_CHURN_APPROVAL_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OPERATORCHURNAPPROVALTYPEHASH is a free data retrieval call binding the contract method 0xca0de882.
//
// Solidity: function OPERATOR_CHURN_APPROVAL_TYPEHASH() view returns(bytes32)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) OPERATORCHURNAPPROVALTYPEHASH() ([32]byte, error) {
	return _ContractSlashingRegistryCoordinator.Contract.OPERATORCHURNAPPROVALTYPEHASH(&_ContractSlashingRegistryCoordinator.CallOpts)
}

// OPERATORCHURNAPPROVALTYPEHASH is a free data retrieval call binding the contract method 0xca0de882.
//
// Solidity: function OPERATOR_CHURN_APPROVAL_TYPEHASH() view returns(bytes32)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCallerSession) OPERATORCHURNAPPROVALTYPEHASH() ([32]byte, error) {
	return _ContractSlashingRegistryCoordinator.Contract.OPERATORCHURNAPPROVALTYPEHASH(&_ContractSlashingRegistryCoordinator.CallOpts)
}

// PUBKEYREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0x9feab859.
//
// Solidity: function PUBKEY_REGISTRATION_TYPEHASH() view returns(bytes32)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCaller) PUBKEYREGISTRATIONTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ContractSlashingRegistryCoordinator.contract.Call(opts, &out, "PUBKEY_REGISTRATION_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PUBKEYREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0x9feab859.
//
// Solidity: function PUBKEY_REGISTRATION_TYPEHASH() view returns(bytes32)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) PUBKEYREGISTRATIONTYPEHASH() ([32]byte, error) {
	return _ContractSlashingRegistryCoordinator.Contract.PUBKEYREGISTRATIONTYPEHASH(&_ContractSlashingRegistryCoordinator.CallOpts)
}

// PUBKEYREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0x9feab859.
//
// Solidity: function PUBKEY_REGISTRATION_TYPEHASH() view returns(bytes32)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCallerSession) PUBKEYREGISTRATIONTYPEHASH() ([32]byte, error) {
	return _ContractSlashingRegistryCoordinator.Contract.PUBKEYREGISTRATIONTYPEHASH(&_ContractSlashingRegistryCoordinator.CallOpts)
}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCaller) AllocationManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractSlashingRegistryCoordinator.contract.Call(opts, &out, "allocationManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) AllocationManager() (common.Address, error) {
	return _ContractSlashingRegistryCoordinator.Contract.AllocationManager(&_ContractSlashingRegistryCoordinator.CallOpts)
}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCallerSession) AllocationManager() (common.Address, error) {
	return _ContractSlashingRegistryCoordinator.Contract.AllocationManager(&_ContractSlashingRegistryCoordinator.CallOpts)
}

// Avs is a free data retrieval call binding the contract method 0xde1164bb.
//
// Solidity: function avs() view returns(address)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCaller) Avs(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractSlashingRegistryCoordinator.contract.Call(opts, &out, "avs")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Avs is a free data retrieval call binding the contract method 0xde1164bb.
//
// Solidity: function avs() view returns(address)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) Avs() (common.Address, error) {
	return _ContractSlashingRegistryCoordinator.Contract.Avs(&_ContractSlashingRegistryCoordinator.CallOpts)
}

// Avs is a free data retrieval call binding the contract method 0xde1164bb.
//
// Solidity: function avs() view returns(address)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCallerSession) Avs() (common.Address, error) {
	return _ContractSlashingRegistryCoordinator.Contract.Avs(&_ContractSlashingRegistryCoordinator.CallOpts)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCaller) BlsApkRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractSlashingRegistryCoordinator.contract.Call(opts, &out, "blsApkRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) BlsApkRegistry() (common.Address, error) {
	return _ContractSlashingRegistryCoordinator.Contract.BlsApkRegistry(&_ContractSlashingRegistryCoordinator.CallOpts)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCallerSession) BlsApkRegistry() (common.Address, error) {
	return _ContractSlashingRegistryCoordinator.Contract.BlsApkRegistry(&_ContractSlashingRegistryCoordinator.CallOpts)
}

// CalculateOperatorChurnApprovalDigestHash is a free data retrieval call binding the contract method 0x84ca5213.
//
// Solidity: function calculateOperatorChurnApprovalDigestHash(address registeringOperator, bytes32 registeringOperatorId, (uint8,address)[] operatorKickParams, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCaller) CalculateOperatorChurnApprovalDigestHash(opts *bind.CallOpts, registeringOperator common.Address, registeringOperatorId [32]byte, operatorKickParams []ISlashingRegistryCoordinatorTypesOperatorKickParam, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _ContractSlashingRegistryCoordinator.contract.Call(opts, &out, "calculateOperatorChurnApprovalDigestHash", registeringOperator, registeringOperatorId, operatorKickParams, salt, expiry)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateOperatorChurnApprovalDigestHash is a free data retrieval call binding the contract method 0x84ca5213.
//
// Solidity: function calculateOperatorChurnApprovalDigestHash(address registeringOperator, bytes32 registeringOperatorId, (uint8,address)[] operatorKickParams, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) CalculateOperatorChurnApprovalDigestHash(registeringOperator common.Address, registeringOperatorId [32]byte, operatorKickParams []ISlashingRegistryCoordinatorTypesOperatorKickParam, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _ContractSlashingRegistryCoordinator.Contract.CalculateOperatorChurnApprovalDigestHash(&_ContractSlashingRegistryCoordinator.CallOpts, registeringOperator, registeringOperatorId, operatorKickParams, salt, expiry)
}

// CalculateOperatorChurnApprovalDigestHash is a free data retrieval call binding the contract method 0x84ca5213.
//
// Solidity: function calculateOperatorChurnApprovalDigestHash(address registeringOperator, bytes32 registeringOperatorId, (uint8,address)[] operatorKickParams, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCallerSession) CalculateOperatorChurnApprovalDigestHash(registeringOperator common.Address, registeringOperatorId [32]byte, operatorKickParams []ISlashingRegistryCoordinatorTypesOperatorKickParam, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _ContractSlashingRegistryCoordinator.Contract.CalculateOperatorChurnApprovalDigestHash(&_ContractSlashingRegistryCoordinator.CallOpts, registeringOperator, registeringOperatorId, operatorKickParams, salt, expiry)
}

// CalculatePubkeyRegistrationMessageHash is a free data retrieval call binding the contract method 0x73447992.
//
// Solidity: function calculatePubkeyRegistrationMessageHash(address operator) view returns(bytes32)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCaller) CalculatePubkeyRegistrationMessageHash(opts *bind.CallOpts, operator common.Address) ([32]byte, error) {
	var out []interface{}
	err := _ContractSlashingRegistryCoordinator.contract.Call(opts, &out, "calculatePubkeyRegistrationMessageHash", operator)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculatePubkeyRegistrationMessageHash is a free data retrieval call binding the contract method 0x73447992.
//
// Solidity: function calculatePubkeyRegistrationMessageHash(address operator) view returns(bytes32)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) CalculatePubkeyRegistrationMessageHash(operator common.Address) ([32]byte, error) {
	return _ContractSlashingRegistryCoordinator.Contract.CalculatePubkeyRegistrationMessageHash(&_ContractSlashingRegistryCoordinator.CallOpts, operator)
}

// CalculatePubkeyRegistrationMessageHash is a free data retrieval call binding the contract method 0x73447992.
//
// Solidity: function calculatePubkeyRegistrationMessageHash(address operator) view returns(bytes32)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCallerSession) CalculatePubkeyRegistrationMessageHash(operator common.Address) ([32]byte, error) {
	return _ContractSlashingRegistryCoordinator.Contract.CalculatePubkeyRegistrationMessageHash(&_ContractSlashingRegistryCoordinator.CallOpts, operator)
}

// ChurnApprover is a free data retrieval call binding the contract method 0x054310e6.
//
// Solidity: function churnApprover() view returns(address)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCaller) ChurnApprover(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractSlashingRegistryCoordinator.contract.Call(opts, &out, "churnApprover")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ChurnApprover is a free data retrieval call binding the contract method 0x054310e6.
//
// Solidity: function churnApprover() view returns(address)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) ChurnApprover() (common.Address, error) {
	return _ContractSlashingRegistryCoordinator.Contract.ChurnApprover(&_ContractSlashingRegistryCoordinator.CallOpts)
}

// ChurnApprover is a free data retrieval call binding the contract method 0x054310e6.
//
// Solidity: function churnApprover() view returns(address)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCallerSession) ChurnApprover() (common.Address, error) {
	return _ContractSlashingRegistryCoordinator.Contract.ChurnApprover(&_ContractSlashingRegistryCoordinator.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCaller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ContractSlashingRegistryCoordinator.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) DomainSeparator() ([32]byte, error) {
	return _ContractSlashingRegistryCoordinator.Contract.DomainSeparator(&_ContractSlashingRegistryCoordinator.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCallerSession) DomainSeparator() ([32]byte, error) {
	return _ContractSlashingRegistryCoordinator.Contract.DomainSeparator(&_ContractSlashingRegistryCoordinator.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _ContractSlashingRegistryCoordinator.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _ContractSlashingRegistryCoordinator.Contract.Eip712Domain(&_ContractSlashingRegistryCoordinator.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _ContractSlashingRegistryCoordinator.Contract.Eip712Domain(&_ContractSlashingRegistryCoordinator.CallOpts)
}

// EjectionCooldown is a free data retrieval call binding the contract method 0xa96f783e.
//
// Solidity: function ejectionCooldown() view returns(uint256)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCaller) EjectionCooldown(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractSlashingRegistryCoordinator.contract.Call(opts, &out, "ejectionCooldown")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EjectionCooldown is a free data retrieval call binding the contract method 0xa96f783e.
//
// Solidity: function ejectionCooldown() view returns(uint256)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) EjectionCooldown() (*big.Int, error) {
	return _ContractSlashingRegistryCoordinator.Contract.EjectionCooldown(&_ContractSlashingRegistryCoordinator.CallOpts)
}

// EjectionCooldown is a free data retrieval call binding the contract method 0xa96f783e.
//
// Solidity: function ejectionCooldown() view returns(uint256)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCallerSession) EjectionCooldown() (*big.Int, error) {
	return _ContractSlashingRegistryCoordinator.Contract.EjectionCooldown(&_ContractSlashingRegistryCoordinator.CallOpts)
}

// Ejector is a free data retrieval call binding the contract method 0x28f61b31.
//
// Solidity: function ejector() view returns(address)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCaller) Ejector(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractSlashingRegistryCoordinator.contract.Call(opts, &out, "ejector")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Ejector is a free data retrieval call binding the contract method 0x28f61b31.
//
// Solidity: function ejector() view returns(address)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) Ejector() (common.Address, error) {
	return _ContractSlashingRegistryCoordinator.Contract.Ejector(&_ContractSlashingRegistryCoordinator.CallOpts)
}

// Ejector is a free data retrieval call binding the contract method 0x28f61b31.
//
// Solidity: function ejector() view returns(address)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCallerSession) Ejector() (common.Address, error) {
	return _ContractSlashingRegistryCoordinator.Contract.Ejector(&_ContractSlashingRegistryCoordinator.CallOpts)
}

// GetCurrentQuorumBitmap is a free data retrieval call binding the contract method 0x871ef049.
//
// Solidity: function getCurrentQuorumBitmap(bytes32 operatorId) view returns(uint192)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCaller) GetCurrentQuorumBitmap(opts *bind.CallOpts, operatorId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ContractSlashingRegistryCoordinator.contract.Call(opts, &out, "getCurrentQuorumBitmap", operatorId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentQuorumBitmap is a free data retrieval call binding the contract method 0x871ef049.
//
// Solidity: function getCurrentQuorumBitmap(bytes32 operatorId) view returns(uint192)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) GetCurrentQuorumBitmap(operatorId [32]byte) (*big.Int, error) {
	return _ContractSlashingRegistryCoordinator.Contract.GetCurrentQuorumBitmap(&_ContractSlashingRegistryCoordinator.CallOpts, operatorId)
}

// GetCurrentQuorumBitmap is a free data retrieval call binding the contract method 0x871ef049.
//
// Solidity: function getCurrentQuorumBitmap(bytes32 operatorId) view returns(uint192)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCallerSession) GetCurrentQuorumBitmap(operatorId [32]byte) (*big.Int, error) {
	return _ContractSlashingRegistryCoordinator.Contract.GetCurrentQuorumBitmap(&_ContractSlashingRegistryCoordinator.CallOpts, operatorId)
}

// GetOperator is a free data retrieval call binding the contract method 0x5865c60c.
//
// Solidity: function getOperator(address operator) view returns((bytes32,uint8))
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCaller) GetOperator(opts *bind.CallOpts, operator common.Address) (ISlashingRegistryCoordinatorTypesOperatorInfo, error) {
	var out []interface{}
	err := _ContractSlashingRegistryCoordinator.contract.Call(opts, &out, "getOperator", operator)

	if err != nil {
		return *new(ISlashingRegistryCoordinatorTypesOperatorInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(ISlashingRegistryCoordinatorTypesOperatorInfo)).(*ISlashingRegistryCoordinatorTypesOperatorInfo)

	return out0, err

}

// GetOperator is a free data retrieval call binding the contract method 0x5865c60c.
//
// Solidity: function getOperator(address operator) view returns((bytes32,uint8))
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) GetOperator(operator common.Address) (ISlashingRegistryCoordinatorTypesOperatorInfo, error) {
	return _ContractSlashingRegistryCoordinator.Contract.GetOperator(&_ContractSlashingRegistryCoordinator.CallOpts, operator)
}

// GetOperator is a free data retrieval call binding the contract method 0x5865c60c.
//
// Solidity: function getOperator(address operator) view returns((bytes32,uint8))
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCallerSession) GetOperator(operator common.Address) (ISlashingRegistryCoordinatorTypesOperatorInfo, error) {
	return _ContractSlashingRegistryCoordinator.Contract.GetOperator(&_ContractSlashingRegistryCoordinator.CallOpts, operator)
}

// GetOperatorFromId is a free data retrieval call binding the contract method 0x296bb064.
//
// Solidity: function getOperatorFromId(bytes32 operatorId) view returns(address)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCaller) GetOperatorFromId(opts *bind.CallOpts, operatorId [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ContractSlashingRegistryCoordinator.contract.Call(opts, &out, "getOperatorFromId", operatorId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOperatorFromId is a free data retrieval call binding the contract method 0x296bb064.
//
// Solidity: function getOperatorFromId(bytes32 operatorId) view returns(address)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) GetOperatorFromId(operatorId [32]byte) (common.Address, error) {
	return _ContractSlashingRegistryCoordinator.Contract.GetOperatorFromId(&_ContractSlashingRegistryCoordinator.CallOpts, operatorId)
}

// GetOperatorFromId is a free data retrieval call binding the contract method 0x296bb064.
//
// Solidity: function getOperatorFromId(bytes32 operatorId) view returns(address)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCallerSession) GetOperatorFromId(operatorId [32]byte) (common.Address, error) {
	return _ContractSlashingRegistryCoordinator.Contract.GetOperatorFromId(&_ContractSlashingRegistryCoordinator.CallOpts, operatorId)
}

// GetOperatorId is a free data retrieval call binding the contract method 0x13542a4e.
//
// Solidity: function getOperatorId(address operator) view returns(bytes32)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCaller) GetOperatorId(opts *bind.CallOpts, operator common.Address) ([32]byte, error) {
	var out []interface{}
	err := _ContractSlashingRegistryCoordinator.contract.Call(opts, &out, "getOperatorId", operator)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetOperatorId is a free data retrieval call binding the contract method 0x13542a4e.
//
// Solidity: function getOperatorId(address operator) view returns(bytes32)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) GetOperatorId(operator common.Address) ([32]byte, error) {
	return _ContractSlashingRegistryCoordinator.Contract.GetOperatorId(&_ContractSlashingRegistryCoordinator.CallOpts, operator)
}

// GetOperatorId is a free data retrieval call binding the contract method 0x13542a4e.
//
// Solidity: function getOperatorId(address operator) view returns(bytes32)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCallerSession) GetOperatorId(operator common.Address) ([32]byte, error) {
	return _ContractSlashingRegistryCoordinator.Contract.GetOperatorId(&_ContractSlashingRegistryCoordinator.CallOpts, operator)
}

// GetOperatorSetParams is a free data retrieval call binding the contract method 0xe65797ad.
//
// Solidity: function getOperatorSetParams(uint8 quorumNumber) view returns((uint32,uint16,uint16))
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCaller) GetOperatorSetParams(opts *bind.CallOpts, quorumNumber uint8) (ISlashingRegistryCoordinatorTypesOperatorSetParam, error) {
	var out []interface{}
	err := _ContractSlashingRegistryCoordinator.contract.Call(opts, &out, "getOperatorSetParams", quorumNumber)

	if err != nil {
		return *new(ISlashingRegistryCoordinatorTypesOperatorSetParam), err
	}

	out0 := *abi.ConvertType(out[0], new(ISlashingRegistryCoordinatorTypesOperatorSetParam)).(*ISlashingRegistryCoordinatorTypesOperatorSetParam)

	return out0, err

}

// GetOperatorSetParams is a free data retrieval call binding the contract method 0xe65797ad.
//
// Solidity: function getOperatorSetParams(uint8 quorumNumber) view returns((uint32,uint16,uint16))
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) GetOperatorSetParams(quorumNumber uint8) (ISlashingRegistryCoordinatorTypesOperatorSetParam, error) {
	return _ContractSlashingRegistryCoordinator.Contract.GetOperatorSetParams(&_ContractSlashingRegistryCoordinator.CallOpts, quorumNumber)
}

// GetOperatorSetParams is a free data retrieval call binding the contract method 0xe65797ad.
//
// Solidity: function getOperatorSetParams(uint8 quorumNumber) view returns((uint32,uint16,uint16))
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCallerSession) GetOperatorSetParams(quorumNumber uint8) (ISlashingRegistryCoordinatorTypesOperatorSetParam, error) {
	return _ContractSlashingRegistryCoordinator.Contract.GetOperatorSetParams(&_ContractSlashingRegistryCoordinator.CallOpts, quorumNumber)
}

// GetOperatorStatus is a free data retrieval call binding the contract method 0xfd39105a.
//
// Solidity: function getOperatorStatus(address operator) view returns(uint8)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCaller) GetOperatorStatus(opts *bind.CallOpts, operator common.Address) (uint8, error) {
	var out []interface{}
	err := _ContractSlashingRegistryCoordinator.contract.Call(opts, &out, "getOperatorStatus", operator)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetOperatorStatus is a free data retrieval call binding the contract method 0xfd39105a.
//
// Solidity: function getOperatorStatus(address operator) view returns(uint8)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) GetOperatorStatus(operator common.Address) (uint8, error) {
	return _ContractSlashingRegistryCoordinator.Contract.GetOperatorStatus(&_ContractSlashingRegistryCoordinator.CallOpts, operator)
}

// GetOperatorStatus is a free data retrieval call binding the contract method 0xfd39105a.
//
// Solidity: function getOperatorStatus(address operator) view returns(uint8)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCallerSession) GetOperatorStatus(operator common.Address) (uint8, error) {
	return _ContractSlashingRegistryCoordinator.Contract.GetOperatorStatus(&_ContractSlashingRegistryCoordinator.CallOpts, operator)
}

// GetQuorumBitmapAtBlockNumberByIndex is a free data retrieval call binding the contract method 0x04ec6351.
//
// Solidity: function getQuorumBitmapAtBlockNumberByIndex(bytes32 operatorId, uint32 blockNumber, uint256 index) view returns(uint192)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCaller) GetQuorumBitmapAtBlockNumberByIndex(opts *bind.CallOpts, operatorId [32]byte, blockNumber uint32, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ContractSlashingRegistryCoordinator.contract.Call(opts, &out, "getQuorumBitmapAtBlockNumberByIndex", operatorId, blockNumber, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetQuorumBitmapAtBlockNumberByIndex is a free data retrieval call binding the contract method 0x04ec6351.
//
// Solidity: function getQuorumBitmapAtBlockNumberByIndex(bytes32 operatorId, uint32 blockNumber, uint256 index) view returns(uint192)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) GetQuorumBitmapAtBlockNumberByIndex(operatorId [32]byte, blockNumber uint32, index *big.Int) (*big.Int, error) {
	return _ContractSlashingRegistryCoordinator.Contract.GetQuorumBitmapAtBlockNumberByIndex(&_ContractSlashingRegistryCoordinator.CallOpts, operatorId, blockNumber, index)
}

// GetQuorumBitmapAtBlockNumberByIndex is a free data retrieval call binding the contract method 0x04ec6351.
//
// Solidity: function getQuorumBitmapAtBlockNumberByIndex(bytes32 operatorId, uint32 blockNumber, uint256 index) view returns(uint192)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCallerSession) GetQuorumBitmapAtBlockNumberByIndex(operatorId [32]byte, blockNumber uint32, index *big.Int) (*big.Int, error) {
	return _ContractSlashingRegistryCoordinator.Contract.GetQuorumBitmapAtBlockNumberByIndex(&_ContractSlashingRegistryCoordinator.CallOpts, operatorId, blockNumber, index)
}

// GetQuorumBitmapHistoryLength is a free data retrieval call binding the contract method 0x03fd3492.
//
// Solidity: function getQuorumBitmapHistoryLength(bytes32 operatorId) view returns(uint256)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCaller) GetQuorumBitmapHistoryLength(opts *bind.CallOpts, operatorId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ContractSlashingRegistryCoordinator.contract.Call(opts, &out, "getQuorumBitmapHistoryLength", operatorId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetQuorumBitmapHistoryLength is a free data retrieval call binding the contract method 0x03fd3492.
//
// Solidity: function getQuorumBitmapHistoryLength(bytes32 operatorId) view returns(uint256)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) GetQuorumBitmapHistoryLength(operatorId [32]byte) (*big.Int, error) {
	return _ContractSlashingRegistryCoordinator.Contract.GetQuorumBitmapHistoryLength(&_ContractSlashingRegistryCoordinator.CallOpts, operatorId)
}

// GetQuorumBitmapHistoryLength is a free data retrieval call binding the contract method 0x03fd3492.
//
// Solidity: function getQuorumBitmapHistoryLength(bytes32 operatorId) view returns(uint256)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCallerSession) GetQuorumBitmapHistoryLength(operatorId [32]byte) (*big.Int, error) {
	return _ContractSlashingRegistryCoordinator.Contract.GetQuorumBitmapHistoryLength(&_ContractSlashingRegistryCoordinator.CallOpts, operatorId)
}

// GetQuorumBitmapIndicesAtBlockNumber is a free data retrieval call binding the contract method 0xc391425e.
//
// Solidity: function getQuorumBitmapIndicesAtBlockNumber(uint32 blockNumber, bytes32[] operatorIds) view returns(uint32[])
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCaller) GetQuorumBitmapIndicesAtBlockNumber(opts *bind.CallOpts, blockNumber uint32, operatorIds [][32]byte) ([]uint32, error) {
	var out []interface{}
	err := _ContractSlashingRegistryCoordinator.contract.Call(opts, &out, "getQuorumBitmapIndicesAtBlockNumber", blockNumber, operatorIds)

	if err != nil {
		return *new([]uint32), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint32)).(*[]uint32)

	return out0, err

}

// GetQuorumBitmapIndicesAtBlockNumber is a free data retrieval call binding the contract method 0xc391425e.
//
// Solidity: function getQuorumBitmapIndicesAtBlockNumber(uint32 blockNumber, bytes32[] operatorIds) view returns(uint32[])
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) GetQuorumBitmapIndicesAtBlockNumber(blockNumber uint32, operatorIds [][32]byte) ([]uint32, error) {
	return _ContractSlashingRegistryCoordinator.Contract.GetQuorumBitmapIndicesAtBlockNumber(&_ContractSlashingRegistryCoordinator.CallOpts, blockNumber, operatorIds)
}

// GetQuorumBitmapIndicesAtBlockNumber is a free data retrieval call binding the contract method 0xc391425e.
//
// Solidity: function getQuorumBitmapIndicesAtBlockNumber(uint32 blockNumber, bytes32[] operatorIds) view returns(uint32[])
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCallerSession) GetQuorumBitmapIndicesAtBlockNumber(blockNumber uint32, operatorIds [][32]byte) ([]uint32, error) {
	return _ContractSlashingRegistryCoordinator.Contract.GetQuorumBitmapIndicesAtBlockNumber(&_ContractSlashingRegistryCoordinator.CallOpts, blockNumber, operatorIds)
}

// GetQuorumBitmapUpdateByIndex is a free data retrieval call binding the contract method 0x1eb812da.
//
// Solidity: function getQuorumBitmapUpdateByIndex(bytes32 operatorId, uint256 index) view returns((uint32,uint32,uint192))
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCaller) GetQuorumBitmapUpdateByIndex(opts *bind.CallOpts, operatorId [32]byte, index *big.Int) (ISlashingRegistryCoordinatorTypesQuorumBitmapUpdate, error) {
	var out []interface{}
	err := _ContractSlashingRegistryCoordinator.contract.Call(opts, &out, "getQuorumBitmapUpdateByIndex", operatorId, index)

	if err != nil {
		return *new(ISlashingRegistryCoordinatorTypesQuorumBitmapUpdate), err
	}

	out0 := *abi.ConvertType(out[0], new(ISlashingRegistryCoordinatorTypesQuorumBitmapUpdate)).(*ISlashingRegistryCoordinatorTypesQuorumBitmapUpdate)

	return out0, err

}

// GetQuorumBitmapUpdateByIndex is a free data retrieval call binding the contract method 0x1eb812da.
//
// Solidity: function getQuorumBitmapUpdateByIndex(bytes32 operatorId, uint256 index) view returns((uint32,uint32,uint192))
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) GetQuorumBitmapUpdateByIndex(operatorId [32]byte, index *big.Int) (ISlashingRegistryCoordinatorTypesQuorumBitmapUpdate, error) {
	return _ContractSlashingRegistryCoordinator.Contract.GetQuorumBitmapUpdateByIndex(&_ContractSlashingRegistryCoordinator.CallOpts, operatorId, index)
}

// GetQuorumBitmapUpdateByIndex is a free data retrieval call binding the contract method 0x1eb812da.
//
// Solidity: function getQuorumBitmapUpdateByIndex(bytes32 operatorId, uint256 index) view returns((uint32,uint32,uint192))
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCallerSession) GetQuorumBitmapUpdateByIndex(operatorId [32]byte, index *big.Int) (ISlashingRegistryCoordinatorTypesQuorumBitmapUpdate, error) {
	return _ContractSlashingRegistryCoordinator.Contract.GetQuorumBitmapUpdateByIndex(&_ContractSlashingRegistryCoordinator.CallOpts, operatorId, index)
}

// IndexRegistry is a free data retrieval call binding the contract method 0x9e9923c2.
//
// Solidity: function indexRegistry() view returns(address)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCaller) IndexRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractSlashingRegistryCoordinator.contract.Call(opts, &out, "indexRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IndexRegistry is a free data retrieval call binding the contract method 0x9e9923c2.
//
// Solidity: function indexRegistry() view returns(address)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) IndexRegistry() (common.Address, error) {
	return _ContractSlashingRegistryCoordinator.Contract.IndexRegistry(&_ContractSlashingRegistryCoordinator.CallOpts)
}

// IndexRegistry is a free data retrieval call binding the contract method 0x9e9923c2.
//
// Solidity: function indexRegistry() view returns(address)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCallerSession) IndexRegistry() (common.Address, error) {
	return _ContractSlashingRegistryCoordinator.Contract.IndexRegistry(&_ContractSlashingRegistryCoordinator.CallOpts)
}

// IsChurnApproverSaltUsed is a free data retrieval call binding the contract method 0x1478851f.
//
// Solidity: function isChurnApproverSaltUsed(bytes32 ) view returns(bool)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCaller) IsChurnApproverSaltUsed(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _ContractSlashingRegistryCoordinator.contract.Call(opts, &out, "isChurnApproverSaltUsed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsChurnApproverSaltUsed is a free data retrieval call binding the contract method 0x1478851f.
//
// Solidity: function isChurnApproverSaltUsed(bytes32 ) view returns(bool)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) IsChurnApproverSaltUsed(arg0 [32]byte) (bool, error) {
	return _ContractSlashingRegistryCoordinator.Contract.IsChurnApproverSaltUsed(&_ContractSlashingRegistryCoordinator.CallOpts, arg0)
}

// IsChurnApproverSaltUsed is a free data retrieval call binding the contract method 0x1478851f.
//
// Solidity: function isChurnApproverSaltUsed(bytes32 ) view returns(bool)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCallerSession) IsChurnApproverSaltUsed(arg0 [32]byte) (bool, error) {
	return _ContractSlashingRegistryCoordinator.Contract.IsChurnApproverSaltUsed(&_ContractSlashingRegistryCoordinator.CallOpts, arg0)
}

// LastEjectionTimestamp is a free data retrieval call binding the contract method 0x125e0584.
//
// Solidity: function lastEjectionTimestamp(address ) view returns(uint256)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCaller) LastEjectionTimestamp(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ContractSlashingRegistryCoordinator.contract.Call(opts, &out, "lastEjectionTimestamp", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastEjectionTimestamp is a free data retrieval call binding the contract method 0x125e0584.
//
// Solidity: function lastEjectionTimestamp(address ) view returns(uint256)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) LastEjectionTimestamp(arg0 common.Address) (*big.Int, error) {
	return _ContractSlashingRegistryCoordinator.Contract.LastEjectionTimestamp(&_ContractSlashingRegistryCoordinator.CallOpts, arg0)
}

// LastEjectionTimestamp is a free data retrieval call binding the contract method 0x125e0584.
//
// Solidity: function lastEjectionTimestamp(address ) view returns(uint256)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCallerSession) LastEjectionTimestamp(arg0 common.Address) (*big.Int, error) {
	return _ContractSlashingRegistryCoordinator.Contract.LastEjectionTimestamp(&_ContractSlashingRegistryCoordinator.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractSlashingRegistryCoordinator.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) Owner() (common.Address, error) {
	return _ContractSlashingRegistryCoordinator.Contract.Owner(&_ContractSlashingRegistryCoordinator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCallerSession) Owner() (common.Address, error) {
	return _ContractSlashingRegistryCoordinator.Contract.Owner(&_ContractSlashingRegistryCoordinator.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _ContractSlashingRegistryCoordinator.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) Paused(index uint8) (bool, error) {
	return _ContractSlashingRegistryCoordinator.Contract.Paused(&_ContractSlashingRegistryCoordinator.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCallerSession) Paused(index uint8) (bool, error) {
	return _ContractSlashingRegistryCoordinator.Contract.Paused(&_ContractSlashingRegistryCoordinator.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractSlashingRegistryCoordinator.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) Paused0() (*big.Int, error) {
	return _ContractSlashingRegistryCoordinator.Contract.Paused0(&_ContractSlashingRegistryCoordinator.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCallerSession) Paused0() (*big.Int, error) {
	return _ContractSlashingRegistryCoordinator.Contract.Paused0(&_ContractSlashingRegistryCoordinator.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractSlashingRegistryCoordinator.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) PauserRegistry() (common.Address, error) {
	return _ContractSlashingRegistryCoordinator.Contract.PauserRegistry(&_ContractSlashingRegistryCoordinator.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCallerSession) PauserRegistry() (common.Address, error) {
	return _ContractSlashingRegistryCoordinator.Contract.PauserRegistry(&_ContractSlashingRegistryCoordinator.CallOpts)
}

// PubkeyRegistrationMessageHash is a free data retrieval call binding the contract method 0x3c2a7f4c.
//
// Solidity: function pubkeyRegistrationMessageHash(address operator) view returns((uint256,uint256))
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCaller) PubkeyRegistrationMessageHash(opts *bind.CallOpts, operator common.Address) (BN254G1Point, error) {
	var out []interface{}
	err := _ContractSlashingRegistryCoordinator.contract.Call(opts, &out, "pubkeyRegistrationMessageHash", operator)

	if err != nil {
		return *new(BN254G1Point), err
	}

	out0 := *abi.ConvertType(out[0], new(BN254G1Point)).(*BN254G1Point)

	return out0, err

}

// PubkeyRegistrationMessageHash is a free data retrieval call binding the contract method 0x3c2a7f4c.
//
// Solidity: function pubkeyRegistrationMessageHash(address operator) view returns((uint256,uint256))
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) PubkeyRegistrationMessageHash(operator common.Address) (BN254G1Point, error) {
	return _ContractSlashingRegistryCoordinator.Contract.PubkeyRegistrationMessageHash(&_ContractSlashingRegistryCoordinator.CallOpts, operator)
}

// PubkeyRegistrationMessageHash is a free data retrieval call binding the contract method 0x3c2a7f4c.
//
// Solidity: function pubkeyRegistrationMessageHash(address operator) view returns((uint256,uint256))
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCallerSession) PubkeyRegistrationMessageHash(operator common.Address) (BN254G1Point, error) {
	return _ContractSlashingRegistryCoordinator.Contract.PubkeyRegistrationMessageHash(&_ContractSlashingRegistryCoordinator.CallOpts, operator)
}

// QuorumCount is a free data retrieval call binding the contract method 0x9aa1653d.
//
// Solidity: function quorumCount() view returns(uint8)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCaller) QuorumCount(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ContractSlashingRegistryCoordinator.contract.Call(opts, &out, "quorumCount")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// QuorumCount is a free data retrieval call binding the contract method 0x9aa1653d.
//
// Solidity: function quorumCount() view returns(uint8)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) QuorumCount() (uint8, error) {
	return _ContractSlashingRegistryCoordinator.Contract.QuorumCount(&_ContractSlashingRegistryCoordinator.CallOpts)
}

// QuorumCount is a free data retrieval call binding the contract method 0x9aa1653d.
//
// Solidity: function quorumCount() view returns(uint8)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCallerSession) QuorumCount() (uint8, error) {
	return _ContractSlashingRegistryCoordinator.Contract.QuorumCount(&_ContractSlashingRegistryCoordinator.CallOpts)
}

// QuorumUpdateBlockNumber is a free data retrieval call binding the contract method 0x249a0c42.
//
// Solidity: function quorumUpdateBlockNumber(uint8 ) view returns(uint256)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCaller) QuorumUpdateBlockNumber(opts *bind.CallOpts, arg0 uint8) (*big.Int, error) {
	var out []interface{}
	err := _ContractSlashingRegistryCoordinator.contract.Call(opts, &out, "quorumUpdateBlockNumber", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuorumUpdateBlockNumber is a free data retrieval call binding the contract method 0x249a0c42.
//
// Solidity: function quorumUpdateBlockNumber(uint8 ) view returns(uint256)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) QuorumUpdateBlockNumber(arg0 uint8) (*big.Int, error) {
	return _ContractSlashingRegistryCoordinator.Contract.QuorumUpdateBlockNumber(&_ContractSlashingRegistryCoordinator.CallOpts, arg0)
}

// QuorumUpdateBlockNumber is a free data retrieval call binding the contract method 0x249a0c42.
//
// Solidity: function quorumUpdateBlockNumber(uint8 ) view returns(uint256)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCallerSession) QuorumUpdateBlockNumber(arg0 uint8) (*big.Int, error) {
	return _ContractSlashingRegistryCoordinator.Contract.QuorumUpdateBlockNumber(&_ContractSlashingRegistryCoordinator.CallOpts, arg0)
}

// SocketRegistry is a free data retrieval call binding the contract method 0xea32afae.
//
// Solidity: function socketRegistry() view returns(address)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCaller) SocketRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractSlashingRegistryCoordinator.contract.Call(opts, &out, "socketRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SocketRegistry is a free data retrieval call binding the contract method 0xea32afae.
//
// Solidity: function socketRegistry() view returns(address)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) SocketRegistry() (common.Address, error) {
	return _ContractSlashingRegistryCoordinator.Contract.SocketRegistry(&_ContractSlashingRegistryCoordinator.CallOpts)
}

// SocketRegistry is a free data retrieval call binding the contract method 0xea32afae.
//
// Solidity: function socketRegistry() view returns(address)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCallerSession) SocketRegistry() (common.Address, error) {
	return _ContractSlashingRegistryCoordinator.Contract.SocketRegistry(&_ContractSlashingRegistryCoordinator.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCaller) StakeRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractSlashingRegistryCoordinator.contract.Call(opts, &out, "stakeRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) StakeRegistry() (common.Address, error) {
	return _ContractSlashingRegistryCoordinator.Contract.StakeRegistry(&_ContractSlashingRegistryCoordinator.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCallerSession) StakeRegistry() (common.Address, error) {
	return _ContractSlashingRegistryCoordinator.Contract.StakeRegistry(&_ContractSlashingRegistryCoordinator.CallOpts)
}

// SupportsAVS is a free data retrieval call binding the contract method 0xb5265787.
//
// Solidity: function supportsAVS(address _avs) view returns(bool)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCaller) SupportsAVS(opts *bind.CallOpts, _avs common.Address) (bool, error) {
	var out []interface{}
	err := _ContractSlashingRegistryCoordinator.contract.Call(opts, &out, "supportsAVS", _avs)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsAVS is a free data retrieval call binding the contract method 0xb5265787.
//
// Solidity: function supportsAVS(address _avs) view returns(bool)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) SupportsAVS(_avs common.Address) (bool, error) {
	return _ContractSlashingRegistryCoordinator.Contract.SupportsAVS(&_ContractSlashingRegistryCoordinator.CallOpts, _avs)
}

// SupportsAVS is a free data retrieval call binding the contract method 0xb5265787.
//
// Solidity: function supportsAVS(address _avs) view returns(bool)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCallerSession) SupportsAVS(_avs common.Address) (bool, error) {
	return _ContractSlashingRegistryCoordinator.Contract.SupportsAVS(&_ContractSlashingRegistryCoordinator.CallOpts, _avs)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ContractSlashingRegistryCoordinator.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) Version() (string, error) {
	return _ContractSlashingRegistryCoordinator.Contract.Version(&_ContractSlashingRegistryCoordinator.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCallerSession) Version() (string, error) {
	return _ContractSlashingRegistryCoordinator.Contract.Version(&_ContractSlashingRegistryCoordinator.CallOpts)
}

// CreateSlashableStakeQuorum is a paid mutator transaction binding the contract method 0x3eef3a51.
//
// Solidity: function createSlashableStakeQuorum((uint32,uint16,uint16) operatorSetParams, uint96 minimumStake, (address,uint96)[] strategyParams, uint32 lookAheadPeriod) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorTransactor) CreateSlashableStakeQuorum(opts *bind.TransactOpts, operatorSetParams ISlashingRegistryCoordinatorTypesOperatorSetParam, minimumStake *big.Int, strategyParams []IStakeRegistryTypesStrategyParams, lookAheadPeriod uint32) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.contract.Transact(opts, "createSlashableStakeQuorum", operatorSetParams, minimumStake, strategyParams, lookAheadPeriod)
}

// CreateSlashableStakeQuorum is a paid mutator transaction binding the contract method 0x3eef3a51.
//
// Solidity: function createSlashableStakeQuorum((uint32,uint16,uint16) operatorSetParams, uint96 minimumStake, (address,uint96)[] strategyParams, uint32 lookAheadPeriod) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) CreateSlashableStakeQuorum(operatorSetParams ISlashingRegistryCoordinatorTypesOperatorSetParam, minimumStake *big.Int, strategyParams []IStakeRegistryTypesStrategyParams, lookAheadPeriod uint32) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.Contract.CreateSlashableStakeQuorum(&_ContractSlashingRegistryCoordinator.TransactOpts, operatorSetParams, minimumStake, strategyParams, lookAheadPeriod)
}

// CreateSlashableStakeQuorum is a paid mutator transaction binding the contract method 0x3eef3a51.
//
// Solidity: function createSlashableStakeQuorum((uint32,uint16,uint16) operatorSetParams, uint96 minimumStake, (address,uint96)[] strategyParams, uint32 lookAheadPeriod) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorTransactorSession) CreateSlashableStakeQuorum(operatorSetParams ISlashingRegistryCoordinatorTypesOperatorSetParam, minimumStake *big.Int, strategyParams []IStakeRegistryTypesStrategyParams, lookAheadPeriod uint32) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.Contract.CreateSlashableStakeQuorum(&_ContractSlashingRegistryCoordinator.TransactOpts, operatorSetParams, minimumStake, strategyParams, lookAheadPeriod)
}

// CreateTotalDelegatedStakeQuorum is a paid mutator transaction binding the contract method 0x8281ab75.
//
// Solidity: function createTotalDelegatedStakeQuorum((uint32,uint16,uint16) operatorSetParams, uint96 minimumStake, (address,uint96)[] strategyParams) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorTransactor) CreateTotalDelegatedStakeQuorum(opts *bind.TransactOpts, operatorSetParams ISlashingRegistryCoordinatorTypesOperatorSetParam, minimumStake *big.Int, strategyParams []IStakeRegistryTypesStrategyParams) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.contract.Transact(opts, "createTotalDelegatedStakeQuorum", operatorSetParams, minimumStake, strategyParams)
}

// CreateTotalDelegatedStakeQuorum is a paid mutator transaction binding the contract method 0x8281ab75.
//
// Solidity: function createTotalDelegatedStakeQuorum((uint32,uint16,uint16) operatorSetParams, uint96 minimumStake, (address,uint96)[] strategyParams) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) CreateTotalDelegatedStakeQuorum(operatorSetParams ISlashingRegistryCoordinatorTypesOperatorSetParam, minimumStake *big.Int, strategyParams []IStakeRegistryTypesStrategyParams) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.Contract.CreateTotalDelegatedStakeQuorum(&_ContractSlashingRegistryCoordinator.TransactOpts, operatorSetParams, minimumStake, strategyParams)
}

// CreateTotalDelegatedStakeQuorum is a paid mutator transaction binding the contract method 0x8281ab75.
//
// Solidity: function createTotalDelegatedStakeQuorum((uint32,uint16,uint16) operatorSetParams, uint96 minimumStake, (address,uint96)[] strategyParams) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorTransactorSession) CreateTotalDelegatedStakeQuorum(operatorSetParams ISlashingRegistryCoordinatorTypesOperatorSetParam, minimumStake *big.Int, strategyParams []IStakeRegistryTypesStrategyParams) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.Contract.CreateTotalDelegatedStakeQuorum(&_ContractSlashingRegistryCoordinator.TransactOpts, operatorSetParams, minimumStake, strategyParams)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0x303ca956.
//
// Solidity: function deregisterOperator(address operator, address avs, uint32[] operatorSetIds) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorTransactor) DeregisterOperator(opts *bind.TransactOpts, operator common.Address, avs common.Address, operatorSetIds []uint32) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.contract.Transact(opts, "deregisterOperator", operator, avs, operatorSetIds)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0x303ca956.
//
// Solidity: function deregisterOperator(address operator, address avs, uint32[] operatorSetIds) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) DeregisterOperator(operator common.Address, avs common.Address, operatorSetIds []uint32) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.Contract.DeregisterOperator(&_ContractSlashingRegistryCoordinator.TransactOpts, operator, avs, operatorSetIds)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0x303ca956.
//
// Solidity: function deregisterOperator(address operator, address avs, uint32[] operatorSetIds) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorTransactorSession) DeregisterOperator(operator common.Address, avs common.Address, operatorSetIds []uint32) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.Contract.DeregisterOperator(&_ContractSlashingRegistryCoordinator.TransactOpts, operator, avs, operatorSetIds)
}

// EjectOperator is a paid mutator transaction binding the contract method 0x6e3b17db.
//
// Solidity: function ejectOperator(address operator, bytes quorumNumbers) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorTransactor) EjectOperator(opts *bind.TransactOpts, operator common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.contract.Transact(opts, "ejectOperator", operator, quorumNumbers)
}

// EjectOperator is a paid mutator transaction binding the contract method 0x6e3b17db.
//
// Solidity: function ejectOperator(address operator, bytes quorumNumbers) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) EjectOperator(operator common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.Contract.EjectOperator(&_ContractSlashingRegistryCoordinator.TransactOpts, operator, quorumNumbers)
}

// EjectOperator is a paid mutator transaction binding the contract method 0x6e3b17db.
//
// Solidity: function ejectOperator(address operator, bytes quorumNumbers) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorTransactorSession) EjectOperator(operator common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.Contract.EjectOperator(&_ContractSlashingRegistryCoordinator.TransactOpts, operator, quorumNumbers)
}

// Initialize is a paid mutator transaction binding the contract method 0x530b97a4.
//
// Solidity: function initialize(address initialOwner, address churnApprover, address ejector, uint256 initialPausedStatus, address avs) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address, churnApprover common.Address, ejector common.Address, initialPausedStatus *big.Int, avs common.Address) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.contract.Transact(opts, "initialize", initialOwner, churnApprover, ejector, initialPausedStatus, avs)
}

// Initialize is a paid mutator transaction binding the contract method 0x530b97a4.
//
// Solidity: function initialize(address initialOwner, address churnApprover, address ejector, uint256 initialPausedStatus, address avs) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) Initialize(initialOwner common.Address, churnApprover common.Address, ejector common.Address, initialPausedStatus *big.Int, avs common.Address) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.Contract.Initialize(&_ContractSlashingRegistryCoordinator.TransactOpts, initialOwner, churnApprover, ejector, initialPausedStatus, avs)
}

// Initialize is a paid mutator transaction binding the contract method 0x530b97a4.
//
// Solidity: function initialize(address initialOwner, address churnApprover, address ejector, uint256 initialPausedStatus, address avs) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorTransactorSession) Initialize(initialOwner common.Address, churnApprover common.Address, ejector common.Address, initialPausedStatus *big.Int, avs common.Address) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.Contract.Initialize(&_ContractSlashingRegistryCoordinator.TransactOpts, initialOwner, churnApprover, ejector, initialPausedStatus, avs)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.Contract.Pause(&_ContractSlashingRegistryCoordinator.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.Contract.Pause(&_ContractSlashingRegistryCoordinator.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) PauseAll() (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.Contract.PauseAll(&_ContractSlashingRegistryCoordinator.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorTransactorSession) PauseAll() (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.Contract.PauseAll(&_ContractSlashingRegistryCoordinator.TransactOpts)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0xc63fd502.
//
// Solidity: function registerOperator(address operator, address avs, uint32[] operatorSetIds, bytes data) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorTransactor) RegisterOperator(opts *bind.TransactOpts, operator common.Address, avs common.Address, operatorSetIds []uint32, data []byte) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.contract.Transact(opts, "registerOperator", operator, avs, operatorSetIds, data)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0xc63fd502.
//
// Solidity: function registerOperator(address operator, address avs, uint32[] operatorSetIds, bytes data) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) RegisterOperator(operator common.Address, avs common.Address, operatorSetIds []uint32, data []byte) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.Contract.RegisterOperator(&_ContractSlashingRegistryCoordinator.TransactOpts, operator, avs, operatorSetIds, data)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0xc63fd502.
//
// Solidity: function registerOperator(address operator, address avs, uint32[] operatorSetIds, bytes data) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorTransactorSession) RegisterOperator(operator common.Address, avs common.Address, operatorSetIds []uint32, data []byte) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.Contract.RegisterOperator(&_ContractSlashingRegistryCoordinator.TransactOpts, operator, avs, operatorSetIds, data)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.Contract.RenounceOwnership(&_ContractSlashingRegistryCoordinator.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.Contract.RenounceOwnership(&_ContractSlashingRegistryCoordinator.TransactOpts)
}

// SetAVS is a paid mutator transaction binding the contract method 0xa65497c6.
//
// Solidity: function setAVS(address _avs) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorTransactor) SetAVS(opts *bind.TransactOpts, _avs common.Address) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.contract.Transact(opts, "setAVS", _avs)
}

// SetAVS is a paid mutator transaction binding the contract method 0xa65497c6.
//
// Solidity: function setAVS(address _avs) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) SetAVS(_avs common.Address) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.Contract.SetAVS(&_ContractSlashingRegistryCoordinator.TransactOpts, _avs)
}

// SetAVS is a paid mutator transaction binding the contract method 0xa65497c6.
//
// Solidity: function setAVS(address _avs) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorTransactorSession) SetAVS(_avs common.Address) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.Contract.SetAVS(&_ContractSlashingRegistryCoordinator.TransactOpts, _avs)
}

// SetChurnApprover is a paid mutator transaction binding the contract method 0x29d1e0c3.
//
// Solidity: function setChurnApprover(address _churnApprover) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorTransactor) SetChurnApprover(opts *bind.TransactOpts, _churnApprover common.Address) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.contract.Transact(opts, "setChurnApprover", _churnApprover)
}

// SetChurnApprover is a paid mutator transaction binding the contract method 0x29d1e0c3.
//
// Solidity: function setChurnApprover(address _churnApprover) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) SetChurnApprover(_churnApprover common.Address) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.Contract.SetChurnApprover(&_ContractSlashingRegistryCoordinator.TransactOpts, _churnApprover)
}

// SetChurnApprover is a paid mutator transaction binding the contract method 0x29d1e0c3.
//
// Solidity: function setChurnApprover(address _churnApprover) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorTransactorSession) SetChurnApprover(_churnApprover common.Address) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.Contract.SetChurnApprover(&_ContractSlashingRegistryCoordinator.TransactOpts, _churnApprover)
}

// SetEjectionCooldown is a paid mutator transaction binding the contract method 0x0d3f2134.
//
// Solidity: function setEjectionCooldown(uint256 _ejectionCooldown) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorTransactor) SetEjectionCooldown(opts *bind.TransactOpts, _ejectionCooldown *big.Int) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.contract.Transact(opts, "setEjectionCooldown", _ejectionCooldown)
}

// SetEjectionCooldown is a paid mutator transaction binding the contract method 0x0d3f2134.
//
// Solidity: function setEjectionCooldown(uint256 _ejectionCooldown) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) SetEjectionCooldown(_ejectionCooldown *big.Int) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.Contract.SetEjectionCooldown(&_ContractSlashingRegistryCoordinator.TransactOpts, _ejectionCooldown)
}

// SetEjectionCooldown is a paid mutator transaction binding the contract method 0x0d3f2134.
//
// Solidity: function setEjectionCooldown(uint256 _ejectionCooldown) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorTransactorSession) SetEjectionCooldown(_ejectionCooldown *big.Int) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.Contract.SetEjectionCooldown(&_ContractSlashingRegistryCoordinator.TransactOpts, _ejectionCooldown)
}

// SetEjector is a paid mutator transaction binding the contract method 0x2cdd1e86.
//
// Solidity: function setEjector(address _ejector) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorTransactor) SetEjector(opts *bind.TransactOpts, _ejector common.Address) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.contract.Transact(opts, "setEjector", _ejector)
}

// SetEjector is a paid mutator transaction binding the contract method 0x2cdd1e86.
//
// Solidity: function setEjector(address _ejector) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) SetEjector(_ejector common.Address) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.Contract.SetEjector(&_ContractSlashingRegistryCoordinator.TransactOpts, _ejector)
}

// SetEjector is a paid mutator transaction binding the contract method 0x2cdd1e86.
//
// Solidity: function setEjector(address _ejector) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorTransactorSession) SetEjector(_ejector common.Address) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.Contract.SetEjector(&_ContractSlashingRegistryCoordinator.TransactOpts, _ejector)
}

// SetOperatorSetParams is a paid mutator transaction binding the contract method 0x5b0b829f.
//
// Solidity: function setOperatorSetParams(uint8 quorumNumber, (uint32,uint16,uint16) operatorSetParams) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorTransactor) SetOperatorSetParams(opts *bind.TransactOpts, quorumNumber uint8, operatorSetParams ISlashingRegistryCoordinatorTypesOperatorSetParam) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.contract.Transact(opts, "setOperatorSetParams", quorumNumber, operatorSetParams)
}

// SetOperatorSetParams is a paid mutator transaction binding the contract method 0x5b0b829f.
//
// Solidity: function setOperatorSetParams(uint8 quorumNumber, (uint32,uint16,uint16) operatorSetParams) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) SetOperatorSetParams(quorumNumber uint8, operatorSetParams ISlashingRegistryCoordinatorTypesOperatorSetParam) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.Contract.SetOperatorSetParams(&_ContractSlashingRegistryCoordinator.TransactOpts, quorumNumber, operatorSetParams)
}

// SetOperatorSetParams is a paid mutator transaction binding the contract method 0x5b0b829f.
//
// Solidity: function setOperatorSetParams(uint8 quorumNumber, (uint32,uint16,uint16) operatorSetParams) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorTransactorSession) SetOperatorSetParams(quorumNumber uint8, operatorSetParams ISlashingRegistryCoordinatorTypesOperatorSetParam) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.Contract.SetOperatorSetParams(&_ContractSlashingRegistryCoordinator.TransactOpts, quorumNumber, operatorSetParams)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.Contract.TransferOwnership(&_ContractSlashingRegistryCoordinator.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.Contract.TransferOwnership(&_ContractSlashingRegistryCoordinator.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.Contract.Unpause(&_ContractSlashingRegistryCoordinator.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.Contract.Unpause(&_ContractSlashingRegistryCoordinator.TransactOpts, newPausedStatus)
}

// UpdateOperators is a paid mutator transaction binding the contract method 0x00cf2ab5.
//
// Solidity: function updateOperators(address[] operators) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorTransactor) UpdateOperators(opts *bind.TransactOpts, operators []common.Address) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.contract.Transact(opts, "updateOperators", operators)
}

// UpdateOperators is a paid mutator transaction binding the contract method 0x00cf2ab5.
//
// Solidity: function updateOperators(address[] operators) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) UpdateOperators(operators []common.Address) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.Contract.UpdateOperators(&_ContractSlashingRegistryCoordinator.TransactOpts, operators)
}

// UpdateOperators is a paid mutator transaction binding the contract method 0x00cf2ab5.
//
// Solidity: function updateOperators(address[] operators) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorTransactorSession) UpdateOperators(operators []common.Address) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.Contract.UpdateOperators(&_ContractSlashingRegistryCoordinator.TransactOpts, operators)
}

// UpdateOperatorsForQuorum is a paid mutator transaction binding the contract method 0x5140a548.
//
// Solidity: function updateOperatorsForQuorum(address[][] operatorsPerQuorum, bytes quorumNumbers) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorTransactor) UpdateOperatorsForQuorum(opts *bind.TransactOpts, operatorsPerQuorum [][]common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.contract.Transact(opts, "updateOperatorsForQuorum", operatorsPerQuorum, quorumNumbers)
}

// UpdateOperatorsForQuorum is a paid mutator transaction binding the contract method 0x5140a548.
//
// Solidity: function updateOperatorsForQuorum(address[][] operatorsPerQuorum, bytes quorumNumbers) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) UpdateOperatorsForQuorum(operatorsPerQuorum [][]common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.Contract.UpdateOperatorsForQuorum(&_ContractSlashingRegistryCoordinator.TransactOpts, operatorsPerQuorum, quorumNumbers)
}

// UpdateOperatorsForQuorum is a paid mutator transaction binding the contract method 0x5140a548.
//
// Solidity: function updateOperatorsForQuorum(address[][] operatorsPerQuorum, bytes quorumNumbers) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorTransactorSession) UpdateOperatorsForQuorum(operatorsPerQuorum [][]common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.Contract.UpdateOperatorsForQuorum(&_ContractSlashingRegistryCoordinator.TransactOpts, operatorsPerQuorum, quorumNumbers)
}

// UpdateSocket is a paid mutator transaction binding the contract method 0x0cf4b767.
//
// Solidity: function updateSocket(string socket) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorTransactor) UpdateSocket(opts *bind.TransactOpts, socket string) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.contract.Transact(opts, "updateSocket", socket)
}

// UpdateSocket is a paid mutator transaction binding the contract method 0x0cf4b767.
//
// Solidity: function updateSocket(string socket) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) UpdateSocket(socket string) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.Contract.UpdateSocket(&_ContractSlashingRegistryCoordinator.TransactOpts, socket)
}

// UpdateSocket is a paid mutator transaction binding the contract method 0x0cf4b767.
//
// Solidity: function updateSocket(string socket) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorTransactorSession) UpdateSocket(socket string) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.Contract.UpdateSocket(&_ContractSlashingRegistryCoordinator.TransactOpts, socket)
}

// ContractSlashingRegistryCoordinatorAVSUpdatedIterator is returned from FilterAVSUpdated and is used to iterate over the raw logs and unpacked data for AVSUpdated events raised by the ContractSlashingRegistryCoordinator contract.
type ContractSlashingRegistryCoordinatorAVSUpdatedIterator struct {
	Event *ContractSlashingRegistryCoordinatorAVSUpdated // Event containing the contract specifics and raw log

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
func (it *ContractSlashingRegistryCoordinatorAVSUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSlashingRegistryCoordinatorAVSUpdated)
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
		it.Event = new(ContractSlashingRegistryCoordinatorAVSUpdated)
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
func (it *ContractSlashingRegistryCoordinatorAVSUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSlashingRegistryCoordinatorAVSUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSlashingRegistryCoordinatorAVSUpdated represents a AVSUpdated event raised by the ContractSlashingRegistryCoordinator contract.
type ContractSlashingRegistryCoordinatorAVSUpdated struct {
	PrevAVS common.Address
	NewAVS  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAVSUpdated is a free log retrieval operation binding the contract event 0x9770f3cadfdcbb6f93af935e86047111590c3768271d237e4a2bc0b874bed693.
//
// Solidity: event AVSUpdated(address prevAVS, address newAVS)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorFilterer) FilterAVSUpdated(opts *bind.FilterOpts) (*ContractSlashingRegistryCoordinatorAVSUpdatedIterator, error) {

	logs, sub, err := _ContractSlashingRegistryCoordinator.contract.FilterLogs(opts, "AVSUpdated")
	if err != nil {
		return nil, err
	}
	return &ContractSlashingRegistryCoordinatorAVSUpdatedIterator{contract: _ContractSlashingRegistryCoordinator.contract, event: "AVSUpdated", logs: logs, sub: sub}, nil
}

// WatchAVSUpdated is a free log subscription operation binding the contract event 0x9770f3cadfdcbb6f93af935e86047111590c3768271d237e4a2bc0b874bed693.
//
// Solidity: event AVSUpdated(address prevAVS, address newAVS)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorFilterer) WatchAVSUpdated(opts *bind.WatchOpts, sink chan<- *ContractSlashingRegistryCoordinatorAVSUpdated) (event.Subscription, error) {

	logs, sub, err := _ContractSlashingRegistryCoordinator.contract.WatchLogs(opts, "AVSUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSlashingRegistryCoordinatorAVSUpdated)
				if err := _ContractSlashingRegistryCoordinator.contract.UnpackLog(event, "AVSUpdated", log); err != nil {
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

// ParseAVSUpdated is a log parse operation binding the contract event 0x9770f3cadfdcbb6f93af935e86047111590c3768271d237e4a2bc0b874bed693.
//
// Solidity: event AVSUpdated(address prevAVS, address newAVS)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorFilterer) ParseAVSUpdated(log types.Log) (*ContractSlashingRegistryCoordinatorAVSUpdated, error) {
	event := new(ContractSlashingRegistryCoordinatorAVSUpdated)
	if err := _ContractSlashingRegistryCoordinator.contract.UnpackLog(event, "AVSUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSlashingRegistryCoordinatorChurnApproverUpdatedIterator is returned from FilterChurnApproverUpdated and is used to iterate over the raw logs and unpacked data for ChurnApproverUpdated events raised by the ContractSlashingRegistryCoordinator contract.
type ContractSlashingRegistryCoordinatorChurnApproverUpdatedIterator struct {
	Event *ContractSlashingRegistryCoordinatorChurnApproverUpdated // Event containing the contract specifics and raw log

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
func (it *ContractSlashingRegistryCoordinatorChurnApproverUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSlashingRegistryCoordinatorChurnApproverUpdated)
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
		it.Event = new(ContractSlashingRegistryCoordinatorChurnApproverUpdated)
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
func (it *ContractSlashingRegistryCoordinatorChurnApproverUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSlashingRegistryCoordinatorChurnApproverUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSlashingRegistryCoordinatorChurnApproverUpdated represents a ChurnApproverUpdated event raised by the ContractSlashingRegistryCoordinator contract.
type ContractSlashingRegistryCoordinatorChurnApproverUpdated struct {
	PrevChurnApprover common.Address
	NewChurnApprover  common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterChurnApproverUpdated is a free log retrieval operation binding the contract event 0x315457d8a8fe60f04af17c16e2f5a5e1db612b31648e58030360759ef8f3528c.
//
// Solidity: event ChurnApproverUpdated(address prevChurnApprover, address newChurnApprover)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorFilterer) FilterChurnApproverUpdated(opts *bind.FilterOpts) (*ContractSlashingRegistryCoordinatorChurnApproverUpdatedIterator, error) {

	logs, sub, err := _ContractSlashingRegistryCoordinator.contract.FilterLogs(opts, "ChurnApproverUpdated")
	if err != nil {
		return nil, err
	}
	return &ContractSlashingRegistryCoordinatorChurnApproverUpdatedIterator{contract: _ContractSlashingRegistryCoordinator.contract, event: "ChurnApproverUpdated", logs: logs, sub: sub}, nil
}

// WatchChurnApproverUpdated is a free log subscription operation binding the contract event 0x315457d8a8fe60f04af17c16e2f5a5e1db612b31648e58030360759ef8f3528c.
//
// Solidity: event ChurnApproverUpdated(address prevChurnApprover, address newChurnApprover)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorFilterer) WatchChurnApproverUpdated(opts *bind.WatchOpts, sink chan<- *ContractSlashingRegistryCoordinatorChurnApproverUpdated) (event.Subscription, error) {

	logs, sub, err := _ContractSlashingRegistryCoordinator.contract.WatchLogs(opts, "ChurnApproverUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSlashingRegistryCoordinatorChurnApproverUpdated)
				if err := _ContractSlashingRegistryCoordinator.contract.UnpackLog(event, "ChurnApproverUpdated", log); err != nil {
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

// ParseChurnApproverUpdated is a log parse operation binding the contract event 0x315457d8a8fe60f04af17c16e2f5a5e1db612b31648e58030360759ef8f3528c.
//
// Solidity: event ChurnApproverUpdated(address prevChurnApprover, address newChurnApprover)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorFilterer) ParseChurnApproverUpdated(log types.Log) (*ContractSlashingRegistryCoordinatorChurnApproverUpdated, error) {
	event := new(ContractSlashingRegistryCoordinatorChurnApproverUpdated)
	if err := _ContractSlashingRegistryCoordinator.contract.UnpackLog(event, "ChurnApproverUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSlashingRegistryCoordinatorEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the ContractSlashingRegistryCoordinator contract.
type ContractSlashingRegistryCoordinatorEIP712DomainChangedIterator struct {
	Event *ContractSlashingRegistryCoordinatorEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *ContractSlashingRegistryCoordinatorEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSlashingRegistryCoordinatorEIP712DomainChanged)
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
		it.Event = new(ContractSlashingRegistryCoordinatorEIP712DomainChanged)
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
func (it *ContractSlashingRegistryCoordinatorEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSlashingRegistryCoordinatorEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSlashingRegistryCoordinatorEIP712DomainChanged represents a EIP712DomainChanged event raised by the ContractSlashingRegistryCoordinator contract.
type ContractSlashingRegistryCoordinatorEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*ContractSlashingRegistryCoordinatorEIP712DomainChangedIterator, error) {

	logs, sub, err := _ContractSlashingRegistryCoordinator.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &ContractSlashingRegistryCoordinatorEIP712DomainChangedIterator{contract: _ContractSlashingRegistryCoordinator.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *ContractSlashingRegistryCoordinatorEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _ContractSlashingRegistryCoordinator.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSlashingRegistryCoordinatorEIP712DomainChanged)
				if err := _ContractSlashingRegistryCoordinator.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorFilterer) ParseEIP712DomainChanged(log types.Log) (*ContractSlashingRegistryCoordinatorEIP712DomainChanged, error) {
	event := new(ContractSlashingRegistryCoordinatorEIP712DomainChanged)
	if err := _ContractSlashingRegistryCoordinator.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSlashingRegistryCoordinatorEjectionCooldownUpdatedIterator is returned from FilterEjectionCooldownUpdated and is used to iterate over the raw logs and unpacked data for EjectionCooldownUpdated events raised by the ContractSlashingRegistryCoordinator contract.
type ContractSlashingRegistryCoordinatorEjectionCooldownUpdatedIterator struct {
	Event *ContractSlashingRegistryCoordinatorEjectionCooldownUpdated // Event containing the contract specifics and raw log

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
func (it *ContractSlashingRegistryCoordinatorEjectionCooldownUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSlashingRegistryCoordinatorEjectionCooldownUpdated)
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
		it.Event = new(ContractSlashingRegistryCoordinatorEjectionCooldownUpdated)
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
func (it *ContractSlashingRegistryCoordinatorEjectionCooldownUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSlashingRegistryCoordinatorEjectionCooldownUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSlashingRegistryCoordinatorEjectionCooldownUpdated represents a EjectionCooldownUpdated event raised by the ContractSlashingRegistryCoordinator contract.
type ContractSlashingRegistryCoordinatorEjectionCooldownUpdated struct {
	PrevEjectionCooldown *big.Int
	NewEjectionCooldown  *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterEjectionCooldownUpdated is a free log retrieval operation binding the contract event 0xa77a91bea7b6d95a8eb5a54878a1d9e3c875e26c86a9b70e3420c5c5db193b62.
//
// Solidity: event EjectionCooldownUpdated(uint256 prevEjectionCooldown, uint256 newEjectionCooldown)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorFilterer) FilterEjectionCooldownUpdated(opts *bind.FilterOpts) (*ContractSlashingRegistryCoordinatorEjectionCooldownUpdatedIterator, error) {

	logs, sub, err := _ContractSlashingRegistryCoordinator.contract.FilterLogs(opts, "EjectionCooldownUpdated")
	if err != nil {
		return nil, err
	}
	return &ContractSlashingRegistryCoordinatorEjectionCooldownUpdatedIterator{contract: _ContractSlashingRegistryCoordinator.contract, event: "EjectionCooldownUpdated", logs: logs, sub: sub}, nil
}

// WatchEjectionCooldownUpdated is a free log subscription operation binding the contract event 0xa77a91bea7b6d95a8eb5a54878a1d9e3c875e26c86a9b70e3420c5c5db193b62.
//
// Solidity: event EjectionCooldownUpdated(uint256 prevEjectionCooldown, uint256 newEjectionCooldown)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorFilterer) WatchEjectionCooldownUpdated(opts *bind.WatchOpts, sink chan<- *ContractSlashingRegistryCoordinatorEjectionCooldownUpdated) (event.Subscription, error) {

	logs, sub, err := _ContractSlashingRegistryCoordinator.contract.WatchLogs(opts, "EjectionCooldownUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSlashingRegistryCoordinatorEjectionCooldownUpdated)
				if err := _ContractSlashingRegistryCoordinator.contract.UnpackLog(event, "EjectionCooldownUpdated", log); err != nil {
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

// ParseEjectionCooldownUpdated is a log parse operation binding the contract event 0xa77a91bea7b6d95a8eb5a54878a1d9e3c875e26c86a9b70e3420c5c5db193b62.
//
// Solidity: event EjectionCooldownUpdated(uint256 prevEjectionCooldown, uint256 newEjectionCooldown)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorFilterer) ParseEjectionCooldownUpdated(log types.Log) (*ContractSlashingRegistryCoordinatorEjectionCooldownUpdated, error) {
	event := new(ContractSlashingRegistryCoordinatorEjectionCooldownUpdated)
	if err := _ContractSlashingRegistryCoordinator.contract.UnpackLog(event, "EjectionCooldownUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSlashingRegistryCoordinatorEjectorUpdatedIterator is returned from FilterEjectorUpdated and is used to iterate over the raw logs and unpacked data for EjectorUpdated events raised by the ContractSlashingRegistryCoordinator contract.
type ContractSlashingRegistryCoordinatorEjectorUpdatedIterator struct {
	Event *ContractSlashingRegistryCoordinatorEjectorUpdated // Event containing the contract specifics and raw log

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
func (it *ContractSlashingRegistryCoordinatorEjectorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSlashingRegistryCoordinatorEjectorUpdated)
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
		it.Event = new(ContractSlashingRegistryCoordinatorEjectorUpdated)
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
func (it *ContractSlashingRegistryCoordinatorEjectorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSlashingRegistryCoordinatorEjectorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSlashingRegistryCoordinatorEjectorUpdated represents a EjectorUpdated event raised by the ContractSlashingRegistryCoordinator contract.
type ContractSlashingRegistryCoordinatorEjectorUpdated struct {
	PrevEjector common.Address
	NewEjector  common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterEjectorUpdated is a free log retrieval operation binding the contract event 0x8f30ab09f43a6c157d7fce7e0a13c003042c1c95e8a72e7a146a21c0caa24dc9.
//
// Solidity: event EjectorUpdated(address prevEjector, address newEjector)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorFilterer) FilterEjectorUpdated(opts *bind.FilterOpts) (*ContractSlashingRegistryCoordinatorEjectorUpdatedIterator, error) {

	logs, sub, err := _ContractSlashingRegistryCoordinator.contract.FilterLogs(opts, "EjectorUpdated")
	if err != nil {
		return nil, err
	}
	return &ContractSlashingRegistryCoordinatorEjectorUpdatedIterator{contract: _ContractSlashingRegistryCoordinator.contract, event: "EjectorUpdated", logs: logs, sub: sub}, nil
}

// WatchEjectorUpdated is a free log subscription operation binding the contract event 0x8f30ab09f43a6c157d7fce7e0a13c003042c1c95e8a72e7a146a21c0caa24dc9.
//
// Solidity: event EjectorUpdated(address prevEjector, address newEjector)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorFilterer) WatchEjectorUpdated(opts *bind.WatchOpts, sink chan<- *ContractSlashingRegistryCoordinatorEjectorUpdated) (event.Subscription, error) {

	logs, sub, err := _ContractSlashingRegistryCoordinator.contract.WatchLogs(opts, "EjectorUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSlashingRegistryCoordinatorEjectorUpdated)
				if err := _ContractSlashingRegistryCoordinator.contract.UnpackLog(event, "EjectorUpdated", log); err != nil {
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

// ParseEjectorUpdated is a log parse operation binding the contract event 0x8f30ab09f43a6c157d7fce7e0a13c003042c1c95e8a72e7a146a21c0caa24dc9.
//
// Solidity: event EjectorUpdated(address prevEjector, address newEjector)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorFilterer) ParseEjectorUpdated(log types.Log) (*ContractSlashingRegistryCoordinatorEjectorUpdated, error) {
	event := new(ContractSlashingRegistryCoordinatorEjectorUpdated)
	if err := _ContractSlashingRegistryCoordinator.contract.UnpackLog(event, "EjectorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSlashingRegistryCoordinatorInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractSlashingRegistryCoordinator contract.
type ContractSlashingRegistryCoordinatorInitializedIterator struct {
	Event *ContractSlashingRegistryCoordinatorInitialized // Event containing the contract specifics and raw log

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
func (it *ContractSlashingRegistryCoordinatorInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSlashingRegistryCoordinatorInitialized)
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
		it.Event = new(ContractSlashingRegistryCoordinatorInitialized)
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
func (it *ContractSlashingRegistryCoordinatorInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSlashingRegistryCoordinatorInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSlashingRegistryCoordinatorInitialized represents a Initialized event raised by the ContractSlashingRegistryCoordinator contract.
type ContractSlashingRegistryCoordinatorInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractSlashingRegistryCoordinatorInitializedIterator, error) {

	logs, sub, err := _ContractSlashingRegistryCoordinator.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractSlashingRegistryCoordinatorInitializedIterator{contract: _ContractSlashingRegistryCoordinator.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractSlashingRegistryCoordinatorInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractSlashingRegistryCoordinator.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSlashingRegistryCoordinatorInitialized)
				if err := _ContractSlashingRegistryCoordinator.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorFilterer) ParseInitialized(log types.Log) (*ContractSlashingRegistryCoordinatorInitialized, error) {
	event := new(ContractSlashingRegistryCoordinatorInitialized)
	if err := _ContractSlashingRegistryCoordinator.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSlashingRegistryCoordinatorOperatorDeregisteredIterator is returned from FilterOperatorDeregistered and is used to iterate over the raw logs and unpacked data for OperatorDeregistered events raised by the ContractSlashingRegistryCoordinator contract.
type ContractSlashingRegistryCoordinatorOperatorDeregisteredIterator struct {
	Event *ContractSlashingRegistryCoordinatorOperatorDeregistered // Event containing the contract specifics and raw log

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
func (it *ContractSlashingRegistryCoordinatorOperatorDeregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSlashingRegistryCoordinatorOperatorDeregistered)
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
		it.Event = new(ContractSlashingRegistryCoordinatorOperatorDeregistered)
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
func (it *ContractSlashingRegistryCoordinatorOperatorDeregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSlashingRegistryCoordinatorOperatorDeregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSlashingRegistryCoordinatorOperatorDeregistered represents a OperatorDeregistered event raised by the ContractSlashingRegistryCoordinator contract.
type ContractSlashingRegistryCoordinatorOperatorDeregistered struct {
	Operator   common.Address
	OperatorId [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOperatorDeregistered is a free log retrieval operation binding the contract event 0x396fdcb180cb0fea26928113fb0fd1c3549863f9cd563e6a184f1d578116c8e4.
//
// Solidity: event OperatorDeregistered(address indexed operator, bytes32 indexed operatorId)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorFilterer) FilterOperatorDeregistered(opts *bind.FilterOpts, operator []common.Address, operatorId [][32]byte) (*ContractSlashingRegistryCoordinatorOperatorDeregisteredIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _ContractSlashingRegistryCoordinator.contract.FilterLogs(opts, "OperatorDeregistered", operatorRule, operatorIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractSlashingRegistryCoordinatorOperatorDeregisteredIterator{contract: _ContractSlashingRegistryCoordinator.contract, event: "OperatorDeregistered", logs: logs, sub: sub}, nil
}

// WatchOperatorDeregistered is a free log subscription operation binding the contract event 0x396fdcb180cb0fea26928113fb0fd1c3549863f9cd563e6a184f1d578116c8e4.
//
// Solidity: event OperatorDeregistered(address indexed operator, bytes32 indexed operatorId)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorFilterer) WatchOperatorDeregistered(opts *bind.WatchOpts, sink chan<- *ContractSlashingRegistryCoordinatorOperatorDeregistered, operator []common.Address, operatorId [][32]byte) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _ContractSlashingRegistryCoordinator.contract.WatchLogs(opts, "OperatorDeregistered", operatorRule, operatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSlashingRegistryCoordinatorOperatorDeregistered)
				if err := _ContractSlashingRegistryCoordinator.contract.UnpackLog(event, "OperatorDeregistered", log); err != nil {
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

// ParseOperatorDeregistered is a log parse operation binding the contract event 0x396fdcb180cb0fea26928113fb0fd1c3549863f9cd563e6a184f1d578116c8e4.
//
// Solidity: event OperatorDeregistered(address indexed operator, bytes32 indexed operatorId)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorFilterer) ParseOperatorDeregistered(log types.Log) (*ContractSlashingRegistryCoordinatorOperatorDeregistered, error) {
	event := new(ContractSlashingRegistryCoordinatorOperatorDeregistered)
	if err := _ContractSlashingRegistryCoordinator.contract.UnpackLog(event, "OperatorDeregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSlashingRegistryCoordinatorOperatorRegisteredIterator is returned from FilterOperatorRegistered and is used to iterate over the raw logs and unpacked data for OperatorRegistered events raised by the ContractSlashingRegistryCoordinator contract.
type ContractSlashingRegistryCoordinatorOperatorRegisteredIterator struct {
	Event *ContractSlashingRegistryCoordinatorOperatorRegistered // Event containing the contract specifics and raw log

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
func (it *ContractSlashingRegistryCoordinatorOperatorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSlashingRegistryCoordinatorOperatorRegistered)
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
		it.Event = new(ContractSlashingRegistryCoordinatorOperatorRegistered)
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
func (it *ContractSlashingRegistryCoordinatorOperatorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSlashingRegistryCoordinatorOperatorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSlashingRegistryCoordinatorOperatorRegistered represents a OperatorRegistered event raised by the ContractSlashingRegistryCoordinator contract.
type ContractSlashingRegistryCoordinatorOperatorRegistered struct {
	Operator   common.Address
	OperatorId [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOperatorRegistered is a free log retrieval operation binding the contract event 0xe8e68cef1c3a761ed7be7e8463a375f27f7bc335e51824223cacce636ec5c3fe.
//
// Solidity: event OperatorRegistered(address indexed operator, bytes32 indexed operatorId)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorFilterer) FilterOperatorRegistered(opts *bind.FilterOpts, operator []common.Address, operatorId [][32]byte) (*ContractSlashingRegistryCoordinatorOperatorRegisteredIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _ContractSlashingRegistryCoordinator.contract.FilterLogs(opts, "OperatorRegistered", operatorRule, operatorIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractSlashingRegistryCoordinatorOperatorRegisteredIterator{contract: _ContractSlashingRegistryCoordinator.contract, event: "OperatorRegistered", logs: logs, sub: sub}, nil
}

// WatchOperatorRegistered is a free log subscription operation binding the contract event 0xe8e68cef1c3a761ed7be7e8463a375f27f7bc335e51824223cacce636ec5c3fe.
//
// Solidity: event OperatorRegistered(address indexed operator, bytes32 indexed operatorId)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorFilterer) WatchOperatorRegistered(opts *bind.WatchOpts, sink chan<- *ContractSlashingRegistryCoordinatorOperatorRegistered, operator []common.Address, operatorId [][32]byte) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _ContractSlashingRegistryCoordinator.contract.WatchLogs(opts, "OperatorRegistered", operatorRule, operatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSlashingRegistryCoordinatorOperatorRegistered)
				if err := _ContractSlashingRegistryCoordinator.contract.UnpackLog(event, "OperatorRegistered", log); err != nil {
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

// ParseOperatorRegistered is a log parse operation binding the contract event 0xe8e68cef1c3a761ed7be7e8463a375f27f7bc335e51824223cacce636ec5c3fe.
//
// Solidity: event OperatorRegistered(address indexed operator, bytes32 indexed operatorId)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorFilterer) ParseOperatorRegistered(log types.Log) (*ContractSlashingRegistryCoordinatorOperatorRegistered, error) {
	event := new(ContractSlashingRegistryCoordinatorOperatorRegistered)
	if err := _ContractSlashingRegistryCoordinator.contract.UnpackLog(event, "OperatorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSlashingRegistryCoordinatorOperatorSetParamsUpdatedIterator is returned from FilterOperatorSetParamsUpdated and is used to iterate over the raw logs and unpacked data for OperatorSetParamsUpdated events raised by the ContractSlashingRegistryCoordinator contract.
type ContractSlashingRegistryCoordinatorOperatorSetParamsUpdatedIterator struct {
	Event *ContractSlashingRegistryCoordinatorOperatorSetParamsUpdated // Event containing the contract specifics and raw log

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
func (it *ContractSlashingRegistryCoordinatorOperatorSetParamsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSlashingRegistryCoordinatorOperatorSetParamsUpdated)
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
		it.Event = new(ContractSlashingRegistryCoordinatorOperatorSetParamsUpdated)
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
func (it *ContractSlashingRegistryCoordinatorOperatorSetParamsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSlashingRegistryCoordinatorOperatorSetParamsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSlashingRegistryCoordinatorOperatorSetParamsUpdated represents a OperatorSetParamsUpdated event raised by the ContractSlashingRegistryCoordinator contract.
type ContractSlashingRegistryCoordinatorOperatorSetParamsUpdated struct {
	QuorumNumber      uint8
	OperatorSetParams ISlashingRegistryCoordinatorTypesOperatorSetParam
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterOperatorSetParamsUpdated is a free log retrieval operation binding the contract event 0x3ee6fe8d54610244c3e9d3c066ae4aee997884aa28f10616ae821925401318ac.
//
// Solidity: event OperatorSetParamsUpdated(uint8 indexed quorumNumber, (uint32,uint16,uint16) operatorSetParams)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorFilterer) FilterOperatorSetParamsUpdated(opts *bind.FilterOpts, quorumNumber []uint8) (*ContractSlashingRegistryCoordinatorOperatorSetParamsUpdatedIterator, error) {

	var quorumNumberRule []interface{}
	for _, quorumNumberItem := range quorumNumber {
		quorumNumberRule = append(quorumNumberRule, quorumNumberItem)
	}

	logs, sub, err := _ContractSlashingRegistryCoordinator.contract.FilterLogs(opts, "OperatorSetParamsUpdated", quorumNumberRule)
	if err != nil {
		return nil, err
	}
	return &ContractSlashingRegistryCoordinatorOperatorSetParamsUpdatedIterator{contract: _ContractSlashingRegistryCoordinator.contract, event: "OperatorSetParamsUpdated", logs: logs, sub: sub}, nil
}

// WatchOperatorSetParamsUpdated is a free log subscription operation binding the contract event 0x3ee6fe8d54610244c3e9d3c066ae4aee997884aa28f10616ae821925401318ac.
//
// Solidity: event OperatorSetParamsUpdated(uint8 indexed quorumNumber, (uint32,uint16,uint16) operatorSetParams)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorFilterer) WatchOperatorSetParamsUpdated(opts *bind.WatchOpts, sink chan<- *ContractSlashingRegistryCoordinatorOperatorSetParamsUpdated, quorumNumber []uint8) (event.Subscription, error) {

	var quorumNumberRule []interface{}
	for _, quorumNumberItem := range quorumNumber {
		quorumNumberRule = append(quorumNumberRule, quorumNumberItem)
	}

	logs, sub, err := _ContractSlashingRegistryCoordinator.contract.WatchLogs(opts, "OperatorSetParamsUpdated", quorumNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSlashingRegistryCoordinatorOperatorSetParamsUpdated)
				if err := _ContractSlashingRegistryCoordinator.contract.UnpackLog(event, "OperatorSetParamsUpdated", log); err != nil {
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

// ParseOperatorSetParamsUpdated is a log parse operation binding the contract event 0x3ee6fe8d54610244c3e9d3c066ae4aee997884aa28f10616ae821925401318ac.
//
// Solidity: event OperatorSetParamsUpdated(uint8 indexed quorumNumber, (uint32,uint16,uint16) operatorSetParams)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorFilterer) ParseOperatorSetParamsUpdated(log types.Log) (*ContractSlashingRegistryCoordinatorOperatorSetParamsUpdated, error) {
	event := new(ContractSlashingRegistryCoordinatorOperatorSetParamsUpdated)
	if err := _ContractSlashingRegistryCoordinator.contract.UnpackLog(event, "OperatorSetParamsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSlashingRegistryCoordinatorOperatorSocketUpdateIterator is returned from FilterOperatorSocketUpdate and is used to iterate over the raw logs and unpacked data for OperatorSocketUpdate events raised by the ContractSlashingRegistryCoordinator contract.
type ContractSlashingRegistryCoordinatorOperatorSocketUpdateIterator struct {
	Event *ContractSlashingRegistryCoordinatorOperatorSocketUpdate // Event containing the contract specifics and raw log

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
func (it *ContractSlashingRegistryCoordinatorOperatorSocketUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSlashingRegistryCoordinatorOperatorSocketUpdate)
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
		it.Event = new(ContractSlashingRegistryCoordinatorOperatorSocketUpdate)
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
func (it *ContractSlashingRegistryCoordinatorOperatorSocketUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSlashingRegistryCoordinatorOperatorSocketUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSlashingRegistryCoordinatorOperatorSocketUpdate represents a OperatorSocketUpdate event raised by the ContractSlashingRegistryCoordinator contract.
type ContractSlashingRegistryCoordinatorOperatorSocketUpdate struct {
	OperatorId [32]byte
	Socket     string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOperatorSocketUpdate is a free log retrieval operation binding the contract event 0xec2963ab21c1e50e1e582aa542af2e4bf7bf38e6e1403c27b42e1c5d6e621eaa.
//
// Solidity: event OperatorSocketUpdate(bytes32 indexed operatorId, string socket)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorFilterer) FilterOperatorSocketUpdate(opts *bind.FilterOpts, operatorId [][32]byte) (*ContractSlashingRegistryCoordinatorOperatorSocketUpdateIterator, error) {

	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _ContractSlashingRegistryCoordinator.contract.FilterLogs(opts, "OperatorSocketUpdate", operatorIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractSlashingRegistryCoordinatorOperatorSocketUpdateIterator{contract: _ContractSlashingRegistryCoordinator.contract, event: "OperatorSocketUpdate", logs: logs, sub: sub}, nil
}

// WatchOperatorSocketUpdate is a free log subscription operation binding the contract event 0xec2963ab21c1e50e1e582aa542af2e4bf7bf38e6e1403c27b42e1c5d6e621eaa.
//
// Solidity: event OperatorSocketUpdate(bytes32 indexed operatorId, string socket)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorFilterer) WatchOperatorSocketUpdate(opts *bind.WatchOpts, sink chan<- *ContractSlashingRegistryCoordinatorOperatorSocketUpdate, operatorId [][32]byte) (event.Subscription, error) {

	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _ContractSlashingRegistryCoordinator.contract.WatchLogs(opts, "OperatorSocketUpdate", operatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSlashingRegistryCoordinatorOperatorSocketUpdate)
				if err := _ContractSlashingRegistryCoordinator.contract.UnpackLog(event, "OperatorSocketUpdate", log); err != nil {
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

// ParseOperatorSocketUpdate is a log parse operation binding the contract event 0xec2963ab21c1e50e1e582aa542af2e4bf7bf38e6e1403c27b42e1c5d6e621eaa.
//
// Solidity: event OperatorSocketUpdate(bytes32 indexed operatorId, string socket)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorFilterer) ParseOperatorSocketUpdate(log types.Log) (*ContractSlashingRegistryCoordinatorOperatorSocketUpdate, error) {
	event := new(ContractSlashingRegistryCoordinatorOperatorSocketUpdate)
	if err := _ContractSlashingRegistryCoordinator.contract.UnpackLog(event, "OperatorSocketUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSlashingRegistryCoordinatorOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContractSlashingRegistryCoordinator contract.
type ContractSlashingRegistryCoordinatorOwnershipTransferredIterator struct {
	Event *ContractSlashingRegistryCoordinatorOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractSlashingRegistryCoordinatorOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSlashingRegistryCoordinatorOwnershipTransferred)
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
		it.Event = new(ContractSlashingRegistryCoordinatorOwnershipTransferred)
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
func (it *ContractSlashingRegistryCoordinatorOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSlashingRegistryCoordinatorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSlashingRegistryCoordinatorOwnershipTransferred represents a OwnershipTransferred event raised by the ContractSlashingRegistryCoordinator contract.
type ContractSlashingRegistryCoordinatorOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractSlashingRegistryCoordinatorOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractSlashingRegistryCoordinator.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractSlashingRegistryCoordinatorOwnershipTransferredIterator{contract: _ContractSlashingRegistryCoordinator.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractSlashingRegistryCoordinatorOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractSlashingRegistryCoordinator.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSlashingRegistryCoordinatorOwnershipTransferred)
				if err := _ContractSlashingRegistryCoordinator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorFilterer) ParseOwnershipTransferred(log types.Log) (*ContractSlashingRegistryCoordinatorOwnershipTransferred, error) {
	event := new(ContractSlashingRegistryCoordinatorOwnershipTransferred)
	if err := _ContractSlashingRegistryCoordinator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSlashingRegistryCoordinatorPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ContractSlashingRegistryCoordinator contract.
type ContractSlashingRegistryCoordinatorPausedIterator struct {
	Event *ContractSlashingRegistryCoordinatorPaused // Event containing the contract specifics and raw log

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
func (it *ContractSlashingRegistryCoordinatorPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSlashingRegistryCoordinatorPaused)
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
		it.Event = new(ContractSlashingRegistryCoordinatorPaused)
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
func (it *ContractSlashingRegistryCoordinatorPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSlashingRegistryCoordinatorPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSlashingRegistryCoordinatorPaused represents a Paused event raised by the ContractSlashingRegistryCoordinator contract.
type ContractSlashingRegistryCoordinatorPaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*ContractSlashingRegistryCoordinatorPausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractSlashingRegistryCoordinator.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractSlashingRegistryCoordinatorPausedIterator{contract: _ContractSlashingRegistryCoordinator.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ContractSlashingRegistryCoordinatorPaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractSlashingRegistryCoordinator.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSlashingRegistryCoordinatorPaused)
				if err := _ContractSlashingRegistryCoordinator.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorFilterer) ParsePaused(log types.Log) (*ContractSlashingRegistryCoordinatorPaused, error) {
	event := new(ContractSlashingRegistryCoordinatorPaused)
	if err := _ContractSlashingRegistryCoordinator.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSlashingRegistryCoordinatorQuorumBlockNumberUpdatedIterator is returned from FilterQuorumBlockNumberUpdated and is used to iterate over the raw logs and unpacked data for QuorumBlockNumberUpdated events raised by the ContractSlashingRegistryCoordinator contract.
type ContractSlashingRegistryCoordinatorQuorumBlockNumberUpdatedIterator struct {
	Event *ContractSlashingRegistryCoordinatorQuorumBlockNumberUpdated // Event containing the contract specifics and raw log

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
func (it *ContractSlashingRegistryCoordinatorQuorumBlockNumberUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSlashingRegistryCoordinatorQuorumBlockNumberUpdated)
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
		it.Event = new(ContractSlashingRegistryCoordinatorQuorumBlockNumberUpdated)
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
func (it *ContractSlashingRegistryCoordinatorQuorumBlockNumberUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSlashingRegistryCoordinatorQuorumBlockNumberUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSlashingRegistryCoordinatorQuorumBlockNumberUpdated represents a QuorumBlockNumberUpdated event raised by the ContractSlashingRegistryCoordinator contract.
type ContractSlashingRegistryCoordinatorQuorumBlockNumberUpdated struct {
	QuorumNumber uint8
	Blocknumber  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterQuorumBlockNumberUpdated is a free log retrieval operation binding the contract event 0x46077d55330763f16269fd75e5761663f4192d2791747c0189b16ad31db07db4.
//
// Solidity: event QuorumBlockNumberUpdated(uint8 indexed quorumNumber, uint256 blocknumber)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorFilterer) FilterQuorumBlockNumberUpdated(opts *bind.FilterOpts, quorumNumber []uint8) (*ContractSlashingRegistryCoordinatorQuorumBlockNumberUpdatedIterator, error) {

	var quorumNumberRule []interface{}
	for _, quorumNumberItem := range quorumNumber {
		quorumNumberRule = append(quorumNumberRule, quorumNumberItem)
	}

	logs, sub, err := _ContractSlashingRegistryCoordinator.contract.FilterLogs(opts, "QuorumBlockNumberUpdated", quorumNumberRule)
	if err != nil {
		return nil, err
	}
	return &ContractSlashingRegistryCoordinatorQuorumBlockNumberUpdatedIterator{contract: _ContractSlashingRegistryCoordinator.contract, event: "QuorumBlockNumberUpdated", logs: logs, sub: sub}, nil
}

// WatchQuorumBlockNumberUpdated is a free log subscription operation binding the contract event 0x46077d55330763f16269fd75e5761663f4192d2791747c0189b16ad31db07db4.
//
// Solidity: event QuorumBlockNumberUpdated(uint8 indexed quorumNumber, uint256 blocknumber)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorFilterer) WatchQuorumBlockNumberUpdated(opts *bind.WatchOpts, sink chan<- *ContractSlashingRegistryCoordinatorQuorumBlockNumberUpdated, quorumNumber []uint8) (event.Subscription, error) {

	var quorumNumberRule []interface{}
	for _, quorumNumberItem := range quorumNumber {
		quorumNumberRule = append(quorumNumberRule, quorumNumberItem)
	}

	logs, sub, err := _ContractSlashingRegistryCoordinator.contract.WatchLogs(opts, "QuorumBlockNumberUpdated", quorumNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSlashingRegistryCoordinatorQuorumBlockNumberUpdated)
				if err := _ContractSlashingRegistryCoordinator.contract.UnpackLog(event, "QuorumBlockNumberUpdated", log); err != nil {
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

// ParseQuorumBlockNumberUpdated is a log parse operation binding the contract event 0x46077d55330763f16269fd75e5761663f4192d2791747c0189b16ad31db07db4.
//
// Solidity: event QuorumBlockNumberUpdated(uint8 indexed quorumNumber, uint256 blocknumber)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorFilterer) ParseQuorumBlockNumberUpdated(log types.Log) (*ContractSlashingRegistryCoordinatorQuorumBlockNumberUpdated, error) {
	event := new(ContractSlashingRegistryCoordinatorQuorumBlockNumberUpdated)
	if err := _ContractSlashingRegistryCoordinator.contract.UnpackLog(event, "QuorumBlockNumberUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSlashingRegistryCoordinatorQuorumCreatedIterator is returned from FilterQuorumCreated and is used to iterate over the raw logs and unpacked data for QuorumCreated events raised by the ContractSlashingRegistryCoordinator contract.
type ContractSlashingRegistryCoordinatorQuorumCreatedIterator struct {
	Event *ContractSlashingRegistryCoordinatorQuorumCreated // Event containing the contract specifics and raw log

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
func (it *ContractSlashingRegistryCoordinatorQuorumCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSlashingRegistryCoordinatorQuorumCreated)
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
		it.Event = new(ContractSlashingRegistryCoordinatorQuorumCreated)
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
func (it *ContractSlashingRegistryCoordinatorQuorumCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSlashingRegistryCoordinatorQuorumCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSlashingRegistryCoordinatorQuorumCreated represents a QuorumCreated event raised by the ContractSlashingRegistryCoordinator contract.
type ContractSlashingRegistryCoordinatorQuorumCreated struct {
	QuorumNumber      uint8
	OperatorSetParams ISlashingRegistryCoordinatorTypesOperatorSetParam
	MinimumStake      *big.Int
	StrategyParams    []IStakeRegistryTypesStrategyParams
	StakeType         uint8
	LookAheadPeriod   uint32
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterQuorumCreated is a free log retrieval operation binding the contract event 0xa34835bc2b673ec37fcf1591a91295b163fc2e181e4ea4e733beb27de1ceac4c.
//
// Solidity: event QuorumCreated(uint8 indexed quorumNumber, (uint32,uint16,uint16) operatorSetParams, uint96 minimumStake, (address,uint96)[] strategyParams, uint8 stakeType, uint32 lookAheadPeriod)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorFilterer) FilterQuorumCreated(opts *bind.FilterOpts, quorumNumber []uint8) (*ContractSlashingRegistryCoordinatorQuorumCreatedIterator, error) {

	var quorumNumberRule []interface{}
	for _, quorumNumberItem := range quorumNumber {
		quorumNumberRule = append(quorumNumberRule, quorumNumberItem)
	}

	logs, sub, err := _ContractSlashingRegistryCoordinator.contract.FilterLogs(opts, "QuorumCreated", quorumNumberRule)
	if err != nil {
		return nil, err
	}
	return &ContractSlashingRegistryCoordinatorQuorumCreatedIterator{contract: _ContractSlashingRegistryCoordinator.contract, event: "QuorumCreated", logs: logs, sub: sub}, nil
}

// WatchQuorumCreated is a free log subscription operation binding the contract event 0xa34835bc2b673ec37fcf1591a91295b163fc2e181e4ea4e733beb27de1ceac4c.
//
// Solidity: event QuorumCreated(uint8 indexed quorumNumber, (uint32,uint16,uint16) operatorSetParams, uint96 minimumStake, (address,uint96)[] strategyParams, uint8 stakeType, uint32 lookAheadPeriod)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorFilterer) WatchQuorumCreated(opts *bind.WatchOpts, sink chan<- *ContractSlashingRegistryCoordinatorQuorumCreated, quorumNumber []uint8) (event.Subscription, error) {

	var quorumNumberRule []interface{}
	for _, quorumNumberItem := range quorumNumber {
		quorumNumberRule = append(quorumNumberRule, quorumNumberItem)
	}

	logs, sub, err := _ContractSlashingRegistryCoordinator.contract.WatchLogs(opts, "QuorumCreated", quorumNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSlashingRegistryCoordinatorQuorumCreated)
				if err := _ContractSlashingRegistryCoordinator.contract.UnpackLog(event, "QuorumCreated", log); err != nil {
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

// ParseQuorumCreated is a log parse operation binding the contract event 0xa34835bc2b673ec37fcf1591a91295b163fc2e181e4ea4e733beb27de1ceac4c.
//
// Solidity: event QuorumCreated(uint8 indexed quorumNumber, (uint32,uint16,uint16) operatorSetParams, uint96 minimumStake, (address,uint96)[] strategyParams, uint8 stakeType, uint32 lookAheadPeriod)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorFilterer) ParseQuorumCreated(log types.Log) (*ContractSlashingRegistryCoordinatorQuorumCreated, error) {
	event := new(ContractSlashingRegistryCoordinatorQuorumCreated)
	if err := _ContractSlashingRegistryCoordinator.contract.UnpackLog(event, "QuorumCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSlashingRegistryCoordinatorUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ContractSlashingRegistryCoordinator contract.
type ContractSlashingRegistryCoordinatorUnpausedIterator struct {
	Event *ContractSlashingRegistryCoordinatorUnpaused // Event containing the contract specifics and raw log

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
func (it *ContractSlashingRegistryCoordinatorUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSlashingRegistryCoordinatorUnpaused)
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
		it.Event = new(ContractSlashingRegistryCoordinatorUnpaused)
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
func (it *ContractSlashingRegistryCoordinatorUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSlashingRegistryCoordinatorUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSlashingRegistryCoordinatorUnpaused represents a Unpaused event raised by the ContractSlashingRegistryCoordinator contract.
type ContractSlashingRegistryCoordinatorUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*ContractSlashingRegistryCoordinatorUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractSlashingRegistryCoordinator.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractSlashingRegistryCoordinatorUnpausedIterator{contract: _ContractSlashingRegistryCoordinator.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ContractSlashingRegistryCoordinatorUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractSlashingRegistryCoordinator.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSlashingRegistryCoordinatorUnpaused)
				if err := _ContractSlashingRegistryCoordinator.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorFilterer) ParseUnpaused(log types.Log) (*ContractSlashingRegistryCoordinatorUnpaused, error) {
	event := new(ContractSlashingRegistryCoordinatorUnpaused)
	if err := _ContractSlashingRegistryCoordinator.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
