// Package utils provides utility functions.
//
// Deprecated: Validation functions and errors have been moved to the types package.
// This file maintains backward compatibility. New code should use the types package directly.
package utils

import (
	"github.com/Layr-Labs/eigensdk-go/types"
)

// Deprecated: Use types.ErrInvalidUrl instead.
var ErrInvalidUrl = types.ErrInvalidUrl

// Deprecated: Use types.ErrInvalidGithubRawUrl instead.
var ErrInvalidGithubRawUrl = types.ErrInvalidGithubRawUrl

// Deprecated: Use types.ErrInvalidText instead.
var ErrInvalidText = types.ErrInvalidText

// Deprecated: Use types.ErrTextTooLong instead.
var ErrTextTooLong = types.ErrTextTooLong

// Deprecated: Use types.ErrEmptyText instead.
var ErrEmptyText = types.ErrEmptyText

// Deprecated: Use types.ErrInvalidImageExtension instead.
var ErrInvalidImageExtension = types.ErrInvalidImageExtension

// Deprecated: Use types.ErrInvalidImageMimeType instead.
var ErrInvalidImageMimeType = types.ErrInvalidImageMimeType

// Deprecated: Use types.ErrInvalidUrlLength instead.
var ErrInvalidUrlLength = types.ErrInvalidUrlLength

// Deprecated: Use types.ErrUrlPointingToLocalServer instead.
var ErrUrlPointingToLocalServer = types.ErrUrlPointingToLocalServer

// Deprecated: Use types.ErrEmptyUrl instead.
var ErrEmptyUrl = types.ErrEmptyUrl

// Deprecated: Use types.ErrInvalidTwitterUrlRegex instead.
var ErrInvalidTwitterUrlRegex = types.ErrInvalidTwitterUrlRegex

// Deprecated: Use types.ErrResponseTooLarge instead.
var ErrResponseTooLarge = types.ErrResponseTooLarge

// Deprecated: Use types.TextCharsLimit instead.
var TextCharsLimit = types.TextCharsLimit

// Deprecated: Use types.ImageExtensions instead.
var ImageExtensions = types.ImageExtensions

// Deprecated: Use types.WrapError instead.
var WrapError = types.WrapError

// Deprecated: Use types.IsValidEthereumAddress instead.
var IsValidEthereumAddress = types.IsValidEthereumAddress

// Deprecated: Use types.ReadPublicURL instead.
var ReadPublicURL = types.ReadPublicURL

// Deprecated: Use types.CheckIfValidTwitterURL instead.
var CheckIfValidTwitterURL = types.CheckIfValidTwitterURL

// Deprecated: Use types.CheckBasicURLValidation instead.
var CheckBasicURLValidation = types.CheckBasicURLValidation

// Deprecated: Use types.CheckIfUrlIsValid instead.
var CheckIfUrlIsValid = types.CheckIfUrlIsValid

// Deprecated: Use types.IsImageURL instead.
var IsImageURL = types.IsImageURL

// Deprecated: Use types.ValidateText instead.
var ValidateText = types.ValidateText

// Deprecated: Use types.ValidateRawGithubUrl instead.
var ValidateRawGithubUrl = types.ValidateRawGithubUrl

// TypedErr is deprecated. Use types.WrapError instead.
//
// Deprecated: This function is deprecated and will be removed in the future.
func TypedErr(e interface{}) error {
	return types.WrapError(e, nil)
}
