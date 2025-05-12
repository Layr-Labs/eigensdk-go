package taskprocessor

import (
	"math/big"
	"strings"
	"testing"

	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
	"github.com/Layr-Labs/eigensdk-go/utils"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	istaskmanager "github.com/Layr-Labs/eigensdk-go/examples/incredible-squaring/bindings/taskManager"
)

// Testing structs

type U32Point struct {
	X uint32
	Y uint32
}

type DoubleBigIntVector struct {
	First  []*big.Int
	Second []*big.Int
}

type DoubleStringVector struct {
	First  []string
	Second []string
}

type Arg0 struct {
	Data []byte
}

type Arg1 struct {
	Arg0 Arg0
	Data []byte
}

type Arg2 struct {
	Arg1 Arg1
	Data []byte
}

func TestIncredibleSquaringAbi(t *testing.T) {
	taskManagerAbi, err := istaskmanager.ContractIncredibleSquaringTaskManagerMetaData.GetAbi()
	require.NoError(t, err)

	originalTaskStruct := sdktypes.GenericInputTask[*big.Int]{
		InputValue:                big.NewInt(10),
		TaskCreatedBlock:          10,
		QuorumNumbers:             []uint8{0},
		QuorumThresholdPercentage: 100,
	}

	originalTaskResponseStruct := sdktypes.GenericOutputTaskResponse[*big.Int]{
		ReferenceTaskIndex: 0,
		OutputValue:        big.NewInt(100),
	}

	nonSigStruct := sdktypes.NonSignerStakesAndSignature{
		NonSignerQuorumBitmapIndices: []uint32{0},
		NonSignerPubkeys:             []sdktypes.BN254G1Point{{X: common.Big0, Y: common.Big0}},
		QuorumApks:                   []sdktypes.BN254G1Point{{X: common.Big0, Y: common.Big0}},
		ApkG2:                        sdktypes.BN254G2Point{X: [2]*big.Int{common.Big0, common.Big0}, Y: [2]*big.Int{common.Big0, common.Big0}},
		Sigma:                        sdktypes.BN254G1Point{X: common.Big0, Y: common.Big0},
		QuorumApkIndices:             []uint32{0},
		TotalStakeIndices:            []uint32{0},
		NonSignerStakeIndices:        [][]uint32{{0}},
	}

	taskReflectStruct := utils.CopyStructAndChangeFieldName(originalTaskStruct, "InputValue", "NumberToBeSquared")
	taskResponseReflectStruct := utils.CopyStructAndChangeFieldName(originalTaskResponseStruct, "OutputValue", "NumberSquared")

	packedBytes, err := taskManagerAbi.Pack("respondToTask", taskReflectStruct, taskResponseReflectStruct, nonSigStruct)
	require.NoError(t, err)
	require.NotZero(t, packedBytes)
}

func TestSimpleStructValueAbi(t *testing.T) {
	rawAbi := "[\n{\n\"inputs\":[\n{\n\"components\":[\n{\n\"components\":[\n{\n\"internalType\":\"uint32\",\n\"name\":\"X\",\n\"type\":\"uint32\"\n},\n{\n\"internalType\":\"uint32\",\n\"name\":\"Y\",\n\"type\":\"uint32\"\n}\n],\n\"internalType\":\"structInputPoint\",\n\"name\":\"inputPoint\",\n\"type\":\"tuple\"\n},\n{\n\"internalType\":\"uint32\",\n\"name\":\"taskCreatedBlock\",\n\"type\":\"uint32\"\n},\n{\n\"internalType\":\"bytes\",\n\"name\":\"quorumNumbers\",\n\"type\":\"bytes\"\n},\n{\n\"internalType\":\"uint32\",\n\"name\":\"quorumThresholdPercentage\",\n\"type\":\"uint32\"\n}\n],\n\"internalType\":\"structIIncredibleSquaringTaskManager.Task\",\n\"name\":\"task\",\n\"type\":\"tuple\"\n},\n{\n\"components\":[\n{\n\"internalType\":\"uint32\",\n\"name\":\"referenceTaskIndex\",\n\"type\":\"uint32\"\n},\n{\n\"components\":[\n{\n\"internalType\":\"uint32\",\n\"name\":\"X\",\n\"type\":\"uint32\"\n},\n{\n\"internalType\":\"uint32\",\n\"name\":\"Y\",\n\"type\":\"uint32\"\n}\n],\n\"internalType\":\"structOutputPoint\",\n\"name\":\"outputPoint\",\n\"type\":\"tuple\"\n}\n],\n\"internalType\":\"structIIncredibleSquaringTaskManager.TaskResponse\",\n\"name\":\"taskResponse\",\n\"type\":\"tuple\"\n},\n{\n\"components\":[\n{\n\"internalType\":\"uint32[]\",\n\"name\":\"nonSignerQuorumBitmapIndices\",\n\"type\":\"uint32[]\"\n},\n{\n\"components\":[\n{\n\"internalType\":\"uint256\",\n\"name\":\"X\",\n\"type\":\"uint256\"\n},\n{\n\"internalType\":\"uint256\",\n\"name\":\"Y\",\n\"type\":\"uint256\"\n}\n],\n\"internalType\":\"structBN254.G1Point[]\",\n\"name\":\"nonSignerPubkeys\",\n\"type\":\"tuple[]\"\n},\n{\n\"components\":[\n{\n\"internalType\":\"uint256\",\n\"name\":\"X\",\n\"type\":\"uint256\"\n},\n{\n\"internalType\":\"uint256\",\n\"name\":\"Y\",\n\"type\":\"uint256\"\n}\n],\n\"internalType\":\"structBN254.G1Point[]\",\n\"name\":\"quorumApks\",\n\"type\":\"tuple[]\"\n},\n{\n\"components\":[\n{\n\"internalType\":\"uint256[2]\",\n\"name\":\"X\",\n\"type\":\"uint256[2]\"\n},\n{\n\"internalType\":\"uint256[2]\",\n\"name\":\"Y\",\n\"type\":\"uint256[2]\"\n}\n],\n\"internalType\":\"structBN254.G2Point\",\n\"name\":\"apkG2\",\n\"type\":\"tuple\"\n},\n{\n\"components\":[\n{\n\"internalType\":\"uint256\",\n\"name\":\"X\",\n\"type\":\"uint256\"\n},\n{\n\"internalType\":\"uint256\",\n\"name\":\"Y\",\n\"type\":\"uint256\"\n}\n],\n\"internalType\":\"structBN254.G1Point\",\n\"name\":\"sigma\",\n\"type\":\"tuple\"\n},\n{\n\"internalType\":\"uint32[]\",\n\"name\":\"quorumApkIndices\",\n\"type\":\"uint32[]\"\n},\n{\n\"internalType\":\"uint32[]\",\n\"name\":\"totalStakeIndices\",\n\"type\":\"uint32[]\"\n},\n{\n\"internalType\":\"uint32[][]\",\n\"name\":\"nonSignerStakeIndices\",\n\"type\":\"uint32[][]\"\n}\n],\n\"internalType\":\"structBLSSignatureChecker.NonSignerStakesAndSignature\",\n\"name\":\"nonSignerStakesAndSignature\",\n\"type\":\"tuple\"\n}\n],\n\"name\":\"respondToTask\",\n\"outputs\":[],\n\"stateMutability\":\"nonpayable\",\n\"type\":\"function\"\n}\n]"
	parsedAbi, err := abi.JSON(strings.NewReader(rawAbi))
	require.NoError(t, err)

	originalTaskStruct := sdktypes.GenericInputTask[U32Point]{
		InputValue:                U32Point{0, 1},
		TaskCreatedBlock:          10,
		QuorumNumbers:             []uint8{0},
		QuorumThresholdPercentage: 100,
	}

	originalTaskResponseStruct := sdktypes.GenericOutputTaskResponse[U32Point]{
		ReferenceTaskIndex: 0,
		OutputValue:        U32Point{2, 3},
	}

	nonSigStruct := sdktypes.NonSignerStakesAndSignature{
		NonSignerQuorumBitmapIndices: []uint32{0},
		NonSignerPubkeys:             []sdktypes.BN254G1Point{{X: common.Big0, Y: common.Big0}},
		QuorumApks:                   []sdktypes.BN254G1Point{{X: common.Big0, Y: common.Big0}},
		ApkG2:                        sdktypes.BN254G2Point{X: [2]*big.Int{common.Big0, common.Big0}, Y: [2]*big.Int{common.Big0, common.Big0}},
		Sigma:                        sdktypes.BN254G1Point{X: common.Big0, Y: common.Big0},
		QuorumApkIndices:             []uint32{0},
		TotalStakeIndices:            []uint32{0},
		NonSignerStakeIndices:        [][]uint32{{0}},
	}

	taskReflectStruct := utils.CopyStructAndChangeFieldName(originalTaskStruct, "InputValue", "InputPoint")
	taskResponseReflectStruct := utils.CopyStructAndChangeFieldName(originalTaskResponseStruct, "OutputValue", "OutputPoint")

	packedBytes, err := parsedAbi.Pack("respondToTask", taskReflectStruct, taskResponseReflectStruct, nonSigStruct)
	require.NoError(t, err)
	require.NotZero(t, packedBytes)
}

func TestStructureOfVectorsAbi(t *testing.T) {
	rawAbi := "[\n{\n\"inputs\":[\n{\n\"components\":[\n{\n\"components\":[\n{\n\"internalType\":\"uint256[]\",\n\"name\":\"First\",\n\"type\":\"uint256[]\"\n},\n{\n\"internalType\":\"uint256[]\",\n\"name\":\"Second\",\n\"type\":\"uint256[]\"\n}\n],\n\"internalType\":\"structDoubleBigIntVector\",\n\"name\":\"doubleVector\",\n\"type\":\"tuple\"\n},\n{\n\"internalType\":\"uint32\",\n\"name\":\"taskCreatedBlock\",\n\"type\":\"uint32\"\n},\n{\n\"internalType\":\"bytes\",\n\"name\":\"quorumNumbers\",\n\"type\":\"bytes\"\n},\n{\n\"internalType\":\"uint32\",\n\"name\":\"quorumThresholdPercentage\",\n\"type\":\"uint32\"\n}\n],\n\"internalType\":\"structIIncredibleSquaringTaskManager.Task\",\n\"name\":\"task\",\n\"type\":\"tuple\"\n},\n{\n\"components\":[\n{\n\"internalType\":\"uint32\",\n\"name\":\"referenceTaskIndex\",\n\"type\":\"uint32\"\n},\n{\n\"components\":[\n{\n\"internalType\":\"string[]\",\n\"name\":\"First\",\n\"type\":\"string[]\"\n},\n{\n\"internalType\":\"string[]\",\n\"name\":\"Second\",\n\"type\":\"string[]\"\n}\n],\n\"internalType\":\"structDoubleStringVector\",\n\"name\":\"doubleVector\",\n\"type\":\"tuple\"\n}\n],\n\"internalType\":\"structIIncredibleSquaringTaskManager.TaskResponse\",\n\"name\":\"taskResponse\",\n\"type\":\"tuple\"\n},\n{\n\"components\":[\n{\n\"internalType\":\"uint32[]\",\n\"name\":\"nonSignerQuorumBitmapIndices\",\n\"type\":\"uint32[]\"\n},\n{\n\"components\":[\n{\n\"internalType\":\"uint256\",\n\"name\":\"X\",\n\"type\":\"uint256\"\n},\n{\n\"internalType\":\"uint256\",\n\"name\":\"Y\",\n\"type\":\"uint256\"\n}\n],\n\"internalType\":\"structBN254.G1Point[]\",\n\"name\":\"nonSignerPubkeys\",\n\"type\":\"tuple[]\"\n},\n{\n\"components\":[\n{\n\"internalType\":\"uint256\",\n\"name\":\"X\",\n\"type\":\"uint256\"\n},\n{\n\"internalType\":\"uint256\",\n\"name\":\"Y\",\n\"type\":\"uint256\"\n}\n],\n\"internalType\":\"structBN254.G1Point[]\",\n\"name\":\"quorumApks\",\n\"type\":\"tuple[]\"\n},\n{\n\"components\":[\n{\n\"internalType\":\"uint256[2]\",\n\"name\":\"X\",\n\"type\":\"uint256[2]\"\n},\n{\n\"internalType\":\"uint256[2]\",\n\"name\":\"Y\",\n\"type\":\"uint256[2]\"\n}\n],\n\"internalType\":\"structBN254.G2Point\",\n\"name\":\"apkG2\",\n\"type\":\"tuple\"\n},\n{\n\"components\":[\n{\n\"internalType\":\"uint256\",\n\"name\":\"X\",\n\"type\":\"uint256\"\n},\n{\n\"internalType\":\"uint256\",\n\"name\":\"Y\",\n\"type\":\"uint256\"\n}\n],\n\"internalType\":\"structBN254.G1Point\",\n\"name\":\"sigma\",\n\"type\":\"tuple\"\n},\n{\n\"internalType\":\"uint32[]\",\n\"name\":\"quorumApkIndices\",\n\"type\":\"uint32[]\"\n},\n{\n\"internalType\":\"uint32[]\",\n\"name\":\"totalStakeIndices\",\n\"type\":\"uint32[]\"\n},\n{\n\"internalType\":\"uint32[][]\",\n\"name\":\"nonSignerStakeIndices\",\n\"type\":\"uint32[][]\"\n}\n],\n\"internalType\":\"structBLSSignatureChecker.NonSignerStakesAndSignature\",\n\"name\":\"nonSignerStakesAndSignature\",\n\"type\":\"tuple\"\n}\n],\n\"name\":\"respondToTask\",\n\"outputs\":[],\n\"stateMutability\":\"nonpayable\",\n\"type\":\"function\"\n}\n]"
	parsedAbi, err := abi.JSON(strings.NewReader(rawAbi))
	require.NoError(t, err)

	originalTaskStruct := sdktypes.GenericInputTask[DoubleBigIntVector]{
		InputValue:                DoubleBigIntVector{[]*big.Int{common.Big1}, []*big.Int{common.Big2}},
		TaskCreatedBlock:          10,
		QuorumNumbers:             []uint8{0},
		QuorumThresholdPercentage: 100,
	}

	originalTaskResponseStruct := sdktypes.GenericOutputTaskResponse[DoubleStringVector]{
		ReferenceTaskIndex: 0,
		OutputValue:        DoubleStringVector{[]string{"1"}, []string{"2"}},
	}

	nonSigStruct := sdktypes.NonSignerStakesAndSignature{
		NonSignerQuorumBitmapIndices: []uint32{0},
		NonSignerPubkeys:             []sdktypes.BN254G1Point{{X: common.Big0, Y: common.Big0}},
		QuorumApks:                   []sdktypes.BN254G1Point{{X: common.Big0, Y: common.Big0}},
		ApkG2:                        sdktypes.BN254G2Point{X: [2]*big.Int{common.Big0, common.Big0}, Y: [2]*big.Int{common.Big0, common.Big0}},
		Sigma:                        sdktypes.BN254G1Point{X: common.Big0, Y: common.Big0},
		QuorumApkIndices:             []uint32{0},
		TotalStakeIndices:            []uint32{0},
		NonSignerStakeIndices:        [][]uint32{{0}},
	}

	taskReflectStruct := utils.CopyStructAndChangeFieldName(originalTaskStruct, "InputValue", "DoubleVector")
	taskResponseReflectStruct := utils.CopyStructAndChangeFieldName(originalTaskResponseStruct, "OutputValue", "DoubleVector")

	packedBytes, err := parsedAbi.Pack("respondToTask", taskReflectStruct, taskResponseReflectStruct, nonSigStruct)
	require.NoError(t, err)
	require.NotZero(t, packedBytes)
}

func TestAddressAndBoolValuesAbi(t *testing.T) {
	rawAbi := "[\n{\n\"inputs\":[\n{\n\"components\":[\n{\n\"internalType\":\"address\",\n\"name\":\"queryAddress\",\n\"type\":\"address\"\n},\n{\n\"internalType\":\"uint32\",\n\"name\":\"taskCreatedBlock\",\n\"type\":\"uint32\"\n},\n{\n\"internalType\":\"bytes\",\n\"name\":\"quorumNumbers\",\n\"type\":\"bytes\"\n},\n{\n\"internalType\":\"uint32\",\n\"name\":\"quorumThresholdPercentage\",\n\"type\":\"uint32\"\n}\n],\n\"internalType\":\"structIIncredibleSquaringTaskManager.Task\",\n\"name\":\"task\",\n\"type\":\"tuple\"\n},\n{\n\"components\":[\n{\n\"internalType\":\"uint32\",\n\"name\":\"referenceTaskIndex\",\n\"type\":\"uint32\"\n},\n{\n\"internalType\":\"bool\",\n\"name\":\"isRegistered\",\n\"type\":\"bool\"\n}\n],\n\"internalType\":\"structIIncredibleSquaringTaskManager.TaskResponse\",\n\"name\":\"taskResponse\",\n\"type\":\"tuple\"\n},\n{\n\"components\":[\n{\n\"internalType\":\"uint32[]\",\n\"name\":\"nonSignerQuorumBitmapIndices\",\n\"type\":\"uint32[]\"\n},\n{\n\"components\":[\n{\n\"internalType\":\"uint256\",\n\"name\":\"X\",\n\"type\":\"uint256\"\n},\n{\n\"internalType\":\"uint256\",\n\"name\":\"Y\",\n\"type\":\"uint256\"\n}\n],\n\"internalType\":\"structBN254.G1Point[]\",\n\"name\":\"nonSignerPubkeys\",\n\"type\":\"tuple[]\"\n},\n{\n\"components\":[\n{\n\"internalType\":\"uint256\",\n\"name\":\"X\",\n\"type\":\"uint256\"\n},\n{\n\"internalType\":\"uint256\",\n\"name\":\"Y\",\n\"type\":\"uint256\"\n}\n],\n\"internalType\":\"structBN254.G1Point[]\",\n\"name\":\"quorumApks\",\n\"type\":\"tuple[]\"\n},\n{\n\"components\":[\n{\n\"internalType\":\"uint256[2]\",\n\"name\":\"X\",\n\"type\":\"uint256[2]\"\n},\n{\n\"internalType\":\"uint256[2]\",\n\"name\":\"Y\",\n\"type\":\"uint256[2]\"\n}\n],\n\"internalType\":\"structBN254.G2Point\",\n\"name\":\"apkG2\",\n\"type\":\"tuple\"\n},\n{\n\"components\":[\n{\n\"internalType\":\"uint256\",\n\"name\":\"X\",\n\"type\":\"uint256\"\n},\n{\n\"internalType\":\"uint256\",\n\"name\":\"Y\",\n\"type\":\"uint256\"\n}\n],\n\"internalType\":\"structBN254.G1Point\",\n\"name\":\"sigma\",\n\"type\":\"tuple\"\n},\n{\n\"internalType\":\"uint32[]\",\n\"name\":\"quorumApkIndices\",\n\"type\":\"uint32[]\"\n},\n{\n\"internalType\":\"uint32[]\",\n\"name\":\"totalStakeIndices\",\n\"type\":\"uint32[]\"\n},\n{\n\"internalType\":\"uint32[][]\",\n\"name\":\"nonSignerStakeIndices\",\n\"type\":\"uint32[][]\"\n}\n],\n\"internalType\":\"structBLSSignatureChecker.NonSignerStakesAndSignature\",\n\"name\":\"nonSignerStakesAndSignature\",\n\"type\":\"tuple\"\n}\n],\n\"name\":\"respondToTask\",\n\"outputs\":[],\n\"stateMutability\":\"nonpayable\",\n\"type\":\"function\"\n}\n]"
	parsedAbi, err := abi.JSON(strings.NewReader(rawAbi))
	require.NoError(t, err)

	originalTaskStruct := sdktypes.GenericInputTask[common.Address]{
		InputValue:                common.HexToAddress("0x01"),
		TaskCreatedBlock:          10,
		QuorumNumbers:             []uint8{0},
		QuorumThresholdPercentage: 100,
	}

	originalTaskResponseStruct := sdktypes.GenericOutputTaskResponse[bool]{
		ReferenceTaskIndex: 0,
		OutputValue:        true,
	}

	nonSigStruct := sdktypes.NonSignerStakesAndSignature{
		NonSignerQuorumBitmapIndices: []uint32{0},
		NonSignerPubkeys:             []sdktypes.BN254G1Point{{X: common.Big0, Y: common.Big0}},
		QuorumApks:                   []sdktypes.BN254G1Point{{X: common.Big0, Y: common.Big0}},
		ApkG2:                        sdktypes.BN254G2Point{X: [2]*big.Int{common.Big0, common.Big0}, Y: [2]*big.Int{common.Big0, common.Big0}},
		Sigma:                        sdktypes.BN254G1Point{X: common.Big0, Y: common.Big0},
		QuorumApkIndices:             []uint32{0},
		TotalStakeIndices:            []uint32{0},
		NonSignerStakeIndices:        [][]uint32{{0}},
	}

	taskReflectStruct := utils.CopyStructAndChangeFieldName(originalTaskStruct, "InputValue", "QueryAddress")
	taskResponseReflectStruct := utils.CopyStructAndChangeFieldName(originalTaskResponseStruct, "OutputValue", "IsRegistered")

	packedBytes, err := parsedAbi.Pack("respondToTask", taskReflectStruct, taskResponseReflectStruct, nonSigStruct)
	require.NoError(t, err)
	require.NotZero(t, packedBytes)
}

func TestNestedStructsAbi(t *testing.T) {
	rawAbi := "[\n{\n\"inputs\":[\n{\n\"components\":[\n{\n\"components\":[\n{\n\"components\":[\n{\n\"components\":[\n{\n\"internalType\":\"bytes\",\n\"name\":\"Data\",\n\"type\":\"bytes\"\n}\n],\n\"internalType\":\"structArg0\",\n\"name\":\"Arg0\",\n\"type\":\"tuple\"\n},\n{\n\"internalType\":\"bytes\",\n\"name\":\"Data\",\n\"type\":\"bytes\"\n}\n],\n\"internalType\":\"structArg1\",\n\"name\":\"Arg1\",\n\"type\":\"tuple\"\n},\n{\n\"internalType\":\"bytes\",\n\"name\":\"Data\",\n\"type\":\"bytes\"\n}\n],\n\"internalType\":\"structArg2\",\n\"name\":\"mainArg\",\n\"type\":\"tuple\"\n},\n{\n\"internalType\":\"uint32\",\n\"name\":\"taskCreatedBlock\",\n\"type\":\"uint32\"\n},\n{\n\"internalType\":\"bytes\",\n\"name\":\"quorumNumbers\",\n\"type\":\"bytes\"\n},\n{\n\"internalType\":\"uint32\",\n\"name\":\"quorumThresholdPercentage\",\n\"type\":\"uint32\"\n}\n],\n\"internalType\":\"structIIncredibleSquaringTaskManager.Task\",\n\"name\":\"task\",\n\"type\":\"tuple\"\n},\n{\n\"components\":[\n{\n\"internalType\":\"uint32\",\n\"name\":\"referenceTaskIndex\",\n\"type\":\"uint32\"\n},\n{\n\"internalType\":\"bytes\",\n\"name\":\"Total\",\n\"type\":\"bytes\"\n}\n],\n\"internalType\":\"structIIncredibleSquaringTaskManager.TaskResponse\",\n\"name\":\"taskResponse\",\n\"type\":\"tuple\"\n},\n{\n\"components\":[\n{\n\"internalType\":\"uint32[]\",\n\"name\":\"nonSignerQuorumBitmapIndices\",\n\"type\":\"uint32[]\"\n},\n{\n\"components\":[\n{\n\"internalType\":\"uint256\",\n\"name\":\"X\",\n\"type\":\"uint256\"\n},\n{\n\"internalType\":\"uint256\",\n\"name\":\"Y\",\n\"type\":\"uint256\"\n}\n],\n\"internalType\":\"structBN254.G1Point[]\",\n\"name\":\"nonSignerPubkeys\",\n\"type\":\"tuple[]\"\n},\n{\n\"components\":[\n{\n\"internalType\":\"uint256\",\n\"name\":\"X\",\n\"type\":\"uint256\"\n},\n{\n\"internalType\":\"uint256\",\n\"name\":\"Y\",\n\"type\":\"uint256\"\n}\n],\n\"internalType\":\"structBN254.G1Point[]\",\n\"name\":\"quorumApks\",\n\"type\":\"tuple[]\"\n},\n{\n\"components\":[\n{\n\"internalType\":\"uint256[2]\",\n\"name\":\"X\",\n\"type\":\"uint256[2]\"\n},\n{\n\"internalType\":\"uint256[2]\",\n\"name\":\"Y\",\n\"type\":\"uint256[2]\"\n}\n],\n\"internalType\":\"structBN254.G2Point\",\n\"name\":\"apkG2\",\n\"type\":\"tuple\"\n},\n{\n\"components\":[\n{\n\"internalType\":\"uint256\",\n\"name\":\"X\",\n\"type\":\"uint256\"\n},\n{\n\"internalType\":\"uint256\",\n\"name\":\"Y\",\n\"type\":\"uint256\"\n}\n],\n\"internalType\":\"structBN254.G1Point\",\n\"name\":\"sigma\",\n\"type\":\"tuple\"\n},\n{\n\"internalType\":\"uint32[]\",\n\"name\":\"quorumApkIndices\",\n\"type\":\"uint32[]\"\n},\n{\n\"internalType\":\"uint32[]\",\n\"name\":\"totalStakeIndices\",\n\"type\":\"uint32[]\"\n},\n{\n\"internalType\":\"uint32[][]\",\n\"name\":\"nonSignerStakeIndices\",\n\"type\":\"uint32[][]\"\n}\n],\n\"internalType\":\"structBLSSignatureChecker.NonSignerStakesAndSignature\",\n\"name\":\"nonSignerStakesAndSignature\",\n\"type\":\"tuple\"\n}\n],\n\"name\":\"respondToTask\",\n\"outputs\":[],\n\"stateMutability\":\"nonpayable\",\n\"type\":\"function\"\n}\n]"

	parsedAbi, err := abi.JSON(strings.NewReader(rawAbi))
	require.NoError(t, err)

	originalTaskStruct := sdktypes.GenericInputTask[Arg2]{
		InputValue:                Arg2{Arg1: Arg1{Arg0: Arg0{Data: []byte{64}}, Data: []byte{32}}, Data: []byte{16}},
		TaskCreatedBlock:          10,
		QuorumNumbers:             []uint8{0},
		QuorumThresholdPercentage: 100,
	}

	originalTaskResponseStruct := sdktypes.GenericOutputTaskResponse[[]byte]{
		ReferenceTaskIndex: 0,
		OutputValue:        []byte{112},
	}

	nonSigStruct := sdktypes.NonSignerStakesAndSignature{
		NonSignerQuorumBitmapIndices: []uint32{0},
		NonSignerPubkeys:             []sdktypes.BN254G1Point{{X: common.Big0, Y: common.Big0}},
		QuorumApks:                   []sdktypes.BN254G1Point{{X: common.Big0, Y: common.Big0}},
		ApkG2:                        sdktypes.BN254G2Point{X: [2]*big.Int{common.Big0, common.Big0}, Y: [2]*big.Int{common.Big0, common.Big0}},
		Sigma:                        sdktypes.BN254G1Point{X: common.Big0, Y: common.Big0},
		QuorumApkIndices:             []uint32{0},
		TotalStakeIndices:            []uint32{0},
		NonSignerStakeIndices:        [][]uint32{{0}},
	}

	taskReflectStruct := utils.CopyStructAndChangeFieldName(originalTaskStruct, "InputValue", "MainArg")
	taskResponseReflectStruct := utils.CopyStructAndChangeFieldName(originalTaskResponseStruct, "OutputValue", "Total")

	packedBytes, err := parsedAbi.Pack("respondToTask", taskReflectStruct, taskResponseReflectStruct, nonSigStruct)
	require.NoError(t, err)
	require.NotZero(t, packedBytes)
}
