package local

import (
	"context"
	"encoding/hex"

	sdkBls "github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/signer/bls/types"
)

type Config struct {
	Path     string
	Password string
}

type Signer struct {
	key *sdkBls.KeyPair
}

func New(cfg Config) (*Signer, error) {
	keyPair, err := sdkBls.ReadPrivateKeyFromFile(cfg.Path, cfg.Password)
	if err != nil {
		return nil, err
	}
	return &Signer{
		key: keyPair,
	}, nil
}

func (s Signer) Sign(ctx context.Context, msg []byte) ([]byte, error) {
	if len(msg) != 32 {
		return nil, types.ErrInvalidMessageLength
	}

	var data [32]byte
	copy(data[:], msg)

	return s.key.SignMessage(data).Serialize(), nil
}

func (s Signer) SignG1(ctx context.Context, msg []byte) ([]byte, error) {
	if len(msg) != 64 {
		return nil, types.ErrInvalidMessageLength
	}

	msgG1 := new(sdkBls.G1Point)
	msgG1 = msgG1.Deserialize(msg)
	return s.key.SignHashedToCurveMessage(msgG1.G1Affine).Serialize(), nil
}

func (s Signer) GetOperatorId() (string, error) {
	return s.key.PubKey.GetOperatorID(), nil
}

func (s Signer) GetPublicKeyG1() string {
	rawBytes := s.key.PubKey.Bytes()
	return hex.EncodeToString(rawBytes[:])
}

func (s Signer) GetPublicKeyG2() string {
	rawBytes := s.key.GetPubKeyG2().Bytes()
	return hex.EncodeToString(rawBytes[:])
}
