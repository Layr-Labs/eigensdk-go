package slashing

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"strings"
)

type StakeSource string

const (
	StakeSourceSlashable    StakeSource = "slashable"
	StakeSourceNonSlashable StakeSource = "non-slashable"
	StakeSourceTotal        StakeSource = "total"
)

type OperatorSet struct {
	AvsAddress    common.Address
	OperatorSetId int32
}

type State struct {
	TotalMagnitude      uint64
	OperatorSets        []OperatorSet
	SlashableMagnitudes []uint64
}

func (s *State) Print() {
	fmt.Println(strings.Repeat("-", 60))
	fmt.Println("Total Magnitude: ", s.TotalMagnitude)
	fmt.Println()
	for i, operatorSet := range s.OperatorSets {
		fmt.Println("Operator Set: ", i)
		fmt.Println("AVS Address: ", operatorSet.AvsAddress.Hex())
		fmt.Println("Operator Set ID: ", operatorSet.OperatorSetId)
		fmt.Println("Magnitude: ", s.SlashableMagnitudes[i])
		fmt.Println()
	}
	fmt.Println(strings.Repeat("-", 60))
}

func CalculateMagnitudes(
	currentState State,
	operatorSet OperatorSet,
	stakeSource StakeSource,
	bips float64,
	debug bool,
) (*State, error) {
	if debug {
		currentState.Print()
	}
	slashableProportion := bips / 10_000
	if stakeSource == StakeSourceSlashable {
		currentSlashableProportion := float64(sumUInt64Array(currentState.SlashableMagnitudes)) / float64(currentState.TotalMagnitude)
		if currentSlashableProportion <= slashableProportion {
			return nil, fmt.Errorf("slashable proportion to allocate %f is less than or equal to available slashable proportion %f", slashableProportion, currentSlashableProportion)
		}

		nonSlashableProportion := (float64(currentState.TotalMagnitude) - float64(sumUInt64Array(currentState.SlashableMagnitudes))) / float64(currentState.TotalMagnitude)
		if debug {
			fmt.Println("nonSlashableProportion: ", nonSlashableProportion)
		}

		allocatedMagnitude := float64(sumUInt64Array(currentState.SlashableMagnitudes))
		if debug {
			fmt.Println("allocatedMagnitude: ", allocatedMagnitude)
		}

		totalMagnitudeNew := allocatedMagnitude / (1 - slashableProportion - nonSlashableProportion)
		if debug {
			fmt.Println("totalMagnitudeNew: ", totalMagnitudeNew)
		}

		slashableMagnitude := slashableProportion * totalMagnitudeNew
		if debug {
			fmt.Println("slashableMagnitude: ", slashableMagnitude)
		}

		nonSlashableMagnitude := nonSlashableProportion * totalMagnitudeNew
		if debug {
			fmt.Println("nonSlashableMagnitude: ", nonSlashableMagnitude)
		}

		opSetToUpdate := []OperatorSet{operatorSet}
		slashableMagnitudeSet := []uint64{uint64(slashableMagnitude)}

		newState := &State{
			TotalMagnitude:      uint64(totalMagnitudeNew),
			OperatorSets:        opSetToUpdate,
			SlashableMagnitudes: slashableMagnitudeSet,
		}
		if debug {
			newState.Print()
		}
		return newState, nil
	} else {
		return nil, fmt.Errorf("unimplemented stake source %s", stakeSource)
	}
}

func sumUInt64Array(arr []uint64) uint64 {
	sum := uint64(0)
	for _, v := range arr {
		sum += v
	}
	return sum
}
