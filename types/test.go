package types

import (
	"crypto/ecdsa"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
)

type TestBlsOperator struct {
	OperatorId     BlsOperatorId
	StakePerQuorum map[QuorumNum]StakeAmount
	BlsKeypair     *bls.KeyPair
}
type TestEcdsaOperator struct {
	OperatorId     EcdsaOperatorId
	StakePerQuorum map[QuorumNum]StakeAmount
	EcdsaPrivKey   *ecdsa.PrivateKey
}
