package fireblocks

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
)

func (f *fireblocksClient) ListContracts(ctx context.Context) ([]WhitelistedContract, error) {
	var contracts []WhitelistedContract
	res, err := f.makeRequest(ctx, "GET", "/v1/contracts", nil)
	if err != nil {
		return contracts, fmt.Errorf("error making request: %w", err)
	}
	body := string(res)
	err = json.NewDecoder(strings.NewReader(body)).Decode(&contracts)
	if err != nil {
		return contracts, fmt.Errorf("error parsing response body: %s: %w", body, err)
	}

	return contracts, nil
}
