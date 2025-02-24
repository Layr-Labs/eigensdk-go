package fireblocks

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/Layr-Labs/eigensdk-go/utils"
)

type CancelTransactionResponse struct {
	Success bool `json:"success"`
}

func (f *client) CancelTransaction(ctx context.Context, txID string) (bool, error) {
	f.logger.Debug("Fireblocks cancel transaction", "txID", txID)
	path := fmt.Sprintf("/v1/transactions/%s/cancel", txID)
	res, err := f.makeRequest(ctx, "POST", path, nil)
	if err != nil {
		return false, utils.WrapError("error making request", err)
	}
	var response CancelTransactionResponse
	err = json.NewDecoder(strings.NewReader(string(res))).Decode(&response)
	if err != nil {
		return false, utils.WrapError("error parsing response body", err)
	}

	return response.Success, nil
}
