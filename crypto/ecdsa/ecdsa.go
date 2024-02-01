package ecdsa

import (
	"crypto/ecdsa"

	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/ethereum/go-ethereum/crypto"
)

type PrivateKey = ecdsa.PrivateKey
type Signature = []byte

// SignMsg is a wrapper around geth's crypto.Sign function which adds 27 to the v value
// see https://github.com/ethereum/go-ethereum/issues/28757#issuecomment-1874525854
// and https://twitter.com/pcaversaccio/status/1671488928262529031
// Currently all of our contracts expect the v value to be 27 or 28, which is why we add it here.
// VerifySignature below thus subtracts 27 from the v value before verifying the signature.
// TODO(samlaf): if and when we figure out some avs contracts use the 0/1 v scheme, or the eip-155 scheme,
// then we should generalize this function to sign based on a scheme passed as argument.
func SignMsg(msg []byte, privateKey *PrivateKey) (Signature, error) {
	sig, err := crypto.Sign(msg, privateKey)
	if err != nil {
		return nil, err
	}
	// TODO(samlaf): if and when we figure out some avs contracts use the 0/1 v scheme, or the eip-155 scheme,
	// then we should generalize this function to sign based on a scheme passed as argument.
	sig[64] += 27
	return sig, nil
}

// VerifySignature is a wrapper around geth's crypto.SigToPub function which subtracts 27 from the v value
// before verifying the signature. See the SignMsg function above for more details.
func VerifySignature(msgHash []byte, sig Signature, operatorId types.EcdsaOperatorId) (bool, error) {
	// make a copy so as not to modify the underlying signature array
	sigCopy := make([]byte, len(sig))
	copy(sigCopy, sig)
	sigCopy[64] -= 27
	recoveredPubKey, err := crypto.SigToPub(msgHash, sigCopy)
	if err != nil {
		return true, err
	}
	return crypto.PubkeyToAddress(*recoveredPubKey) == operatorId, nil

}
