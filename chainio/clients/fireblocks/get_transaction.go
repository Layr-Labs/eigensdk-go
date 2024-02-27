package fireblocks

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
)

// Transaction is a type for the transaction response from Fireblocks
type Transaction struct {
	ID                            string               `json:"id"`
	ExternalID                    string               `json:"externalId"`
	Status                        Status               `json:"status"`
	SubStatus                     string               `json:"subStatus"`
	TxHash                        string               `json:"txHash"`
	Operation                     TransactionOperation `json:"operation"`
	CreatedAt                     int64                `json:"createdAt"`
	LastUpdated                   int64                `json:"lastUpdated"`
	AssetID                       AssetID              `json:"assetId"`
	Source                        account              `json:"source"`
	SourceAddress                 string               `json:"sourceAddress"`
	Destination                   account              `json:"destination"`
	DestinationAddress            string               `json:"destinationAddress"`
	DestinationAddressDescription string               `json:"destinationAddressDescription"`
	DestinationTag                string               `json:"destinationTag"`
	AmountInfo                    struct {
		Amount          string `json:"amount"`
		RequestedAmount string `json:"requestedAmount"`
		NetAmount       string `json:"netAmount"`
		AmountUSD       string `json:"amountUSD"`
	} `json:"amountInfo"`
	FeeInfo struct {
		NetworkFee string `json:"networkFee"`
		ServiceFee string `json:"serviceFee"`
		GasPrice   string `json:"gasPrice"`
	} `json:"feeInfo"`
	FeeCurrency     string `json:"feeCurrency"`
	ExtraParameters struct {
		ContractCallData string `json:"contractCallData"`
	} `json:"extraParameters"`
	NumOfConfirmations int `json:"numOfConfirmations"`
	// The block hash and height of the block that this transaction was mined in.
	BlockInfo struct {
		BlockHeight string `json:"blockHeight"`
		BlockHash   string `json:"blockHash"`
	} `json:"blockInfo"`
}

func (f *client) GetTransaction(ctx context.Context, txID string) (*Transaction, error) {
	f.logger.Debug("Fireblocks get transaction", "txID", txID)
	url := fmt.Sprintf("/v1/transactions/%s", txID)
	res, err := f.makeRequest(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	var tx Transaction
	err = json.NewDecoder(strings.NewReader(string(res))).Decode(&tx)
	if err != nil {
		return nil, fmt.Errorf("error parsing response body: %w", err)
	}

	return &tx, nil
}
