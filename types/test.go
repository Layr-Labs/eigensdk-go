package types

import "github.com/Layr-Labs/eigensdk-go/crypto/bls"

type TestOperator struct {
	OperatorId     OperatorId
	StakePerQuorum map[QuorumNum]StakeAmount
	BlsKeypair     *bls.KeyPair
}
