package avsregistry

import (
	"context"
	"log/slog"
	"os"
	"testing"

	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/testutils"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
)

func TestReaderMethods(t *testing.T) {
	anvilStateFileName := "contracts-deployed-anvil-state.json"
	anvilC, err := testutils.StartAnvilContainer(anvilStateFileName)
	require.NoError(t, err)

	anvilHttpEndpoint, err := anvilC.Endpoint(context.Background(), "http")
	require.NoError(t, err)

	contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilHttpEndpoint)
	ethHttpClient, err := ethclient.Dial(anvilHttpEndpoint)
	require.NoError(t, err)

	quorumNumbers := types.QuorumNums{0}

	logger := logging.NewTextSLogger(os.Stdout, &logging.SLoggerOptions{Level: slog.LevelDebug})

	chainReader, err := NewReaderFromConfig(Config{
		RegistryCoordinatorAddress:    contractAddrs.RegistryCoordinator,
		OperatorStateRetrieverAddress: contractAddrs.OperatorStateRetriever,
	}, ethHttpClient, logger)
	require.NoError(t, err)

	t.Run("get quorum state", func(t *testing.T) {
		count, err := chainReader.GetQuorumCount(&bind.CallOpts{})
		require.NoError(t, err)
		require.NotNil(t, count)
	})

	t.Run("get operator stake in quorums at current block", func(t *testing.T) {
		stake, err := chainReader.GetOperatorsStakeInQuorumsAtCurrentBlock(&bind.CallOpts{}, quorumNumbers)
		require.NoError(t, err)
		require.NotNil(t, stake)
	})

	t.Run("get operator stake in quorums at block", func(t *testing.T) {
		stake, err := chainReader.GetOperatorsStakeInQuorumsAtBlock(&bind.CallOpts{}, quorumNumbers, 100)
		require.NoError(t, err)
		require.NotNil(t, stake)
	})

	t.Run("get operator address in quorums at current block", func(t *testing.T) {
		addresses, err := chainReader.GetOperatorAddrsInQuorumsAtCurrentBlock(&bind.CallOpts{}, quorumNumbers)
		require.NoError(t, err)
		require.NotNil(t, addresses)
	})

	t.Run("get operators stake in quorums of operator at block returns error for non-registered operator", func(t *testing.T) {
		operatorAddress := common.Address{0x1}
		operatorId, err := chainReader.GetOperatorId(&bind.CallOpts{}, operatorAddress)
		require.NoError(t, err)

		_, _, err = chainReader.GetOperatorsStakeInQuorumsOfOperatorAtBlock(&bind.CallOpts{}, operatorId, 100)
		require.Error(t, err)
		require.Contains(t, err.Error(), "Failed to get operators state")
	})

	t.Run("get single operator stake in quorums of operator at current block returns error for non-registered operator", func(t *testing.T) {
		operatorAddress := common.Address{0x1}
		operatorId, err := chainReader.GetOperatorId(&bind.CallOpts{}, operatorAddress)
		require.NoError(t, err)

		stakes, err := chainReader.GetOperatorStakeInQuorumsOfOperatorAtCurrentBlock(&bind.CallOpts{}, operatorId)
		require.NoError(t, err)
		require.Equal(t, 0, len(stakes))
	})

	t.Run("get check signatures indices", func(t *testing.T) {
		operatorAddress := common.Address{0x1}
		operatorId, err := chainReader.GetOperatorId(&bind.CallOpts{}, operatorAddress)
		require.NoError(t, err)

		checkSignaturesIndices, err := chainReader.GetCheckSignaturesIndices(&bind.CallOpts{}, 100, quorumNumbers, []types.OperatorId{operatorId})
		require.NoError(t, err)
		require.Equal(t, 0, len(checkSignaturesIndices.NonSignerQuorumBitmapIndices))
		require.Equal(t, 0, len(checkSignaturesIndices.NonSignerStakeIndices))
		require.Equal(t, 0, len(checkSignaturesIndices.QuorumApkIndices))
		require.Equal(t, 0, len(checkSignaturesIndices.TotalStakeIndices))
	})

	t.Run("get operator id", func(t *testing.T) {
		operatorAddress := common.Address{0x1}
		operatorId, err := chainReader.GetOperatorId(&bind.CallOpts{}, operatorAddress)
		require.NoError(t, err)
		require.NotNil(t, operatorId)
	})

	t.Run("get operator from id returns zero address for non-registered operator", func(t *testing.T) {
		operatorAddress := common.HexToAddress("0x1234567890123456789012345678901234567890")
		operatorId, err := chainReader.GetOperatorId(&bind.CallOpts{}, operatorAddress)
		require.NoError(t, err)

		retrievedAddress, err := chainReader.GetOperatorFromId(&bind.CallOpts{}, operatorId)
		require.NoError(t, err)
		require.Equal(t, retrievedAddress, common.Address{0x0})
	})
}
