// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractAwesomeVaultTaskManager

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

// BN254G2Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G2Point struct {
	X [2]*big.Int
	Y [2]*big.Int
}

// IAwesomeVaultTaskManagerTask is an auto generated low-level Go binding around an user-defined struct.
type IAwesomeVaultTaskManagerTask struct {
	Input                     IAwesomeVaultTaskManagerTaskInput
	TaskCreatedBlock          uint32
	QuorumNumbers             []byte
	QuorumThresholdPercentage uint32
}

// IAwesomeVaultTaskManagerTaskInput is an auto generated low-level Go binding around an user-defined struct.
type IAwesomeVaultTaskManagerTaskInput struct {
	Key   string
	Value string
}

// IAwesomeVaultTaskManagerTaskResponse is an auto generated low-level Go binding around an user-defined struct.
type IAwesomeVaultTaskManagerTaskResponse struct {
	ReferenceTaskIndex uint32
	Result             [32]byte
}

// IAwesomeVaultTaskManagerTaskResponseMetadata is an auto generated low-level Go binding around an user-defined struct.
type IAwesomeVaultTaskManagerTaskResponseMetadata struct {
	TaskRespondedBlock uint32
	HashOfNonSigners   [32]byte
}

// IBLSSignatureCheckerTypesNonSignerStakesAndSignature is an auto generated low-level Go binding around an user-defined struct.
type IBLSSignatureCheckerTypesNonSignerStakesAndSignature struct {
	NonSignerQuorumBitmapIndices []uint32
	NonSignerPubkeys             []BN254G1Point
	QuorumApks                   []BN254G1Point
	ApkG2                        BN254G2Point
	Sigma                        BN254G1Point
	QuorumApkIndices             []uint32
	TotalStakeIndices            []uint32
	NonSignerStakeIndices        [][]uint32
}

// IBLSSignatureCheckerTypesQuorumStakeTotals is an auto generated low-level Go binding around an user-defined struct.
type IBLSSignatureCheckerTypesQuorumStakeTotals struct {
	SignedStakeForQuorum []*big.Int
	TotalStakeForQuorum  []*big.Int
}

// OperatorStateRetrieverCheckSignaturesIndices is an auto generated low-level Go binding around an user-defined struct.
type OperatorStateRetrieverCheckSignaturesIndices struct {
	NonSignerQuorumBitmapIndices []uint32
	QuorumApkIndices             []uint32
	TotalStakeIndices            []uint32
	NonSignerStakeIndices        [][]uint32
}

// OperatorStateRetrieverOperator is an auto generated low-level Go binding around an user-defined struct.
type OperatorStateRetrieverOperator struct {
	Operator   common.Address
	OperatorId [32]byte
	Stake      *big.Int
}

// ContractAwesomeVaultTaskManagerMetaData contains all meta data concerning the ContractAwesomeVaultTaskManager contract.
var ContractAwesomeVaultTaskManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractISlashingRegistryCoordinator\"},{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"_taskResponseWindowBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"TASK_CHALLENGE_WINDOW_BLOCK\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TASK_RESPONSE_WINDOW_BLOCK\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"WADS_TO_SLASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"aggregator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allTaskHashes\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allTaskResponses\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allocationManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"blsApkRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBLSApkRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkSignatures\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureCheckerTypes.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureCheckerTypes.QuorumStakeTotals\",\"components\":[{\"name\":\"signedStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"},{\"name\":\"totalStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"}]},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createNewTask\",\"inputs\":[{\"name\":\"input\",\"type\":\"tuple\",\"internalType\":\"structIAwesomeVaultTaskManager.TaskInput\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"generator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBatchOperatorFromId\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractISlashingRegistryCoordinator\"},{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBatchOperatorId\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractISlashingRegistryCoordinator\"},{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCheckSignaturesIndices\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractISlashingRegistryCoordinator\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"nonSignerOperatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOperatorStateRetriever.CheckSignaturesIndices\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractISlashingRegistryCoordinator\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractISlashingRegistryCoordinator\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumBitmapsAtBlockNumber\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractISlashingRegistryCoordinator\"},{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTaskResponseWindowBlock\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_aggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_generator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_allocationManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_slasher\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_serviceManager\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"instantSlasher\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"latestTaskNum\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"raiseAndResolveChallenge\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structIAwesomeVaultTaskManager.Task\",\"components\":[{\"name\":\"input\",\"type\":\"tuple\",\"internalType\":\"structIAwesomeVaultTaskManager.TaskInput\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"internalType\":\"structIAwesomeVaultTaskManager.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"result\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"taskResponseMetadata\",\"type\":\"tuple\",\"internalType\":\"structIAwesomeVaultTaskManager.TaskResponseMetadata\",\"components\":[{\"name\":\"taskRespondedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"pubkeysOfNonSigningOperators\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registryCoordinator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractISlashingRegistryCoordinator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"respondToTask\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structIAwesomeVaultTaskManager.Task\",\"components\":[{\"name\":\"input\",\"type\":\"tuple\",\"internalType\":\"structIAwesomeVaultTaskManager.TaskInput\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"internalType\":\"structIAwesomeVaultTaskManager.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"result\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"nonSignerStakesAndSignature\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureCheckerTypes.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"serviceManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setStaleStakesForbidden\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakeRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"staleStakesForbidden\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"taskNumber\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"taskSuccesfullyChallenged\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"trySignatureAndApkVerification\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"apk\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"pairingSuccessful\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"siganatureIsValid\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewTaskCreated\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"task\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIAwesomeVaultTaskManager.Task\",\"components\":[{\"name\":\"input\",\"type\":\"tuple\",\"internalType\":\"structIAwesomeVaultTaskManager.TaskInput\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaleStakesForbiddenUpdate\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskChallengedSuccessfully\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"challenger\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskChallengedUnsuccessfully\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"challenger\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskCompleted\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskResponded\",\"inputs\":[{\"name\":\"taskResponse\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIAwesomeVaultTaskManager.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"result\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"taskResponseMetadata\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIAwesomeVaultTaskManager.TaskResponseMetadata\",\"components\":[{\"name\":\"taskRespondedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BitmapValueTooLarge\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BytesArrayLengthTooLong\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BytesArrayNotOrdered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CurrentlyPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ECAddFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ECMulFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpModFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputAddressZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputArrayLengthMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputEmptyQuorumNumbers\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputNonSignerLengthMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidBLSPairingKey\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidBLSSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidNewPausedStatus\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidQuorumApkHash\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidReferenceBlocknumber\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonSignerPubkeysNotSorted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyPauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyRegistryCoordinatorOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyUnpauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorNotRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ScalarTooLarge\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StaleStakesForbidden\",\"inputs\":[]}]",
	Bin: "0x61014080604052346101d857606081614b2f80380380916100208285610283565b8339810103126101d85780516001600160a01b038116908181036101d85760208301516001600160a01b038116938482036101d857604001519363ffffffff851685036101d857156102745760805260a052604051636830483560e01b8152602081600481855afa9081156101e4575f91610231575b5060c052604051632efa2ca360e11b815290602090829060049082905afa9081156101e4575f916101ef575b5060e05260c05160405163df5cf72360e01b815290602090829060049082906001600160a01b03165afa9081156101e4575f9161019e575b50610100526101205260405161487490816102bb82396080518181816102ca01528181610d6e015281816114f00152611f65015260a051818181610963015281816110c00152818161309b015281816131840152613716015260c05181818161107c0152818161348601526135d1015260e05181818161103801526133c2015261010051818181611ccc0152613297015261012051818181610595015261186c0152f35b90506020813d6020116101dc575b816101b960209383610283565b810103126101d857516001600160a01b03811681036101d8575f6100fa565b5f80fd5b3d91506101ac565b6040513d5f823e3d90fd5b90506020813d602011610229575b8161020a60209383610283565b810103126101d857516001600160a01b03811681036101d8575f6100c2565b3d91506101fd565b90506020813d60201161026c575b8161024c60209383610283565b810103126101d857516001600160a01b03811681036101d8576004610096565b3d915061023f565b6339b190bb60e11b5f5260045ffd5b601f909101601f19168101906001600160401b038211908210176102a657604052565b634e487b7160e01b5f52604160045260245ffdfe60806040526004361015610011575f80fd5b5f3560e01c8063136439dd1461029a578063171f1d5b146102955780631ad43189146101e6578063245a7bfc146102905780632cb223d51461028b5780632d89f6fc1461028657806331b36bd9146102815780633563b0d11461027c5780633998fdd314610277578063416c7e5e146102725780634d2b57fe1461026d5780634f739f7414610268578063595c6a67146102635780635a2d7f021461025e5780635ac86ab7146102595780635c155662146102545780635c975abb1461024f5780635decc3f51461024a5780635df459461461024557806368304835146102405780636d14a9871461023b5780636efb463614610236578063715018a61461023157806372d18e8d146102225780637afa1eed1461022c578063886f1195146102275780638b00ce7c146102225780638da5cb5b1461021d57806392a51824146102185780639b290e9814610213578063b70239491461020e578063b98d090814610209578063ca8aa7c714610204578063cc2a9a5b146101ff578063cefdc1d4146101fa578063df5cf723146101f5578063ef029dbc146101f0578063f2fde38b146101eb578063f5c9899d146101e6578063f63c5bab146101e15763fabc1cbc146101dc575f80fd5b611f3c565b611f21565b610579565b611e90565b611cfb565b611cb7565b611b73565b611a33565b611a0b565b6119e9565b61174f565b611727565b611574565b61151f565b611490565b6114db565b6114b3565b611435565b611388565b6110ab565b611067565b611023565b610fe5565b610fc8565b610e4f565b610de3565b610db6565b610d43565b610c9c565b610a91565b610931565b6108ff565b610885565b6106db565b610633565b6105fa565b6105b9565b610511565b3461035a57602036600319011261035a5760043560405163237dfb4760e11b8152336004820152906020826024817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa9182156103555761032492610310915f91610326575b5061203e565b61031f60665482811614612054565b613fcd565b005b610348915060203d60201161034e575b61034081836103ad565b81019061201e565b5f61030a565b503d610336565b612033565b5f80fd5b634e487b7160e01b5f52604160045260245ffd5b604081019081106001600160401b0382111761038d57604052565b61035e565b608081019081106001600160401b0382111761038d57604052565b90601f801991011681019081106001600160401b0382111761038d57604052565b604051906103de610100836103ad565b565b604051906103de6040836103ad565b604051906103de6060836103ad565b906103de60405192836103ad565b60409060e319011261035a576040519061042582610372565b60e4358252610104356020830152565b919082604091031261035a5760405161044d81610372565b6020808294803584520135910152565b9080601f8301121561035a57604051916104786040846103ad565b82906040810192831161035a57905b8282106104945750505090565b8135815260209182019101610487565b90608060631983011261035a576040516104bd81610372565b60206104d882946104cf81606461045d565b845260a461045d565b910152565b919060808382031261035a5760206104d8604051926104fb84610372565b60408496610509838261045d565b86520161045d565b3461035a5761012036600319011261035a57600435604036602319011261035a57610569604091825161054381610372565b60243581526044356020820152610559366104a4565b906105633661040c565b926120a8565b8251911515825215156020820152f35b3461035a575f36600319011261035a57602060405163ffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b3461035a575f36600319011261035a5760cd546040516001600160a01b039091168152602090f35b63ffffffff81160361035a57565b35906103de826105e1565b3461035a57602036600319011261035a5763ffffffff60043561061c816105e1565b165f5260cb602052602060405f2054604051908152f35b3461035a57602036600319011261035a5763ffffffff600435610655816105e1565b165f5260ca602052602060405f2054604051908152f35b6001600160a01b0381160361035a57565b6001600160401b03811161038d5760051b60200190565b90602080835192838152019201905f5b8181106106b15750505090565b82518452602093840193909201916001016106a4565b9060206106d8928181520190610694565b90565b3461035a57604036600319011261035a576004356106f88161066c565b602435906001600160401b03821161035a573660238301121561035a578160040135916107248361067d565b9261073260405194856103ad565b8084526024602085019160051b8301019136831161035a57602401905b8282106107735761076f6107638686612212565b604051918291826106c7565b0390f35b6020809183356107828161066c565b81520191019061074f565b6001600160401b03811161038d57601f01601f191660200190565b9291926107b48261078d565b916107c260405193846103ad565b82948184528183011161035a578281602093845f960137010152565b9080602083519182815201916020808360051b8301019401925f915b83831061080957505050505090565b9091929394601f19828203018352855190602080835192838152019201905f905b80821061084957505050602080600192970193019301919392906107fa565b909192602060606001926001600160601b0360408851868060a01b0381511684528581015186850152015116604082015201940192019061082a565b3461035a57606036600319011261035a576004356108a28161066c565b6024356001600160401b03811161035a573660238201121561035a5761076f916108d96108eb9236906024816004013591016107a8565b604435916108e6836105e1565b612450565b6040519182916020835260208301906107de565b3461035a575f36600319011261035a5760d1546040516001600160a01b039091168152602090f35b8015150361035a57565b3461035a57602036600319011261035a5760043561094e81610927565b604051638da5cb5b60e01b81526020816004817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa908115610355575f916109c3575b506001600160a01b031633036109b45761032490614460565b637070f3b160e11b5f5260045ffd5b6109e5915060203d6020116109eb575b6109dd81836103ad565b8101906122d7565b5f61099b565b503d6109d3565b9080601f8301121561035a578135610a098161067d565b92610a1760405194856103ad565b81845260208085019260051b82010192831161035a57602001905b828210610a3f5750505090565b8135815260209182019101610a32565b60206040818301928281528451809452019201905f5b818110610a725750505090565b82516001600160a01b0316845260209384019390920191600101610a65565b3461035a57604036600319011261035a57600435610aae8161066c565b6024356001600160401b03811161035a57610acd9036906004016109f2565b610ad781516121b0565b916001600160a01b03165f5b8251811015610b7457806020610afc610b1c93866121ef565b5160405180948192630a5aec1960e21b8352600483019190602083019252565b0381865afa91821561035557600192610b50915f91610b56575b50610b4183886121ef565b6001600160a01b039091169052565b01610ae3565b610b6e915060203d81116109eb576109dd81836103ad565b5f610b36565b6040518061076f8682610a4f565b9181601f8401121561035a578235916001600160401b03831161035a576020838186019501011161035a57565b90602080835192838152019201905f5b818110610bcc5750505090565b825163ffffffff16845260209384019390920191600101610bbf565b90602082526060610c36610c21610c0b84516080602088015260a0870190610baf565b6020850151868203601f19016040880152610baf565b6040840151858203601f190184870152610baf565b910151916080601f1982840301910152815180825260208201916020808360051b8301019401925f915b838310610c6f57505050505090565b9091929394602080610c8d600193601f198682030187528951610baf565b97019301930191939290610c60565b3461035a57608036600319011261035a57600435610cb98161066c565b60243590610cc6826105e1565b6044356001600160401b03811161035a57610ce5903690600401610b82565b91606435926001600160401b03841161035a573660238501121561035a578360040135926001600160401b03841161035a573660248560051b8701011161035a5761076f956024610d37960193612964565b60405191829182610be8565b3461035a575f36600319011261035a5760405163237dfb4760e11b81523360048201526020816024817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa801561035557610dae915f91610326575061203e565b610324613f99565b3461035a575f36600319011261035a57602060405167016345785d8a00008152f35b60ff81160361035a57565b3461035a57602036600319011261035a576020600160ff600435610e0681610dd8565b161b806066541614604051908152f35b60206040818301928281528451809452019201905f5b818110610e395750505090565b8251845260209384019390920191600101610e2c565b3461035a57606036600319011261035a57600435610e6c8161066c565b6024356001600160401b03811161035a57610e8b9036906004016109f2565b60443591610e98836105e1565b6040516361c8a12f60e11b8152906001600160a01b03165f8280610ec0868860048401612dcc565b0381845afa918215610355575f92610fa4575b50610ede83516121b0565b935f5b8451811015610f9657610ef481866121ef565b5190602083610f10610f0684896121ef565b5163ffffffff1690565b6040516304ec635160e01b8152600481019590955263ffffffff918216602486015216604484015282606481875afa8015610355576001925f91610f68575b50828060c01b0316610f6182896121ef565b5201610ee1565b610f89915060203d8111610f8f575b610f8181836103ad565b8101906128db565b5f610f4f565b503d610f77565b6040518061076f8882610e16565b610fc19192503d805f833e610fb981836103ad565b8101906127aa565b905f610ed3565b3461035a575f36600319011261035a576020606654604051908152f35b3461035a57602036600319011261035a5763ffffffff600435611007816105e1565b165f5260cc602052602060ff60405f2054166040519015158152f35b3461035a575f36600319011261035a576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b3461035a575f36600319011261035a576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b3461035a575f36600319011261035a576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b9080601f8301121561035a5781356111068161067d565b9261111460405194856103ad565b81845260208085019260051b82010192831161035a57602001905b82821061113c5750505090565b60208091833561114b816105e1565b81520191019061112f565b81601f8201121561035a57803561116c8161067d565b9261117a60405194856103ad565b81845260208085019260061b8401019281841161035a57602001915b8383106111a4575050505090565b60206040916111b38486610435565b815201920191611196565b9080601f8301121561035a5781356111d58161067d565b926111e360405194856103ad565b81845260208085019260051b8201019183831161035a5760208201905b83821061120f57505050505090565b81356001600160401b03811161035a57602091611231878480948801016110ef565b815201910190611200565b9190916101808184031261035a576112526103ce565b9281356001600160401b03811161035a578161126f9184016110ef565b845260208201356001600160401b03811161035a5781611290918401611156565b602085015260408201356001600160401b03811161035a57816112b4918401611156565b60408501526112c681606084016104dd565b60608501526112d88160e08401610435565b60808501526101208201356001600160401b03811161035a57816112fd9184016110ef565b60a08501526101408201356001600160401b03811161035a57816113229184016110ef565b60c08501526101608201356001600160401b03811161035a5761134592016111be565b60e0830152565b90602080835192838152019201905f5b8181106113695750505090565b82516001600160601b031684526020938401939092019160010161135c565b3461035a57608036600319011261035a576004356024356001600160401b03811161035a576113bb903690600401610b82565b90916044356113c9816105e1565b606435926001600160401b03841161035a5761142b946113f06113f695369060040161123c565b93612fbe565b6040519283926040845260206114178251604080880152608087019061134c565b910151848203603f1901606086015261134c565b9060208301520390f35b3461035a575f36600319011261035a5761144d61463c565b603380546001600160a01b031981169091555f906001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a3005b3461035a575f36600319011261035a57602063ffffffff60c95416604051908152f35b3461035a575f36600319011261035a5760ce546040516001600160a01b039091168152602090f35b3461035a575f36600319011261035a576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b3461035a575f36600319011261035a576033546040516001600160a01b039091168152602090f35b9081608091031261035a5790565b604090602319011261035a57602490565b9081604091031261035a5790565b3461035a5760c036600319011261035a576004356001600160401b03811161035a576115a4903690600401611547565b506115ae36611555565b604036606319011261035a5760a4356001600160401b03811161035a576115d9903690600401611156565b508035906115e6826105e1565b6115fe8263ffffffff165f5260cb60205260405f2090565b54156116d857611650906116208363ffffffff165f5260cb60205260405f2090565b5490604051611647816116396064602083019586613883565b03601f1981018352826103ad565b519020146138a1565b61167b61167561166e8363ffffffff165f5260cc60205260405f2090565b5460ff1690565b15613913565b6116aa61169c61169361168e606461385b565b6139a9565b63ffffffff1690565b63ffffffff431611156139db565b63ffffffff3391167ffd3e26beeb5967fc5a57a0446914eabc45b4aa474c67a51b4b5160cac60ddb055f80a3005b60405162461bcd60e51b815260206004820152602160248201527f5461736b206861736e2774206265656e20726573706f6e64656420746f2079656044820152601d60fa1b6064820152608490fd5b3461035a575f36600319011261035a5760cf546040516001600160a01b039091168152602090f35b3461035a57608036600319011261035a576004356001600160401b03811161035a5761177f903690600401611547565b61178836611555565b906064356001600160401b03811161035a576117a890369060040161123c565b60cd549092906001600160a01b031633036119a4576117cb60208394930161385b565b916118b76117dc6040860186613a4d565b92909461182e6117ee6060890161385b565b9760405161180481611639602082019485613ad4565b5190206118276118138861385b565b63ffffffff165f5260ca60205260405f2090565b5414613b8c565b61185861185161183d8761385b565b63ffffffff165f5260cb60205260405f2090565b5415613bfe565b8363ffffffff4316966118996118916116937f0000000000000000000000000000000000000000000000000000000000000000866139c1565b891115613c5f565b60405160208101906118af816116398b85613cc1565b519020612fbe565b919060ff5f9616955b828110611942577ff2af11fad73d4349c99cf62f298d337641ea0bb7c0f5a8db92a98a275f734f588686866119026118f66103e0565b63ffffffff9094168452565b6020830152604051602081019061191e81611639868686613d91565b51902061192d61183d8361385b565b5561193d60405192839283613d91565b0390a1005b8061199e61197a61197561196961195c60019688516121ef565b516001600160601b031690565b6001600160601b031690565b613cd1565b6119976119698b61199261195c8760208b01516121ef565b613cfd565b1115613d20565b016118c0565b60405162461bcd60e51b815260206004820152601d60248201527f41676772656761746f72206d757374206265207468652063616c6c65720000006044820152606490fd5b3461035a575f36600319011261035a57602060ff609754166040519015158152f35b3461035a575f36600319011261035a5760d0546040516001600160a01b039091168152602090f35b3461035a5760c036600319011261035a57600435611a508161066c565b611ad0602435611a5f8161066c565b604435611a6b8161066c565b606435611a778161066c565b60843591611a848361066c565b60a43593611a918561066c565b5f5496611ab660ff60088a901c16158099819a611b4e575b8115611b2e575b50613dbb565b87611ac7600160ff195f5416175f55565b611b1757613e1e565b611ad657005b611ae461ff00195f54165f55565b604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249890806020810161193d565b611b2961010061ff00195f5416175f55565b613e1e565b303b15915081611b40575b505f611ab0565b60ff1660011490505f611b39565b600160ff8216109150611aa9565b6040906106d89392815281602082015201906107de565b3461035a57606036600319011261035a57600435611b908161066c565b602435604435611b9f816105e1565b611be0611baa61218e565b9280611bb5856121e2565b526040516361c8a12f60e11b81526001600160a01b0386169490925f91849182918760048401612dcc565b0381875afa9384156103555783611c0a611693610f06611c3f986020975f91611c9d575b506121e2565b92604051968794859384936304ec635160e01b85526004850163ffffffff604092959493606083019683521660208201520152565b03915afa801561035557611c6e925f91611c7e575b506001600160c01b031692611c68846146dc565b90612450565b9061076f60405192839283611b5c565b611c97915060203d602011610f8f57610f8181836103ad565b5f611c54565b611cb191503d805f833e610fb981836103ad565b5f611c04565b3461035a575f36600319011261035a576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b3461035a57606036600319011261035a576004356001600160401b03811161035a57611d2b903690600401611566565b60243590611d38826105e1565b6044356001600160401b03811161035a57611d57903690600401610b82565b60ce5491939092916001600160a01b03163303611e415761032493611e2b93611dab611db293611d90611d88613e87565b963690613ecd565b86524363ffffffff16602087015263ffffffff166060860152565b36916107a8565b60408201526040516020810190611dcd816116398585613f27565b519020611de261181360c95463ffffffff1690565b5560c95463ffffffff16907fa6b1912d066fd3a5413a6222a5785ae43f242458f3eef45c50dd808987b878fd60405180611e2363ffffffff86169482613f27565b0390a2613991565b63ffffffff1663ffffffff1960c954161760c955565b60405162461bcd60e51b815260206004820152602160248201527f5461736b2067656e657261746f72206d757374206265207468652063616c6c656044820152603960f91b6064820152608490fd5b3461035a57602036600319011261035a57600435611ead8161066c565b611eb561463c565b6001600160a01b03811615611ecd5761032490614694565b60405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608490fd5b3461035a575f36600319011261035a57602060405160648152f35b3461035a57602036600319011261035a5760043560405163755b36bd60e11b81526020816004817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa908115610355575f91611fff575b506001600160a01b03163303611ff057611fbe606654198219811614612054565b806066556040519081527f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c60203392a2005b63794821ff60e01b5f5260045ffd5b612018915060203d6020116109eb576109dd81836103ad565b5f611f9d565b9081602091031261035a57516106d881610927565b6040513d5f823e3d90fd5b1561204557565b631d77d47760e21b5f5260045ffd5b1561205b57565b63c61dca5d60e01b5f5260045ffd5b634e487b7160e01b5f52603260045260245ffd5b90600281101561208f5760051b0190565b61206a565b634e487b7160e01b5f52601260045260245ffd5b61218461216161218a9561215b61215485875160208901518a515160208c51015160208d016020815151915101519189519360208b0151956040519760208901998a5260208a015260408901526060880152608087015260a086015260c085015260e084015261010083015261212b81610120840103601f1981018352826103ad565b5190207f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001900690565b8096614043565b90614089565b9261215b6121766121706140eb565b946141e2565b9161217f6142fe565b614043565b91614332565b9091565b6040805190919061219f83826103ad565b6001815291601f1901366020840137565b906121ba8261067d565b6121c760405191826103ad565b82815280926121d8601f199161067d565b0190602036910137565b80511561208f5760200190565b805182101561208f5760209160051b010190565b9081602091031261035a575190565b91909161221f83516121b0565b925f5b81518110156122d25780602061224b61223e61227494866121ef565b516001600160a01b031690565b6040516309aa152760e11b81526001600160a01b03909116600482015292839081906024820190565b03816001600160a01b0388165afa8015610355576001925f916122a4575b5061229d82886121ef565b5201612222565b6122c5915060203d81116122cb575b6122bd81836103ad565b810190612203565b5f612292565b503d6122b3565b505050565b9081602091031261035a57516106d88161066c565b906122f68261067d565b61230360405191826103ad565b82815260208193612316601f199161067d565b0191015f5b82811061232757505050565b60608282015260200161231b565b90815181101561208f570160200190565b60208183031261035a578051906001600160401b03821161035a57019080601f8301121561035a5781516123798161067d565b9261238760405194856103ad565b81845260208085019260051b82010192831161035a57602001905b8282106123af5750505090565b81518152602091820191016123a2565b906123c98261067d565b6123d660405191826103ad565b82815280926123e7601f199161067d565b015f5b8181106123f657505050565b6040519060608201918083106001600160401b0384111761038d576020926040525f81525f838201525f6040820152828286010152016123ea565b9081602091031261035a57516001600160601b038116810361035a5790565b604051636830483560e01b815293919291906001600160a01b0316602085600481845afa948515610355575f95612765575b50604051634f4c91e160e11b815294602086600481855afa918215610355576004965f93612743575b5060209060405197888092632efa2ca360e11b82525afa958615610355575f96612722575b506124de85939295516122ec565b945f935b80518510156127185761250f6125096124fb8784612335565b516001600160f81b03191690565b60f81c90565b604051638902624560e01b815260ff8216600482015263ffffffff88166024820152909490925f846044816001600160a01b0385165afa938415610355575f946126f4575b5061255f84516123bf565b612569888b6121ef565b52612574878a6121ef565b505f5b84518110156126e3578060206125906125b293886121ef565b518d60405180809681946308f6629d60e31b8352600483019190602083019252565b03916001600160a01b03165afa918215610355575f926126c3575b506125d881876121ef565b518a60208a6125e7858b6121ef565b5160405163fa28c62760e01b8152600481019190915260ff91909116602482015263ffffffff929092166044830152816064816001600160a01b038d165afa9384156103555761267a8c8f61267560019861268c9789975f92612693575b506126606126516103ef565b6001600160a01b039098168852565b60208701526001600160601b03166040860152565b6121ef565b519061268683836121ef565b526121ef565b5001612577565b6126b591925060203d81116126bc575b6126ad81836103ad565b810190612431565b905f612645565b503d6126a3565b6126dc91925060203d81116109eb576109dd81836103ad565b905f6125cd565b5060019096019590945091506124e2565b6127119194503d805f833e61270981836103ad565b810190612346565b925f612554565b5050509350505090565b61273c91965060203d6020116109eb576109dd81836103ad565b945f6124d0565b602091935061275e90823d84116109eb576109dd81836103ad565b92906124ab565b61277f91955060203d6020116109eb576109dd81836103ad565b935f612482565b6040519061279382610392565b606080838181528160208201528160408201520152565b60208183031261035a578051906001600160401b03821161035a57019080601f8301121561035a5781516127dd8161067d565b926127eb60405194856103ad565b81845260208085019260051b82010192831161035a57602001905b8282106128135750505090565b602080918351612822816105e1565b815201910190612806565b63ffffffff909116815260406020820181905281018390526001600160fb1b03831161035a5760609260051b809284830137010190565b908060209392818452848401375f828201840152601f01601f1916010190565b60409063ffffffff6106d895931681528160208201520191612864565b634e487b7160e01b5f52601160045260245ffd5b60ff1660ff81146128c65760010190565b6128a1565b919081101561208f5760051b0190565b9081602091031261035a57516001600160c01b038116810361035a5790565b1561290157565b6325ec6c1f60e01b5f5260045ffd5b9082101561208f570190565b9081602091031261035a57516106d8816105e1565b5f1981146128c65760010190565b9161295d60209263ffffffff92969596604086526040860191612864565b9416910152565b9593949592909192612974612786565b50604051636830483560e01b8152936001600160a01b03919091169190602085600481865afa948515610355575f95612dab575b506129b1612786565b946040516361c8a12f60e11b81525f81806129d18d8d8b6004850161282d565b0381885afa908115610355575f91612d91575b5086526040516340e03a8160e11b81526001600160a01b039190911692905f8180612a1485878b60048501612884565b0381875afa908115610355575f91612d77575b506040870152612a36816122ec565b9860608701998a525f5b60ff811683811015612cc257885f612a69838f612a5c886121b0565b90519061268683836121ef565b505f8a868f5b818410612aec575050505090508c612a86826121b0565b915f5b818110612ab357505091612aa891612aae9493519061268683836121ef565b506128b5565b612a40565b80612ae6612ad1610f06600194612acb8a89516121ef565b516121ef565b612adb83886121ef565b9063ffffffff169052565b01612a89565b610f0684612b018160209695612b09956128cb565b3597516121ef565b6040516304ec635160e01b8152600481019690965263ffffffff9182166024870152166044850152836064818d5afa801561035557888f888a918f94612bae6001612ba181938d809d5f92612c96575b50612509612b7d612b8b92612b76878060c01b03861615156128fa565b8b8d612910565b356001600160f81b03191690565b6001600160c01b0391821660ff919091161c1690565b166001600160c01b031690565b14612bca575b5050505050600191925001908a918a868f612a6f565b8597612bec93612be5602097999861250995612b7d956128cb565b3595612910565b60405163dd9846b960e01b8152600481019290925260ff16602482015263ffffffff939093166044840152826064818c5afa908115610355578f612c4a90612c4f9383886001975f93612c5e575b50612acb90612adb9394516121ef565b612931565b905082918a888f888a91612bb4565b612adb935090612c87612acb9260203d8111612c8f575b612c7f81836103ad565b81019061291c565b935090612c3a565b503d612c75565b612b8b919250612b7d612cb96125099260203d8111610f8f57610f8181836103ad565b93925050612b59565b505050929095975060049496506020915060405194858092632efa2ca360e11b82525afa90811561035557612d18945f948593612d56575b5060405163354952a360e21b8152958694859384936004850161293f565b03916001600160a01b03165afa908115610355575f91612d3c575b50602082015290565b612d5091503d805f833e610fb981836103ad565b5f612d33565b612d7091935060203d6020116109eb576109dd81836103ad565b915f612cfa565b612d8b91503d805f833e610fb981836103ad565b5f612a27565b612da591503d805f833e610fb981836103ad565b5f6129e4565b612dc591955060203d6020116109eb576109dd81836103ad565b935f6129a8565b60409063ffffffff6106d894931681528160208201520190610694565b60405190612df682610372565b60606020838281520152565b15612e0957565b62f8202d60e51b5f5260045ffd5b15612e1e57565b6343714afd60e01b5f5260045ffd5b15612e3457565b635f832f4160e01b5f5260045ffd5b15612e4a57565b634b874f4560e01b5f5260045ffd5b9081602091031261035a57516106d881610dd8565b5f198101919082116128c657565b15612e8357565b633fdc650560e21b5f5260045ffd5b90600182018092116128c657565b90600282018092116128c657565b90600382018092116128c657565b90600482018092116128c657565b90600582018092116128c657565b919082018092116128c657565b15612eec57565b63affc5edb60e01b5f5260045ffd5b9081602091031261035a575167ffffffffffffffff198116810361035a5790565b15612f2357565b63e1310aed60e01b5f5260045ffd5b906001600160601b03809116911603906001600160601b0382116128c657565b15612f5957565b6367988d3360e01b5f5260045ffd5b15612f6f57565b63ab1b236b60e01b5f5260045ffd5b60049163ffffffff60e01b9060e01b1681520160208251919201905f5b818110612fa85750505090565b8251845260209384019390920191600101612f9b565b949392909193612fcc612de9565b50612fd8851515612e02565b60408401515185148061384d575b8061383f575b80613831575b612ffb90612e17565b61300d60208501515185515114612e2d565b61302463ffffffff431663ffffffff841610612e43565b61302c6103e0565b5f81525f60208201529261303e612de9565b613047876121b0565b6020820152613055876121b0565b815261305f612de9565b9261306e6020880151516121b0565b845261307e6020880151516121b0565b602085810191909152604051639aa1653d60e01b815290816004817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa8015610355576130e7915f91613802575b506130e2368b876107a8565b61449e565b985f965b602089015180518910156132595760208861314e610f068c6131468f96868e61312b6131188680956121ef565b5180515f526020015160205260405f2090565b61313884848401516121ef565b5282613226575b01516121ef565b5195516121ef565b6040516304ec635160e01b8152600481019490945263ffffffff9182166024850152166044830152816064816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa9182156103555761215b8a6131fb8f6131f48f8460208f926131eb936131e38460019e6132019e5f91613209575b508f8060c01b031692516121ef565b5201516121ef565b51938d516121ef565b51166144c9565b906144fa565b9701966130eb565b6132209150863d8111610f8f57610f8181836103ad565b5f6131d4565b61325461323684848401516121ef565b5161324d8484015161324787612e6e565b906121ef565b5110612e7c565b61313f565b5090959794965061326e9198939299506145b7565b9161327b60975460ff1690565b9081156137fa576040516318891fd760e31b81526020816004817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa908115610355575f916137db575b5091905b5f925b81841061332c5750505050509261331361330e613307613326958561163998608060606020990151920151926120a8565b9190612f52565b612f68565b0151604051928391602083019586612f7e565b51902090565b92989596909399919794878b888c888d6136d5575b610f068260a0613381612509612b7d846133899761337b61336d6131188f9c604060209f9e01516121ef565b67ffffffffffffffff191690565b9b612910565b9701516121ef565b604051631a2f32ab60e21b815260ff95909516600486015263ffffffff9182166024860152166044840152826064816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa9081156103555761344d610f068f958f906134458f978f96848f61343f60c096613438848f60209f9061313f612b7d996040936125099c5f916136a7575b5067ffffffffffffffff19918216911614612f1c565b5190614089565b9c612910565b9601516121ef565b604051636414a62b60e11b815260ff94909416600485015263ffffffff9182166024850152166044830152816064816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa908115610355576134da918c8f925f92613683575b5060206134cc929301516121ef565b906001600160601b03169052565b6134fa8c6134cc8c6134f361195c8260208601516121ef565b92516121ef565b5f985f5b60208a01515181101561366a578b8d61353c8961352f612509612b7d868f8961352791516121ef565b519487612910565b60ff161c60019081161490565b61354b575b50506001016134fe565b8a8a6135cd859f948f9686612acb8f9360e0613584610f0695602061357c612509612b7d839f61358d9c8991612910565b9a01516121ef565b519b01516121ef565b60405163795f4a5760e11b815260ff909316600484015263ffffffff93841660248401526044830196909652919094166064850152839081906084820190565b03817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa908115610355578f613639908f936001959486955f92613644575b506136336134cc9293519361362e61195c84876121ef565b612f32565b926121ef565b019a90508b8d613541565b6134cc92506136636136339160203d81116126bc576126ad81836103ad565b9250613616565b5093919796996001919699509a94929a019291906132d6565b6134cc92506136a0602091823d81116126bc576126ad81836103ad565b92506134bd565b60206136c892503d81116136ce575b6136c081836103ad565b810190612efb565b5f613422565b503d6136b6565b61371294506136ef925061250991612b7d91602095612910565b60405163124d062160e11b815260ff909116600482015291829081906024820190565b03817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa8015610355576020896133898f938f60a08f97612509612b7d8f8f9061337b61336d6131188f60408b96918f8893610f069f6137969061379c936133819f5f926137b2575b5063ffffffff809116931690612ed8565b11612ee5565b5050505050509750505050505092935050613341565b602063ffffffff92935082916137d3913d81116122cb576122bd81836103ad565b929150613785565b6137f4915060203d602011612c8f57612c7f81836103ad565b5f6132cf565b5f91906132d3565b613824915060203d60201161382a575b61381c81836103ad565b810190612e59565b5f6130d6565b503d613812565b5060e0840151518514612ff2565b5060c0840151518514612fec565b5060a0840151518514612fe6565b356106d8816105e1565b6020809163ffffffff8135613879816105e1565b1684520135910152565b9160406103de92949361389a816080810197613865565b0190613865565b156138a857565b60405162461bcd60e51b815260206004820152603d60248201527f5461736b20726573706f6e736520646f6573206e6f74206d617463682074686560448201527f206f6e65207265636f7264656420696e2074686520636f6e74726163740000006064820152608490fd5b1561391a57565b60405162461bcd60e51b815260206004820152604360248201527f54686520726573706f6e736520746f2074686973207461736b2068617320616c60448201527f7265616479206265656e206368616c6c656e676564207375636365737366756c606482015262363c9760e91b608482015260a490fd5b63ffffffff60019116019063ffffffff82116128c657565b63ffffffff60649116019063ffffffff82116128c657565b9063ffffffff8091169116019063ffffffff82116128c657565b156139e257565b60405162461bcd60e51b815260206004820152603760248201527f546865206368616c6c656e676520706572696f6420666f72207468697320746160448201527f736b2068617320616c726561647920657870697265642e0000000000000000006064820152608490fd5b903590601e198136030182121561035a57018035906001600160401b03821161035a5760200191813603831361035a57565b805180835260209291819084018484015e5f828201840152601f01601f1916010190565b9035601e198236030181121561035a5701602081359101916001600160401b03821161035a57813603831361035a57565b60208152813590603e198336030182121561035a576080613b816060613b7a613b40876106d89701856020880152613b2e613b23613b128380613aa3565b604060a08c015260e08b0191612864565b916020810190613aa3565b888303609f190160c08a015290612864565b613b5c613b4f60208a016105ef565b63ffffffff166040880152565b613b696040890189613aa3565b878303601f19018589015290612864565b95016105ef565b63ffffffff16910152565b15613b9357565b60405162461bcd60e51b815260206004820152603d60248201527f737570706c696564207461736b20646f6573206e6f74206d617463682074686560448201527f206f6e65207265636f7264656420696e2074686520636f6e74726163740000006064820152608490fd5b15613c0557565b60405162461bcd60e51b815260206004820152602c60248201527f41676772656761746f722068617320616c726561647920726573706f6e64656460448201526b20746f20746865207461736b60a01b6064820152608490fd5b15613c6657565b60405162461bcd60e51b815260206004820152602d60248201527f41676772656761746f722068617320726573706f6e64656420746f207468652060448201526c7461736b20746f6f206c61746560981b6064820152608490fd5b6040810192916103de9190613865565b906064820291808304606414901517156128c657565b906006820291808304600614901517156128c657565b906001600160601b03809116911602906001600160601b0382169182036128c657565b15613d2757565b608460405162461bcd60e51b815260206004820152604060248201527f5369676e61746f7269657320646f206e6f74206f776e206174206c656173742060448201527f7468726573686f6c642070657263656e74616765206f6620612071756f72756d6064820152fd5b9092916020606091613da7846080810197613865565b63ffffffff81511660408501520151910152565b15613dc257565b60405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608490fd5b613e2790614694565b60cd80546001600160a01b03199081166001600160a01b039384161790915560ce805482169383169390931790925560d0805483169382169390931790925560cf805482169383169390931790925560d180549092169216919091179055565b60405190613e9482610392565b5f606083613ea0612de9565b81528260208201528160408201520152565b9080601f8301121561035a578160206106d8933591016107a8565b919060408382031261035a5760405190613ee682610372565b819380356001600160401b03811161035a5782613f04918301613eb2565b83526020810135916001600160401b03831161035a576020926104d89201613eb2565b60208152608063ffffffff6060613f8f613f6e86518560208801526020613f5a8251604060a08b015260e08a0190613a7f565b910151878203609f190160c0890152613a7f565b8360208801511660408701526040870151601f198783030184880152613a7f565b9401511691015290565b5f196066556040515f1981527fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d60203392a2565b806066556040519081527fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d60203392a2565b6040519061400c82610372565b5f6020838281520152565b6040519061018061402881846103ad565b368337565b6040519061403c6020836103ad565b6020368337565b91906040906060614052613fff565b948592602085519261406485856103ad565b8436853780518452015160208301528482015260076107cf195a01fa1561408757565bfe5b60209291608060409261409a613fff565b958693818651936140ab86866103ad565b85368637805185520151828401528051868401520151606082015260066107cf195a01fa801561408757156140dc57565b63d4b68fd760e01b5f5260045ffd5b6040516140f781610372565b604090815161410683826103ad565b823682378152602082519161411b84846103ad565b833684370152805161412d82826103ad565b7f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c281527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed602082015281519061418383836103ad565b7f275dc4a288d1afb3cbb1ac09187524c7db36395df7be3b99e673b13a075a65ec82527f1d9befcd05a5323e6da4d435f3b617cdb3af83285c2df711ef39c01571827f9d60208301526141d8835193846103ad565b8252602082015290565b5f51602061481f5f395f51905f52906141f9613fff565b505f919006602060c0835b6142f9575f935f51602061481f5f395f51905f526003818681818009090860405161422f85826103ad565b8436823784818560405161424382826103ad565b813682378381528360208201528360408201528560608201527f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f5260808201525f51602061481f5f395f51905f5260a082015260056107cf195a01fa8015614087576142ad90614808565b51916142f9575f51602061481f5f395f51905f52828009146142e457505f51602061481f5f395f51905f5260015f94089293614204565b929350506142f06103e0565b92835282015290565b612094565b614306613fff565b5060405161431381610372565b600181526002602082015290565b90600c81101561208f5760051b0190565b9392909161434060406103fe565b948552602085015261435260406103fe565b9182526020820152614362614017565b925f5b6002811061438f5750505060206101809261437e61402d565b93849160086201d4c0fa9151151590565b8061439b600192613ce7565b6143a5828561207e565b51516143b18289614321565b5260206143be838661207e565b5101516143d36143cd83612e92565b89614321565b526143de828661207e565b5151516143ed6143cd83612ea0565b526144036143fb838761207e565b515160200190565b516144106143cd83612eae565b52602061441d838761207e565b5101515161442d6143cd83612ebc565b5261445961445361444c6020614443868a61207e565b51015160200190565b5192612eca565b88614321565b5201614365565b60207f40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc91151560ff196097541660ff821617609755604051908152a1565b9060016144ac60ff93614790565b928392161b11156144ba5790565b63ca95733360e01b5f5260045ffd5b805f915b6144d5575090565b5f1981018181116128c65761ffff9116911661ffff81146128c65760010190806144cd565b90614503613fff565b5061ffff8116906102008210156145a857600182146145a3576145246103e0565b5f81525f602082015292906001905f925b61ffff831685101561454957505050505090565b600161ffff831660ff86161c811614614583575b600161457961456e8360ff94614089565b9460011b61fffe1690565b9401169291614535565b94600161457961456e6145988960ff95614089565b98935050505061455d565b505090565b637fc4ea7d60e11b5f5260045ffd5b6145bf613fff565b50805190811580614630575b156145ec5750506040516145e06040826103ad565b5f81525f602082015290565b60205f51602061481f5f395f51905f52910151065f51602061481f5f395f51905f52035f51602061481f5f395f51905f5281116128c657604051916141d883610372565b506020810151156145cb565b6033546001600160a01b0316330361465057565b606460405162461bcd60e51b815260206004820152602060248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152fd5b603380546001600160a01b039283166001600160a01b0319821681179092559091167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e05f80a3565b61ffff6146e8826144c9565b166146f28161078d565b9061470060405192836103ad565b80825261470f601f199161078d565b013660208301375f5f5b825182108061476f575b15614768576001811b8416614741575b61473c90612931565b614719565b90600161473c9160ff60f81b8460f81b165f1a61475e8287612335565b5301919050614733565b5050905090565b506101008110614723565b1561478157565b631019106960e31b5f5260045ffd5b906101008251116147f9578151156147f457602082015160019060f81c81901b5b83518210156147ef576001906147da6147d06125096124fb8689612335565b60ff600191161b90565b906147e681831161477a565b179101906147b1565b925050565b5f9150565b637da54e4760e11b5f5260045ffd5b1561480f57565b63d51edae360e01b5f5260045ffdfe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47a2646970667358221220674d184bd18b5735580bbb6fd61037109fbbc254b45a6b4ab542568477980e8864736f6c634300081b0033",
}

// ContractAwesomeVaultTaskManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractAwesomeVaultTaskManagerMetaData.ABI instead.
var ContractAwesomeVaultTaskManagerABI = ContractAwesomeVaultTaskManagerMetaData.ABI

// ContractAwesomeVaultTaskManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractAwesomeVaultTaskManagerMetaData.Bin instead.
var ContractAwesomeVaultTaskManagerBin = ContractAwesomeVaultTaskManagerMetaData.Bin

// DeployContractAwesomeVaultTaskManager deploys a new Ethereum contract, binding an instance of ContractAwesomeVaultTaskManager to it.
func DeployContractAwesomeVaultTaskManager(auth *bind.TransactOpts, backend bind.ContractBackend, _registryCoordinator common.Address, _pauserRegistry common.Address, _taskResponseWindowBlock uint32) (common.Address, *types.Transaction, *ContractAwesomeVaultTaskManager, error) {
	parsed, err := ContractAwesomeVaultTaskManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractAwesomeVaultTaskManagerBin), backend, _registryCoordinator, _pauserRegistry, _taskResponseWindowBlock)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractAwesomeVaultTaskManager{ContractAwesomeVaultTaskManagerCaller: ContractAwesomeVaultTaskManagerCaller{contract: contract}, ContractAwesomeVaultTaskManagerTransactor: ContractAwesomeVaultTaskManagerTransactor{contract: contract}, ContractAwesomeVaultTaskManagerFilterer: ContractAwesomeVaultTaskManagerFilterer{contract: contract}}, nil
}

// ContractAwesomeVaultTaskManager is an auto generated Go binding around an Ethereum contract.
type ContractAwesomeVaultTaskManager struct {
	ContractAwesomeVaultTaskManagerCaller     // Read-only binding to the contract
	ContractAwesomeVaultTaskManagerTransactor // Write-only binding to the contract
	ContractAwesomeVaultTaskManagerFilterer   // Log filterer for contract events
}

// ContractAwesomeVaultTaskManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractAwesomeVaultTaskManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractAwesomeVaultTaskManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractAwesomeVaultTaskManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractAwesomeVaultTaskManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractAwesomeVaultTaskManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractAwesomeVaultTaskManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractAwesomeVaultTaskManagerSession struct {
	Contract     *ContractAwesomeVaultTaskManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                    // Call options to use throughout this session
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// ContractAwesomeVaultTaskManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractAwesomeVaultTaskManagerCallerSession struct {
	Contract *ContractAwesomeVaultTaskManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                          // Call options to use throughout this session
}

// ContractAwesomeVaultTaskManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractAwesomeVaultTaskManagerTransactorSession struct {
	Contract     *ContractAwesomeVaultTaskManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                          // Transaction auth options to use throughout this session
}

// ContractAwesomeVaultTaskManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractAwesomeVaultTaskManagerRaw struct {
	Contract *ContractAwesomeVaultTaskManager // Generic contract binding to access the raw methods on
}

// ContractAwesomeVaultTaskManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractAwesomeVaultTaskManagerCallerRaw struct {
	Contract *ContractAwesomeVaultTaskManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ContractAwesomeVaultTaskManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractAwesomeVaultTaskManagerTransactorRaw struct {
	Contract *ContractAwesomeVaultTaskManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractAwesomeVaultTaskManager creates a new instance of ContractAwesomeVaultTaskManager, bound to a specific deployed contract.
func NewContractAwesomeVaultTaskManager(address common.Address, backend bind.ContractBackend) (*ContractAwesomeVaultTaskManager, error) {
	contract, err := bindContractAwesomeVaultTaskManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractAwesomeVaultTaskManager{ContractAwesomeVaultTaskManagerCaller: ContractAwesomeVaultTaskManagerCaller{contract: contract}, ContractAwesomeVaultTaskManagerTransactor: ContractAwesomeVaultTaskManagerTransactor{contract: contract}, ContractAwesomeVaultTaskManagerFilterer: ContractAwesomeVaultTaskManagerFilterer{contract: contract}}, nil
}

// NewContractAwesomeVaultTaskManagerCaller creates a new read-only instance of ContractAwesomeVaultTaskManager, bound to a specific deployed contract.
func NewContractAwesomeVaultTaskManagerCaller(address common.Address, caller bind.ContractCaller) (*ContractAwesomeVaultTaskManagerCaller, error) {
	contract, err := bindContractAwesomeVaultTaskManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractAwesomeVaultTaskManagerCaller{contract: contract}, nil
}

// NewContractAwesomeVaultTaskManagerTransactor creates a new write-only instance of ContractAwesomeVaultTaskManager, bound to a specific deployed contract.
func NewContractAwesomeVaultTaskManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractAwesomeVaultTaskManagerTransactor, error) {
	contract, err := bindContractAwesomeVaultTaskManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractAwesomeVaultTaskManagerTransactor{contract: contract}, nil
}

// NewContractAwesomeVaultTaskManagerFilterer creates a new log filterer instance of ContractAwesomeVaultTaskManager, bound to a specific deployed contract.
func NewContractAwesomeVaultTaskManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractAwesomeVaultTaskManagerFilterer, error) {
	contract, err := bindContractAwesomeVaultTaskManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractAwesomeVaultTaskManagerFilterer{contract: contract}, nil
}

// bindContractAwesomeVaultTaskManager binds a generic wrapper to an already deployed contract.
func bindContractAwesomeVaultTaskManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractAwesomeVaultTaskManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractAwesomeVaultTaskManager.Contract.ContractAwesomeVaultTaskManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractAwesomeVaultTaskManager.Contract.ContractAwesomeVaultTaskManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractAwesomeVaultTaskManager.Contract.ContractAwesomeVaultTaskManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractAwesomeVaultTaskManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractAwesomeVaultTaskManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractAwesomeVaultTaskManager.Contract.contract.Transact(opts, method, params...)
}

// TASKCHALLENGEWINDOWBLOCK is a free data retrieval call binding the contract method 0xf63c5bab.
//
// Solidity: function TASK_CHALLENGE_WINDOW_BLOCK() view returns(uint32)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerCaller) TASKCHALLENGEWINDOWBLOCK(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractAwesomeVaultTaskManager.contract.Call(opts, &out, "TASK_CHALLENGE_WINDOW_BLOCK")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TASKCHALLENGEWINDOWBLOCK is a free data retrieval call binding the contract method 0xf63c5bab.
//
// Solidity: function TASK_CHALLENGE_WINDOW_BLOCK() view returns(uint32)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerSession) TASKCHALLENGEWINDOWBLOCK() (uint32, error) {
	return _ContractAwesomeVaultTaskManager.Contract.TASKCHALLENGEWINDOWBLOCK(&_ContractAwesomeVaultTaskManager.CallOpts)
}

// TASKCHALLENGEWINDOWBLOCK is a free data retrieval call binding the contract method 0xf63c5bab.
//
// Solidity: function TASK_CHALLENGE_WINDOW_BLOCK() view returns(uint32)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerCallerSession) TASKCHALLENGEWINDOWBLOCK() (uint32, error) {
	return _ContractAwesomeVaultTaskManager.Contract.TASKCHALLENGEWINDOWBLOCK(&_ContractAwesomeVaultTaskManager.CallOpts)
}

// TASKRESPONSEWINDOWBLOCK is a free data retrieval call binding the contract method 0x1ad43189.
//
// Solidity: function TASK_RESPONSE_WINDOW_BLOCK() view returns(uint32)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerCaller) TASKRESPONSEWINDOWBLOCK(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractAwesomeVaultTaskManager.contract.Call(opts, &out, "TASK_RESPONSE_WINDOW_BLOCK")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TASKRESPONSEWINDOWBLOCK is a free data retrieval call binding the contract method 0x1ad43189.
//
// Solidity: function TASK_RESPONSE_WINDOW_BLOCK() view returns(uint32)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerSession) TASKRESPONSEWINDOWBLOCK() (uint32, error) {
	return _ContractAwesomeVaultTaskManager.Contract.TASKRESPONSEWINDOWBLOCK(&_ContractAwesomeVaultTaskManager.CallOpts)
}

// TASKRESPONSEWINDOWBLOCK is a free data retrieval call binding the contract method 0x1ad43189.
//
// Solidity: function TASK_RESPONSE_WINDOW_BLOCK() view returns(uint32)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerCallerSession) TASKRESPONSEWINDOWBLOCK() (uint32, error) {
	return _ContractAwesomeVaultTaskManager.Contract.TASKRESPONSEWINDOWBLOCK(&_ContractAwesomeVaultTaskManager.CallOpts)
}

// WADSTOSLASH is a free data retrieval call binding the contract method 0x5a2d7f02.
//
// Solidity: function WADS_TO_SLASH() view returns(uint256)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerCaller) WADSTOSLASH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractAwesomeVaultTaskManager.contract.Call(opts, &out, "WADS_TO_SLASH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WADSTOSLASH is a free data retrieval call binding the contract method 0x5a2d7f02.
//
// Solidity: function WADS_TO_SLASH() view returns(uint256)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerSession) WADSTOSLASH() (*big.Int, error) {
	return _ContractAwesomeVaultTaskManager.Contract.WADSTOSLASH(&_ContractAwesomeVaultTaskManager.CallOpts)
}

// WADSTOSLASH is a free data retrieval call binding the contract method 0x5a2d7f02.
//
// Solidity: function WADS_TO_SLASH() view returns(uint256)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerCallerSession) WADSTOSLASH() (*big.Int, error) {
	return _ContractAwesomeVaultTaskManager.Contract.WADSTOSLASH(&_ContractAwesomeVaultTaskManager.CallOpts)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerCaller) Aggregator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractAwesomeVaultTaskManager.contract.Call(opts, &out, "aggregator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerSession) Aggregator() (common.Address, error) {
	return _ContractAwesomeVaultTaskManager.Contract.Aggregator(&_ContractAwesomeVaultTaskManager.CallOpts)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerCallerSession) Aggregator() (common.Address, error) {
	return _ContractAwesomeVaultTaskManager.Contract.Aggregator(&_ContractAwesomeVaultTaskManager.CallOpts)
}

// AllTaskHashes is a free data retrieval call binding the contract method 0x2d89f6fc.
//
// Solidity: function allTaskHashes(uint32 ) view returns(bytes32)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerCaller) AllTaskHashes(opts *bind.CallOpts, arg0 uint32) ([32]byte, error) {
	var out []interface{}
	err := _ContractAwesomeVaultTaskManager.contract.Call(opts, &out, "allTaskHashes", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AllTaskHashes is a free data retrieval call binding the contract method 0x2d89f6fc.
//
// Solidity: function allTaskHashes(uint32 ) view returns(bytes32)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerSession) AllTaskHashes(arg0 uint32) ([32]byte, error) {
	return _ContractAwesomeVaultTaskManager.Contract.AllTaskHashes(&_ContractAwesomeVaultTaskManager.CallOpts, arg0)
}

// AllTaskHashes is a free data retrieval call binding the contract method 0x2d89f6fc.
//
// Solidity: function allTaskHashes(uint32 ) view returns(bytes32)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerCallerSession) AllTaskHashes(arg0 uint32) ([32]byte, error) {
	return _ContractAwesomeVaultTaskManager.Contract.AllTaskHashes(&_ContractAwesomeVaultTaskManager.CallOpts, arg0)
}

// AllTaskResponses is a free data retrieval call binding the contract method 0x2cb223d5.
//
// Solidity: function allTaskResponses(uint32 ) view returns(bytes32)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerCaller) AllTaskResponses(opts *bind.CallOpts, arg0 uint32) ([32]byte, error) {
	var out []interface{}
	err := _ContractAwesomeVaultTaskManager.contract.Call(opts, &out, "allTaskResponses", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AllTaskResponses is a free data retrieval call binding the contract method 0x2cb223d5.
//
// Solidity: function allTaskResponses(uint32 ) view returns(bytes32)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerSession) AllTaskResponses(arg0 uint32) ([32]byte, error) {
	return _ContractAwesomeVaultTaskManager.Contract.AllTaskResponses(&_ContractAwesomeVaultTaskManager.CallOpts, arg0)
}

// AllTaskResponses is a free data retrieval call binding the contract method 0x2cb223d5.
//
// Solidity: function allTaskResponses(uint32 ) view returns(bytes32)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerCallerSession) AllTaskResponses(arg0 uint32) ([32]byte, error) {
	return _ContractAwesomeVaultTaskManager.Contract.AllTaskResponses(&_ContractAwesomeVaultTaskManager.CallOpts, arg0)
}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerCaller) AllocationManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractAwesomeVaultTaskManager.contract.Call(opts, &out, "allocationManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerSession) AllocationManager() (common.Address, error) {
	return _ContractAwesomeVaultTaskManager.Contract.AllocationManager(&_ContractAwesomeVaultTaskManager.CallOpts)
}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerCallerSession) AllocationManager() (common.Address, error) {
	return _ContractAwesomeVaultTaskManager.Contract.AllocationManager(&_ContractAwesomeVaultTaskManager.CallOpts)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerCaller) BlsApkRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractAwesomeVaultTaskManager.contract.Call(opts, &out, "blsApkRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerSession) BlsApkRegistry() (common.Address, error) {
	return _ContractAwesomeVaultTaskManager.Contract.BlsApkRegistry(&_ContractAwesomeVaultTaskManager.CallOpts)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerCallerSession) BlsApkRegistry() (common.Address, error) {
	return _ContractAwesomeVaultTaskManager.Contract.BlsApkRegistry(&_ContractAwesomeVaultTaskManager.CallOpts)
}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerCaller) CheckSignatures(opts *bind.CallOpts, msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerTypesNonSignerStakesAndSignature) (IBLSSignatureCheckerTypesQuorumStakeTotals, [32]byte, error) {
	var out []interface{}
	err := _ContractAwesomeVaultTaskManager.contract.Call(opts, &out, "checkSignatures", msgHash, quorumNumbers, referenceBlockNumber, params)

	if err != nil {
		return *new(IBLSSignatureCheckerTypesQuorumStakeTotals), *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(IBLSSignatureCheckerTypesQuorumStakeTotals)).(*IBLSSignatureCheckerTypesQuorumStakeTotals)
	out1 := *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return out0, out1, err

}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerSession) CheckSignatures(msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerTypesNonSignerStakesAndSignature) (IBLSSignatureCheckerTypesQuorumStakeTotals, [32]byte, error) {
	return _ContractAwesomeVaultTaskManager.Contract.CheckSignatures(&_ContractAwesomeVaultTaskManager.CallOpts, msgHash, quorumNumbers, referenceBlockNumber, params)
}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerCallerSession) CheckSignatures(msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerTypesNonSignerStakesAndSignature) (IBLSSignatureCheckerTypesQuorumStakeTotals, [32]byte, error) {
	return _ContractAwesomeVaultTaskManager.Contract.CheckSignatures(&_ContractAwesomeVaultTaskManager.CallOpts, msgHash, quorumNumbers, referenceBlockNumber, params)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerCaller) Delegation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractAwesomeVaultTaskManager.contract.Call(opts, &out, "delegation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerSession) Delegation() (common.Address, error) {
	return _ContractAwesomeVaultTaskManager.Contract.Delegation(&_ContractAwesomeVaultTaskManager.CallOpts)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerCallerSession) Delegation() (common.Address, error) {
	return _ContractAwesomeVaultTaskManager.Contract.Delegation(&_ContractAwesomeVaultTaskManager.CallOpts)
}

// Generator is a free data retrieval call binding the contract method 0x7afa1eed.
//
// Solidity: function generator() view returns(address)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerCaller) Generator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractAwesomeVaultTaskManager.contract.Call(opts, &out, "generator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Generator is a free data retrieval call binding the contract method 0x7afa1eed.
//
// Solidity: function generator() view returns(address)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerSession) Generator() (common.Address, error) {
	return _ContractAwesomeVaultTaskManager.Contract.Generator(&_ContractAwesomeVaultTaskManager.CallOpts)
}

// Generator is a free data retrieval call binding the contract method 0x7afa1eed.
//
// Solidity: function generator() view returns(address)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerCallerSession) Generator() (common.Address, error) {
	return _ContractAwesomeVaultTaskManager.Contract.Generator(&_ContractAwesomeVaultTaskManager.CallOpts)
}

// GetBatchOperatorFromId is a free data retrieval call binding the contract method 0x4d2b57fe.
//
// Solidity: function getBatchOperatorFromId(address registryCoordinator, bytes32[] operatorIds) view returns(address[] operators)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerCaller) GetBatchOperatorFromId(opts *bind.CallOpts, registryCoordinator common.Address, operatorIds [][32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _ContractAwesomeVaultTaskManager.contract.Call(opts, &out, "getBatchOperatorFromId", registryCoordinator, operatorIds)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetBatchOperatorFromId is a free data retrieval call binding the contract method 0x4d2b57fe.
//
// Solidity: function getBatchOperatorFromId(address registryCoordinator, bytes32[] operatorIds) view returns(address[] operators)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerSession) GetBatchOperatorFromId(registryCoordinator common.Address, operatorIds [][32]byte) ([]common.Address, error) {
	return _ContractAwesomeVaultTaskManager.Contract.GetBatchOperatorFromId(&_ContractAwesomeVaultTaskManager.CallOpts, registryCoordinator, operatorIds)
}

// GetBatchOperatorFromId is a free data retrieval call binding the contract method 0x4d2b57fe.
//
// Solidity: function getBatchOperatorFromId(address registryCoordinator, bytes32[] operatorIds) view returns(address[] operators)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerCallerSession) GetBatchOperatorFromId(registryCoordinator common.Address, operatorIds [][32]byte) ([]common.Address, error) {
	return _ContractAwesomeVaultTaskManager.Contract.GetBatchOperatorFromId(&_ContractAwesomeVaultTaskManager.CallOpts, registryCoordinator, operatorIds)
}

// GetBatchOperatorId is a free data retrieval call binding the contract method 0x31b36bd9.
//
// Solidity: function getBatchOperatorId(address registryCoordinator, address[] operators) view returns(bytes32[] operatorIds)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerCaller) GetBatchOperatorId(opts *bind.CallOpts, registryCoordinator common.Address, operators []common.Address) ([][32]byte, error) {
	var out []interface{}
	err := _ContractAwesomeVaultTaskManager.contract.Call(opts, &out, "getBatchOperatorId", registryCoordinator, operators)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetBatchOperatorId is a free data retrieval call binding the contract method 0x31b36bd9.
//
// Solidity: function getBatchOperatorId(address registryCoordinator, address[] operators) view returns(bytes32[] operatorIds)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerSession) GetBatchOperatorId(registryCoordinator common.Address, operators []common.Address) ([][32]byte, error) {
	return _ContractAwesomeVaultTaskManager.Contract.GetBatchOperatorId(&_ContractAwesomeVaultTaskManager.CallOpts, registryCoordinator, operators)
}

// GetBatchOperatorId is a free data retrieval call binding the contract method 0x31b36bd9.
//
// Solidity: function getBatchOperatorId(address registryCoordinator, address[] operators) view returns(bytes32[] operatorIds)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerCallerSession) GetBatchOperatorId(registryCoordinator common.Address, operators []common.Address) ([][32]byte, error) {
	return _ContractAwesomeVaultTaskManager.Contract.GetBatchOperatorId(&_ContractAwesomeVaultTaskManager.CallOpts, registryCoordinator, operators)
}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerCaller) GetCheckSignaturesIndices(opts *bind.CallOpts, registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	var out []interface{}
	err := _ContractAwesomeVaultTaskManager.contract.Call(opts, &out, "getCheckSignaturesIndices", registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)

	if err != nil {
		return *new(OperatorStateRetrieverCheckSignaturesIndices), err
	}

	out0 := *abi.ConvertType(out[0], new(OperatorStateRetrieverCheckSignaturesIndices)).(*OperatorStateRetrieverCheckSignaturesIndices)

	return out0, err

}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerSession) GetCheckSignaturesIndices(registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	return _ContractAwesomeVaultTaskManager.Contract.GetCheckSignaturesIndices(&_ContractAwesomeVaultTaskManager.CallOpts, registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)
}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerCallerSession) GetCheckSignaturesIndices(registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	return _ContractAwesomeVaultTaskManager.Contract.GetCheckSignaturesIndices(&_ContractAwesomeVaultTaskManager.CallOpts, registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)
}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerCaller) GetOperatorState(opts *bind.CallOpts, registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	var out []interface{}
	err := _ContractAwesomeVaultTaskManager.contract.Call(opts, &out, "getOperatorState", registryCoordinator, quorumNumbers, blockNumber)

	if err != nil {
		return *new([][]OperatorStateRetrieverOperator), err
	}

	out0 := *abi.ConvertType(out[0], new([][]OperatorStateRetrieverOperator)).(*[][]OperatorStateRetrieverOperator)

	return out0, err

}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerSession) GetOperatorState(registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	return _ContractAwesomeVaultTaskManager.Contract.GetOperatorState(&_ContractAwesomeVaultTaskManager.CallOpts, registryCoordinator, quorumNumbers, blockNumber)
}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerCallerSession) GetOperatorState(registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	return _ContractAwesomeVaultTaskManager.Contract.GetOperatorState(&_ContractAwesomeVaultTaskManager.CallOpts, registryCoordinator, quorumNumbers, blockNumber)
}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (address,bytes32,uint96)[][])
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerCaller) GetOperatorState0(opts *bind.CallOpts, registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	var out []interface{}
	err := _ContractAwesomeVaultTaskManager.contract.Call(opts, &out, "getOperatorState0", registryCoordinator, operatorId, blockNumber)

	if err != nil {
		return *new(*big.Int), *new([][]OperatorStateRetrieverOperator), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new([][]OperatorStateRetrieverOperator)).(*[][]OperatorStateRetrieverOperator)

	return out0, out1, err

}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (address,bytes32,uint96)[][])
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerSession) GetOperatorState0(registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	return _ContractAwesomeVaultTaskManager.Contract.GetOperatorState0(&_ContractAwesomeVaultTaskManager.CallOpts, registryCoordinator, operatorId, blockNumber)
}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (address,bytes32,uint96)[][])
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerCallerSession) GetOperatorState0(registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	return _ContractAwesomeVaultTaskManager.Contract.GetOperatorState0(&_ContractAwesomeVaultTaskManager.CallOpts, registryCoordinator, operatorId, blockNumber)
}

// GetQuorumBitmapsAtBlockNumber is a free data retrieval call binding the contract method 0x5c155662.
//
// Solidity: function getQuorumBitmapsAtBlockNumber(address registryCoordinator, bytes32[] operatorIds, uint32 blockNumber) view returns(uint256[])
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerCaller) GetQuorumBitmapsAtBlockNumber(opts *bind.CallOpts, registryCoordinator common.Address, operatorIds [][32]byte, blockNumber uint32) ([]*big.Int, error) {
	var out []interface{}
	err := _ContractAwesomeVaultTaskManager.contract.Call(opts, &out, "getQuorumBitmapsAtBlockNumber", registryCoordinator, operatorIds, blockNumber)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetQuorumBitmapsAtBlockNumber is a free data retrieval call binding the contract method 0x5c155662.
//
// Solidity: function getQuorumBitmapsAtBlockNumber(address registryCoordinator, bytes32[] operatorIds, uint32 blockNumber) view returns(uint256[])
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerSession) GetQuorumBitmapsAtBlockNumber(registryCoordinator common.Address, operatorIds [][32]byte, blockNumber uint32) ([]*big.Int, error) {
	return _ContractAwesomeVaultTaskManager.Contract.GetQuorumBitmapsAtBlockNumber(&_ContractAwesomeVaultTaskManager.CallOpts, registryCoordinator, operatorIds, blockNumber)
}

// GetQuorumBitmapsAtBlockNumber is a free data retrieval call binding the contract method 0x5c155662.
//
// Solidity: function getQuorumBitmapsAtBlockNumber(address registryCoordinator, bytes32[] operatorIds, uint32 blockNumber) view returns(uint256[])
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerCallerSession) GetQuorumBitmapsAtBlockNumber(registryCoordinator common.Address, operatorIds [][32]byte, blockNumber uint32) ([]*big.Int, error) {
	return _ContractAwesomeVaultTaskManager.Contract.GetQuorumBitmapsAtBlockNumber(&_ContractAwesomeVaultTaskManager.CallOpts, registryCoordinator, operatorIds, blockNumber)
}

// GetTaskResponseWindowBlock is a free data retrieval call binding the contract method 0xf5c9899d.
//
// Solidity: function getTaskResponseWindowBlock() view returns(uint32)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerCaller) GetTaskResponseWindowBlock(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractAwesomeVaultTaskManager.contract.Call(opts, &out, "getTaskResponseWindowBlock")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetTaskResponseWindowBlock is a free data retrieval call binding the contract method 0xf5c9899d.
//
// Solidity: function getTaskResponseWindowBlock() view returns(uint32)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerSession) GetTaskResponseWindowBlock() (uint32, error) {
	return _ContractAwesomeVaultTaskManager.Contract.GetTaskResponseWindowBlock(&_ContractAwesomeVaultTaskManager.CallOpts)
}

// GetTaskResponseWindowBlock is a free data retrieval call binding the contract method 0xf5c9899d.
//
// Solidity: function getTaskResponseWindowBlock() view returns(uint32)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerCallerSession) GetTaskResponseWindowBlock() (uint32, error) {
	return _ContractAwesomeVaultTaskManager.Contract.GetTaskResponseWindowBlock(&_ContractAwesomeVaultTaskManager.CallOpts)
}

// InstantSlasher is a free data retrieval call binding the contract method 0x9b290e98.
//
// Solidity: function instantSlasher() view returns(address)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerCaller) InstantSlasher(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractAwesomeVaultTaskManager.contract.Call(opts, &out, "instantSlasher")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// InstantSlasher is a free data retrieval call binding the contract method 0x9b290e98.
//
// Solidity: function instantSlasher() view returns(address)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerSession) InstantSlasher() (common.Address, error) {
	return _ContractAwesomeVaultTaskManager.Contract.InstantSlasher(&_ContractAwesomeVaultTaskManager.CallOpts)
}

// InstantSlasher is a free data retrieval call binding the contract method 0x9b290e98.
//
// Solidity: function instantSlasher() view returns(address)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerCallerSession) InstantSlasher() (common.Address, error) {
	return _ContractAwesomeVaultTaskManager.Contract.InstantSlasher(&_ContractAwesomeVaultTaskManager.CallOpts)
}

// LatestTaskNum is a free data retrieval call binding the contract method 0x8b00ce7c.
//
// Solidity: function latestTaskNum() view returns(uint32)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerCaller) LatestTaskNum(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractAwesomeVaultTaskManager.contract.Call(opts, &out, "latestTaskNum")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LatestTaskNum is a free data retrieval call binding the contract method 0x8b00ce7c.
//
// Solidity: function latestTaskNum() view returns(uint32)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerSession) LatestTaskNum() (uint32, error) {
	return _ContractAwesomeVaultTaskManager.Contract.LatestTaskNum(&_ContractAwesomeVaultTaskManager.CallOpts)
}

// LatestTaskNum is a free data retrieval call binding the contract method 0x8b00ce7c.
//
// Solidity: function latestTaskNum() view returns(uint32)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerCallerSession) LatestTaskNum() (uint32, error) {
	return _ContractAwesomeVaultTaskManager.Contract.LatestTaskNum(&_ContractAwesomeVaultTaskManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractAwesomeVaultTaskManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerSession) Owner() (common.Address, error) {
	return _ContractAwesomeVaultTaskManager.Contract.Owner(&_ContractAwesomeVaultTaskManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerCallerSession) Owner() (common.Address, error) {
	return _ContractAwesomeVaultTaskManager.Contract.Owner(&_ContractAwesomeVaultTaskManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _ContractAwesomeVaultTaskManager.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerSession) Paused(index uint8) (bool, error) {
	return _ContractAwesomeVaultTaskManager.Contract.Paused(&_ContractAwesomeVaultTaskManager.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerCallerSession) Paused(index uint8) (bool, error) {
	return _ContractAwesomeVaultTaskManager.Contract.Paused(&_ContractAwesomeVaultTaskManager.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractAwesomeVaultTaskManager.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerSession) Paused0() (*big.Int, error) {
	return _ContractAwesomeVaultTaskManager.Contract.Paused0(&_ContractAwesomeVaultTaskManager.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerCallerSession) Paused0() (*big.Int, error) {
	return _ContractAwesomeVaultTaskManager.Contract.Paused0(&_ContractAwesomeVaultTaskManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractAwesomeVaultTaskManager.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerSession) PauserRegistry() (common.Address, error) {
	return _ContractAwesomeVaultTaskManager.Contract.PauserRegistry(&_ContractAwesomeVaultTaskManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerCallerSession) PauserRegistry() (common.Address, error) {
	return _ContractAwesomeVaultTaskManager.Contract.PauserRegistry(&_ContractAwesomeVaultTaskManager.CallOpts)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerCaller) RegistryCoordinator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractAwesomeVaultTaskManager.contract.Call(opts, &out, "registryCoordinator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerSession) RegistryCoordinator() (common.Address, error) {
	return _ContractAwesomeVaultTaskManager.Contract.RegistryCoordinator(&_ContractAwesomeVaultTaskManager.CallOpts)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerCallerSession) RegistryCoordinator() (common.Address, error) {
	return _ContractAwesomeVaultTaskManager.Contract.RegistryCoordinator(&_ContractAwesomeVaultTaskManager.CallOpts)
}

// ServiceManager is a free data retrieval call binding the contract method 0x3998fdd3.
//
// Solidity: function serviceManager() view returns(address)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerCaller) ServiceManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractAwesomeVaultTaskManager.contract.Call(opts, &out, "serviceManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ServiceManager is a free data retrieval call binding the contract method 0x3998fdd3.
//
// Solidity: function serviceManager() view returns(address)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerSession) ServiceManager() (common.Address, error) {
	return _ContractAwesomeVaultTaskManager.Contract.ServiceManager(&_ContractAwesomeVaultTaskManager.CallOpts)
}

// ServiceManager is a free data retrieval call binding the contract method 0x3998fdd3.
//
// Solidity: function serviceManager() view returns(address)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerCallerSession) ServiceManager() (common.Address, error) {
	return _ContractAwesomeVaultTaskManager.Contract.ServiceManager(&_ContractAwesomeVaultTaskManager.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerCaller) StakeRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractAwesomeVaultTaskManager.contract.Call(opts, &out, "stakeRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerSession) StakeRegistry() (common.Address, error) {
	return _ContractAwesomeVaultTaskManager.Contract.StakeRegistry(&_ContractAwesomeVaultTaskManager.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerCallerSession) StakeRegistry() (common.Address, error) {
	return _ContractAwesomeVaultTaskManager.Contract.StakeRegistry(&_ContractAwesomeVaultTaskManager.CallOpts)
}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerCaller) StaleStakesForbidden(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ContractAwesomeVaultTaskManager.contract.Call(opts, &out, "staleStakesForbidden")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerSession) StaleStakesForbidden() (bool, error) {
	return _ContractAwesomeVaultTaskManager.Contract.StaleStakesForbidden(&_ContractAwesomeVaultTaskManager.CallOpts)
}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerCallerSession) StaleStakesForbidden() (bool, error) {
	return _ContractAwesomeVaultTaskManager.Contract.StaleStakesForbidden(&_ContractAwesomeVaultTaskManager.CallOpts)
}

// TaskNumber is a free data retrieval call binding the contract method 0x72d18e8d.
//
// Solidity: function taskNumber() view returns(uint32)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerCaller) TaskNumber(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractAwesomeVaultTaskManager.contract.Call(opts, &out, "taskNumber")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TaskNumber is a free data retrieval call binding the contract method 0x72d18e8d.
//
// Solidity: function taskNumber() view returns(uint32)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerSession) TaskNumber() (uint32, error) {
	return _ContractAwesomeVaultTaskManager.Contract.TaskNumber(&_ContractAwesomeVaultTaskManager.CallOpts)
}

// TaskNumber is a free data retrieval call binding the contract method 0x72d18e8d.
//
// Solidity: function taskNumber() view returns(uint32)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerCallerSession) TaskNumber() (uint32, error) {
	return _ContractAwesomeVaultTaskManager.Contract.TaskNumber(&_ContractAwesomeVaultTaskManager.CallOpts)
}

// TaskSuccesfullyChallenged is a free data retrieval call binding the contract method 0x5decc3f5.
//
// Solidity: function taskSuccesfullyChallenged(uint32 ) view returns(bool)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerCaller) TaskSuccesfullyChallenged(opts *bind.CallOpts, arg0 uint32) (bool, error) {
	var out []interface{}
	err := _ContractAwesomeVaultTaskManager.contract.Call(opts, &out, "taskSuccesfullyChallenged", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TaskSuccesfullyChallenged is a free data retrieval call binding the contract method 0x5decc3f5.
//
// Solidity: function taskSuccesfullyChallenged(uint32 ) view returns(bool)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerSession) TaskSuccesfullyChallenged(arg0 uint32) (bool, error) {
	return _ContractAwesomeVaultTaskManager.Contract.TaskSuccesfullyChallenged(&_ContractAwesomeVaultTaskManager.CallOpts, arg0)
}

// TaskSuccesfullyChallenged is a free data retrieval call binding the contract method 0x5decc3f5.
//
// Solidity: function taskSuccesfullyChallenged(uint32 ) view returns(bool)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerCallerSession) TaskSuccesfullyChallenged(arg0 uint32) (bool, error) {
	return _ContractAwesomeVaultTaskManager.Contract.TaskSuccesfullyChallenged(&_ContractAwesomeVaultTaskManager.CallOpts, arg0)
}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerCaller) TrySignatureAndApkVerification(opts *bind.CallOpts, msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	var out []interface{}
	err := _ContractAwesomeVaultTaskManager.contract.Call(opts, &out, "trySignatureAndApkVerification", msgHash, apk, apkG2, sigma)

	outstruct := new(struct {
		PairingSuccessful bool
		SiganatureIsValid bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PairingSuccessful = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.SiganatureIsValid = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerSession) TrySignatureAndApkVerification(msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	return _ContractAwesomeVaultTaskManager.Contract.TrySignatureAndApkVerification(&_ContractAwesomeVaultTaskManager.CallOpts, msgHash, apk, apkG2, sigma)
}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerCallerSession) TrySignatureAndApkVerification(msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	return _ContractAwesomeVaultTaskManager.Contract.TrySignatureAndApkVerification(&_ContractAwesomeVaultTaskManager.CallOpts, msgHash, apk, apkG2, sigma)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0xef029dbc.
//
// Solidity: function createNewTask((string,string) input, uint32 quorumThresholdPercentage, bytes quorumNumbers) returns()
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerTransactor) CreateNewTask(opts *bind.TransactOpts, input IAwesomeVaultTaskManagerTaskInput, quorumThresholdPercentage uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractAwesomeVaultTaskManager.contract.Transact(opts, "createNewTask", input, quorumThresholdPercentage, quorumNumbers)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0xef029dbc.
//
// Solidity: function createNewTask((string,string) input, uint32 quorumThresholdPercentage, bytes quorumNumbers) returns()
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerSession) CreateNewTask(input IAwesomeVaultTaskManagerTaskInput, quorumThresholdPercentage uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractAwesomeVaultTaskManager.Contract.CreateNewTask(&_ContractAwesomeVaultTaskManager.TransactOpts, input, quorumThresholdPercentage, quorumNumbers)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0xef029dbc.
//
// Solidity: function createNewTask((string,string) input, uint32 quorumThresholdPercentage, bytes quorumNumbers) returns()
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerTransactorSession) CreateNewTask(input IAwesomeVaultTaskManagerTaskInput, quorumThresholdPercentage uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractAwesomeVaultTaskManager.Contract.CreateNewTask(&_ContractAwesomeVaultTaskManager.TransactOpts, input, quorumThresholdPercentage, quorumNumbers)
}

// Initialize is a paid mutator transaction binding the contract method 0xcc2a9a5b.
//
// Solidity: function initialize(address initialOwner, address _aggregator, address _generator, address _allocationManager, address _slasher, address _serviceManager) returns()
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address, _aggregator common.Address, _generator common.Address, _allocationManager common.Address, _slasher common.Address, _serviceManager common.Address) (*types.Transaction, error) {
	return _ContractAwesomeVaultTaskManager.contract.Transact(opts, "initialize", initialOwner, _aggregator, _generator, _allocationManager, _slasher, _serviceManager)
}

// Initialize is a paid mutator transaction binding the contract method 0xcc2a9a5b.
//
// Solidity: function initialize(address initialOwner, address _aggregator, address _generator, address _allocationManager, address _slasher, address _serviceManager) returns()
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerSession) Initialize(initialOwner common.Address, _aggregator common.Address, _generator common.Address, _allocationManager common.Address, _slasher common.Address, _serviceManager common.Address) (*types.Transaction, error) {
	return _ContractAwesomeVaultTaskManager.Contract.Initialize(&_ContractAwesomeVaultTaskManager.TransactOpts, initialOwner, _aggregator, _generator, _allocationManager, _slasher, _serviceManager)
}

// Initialize is a paid mutator transaction binding the contract method 0xcc2a9a5b.
//
// Solidity: function initialize(address initialOwner, address _aggregator, address _generator, address _allocationManager, address _slasher, address _serviceManager) returns()
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerTransactorSession) Initialize(initialOwner common.Address, _aggregator common.Address, _generator common.Address, _allocationManager common.Address, _slasher common.Address, _serviceManager common.Address) (*types.Transaction, error) {
	return _ContractAwesomeVaultTaskManager.Contract.Initialize(&_ContractAwesomeVaultTaskManager.TransactOpts, initialOwner, _aggregator, _generator, _allocationManager, _slasher, _serviceManager)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractAwesomeVaultTaskManager.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractAwesomeVaultTaskManager.Contract.Pause(&_ContractAwesomeVaultTaskManager.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractAwesomeVaultTaskManager.Contract.Pause(&_ContractAwesomeVaultTaskManager.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractAwesomeVaultTaskManager.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerSession) PauseAll() (*types.Transaction, error) {
	return _ContractAwesomeVaultTaskManager.Contract.PauseAll(&_ContractAwesomeVaultTaskManager.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerTransactorSession) PauseAll() (*types.Transaction, error) {
	return _ContractAwesomeVaultTaskManager.Contract.PauseAll(&_ContractAwesomeVaultTaskManager.TransactOpts)
}

// RaiseAndResolveChallenge is a paid mutator transaction binding the contract method 0x92a51824.
//
// Solidity: function raiseAndResolveChallenge(((string,string),uint32,bytes,uint32) task, (uint32,bytes32) taskResponse, (uint32,bytes32) taskResponseMetadata, (uint256,uint256)[] pubkeysOfNonSigningOperators) returns()
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerTransactor) RaiseAndResolveChallenge(opts *bind.TransactOpts, task IAwesomeVaultTaskManagerTask, taskResponse IAwesomeVaultTaskManagerTaskResponse, taskResponseMetadata IAwesomeVaultTaskManagerTaskResponseMetadata, pubkeysOfNonSigningOperators []BN254G1Point) (*types.Transaction, error) {
	return _ContractAwesomeVaultTaskManager.contract.Transact(opts, "raiseAndResolveChallenge", task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators)
}

// RaiseAndResolveChallenge is a paid mutator transaction binding the contract method 0x92a51824.
//
// Solidity: function raiseAndResolveChallenge(((string,string),uint32,bytes,uint32) task, (uint32,bytes32) taskResponse, (uint32,bytes32) taskResponseMetadata, (uint256,uint256)[] pubkeysOfNonSigningOperators) returns()
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerSession) RaiseAndResolveChallenge(task IAwesomeVaultTaskManagerTask, taskResponse IAwesomeVaultTaskManagerTaskResponse, taskResponseMetadata IAwesomeVaultTaskManagerTaskResponseMetadata, pubkeysOfNonSigningOperators []BN254G1Point) (*types.Transaction, error) {
	return _ContractAwesomeVaultTaskManager.Contract.RaiseAndResolveChallenge(&_ContractAwesomeVaultTaskManager.TransactOpts, task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators)
}

// RaiseAndResolveChallenge is a paid mutator transaction binding the contract method 0x92a51824.
//
// Solidity: function raiseAndResolveChallenge(((string,string),uint32,bytes,uint32) task, (uint32,bytes32) taskResponse, (uint32,bytes32) taskResponseMetadata, (uint256,uint256)[] pubkeysOfNonSigningOperators) returns()
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerTransactorSession) RaiseAndResolveChallenge(task IAwesomeVaultTaskManagerTask, taskResponse IAwesomeVaultTaskManagerTaskResponse, taskResponseMetadata IAwesomeVaultTaskManagerTaskResponseMetadata, pubkeysOfNonSigningOperators []BN254G1Point) (*types.Transaction, error) {
	return _ContractAwesomeVaultTaskManager.Contract.RaiseAndResolveChallenge(&_ContractAwesomeVaultTaskManager.TransactOpts, task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractAwesomeVaultTaskManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractAwesomeVaultTaskManager.Contract.RenounceOwnership(&_ContractAwesomeVaultTaskManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractAwesomeVaultTaskManager.Contract.RenounceOwnership(&_ContractAwesomeVaultTaskManager.TransactOpts)
}

// RespondToTask is a paid mutator transaction binding the contract method 0xb7023949.
//
// Solidity: function respondToTask(((string,string),uint32,bytes,uint32) task, (uint32,bytes32) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerTransactor) RespondToTask(opts *bind.TransactOpts, task IAwesomeVaultTaskManagerTask, taskResponse IAwesomeVaultTaskManagerTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerTypesNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractAwesomeVaultTaskManager.contract.Transact(opts, "respondToTask", task, taskResponse, nonSignerStakesAndSignature)
}

// RespondToTask is a paid mutator transaction binding the contract method 0xb7023949.
//
// Solidity: function respondToTask(((string,string),uint32,bytes,uint32) task, (uint32,bytes32) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerSession) RespondToTask(task IAwesomeVaultTaskManagerTask, taskResponse IAwesomeVaultTaskManagerTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerTypesNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractAwesomeVaultTaskManager.Contract.RespondToTask(&_ContractAwesomeVaultTaskManager.TransactOpts, task, taskResponse, nonSignerStakesAndSignature)
}

// RespondToTask is a paid mutator transaction binding the contract method 0xb7023949.
//
// Solidity: function respondToTask(((string,string),uint32,bytes,uint32) task, (uint32,bytes32) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerTransactorSession) RespondToTask(task IAwesomeVaultTaskManagerTask, taskResponse IAwesomeVaultTaskManagerTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerTypesNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractAwesomeVaultTaskManager.Contract.RespondToTask(&_ContractAwesomeVaultTaskManager.TransactOpts, task, taskResponse, nonSignerStakesAndSignature)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerTransactor) SetStaleStakesForbidden(opts *bind.TransactOpts, value bool) (*types.Transaction, error) {
	return _ContractAwesomeVaultTaskManager.contract.Transact(opts, "setStaleStakesForbidden", value)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerSession) SetStaleStakesForbidden(value bool) (*types.Transaction, error) {
	return _ContractAwesomeVaultTaskManager.Contract.SetStaleStakesForbidden(&_ContractAwesomeVaultTaskManager.TransactOpts, value)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerTransactorSession) SetStaleStakesForbidden(value bool) (*types.Transaction, error) {
	return _ContractAwesomeVaultTaskManager.Contract.SetStaleStakesForbidden(&_ContractAwesomeVaultTaskManager.TransactOpts, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractAwesomeVaultTaskManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractAwesomeVaultTaskManager.Contract.TransferOwnership(&_ContractAwesomeVaultTaskManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractAwesomeVaultTaskManager.Contract.TransferOwnership(&_ContractAwesomeVaultTaskManager.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractAwesomeVaultTaskManager.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractAwesomeVaultTaskManager.Contract.Unpause(&_ContractAwesomeVaultTaskManager.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractAwesomeVaultTaskManager.Contract.Unpause(&_ContractAwesomeVaultTaskManager.TransactOpts, newPausedStatus)
}

// ContractAwesomeVaultTaskManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractAwesomeVaultTaskManager contract.
type ContractAwesomeVaultTaskManagerInitializedIterator struct {
	Event *ContractAwesomeVaultTaskManagerInitialized // Event containing the contract specifics and raw log

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
func (it *ContractAwesomeVaultTaskManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAwesomeVaultTaskManagerInitialized)
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
		it.Event = new(ContractAwesomeVaultTaskManagerInitialized)
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
func (it *ContractAwesomeVaultTaskManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAwesomeVaultTaskManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAwesomeVaultTaskManagerInitialized represents a Initialized event raised by the ContractAwesomeVaultTaskManager contract.
type ContractAwesomeVaultTaskManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractAwesomeVaultTaskManagerInitializedIterator, error) {

	logs, sub, err := _ContractAwesomeVaultTaskManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractAwesomeVaultTaskManagerInitializedIterator{contract: _ContractAwesomeVaultTaskManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractAwesomeVaultTaskManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractAwesomeVaultTaskManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAwesomeVaultTaskManagerInitialized)
				if err := _ContractAwesomeVaultTaskManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerFilterer) ParseInitialized(log types.Log) (*ContractAwesomeVaultTaskManagerInitialized, error) {
	event := new(ContractAwesomeVaultTaskManagerInitialized)
	if err := _ContractAwesomeVaultTaskManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAwesomeVaultTaskManagerNewTaskCreatedIterator is returned from FilterNewTaskCreated and is used to iterate over the raw logs and unpacked data for NewTaskCreated events raised by the ContractAwesomeVaultTaskManager contract.
type ContractAwesomeVaultTaskManagerNewTaskCreatedIterator struct {
	Event *ContractAwesomeVaultTaskManagerNewTaskCreated // Event containing the contract specifics and raw log

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
func (it *ContractAwesomeVaultTaskManagerNewTaskCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAwesomeVaultTaskManagerNewTaskCreated)
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
		it.Event = new(ContractAwesomeVaultTaskManagerNewTaskCreated)
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
func (it *ContractAwesomeVaultTaskManagerNewTaskCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAwesomeVaultTaskManagerNewTaskCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAwesomeVaultTaskManagerNewTaskCreated represents a NewTaskCreated event raised by the ContractAwesomeVaultTaskManager contract.
type ContractAwesomeVaultTaskManagerNewTaskCreated struct {
	TaskIndex uint32
	Task      IAwesomeVaultTaskManagerTask
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewTaskCreated is a free log retrieval operation binding the contract event 0xa6b1912d066fd3a5413a6222a5785ae43f242458f3eef45c50dd808987b878fd.
//
// Solidity: event NewTaskCreated(uint32 indexed taskIndex, ((string,string),uint32,bytes,uint32) task)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerFilterer) FilterNewTaskCreated(opts *bind.FilterOpts, taskIndex []uint32) (*ContractAwesomeVaultTaskManagerNewTaskCreatedIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractAwesomeVaultTaskManager.contract.FilterLogs(opts, "NewTaskCreated", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return &ContractAwesomeVaultTaskManagerNewTaskCreatedIterator{contract: _ContractAwesomeVaultTaskManager.contract, event: "NewTaskCreated", logs: logs, sub: sub}, nil
}

// WatchNewTaskCreated is a free log subscription operation binding the contract event 0xa6b1912d066fd3a5413a6222a5785ae43f242458f3eef45c50dd808987b878fd.
//
// Solidity: event NewTaskCreated(uint32 indexed taskIndex, ((string,string),uint32,bytes,uint32) task)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerFilterer) WatchNewTaskCreated(opts *bind.WatchOpts, sink chan<- *ContractAwesomeVaultTaskManagerNewTaskCreated, taskIndex []uint32) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractAwesomeVaultTaskManager.contract.WatchLogs(opts, "NewTaskCreated", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAwesomeVaultTaskManagerNewTaskCreated)
				if err := _ContractAwesomeVaultTaskManager.contract.UnpackLog(event, "NewTaskCreated", log); err != nil {
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

// ParseNewTaskCreated is a log parse operation binding the contract event 0xa6b1912d066fd3a5413a6222a5785ae43f242458f3eef45c50dd808987b878fd.
//
// Solidity: event NewTaskCreated(uint32 indexed taskIndex, ((string,string),uint32,bytes,uint32) task)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerFilterer) ParseNewTaskCreated(log types.Log) (*ContractAwesomeVaultTaskManagerNewTaskCreated, error) {
	event := new(ContractAwesomeVaultTaskManagerNewTaskCreated)
	if err := _ContractAwesomeVaultTaskManager.contract.UnpackLog(event, "NewTaskCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAwesomeVaultTaskManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContractAwesomeVaultTaskManager contract.
type ContractAwesomeVaultTaskManagerOwnershipTransferredIterator struct {
	Event *ContractAwesomeVaultTaskManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractAwesomeVaultTaskManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAwesomeVaultTaskManagerOwnershipTransferred)
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
		it.Event = new(ContractAwesomeVaultTaskManagerOwnershipTransferred)
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
func (it *ContractAwesomeVaultTaskManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAwesomeVaultTaskManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAwesomeVaultTaskManagerOwnershipTransferred represents a OwnershipTransferred event raised by the ContractAwesomeVaultTaskManager contract.
type ContractAwesomeVaultTaskManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractAwesomeVaultTaskManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractAwesomeVaultTaskManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractAwesomeVaultTaskManagerOwnershipTransferredIterator{contract: _ContractAwesomeVaultTaskManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractAwesomeVaultTaskManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractAwesomeVaultTaskManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAwesomeVaultTaskManagerOwnershipTransferred)
				if err := _ContractAwesomeVaultTaskManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerFilterer) ParseOwnershipTransferred(log types.Log) (*ContractAwesomeVaultTaskManagerOwnershipTransferred, error) {
	event := new(ContractAwesomeVaultTaskManagerOwnershipTransferred)
	if err := _ContractAwesomeVaultTaskManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAwesomeVaultTaskManagerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ContractAwesomeVaultTaskManager contract.
type ContractAwesomeVaultTaskManagerPausedIterator struct {
	Event *ContractAwesomeVaultTaskManagerPaused // Event containing the contract specifics and raw log

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
func (it *ContractAwesomeVaultTaskManagerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAwesomeVaultTaskManagerPaused)
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
		it.Event = new(ContractAwesomeVaultTaskManagerPaused)
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
func (it *ContractAwesomeVaultTaskManagerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAwesomeVaultTaskManagerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAwesomeVaultTaskManagerPaused represents a Paused event raised by the ContractAwesomeVaultTaskManager contract.
type ContractAwesomeVaultTaskManagerPaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*ContractAwesomeVaultTaskManagerPausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractAwesomeVaultTaskManager.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractAwesomeVaultTaskManagerPausedIterator{contract: _ContractAwesomeVaultTaskManager.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ContractAwesomeVaultTaskManagerPaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractAwesomeVaultTaskManager.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAwesomeVaultTaskManagerPaused)
				if err := _ContractAwesomeVaultTaskManager.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerFilterer) ParsePaused(log types.Log) (*ContractAwesomeVaultTaskManagerPaused, error) {
	event := new(ContractAwesomeVaultTaskManagerPaused)
	if err := _ContractAwesomeVaultTaskManager.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAwesomeVaultTaskManagerStaleStakesForbiddenUpdateIterator is returned from FilterStaleStakesForbiddenUpdate and is used to iterate over the raw logs and unpacked data for StaleStakesForbiddenUpdate events raised by the ContractAwesomeVaultTaskManager contract.
type ContractAwesomeVaultTaskManagerStaleStakesForbiddenUpdateIterator struct {
	Event *ContractAwesomeVaultTaskManagerStaleStakesForbiddenUpdate // Event containing the contract specifics and raw log

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
func (it *ContractAwesomeVaultTaskManagerStaleStakesForbiddenUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAwesomeVaultTaskManagerStaleStakesForbiddenUpdate)
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
		it.Event = new(ContractAwesomeVaultTaskManagerStaleStakesForbiddenUpdate)
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
func (it *ContractAwesomeVaultTaskManagerStaleStakesForbiddenUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAwesomeVaultTaskManagerStaleStakesForbiddenUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAwesomeVaultTaskManagerStaleStakesForbiddenUpdate represents a StaleStakesForbiddenUpdate event raised by the ContractAwesomeVaultTaskManager contract.
type ContractAwesomeVaultTaskManagerStaleStakesForbiddenUpdate struct {
	Value bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterStaleStakesForbiddenUpdate is a free log retrieval operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerFilterer) FilterStaleStakesForbiddenUpdate(opts *bind.FilterOpts) (*ContractAwesomeVaultTaskManagerStaleStakesForbiddenUpdateIterator, error) {

	logs, sub, err := _ContractAwesomeVaultTaskManager.contract.FilterLogs(opts, "StaleStakesForbiddenUpdate")
	if err != nil {
		return nil, err
	}
	return &ContractAwesomeVaultTaskManagerStaleStakesForbiddenUpdateIterator{contract: _ContractAwesomeVaultTaskManager.contract, event: "StaleStakesForbiddenUpdate", logs: logs, sub: sub}, nil
}

// WatchStaleStakesForbiddenUpdate is a free log subscription operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerFilterer) WatchStaleStakesForbiddenUpdate(opts *bind.WatchOpts, sink chan<- *ContractAwesomeVaultTaskManagerStaleStakesForbiddenUpdate) (event.Subscription, error) {

	logs, sub, err := _ContractAwesomeVaultTaskManager.contract.WatchLogs(opts, "StaleStakesForbiddenUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAwesomeVaultTaskManagerStaleStakesForbiddenUpdate)
				if err := _ContractAwesomeVaultTaskManager.contract.UnpackLog(event, "StaleStakesForbiddenUpdate", log); err != nil {
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

// ParseStaleStakesForbiddenUpdate is a log parse operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerFilterer) ParseStaleStakesForbiddenUpdate(log types.Log) (*ContractAwesomeVaultTaskManagerStaleStakesForbiddenUpdate, error) {
	event := new(ContractAwesomeVaultTaskManagerStaleStakesForbiddenUpdate)
	if err := _ContractAwesomeVaultTaskManager.contract.UnpackLog(event, "StaleStakesForbiddenUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAwesomeVaultTaskManagerTaskChallengedSuccessfullyIterator is returned from FilterTaskChallengedSuccessfully and is used to iterate over the raw logs and unpacked data for TaskChallengedSuccessfully events raised by the ContractAwesomeVaultTaskManager contract.
type ContractAwesomeVaultTaskManagerTaskChallengedSuccessfullyIterator struct {
	Event *ContractAwesomeVaultTaskManagerTaskChallengedSuccessfully // Event containing the contract specifics and raw log

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
func (it *ContractAwesomeVaultTaskManagerTaskChallengedSuccessfullyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAwesomeVaultTaskManagerTaskChallengedSuccessfully)
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
		it.Event = new(ContractAwesomeVaultTaskManagerTaskChallengedSuccessfully)
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
func (it *ContractAwesomeVaultTaskManagerTaskChallengedSuccessfullyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAwesomeVaultTaskManagerTaskChallengedSuccessfullyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAwesomeVaultTaskManagerTaskChallengedSuccessfully represents a TaskChallengedSuccessfully event raised by the ContractAwesomeVaultTaskManager contract.
type ContractAwesomeVaultTaskManagerTaskChallengedSuccessfully struct {
	TaskIndex  uint32
	Challenger common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTaskChallengedSuccessfully is a free log retrieval operation binding the contract event 0xc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec.
//
// Solidity: event TaskChallengedSuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerFilterer) FilterTaskChallengedSuccessfully(opts *bind.FilterOpts, taskIndex []uint32, challenger []common.Address) (*ContractAwesomeVaultTaskManagerTaskChallengedSuccessfullyIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractAwesomeVaultTaskManager.contract.FilterLogs(opts, "TaskChallengedSuccessfully", taskIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return &ContractAwesomeVaultTaskManagerTaskChallengedSuccessfullyIterator{contract: _ContractAwesomeVaultTaskManager.contract, event: "TaskChallengedSuccessfully", logs: logs, sub: sub}, nil
}

// WatchTaskChallengedSuccessfully is a free log subscription operation binding the contract event 0xc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec.
//
// Solidity: event TaskChallengedSuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerFilterer) WatchTaskChallengedSuccessfully(opts *bind.WatchOpts, sink chan<- *ContractAwesomeVaultTaskManagerTaskChallengedSuccessfully, taskIndex []uint32, challenger []common.Address) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractAwesomeVaultTaskManager.contract.WatchLogs(opts, "TaskChallengedSuccessfully", taskIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAwesomeVaultTaskManagerTaskChallengedSuccessfully)
				if err := _ContractAwesomeVaultTaskManager.contract.UnpackLog(event, "TaskChallengedSuccessfully", log); err != nil {
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

// ParseTaskChallengedSuccessfully is a log parse operation binding the contract event 0xc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec.
//
// Solidity: event TaskChallengedSuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerFilterer) ParseTaskChallengedSuccessfully(log types.Log) (*ContractAwesomeVaultTaskManagerTaskChallengedSuccessfully, error) {
	event := new(ContractAwesomeVaultTaskManagerTaskChallengedSuccessfully)
	if err := _ContractAwesomeVaultTaskManager.contract.UnpackLog(event, "TaskChallengedSuccessfully", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAwesomeVaultTaskManagerTaskChallengedUnsuccessfullyIterator is returned from FilterTaskChallengedUnsuccessfully and is used to iterate over the raw logs and unpacked data for TaskChallengedUnsuccessfully events raised by the ContractAwesomeVaultTaskManager contract.
type ContractAwesomeVaultTaskManagerTaskChallengedUnsuccessfullyIterator struct {
	Event *ContractAwesomeVaultTaskManagerTaskChallengedUnsuccessfully // Event containing the contract specifics and raw log

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
func (it *ContractAwesomeVaultTaskManagerTaskChallengedUnsuccessfullyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAwesomeVaultTaskManagerTaskChallengedUnsuccessfully)
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
		it.Event = new(ContractAwesomeVaultTaskManagerTaskChallengedUnsuccessfully)
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
func (it *ContractAwesomeVaultTaskManagerTaskChallengedUnsuccessfullyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAwesomeVaultTaskManagerTaskChallengedUnsuccessfullyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAwesomeVaultTaskManagerTaskChallengedUnsuccessfully represents a TaskChallengedUnsuccessfully event raised by the ContractAwesomeVaultTaskManager contract.
type ContractAwesomeVaultTaskManagerTaskChallengedUnsuccessfully struct {
	TaskIndex  uint32
	Challenger common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTaskChallengedUnsuccessfully is a free log retrieval operation binding the contract event 0xfd3e26beeb5967fc5a57a0446914eabc45b4aa474c67a51b4b5160cac60ddb05.
//
// Solidity: event TaskChallengedUnsuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerFilterer) FilterTaskChallengedUnsuccessfully(opts *bind.FilterOpts, taskIndex []uint32, challenger []common.Address) (*ContractAwesomeVaultTaskManagerTaskChallengedUnsuccessfullyIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractAwesomeVaultTaskManager.contract.FilterLogs(opts, "TaskChallengedUnsuccessfully", taskIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return &ContractAwesomeVaultTaskManagerTaskChallengedUnsuccessfullyIterator{contract: _ContractAwesomeVaultTaskManager.contract, event: "TaskChallengedUnsuccessfully", logs: logs, sub: sub}, nil
}

// WatchTaskChallengedUnsuccessfully is a free log subscription operation binding the contract event 0xfd3e26beeb5967fc5a57a0446914eabc45b4aa474c67a51b4b5160cac60ddb05.
//
// Solidity: event TaskChallengedUnsuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerFilterer) WatchTaskChallengedUnsuccessfully(opts *bind.WatchOpts, sink chan<- *ContractAwesomeVaultTaskManagerTaskChallengedUnsuccessfully, taskIndex []uint32, challenger []common.Address) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractAwesomeVaultTaskManager.contract.WatchLogs(opts, "TaskChallengedUnsuccessfully", taskIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAwesomeVaultTaskManagerTaskChallengedUnsuccessfully)
				if err := _ContractAwesomeVaultTaskManager.contract.UnpackLog(event, "TaskChallengedUnsuccessfully", log); err != nil {
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

// ParseTaskChallengedUnsuccessfully is a log parse operation binding the contract event 0xfd3e26beeb5967fc5a57a0446914eabc45b4aa474c67a51b4b5160cac60ddb05.
//
// Solidity: event TaskChallengedUnsuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerFilterer) ParseTaskChallengedUnsuccessfully(log types.Log) (*ContractAwesomeVaultTaskManagerTaskChallengedUnsuccessfully, error) {
	event := new(ContractAwesomeVaultTaskManagerTaskChallengedUnsuccessfully)
	if err := _ContractAwesomeVaultTaskManager.contract.UnpackLog(event, "TaskChallengedUnsuccessfully", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAwesomeVaultTaskManagerTaskCompletedIterator is returned from FilterTaskCompleted and is used to iterate over the raw logs and unpacked data for TaskCompleted events raised by the ContractAwesomeVaultTaskManager contract.
type ContractAwesomeVaultTaskManagerTaskCompletedIterator struct {
	Event *ContractAwesomeVaultTaskManagerTaskCompleted // Event containing the contract specifics and raw log

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
func (it *ContractAwesomeVaultTaskManagerTaskCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAwesomeVaultTaskManagerTaskCompleted)
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
		it.Event = new(ContractAwesomeVaultTaskManagerTaskCompleted)
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
func (it *ContractAwesomeVaultTaskManagerTaskCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAwesomeVaultTaskManagerTaskCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAwesomeVaultTaskManagerTaskCompleted represents a TaskCompleted event raised by the ContractAwesomeVaultTaskManager contract.
type ContractAwesomeVaultTaskManagerTaskCompleted struct {
	TaskIndex uint32
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTaskCompleted is a free log retrieval operation binding the contract event 0x9a144f228a931b9d0d1696fbcdaf310b24b5d2d21e799db623fc986a0f547430.
//
// Solidity: event TaskCompleted(uint32 indexed taskIndex)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerFilterer) FilterTaskCompleted(opts *bind.FilterOpts, taskIndex []uint32) (*ContractAwesomeVaultTaskManagerTaskCompletedIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractAwesomeVaultTaskManager.contract.FilterLogs(opts, "TaskCompleted", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return &ContractAwesomeVaultTaskManagerTaskCompletedIterator{contract: _ContractAwesomeVaultTaskManager.contract, event: "TaskCompleted", logs: logs, sub: sub}, nil
}

// WatchTaskCompleted is a free log subscription operation binding the contract event 0x9a144f228a931b9d0d1696fbcdaf310b24b5d2d21e799db623fc986a0f547430.
//
// Solidity: event TaskCompleted(uint32 indexed taskIndex)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerFilterer) WatchTaskCompleted(opts *bind.WatchOpts, sink chan<- *ContractAwesomeVaultTaskManagerTaskCompleted, taskIndex []uint32) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractAwesomeVaultTaskManager.contract.WatchLogs(opts, "TaskCompleted", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAwesomeVaultTaskManagerTaskCompleted)
				if err := _ContractAwesomeVaultTaskManager.contract.UnpackLog(event, "TaskCompleted", log); err != nil {
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

// ParseTaskCompleted is a log parse operation binding the contract event 0x9a144f228a931b9d0d1696fbcdaf310b24b5d2d21e799db623fc986a0f547430.
//
// Solidity: event TaskCompleted(uint32 indexed taskIndex)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerFilterer) ParseTaskCompleted(log types.Log) (*ContractAwesomeVaultTaskManagerTaskCompleted, error) {
	event := new(ContractAwesomeVaultTaskManagerTaskCompleted)
	if err := _ContractAwesomeVaultTaskManager.contract.UnpackLog(event, "TaskCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAwesomeVaultTaskManagerTaskRespondedIterator is returned from FilterTaskResponded and is used to iterate over the raw logs and unpacked data for TaskResponded events raised by the ContractAwesomeVaultTaskManager contract.
type ContractAwesomeVaultTaskManagerTaskRespondedIterator struct {
	Event *ContractAwesomeVaultTaskManagerTaskResponded // Event containing the contract specifics and raw log

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
func (it *ContractAwesomeVaultTaskManagerTaskRespondedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAwesomeVaultTaskManagerTaskResponded)
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
		it.Event = new(ContractAwesomeVaultTaskManagerTaskResponded)
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
func (it *ContractAwesomeVaultTaskManagerTaskRespondedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAwesomeVaultTaskManagerTaskRespondedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAwesomeVaultTaskManagerTaskResponded represents a TaskResponded event raised by the ContractAwesomeVaultTaskManager contract.
type ContractAwesomeVaultTaskManagerTaskResponded struct {
	TaskResponse         IAwesomeVaultTaskManagerTaskResponse
	TaskResponseMetadata IAwesomeVaultTaskManagerTaskResponseMetadata
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterTaskResponded is a free log retrieval operation binding the contract event 0xf2af11fad73d4349c99cf62f298d337641ea0bb7c0f5a8db92a98a275f734f58.
//
// Solidity: event TaskResponded((uint32,bytes32) taskResponse, (uint32,bytes32) taskResponseMetadata)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerFilterer) FilterTaskResponded(opts *bind.FilterOpts) (*ContractAwesomeVaultTaskManagerTaskRespondedIterator, error) {

	logs, sub, err := _ContractAwesomeVaultTaskManager.contract.FilterLogs(opts, "TaskResponded")
	if err != nil {
		return nil, err
	}
	return &ContractAwesomeVaultTaskManagerTaskRespondedIterator{contract: _ContractAwesomeVaultTaskManager.contract, event: "TaskResponded", logs: logs, sub: sub}, nil
}

// WatchTaskResponded is a free log subscription operation binding the contract event 0xf2af11fad73d4349c99cf62f298d337641ea0bb7c0f5a8db92a98a275f734f58.
//
// Solidity: event TaskResponded((uint32,bytes32) taskResponse, (uint32,bytes32) taskResponseMetadata)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerFilterer) WatchTaskResponded(opts *bind.WatchOpts, sink chan<- *ContractAwesomeVaultTaskManagerTaskResponded) (event.Subscription, error) {

	logs, sub, err := _ContractAwesomeVaultTaskManager.contract.WatchLogs(opts, "TaskResponded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAwesomeVaultTaskManagerTaskResponded)
				if err := _ContractAwesomeVaultTaskManager.contract.UnpackLog(event, "TaskResponded", log); err != nil {
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

// ParseTaskResponded is a log parse operation binding the contract event 0xf2af11fad73d4349c99cf62f298d337641ea0bb7c0f5a8db92a98a275f734f58.
//
// Solidity: event TaskResponded((uint32,bytes32) taskResponse, (uint32,bytes32) taskResponseMetadata)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerFilterer) ParseTaskResponded(log types.Log) (*ContractAwesomeVaultTaskManagerTaskResponded, error) {
	event := new(ContractAwesomeVaultTaskManagerTaskResponded)
	if err := _ContractAwesomeVaultTaskManager.contract.UnpackLog(event, "TaskResponded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAwesomeVaultTaskManagerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ContractAwesomeVaultTaskManager contract.
type ContractAwesomeVaultTaskManagerUnpausedIterator struct {
	Event *ContractAwesomeVaultTaskManagerUnpaused // Event containing the contract specifics and raw log

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
func (it *ContractAwesomeVaultTaskManagerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAwesomeVaultTaskManagerUnpaused)
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
		it.Event = new(ContractAwesomeVaultTaskManagerUnpaused)
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
func (it *ContractAwesomeVaultTaskManagerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAwesomeVaultTaskManagerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAwesomeVaultTaskManagerUnpaused represents a Unpaused event raised by the ContractAwesomeVaultTaskManager contract.
type ContractAwesomeVaultTaskManagerUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*ContractAwesomeVaultTaskManagerUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractAwesomeVaultTaskManager.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractAwesomeVaultTaskManagerUnpausedIterator{contract: _ContractAwesomeVaultTaskManager.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ContractAwesomeVaultTaskManagerUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractAwesomeVaultTaskManager.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAwesomeVaultTaskManagerUnpaused)
				if err := _ContractAwesomeVaultTaskManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerFilterer) ParseUnpaused(log types.Log) (*ContractAwesomeVaultTaskManagerUnpaused, error) {
	event := new(ContractAwesomeVaultTaskManagerUnpaused)
	if err := _ContractAwesomeVaultTaskManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
