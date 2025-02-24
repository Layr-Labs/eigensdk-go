package fireblocks

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/Layr-Labs/eigensdk-go/utils"
	"github.com/ethereum/go-ethereum/common"
)

type WhitelistedContract struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Assets []struct {
		ID      AssetID        `json:"id"`
		Status  string         `json:"status"`
		Address common.Address `json:"address"`
		Tag     string         `json:"tag"`
	} `json:"assets"`
}

func (f *client) ListContracts(ctx context.Context) ([]WhitelistedContract, error) {
	var contracts []WhitelistedContract
	res, err := f.makeRequest(ctx, "GET", "/v1/contracts", nil)
	if err != nil {
		return contracts, utils.WrapError("error making request", err)
	}
	body := string(res)
	err = json.NewDecoder(strings.NewReader(body)).Decode(&contracts)
	if err != nil {
		text := fmt.Sprintf("error parsing response body: %s", body)
		return contracts, utils.WrapError(text, err)
	}

	return contracts, nil
}
