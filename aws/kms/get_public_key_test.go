package kms_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	eigenkms "github.com/Layr-Labs/eigensdk-go/aws/kms"
	"github.com/Layr-Labs/eigensdk-go/testutils"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
)

var (
	mappedLocalstackPort string
	keyMetadata          *types.KeyMetadata
	container            testcontainers.Container
)

func TestMain(m *testing.M) {
	err := setup()
	if err != nil {
		fmt.Println("Error setting up test environment:", err)
		teardown()
		os.Exit(1)
	}
	exitCode := m.Run()
	teardown()
	os.Exit(exitCode)
}

func setup() error {
	var err error
	container, err = testutils.StartLocalstackContainer("get_public_key_test")
	if err != nil {
		return err
	}
	mappedPort, err := container.MappedPort(context.Background(), testutils.LocalStackPort)
	if err != nil {
		return err
	}
	mappedLocalstackPort = string(mappedPort)
	keyMetadata, err = testutils.CreateKMSKey(mappedLocalstackPort)
	if err != nil {
		return err
	}
	return nil
}

func teardown() {
	_ = container.Terminate(context.Background())
}

func TestGetPublicKey(t *testing.T) {
	c, err := testutils.NewKMSClient(mappedLocalstackPort)
	assert.Nil(t, err)
	assert.NotNil(t, keyMetadata.KeyId)
	pk, err := eigenkms.GetPublicKey(context.Background(), c, *keyMetadata.KeyId)
	assert.Nil(t, err)
	assert.NotNil(t, pk)
	keyAddr := crypto.PubkeyToAddress(*pk)
	t.Logf("Public key address: %s", keyAddr.String())
	assert.NotEqual(t, keyAddr, common.Address{0})
}
