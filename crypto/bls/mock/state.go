package mock

import (
	"context"
	"fmt"
	"math/big"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
)

type ChainDataMock struct {
	KeyPairs     []*bls.KeyPair
	NumOperators bls.OperatorIndex
}

var _ bls.ChainState = (*ChainDataMock)(nil)
var _ bls.IndexedChainState = (*ChainDataMock)(nil)

type PrivateOperatorInfo struct {
	*bls.OperatorInfo
	*bls.IndexedOperatorInfo
	KeyPair *bls.KeyPair
	Host    string
	Port    string
}

type PrivateOperatorState struct {
	*bls.OperatorState
	*bls.IndexedOperatorState
	PrivateOperators map[bls.OperatorId]*PrivateOperatorInfo
}

func makeOperatorId(id int) bls.OperatorId {
	data := [32]byte{}
	copy(data[:], []byte(fmt.Sprintf("%d", id)))
	return data
}

func NewChainDataMock(numOperators bls.OperatorIndex) (*ChainDataMock, error) {

	keyPairs := make([]*bls.KeyPair, numOperators)
	for ind := bls.OperatorIndex(0); ind < numOperators; ind++ {
		keyPair, err := bls.GenRandomBlsKeys()
		if err != nil {
			return nil, err
		}
		keyPairs[ind] = keyPair
	}

	return &ChainDataMock{
		NumOperators: numOperators,
		KeyPairs:     keyPairs,
	}, nil
}

func (d *ChainDataMock) GetTotalOperatorState(blockNumber uint) *PrivateOperatorState {

	indexedOperators := make(map[bls.OperatorId]*bls.IndexedOperatorInfo, d.NumOperators)
	storedOperators := make(map[bls.OperatorId]*bls.OperatorInfo)
	privateOperators := make(map[bls.OperatorId]*PrivateOperatorInfo, d.NumOperators)

	var aggPubKey *bls.G1Point

	quorumStake := 0

	for ind := bls.OperatorIndex(0); ind < d.NumOperators; ind++ {

		if ind == 0 {
			key := d.KeyPairs[ind].GetPubKeyG1()
			aggPubKey = key.Deserialize(key.Serialize())
		} else {
			aggPubKey.Add(d.KeyPairs[ind].GetPubKeyG1())
		}

		stake := ind + 1

		host := "0.0.0.0"
		port := fmt.Sprintf("3%03v", int(ind))

		stored := &bls.OperatorInfo{
			Stake: big.NewInt(int64(stake)),
			Index: ind,
		}

		indexed := &bls.IndexedOperatorInfo{
			Socket:   fmt.Sprintf("%v:%v", host, port),
			PubkeyG1: d.KeyPairs[ind].GetPubKeyG1(),
			PubkeyG2: d.KeyPairs[ind].GetPubKeyG2(),
		}

		private := &PrivateOperatorInfo{
			OperatorInfo:        stored,
			IndexedOperatorInfo: indexed,
			KeyPair:             d.KeyPairs[ind],
			Host:                host,
			Port:                port,
		}

		id := makeOperatorId(int(ind))
		storedOperators[id] = stored
		indexedOperators[id] = indexed
		privateOperators[id] = private

		quorumStake += int(stake)

	}

	totals := map[bls.QuorumID]*bls.OperatorInfo{
		0: {
			Stake: big.NewInt(int64(quorumStake)),
			Index: d.NumOperators,
		},
	}

	operatorState := &bls.OperatorState{
		Operators: map[bls.QuorumID]map[bls.OperatorId]*bls.OperatorInfo{
			0: storedOperators,
		},
		Totals:      totals,
		BlockNumber: blockNumber,
	}

	indexedState := &bls.IndexedOperatorState{
		OperatorState:    operatorState,
		IndexedOperators: indexedOperators,
		AggKeys: map[bls.QuorumID]*bls.G1Point{
			0: aggPubKey,
		},
	}

	privateOperatorState := &PrivateOperatorState{
		OperatorState:        operatorState,
		IndexedOperatorState: indexedState,
		PrivateOperators:     privateOperators,
	}

	return privateOperatorState

}

func (d *ChainDataMock) GetOperatorState(blockNumber uint, quorums []bls.QuorumID) (*bls.OperatorState, error) {

	state := d.GetTotalOperatorState(blockNumber)

	return state.OperatorState, nil

}

func (d *ChainDataMock) GetOperatorStateByOperator(
	blockNumber uint,
	operator bls.OperatorId,
) (*bls.OperatorState, error) {

	state := d.GetTotalOperatorState(blockNumber)

	return state.OperatorState, nil

}

func (d *ChainDataMock) GetIndexedOperatorState(
	blockNumber uint,
	quorums []bls.QuorumID,
) (*bls.IndexedOperatorState, error) {

	state := d.GetTotalOperatorState(blockNumber)

	return state.IndexedOperatorState, nil

}

func (d *ChainDataMock) GetCurrentBlockNumber() (uint, error) {
	return 0, nil
}

func (d *ChainDataMock) Start(context.Context) error {
	return nil
}
