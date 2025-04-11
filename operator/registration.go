package operator

import (
	"context"

	allocationmanager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/AllocationManager"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/elcontracts"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/metrics"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/Layr-Labs/eigensdk-go/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// The idea is to have a util function that just registers an operator for an operator set if the avs needs it
func RegisterOperatorWithEigenlayer(
	operatorAddr common.Address,
	elcontractsConfig elcontracts.Config,
	ethClient *ethclient.Client,
	logger logging.Logger,
	txMgr txmgr.TxManager,
) error {

	// Register operator with EigenLayer
	op := types.Operator{
		Address:                   operatorAddr.String(),
		DelegationApproverAddress: operatorAddr.String(),
	}

	elWriter, err := elcontracts.NewWriterFromConfig(elcontractsConfig, ethClient, logger, &metrics.EigenMetrics{}, txMgr)
	if err != nil {
		logger.Error("Error creating eigenlayer chain writer", "err", err)
		return err
	}

	_, err = elWriter.RegisterAsOperator(context.Background(), op, true)
	if err != nil {
		logger.Error("Error registering operator with eigenlayer", "err", err)
		return err
	}

	return nil
}

func RegisterForOperatorSets(
	operatorAddr common.Address,
	logger logging.Logger,
	elcontractsConfig elcontracts.Config,
	ethClient *ethclient.Client,
	txMgr txmgr.TxManager,
	registryCoordinatorAddr common.Address,
	serviceManagerAddr common.Address,
	operatorSetsIds []uint32,
	blsKeyPair bls.KeyPair,
	socket string,
) error {
	elWriter, err := elcontracts.NewWriterFromConfig(elcontractsConfig, ethClient, logger, &metrics.EigenMetrics{}, txMgr)
	if err != nil {
		logger.Error("Error creating eigenlayer chain writer", "err", err)
		return err
	}
	// Register operator for operator sets
	registrationRequest := elcontracts.RegistrationRequest{
		OperatorAddress: operatorAddr,
		AVSAddress:      serviceManagerAddr,
		OperatorSetIds:  operatorSetsIds,
		WaitForReceipt:  true,
		BlsKeyPair:      &blsKeyPair,
		Socket:          socket,
	}

	_, err = elWriter.RegisterForOperatorSets(
		context.Background(),
		registryCoordinatorAddr,
		registrationRequest,
	)

	if err != nil {
		logger.Errorf("Unable to register operator with the operator set")
		return err
	}
	logger.Info("Registered operator with operator set")

	return nil
}

func SetAllocationDelay(
	logger logging.Logger,
	operatorAddr common.Address,
	ethClient *ethclient.Client,
	allocationManagerAddr common.Address,
	txMgr txmgr.TxManager,
	delay uint32,
) error {
	txOpts, _ := txMgr.GetNoSendTxOpts()

	waitForReceipt := true
	allocationManagerContract, _ := allocationmanager.NewContractAllocationManager(allocationManagerAddr, ethClient)

	tx, err := allocationManagerContract.SetAllocationDelay(txOpts, operatorAddr, delay)
	if err != nil {
		return err
	}
	receipt, err := txMgr.Send(context.Background(), tx, waitForReceipt)
	if err != nil {
		return utils.WrapError("failed to send setAllocationDelay tx with err", err)
	}
	logger.Info(
		"tx successfully included for SetAllocationDelay  ",
		"txHash",
		receipt.TxHash.String(),
	)

	return nil
}
