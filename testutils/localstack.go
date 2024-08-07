package testutils

import (
	"context"
	"fmt"

	"github.com/Layr-Labs/eigensdk-go/aws"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

const LocalStackPort = "4566"

func StartLocalstackContainer(name string) (testcontainers.Container, error) {
	fmt.Println("Starting Localstack container")
	req := testcontainers.ContainerRequest{
		Image: "localstack/localstack:latest",
		Name:  fmt.Sprintf("localstack-test-%s", name),
		Env: map[string]string{
			"LOCALSTACK_HOST": fmt.Sprintf("localhost.localstack.cloud:%s", LocalStackPort),
		},
		ExposedPorts: []string{LocalStackPort},
		WaitingFor:   wait.ForLog("Ready."),
		AutoRemove:   true,
	}
	return testcontainers.GenericContainer(context.Background(), testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
}

func NewKMSClient(localStackPort string) (*kms.Client, error) {
	cfg, err := aws.GetAWSConfig(
		"localstack",
		"localstack",
		"us-east-1",
		fmt.Sprintf("http://127.0.0.1:%s", localStackPort),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to load AWS config: %w", err)
	}

	c := kms.NewFromConfig(*cfg)
	return c, nil
}

func CreateKMSKey(localStackPort string) (*types.KeyMetadata, error) {
	c, err := NewKMSClient(localStackPort)
	if err != nil {
		return nil, fmt.Errorf("failed to create KMS client: %w", err)
	}
	res, err := c.CreateKey(context.Background(), &kms.CreateKeyInput{
		KeySpec:  types.KeySpecEccSecgP256k1,
		KeyUsage: types.KeyUsageTypeSignVerify,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create key: %w", err)
	}
	return res.KeyMetadata, nil
}
