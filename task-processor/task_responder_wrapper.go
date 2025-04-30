package taskprocessor

import (
	"context"

	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
	"github.com/Layr-Labs/eigensdk-go/utils"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

var _ TaskResponder[any, any] = (*taskResponderContractWrapper[any, any])(nil)

type taskResponderContractWrapper[Input any, Output any] struct {
	contract taskManagerAbiContract[Input, Output]
	txMgr    txmgr.TxManager
}

type taskManagerAbiContract[Input any, Output any] struct {
	contract *bind.BoundContract
}

func (tm taskManagerAbiContract[Input, Output]) RespondToTask(opts *bind.TransactOpts, task sdktypes.GenericInputTask[Input], taskResponse sdktypes.GenericOutputTaskResponse[Output], nonSignerStakesAndSignature sdktypes.NonSignerStakesAndSignature) (*types.Transaction, error) {
	return tm.contract.Transact(opts, "respondToTask", task, taskResponse, nonSignerStakesAndSignature)
}

// Creates a TaskResponder from an address and ABI.
// Returns an error in case the ABI is not compatible.
func NewTaskResponderFromAbi[Input any, Output any](address common.Address, abi abi.ABI, txMgr txmgr.TxManager, httpClient bind.ContractBackend) (TaskResponder[Input, Output], error) {
	boundContract := bind.NewBoundContract(address, abi, httpClient, httpClient, httpClient)
	// TODO: check if the ABI is compatible
	contract := taskManagerAbiContract[Input, Output]{boundContract}
	return NewTaskResponderFromContract(contract, txMgr), nil
}

// Creates a TaskResponder from a contract implementing the given interface.
func NewTaskResponderFromContract[Input any, Output any](contract taskManagerAbiContract[Input, Output], txMgr txmgr.TxManager) TaskResponder[Input, Output] {
	return &taskResponderContractWrapper[Input, Output]{contract, txMgr}
}

func (senderWrapper *taskResponderContractWrapper[Input, Output]) RespondToTask(
	task sdktypes.GenericInputTask[Input], 
	taskResponse sdktypes.GenericOutputTaskResponse[Output], 
	nonSignersStakesAndSig sdktypes.NonSignerStakesAndSignature,
) error {
	println("Respond to task del senderWrapper 0")
	txOpts, err := senderWrapper.txMgr.GetNoSendTxOpts()
	if err != nil {
		return utils.WrapError("Error getting tx opts", err)
	}

	println("Respond to task del senderWrapper 1")

	tx, err := senderWrapper.contract.RespondToTask(txOpts, task, taskResponse, nonSignersStakesAndSig)
	if err != nil {
		println("Respond to task del senderWrapper 1.5, err es ", err.Error())
		return utils.WrapError("Error assembling CreateNewTask tx", err)
	}

	println("Respond to task del senderWrapper 2")

	receipt, err := senderWrapper.txMgr.Send(context.Background(), tx, true)
	if err != nil {
		return utils.WrapError("Error submitting CreateNewTask tx", err)
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		return utils.WrapError("CreateNewTask tx failed", nil)
	}

	println("Respond to task del senderWrapper 3")
	
	return nil
}

