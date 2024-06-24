package slashing

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"

	"github.com/stretchr/testify/assert"
)

func TestCalculateNewState(t *testing.T) {

	var tests = []struct {
		name        string
		oldState    State
		stakeSource StakeSource
		operatorSet OperatorSet
		bips        float64
		newState    *State
		debug       bool
		expectErr   bool
	}{
		{
			name: "Simple case where stake is sourced from slashable stake. A new operator set is added",
			oldState: State{
				TotalMagnitude: 10,
				OperatorSets: []OperatorSet{
					{
						AvsAddress:    common.HexToAddress("0xabc"),
						OperatorSetId: 1,
					},
					{
						AvsAddress:    common.HexToAddress("0xabc"),
						OperatorSetId: 2,
					},
				},
				SlashableMagnitudes: []uint64{1, 1},
			},
			stakeSource: StakeSourceSlashable,
			operatorSet: OperatorSet{
				AvsAddress:    common.HexToAddress("0xabc"),
				OperatorSetId: 4,
			},
			bips: 1000,
			newState: &State{
				TotalMagnitude: 20.0,
				OperatorSets: []OperatorSet{
					{
						AvsAddress:    common.HexToAddress("0xabc"),
						OperatorSetId: 4,
					},
				},
				SlashableMagnitudes: []uint64{2},
			},
		},
		{
			name: "Simple case where stake is sourced from slashable stake. Existing operator set is updated",
			oldState: State{
				TotalMagnitude: 10,
				OperatorSets: []OperatorSet{
					{
						AvsAddress:    common.HexToAddress("0xabc"),
						OperatorSetId: 1,
					},
					{
						AvsAddress:    common.HexToAddress("0xabc"),
						OperatorSetId: 2,
					},
				},
				SlashableMagnitudes: []uint64{1, 1},
			},
			stakeSource: StakeSourceSlashable,
			operatorSet: OperatorSet{
				AvsAddress:    common.HexToAddress("0xabc"),
				OperatorSetId: 1,
			},
			bips: 1500,
			newState: &State{
				TotalMagnitude: 40,
				OperatorSets: []OperatorSet{
					{
						AvsAddress:    common.HexToAddress("0xabc"),
						OperatorSetId: 1,
					},
				},
				SlashableMagnitudes: []uint64{6},
			},
		},
		{
			name: "New slashable stake is equal to available slashable stake so it should fail",
			oldState: State{
				TotalMagnitude: 10,
				OperatorSets: []OperatorSet{
					{
						AvsAddress:    common.HexToAddress("0xabc"),
						OperatorSetId: 1,
					},
					{
						AvsAddress:    common.HexToAddress("0xabc"),
						OperatorSetId: 2,
					},
				},
				SlashableMagnitudes: []uint64{1, 1},
			},
			stakeSource: StakeSourceSlashable,
			operatorSet: OperatorSet{
				AvsAddress:    common.HexToAddress("0xabc"),
				OperatorSetId: 4,
			},
			bips:      2000,
			expectErr: true,
		},
		//{
		//	name: "Simple case where stake is sourced from non slashable stake",
		//	oldState: State{
		//		TotalMagnitude:      10,
		//		OperatorSets:        []int{1, 2},
		//		SlashableMagnitudes: []float64{1, 1},
		//	},
		//	stakeSource: StakeSourceNonSlashable,
		//	operatorSet: 4,
		//	bips:        100,
		//	newState: &State{
		//		TotalMagnitude:      10.0,
		//		OperatorSets:        []int{4},
		//		SlashableMagnitudes: []float64{1.0},
		//	},
		//},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			newState, err := CalculateMagnitudes(tt.oldState, tt.operatorSet, tt.stakeSource, tt.bips, tt.debug)
			if tt.expectErr {
				assert.Error(t, err)

			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.newState, newState)
			}

		})

	}

}
