package types

import (
	"errors"
	"fmt"
	"net/url"
	"path/filepath"
	"regexp"
	"strings"
)

// OperatorMetadata is the metadata operator uploads while registering
// itself to eigenlayer
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
		return errors.New("name is required")
	}

	if len(om.Description) == 0 {
		return errors.New("description is required")
	}

	if len(om.Description) > 200 {
		return errors.New("description should be less than 200 characters")
	}

	if len(om.Logo) == 0 {
		return errors.New("logo is required")
	}

	if err := isImageURL(om.Logo); err != nil {
		return err
	}

	if len(om.Website) != 0 {
		err := checkIfUrlIsValid(om.Website)
		if err != nil {
			fmt.Println("error validating website url")
			return err
		}
	}

	if len(om.Twitter) != 0 {
		err := checkIfUrlIsValid(om.Twitter)
		if err != nil {
			fmt.Println("error validating twitter profile url")
			return err
		}
	}

	return nil
}

func checkIfUrlIsValid(rawUrl string) error {
	// Regular expression to validate URLs
	urlPattern := regexp.MustCompile(`^(https?|ftp)://[^\s/$.?#].[^\s]*$`)

	// Check if the URL matches the regular expression
	if !urlPattern.MatchString(rawUrl) {
		return errors.New("invalid url")
	}

	parsedURL, err := url.Parse(rawUrl)
	if err != nil {
		return err
	}

	// Check if the URL is valid
	if parsedURL.Scheme != "" && parsedURL.Host != "" {
		return nil
	} else {
		return errors.New("invalid url")
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

	// List of common image file extensions
	// Only support PNG for now to reduce surface area of image validation
	// We do NOT want to support formats like SVG since they can be used for javascript injection
	// If we get pushback on only supporting png, we can support jpg, jpeg, gif, etc. later
	imageExtensions := []string{".png"}

	// Check if the extension is in the list of image extensions
	for _, imgExt := range imageExtensions {
		if strings.EqualFold(extension, imgExt) {
			return nil
		}
	}

	return errors.New("invalid image extension. only " + strings.Join(imageExtensions, ",") + "is supported")
}
