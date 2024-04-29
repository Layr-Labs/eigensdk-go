package utils

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrInvalidUrl            = errors.New("invalid url")
	ErrInvalidGithubRawUrl   = errors.New("invalid github raw url")
	ErrInvalidText           = fmt.Errorf("invalid text format, doesn't conform to regex %s", TextRegex)
	ErrTextTooLong           = errors.New("text should be less than 500 characters")
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

func TypedErr(e interface{}) error {
	switch t := e.(type) {
	case error:
		return t
	case string:
		return errors.New(t)
	default:
		return nil
	}
}

func WrapError(mainErr interface{}, subErr interface{}) error {
	var main, sub error
	main = TypedErr(mainErr)
	sub = TypedErr(subErr)
	// Some times the wrap will wrap a nil error
	if main == nil && sub == nil {
		return nil
	}

	if main == nil && sub != nil {
		return sub
	}

	if main != nil && sub == nil {
		return main
	}

	return main
}
