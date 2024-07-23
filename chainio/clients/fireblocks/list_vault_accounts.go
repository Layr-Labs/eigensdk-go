package fireblocks

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

type Asset struct {
	ID        AssetID `json:"id"`
	Total     string  `json:"total"`
	Balance   string  `json:"balance"`
	Available string  `json:"available"`
}

type VaultAccount struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Assets []Asset `json:"assets"`
}

func (f *client) ListVaultAccounts(ctx context.Context) ([]VaultAccount, error) {
	var accounts []VaultAccount
	type paging struct {
		Before string `json:"before"`
		After  string `json:"after"`
	}
	var response struct {
		Accounts []VaultAccount `json:"accounts"`
		Paging   paging         `json:"paging"`
	}
	p := paging{}
	next := true
	for next {
		u, err := url.Parse("/v1/vault/accounts_paged")
		if err != nil {
			return accounts, fmt.Errorf("error parsing URL: %w", err)
		}
		q := u.Query()
		q.Set("before", p.Before)
		q.Set("after", p.After)
		u.RawQuery = q.Encode()
		res, err := f.makeRequest(ctx, "GET", u.String(), nil)
		if err != nil {
			return accounts, fmt.Errorf("error making request: %w", err)
		}
		body := string(res)
		err = json.NewDecoder(strings.NewReader(body)).Decode(&response)
		if err != nil {
			return accounts, fmt.Errorf("error parsing response body: %s: %w", body, err)
		}

		accounts = append(accounts, response.Accounts...)
		p = response.Paging
		if p.After == "" {
			next = false
		}
	}

	return accounts, nil
}
