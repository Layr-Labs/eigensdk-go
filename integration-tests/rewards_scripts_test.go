package integration_test

import (
	"context"
	"crypto/sha256"
	"fmt"
	"log"
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
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIntegrationRewards(t *testing.T) {
	log.Println("This test takes ~50 seconds to run...")

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

	// This is the avsRegistry writer address (anvil nineth address)
	tx, err := mockToken.Mint(noSendTxOpts, common.HexToAddress("0xa0Ee7A142d267C1f36714E4a8F75612F20a79720"), big.NewInt(1000))
	require.NoError(t, err)

	receipt, err := txMgr.Send(context.Background(), tx, true)
	require.NoError(t, err)
	require.Equal(t, receipt.Status, uint64(1))

	amountPerPayment := int64(100)
	numPayments := int64(8)
	tx, err = mockToken.IncreaseAllowance(noSendTxOpts, contractAddrs.ServiceManager,
		big.NewInt(amountPerPayment*numPayments),
	)
	require.NoError(t, err)

	receipt, err = txMgr.Send(context.Background(), tx, true)
	require.NoError(t, err)
	require.Equal(t, receipt.Status, uint64(1))

	// Initially, claimer balance in strategy is zero
	initialBalance, err := mockToken.BalanceOf(&bind.CallOpts{}, common.HexToAddress("0x0000000000000000000000000000000000000001"))
	require.NoError(t, err)
	assert.Zero(t, initialBalance.Int64())

	// Now call the rewards scripts or related and assert the balance has changed.
	stratAndMul := []servicemanager.IRewardsCoordinatorTypesStrategyAndMultiplier{
		{
			Strategy:   contractAddrs.Erc20MockStrategy,
			Multiplier: big.NewInt(1_000_000),
		},
	}
	// These values were taken from Go Incredible Squaring AVS's rewards scripts
	rewardsSubmission := []servicemanager.IRewardsCoordinatorTypesRewardsSubmission{
		{
			StrategiesAndMultipliers: stratAndMul,
			Token:                    tokenAddr,
			Amount:                   big.NewInt(100),
			StartTimestamp:           1743033600,
			Duration:                 6048000,
		},
	}

	receipt, err = clients.AvsRegistryChainWriter.CreateAVSRewardsSubmission(context.Background(), rewardsSubmission, true)
	require.NoError(t, err)
	require.Equal(t, receipt.Status, uint64(1))

	contractRewardsCoordinator, err := rewardsCoordinator.NewContractRewardsCoordinator(contractAddrs.RewardsCoordinator, clients.EthHttpClient)
	require.NoError(t, err)

	tx, err = contractRewardsCoordinator.SetActivationDelay(noSendTxOpts, 0)
	require.NoError(t, err)

	receipt, err = txMgr.Send(context.Background(), tx, true)
	require.NoError(t, err)
	require.Equal(t, receipt.Status, uint64(1))

	tokenLeaves, err := CreateTokenLeaves(contractRewardsCoordinator, 1, 100, tokenAddr)
	require.NoError(t, err)

	earners := getEarners(common.HexToAddress("0x01")) // Maybe should be the deployer addr
	earnerLeaves := CreateEarnerLeaves(earners, tokenLeaves)
	require.NoError(t, err)

	leaves, root, err := createPaymentRoot(contractRewardsCoordinator, tokenLeaves, earnerLeaves, 8, 1)
	require.NoError(t, err)

	tx, err = contractRewardsCoordinator.SubmitRoot(noSendTxOpts, root, 100)
	require.NoError(t, err)

	receipt, err = txMgr.Send(context.Background(), tx, true)
	require.NoError(t, err)
	require.Equal(t, receipt.Status, uint64(1))

	// this workflow assumes a new root submitted for every payment claimed.  So we get the latest rood index to process a claim for
	rootLength, err := contractRewardsCoordinator.GetDistributionRootsLength(&bind.CallOpts{})
	require.NoError(t, err)
	rootIndex := rootLength.Uint64() - 1

	indexToProve := 0
	proof, err := generateMerkleProof(leaves, indexToProve)
	require.NoError(t, err)

	tokenProof, err := generateMerkleProof(tokenLeaves, indexToProve)
	require.NoError(t, err)

	tokenIndices := make([]uint32, 1)
	tokenProofs := make([][]byte, 1)
	tokenProofs[0] = tokenProof

	newTokenLeaves := make([]rewardsCoordinator.IRewardsCoordinatorTypesTokenTreeMerkleLeaf, 1)
	newTokenLeaves[0] = defaultTokenLeaf(100, contractAddrs.Erc20MockStrategy)

	claim := rewardsCoordinator.IRewardsCoordinatorTypesRewardsMerkleClaim{
		RootIndex:       uint32(rootIndex),
		EarnerIndex:     0,
		EarnerTreeProof: proof,
		EarnerLeaf:      earnerLeaves[indexToProve],
		TokenIndices:    tokenIndices,
		TokenTreeProofs: tokenProofs,
		TokenLeaves:     newTokenLeaves,
	}

	testutils.AdvanceChainByNBlocks(1, anvilHttpEndpoint)

	receipt, err = clients.ElChainWriter.ProcessClaim(context.Background(), claim, common.HexToAddress("0x01"), true)
	require.NoError(t, err)
	require.Equal(t, receipt.Status, uint64(1))
}

// These utils were inspired in those used in rewards scripts in Go Inc Squaring:
// https://github.com/Layr-Labs/incredible-squaring-avs/blob/dev/contracts/script/utils/SetupDistributionsLib.sol
func getEarners(deployer common.Address) []common.Address {
	earners := make([]common.Address, 8)
	for i := range earners {
		earners[i] = deployer
	}
	return earners
}

func CreateEarnerLeaves(earners []common.Address, tokenLeaves [][32]byte) []rewardsCoordinator.IRewardsCoordinatorTypesEarnerTreeMerkleLeaf {
	leaves := make([]rewardsCoordinator.IRewardsCoordinatorTypesEarnerTreeMerkleLeaf, len(earners))
	tokenRoot := CreateTokenRoot(tokenLeaves)

	for i, earner := range earners {
		leaves[i] = rewardsCoordinator.IRewardsCoordinatorTypesEarnerTreeMerkleLeaf{
			Earner:          earner,
			EarnerTokenRoot: tokenRoot,
		}
	}
	return leaves
}

func CreateTokenLeaves(
	rewardsCoordinator *rewardsCoordinator.ContractRewardsCoordinator,
	numTokenEarnings int,
	tokenEarnings uint64,
	tokenAddr common.Address,
) ([][32]byte, error) {
	leaves := make([][32]byte, numTokenEarnings)

	for i := 0; i < numTokenEarnings; i++ {
		leaf := defaultTokenLeaf(tokenEarnings, tokenAddr)
		leafBytes, err := rewardsCoordinator.CalculateTokenLeafHash(&bind.CallOpts{}, leaf)
		if err != nil {
			return [][32]byte{}, sdkutils.WrapError("Failed to call CalculateEarnerLeafHash", err)
		}
		leaves[i] = leafBytes
	}
	return leaves, nil
}

func defaultTokenLeaf(tokenEarnings uint64, tokenAddr common.Address) rewardsCoordinator.IRewardsCoordinatorTypesTokenTreeMerkleLeaf {
	return rewardsCoordinator.IRewardsCoordinatorTypesTokenTreeMerkleLeaf{
		Token:              tokenAddr,
		CumulativeEarnings: big.NewInt(int64(tokenEarnings)),
	}
}

func createPaymentRoot(
	rewardsCoordinator *rewardsCoordinator.ContractRewardsCoordinator,
	tokenLeaves [][32]byte,
	earnerLeaves []rewardsCoordinator.IRewardsCoordinatorTypesEarnerTreeMerkleLeaf,
	numPayments int,
	numTokenEarnings int,
) ([][32]byte, [32]byte, error) {
	if len(earnerLeaves) != numPayments {
		return [][32]byte{}, [32]byte{}, fmt.Errorf("Number of earners must match number of payments")
	}
	if len(tokenLeaves) != numTokenEarnings {
		return [][32]byte{}, [32]byte{}, fmt.Errorf("Number of token leaves must match number of token earnings")
	}

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

func generateMerkleProof(leaves [][32]byte, index int) ([]byte, error) {
	if len(leaves) == 0 {
		return nil, fmt.Errorf("leaves array cannot be empty")
	}
	if index < 0 || index >= len(leaves) {
		return nil, fmt.Errorf("index out of bounds")
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

	return proofBytes, nil
}

func merkleizeKeccak(leaves [][32]byte) [32]byte {
	leaves = padLeaves(leaves)

	numNodesInLayer := len(leaves) / 2
	layer := make([][32]byte, numNodesInLayer)

	for i := 0; i < numNodesInLayer; i++ {
		layer[i] = keccak256(append(leaves[2*i][:], leaves[2*i+1][:]...))
	}

	for numNodesInLayer > 1 {
		numNodesInLayer /= 2
		for i := 0; i < numNodesInLayer; i++ {
			layer[i] = keccak256(append(layer[2*i][:], layer[2*i+1][:]...))
		}
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
	hash := sha256.Sum256(data)
	return hash
}

func CreateTokenRoot(tokenLeaves [][32]byte) [32]byte {
	return merkleizeKeccak(tokenLeaves)
}
