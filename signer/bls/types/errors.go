package types

import "errors"

var (
	ErrInvalidMessageLength = errors.New("invalid message length. must be 32 bytes")
	ErrInvalidSignerType    = errors.New("invalid signer type")
)
