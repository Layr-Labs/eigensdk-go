package types

import (
	"errors"
	"net/url"
	"regexp"
)

func isValidEthereumAddress(address string) bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	return re.MatchString(address)
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
