package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
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
