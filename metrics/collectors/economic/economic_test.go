package economic

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/prometheus/client_golang/prometheus/testutil"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	chainioMocks "github.com/Layr-Labs/eigensdk-go/chainio/mocks"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/types"
)

const registeredOpAddress = "\"0xb81b18c988bfc7d131fca985a9c531f325e98a2f\""

type FakeELReader struct {
	registeredOperators map[common.Address]bool
}

func NewFakeELReader() *FakeELReader {
	registeredOperators := make(map[common.Address]bool)
	registeredOperators[common.HexToAddress(registeredOpAddress)] = false
	return &FakeELReader{
		registeredOperators: registeredOperators,
	}
}

func (f *FakeELReader) OperatorIsFrozen(opts *bind.CallOpts, operatorAddr common.Address) (bool, error) {
	return f.registeredOperators[operatorAddr], nil
}

func TestEconomicCollector(t *testing.T) {
	operatorAddr := common.HexToAddress(registeredOpAddress)
	operatorId := types.OperatorId{1}
	quorumNames := map[types.QuorumNum]string{
		0: "ethQuorum",
		1: "someOtherTokenQuorum",
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	elReader := NewFakeELReader()
	avsRegistryReader := chainioMocks.NewMockAVSReader(mockCtrl)
	avsRegistryReader.EXPECT().GetOperatorId(gomock.Any(), operatorAddr).Return(operatorId, nil)
	avsRegistryReader.EXPECT().GetOperatorStakeInQuorumsOfOperatorAtCurrentBlock(gomock.Any(), gomock.Any()).Return(
		map[types.QuorumNum]*big.Int{
			0: big.NewInt(1000),
			1: big.NewInt(2000),
		},
		nil,
	)

	logger := logging.NewNoopLogger()
	economicCollector := NewCollector(elReader, avsRegistryReader, "testavs", logger, operatorAddr, quorumNames)

	count := testutil.CollectAndCount(economicCollector, "eigen_slashing_status", "eigen_registered_stakes")
	// 1 for eigen_slashing_status, and 2 for eigen_registered_stakes (1 per quorum)
	assert.Equal(t, 3, count)
}
