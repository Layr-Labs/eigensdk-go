package challenger

import (
	"context"
	"reflect"
	"strings"

	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
	"github.com/Layr-Labs/eigensdk-go/utils"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

var _ ChallengerRaiser[any, any] = (*challengerRaiserContractWrapper[any, any])(nil)

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
func NewChallengerRaiserFromAbi[Input any, Output any](address common.Address, abi *abi.ABI, txMgr txmgr.TxManager, httpClient bind.ContractBackend) (ChallengerRaiser[Input, Output], error) {
	boundContract := bind.NewBoundContract(address, *abi, httpClient, httpClient, httpClient)
	// TODO: check if the ABI is compatible
	contract := taskManagerAbiContract[Input, Output]{boundContract}
	return NewChallengerRaiserFromContract(abi, contract, txMgr), nil
}

// Creates a ChallengerRaiser from a contract implementing the given interface.
func NewChallengerRaiserFromContract[Input any, Output any](taskManagerAbi *abi.ABI, contract taskManagerAbiContract[Input, Output], txMgr txmgr.TxManager) ChallengerRaiser[Input, Output] {
	return &challengerRaiserContractWrapper[Input, Output]{taskManagerAbi, contract, txMgr}
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

	inputFieldName := capitalizeFieldName(challengerRaiser.taskManagerAbi.Methods["respondToTask"].Inputs[0].Type.TupleRawNames[0])
	outputFieldName := capitalizeFieldName(challengerRaiser.taskManagerAbi.Methods["respondToTask"].Inputs[1].Type.TupleRawNames[1])

	newTaskStruct := copyStructAndChangeFieldName(task, "InputValue", inputFieldName)
	newTaskResponseStruct := copyStructAndChangeFieldName(taskResponse, "OutputValue", outputFieldName)

	tx, err := challengerRaiser.contract.RaiseChallenge(txOpts, newTaskStruct, newTaskResponseStruct, TaskResponseMetadata, NonSigningOperatorPubKeys)
	if err != nil {
		return utils.WrapError("Error assembling RaiseChallenge tx", err)
	}

	receipt, err := challengerRaiser.txMgr.Send(context.Background(), tx, true)
	if err != nil {
		return utils.WrapError("Error submitting RaiseChallenge tx", err)
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		return utils.WrapError("RaiseChallenge tx failed", nil)
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
