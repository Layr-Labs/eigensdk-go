package avsregistry_test

import (
	"context"
	"testing"
	"time"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/elcontracts"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/testutils"
	"github.com/Layr-Labs/eigensdk-go/testutils/testclients"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSubscriberAvsRegistry(t *testing.T) {
	clients, anvilHttpEndpoint := testclients.BuildTestClients(t)
	chainSubscriber := clients.AvsRegistryChainSubscriber
	chainWriter := clients.AvsRegistryChainWriter

	contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilHttpEndpoint)

	t.Run("subscribe to new pubkey registrations", func(t *testing.T) {
		pubKeyRegistrationsC, event, err := chainSubscriber.SubscribeToNewPubkeyRegistrations()
		require.NoError(t, err)
		defer event.Unsubscribe()

		// Emit a NewPubkeyRegistration event creating a new operator
		keypair, err := bls.NewKeyPairFromString("0x01")
		require.NoError(t, err)

		ecdsaPrivateKey, err := crypto.HexToECDSA(testutils.ANVIL_FIRST_PRIVATE_KEY)
		require.NoError(t, err)

		//quorumNumbers := types.QuorumNums{0}

		require.NoError(t, err)
		request := elcontracts.RegistrationRequest{
			OperatorAddress: common.HexToAddress(testutils.ANVIL_FIRST_ADDRESS),
			AVSAddress:      contractAddrs.ServiceManager,
			OperatorSetIds:  []uint32{0},
			WaitForReceipt:  true,
			Socket:          "socket",
			BlsKeyPair:      keypair,
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

		receipt, err = elWriter.RegisterForOperatorSets(
			context.Background(),
			contractAddrs.RegistryCoordinator,
			request,
		)
		require.NoError(t, err)
		require.NotNil(t, receipt)

		select {
		case newPubkeyRegistration := <-pubKeyRegistrationsC:
			expectedOperator := crypto.PubkeyToAddress(ecdsaPrivateKey.PublicKey)
			assert.Equal(t, expectedOperator, newPubkeyRegistration.Operator)
		case <-time.After(10 * time.Second):
			// Throw an error if the event is not received within 10 seconds, making the test fail
			t.Fatal("Timed out waiting for NewPubkeyRegistration event")
		}
	})

	t.Run("subscribe to operator socket updates", func(t *testing.T) {
		socketC, event, err := chainSubscriber.SubscribeToOperatorSocketUpdates()
		require.NoError(t, err)
		defer event.Unsubscribe()

		// Emit a SocketUpdate event
		socketUpdate := "socket-update"
		receipt, err := chainWriter.UpdateSocket(context.Background(), types.Socket(socketUpdate), true)
		require.NoError(t, err)
		require.NotNil(t, receipt)

		select {
		case operatorSocketUpdate := <-socketC:
			assert.Equal(t, socketUpdate, operatorSocketUpdate.Socket)
		case <-time.After(10 * time.Second):
			// Throw an error if the event is not received within 10 seconds, making the test fail
			t.Fatal("Timed out waiting for OperatorSocketUpdate event")
		}
	})
}
