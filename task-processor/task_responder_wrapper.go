package taskprocessor

import (
	"context"
	"reflect"
	"strings"

	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	internalutils "github.com/Layr-Labs/eigensdk-go/internal/utils"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
	"github.com/Layr-Labs/eigensdk-go/utils"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

var _ TaskResponder[any, any] = (*taskResponderContractWrapper[any, any])(nil)

type taskResponderContractWrapper[Input any, Output any] struct {
	taskManagerAbi *abi.ABI
	contract       taskManagerAbiContract[Input, Output]
	txMgr          txmgr.TxManager
}

func (tr taskResponderContractWrapper[Input, Output]) ProcessTaskResponse(taskResponse sdktypes.GenericOutputTaskResponse[Output]) (sdktypes.Bytes32, error) {
	abiType, err := internalutils.ExtractTypeFromAbi(tr.taskManagerAbi)
	if err != nil {
		return [32]byte{}, err
	}
	hashFn := internalutils.GetDefaultHashFunction(abiType)

	return hashFn(taskResponse)
}

// Creates a TaskResponder from an address and ABI.
// Returns an error in case the ABI is not compatible.
func NewTaskResponderFromAbi[Input any, Output any](address common.Address, abi *abi.ABI, txMgr txmgr.TxManager, httpClient bind.ContractBackend) (TaskResponder[Input, Output], error) {
	boundContract := bind.NewBoundContract(address, *abi, httpClient, httpClient, httpClient)
	// TODO: check if the ABI is compatible
	contract := taskManagerAbiContract[Input, Output]{boundContract}
	return NewTaskResponderFromContract(abi, contract, txMgr), nil
}

// Creates a TaskResponder from a contract implementing the given interface.
func NewTaskResponderFromContract[Input any, Output any](taskManagerAbi *abi.ABI, contract taskManagerAbiContract[Input, Output], txMgr txmgr.TxManager) TaskResponder[Input, Output] {
	return &taskResponderContractWrapper[Input, Output]{taskManagerAbi, contract, txMgr}
}

func (senderWrapper *taskResponderContractWrapper[Input, Output]) RespondToTask(
	task sdktypes.GenericInputTask[Input],
	taskResponse sdktypes.GenericOutputTaskResponse[Output],
	nonSignersStakesAndSig sdktypes.NonSignerStakesAndSignature,
) error {
	txOpts, err := senderWrapper.txMgr.GetNoSendTxOpts()
	if err != nil {
		return utils.WrapError("Error getting tx opts", err)
	}

	inputFieldName := capitalizeFieldName(senderWrapper.taskManagerAbi.Methods["respondToTask"].Inputs[0].Type.TupleRawNames[0])
	outputFieldName := capitalizeFieldName(senderWrapper.taskManagerAbi.Methods["respondToTask"].Inputs[1].Type.TupleRawNames[1])

	newTaskStruct := copyStructAndChangeFieldName(task, "InputValue", inputFieldName)
	newTaskResponseStruct := copyStructAndChangeFieldName(taskResponse, "OutputValue", outputFieldName)

	tx, err := senderWrapper.contract.RespondToTask(txOpts, newTaskStruct, newTaskResponseStruct, nonSignersStakesAndSig)
	if err != nil {
		return utils.WrapError("Error assembling CreateNewTask tx", err)
	}

	receipt, err := senderWrapper.txMgr.Send(context.Background(), tx, true)
	if err != nil {
		return utils.WrapError("Error submitting CreateNewTask tx", err)
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		return utils.WrapError("CreateNewTask tx failed", nil)
	}

	return nil
}

func capitalizeFieldName(name string) string {
	return strings.ToUpper(name[:1]) + name[1:]
}

func copyStructAndChangeFieldName(originalStruct any, previousName string, newName string) any {
	val := reflect.ValueOf(originalStruct)
	typ := reflect.TypeOf(originalStruct)

	var newFields []reflect.StructField
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		if field.Name == previousName {
			field.Name = newName
		}
		newFields = append(newFields, field)
	}

	newStructType := reflect.StructOf(newFields)
	newStruct := reflect.New(newStructType).Elem()

	for i := 0; i < newStruct.NumField(); i++ {
		originalField2 := val.Field(i)
		newStruct.Field(i).Set(originalField2)
	}

	return newStruct.Interface()
}
