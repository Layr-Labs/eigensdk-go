package fireblocks

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

type WhitelistedAccount struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Assets []struct {
		ID           AssetID        `json:"id"`
		Balance      string         `json:"balance"`
		LockedAmount string         `json:"lockedAmount"`
		Status       string         `json:"status"`
		Address      common.Address `json:"address"`
		Tag          string         `json:"tag"`
	} `json:"assets"`
}

func (f *client) ListExternalWallets(ctx context.Context) ([]WhitelistedAccount, error) {
	var accounts []WhitelistedAccount
	res, err := f.makeRequest(ctx, "GET", "/v1/external_wallets", nil)
	if err != nil {
		return accounts, fmt.Errorf("error making request: %w", err)
	}
	body := string(res)
	err = json.NewDecoder(strings.NewReader(body)).Decode(&accounts)
	if err != nil {
		return accounts, fmt.Errorf("error parsing response body: %s: %w", body, err)
	}

	return accounts, nil
}
