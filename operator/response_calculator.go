package operator

import (
	"fmt"
	"math/rand/v2"
)

// Interface used to compute the response of a task.
type ResponseCalculator[Input any, Output any] interface {
	// Computes the response of a task given its index and input.
	// It may return an error if the computation fails.
	ComputeResponse(taskIndex uint32, input Input) (Output, error)
}

type responseCalculatorFunction[Input any, Output any] struct {
	computeFn func(taskIndex uint32, input Input) (Output, error)
}

// Turns a function that receives a task index and an input value and returns an output value into a ResponseCalculator.
func NewResponseCalculatorFunction[Input any, Output any](
	computeFn func(taskIndex uint32, input Input) (Output, error),
) ResponseCalculator[Input, Output] {
	return responseCalculatorFunction[Input, Output]{computeFn: computeFn}
}

func (r responseCalculatorFunction[Input, Output]) ComputeResponse(
	taskIndex uint32, input Input,
) (Output, error) {
	return r.computeFn(taskIndex, input)
}

// Takes a ResponseCalculator and a value that is always an incorrect response.
// It returns a new ResponseCalculator that will compute a response using the received response calculator,
// but swap the response for the incorrect value with a certain probability.
//
// The function will return an error if the failure rate percentage is over 100.
func ComputeWithFailures[Input any, Output any](
	responseCalculator ResponseCalculator[Input, Output],
	incorrectValue Output,
	failureRatePercentage uint32,
) (ResponseCalculator[Input, Output], error) {
	if failureRatePercentage > 100 {
		return nil, fmt.Errorf("failure rate is over 100, should be a number between 0 and 100")
	}

	return NewResponseCalculatorFunction(func(taskIndex uint32, input Input) (Output, error) {
		correctResult, err := responseCalculator.ComputeResponse(taskIndex, input)
		if rand.Uint32()%100 < failureRatePercentage {
			return incorrectValue, err
		} else {
			return correctResult, err
		}
	}), nil
}
