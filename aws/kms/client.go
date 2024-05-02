package kms

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/kms"
)

func NewKMSClient(ctx context.Context, secretName, region string) (*kms.Client, error) {
	config, err := config.LoadDefaultConfig(ctx, config.WithRegion(region))
	if err != nil {
		return nil, fmt.Errorf("failed to load AWS config: %w", err)
	}

	c := kms.NewFromConfig(config)
	return c, nil
}
