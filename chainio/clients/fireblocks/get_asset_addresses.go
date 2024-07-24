package fireblocks

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

type AssetAddress struct {
	AssetID           AssetID `json:"assetId"`
	Address           string  `json:"address"`
	Description       string  `json:"description"`
	Tag               string  `json:"tag"`
	Type              string  `json:"type"`
	CustomerRefID     string  `json:"customerRefId"`
	AddressFormat     string  `json:"addressFormat"`
	LegacyAddress     string  `json:"legacyAddress"`
	EnterpriseAddress string  `json:"enterpriseAddress"`
	BIP44AddressIndex int     `json:"bip44AddressIndex"`
	UserDefined       bool    `json:"userDefined"`
}

func (f *client) GetAssetAddresses(ctx context.Context, vaultID string, assetID AssetID) ([]AssetAddress, error) {
	f.logger.Debug("Fireblocks get asset addressees", "vaultID", vaultID, "assetID", assetID)
	var addresses []AssetAddress
	type paging struct {
		Before string `json:"before"`
		After  string `json:"after"`
	}
	var response struct {
		Addresses []AssetAddress `json:"addresses"`
		Paging    paging         `json:"paging"`
	}

	p := paging{}
	next := true
	for next {
		path := fmt.Sprintf("/v1/vault/accounts/%s/%s/addresses_paginated", vaultID, assetID)
		u, err := url.Parse(path)
		if err != nil {
			return addresses, fmt.Errorf("error parsing URL: %w", err)
		}
		q := u.Query()
		q.Set("before", p.Before)
		q.Set("after", p.After)
		u.RawQuery = q.Encode()

		res, err := f.makeRequest(ctx, "GET", u.String(), nil)
		if err != nil {
			return nil, fmt.Errorf("error making request: %w", err)
		}
		body := string(res)
		err = json.NewDecoder(strings.NewReader(body)).Decode(&response)
		if err != nil {
			return addresses, fmt.Errorf("error parsing response body: %s: %w", body, err)
		}

		addresses = append(addresses, response.Addresses...)
		p = response.Paging
		if p.After == "" {
			next = false
		}
	}

	return addresses, nil
}
