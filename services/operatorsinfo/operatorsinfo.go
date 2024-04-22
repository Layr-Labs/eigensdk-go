package operatorsinfo

import (
	"context"

	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/ethereum/go-ethereum/common"
)

// OperatorsInfoService is a service that indexes the BLSApkRegistry contract and provides a way to query for operator pubkeys.
// Currently BLSApkRegistry only stores the hash of the G1 and G2 pubkeys, so this service needs to listen to the
// event NewPubkeyRegistration(address indexed operator, BN254.G1Point pubkeyG1, BN254.G2Point pubkeyG2)
// and store the actual pubkeys, so that AVS aggregators can get the pubkeys of the operators registered with their AVS.
//
// TODO: having this service as a separate service (instead of merged with avsregistry service) is a vestige of the past when
// we had a separate blsPublicKeyCompendium shared contract between all AVSs. We should eventually merge this back into avsregistry.
type OperatorsInfoService interface {
	// GetOperatorInfo returns the info of the operator with the given address.
	// it returns operatorFound=false if the operator is not found.
	GetOperatorInfo(ctx context.Context, operator common.Address) (operatorInfo types.OperatorInfo, operatorFound bool)
}
