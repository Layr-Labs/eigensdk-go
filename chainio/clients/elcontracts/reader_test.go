package elcontracts_test

import (
	"context"
	"math/big"
	"os"
	"testing"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/elcontracts"
	allocationmanager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/AllocationManager"
	erc20 "github.com/Layr-Labs/eigensdk-go/contracts/bindings/IERC20"
	rewardscoordinator "github.com/Layr-Labs/eigensdk-go/contracts/bindings/RewardsCoordinator"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/testutils"
	"github.com/Layr-Labs/eigensdk-go/testutils/testclients"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestChainReader(t *testing.T) {
	clients, anvilHttpEndpoint := testclients.BuildTestClients(t)
	ctx := context.Background()

	contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilHttpEndpoint)
	operator := types.Operator{
		Address: testutils.ANVIL_FIRST_ADDRESS,
	}

	rewardsCoordinator, err := rewardscoordinator.NewContractRewardsCoordinator(
		contractAddrs.RewardsCoordinator,
		clients.EthHttpClient,
	)
	require.NoError(t, err)

	t.Run("is operator registered", func(t *testing.T) {
		isOperator, err := clients.ElChainReader.IsOperatorRegistered(ctx, operator)
		assert.NoError(t, err)
		assert.Equal(t, isOperator, true)
	})

	t.Run("get operator details", func(t *testing.T) {
		operatorDetails, err := clients.ElChainReader.GetOperatorDetails(ctx, operator)
		assert.NoError(t, err)
		assert.NotNil(t, operatorDetails)
		assert.Equal(t, operator.Address, operatorDetails.Address)
	})

	t.Run("get strategy and underlying token", func(t *testing.T) {
		strategyAddr := contractAddrs.Erc20MockStrategy
		strategy, underlyingTokenAddr, err := clients.ElChainReader.GetStrategyAndUnderlyingToken(
			ctx,
			strategyAddr,
		)
		assert.NoError(t, err)
		assert.NotNil(t, strategy)
		assert.NotEqual(t, common.Address{}, underlyingTokenAddr)

		erc20Token, err := erc20.NewContractIERC20(underlyingTokenAddr, clients.EthHttpClient)
		assert.NoError(t, err)

		tokenName, err := erc20Token.Name(&bind.CallOpts{})
		assert.NoError(t, err)
		assert.NotEmpty(t, tokenName)
	})

	t.Run("get strategy and underlying ERC20 token", func(t *testing.T) {
		strategyAddr := contractAddrs.Erc20MockStrategy
		strategy, contractUnderlyingToken, underlyingTokenAddr, err := clients.ElChainReader.GetStrategyAndUnderlyingERC20Token(
			ctx,
			strategyAddr,
		)
		assert.NoError(t, err)
		assert.NotNil(t, strategy)
		assert.NotEqual(t, common.Address{}, underlyingTokenAddr)
		assert.NotNil(t, contractUnderlyingToken)

		tokenName, err := contractUnderlyingToken.Name(&bind.CallOpts{})
		assert.NoError(t, err)
		assert.NotEmpty(t, tokenName)
	})

	t.Run("get operator shares in strategy", func(t *testing.T) {
		shares, err := clients.ElChainReader.GetOperatorSharesInStrategy(
			ctx,
			common.HexToAddress(operator.Address),
			contractAddrs.Erc20MockStrategy,
		)
		assert.NoError(t, err)
		assert.NotZero(t, shares)
	})

	t.Run("calculate delegation approval digest hash", func(t *testing.T) {
		staker := common.Address{0x0}
		delegationApprover := common.Address{0x0}
		approverSalt := [32]byte{}
		expiry := big.NewInt(0)
		digest, err := clients.ElChainReader.CalculateDelegationApprovalDigestHash(
			ctx,
			staker,
			common.HexToAddress(operator.Address),
			delegationApprover,
			approverSalt,
			expiry,
		)
		assert.NoError(t, err)
		assert.NotEmpty(t, digest)
	})

	t.Run("calculate operator AVS registration digest hash", func(t *testing.T) {
		avs := common.Address{0x0}
		salt := [32]byte{}
		expiry := big.NewInt(0)
		digest, err := clients.ElChainReader.CalculateOperatorAVSRegistrationDigestHash(
			ctx,
			common.HexToAddress(operator.Address),
			avs,
			salt,
			expiry,
		)
		assert.NoError(t, err)
		assert.NotEmpty(t, digest)
	})

	t.Run("get staker shares", func(t *testing.T) {
		strategies, shares, err := clients.ElChainReader.GetStakerShares(
			ctx,
			common.HexToAddress(operator.Address),
		)
		assert.NotZero(t, len(strategies), "Strategies has at least one element")
		assert.NotZero(t, len(shares), "Shares has at least one element")
		assert.Equal(t, len(strategies), len(shares), "Strategies has the same ammount of elements as shares")
		assert.NoError(t, err)
	})

	t.Run("get delegated operator", func(t *testing.T) {
		blockNumber := big.NewInt(0)
		address, err := clients.ElChainReader.GetDelegatedOperator(
			ctx,
			common.HexToAddress(operator.Address),
			blockNumber,
		)

		assert.NoError(t, err)
		// The delegated operator of an operator is the operator itself
		assert.Equal(t, address.String(), operator.Address)
	})

	t.Run("GetOperatorShares", func(t *testing.T) {
		strategyAddr := contractAddrs.Erc20MockStrategy
		strategies := []common.Address{strategyAddr}
		shares, err := clients.ElChainReader.GetOperatorShares(
			ctx,
			common.HexToAddress(operator.Address),
			strategies,
		)
		assert.NoError(t, err)
		assert.Len(t, shares, 1)

		// with n strategies, response's list length is n
		strategies = []common.Address{strategyAddr, strategyAddr, strategyAddr}
		shares, err = clients.ElChainReader.GetOperatorShares(
			ctx,
			common.HexToAddress(operator.Address),
			strategies,
		)
		assert.NoError(t, err)
		assert.Len(t, shares, 3)

		// We could test modify the shares and verify the diff is the expected
	})

	t.Run("GetOperatorsShares", func(t *testing.T) {
		operatorAddr := common.HexToAddress(operator.Address)
		operators := []common.Address{operatorAddr}
		strategyAddr := contractAddrs.Erc20MockStrategy
		strategies := []common.Address{strategyAddr}
		shares, err := clients.ElChainReader.GetOperatorsShares(
			ctx,
			operators,
			strategies,
		)
		assert.NoError(t, err)
		assert.Len(t, shares, 1)

		// with n strategies, response's list length is [1][n]
		mult_strategies := []common.Address{strategyAddr, strategyAddr, strategyAddr}
		shares, err = clients.ElChainReader.GetOperatorsShares(
			ctx,
			operators,
			mult_strategies,
		)
		assert.NoError(t, err)
		assert.Len(t, shares, 1)
		assert.Len(t, shares[0], 3)

		// with n strategies, response's list length is [n][1]
		mult_operators := []common.Address{operatorAddr, operatorAddr, operatorAddr}
		shares, err = clients.ElChainReader.GetOperatorsShares(
			ctx,
			mult_operators,
			strategies,
		)
		assert.NoError(t, err)
		assert.Len(t, shares, 3)
		assert.Len(t, shares[0], 1)

		// with n strategies and n operators, response's list length is [n][n]
		shares, err = clients.ElChainReader.GetOperatorsShares(
			ctx,
			mult_operators,
			mult_strategies,
		)
		assert.NoError(t, err)
		assert.Len(t, shares, 3)
		assert.Len(t, shares[2], 3)
	})

	t.Run("Get delegationApproverSaltIsSpent", func(t *testing.T) {
		approverSalt := [32]byte{}
		isSpent, err := clients.ElChainReader.GetDelegationApproverSaltIsSpent(
			ctx,
			common.HexToAddress(operator.Address),
			approverSalt,
		)
		assert.NoError(t, err)
		assert.False(t, isSpent)
	})

	t.Run("Get pending withdrawal status", func(t *testing.T) {
		withdrawalRoot := [32]byte{}
		pending, err := clients.ElChainReader.GetPendingWithdrawalStatus(
			ctx,
			withdrawalRoot,
		)
		assert.NoError(t, err)
		assert.False(t, pending)
	})

	t.Run("Get cumulative withdrawals queued", func(t *testing.T) {
		staker := common.HexToAddress(operator.Address)
		cumulative, err := clients.ElChainReader.GetCumulativeWithdrawalsQueued(
			ctx,
			staker,
		)
		assert.NoError(t, err)
		assert.Zero(t, cumulative.Cmp(big.NewInt(0)))
	})

	t.Run("Get deallocation delay", func(t *testing.T) {
		delay, err := clients.ElChainReader.GetDeallocationDelay(
			ctx,
		)
		assert.NoError(t, err)
		assert.NotZero(t, delay)
	})

	t.Run("Get allocation configuration delay", func(t *testing.T) {
		delay, err := clients.ElChainReader.GetAllocationConfigurationDelay(
			ctx,
		)
		assert.NoError(t, err)
		assert.NotZero(t, delay)
	})

	t.Run("get operator avs", func(t *testing.T) {
		split, err := clients.ElChainReader.GetOperatorAVSSplit(
			context.Background(),
			common.HexToAddress(operator.Address),
			common.MaxAddress,
		)
		require.NoError(t, err)
		require.NotZero(t, split)

		split, err = clients.ElChainReader.GetOperatorPISplit(
			context.Background(),
			common.HexToAddress(operator.Address),
		)
		require.NoError(t, err)
		require.NotZero(t, split)
	})

	t.Run("get operator set split", func(t *testing.T) {
		testAddr := common.HexToAddress(testutils.ANVIL_FIRST_ADDRESS)
		operatorSetId := uint32(1)
		operatorSet := rewardscoordinator.OperatorSet{
			Avs: testAddr,
			Id:  operatorSetId,
		}
		split, err := clients.ElChainReader.GetOperatorSetSplit(
			context.Background(),
			common.HexToAddress(operator.Address),
			operatorSet,
		)
		require.NoError(t, err)
		require.NotZero(t, split)
	})

	// Get the interval in seconds at which the calculation for rewards distribution is done.
	t.Run("get calculation interval seconds", func(t *testing.T) {
		interval, err := clients.ElChainReader.GetCalculationIntervalSeconds(context.Background())
		require.NoError(t, err)
		// currently this is configured to zero but may be configured to 1 week in a future release, based on this
		// comment:
		// https://github.com/Layr-Labs/eigenlayer-contracts/blob/441339cbd570ad0d650a9c11bea9eed7f70a490d/src/contracts/core/RewardsCoordinatorStorage.sol#L64
		require.NotZero(t, interval)

		intervalActual, err := rewardsCoordinator.CALCULATIONINTERVALSECONDS(&bind.CallOpts{})
		require.NoError(t, err)
		require.Equal(t, interval, intervalActual)
	})

	t.Run("get max duration  seconds", func(t *testing.T) {
		duration, err := clients.ElChainReader.GetMaxRewardsDuration(context.Background())
		require.NoError(t, err)
		require.NotZero(t, duration)

		durationActual, err := rewardsCoordinator.MAXREWARDSDURATION(&bind.CallOpts{})
		require.NoError(t, err)
		require.Equal(t, duration, durationActual)
	})

	t.Run("get max retroactive length", func(t *testing.T) {
		length, err := clients.ElChainReader.GetMaxRetroactiveLength(context.Background())
		require.NoError(t, err)
		require.NotZero(t, length)

		lengthActual, err := rewardsCoordinator.MAXRETROACTIVELENGTH(&bind.CallOpts{})
		require.NoError(t, err)
		require.Equal(t, length, lengthActual)
	})

	t.Run("get max future length", func(t *testing.T) {
		length, err := clients.ElChainReader.GetMaxFutureLength(context.Background())
		require.NoError(t, err)
		require.NotZero(t, length)

		lengthActual, err := rewardsCoordinator.MAXFUTURELENGTH(&bind.CallOpts{})
		require.NoError(t, err)
		require.Equal(t, length, lengthActual)
	})

	t.Run("get genesis rewards timestamp", func(t *testing.T) {
		timestamp, err := clients.ElChainReader.GetGenesisRewardsTimestamp(context.Background())
		require.NoError(t, err)
		require.NotZero(t, timestamp)

		timestampActual, err := rewardsCoordinator.GENESISREWARDSTIMESTAMP(&bind.CallOpts{})
		require.NoError(t, err)
		require.Equal(t, timestamp, timestampActual)
	})

	t.Run("Get rewards updater", func(t *testing.T) {
		updater, err := clients.ElChainReader.GetRewardsUpdater(context.Background())
		require.NoError(t, err)

		updaterActual, err := rewardsCoordinator.RewardsUpdater(&bind.CallOpts{})
		require.NoError(t, err)
		require.Equal(t, updaterActual, updater)
	})

	t.Run("get current rewards calculation end timestamp", func(t *testing.T) {
		timestamp, err := clients.ElChainReader.GetCurrRewardsCalculationEndTimestamp(context.Background())
		require.NoError(t, err)

		timestampActual, err := rewardsCoordinator.CurrRewardsCalculationEndTimestamp(&bind.CallOpts{})
		require.NoError(t, err)
		require.Equal(t, timestamp, timestampActual)
	})

	t.Run("get default operator split bips", func(t *testing.T) {
		splitBips, err := clients.ElChainReader.GetDefaultOperatorSplitBips(context.Background())
		require.NoError(t, err)
		require.NotZero(t, splitBips)

		splitBipsActual, err := rewardsCoordinator.DefaultOperatorSplitBips(&bind.CallOpts{})
		require.NoError(t, err)
		require.Equal(t, splitBips, splitBipsActual)
	})

	t.Run("get claimer for operator", func(t *testing.T) {
		claimer, err := clients.ElChainReader.GetClaimerFor(context.Background(), common.HexToAddress(operator.Address))
		require.NoError(t, err)

		receipt, err := clients.ElChainWriter.SetClaimerFor(
			context.Background(),
			common.HexToAddress(operator.Address),
			true,
		)
		require.NoError(t, err)
		require.Equal(t, gethtypes.ReceiptStatusSuccessful, receipt.Status)

		newClaimer, err := clients.ElChainReader.GetClaimerFor(
			context.Background(),
			common.HexToAddress(operator.Address),
		)
		require.NoError(t, err)
		require.NotEqual(t, claimer, newClaimer)
		require.Equal(t, common.HexToAddress(operator.Address), newClaimer)
	})

	t.Run("get cumulative claimed", func(t *testing.T) {
		strategyAddr := contractAddrs.Erc20MockStrategy
		strategy, underlyingTokenAddr, err := clients.ElChainReader.GetStrategyAndUnderlyingToken(
			ctx,
			strategyAddr,
		)
		assert.NoError(t, err)
		assert.NotNil(t, strategy)
		assert.NotEqual(t, common.Address{}, underlyingTokenAddr)
		claimed, err := clients.ElChainReader.GetCumulativeClaimed(
			context.Background(),
			common.HexToAddress(operator.Address),
			underlyingTokenAddr,
		)
		require.NoError(t, err)
		require.NotNil(t, claimed)
	})

	t.Run("get submission nonce", func(t *testing.T) {
		nonce, err := clients.ElChainReader.GetSubmissionNonce(context.Background(), contractAddrs.ServiceManager)
		require.NoError(t, err)
		require.NotNil(t, nonce)
	})

	t.Run("get is AVS rewards submission hash", func(t *testing.T) {
		isValid, err := clients.ElChainReader.GetIsAVSRewardsSubmissionHash(
			context.Background(),
			contractAddrs.ServiceManager,
			[32]byte{},
		)
		require.NoError(t, err)
		require.False(t, isValid)
	})

	t.Run("get is rewards submission for all hash", func(t *testing.T) {
		isValid, err := clients.ElChainReader.GetIsRewardsSubmissionForAllHash(
			context.Background(),
			contractAddrs.ServiceManager,
			[32]byte{},
		)
		require.NoError(t, err)
		require.False(t, isValid)
	})

	t.Run("get is rewards for all submitter", func(t *testing.T) {
		isValid, err := clients.ElChainReader.GetIsRewardsForAllSubmitter(
			context.Background(),
			common.HexToAddress(operator.Address),
		)
		require.NoError(t, err)
		require.False(t, isValid)
	})

	t.Run("get is rewards submission for all earners hash", func(t *testing.T) {
		isValid, err := clients.ElChainReader.GetIsRewardsSubmissionForAllEarnersHash(
			context.Background(),
			common.HexToAddress(operator.Address),
			[32]byte{},
		)
		require.NoError(t, err)
		require.False(t, isValid)
	})

	t.Run("get is operator directed AVS rewards submission hash", func(t *testing.T) {
		isValid, err := clients.ElChainReader.GetIsOperatorDirectedAVSRewardsSubmissionHash(
			context.Background(),
			common.HexToAddress(operator.Address),
			[32]byte{},
		)
		require.NoError(t, err)
		require.False(t, isValid)
	})

	t.Run("get is operator directed operator set rewards submission hash", func(t *testing.T) {
		isValid, err := clients.ElChainReader.GetIsOperatorDirectedOperatorSetRewardsSubmissionHash(
			context.Background(),
			common.HexToAddress(operator.Address),
			[32]byte{},
		)
		require.NoError(t, err)
		require.False(t, isValid)
	})

}

func TestGetCurrentClaimableDistributionRoot(t *testing.T) {
	// Verifies GetCurrentClaimableDistributionRoot returns 0 if no root and the root if there's one
	_, anvilHttpEndpoint := testclients.BuildTestClients(t)
	ctx := context.Background()

	contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilHttpEndpoint)

	root := [32]byte{
		0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x01, 0x01,
		0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01,
		0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01,
		0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01,
	}

	rewardsCoordinatorAddr := contractAddrs.RewardsCoordinator
	config := elcontracts.Config{
		DelegationManagerAddress:  contractAddrs.DelegationManager,
		RewardsCoordinatorAddress: rewardsCoordinatorAddr,
	}

	chainReader, err := testclients.NewTestChainReaderFromConfig(anvilHttpEndpoint, config)
	require.NoError(t, err)

	// Create and configure rewards coordinator
	ethClient, err := ethclient.Dial(anvilHttpEndpoint)
	require.NoError(t, err)
	rewardsCoordinator, err := rewardscoordinator.NewContractRewardsCoordinator(rewardsCoordinatorAddr, ethClient)
	require.NoError(t, err)

	ecdsaPrivKeyHex := testutils.ANVIL_FIRST_PRIVATE_KEY

	// Set delay to zero to inmediatly operate with coordinator
	receipt, err := setTestRewardsCoordinatorActivationDelay(anvilHttpEndpoint, ecdsaPrivKeyHex, uint32(0))
	require.NoError(t, err)
	require.Equal(t, receipt.Status, gethtypes.ReceiptStatusSuccessful)

	// Create txManager to send transactions to the Ethereum node
	txManager, err := testclients.NewTestTxManager(anvilHttpEndpoint, ecdsaPrivKeyHex)
	require.NoError(t, err)
	noSendTxOpts, err := txManager.GetNoSendTxOpts()
	require.NoError(t, err)

	rewardsUpdater := common.HexToAddress(testutils.ANVIL_FIRST_ADDRESS)

	// Change the rewards updater to be able to submit the new root
	tx, err := rewardsCoordinator.SetRewardsUpdater(noSendTxOpts, rewardsUpdater)
	require.NoError(t, err)

	waitForReceipt := true
	_, err = txManager.Send(context.Background(), tx, waitForReceipt)
	require.NoError(t, err)

	// Check that if there is no root submitted the result is zero
	distr_root, err := chainReader.GetCurrentClaimableDistributionRoot(
		ctx,
	)
	assert.NoError(t, err)
	assert.Zero(t, distr_root.Root)

	currRewardsCalculationEndTimestamp, err := chainReader.CurrRewardsCalculationEndTimestamp(context.Background())
	require.NoError(t, err)

	tx, err = rewardsCoordinator.SubmitRoot(noSendTxOpts, root, currRewardsCalculationEndTimestamp+1)
	require.NoError(t, err)

	_, err = txManager.Send(context.Background(), tx, waitForReceipt)
	require.NoError(t, err)

	// Check that if there is a root submitted the result is that root
	distr_root, err = chainReader.GetCurrentClaimableDistributionRoot(
		ctx,
	)
	assert.NoError(t, err)
	assert.Equal(t, distr_root.Root, root)
}

func TestGetRootIndexFromRootHash(t *testing.T) {
	_, anvilHttpEndpoint := testclients.BuildTestClients(t)
	ctx := context.Background()

	contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilHttpEndpoint)

	rewardsCoordinatorAddr := contractAddrs.RewardsCoordinator
	config := elcontracts.Config{
		DelegationManagerAddress:  contractAddrs.DelegationManager,
		RewardsCoordinatorAddress: rewardsCoordinatorAddr,
	}

	chainReader, err := testclients.NewTestChainReaderFromConfig(anvilHttpEndpoint, config)
	require.NoError(t, err)

	// Create and configure rewards coordinator
	ethClient, err := ethclient.Dial(anvilHttpEndpoint)
	require.NoError(t, err)
	rewardsCoordinator, err := rewardscoordinator.NewContractRewardsCoordinator(rewardsCoordinatorAddr, ethClient)
	require.NoError(t, err)
	ecdsaPrivKeyHex := testutils.ANVIL_FIRST_PRIVATE_KEY

	// Set delay to zero to inmediatly operate with coordinator
	receipt, err := setTestRewardsCoordinatorActivationDelay(anvilHttpEndpoint, ecdsaPrivKeyHex, uint32(0))
	require.NoError(t, err)
	require.Equal(t, receipt.Status, gethtypes.ReceiptStatusSuccessful)

	// Create txManager to send transactions to the Ethereum node
	txManager, err := testclients.NewTestTxManager(anvilHttpEndpoint, ecdsaPrivKeyHex)
	require.NoError(t, err)
	noSendTxOpts, err := txManager.GetNoSendTxOpts()
	require.NoError(t, err)

	rewardsUpdater := common.HexToAddress(testutils.ANVIL_FIRST_ADDRESS)

	// Change the rewards updater to be able to submit the new root
	tx, err := rewardsCoordinator.SetRewardsUpdater(noSendTxOpts, rewardsUpdater)
	require.NoError(t, err)

	waitForReceipt := true
	_, err = txManager.Send(context.Background(), tx, waitForReceipt)
	require.NoError(t, err)

	root := [32]byte{
		0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x01, 0x01,
		0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01,
		0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01,
		0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01,
	}

	// Check that if there is no root submitted the result is an InvalidRoot error
	root_index, err := chainReader.GetRootIndexFromHash(
		ctx,
		root,
	)
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "execution reverted: custom error 0x504570e3",
		"GetRootIndexFromHash should return an InvalidRoot() error",
	)
	assert.Zero(t, root_index)

	currRewardsCalculationEndTimestamp, err := chainReader.CurrRewardsCalculationEndTimestamp(context.Background())
	require.NoError(t, err)

	tx, err = rewardsCoordinator.SubmitRoot(noSendTxOpts, root, currRewardsCalculationEndTimestamp+1)
	require.NoError(t, err)

	_, err = txManager.Send(context.Background(), tx, waitForReceipt)
	require.NoError(t, err)

	root2 := [32]byte{
		0x00, 0x00, 0x00, 0x00, 0x01, 0x01, 0x01, 0x01,
		0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01,
		0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01,
		0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01,
	}

	currRewardsCalculationEndTimestamp2, err := chainReader.CurrRewardsCalculationEndTimestamp(context.Background())
	require.NoError(t, err)

	tx, err = rewardsCoordinator.SubmitRoot(noSendTxOpts, root2, currRewardsCalculationEndTimestamp2+1)
	require.NoError(t, err)

	_, err = txManager.Send(context.Background(), tx, waitForReceipt)
	require.NoError(t, err)

	// Check that the first root inserted is the first indexed (zero)
	root_index, err = chainReader.GetRootIndexFromHash(
		ctx,
		root,
	)
	assert.NoError(t, err)
	assert.Equal(t, root_index, uint32(0))

	// Check that the second root inserted is the second indexed (zero)
	root_index, err = chainReader.GetRootIndexFromHash(
		ctx,
		root2,
	)
	assert.NoError(t, err)
	assert.Equal(t, root_index, uint32(1))
}

func TestGetCumulativeClaimedRewards(t *testing.T) {
	clients, anvilHttpEndpoint := testclients.BuildTestClients(t)
	ctx := context.Background()

	contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilHttpEndpoint)

	rewardsCoordinatorAddr := contractAddrs.RewardsCoordinator
	config := elcontracts.Config{
		DelegationManagerAddress:  contractAddrs.DelegationManager,
		RewardsCoordinatorAddress: rewardsCoordinatorAddr,
	}
	privateKeyHex := testutils.ANVIL_FIRST_PRIVATE_KEY

	// Create ChainWriter
	chainWriter, err := testclients.NewTestChainWriterFromConfig(anvilHttpEndpoint, privateKeyHex, config)
	require.NoError(t, err)

	chainReader, err := testclients.NewTestChainReaderFromConfig(anvilHttpEndpoint, config)
	require.NoError(t, err)

	activationDelay := uint32(0)
	// Set activation delay to zero so that the earnings can be claimed right after submitting the root
	receipt, err := setTestRewardsCoordinatorActivationDelay(anvilHttpEndpoint, privateKeyHex, activationDelay)
	require.NoError(t, err)
	require.True(t, receipt.Status == gethtypes.ReceiptStatusSuccessful)

	strategyAddr := contractAddrs.Erc20MockStrategy
	strategy, contractUnderlyingToken, underlyingTokenAddr, err := clients.ElChainReader.GetStrategyAndUnderlyingERC20Token(
		ctx,
		strategyAddr,
	)
	assert.NoError(t, err)
	assert.NotNil(t, strategy)
	assert.NotEqual(t, common.Address{}, underlyingTokenAddr)
	assert.NotNil(t, contractUnderlyingToken)

	anvil_address := common.HexToAddress(testutils.ANVIL_FIRST_ADDRESS)

	// This tests that without claims result is zero
	claimed, err := chainReader.GetCumulativeClaimed(ctx, anvil_address, underlyingTokenAddr)
	assert.Zero(t, claimed.Cmp(big.NewInt(0)))
	assert.NoError(t, err)

	cumulativeEarnings := int64(45)
	claim, err := newTestClaim(chainReader, anvilHttpEndpoint, cumulativeEarnings, privateKeyHex)
	require.NoError(t, err)

	receipt, err = chainWriter.ProcessClaim(context.Background(), *claim, rewardsCoordinatorAddr, true)
	require.NoError(t, err)
	require.True(t, receipt.Status == gethtypes.ReceiptStatusSuccessful)

	// This tests that with a claim result is cumulativeEarnings
	claimed, err = chainReader.GetCumulativeClaimed(ctx, anvil_address, underlyingTokenAddr)
	assert.Equal(t, claimed, big.NewInt(cumulativeEarnings))
	assert.NoError(t, err)
}

func TestCheckClaim(t *testing.T) {
	clients, anvilHttpEndpoint := testclients.BuildTestClients(t)
	ctx := context.Background()

	contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilHttpEndpoint)

	rewardsCoordinatorAddr := contractAddrs.RewardsCoordinator
	config := elcontracts.Config{
		DelegationManagerAddress:  contractAddrs.DelegationManager,
		RewardsCoordinatorAddress: rewardsCoordinatorAddr,
	}
	privateKeyHex := testutils.ANVIL_FIRST_PRIVATE_KEY

	// Create ChainWriter and chain reader
	chainWriter, err := testclients.NewTestChainWriterFromConfig(anvilHttpEndpoint, privateKeyHex, config)
	require.NoError(t, err)

	chainReader, err := testclients.NewTestChainReaderFromConfig(anvilHttpEndpoint, config)
	require.NoError(t, err)

	activationDelay := uint32(0)
	// Set activation delay to zero so that the earnings can be claimed right after submitting the root
	receipt, err := setTestRewardsCoordinatorActivationDelay(anvilHttpEndpoint, privateKeyHex, activationDelay)
	require.NoError(t, err)
	require.True(t, receipt.Status == gethtypes.ReceiptStatusSuccessful)

	cumulativeEarnings := int64(45)
	claim, err := newTestClaim(chainReader, anvilHttpEndpoint, cumulativeEarnings, privateKeyHex)
	require.NoError(t, err)

	receipt, err = chainWriter.ProcessClaim(context.Background(), *claim, rewardsCoordinatorAddr, true)
	require.NoError(t, err)
	require.True(t, receipt.Status == gethtypes.ReceiptStatusSuccessful)

	strategyAddr := contractAddrs.Erc20MockStrategy
	strategy, contractUnderlyingToken, underlyingTokenAddr, err := clients.ElChainReader.GetStrategyAndUnderlyingERC20Token(
		ctx,
		strategyAddr,
	)
	assert.NoError(t, err)
	assert.NotNil(t, strategy)
	assert.NotEqual(t, common.Address{}, underlyingTokenAddr)
	assert.NotNil(t, contractUnderlyingToken)

	checked, err := chainReader.CheckClaim(ctx, *claim)
	require.NoError(t, err)
	assert.True(t, checked)
}

func TestGetAllocatableMagnitudeAndEncumberedMagnitudeAndGetMaxMagnitudes(t *testing.T) {
	// Without changes, Allocable magnitude is max magnitude

	// Test setup
	ctx := context.Background()

	clients, anvilHttpEndpoint := testclients.BuildTestClients(t)
	contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilHttpEndpoint)

	anvilC := clients.AnvilC
	operatorAddr := common.HexToAddress(testutils.ANVIL_FIRST_ADDRESS)

	chainReader := clients.ElChainReader

	strategyAddr := contractAddrs.Erc20MockStrategy
	testAddr := common.HexToAddress(testutils.ANVIL_FIRST_ADDRESS)

	strategies := []common.Address{strategyAddr}
	maxMagnitudes, err := chainReader.GetMaxMagnitudes(ctx, testAddr, strategies)
	assert.NoError(t, err)

	// Assert that at the beginning, Allocatable Magnitude is Max allocatable magnitude
	allocable, err := chainReader.GetAllocatableMagnitude(ctx, testAddr, strategyAddr)
	assert.NoError(t, err)

	encumberedMagnitude, err := chainReader.GetEncumberedMagnitude(ctx, testAddr, strategyAddr)
	assert.NoError(t, err)
	assert.Zero(t, encumberedMagnitude)

	assert.Equal(t, maxMagnitudes[0], allocable)

	// Reduce allocatable magnitude for testAddr
	chainWriter := clients.ElChainWriter

	waitForReceipt := true
	delay := uint32(1)
	receipt, err := chainWriter.SetAllocationDelay(context.Background(), operatorAddr, delay, waitForReceipt)
	require.NoError(t, err)
	require.Equal(t, gethtypes.ReceiptStatusSuccessful, receipt.Status)

	allocationConfigurationDelay := 1200
	testutils.AdvanceChainByNBlocksExecInContainer(context.Background(), allocationConfigurationDelay+1, anvilC)

	// Check that Allocation delay has been applied
	_, err = chainReader.GetAllocationDelay(context.Background(), operatorAddr)
	require.NoError(t, err)

	operatorSetId := uint32(1)
	err = createOperatorSet(clients, anvilHttpEndpoint, contractAddrs.ServiceManager, strategyAddr)
	require.NoError(t, err)

	operatorSet := allocationmanager.OperatorSet{
		Avs: contractAddrs.ServiceManager,
		Id:  operatorSetId,
	}
	allocatable_reduction := uint64(100)
	allocateParams := []allocationmanager.IAllocationManagerTypesAllocateParams{
		{
			OperatorSet:   operatorSet,
			Strategies:    []common.Address{strategyAddr},
			NewMagnitudes: []uint64{allocatable_reduction},
		},
	}

	receipt, err = chainWriter.ModifyAllocations(context.Background(), operatorAddr, allocateParams, waitForReceipt)
	require.NoError(t, err)
	require.Equal(t, gethtypes.ReceiptStatusSuccessful, receipt.Status)

	// Assert that after stake reduction, Allocatable Magnitude + reduction ammount equals Max allocatable magnitude
	allocable, err = chainReader.GetAllocatableMagnitude(ctx, testAddr, strategyAddr)
	assert.NoError(t, err)
	assert.Equal(t, maxMagnitudes[0], allocable+allocatable_reduction)

	encumberedMagnitude, err = chainReader.GetEncumberedMagnitude(ctx, testAddr, strategyAddr)
	assert.Equal(t, encumberedMagnitude, allocatable_reduction)

	assert.NoError(t, err)

	// Check that the new allocationDelay is equal to delay
	op := types.Operator{
		Address: operatorAddr.String(),
	}

	operatorDetails, err := chainReader.GetOperatorDetails(ctx, op)
	assert.NoError(t, err)
	assert.NotNil(t, operatorDetails)
	assert.Equal(t, op.Address, operatorDetails.Address)
	assert.Equal(t, delay, operatorDetails.AllocationDelay)
}

func TestAdminFunctions(t *testing.T) {
	testConfig := testutils.GetDefaultTestConfig()
	anvilC, err := testutils.StartAnvilContainer(testConfig.AnvilStateFileName)
	assert.NoError(t, err)

	anvilHttpEndpoint, err := anvilC.Endpoint(context.Background(), "http")
	assert.NoError(t, err)

	contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilHttpEndpoint)

	config := elcontracts.Config{PermissionControllerAddress: contractAddrs.PermissionController}

	operatorAddr := common.HexToAddress(testutils.ANVIL_FIRST_ADDRESS)
	privateKeyHex := testutils.ANVIL_FIRST_PRIVATE_KEY
	accountChainWriter, err := testclients.NewTestChainWriterFromConfig(anvilHttpEndpoint, privateKeyHex, config)
	assert.NoError(t, err)

	pendingAdminAddr := common.HexToAddress(testutils.ANVIL_SECOND_ADDRESS)
	pendingAdminPrivateKeyHex := testutils.ANVIL_SECOND_PRIVATE_KEY
	adminChainWriter, err := testclients.NewTestChainWriterFromConfig(
		anvilHttpEndpoint,
		pendingAdminPrivateKeyHex,
		config,
	)
	assert.NoError(t, err)

	chainReader, err := testclients.NewTestChainReaderFromConfig(anvilHttpEndpoint, config)
	assert.NoError(t, err)

	t.Run("non-existent pending admin", func(t *testing.T) {
		isPendingAdmin, err := chainReader.IsPendingAdmin(context.Background(), operatorAddr, pendingAdminAddr)
		assert.NoError(t, err)
		assert.False(t, isPendingAdmin)
	})

	t.Run("list pending admins when empty", func(t *testing.T) {
		listPendingAdmins, err := chainReader.ListPendingAdmins(context.Background(), operatorAddr)
		assert.NoError(t, err)
		assert.Empty(t, listPendingAdmins)
	})

	t.Run("add pending admin and list", func(t *testing.T) {
		request := elcontracts.AddPendingAdminRequest{
			AccountAddress: operatorAddr,
			AdminAddress:   pendingAdminAddr,
			WaitForReceipt: true,
		}

		receipt, err := accountChainWriter.AddPendingAdmin(context.Background(), request)
		assert.NoError(t, err)
		assert.Equal(t, receipt.Status, gethtypes.ReceiptStatusSuccessful)

		isPendingAdmin, err := chainReader.IsPendingAdmin(context.Background(), operatorAddr, pendingAdminAddr)
		assert.NoError(t, err)
		assert.True(t, isPendingAdmin)

		listPendingAdmins, err := chainReader.ListPendingAdmins(context.Background(), operatorAddr)
		assert.NoError(t, err)
		assert.NotEmpty(t, listPendingAdmins)
	})

	t.Run("non-existent admin", func(t *testing.T) {
		isAdmin, err := chainReader.IsAdmin(context.Background(), operatorAddr, pendingAdminAddr)
		assert.NoError(t, err)
		assert.False(t, isAdmin)
	})

	t.Run("list admins", func(t *testing.T) {
		acceptAdminRequest := elcontracts.AcceptAdminRequest{
			AccountAddress: operatorAddr,
			WaitForReceipt: true,
		}

		receipt, err := adminChainWriter.AcceptAdmin(context.Background(), acceptAdminRequest)
		assert.NoError(t, err)
		assert.Equal(t, receipt.Status, gethtypes.ReceiptStatusSuccessful)

		listAdmins, err := chainReader.ListAdmins(context.Background(), operatorAddr)
		assert.NoError(t, err)
		assert.Len(t, listAdmins, 1)

		admin := listAdmins[0]
		isAdmin, err := chainReader.IsAdmin(context.Background(), operatorAddr, admin)
		assert.NoError(t, err)
		assert.True(t, isAdmin)
	})
}

func TestAppointeesFunctions(t *testing.T) {
	testConfig := testutils.GetDefaultTestConfig()
	anvilC, err := testutils.StartAnvilContainer(testConfig.AnvilStateFileName)
	assert.NoError(t, err)

	anvilHttpEndpoint, err := anvilC.Endpoint(context.Background(), "http")
	assert.NoError(t, err)

	contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilHttpEndpoint)

	config := elcontracts.Config{PermissionControllerAddress: contractAddrs.PermissionController}

	chainReader, err := testclients.NewTestChainReaderFromConfig(anvilHttpEndpoint, config)
	assert.NoError(t, err)

	privateKey := testutils.ANVIL_FIRST_PRIVATE_KEY
	chainWriter, err := testclients.NewTestChainWriterFromConfig(anvilHttpEndpoint, privateKey, config)
	assert.NoError(t, err)

	accountAddress := common.HexToAddress(testutils.ANVIL_FIRST_ADDRESS)

	appointeeAddress := common.HexToAddress(testutils.ANVIL_SECOND_ADDRESS)
	target := common.HexToAddress(testutils.ANVIL_THIRD_ADDRESS)
	selector := [4]byte{0, 1, 2, 3}

	t.Run("list appointees when empty", func(t *testing.T) {
		appointees, err := chainReader.ListAppointees(context.Background(), accountAddress, target, selector)
		assert.NoError(t, err)
		assert.Empty(t, appointees)
	})

	t.Run("list appointees", func(t *testing.T) {
		setPermissionRequest := elcontracts.SetPermissionRequest{
			AccountAddress:   accountAddress,
			AppointeeAddress: appointeeAddress,
			Target:           target,
			Selector:         selector,
			WaitForReceipt:   true,
		}

		receipt, err := chainWriter.SetPermission(context.Background(), setPermissionRequest)
		require.NoError(t, err)
		require.Equal(t, receipt.Status, gethtypes.ReceiptStatusSuccessful)

		canCall, err := chainReader.CanCall(context.Background(), accountAddress, appointeeAddress, target, selector)
		require.NoError(t, err)
		require.True(t, canCall)

		appointees, err := chainReader.ListAppointees(context.Background(), accountAddress, target, selector)
		assert.NoError(t, err)
		assert.NotEmpty(t, appointees)
	})

	t.Run("list appointees permissions", func(t *testing.T) {
		appointeesPermission, _, err := chainReader.ListAppointeePermissions(
			context.Background(),
			accountAddress,
			appointeeAddress,
		)
		assert.NoError(t, err)
		assert.NotEmpty(t, appointeesPermission)
	})
}

func TestContractErrorCases(t *testing.T) {
	ctx := context.Background()

	testConfig := testutils.GetDefaultTestConfig()
	anvilC, err := testutils.StartAnvilContainer(testConfig.AnvilStateFileName)
	require.NoError(t, err)

	anvilHttpEndpoint, err := anvilC.Endpoint(context.Background(), "http")
	require.NoError(t, err)
	contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilHttpEndpoint)

	config := elcontracts.Config{
		DelegationManagerAddress: contractAddrs.DelegationManager,
	}

	chainReader, err := testclients.NewTestChainReaderFromConfig(anvilHttpEndpoint, config)
	require.NoError(t, err)

	// This address does not belong to a Token contract
	strategyAddr := common.HexToAddress("34634374736473673643")

	t.Run("GetStrategyAndUnderlyingToken", func(t *testing.T) {
		_, _, err := chainReader.GetStrategyAndUnderlyingToken(ctx, strategyAddr)
		assert.Error(t, err)
		assert.Equal(t, err.Error(), "Failed to fetch token contract: no contract code at given address")
	})

	t.Run("GetStrategyAndUnderlyingERC20Token", func(t *testing.T) {
		_, _, _, err := chainReader.GetStrategyAndUnderlyingERC20Token(ctx, strategyAddr)
		assert.Error(t, err)
		assert.Equal(t, err.Error(), "Failed to fetch token contract: no contract code at given address")
	})
}

// TestInvalidConfig tests the behavior of the chainReader when the config is invalid (e.g. missing addresses, wrong
// addresses)
func TestInvalidConfig(t *testing.T) {
	testConfig := testutils.GetDefaultTestConfig()
	anvilC, err := testutils.StartAnvilContainer(testConfig.AnvilStateFileName)
	require.NoError(t, err)

	anvilHttpEndpoint, err := anvilC.Endpoint(context.Background(), "http")
	require.NoError(t, err)

	operatorAddr := testutils.ANVIL_FIRST_ADDRESS
	operator := types.Operator{
		Address: operatorAddr,
	}

	config := elcontracts.Config{}
	chainReader, err := testclients.NewTestChainReaderFromConfig(anvilHttpEndpoint, config)
	require.NoError(t, err)

	t.Run("try to check if operator is registered with invalid config", func(t *testing.T) {
		// IsOperatorRegistered needs a correct DelegationManagerAddress
		_, err := chainReader.IsOperatorRegistered(context.Background(), operator)
		require.Error(t, err)
	})

	t.Run("get operator details with invalid config", func(t *testing.T) {
		// GetOperatorDetails needs a correct DelegationManagerAddress
		_, err := chainReader.GetOperatorDetails(context.Background(), operator)
		require.Error(t, err)
	})

	t.Run("try to get strategy and underlying token with wrong strategy address", func(t *testing.T) {
		// Invalid strategy address
		strategyAddr := common.HexToAddress(testutils.ANVIL_FIRST_ADDRESS)
		operatorAddr := common.HexToAddress(testutils.ANVIL_SECOND_ADDRESS)

		// GetOperatorSharesInStrategy needs a correct DelegationManagerAddress
		_, err := chainReader.GetOperatorSharesInStrategy(context.Background(), operatorAddr, strategyAddr)
		require.Error(t, err)

		// GetStrategyAndUnderlyingToken needs a correct StrategyAddress
		_, _, err = chainReader.GetStrategyAndUnderlyingToken(context.Background(), strategyAddr)
		require.Error(t, err)

		_, _, _, err = chainReader.GetStrategyAndUnderlyingERC20Token(context.Background(), strategyAddr)
		require.Error(t, err)
	})

	t.Run("calculate digest hash with invalid config", func(t *testing.T) {
		staker := common.Address{0x0}
		delegationApprover := common.Address{0x0}
		approverSalt := [32]byte{}
		expiry := big.NewInt(0)

		// CalculateDelegationApprovalDigestHash needs a correct DelegationManagerAddress
		_, err := chainReader.CalculateDelegationApprovalDigestHash(
			context.Background(),
			staker,
			common.HexToAddress(operatorAddr),
			delegationApprover,
			approverSalt,
			expiry,
		)
		require.Error(t, err)

		// CalculateOperatorAVSRegistrationDigestHash needs a correct AvsDirectoryAddress
		_, err = chainReader.CalculateOperatorAVSRegistrationDigestHash(context.Background(),
			common.HexToAddress(operatorAddr),
			staker,
			approverSalt, expiry)
		require.Error(t, err)
	})

	t.Run("get root with invalid config", func(t *testing.T) {
		// GetDistributionRootsLength needs a correct RewardsCoordinatorAddress
		_, err := chainReader.GetDistributionRootsLength(context.Background())
		require.Error(t, err)

		// GetRootIndexFromHash needs a correct RewardsCoordinatorAddress
		_, err = chainReader.GetRootIndexFromHash(context.Background(), [32]byte{})
		require.Error(t, err)

		_, err = chainReader.GetCurrentClaimableDistributionRoot(context.Background())
		require.Error(t, err)
	})

	t.Run("get magnitudes, rewards and claims with invalid config", func(t *testing.T) {
		contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilHttpEndpoint)
		strategyAddr := contractAddrs.Erc20MockStrategy

		_, err = chainReader.GetCurrentClaimableDistributionRoot(context.Background())
		require.Error(t, err)

		_, err := chainReader.GetCumulativeClaimed(
			context.Background(),
			common.HexToAddress(testutils.ANVIL_THIRD_ADDRESS),
			common.HexToAddress(testutils.ANVIL_SECOND_ADDRESS),
		)
		require.Error(t, err)

		_, err = chainReader.GetMaxMagnitudes(
			context.Background(),
			common.HexToAddress(operatorAddr),
			[]common.Address{strategyAddr},
		)
		require.Error(t, err)

		_, err = chainReader.GetAllocatableMagnitude(
			context.Background(),
			common.HexToAddress(operatorAddr),
			strategyAddr,
		)
		require.Error(t, err)

		_, err = chainReader.GetAllocationInfo(context.Background(), common.HexToAddress(operatorAddr), strategyAddr)
		require.Error(t, err)

		_, err = chainReader.GetAllocationDelay(context.Background(), common.HexToAddress(operatorAddr))
		require.Error(t, err)

		_, err = chainReader.CheckClaim(
			context.Background(),
			rewardscoordinator.IRewardsCoordinatorTypesRewardsMerkleClaim{},
		)
		require.Error(t, err)

		_, err = chainReader.CurrRewardsCalculationEndTimestamp(context.Background())
		require.Error(t, err)
	})

	t.Run("try to get a staker shares with invalid config", func(t *testing.T) {
		// GetStakerShares needs a correct DelegationManagerAddress
		_, _, err := chainReader.GetStakerShares(context.Background(), common.HexToAddress(operator.Address))
		require.Error(t, err)
	})

	t.Run("try to get the delegated operator shares with invalid config", func(t *testing.T) {
		// GetDelegatedOperator needs a correct DelegationManagerAddress
		_, err := chainReader.GetDelegatedOperator(
			context.Background(),
			common.HexToAddress(operator.Address),
			big.NewInt(0),
		)
		require.Error(t, err)
	})

	t.Run("try to get the number of operator sets for an operator with invalid config", func(t *testing.T) {
		// GetNumOperatorSetsForOperator needs a correct AllocationManagerAddress
		_, err := chainReader.GetNumOperatorSetsForOperator(context.Background(), common.HexToAddress(operator.Address))
		require.Error(t, err)
	})

	t.Run("try to get the operator sets for an operator with invalid config", func(t *testing.T) {
		// GetOperatorSetsForOperator needs a correct AllocationManagerAddress
		_, err := chainReader.GetOperatorSetsForOperator(context.Background(), common.HexToAddress(operator.Address))
		require.Error(t, err)
	})

	t.Run("try to check if the operator is registered in an operator set with set id 0 and an invalid config",
		func(t *testing.T) {
			// IsOperatorRegisteredWithOperatorSet with setId 0 needs a correct AVSDirectory
			testAddr := common.HexToAddress(testutils.ANVIL_FIRST_ADDRESS)
			operatorSetId := uint32(0)
			operatorSet := allocationmanager.OperatorSet{
				Avs: testAddr,
				Id:  operatorSetId,
			}
			_, err := chainReader.IsOperatorRegisteredWithOperatorSet(
				context.Background(),
				common.HexToAddress(operator.Address),
				operatorSet,
			)
			require.Error(t, err)
		},
	)

	t.Run("try to check if the operator is registered in an operator set with set id 1 and an invalid config",
		func(t *testing.T) {
			// IsOperatorRegisteredWithOperatorSet with setId 1 needs a correct AllocationManagerAddress
			testAddr := common.HexToAddress(testutils.ANVIL_FIRST_ADDRESS)
			operatorSetId := uint32(1)
			operatorSet := allocationmanager.OperatorSet{
				Avs: testAddr,
				Id:  operatorSetId,
			}
			_, err := chainReader.IsOperatorRegisteredWithOperatorSet(
				context.Background(),
				common.HexToAddress(operator.Address),
				operatorSet,
			)
			require.Error(t, err)
		},
	)

	t.Run("try to get the operators for an operator set with set id 1 and an invalid config",
		func(t *testing.T) {
			// GetOperatorsForOperatorSet needs a correct AllocationManagerAddress
			testAddr := common.HexToAddress(testutils.ANVIL_FIRST_ADDRESS)
			operatorSetId := uint32(1)
			operatorSet := allocationmanager.OperatorSet{
				Avs: testAddr,
				Id:  operatorSetId,
			}
			_, err := chainReader.GetOperatorsForOperatorSet(
				context.Background(),
				operatorSet,
			)
			require.Error(t, err)
		},
	)

	t.Run("try to get the number of operators for an operator set with set id 1 and an invalid config",
		func(t *testing.T) {
			// GetNumOperatorsForOperatorSet needs a correct AllocationManagerAddress
			testAddr := common.HexToAddress(testutils.ANVIL_FIRST_ADDRESS)
			operatorSetId := uint32(1)
			operatorSet := allocationmanager.OperatorSet{
				Avs: testAddr,
				Id:  operatorSetId,
			}
			_, err := chainReader.GetNumOperatorsForOperatorSet(
				context.Background(),
				operatorSet,
			)
			require.Error(t, err)
		},
	)

	t.Run("try to get the strategies for an operator set with set id 1 and an invalid config",
		func(t *testing.T) {
			// GetStrategiesForOperatorSet needs a correct AllocationManagerAddress
			testAddr := common.HexToAddress(testutils.ANVIL_FIRST_ADDRESS)
			operatorSetId := uint32(1)
			operatorSet := allocationmanager.OperatorSet{
				Avs: testAddr,
				Id:  operatorSetId,
			}
			_, err := chainReader.GetStrategiesForOperatorSet(
				context.Background(),
				operatorSet,
			)
			require.Error(t, err)
		},
	)
}

func TestOperatorSetsAndSlashableShares(t *testing.T) {
	testConfig := testutils.GetDefaultTestConfig()
	anvilC, err := testutils.StartAnvilContainer(testConfig.AnvilStateFileName)
	require.NoError(t, err)

	clients, anvilHttpEndpoint := testclients.BuildTestClients(t)
	contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilHttpEndpoint)

	chainReader := clients.ElChainReader

	operatorAddr := common.HexToAddress(testutils.ANVIL_SECOND_ADDRESS)
	chainWriter := clients.ElChainWriter

	avsAddr := contractAddrs.ServiceManager
	operatorSetId := uint32(1)
	operatorSet := allocationmanager.OperatorSet{
		Avs: avsAddr,
		Id:  operatorSetId,
	}

	strategyAddr := contractAddrs.Erc20MockStrategy
	strategies := []common.Address{strategyAddr}

	err = createOperatorSet(clients, anvilHttpEndpoint, avsAddr, strategyAddr)
	require.NoError(t, err)

	keypair, err := bls.NewKeyPairFromString("0x01")
	require.NoError(t, err)

	request := elcontracts.RegistrationRequest{
		OperatorAddress: operatorAddr,
		AVSAddress:      avsAddr,
		OperatorSetIds:  []uint32{operatorSetId},
		WaitForReceipt:  true,
		Socket:          "socket",
		BlsKeyPair:      keypair,
	}

	registryCoordinatorAddress := contractAddrs.RegistryCoordinator
	receipt, err := chainWriter.RegisterForOperatorSets(
		context.Background(),
		registryCoordinatorAddress,
		request,
	)
	require.NoError(t, err)
	require.Equal(t, uint64(1), receipt.Status)

	allocationDelay := 1
	allocationMagnitude := 100
	allocationConfigurationDelay := 1200

	receipt, err = chainWriter.SetAllocationDelay(
		context.Background(),
		operatorAddr,
		uint32(allocationDelay),
		true,
	)
	require.NoError(t, err)
	require.Equal(t, gethtypes.ReceiptStatusSuccessful, receipt.Status)

	testutils.AdvanceChainByNBlocksExecInContainer(
		context.Background(),
		allocationConfigurationDelay+1,
		anvilC,
	)

	allocationParams := []allocationmanager.IAllocationManagerTypesAllocateParams{
		{
			OperatorSet:   operatorSet,
			Strategies:    strategies,
			NewMagnitudes: []uint64{uint64(allocationMagnitude)},
		},
	}

	receipt, err = chainWriter.ModifyAllocations(
		context.Background(),
		operatorAddr,
		allocationParams,
		true,
	)
	require.NoError(t, err)
	require.Equal(t, gethtypes.ReceiptStatusSuccessful, receipt.Status)

	t.Run("get operators and operator sets", func(t *testing.T) {
		t.Run("validate strategies for operatorSet", func(t *testing.T) {
			strats, err := chainReader.GetStrategiesForOperatorSet(context.Background(), operatorSet)
			require.NoError(t, err)
			require.Len(t, strats, 1)
			require.Equal(t, strats[0].Hex(), strategyAddr.Hex())
		})

		t.Run("get registered sets", func(t *testing.T) {
			registeredSets, err := chainReader.GetRegisteredSets(context.Background(), operatorAddr)
			require.NoError(t, err)
			require.NotEmpty(t, registeredSets)
		})

		t.Run("get operator sets for operator", func(t *testing.T) {
			opSets, err := chainReader.GetOperatorSetsForOperator(context.Background(), operatorAddr)
			require.NoError(t, err)
			require.NotEmpty(t, opSets)
		})

		t.Run("get amount operatorSets for operator", func(t *testing.T) {
			opSetsCount, err := chainReader.GetNumOperatorSetsForOperator(
				context.Background(),
				operatorAddr,
			)
			require.NoError(t, err)
			require.NotZero(t, opSetsCount)
		})

		t.Run("get operator for operatorsets", func(t *testing.T) {
			operators, err := chainReader.GetOperatorsForOperatorSet(context.Background(), operatorSet)
			require.NoError(t, err)
			require.NotEmpty(t, operators)
		})

		t.Run("get amount of operators for operatorsets", func(t *testing.T) {
			operatorsCount, err := chainReader.GetNumOperatorsForOperatorSet(context.Background(), operatorSet)
			require.NoError(t, err)
			require.NotZero(t, operatorsCount)
		})
	})

	t.Run("slashable shares tests", func(t *testing.T) {
		t.Run("get slashable shares for single operator", func(t *testing.T) {
			shares, err := chainReader.GetSlashableShares(
				context.Background(),
				operatorAddr,
				operatorSet,
				strategies,
			)
			require.NoError(t, err)
			require.NotEmpty(t, shares)
		})

		t.Run("get slashable shares for multiple operatorSets", func(t *testing.T) {
			shares, err := chainReader.GetSlashableSharesForOperatorSets(
				context.Background(),
				[]allocationmanager.OperatorSet{operatorSet},
			)
			require.NoError(t, err)
			require.NotEmpty(t, shares)
		})

		t.Run("get slashable shares before specific block", func(t *testing.T) {
			shares, err := chainReader.GetSlashableSharesForOperatorSetsBefore(
				context.Background(),
				[]allocationmanager.OperatorSet{operatorSet},
				2,
			)
			require.NoError(t, err)
			require.NotEmpty(t, shares)
		})
	})
}

func TestOperatorSetsWithWrongInput(t *testing.T) {
	_, anvilHttpEndpoint := testclients.BuildTestClients(t)
	ctx := context.Background()

	contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilHttpEndpoint)
	operatorAddr := common.HexToAddress(testutils.ANVIL_FIRST_ADDRESS)

	config := elcontracts.Config{}
	operatorSet := allocationmanager.OperatorSet{
		Avs: common.HexToAddress(testutils.ANVIL_SECOND_ADDRESS),
		Id:  0,
	}

	chainReader, err := testclients.NewTestChainReaderFromConfig(anvilHttpEndpoint, config)
	require.NoError(t, err)

	t.Run("test operator set with invalid id", func(t *testing.T) {
		_, err := chainReader.GetOperatorsForOperatorSet(ctx, operatorSet)
		require.Error(t, err)

		_, err = chainReader.GetNumOperatorsForOperatorSet(ctx, operatorSet)
		require.Error(t, err)

		_, err = chainReader.GetStrategiesForOperatorSet(ctx, operatorSet)
		require.Error(t, err)

		strategies := []common.Address{contractAddrs.Erc20MockStrategy}

		_, err = chainReader.GetSlashableShares(
			ctx,
			operatorAddr,
			operatorSet,
			strategies,
		)
		require.Error(t, err)
	})

	t.Run("get slashable shares with invalid operatorSet", func(t *testing.T) {
		config := elcontracts.Config{
			DelegationManagerAddress: contractAddrs.DelegationManager,
		}

		chainReader, err = testclients.NewTestChainReaderFromConfig(anvilHttpEndpoint, config)
		require.NoError(t, err)

		operatorSets := []allocationmanager.OperatorSet{operatorSet}

		_, err = chainReader.GetSlashableSharesForOperatorSetsBefore(context.Background(), operatorSets, 10)
		require.Error(t, err)
	})
}

// The idea is to cover some cases where network can fail by passing a
// cancelled context, so the binding returns an error
func TestFailingNetwork(t *testing.T) {
	clients, anvilHttpEndpoint := testclients.BuildTestReadClients(t)
	ctx := context.Background()

	subCtx, cancelFn := context.WithCancel(ctx)
	cancelFn()

	contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilHttpEndpoint)
	operator := types.Operator{
		Address: testutils.ANVIL_FIRST_ADDRESS,
	}

	t.Run("is operator registered", func(t *testing.T) {
		isOperator, err := clients.ElChainReader.IsOperatorRegistered(subCtx, operator)
		assert.Error(t, err)
		assert.False(t, isOperator)
	})

	t.Run("get operator details", func(t *testing.T) {
		operatorDetails, err := clients.ElChainReader.GetOperatorDetails(subCtx, operator)
		assert.Error(t, err)
		assert.Zero(t, operatorDetails)
	})

	t.Run("get strategy and underlying token", func(t *testing.T) {
		strategyAddr := contractAddrs.Erc20MockStrategy
		strategy, underlyingTokenAddr, err := clients.ElChainReader.GetStrategyAndUnderlyingToken(
			subCtx,
			strategyAddr,
		)
		assert.Error(t, err)
		assert.Nil(t, strategy)
		assert.Zero(t, underlyingTokenAddr)
	})

	t.Run("get strategy and underlying ERC20 token", func(t *testing.T) {
		strategyAddr := contractAddrs.Erc20MockStrategy
		strategy, contractUnderlyingToken, underlyingTokenAddr, err := clients.ElChainReader.GetStrategyAndUnderlyingERC20Token(
			subCtx,
			strategyAddr,
		)
		assert.Error(t, err)
		assert.Nil(t, strategy)
		assert.Zero(t, underlyingTokenAddr)
		assert.Nil(t, contractUnderlyingToken)
	})

	t.Run("get operator shares in strategy", func(t *testing.T) {
		shares, err := clients.ElChainReader.GetOperatorSharesInStrategy(
			subCtx,
			common.HexToAddress(operator.Address),
			contractAddrs.Erc20MockStrategy,
		)
		assert.Error(t, err)
		assert.Zero(t, shares)
	})

	t.Run("calculate delegation approval digest hash", func(t *testing.T) {
		staker := common.Address{0x0}
		delegationApprover := common.Address{0x0}
		approverSalt := [32]byte{}
		expiry := big.NewInt(0)
		digest, err := clients.ElChainReader.CalculateDelegationApprovalDigestHash(
			subCtx,
			staker,
			common.HexToAddress(operator.Address),
			delegationApprover,
			approverSalt,
			expiry,
		)
		assert.Error(t, err)
		assert.Empty(t, digest)
	})

	t.Run("calculate operator AVS registration digest hash", func(t *testing.T) {
		avs := common.Address{0x0}
		salt := [32]byte{}
		expiry := big.NewInt(0)
		digest, err := clients.ElChainReader.CalculateOperatorAVSRegistrationDigestHash(
			subCtx,
			common.HexToAddress(operator.Address),
			avs,
			salt,
			expiry,
		)
		assert.Error(t, err)
		assert.Empty(t, digest)
	})

	t.Run("get staker shares", func(t *testing.T) {
		strategies, shares, err := clients.ElChainReader.GetStakerShares(
			subCtx,
			common.HexToAddress(operator.Address),
		)
		assert.Empty(t, strategies)
		assert.Empty(t, shares)
		assert.Error(t, err)
	})

	t.Run("get delegated operator", func(t *testing.T) {
		blockNumber := big.NewInt(0)
		address, err := clients.ElChainReader.GetDelegatedOperator(
			subCtx,
			common.HexToAddress(operator.Address),
			blockNumber,
		)

		assert.Error(t, err)
		assert.Zero(t, address)
	})

	t.Run("GetOperatorShares", func(t *testing.T) {
		strategyAddr := contractAddrs.Erc20MockStrategy
		strategies := []common.Address{strategyAddr}
		shares, err := clients.ElChainReader.GetOperatorShares(
			subCtx,
			common.HexToAddress(operator.Address),
			strategies,
		)
		assert.Error(t, err)
		assert.Nil(t, shares)
	})

	t.Run("GetOperatorsShares", func(t *testing.T) {
		operatorAddr := common.HexToAddress(operator.Address)
		operators := []common.Address{operatorAddr}
		strategyAddr := contractAddrs.Erc20MockStrategy
		strategies := []common.Address{strategyAddr}
		shares, err := clients.ElChainReader.GetOperatorsShares(
			subCtx,
			operators,
			strategies,
		)
		assert.Error(t, err)
		assert.Nil(t, shares)
	})
}

func TestCreateRederFromConfig(t *testing.T) {
	_, anvilHttpEndpoint := testclients.BuildTestClients(t)
	testConfig := testutils.GetDefaultTestConfig()
	logger := logging.NewTextSLogger(os.Stdout, &logging.SLoggerOptions{Level: testConfig.LogLevel})

	contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilHttpEndpoint)

	ethHttpClient, err := ethclient.Dial(anvilHttpEndpoint)
	require.NoError(t, err)

	t.Run("create a reader client", func(t *testing.T) {
		config := elcontracts.Config{
			DelegationManagerAddress: contractAddrs.DelegationManager,
		}

		_, err = elcontracts.NewReaderFromConfig(config, ethHttpClient, logger)
		require.NoError(t, err)
	})

	t.Run("try to create a reader with an invalid config", func(t *testing.T) {
		config := elcontracts.Config{
			DelegationManagerAddress: common.HexToAddress(testutils.ANVIL_FIRST_ADDRESS),
		}

		_, err = elcontracts.NewReaderFromConfig(config, ethHttpClient, logger)
		require.Error(t, err)
	})
}
