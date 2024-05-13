package operatorsinfo

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/consensys/gnark-crypto/ecc/bn254"
	"github.com/ethereum/go-ethereum/common"
	"github.com/shurcooL/graphql"
)

type (
	QueryOperatorByAddressGql struct {
		Operator IndexedOperatorInfoGql `graphql:"operator(address: $address)"`
	}
	OperatorsInfoServiceSubgraph struct {
		logger logging.Logger
		client *graphql.Client
	}
	SocketUpdates struct {
		Socket graphql.String
	}
	IndexedOperatorInfoGql struct {
		Address    graphql.String
		PubkeyG1_X graphql.String   `graphql:"pubkeyG1_X"`
		PubkeyG1_Y graphql.String   `graphql:"pubkeyG1_Y"`
		PubkeyG2_X []graphql.String `graphql:"pubkeyG2_X"`
		PubkeyG2_Y []graphql.String `graphql:"pubkeyG2_Y"`
		// Socket is the socket address of the operator, in the form "host:port"
		SocketUpdates []SocketUpdates `graphql:"socketUpdates(first: 1, orderBy: blockNumber, orderDirection: desc)"`
	}
	IndexedOperatorInfo struct {
		// PubKeyG1 and PubKeyG2 are the public keys of the operator, which are retrieved from the EigenDAPubKeyCompendium smart contract
		PubkeyG1 *G1Point
		PubkeyG2 *G2Point
		// Socket is the socket address of the operator, in the form "host:port"
		Socket string
	}
	G2Point struct {
		*bn254.G2Affine
	}
	G1Point struct {
		*bn254.G1Affine
	}
)

var _ OperatorsInfoService = (*OperatorsInfoServiceSubgraph)(nil)

// NewOperatorsInfoServiceSubgraph constructs a OperatorsInfoServiceSubgraph and starts it in a goroutine.
// It takes a context as argument because the "backfilling" of the database is done inside this constructor,
// so we wait for all past NewPubkeyRegistration/OperatorSocketUpdate events to be queried and the db to be filled before returning the service.
// The constructor is thus following a RAII-like pattern, of initializing the serving during construction.
// Using a separate initialize() function might lead to some users forgetting to call it and the service not behaving properly.
func NewOperatorsInfoServiceSubgraph(
	ctx context.Context,
	url string,
	logger logging.Logger,
) *OperatorsInfoServiceSubgraph {
	client := graphql.NewClient(url, nil)

	return &OperatorsInfoServiceSubgraph{
		logger: logger,
		client: client,
	}
}

// TODO(samlaf): we might want to also add an async version of this method that returns a channel of operator pubkeys?
func (ops *OperatorsInfoServiceSubgraph) GetOperatorInfo(ctx context.Context, operator common.Address) (operatorPubkeys types.OperatorInfo, operatorFound bool) {
	operatorInfo, err := ops.getIndexedOperatorInfoByOperatorId(ctx, operator)
	if err != nil {
		return types.OperatorInfo{}, false
	}
	return *operatorInfo, true
}

func (ops *OperatorsInfoServiceSubgraph) getIndexedOperatorInfoByOperatorId(ctx context.Context, operator common.Address) (*types.OperatorInfo, error) {
	var (
		query     QueryOperatorByAddressGql
		variables = map[string]any{
			"id": graphql.String(fmt.Sprintf("0x%s", hex.EncodeToString(operator[:]))),
		}
	)
	err := ops.client.Query(ctx, &query, variables)
	if err != nil {
		ops.logger.Error("Error requesting info for operator", "err", err, "operator", hex.EncodeToString(operator[:]))
		return nil, err
	}

	return convertIndexedOperatorInfoGqlToOperatorInfo(&query.Operator)
}

func convertIndexedOperatorInfoGqlToOperatorInfo(operator *IndexedOperatorInfoGql) (*types.OperatorInfo, error) {

	if len(operator.SocketUpdates) == 0 {
		return nil, errors.New("no socket found for operator")
	}

	pubkeyG1 := new(bn254.G1Affine)
	_, err := pubkeyG1.X.SetString(string(operator.PubkeyG1_X))
	if err != nil {
		return nil, err
	}
	_, err = pubkeyG1.Y.SetString(string(operator.PubkeyG1_Y))
	if err != nil {
		return nil, err
	}

	pubkeyG2 := new(bn254.G2Affine)
	_, err = pubkeyG2.X.A1.SetString(string(operator.PubkeyG2_X[0]))
	if err != nil {
		return nil, err
	}
	_, err = pubkeyG2.X.A0.SetString(string(operator.PubkeyG2_X[1]))
	if err != nil {
		return nil, err
	}
	_, err = pubkeyG2.Y.A1.SetString(string(operator.PubkeyG2_Y[0]))
	if err != nil {
		return nil, err
	}
	_, err = pubkeyG2.Y.A0.SetString(string(operator.PubkeyG2_Y[1]))
	if err != nil {
		return nil, err
	}

	return &types.OperatorInfo{
		Socket: types.Socket(string(operator.SocketUpdates[0].Socket)),
		Pubkeys: types.OperatorPubkeys{
			G1Pubkey: &bls.G1Point{G1Affine: pubkeyG1},
			G2Pubkey: &bls.G2Point{G2Affine: pubkeyG2},
		},
	}, nil

}
