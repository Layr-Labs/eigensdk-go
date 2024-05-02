package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
)

func GetAWSConfig(accessKey, secretAccessKey, region, endpointURL string) (*aws.Config, error) {
	createClient := func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		if endpointURL != "" {
			return aws.Endpoint{
				PartitionID:   "aws",
				URL:           endpointURL,
				SigningRegion: region,
			}, nil
		}

		// returning EndpointNotFoundError will allow the service to fallback to its default resolution
		return aws.Endpoint{}, &aws.EndpointNotFoundError{}
	}
	customResolver := aws.EndpointResolverWithOptionsFunc(createClient)

	cfg, errCfg := config.LoadDefaultConfig(context.Background(),
		config.WithRegion(region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKey, secretAccessKey, "")),
		config.WithEndpointResolverWithOptions(customResolver),
		config.WithRetryMode(aws.RetryModeStandard),
	)
	if errCfg != nil {
		return nil, errCfg
	}
	return &cfg, nil
}
