package examples

import (
	"context"
	"fmt"
	"math/big"

	"github.com/Layr-Labs/eigensdk-go/aggregator"
	"github.com/Layr-Labs/eigensdk-go/logging"
	sdkoperator "github.com/Layr-Labs/eigensdk-go/operator"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
	"github.com/ethereum/go-ethereum/core/types"

	cstaskmanager "github.com/Layr-Labs/eigensdk-go/examples/bindings/taskManager"
)

// The idea of this example is to show how to create a custom operator using the SDK generic implementation
type IncredibleSquaringTaskResponse struct {
	ReferenceTaskIndex uint32
	NumberSquared      *big.Int
}

func (tr IncredibleSquaringTaskResponse) TaskIndex() sdktypes.TaskIndex {
	return tr.ReferenceTaskIndex
}

func (tr IncredibleSquaringTaskResponse) Digest() [32]byte {
	return [32]byte(tr.NumberSquared.Bytes())
}

type OperatorTaskProcessor struct {
	logger logging.Logger
}

func NewOperatorTaskProcessor(c sdkoperator.OperatorConfig, logger logging.Logger) OperatorTaskProcessor {
	return OperatorTaskProcessor{
		logger: logger,
	}
}

// Takes a NewTaskCreatedLog struct as input and returns a TaskResponseHeader struct.
// The TaskResponseHeader struct is the struct that is signed and sent to the contract as a task response.
func (otp OperatorTaskProcessor) ProcessNewTaskCreatedLog(
	log types.Log,
) (aggregator.TaskResponse, error) {
	var newTaskCreatedLog cstaskmanager.ContractIncredibleSquaringTaskManagerNewTaskCreated

	taskManagerAbi, err := cstaskmanager.ContractIncredibleSquaringTaskManagerMetaData.GetAbi()
	if err != nil {
		otp.logger.Fatalf("Error obtaining task manager ABI: %v", err)
	}

	err = taskManagerAbi.UnpackIntoInterface(&newTaskCreatedLog, "NewTaskCreated", log.Data)
	if err != nil {
		return nil, fmt.Errorf("error unpacking the log: %w", err)
	}

	numberSquared := big.NewInt(0).Exp(newTaskCreatedLog.Task.NumberToBeSquared, big.NewInt(2), nil)

	taskResponse := IncredibleSquaringTaskResponse{
		ReferenceTaskIndex: newTaskCreatedLog.TaskIndex,
		NumberSquared:      numberSquared,
	}
	return taskResponse, nil
}

func main() {
	logger, err := logging.NewZapLogger(logging.Production) // Change here if want to change logging level
	if err != nil {
		println("Failure creating logger")
		return
	}

	taskManagerAbi, err := cstaskmanager.ContractIncredibleSquaringTaskManagerMetaData.GetAbi()
	if err != nil {
		logger.Fatalf(err.Error())
	}

	blockHash := taskManagerAbi.Events["NewTaskCreated"].ID

	// The values from this config are extracted from an incredible squaring config file:
	// https://github.com/Layr-Labs/incredible-squaring-avs/blob/dev/config-files/operator.anvil.yaml
	operatorConfig := sdkoperator.OperatorConfig{
		OperatorAddress:               "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266",
		OperatorStateRetrieverAddress: "0x4c5859f0f772848b2d91f1d83e2fe57935348029",
		ServiceManagerAddress:         "0x5f3f1dbd7b74c6b46e8c44f98792a1daf8d69154",
		AVSRegistryCoordinatorAddress: "0x7bc06c482dead17c0e297afbc32f6e63d3846650",
		EthRpcUrl:                     "http://localhost:8545",
		EthWsUrl:                      "ws://localhost:8545",
		BlsPrivateKeyStorePath:        "tests/keys/test.bls.key.json",
		AggregatorServerIpPortAddress: "localhost:8090",
	}
	operatorTaskProcessor := NewOperatorTaskProcessor(operatorConfig, logger)
	operator, err := sdkoperator.NewOperatorFromConfig(operatorConfig, blockHash, operatorTaskProcessor, logger)
	if err != nil {
		logger.Errorf("Failed to create operator from config: %v", err)
		return
	}

	err = operator.Start(context.Background(), &IncredibleSquaringTaskResponse{})
	if err != nil {
		logger.Errorf("Error while running operator: %v", err)
		return
	}
}
