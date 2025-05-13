package taskprocessor

import (
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
)

type TaskManagerContract[Input any, Output any] interface {
	CreateNewTask(opts *bind.TransactOpts, input Input, quorumThresholdPercentage uint32, quorumNumbers []byte) (*gethtypes.Transaction, error)

	RespondToTask(task sdktypes.GenericInputTask[Input], taskResponse sdktypes.GenericOutputTaskResponse[Output], nonSignersStakesAndSig sdktypes.NonSignerStakesAndSignature) error
	ProcessTaskResponse(taskResponse sdktypes.GenericOutputTaskResponse[Output]) (sdktypes.TaskResponseDigest, error)

	RaiseChallenge(task sdktypes.GenericInputTask[Input], taskResponse sdktypes.GenericOutputTaskResponse[Output], TaskResponseMetadata sdktypes.GenericTaskResponseMetadata, NonSigningOperatorPubKeys []sdktypes.BN254G1Point) error
}
