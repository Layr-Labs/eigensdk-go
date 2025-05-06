package challenger

import (
	"math/big"

	"github.com/Layr-Labs/eigensdk-go/challenger"
	"github.com/Layr-Labs/eigensdk-go/logging"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
	"github.com/ethereum/go-ethereum/common"
)

type DotProductInput struct {
	X []*big.Int
	Y []*big.Int
}

type DotProductOutput struct {
	Result *big.Int
}

type ChallengeVerifier struct {
	logger logging.Logger
}

func NewChallengeVerifier(logger logging.Logger, taskManagerAddr common.Address) (ChallengeVerifier, error) {
	// Get the task manager contract with the given address
	return ChallengeVerifier{
		logger: logger,
	}, nil
}

var _ challenger.ChallengeVerifier[DotProductInput, DotProductOutput] = (*ChallengeVerifier)(nil)

func (cv ChallengeVerifier) VerifyChallenge(taskIndex uint32, task sdktypes.GenericInputTask[DotProductInput], taskResponse sdktypes.TaskResponseData[DotProductOutput]) error {
	// Calculate response
	totalSum := big.NewInt(0)
	for i := range task.InputValue.X {
		currentSum := big.NewInt(0).Mul(task.InputValue.X[i], task.InputValue.Y[i])
		totalSum.Add(totalSum, currentSum)
	}

	// Compare submitted response with calculated here
	receivedResponse := taskResponse.TaskResponse.OutputValue.Result
	shouldRaiseChallenge := totalSum == receivedResponse

	if shouldRaiseChallenge {
		cv.logger.Infof("Response was not correct, expected %v and got %v", totalSum, receivedResponse)
		// Call the Raise Challenge method of the on-chain contract

	}

	return nil
}
