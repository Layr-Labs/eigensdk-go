package kms

import (
	"context"

	"github.com/Layr-Labs/eigensdk-go/utils"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/kms"
)

func NewKMSClient(ctx context.Context, region string) (*kms.Client, error) {
	config, err := config.LoadDefaultConfig(ctx, config.WithRegion(region))
	if err != nil {
		return nil, utils.WrapError("failed to load AWS config", err)
	}

	c := kms.NewFromConfig(config)
	return c, nil
}
