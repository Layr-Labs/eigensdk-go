package wallet_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/fireblocks"
	cmocks "github.com/Layr-Labs/eigensdk-go/chainio/clients/mocks"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/wallet"
	"github.com/Layr-Labs/eigensdk-go/chainio/mocks"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

const (
	vaultAccountName = "batcher"
	contractAddress  = "0x5f9ef6e1bb2acb8f592a483052b732ceb78e58ca"
)

func TestSendTransaction(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	fireblocksClient := cmocks.NewMockFireblocksClient(ctrl)
	ethClient := mocks.NewMockEthClient(ctrl)
	logger, err := logging.NewZapLogger(logging.Development)
	assert.NoError(t, err)
	ethClient.EXPECT().ChainID(gomock.Any()).Return(big.NewInt(5), nil)
	sender, err := wallet.NewFireblocksWallet(fireblocksClient, ethClient, vaultAccountName, logger)
	assert.NoError(t, err)

	fireblocksClient.EXPECT().ListContracts(gomock.Any()).Return([]fireblocks.WhitelistedContract{
		{
			ID:   "contractID",
			Name: "TestContract",
			Assets: []struct {
				ID      fireblocks.AssetID `json:"id"`
				Status  string             `json:"status"`
				Address common.Address     `json:"address"`
				Tag     string             `json:"tag"`
			}{{
				ID:      "ETH_TEST3",
				Status:  "APPROVED",
				Address: common.HexToAddress(contractAddress),
				Tag:     "",
			},
			},
		},
	}, nil)
	fireblocksClient.EXPECT().ContractCall(gomock.Any(), gomock.Any()).Return(&fireblocks.ContractCallResponse{
		ID:     "1234",
		Status: fireblocks.Confirming,
	}, nil)
	fireblocksClient.EXPECT().ListVaultAccounts(gomock.Any()).Return([]fireblocks.VaultAccount{
		{
			ID:   "vaultAccountID",
			Name: vaultAccountName,
			Assets: []fireblocks.Asset{
				{
					ID:        "ETH_TEST3",
					Total:     "1",
					Balance:   "1",
					Available: "1",
				},
			},
		},
	}, nil)

	txID, err := sender.SendTransaction(context.Background(), types.NewTransaction(
		0,                                    // nonce
		common.HexToAddress(contractAddress), // to
		big.NewInt(0),                        // value
		100000,                               // gas
		big.NewInt(100),                      // gasPrice
		common.Hex2Bytes("0x6057361d00000000000000000000000000000000000000000000000000000000000f4240"), // data
	))
	assert.NoError(t, err)
	assert.Equal(t, "1234", txID)
}

func TestSendTransactionNoValidContract(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	fireblocksClient := cmocks.NewMockFireblocksClient(ctrl)
	ethClient := mocks.NewMockEthClient(ctrl)
	logger, err := logging.NewZapLogger(logging.Development)
	assert.NoError(t, err)
	ethClient.EXPECT().ChainID(gomock.Any()).Return(big.NewInt(5), nil)
	sender, err := wallet.NewFireblocksWallet(fireblocksClient, ethClient, vaultAccountName, logger)
	assert.NoError(t, err)

	fireblocksClient.EXPECT().ListContracts(gomock.Any()).Return([]fireblocks.WhitelistedContract{
		{
			ID:   "contractID",
			Name: "TestContract",
			Assets: []struct {
				ID      fireblocks.AssetID `json:"id"`
				Status  string             `json:"status"`
				Address common.Address     `json:"address"`
				Tag     string             `json:"tag"`
			}{{
				ID:      "ETH_TEST123123", // wrong asset ID
				Status:  "APPROVED",
				Address: common.HexToAddress(contractAddress),
				Tag:     "",
			},
			},
		},
	}, nil)
	fireblocksClient.EXPECT().ListVaultAccounts(gomock.Any()).Return([]fireblocks.VaultAccount{
		{
			ID:   "vaultAccountID",
			Name: vaultAccountName,
			Assets: []fireblocks.Asset{
				{
					ID:        "ETH_TEST3",
					Total:     "1",
					Balance:   "1",
					Available: "1",
				},
			},
		},
	}, nil)

	txID, err := sender.SendTransaction(context.Background(), types.NewTransaction(
		0,                                    // nonce
		common.HexToAddress(contractAddress), // to
		big.NewInt(0),                        // value
		100000,                               // gas
		big.NewInt(100),                      // gasPrice
		common.Hex2Bytes("0x6057361d00000000000000000000000000000000000000000000000000000000000f4240"), // data
	))
	assert.Error(t, err)
	assert.Equal(t, "", txID)
}

func TestSendTransactionInvalidVault(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	fireblocksClient := cmocks.NewMockFireblocksClient(ctrl)
	ethClient := mocks.NewMockEthClient(ctrl)
	logger, err := logging.NewZapLogger(logging.Development)
	assert.NoError(t, err)
	ethClient.EXPECT().ChainID(gomock.Any()).Return(big.NewInt(5), nil)
	sender, err := wallet.NewFireblocksWallet(fireblocksClient, ethClient, vaultAccountName, logger)
	assert.NoError(t, err)

	fireblocksClient.EXPECT().ListVaultAccounts(gomock.Any()).Return([]fireblocks.VaultAccount{
		{
			ID:   "vaultAccountID",
			Name: vaultAccountName,
			Assets: []fireblocks.Asset{
				{
					ID:        "ETH_TEST123123", // wrong asset ID
					Total:     "1",
					Balance:   "1",
					Available: "1",
				},
			},
		},
	}, nil)

	txID, err := sender.SendTransaction(context.Background(), types.NewTransaction(
		0,                                    // nonce
		common.HexToAddress(contractAddress), // to
		big.NewInt(0),                        // value
		100000,                               // gas
		big.NewInt(100),                      // gasPrice
		common.Hex2Bytes("0x6057361d00000000000000000000000000000000000000000000000000000000000f4240"), // data
	))
	assert.Error(t, err)
	assert.Equal(t, "", txID)
}

func TestSendTransactionReplaceTx(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	fireblocksClient := cmocks.NewMockFireblocksClient(ctrl)
	ethClient := mocks.NewMockEthClient(ctrl)
	logger, err := logging.NewZapLogger(logging.Development)
	assert.NoError(t, err)
	ethClient.EXPECT().ChainID(gomock.Any()).Return(big.NewInt(5), nil)
	sender, err := wallet.NewFireblocksWallet(fireblocksClient, ethClient, vaultAccountName, logger)
	assert.NoError(t, err)

	fireblocksClient.EXPECT().ListContracts(gomock.Any()).Return([]fireblocks.WhitelistedContract{
		{
			ID:   "contractID",
			Name: "TestContract",
			Assets: []struct {
				ID      fireblocks.AssetID `json:"id"`
				Status  string             `json:"status"`
				Address common.Address     `json:"address"`
				Tag     string             `json:"tag"`
			}{{
				ID:      "ETH_TEST3",
				Status:  "APPROVED",
				Address: common.HexToAddress(contractAddress),
				Tag:     "",
			},
			},
		},
	}, nil)
	fireblocksClient.EXPECT().ContractCall(gomock.Any(), gomock.Any()).Return(&fireblocks.ContractCallResponse{
		ID:     "1234",
		Status: fireblocks.Confirming,
	}, nil)
	fireblocksClient.EXPECT().ListVaultAccounts(gomock.Any()).Return([]fireblocks.VaultAccount{
		{
			ID:   "vaultAccountID",
			Name: vaultAccountName,
			Assets: []fireblocks.Asset{
				{
					ID:        "ETH_TEST3",
					Total:     "1",
					Balance:   "1",
					Available: "1",
				},
			},
		},
	}, nil)

	txID, err := sender.SendTransaction(context.Background(), types.NewTransaction(
		0,                                    // nonce
		common.HexToAddress(contractAddress), // to
		big.NewInt(0),                        // value
		100000,                               // gas
		big.NewInt(100),                      // gasPrice
		common.Hex2Bytes("0x6057361d00000000000000000000000000000000000000000000000000000000000f4240"), // data
	))
	assert.NoError(t, err)
	assert.Equal(t, "1234", txID)

	addr := common.HexToAddress(contractAddress)
	gasLimit := uint64(1000000)
	baseTx := &types.DynamicFeeTx{
		To:        &addr,
		Nonce:     0,
		GasFeeCap: big.NewInt(10_000_000_000),
		GasTipCap: big.NewInt(1_000_000_000),
		Gas:       gasLimit,
		Value:     big.NewInt(0),
		Data:      common.Hex2Bytes("0x6057361d00000000000000000000000000000000000000000000000000000000000f4240"),
	}
	replacementTx := types.NewTx(baseTx)
	expectedTxHash := "0xdeadbeef"
	fireblocksClient.EXPECT().GetTransaction(gomock.Any(), "1234").Return(&fireblocks.Transaction{
		ID:     expectedTxHash,
		Status: fireblocks.Completed,
		TxHash: expectedTxHash,
	}, nil)
	fireblocksClient.EXPECT().ContractCall(gomock.Any(), fireblocks.NewContractCallRequest(
		"",
		"ETH_TEST3",
		"vaultAccountID",
		"contractID",
		"0",
		"0x",
		expectedTxHash,
		"",        // gasPrice
		"1000000", // gasLimit
		"10",      // maxFee
		"1",       // priorityFee
		"",        // feeLevel
	)).Return(&fireblocks.ContractCallResponse{
		ID:     "5678",
		Status: fireblocks.Confirming,
	}, nil)
	// send another tx with the same nonce
	txID, err = sender.SendTransaction(context.Background(), replacementTx)
	assert.NoError(t, err)
	assert.Equal(t, "5678", txID)
}

func TestWaitForTransactionReceipt(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	fireblocksClient := cmocks.NewMockFireblocksClient(ctrl)
	ethClient := mocks.NewMockEthClient(ctrl)
	logger, err := logging.NewZapLogger(logging.Development)
	assert.NoError(t, err)
	ethClient.EXPECT().ChainID(gomock.Any()).Return(big.NewInt(5), nil)
	sender, err := wallet.NewFireblocksWallet(fireblocksClient, ethClient, vaultAccountName, logger)
	assert.NoError(t, err)

	expectedTxHash := "0x0000000000000000000000000000000000000000000000000000000000001234"
	fireblocksClient.EXPECT().GetTransaction(gomock.Any(), expectedTxHash).Return(&fireblocks.Transaction{
		ID:     expectedTxHash,
		Status: fireblocks.Completed,
		TxHash: expectedTxHash,
	}, nil)
	ethClient.EXPECT().TransactionReceipt(gomock.Any(), common.HexToHash(expectedTxHash)).Return(&types.Receipt{
		TxHash:      common.HexToHash(expectedTxHash),
		BlockNumber: big.NewInt(1234),
	}, nil)

	receipt, err := sender.GetTransactionReceipt(context.Background(), expectedTxHash)
	assert.NoError(t, err)
	assert.Equal(t, expectedTxHash, receipt.TxHash.String())
	assert.Equal(t, big.NewInt(1234), receipt.BlockNumber)
}

func TestWaitForTransactionReceiptFailFromFireblocks(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	fireblocksClient := cmocks.NewMockFireblocksClient(ctrl)
	ethClient := mocks.NewMockEthClient(ctrl)
	logger, err := logging.NewZapLogger(logging.Development)
	assert.NoError(t, err)
	ethClient.EXPECT().ChainID(gomock.Any()).Return(big.NewInt(5), nil)
	sender, err := wallet.NewFireblocksWallet(fireblocksClient, ethClient, vaultAccountName, logger)
	assert.NoError(t, err)

	expectedTxHash := "0x0000000000000000000000000000000000000000000000000000000000001234"
	fireblocksClient.EXPECT().GetTransaction(gomock.Any(), expectedTxHash).Return(&fireblocks.Transaction{
		ID:     expectedTxHash,
		Status: fireblocks.Confirming, // not completed
		TxHash: expectedTxHash,
	}, nil)

	receipt, err := sender.GetTransactionReceipt(context.Background(), expectedTxHash)
	assert.ErrorAs(t, err, &wallet.ErrReceiptNotYetAvailable)
	assert.Nil(t, receipt)
}

func TestWaitForTransactionReceiptStuckAtFireblocks(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	fireblocksClient := cmocks.NewMockFireblocksClient(ctrl)
	ethClient := mocks.NewMockEthClient(ctrl)
	logger, err := logging.NewZapLogger(logging.Development)
	assert.NoError(t, err)
	ethClient.EXPECT().ChainID(gomock.Any()).Return(big.NewInt(5), nil)
	sender, err := wallet.NewFireblocksWallet(fireblocksClient, ethClient, vaultAccountName, logger)
	assert.NoError(t, err)

	expectedTxHash := "0x0000000000000000000000000000000000000000000000000000000000001234"
	fireblocksClient.EXPECT().GetTransaction(gomock.Any(), expectedTxHash).Return(&fireblocks.Transaction{
		ID:     expectedTxHash,
		Status: fireblocks.PendingSignature, // not completed
		TxHash: expectedTxHash,
	}, nil)

	receipt, err := sender.GetTransactionReceipt(context.Background(), expectedTxHash)
	assert.ErrorAs(t, err, &wallet.ErrNotYetBroadcasted)
	assert.Nil(t, receipt)
}

func TestWaitForTransactionReceiptFailFromChain(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	fireblocksClient := cmocks.NewMockFireblocksClient(ctrl)
	ethClient := mocks.NewMockEthClient(ctrl)
	logger, err := logging.NewZapLogger(logging.Development)
	assert.NoError(t, err)
	ethClient.EXPECT().ChainID(gomock.Any()).Return(big.NewInt(5), nil)
	sender, err := wallet.NewFireblocksWallet(fireblocksClient, ethClient, vaultAccountName, logger)
	assert.NoError(t, err)

	expectedTxHash := "0x0000000000000000000000000000000000000000000000000000000000001234"
	fireblocksClient.EXPECT().GetTransaction(gomock.Any(), expectedTxHash).Return(&fireblocks.Transaction{
		ID:     expectedTxHash,
		Status: fireblocks.Completed,
		TxHash: expectedTxHash,
	}, nil)
	ethClient.EXPECT().TransactionReceipt(gomock.Any(), common.HexToHash(expectedTxHash)).Return(nil, ethereum.NotFound)

	receipt, err := sender.GetTransactionReceipt(context.Background(), expectedTxHash)
	assert.Error(t, err)
	assert.Nil(t, receipt)
}

func TestSenderAddress(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	fireblocksClient := cmocks.NewMockFireblocksClient(ctrl)
	ethClient := mocks.NewMockEthClient(ctrl)
	logger, err := logging.NewZapLogger(logging.Development)
	assert.NoError(t, err)
	ethClient.EXPECT().ChainID(gomock.Any()).Return(big.NewInt(5), nil)
	w, err := wallet.NewFireblocksWallet(fireblocksClient, ethClient, vaultAccountName, logger)
	assert.NoError(t, err)
	assetID := fireblocks.AssetIDByChain[5]
	fireblocksClient.EXPECT().ListVaultAccounts(gomock.Any()).Return([]fireblocks.VaultAccount{
		{
			ID:   "vaultAccountID",
			Name: vaultAccountName,
			Assets: []fireblocks.Asset{
				{
					ID:        assetID,
					Total:     "1",
					Balance:   "1",
					Available: "1",
				},
			},
		},
	}, nil)
	expectedSenderAddr := "0x0000000000000000000000000000000000000000"
	fireblocksClient.EXPECT().GetAssetAddresses(gomock.Any(), "vaultAccountID", assetID).Return([]fireblocks.AssetAddress{
		{
			AssetID: assetID,
			Address: expectedSenderAddr,
		},
	}, nil)

	addr, err := w.SenderAddress(context.Background())
	assert.Nil(t, err)
	assert.Equal(t, expectedSenderAddr, addr.String())
}
