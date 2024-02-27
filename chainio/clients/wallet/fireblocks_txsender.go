package wallet

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/fireblocks"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
)

var _ Wallet = (*fireblocksWallet)(nil)

type fireblocksWallet struct {
	fireblocksClient fireblocks.Client
	ethClient        eth.Client
	vaultAccountName string
	logger           logging.Logger
	chainID          *big.Int

	// caches
	account              *fireblocks.VaultAccount
	whitelistedContracts map[common.Address]*fireblocks.WhitelistedContract
}

func NewFireblocksWallet(fireblocksClient fireblocks.Client, ethClient eth.Client, vaultAccountName string, logger logging.Logger) (Wallet, error) {
	chainID, err := ethClient.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("error getting chain ID: %w", err)
	}
	return &fireblocksWallet{
		fireblocksClient: fireblocksClient,
		ethClient:        ethClient,
		vaultAccountName: vaultAccountName,
		logger:           logger,
		chainID:          chainID,

		// caches
		account:              nil,
		whitelistedContracts: make(map[common.Address]*fireblocks.WhitelistedContract),
	}, nil
}

func (t *fireblocksWallet) getAccount(ctx context.Context) (*fireblocks.VaultAccount, error) {
	if t.account == nil {
		accounts, err := t.fireblocksClient.ListVaultAccounts(ctx)
		if err != nil {
			return nil, fmt.Errorf("error listing vault accounts: %w", err)
		}
		for _, a := range accounts {
			if a.Name == t.vaultAccountName {
				t.account = &a
				break
			}
		}
	}
	return t.account, nil
}

func (t *fireblocksWallet) getWhitelistedContract(ctx context.Context, address common.Address) (*fireblocks.WhitelistedContract, error) {
	assetID, ok := fireblocks.AssetIDByChain[t.chainID.Uint64()]
	if !ok {
		return nil, fmt.Errorf("unsupported chain %d", t.chainID.Uint64())
	}
	contract, ok := t.whitelistedContracts[address]
	if !ok {
		contracts, err := t.fireblocksClient.ListContracts(ctx)
		if err != nil {
			return nil, fmt.Errorf("error listing contracts: %w", err)
		}
		for _, c := range contracts {
			for _, a := range c.Assets {
				if a.Address == address && a.Status == "ACTIVE" && a.ID == assetID {
					t.whitelistedContracts[address] = &c
					contract = &c
					break
				}
			}
		}
	}

	if contract == nil {
		return nil, fmt.Errorf("contract %s not found in whitelisted contracts", address.Hex())
	}
	return contract, nil
}

func (t *fireblocksWallet) SendTransaction(ctx context.Context, tx *types.Transaction) (TxID, error) {
	assetID, ok := fireblocks.AssetIDByChain[t.chainID.Uint64()]
	if !ok {
		return "", fmt.Errorf("unsupported chain %d", t.chainID.Uint64())
	}
	account, err := t.getAccount(ctx)
	if err != nil {
		return "", fmt.Errorf("error getting account: %w", err)
	}
	foundAsset := false
	for _, a := range account.Assets {
		if a.ID == assetID {
			if a.Available == "0" {
				return "", errors.New("insufficient funds")
			}
			foundAsset = true
			break
		}
	}
	if !foundAsset {
		return "", fmt.Errorf("asset %s not found in account %s", assetID, t.vaultAccountName)
	}

	contract, err := t.getWhitelistedContract(ctx, *tx.To())
	if err != nil {
		return "", fmt.Errorf("error getting whitelisted contract %s: %w", tx.To().Hex(), err)
	}
	req := fireblocks.NewContractCallRequest(
		tx.To().Hex(),
		assetID,
		account.ID,                // source account ID
		contract.ID,               // destination account ID
		tx.Value().String(),       // amount
		hexutil.Encode(tx.Data()), // calldata
	)
	res, err := t.fireblocksClient.ContractCall(ctx, req)
	if err != nil {
		return "", fmt.Errorf("error calling contract %s: %w", tx.To().Hex(), err)
	}

	return res.ID, nil
}

func (t *fireblocksWallet) GetTransactionReceipt(ctx context.Context, txID TxID) (*types.Receipt, error) {
	fireblockTx, err := t.fireblocksClient.GetTransaction(ctx, txID)
	if err != nil {
		return nil, fmt.Errorf("error getting fireblocks transaction %s: %w", txID, err)
	}
	if fireblockTx.Status == "COMPLETED" {
		txHash := common.HexToHash(fireblockTx.TxHash)
		receipt, err := t.ethClient.TransactionReceipt(ctx, txHash)
		if err == nil {
			return receipt, nil
		}
		if errors.Is(err, ethereum.NotFound) {
			return nil, fmt.Errorf("transaction receipt %s not yet available", txID)
		} else {
			return nil, fmt.Errorf("Transaction receipt retrieval failed: %w", err)
		}
	}

	return nil, fmt.Errorf("transaction %s not yet completed: status %s", txID, fireblockTx.Status)
}
