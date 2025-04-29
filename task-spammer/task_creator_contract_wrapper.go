package taskspammer

import (
	"context"

	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	"github.com/Layr-Labs/eigensdk-go/utils"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
)

type TaskManagerTaskContract[Input any] interface {
	CreateNewTask(opts *bind.TransactOpts, input Input, quorumThresholdPercentage uint32, quorumNumbers []byte) (*gethtypes.Transaction, error)
}

type taskCreatorContractWrapper[Input any] struct {
	contract TaskManagerTaskContract[Input]
	txMgr    txmgr.TxManager
}

type taskManagerAbiContract[Input any] struct {
	contract *bind.BoundContract
}

func (tm taskManagerAbiContract[Input]) CreateNewTask(opts *bind.TransactOpts, input Input, quorumThresholdPercentage uint32, quorumNumbers []byte) (*gethtypes.Transaction, error) {
	return tm.contract.Transact(opts, "createNewTask", input, quorumThresholdPercentage, quorumNumbers)
}

func NewTaskCreatorFromAbi[Input any](address common.Address, abi abi.ABI, txMgr txmgr.TxManager) TaskCreator[Input] {
	boundContract := bind.NewBoundContract(address, abi, nil, nil, nil)
	contract := taskManagerAbiContract[Input]{boundContract}
	return &taskCreatorContractWrapper[Input]{contract, txMgr}
}

func NewTaskCreatorFromContract[Input any](contract TaskManagerTaskContract[Input], txMgr txmgr.TxManager) TaskCreator[Input] {
	return &taskCreatorContractWrapper[Input]{contract, txMgr}
}

func (senderWrapper *taskCreatorContractWrapper[Input]) CreateNewTask(ctx context.Context, input Input) error {
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
