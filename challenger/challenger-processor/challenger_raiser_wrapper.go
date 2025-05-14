package challengerprocessor

import (
	"context"

	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	internalutils "github.com/Layr-Labs/eigensdk-go/internal/utils"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
	"github.com/Layr-Labs/eigensdk-go/utils"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

var _ ChallengeRaiser[any, any] = (*challengerRaiserContractWrapper[any, any])(nil)

type challengerRaiserContractWrapper[Input any, Output any] struct {
	taskManagerAbi *abi.ABI
	contract       taskManagerAbiContract[Input, Output]
	txMgr          txmgr.TxManager
}

type taskManagerAbiContract[Input any, Output any] struct {
	contract *bind.BoundContract
}

func (tm taskManagerAbiContract[Input, Output]) RaiseChallenge(opts *bind.TransactOpts, task any, taskResponse any, taskResponseMetadata any, nonSigningOperatorPubKeys any) (*types.Transaction, error) {
	return tm.contract.Transact(opts, "raiseAndResolveChallenge", task, taskResponse, taskResponseMetadata, nonSigningOperatorPubKeys)
}

// Creates a ChallengerRaiser from an address and ABI.
// Returns an error in case the ABI is not compatible.
func NewChallengeRaiserFromAbi[Input any, Output any](address common.Address, abi *abi.ABI, txMgr txmgr.TxManager, httpClient bind.ContractBackend) (ChallengeRaiser[Input, Output], error) {
	boundContract := bind.NewBoundContract(address, *abi, httpClient, httpClient, httpClient)
	// TODO: check if the ABI is compatible
	contract := taskManagerAbiContract[Input, Output]{boundContract}
	return &challengerRaiserContractWrapper[Input, Output]{abi, contract, txMgr}, nil
}

func (challengerRaiser *challengerRaiserContractWrapper[Input, Output]) RaiseChallenge(
	task sdktypes.GenericInputTask[Input],
	taskResponse sdktypes.GenericOutputTaskResponse[Output],
	TaskResponseMetadata sdktypes.GenericTaskResponseMetadata,
	NonSigningOperatorPubKeys []sdktypes.BN254G1Point,
) error {
	txOpts, err := challengerRaiser.txMgr.GetNoSendTxOpts()
	if err != nil {
		return utils.WrapError("Error getting tx opts", err)
	}

	inputFieldName := internalutils.CapitalizeFieldName(challengerRaiser.taskManagerAbi.Methods["respondToTask"].Inputs[0].Type.TupleRawNames[0])
	outputFieldName := internalutils.CapitalizeFieldName(challengerRaiser.taskManagerAbi.Methods["respondToTask"].Inputs[1].Type.TupleRawNames[1])

	newTaskStruct := internalutils.CopyStructAndChangeFieldName(task, "InputValue", inputFieldName)
	newTaskResponseStruct := internalutils.CopyStructAndChangeFieldName(taskResponse, "OutputValue", outputFieldName)

	tx, err := challengerRaiser.contract.RaiseChallenge(txOpts, newTaskStruct, newTaskResponseStruct, TaskResponseMetadata, NonSigningOperatorPubKeys)
	if err != nil {
		return utils.WrapError("Error assembling RaiseChallenge tx", err)
	}

	receipt, err := challengerRaiser.txMgr.Send(context.Background(), tx, true)
	if err != nil {
		return utils.WrapError("Error submitting RaiseChallenge tx", err)
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		return utils.WrapError("RaiseChallenge tx reverted", nil)
	}

	return nil
}
