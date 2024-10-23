package economic

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/prometheus/client_golang/prometheus/testutil"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/Layr-Labs/eigensdk-go/testutils"
	"github.com/Layr-Labs/eigensdk-go/types"
)

const registeredOpAddress = "0xb81b18c988bfc7d131fca985a9c531f325e98a2f"

type fakeELReader struct {
	registeredOperators map[common.Address]bool
}

func newFakeELReader() *fakeELReader {
	registeredOperators := make(map[common.Address]bool)
	registeredOperators[common.HexToAddress(registeredOpAddress)] = false
	return &fakeELReader{
		registeredOperators: registeredOperators,
	}
}

func (f *fakeELReader) OperatorIsFrozen(ctx context.Context, operatorAddr common.Address) (bool, error) {
	return f.registeredOperators[operatorAddr], nil
}

type fakeAVSRegistryReader struct {
	operatorId types.OperatorId
	stakes     map[types.QuorumNum]*big.Int
}

func (f *fakeAVSRegistryReader) GetOperatorId(opts *bind.CallOpts, operatorAddr common.Address) ([32]byte, error) {
	return f.operatorId, nil
}

func (f *fakeAVSRegistryReader) GetOperatorStakeInQuorumsOfOperatorAtCurrentBlock(
	opts *bind.CallOpts,
	operatorId types.OperatorId,
) (map[types.QuorumNum]*big.Int, error) {
	return f.stakes, nil
}

func newFakeAVSRegistryReader() *fakeAVSRegistryReader {
	operatorId := types.OperatorId{1}
	stakes := map[types.QuorumNum]*big.Int{
		0: big.NewInt(1000),
		1: big.NewInt(2000),
	}
	return &fakeAVSRegistryReader{
		operatorId: operatorId,
		stakes:     stakes,
	}
}

func TestEconomicCollector(t *testing.T) {
	operatorAddr := common.HexToAddress(registeredOpAddress)
	quorumNames := map[types.QuorumNum]string{
		0: "ethQuorum",
		1: "someOtherTokenQuorum",
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	elReader := newFakeELReader()
	avsRegistryReader := newFakeAVSRegistryReader()

	logger := testutils.GetTestLogger()
	economicCollector := NewCollector(elReader, avsRegistryReader, "testavs", logger, operatorAddr, quorumNames)

	count := testutil.CollectAndCount(economicCollector, "eigen_slashing_status", "eigen_registered_stakes")
	// 1 for eigen_slashing_status, and 2 for eigen_registered_stakes (1 per quorum)
	assert.Equal(t, 3, count)
}
