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

var _ TaskManager[any, any, any] = (*taskManagerContractWrapper[any, any, any])(nil)

type taskManagerContractWrapper[Input any, Output any, Proof any] struct {
	taskManagerAbi  *abi.ABI
	contract        taskManagerAbiContract[Input, Output, Proof]
	txMgr           txmgr.TxManager
	hashFunction    sdktypes.TaskResponseHashFunction
	inputFieldName  string
	outputFieldName string
}

type taskManagerAbiContract[Input any, Output any, Proof any] struct {
	contract *bind.BoundContract
}

func (tm taskManagerAbiContract[Input, Output, Proof]) RaiseChallenge(opts *bind.TransactOpts, task any, taskResponse any, taskResponseMetadata any, nonSigningOperatorPubKeys any, proof any) (*gethtypes.Transaction, error) {
	if proof == nil {
		return tm.contract.Transact(opts, "raiseAndResolveChallenge", task, taskResponse, taskResponseMetadata, nonSigningOperatorPubKeys)
	} else {
		return tm.contract.Transact(opts, "raiseAndResolveChallenge", task, taskResponse, taskResponseMetadata, nonSigningOperatorPubKeys, proof)
	}
}

func (tm taskManagerAbiContract[Input, Output, Proof]) RespondToTask(opts *bind.TransactOpts, task any, taskResponse any, nonSignerStakesAndSignature any) (*gethtypes.Transaction, error) {
	return tm.contract.Transact(opts, "respondToTask", task, taskResponse, nonSignerStakesAndSignature)
}

func (tm taskManagerAbiContract[Input, Output, Proof]) CreateNewTask(opts *bind.TransactOpts, input Input, quorumThresholdPercentage uint32, quorumNumbers []byte) (*gethtypes.Transaction, error) {
	return tm.contract.Transact(opts, "createNewTask", input, quorumThresholdPercentage, quorumNumbers)
}

// Creates a taskManager wrapper from an address and ABI.
// Returns an error in case the ABI is not compatible.
func NewTaskManagerFromAbi[Input any, Output any, Proof any](address common.Address, abi *abi.ABI, txMgr txmgr.TxManager, httpClient bind.ContractBackend) (TaskManager[Input, Output, Proof], error) {
	boundContract := bind.NewBoundContract(address, *abi, httpClient, httpClient, httpClient)

	abiType, err := internalutils.ExtractTypeFromAbi(abi)
	if err != nil {
		return nil, err
	}
	hashFn := internalutils.GetDefaultHashFunction(abiType)

	inputFieldName := internalutils.CapitalizeFieldName(abi.Methods["respondToTask"].Inputs[0].Type.TupleRawNames[0])
	outputFieldName := internalutils.CapitalizeFieldName(abi.Methods["respondToTask"].Inputs[1].Type.TupleRawNames[1])

	// TODO: check if the ABI is compatible
	contract := taskManagerAbiContract[Input, Output, Proof]{boundContract}
	return &taskManagerContractWrapper[Input, Output, Proof]{abi, contract, txMgr, hashFn, inputFieldName, outputFieldName}, nil
}

func (senderWrapper *taskManagerContractWrapper[Input, Output, Proof]) RaiseChallenge(
	task Task[Input],
	taskResponse TaskResponse[Output],
	taskResponseMetadata sdktypes.TaskResponseMetadata,
	nonSigningOperatorPubKeys []sdktypes.BN254G1Point,
	proof Proof,
) error {
	txOpts, err := senderWrapper.txMgr.GetNoSendTxOpts()
	if err != nil {
		return utils.WrapError("Error getting tx opts", err)
	}

	newTaskStruct := internalutils.CopyStructAndChangeFieldName(task, "InputValue", senderWrapper.inputFieldName)
	newTaskResponseStruct := internalutils.CopyStructAndChangeFieldName(taskResponse, "OutputValue", senderWrapper.outputFieldName)

	tx, err := senderWrapper.contract.RaiseChallenge(txOpts, newTaskStruct, newTaskResponseStruct, taskResponseMetadata, nonSigningOperatorPubKeys, proof)
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

func (senderWrapper *taskManagerContractWrapper[Input, Output, Proof]) CreateNewTask(ctx context.Context, input Input, quorumThresholdPercentage uint32, quorumNumbers []uint8) error {
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

func (senderWrapper *taskManagerContractWrapper[Input, Output, Proof]) RespondToTask(
	task Task[Input],
	taskResponse TaskResponse[Output],
	nonSignersStakesAndSig sdktypes.NonSignerStakesAndSignature,
) error {
	txOpts, err := senderWrapper.txMgr.GetNoSendTxOpts()
	if err != nil {
		return utils.WrapError("Error getting tx opts", err)
	}

	newTaskStruct := internalutils.CopyStructAndChangeFieldName(task, "InputValue", senderWrapper.inputFieldName)
	newTaskResponseStruct := internalutils.CopyStructAndChangeFieldName(taskResponse, "OutputValue", senderWrapper.outputFieldName)

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

func (tr taskManagerContractWrapper[Input, Output, Proof]) HashTaskResponse(taskResponse TaskResponse[Output]) (sdktypes.Bytes32, error) {
	return tr.hashFunction(taskResponse)
}
