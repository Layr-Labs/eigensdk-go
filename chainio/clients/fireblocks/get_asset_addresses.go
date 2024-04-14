package fireblocks

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
)

type AssetAddress struct {
	AssetID           AssetID `json:"assetId"`
	Address           string  `json:"address"`
	Tag               string  `json:"tag"`
	Description       string  `json:"description"`
	Type              string  `json:"type"`
	LegacyAddress     string  `json:"legacyAddress"`
	EnterpriseAddress string  `json:"enterpriseAddress"`
	UserDefined       bool    `json:"userDefined"`
}

func (f *client) GetAssetAddresses(ctx context.Context, vaultID string, assetID AssetID) ([]AssetAddress, error) {
	f.logger.Debug("Fireblocks get asset addressees", "vaultID", vaultID, "assetID", assetID)
	path := fmt.Sprintf("/v1/vault/accounts/%s/%s/addresses", vaultID, assetID)
	res, err := f.makeRequest(ctx, "GET", path, nil)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	var addresses []AssetAddress
	err = json.NewDecoder(strings.NewReader(string(res))).Decode(&addresses)
	if err != nil {
		return nil, fmt.Errorf("error parsing response body: %w", err)
	}

	return addresses, nil
}
