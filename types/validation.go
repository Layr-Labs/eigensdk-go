package types

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

const (
	PngMimeType = "image/png"

	TextRegex = `^[a-zA-Z0-9 +.,;:?!''"""\-_/()\[\]~&#$—%]+$`

	// httpResponseLimitBytes limits HTTP response to 1 MB
	httpResponseLimitBytes = 1 * 1024 * 1024

	TextCharsLimit = 500
)

var (
	// ImageExtensions List of common image file extensions
	// Only support PNG for now to reduce surface area of image validation
	ImageExtensions = []string{".png"}

	// ethAddrPattern is a regular expression to validate ethereum addresses
	ethAddrPattern = regexp.MustCompile("^0x[0-9a-fA-F]{40}$")

	// textPattern is a regular expression to validate text
	textPattern = regexp.MustCompile(TextRegex)

	// rawGitHubUrlPattern validates raw GitHub URLs
	rawGitHubUrlPattern = regexp.MustCompile(`^https?://raw\.githubusercontent\.com/.*$`)

	// twitterUrlPattern validates Twitter/X URLs
	twitterUrlPattern = regexp.MustCompile(`^(?:https?://)?(?:www\.)?(?:twitter\.com/\w+|x\.com/\w+)(?:/?|$)`)

	// urlPattern validates generic URLs
	urlPattern = regexp.MustCompile(`^(https?)://[^\s/$.?#].[^\s]*$`)
)

// Validation errors
var (
	ErrInvalidUrl          = errors.New("invalid url")
	ErrInvalidGithubRawUrl = errors.New("invalid github raw url")
	ErrInvalidText         = fmt.Errorf("invalid text format, doesn't conform to regex %s", TextRegex)
	ErrTextTooLong         = func(limit int) error {
		return fmt.Errorf("text should be less than %d characters", limit)
	}
	ErrEmptyText             = errors.New("text is empty")
	ErrInvalidImageExtension = errors.New(
		"invalid image extension. only " + strings.Join(ImageExtensions, ",") + " is supported",
	)
	ErrInvalidImageMimeType     = errors.New("invalid image mime-type. only png is supported")
	ErrInvalidUrlLength         = errors.New("url length should be no larger than 1024 character")
	ErrUrlPointingToLocalServer = errors.New("url should not point to local server")
	ErrEmptyUrl                 = errors.New("url is empty")
	ErrInvalidTwitterUrlRegex   = errors.New(
		"invalid twitter url, it should be of the format https://twitter.com/<username> or https://x.com/<username>",
	)
	ErrResponseTooLarge = errors.New("response too large, allowed size is 1 MB")
)

// IsValidEthereumAddress checks if the given string is a valid Ethereum address
func IsValidEthereumAddress(address string) bool {
	return ethAddrPattern.MatchString(address)
}

// ReadPublicURL fetches the content of a public URL with safety limits
func ReadPublicURL(url string) ([]byte, error) {
	httpClient := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Timeout: 3 * time.Second,
	}

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

	response, err := io.ReadAll(http.MaxBytesReader(nil, resp.Body, httpResponseLimitBytes))
	if err != nil {
		maxByteErr := http.MaxBytesError{}
		if err.Error() == maxByteErr.Error() {
			return nil, ErrResponseTooLarge
		}
		return nil, err
	}
	return response, nil
}

// CheckIfValidTwitterURL validates a Twitter/X URL
func CheckIfValidTwitterURL(twitterURL string) error {
	err := CheckBasicURLValidation(twitterURL)
	if err != nil {
		return err
	}

	if !twitterUrlPattern.MatchString(twitterURL) {
		return ErrInvalidTwitterUrlRegex
	}

	return nil
}

// CheckBasicURLValidation performs basic URL validation checks
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

	if parsedURL.Scheme != "" && parsedURL.Host != "" {
		return nil
	}
	return ErrInvalidUrl
}

// CheckIfUrlIsValid validates a URL with pattern matching
func CheckIfUrlIsValid(rawUrl string) error {
	err := CheckBasicURLValidation(rawUrl)
	if err != nil {
		return err
	}

	if !urlPattern.MatchString(rawUrl) {
		return ErrInvalidUrl
	}

	return nil
}

// IsImageURL validates that a URL points to a valid image
func IsImageURL(urlString string) error {
	parsedURL, err := url.Parse(urlString)
	if err != nil {
		return err
	}

	path := parsedURL.Path
	extension := filepath.Ext(path)

	for _, imgExt := range ImageExtensions {
		if strings.EqualFold(extension, imgExt) {
			imageBytes, err := ReadPublicURL(urlString)
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

// ValidateText validates text content against format and length rules
func ValidateText(text string) error {
	if len(text) == 0 {
		return ErrEmptyText
	}

	if len(text) > TextCharsLimit {
		return ErrTextTooLong(TextCharsLimit)
	}

	if !textPattern.MatchString(text) {
		return ErrInvalidText
	}

	return nil
}

// ValidateRawGithubUrl validates that a URL is a valid raw GitHub URL
func ValidateRawGithubUrl(url string) error {
	err := CheckBasicURLValidation(url)
	if err != nil {
		return err
	}

	if !rawGitHubUrlPattern.MatchString(url) {
		return ErrInvalidGithubRawUrl
	}

	return nil
}

// WrapError wraps two errors together
func WrapError(mainErr interface{}, subErr interface{}) error {
	var main, sub error
	main = typedErr(mainErr)
	sub = typedErr(subErr)

	if main == nil && sub == nil {
		return nil
	}

	if main == nil && sub != nil {
		return sub
	}

	if main != nil && sub == nil {
		return main
	}

	return fmt.Errorf("%w: %w", main, sub)
}

// typedErr converts an interface to an error
func typedErr(e interface{}) error {
	switch t := e.(type) {
	case error:
		return t
	case string:
		return errors.New(t)
	default:
		return nil
	}
}
