package utils

import (
	"errors"
	"strings"
)

var (
	ErrInvalidUrl            = errors.New("invalid url")
	ErrInvalidGithubRawUrl   = errors.New("invalid github raw url")
	ErrInvalidText           = errors.New("text is invalid")
	ErrTextTooLong           = errors.New("text should be less than 200 characters")
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
)
