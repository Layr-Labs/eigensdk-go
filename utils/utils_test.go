package utils

import (
	"math/big"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
)

func TestReadFile(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		content     []byte
		expectError bool
	}{
		{
			name:    "Read existing file",
			content: []byte("Hello, world!"),
		},
		{
			name:        "File does not exist",
			content:     nil,
			expectError: true,
		},
		{
			name:    "Empty file",
			content: []byte(""),
		},
		{
			name:    "Large file",
			content: make([]byte, 1024*1024),
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			var filePath string
			if tc.content != nil {
				// Create a temporary file
				tmpFile, err := os.CreateTemp("", "testfile")
				assert.NoError(t, err)
				defer os.Remove(tmpFile.Name())

				// Write content to the file
				_, err = tmpFile.Write(tc.content)
				assert.NoError(t, err)

				// Close the file
				err = tmpFile.Close()
				assert.NoError(t, err)

				filePath = tmpFile.Name()
			} else {
				// Non-existent file path
				filePath = "nonexistentfile.txt"
			}

			// Read the file using ReadFile function
			readContent, err := ReadFile(filePath)
			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.content, readContent)
			}
		})
	}
}

func TestReadYamlConfig(t *testing.T) {
	t.Parallel()

	type Config struct {
		Name string `yaml:"name"`
		Age  int    `yaml:"age"`
	}

	tests := []struct {
		name        string
		yamlContent string
		expectError bool
		expected    Config
	}{
		{
			name: "Valid YAML",
			yamlContent: `
name: John Doe
age: 30
`,
			expected: Config{Name: "John Doe", Age: 30},
		},
		{
			name: "Missing fields",
			yamlContent: `
name: Alice
`,
			expected: Config{Name: "Alice", Age: 0},
		},
		{
			name:        "Empty YAML",
			yamlContent: ``,
			expected:    Config{},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			// Create a temporary YAML file
			tmpFile, err := os.CreateTemp("", "testconfig*.yaml")
			assert.NoError(t, err)
			defer os.Remove(tmpFile.Name())

			_, err = tmpFile.Write([]byte(tc.yamlContent))
			assert.NoError(t, err)
			err = tmpFile.Close()
			assert.NoError(t, err)

			var config Config
			err = ReadYamlConfig(tmpFile.Name(), &config)
			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, config)
			}
		})
	}
}

func TestReadJsonConfig(t *testing.T) {
	t.Parallel()

	type Config struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	tests := []struct {
		name        string
		jsonContent string
		expectError bool
		expected    Config
	}{
		{
			name: "Valid JSON",
			jsonContent: `{
	"name": "Jane Doe",
	"age": 25
}`,
			expected: Config{Name: "Jane Doe", Age: 25},
		},
		{
			name:        "Missing fields",
			jsonContent: `{"name": "Alice"}`,
			expected:    Config{Name: "Alice", Age: 0},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			// Create a temporary JSON file
			tmpFile, err := os.CreateTemp("", "testconfig*.json")
			assert.NoError(t, err)
			defer os.Remove(tmpFile.Name())

			_, err = tmpFile.Write([]byte(tc.jsonContent))
			assert.NoError(t, err)
			err = tmpFile.Close()
			assert.NoError(t, err)

			var config Config
			err = ReadJsonConfig(tmpFile.Name(), &config)
			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, config)
			}
		})
	}
}

func TestEcdsaPrivateKeyToAddress(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name          string
		privateKeyHex string
		expectError   bool
		expectedAddr  gethcommon.Address
	}{
		{
			name:          "Valid private key",
			privateKeyHex: "4c0883a69102937d6231471b5dbb6204fe512961708279e6e82cc073aa8aa1a9",
			expectedAddr:  gethcommon.HexToAddress("0x8019FFe7A44A943c3a507C94D418DA3eD829f04d"),
		},
		{
			name:          "Invalid private key (non-hex)",
			privateKeyHex: "invalidkey",
			expectError:   true,
		},
		{
			name:          "Invalid private key (empty)",
			privateKeyHex: "",
			expectError:   true,
		},
		{
			name:          "Valid private key with 0x prefix",
			privateKeyHex: "0x8f2a5594902d6e44c069ad1c6e42c3e92a0416b4cf8ba689344844ab63b0f5f8",
			expectedAddr:  gethcommon.HexToAddress("0x8089DD5E304F65219dF307357256C07896f87703"),
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			privateKeyHex := strings.TrimPrefix(tc.privateKeyHex, "0x")
			privateKey, err := crypto.HexToECDSA(privateKeyHex)
			if tc.expectError {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)

			address, err := EcdsaPrivateKeyToAddress(privateKey)
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedAddr.Hex(), address.Hex())
		})
	}
}

func TestRoundUpDivideBig(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		a           *big.Int
		b           *big.Int
		expected    *big.Int
		expectPanic bool
	}{
		{
			name:        "Divide 5 by 2",
			a:           big.NewInt(5),
			b:           big.NewInt(2),
			expected:    big.NewInt(3),
			expectPanic: false,
		},
		{
			name:        "Divide 10 by 3",
			a:           big.NewInt(10),
			b:           big.NewInt(3),
			expected:    big.NewInt(4),
			expectPanic: false,
		},
		{
			name:        "Divide 100 by 33",
			a:           big.NewInt(100),
			b:           big.NewInt(33),
			expected:    big.NewInt(4),
			expectPanic: false,
		},
		{
			name:        "Divide 0 by 1",
			a:           big.NewInt(0),
			b:           big.NewInt(1),
			expected:    big.NewInt(0),
			expectPanic: false,
		},
		{
			name:        "Divide 1 by 1",
			a:           big.NewInt(1),
			b:           big.NewInt(1),
			expected:    big.NewInt(1),
			expectPanic: false,
		},
		{
			name:        "Divide by zero (100/0)",
			a:           big.NewInt(100),
			b:           big.NewInt(0),
			expected:    nil,
			expectPanic: true,
		},
		{
			name:        "Both zero (0/0)",
			a:           big.NewInt(0),
			b:           big.NewInt(0),
			expected:    nil,
			expectPanic: true,
		},
		{
			name:        "Large division",
			a:           big.NewInt(999999999999999999),
			b:           big.NewInt(3),
			expected:    big.NewInt(333333333333333333),
			expectPanic: false,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			if tc.expectPanic {
				assert.Panics(t, func() {
					RoundUpDivideBig(new(big.Int).Set(tc.a), tc.b)
				})
			} else {
				result := RoundUpDivideBig(new(big.Int).Set(tc.a), tc.b)
				assert.Equal(t, tc.expected.String(), result.String(), "RoundUpDivideBig(%s, %s)", tc.a.String(), tc.b.String())
			}
		})
	}
}

func TestReadPublicURL(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Serve different content based on the request URL
		if r.URL.Path == "/small" {
			w.WriteHeader(http.StatusOK)
			w.Write(make([]byte, 1024)) //nolint:errcheck
		} else if r.URL.Path == "/large" {
			w.WriteHeader(http.StatusOK)
			w.Write(make([]byte, 2*1024*1024)) //nolint:errcheck
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer server.Close()

	tests := []struct {
		name        string
		urlPath     string
		expectedErr error
	}{
		{
			name:        "request < 1mb",
			urlPath:     "/small",
			expectedErr: nil,
		},
		{
			name:        "request too large",
			urlPath:     "/large",
			expectedErr: ErrResponseTooLarge,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := server.URL + tt.urlPath
			client := &http.Client{}
			_, err := ReadPublicURL(url, client)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}

func TestIsImageURL(t *testing.T) {
	t.Parallel()

	// Create an httptest server that serves a valid PNG image
	pngData := []byte{137, 80, 78, 71, 13, 10, 26, 10}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, ".png") {
			w.Header().Set("Content-Type", "image/png")
			w.WriteHeader(http.StatusOK)
			w.Write(pngData) //nolint:errcheck
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer server.Close()

	tests := []struct {
		name        string
		url         string
		expectedErr error
	}{
		{
			name:        "valid PNG image URL",
			url:         server.URL + "/image.png",
			expectedErr: nil,
		},
		{
			name:        "invalid image extension",
			url:         server.URL + "/image.jpg",
			expectedErr: ErrInvalidImageExtension,
		},
		{
			name:        "non-image URL",
			url:         server.URL + "/notfound",
			expectedErr: ErrInvalidImageExtension,
		},
		{
			name:        "image with wrong MIME type",
			url:         server.URL + "/image.jpg",
			expectedErr: ErrInvalidImageExtension,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &http.Client{}
			err := IsImageURL(tt.url, client)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}

func TestIsValidEthereumAddress(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		address string
		isValid bool
	}{
		{
			name:    "valid address",
			address: "0x1234567890abcdef1234567890abcdef12345678",
			isValid: true,
		},
		{
			name:    "uppercase",
			address: "0x1234567890ABCDEF1234567890ABCDEF12345678",
			isValid: true,
		},
		{
			name:    "too short",
			address: "0x1234567890abcdef1234567890abcdef123456",
			isValid: false,
		},
		{
			name:    "missing 0x prefix",
			address: "001234567890abcdef1234567890abcdef12345678",
			isValid: false,
		},
		{
			name:    "non-hex characters",
			address: "0x1234567890abcdef1234567890abcdef123ÅÅÅÅÅ",
			isValid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsValidEthereumAddress(tt.address)
			assert.Equal(t, tt.isValid, result)
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
