package bls

// TODO: move these somewhere more appropriate.
//
//	these are avs types, nothing to do with bls
//
// also there's already an OperatorId in types, but we can't use it here because it would form a circular import
// since types using bls.Signature.. we prob just need to move all of these types to types/avs.go or something
type OperatorId = [32]byte
type OperatorIndex uint

// Security and Quorum Parameters

// QuorumID is a unique identifier for a quorum; initially EigenDA will support up to 256 quorums
type QuorumID uint8

// SecurityParam contains the quorum ID and the adversary threshold for the quorum;
type SecurityParam struct {
	QuorumID QuorumID
	// AdversaryThreshold is the maximum amount of stake that can be controlled by an adversary in the quorum as a
	// percentage of the total stake in the quorum
	AdversaryThreshold uint8
}

// QuorumParam contains the quorum ID and the quorum threshold for the quorum
type QuorumParam struct {
	QuorumID QuorumID
	// QuorumThreshold is the amount of stake that must sign a message for it to be considered valid as a percentage of
	// the total stake in the quorum
	QuorumThreshold uint8
}
