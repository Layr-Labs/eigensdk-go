package types

import (
	"io"
	"net/http"
	"net/url"
	"path/filepath"
	"regexp"
	"strings"
)

const (
	PngMimeType = "image/png"
)

var (
	// ImageExtensions List of common image file extensions
	// Only support PNG for now to reduce surface area of image validation
	// We do NOT want to support formats like SVG since they can be used for javascript injection
	// If we get pushback on only supporting png, we can support jpg, jpeg, gif, etc. later
	ImageExtensions = []string{".png"}
)

// OperatorMetadata is the metadata operator uploads while registering
// itself to EigenLayer
type OperatorMetadata struct {

	// Name of the operator
	// It is a required field
	Name string `yaml:"name" json:"name"`

	// Website of the operator
	// It is a required field
	Website string `yaml:"website" json:"website"`

	// Description of the operator. There is a 200-character limit
	// It is a required field
	Description string `yaml:"description" json:"description"`

	// Logo of the operator. This should be a link to a image file
	// which is publicly accessible
	// It is a required field
	Logo string `yaml:"logo" json:"logo"`

	// Twitter handle of the operator
	// It is an optional field
	Twitter string `yaml:"twitter" json:"twitter"`
}

func (om *OperatorMetadata) Validate() error {
	if len(om.Name) == 0 {
		return ErrNameRequired
	}

	if len(om.Description) == 0 {
		return ErrDescriptionRequired
	}

	if len(om.Description) > 200 {
		return ErrDescriptionTooLong
	}

	if len(om.Logo) == 0 {
		return ErrLogoRequired
	}

	if err := isImageURL(om.Logo); err != nil {
		return err
	}

	if len(om.Website) != 0 {
		err := checkIfUrlIsValid(om.Website)
		if err != nil {
			return wrapError(ErrInvalidWebsiteUrl, err)
		}
	}

	if len(om.Twitter) != 0 {
		err := checkIfUrlIsValid(om.Twitter)
		if err != nil {
			return wrapError(ErrInvalidTwitterUrl, err)
		}
	}

	return nil
}

func checkIfUrlIsValid(rawUrl string) error {
	if len(rawUrl) == 0 {
		return ErrEmptyUrl
	}

	if strings.Contains(rawUrl, "localhost") || strings.Contains(rawUrl, "127.0.0.1") {
		return ErrUrlPointingToLocalServer
	}

	if len(rawUrl) > 1024 {
		return ErrInvalidUrlLength
	}

	// Regular expression to validate URLs
	urlPattern := regexp.MustCompile(`^(https?|ftp)://[^\s/$.?#].[^\s]*$`)

	// Check if the URL matches the regular expression
	if !urlPattern.MatchString(rawUrl) {
		return ErrInvalidUrl
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

func isImageURL(urlString string) error {
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
			imageResponse, err := http.Get(urlString)
			if err != nil {
				return err
			}
			imageBytes, err := io.ReadAll(imageResponse.Body)
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
