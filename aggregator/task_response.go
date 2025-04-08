package aggregator

import sdktypes "github.com/Layr-Labs/eigensdk-go/types"

type TaskResponse interface {
	TaskIndex() sdktypes.TaskIndex
	Digest() [256]byte
}
