package fireblocks

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
)

type TransactionOperation string

const (
	ContractCall TransactionOperation = "CONTRACT_CALL"
	Transfer     TransactionOperation = "TRANSFER"
	Mint         TransactionOperation = "MINT"
	Burn         TransactionOperation = "BURN"
	TypedMessage TransactionOperation = "TYPED_MESSAGE"
	Raw          TransactionOperation = "RAW"
)

type account struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

type extraParams struct {
	Calldata string `json:"contractCallData"`
}

type ContractCallRequest struct {
	Operation       TransactionOperation `json:"operation"`
	ExternalTxID    string               `json:"externalTxId"`
	AssetID         AssetID              `json:"assetId"`
	Source          account              `json:"source"`
	Destination     account              `json:"destination"`
	Amount          string               `json:"amount"`
	ExtraParameters extraParams          `json:"extraParameters"`
}

type ContractCallResponse struct {
	ID     string `json:"id"`
	Status Status `json:"status"`
}

func NewContractCallRequest(
	externalTxID string,
	assetID AssetID,
	sourceAccountID string,
	destinationAccountID string,
	amount string,
	calldata string,
) *ContractCallRequest {
	return &ContractCallRequest{
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
	}
}

func (f *client) ContractCall(ctx context.Context, req *ContractCallRequest) (*ContractCallResponse, error) {
	f.logger.Debug("Fireblocks call contract", "req", req)
	res, err := f.makeRequest(ctx, "POST", "/v1/transactions", req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	var response ContractCallResponse
	err = json.NewDecoder(strings.NewReader(string(res))).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("error parsing response body: %w", err)
	}

	return &ContractCallResponse{
		ID:     response.ID,
		Status: response.Status,
	}, nil
}
