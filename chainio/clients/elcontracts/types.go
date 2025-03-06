package elcontracts

import (
	"math/big"

	allocationmanager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/AllocationManager"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"

	"github.com/ethereum/go-ethereum/common"
)

type RegistrationType uint8

const (
	RegistrationTypeNormal RegistrationType = 0
	RegistrationTypeChurn  RegistrationType = 1
)

type OperatorSetStakes struct {
	OperatorSet     allocationmanager.OperatorSet
	Strategies      []common.Address
	Operators       []common.Address
	SlashableStakes [][]*big.Int
}

type PendingDeallocation struct {
	MagnitudeDiff        uint64
	CompletableTimestamp uint32
}

type AllocationInfo struct {
	CurrentMagnitude *big.Int
	PendingDiff      *big.Int
	EffectBlock      uint32
	OperatorSetId    uint32
	AvsAddress       common.Address
}

type DeregistrationRequest struct {
	AVSAddress     common.Address
	OperatorSetIds []uint32
	WaitForReceipt bool
}

type RegistrationRequest struct {
	OperatorAddress common.Address
	AVSAddress      common.Address
	OperatorSetIds  []uint32
	WaitForReceipt  bool
	BlsKeyPair      *bls.KeyPair
	Socket          string
}

// Parameters for removing an operator during churn.
// Used in RegistrationRequest to specify which operator to replace.
type OperatorKickParam struct {
	// Quorum from which to remove the operator.
	QuorumNumber uint8
	// Address of the operator to be removed.
	Operator common.Address
}

// Struct that bundles together a signature, a salt for uniqueness, and an expiration time for the signature. Used primarily for stack management.
type SignatureWithSaltAndExpiry struct {
	// the signature itself, formatted as a single bytes object
	Signature []byte
	// the salt used to generate the signature
	Salt [32]byte
	// the expiration timestamp (UTC) of the signature
	Expiry *big.Int
}

type RemovePermissionRequest struct {
	AccountAddress   common.Address
	AppointeeAddress common.Address
	Target           common.Address
	Selector         [4]byte
	WaitForReceipt   bool
}

type SetPermissionRequest struct {
	AccountAddress   common.Address
	AppointeeAddress common.Address
	Target           common.Address
	Selector         [4]byte
	WaitForReceipt   bool
}

type AcceptAdminRequest struct {
	AccountAddress common.Address
	WaitForReceipt bool
}

type AddPendingAdminRequest struct {
	AccountAddress common.Address
	AdminAddress   common.Address
	WaitForReceipt bool
}

type RemoveAdminRequest struct {
	AccountAddress common.Address
	AdminAddress   common.Address
	WaitForReceipt bool
}

type RemovePendingAdminRequest struct {
	AccountAddress common.Address
	AdminAddress   common.Address
	WaitForReceipt bool
}
