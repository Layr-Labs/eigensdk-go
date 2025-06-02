package fireblocks

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/Layr-Labs/eigensdk-go/utils"
)

func NewContractCallRequest(
	externalTxID string,
	assetID AssetID,
	sourceAccountID string,
	destinationAccountID string,
	amount string, // amount in ETH
	calldata string,
	replaceTxByHash string,
	gasPrice string,
	gasLimit string,
	maxFee string,
	priorityFee string,
	feeLevel FeeLevel,
) *TransactionRequest {
	req := &TransactionRequest{
		Operation:    ContractCall,
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
		Amount: amount,
		ExtraParameters: extraParams{
			Calldata: calldata,
		},
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

func (f *client) ContractCall(ctx context.Context, req *TransactionRequest) (*TransactionResponse, error) {
	f.logger.Debug("Fireblocks call contract", "req", req)
	res, err := f.makeRequest(ctx, "POST", "/v1/transactions", req)
	if err != nil {
		return nil, utils.WrapError("error making request", err)
	}
	var response TransactionResponse
	err = json.NewDecoder(strings.NewReader(string(res))).Decode(&response)
	if err != nil {
		return nil, utils.WrapError("error parsing response body", err)
	}

	return &TransactionResponse{
		ID:     response.ID,
		Status: response.Status,
	}, nil
}
