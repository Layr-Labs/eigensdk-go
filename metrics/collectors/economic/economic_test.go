package economic

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/prometheus/client_golang/prometheus/testutil"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/Layr-Labs/eigensdk-go/chainio/mocks"
	blsoperatorstateretrievar "github.com/Layr-Labs/eigensdk-go/contracts/bindings/BLSOperatorStateRetriever"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/types"
)

func TestEconomicCollector(t *testing.T) {
	operatorAddr := common.HexToAddress("0x0")
	operatorId := types.OperatorId{0}
	quorumNames := map[types.QuorumNum]string{
		0: "ethQuorum",
		1: "someOtherTokenQuorum",
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	ethReader := mocks.NewMockELReader(mockCtrl)
	ethReader.EXPECT().OperatorIsFrozen(gomock.Any(), operatorAddr).Return(false, nil)
	// ethReader.EXPECT().GetOperatorSharesInStrategy(gomock.Any(), operatorAddr, strategyAddrs[0]).Return(big.NewInt(1000), nil)
	// ethReader.EXPECT().GetOperatorSharesInStrategy(gomock.Any(), operatorAddr, strategyAddrs[1]).Return(big.NewInt(2000), nil)

	avsRegistryReader := mocks.NewMockAvsRegistryReader(mockCtrl)
	avsRegistryReader.EXPECT().GetOperatorsStakeInQuorumsOfOperatorAtCurrentBlock(gomock.Any(), gomock.Any()).Return(
		[]types.QuorumNum{0, 1},
		[][]blsoperatorstateretrievar.BLSOperatorStateRetrieverOperator{
			{
				{
					OperatorId: operatorId,
					Stake:      big.NewInt(1000),
				},
			},
			{
				{
					OperatorId: operatorId,
					Stake:      big.NewInt(2000),
				},
			},
		},
		nil,
	)
	avsRegistryReader.EXPECT().GetOperatorId(gomock.Any(), operatorAddr).Return(operatorId, nil)

	logger := logging.NewNoopLogger()
	economicCollector := NewCollector(ethReader, avsRegistryReader, "testavs", logger, operatorAddr, quorumNames)

	count := testutil.CollectAndCount(economicCollector, "eigen_slashing_status", "eigen_registered_stakes")
	// 1 for eigen_slashing_status, and 2 for eigen_registered_stakes (1 per quorum)
	assert.Equal(t, 3, count)
}
