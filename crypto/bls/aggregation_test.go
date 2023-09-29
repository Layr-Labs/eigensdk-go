package bls_test

import (
	"errors"
	"log"
	"strconv"
	"testing"

	loggingmock "github.com/Layr-Labs/eigensdk-go/logging/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls/mock"
)

func TestSignatureAggregation(t *testing.T) {
	t.Parallel()
	const NUM_OPERATORS = 10

	type fields struct {
		logger *loggingmock.MockLogger
	}

	dat, err := mock.NewChainDataMock(NUM_OPERATORS)
	if err != nil {
		log.Fatal(err)
	}

	var tests = map[string]struct {
		quorums  []bls.QuorumParam
		message  [32]byte
		advCount uint
		success  bool
		prepare  func(f *fields)
	}{
		"Succeeds when all operators sign at quorum threshold 100": {
			quorums: []bls.QuorumParam{
				{
					QuorumID:        0,
					QuorumThreshold: 100,
				},
			},
			message:  [32]byte{1, 2, 3, 4, 5, 6},
			advCount: 0,
			success:  true,
		},
		"Succeeds when 9 out of 10 operators sign at quorum threshold 80": {
			quorums: []bls.QuorumParam{
				{
					QuorumID:        0,
					QuorumThreshold: 80,
				},
			},
			message:  [32]byte{1, 2, 3, 4, 5, 6},
			advCount: 1,
			prepare: func(f *fields) {
				// 1 node operators are adversarial so we expect 1 such warning
				f.logger.EXPECT().
					Warn("Error returned from messageChan", "err", errors.New("adversary"))
			},
			success: true,
		},
		"Fails when 8 out of 10 operators sign at quorum threshold 90": {
			quorums: []bls.QuorumParam{
				{
					QuorumID:        0,
					QuorumThreshold: 90,
				},
			},
			message:  [32]byte{1, 2, 3, 4, 5, 6},
			advCount: 2,
			prepare: func(f *fields) {
				// 2 node operators are adversarial so we expect 2 such warnings
				f.logger.EXPECT().
					Warn("Error returned from messageChan", "err", errors.New("adversary")).
					Times(2)
			},
			success: false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			f := fields{
				logger: loggingmock.NewMockLogger(mockCtrl),
			}
			if test.prepare != nil {
				test.prepare(&f)
			}

			agg := bls.NewStdSignatureAggregator(f.logger)
			state := dat.GetTotalOperatorState(0)
			update := make(chan bls.SignerMessage)

			go simulateOperators(*state, test.message, update, test.advCount)

			_, err := agg.AggregateSignatures(state.IndexedOperatorState, test.quorums, test.message, update)
			if test.success {
				assert.Nil(t, err)
			} else {
				assert.NotNil(t, err)
			}
		})
	}
}

func makeOperatorId(id int) bls.OperatorId {
	data := [32]byte{}
	copy(data[:], []byte(strconv.FormatInt(int64(id), 10)))
	return data
}

func simulateOperators(
	state mock.PrivateOperatorState,
	message [32]byte,
	update chan bls.SignerMessage,
	advCount uint,
) {

	count := 0

	// Simulate the operators signing the message.
	// In real life, the ordering will be random, but we simulate the signing in a fixed order
	// to simulate stakes deterministically
	for i := 0; i < len(state.PrivateOperators); i++ {
		id := makeOperatorId(i)
		op := state.PrivateOperators[id]
		sig := op.KeyPair.SignMessage(message)
		if count < len(state.IndexedOperators)-int(advCount) {
			update <- bls.SignerMessage{
				Signature: sig,
				Operator:  id,
				Err:       nil,
			}
		} else {
			update <- bls.SignerMessage{
				Signature: nil,
				Operator:  id,
				Err:       errors.New("adversary"),
			}
		}

		count += 1
	}
}
