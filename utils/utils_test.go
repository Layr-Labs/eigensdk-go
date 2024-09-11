package utils

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidEthereumAddress(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		address  string
		expected bool
	}{
		{
			name:     "valid address",
			address:  "0x1234567890abcdef1234567890abcdef12345678",
			expected: true,
		},
		{
			name:     "uppercase",
			address:  "0x1234567890ABCDEF1234567890ABCDEF12345678",
			expected: true,
		},
		{
			name:     "too short",
			address:  "0x1234567890abcdef1234567890abcdef123456",
			expected: false,
		},
		{
			name:     "missing 0x prefix",
			address:  "001234567890abcdef1234567890abcdef12345678",
			expected: false,
		},
		{
			name:     "non-hex characters",
			address:  "0x1234567890abcdef1234567890abcdef123ÅÅÅÅÅ",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsValidEthereumAddress(tt.address)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestCheckIfValidTwitterURL(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		twitterURL  string
		expectedErr error
	}{
		{
			name:        "valid Twitter URL",
			twitterURL:  "https://twitter.com/user",
			expectedErr: nil,
		},
		{
			name:        "valid X URL",
			twitterURL:  "https://x.com/user",
			expectedErr: nil,
		},
		{
			name:        "invalid non-Twitter URL",
			twitterURL:  "https://example.com/user",
			expectedErr: ErrInvalidTwitterUrlRegex,
		},
		{
			name:        "empty URL",
			twitterURL:  "",
			expectedErr: ErrEmptyUrl,
		},
		{
			name:        "URL pointing to localhost",
			twitterURL:  "http://localhost",
			expectedErr: ErrUrlPointingToLocalServer,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CheckIfValidTwitterURL(tt.twitterURL)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}

func TestCheckIfUrlIsValid(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		rawUrl      string
		expectedErr error
	}{
		{
			name:        "valid http URL",
			rawUrl:      "http://example.com",
			expectedErr: nil,
		},
		{
			name:        "valid https URL",
			rawUrl:      "https://example.com",
			expectedErr: nil,
		},
		{
			name:        "empty URL",
			rawUrl:      "",
			expectedErr: ErrEmptyUrl,
		},
		{
			name:        "localhost",
			rawUrl:      "http://localhost",
			expectedErr: ErrUrlPointingToLocalServer,
		},
		{
			name:        "too long",
			rawUrl:      "http://example.com/" + string(make([]byte, 1025)),
			expectedErr: ErrInvalidUrlLength,
		},
		{
			name:        "invalid basic URL format",
			rawUrl:      "invalid-url",
			expectedErr: ErrInvalidUrl,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CheckIfUrlIsValid(tt.rawUrl)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}

func TestReadPublicURL(t *testing.T) {
	t.Parallel()

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

func TestValidateText(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		text        string
		expectedErr error
	}{
		{
			name:        "valid text",
			text:        "Hello, World!",
			expectedErr: nil,
		},
		{
			name:        "empty text",
			text:        "",
			expectedErr: ErrEmptyText,
		},
		{
			name:        "text too long",
			text:        strings.Repeat("a", TextCharsLimit+1),
			expectedErr: ErrTextTooLong(TextCharsLimit),
		},
		{
			name:        "text not too long",
			text:        strings.Repeat("a", TextCharsLimit-1),
			expectedErr: nil,
		},
		{
			name:        "invalid text with special characters",
			text:        "Invalid chars: @#$%^",
			expectedErr: ErrInvalidText,
		},
		{
			name:        "valid text with allowed special characters",
			text:        "Valid chars: -_(),.!?",
			expectedErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateText(tt.text)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}

func TestValidateRawGithubUrl(t *testing.T) {
	t.Parallel()

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
			name:        "localhost",
			url:         "localhost",
			expectedErr: ErrUrlPointingToLocalServer,
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

func TestAdd0x(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		address  string
		expected string
	}{
		{
			name:     "address with 0x prefix",
			address:  "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045",
			expected: "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045",
		},
		{
			name:     "address without 0x prefix",
			address:  "d8dA6BF26964aF9D7eEd9e03E53415D37aA96045",
			expected: "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Add0x(tt.address)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestTrim0x(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		address  string
		expected string
	}{
		{
			name:     "address with 0x prefix",
			address:  "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045",
			expected: "d8dA6BF26964aF9D7eEd9e03E53415D37aA96045",
		},
		{
			name:     "address without 0x prefix",
			address:  "d8dA6BF26964aF9D7eEd9e03E53415D37aA96045",
			expected: "d8dA6BF26964aF9D7eEd9e03E53415D37aA96045",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Trim0x(tt.address)
			assert.Equal(t, tt.expected, result)
		})
	}
}
