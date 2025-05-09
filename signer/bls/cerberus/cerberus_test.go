package cerberus

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO: add test for cerberus when the docker is released

func TestOperatorId(t *testing.T) {
	signer := Signer{
		pubKeyHex: "dc8f9427033e5ff392f5cc97cc3d6a5472cff2778eee0961a497bd7dbb629a36",
	}

	operatorId, err := signer.GetOperatorId()
	assert.NoError(t, err)
	assert.Equal(t, operatorId, "0x2f6d54c4ef253d4ac1b6bfa672a8f449a0aa6ab1c20ee97a74f8e521fe57604b")
}
