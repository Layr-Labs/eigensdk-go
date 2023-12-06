package types

import (
	"errors"
	"fmt"
	"net/url"
	"path/filepath"
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

	if !isImageURL(om.Logo) {
		return errors.New("logo must be a valid image url")
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

func isImageURL(urlString string) bool {
	// Parse the URL
	parsedURL, err := url.Parse(urlString)
	if err != nil {
		return false
	}

	// Extract the path component from the URL
	path := parsedURL.Path

	// Get the file extension
	extension := filepath.Ext(path)

	// List of common image file extensions
	imageExtensions := []string{".jpg", ".jpeg", ".png", ".gif", ".bmp", ".svg", ".webp"}

	// Check if the extension is in the list of image extensions
	for _, imgExt := range imageExtensions {
		if strings.EqualFold(extension, imgExt) {
			return true
		}
	}

	return false
}
