package fireblocks

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
)

func NewTransferRequest(
	externalTxID string,
	assetID AssetID,
	sourceAccountID string,
	destinationAccountID string,
	amount string, // amount in ETH
	replaceTxByHash string,
	gasPrice string,
	gasLimit string,
	maxFee string,
	priorityFee string,
	feeLevel FeeLevel,
) *TransactionRequest {
	req := &TransactionRequest{
		Operation:    Transfer,
		ExternalTxID: externalTxID,
		AssetID:      assetID,
		Source: account{
			Type: "VAULT_ACCOUNT",
			ID:   sourceAccountID,
		},
		// https://developers.fireblocks.com/reference/transaction-sources-destinations
		Destination: account{
			Type: "EXTERNAL_WALLET",
			ID:   destinationAccountID,
		},
		Amount:          amount,
		ReplaceTxByHash: replaceTxByHash,
		GasLimit:        gasLimit,
	}

	if maxFee != "" && priorityFee != "" {
		req.MaxFee = maxFee
		req.PriorityFee = priorityFee
	} else if gasPrice != "" {
		req.GasPrice = gasPrice
	} else {
		req.FeeLevel = feeLevel
	}

	return req
}

func (f *client) Transfer(ctx context.Context, req *TransactionRequest) (*TransactionResponse, error) {
	f.logger.Debug("Fireblocks transfer", "req", req)
	res, err := f.makeRequest(ctx, "POST", "/v1/transactions", req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	var response TransactionResponse
	err = json.NewDecoder(strings.NewReader(string(res))).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("error parsing response body: %w", err)
	}

	return &TransactionResponse{
		ID:     response.ID,
		Status: response.Status,
	}, nil
}
