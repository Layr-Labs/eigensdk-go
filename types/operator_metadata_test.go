package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOperatorMetadata(t *testing.T) {
	var tests = []struct {
		name          string
		metadata      OperatorMetadata
		expectedError error
	}{
		{
			name: "Valid metadata",
			metadata: OperatorMetadata{
				Name:        "test",
				Description: "test",
				Logo:        "https://goerli-operator-metadata.s3.amazonaws.com/eigenlayer.png",
				Twitter:     "https://twitter.com/test",
				Website:     "https://test.com",
			},
		},
		{
			name: "Invalid metadata - no name",
			metadata: OperatorMetadata{
				Name:        "",
				Description: "test",
				Logo:        "https://goerli-operator-metadata.s3.amazonaws.com/eigenlayer.png",
				Twitter:     "https://twitter.com/test",
				Website:     "https://test.com",
			},
			expectedError: ErrNameRequired,
		},
		{
			name: "Invalid metadata - no description",
			metadata: OperatorMetadata{
				Name:        "test",
				Description: "",
				Logo:        "https://goerli-operator-metadata.s3.amazonaws.com/eigenlayer.png",
				Twitter:     "https://twitter.com/test",
				Website:     "https://test.com",
			},
			expectedError: ErrDescriptionRequired,
		},
		{
			name: "Invalid metadata - wrong image format",
			metadata: OperatorMetadata{
				Name:        "test",
				Description: "My operator",
				Logo:        "https://test.com/test.svg",
				Twitter:     "https://twitter.com/test",
				Website:     "https://test.com",
			},
			expectedError: ErrInvalidImageExtension,
		},
		{
			name: "Invalid metadata - invalid mime type with correct extension",
			metadata: OperatorMetadata{
				Name:        "test",
				Description: "My operator",
				Logo:        "https://goerli-operator-metadata.s3.amazonaws.com/cat.png",
				Twitter:     "https://twitter.com/test",
				Website:     "https://test.com",
			},
			expectedError: ErrInvalidImageMimeType,
		},
		{
			name: "Invalid metadata - description > 200 characters",
			metadata: OperatorMetadata{
				Name:        "test",
				Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.",
				Logo:        "https://goerli-operator-metadata.s3.amazonaws.com/eigenlayer.png",
				Twitter:     "https://twitter.com/test",
				Website:     "https://test.com",
			},
			expectedError: ErrDescriptionTooLong,
		},
		{
			name: "Invalid metadata - no logo",
			metadata: OperatorMetadata{
				Name:        "test",
				Description: "test",
				Logo:        "",
				Twitter:     "https://twitter.com/test",
				Website:     "https://test.com",
			},
			expectedError: ErrLogoRequired,
		},
		{
			name: "Invalid metadata - invalid logo no extension",
			metadata: OperatorMetadata{
				Name:        "test",
				Description: "test",
				Logo:        "https://test.com/test",
				Twitter:     "https://twitter.com/test",
				Website:     "https://test.com",
			},
			expectedError: ErrInvalidImageExtension,
		},
		{
			name: "Invalid metadata - invalid website url #1",
			metadata: OperatorMetadata{
				Name:        "test",
				Description: "test",
				Logo:        "https://goerli-operator-metadata.s3.amazonaws.com/eigenlayer.png",
				Twitter:     "https://twitter.com/test",
				Website:     "https",
			},
			expectedError: WrapError(ErrInvalidWebsiteUrl, ErrInvalidUrl),
		},
		{
			name: "Invalid metadata - invalid website url #2",
			metadata: OperatorMetadata{
				Name:        "test",
				Description: "test",
				Logo:        "https://goerli-operator-metadata.s3.amazonaws.com/eigenlayer.png",
				Twitter:     "https://twitter.com/test",
				Website:     "https:/test.com",
			},
			expectedError: WrapError(ErrInvalidWebsiteUrl, ErrInvalidUrl),
		},
		{
			name: "Invalid metadata - invalid website url #3",
			metadata: OperatorMetadata{
				Name:        "test",
				Description: "test",
				Logo:        "https://goerli-operator-metadata.s3.amazonaws.com/eigenlayer.png",
				Twitter:     "https://twitter.com/test",
				Website:     "ps://test.com",
			},
			expectedError: WrapError(ErrInvalidWebsiteUrl, ErrInvalidUrl),
		},
		{
			name: "Invalid metadata - invalid twitter url #1",
			metadata: OperatorMetadata{
				Name:        "test",
				Description: "test",
				Logo:        "https://goerli-operator-metadata.s3.amazonaws.com/eigenlayer.png",
				Twitter:     "http",
				Website:     "https://test.com",
			},
			expectedError: WrapError(ErrInvalidTwitterUrl, ErrInvalidUrl),
		},
		{
			name: "Invalid metadata - invalid twitter url #2",
			metadata: OperatorMetadata{
				Name:        "test",
				Description: "test",
				Logo:        "https://goerli-operator-metadata.s3.amazonaws.com/eigenlayer.png",
				Twitter:     "ht://twitter.com/test",
				Website:     "https://test.com",
			},
			expectedError: WrapError(ErrInvalidTwitterUrl, ErrInvalidUrl),
		},
		{
			name: "Invalid metadata - invalid twitter url #3",
			metadata: OperatorMetadata{
				Name:        "test",
				Description: "test",
				Logo:        "https://goerli-operator-metadata.s3.amazonaws.com/eigenlayer.png",
				Twitter:     "https:/twitt",
				Website:     "https://test.com",
			},
			expectedError: WrapError(ErrInvalidTwitterUrl, ErrInvalidUrl),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.metadata.Validate()
			assert.Equal(t, tt.expectedError, err, "error should be equal")
		})
	}
}
