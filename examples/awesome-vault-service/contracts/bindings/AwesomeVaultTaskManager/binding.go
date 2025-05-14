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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractISlashingRegistryCoordinator\"},{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"_taskResponseWindowBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"TASK_CHALLENGE_WINDOW_BLOCK\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TASK_RESPONSE_WINDOW_BLOCK\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"WADS_TO_SLASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"aggregator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allStateRoots\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allTaskHashes\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allTaskResponses\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allocationManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"blsApkRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBLSApkRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkSignatures\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureCheckerTypes.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureCheckerTypes.QuorumStakeTotals\",\"components\":[{\"name\":\"signedStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"},{\"name\":\"totalStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"}]},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createNewTask\",\"inputs\":[{\"name\":\"input\",\"type\":\"tuple\",\"internalType\":\"structIAwesomeVaultTaskManager.TaskInput\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"generator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBatchOperatorFromId\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractISlashingRegistryCoordinator\"},{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBatchOperatorId\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractISlashingRegistryCoordinator\"},{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCheckSignaturesIndices\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractISlashingRegistryCoordinator\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"nonSignerOperatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOperatorStateRetriever.CheckSignaturesIndices\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractISlashingRegistryCoordinator\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractISlashingRegistryCoordinator\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumBitmapsAtBlockNumber\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractISlashingRegistryCoordinator\"},{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTaskResponseWindowBlock\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_aggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_generator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_allocationManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_slasher\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_serviceManager\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"instantSlasher\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"latestTaskNum\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"raiseAndResolveChallenge\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structIAwesomeVaultTaskManager.Task\",\"components\":[{\"name\":\"input\",\"type\":\"tuple\",\"internalType\":\"structIAwesomeVaultTaskManager.TaskInput\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"internalType\":\"structIAwesomeVaultTaskManager.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"result\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"taskResponseMetadata\",\"type\":\"tuple\",\"internalType\":\"structIAwesomeVaultTaskManager.TaskResponseMetadata\",\"components\":[{\"name\":\"taskRespondedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"pubkeysOfNonSigningOperators\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"prevStateLeaves\",\"type\":\"tuple[]\",\"internalType\":\"structIAwesomeVaultTaskManager.TaskInput[]\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registryCoordinator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractISlashingRegistryCoordinator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"respondToTask\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structIAwesomeVaultTaskManager.Task\",\"components\":[{\"name\":\"input\",\"type\":\"tuple\",\"internalType\":\"structIAwesomeVaultTaskManager.TaskInput\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"internalType\":\"structIAwesomeVaultTaskManager.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"result\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"nonSignerStakesAndSignature\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureCheckerTypes.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"serviceManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stakeRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"taskNumber\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"taskSuccessfullyChallenged\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"trySignatureAndApkVerification\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"apk\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"pairingSuccessful\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"siganatureIsValid\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewTaskCreated\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"task\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIAwesomeVaultTaskManager.Task\",\"components\":[{\"name\":\"input\",\"type\":\"tuple\",\"internalType\":\"structIAwesomeVaultTaskManager.TaskInput\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskChallengedSuccessfully\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"challenger\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskChallengedUnsuccessfully\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"challenger\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskCompleted\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskResponded\",\"inputs\":[{\"name\":\"taskResponse\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIAwesomeVaultTaskManager.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"result\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"taskResponseMetadata\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIAwesomeVaultTaskManager.TaskResponseMetadata\",\"components\":[{\"name\":\"taskRespondedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BitmapValueTooLarge\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BytesArrayLengthTooLong\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BytesArrayNotOrdered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CurrentlyPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ECAddFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ECMulFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpModFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputAddressZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputArrayLengthMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputEmptyQuorumNumbers\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputNonSignerLengthMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidBLSPairingKey\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidBLSSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidNewPausedStatus\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidQuorumApkHash\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidReferenceBlocknumber\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonSignerPubkeysNotSorted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyPauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyRegistryCoordinatorOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyUnpauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorNotRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ScalarTooLarge\",\"inputs\":[]}]",
	Bin: "0x61014080604052346101d1576060816153da8038038091610020828561027c565b8339810103126101d15780516001600160a01b038116908181036101d15760208301516001600160a01b038116938482036101d157604001519363ffffffff851685036101d1571561026d5760805260a052604051636830483560e01b8152602081600481855afa9081156101dd575f9161022a575b5060c052604051632efa2ca360e11b815290602090829060049082905afa9081156101dd575f916101e8575b5060e05260c05160405163df5cf72360e01b815290602090829060049082906001600160a01b03165afa9081156101dd575f91610197575b50610100526101205260405161512690816102b482396080518181816102ba0152818161112b015281816118860152612223015260a05181818161145601528181612e3801528181612f210152613bd7015260c051818181611412015281816131af01526132eb015260e0518181816113ce015281816130ed0152613af101526101005181611f8a01526101205181818161059e0152611b1d0152f35b90506020813d6020116101d5575b816101b26020938361027c565b810103126101d157516001600160a01b03811681036101d1575f6100fa565b5f80fd5b3d91506101a5565b6040513d5f823e3d90fd5b90506020813d602011610222575b816102036020938361027c565b810103126101d157516001600160a01b03811681036101d1575f6100c2565b3d91506101f6565b90506020813d602011610265575b816102456020938361027c565b810103126101d157516001600160a01b03811681036101d1576004610096565b3d9150610238565b6339b190bb60e11b5f5260045ffd5b601f909101601f19168101906001600160401b0382119082101761029f57604052565b634e487b7160e01b5f52604160045260245ffdfe60806040526004361015610011575f80fd5b5f3560e01c8063136439dd1461028a578063171f1d5b146102855780631ad43189146101db5780631ce8a3e314610280578063245a7bfc1461027b5780632cb223d5146102765780632d89f6fc1461027157806331b36bd91461026c5780633563b0d1146102675780633998fdd3146102625780634d2b57fe1461025d5780634f739f7414610258578063595c6a67146102535780635a2d7f021461024e5780635ac86ab7146102495780635bcaa381146102445780635c1556621461023f5780635c975abb1461023a5780635df459461461023557806368304835146102305780636d14a9871461022b5780636efb463614610226578063715018a61461022157806372d18e8d146102125780637afa1eed1461021c578063886f1195146102175780638b00ce7c146102125780638da5cb5b1461020d5780639b290e9814610208578063a8ec579b14610203578063b7023949146101fe578063ca8aa7c7146101f9578063cc2a9a5b146101f4578063cefdc1d4146101ef578063df5cf723146101ea578063ef029dbc146101e5578063f2fde38b146101e0578063f5c9899d146101db578063f63c5bab146101d65763fabc1cbc146101d1575f80fd5b6121fa565b6121df565b610582565b61214e565b611fb9565b611f75565b611e31565b611cf1565b611cc9565b6119f2565b611943565b6118dd565b6118b5565b611826565b611871565b611849565b6117cb565b61171e565b611441565b6113fd565b6113b9565b61139c565b611245565b6111d3565b6111a0565b611173565b611100565b610c1d565b610a02565b61093b565b6108c1565b610717565b61066f565b610636565b61060e565b6105d0565b610510565b3461034a57602036600319011261034a5760043560405163237dfb4760e11b8152336004820152906020826024817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa9182156103455761031492610300915f91610316575b506122ff565b61030f60665482811614612315565b6144ef565b005b610338915060203d60201161033e575b610330818361039d565b8101906122dc565b5f6102fa565b503d610326565b6122f4565b5f80fd5b634e487b7160e01b5f52604160045260245ffd5b604081019081106001600160401b0382111761037d57604052565b61034e565b608081019081106001600160401b0382111761037d57604052565b90601f801991011681019081106001600160401b0382111761037d57604052565b604051906103ce6101008361039d565b565b604051906103ce60408361039d565b604051906103ce60608361039d565b604051906103ce60a08361039d565b906103ce604051928361039d565b60409060e319011261034a576040519061042482610362565b60e4358252610104356020830152565b919082604091031261034a5760405161044c81610362565b6020808294803584520135910152565b9080601f8301121561034a576040519161047760408461039d565b82906040810192831161034a57905b8282106104935750505090565b8135815260209182019101610486565b90608060631983011261034a576040516104bc81610362565b60206104d782946104ce81606461045c565b845260a461045c565b910152565b919060808382031261034a5760206104d7604051926104fa84610362565b60408496610508838261045c565b86520161045c565b3461034a5761012036600319011261034a57600435604036602319011261034a57610568604091825161054281610362565b60243581526044356020820152610558366104a3565b906105623661040b565b92612369565b8251911515825215156020820152f35b5f91031261034a57565b3461034a575f36600319011261034a57602060405163ffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b63ffffffff81160361034a57565b3461034a57602036600319011261034a5763ffffffff6004356105f2816105c2565b165f5260cd602052602060ff60405f2054166040519015158152f35b3461034a575f36600319011261034a5760ce546040516001600160a01b039091168152602090f35b3461034a57602036600319011261034a5763ffffffff600435610658816105c2565b165f5260cb602052602060405f2054604051908152f35b3461034a57602036600319011261034a5763ffffffff600435610691816105c2565b165f5260ca602052602060405f2054604051908152f35b6001600160a01b0381160361034a57565b6001600160401b03811161037d5760051b60200190565b90602080835192838152019201905f5b8181106106ed5750505090565b82518452602093840193909201916001016106e0565b9060206107149281815201906106d0565b90565b3461034a57604036600319011261034a57600435610734816106a8565b602435906001600160401b03821161034a573660238301121561034a57816004013591610760836106b9565b9261076e604051948561039d565b8084526024602085019160051b8301019136831161034a57602401905b8282106107af576107ab61079f86866124d3565b60405191829182610703565b0390f35b6020809183356107be816106a8565b81520191019061078b565b6001600160401b03811161037d57601f01601f191660200190565b9291926107f0826107c9565b916107fe604051938461039d565b82948184528183011161034a578281602093845f960137010152565b9080602083519182815201916020808360051b8301019401925f915b83831061084557505050505090565b9091929394601f19828203018352855190602080835192838152019201905f905b8082106108855750505060208060019297019301930191939290610836565b909192602060606001926001600160601b0360408851868060a01b03815116845285810151868501520151166040820152019401920190610866565b3461034a57606036600319011261034a576004356108de816106a8565b6024356001600160401b03811161034a573660238201121561034a576107ab916109156109279236906024816004013591016107e4565b60443591610922836105c2565b612711565b60405191829160208352602083019061081a565b3461034a575f36600319011261034a5760d2546040516001600160a01b039091168152602090f35b9080601f8301121561034a57813561097a816106b9565b92610988604051948561039d565b81845260208085019260051b82010192831161034a57602001905b8282106109b05750505090565b81358152602091820191016109a3565b60206040818301928281528451809452019201905f5b8181106109e35750505090565b82516001600160a01b03168452602093840193909201916001016109d6565b3461034a57604036600319011261034a57600435610a1f816106a8565b6024356001600160401b03811161034a57610a3e903690600401610963565b610a488151612471565b916001600160a01b03165f5b8251811015610af557806020610a6d610a8d93866124b0565b5160405180948192630a5aec1960e21b8352600483019190602083019252565b0381865afa91821561034557600192610ac1915f91610ac7575b50610ab283886124b0565b6001600160a01b039091169052565b01610a54565b610ae8915060203d8111610aee575b610ae0818361039d565b810190612598565b5f610aa7565b503d610ad6565b604051806107ab86826109c0565b9181601f8401121561034a578235916001600160401b03831161034a576020838186019501011161034a57565b90602080835192838152019201905f5b818110610b4d5750505090565b825163ffffffff16845260209384019390920191600101610b40565b90602082526060610bb7610ba2610b8c84516080602088015260a0870190610b30565b6020850151868203601f19016040880152610b30565b6040840151858203601f190184870152610b30565b910151916080601f1982840301910152815180825260208201916020808360051b8301019401925f915b838310610bf057505050505090565b9091929394602080610c0e600193601f198682030187528951610b30565b97019301930191939290610be1565b3461034a57608036600319011261034a57600435610c3a816106a8565b60243590610c47826105c2565b6044356001600160401b03811161034a57610c66903690600401610b03565b906064356001600160401b03811161034a57610c86903690600401610963565b92610c8f612a3b565b50604051636830483560e01b81526001600160a01b039190911691602082600481865afa918215610345575f926110df575b50610cca612a3b565b916040516361c8a12f60e11b81525f8180610ce98a8c60048401612ae2565b0381885afa908115610345575f916110c5575b5083526040516340e03a8160e11b81526001600160a01b039190911691905f8180610d2c89868d60048501612b1f565b0381865afa908115610345575f916110ab575b506040840152610d4e856125ad565b94606084019586525f965b60ff88169482861015610feb575f96610d8987610d768551612471565b8b5190610d8383836124b0565b526124b0565b505f5b8351811015610f8257610d9f81856124b0565b5160208d610dbb610db18588516124b0565b5163ffffffff1690565b6040516304ec635160e01b8152600481019490945263ffffffff918216602485015216604483015281606481865afa908115610345576001610e5481938c8a8c610e61965f94610f42575b50610e3e92610e3892610e2a92610e25898060c01b0388161515612b85565b612b9b565b356001600160f81b03191690565b60f81c90565b6001600160c01b0391821660ff919091161c1690565b166001600160c01b031690565b14610e6f575b600101610d8c565b978b6020610e7d8b876124b0565b51610e8f610e38610e2a8d8b8d612b9b565b60405163dd9846b960e01b8152600481019290925260ff16602482015263ffffffff929092166044830152816064818b5afa91821561034557610efd8c600194838d610f02965f93610f0a575b50610eec90610ef29394516124b0565b516124b0565b9063ffffffff169052565b612bbc565b989050610e67565b610ef2935090610f33610eec9260203d8111610f3b575b610f2b818361039d565b810190612ba7565b935090610edc565b503d610f21565b610e2a91945092610e3892610f70610e3e9560203d8111610f7b575b610f68818361039d565b810190612b66565b959250925092610e06565b503d610f5e565b5096989198959095610f9381612471565b905f5b8a828210610fc357610fbd95949250610fb79391505190610d8383836124b0565b50612b50565b96610d59565b90610fe5610fdb610db183610eec89600197516124b0565b610ef283876124b0565b01610f96565b604051632efa2ca360e11b81528491908b856020836004818e5afa90811561034557611038955f94859361108a575b5060405163354952a360e21b81529687948593849360048501612bca565b03916001600160a01b03165afa8015610345576107ab925f91611068575b50602082015260405191829182610b69565b61108491503d805f833e61107c818361039d565b810190612a5f565b83611056565b6110a491935060203d602011610aee57610ae0818361039d565b918761101a565b6110bf91503d805f833e61107c818361039d565b5f610d3f565b6110d991503d805f833e61107c818361039d565b5f610cfc565b6110f991925060203d602011610aee57610ae0818361039d565b905f610cc1565b3461034a575f36600319011261034a5760405163237dfb4760e11b81523360048201526020816024817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa80156103455761116b915f9161031657506122ff565b6103146144bb565b3461034a575f36600319011261034a57602060405167016345785d8a00008152f35b60ff81160361034a57565b3461034a57602036600319011261034a576020600160ff6004356111c381611195565b161b806066541614604051908152f35b3461034a57602036600319011261034a5763ffffffff6004356111f5816105c2565b165f5260cc602052602060405f2054604051908152f35b60206040818301928281528451809452019201905f5b81811061122f5750505090565b8251845260209384019390920191600101611222565b3461034a57606036600319011261034a57600435611262816106a8565b6024356001600160401b03811161034a57611281903690600401610963565b6044359161128e836105c2565b6040516361c8a12f60e11b8152906001600160a01b03165f82806112b6868860048401612ae2565b0381845afa918215610345575f92611380575b506112d48351612471565b935f5b8451811015611372576112ea81866124b0565b51906020836112fc610db184896124b0565b6040516304ec635160e01b8152600481019590955263ffffffff918216602486015216604484015282606481875afa8015610345576001925f91611354575b50828060c01b031661134d82896124b0565b52016112d7565b61136c915060203d8111610f7b57610f68818361039d565b5f61133b565b604051806107ab888261120c565b6113959192503d805f833e61107c818361039d565b905f6112c9565b3461034a575f36600319011261034a576020606654604051908152f35b3461034a575f36600319011261034a576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b3461034a575f36600319011261034a576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b3461034a575f36600319011261034a576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b9080601f8301121561034a57813561149c816106b9565b926114aa604051948561039d565b81845260208085019260051b82010192831161034a57602001905b8282106114d25750505090565b6020809183356114e1816105c2565b8152019101906114c5565b81601f8201121561034a578035611502816106b9565b92611510604051948561039d565b81845260208085019260061b8401019281841161034a57602001915b83831061153a575050505090565b60206040916115498486610434565b81520192019161152c565b9080601f8301121561034a57813561156b816106b9565b92611579604051948561039d565b81845260208085019260051b8201019183831161034a5760208201905b8382106115a557505050505090565b81356001600160401b03811161034a576020916115c787848094880101611485565b815201910190611596565b9190916101808184031261034a576115e86103be565b9281356001600160401b03811161034a5781611605918401611485565b845260208201356001600160401b03811161034a57816116269184016114ec565b602085015260408201356001600160401b03811161034a578161164a9184016114ec565b604085015261165c81606084016104dc565b606085015261166e8160e08401610434565b60808501526101208201356001600160401b03811161034a5781611693918401611485565b60a08501526101408201356001600160401b03811161034a57816116b8918401611485565b60c08501526101608201356001600160401b03811161034a576116db9201611554565b60e0830152565b90602080835192838152019201905f5b8181106116ff5750505090565b82516001600160601b03168452602093840193909201916001016116f2565b3461034a57608036600319011261034a576004356024356001600160401b03811161034a57611751903690600401610b03565b909160443561175f816105c2565b606435926001600160401b03841161034a576117c19461178661178c9536906004016115d2565b93612d5b565b6040519283926040845260206117ad825160408088015260808701906116e2565b910151848203603f190160608601526116e2565b9060208301520390f35b3461034a575f36600319011261034a576117e3614b66565b603380546001600160a01b031981169091555f906001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a3005b3461034a575f36600319011261034a57602063ffffffff60c95416604051908152f35b3461034a575f36600319011261034a5760cf546040516001600160a01b039091168152602090f35b3461034a575f36600319011261034a576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b3461034a575f36600319011261034a576033546040516001600160a01b039091168152602090f35b3461034a575f36600319011261034a5760d0546040516001600160a01b039091168152602090f35b9081608091031261034a5790565b604090602319011261034a57602490565b604090606319011261034a57606490565b9081604091031261034a5790565b3461034a5760e036600319011261034a576004356001600160401b03811161034a57611973903690600401611905565b61197c36611913565b61198536611924565b9160a4356001600160401b03811161034a576119a59036906004016114ec565b60c435936001600160401b03851161034a573660238601121561034a578460040135936001600160401b03851161034a573660248660051b8801011161034a576024610314960193613974565b3461034a57608036600319011261034a576004356001600160401b03811161034a57611a22903690600401611905565b611a2b36611913565b906064356001600160401b03811161034a57611a4b9036906004016115d2565b60ce549092906001600160a01b03163303611c8457611a6e602083949301613439565b91611b71611a7f60408601866137c7565b929094611adf611a9160608901613439565b97604051611ab581611aa760208201948561401c565b03601f19810183528261039d565b519020611ad8611ac488613439565b63ffffffff165f5260ca60205260405f2090565b5414614098565b611b09611b02611aee87613439565b63ffffffff165f5260cb60205260405f2090565b541561410a565b8363ffffffff431696611b53611b4b611b427f0000000000000000000000000000000000000000000000000000000000000000866135f5565b63ffffffff1690565b89111561416b565b6040516020810190611b6981611aa78b856141cd565b519020612d5b565b919060ff5f9616955b828110611c22577ff2af11fad73d4349c99cf62f298d337641ea0bb7c0f5a8db92a98a275f734f58868686611bbc611bb06103d0565b63ffffffff9094168452565b60208301526040516020810190611bd881611aa78686866142b3565b519020611be7611aee83613439565b556020810135611c0d611bf983613439565b63ffffffff165f5260cc60205260405f2090565b55611c1d604051928392836142b3565b0390a1005b80611c7e611c5a611c55611c49611c3c60019688516124b0565b516001600160601b031690565b6001600160601b031690565b6141dd565b611c77611c498b611c72611c3c8760208b01516124b0565b61421f565b1115614242565b01611b7a565b60405162461bcd60e51b815260206004820152601d60248201527f41676772656761746f72206d757374206265207468652063616c6c65720000006044820152606490fd5b3461034a575f36600319011261034a5760d1546040516001600160a01b039091168152602090f35b3461034a5760c036600319011261034a57600435611d0e816106a8565b611d8e602435611d1d816106a8565b604435611d29816106a8565b606435611d35816106a8565b60843591611d42836106a8565b60a43593611d4f856106a8565b5f5496611d7460ff60088a901c16158099819a611e0c575b8115611dec575b506142dd565b87611d85600160ff195f5416175f55565b611dd557614340565b611d9457005b611da261ff00195f54165f55565b604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498908060208101611c1d565b611de761010061ff00195f5416175f55565b614340565b303b15915081611dfe575b505f611d6e565b60ff1660011490505f611df7565b600160ff8216109150611d67565b60409061071493928152816020820152019061081a565b3461034a57606036600319011261034a57600435611e4e816106a8565b602435604435611e5d816105c2565b611e9e611e6861244f565b9280611e73856124a3565b526040516361c8a12f60e11b81526001600160a01b0386169490925f91849182918760048401612ae2565b0381875afa9384156103455783611ec8611b42610db1611efd986020975f91611f5b575b506124a3565b92604051968794859384936304ec635160e01b85526004850163ffffffff604092959493606083019683521660208201520152565b03915afa801561034557611f2c925f91611f3c575b506001600160c01b031692611f2684614dc0565b90612711565b906107ab60405192839283611e1a565b611f55915060203d602011610f7b57610f68818361039d565b5f611f12565b611f6f91503d805f833e61107c818361039d565b5f611ec2565b3461034a575f36600319011261034a576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b3461034a57606036600319011261034a576004356001600160401b03811161034a57611fe9903690600401611935565b60243590611ff6826105c2565b6044356001600160401b03811161034a57612015903690600401610b03565b60cf5491939092916001600160a01b031633036120ff57610314936120e9936120696120709361204e6120466143a9565b9636906143ef565b86524363ffffffff16602087015263ffffffff166060860152565b36916107e4565b6040820152604051602081019061208b81611aa78585614449565b5190206120a0611ac460c95463ffffffff1690565b5560c95463ffffffff16907fa6b1912d066fd3a5413a6222a5785ae43f242458f3eef45c50dd808987b878fd604051806120e163ffffffff86169482614449565b0390a26135c5565b63ffffffff1663ffffffff1960c954161760c955565b60405162461bcd60e51b815260206004820152602160248201527f5461736b2067656e657261746f72206d757374206265207468652063616c6c656044820152603960f91b6064820152608490fd5b3461034a57602036600319011261034a5760043561216b816106a8565b612173614b66565b6001600160a01b0381161561218b5761031490614bbe565b60405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608490fd5b3461034a575f36600319011261034a57602060405160648152f35b3461034a57602036600319011261034a5760043560405163755b36bd60e11b81526020816004817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa908115610345575f916122bd575b506001600160a01b031633036122ae5761227c606654198219811614612315565b806066556040519081527f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c60203392a2005b63794821ff60e01b5f5260045ffd5b6122d6915060203d602011610aee57610ae0818361039d565b5f61225b565b9081602091031261034a5751801515810361034a5790565b6040513d5f823e3d90fd5b1561230657565b631d77d47760e21b5f5260045ffd5b1561231c57565b63c61dca5d60e01b5f5260045ffd5b634e487b7160e01b5f52603260045260245ffd5b9060028110156123505760051b0190565b61232b565b634e487b7160e01b5f52601260045260245ffd5b61244561242261244b9561241c61241585875160208901518a515160208c51015160208d016020815151915101519189519360208b0151956040519760208901998a5260208a015260408901526060880152608087015260a086015260c085015260e08401526101008301526123ec81610120840103601f19810183528261039d565b5190207f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001900690565b8096614565565b906145ab565b9261241c61243761243161460d565b94614704565b91612440614820565b614565565b9161489a565b9091565b60408051909190612460838261039d565b6001815291601f1901366020840137565b9061247b826106b9565b612488604051918261039d565b8281528092612499601f19916106b9565b0190602036910137565b8051156123505760200190565b80518210156123505760209160051b010190565b9081602091031261034a575190565b9190916124e08351612471565b925f5b81518110156125935780602061250c6124ff61253594866124b0565b516001600160a01b031690565b6040516309aa152760e11b81526001600160a01b03909116600482015292839081906024820190565b03816001600160a01b0388165afa8015610345576001925f91612565575b5061255e82886124b0565b52016124e3565b612586915060203d811161258c575b61257e818361039d565b8101906124c4565b5f612553565b503d612574565b505050565b9081602091031261034a5751610714816106a8565b906125b7826106b9565b6125c4604051918261039d565b828152602081936125d7601f19916106b9565b0191015f5b8281106125e857505050565b6060828201526020016125dc565b908151811015612350570160200190565b60208183031261034a578051906001600160401b03821161034a57019080601f8301121561034a57815161263a816106b9565b92612648604051948561039d565b81845260208085019260051b82010192831161034a57602001905b8282106126705750505090565b8151815260209182019101612663565b9061268a826106b9565b612697604051918261039d565b82815280926126a8601f19916106b9565b015f5b8181106126b757505050565b6040519060608201918083106001600160401b0384111761037d576020926040525f81525f838201525f6040820152828286010152016126ab565b9081602091031261034a57516001600160601b038116810361034a5790565b604051636830483560e01b815293919291906001600160a01b0316602085600481845afa948515610345575f95612a1a575b50604051634f4c91e160e11b815294602086600481855afa918215610345576004965f936129f8575b5060209060405197888092632efa2ca360e11b82525afa958615610345575f966129d7575b5061279f85939295516125ad565b945f935b80518510156129cd576127ca610e386127bc87846125f6565b516001600160f81b03191690565b604051638902624560e01b815260ff8216600482015263ffffffff88166024820152909490925f846044816001600160a01b0385165afa938415610345575f946129a9575b5061281a8451612680565b612824888b6124b0565b5261282f878a6124b0565b505f5b84518110156129985780602061284b61286d93886124b0565b518d60405180809681946308f6629d60e31b8352600483019190602083019252565b03916001600160a01b03165afa918215610345575f92612978575b5061289381876124b0565b518a60208a6128a2858b6124b0565b5160405163fa28c62760e01b8152600481019190915260ff91909116602482015263ffffffff929092166044830152816064816001600160a01b038d165afa938415610345576129358c8f6129306001986129419789975f92612948575b5061291b61290c6103df565b6001600160a01b039098168852565b60208701526001600160601b03166040860152565b6124b0565b5190610d8383836124b0565b5001612832565b61296a91925060203d8111612971575b612962818361039d565b8101906126f2565b905f612900565b503d612958565b61299191925060203d8111610aee57610ae0818361039d565b905f612888565b5060019096019590945091506127a3565b6129c69194503d805f833e6129be818361039d565b810190612607565b925f61280f565b5050509350505090565b6129f191965060203d602011610aee57610ae0818361039d565b945f612791565b6020919350612a1390823d8411610aee57610ae0818361039d565b929061276c565b612a3491955060203d602011610aee57610ae0818361039d565b935f612743565b60405190612a4882610382565b606080838181528160208201528160408201520152565b60208183031261034a578051906001600160401b03821161034a57019080601f8301121561034a578151612a92816106b9565b92612aa0604051948561039d565b81845260208085019260051b82010192831161034a57602001905b828210612ac85750505090565b602080918351612ad7816105c2565b815201910190612abb565b60409063ffffffff610714949316815281602082015201906106d0565b908060209392818452848401375f828201840152601f01601f1916010190565b60409063ffffffff61071495931681528160208201520191612aff565b634e487b7160e01b5f52601160045260245ffd5b60ff1660ff8114612b615760010190565b612b3c565b9081602091031261034a57516001600160c01b038116810361034a5790565b15612b8c57565b6325ec6c1f60e01b5f5260045ffd5b90821015612350570190565b9081602091031261034a5751610714816105c2565b5f198114612b615760010190565b91612be860209263ffffffff92969596604086526040860191612aff565b9416910152565b60405190612bfc82610362565b60606020838281520152565b15612c0f57565b62f8202d60e51b5f5260045ffd5b15612c2457565b6343714afd60e01b5f5260045ffd5b15612c3a57565b635f832f4160e01b5f5260045ffd5b15612c5057565b634b874f4560e01b5f5260045ffd5b9081602091031261034a575161071481611195565b5f19810191908211612b6157565b15612c8957565b633fdc650560e21b5f5260045ffd5b9081602091031261034a575167ffffffffffffffff198116810361034a5790565b15612cc057565b63e1310aed60e01b5f5260045ffd5b906001600160601b03809116911603906001600160601b038211612b6157565b15612cf657565b6367988d3360e01b5f5260045ffd5b15612d0c57565b63ab1b236b60e01b5f5260045ffd5b60049163ffffffff60e01b9060e01b1681520160208251919201905f5b818110612d455750505090565b8251845260209384019390920191600101612d38565b949392909193612d69612bef565b50612d75851515612c08565b60408401515185148061342b575b8061341d575b8061340f575b612d9890612c1d565b612daa60208501515185515114612c33565b612dc163ffffffff431663ffffffff841610612c49565b612dc96103d0565b5f81525f602082015292612ddb612bef565b612de487612471565b6020820152612df287612471565b8152612dfc612bef565b92612e0b602088015151612471565b8452612e1b602088015151612471565b602085810191909152604051639aa1653d60e01b815290816004817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa801561034557612e84915f916133e0575b50612e7f368b876107e4565b6149c8565b985f965b60208901518051891015612ff657602088612eeb610db18c612ee38f96868e612ec8612eb58680956124b0565b5180515f526020015160205260405f2090565b612ed584848401516124b0565b5282612fc3575b01516124b0565b5195516124b0565b6040516304ec635160e01b8152600481019490945263ffffffff9182166024850152166044830152816064816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa9182156103455761241c8a612f988f612f918f8460208f92612f8893612f808460019e612f9e9e5f91612fa6575b508f8060c01b031692516124b0565b5201516124b0565b51938d516124b0565b51166149f3565b90614a24565b970196612e88565b612fbd9150863d8111610f7b57610f68818361039d565b5f612f71565b612ff1612fd384848401516124b0565b51612fea84840151612fe487612c74565b906124b0565b5110612c82565b612edc565b5090959794965061300b919893929950614ae1565b915f905b808210613063575050509261304a61304561303e61305d9585611aa79860806060602099015192015192612369565b9190612cef565b612d05565b0151604051928391602083019586612d1b565b51902090565b91849596936130a0610e38610e2a858761309a879f989e9b612eb584604061308c9301516124b0565b67ffffffffffffffff191690565b96612b9b565b6020876130b4610db18d60a08d01516124b0565b604051631a2f32ab60e21b815260ff94909416600485015263ffffffff9182166024850152166044830152816064816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa928315610345576131509361313b925f916133b2575b5067ffffffffffffffff19918216911614612cb9565b6131498960408901516124b0565b51906145ab565b90613162610e38610e2a8a868d612b9b565b602086613176610db18c60c08c01516124b0565b604051636414a62b60e11b815260ff94909416600485015263ffffffff9182166024850152166044830152816064816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa801561034557611c3c8a60208e612edc61320c95613217975f91613395575b506131fe85858501516124b0565b906001600160601b03169052565b6131fe8a8d516124b0565b5f965f5b60208801515181101561338157898b613257613238848a516124b0565b5161324a610e38610e2a868c87612b9b565b60ff161c60019081161490565b613266575b505060010161321b565b88886132e76132a7610db1879f8f978060e08f849c6020613296610e38610e2a839f9861329e96610eec9a612b9b565b9a01516124b0565b519a01516124b0565b60405163795f4a5760e11b815260ff909316600484015263ffffffff93841660248401526044830195909552919093166064840152829081906084820190565b03817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa918215610345576133508e6001948e86955f9261335b575b5061334a6131fe92935193613345611c3c84876124b0565b612ccf565b926124b0565b01989050898b61325c565b6131fe925061337a61334a9160203d811161297157612962818361039d565b925061332d565b50959497600191999392949750019061300f565b6133ac9150843d811161297157612962818361039d565b5f6131f0565b6133d3915060203d81116133d9575b6133cb818361039d565b810190612c98565b5f613125565b503d6133c1565b613402915060203d602011613408575b6133fa818361039d565b810190612c5f565b5f612e73565b503d6133f0565b5060e0840151518514612d8f565b5060c0840151518514612d89565b5060a0840151518514612d83565b35610714816105c2565b1561344a57565b60405162461bcd60e51b815260206004820152602160248201527f5461736b206861736e2774206265656e20726573706f6e64656420746f2079656044820152601d60fa1b6064820152608490fd5b6020809163ffffffff81356134ad816105c2565b1684520135910152565b9160406103ce9294936134ce816080810197613499565b0190613499565b156134dc57565b60405162461bcd60e51b815260206004820152603d60248201527f5461736b20726573706f6e736520646f6573206e6f74206d617463682074686560448201527f206f6e65207265636f7264656420696e2074686520636f6e74726163740000006064820152608490fd5b1561354e57565b60405162461bcd60e51b815260206004820152604360248201527f54686520726573706f6e736520746f2074686973207461736b2068617320616c60448201527f7265616479206265656e206368616c6c656e676564207375636365737366756c606482015262363c9760e91b608482015260a490fd5b63ffffffff60019116019063ffffffff8211612b6157565b63ffffffff60649116019063ffffffff8211612b6157565b9063ffffffff8091169116019063ffffffff8211612b6157565b1561361657565b60405162461bcd60e51b815260206004820152603760248201527f546865206368616c6c656e676520706572696f6420666f72207468697320746160448201527f736b2068617320616c726561647920657870697265642e0000000000000000006064820152608490fd5b63ffffffff5f199116019063ffffffff8211612b6157565b156136a057565b60405162461bcd60e51b815260206004820152605360248201527f54686520737461746520726f6f74206f66207468652070726576696f7573207460448201527f61736b20646f6573206e6f74206d6174636820746865206f6e65207265636f72606482015272191959081a5b881d1a194818dbdb9d1c9858dd606a1b608482015260a490fd5b903590603e198136030182121561034a570190565b1561374357565b60405162461bcd60e51b815260206004820152605060248201527f546865207075626b657973206f66206e6f6e2d7369676e696e67206f7065726160448201527f746f727320737570706c69656420627920746865206368616c6c656e6765722060648201526f30b932903737ba1031b7b93932b1ba1760811b608482015260a490fd5b903590601e198136030182121561034a57018035906001600160401b03821161034a5760200191813603831361034a57565b60208183031261034a578051906001600160401b03821161034a57019080601f8301121561034a57815161382c816106b9565b9261383a604051948561039d565b81845260208085019260051b82010192831161034a57602001905b8282106138625750505090565b602080918351613871816106a8565b815201910190613855565b6040519061388b60408361039d565b601282527139b630b9b42fba3432afb7b832b930ba37b960711b6020830152565b805180835260209291819084018484015e5f828201840152601f01601f1916010190565b91906020835260c083019260018060a01b03825116602082015263ffffffff602083015116604082015260408201519360a060608301528451809152602060e083019501905f5b81811061395557505050608061394061071494956060850151601f1985830301848601526106d0565b9201519060a0601f19828503019101526138ac565b82516001600160a01b0316875260209687019690920191600101613917565b959192936020613a6860019261398985613439565b986139ad6139a58b63ffffffff165f5260cb60205260405f2090565b541515613443565b6139e86139c88b63ffffffff165f5260cb60205260405f2090565b54896040516139df81611aa78a8201948d866134b7565b519020146134d5565b613a13613a0d613a068c63ffffffff165f5260cd60205260405f2090565b5460ff1690565b15613547565b613a38613a2a611b42613a258b613439565b6135dd565b63ffffffff4316111561360f565b613a58613a458284614c2e565b613a51611bf98d613681565b5414613699565b613a628b80613727565b91614c9b565b920135821414613f7d57613a7c8451612471565b945f5b8551811015613aa95780613a98612eb5600193896124b0565b613aa2828a6124b0565b5201613a7f565b50909193949295613ae460208201976020613ac38a613439565b604051613ad881611aa78b8683019586612d1b565b5190209101351461373c565b613aee8651612471565b967f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316955f5b8851811015613ba357806020613b35613b55938a6124b0565b516040518094819263745dcd7360e11b8352600483019190602083019252565b03818c5afa8015610345578b613b7b91836001955f92613b81575b50610ab291926124b0565b01613b1c565b610ab29250613b9d9060203d8111610aee57610ae0818361039d565b91613b70565b50929596509296909350613c04613bcd613bd56040870195613bc587896137c7565b939091613439565b9236916107e4565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316612711565b955f925b8751841015613f055795965f965b613c2085836124b0565b5151881015613ef657613c6099602080613c3e8b610eec8a886124b0565b510151604051809d81926308f6629d60e31b8352600483019190602083019252565b0381875afa9a8b15610345575f9b613ed6575b506001995f5b8651811015613ec657613c9e613c926124ff838a6124b0565b6001600160a01b031690565b6001600160a01b038e1614613cb557600101613c79565b5098919299509960015f5b151514613cd4575b50600101969790613c16565b99918694925f89613d3b60ff613d139c9e969c610e38610e2a8d613d909f613d0d90613d0760d25460018060a01b031690565b986137c7565b90612b9b565b613d2d613d1e6103d0565b6001600160a01b039095168552565b1663ffffffff166020830152565b60d154613d5290613c92906001600160a01b031681565b60405163105dea1f60e21b815282516001600160a01b0316600482015260209092015163ffffffff16602483015290998a9190829081906044820190565b03915afa978815610345575f98613ea2575b50613dad8851612471565b995f5b8b51811015613dd6578067016345785d8a0000613dcf6001938f6124b0565b5201613db0565b5099909b929a9891939597613e1460ff613dfb610e38610e2a8c8f8f613d0d916137c7565b613e0661290c6103ee565b1663ffffffff166020860152565b60408401526060830152613e2661387c565b608083015260d054613e4290613c92906001600160a01b031681565b803b1561034a57604051636a669b4160e01b8152925f918491829084908290613e6e90600483016138d0565b03925af191821561034557600192613e88575b5090613cc8565b80613e965f613e9c9361039d565b80610578565b5f613e81565b613ebf9198503d805f833e613eb7818361039d565b8101906137f9565b965f613da2565b50989192996001909b919b613cc0565b613eef919b5060203d8111610aee57610ae0818361039d565b995f613c73565b90979650600190930192613c08565b979596505050505050613f36613f298363ffffffff165f5260cd60205260405f2090565b805460ff19166001179055565b613f4e8263ffffffff165f5260cc60205260405f2090565b5563ffffffff3391167fc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec5f80a3565b509350505063ffffffff3391167ffd3e26beeb5967fc5a57a0446914eabc45b4aa474c67a51b4b5160cac60ddb055f80a3565b9035601e198236030181121561034a5701602081359101916001600160401b03821161034a57813603831361034a57565b6107149161400e614003613ff58480613fb0565b604085526040850191612aff565b926020810190613fb0565b916020818503910152612aff565b602081528135603e198336030181121561034a5763ffffffff60606140856140526080948560208801528760a088019101613fe1565b836020880135614061816105c2565b1660408701526140746040880188613fb0565b878303601f19018589015290612aff565b940135614091816105c2565b1691015290565b1561409f57565b60405162461bcd60e51b815260206004820152603d60248201527f737570706c696564207461736b20646f6573206e6f74206d617463682074686560448201527f206f6e65207265636f7264656420696e2074686520636f6e74726163740000006064820152608490fd5b1561411157565b60405162461bcd60e51b815260206004820152602c60248201527f41676772656761746f722068617320616c726561647920726573706f6e64656460448201526b20746f20746865207461736b60a01b6064820152608490fd5b1561417257565b60405162461bcd60e51b815260206004820152602d60248201527f41676772656761746f722068617320726573706f6e64656420746f207468652060448201526c7461736b20746f6f206c61746560981b6064820152608490fd5b6040810192916103ce9190613499565b90606482029180830460641490151715612b6157565b90600682029180830460061490151715612b6157565b908160011b9180830460021490151715612b6157565b906001600160601b03809116911602906001600160601b038216918203612b6157565b1561424957565b608460405162461bcd60e51b815260206004820152604060248201527f5369676e61746f7269657320646f206e6f74206f776e206174206c656173742060448201527f7468726573686f6c642070657263656e74616765206f6620612071756f72756d6064820152fd5b90929160206060916142c9846080810197613499565b63ffffffff81511660408501520151910152565b156142e457565b60405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608490fd5b61434990614bbe565b60ce80546001600160a01b03199081166001600160a01b039384161790915560cf805482169383169390931790925560d1805483169382169390931790925560d0805482169383169390931790925560d280549092169216919091179055565b604051906143b682610382565b5f6060836143c2612bef565b81528260208201528160408201520152565b9080601f8301121561034a57816020610714933591016107e4565b919060408382031261034a576040519061440882610362565b819380356001600160401b03811161034a57826144269183016143d4565b83526020810135916001600160401b03831161034a576020926104d792016143d4565b60208152608063ffffffff60606144b16144908651856020880152602061447c8251604060a08b015260e08a01906138ac565b910151878203609f190160c08901526138ac565b8360208801511660408701526040870151601f1987830301848801526138ac565b9401511691015290565b5f196066556040515f1981527fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d60203392a2565b806066556040519081527fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d60203392a2565b6040519061452e82610362565b5f6020838281520152565b6040519061018061454a818461039d565b368337565b6040519061455e60208361039d565b6020368337565b91906040906060614574614521565b9485926020855192614586858561039d565b8436853780518452015160208301528482015260076107cf195a01fa156145a957565bfe5b6020929160806040926145bc614521565b958693818651936145cd868661039d565b85368637805185520151828401528051868401520151606082015260066107cf195a01fa80156145a957156145fe57565b63d4b68fd760e01b5f5260045ffd5b60405161461981610362565b6040908151614628838261039d565b823682378152602082519161463d848461039d565b833684370152805161464f828261039d565b7f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c281527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed60208201528151906146a5838361039d565b7f275dc4a288d1afb3cbb1ac09187524c7db36395df7be3b99e673b13a075a65ec82527f1d9befcd05a5323e6da4d435f3b617cdb3af83285c2df711ef39c01571827f9d60208301526146fa8351938461039d565b8252602082015290565b5f5160206150d15f395f51905f529061471b614521565b505f919006602060c0835b61481b575f935f5160206150d15f395f51905f5260038186818180090908604051614751858261039d565b84368237848185604051614765828261039d565b813682378381528360208201528360408201528560608201527f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f5260808201525f5160206150d15f395f51905f5260a082015260056107cf195a01fa80156145a9576147cf90615073565b519161481b575f5160206150d15f395f51905f528280091461480657505f5160206150d15f395f51905f5260015f94089293614726565b929350506148126103d0565b92835282015290565b612355565b614828614521565b5060405161483581610362565b600181526002602082015290565b9060018201809211612b6157565b9060028201809211612b6157565b9060038201809211612b6157565b9060048201809211612b6157565b9060058201809211612b6157565b90600c8110156123505760051b0190565b939290916148a860406103fd565b94855260208501526148ba60406103fd565b91825260208201526148ca614539565b925f5b600281106148f7575050506020610180926148e661454f565b93849160086201d4c0fa9151151590565b806149036001926141f3565b61490d828561233f565b51516149198289614889565b526020614926838661233f565b51015161493b61493583614843565b89614889565b52614946828661233f565b51515161495561493583614851565b5261496b614963838761233f565b515160200190565b516149786149358361485f565b526020614985838761233f565b510151516149956149358361486d565b526149c16149bb6149b460206149ab868a61233f565b51015160200190565b519261487b565b88614889565b52016148cd565b9060016149d660ff93614e74565b928392161b11156149e45790565b63ca95733360e01b5f5260045ffd5b805f915b6149ff575090565b5f198101818111612b615761ffff9116911661ffff8114612b615760010190806149f7565b90614a2d614521565b5061ffff811690610200821015614ad25760018214614acd57614a4e6103d0565b5f81525f602082015292906001905f925b61ffff8316851015614a7357505050505090565b600161ffff831660ff86161c811614614aad575b6001614aa3614a988360ff946145ab565b9460011b61fffe1690565b9401169291614a5f565b946001614aa3614a98614ac28960ff956145ab565b989350505050614a87565b505090565b637fc4ea7d60e11b5f5260045ffd5b614ae9614521565b50805190811580614b5a575b15614b16575050604051614b0a60408261039d565b5f81525f602082015290565b60205f5160206150d15f395f51905f52910151065f5160206150d15f395f51905f52035f5160206150d15f395f51905f528111612b6157604051916146fa83610362565b50602081015115614af5565b6033546001600160a01b03163303614b7a57565b606460405162461bcd60e51b815260206004820152602060248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152fd5b603380546001600160a01b039283166001600160a01b0319821681179092559091167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e05f80a3565b90821015612350576107149160051b810190613727565b906020610714928181520190613fe1565b614c3782612471565b915f5b818110614c5c57505050614c4d81614eec565b50805115612350576020015190565b80614c6a6001928486614c06565b604051614c8781611aa76020820194602086526020860190613fe1565b519020614c9482876124b0565b5201614c3a565b929190614caf614caa82614843565b612471565b915f945f925f915b818310614d0657505050614cd7939415614cdb575b5050611ec281614eec565b5190565b614cfe90604051614cf481611aa7602082019485614c1d565b51902091836124b0565b525f80614ccc565b9091938715614d53575b614d49600191614d21878686614c06565b604051614d3681611aa7602082019485614c1d565b519020614d43828a6124b0565b52612bbc565b945b019190614cb7565b614d7d614d6a614d64878686614c06565b806137c7565b90614d7587806137c7565b929091614f9c565b5f811215614d8c575b50614d10565b909750614dab6001986040516020810190614d3681611aa78a85614c1d565b9015614db7575f614d86565b93600190614d4b565b61ffff614dcc826149f3565b16614dd6816107c9565b90614de4604051928361039d565b808252614df3601f19916107c9565b013660208301375f5f5b8251821080614e53575b15614e4c576001811b8416614e25575b614e2090612bbc565b614dfd565b906001614e209160ff60f81b8460f81b165f1a614e4282876125f6565b5301919050614e17565b5050905090565b506101008110614e07565b15614e6557565b631019106960e31b5f5260045ffd5b90610100825111614edd57815115614ed857602082015160019060f81c81901b5b8351821015614ed357600190614ebe614eb4610e386127bc86896125f6565b60ff600191161b90565b90614eca818311614e5e565b17910190614e95565b925050565b5f9150565b637da54e4760e11b5f5260045ffd5b80515b60018111614f015750614cd7906124a3565b5f5b818110614f225750614f17614f1d91614843565b60011c90565b614eef565b80614f37614f31600193614209565b856124b0565b5183614f4a614f4584614209565b614843565b1015614f8257614f7090614f69614f63614f4585614209565b876124b0565b5190615089565b614f7a82866124b0565b525b01614f03565b80614f8c91615089565b614f9682866124b0565b52614f7c565b613bcd90614fac939236916107e4565b90805182518082105f1461506c57505b5f5b818110614fe85750505190519081811015614fda5750505f1990565b11614fe3575f90565b600190565b614ff56127bc82856125f6565b6150126150056127bc84886125f6565b6001600160f81b03191690565b6001600160f81b0319909116101561502d57505050505f1990565b61503a6127bc82856125f6565b61504a6150056127bc84886125f6565b6001600160f81b03199091161161506357600101614fbe565b50505050600190565b9050614fbc565b1561507a57565b63d51edae360e01b5f5260045ffd5b908082116150b05760408051602081019384529081019190915261305d8160608101611aa7565b6040805160208101928352908101929092529061305d8160608101611aa756fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47a264697066735822122067364904154f04e81ae76173c41ed8734a9692cb2a454e1fa7928aa52c1ddc8064736f6c634300081b0033",
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

// AllStateRoots is a free data retrieval call binding the contract method 0x5bcaa381.
//
// Solidity: function allStateRoots(uint32 ) view returns(bytes32)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerCaller) AllStateRoots(opts *bind.CallOpts, arg0 uint32) ([32]byte, error) {
	var out []interface{}
	err := _ContractAwesomeVaultTaskManager.contract.Call(opts, &out, "allStateRoots", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AllStateRoots is a free data retrieval call binding the contract method 0x5bcaa381.
//
// Solidity: function allStateRoots(uint32 ) view returns(bytes32)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerSession) AllStateRoots(arg0 uint32) ([32]byte, error) {
	return _ContractAwesomeVaultTaskManager.Contract.AllStateRoots(&_ContractAwesomeVaultTaskManager.CallOpts, arg0)
}

// AllStateRoots is a free data retrieval call binding the contract method 0x5bcaa381.
//
// Solidity: function allStateRoots(uint32 ) view returns(bytes32)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerCallerSession) AllStateRoots(arg0 uint32) ([32]byte, error) {
	return _ContractAwesomeVaultTaskManager.Contract.AllStateRoots(&_ContractAwesomeVaultTaskManager.CallOpts, arg0)
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

// TaskSuccessfullyChallenged is a free data retrieval call binding the contract method 0x1ce8a3e3.
//
// Solidity: function taskSuccessfullyChallenged(uint32 ) view returns(bool)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerCaller) TaskSuccessfullyChallenged(opts *bind.CallOpts, arg0 uint32) (bool, error) {
	var out []interface{}
	err := _ContractAwesomeVaultTaskManager.contract.Call(opts, &out, "taskSuccessfullyChallenged", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TaskSuccessfullyChallenged is a free data retrieval call binding the contract method 0x1ce8a3e3.
//
// Solidity: function taskSuccessfullyChallenged(uint32 ) view returns(bool)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerSession) TaskSuccessfullyChallenged(arg0 uint32) (bool, error) {
	return _ContractAwesomeVaultTaskManager.Contract.TaskSuccessfullyChallenged(&_ContractAwesomeVaultTaskManager.CallOpts, arg0)
}

// TaskSuccessfullyChallenged is a free data retrieval call binding the contract method 0x1ce8a3e3.
//
// Solidity: function taskSuccessfullyChallenged(uint32 ) view returns(bool)
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerCallerSession) TaskSuccessfullyChallenged(arg0 uint32) (bool, error) {
	return _ContractAwesomeVaultTaskManager.Contract.TaskSuccessfullyChallenged(&_ContractAwesomeVaultTaskManager.CallOpts, arg0)
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

// RaiseAndResolveChallenge is a paid mutator transaction binding the contract method 0xa8ec579b.
//
// Solidity: function raiseAndResolveChallenge(((string,string),uint32,bytes,uint32) task, (uint32,bytes32) taskResponse, (uint32,bytes32) taskResponseMetadata, (uint256,uint256)[] pubkeysOfNonSigningOperators, (string,string)[] prevStateLeaves) returns()
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerTransactor) RaiseAndResolveChallenge(opts *bind.TransactOpts, task IAwesomeVaultTaskManagerTask, taskResponse IAwesomeVaultTaskManagerTaskResponse, taskResponseMetadata IAwesomeVaultTaskManagerTaskResponseMetadata, pubkeysOfNonSigningOperators []BN254G1Point, prevStateLeaves []IAwesomeVaultTaskManagerTaskInput) (*types.Transaction, error) {
	return _ContractAwesomeVaultTaskManager.contract.Transact(opts, "raiseAndResolveChallenge", task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators, prevStateLeaves)
}

// RaiseAndResolveChallenge is a paid mutator transaction binding the contract method 0xa8ec579b.
//
// Solidity: function raiseAndResolveChallenge(((string,string),uint32,bytes,uint32) task, (uint32,bytes32) taskResponse, (uint32,bytes32) taskResponseMetadata, (uint256,uint256)[] pubkeysOfNonSigningOperators, (string,string)[] prevStateLeaves) returns()
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerSession) RaiseAndResolveChallenge(task IAwesomeVaultTaskManagerTask, taskResponse IAwesomeVaultTaskManagerTaskResponse, taskResponseMetadata IAwesomeVaultTaskManagerTaskResponseMetadata, pubkeysOfNonSigningOperators []BN254G1Point, prevStateLeaves []IAwesomeVaultTaskManagerTaskInput) (*types.Transaction, error) {
	return _ContractAwesomeVaultTaskManager.Contract.RaiseAndResolveChallenge(&_ContractAwesomeVaultTaskManager.TransactOpts, task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators, prevStateLeaves)
}

// RaiseAndResolveChallenge is a paid mutator transaction binding the contract method 0xa8ec579b.
//
// Solidity: function raiseAndResolveChallenge(((string,string),uint32,bytes,uint32) task, (uint32,bytes32) taskResponse, (uint32,bytes32) taskResponseMetadata, (uint256,uint256)[] pubkeysOfNonSigningOperators, (string,string)[] prevStateLeaves) returns()
func (_ContractAwesomeVaultTaskManager *ContractAwesomeVaultTaskManagerTransactorSession) RaiseAndResolveChallenge(task IAwesomeVaultTaskManagerTask, taskResponse IAwesomeVaultTaskManagerTaskResponse, taskResponseMetadata IAwesomeVaultTaskManagerTaskResponseMetadata, pubkeysOfNonSigningOperators []BN254G1Point, prevStateLeaves []IAwesomeVaultTaskManagerTaskInput) (*types.Transaction, error) {
	return _ContractAwesomeVaultTaskManager.Contract.RaiseAndResolveChallenge(&_ContractAwesomeVaultTaskManager.TransactOpts, task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators, prevStateLeaves)
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
