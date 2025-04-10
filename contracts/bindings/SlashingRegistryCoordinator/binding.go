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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_stakeRegistry\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"},{\"name\":\"_blsApkRegistry\",\"type\":\"address\",\"internalType\":\"contractIBLSApkRegistry\"},{\"name\":\"_indexRegistry\",\"type\":\"address\",\"internalType\":\"contractIIndexRegistry\"},{\"name\":\"_socketRegistry\",\"type\":\"address\",\"internalType\":\"contractISocketRegistry\"},{\"name\":\"_allocationManager\",\"type\":\"address\",\"internalType\":\"contractIAllocationManager\"},{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"_version\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"OPERATOR_CHURN_APPROVAL_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PUBKEY_REGISTRATION_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allocationManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIAllocationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"avs\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"blsApkRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBLSApkRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateOperatorChurnApprovalDigestHash\",\"inputs\":[{\"name\":\"registeringOperator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"registeringOperatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"operatorKickParams\",\"type\":\"tuple[]\",\"internalType\":\"structISlashingRegistryCoordinatorTypes.OperatorKickParam[]\",\"components\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculatePubkeyRegistrationMessageHash\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"churnApprover\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createSlashableStakeQuorum\",\"inputs\":[{\"name\":\"operatorSetParams\",\"type\":\"tuple\",\"internalType\":\"structISlashingRegistryCoordinatorTypes.OperatorSetParam\",\"components\":[{\"name\":\"maxOperatorCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"kickBIPsOfOperatorStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"kickBIPsOfTotalStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"name\":\"minimumStake\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"strategyParams\",\"type\":\"tuple[]\",\"internalType\":\"structIStakeRegistryTypes.StrategyParams[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"lookAheadPeriod\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createTotalDelegatedStakeQuorum\",\"inputs\":[{\"name\":\"operatorSetParams\",\"type\":\"tuple\",\"internalType\":\"structISlashingRegistryCoordinatorTypes.OperatorSetParam\",\"components\":[{\"name\":\"maxOperatorCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"kickBIPsOfOperatorStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"kickBIPsOfTotalStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"name\":\"minimumStake\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"strategyParams\",\"type\":\"tuple[]\",\"internalType\":\"structIStakeRegistryTypes.StrategyParams[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deregisterOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"domainSeparator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"eip712Domain\",\"inputs\":[],\"outputs\":[{\"name\":\"fields\",\"type\":\"bytes1\",\"internalType\":\"bytes1\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"verifyingContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"extensions\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ejectOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ejectionCooldown\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ejector\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentQuorumBitmap\",\"inputs\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint192\",\"internalType\":\"uint192\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structISlashingRegistryCoordinatorTypes.OperatorInfo\",\"components\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumISlashingRegistryCoordinatorTypes.OperatorStatus\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorFromId\",\"inputs\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorId\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorSetParams\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structISlashingRegistryCoordinatorTypes.OperatorSetParam\",\"components\":[{\"name\":\"maxOperatorCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"kickBIPsOfOperatorStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"kickBIPsOfTotalStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorStatus\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumISlashingRegistryCoordinatorTypes.OperatorStatus\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumBitmapAtBlockNumberByIndex\",\"inputs\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint192\",\"internalType\":\"uint192\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumBitmapHistoryLength\",\"inputs\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumBitmapIndicesAtBlockNumber\",\"inputs\":[{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumBitmapUpdateByIndex\",\"inputs\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structISlashingRegistryCoordinatorTypes.QuorumBitmapUpdate\",\"components\":[{\"name\":\"updateBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"nextUpdateBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumBitmap\",\"type\":\"uint192\",\"internalType\":\"uint192\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"indexRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIIndexRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"churnApprover\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"ejector\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isChurnApproverSaltUsed\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lastEjectionTimestamp\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pubkeyRegistrationMessageHash\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"quorumCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"quorumUpdateBlockNumber\",\"inputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setAVS\",\"inputs\":[{\"name\":\"_avs\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChurnApprover\",\"inputs\":[{\"name\":\"_churnApprover\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setEjectionCooldown\",\"inputs\":[{\"name\":\"_ejectionCooldown\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setEjector\",\"inputs\":[{\"name\":\"_ejector\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setOperatorSetParams\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"operatorSetParams\",\"type\":\"tuple\",\"internalType\":\"structISlashingRegistryCoordinatorTypes.OperatorSetParam\",\"components\":[{\"name\":\"maxOperatorCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"kickBIPsOfOperatorStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"kickBIPsOfTotalStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"socketRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractISocketRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stakeRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"supportsAVS\",\"inputs\":[{\"name\":\"_avs\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateOperators\",\"inputs\":[{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateOperatorsForQuorum\",\"inputs\":[{\"name\":\"operatorsPerQuorum\",\"type\":\"address[][]\",\"internalType\":\"address[][]\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateSocket\",\"inputs\":[{\"name\":\"socket\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"AVSUpdated\",\"inputs\":[{\"name\":\"prevAVS\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newAVS\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChurnApproverUpdated\",\"inputs\":[{\"name\":\"prevChurnApprover\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newChurnApprover\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EIP712DomainChanged\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EjectionCooldownUpdated\",\"inputs\":[{\"name\":\"prevEjectionCooldown\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newEjectionCooldown\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EjectorUpdated\",\"inputs\":[{\"name\":\"prevEjector\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newEjector\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorDeregistered\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorRegistered\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSetParamsUpdated\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"uint8\"},{\"name\":\"operatorSetParams\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structISlashingRegistryCoordinatorTypes.OperatorSetParam\",\"components\":[{\"name\":\"maxOperatorCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"kickBIPsOfOperatorStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"kickBIPsOfTotalStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSocketUpdate\",\"inputs\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"socket\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"QuorumBlockNumberUpdated\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"uint8\"},{\"name\":\"blocknumber\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"QuorumCreated\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"uint8\"},{\"name\":\"operatorSetParams\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structISlashingRegistryCoordinatorTypes.OperatorSetParam\",\"components\":[{\"name\":\"maxOperatorCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"kickBIPsOfOperatorStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"kickBIPsOfTotalStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"name\":\"minimumStake\",\"type\":\"uint96\",\"indexed\":false,\"internalType\":\"uint96\"},{\"name\":\"strategyParams\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"structIStakeRegistryTypes.StrategyParams[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"stakeType\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumIStakeRegistryTypes.StakeType\"},{\"name\":\"lookAheadPeriod\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyRegisteredForQuorums\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BitmapCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BitmapEmpty\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BitmapValueTooLarge\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BytesArrayLengthTooLong\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BytesArrayNotOrdered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotChurnSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotKickOperatorAboveThreshold\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotReregisterYet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChurnApproverSaltUsed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CurrentlyPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpModFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputAddressZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputLengthMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientStakeForChurn\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAVS\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidNewPausedStatus\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRegistrationType\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidShortString\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LookAheadPeriodTooLong\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MaxOperatorCountReached\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MaxQuorumsReached\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotRegisteredForQuorum\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotSorted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyAllocationManager\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyEjector\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyPauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyUnpauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorNotRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"QuorumDoesNotExist\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"QuorumOperatorCountMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SignatureExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StringTooLong\",\"inputs\":[{\"name\":\"str\",\"type\":\"string\",\"internalType\":\"string\"}]}]",
	Bin: "0x610160604052348015610010575f5ffd5b5060405161576838038061576883398101604081905261002f916101d3565b8686868686868661003f816100aa565b608052506001600160a01b03811661006a576040516339b190bb60e11b815260040160405180910390fd5b6001600160a01b0390811660a0529485166101005292841660e05290831661012052821660c052166101405261009e6100f0565b50505050505050610351565b5f5f829050601f815111156100dd578260405163305a27a960e01b81526004016100d491906102f6565b60405180910390fd5b80516100e88261032b565b179392505050565b5f54610100900460ff16156101575760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b60648201526084016100d4565b5f5460ff908116146101a6575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b03811681146101bc575f5ffd5b50565b634e487b7160e01b5f52604160045260245ffd5b5f5f5f5f5f5f5f60e0888a0312156101e9575f5ffd5b87516101f4816101a8565b6020890151909750610205816101a8565b6040890151909650610216816101a8565b6060890151909550610227816101a8565b6080890151909450610238816101a8565b60a0890151909350610249816101a8565b60c08901519092506001600160401b03811115610264575f5ffd5b8801601f81018a13610274575f5ffd5b80516001600160401b0381111561028d5761028d6101bf565b604051601f8201601f19908116603f011681016001600160401b03811182821017156102bb576102bb6101bf565b6040528181528282016020018c10156102d2575f5ffd5b8160208401602083015e5f6020838301015280935050505092959891949750929550565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b8051602080830151919081101561034b575f198160200360031b1b821691505b50919050565b60805160a05160c05160e0516101005161012051610140516153256104435f395f818161084801528181611f740152818161252a0152818161265501526135a401525f818161076201528181611037015281816122a80152818161278f0152612feb01525f818161066201528181611c830152818161222d015281816125c80152818161270e01528181612f4901526137b201525f818161063b01528181610df0015281816121b40152818161280401528181612c5c0152612ed001525f81816108fe0152611d6c01525f818161070b01528181610c84015281816114ab01526119c501525f6113fb01526153255ff3fe608060405234801561000f575f5ffd5b506004361061034c575f3560e01c80635df45946116101c9578063a65497c6116100fe578063de1164bb1161009e578063f2fde38b11610079578063f2fde38b14610920578063f698da2514610933578063fabc1cbc1461093b578063fd39105a1461094e575f5ffd5b8063de1164bb1461086a578063e65797ad1461087d578063ea32afae146108f9575f5ffd5b8063c391425e116100d9578063c391425e146107e9578063c63fd50214610809578063ca0de8821461081c578063ca8aa7c714610843575f5ffd5b8063a65497c6146107ab578063a96f783e146107be578063b5265787146107c7575f5ffd5b806384ca5213116101695780638da5cb5b116101445780638da5cb5b1461072d5780639aa1653d1461073e5780639e9923c21461075d5780639feab85914610784575f5ffd5b806384ca5213146106e0578063871ef049146106f3578063886f119514610706575f5ffd5b8063715018a6116101a4578063715018a614610697578063734479921461069f5780638281ab75146106b257806384b0196e146106c5575f5ffd5b80635df4594614610636578063683048351461065d5780636e3b17db14610684575f5ffd5b8063296bb0641161029f578063530b97a41161023f578063595c6a671161021a578063595c6a67146105f45780635ac86ab7146105fc5780635b0b829f1461061b5780635c975abb1461062e575f5ffd5b8063530b97a4146105ac57806354fd4d50146105bf5780635865c60c146105d4575f5ffd5b8063303ca9561161027a578063303ca956146105535780633c2a7f4c146105665780633eef3a51146105865780635140a54814610599575f5ffd5b8063296bb0641461051a57806329d1e0c31461052d5780632cdd1e8614610540575f5ffd5b8063125e05841161030a5780631478851f116102e55780631478851f1461046d5780631eb812da1461049f578063249a0c42146104e857806328f61b3114610507575f5ffd5b8063125e05841461041357806313542a4e14610432578063136439dd1461045a575f5ffd5b8062cf2ab51461035057806303fd34921461036557806304ec635114610397578063054310e6146103c25780630cf4b767146103ed5780630d3f213414610400575b5f5ffd5b61036361035e366004613c6e565b610989565b005b610384610373366004613c9f565b5f9081526098602052604090205490565b6040519081526020015b60405180910390f35b6103aa6103a5366004613cc7565b610b2a565b6040516001600160c01b03909116815260200161038e565b609d546103d5906001600160a01b031681565b6040516001600160a01b03909116815260200161038e565b6103636103fb366004613d6d565b610bc0565b61036361040e366004613c9f565b610c22565b610384610421366004613d9e565b609f6020525f908152604090205481565b610384610440366004613d9e565b6001600160a01b03165f9081526099602052604090205490565b610363610468366004613c9f565b610c6f565b61048f61047b366004613c9f565b609a6020525f908152604090205460ff1681565b604051901515815260200161038e565b6104b26104ad366004613db9565b610d44565b60408051825163ffffffff908116825260208085015190911690820152918101516001600160c01b03169082015260600161038e565b6103846104f6366004613dee565b609b6020525f908152604090205481565b609e546103d5906001600160a01b031681565b6103d5610528366004613c9f565b610dd8565b61036361053b366004613d9e565b610e61565b61036361054e366004613d9e565b610e72565b610363610561366004613e6b565b610e83565b610579610574366004613d9e565b610efd565b60405161038e9190613ec8565b610363610594366004613fe9565b610f21565b6103636105a736600461409c565b610f3d565b6103636105ba366004614179565b611264565b6105c76113f4565b60405161038e919061420b565b6105e76105e2366004613d9e565b611424565b60405161038e9190614245565b610363611496565b61048f61060a366004613dee565b6001805460ff9092161b9081161490565b610363610629366004614260565b611545565b600154610384565b6103d57f000000000000000000000000000000000000000000000000000000000000000081565b6103d57f000000000000000000000000000000000000000000000000000000000000000081565b610363610692366004614292565b611561565b61036361158d565b6103846106ad366004613d9e565b61159e565b6103636106c03660046142de565b611602565b6106cd611617565b60405161038e9796959493929190614330565b6103846106ee36600461445a565b6116b0565b6103aa610701366004613c9f565b6116f9565b6103d57f000000000000000000000000000000000000000000000000000000000000000081565b6064546001600160a01b03166103d5565b60965461074b9060ff1681565b60405160ff909116815260200161038e565b6103d57f000000000000000000000000000000000000000000000000000000000000000081565b6103847f2bd82124057f0913bc3b772ce7b83e8057c1ad1f3510fc83778be20f10ec5de681565b6103636107b9366004613d9e565b611703565b61038460a05481565b61048f6107d5366004613d9e565b60a1546001600160a01b0391821691161490565b6107fc6107f73660046144bf565b611714565b60405161038e9190614564565b6103636108173660046145ac565b61179b565b6103847f4d404e3276e7ac2163d8ee476afa6a41d1f68fb71f2d8b6546b24e55ce01b72a81565b6103d57f000000000000000000000000000000000000000000000000000000000000000081565b60a1546103d5906001600160a01b031681565b6108ec61088b366004613dee565b60408051606080820183525f808352602080840182905292840181905260ff9490941684526097825292829020825193840183525463ffffffff8116845261ffff6401000000008204811692850192909252600160301b9004169082015290565b60405161038e9190614639565b6103d57f000000000000000000000000000000000000000000000000000000000000000081565b61036361092e366004613d9e565b611944565b6103846119ba565b610363610949366004613c9f565b6119c3565b61097c61095c366004613d9e565b6001600160a01b03165f9081526099602052604090206001015460ff1690565b60405161038e919061466c565b6001546002906004908116036109b25760405163840a48d560e01b815260040160405180910390fd5b5f5b8251811015610b25576040805160018082528183019092525f91602080830190803683370190505090508382815181106109f0576109f061467a565b6020026020010151815f81518110610a0a57610a0a61467a565b6001600160a01b0392909216602092830291909101909101526040805160018082528183019092525f9181602001602082028036833701905050905060995f868581518110610a5b57610a5b61467a565b60200260200101516001600160a01b03166001600160a01b031681526020019081526020015f205f0154815f81518110610a9757610a9761467a565b6020026020010181815250505f610ac6825f81518110610ab957610ab961467a565b6020026020010151611ada565b90505f610adb826001600160c01b0316611b55565b90505f5b8151811015610b1457610b0c8585848481518110610aff57610aff61467a565b016020015160f81c611c1e565b600101610adf565b5050600190930192506109b4915050565b505050565b604051638192610f60e01b8152609860048201526024810184905263ffffffff83166044820152606481018290525f9073__$011c7a59801e6497cf0e5211ce0443418d$__90638192610f90608401602060405180830381865af4158015610b94573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610bb8919061468e565b949350505050565b6001335f9081526099602052604090206001015460ff166002811115610be857610be861421d565b14610c065760405163aba4733960e01b815260040160405180910390fd5b335f90815260996020526040902054610c1f9082611d55565b50565b610c2a611e00565b60a080549082905560408051828152602081018490527fa77a91bea7b6d95a8eb5a54878a1d9e3c875e26c86a9b70e3420c5c5db193b62910160405180910390a15050565b60405163237dfb4760e11b81523360048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906346fbf68e90602401602060405180830381865afa158015610cd1573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610cf591906146c3565b610d1257604051631d77d47760e21b815260040160405180910390fd5b6001548181168114610d375760405163c61dca5d60e01b815260040160405180910390fd5b610d4082611e5a565b5050565b604080516060810182525f80825260208201819052918101919091525f838152609860205260409020805483908110610d7f57610d7f61467a565b5f91825260209182902060408051606081018252919092015463ffffffff8082168352640100000000820416938201939093526001600160c01b0368010000000000000000909304929092169082015290505b92915050565b6040516308f6629d60e31b8152600481018290525f907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906347b314e890602401602060405180830381865afa158015610e3d573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610dd291906146dc565b610e69611e00565b610c1f81611e97565b610e7a611e00565b610c1f81611f00565b610e8b611f69565b60018054600290811603610eb25760405163840a48d560e01b815260040160405180910390fd5b60a1546001600160a01b03848116911614610ee0576040516366e565df60e01b815260040160405180910390fd5b5f610eea83611fb2565b9050610ef6858261205a565b5050505050565b604080518082019091525f8082526020820152610dd2610f1c8361159e565b61230a565b610f29611e00565b610f37848484600185612394565b50505050565b600154600290600490811603610f665760405163840a48d560e01b815260040160405180910390fd5b610fab83838080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152505060965460ff1691506128b39050565b5083518214610fcd5760405163aaad13f760e01b815260040160405180910390fd5b5f5b82811015610ef6575f848483818110610fea57610fea61467a565b885192013560f81c92505f91889150849081106110095761100961467a565b60209081029190910101516040516379a0849160e11b815260ff841660048201529091506001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063f341092290602401602060405180830381865afa15801561107c573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906110a091906146f7565b63ffffffff168151146110c657604051638e5aeee760e01b815260040160405180910390fd5b5f81516001600160401b038111156110e0576110e0613b37565b604051908082528060200260200182016040528015611109578160200160208202803683370190505b5090505f805b83518110156111ff575f84828151811061112b5761112b61467a565b6020026020010151905060995f826001600160a01b03166001600160a01b031681526020019081526020015f205f015484838151811061116d5761116d61467a565b6020026020010181815250505f61118f858481518110610ab957610ab961467a565b905060016001600160c01b03821660ff89161c8116146111c25760405163d053aa2160e01b815260040160405180910390fd5b836001600160a01b0316826001600160a01b0316116111f45760405163ba50f91160e01b815260040160405180910390fd5b50915060010161110f565b5061120b838386611c1e565b60ff84165f818152609b6020908152604091829020439081905591519182527f46077d55330763f16269fd75e5761663f4192d2791747c0189b16ad31db07db4910160405180910390a250505050806001019050610fcf565b5f54610100900460ff161580801561128257505f54600160ff909116105b8061129b5750303b15801561129b57505f5460ff166001145b6113035760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b5f805460ff191660011790558015611324575f805461ff0019166101001790555b61137a6040518060400160405280601681526020017520ab29a932b3b4b9ba393ca1b7b7b93234b730ba37b960511b8152506040518060400160405280600681526020016576302e302e3160d01b8152506128e7565b61138386612917565b61138c85611e97565b61139583611e5a565b61139e84611f00565b6113a782612968565b80156113ec575f805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050505050565b606061141f7f00000000000000000000000000000000000000000000000000000000000000006129d3565b905090565b604080518082019091525f80825260208201526001600160a01b0382165f908152609960209081526040918290208251808401909352805483526001810154909183019060ff16600281111561147c5761147c61421d565b600281111561148d5761148d61421d565b90525092915050565b60405163237dfb4760e11b81523360048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906346fbf68e90602401602060405180830381865afa1580156114f8573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061151c91906146c3565b61153957604051631d77d47760e21b815260040160405180910390fd5b6115435f19611e5a565b565b61154d611e00565b8161155781612a10565b610b258383612a39565b611569612acb565b6001600160a01b0382165f908152609f60205260409020429055610d408282612af6565b611595611e00565b6115435f612917565b5f610dd27f2bd82124057f0913bc3b772ce7b83e8057c1ad1f3510fc83778be20f10ec5de6836040516020016115e79291909182526001600160a01b0316602082015260400190565b60405160208183030381529060405280519060200120612b8e565b61160a611e00565b610b258383835f5f612394565b5f6060805f5f5f606060c8545f5f1b148015611633575060c954155b6116775760405162461bcd60e51b81526020600482015260156024820152741152540dcc4c8e88155b9a5b9a5d1a585b1a5e9959605a1b60448201526064016112fa565b61167f612bba565b611687612c4a565b604080515f80825260208201909252600f60f81b9b939a50919850469750309650945092509050565b5f6116ef7f4d404e3276e7ac2163d8ee476afa6a41d1f68fb71f2d8b6546b24e55ce01b72a87878787876040516020016115e796959493929190614712565b9695505050505050565b5f610dd282611ada565b61170b611e00565b610c1f81612968565b604051630449895d60e31b815260609073__$011c7a59801e6497cf0e5211ce0443418d$__9063224c4ae8906117539060989087908790600401614797565b5f60405180830381865af415801561176d573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261179491908101906147f1565b9392505050565b6117a3611f69565b600180545f91908116036117ca5760405163840a48d560e01b815260040160405180910390fd5b60a1546001600160a01b038681169116146117f8576040516366e565df60e01b815260040160405180910390fd5b5f61180285611fb2565b90505f80806118138688018861497f565b9250925092505f6118248b83612c59565b90505f8460018111156118395761183961421d565b036118e2575f61184d8c8388876001612cf2565b5190505f5b86518110156118db575f87828151811061186e5761186e61467a565b0160209081015160f81c5f8181526097909252604090912054845191925063ffffffff16908490849081106118a5576118a561467a565b602002602001015163ffffffff1611156118d25760405163c6b9e76760e01b815260040160405180910390fd5b50600101611852565b5050611937565b60018460018111156118f6576118f661421d565b0361191e575f80611909898b018b6149da565b945094505050506118db8d848988868661313c565b60405163354bb8ab60e01b815260040160405180910390fd5b5050505050505050505050565b61194c611e00565b6001600160a01b0381166119b15760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016112fa565b610c1f81612917565b5f61141f613332565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611a1f573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611a4391906146dc565b6001600160a01b0316336001600160a01b031614611a745760405163794821ff60e01b815260040160405180910390fd5b60015480198219811614611a9b5760405163c61dca5d60e01b815260040160405180910390fd5b600182905560405182815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c906020015b60405180910390a25050565b604051631a3894e960e01b815260986004820152602481018290525f9073__$011c7a59801e6497cf0e5211ce0443418d$__90631a3894e990604401602060405180830381865af4158015611b31573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610dd2919061468e565b60605f5f611b628461333b565b61ffff166001600160401b03811115611b7d57611b7d613b37565b6040519080825280601f01601f191660200182016040528015611ba7576020820181803683370190505b5090505f805b825182108015611bbe575061010081105b15611c14576001811b935085841615611c04578060f81b838381518110611be757611be761467a565b60200101906001600160f81b03191690815f1a9053508160010191505b611c0d81614ae7565b9050611bad565b5090949350505050565b6040805160018082528183019092525f916020820181803683370190505090508160f81b815f81518110611c5457611c5461467a565b60200101906001600160f81b03191690815f1a905350604051636c3fb4bf60e01b81525f906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690636c3fb4bf90611cbc90889088908890600401614aff565b5f604051808303815f875af1158015611cd7573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052611cfe9190810190614b8f565b90505f5b85518110156113ec57818181518110611d1d57611d1d61467a565b602002602001015115611d4d57611d4d868281518110611d3f57611d3f61467a565b602002602001015184612af6565b600101611d02565b6040516378219b3f60e11b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063f043367e90611da39085908590600401614c1c565b5f604051808303815f87803b158015611dba575f5ffd5b505af1158015611dcc573d5f5f3e3d5ffd5b50505050817fec2963ab21c1e50e1e582aa542af2e4bf7bf38e6e1403c27b42e1c5d6e621eaa82604051611ace919061420b565b6064546001600160a01b031633146115435760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016112fa565b600181905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a250565b609d54604080516001600160a01b03928316815291831660208301527f315457d8a8fe60f04af17c16e2f5a5e1db612b31648e58030360759ef8f3528c910160405180910390a1609d80546001600160a01b0319166001600160a01b0392909216919091179055565b609e54604080516001600160a01b03928316815291831660208301527f8f30ab09f43a6c157d7fce7e0a13c003042c1c95e8a72e7a146a21c0caa24dc9910160405180910390a1609e80546001600160a01b0319166001600160a01b0392909216919091179055565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614611543576040516323d871a560e01b815260040160405180910390fd5b60605f82516001600160401b03811115611fce57611fce613b37565b6040519080825280601f01601f191660200182016040528015611ff8576020820181803683370190505b5090505f5b8351811015612053578381815181106120185761201861467a565b602002602001015160f81b8282815181106120355761203561467a565b60200101906001600160f81b03191690815f1a905350600101611ffd565b5092915050565b6001600160a01b0382165f9081526099602052604081208054909161207e82611ada565b905060018084015460ff16600281111561209a5761209a61421d565b146120b85760405163aba4733960e01b815260040160405180910390fd5b6096545f906120cb90869060ff166128b3565b90506001600160c01b0381166120f4576040516368b6a87560e11b815260040160405180910390fd5b61210b6001600160c01b0382811690841681161490565b6121285760405163d053aa2160e01b815260040160405180910390fd5b6001600160c01b03818116198316166121418482613365565b6001600160c01b03811661219d576001600160a01b0387165f81815260996020526040808220600101805460ff19166002179055518692917f396fdcb180cb0fea26928113fb0fd1c3549863f9cd563e6a184f1d578116c8e491a35b60405163f4e24fe560e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063f4e24fe5906121eb908a908a90600401614c34565b5f604051808303815f87803b158015612202575f5ffd5b505af1158015612214573d5f5f3e3d5ffd5b505060405163bd29b8cd60e01b81526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016925063bd29b8cd91506122669087908a90600401614c1c565b5f604051808303815f87803b15801561227d575f5ffd5b505af115801561228f573d5f5f3e3d5ffd5b505060405163bd29b8cd60e01b81526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016925063bd29b8cd91506122e19087908a90600401614c1c565b5f604051808303815f87803b1580156122f8575f5ffd5b505af1158015611937573d5f5f3e3d5ffd5b604080518082019091525f80825260208201525f80806123375f5160206152d05f395f51905f5286614c6b565b90505b612343816133d6565b90935091505f5160206152d05f395f51905f52828309830361237b576040805180820190915290815260208101919091529392505050565b5f5160206152d05f395f51905f5260018208905061233a565b60965460ff1660c081106123bb57604051633cb89c9760e01b815260040160405180910390fd5b60968054600191905f906123d390849060ff16614c7e565b92506101000a81548160ff021916908360ff1602179055506123f58187612a39565b6040805160018082528183019092525f91816020015b604080518082019091525f81526060602082015281526020019060019003908161240b5790505090505f85516001600160401b0381111561244e5761244e613b37565b604051908082528060200260200182016040528015612477578160200160208202803683370190505b5090505f5b86518110156124d4578681815181106124975761249761467a565b60200260200101515f01518282815181106124b4576124b461467a565b6001600160a01b039092166020928302919091019091015260010161247c565b5060405180604001604052808460ff1663ffffffff16815260200182815250825f815181106125055761250561467a565b602090810291909101015260a154604051630130fc2760e51b81526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000081169263261f84e0926125659291909116908690600401614c97565b5f604051808303815f87803b15801561257c575f5ffd5b505af115801561258e573d5f5f3e3d5ffd5b505f925061259a915050565b8560018111156125ac576125ac61421d565b0361263357604051633aea0b9d60e11b81526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906375d4173a906126019086908b908b90600401614db1565b5f604051808303815f87803b158015612618575f5ffd5b505af115801561262a573d5f5f3e3d5ffd5b50505050612777565b60018560018111156126475761264761421d565b03612777578363ffffffff167f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316632981eb776040518163ffffffff1660e01b8152600401602060405180830381865afa1580156126af573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906126d391906146f7565b63ffffffff16116126f757604051630bd441b960e21b815260040160405180910390fd5b604051630662d3e160e51b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063cc5a7c20906127499086908b9089908c90600401614de4565b5f604051808303815f87803b158015612760575f5ffd5b505af1158015612772573d5f5f3e3d5ffd5b505050505b60405163136ca0f960e11b815260ff841660048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906326d941f2906024015f604051808303815f87803b1580156127d8575f5ffd5b505af11580156127ea573d5f5f3e3d5ffd5b505060405163136ca0f960e11b815260ff861660048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031692506326d941f291506024015f604051808303815f87803b15801561284f575f5ffd5b505af1158015612861573d5f5f3e3d5ffd5b505050508260ff167fa34835bc2b673ec37fcf1591a91295b163fc2e181e4ea4e733beb27de1ceac4c89898989896040516128a0959493929190614e1a565b60405180910390a25b5050505050505050565b5f5f6128be84613452565b9050808360ff166001901b116117945760405163ca95733360e01b815260040160405180910390fd5b5f54610100900460ff1661290d5760405162461bcd60e51b81526004016112fa90614e9b565b610d40828261350d565b606480546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b60a154604080516001600160a01b0392831680825292841660208201527f9770f3cadfdcbb6f93af935e86047111590c3768271d237e4a2bc0b874bed693910160405180910390a15060a180546001600160a01b0319166001600160a01b0392909216919091179055565b60605f6129df8361355a565b6040805160208082528183019092529192505f91906020820181803683375050509182525060208101929092525090565b60965460ff90811690821610610c1f57604051637310cff560e11b815260040160405180910390fd5b60ff82165f8181526097602090815260409182902084518154928601518487015161ffff908116600160301b0267ffff00000000000019919092166401000000000265ffffffffffff1990951663ffffffff909316929092179390931716919091179055517f3ee6fe8d54610244c3e9d3c066ae4aee997884aa28f10616ae821925401318ac90611ace908490614639565b609e546001600160a01b03163314611543576040516376d8ab1760e11b815260040160405180910390fd5b6001600160a01b0382165f90815260996020526040902060018082015460ff166002811115612b2757612b2761421d565b14612b45576040516325ec6c1f60e01b815260040160405180910390fd5b80546096545f90612b5a90859060ff166128b3565b90505f612b6683611ada565b9050612b7f6001600160c01b0383811690831681161490565b156113ec576113ec8686613581565b5f610dd2612b9a613332565b8360405161190160f01b8152600281019290925260228201526042902090565b606060ca8054612bc990614ee6565b80601f0160208091040260200160405190810160405280929190818152602001828054612bf590614ee6565b8015612c405780601f10612c1757610100808354040283529160200191612c40565b820191905f5260205f20905b815481529060010190602001808311612c2357829003601f168201915b5050505050905090565b606060cb8054612bc990614ee6565b5f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166303c5a6b68484612c9487610efd565b6040518463ffffffff1660e01b8152600401612cb293929190614f40565b6020604051808303815f875af1158015612cce573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906117949190614fb9565b612d1660405180606001604052806060815260200160608152602001606081525090565b6096545f90612d2990869060ff166128b3565b90505f612d3587611ada565b90506001600160c01b038216612d5e576040516313ca465760e01b815260040160405180910390fd5b8082166001600160c01b031615612d8857604051630c6816cd60e01b815260040160405180910390fd5b60a0546001600160a01b0389165f908152609f60205260409020546001600160c01b0383811690851617914291612dbf9190614fd0565b10612ddd57604051631968677d60e11b815260040160405180910390fd5b612de78882613365565b612df18887611d55565b60016001600160a01b038a165f9081526099602052604090206001015460ff166002811115612e2257612e2261421d565b14612eb957604080518082018252898152600160208083018281526001600160a01b038e165f908152609990925293902082518155925183820180549394939192909160ff191690836002811115612e7c57612e7c61421d565b0217905550506040518991506001600160a01b038b16907fe8e68cef1c3a761ed7be7e8463a375f27f7bc335e51824223cacce636ec5c3fe905f90a35b604051631fd93ca960e11b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690633fb2795290612f07908c908b90600401614c34565b5f604051808303815f87803b158015612f1e575f5ffd5b505af1158015612f30573d5f5f3e3d5ffd5b5050604051632550477760e01b81526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016925063255047779150612f84908c908c908c90600401614fe3565b5f604051808303815f875af1158015612f9f573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052612fc6919081019061506d565b60408087019190915260208601919091525162bff04d60e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169062bff04d90613021908b908b90600401614c1c565b5f604051808303815f875af115801561303c573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261306391908101906147f1565b84528415613130575f5b875181101561312e575f60975f8a848151811061308c5761308c61467a565b0160209081015160f81c82528181019290925260409081015f208151606081018352905463ffffffff811680835261ffff6401000000008304811695840195909552600160301b909104909316918101919091528751805191935090849081106130f8576130f861467a565b602002602001015163ffffffff1611156131255760405163c6b9e76760e01b815260040160405180910390fd5b5060010161306d565b505b50505095945050505050565b835182511461315e5760405163aaad13f760e01b815260040160405180910390fd5b61316a8686848461361f565b5f613178878787875f612cf2565b90505f5b85518110156128a9575f60975f88848151811061319b5761319b61467a565b0160209081015160f81c82528181019290925260409081015f208151606081018352905463ffffffff811680835261ffff6401000000008304811695840195909552600160301b909104909316918101919091528451805191935090849081106132075761320761467a565b602002602001015163ffffffff1611156133295761329b8783815181106132305761323061467a565b602001015160f81c60f81b60f81c846040015184815181106132545761325461467a565b60200260200101518b866020015186815181106132735761327361467a565b602002602001015189878151811061328d5761328d61467a565b602002602001015186613725565b6040805160018082528183019092525f916020820181803683370190505090508783815181106132cd576132cd61467a565b602001015160f81c60f81b815f815181106132ea576132ea61467a565b60200101906001600160f81b03191690815f1a9053506133278684815181106133155761331561467a565b60200260200101516020015182612af6565b505b5060010161317c565b5f61141f6138a6565b5f805b8215610dd25761334f6001846150c6565b909216918061335d816150d9565b91505061333e565b60405163ae1d62db60e01b815260986004820152602481018390526001600160c01b038216604482015273__$011c7a59801e6497cf0e5211ce0443418d$__9063ae1d62db906064015f6040518083038186803b1580156133c4575f5ffd5b505af41580156113ec573d5f5f3e3d5ffd5b5f80805f5160206152d05f395f51905f5260035f5160206152d05f395f51905f52865f5160206152d05f395f51905f52888909090890505f613446827f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f525f5160206152d05f395f51905f52613919565b91959194509092505050565b5f6101008251111561347757604051637da54e4760e11b815260040160405180910390fd5b81515f0361348657505f919050565b5f5f835f8151811061349a5761349a61467a565b0160200151600160f89190911c81901b92505b8451811015613504578481815181106134c8576134c861467a565b0160200151600160f89190911c1b91508282116134f857604051631019106960e31b815260040160405180910390fd5b918117916001016134ad565b50909392505050565b5f54610100900460ff166135335760405162461bcd60e51b81526004016112fa90614e9b565b60ca61353f838261513d565b5060cb61354c828261513d565b50505f60c881905560c95550565b5f60ff8216601f811115610dd257604051632cd44ac360e21b815260040160405180910390fd5b604080516060810182526001600160a01b03848116825260a154811660208301527f00000000000000000000000000000000000000000000000000000000000000001691636e3492b5919081016135d785613992565b8152506040518263ffffffff1660e01b81526004016135f691906151f7565b5f604051808303815f87803b15801561360d575f5ffd5b505af11580156113ec573d5f5f3e3d5ffd5b6020808201515f908152609a909152604090205460ff161561365457604051636fbefec360e11b815260040160405180910390fd5b428160400151101561367957604051630819bdcd60e01b815260040160405180910390fd5b602080820180515f908152609a909252604091829020805460ff19166001179055609d5490519183015173__$c450b067f122b1edf84c92f7cfafc3aa17$__9263238a4d1e926001600160a01b0316916136d8918991899189916116b0565b84516040516001600160e01b031960e086901b1681526136fd93929190600401614fe3565b5f6040518083038186803b158015613713575f5ffd5b505af41580156128a9573d5f5f3e3d5ffd5b6020808301516001600160a01b038082165f8181526099909452604090932054919290871603613768576040516356168b4160e11b815260040160405180910390fd5b8760ff16845f015160ff161461379157604051638e5aeee760e01b815260040160405180910390fd5b604051635401ed2760e01b81526004810182905260ff891660248201525f907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690635401ed2790604401602060405180830381865afa1580156137ff573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906138239190615265565b905061382f8185613a37565b6001600160601b0316866001600160601b03161161386057604051634c44995d60e01b815260040160405180910390fd5b61386a8885613a5a565b6001600160601b0316816001600160601b03161061389b5760405163b187e86960e01b815260040160405180910390fd5b505050505050505050565b5f7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f6138d0613a73565b6138d8613acb565b60408051602081019490945283019190915260608201524660808201523060a082015260c00160405160208183030381529060405280519060200120905090565b5f5f613923613afb565b61392b613b19565b602080825281810181905260408201819052606082018890526080820187905260a082018690528260c08360056107d05a03fa9250828061396857fe5b50826139875760405163d51edae360e01b815260040160405180910390fd5b505195945050505050565b60605f82516001600160401b038111156139ae576139ae613b37565b6040519080825280602002602001820160405280156139d7578160200160208202803683370190505b5090505f5b8351811015612053578381815181106139f7576139f761467a565b602001015160f81c60f81b60f81c60ff16828281518110613a1a57613a1a61467a565b63ffffffff909216602092830291909101909101526001016139dc565b60208101515f9061271090613a509061ffff1685615280565b61179491906152a2565b60408101515f9061271090613a509061ffff1685615280565b5f5f613a7d612bba565b805190915015613a94578051602090910120919050565b60c8548015613aa35792915050565b7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a4709250505090565b5f5f613ad5612c4a565b805190915015613aec578051602090910120919050565b60c9548015613aa35792915050565b60405180602001604052806001906020820280368337509192915050565b6040518060c001604052806006906020820280368337509192915050565b634e487b7160e01b5f52604160045260245ffd5b604051606081016001600160401b0381118282101715613b6d57613b6d613b37565b60405290565b604080519081016001600160401b0381118282101715613b6d57613b6d613b37565b604051601f8201601f191681016001600160401b0381118282101715613bbd57613bbd613b37565b604052919050565b5f6001600160401b03821115613bdd57613bdd613b37565b5060051b60200190565b6001600160a01b0381168114610c1f575f5ffd5b5f82601f830112613c0a575f5ffd5b8135613c1d613c1882613bc5565b613b95565b8082825260208201915060208360051b860101925085831115613c3e575f5ffd5b602085015b83811015613c64578035613c5681613be7565b835260209283019201613c43565b5095945050505050565b5f60208284031215613c7e575f5ffd5b81356001600160401b03811115613c93575f5ffd5b610bb884828501613bfb565b5f60208284031215613caf575f5ffd5b5035919050565b63ffffffff81168114610c1f575f5ffd5b5f5f5f60608486031215613cd9575f5ffd5b833592506020840135613ceb81613cb6565b929592945050506040919091013590565b5f82601f830112613d0b575f5ffd5b8135602083015f5f6001600160401b03841115613d2a57613d2a613b37565b50601f8301601f1916602001613d3f81613b95565b915050828152858383011115613d53575f5ffd5b828260208301375f92810160200192909252509392505050565b5f60208284031215613d7d575f5ffd5b81356001600160401b03811115613d92575f5ffd5b610bb884828501613cfc565b5f60208284031215613dae575f5ffd5b813561179481613be7565b5f5f60408385031215613dca575f5ffd5b50508035926020909101359150565b803560ff81168114613de9575f5ffd5b919050565b5f60208284031215613dfe575f5ffd5b61179482613dd9565b5f82601f830112613e16575f5ffd5b8135613e24613c1882613bc5565b8082825260208201915060208360051b860101925085831115613e45575f5ffd5b602085015b83811015613c64578035613e5d81613cb6565b835260209283019201613e4a565b5f5f5f60608486031215613e7d575f5ffd5b8335613e8881613be7565b92506020840135613e9881613be7565b915060408401356001600160401b03811115613eb2575f5ffd5b613ebe86828701613e07565b9150509250925092565b815181526020808301519082015260408101610dd2565b803561ffff81168114613de9575f5ffd5b5f60608284031215613f00575f5ffd5b613f08613b4b565b90508135613f1581613cb6565b8152613f2360208301613edf565b6020820152613f3460408301613edf565b604082015292915050565b6001600160601b0381168114610c1f575f5ffd5b5f82601f830112613f62575f5ffd5b8135613f70613c1882613bc5565b8082825260208201915060208360061b860101925085831115613f91575f5ffd5b602085015b83811015613c645760408188031215613fad575f5ffd5b613fb5613b73565b8135613fc081613be7565b81526020820135613fd081613f3f565b6020828101919091529084529290920191604001613f96565b5f5f5f5f60c08587031215613ffc575f5ffd5b6140068686613ef0565b9350606085013561401681613f3f565b925060808501356001600160401b03811115614030575f5ffd5b61403c87828801613f53565b92505060a085013561404d81613cb6565b939692955090935050565b5f5f83601f840112614068575f5ffd5b5081356001600160401b0381111561407e575f5ffd5b602083019150836020828501011115614095575f5ffd5b9250929050565b5f5f5f604084860312156140ae575f5ffd5b83356001600160401b038111156140c3575f5ffd5b8401601f810186136140d3575f5ffd5b80356140e1613c1882613bc5565b8082825260208201915060208360051b850101925088831115614102575f5ffd5b602084015b838110156141425780356001600160401b03811115614124575f5ffd5b6141338b602083890101613bfb565b84525060209283019201614107565b50955050505060208401356001600160401b03811115614160575f5ffd5b61416c86828701614058565b9497909650939450505050565b5f5f5f5f5f60a0868803121561418d575f5ffd5b853561419881613be7565b945060208601356141a881613be7565b935060408601356141b881613be7565b92506060860135915060808601356141cf81613be7565b809150509295509295909350565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b602081525f61179460208301846141dd565b634e487b7160e01b5f52602160045260245ffd5b600381106142415761424161421d565b9052565b81518152602080830151604083019161205390840182614231565b5f5f60808385031215614271575f5ffd5b61427a83613dd9565b91506142898460208501613ef0565b90509250929050565b5f5f604083850312156142a3575f5ffd5b82356142ae81613be7565b915060208301356001600160401b038111156142c8575f5ffd5b6142d485828601613cfc565b9150509250929050565b5f5f5f60a084860312156142f0575f5ffd5b6142fa8585613ef0565b9250606084013561430a81613f3f565b915060808401356001600160401b03811115614324575f5ffd5b613ebe86828701613f53565b60ff60f81b8816815260e060208201525f61434e60e08301896141dd565b828103604084015261436081896141dd565b606084018890526001600160a01b038716608085015260a0840186905283810360c0850152845180825260208087019350909101905f5b818110156143b5578351835260209384019390920191600101614397565b50909b9a5050505050505050505050565b5f82601f8301126143d5575f5ffd5b81356143e3613c1882613bc5565b8082825260208201915060208360061b860101925085831115614404575f5ffd5b602085015b83811015613c645760408188031215614420575f5ffd5b614428613b73565b61443182613dd9565b8152602082013561444181613be7565b6020828101919091529084529290920191604001614409565b5f5f5f5f5f60a0868803121561446e575f5ffd5b853561447981613be7565b94506020860135935060408601356001600160401b0381111561449a575f5ffd5b6144a6888289016143c6565b9598949750949560608101359550608001359392505050565b5f5f604083850312156144d0575f5ffd5b82356144db81613cb6565b915060208301356001600160401b038111156144f5575f5ffd5b8301601f81018513614505575f5ffd5b8035614513613c1882613bc5565b8082825260208201915060208360051b850101925087831115614534575f5ffd5b6020840193505b8284101561455657833582526020938401939091019061453b565b809450505050509250929050565b602080825282518282018190525f918401906040840190835b818110156145a157835163ffffffff1683526020938401939092019160010161457d565b509095945050505050565b5f5f5f5f5f608086880312156145c0575f5ffd5b85356145cb81613be7565b945060208601356145db81613be7565b935060408601356001600160401b038111156145f5575f5ffd5b61460188828901613e07565b93505060608601356001600160401b0381111561461c575f5ffd5b61462888828901614058565b969995985093965092949392505050565b60608101610dd2828463ffffffff815116825261ffff602082015116602083015261ffff60408201511660408301525050565b60208101610dd28284614231565b634e487b7160e01b5f52603260045260245ffd5b5f6020828403121561469e575f5ffd5b81516001600160c01b0381168114611794575f5ffd5b80518015158114613de9575f5ffd5b5f602082840312156146d3575f5ffd5b611794826146b4565b5f602082840312156146ec575f5ffd5b815161179481613be7565b5f60208284031215614707575f5ffd5b815161179481613cb6565b5f60c0820188835260018060a01b038816602084015286604084015260c0606084015280865180835260e0850191506020880192505f5b8181101561477f578351805160ff1684526020908101516001600160a01b03168185015290930192604090920191600101614749565b50506080840195909552505060a00152949350505050565b5f6060820185835263ffffffff85166020840152606060408401528084518083526080850191506020860192505f5b818110156147e45783518352602093840193909201916001016147c6565b5090979650505050505050565b5f60208284031215614801575f5ffd5b81516001600160401b03811115614816575f5ffd5b8201601f81018413614826575f5ffd5b8051614834613c1882613bc5565b8082825260208201915060208360051b850101925086831115614855575f5ffd5b6020840193505b828410156116ef57835161486f81613cb6565b82526020938401939091019061485c565b803560028110613de9575f5ffd5b5f6040828403121561489e575f5ffd5b6148a6613b73565b823581526020928301359281019290925250919050565b5f82601f8301126148cc575f5ffd5b6148d4613b73565b8060408401858111156148e5575f5ffd5b845b818110156145a15780358452602093840193016148e7565b5f818303610100811215614911575f5ffd5b614919613b4b565b9150614925848461488e565b8252614934846040850161488e565b60208301526080607f198201121561494a575f5ffd5b50614953613b73565b61496084608085016148bd565b815261496f8460c085016148bd565b6020820152604082015292915050565b5f5f5f6101408486031215614992575f5ffd5b61499b84614880565b925060208401356001600160401b038111156149b5575f5ffd5b6149c186828701613cfc565b9250506149d185604086016148ff565b90509250925092565b5f5f5f5f5f61018086880312156149ef575f5ffd5b6149f886614880565b945060208601356001600160401b03811115614a12575f5ffd5b614a1e88828901613cfc565b945050614a2e87604088016148ff565b92506101408601356001600160401b03811115614a49575f5ffd5b614a55888289016143c6565b9250506101608601356001600160401b03811115614a71575f5ffd5b860160608189031215614a82575f5ffd5b614a8a613b4b565b81356001600160401b03811115614a9f575f5ffd5b614aab8a828501613cfc565b8252506020828101359082015260409182013591810191909152949793965091945092919050565b634e487b7160e01b5f52601160045260245ffd5b5f60018201614af857614af8614ad3565b5060010190565b606080825284519082018190525f9060208601906080840190835b81811015614b415783516001600160a01b0316835260209384019390920191600101614b1a565b5050838103602080860191909152865180835291810192508601905f5b81811015614b7c578251845260209384019390920191600101614b5e565b50505060ff841660408401529050610bb8565b5f60208284031215614b9f575f5ffd5b81516001600160401b03811115614bb4575f5ffd5b8201601f81018413614bc4575f5ffd5b8051614bd2613c1882613bc5565b8082825260208201915060208360051b850101925086831115614bf3575f5ffd5b6020840193505b828410156116ef57614c0b846146b4565b825260209384019390910190614bfa565b828152604060208201525f610bb860408301846141dd565b6001600160a01b03831681526040602082018190525f90610bb8908301846141dd565b634e487b7160e01b5f52601260045260245ffd5b5f82614c7957614c79614c57565b500690565b60ff8181168382160190811115610dd257610dd2614ad3565b5f6040820160018060a01b03851683526040602084015280845180835260608501915060608160051b8601019250602086015f5b82811015614d4c57868503605f190184528151805163ffffffff168652602090810151604082880181905281519088018190529101905f9060608801905b80831015614d345783516001600160a01b031682526020938401936001939093019290910190614d09565b50965050506020938401939190910190600101614ccb565b5092979650505050505050565b5f8151808452602084019350602083015f5b82811015614da757815180516001600160a01b031687526020908101516001600160601b03168188015260409096019590910190600101614d6b565b5093949350505050565b60ff841681526001600160601b0383166020820152606060408201525f614ddb6060830184614d59565b95945050505050565b60ff851681526001600160601b038416602082015263ffffffff83166040820152608060608201525f6116ef6080830184614d59565b614e49818763ffffffff815116825261ffff602082015116602083015261ffff60408201511660408301525050565b6001600160601b038516606082015260e060808201525f614e6d60e0830186614d59565b905060028410614e7f57614e7f61421d565b8360a083015263ffffffff831660c08301529695505050505050565b6020808252602b908201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960408201526a6e697469616c697a696e6760a81b606082015260800190565b600181811c90821680614efa57607f821691505b602082108103614f1857634e487b7160e01b5f52602260045260245ffd5b50919050565b805f5b6002811015610f37578151845260209384019390910190600101614f21565b6001600160a01b03841681528251805160208084019190915201516040820152610160810160208481015180516060850152908101516080840152506040840151614f8f60a084018251614f1e565b60200151614fa060e0840182614f1e565b5082516101208301526020830151610140830152610bb8565b5f60208284031215614fc9575f5ffd5b5051919050565b80820180821115610dd257610dd2614ad3565b60018060a01b0384168152826020820152606060408201525f614ddb60608301846141dd565b5f82601f830112615018575f5ffd5b8151615026613c1882613bc5565b8082825260208201915060208360051b860101925085831115615047575f5ffd5b602085015b83811015613c6457805161505f81613f3f565b83526020928301920161504c565b5f5f6040838503121561507e575f5ffd5b82516001600160401b03811115615093575f5ffd5b61509f85828601615009565b92505060208301516001600160401b038111156150ba575f5ffd5b6142d485828601615009565b81810381811115610dd257610dd2614ad3565b5f61ffff821661ffff81036150f0576150f0614ad3565b60010192915050565b601f821115610b2557805f5260205f20601f840160051c8101602085101561511e5750805b601f840160051c820191505b81811015610ef6575f815560010161512a565b81516001600160401b0381111561515657615156613b37565b61516a816151648454614ee6565b846150f9565b6020601f82116001811461519c575f83156151855750848201515b5f19600385901b1c1916600184901b178455610ef6565b5f84815260208120601f198516915b828110156151cb57878501518255602094850194600190920191016151ab565b50848210156151e857868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b602080825282516001600160a01b039081168383015283820151166040808401919091528301516060808401528051608084018190525f929190910190829060a08501905b80831015613c645763ffffffff845116825260208201915060208401935060018301925061523c565b5f60208284031215615275575f5ffd5b815161179481613f3f565b6001600160601b03818116838216029081169081811461205357612053614ad3565b5f6001600160601b038316806152ba576152ba614c57565b806001600160601b038416049150509291505056fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47a2646970667358221220e4f94f6b188e56d8e146add694617e7663604380f5a4ff805cbe3ce583060fb864736f6c634300081b0033",
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
