package types

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidOperatorAddress           = errors.New("invalid operator address")
	ErrInvalidEarningsReceiverAddress   = errors.New("invalid EarningsReceiverAddress address")
	ErrInvalidDelegationApproverAddress = fmt.Errorf(
		"invalid DelegationApproverAddress address, it should be either %s or a valid non zero ethereum address",
		ZeroAddress,
	)

	ErrInvalidName        = errors.New("name is invalid")
	ErrInvalidDescription = errors.New("description is invalid")
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
