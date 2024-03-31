package fireblocks

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
)

type SetConfirmationThresholdRequest struct {
	NumOfConfirmations int `json:"numOfConfirmations"`
}

type SetConfirmationThresholdResponse struct {
	Success      bool     `json:"success"`
	Transactions []string `json:"transactions"`
}

func (f *client) SetConfirmationThreshold(ctx context.Context, txID string, threshold int) (bool, error) {
	f.logger.Debug("Fireblocks set confirmation threshold", "txID", txID, "threshold", threshold)
	url := fmt.Sprintf("/v1/transactions/%s/set_confirmation_threshold", txID)
	res, err := f.makeRequest(ctx, "POST", url, &SetConfirmationThresholdRequest{
		NumOfConfirmations: threshold,
	})
	if err != nil {
		return false, fmt.Errorf("error making request: %w", err)
	}
	var response SetConfirmationThresholdResponse
	err = json.NewDecoder(strings.NewReader(string(res))).Decode(&response)
	if err != nil {
		return false, fmt.Errorf("error parsing response body: %w", err)
	}
	f.logger.Debug("Fireblocks set confirmation threshold response", "response", string(res))

	return response.Success, nil
}
