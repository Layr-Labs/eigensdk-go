package operatorsinfo

import (
	"context"
	"testing"

	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/shurcooL/graphql"
	"github.com/stretchr/testify/assert"
)

type mockGraphQLQuerier struct {
	QueryFn func(ctx context.Context, q any, variables map[string]any) error
}

func (m mockGraphQLQuerier) Query(ctx context.Context, q any, variables map[string]any) error {
	return m.QueryFn(ctx, q, variables)
}

func TestIndexedChainState_GetIndexedOperatorState(t *testing.T) {
	logger := logging.NewNoopLogger()

	operatorAddress := common.Address{0x1}

	operatorsQueryCalled := false
	querier := &mockGraphQLQuerier{}
	querier.QueryFn = func(ctx context.Context, q any, variables map[string]any) error {
		switch res := q.(type) {
		case *QueryOperatorByAddressGql:
			if operatorsQueryCalled {
				return nil
			}
			res.Operator = IndexedOperatorInfoGql{
				Address:    graphql.String(operatorAddress.String()),
				PubkeyG1_X: "3336192159512049190945679273141887248666932624338963482128432381981287252980",
				PubkeyG1_Y: "15195175002875833468883745675063986308012687914999552116603423331534089122704",
				PubkeyG2_X: []graphql.String{
					"21597023645215426396093421944506635812143308313031252511177204078669540440732",
					"11405255666568400552575831267661419473985517916677491029848981743882451844775",
				},
				PubkeyG2_Y: []graphql.String{
					"9416989242565286095121881312760798075882411191579108217086927390793923664442",
					"13612061731370453436662267863740141021994163834412349567410746669651828926551",
				},
				SocketUpdates: []SocketUpdates{{Socket: "localhost:32006;32007"}},
			}
			operatorsQueryCalled = true
			return nil
		default:
			return nil
		}
	}

	cs := NewOperatorsInfoServiceSubgraph(context.Background(), querier, logger)
	operatorPubkeys, err := cs.GetOperatorInfo(context.Background(), operatorAddress)
	assert.True(t, err)
	assert.Equal(t, operatorPubkeys.Socket, types.Socket("localhost:32006;32007"))
}
