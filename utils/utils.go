package utils

import (
	"crypto/ecdsa"
	"errors"
	"net/url"
	"strings"
	"time"

	"math/big"

	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
)

const (
	PngMimeType = "image/png"

	TextRegex = `^[a-zA-Z0-9 +.,;:?!'’"“”\-_/()\[\]~&#$—%]+$`

	// Limit Http response to 1 MB
	httpResponseLimitBytes = 1 * 1024 * 1024

	TextCharsLimit = 500
)

var (
	// ImageExtensions List of common image file extensions
	// Only support PNG for now to reduce surface area of image validation
	// We do NOT want to support formats like SVG since they can be used for javascript injection
	// If we get pushback on only supporting png, we can support jpg, jpeg, gif, etc. later
	ImageExtensions = []string{".png"}

	// Regular expression to ethereum address
	ethAddrPattern = regexp.MustCompile("^0x[0-9a-fA-F]{40}$")

	// Regular expression to validate text
	textPattern = regexp.MustCompile(TextRegex)

	// Regular expression to validate URLs
	rawGitHubUrlPattern = regexp.MustCompile(`^https?://raw\.githubusercontent\.com/.*$`)

	// Regular expression to validate URLs
	twitterUrlPattern = regexp.MustCompile(`^(?:https?://)?(?:www\.)?(?:twitter\.com/\w+|x\.com/\w+)(?:/?|$)`)

	// Regular expression to validate URLs
	urlPattern = regexp.MustCompile(`^(https?)://[^\s/$.?#].[^\s]*$`)
)

func ReadFile(path string) ([]byte, error) {
	b, err := os.ReadFile(filepath.Clean(path))
	if err != nil {
		return nil, err
	}
	return b, nil
}

func EcdsaPrivateKeyToAddress(privateKey *ecdsa.PrivateKey) (gethcommon.Address, error) {
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return gethcommon.Address{}, errors.New("ErrCannotGetECDSAPubKey")
	}
	return crypto.PubkeyToAddress(*publicKeyECDSA), nil
}

// RoundUpDivideBig divides two positive big.Int numbers and rounds up the result.
func RoundUpDivideBig(a, b *big.Int) *big.Int {
	one := new(big.Int).SetUint64(1)
	res := new(big.Int)
	a.Add(a, b)
	a.Sub(a, one)
	res.Div(a, b)
	return res
}

func IsValidEthereumAddress(address string) bool {
	return ethAddrPattern.MatchString(address)
}

func ReadPublicURL(url string, httpClient *http.Client) ([]byte, error) {
	// Allow no redirects
	httpClient.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}
	httpClient.Timeout = 3 * time.Second

	resp, err := httpClient.Get(url)
	if err != nil {
		return []byte{}, err
	}

	if resp.StatusCode >= 400 {
		return []byte{}, fmt.Errorf("error fetching url: %s", resp.Status)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("error closing url body")
		}
	}(resp.Body)

	// allow images of up to 1 MiB
	response, err := io.ReadAll(http.MaxBytesReader(nil, resp.Body, httpResponseLimitBytes))
	if err != nil {
		// We are doing this because errors.Is(err) check doesn't work for this
		// since MaxBytesError has pointer receiver. Not sure what is the correct
		// way to do this.
		maxByteErr := http.MaxBytesError{}
		if err.Error() == maxByteErr.Error() {
			return nil, ErrResponseTooLarge
		}
		return nil, err
	}
	return response, nil
}

func CheckIfValidTwitterURL(twitterURL string) error {
	// Basic validation
	err := CheckBasicURLValidation(twitterURL)
	if err != nil {
		return err
	}

	// Check if the URL matches the regular expression
	if !twitterUrlPattern.MatchString(twitterURL) {
		return ErrInvalidTwitterUrlRegex
	}

	return nil
}

func CheckBasicURLValidation(rawUrl string) error {
	if len(rawUrl) == 0 {
		return ErrEmptyUrl
	}

	if strings.Contains(rawUrl, "localhost") || strings.Contains(rawUrl, "127.0.0.1") {
		return ErrUrlPointingToLocalServer
	}

	if len(rawUrl) > 1024 {
		return ErrInvalidUrlLength
	}

	parsedURL, err := url.Parse(rawUrl)
	if err != nil {
		return err
	}

	// Check if the URL is valid
	if parsedURL.Scheme != "" && parsedURL.Host != "" {
		return nil
	} else {
		return ErrInvalidUrl
	}
}

func CheckIfUrlIsValid(rawUrl string) error {
	// Basic validation
	err := CheckBasicURLValidation(rawUrl)
	if err != nil {
		return err
	}

	// Check if the URL matches the regular expression
	if !urlPattern.MatchString(rawUrl) {
		return ErrInvalidUrl
	}

	return nil
}

func IsImageURL(urlString string, httpClient *http.Client) error {
	// Parse the URL
	parsedURL, err := url.Parse(urlString)
	if err != nil {
		return err
	}

	// Extract the path component from the URL
	path := parsedURL.Path

	// Get the file extension
	extension := filepath.Ext(path)

	// Check if the extension is in the list of image extensions
	for _, imgExt := range ImageExtensions {
		if strings.EqualFold(extension, imgExt) {
			imageBytes, err := ReadPublicURL(urlString, httpClient)
			if err != nil {
				return err
			}

			contentType := http.DetectContentType(imageBytes)
			if contentType != PngMimeType {
				return ErrInvalidImageMimeType
			}
			return nil
		}
	}

	return ErrInvalidImageExtension
}

func ValidateText(text string) error {
	if len(text) == 0 {
		return ErrEmptyText
	}

	if len(text) > TextCharsLimit {
		return ErrTextTooLong(TextCharsLimit)
	}

	// Check if the URL matches the regular expression
	if !textPattern.MatchString(text) {
		return ErrInvalidText
	}

	return nil
}

func ValidateRawGithubUrl(url string) error {
	// Basic validation
	err := CheckBasicURLValidation(url)
	if err != nil {
		return err
	}

	// Check if the URL matches the regular expression
	if !rawGitHubUrlPattern.MatchString(url) {
		return ErrInvalidGithubRawUrl
	}

	return nil
}

func Add0x(address string) string {
	if strings.HasPrefix(address, "0x") {
		return address
	}
	return "0x" + address
}

func Trim0x(address string) string {
	return strings.TrimPrefix(address, "0x")
}
