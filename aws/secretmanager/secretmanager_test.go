package secretmanager_test

import (
	"context"
	"testing"

	"github.com/Layr-Labs/eigensdk-go/aws/secretmanager"
)

func TestReadString(t *testing.T) {
	t.Skip("skipping test as it's meant for manual runs only")

	secretName := "SECRET_NAME"
	region := "REGION"
	value, err := secretmanager.ReadStringFromSecretManager(context.Background(), secretName, region)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Secret value: %s", value)
}
