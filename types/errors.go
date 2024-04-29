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
