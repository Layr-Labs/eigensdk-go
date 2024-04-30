package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateRawGithubUrl(t *testing.T) {
	tests := []struct {
		name        string
		url         string
		expectedErr error
	}{
		{
			name: "valid raw github url",
			url:  "https://raw.githubusercontent.com/Layr-Labs/eigensdk-go/main/README.md",
		},
		{
			name:        "invalid raw github url",
			url:         "https://facebook.com",
			expectedErr: ErrInvalidGithubRawUrl,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateRawGithubUrl(tt.url)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}

func TestReadPublicURL(t *testing.T) {
	tests := []struct {
		name        string
		url         string
		expectedErr error
	}{
		{
			name:        "request < 1mb",
			url:         "https://raw.githubusercontent.com/shrimalmadhur/metadata/main/logo.png",
			expectedErr: nil,
		},
		{
			name:        "request too large",
			url:         "https://raw.githubusercontent.com/shrimalmadhur/metadata/main/2mb.png",
			expectedErr: ErrResponseTooLarge,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ReadPublicURL(tt.url)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}
