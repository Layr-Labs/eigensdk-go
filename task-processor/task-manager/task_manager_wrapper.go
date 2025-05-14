package taskmanager

import (
	"context"

	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	internalutils "github.com/Layr-Labs/eigensdk-go/internal/utils"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
	"github.com/Layr-Labs/eigensdk-go/utils"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
)

var _ TaskManager[any, any] = (*taskManagerContractWrapper[any, any])(nil)

type taskManagerContractWrapper[Input any, Output any] struct {
	taskManagerAbi *abi.ABI
	contract       taskManagerAbiContract[Input, Output]
	txMgr          txmgr.TxManager
	hashFunction   sdktypes.TaskResponseHashFunction
}

type taskManagerAbiContract[Input any, Output any] struct {
	contract *bind.BoundContract
}

func (tm taskManagerAbiContract[Input, Output]) RaiseChallenge(opts *bind.TransactOpts, task any, taskResponse any, taskResponseMetadata any, nonSigningOperatorPubKeys any) (*gethtypes.Transaction, error) {
	return tm.contract.Transact(opts, "raiseAndResolveChallenge", task, taskResponse, taskResponseMetadata, nonSigningOperatorPubKeys)
}

func (tm taskManagerAbiContract[Input, Output]) RespondToTask(opts *bind.TransactOpts, task any, taskResponse any, nonSignerStakesAndSignature any) (*gethtypes.Transaction, error) {
	return tm.contract.Transact(opts, "respondToTask", task, taskResponse, nonSignerStakesAndSignature)
}

func (tm taskManagerAbiContract[Input, Output]) CreateNewTask(opts *bind.TransactOpts, input Input, quorumThresholdPercentage uint32, quorumNumbers []byte) (*gethtypes.Transaction, error) {
	return tm.contract.Transact(opts, "createNewTask", input, quorumThresholdPercentage, quorumNumbers)
}

// Creates a taskManager wrapper from an address and ABI.
// Returns an error in case the ABI is not compatible.
func NewTaskManagerFromAbi[Input any, Output any](address common.Address, abi *abi.ABI, txMgr txmgr.TxManager, httpClient bind.ContractBackend) (TaskManager[Input, Output], error) {
	boundContract := bind.NewBoundContract(address, *abi, httpClient, httpClient, httpClient)

	abiType, err := internalutils.ExtractTypeFromAbi(abi)
	if err != nil {
		return nil, err
	}
	hashFn := internalutils.GetDefaultHashFunction(abiType)

	// TODO: check if the ABI is compatible
	contract := taskManagerAbiContract[Input, Output]{boundContract}
	return &taskManagerContractWrapper[Input, Output]{abi, contract, txMgr, hashFn}, nil
}

func (senderWrapper *taskManagerContractWrapper[Input, Output]) RaiseChallenge(
	task sdktypes.GenericInputTask[Input],
	taskResponse sdktypes.GenericOutputTaskResponse[Output],
	TaskResponseMetadata sdktypes.GenericTaskResponseMetadata,
	NonSigningOperatorPubKeys []sdktypes.BN254G1Point,
) error {
	txOpts, err := senderWrapper.txMgr.GetNoSendTxOpts()
	if err != nil {
		return utils.WrapError("Error getting tx opts", err)
	}

	inputFieldName := internalutils.CapitalizeFieldName(senderWrapper.taskManagerAbi.Methods["respondToTask"].Inputs[0].Type.TupleRawNames[0])
	outputFieldName := internalutils.CapitalizeFieldName(senderWrapper.taskManagerAbi.Methods["respondToTask"].Inputs[1].Type.TupleRawNames[1])

	newTaskStruct := internalutils.CopyStructAndChangeFieldName(task, "InputValue", inputFieldName)
	newTaskResponseStruct := internalutils.CopyStructAndChangeFieldName(taskResponse, "OutputValue", outputFieldName)

	tx, err := senderWrapper.contract.RaiseChallenge(txOpts, newTaskStruct, newTaskResponseStruct, TaskResponseMetadata, NonSigningOperatorPubKeys)
	if err != nil {
		return utils.WrapError("Error assembling RaiseChallenge tx", err)
	}

	receipt, err := senderWrapper.txMgr.Send(context.Background(), tx, true)
	if err != nil {
		return utils.WrapError("Error submitting RaiseChallenge tx", err)
	}
	if receipt.Status != gethtypes.ReceiptStatusSuccessful {
		return utils.WrapError("RaiseChallenge tx failed", nil)
	}

	return nil
}

func (senderWrapper *taskManagerContractWrapper[Input, Output]) CreateNewTask(ctx context.Context, input Input) error {
	var quorumThresholdPercentage uint32 = 100
	var quorumNumbers []uint8 = []uint8{0}

	txOpts, err := senderWrapper.txMgr.GetNoSendTxOpts()
	if err != nil {
		return utils.WrapError("Error getting tx opts", err)
	}

	tx, err := senderWrapper.contract.CreateNewTask(txOpts, input, quorumThresholdPercentage, quorumNumbers)
	if err != nil {
		return utils.WrapError("Error assembling CreateNewTask tx", err)
	}

	receipt, err := senderWrapper.txMgr.Send(ctx, tx, true)
	if err != nil {
		return utils.WrapError("Error submitting CreateNewTask tx", err)
	}
	if receipt.Status != gethtypes.ReceiptStatusSuccessful {
		return utils.WrapError("CreateNewTask tx failed", nil)
	}
	return nil
}

func (senderWrapper *taskManagerContractWrapper[Input, Output]) RespondToTask(
	task sdktypes.GenericInputTask[Input],
	taskResponse sdktypes.GenericOutputTaskResponse[Output],
	nonSignersStakesAndSig sdktypes.NonSignerStakesAndSignature,
) error {
	txOpts, err := senderWrapper.txMgr.GetNoSendTxOpts()
	if err != nil {
		return utils.WrapError("Error getting tx opts", err)
	}

	inputFieldName := internalutils.CapitalizeFieldName(senderWrapper.taskManagerAbi.Methods["respondToTask"].Inputs[0].Type.TupleRawNames[0])
	outputFieldName := internalutils.CapitalizeFieldName(senderWrapper.taskManagerAbi.Methods["respondToTask"].Inputs[1].Type.TupleRawNames[1])

	newTaskStruct := internalutils.CopyStructAndChangeFieldName(task, "InputValue", inputFieldName)
	newTaskResponseStruct := internalutils.CopyStructAndChangeFieldName(taskResponse, "OutputValue", outputFieldName)

	tx, err := senderWrapper.contract.RespondToTask(txOpts, newTaskStruct, newTaskResponseStruct, nonSignersStakesAndSig)
	if err != nil {
		return utils.WrapError("Error assembling CreateNewTask tx", err)
	}

	receipt, err := senderWrapper.txMgr.Send(context.Background(), tx, true)
	if err != nil {
		return utils.WrapError("Error submitting CreateNewTask tx", err)
	}
	if receipt.Status != gethtypes.ReceiptStatusSuccessful {
		return utils.WrapError("CreateNewTask tx failed", nil)
	}

	return nil
}

func (tr taskManagerContractWrapper[Input, Output]) HashTaskResponse(taskResponse sdktypes.GenericOutputTaskResponse[Output]) (sdktypes.Bytes32, error) {
	return tr.hashFunction(taskResponse)
}
