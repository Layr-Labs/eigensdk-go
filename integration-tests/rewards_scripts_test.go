package integration_test

import (
	"context"
	"crypto/sha256"
	"errors"
	"math/big"
	"testing"

	strategy "github.com/Layr-Labs/eigensdk-go/contracts/bindings/IStrategy"
	mockerc20 "github.com/Layr-Labs/eigensdk-go/contracts/bindings/MockERC20"
	rewardsCoordinator "github.com/Layr-Labs/eigensdk-go/contracts/bindings/RewardsCoordinator"
	servicemanager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/ServiceManagerBase"
	"github.com/Layr-Labs/eigensdk-go/testutils"
	"github.com/Layr-Labs/eigensdk-go/testutils/testclients"
	sdkutils "github.com/Layr-Labs/eigensdk-go/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/sha3"
)

func TestIntegrationRewards(t *testing.T) {
	// Test set up
	clients, anvilHttpEndpoint := testclients.BuildTestClients(t)
	contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilHttpEndpoint)

	contractStrategy, err := strategy.NewContractIStrategy(contractAddrs.Erc20MockStrategy, clients.EthHttpClient)
	require.NoError(t, err)

	tokenAddr, err := contractStrategy.UnderlyingToken(&bind.CallOpts{Context: context.Background()})
	require.NoError(t, err)

	mockToken, err := mockerc20.NewContractMockERC20(tokenAddr, clients.EthHttpClient)
	require.NoError(t, err)

	txMgr := clients.TxManager
	noSendTxOpts, err := txMgr.GetNoSendTxOpts()
	require.NoError(t, err)

	amountPerPayment := int64(100)
	numPayments := int64(8)

	// Initially, claimer balance in strategy is zero
	initialBalance, err := mockToken.BalanceOf(&bind.CallOpts{}, common.HexToAddress("0x0000000000000000000000000000000000000001"))
	require.NoError(t, err)
	assert.Zero(t, initialBalance.Int64())

	// Create AVS rewards submission
	stratAndMul := []servicemanager.IRewardsCoordinatorTypesStrategyAndMultiplier{
		{
			Strategy:   contractAddrs.Erc20MockStrategy,
			Multiplier: big.NewInt(1_000_000),
		},
	}

	calculationInterval, err := clients.EigenlayerContractBindings.RewardsCoordinator.CALCULATIONINTERVALSECONDS(nil)
	require.NoError(t, err)

	header, err := clients.EthHttpClient.HeaderByNumber(context.Background(), nil)
	require.NoError(t, err)

	// These values are set to align with the contract's requirements for the `OperatorDirectedRewardsSubmission`.
	// https://github.com/Layr-Labs/eigenlayer-contracts/blob/ecaff6304de6cb0f43b42024ad55d0e8a0430790/src/contracts/core/RewardsCoordinator.sol#L414
	// https://github.com/Layr-Labs/eigenlayer-contracts/blob/ecaff6304de6cb0f43b42024ad55d0e8a0430790/src/contracts/core/RewardsCoordinator.sol#L482
	var duration uint32 = calculationInterval
	var startTimestamp uint32 = ((uint32(header.Time) / calculationInterval) + 1) * calculationInterval

	// These values were taken from Go Incredible Squaring AVS's rewards scripts
	rewardsSubmission := []servicemanager.IRewardsCoordinatorTypesRewardsSubmission{
		{
			StrategiesAndMultipliers: stratAndMul,
			Token:                    tokenAddr,
			Amount:                   big.NewInt(amountPerPayment),
			StartTimestamp:           startTimestamp,
			Duration:                 duration,
		},
	}

	receipt, err := clients.AvsRegistryChainWriter.CreateAVSRewardsSubmission(context.Background(), rewardsSubmission, true)
	require.NoError(t, err)
	require.Equal(t, receipt.Status, gethtypes.ReceiptStatusSuccessful)

	// Submit the root for the submission
	contractRewardsCoordinator, err := rewardsCoordinator.NewContractRewardsCoordinator(contractAddrs.RewardsCoordinator, clients.EthHttpClient)
	require.NoError(t, err)

	tx, err := contractRewardsCoordinator.SetActivationDelay(noSendTxOpts, 0)
	require.NoError(t, err)

	receipt, err = txMgr.Send(context.Background(), tx, true)
	require.NoError(t, err)
	require.Equal(t, receipt.Status, gethtypes.ReceiptStatusSuccessful)

	tokenLeaves := createTokenLeaves(contractRewardsCoordinator, 1, 100, tokenAddr)
	require.NoError(t, err)

	deployerAddress := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	earners := getEarners(deployerAddress, int(numPayments))
	earnerLeaves := createEarnerLeaves(earners, tokenLeaves)
	require.NoError(t, err)

	leaves, root, err := createPaymentRoot(contractRewardsCoordinator, earnerLeaves, int(numPayments))
	require.NoError(t, err)

	tx, err = contractRewardsCoordinator.SubmitRoot(noSendTxOpts, root, 100)
	require.NoError(t, err)

	receipt, err = txMgr.Send(context.Background(), tx, true)
	require.NoError(t, err)
	require.Equal(t, receipt.Status, gethtypes.ReceiptStatusSuccessful)

	// Process the claim for the submitted root (and claim the rewards for the claimer)
	indexToProve := 0

	proof, err := generateEarnerMerkleProof(leaves, indexToProve)
	require.NoError(t, err)

	tokenProof := generateTokenMerkleProof(tokenLeaves, indexToProve)
	require.NoError(t, err)

	tokenIndices := make([]uint32, 1)
	tokenProofs := make([][]byte, 1)
	tokenProofs[0] = tokenProof

	newTokenLeaves := make([]rewardsCoordinator.IRewardsCoordinatorTypesTokenTreeMerkleLeaf, 1)
	newTokenLeaves[0] = defaultTokenLeaf(100, tokenAddr)

	// this workflow assumes a new root submitted for every payment claimed.  So we get the latest rood index to process a claim for
	rootLength, err := contractRewardsCoordinator.GetDistributionRootsLength(&bind.CallOpts{})
	require.NoError(t, err)
	rootIndex := rootLength.Uint64() - 1

	claim := rewardsCoordinator.IRewardsCoordinatorTypesRewardsMerkleClaim{
		RootIndex:       uint32(rootIndex),
		EarnerIndex:     0,
		EarnerTreeProof: proof,
		EarnerLeaf:      earnerLeaves[indexToProve],
		TokenIndices:    tokenIndices,
		TokenTreeProofs: tokenProofs,
		TokenLeaves:     newTokenLeaves,
	}

	allocationConfigurationDelay := 1200
	testutils.AdvanceChainByNBlocksExecInContainer(
		context.Background(),
		allocationConfigurationDelay+1,
		clients.AnvilC,
	)

	receipt, err = clients.ElChainWriter.ProcessClaim(context.Background(), claim, common.HexToAddress("0x01"), true)
	require.NoError(t, err)
	require.Equal(t, receipt.Status, gethtypes.ReceiptStatusSuccessful)

	// After claim, claimer balance in strategy is zero
	balanceAfterClaim, err := mockToken.BalanceOf(&bind.CallOpts{}, common.HexToAddress("0x0000000000000000000000000000000000000001"))
	require.NoError(t, err)
	assert.Equal(t, balanceAfterClaim.Int64(), amountPerPayment)
}

// These utils were inspired in those used in rewards scripts in Go Inc Squaring:
// https://github.com/Layr-Labs/incredible-squaring-avs/blob/dev/contracts/script/utils/SetupDistributionsLib.sol
func getEarners(deployer common.Address, numPayments int) []common.Address {
	earners := make([]common.Address, numPayments)
	for i := range earners {
		earners[i] = deployer
	}
	return earners
}

func createEarnerLeaves(earners []common.Address, tokenLeaves [][32]byte) []rewardsCoordinator.IRewardsCoordinatorTypesEarnerTreeMerkleLeaf {
	leaves := make([]rewardsCoordinator.IRewardsCoordinatorTypesEarnerTreeMerkleLeaf, len(earners))
	tokenRoot := createTokenRoot(tokenLeaves)

	for i, earner := range earners {
		leaves[i] = rewardsCoordinator.IRewardsCoordinatorTypesEarnerTreeMerkleLeaf{
			Earner:          earner,
			EarnerTokenRoot: tokenRoot,
		}
	}
	return leaves
}

func createTokenLeaves(
	rewardsCoordinator *rewardsCoordinator.ContractRewardsCoordinator,
	numberOfTokenLeaves int,
	tokensEarned uint64,
	tokenAddr common.Address,
) [][32]byte {
	leaves := make([][32]byte, numberOfTokenLeaves)

	for i := 0; i < numberOfTokenLeaves; i++ {
		leaf := defaultTokenLeaf(tokensEarned, tokenAddr)
		leafBytes, err := rewardsCoordinator.CalculateTokenLeafHash(&bind.CallOpts{}, leaf)
		if err != nil {
			panic(sdkutils.WrapError("Failed to call CalculateEarnerLeafHash", err))
		}
		leaves[i] = leafBytes
	}
	return leaves
}

func defaultTokenLeaf(
	tokensEarned uint64,
	tokenAddr common.Address,
) rewardsCoordinator.IRewardsCoordinatorTypesTokenTreeMerkleLeaf {
	return rewardsCoordinator.IRewardsCoordinatorTypesTokenTreeMerkleLeaf{
		Token:              tokenAddr,
		CumulativeEarnings: big.NewInt(int64(tokensEarned)),
	}
}

func createPaymentRoot(
	rewardsCoordinator *rewardsCoordinator.ContractRewardsCoordinator,
	earnerLeaves []rewardsCoordinator.IRewardsCoordinatorTypesEarnerTreeMerkleLeaf,
	numPayments int,
) ([][32]byte, [32]byte, error) {
	leaves := make([][32]byte, numPayments)
	for i := 0; i < numPayments; i++ {
		leaf, err := rewardsCoordinator.CalculateEarnerLeafHash(&bind.CallOpts{}, earnerLeaves[i])
		if err != nil {
			return [][32]byte{}, [32]byte{}, sdkutils.WrapError("Failed to call CalculateEarnerLeafHash", err)
		}
		leaves[i] = leaf
	}

	//writeLeavesToJson(leaves, tokenLeaves, filePath)
	return leaves, merkleizeKeccak(leaves), nil
}

func generateTokenMerkleProof(leaves [][32]byte, index int) []byte {
	if len(leaves) == 0 {
		panic("leaves array cannot be empty")
	}
	if index < 0 || index >= len(leaves) {
		panic("index out of bounds")
	}

	leaves = padLeaves(leaves)

	n := len(leaves)
	depth := 0
	for (1 << depth) < n {
		depth++
	}

	var proofBytes []byte

	for i := 0; i < depth; i++ {
		levelSize := (n + 1) / 2
		siblingIndex := index ^ 1

		if siblingIndex < n {
			proofBytes = append(proofBytes, leaves[siblingIndex][:]...)
		}

		for j := 0; j < levelSize; j++ {
			if 2*j+1 < n {
				concatenated := append(leaves[2*j][:], leaves[2*j+1][:]...)
				hash := sha256.Sum256(concatenated)
				leaves[j] = hash
			} else {
				leaves[j] = leaves[2*j]
			}
		}

		n = levelSize
		index /= 2
	}

	return proofBytes
}

func merkleizeKeccak(leaves [][32]byte) [32]byte {
	leaves = padLeaves(leaves)

	numNodesInLayer := len(leaves) / 2
	layer := make([][32]byte, numNodesInLayer)

	for i := 0; i < numNodesInLayer; i++ {
		layer[i] = keccak256(append(leaves[2*i][:], leaves[2*i+1][:]...))
	}

	numNodesInLayer /= 2
	for numNodesInLayer != 0 {
		for i := 0; i < numNodesInLayer; i++ {
			layer[i] = keccak256(append(layer[2*i][:], layer[2*i+1][:]...))
		}
		numNodesInLayer /= 2
	}

	return layer[0]
}

func padLeaves(leaves [][32]byte) [][32]byte {
	paddedLength := 2
	for paddedLength < len(leaves) {
		paddedLength *= 2
	}

	paddedLeaves := make([][32]byte, paddedLength)
	copy(paddedLeaves, leaves)
	return paddedLeaves
}

func keccak256(data []byte) [32]byte {
	var hash [32]byte

	h := sha3.NewLegacyKeccak256()

	h.Write(data)

	h.Sum(hash[:0])

	return hash
}

func createTokenRoot(tokenLeaves [][32]byte) [32]byte {
	return merkleizeKeccak(tokenLeaves)
}

func generateEarnerMerkleProof(leaves [][32]byte, index int) ([]byte, error) {
	if len(leaves) == 0 {
		return nil, errors.New("empty leaves array")
	}

	if index < 0 || index >= len(leaves) {
		return nil, errors.New("index out of bounds")
	}

	tree := buildMerkleTree(leaves)

	proof := make([]byte, 0)
	currentIndex := index

	for level := 0; level < len(tree)-1; level++ {
		levelNodes := tree[level]
		levelLength := len(levelNodes)

		var sibling [32]byte
		if currentIndex%2 == 1 {
			sibling = levelNodes[currentIndex-1]
		} else {
			if currentIndex+1 < levelLength {
				sibling = levelNodes[currentIndex+1]
			} else {
				currentIndex = currentIndex / 2
				continue
			}
		}

		proof = append(proof, sibling[:]...)

		currentIndex = currentIndex / 2
	}

	return proof, nil
}

func buildMerkleTree(leaves [][32]byte) [][][32]byte {
	tree := make([][][32]byte, 0)
	tree = append(tree, leaves)

	for len(tree[len(tree)-1]) > 1 {
		lastLevel := tree[len(tree)-1]
		newLevel := make([][32]byte, 0)

		for i := 0; i < len(lastLevel); i += 2 {
			var left, right [32]byte
			left = lastLevel[i]

			if i+1 < len(lastLevel) {
				right = lastLevel[i+1]
			} else {
				right = left
			}

			combined := append(left[:], right[:]...)
			parent := keccak256(combined)
			newLevel = append(newLevel, parent)
		}

		tree = append(tree, newLevel)
	}

	return tree
}
