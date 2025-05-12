package taskprocessor

import (
	"context"
	"fmt"

	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
	"github.com/Layr-Labs/eigensdk-go/utils"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"golang.org/x/crypto/sha3"
)

var _ TaskResponder[any, any] = (*taskResponderContractWrapper[any, any])(nil)

type taskResponderContractWrapper[Input any, Output any] struct {
	taskManagerAbi *abi.ABI
	contract       taskManagerAbiContract[Input, Output]
	txMgr          txmgr.TxManager
}

func (tr taskResponderContractWrapper[Input, Output]) ProcessTaskResponse(taskResponse sdktypes.GenericOutputTaskResponse[Output]) (sdktypes.Bytes32, error) {
	abiType, err := extractTypeFromAbi(tr.taskManagerAbi)
	if err != nil {
		return [32]byte{}, err
	}
	hashFn := getDefaultHashFunction(abiType)

	return hashFn(taskResponse)
}

func extractTypeFromAbi(taskManagerAbi *abi.ABI) (abi.Type, error) {
	taskResponseType, err := abi.NewType("tuple", "", []abi.ArgumentMarshaling{
		{
			Name: "referenceTaskIndex",
			Type: "uint32",
		},
		{
			Name: "OutputValue", // Left because abi does not support purely anonymous or underscored fields
			Type: taskManagerAbi.Events["TaskResponded"].Inputs[0].Type.TupleElems[1].String(),
		},
	})
	if err != nil {
		return abi.Type{}, fmt.Errorf("error creating abi task response type: %w", err)
	}

	return taskResponseType, nil
}

func getDefaultHashFunction(taskResponseType abi.Type) sdktypes.TaskResponseHashFunction {
	return func(taskResponse sdktypes.TaskResponse) (sdktypes.TaskResponseDigest, error) {
		arguments := abi.Arguments{
			{
				Type: taskResponseType,
			},
		}

		encodeTaskResponseByte, err := arguments.Pack(taskResponse)
		if err != nil {
			return sdktypes.Bytes32{}, fmt.Errorf("error encoding task response: %w", err)
		}

		var taskResponseDigest [32]byte
		hasher := sha3.NewLegacyKeccak256()
		hasher.Write(encodeTaskResponseByte)
		copy(taskResponseDigest[:], hasher.Sum(nil)[:32])

		return taskResponseDigest, nil
	}
}

type taskManagerAbiContract[Input any, Output any] struct {
	contract *bind.BoundContract
}

func (tm taskManagerAbiContract[Input, Output]) RespondToTask(opts *bind.TransactOpts, task any, taskResponse any, nonSignerStakesAndSignature any) (*types.Transaction, error) {
	return tm.contract.Transact(opts, "respondToTask", task, taskResponse, nonSignerStakesAndSignature)
}

// Creates a TaskResponder from an address and ABI.
// Returns an error in case the ABI is not compatible.
func NewTaskResponderFromAbi[Input any, Output any](address common.Address, abi *abi.ABI, txMgr txmgr.TxManager, httpClient bind.ContractBackend) (TaskResponder[Input, Output], error) {
	boundContract := bind.NewBoundContract(address, *abi, httpClient, httpClient, httpClient)
	// TODO: check if the ABI is compatible
	contract := taskManagerAbiContract[Input, Output]{boundContract}
	return &taskResponderContractWrapper[Input, Output]{abi, contract, txMgr}, nil
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

	inputFieldName := utils.CapitalizeFieldName(senderWrapper.taskManagerAbi.Methods["respondToTask"].Inputs[0].Type.TupleRawNames[0])
	outputFieldName := utils.CapitalizeFieldName(senderWrapper.taskManagerAbi.Methods["respondToTask"].Inputs[1].Type.TupleRawNames[1])

	newTaskStruct := utils.CopyStructAndChangeFieldName(task, "InputValue", inputFieldName)
	newTaskResponseStruct := utils.CopyStructAndChangeFieldName(taskResponse, "OutputValue", outputFieldName)

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
