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
			name: "Valid metadata with twitter.com url",
			metadata: OperatorMetadata{
				Name:        "Ethereum Utopia",
				Description: "test",
				Logo:        "https://goerli-operator-metadata.s3.amazonaws.com/eigenlayer.png",
				Twitter:     "https://twitter.com/test",
				Website:     "https://test.com",
			},
		},
		{
			name: "Valid metadata with twitter.com url with /",
			metadata: OperatorMetadata{
				Name:        "test",
				Description: "test",
				Logo:        "https://goerli-operator-metadata.s3.amazonaws.com/eigenlayer.png",
				Twitter:     "https://twitter.com/test/",
				Website:     "https://test.com",
			},
		},
		{
			name: "Valid metadata with x.com twitter url",
			metadata: OperatorMetadata{
				Name:        "test",
				Description: "test",
				Logo:        "https://goerli-operator-metadata.s3.amazonaws.com/eigenlayer.png",
				Twitter:     "https://x.com/test",
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
			expectedError: WrapError(ErrInvalidName, ErrEmptyText),
		},
		{
			name: "Invalid metadata - name has js script",
			metadata: OperatorMetadata{
				Name:        "<script> alert('test') </script>",
				Description: "test",
				Logo:        "https://goerli-operator-metadata.s3.amazonaws.com/eigenlayer.png",
				Twitter:     "https://twitter.com/test",
				Website:     "https://test.com",
			},
			expectedError: WrapError(ErrInvalidName, ErrInvalidText),
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
			expectedError: WrapError(ErrInvalidDescription, ErrEmptyText),
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
			name: "Invalid metadata - name > 200 characters",
			metadata: OperatorMetadata{
				Name:        "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.",
				Description: "test",
				Logo:        "https://goerli-operator-metadata.s3.amazonaws.com/eigenlayer.png",
				Twitter:     "https://twitter.com/test",
				Website:     "https://test.com",
			},
			expectedError: WrapError(ErrInvalidName, ErrTextTooLong),
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
			expectedError: WrapError(ErrInvalidDescription, ErrTextTooLong),
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
			name: "Invalid url - FTP url",
			metadata: OperatorMetadata{
				Name:        "test",
				Description: "test",
				Logo:        "https://goerli-operator-metadata.s3.amazonaws.com/eigenlayer.png",
				Twitter:     "ftp://twitter.com/test",
				Website:     "https://test.com",
			},
			expectedError: WrapError(ErrInvalidTwitterUrl, ErrInvalidTwitterUrlRegex),
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
			expectedError: WrapError(ErrInvalidTwitterUrl, ErrInvalidTwitterUrlRegex),
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
		{
			name: "Invalid metadata - invalid twitter url #4 - not twitter url",
			metadata: OperatorMetadata{
				Name:        "test",
				Description: "test",
				Logo:        "https://goerli-operator-metadata.s3.amazonaws.com/eigenlayer.png",
				Twitter:     "https://facebook.com/test",
				Website:     "https://test.com",
			},
			expectedError: WrapError(ErrInvalidTwitterUrl, ErrInvalidTwitterUrlRegex),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.metadata.Validate()
			assert.Equal(t, tt.expectedError, err, "error should be equal")
		})
	}
}
