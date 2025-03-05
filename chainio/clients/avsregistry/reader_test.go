package avsregistry_test

import (
	"context"
	"fmt"
	"math/big"
	"testing"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/avsregistry"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/elcontracts"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"

	"github.com/Layr-Labs/eigensdk-go/testutils"
	"github.com/Layr-Labs/eigensdk-go/testutils/testclients"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/stretchr/testify/require"
)

func TestReaderMethods(t *testing.T) {
	clients, anvilHttpEndpoint := testclients.BuildTestClients(t)
	chainReader := clients.ReadClients.AvsRegistryChainReader
	contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilHttpEndpoint)
	strategy := contractAddrs.Erc20MockStrategy
	quorumNumber := types.QuorumNum(0)
	quorumNumbers := types.QuorumNums{0}

	operatorPrivateKeyHex := testutils.ANVIL_FIRST_PRIVATE_KEY

	config := avsregistry.Config{
		RegistryCoordinatorAddress:    contractAddrs.RegistryCoordinator,
		OperatorStateRetrieverAddress: contractAddrs.OperatorStateRetriever,
		ServiceManagerAddress:         contractAddrs.ServiceManager,
	}

	chainWriter, err := testclients.NewTestAvsRegistryWriterFromConfig(anvilHttpEndpoint, operatorPrivateKeyHex, config)
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

	t.Run(
		"get operators stake in quorums of operator at block returns error for non-registered operator",
		func(t *testing.T) {
			operatorAddress := common.Address{0x1}
			operatorId, err := chainReader.GetOperatorId(&bind.CallOpts{}, operatorAddress)
			require.NoError(t, err)

			_, _, err = chainReader.GetOperatorsStakeInQuorumsOfOperatorAtBlock(&bind.CallOpts{}, operatorId, 100)
			require.Error(t, err)
			require.Contains(t, err.Error(), "Failed to get operators state")
		})

	t.Run(
		"get single operator stake in quorums of operator at current block returns error for non-registered operator",
		func(t *testing.T) {
			operatorAddress := common.Address{0x1}
			operatorId, err := chainReader.GetOperatorId(&bind.CallOpts{}, operatorAddress)
			require.NoError(t, err)

			stakes, err := chainReader.GetOperatorStakeInQuorumsOfOperatorAtCurrentBlock(&bind.CallOpts{}, operatorId)
			require.NoError(t, err)
			require.Equal(t, 0, len(stakes))
		})

	t.Run("get check signatures indices returns error for non-registered operator", func(t *testing.T) {
		operatorAddress := common.Address{0x1}
		operatorId, err := chainReader.GetOperatorId(&bind.CallOpts{}, operatorAddress)
		require.NoError(t, err)

		_, err = chainReader.GetCheckSignaturesIndices(
			&bind.CallOpts{},
			100,
			quorumNumbers,
			[]types.OperatorId{operatorId},
		)
		require.Contains(t, err.Error(), "Failed to get check signatures indices")
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

	t.Run("query registration detail", func(t *testing.T) {
		operatorAddress := common.HexToAddress("0x1234567890123456789012345678901234567890")
		quorums, err := chainReader.QueryRegistrationDetail(&bind.CallOpts{}, operatorAddress)
		require.NoError(t, err)
		require.Equal(t, 1, len(quorums))
	})

	t.Run("is operator registered", func(t *testing.T) {
		operatorAddress := common.HexToAddress("0x1234567890123456789012345678901234567890")
		isRegistered, err := chainReader.IsOperatorRegistered(&bind.CallOpts{}, operatorAddress)
		require.NoError(t, err)
		require.False(t, isRegistered)
	})

	t.Run(
		"query existing registered operator pub keys", func(t *testing.T) {
			addresses, pubKeys, err := chainReader.QueryExistingRegisteredOperatorPubKeys(
				context.Background(),
				big.NewInt(0),
				nil,
				nil,
			)
			require.NoError(t, err)
			require.Equal(t, 0, len(pubKeys))
			require.Equal(t, 0, len(addresses))
		})

	t.Run(
		"query existing registered operator sockets", func(t *testing.T) {
			address_to_sockets, err := chainReader.QueryExistingRegisteredOperatorSockets(
				context.Background(),
				big.NewInt(0),
				nil,
				nil,
			)
			require.NoError(t, err)
			require.Equal(t, 0, len(address_to_sockets))
		})

	t.Run("Get weight of operator for quorum", func(t *testing.T) {
		operatorAddress := common.HexToAddress("0x1234567890123456789012345678901234567890")
		// A quorum registered in the old workflow should return 0
		weight, err := chainReader.WeightOfOperatorForQuorum(
			&bind.CallOpts{},
			0,
			operatorAddress,
		)
		require.NoError(t, err)
		require.Equal(t, int64(0), weight.Int64())
	})

	t.Run("Get strategy params length", func(t *testing.T) {
		quorumNumber := types.QuorumNum(0)
		length, err := chainReader.StrategyParamsLength(&bind.CallOpts{}, quorumNumber.UnderlyingType())
		require.NoError(t, err)
		require.Equal(t, int64(1), length.Int64())
	})

	t.Run("Get Stake History", func(t *testing.T) {
		operatorAddress := common.HexToAddress("0x1234567890123456789012345678901234567890")
		operatorId, err := chainReader.GetOperatorId(&bind.CallOpts{}, operatorAddress)
		require.NoError(t, err)
		stakeHistory, err := chainReader.GetStakeHistory(&bind.CallOpts{}, operatorId, quorumNumber.UnderlyingType())
		require.NoError(t, err)
		require.Equal(t, 0, len(stakeHistory))
	})

	t.Run("Get strategy params by index", func(t *testing.T) {
		params, err := chainReader.StrategyParamsByIndex(&bind.CallOpts{}, quorumNumber.UnderlyingType(), big.NewInt(0))
		require.NoError(t, err)
		require.Equal(t, strategy, params.Strategy)
	})

	operatorAddress := common.HexToAddress(testutils.ANVIL_FIRST_ADDRESS)

	//REGISTER OPERATOR
	otherKeyPair, err := bls.NewKeyPairFromString("0x01")
	require.NoError(t, err)
	request := elcontracts.RegistrationRequest{
		OperatorAddress: operatorAddress,
		AVSAddress:      contractAddrs.ServiceManager,
		OperatorSetIds:  []uint32{0},
		WaitForReceipt:  true,
		Socket:          "socket",
		BlsKeyPair:      otherKeyPair,
	}

	// Register operator
	elWriter := clients.ElChainWriter
	receipt, err := elWriter.SetAVSRegistrar(
		context.Background(),
		contractAddrs.ServiceManager,
		contractAddrs.RegistryCoordinator,
		true,
	)
	require.NoError(t, err)
	require.NotNil(t, receipt)

	receipt, err = elWriter.RegisterForOperatorSets(context.Background(), contractAddrs.RegistryCoordinator, request)
	require.NoError(t, err)
	require.NotNil(t, receipt)

	t.Run("get operators stake in quorums", func(t *testing.T) {
		blockNumber := uint32(receipt.BlockNumber.Uint64())
		operatorId, err := chainReader.GetOperatorId(&bind.CallOpts{}, operatorAddress)
		require.NoError(t, err)

		t.Run("get operators stake in quorums at block", func(t *testing.T) {
			stake, operators, err := chainReader.GetOperatorsStakeInQuorumsOfOperatorAtBlock(
				&bind.CallOpts{},
				operatorId,
				blockNumber,
			)
			require.NoError(t, err)
			require.Equal(t, 1, len(stake))
			require.Equal(t, 1, len(operators))
		})

		t.Run("get operators stake in quorums at current block", func(t *testing.T) {
			stake, operators, err := chainReader.GetOperatorsStakeInQuorumsOfOperatorAtCurrentBlock(
				&bind.CallOpts{},
				operatorId,
			)
			require.NoError(t, err)
			require.Equal(t, 1, len(stake))
			require.Equal(t, 1, len(operators))
		})

		t.Run("get operator stake in quorums at current block", func(t *testing.T) {
			stakeMap, err := chainReader.GetOperatorStakeInQuorumsOfOperatorAtCurrentBlock(&bind.CallOpts{}, operatorId)
			require.NoError(t, err)
			require.Equal(t, 1, len(stakeMap))
		})

		t.Run("get check signatures indices ", func(t *testing.T) {
			indices, err := chainReader.GetCheckSignaturesIndices(
				&bind.CallOpts{},
				blockNumber,
				quorumNumbers,
				[]types.OperatorId{operatorId},
			)
			require.NoError(t, err)
			require.NotNil(t, indices)
		})

		t.Run(
			"query existing registered operator pub keys", func(t *testing.T) {
				addresses, pubKeys, err := chainReader.QueryExistingRegisteredOperatorPubKeys(
					context.Background(),
					big.NewInt(0),
					nil,
					nil,
				)
				require.NoError(t, err)
				require.Equal(t, 1, len(pubKeys))
				require.Equal(t, 1, len(addresses))
			})
	})

	t.Run("Get stakeHistory length", func(t *testing.T) {
		operatorId, err := chainReader.GetOperatorId(&bind.CallOpts{}, operatorAddress)
		require.NoError(t, err)
		length, err := chainReader.GetStakeHistoryLength(&bind.CallOpts{}, operatorId, quorumNumber.UnderlyingType())
		require.NoError(t, err)
		require.Equal(t, int64(1), length.Int64())
	})

	t.Run("Get latest stake update", func(t *testing.T) {
		operatorId, err := chainReader.GetOperatorId(&bind.CallOpts{}, operatorAddress)
		require.NoError(t, err)

		stakeUpdate2, err := chainReader.GetLatestStakeUpdate(
			&bind.CallOpts{},
			operatorId,
			quorumNumber.UnderlyingType(),
		)
		fmt.Println("STAKE", stakeUpdate2.Stake)
		require.NoError(t, err)
		require.NotEqual(t, uint32(0), uint32(stakeUpdate2.Stake.Uint64()))
		require.Equal(t, uint32(0), stakeUpdate2.NextUpdateBlockNumber)
	})

	t.Run("Get stake update at index", func(t *testing.T) {
		operatorId, err := chainReader.GetOperatorId(&bind.CallOpts{}, operatorAddress)
		require.NoError(t, err)

		stakeUpdate2, err := chainReader.GetStakeUpdateAtIndex(
			&bind.CallOpts{},
			operatorId,
			quorumNumber.UnderlyingType(),
			big.NewInt(0),
		)
		require.NoError(t, err)
		require.NotEqual(t, uint32(0), uint32(stakeUpdate2.Stake.Uint64()))
		require.Equal(t, uint32(0), stakeUpdate2.NextUpdateBlockNumber)
	})

	t.Run("Get stake at block number", func(t *testing.T) {
		operatorId, err := chainReader.GetOperatorId(&bind.CallOpts{}, operatorAddress)
		require.NoError(t, err)

		stakeUpdate2, err := chainReader.GetLatestStakeUpdate(
			&bind.CallOpts{},
			operatorId,
			quorumNumber.UnderlyingType(),
		)
		require.NoError(t, err)

		updateBlockNumber := stakeUpdate2.UpdateBlockNumber
		stakeActual := stakeUpdate2.Stake

		stake, err := chainReader.GetStakeAtBlockNumber(
			&bind.CallOpts{},
			operatorId,
			quorumNumber.UnderlyingType(),
			updateBlockNumber,
		)
		require.NoError(t, err)
		require.Equal(t, stakeActual.Int64(), stake.Int64())
	})

	t.Run("Get stake updated index at block number", func(t *testing.T) {
		operatorId, err := chainReader.GetOperatorId(&bind.CallOpts{}, operatorAddress)
		require.NoError(t, err)

		stakeUpdate2, err := chainReader.GetLatestStakeUpdate(
			&bind.CallOpts{},
			operatorId,
			quorumNumber.UnderlyingType(),
		)
		require.NoError(t, err)

		updateBlockNumber := stakeUpdate2.UpdateBlockNumber

		stake, err := chainReader.GetStakeUpdateIndexAtBlockNumber(
			&bind.CallOpts{},
			operatorId,
			quorumNumber.UnderlyingType(),
			updateBlockNumber,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), stake)
	})

	t.Run("Get stake update index at block number", func(t *testing.T) {
		operatorId, err := chainReader.GetOperatorId(&bind.CallOpts{}, operatorAddress)
		require.NoError(t, err)

		stakeUpdate2, err := chainReader.GetLatestStakeUpdate(
			&bind.CallOpts{},
			operatorId,
			quorumNumber.UnderlyingType(),
		)
		require.NoError(t, err)

		updateBlockNumber := stakeUpdate2.UpdateBlockNumber
		stakeActual := stakeUpdate2.Stake

		stake, err := chainReader.GetStakeAtBlockNumberAndIndex(
			&bind.CallOpts{},
			operatorId,
			quorumNumber.UnderlyingType(),
			updateBlockNumber,
			big.NewInt(0),
		)
		require.NoError(t, err)
		require.Equal(t, stakeActual.Int64(), stake.Int64())
	})

	t.Run("Get total current stake", func(t *testing.T) {
		operatorId, err := chainReader.GetOperatorId(&bind.CallOpts{}, operatorAddress)
		require.NoError(t, err)

		stakeUpdate2, err := chainReader.GetLatestStakeUpdate(
			&bind.CallOpts{},
			operatorId,
			quorumNumber.UnderlyingType(),
		)
		require.NoError(t, err)

		stakeActual := stakeUpdate2.Stake

		currentTotalStake, err := chainReader.GetCurrentTotalStake(&bind.CallOpts{}, quorumNumber.UnderlyingType())
		require.NoError(t, err)
		require.Equal(t, stakeActual.Int64(), currentTotalStake.Int64())

	})

	t.Run("Get total stake history length", func(t *testing.T) {

		totalStakeHistoryLength, err := chainReader.GetTotalStakeHistoryLength(
			&bind.CallOpts{},
			quorumNumber.UnderlyingType(),
		)
		require.NoError(t, err)
		require.Equal(t, int64(2), totalStakeHistoryLength.Int64())
	})

	t.Run("Get total stake update at index", func(t *testing.T) {
		operatorId, err := chainReader.GetOperatorId(&bind.CallOpts{}, operatorAddress)
		require.NoError(t, err)

		stakeUpdate2, err := chainReader.GetLatestStakeUpdate(
			&bind.CallOpts{},
			operatorId,
			quorumNumber.UnderlyingType(),
		)
		require.NoError(t, err)

		stakeActual := stakeUpdate2.Stake

		totalStakeUpdateAtIndex, err := chainReader.GetTotalStakeUpdateAtIndex(
			&bind.CallOpts{},
			quorumNumber.UnderlyingType(),
			big.NewInt(1),
		)
		require.NoError(t, err)
		require.Equal(t, stakeActual.Int64(), totalStakeUpdateAtIndex.Stake.Int64())

	})

	t.Run("Get total stake update at block number from index", func(t *testing.T) {
		operatorId, err := chainReader.GetOperatorId(&bind.CallOpts{}, operatorAddress)
		require.NoError(t, err)

		stakeUpdate2, err := chainReader.GetLatestStakeUpdate(
			&bind.CallOpts{},
			operatorId,
			quorumNumber.UnderlyingType(),
		)
		require.NoError(t, err)

		stakeActual := stakeUpdate2.Stake
		updateBlockNumber := stakeUpdate2.UpdateBlockNumber

		totalStakeUpdateAtIndex, err := chainReader.GetTotalStakeAtBlockNumberFromIndex(
			&bind.CallOpts{},
			quorumNumber.UnderlyingType(),
			updateBlockNumber,
			big.NewInt(1),
		)
		require.NoError(t, err)
		require.Equal(t, stakeActual.Int64(), totalStakeUpdateAtIndex.Int64())
	})

	t.Run("Get total stake indices at block number", func(t *testing.T) {
		operatorId, err := chainReader.GetOperatorId(&bind.CallOpts{}, operatorAddress)
		require.NoError(t, err)

		stakeUpdate2, err := chainReader.GetLatestStakeUpdate(
			&bind.CallOpts{},
			operatorId,
			quorumNumber.UnderlyingType(),
		)
		require.NoError(t, err)

		updateBlockNumber := stakeUpdate2.UpdateBlockNumber

		totalStakeIndices, err := chainReader.GetTotalStakeIndicesAtBlockNumber(
			&bind.CallOpts{},
			quorumNumbers,
			updateBlockNumber,
		)
		require.NoError(t, err)
		require.Equal(t, 1, len(totalStakeIndices))
		require.Equal(t, uint32(1), totalStakeIndices[0])
	})

	t.Run("Get minimum stake for quorum", func(t *testing.T) {
		receipt, err := chainWriter.SetMinimumStakeForQuorum(
			context.Background(),
			quorumNumber.UnderlyingType(),
			big.NewInt(100),
			true,
		)
		require.NoError(t, err)
		require.NotNil(t, receipt)

		minimumStake, err := chainReader.GetMinimumStakeForQuorum(&bind.CallOpts{}, quorumNumber.UnderlyingType())
		require.NoError(t, err)
		require.Equal(t, big.NewInt(100), minimumStake)
	})

	t.Run("Get strategy params at index", func(t *testing.T) {
		params, err := chainReader.StrategyParamsByIndex(&bind.CallOpts{}, quorumNumber.UnderlyingType(), big.NewInt(0))
		require.NoError(t, err)
		require.Equal(t, strategy, params.Strategy)
		require.Equal(t, big.NewInt(1e18), params.Multiplier)
	})

	t.Run("Get strategy per quorum at index", func(t *testing.T) {
		retrievedStrat, err := chainReader.GetStrategyPerQuorumAtIndex(
			&bind.CallOpts{},
			quorumNumber.UnderlyingType(),
			big.NewInt(0),
		)
		require.NoError(t, err)
		require.Equal(t, strategy, retrievedStrat)
	})

	t.Run("Get restakeable strategies", func(t *testing.T) {
		retrievedStrat, err := chainReader.GetRestakeableStrategies(&bind.CallOpts{})
		require.NoError(t, err)
		require.Equal(t, []common.Address{strategy}, retrievedStrat)
	})

	t.Run("Get operator restaked strategies", func(t *testing.T) {
		retrievedStrat, err := chainReader.GetOperatorRestakedStrategies(&bind.CallOpts{}, operatorAddress)
		require.NoError(t, err)
		require.Equal(t, []common.Address{strategy}, retrievedStrat)
	})

	t.Run("Get restakeable strategies", func(t *testing.T) {
		retrievedStrat, err := chainReader.GetRestakeableStrategies(&bind.CallOpts{})
		require.NoError(t, err)
		require.Equal(t, []common.Address{strategy}, retrievedStrat)
	})

	t.Run("Get stake type per quorum", func(t *testing.T) {
		stakeType, err := chainReader.GetStakeTypePerQuorum(&bind.CallOpts{}, quorumNumber.UnderlyingType())
		require.NoError(t, err)
		require.Equal(t, uint8(0), stakeType)
	})

	t.Run("Get slashable stake look ahead per quorum", func(t *testing.T) {
		lookAhead, err := chainReader.GetSlashableStakeLookAheadPerQuorum(
			&bind.CallOpts{},
			quorumNumber.UnderlyingType(),
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), lookAhead)
	})

	t.Run("Get operatorPubkeyHash", func(t *testing.T) {
		operatorId, err := chainReader.GetOperatorId(&bind.CallOpts{}, operatorAddress)
		require.NoError(t, err)

		pubKeyHash, err := chainReader.GetOperatorIdFromOperatorAddress(&bind.CallOpts{}, operatorAddress)
		require.NoError(t, err)
		require.NotNil(t, pubKeyHash)
		require.Equal(t, operatorId, pubKeyHash)
	})

	t.Run("Get operatorPubkeyHash from operatorId", func(t *testing.T) {
		operatorId, err := chainReader.GetOperatorId(&bind.CallOpts{}, operatorAddress)
		require.NoError(t, err)

		gotOperatorAddress, err := chainReader.GetOperatorAddressFromOperatorId(&bind.CallOpts{}, operatorId)
		require.NoError(t, err)
		require.Equal(t, operatorAddress, gotOperatorAddress)
	})

	t.Run("Get operator bls pubkey", func(t *testing.T) {
		pubKey, err := chainReader.GetPubkeyFromOperatorAddress(
			&bind.CallOpts{},
			operatorAddress,
		)
		require.NoError(t, err)
		require.NotNil(t, pubKey)
	})

	t.Run("Get apk update", func(t *testing.T) {
		apkUpdateAfter, err := chainReader.GetApkUpdate(&bind.CallOpts{}, quorumNumber.UnderlyingType(), big.NewInt(0))
		require.NoError(t, err)
		require.NotNil(t, apkUpdateAfter)
		require.Greater(t, apkUpdateAfter.NextUpdateBlockNumber, apkUpdateAfter.UpdateBlockNumber)
	})

	t.Run("Get current apk", func(t *testing.T) {
		apk, err := chainReader.GetCurrentApk(&bind.CallOpts{}, quorumNumber.UnderlyingType())
		require.NoError(t, err)
		require.NotNil(t, apk)
	})
}

// Test that the reader returns an error when the configuration is invalid.
func TestReaderWithInvalidConfiguration(t *testing.T) {
	_, anvilHttpEndpoint := testclients.BuildTestClients(t)

	config := avsregistry.Config{}
	chainReader, err := testclients.NewTestAvsRegistryReaderFromConfig(anvilHttpEndpoint, config)
	require.NoError(t, err)

	quorumNumbers := types.QuorumNums{0}
	randomOperatorId := types.OperatorId{99}

	tests := []struct {
		name    string
		runFunc func() error
	}{
		{
			name: "get operator id",
			runFunc: func() error {
				_, err := chainReader.GetOperatorId(&bind.CallOpts{}, common.Address{})
				return err
			},
		},
		{
			name: "get operator from id",
			runFunc: func() error {
				_, err := chainReader.GetOperatorFromId(&bind.CallOpts{}, randomOperatorId)
				return err
			},
		},
		{
			name: "check if operator is registered",
			runFunc: func() error {
				_, err := chainReader.IsOperatorRegistered(&bind.CallOpts{}, common.Address{})
				return err
			},
		},
		{
			name: "get quorum state",
			runFunc: func() error {
				_, err := chainReader.GetQuorumCount(&bind.CallOpts{})
				return err
			},
		},
		{
			name: "get operator stake in quorums at current block",
			runFunc: func() error {
				_, err := chainReader.GetOperatorsStakeInQuorumsAtBlock(&bind.CallOpts{}, quorumNumbers, 100)
				return err
			},
		},
		{
			name: "get operator address in quorums at current block",
			runFunc: func() error {
				_, err := chainReader.GetOperatorAddrsInQuorumsAtCurrentBlock(&bind.CallOpts{}, quorumNumbers)
				return err
			},
		},
		{
			name: "get operators stake in quorums of operator at block",
			runFunc: func() error {
				_, _, err := chainReader.GetOperatorsStakeInQuorumsOfOperatorAtBlock(
					&bind.CallOpts{},
					randomOperatorId,
					100,
				)
				return err
			},
		},
		{
			name: "get single operator stake in quorums of operator at current block",
			runFunc: func() error {
				_, err := chainReader.GetOperatorStakeInQuorumsOfOperatorAtCurrentBlock(
					&bind.CallOpts{},
					randomOperatorId,
				)
				return err
			},
		}, {
			name: "check signatures indices",
			runFunc: func() error {
				_, err := chainReader.GetCheckSignaturesIndices(
					&bind.CallOpts{},
					100,
					quorumNumbers,
					[]types.OperatorId{randomOperatorId},
				)
				return err
			},
		},
		{
			name: "query registered operator sockets",
			runFunc: func() error {
				_, err := chainReader.QueryExistingRegisteredOperatorSockets(
					context.Background(),
					big.NewInt(0),
					nil,
					nil,
				)
				return err
			},
		},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%s with invalid config", tc.name), func(t *testing.T) {
			err := tc.runFunc()
			require.Error(t, err, "Expected error for %s", tc.name)
		})
	}
}
