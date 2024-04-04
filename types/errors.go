package types

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrInvalidOperatorAddress           = errors.New("invalid operator address")
	ErrInvalidEarningsReceiverAddress   = errors.New("invalid EarningsReceiverAddress address")
	ErrInvalidDelegationApproverAddress = fmt.Errorf(
		"invalid DelegationApproverAddress address, it should be either %s or a valid non zero ethereum address",
		ZeroAddress,
	)

	ErrEmptyText                = errors.New("text is empty")
	ErrInvalidText              = errors.New("text is invalid")
	ErrEmptyUrl                 = errors.New("url is empty")
	ErrUrlPointingToLocalServer = errors.New("url should not point to local server")
	ErrInvalidUrlLength         = errors.New("url length should be no larger than 1024 character")
	ErrInvalidUrl               = errors.New("invalid url")

	ErrImageFetchFailed      = errors.New("error fetching operator logo")
	ErrInvalidImageMimeType  = errors.New("invalid image mime-type. only png is supported")
	ErrInvalidImageExtension = errors.New(
		"invalid image extension. only " + strings.Join(ImageExtensions, ",") + " is supported",
	)

	// Metadata Errors

	ErrNameRequired           = errors.New("name is required")
	ErrDescriptionRequired    = errors.New("description is required")
	ErrDescriptionTooLong     = errors.New("description should be less than 200 characters")
	ErrInvalidTwitterUrlRegex = errors.New(
		"invalid twitter url, it should be of the format https://twitter.com/<username> or https://x.com/<username>",
	)

	ErrInvalidName        = errors.New("name is invalid")
	ErrInvalidDescription = errors.New("description is invalid")
	ErrTextTooLong        = errors.New("text should be less than 200 characters")
	ErrLogoRequired       = errors.New("logo is required")
	ErrInvalidWebsiteUrl  = errors.New("invalid website url")
	ErrInvalidTwitterUrl  = errors.New("invalid twitter url")

	ErrInvalidMetadataUrl         = errors.New("invalid metadata url")
	ErrUnmarshalOperatorMetadata  = errors.New("unable to unmarshal operator metadata")
	ErrReadingMetadataUrlResponse = errors.New("error reading metadata url body")
)

func WrapError(mainErr error, subErr error) error {
	// Some times the wrap will wrap a nil error
	if mainErr == nil && subErr == nil {
		return nil
	}

	if mainErr == nil && subErr != nil {
		return fmt.Errorf("sub error: %s", subErr.Error())
	}

	if mainErr != nil && subErr == nil {
		return fmt.Errorf("%s: unknown sub error", mainErr.Error())
	}

	return fmt.Errorf("%s: %s", mainErr.Error(), subErr.Error())
}
