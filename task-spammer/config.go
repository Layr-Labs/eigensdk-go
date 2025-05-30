package taskspammer

import (
	"time"
)

// Task spammer configuration struct.
//
// Contains optional parameters for the task spammer.
// TODO: have default values
type Config struct {
	// The duration of the period between tasks
	TimeBetweenTasks time.Duration `toml:"time_between_tasks"`

	// The percentage of the total quorum stake needed by the signers to make the aggregated response valid
	QuorumThresholdPercentage uint32 `toml:"quorum_threshold_percentage"`

	// The numbers of the quorums required to respond to tasks for the response to be valid
	QuorumNumbers []uint8 `toml:"quorum_number"`
}
