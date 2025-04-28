package taskspammer

import (
	"time"

	"github.com/Layr-Labs/eigensdk-go/logging"
)

// Task spammer configuration struct.
//
// Contains optional parameters for the task spammer.
// TODO: have default values
type Config struct {
	Logger           logging.Logger
	TimeBetweenTasks time.Duration

	QuorumThresholdPercentage uint32
	QuorumNumbers             []uint8
}
