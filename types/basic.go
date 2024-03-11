package types

// This file defines internal types that have custom LogValues (for go's slog), to make debugging easier

import (
	"encoding/hex"
	"log/slog"
)

// make TaskResponseDigests print as hex encoded strings instead of a sequence of bytes
type Bytes32 [32]byte

func (m Bytes32) LogValue() slog.Value {
	return slog.StringValue(hex.EncodeToString(m[:]))
}

func (m *Bytes32) UnderlyingType() [32]byte {
	return *m
}
