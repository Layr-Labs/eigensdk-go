package avsregistry_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/elcontracts"
	chainioutils "github.com/Layr-Labs/eigensdk-go/chainio/utils"
	avsdirectory "github.com/Layr-Labs/eigensdk-go/contracts/bindings/AVSDirectory"
	regcoord "github.com/Layr-Labs/eigensdk-go/contracts/bindings/RegistryCoordinator"
	servicemanager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/ServiceManagerBase"
	stakeregistry "github.com/Layr-Labs/eigensdk-go/contracts/bindings/StakeRegistry"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/testutils"
	"github.com/Layr-Labs/eigensdk-go/testutils/testclients"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethcommon "github.com/ethereum/go-ethereum/common"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWriterMethods(t *testing.T) {
	clients, anvilHttpEndpoint := testclients.BuildTestClients(t)
	contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilHttpEndpoint)

	chainWriter := clients.AvsRegistryChainWriter

	keypair, err := bls.NewKeyPairFromString("0x01")
	require.NoError(t, err)

	addr := gethcommon.HexToAddress(testutils.ANVIL_FIRST_ADDRESS)
	ecdsaPrivateKey, err := crypto.HexToECDSA(testutils.ANVIL_FIRST_PRIVATE_KEY)
	require.NoError(t, err)

	quorumNumbers := types.QuorumNums{0}
	quorumNumber := types.QuorumNum(0)

	subCtx, cancelFn := context.WithCancel(context.Background())
	cancelFn()

	t.Run("update socket without being registered", func(t *testing.T) {
		receipt, err := chainWriter.UpdateSocket(
			context.Background(),
			types.Socket("102901920192019201902910291209"),
			true,
		)
		assert.Error(t, err)
		assert.Nil(t, receipt)
	})

	otherKeyPair, err := bls.NewKeyPairFromString("0x01")
	require.NoError(t, err)
	request := elcontracts.RegistrationRequest{
		OperatorAddress: addr,
		AVSAddress:      contractAddrs.ServiceManager,
		OperatorSetIds:  []uint32{0},
		WaitForReceipt:  true,
		Socket:          "socket",
		BlsKeyPair:      otherKeyPair,
	}

	// Register operator
	elWriter := clients.ElChainWriter
	receipt, err := elWriter.RegisterForOperatorSets(context.Background(), contractAddrs.RegistryCoordinator, request)
	require.NoError(t, err)
	require.NotNil(t, receipt)

	t.Run("update stake of operator subset", func(t *testing.T) {
		receipt, err := chainWriter.UpdateStakesOfOperatorSubsetForAllQuorums(
			context.Background(),
			[]gethcommon.Address{addr},
			true,
		)
		require.NoError(t, err)
		require.NotNil(t, receipt)
	})

	t.Run("update stake of entire operator set", func(t *testing.T) {
		receipt, err := chainWriter.UpdateStakesOfEntireOperatorSetForQuorums(
			context.Background(),
			[][]gethcommon.Address{{addr}},
			quorumNumbers,
			true,
		)
		require.NoError(t, err)
		require.NotNil(t, receipt)
	})

	t.Run("update socket", func(t *testing.T) {
		receipt, err = chainWriter.UpdateSocket(
			context.Background(),
			types.Socket(""),
			true,
		)
		require.NoError(t, err)
		require.NotNil(t, receipt)
	})

	t.Run("set rewards initiator", func(t *testing.T) {
		// Set up to create a ServiceManager binding
		ethHttpClient, err := ethclient.Dial(anvilHttpEndpoint)
		require.NoError(t, err)

		serviceManager, err := servicemanager.NewContractServiceManagerBase(
			contractAddrs.ServiceManager,
			ethHttpClient,
		)
		require.NoError(t, err)

		// Check that, first, initiator address is anvil first address
		initiator_initial_addr, err := serviceManager.RewardsInitiator(&bind.CallOpts{})
		require.NoError(t, err)
		assert.Equal(t, initiator_initial_addr.String(), testutils.ANVIL_FIRST_ADDRESS)

		// Modify initiator address, set it as anvil second address
		new_initiator_addr := gethcommon.HexToAddress(testutils.ANVIL_SECOND_ADDRESS)
		receipt, err := chainWriter.SetRewardsInitiator(
			context.Background(),
			new_initiator_addr,
			true,
		)
		require.NoError(t, err)
		require.NotNil(t, receipt)

		// Check that, after modifying it, initiator address is now anvil second address
		initiator_modified_addr, err := serviceManager.RewardsInitiator(&bind.CallOpts{})
		require.NoError(t, err)
		assert.Equal(t, initiator_modified_addr, new_initiator_addr)
	})

	// Error cases
	t.Run("fail register operator cancelling context", func(t *testing.T) {
		receipt, err := chainWriter.RegisterOperator(
			subCtx,
			ecdsaPrivateKey,
			keypair,
			quorumNumbers,
			"",
			true,
		)
		assert.Error(t, err)
		assert.Nil(t, receipt)
	})

	t.Run("fail update stake of operator subset cancelling context", func(t *testing.T) {
		receipt, err := chainWriter.UpdateStakesOfOperatorSubsetForAllQuorums(
			subCtx,
			[]gethcommon.Address{addr},
			true,
		)
		assert.Error(t, err)
		assert.Nil(t, receipt)
	})

	t.Run("fail update stake of entire operator set cancelling context", func(t *testing.T) {
		receipt, err := chainWriter.UpdateStakesOfEntireOperatorSetForQuorums(
			subCtx,
			[][]gethcommon.Address{{addr}},
			quorumNumbers,
			true,
		)
		assert.Error(t, err)
		assert.Nil(t, receipt)
	})

	t.Run("fail update stake of entire operator set because of quorum length", func(t *testing.T) {
		// Fails because operators per quorum length is distinct from quorum numbers
		receipt, err := chainWriter.UpdateStakesOfEntireOperatorSetForQuorums(
			context.Background(),
			[][]gethcommon.Address{{addr, addr}},
			quorumNumbers,
			true,
		)
		assert.Error(t, err)
		assert.Nil(t, receipt)
	})

	t.Run("fail deregister operator cancelling context", func(t *testing.T) {
		receipt, err := chainWriter.DeregisterOperator(
			subCtx,
			quorumNumbers,
			chainioutils.ConvertToBN254G1Point(keypair.PubKey),
			true,
		)
		assert.Error(t, err)
		assert.Nil(t, receipt)
	})

	t.Run("fail update socket cancelling context", func(t *testing.T) {
		receipt, err := chainWriter.UpdateSocket(
			subCtx,
			types.Socket(""),
			true,
		)
		assert.Error(t, err)
		assert.Nil(t, receipt)
	})

	t.Run("set slashable stake lookahead", func(t *testing.T) {
		operatorSetParams := regcoord.ISlashingRegistryCoordinatorTypesOperatorSetParam{
			MaxOperatorCount: 5,
		}
		minimumStakeNeeded := big.NewInt(0)

		strategyAddr := contractAddrs.Erc20MockStrategy
		strategyParam := regcoord.IStakeRegistryTypesStrategyParams{
			Strategy:   strategyAddr,
			Multiplier: big.NewInt(1e18),
		}

		lookAheadPeriod := uint32(0)

		receipt, err = chainWriter.CreateSlashableStakeQuorum(
			context.Background(),
			operatorSetParams,
			minimumStakeNeeded,
			[]regcoord.IStakeRegistryTypesStrategyParams{strategyParam},
			lookAheadPeriod,
			true,
		)

		// When not set, lookAheadPeriod is Zero
		lookAheadPeriod, err := clients.AvsRegistryChainReader.GetSlashableStakeLookAheadPerQuorum(&bind.CallOpts{}, 1)
		require.NoError(t, err)
		assert.Zero(t, lookAheadPeriod)

		// Modify lookAheadPeriod, set it as 32
		newLookAheadPeriod := 32
		receipt, err := chainWriter.SetSlashableStakeLookahead(
			context.Background(),
			1,
			uint32(newLookAheadPeriod),
			true,
		)
		require.NoError(t, err)
		require.NotNil(t, receipt)

		// After modify, lookAheadPeriod's value is 32
		lookAheadPeriod, err = clients.AvsRegistryChainReader.GetSlashableStakeLookAheadPerQuorum(&bind.CallOpts{}, 1)
		require.NoError(t, err)

		assert.Equal(t, lookAheadPeriod, uint32(newLookAheadPeriod))
	})

	t.Run("set minimum stake for quorum", func(t *testing.T) {
		// Create stakeRegistry contract
		ethHttpClient, err := ethclient.Dial(anvilHttpEndpoint)
		require.NoError(t, err)

		contractRegistryCoordinator, err := regcoord.NewContractRegistryCoordinator(
			contractAddrs.RegistryCoordinator,
			ethHttpClient,
		)
		require.NoError(t, err)

		stakeRegistryAddr, err := contractRegistryCoordinator.StakeRegistry(&bind.CallOpts{})
		require.NoError(t, err)

		stakeRegistry, err := stakeregistry.NewContractStakeRegistry(
			stakeRegistryAddr,
			ethHttpClient,
		)
		require.NoError(t, err)
		receipt, err := chainWriter.SetMinimumStakeForQuorum(
			context.Background(),
			quorumNumber.UnderlyingType(),
			big.NewInt(100),
			true,
		)
		require.NoError(t, err)
		require.NotNil(t, receipt)

		newMinimumStakeForQuorum, err := stakeRegistry.MinimumStakeForQuorum(
			&bind.CallOpts{},
			quorumNumber.UnderlyingType(),
		)
		require.NoError(t, err)

		assert.Equal(t, newMinimumStakeForQuorum, big.NewInt(100))
	})
}

/*
This test is commented because we need to use the new flow functions, and RegisterOperatorWithChurn belongs to the old one. We can
use RegisterOperatorForOperatorSet to register with churn, but we should expose a function registerOperatorForOperatorSetsWithChurn

func TestRegisterOperatorWithChurn(t *testing.T) {
	clients, anvilHttpEndpoint := testclients.BuildTestClients(t)

	contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilHttpEndpoint)

	chainWriter := clients.AvsRegistryChainWriter
	chainReader := clients.AvsRegistryChainReader

	firstOperatorAddress := gethcommon.HexToAddress(testutils.ANVIL_FIRST_ADDRESS)
	firstOperatorECDSAPrivateKey, err := crypto.HexToECDSA(testutils.ANVIL_FIRST_PRIVATE_KEY)
	require.NoError(t, err)
	firstOperatorKeyPair, err := bls.NewKeyPairFromString("0x01")
	require.NoError(t, err)

	quorumNumbers := types.QuorumNums{0}

	ethHttpClient := clients.EthHttpClient

	registryCoordinatorContract, err := regcoord.NewContractRegistryCoordinator(
		contractAddrs.RegistryCoordinator,
		ethHttpClient,
	)
	require.NoError(t, err)

	// At first, churnApprover is ANVIL_FIRST_ADDRESS
	approver, err := registryCoordinatorContract.ChurnApprover(&bind.CallOpts{})
	require.NoError(t, err)
	assert.Equal(t, approver.String(), testutils.ANVIL_FIRST_ADDRESS)

	// Set ANVIL_SECOND_ADDRESS as the new churnApprover
	churnApproverAddress := gethcommon.HexToAddress(testutils.ANVIL_SECOND_ADDRESS)
	churnECDSAPrivateKey, err := crypto.HexToECDSA(testutils.ANVIL_SECOND_PRIVATE_KEY)
	require.NoError(t, err)

	receipt, err := chainWriter.SetChurnApprover(context.Background(), churnApproverAddress, true)
	require.NoError(t, err)
	require.Equal(t, receipt.Status, gethtypes.ReceiptStatusSuccessful)

	// After change, churnApprover is ANVIL_SECOND_ADDRESS
	newApprover, err := registryCoordinatorContract.ChurnApprover(&bind.CallOpts{})
	require.NoError(t, err)
	assert.Equal(t, newApprover.String(), testutils.ANVIL_SECOND_ADDRESS)

	//Register ANVIL_FIRST_ADDRESS as operator
	receipt, err = chainWriter.RegisterOperator(
		context.Background(),
		firstOperatorECDSAPrivateKey,
		firstOperatorKeyPair,
		quorumNumbers,
		"",
		true,
	)
	require.NoError(t, err)
	require.NotNil(t, receipt)

	// Change the OperatorSetParams to allow only 1 operator
	receipt, err = chainWriter.SetOperatorSetParams(
		context.Background(),
		0,
		regcoord.ISlashingRegistryCoordinatorTypesOperatorSetParam{
			MaxOperatorCount:        1,
			KickBIPsOfOperatorStake: 10,
			KickBIPsOfTotalStake:    10000,
		},
		true,
	)
	require.NoError(t, err)
	require.NotNil(t, receipt)

	// We want to kick the first operator
	operatorsToKick := []gethcommon.Address{firstOperatorAddress}

	thirdOperatorAddress := gethcommon.HexToAddress(testutils.ANVIL_THIRD_ADDRESS)
	thirdOperatorECDSAPrivateKey, err := crypto.HexToECDSA(testutils.ANVIL_THIRD_PRIVATE_KEY)
	require.NoError(t, err)
	thirdOperatorPrivateKey := testutils.ANVIL_THIRD_PRIVATE_KEY
	newKeyPair, err := bls.NewKeyPairFromString("0x03")
	require.NoError(t, err)

	config := avsregistry.Config{
		RegistryCoordinatorAddress:    contractAddrs.RegistryCoordinator,
		OperatorStateRetrieverAddress: contractAddrs.OperatorStateRetriever,
		ServiceManagerAddress:         contractAddrs.ServiceManager,
	}
	chainWriter3, err := testclients.NewTestAvsRegistryWriterFromConfig(
		anvilHttpEndpoint,
		thirdOperatorPrivateKey,
		config,
	)
	require.NoError(t, err)

	// Register ANVIL_THIRD_ADDRESS as operator. Since there is only one slot available, ANVIL_FIRST_ADDRESS should be
	// kicked
	receipt, err = chainWriter3.RegisterOperatorWithChurn(
		context.Background(),
		thirdOperatorECDSAPrivateKey,
		churnECDSAPrivateKey,
		newKeyPair,
		quorumNumbers,
		quorumNumbers,
		operatorsToKick,
		"",
		true,
	)
	require.NoError(t, err)
	require.NotNil(t, receipt)

	// ANVIL_FIRST_ADDRESS should be deregistered
	deregisteredOperatorWithChurn, err := chainReader.IsOperatorRegistered(&bind.CallOpts{}, firstOperatorAddress)
	require.NoError(t, err)
	require.False(t, deregisteredOperatorWithChurn)

	// ANVIL_THIRD_ADDRESS should be registered
	registeredOperatorWithChurn, err := chainReader.IsOperatorRegistered(&bind.CallOpts{}, thirdOperatorAddress)
	require.NoError(t, err)
	require.True(t, registeredOperatorWithChurn)
}
*/

// Compliance test for BLS signature
func TestBlsSignature(t *testing.T) {
	// read input from JSON if available, otherwise use default values
	// Data taken from
	// https://github.com/Layr-Labs/eigensdk-compliance/blob/429459572302d9fd42c1184b7257703460ba09ca/data_files/bls_signature.json
	var defaultInput = struct {
		Message    string `json:"message"`
		BlsPrivKey string `json:"bls_priv_key"`
	}{
		Message:    "Hello, world!Hello, world!123456",
		BlsPrivKey: "12248929636257230549931416853095037629726205319386239410403476017439825112537",
	}

	testData := testutils.NewTestData(defaultInput)
	// The message to sign
	messageArray := []byte(testData.Input.Message)

	var messageArray32 [32]byte
	copy(messageArray32[:], messageArray)

	// The private key as a string
	privKey, _ := bls.NewPrivateKey(testData.Input.BlsPrivKey)
	keyPair := bls.NewKeyPair(privKey)

	sig := keyPair.SignMessage(messageArray32)

	x := sig.G1Affine.X.String()
	y := sig.G1Affine.Y.String()

	// Values taken from previous run of this test
	assert.Equal(t, x, "15790168376429033610067099039091292283117017641532256477437243974517959682102")
	assert.Equal(t, y, "4960450323239587206117776989095741074887370703941588742100855592356200866613")
}

func TestCreateDelegatedAndSlashableStakeQuorums(t *testing.T) {
	clients, anvilHttpEndpoint := testclients.BuildTestClients(t)
	chainReader := clients.ReadClients.AvsRegistryChainReader

	contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilHttpEndpoint)

	chainWriter := clients.AvsRegistryChainWriter

	operatorSetParams := regcoord.ISlashingRegistryCoordinatorTypesOperatorSetParam{
		MaxOperatorCount: 5,
	}
	minimumStakeNeeded := big.NewInt(0)

	strategyAddr := contractAddrs.Erc20MockStrategy
	strategyParam := regcoord.IStakeRegistryTypesStrategyParams{
		Strategy:   strategyAddr,
		Multiplier: big.NewInt(1e18),
	}

	lookAheadPeriod := uint32(0)

	// First, quorum count is 1 because Registry is initialized with 1 quorum
	count, err := chainReader.GetQuorumCount(&bind.CallOpts{})
	require.NoError(t, err)
	assert.Equal(t, count, uint8(1))

	// Create a new total delegated stake quorum
	receipt, err := chainWriter.CreateTotalDelegatedStakeQuorum(
		context.Background(),
		operatorSetParams,
		minimumStakeNeeded,
		[]regcoord.IStakeRegistryTypesStrategyParams{strategyParam},
		true,
	)
	require.NoError(t, err)
	require.Equal(t, receipt.Status, gethtypes.ReceiptStatusSuccessful)

	// After creating first quorum, count is 2
	count, err = chainReader.GetQuorumCount(&bind.CallOpts{})
	require.NoError(t, err)
	assert.Equal(t, count, uint8(2))

	// Create a new slashable stake quorum
	receipt, err = chainWriter.CreateSlashableStakeQuorum(
		context.Background(),
		operatorSetParams,
		minimumStakeNeeded,
		[]regcoord.IStakeRegistryTypesStrategyParams{strategyParam},
		lookAheadPeriod,
		true,
	)
	require.NoError(t, err)
	require.Equal(t, receipt.Status, gethtypes.ReceiptStatusSuccessful)

	// After creating a new one, quorum count is 3
	count, err = chainReader.GetQuorumCount(&bind.CallOpts{})
	require.NoError(t, err)
	assert.Equal(t, count, uint8(3))
}

func TestEjectOperator(t *testing.T) {
	// Test set up
	clients, anvilHttpEndpoint := testclients.BuildTestClients(t)

	contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilHttpEndpoint)

	chainReader := clients.ReadClients.AvsRegistryChainReader
	chainWriter := clients.AvsRegistryChainWriter

	operatorAddr := gethcommon.HexToAddress(testutils.ANVIL_FIRST_ADDRESS)

	quorumNumbers := types.QuorumNums{0}

	// At the beginning, operator is not registered
	isRegisterd, err := chainReader.IsOperatorRegistered(&bind.CallOpts{}, operatorAddr)
	require.NoError(t, err)
	require.False(t, isRegisterd)

	// After registration, operator is registered
	elWriter := clients.ElChainWriter
	otherKeyPair, err := bls.NewKeyPairFromString("0x01")
	require.NoError(t, err)
	request := elcontracts.RegistrationRequest{
		OperatorAddress: operatorAddr,
		AVSAddress:      contractAddrs.ServiceManager,
		OperatorSetIds:  []uint32{0},
		WaitForReceipt:  true,
		Socket:          "socket",
		BlsKeyPair:      otherKeyPair,
	}

	receipt, err := elWriter.RegisterForOperatorSets(context.Background(), contractAddrs.RegistryCoordinator, request)
	require.NoError(t, err)
	require.NotNil(t, receipt)

	isRegisterd, err = chainReader.IsOperatorRegistered(&bind.CallOpts{}, operatorAddr)
	require.NoError(t, err)
	require.True(t, isRegisterd)

	// After being ejected, operator is not registered anymore
	receipt, err = chainWriter.EjectOperator(context.Background(), operatorAddr, quorumNumbers, true)
	require.NoError(t, err)
	require.Equal(t, receipt.Status, gethtypes.ReceiptStatusSuccessful)

	isRegisterd, err = chainReader.IsOperatorRegistered(&bind.CallOpts{}, operatorAddr)
	require.NoError(t, err)
	require.False(t, isRegisterd)
}

func TestSetOperatorSetParams(t *testing.T) {
	// Test set up
	clients, anvilHttpEndpoint := testclients.BuildTestClients(t)

	contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilHttpEndpoint)

	chainWriter := clients.AvsRegistryChainWriter

	registryCoordinatorAddress := contractAddrs.RegistryCoordinator
	registryCoordinator, err := regcoord.NewContractRegistryCoordinator(
		registryCoordinatorAddress,
		clients.EthHttpClient,
	)
	require.NoError(t, err)

	// This parameters are seted to the quorum created on reg coordinator initialization
	initialParams := regcoord.ISlashingRegistryCoordinatorTypesOperatorSetParam{
		MaxOperatorCount:        10000,
		KickBIPsOfOperatorStake: 15000,
		KickBIPsOfTotalStake:    100,
	}

	// At the beginning, params are the set on initialization
	params, err := registryCoordinator.GetOperatorSetParams(&bind.CallOpts{}, 0)
	require.NoError(t, err)
	require.Equal(t, params, initialParams)

	newOperatorSetParams := regcoord.ISlashingRegistryCoordinatorTypesOperatorSetParam{
		MaxOperatorCount: 5,
	}

	receipt, err := chainWriter.SetOperatorSetParams(
		context.Background(),
		0,
		newOperatorSetParams,
		true,
	)
	require.NoError(t, err)
	require.Equal(t, receipt.Status, gethtypes.ReceiptStatusSuccessful)

	// After setting operator set params, params are the setted ones
	params, err = registryCoordinator.GetOperatorSetParams(&bind.CallOpts{}, 0)
	require.NoError(t, err)
	require.Equal(t, params, newOperatorSetParams)
}

func TestSetChurnApprover(t *testing.T) {
	// Test set up
	clients, anvilHttpEndpoint := testclients.BuildTestClients(t)

	contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilHttpEndpoint)

	chainWriter := clients.AvsRegistryChainWriter

	churnApproverAddress := gethcommon.HexToAddress(testutils.ANVIL_SECOND_ADDRESS)

	ethHttpClient := clients.EthHttpClient

	registryCoordinatorContract, err := regcoord.NewContractRegistryCoordinator(
		contractAddrs.RegistryCoordinator,
		ethHttpClient,
	)
	require.NoError(t, err)

	// At first, churnApprover is anvil first address
	approver, err := registryCoordinatorContract.ChurnApprover(&bind.CallOpts{})
	require.NoError(t, err)
	assert.Equal(t, approver.String(), testutils.ANVIL_FIRST_ADDRESS)

	// Set a new churnApprover
	receipt, err := chainWriter.SetChurnApprover(context.Background(), churnApproverAddress, true)
	require.NoError(t, err)
	require.Equal(t, receipt.Status, gethtypes.ReceiptStatusSuccessful)

	// After change, churnApprover is the setted value
	newApprover, err := registryCoordinatorContract.ChurnApprover(&bind.CallOpts{})
	require.NoError(t, err)
	assert.Equal(t, newApprover.String(), testutils.ANVIL_SECOND_ADDRESS)
}

func TestSetEjector(t *testing.T) {
	// Test set up
	clients, anvilHttpEndpoint := testclients.BuildTestClients(t)

	contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilHttpEndpoint)

	chainWriter := clients.AvsRegistryChainWriter

	ejectorAddress := gethcommon.HexToAddress(testutils.ANVIL_SECOND_ADDRESS)

	ethHttpClient := clients.EthHttpClient

	registryCoordinatorContract, err := regcoord.NewContractRegistryCoordinator(
		contractAddrs.RegistryCoordinator,
		ethHttpClient,
	)
	require.NoError(t, err)

	// At first, ejector is anvil first address
	ejector, err := registryCoordinatorContract.Ejector(&bind.CallOpts{})
	require.NoError(t, err)
	assert.Equal(t, ejector.String(), testutils.ANVIL_FIRST_ADDRESS)

	// Set a new ejector
	receipt, err := chainWriter.SetEjector(context.Background(), ejectorAddress, true)
	require.NoError(t, err)
	require.Equal(t, receipt.Status, gethtypes.ReceiptStatusSuccessful)

	// After change, ejector is the setted value
	newEjector, err := registryCoordinatorContract.Ejector(&bind.CallOpts{})
	require.NoError(t, err)
	assert.Equal(t, newEjector.String(), testutils.ANVIL_SECOND_ADDRESS)
}

func TestSetAccountIdentifier(t *testing.T) {
	// Test set up
	clients, anvilHttpEndpoint := testclients.BuildTestClients(t)

	contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilHttpEndpoint)

	chainWriter := clients.AvsRegistryChainWriter

	accountIdentifierAddress := gethcommon.HexToAddress(testutils.ANVIL_SECOND_ADDRESS)

	ethHttpClient := clients.EthHttpClient

	registryCoordinatorContract, err := regcoord.NewContractRegistryCoordinator(
		contractAddrs.RegistryCoordinator,
		ethHttpClient,
	)
	require.NoError(t, err)

	// At first, accountIdentifier is service manager address
	accountIdentifier, err := registryCoordinatorContract.Avs(&bind.CallOpts{})
	require.NoError(t, err)
	assert.Equal(t, accountIdentifier, contractAddrs.ServiceManager)

	// Set a new accountIdentifier
	receipt, err := chainWriter.SetAccountIdentifier(context.Background(), accountIdentifierAddress, true)
	require.NoError(t, err)
	require.Equal(t, receipt.Status, gethtypes.ReceiptStatusSuccessful)

	// After change, accountIdentifier is the value set
	newAccountIdentifier, err := registryCoordinatorContract.Avs(&bind.CallOpts{})
	require.NoError(t, err)
	assert.Equal(t, newAccountIdentifier.String(), testutils.ANVIL_SECOND_ADDRESS)
}

func TestRemoveStrategies(t *testing.T) {
	clients, _ := testclients.BuildTestClients(t)
	chainWriter := clients.AvsRegistryChainWriter
	chainReader := clients.AvsRegistryChainReader

	quorumNumber := types.QuorumNum(0)
	indices := []*big.Int{big.NewInt(0)}

	_, err := clients.AvsRegistryChainReader.GetStrategyParamsAtIndex(
		&bind.CallOpts{Context: context.Background()},
		quorumNumber.UnderlyingType(),
		indices[0],
	)
	require.NoError(t, err)

	// There is a strategy at index 0. We will remove it
	receipt, err := chainWriter.RemoveStrategies(context.Background(), quorumNumber, indices, true)
	require.NoError(t, err)
	require.Equal(t, receipt.Status, gethtypes.ReceiptStatusSuccessful)

	// After removing, there are no strategies in quorum
	length, err := chainReader.StrategyParamsLength(&bind.CallOpts{}, quorumNumber.UnderlyingType())
	require.NoError(t, err)
	require.Zero(t, length.Cmp(big.NewInt(0)))
}

func TestSetEjectionCooldown(t *testing.T) {
	// Test set up
	clients, anvilHttpEndpoint := testclients.BuildTestClients(t)

	contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilHttpEndpoint)

	chainWriter := clients.AvsRegistryChainWriter

	ejectionCooldown := big.NewInt(2873)

	ethHttpClient := clients.EthHttpClient

	registryCoordinatorContract, err := regcoord.NewContractRegistryCoordinator(
		contractAddrs.RegistryCoordinator,
		ethHttpClient,
	)
	require.NoError(t, err)

	// At first, ejectionCooldown is zero
	cooldown, err := registryCoordinatorContract.EjectionCooldown(&bind.CallOpts{})
	require.NoError(t, err)
	assert.Equal(t, cooldown.Int64(), int64(0))

	// Set a new ejectionCooldown
	receipt, err := chainWriter.SetEjectionCooldown(context.Background(), ejectionCooldown, true)
	require.NoError(t, err)
	require.Equal(t, receipt.Status, gethtypes.ReceiptStatusSuccessful)

	// After change, ejectionCooldown is the value set
	newCooldown, err := registryCoordinatorContract.EjectionCooldown(&bind.CallOpts{})
	require.NoError(t, err)
	assert.Equal(t, newCooldown, ejectionCooldown)
}

func TestCreateAVSRewardsSubmission(t *testing.T) {
	clients, _ := testclients.BuildTestClients(t)
	chainWriter := clients.AvsRegistryChainWriter

	strategies, err := clients.AvsRegistryChainReader.StrategyParamsByIndex(nil, 0, big.NewInt(0))
	require.NoError(t, err)

	calculationInterval, err := clients.EigenlayerContractBindings.RewardsCoordinator.CALCULATIONINTERVALSECONDS(nil)
	require.NoError(t, err)

	strategy := strategies.Strategy

	_, token, err := clients.ElChainReader.GetStrategyAndUnderlyingToken(context.TODO(), strategies.Strategy)
	require.NoError(t, err)

	strategiesAndMultipliers := []servicemanager.IRewardsCoordinatorTypesStrategyAndMultiplier{
		{
			Strategy:   strategy,
			Multiplier: big.NewInt(1),
		},
	}
	header, err := clients.EthHttpClient.HeaderByNumber(context.TODO(), nil)
	require.NoError(t, err)

	// These values are set to align with the contract's requirements for the `OperatorDirectedRewardsSubmission`.
	// https://github.com/Layr-Labs/eigenlayer-contracts/blob/ecaff6304de6cb0f43b42024ad55d0e8a0430790/src/contracts/core/RewardsCoordinator.sol#L414
	// https://github.com/Layr-Labs/eigenlayer-contracts/blob/ecaff6304de6cb0f43b42024ad55d0e8a0430790/src/contracts/core/RewardsCoordinator.sol#L482
	var duration uint32 = calculationInterval
	var startTimestamp uint32 = ((uint32(header.Time) / calculationInterval) + 1) * calculationInterval

	rewardsSubmission := []servicemanager.IRewardsCoordinatorTypesRewardsSubmission{{
		StrategiesAndMultipliers: strategiesAndMultipliers,
		Token:                    token,
		Amount:                   big.NewInt(1000),
		StartTimestamp:           startTimestamp,
		Duration:                 duration,
	}}
	receipt, err := chainWriter.CreateAVSRewardsSubmission(context.TODO(), rewardsSubmission, true)
	require.NoError(t, err)
	require.Equal(t, gethtypes.ReceiptStatusSuccessful, receipt.Status)
}

func TestAddStrategies(t *testing.T) {
	clients, _ := testclients.BuildTestClients(t)
	chainWriter := clients.AvsRegistryChainWriter
	chainReader := clients.AvsRegistryChainReader

	// contractAddrs.Erc20MockStrategy is already set as a strategy at index 0
	strategyParam := stakeregistry.IStakeRegistryTypesStrategyParams{
		Strategy:   gethcommon.HexToAddress("0x1"),
		Multiplier: big.NewInt(100),
	}

	strategiesParams := []stakeregistry.IStakeRegistryTypesStrategyParams{strategyParam}
	quorumNumber := types.QuorumNum(0)

	receipt, err := chainWriter.AddStrategies(
		context.Background(),
		quorumNumber,
		strategiesParams,
		true,
	)
	require.NoError(t, err)
	require.Equal(t, receipt.Status, gethtypes.ReceiptStatusSuccessful)

	// New strategy is set at index 1
	params, err := chainReader.GetStrategyParamsAtIndex(nil, 0, big.NewInt(1))
	require.NoError(t, err)
	require.Equal(t, params.Strategy, strategyParam.Strategy)
	require.Equal(t, params.Multiplier, strategyParam.Multiplier)
}

func TestModifyStrategyParams(t *testing.T) {
	clients, _ := testclients.BuildTestClients(t)
	chainWriter := clients.AvsRegistryChainWriter
	chainReader := clients.AvsRegistryChainReader

	indices := []*big.Int{big.NewInt(0)}
	multiplier := []*big.Int{big.NewInt(5e18)}

	receipt, err := chainWriter.ModifyStrategyParams(context.Background(), 0, indices, multiplier, true)
	require.NoError(t, err)
	require.Equal(t, gethtypes.ReceiptStatusSuccessful, receipt.Status)

	strategies, err := chainReader.StrategyParamsByIndex(nil, 0, indices[0])
	require.NoError(t, err)
	require.Equal(t, strategies.Multiplier, multiplier[0])
}

func TestCreateOperatorDirectedAVSRewardsSubmission(t *testing.T) {
	clients, _ := testclients.BuildTestClients(t)
	chainWriter := clients.AvsRegistryChainWriter

	strategies, err := clients.AvsRegistryChainReader.StrategyParamsByIndex(nil, 0, big.NewInt(0))
	require.NoError(t, err)

	calculationInterval, err := clients.EigenlayerContractBindings.RewardsCoordinator.CALCULATIONINTERVALSECONDS(nil)
	require.NoError(t, err)

	strategy := strategies.Strategy

	_, token, err := clients.ElChainReader.GetStrategyAndUnderlyingToken(context.TODO(), strategies.Strategy)
	require.NoError(t, err)

	strategiesAndMultipliers := []servicemanager.IRewardsCoordinatorTypesStrategyAndMultiplier{
		{
			Strategy:   strategy,
			Multiplier: big.NewInt(1),
		},
	}
	header, err := clients.EthHttpClient.HeaderByNumber(context.TODO(), nil)
	require.NoError(t, err)

	// These values are set to align with the contract's requirements for the `OperatorDirectedRewardsSubmission`.
	// https://github.com/Layr-labs/eigenlayer-contracts/blob/5341ef83500476c62a4406ff00cdde7f5c2cc11f/src/contracts/core/RewardsCoordinator.sol#L485
	// https://github.com/Layr-labs/eigenlayer-contracts/blob/5341ef83500476c62a4406ff00cdde7f5c2cc11f/src/contracts/core/RewardsCoordinator.sol#L438
	var duration uint32 = calculationInterval
	var startTimestamp uint32 = ((uint32(header.Time) / calculationInterval) - 2) * calculationInterval

	operatorRewards := []servicemanager.IRewardsCoordinatorTypesOperatorReward{{
		Operator: gethcommon.HexToAddress(testutils.ANVIL_FIRST_ADDRESS),
		Amount:   big.NewInt(1000),
	}}

	rewardsSubmission := []servicemanager.IRewardsCoordinatorTypesOperatorDirectedRewardsSubmission{{
		StrategiesAndMultipliers: strategiesAndMultipliers,
		Token:                    token,
		OperatorRewards:          operatorRewards,
		StartTimestamp:           startTimestamp,
		Duration:                 duration,
		Description:              "some description here",
	}}
	receipt, err := chainWriter.CreateOperatorDirectedAVSRewardsSubmission(context.TODO(), rewardsSubmission, true)
	require.NoError(t, err)
	require.Equal(t, gethtypes.ReceiptStatusSuccessful, receipt.Status)
}

func TestUpdateAVSMetadataURI(t *testing.T) {
	clients, _ := testclients.BuildTestClients(t)
	chainWriter := clients.AvsRegistryChainWriter

	svcManagerAddr := clients.AvsRegistryContractBindings.ServiceManagerAddr
	avsDirectoryAddr := clients.EigenlayerContractBindings.AvsDirectoryAddr

	avsDirectory, err := avsdirectory.NewContractAVSDirectory(
		avsDirectoryAddr,
		clients.EthHttpClient,
	)
	require.NoError(t, err)

	// Update the metadata URI
	newMetadata := "https://new-metadata-uri.com"
	receipt, err := chainWriter.UpdateAVSMetadataURI(context.TODO(), newMetadata, true)
	require.NoError(t, err)
	require.Equal(t, gethtypes.ReceiptStatusSuccessful, receipt.Status)

	// Assert the event was emitted
	iter, err := avsDirectory.FilterAVSMetadataURIUpdated(nil, []gethcommon.Address{svcManagerAddr})
	require.NoError(t, err)
	require.True(t, iter.Next())
	require.Equal(t, newMetadata, iter.Event.MetadataURI)
}
