package operatorpubkeys

import (
	"context"

	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/ethereum/go-ethereum/common"
)

// OperatorPubkeysService is a service that indexes the BLSApkRegistry contract and provides a way to query for operator pubkeys.
// Currently BLSApkRegistry only stores the hash of the G1 and G2 pubkeys, so this service needs to listen to the
// event NewPubkeyRegistration(address indexed operator, BN254.G1Point pubkeyG1, BN254.G2Point pubkeyG2)
// and store the actual pubkeys, so that AVS aggregators can get the pubkeys of the operators registered with their AVS.
type OperatorPubkeysService interface {
	// GetOperatorPubkeys returns the pubkeys of the operator with the given address.
	// it returns false is the operator is not found.
	GetOperatorPubkeys(ctx context.Context, operator common.Address) (operatorPubkeys types.OperatorPubkeys, operatorFound bool)
}
