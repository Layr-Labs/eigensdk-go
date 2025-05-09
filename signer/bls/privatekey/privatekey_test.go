package privatekey

import (
	"context"
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

const privateKey = "13718104057011380243349384062412322292853638010146548074368241565852610884213"

func TestOperatorId(t *testing.T) {
	signer, err := New(Config{
		PrivateKey: privateKey,
	})
	assert.NoError(t, err)

	operatorId, err := signer.GetOperatorId()
	assert.NoError(t, err)
	assert.Equal(t, operatorId, "0x2f6d54c4ef253d4ac1b6bfa672a8f449a0aa6ab1c20ee97a74f8e521fe57604b")
}

func TestPublicKeyG1(t *testing.T) {
	signer, err := New(Config{
		PrivateKey: privateKey,
	})
	assert.NoError(t, err)

	pubKeyG1 := signer.GetPublicKeyG1()
	assert.Equal(t, pubKeyG1, "0xdc8f9427033e5ff392f5cc97cc3d6a5472cff2778eee0961a497bd7dbb629a36")
}

func TestPublicKeyG2(t *testing.T) {
	signer, err := New(Config{
		PrivateKey: privateKey,
	})
	assert.NoError(t, err)

	pubKeyG2 := signer.GetPublicKeyG2()
	assert.Equal(
		t,
		pubKeyG2,
		"0xe9b0f889a847f8dc2ed0514a6b7e11043679491052502e0a68ccc9a410f524e01e9d4863b49ca41d0a94928290b15aed25bfe097e266bdbb9106a09f689b4ea8",
	)
}

func TestSign(t *testing.T) {
	signer, err := New(Config{
		PrivateKey: privateKey,
	})
	assert.NoError(t, err)

	msg := [32]byte{}
	copy(msg[:], []byte("hello"))

	signature, err := signer.Sign(context.Background(), msg[:])
	signatureHex := hex.EncodeToString(signature)
	assert.NoError(t, err)
	assert.Equal(
		t,
		signatureHex,
		"2a21c157c51d13aa5922f67751520287088aa5336ae6d2ba06d0423fe7166ed811b48b5a3ff129d4c645f5f4cb33177a3503e26dc2d508b4b943bb8dd45085e2",
	)
}
