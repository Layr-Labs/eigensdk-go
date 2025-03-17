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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_stakeRegistry\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"},{\"name\":\"_blsApkRegistry\",\"type\":\"address\",\"internalType\":\"contractIBLSApkRegistry\"},{\"name\":\"_indexRegistry\",\"type\":\"address\",\"internalType\":\"contractIIndexRegistry\"},{\"name\":\"_socketRegistry\",\"type\":\"address\",\"internalType\":\"contractISocketRegistry\"},{\"name\":\"_allocationManager\",\"type\":\"address\",\"internalType\":\"contractIAllocationManager\"},{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"OPERATOR_CHURN_APPROVAL_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PUBKEY_REGISTRATION_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allocationManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIAllocationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"avs\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"blsApkRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBLSApkRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateOperatorChurnApprovalDigestHash\",\"inputs\":[{\"name\":\"registeringOperator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"registeringOperatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"operatorKickParams\",\"type\":\"tuple[]\",\"internalType\":\"structISlashingRegistryCoordinatorTypes.OperatorKickParam[]\",\"components\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculatePubkeyRegistrationMessageHash\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"churnApprover\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createSlashableStakeQuorum\",\"inputs\":[{\"name\":\"operatorSetParams\",\"type\":\"tuple\",\"internalType\":\"structISlashingRegistryCoordinatorTypes.OperatorSetParam\",\"components\":[{\"name\":\"maxOperatorCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"kickBIPsOfOperatorStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"kickBIPsOfTotalStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"name\":\"minimumStake\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"strategyParams\",\"type\":\"tuple[]\",\"internalType\":\"structIStakeRegistryTypes.StrategyParams[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"lookAheadPeriod\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createTotalDelegatedStakeQuorum\",\"inputs\":[{\"name\":\"operatorSetParams\",\"type\":\"tuple\",\"internalType\":\"structISlashingRegistryCoordinatorTypes.OperatorSetParam\",\"components\":[{\"name\":\"maxOperatorCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"kickBIPsOfOperatorStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"kickBIPsOfTotalStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"name\":\"minimumStake\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"strategyParams\",\"type\":\"tuple[]\",\"internalType\":\"structIStakeRegistryTypes.StrategyParams[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deregisterOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"domainSeparator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"eip712Domain\",\"inputs\":[],\"outputs\":[{\"name\":\"fields\",\"type\":\"bytes1\",\"internalType\":\"bytes1\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"verifyingContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"extensions\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ejectOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ejectionCooldown\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ejector\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentQuorumBitmap\",\"inputs\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint192\",\"internalType\":\"uint192\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structISlashingRegistryCoordinatorTypes.OperatorInfo\",\"components\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumISlashingRegistryCoordinatorTypes.OperatorStatus\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorFromId\",\"inputs\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorId\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorSetParams\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structISlashingRegistryCoordinatorTypes.OperatorSetParam\",\"components\":[{\"name\":\"maxOperatorCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"kickBIPsOfOperatorStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"kickBIPsOfTotalStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorStatus\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumISlashingRegistryCoordinatorTypes.OperatorStatus\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumBitmapAtBlockNumberByIndex\",\"inputs\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint192\",\"internalType\":\"uint192\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumBitmapHistoryLength\",\"inputs\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumBitmapIndicesAtBlockNumber\",\"inputs\":[{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumBitmapUpdateByIndex\",\"inputs\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structISlashingRegistryCoordinatorTypes.QuorumBitmapUpdate\",\"components\":[{\"name\":\"updateBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"nextUpdateBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumBitmap\",\"type\":\"uint192\",\"internalType\":\"uint192\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"indexRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIIndexRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_churnApprover\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_ejector\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_initialPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_avs\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isChurnApproverSaltUsed\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lastEjectionTimestamp\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"numRegistries\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pubkeyRegistrationMessageHash\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"quorumCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"quorumUpdateBlockNumber\",\"inputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registries\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setAVS\",\"inputs\":[{\"name\":\"_avs\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChurnApprover\",\"inputs\":[{\"name\":\"_churnApprover\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setEjectionCooldown\",\"inputs\":[{\"name\":\"_ejectionCooldown\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setEjector\",\"inputs\":[{\"name\":\"_ejector\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setOperatorSetParams\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"operatorSetParams\",\"type\":\"tuple\",\"internalType\":\"structISlashingRegistryCoordinatorTypes.OperatorSetParam\",\"components\":[{\"name\":\"maxOperatorCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"kickBIPsOfOperatorStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"kickBIPsOfTotalStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"socketRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractISocketRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stakeRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"supportsAVS\",\"inputs\":[{\"name\":\"_avs\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateOperators\",\"inputs\":[{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateOperatorsForQuorum\",\"inputs\":[{\"name\":\"operatorsPerQuorum\",\"type\":\"address[][]\",\"internalType\":\"address[][]\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateSocket\",\"inputs\":[{\"name\":\"socket\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"event\",\"name\":\"ChurnApproverUpdated\",\"inputs\":[{\"name\":\"prevChurnApprover\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newChurnApprover\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EIP712DomainChanged\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EjectorUpdated\",\"inputs\":[{\"name\":\"prevEjector\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newEjector\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorDeregistered\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorRegistered\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSetParamsUpdated\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"uint8\"},{\"name\":\"operatorSetParams\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structISlashingRegistryCoordinatorTypes.OperatorSetParam\",\"components\":[{\"name\":\"maxOperatorCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"kickBIPsOfOperatorStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"kickBIPsOfTotalStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSocketUpdate\",\"inputs\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"socket\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"QuorumBlockNumberUpdated\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"uint8\"},{\"name\":\"blocknumber\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyRegisteredForQuorums\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BitmapCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BitmapEmpty\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BitmapUpdateIsAfterBlockNumber\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BitmapValueTooLarge\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BytesArrayLengthTooLong\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BytesArrayNotOrdered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotChurnSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotKickOperatorAboveThreshold\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotReregisterYet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChurnApproverSaltUsed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CurrentlyPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpModFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputAddressZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputLengthMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientStakeForChurn\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAVS\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidNewPausedStatus\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRegistrationType\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidShortString\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MaxQuorumsReached\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NextBitmapUpdateIsBeforeBlockNumber\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotRegisteredForQuorum\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotSorted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyAllocationManager\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyEjector\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyPauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyUnpauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"QuorumDoesNotExist\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"QuorumOperatorCountMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SignatureExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StringTooLong\",\"inputs\":[{\"name\":\"str\",\"type\":\"string\",\"internalType\":\"string\"}]}]",
	Bin: "0x610220604052348015610010575f5ffd5b50604051615c3e380380615c3e83398101604081905261002f916102e9565b604080518082018252601681527f4156535265676973747279436f6f7264696e61746f72000000000000000000006020808301919091528251808401909352600683526576302e302e3160d01b908301526001600160a01b0380891660c05280881660a05280871660e05280861660805284166101005282916100b38260336101a3565b6101c0526100c28160346101a3565b6101e0528151602080840191909120610180528151908201206101a0524661014052610151610180516101a051604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60208201529081019290925260608201524660808201523060a08201525f9060c00160405160208183030381529060405280519060200120905090565b61012052505030610160526001600160a01b038116610183576040516339b190bb60e11b815260040160405180910390fd5b6001600160a01b0316610200526101986101d5565b505050505050610516565b5f6020835110156101be576101b783610295565b90506101cf565b816101c98482610404565b5060ff90505b92915050565b603254610100900460ff16156102425760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b60648201526084015b60405180910390fd5b60325460ff90811614610293576032805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b5f5f829050601f815111156102bf578260405163305a27a960e01b815260040161023991906104be565b80516102ca826104f3565b179392505050565b6001600160a01b03811681146102e6575f5ffd5b50565b5f5f5f5f5f5f60c087890312156102fe575f5ffd5b8651610309816102d2565b602088015190965061031a816102d2565b604088015190955061032b816102d2565b606088015190945061033c816102d2565b608088015190935061034d816102d2565b60a088015190925061035e816102d2565b809150509295509295509295565b634e487b7160e01b5f52604160045260245ffd5b600181811c9082168061039457607f821691505b6020821081036103b257634e487b7160e01b5f52602260045260245ffd5b50919050565b601f8211156103ff57805f5260205f20601f840160051c810160208510156103dd5750805b601f840160051c820191505b818110156103fc575f81556001016103e9565b50505b505050565b81516001600160401b0381111561041d5761041d61036c565b6104318161042b8454610380565b846103b8565b6020601f821160018114610463575f831561044c5750848201515b5f19600385901b1c1916600184901b1784556103fc565b5f84815260208120601f198516915b828110156104925787850151825560209485019460019092019101610472565b50848210156104af57868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b805160208083015191908110156103b2575f1960209190910360031b1b16919050565b60805160a05160c05160e05161010051610120516101405161016051610180516101a0516101c0516101e051610200516155de6106605f395f818161075a01528181610c420152818161154b01526119ed01525f6116fc01525f6116d101525f61345b01525f61343301525f61338e01525f6133b801525f6133e201525f818161089601528181612011015281816125ef01526136a501525f81816107b0015281816110480152818161141801528181612343015281816127a9015261303d01525f81816106b1015281816113a601528181611c3c015281816122c80152818161268d0152818161272801528181612f9b01526139ad01525f818161067701528181610da8015281816113e40152818161224f0152818161281e01528181612c3901528181612cb00152612f2201525f818161097a0152818161144d0152611e0901526155de5ff3fe608060405234801561000f575f5ffd5b5060043610610371575f3560e01c80636347c900116101d4578063a65497c611610109578063d72d8dd6116100a9578063f2fde38b11610079578063f2fde38b1461099c578063f698da25146109af578063fabc1cbc146109b7578063fd39105a146109ca575f5ffd5b8063d72d8dd6146108b8578063de1164bb146108c0578063e65797ad146108d3578063ea32afae14610975575f5ffd5b8063c391425e116100e4578063c391425e14610837578063c63fd50214610857578063ca0de8821461086a578063ca8aa7c714610891575f5ffd5b8063a65497c6146107f9578063a96f783e1461080c578063b526578714610815575f5ffd5b806384ca5213116101745780638da5cb5b1161014f5780638da5cb5b1461077c5780639aa1653d1461078d5780639e9923c2146107ab5780639feab859146107d2575f5ffd5b806384ca52131461072f578063871ef04914610742578063886f119514610755575f5ffd5b8063715018a6116101af578063715018a6146106e657806373447992146106ee5780638281ab751461070157806384b0196e14610714575f5ffd5b80636347c9001461069957806368304835146106ac5780636e3b17db146106d3575f5ffd5b806329d1e0c3116102aa57806354fd4d501161024a5780635ac86ab7116102255780635ac86ab7146106345780635b0b829f146106575780635c975abb1461066a5780635df4594614610672575f5ffd5b806354fd4d50146105e45780635865c60c1461060c578063595c6a671461062c575f5ffd5b80633c2a7f4c116102855780633c2a7f4c1461058b5780633eef3a51146105ab5780635140a548146105be578063530b97a4146105d1575f5ffd5b806329d1e0c3146105525780632cdd1e8614610565578063303ca95614610578575f5ffd5b806313542a4e116103155780631eb812da116102f05780631eb812da146104c4578063249a0c421461050d57806328f61b311461052c578063296bb0641461053f575f5ffd5b806313542a4e14610457578063136439dd1461047f5780631478851f14610492575f5ffd5b8063054310e611610350578063054310e6146103e75780630cf4b767146104125780630d3f213414610425578063125e058414610438575f5ffd5b8062cf2ab51461037557806303fd34921461038a57806304ec6351146103bc575b5f5ffd5b6103886103833660046141b0565b610a05565b005b6103a96103983660046141e1565b5f9081526002602052604090205490565b6040519081526020015b60405180910390f35b6103cf6103ca366004614209565b610ba6565b6040516001600160c01b0390911681526020016103b3565b6007546103fa906001600160a01b031681565b6040516001600160a01b0390911681526020016103b3565b6103886104203660046142af565b610bbe565b6103886104333660046141e1565b610c20565b6103a96104463660046142e0565b60096020525f908152604090205481565b6103a96104653660046142e0565b6001600160a01b03165f9081526003602052604090205490565b61038861048d3660046141e1565b610c2d565b6104b46104a03660046141e1565b60046020525f908152604090205460ff1681565b60405190151581526020016103b3565b6104d76104d23660046142fb565b610d02565b60408051825163ffffffff908116825260208085015190911690820152918101516001600160c01b0316908201526060016103b3565b6103a961051b366004614330565b60056020525f908152604090205481565b6008546103fa906001600160a01b031681565b6103fa61054d3660046141e1565b610d90565b6103886105603660046142e0565b610e19565b6103886105733660046142e0565b610e2a565b6103886105863660046143ad565b610e3b565b61059e6105993660046142e0565b610eb7565b6040516103b3919061440a565b6103886105b936600461452b565b610f35565b6103886105cc3660046145d7565b610f51565b6103886105df3660046146b4565b611275565b604080518082018252600681526576302e302e3160d01b602082015290516103b39190614746565b61061f61061a3660046142e0565b6114c4565b6040516103b3919061478c565b610388611536565b6104b4610642366004614330565b603654600160ff9092169190911b9081161490565b6103886106653660046147a7565b6115e5565b6036546103a9565b6103fa7f000000000000000000000000000000000000000000000000000000000000000081565b6103fa6106a73660046141e1565b611601565b6103fa7f000000000000000000000000000000000000000000000000000000000000000081565b6103886106e13660046147d9565b611629565b610388611655565b6103a96106fc3660046142e0565b611666565b61038861070f366004614825565b6116af565b61071c6116c4565b6040516103b39796959493929190614877565b6103a961073d3660046149a1565b61174b565b6103cf6107503660046141e1565b611794565b6103fa7f000000000000000000000000000000000000000000000000000000000000000081565b6099546001600160a01b03166103fa565b5f546107999060ff1681565b60405160ff90911681526020016103b3565b6103fa7f000000000000000000000000000000000000000000000000000000000000000081565b6103a97f2bd82124057f0913bc3b772ce7b83e8057c1ad1f3510fc83778be20f10ec5de681565b6103886108073660046142e0565b61179e565b6103a9600a5481565b6104b46108233660046142e0565b600b546001600160a01b0391821691161490565b61084a610845366004614a06565b6117af565b6040516103b39190614aeb565b610388610865366004614afd565b6117bd565b6103a97f4d404e3276e7ac2163d8ee476afa6a41d1f68fb71f2d8b6546b24e55ce01b72a81565b6103fa7f000000000000000000000000000000000000000000000000000000000000000081565b6006546103a9565b600b546103fa906001600160a01b031681565b6109416108e1366004614330565b60408051606080820183525f808352602080840182905292840181905260ff9490941684526001825292829020825193840183525463ffffffff8116845261ffff600160201b8204811692850192909252600160301b9004169082015290565b60408051825163ffffffff16815260208084015161ffff9081169183019190915292820151909216908201526060016103b3565b6103fa7f000000000000000000000000000000000000000000000000000000000000000081565b6103886109aa3660046142e0565b611967565b6103a96119dd565b6103886109c53660046141e1565b6119eb565b6109f86109d83660046142e0565b6001600160a01b03165f9081526003602052604090206001015460ff1690565b6040516103b39190614b8a565b603654600290600490811603610a2e5760405163840a48d560e01b815260040160405180910390fd5b5f5b8251811015610ba1576040805160018082528183019092525f9160208083019080368337019050509050838281518110610a6c57610a6c614b98565b6020026020010151815f81518110610a8657610a86614b98565b6001600160a01b0392909216602092830291909101909101526040805160018082528183019092525f9181602001602082028036833701905050905060035f868581518110610ad757610ad7614b98565b60200260200101516001600160a01b03166001600160a01b031681526020019081526020015f205f0154815f81518110610b1357610b13614b98565b6020026020010181815250505f610b42825f81518110610b3557610b35614b98565b6020026020010151611b02565b90505f610b57826001600160c01b0316611b0e565b90505f5b8151811015610b9057610b888585848481518110610b7b57610b7b614b98565b016020015160f81c611bd7565b600101610b5b565b505060019093019250610a30915050565b505050565b5f610bb46002858585611d0e565b90505b9392505050565b6001335f9081526003602052604090206001015460ff166002811115610be657610be6614758565b14610c045760405163aba4733960e01b815260040160405180910390fd5b335f90815260036020526040902054610c1d9082611df2565b50565b610c28611e9d565b600a55565b60405163237dfb4760e11b81523360048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906346fbf68e90602401602060405180830381865afa158015610c8f573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610cb39190614bbb565b610cd057604051631d77d47760e21b815260040160405180910390fd5b6036548181168114610cf55760405163c61dca5d60e01b815260040160405180910390fd5b610cfe82611ef7565b5050565b604080516060810182525f80825260208201819052918101919091525f838152600260205260409020805483908110610d3d57610d3d614b98565b5f91825260209182902060408051606081018252919092015463ffffffff8082168352600160201b820416938201939093526001600160c01b03600160401b909304929092169082015290505b92915050565b6040516308f6629d60e31b8152600481018290525f907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906347b314e890602401602060405180830381865afa158015610df5573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610d8a9190614bd4565b610e21611e9d565b610c1d81611f34565b610e32611e9d565b610c1d81611f9d565b610e43612006565b603654600190600290811603610e6c5760405163840a48d560e01b815260040160405180910390fd5b600b546001600160a01b03848116911614610e9a576040516366e565df60e01b815260040160405180910390fd5b5f610ea48361204f565b9050610eb085826120f7565b5050505050565b604080518082019091525f8082526020820152610d8a610f307f2bd82124057f0913bc3b772ce7b83e8057c1ad1f3510fc83778be20f10ec5de684604051602001610f159291909182526001600160a01b0316602082015260400190565b604051602081830303815290604052805190602001206123a5565b6123d1565b610f3d611e9d565b610f4b84848460018561245b565b50505050565b603654600290600490811603610f7a5760405163840a48d560e01b815260040160405180910390fd5b610fbc83838080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201829052505460ff16925061288a915050565b5083518214610fde5760405163aaad13f760e01b815260040160405180910390fd5b5f5b82811015610eb0575f848483818110610ffb57610ffb614b98565b885192013560f81c92505f918891508490811061101a5761101a614b98565b60209081029190910101516040516379a0849160e11b815260ff841660048201529091506001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063f341092290602401602060405180830381865afa15801561108d573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906110b19190614bef565b63ffffffff168151146110d757604051638e5aeee760e01b815260040160405180910390fd5b5f81516001600160401b038111156110f1576110f1614079565b60405190808252806020026020018201604052801561111a578160200160208202803683370190505b5090505f805b8351811015611210575f84828151811061113c5761113c614b98565b6020026020010151905060035f826001600160a01b03166001600160a01b031681526020019081526020015f205f015484838151811061117e5761117e614b98565b6020026020010181815250505f6111a0858481518110610b3557610b35614b98565b905060016001600160c01b03821660ff89161c8116146111d35760405163d053aa2160e01b815260040160405180910390fd5b836001600160a01b0316826001600160a01b0316116112055760405163ba50f91160e01b815260040160405180910390fd5b509150600101611120565b5061121c838386611bd7565b60ff84165f81815260056020908152604091829020439081905591519182527f46077d55330763f16269fd75e5761663f4192d2791747c0189b16ad31db07db4910160405180910390a250505050806001019050610fe0565b603254610100900460ff16158080156112955750603254600160ff909116105b806112af5750303b1580156112af575060325460ff166001145b6113175760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b6032805460ff19166001179055801561133a576032805461ff0019166101001790555b611343866128be565b61134c85611f34565b61135583611ef7565b61135e84611f9d565b6113678261290f565b60068054600181810183555f8390527ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d3f91820180546001600160a01b037f000000000000000000000000000000000000000000000000000000000000000081166001600160a01b03199283161790925584548084018655840180547f0000000000000000000000000000000000000000000000000000000000000000841690831617905584548084018655840180547f000000000000000000000000000000000000000000000000000000000000000084169083161790558454928301909455910180547f00000000000000000000000000000000000000000000000000000000000000009092169190921617905580156114bc576032805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050505050565b604080518082019091525f80825260208201526001600160a01b0382165f908152600360209081526040918290208251808401909352805483526001810154909183019060ff16600281111561151c5761151c614758565b600281111561152d5761152d614758565b90525092915050565b60405163237dfb4760e11b81523360048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906346fbf68e90602401602060405180830381865afa158015611598573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906115bc9190614bbb565b6115d957604051631d77d47760e21b815260040160405180910390fd5b6115e35f19611ef7565b565b6115ed611e9d565b816115f781612931565b610ba18383612959565b60068181548110611610575f80fd5b5f918252602090912001546001600160a01b0316905081565b6116316129fe565b6001600160a01b0382165f908152600960205260409020429055610cfe8282612a29565b61165d611e9d565b6115e35f6128be565b5f610d8a7f2bd82124057f0913bc3b772ce7b83e8057c1ad1f3510fc83778be20f10ec5de683604051602001610f159291909182526001600160a01b0316602082015260400190565b6116b7611e9d565b610ba18383835f5f61245b565b5f606080828080836116f77f00000000000000000000000000000000000000000000000000000000000000006033612ac0565b6117227f00000000000000000000000000000000000000000000000000000000000000006034612ac0565b604080515f80825260208201909252600f60f81b9b939a50919850469750309650945092509050565b5f61178a7f4d404e3276e7ac2163d8ee476afa6a41d1f68fb71f2d8b6546b24e55ce01b72a8787878787604051602001610f1596959493929190614c0a565b9695505050505050565b5f610d8a82611b02565b6117a6611e9d565b610c1d8161290f565b6060610bb760028484612b69565b6117c5612006565b6036545f906001908116036117ed5760405163840a48d560e01b815260040160405180910390fd5b600b546001600160a01b0386811691161461181b576040516366e565df60e01b815260040160405180910390fd5b5f6118258561204f565b90505f808061183686880188614d99565b9250925092505f6118478b83612c18565b90505f84600181111561185c5761185c614758565b03611905575f6118708c8388876001612d46565b5190505f5b86518110156118fe575f87828151811061189157611891614b98565b0160209081015160f81c5f8181526001909252604090912054845191925063ffffffff16908490849081106118c8576118c8614b98565b602002602001015163ffffffff1611156118f557604051633cb89c9760e01b815260040160405180910390fd5b50600101611875565b505061195a565b600184600181111561191957611919614758565b03611941575f8061192c898b018b614df4565b945094505050506118fe8d848988868661318d565b60405163354bb8ab60e01b815260040160405180910390fd5b5050505050505050505050565b61196f611e9d565b6001600160a01b0381166119d45760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b606482015260840161130e565b610c1d816128be565b5f6119e6613382565b905090565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611a47573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611a6b9190614bd4565b6001600160a01b0316336001600160a01b031614611a9c5760405163794821ff60e01b815260040160405180910390fd5b60365480198219811614611ac35760405163c61dca5d60e01b815260040160405180910390fd5b603682905560405182815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c906020015b60405180910390a25050565b5f610d8a6002836134ab565b60605f5f611b1b84613515565b61ffff166001600160401b03811115611b3657611b36614079565b6040519080825280601f01601f191660200182016040528015611b60576020820181803683370190505b5090505f805b825182108015611b77575061010081105b15611bcd576001811b935085841615611bbd578060f81b838381518110611ba057611ba0614b98565b60200101906001600160f81b03191690815f1a9053508160010191505b611bc681614f01565b9050611b66565b5090949350505050565b6040805160018082528183019092525f916020820181803683370190505090508160f81b815f81518110611c0d57611c0d614b98565b60200101906001600160f81b03191690815f1a905350604051636c3fb4bf60e01b81525f906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690636c3fb4bf90611c7590889088908890600401614f19565b5f604051808303815f875af1158015611c90573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052611cb79190810190614fa9565b90505f5b85518110156114bc57818181518110611cd657611cd6614b98565b602002602001015115611d0657611d06868281518110611cf857611cf8614b98565b602002602001015184612a29565b600101611cbb565b5f838152602085905260408120805482919084908110611d3057611d30614b98565b5f91825260209182902060408051606081018252929091015463ffffffff808216808552600160201b8304821695850195909552600160401b9091046001600160c01b03169183019190915290925085161015611da057604051636cb19aff60e01b815260040160405180910390fd5b602081015163ffffffff161580611dc65750806020015163ffffffff168463ffffffff16105b611de35760405163bbba60cb60e01b815260040160405180910390fd5b6040015190505b949350505050565b6040516378219b3f60e11b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063f043367e90611e409085908590600401615036565b5f604051808303815f87803b158015611e57575f5ffd5b505af1158015611e69573d5f5f3e3d5ffd5b50505050817fec2963ab21c1e50e1e582aa542af2e4bf7bf38e6e1403c27b42e1c5d6e621eaa82604051611af69190614746565b6099546001600160a01b031633146115e35760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161130e565b603681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a250565b600754604080516001600160a01b03928316815291831660208301527f315457d8a8fe60f04af17c16e2f5a5e1db612b31648e58030360759ef8f3528c910160405180910390a1600780546001600160a01b0319166001600160a01b0392909216919091179055565b600854604080516001600160a01b03928316815291831660208301527f8f30ab09f43a6c157d7fce7e0a13c003042c1c95e8a72e7a146a21c0caa24dc9910160405180910390a1600880546001600160a01b0319166001600160a01b0392909216919091179055565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146115e3576040516323d871a560e01b815260040160405180910390fd5b60605f82516001600160401b0381111561206b5761206b614079565b6040519080825280601f01601f191660200182016040528015612095576020820181803683370190505b5090505f5b83518110156120f0578381815181106120b5576120b5614b98565b602002602001015160f81b8282815181106120d2576120d2614b98565b60200101906001600160f81b03191690815f1a90535060010161209a565b5092915050565b6001600160a01b0382165f9081526003602052604081208054909161211b82611b02565b905060018084015460ff16600281111561213757612137614758565b146121555760405163aba4733960e01b815260040160405180910390fd5b5f805461216690869060ff1661288a565b90506001600160c01b03811661218f576040516368b6a87560e11b815260040160405180910390fd5b6121a66001600160c01b0382811690841681161490565b6121c35760405163d053aa2160e01b815260040160405180910390fd5b6001600160c01b03818116198316166121dc848261353f565b6001600160c01b038116612238576001600160a01b0387165f81815260036020526040808220600101805460ff19166002179055518692917f396fdcb180cb0fea26928113fb0fd1c3549863f9cd563e6a184f1d578116c8e491a35b60405163f4e24fe560e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063f4e24fe590612286908a908a9060040161504e565b5f604051808303815f87803b15801561229d575f5ffd5b505af11580156122af573d5f5f3e3d5ffd5b505060405163bd29b8cd60e01b81526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016925063bd29b8cd91506123019087908a90600401615036565b5f604051808303815f87803b158015612318575f5ffd5b505af115801561232a573d5f5f3e3d5ffd5b505060405163bd29b8cd60e01b81526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016925063bd29b8cd915061237c9087908a90600401615036565b5f604051808303815f87803b158015612393575f5ffd5b505af115801561195a573d5f5f3e3d5ffd5b5f610d8a6123b1613382565b8360405161190160f01b8152600281019290925260228201526042902090565b604080518082019091525f80825260208201525f80806123fe5f5160206155895f395f51905f5286615085565b90505b61240a8161354b565b90935091505f5160206155895f395f51905f528283098303612442576040805180820190915290815260208101919091529392505050565b5f5160206155895f395f51905f52600182089050612401565b5f5460ff1660c0811061248157604051633cb89c9760e01b815260040160405180910390fd5b5f805460019190819061249890849060ff16615098565b92506101000a81548160ff021916908360ff1602179055506124ba8187612959565b6040805160018082528183019092525f91816020015b604080518082019091525f8152606060208201528152602001906001900390816124d05790505090505f85516001600160401b0381111561251357612513614079565b60405190808252806020026020018201604052801561253c578160200160208202803683370190505b5090505f5b86518110156125995786818151811061255c5761255c614b98565b60200260200101515f015182828151811061257957612579614b98565b6001600160a01b0390921660209283029190910190910152600101612541565b5060405180604001604052808460ff1663ffffffff16815260200182815250825f815181106125ca576125ca614b98565b6020908102919091010152600b54604051630130fc2760e51b81526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000081169263261f84e09261262a92919091169086906004016150b1565b5f604051808303815f87803b158015612641575f5ffd5b505af1158015612653573d5f5f3e3d5ffd5b505f925061265f915050565b85600181111561267157612671614758565b036126f857604051633aea0b9d60e11b81526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906375d4173a906126c69086908b908b906004016151c1565b5f604051808303815f87803b1580156126dd575f5ffd5b505af11580156126ef573d5f5f3e3d5ffd5b50505050612791565b600185600181111561270c5761270c614758565b0361279157604051630662d3e160e51b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063cc5a7c20906127639086908b9089908c906004016151f4565b5f604051808303815f87803b15801561277a575f5ffd5b505af115801561278c573d5f5f3e3d5ffd5b505050505b60405163136ca0f960e11b815260ff841660048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906326d941f2906024015f604051808303815f87803b1580156127f2575f5ffd5b505af1158015612804573d5f5f3e3d5ffd5b505060405163136ca0f960e11b815260ff861660048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031692506326d941f291506024015f604051808303815f87803b158015612869575f5ffd5b505af115801561287b573d5f5f3e3d5ffd5b505050505b5050505050505050565b5f5f612895846135c7565b9050808360ff166001901b11610bb75760405163ca95733360e01b815260040160405180910390fd5b609980546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b600b80546001600160a01b0319166001600160a01b0392909216919091179055565b5f5460ff90811690821610610c1d57604051637310cff560e11b815260040160405180910390fd5b60ff82165f81815260016020908152604091829020845181548684018051888701805163ffffffff90951665ffffffffffff199094168417600160201b61ffff938416021767ffff0000000000001916600160301b95831695909502949094179094558551918252518316938101939093525116918101919091527f3ee6fe8d54610244c3e9d3c066ae4aee997884aa28f10616ae821925401318ac90606001611af6565b6008546001600160a01b031633146115e3576040516376d8ab1760e11b815260040160405180910390fd5b6001600160a01b0382165f9081526003602052604081208054825491929091612a5690859060ff1661288a565b90505f612a6283611b02565b905060018085015460ff166002811115612a7e57612a7e614758565b148015612a9357506001600160c01b03821615155b8015612ab15750612ab16001600160c01b0383811690831681161490565b156114bc576114bc8686613682565b606060ff8314612ada57612ad383613720565b9050610d8a565b818054612ae69061522a565b80601f0160208091040260200160405190810160405280929190818152602001828054612b129061522a565b8015612b5d5780601f10612b3457610100808354040283529160200191612b5d565b820191905f5260205f20905b815481529060010190602001808311612b4057829003601f168201915b50505050509050610d8a565b60605f82516001600160401b03811115612b8557612b85614079565b604051908082528060200260200182016040528015612bae578160200160208202803683370190505b5090505f5b8351811015612c0f57612be08686868481518110612bd357612bd3614b98565b602002602001015161375d565b828281518110612bf257612bf2614b98565b63ffffffff90921660209283029190910190910152600101612bb3565b50949350505050565b6040516309aa152760e11b81526001600160a01b0383811660048301525f917f0000000000000000000000000000000000000000000000000000000000000000909116906313542a4e90602401602060405180830381865afa158015612c80573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612ca49190615262565b90505f819003610d8a577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663bf79ce588484612ce887610eb7565b6040518463ffffffff1660e01b8152600401612d069392919061529b565b6020604051808303815f875af1158015612d22573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610bb79190615262565b612d6a60405180606001604052806060815260200160608152602001606081525090565b5f8054612d7b90869060ff1661288a565b90505f612d8787611b02565b90506001600160c01b038216612db0576040516313ca465760e01b815260040160405180910390fd5b8082166001600160c01b031615612dda57604051630c6816cd60e01b815260040160405180910390fd5b600a546001600160a01b0389165f908152600960205260409020546001600160c01b0383811690851617914291612e119190615314565b10612e2f57604051631968677d60e11b815260040160405180910390fd5b612e39888261353f565b612e438887611df2565b60016001600160a01b038a165f9081526003602052604090206001015460ff166002811115612e7457612e74614758565b14612f0b57604080518082018252898152600160208083018281526001600160a01b038e165f908152600390925293902082518155925183820180549394939192909160ff191690836002811115612ece57612ece614758565b0217905550506040518991506001600160a01b038b16907fe8e68cef1c3a761ed7be7e8463a375f27f7bc335e51824223cacce636ec5c3fe905f90a35b604051631fd93ca960e11b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690633fb2795290612f59908c908b9060040161504e565b5f604051808303815f87803b158015612f70575f5ffd5b505af1158015612f82573d5f5f3e3d5ffd5b5050604051632550477760e01b81526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016925063255047779150612fd6908c908c908c90600401615327565b5f604051808303815f875af1158015612ff1573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261301891908101906153b1565b60408087019190915260208601919091525162bff04d60e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169062bff04d90613073908b908b90600401615036565b5f604051808303815f875af115801561308e573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526130b5919081019061540a565b84528415613181575f5b875181101561317f575f60015f8a84815181106130de576130de614b98565b0160209081015160f81c82528181019290925260409081015f208151606081018352905463ffffffff811680835261ffff600160201b8304811695840195909552600160301b9091049093169181019190915287518051919350908490811061314957613149614b98565b602002602001015163ffffffff16111561317657604051633cb89c9760e01b815260040160405180910390fd5b506001016130bf565b505b50505095945050505050565b83518251146131af5760405163aaad13f760e01b815260040160405180910390fd5b6131bb86868484613875565b5f6131c9878787875f612d46565b90505f5b8551811015612880575f60015f8884815181106131ec576131ec614b98565b0160209081015160f81c82528181019290925260409081015f208151606081018352905463ffffffff811680835261ffff600160201b8304811695840195909552600160301b9091049093169181019190915284518051919350908490811061325757613257614b98565b602002602001015163ffffffff161115613379576132eb87838151811061328057613280614b98565b602001015160f81c60f81b60f81c846040015184815181106132a4576132a4614b98565b60200260200101518b866020015186815181106132c3576132c3614b98565b60200260200101518987815181106132dd576132dd614b98565b602002602001015186613920565b6040805160018082528183019092525f9160208201818036833701905050905087838151811061331d5761331d614b98565b602001015160f81c60f81b815f8151811061333a5761333a614b98565b60200101906001600160f81b03191690815f1a90535061337786848151811061336557613365614b98565b60200260200101516020015182612a29565b505b506001016131cd565b5f306001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161480156133da57507f000000000000000000000000000000000000000000000000000000000000000046145b1561340457507f000000000000000000000000000000000000000000000000000000000000000090565b6119e6604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60208201527f0000000000000000000000000000000000000000000000000000000000000000918101919091527f000000000000000000000000000000000000000000000000000000000000000060608201524660808201523060a08201525f9060c00160405160208183030381529060405280519060200120905090565b5f818152602083905260408120548082036134c9575f915050610d8a565b5f8381526020859052604090206134e1600183615499565b815481106134f1576134f1614b98565b5f91825260209091200154600160401b90046001600160c01b03169150610d8a9050565b5f805b8215610d8a57613529600184615499565b9092169180613537816154ac565b915050613518565b610cfe60028383613aa1565b5f80805f5160206155895f395f51905f5260035f5160206155895f395f51905f52865f5160206155895f395f51905f52888909090890505f6135bb827f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f525f5160206155895f395f51905f52613c5a565b91959194509092505050565b5f610100825111156135ec57604051637da54e4760e11b815260040160405180910390fd5b81515f036135fb57505f919050565b5f5f835f8151811061360f5761360f614b98565b0160200151600160f89190911c81901b92505b84518110156136795784818151811061363d5761363d614b98565b0160200151600160f89190911c1b915082821161366d57604051631019106960e31b815260040160405180910390fd5b91811791600101613622565b50909392505050565b604080516060810182526001600160a01b038481168252600b54811660208301527f00000000000000000000000000000000000000000000000000000000000000001691636e3492b5919081016136d885613cd3565b8152506040518263ffffffff1660e01b81526004016136f791906154cc565b5f604051808303815f87803b15801561370e575f5ffd5b505af11580156114bc573d5f5f3e3d5ffd5b60605f61372c83613d78565b6040805160208082528183019092529192505f91906020820181803683375050509182525060208101929092525090565b5f81815260208490526040812054815b818110156137e05760016137818284615499565b61378b9190615499565b92508463ffffffff16865f8681526020019081526020015f208463ffffffff16815481106137bb576137bb614b98565b5f9182526020909120015463ffffffff16116137d8575050610bb7565b60010161376d565b5060405162461bcd60e51b815260206004820152605c60248201527f5265676973747279436f6f7264696e61746f722e67657451756f72756d42697460448201527f6d6170496e6465784174426c6f636b4e756d6265723a206e6f206269746d617060648201527f2075706461746520666f756e6420666f72206f70657261746f72496400000000608482015260a40161130e565b6020808201515f9081526004909152604090205460ff16156138aa57604051636fbefec360e11b815260040160405180910390fd5b42816040015110156138cf57604051630819bdcd60e01b815260040160405180910390fd5b602080820180515f9081526004909252604091829020805460ff19166001179055600754905191830151610f4b926001600160a01b0390921691613919918891889188919061174b565b8351613d9f565b6020808301516001600160a01b038082165f8181526003909452604090932054919290871603613963576040516356168b4160e11b815260040160405180910390fd5b8760ff16845f015160ff161461398c57604051638e5aeee760e01b815260040160405180910390fd5b604051635401ed2760e01b81526004810182905260ff891660248201525f907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690635401ed2790604401602060405180830381865afa1580156139fa573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613a1e9190615508565b9050613a2a8185613dc7565b6001600160601b0316866001600160601b031611613a5b57604051634c44995d60e01b815260040160405180910390fd5b613a658885613dea565b6001600160601b0316816001600160601b031610613a965760405163b187e86960e01b815260040160405180910390fd5b505050505050505050565b5f8281526020849052604081205490819003613b45575f83815260208581526040808320815160608101835263ffffffff43811682528185018681526001600160c01b03808a16958401958652845460018101865594885295909620915191909201805495519351909416600160401b026001600160401b03938316600160201b0267ffffffffffffffff1990961691909216179390931716919091179055610f4b565b5f838152602085905260408120613b5d600184615499565b81548110613b6d57613b6d614b98565b5f918252602090912001805490915063ffffffff438116911603613bae5780546001600160401b0316600160401b6001600160c01b03851602178155610eb0565b805463ffffffff438116600160201b81810267ffffffff00000000199094169390931784555f8781526020898152604080832081516060810183529485528483018481526001600160c01b03808c1693870193845282546001810184559286529390942094519401805493519151909216600160401b026001600160401b0391861690960267ffffffffffffffff19909316939094169290921717919091169190911790555050505050565b5f5f613c6461403d565b613c6c61405b565b602080825281810181905260408201819052606082018890526080820187905260a082018690528260c08360056107d05a03fa92508280613ca957fe5b5082613cc85760405163d51edae360e01b815260040160405180910390fd5b505195945050505050565b60605f82516001600160401b03811115613cef57613cef614079565b604051908082528060200260200182016040528015613d18578160200160208202803683370190505b5090505f5b83518110156120f057838181518110613d3857613d38614b98565b602001015160f81c60f81b60f81c60ff16828281518110613d5b57613d5b614b98565b63ffffffff90921660209283029190910190910152600101613d1d565b5f60ff8216601f811115610d8a57604051632cd44ac360e21b815260040160405180910390fd5b613daa838383613e03565b610ba157604051638baa579f60e01b815260040160405180910390fd5b60208101515f9061271090613de09061ffff1685615523565b610bb79190615545565b60408101515f9061271090613de09061ffff1685615523565b5f5f5f613e108585613e57565b90925090505f816004811115613e2857613e28614758565b148015613e465750856001600160a01b0316826001600160a01b0316145b8061178a575061178a868686613e99565b5f5f8251604103613e8b576020830151604084015160608501515f1a613e7f87828585613f80565b94509450505050613e92565b505f905060025b9250929050565b5f5f5f856001600160a01b0316631626ba7e60e01b8686604051602401613ec1929190615036565b60408051601f198184030181529181526020820180516001600160e01b03166001600160e01b0319909416939093179092529051613eff9190615572565b5f60405180830381855afa9150503d805f8114613f37576040519150601f19603f3d011682016040523d82523d5f602084013e613f3c565b606091505b5091509150818015613f5057506020815110155b801561178a57508051630b135d3f60e11b90613f759083016020908101908401615262565b149695505050505050565b5f807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0831115613fb557505f90506003614034565b604080515f8082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015614006573d5f5f3e3d5ffd5b5050604051601f1901519150506001600160a01b03811661402e575f60019250925050614034565b91505f90505b94509492505050565b60405180602001604052806001906020820280368337509192915050565b6040518060c001604052806006906020820280368337509192915050565b634e487b7160e01b5f52604160045260245ffd5b604051606081016001600160401b03811182821017156140af576140af614079565b60405290565b604080519081016001600160401b03811182821017156140af576140af614079565b604051601f8201601f191681016001600160401b03811182821017156140ff576140ff614079565b604052919050565b5f6001600160401b0382111561411f5761411f614079565b5060051b60200190565b6001600160a01b0381168114610c1d575f5ffd5b5f82601f83011261414c575f5ffd5b813561415f61415a82614107565b6140d7565b8082825260208201915060208360051b860101925085831115614180575f5ffd5b602085015b838110156141a657803561419881614129565b835260209283019201614185565b5095945050505050565b5f602082840312156141c0575f5ffd5b81356001600160401b038111156141d5575f5ffd5b611dea8482850161413d565b5f602082840312156141f1575f5ffd5b5035919050565b63ffffffff81168114610c1d575f5ffd5b5f5f5f6060848603121561421b575f5ffd5b83359250602084013561422d816141f8565b929592945050506040919091013590565b5f82601f83011261424d575f5ffd5b8135602083015f5f6001600160401b0384111561426c5761426c614079565b50601f8301601f1916602001614281816140d7565b915050828152858383011115614295575f5ffd5b828260208301375f92810160200192909252509392505050565b5f602082840312156142bf575f5ffd5b81356001600160401b038111156142d4575f5ffd5b611dea8482850161423e565b5f602082840312156142f0575f5ffd5b8135610bb781614129565b5f5f6040838503121561430c575f5ffd5b50508035926020909101359150565b803560ff8116811461432b575f5ffd5b919050565b5f60208284031215614340575f5ffd5b610bb78261431b565b5f82601f830112614358575f5ffd5b813561436661415a82614107565b8082825260208201915060208360051b860101925085831115614387575f5ffd5b602085015b838110156141a657803561439f816141f8565b83526020928301920161438c565b5f5f5f606084860312156143bf575f5ffd5b83356143ca81614129565b925060208401356143da81614129565b915060408401356001600160401b038111156143f4575f5ffd5b61440086828701614349565b9150509250925092565b815181526020808301519082015260408101610d8a565b803561ffff8116811461432b575f5ffd5b5f60608284031215614442575f5ffd5b61444a61408d565b90508135614457816141f8565b815261446560208301614421565b602082015261447660408301614421565b604082015292915050565b6001600160601b0381168114610c1d575f5ffd5b5f82601f8301126144a4575f5ffd5b81356144b261415a82614107565b8082825260208201915060208360061b8601019250858311156144d3575f5ffd5b602085015b838110156141a657604081880312156144ef575f5ffd5b6144f76140b5565b813561450281614129565b8152602082013561451281614481565b60208281019190915290845292909201916040016144d8565b5f5f5f5f60c0858703121561453e575f5ffd5b6145488686614432565b9350606085013561455881614481565b925060808501356001600160401b03811115614572575f5ffd5b61457e87828801614495565b92505060a085013561458f816141f8565b939692955090935050565b5f5f83601f8401126145aa575f5ffd5b5081356001600160401b038111156145c0575f5ffd5b602083019150836020828501011115613e92575f5ffd5b5f5f5f604084860312156145e9575f5ffd5b83356001600160401b038111156145fe575f5ffd5b8401601f8101861361460e575f5ffd5b803561461c61415a82614107565b8082825260208201915060208360051b85010192508883111561463d575f5ffd5b602084015b8381101561467d5780356001600160401b0381111561465f575f5ffd5b61466e8b60208389010161413d565b84525060209283019201614642565b50955050505060208401356001600160401b0381111561469b575f5ffd5b6146a78682870161459a565b9497909650939450505050565b5f5f5f5f5f60a086880312156146c8575f5ffd5b85356146d381614129565b945060208601356146e381614129565b935060408601356146f381614129565b925060608601359150608086013561470a81614129565b809150509295509295909350565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b602081525f610bb76020830184614718565b634e487b7160e01b5f52602160045260245ffd5b6003811061478857634e487b7160e01b5f52602160045260245ffd5b9052565b8151815260208083015160408301916120f09084018261476c565b5f5f608083850312156147b8575f5ffd5b6147c18361431b565b91506147d08460208501614432565b90509250929050565b5f5f604083850312156147ea575f5ffd5b82356147f581614129565b915060208301356001600160401b0381111561480f575f5ffd5b61481b8582860161423e565b9150509250929050565b5f5f5f60a08486031215614837575f5ffd5b6148418585614432565b9250606084013561485181614481565b915060808401356001600160401b0381111561486b575f5ffd5b61440086828701614495565b60ff60f81b8816815260e060208201525f61489560e0830189614718565b82810360408401526148a78189614718565b606084018890526001600160a01b038716608085015260a0840186905283810360c0850152845180825260208087019350909101905f5b818110156148fc5783518352602093840193909201916001016148de565b50909b9a5050505050505050505050565b5f82601f83011261491c575f5ffd5b813561492a61415a82614107565b8082825260208201915060208360061b86010192508583111561494b575f5ffd5b602085015b838110156141a65760408188031215614967575f5ffd5b61496f6140b5565b6149788261431b565b8152602082013561498881614129565b6020828101919091529084529290920191604001614950565b5f5f5f5f5f60a086880312156149b5575f5ffd5b85356149c081614129565b94506020860135935060408601356001600160401b038111156149e1575f5ffd5b6149ed8882890161490d565b9598949750949560608101359550608001359392505050565b5f5f60408385031215614a17575f5ffd5b8235614a22816141f8565b915060208301356001600160401b03811115614a3c575f5ffd5b8301601f81018513614a4c575f5ffd5b8035614a5a61415a82614107565b8082825260208201915060208360051b850101925087831115614a7b575f5ffd5b6020840193505b82841015614a9d578335825260209384019390910190614a82565b809450505050509250929050565b5f8151808452602084019350602083015f5b82811015614ae157815163ffffffff16865260209586019590910190600101614abd565b5093949350505050565b602081525f610bb76020830184614aab565b5f5f5f5f5f60808688031215614b11575f5ffd5b8535614b1c81614129565b94506020860135614b2c81614129565b935060408601356001600160401b03811115614b46575f5ffd5b614b5288828901614349565b93505060608601356001600160401b03811115614b6d575f5ffd5b614b798882890161459a565b969995985093965092949392505050565b60208101610d8a828461476c565b634e487b7160e01b5f52603260045260245ffd5b8051801515811461432b575f5ffd5b5f60208284031215614bcb575f5ffd5b610bb782614bac565b5f60208284031215614be4575f5ffd5b8151610bb781614129565b5f60208284031215614bff575f5ffd5b8151610bb7816141f8565b5f60c0820188835260018060a01b038816602084015286604084015260c0606084015280865180835260e0850191506020880192505f5b81811015614c77578351805160ff1684526020908101516001600160a01b03168185015290930192604090920191600101614c41565b50506080840195909552505060a00152949350505050565b80356002811061432b575f5ffd5b5f60408284031215614cad575f5ffd5b614cb56140b5565b823581526020928301359281019290925250919050565b5f82601f830112614cdb575f5ffd5b614ce36140b5565b806040840185811115614cf4575f5ffd5b845b81811015614d0e578035845260209384019301614cf6565b509095945050505050565b5f818303610100811215614d2b575f5ffd5b614d3361408d565b9150614d3f8484614c9d565b8252614d4e8460408501614c9d565b60208301526080607f1982011215614d64575f5ffd5b50614d6d6140b5565b614d7a8460808501614ccc565b8152614d898460c08501614ccc565b6020820152604082015292915050565b5f5f5f6101408486031215614dac575f5ffd5b614db584614c8f565b925060208401356001600160401b03811115614dcf575f5ffd5b614ddb8682870161423e565b925050614deb8560408601614d19565b90509250925092565b5f5f5f5f5f6101808688031215614e09575f5ffd5b614e1286614c8f565b945060208601356001600160401b03811115614e2c575f5ffd5b614e388882890161423e565b945050614e488760408801614d19565b92506101408601356001600160401b03811115614e63575f5ffd5b614e6f8882890161490d565b9250506101608601356001600160401b03811115614e8b575f5ffd5b860160608189031215614e9c575f5ffd5b614ea461408d565b81356001600160401b03811115614eb9575f5ffd5b614ec58a82850161423e565b8252506020828101359082015260409182013591810191909152949793965091945092919050565b634e487b7160e01b5f52601160045260245ffd5b5f60018201614f1257614f12614eed565b5060010190565b606080825284519082018190525f9060208601906080840190835b81811015614f5b5783516001600160a01b0316835260209384019390920191600101614f34565b5050838103602080860191909152865180835291810192508601905f5b81811015614f96578251845260209384019390920191600101614f78565b50505060ff841660408401529050611dea565b5f60208284031215614fb9575f5ffd5b81516001600160401b03811115614fce575f5ffd5b8201601f81018413614fde575f5ffd5b8051614fec61415a82614107565b8082825260208201915060208360051b85010192508683111561500d575f5ffd5b6020840193505b8284101561178a5761502584614bac565b825260209384019390910190615014565b828152604060208201525f610bb46040830184614718565b6001600160a01b03831681526040602082018190525f90610bb490830184614718565b634e487b7160e01b5f52601260045260245ffd5b5f8261509357615093615071565b500690565b60ff8181168382160190811115610d8a57610d8a614eed565b5f6040820160018060a01b03851683526040602084015280845180835260608501915060608160051b8601019250602086015f5b8281101561516657868503605f190184528151805163ffffffff168652602090810151604082880181905281519088018190529101905f9060608801905b8083101561514e5783516001600160a01b031682526020938401936001939093019290910190615123565b509650505060209384019391909101906001016150e5565b5092979650505050505050565b5f8151808452602084019350602083015f5b82811015614ae157815180516001600160a01b031687526020908101516001600160601b03168188015260409096019590910190600101615185565b60ff841681526001600160601b0383166020820152606060408201525f6151eb6060830184615173565b95945050505050565b60ff851681526001600160601b038416602082015263ffffffff83166040820152608060608201525f61178a6080830184615173565b600181811c9082168061523e57607f821691505b60208210810361525c57634e487b7160e01b5f52602260045260245ffd5b50919050565b5f60208284031215615272575f5ffd5b5051919050565b805f5b6002811015610f4b57815184526020938401939091019060010161527c565b6001600160a01b038416815282518051602080840191909152015160408201526101608101602084810151805160608501529081015160808401525060408401516152ea60a084018251615279565b602001516152fb60e0840182615279565b5082516101208301526020830151610140830152611dea565b80820180821115610d8a57610d8a614eed565b60018060a01b0384168152826020820152606060408201525f6151eb6060830184614718565b5f82601f83011261535c575f5ffd5b815161536a61415a82614107565b8082825260208201915060208360051b86010192508583111561538b575f5ffd5b602085015b838110156141a65780516153a381614481565b835260209283019201615390565b5f5f604083850312156153c2575f5ffd5b82516001600160401b038111156153d7575f5ffd5b6153e38582860161534d565b92505060208301516001600160401b038111156153fe575f5ffd5b61481b8582860161534d565b5f6020828403121561541a575f5ffd5b81516001600160401b0381111561542f575f5ffd5b8201601f8101841361543f575f5ffd5b805161544d61415a82614107565b8082825260208201915060208360051b85010192508683111561546e575f5ffd5b6020840193505b8284101561178a578351615488816141f8565b825260209384019390910190615475565b81810381811115610d8a57610d8a614eed565b5f61ffff821661ffff81036154c3576154c3614eed565b60010192915050565b602080825282516001600160a01b039081168383015290830151166040808301919091528201516060808301525f90611dea6080840182614aab565b5f60208284031215615518575f5ffd5b8151610bb781614481565b6001600160601b0381811683821602908116908181146120f0576120f0614eed565b5f6001600160601b0383168061555d5761555d615071565b806001600160601b0384160491505092915050565b5f82518060208501845e5f92019182525091905056fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47a2646970667358221220c9d375a14de7aa5f8dc1985a9018a338de19ebef58a98e4424280fe9c775906f64736f6c634300081b0033",
}

// ContractSlashingRegistryCoordinatorABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractSlashingRegistryCoordinatorMetaData.ABI instead.
var ContractSlashingRegistryCoordinatorABI = ContractSlashingRegistryCoordinatorMetaData.ABI

// ContractSlashingRegistryCoordinatorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractSlashingRegistryCoordinatorMetaData.Bin instead.
var ContractSlashingRegistryCoordinatorBin = ContractSlashingRegistryCoordinatorMetaData.Bin

// DeployContractSlashingRegistryCoordinator deploys a new Ethereum contract, binding an instance of ContractSlashingRegistryCoordinator to it.
func DeployContractSlashingRegistryCoordinator(auth *bind.TransactOpts, backend bind.ContractBackend, _stakeRegistry common.Address, _blsApkRegistry common.Address, _indexRegistry common.Address, _socketRegistry common.Address, _allocationManager common.Address, _pauserRegistry common.Address) (common.Address, *types.Transaction, *ContractSlashingRegistryCoordinator, error) {
	parsed, err := ContractSlashingRegistryCoordinatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractSlashingRegistryCoordinatorBin), backend, _stakeRegistry, _blsApkRegistry, _indexRegistry, _socketRegistry, _allocationManager, _pauserRegistry)
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

	NumRegistries(opts *bind.CallOpts) (*big.Int, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	Paused(opts *bind.CallOpts, index uint8) (bool, error)

	Paused0(opts *bind.CallOpts) (*big.Int, error)

	PauserRegistry(opts *bind.CallOpts) (common.Address, error)

	PubkeyRegistrationMessageHash(opts *bind.CallOpts, operator common.Address) (BN254G1Point, error)

	QuorumCount(opts *bind.CallOpts) (uint8, error)

	QuorumUpdateBlockNumber(opts *bind.CallOpts, arg0 uint8) (*big.Int, error)

	Registries(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error)

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

	Initialize(opts *bind.TransactOpts, _initialOwner common.Address, _churnApprover common.Address, _ejector common.Address, _initialPausedStatus *big.Int, _avs common.Address) (*types.Transaction, error)

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
	FilterChurnApproverUpdated(opts *bind.FilterOpts) (*ContractSlashingRegistryCoordinatorChurnApproverUpdatedIterator, error)
	WatchChurnApproverUpdated(opts *bind.WatchOpts, sink chan<- *ContractSlashingRegistryCoordinatorChurnApproverUpdated) (event.Subscription, error)
	ParseChurnApproverUpdated(log types.Log) (*ContractSlashingRegistryCoordinatorChurnApproverUpdated, error)

	FilterEIP712DomainChanged(opts *bind.FilterOpts) (*ContractSlashingRegistryCoordinatorEIP712DomainChangedIterator, error)
	WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *ContractSlashingRegistryCoordinatorEIP712DomainChanged) (event.Subscription, error)
	ParseEIP712DomainChanged(log types.Log) (*ContractSlashingRegistryCoordinatorEIP712DomainChanged, error)

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

// NumRegistries is a free data retrieval call binding the contract method 0xd72d8dd6.
//
// Solidity: function numRegistries() view returns(uint256)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCaller) NumRegistries(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractSlashingRegistryCoordinator.contract.Call(opts, &out, "numRegistries")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumRegistries is a free data retrieval call binding the contract method 0xd72d8dd6.
//
// Solidity: function numRegistries() view returns(uint256)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) NumRegistries() (*big.Int, error) {
	return _ContractSlashingRegistryCoordinator.Contract.NumRegistries(&_ContractSlashingRegistryCoordinator.CallOpts)
}

// NumRegistries is a free data retrieval call binding the contract method 0xd72d8dd6.
//
// Solidity: function numRegistries() view returns(uint256)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCallerSession) NumRegistries() (*big.Int, error) {
	return _ContractSlashingRegistryCoordinator.Contract.NumRegistries(&_ContractSlashingRegistryCoordinator.CallOpts)
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

// Registries is a free data retrieval call binding the contract method 0x6347c900.
//
// Solidity: function registries(uint256 ) view returns(address)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCaller) Registries(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ContractSlashingRegistryCoordinator.contract.Call(opts, &out, "registries", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Registries is a free data retrieval call binding the contract method 0x6347c900.
//
// Solidity: function registries(uint256 ) view returns(address)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) Registries(arg0 *big.Int) (common.Address, error) {
	return _ContractSlashingRegistryCoordinator.Contract.Registries(&_ContractSlashingRegistryCoordinator.CallOpts, arg0)
}

// Registries is a free data retrieval call binding the contract method 0x6347c900.
//
// Solidity: function registries(uint256 ) view returns(address)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCallerSession) Registries(arg0 *big.Int) (common.Address, error) {
	return _ContractSlashingRegistryCoordinator.Contract.Registries(&_ContractSlashingRegistryCoordinator.CallOpts, arg0)
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
// Solidity: function version() pure returns(string)
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
// Solidity: function version() pure returns(string)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) Version() (string, error) {
	return _ContractSlashingRegistryCoordinator.Contract.Version(&_ContractSlashingRegistryCoordinator.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
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
// Solidity: function initialize(address _initialOwner, address _churnApprover, address _ejector, uint256 _initialPausedStatus, address _avs) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorTransactor) Initialize(opts *bind.TransactOpts, _initialOwner common.Address, _churnApprover common.Address, _ejector common.Address, _initialPausedStatus *big.Int, _avs common.Address) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.contract.Transact(opts, "initialize", _initialOwner, _churnApprover, _ejector, _initialPausedStatus, _avs)
}

// Initialize is a paid mutator transaction binding the contract method 0x530b97a4.
//
// Solidity: function initialize(address _initialOwner, address _churnApprover, address _ejector, uint256 _initialPausedStatus, address _avs) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) Initialize(_initialOwner common.Address, _churnApprover common.Address, _ejector common.Address, _initialPausedStatus *big.Int, _avs common.Address) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.Contract.Initialize(&_ContractSlashingRegistryCoordinator.TransactOpts, _initialOwner, _churnApprover, _ejector, _initialPausedStatus, _avs)
}

// Initialize is a paid mutator transaction binding the contract method 0x530b97a4.
//
// Solidity: function initialize(address _initialOwner, address _churnApprover, address _ejector, uint256 _initialPausedStatus, address _avs) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorTransactorSession) Initialize(_initialOwner common.Address, _churnApprover common.Address, _ejector common.Address, _initialPausedStatus *big.Int, _avs common.Address) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.Contract.Initialize(&_ContractSlashingRegistryCoordinator.TransactOpts, _initialOwner, _churnApprover, _ejector, _initialPausedStatus, _avs)
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
