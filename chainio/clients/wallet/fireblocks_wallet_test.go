package wallet_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/Layr-Labs/eigensdk-go/internal/fakes"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/fireblocks"
	cmocks "github.com/Layr-Labs/eigensdk-go/chainio/clients/mocks"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/wallet"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

const (
	vaultAccountName = "batcher"
	contractAddress  = "0x5f9ef6e1bb2acb8f592a483052b732ceb78e58ca"
	externalAccount  = "0x1111111111111111111111111111111111111111"
)

func TestSendTransaction(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	fireblocksClient := cmocks.NewMockFireblocksClient(ctrl)
	ethClient := fakes.NewEthClient()
	logger, err := logging.NewZapLogger(logging.Development)
	assert.NoError(t, err)
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
	fireblocksClient.EXPECT().ContractCall(gomock.Any(), gomock.Any()).Return(&fireblocks.TransactionResponse{
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
		common.Hex2Bytes("6057361d00000000000000000000000000000000000000000000000000000000000f4240"), // data
	))
	assert.NoError(t, err)
	assert.Equal(t, "1234", txID)
}

func TestTransfer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	fireblocksClient := cmocks.NewMockFireblocksClient(ctrl)
	ethClient := fakes.NewEthClient()
	logger, err := logging.NewZapLogger(logging.Development)
	assert.NoError(t, err)
	sender, err := wallet.NewFireblocksWallet(fireblocksClient, ethClient, vaultAccountName, logger)
	assert.NoError(t, err)

	fireblocksClient.EXPECT().ListExternalWallets(gomock.Any()).Return([]fireblocks.WhitelistedAccount{
		{
			ID:   "accountID",
			Name: "Test Account",
			Assets: []struct {
				ID           fireblocks.AssetID `json:"id"`
				Balance      string             `json:"balance"`
				LockedAmount string             `json:"lockedAmount"`
				Status       string             `json:"status"`
				Address      common.Address     `json:"address"`
				Tag          string             `json:"tag"`
			}{{
				ID:           "ETH_TEST3",
				Balance:      "",
				LockedAmount: "",
				Status:       "APPROVED",
				Address:      common.HexToAddress(externalAccount),
				Tag:          "",
			},
			},
		},
	}, nil)
	fireblocksClient.EXPECT().Transfer(gomock.Any(), &fireblocks.TransactionRequest{
		Operation:    fireblocks.Transfer,
		ExternalTxID: "",
		AssetID:      "ETH_TEST3",
		Source: struct {
			Type string `json:"type"`
			ID   string `json:"id"`
		}{
			Type: "VAULT_ACCOUNT",
			ID:   "vaultAccountID",
		},
		Destination: struct {
			Type string `json:"type"`
			ID   string `json:"id"`
		}{
			Type: "EXTERNAL_WALLET",
			ID:   "accountID",
		},
		Amount:          "1",
		ReplaceTxByHash: "",
		GasPrice:        "",
		GasLimit:        "100000",
		MaxFee:          "1e-07",
		PriorityFee:     "1e-07",
	}).Return(&fireblocks.TransactionResponse{
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
		common.HexToAddress(externalAccount), // to
		big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil), // value 1 ETH
		100000,          // gas
		big.NewInt(100), // gasPrice
		[]byte{},        // data
	))
	assert.NoError(t, err)
	assert.Equal(t, "1234", txID)
}

func TestSendTransactionNoValidContract(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	fireblocksClient := cmocks.NewMockFireblocksClient(ctrl)
	ethClient := fakes.NewEthClient()
	logger, err := logging.NewZapLogger(logging.Development)
	assert.NoError(t, err)
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
		common.Hex2Bytes("6057361d00000000000000000000000000000000000000000000000000000000000f4240"), // data
	))
	assert.Error(t, err)
	assert.Equal(t, "", txID)
}

func TestSendTransactionNoValidExternalAccount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	fireblocksClient := cmocks.NewMockFireblocksClient(ctrl)
	ethClient := fakes.NewEthClient()
	logger, err := logging.NewZapLogger(logging.Development)
	assert.NoError(t, err)
	sender, err := wallet.NewFireblocksWallet(fireblocksClient, ethClient, vaultAccountName, logger)
	assert.NoError(t, err)

	fireblocksClient.EXPECT().ListExternalWallets(gomock.Any()).Return([]fireblocks.WhitelistedAccount{
		{
			ID:   "accountID",
			Name: "TestAccount",
			Assets: []struct {
				ID           fireblocks.AssetID `json:"id"`
				Balance      string             `json:"balance"`
				LockedAmount string             `json:"lockedAmount"`
				Status       string             `json:"status"`
				Address      common.Address     `json:"address"`
				Tag          string             `json:"tag"`
			}{{
				ID:           "ETH_TEST123123", // wrong asset ID
				Balance:      "",
				LockedAmount: "",
				Status:       "APPROVED",
				Address:      common.HexToAddress(externalAccount),
				Tag:          "",
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
		common.HexToAddress(externalAccount), // to
		big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil), // value 1 ETH
		100000,          // gas
		big.NewInt(100), // gasPrice
		[]byte{},        // data
	))
	assert.Error(t, err)
	assert.Equal(t, "", txID)
}

func TestSendTransactionInvalidVault(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	fireblocksClient := cmocks.NewMockFireblocksClient(ctrl)
	ethClient := fakes.NewEthClient()
	logger, err := logging.NewZapLogger(logging.Development)
	assert.NoError(t, err)
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
		common.Hex2Bytes("6057361d00000000000000000000000000000000000000000000000000000000000f4240"), // data
	))
	assert.Error(t, err)
	assert.Equal(t, "", txID)
}

func TestSendTransactionReplaceTx(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	fireblocksClient := cmocks.NewMockFireblocksClient(ctrl)
	ethClient := fakes.NewEthClient()
	logger, err := logging.NewZapLogger(logging.Development)
	assert.NoError(t, err)
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
	fireblocksClient.EXPECT().ContractCall(gomock.Any(), gomock.Any()).Return(&fireblocks.TransactionResponse{
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
		common.Hex2Bytes("6057361d00000000000000000000000000000000000000000000000000000000000f4240"), // data
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
		Data:      common.Hex2Bytes("6057361d00000000000000000000000000000000000000000000000000000000000f4240"),
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
		"0x6057361d00000000000000000000000000000000000000000000000000000000000f4240",
		expectedTxHash,
		"",        // gasPrice
		"1000000", // gasLimit
		"10",      // maxFee
		"1",       // priorityFee
		"",        // feeLevel
	)).Return(&fireblocks.TransactionResponse{
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
	ethClient := fakes.NewEthClient()
	logger, err := logging.NewZapLogger(logging.Development)
	assert.NoError(t, err)
	sender, err := wallet.NewFireblocksWallet(fireblocksClient, ethClient, vaultAccountName, logger)
	assert.NoError(t, err)

	fireblocksClient.EXPECT().GetTransaction(gomock.Any(), fakes.TransactionHash).Return(&fireblocks.Transaction{
		ID:     fakes.TransactionHash,
		Status: fireblocks.Completed,
		TxHash: fakes.TransactionHash,
	}, nil)

	receipt, err := sender.GetTransactionReceipt(context.Background(), fakes.TransactionHash)
	assert.NoError(t, err)
	assert.Equal(t, fakes.TransactionHash, receipt.TxHash.String())
	assert.Equal(t, big.NewInt(1234), receipt.BlockNumber)
}

func TestWaitForTransactionReceiptFailFromFireblocks(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	fireblocksClient := cmocks.NewMockFireblocksClient(ctrl)
	ethClient := fakes.NewEthClient()
	logger, err := logging.NewZapLogger(logging.Development)
	assert.NoError(t, err)
	sender, err := wallet.NewFireblocksWallet(fireblocksClient, ethClient, vaultAccountName, logger)
	assert.NoError(t, err)

	fireblocksClient.EXPECT().GetTransaction(gomock.Any(), fakes.TransactionHash).Return(&fireblocks.Transaction{
		ID:     fakes.TransactionHash,
		Status: fireblocks.Confirming, // not completed
		TxHash: fakes.TransactionHash,
	}, nil)

	receipt, err := sender.GetTransactionReceipt(context.Background(), fakes.TransactionHash)
	assert.ErrorAs(t, err, &wallet.ErrReceiptNotYetAvailable)
	assert.Nil(t, receipt)
}

func TestWaitForTransactionReceiptStuckAtFireblocks(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	fireblocksClient := cmocks.NewMockFireblocksClient(ctrl)
	ethClient := fakes.NewEthClient()
	logger, err := logging.NewZapLogger(logging.Development)
	assert.NoError(t, err)
	sender, err := wallet.NewFireblocksWallet(fireblocksClient, ethClient, vaultAccountName, logger)
	assert.NoError(t, err)

	fireblocksClient.EXPECT().GetTransaction(gomock.Any(), fakes.TransactionHash).Return(&fireblocks.Transaction{
		ID:     fakes.TransactionHash,
		Status: fireblocks.PendingSignature, // not completed
		TxHash: fakes.TransactionHash,
	}, nil)

	receipt, err := sender.GetTransactionReceipt(context.Background(), fakes.TransactionHash)
	assert.ErrorAs(t, err, &wallet.ErrNotYetBroadcasted)
	assert.Nil(t, receipt)
}

func TestWaitForTransactionReceiptFailFromChain(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	fireblocksClient := cmocks.NewMockFireblocksClient(ctrl)
	ethClient := fakes.NewEthClient()
	logger, err := logging.NewZapLogger(logging.Development)
	assert.NoError(t, err)
	sender, err := wallet.NewFireblocksWallet(fireblocksClient, ethClient, vaultAccountName, logger)
	assert.NoError(t, err)

	fireblocksClient.EXPECT().
		GetTransaction(gomock.Any(), fakes.TransactionNashNotInFake).
		Return(&fireblocks.Transaction{
			ID:     fakes.TransactionNashNotInFake,
			Status: fireblocks.Completed,
			TxHash: fakes.TransactionNashNotInFake,
		}, nil)

	receipt, err := sender.GetTransactionReceipt(context.Background(), fakes.TransactionNashNotInFake)
	assert.Error(t, err)
	assert.Nil(t, receipt)
}

func TestSenderAddress(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	fireblocksClient := cmocks.NewMockFireblocksClient(ctrl)
	ethClient := fakes.NewEthClient()
	logger, err := logging.NewZapLogger(logging.Development)
	assert.NoError(t, err)
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
	fireblocksClient.EXPECT().
		GetAssetAddresses(gomock.Any(), "vaultAccountID", assetID).
		Return([]fireblocks.AssetAddress{
			{
				AssetID: assetID,
				Address: expectedSenderAddr,
			},
		}, nil)

	addr, err := w.SenderAddress(context.Background())
	assert.Nil(t, err)
	assert.Equal(t, expectedSenderAddr, addr.String())
}
