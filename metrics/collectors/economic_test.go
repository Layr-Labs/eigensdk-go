package collectors

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/prometheus/client_golang/prometheus/testutil"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/Layr-Labs/eigensdk-go/chainio/mocks"
	"github.com/Layr-Labs/eigensdk-go/logging"
)

func TestEconomicCollector(t *testing.T) {
	operatorAddr := common.HexToAddress("0x0")
	strategyAddrs := []common.Address{
		common.HexToAddress("0x1"),
		common.HexToAddress("0x2"),
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	ethReader := mocks.NewMockELReader(mockCtrl)
	ethReader.EXPECT().OperatorIsFrozen(gomock.Any(), operatorAddr).Return(false, nil)
	ethReader.EXPECT().GetOperatorSharesInStrategy(gomock.Any(), operatorAddr, strategyAddrs[0]).Return(big.NewInt(1000), nil)
	ethReader.EXPECT().GetOperatorSharesInStrategy(gomock.Any(), operatorAddr, strategyAddrs[1]).Return(big.NewInt(2000), nil)

	logger := logging.NewNoopLogger()
	economicCollector := NewEconomicCollector(ethReader, "testavs", logger, operatorAddr, strategyAddrs)

	count := testutil.CollectAndCount(economicCollector, "eigen_slashing_status", "eigen_delegated_shares")
	// 1 for eigen_slashing_status, and 6 for eigen_delegated_shares (2 strategies * 3 unit labels each - wei,gwei,ether)
	assert.Equal(t, 7, count)
}
