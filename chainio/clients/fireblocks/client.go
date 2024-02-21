package fireblocks

import (
	"bytes"
	"context"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/ethereum/go-ethereum/common"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type FireblocksTxID string

type FireblocksClient interface {
	// ContractCall makes a ContractCall request to the Fireblocks API.
	// It signs and broadcasts a transaction and returns the transaction ID and status.
	// ref: https://developers.fireblocks.com/reference/post_transactions
	ContractCall(ctx context.Context, body *ContractCallRequest) (*ContractCallResponse, error)
	// ListContracts makes a ListContracts request to the Fireblocks API
	// It returns a list of whitelisted contracts and their assets for the account.
	// This call is used to get the contract ID for a whitelisted contract, which is needed as destination account ID by NewContractCallRequest in a ContractCall
	// ref: https://developers.fireblocks.com/reference/get_contracts
	ListContracts(ctx context.Context) ([]WhitelistedContract, error)
}

type fireblocksClient struct {
	apiKey     string
	privateKey *rsa.PrivateKey
	baseURL    string
	timeout    time.Duration
	client     *http.Client
	logger     logging.Logger
}

type Asset struct {
	ID      string         `json:"id"`
	Status  string         `json:"status"`
	Address common.Address `json:"address"`
	Tag     string         `json:"tag"`
}

type ErrorResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type WhitelistedContract struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Assets []Asset `json:"assets"`
}

func NewFireblocksClient(apiKey string, secretKey []byte, baseURL string, timeout time.Duration, logger logging.Logger) (FireblocksClient, error) {
	c := http.Client{Timeout: timeout}
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(secretKey)
	if err != nil {
		return nil, fmt.Errorf("error parsing RSA private key: %w", err)
	}

	return &fireblocksClient{
		apiKey:     apiKey,
		privateKey: privateKey,
		baseURL:    baseURL,
		timeout:    timeout,
		client:     &c,
		logger:     logger,
	}, nil
}

// signJwt signs a JWT token for the Fireblocks API
// mostly copied from the Fireblocks example: https://github.com/fireblocks/developers-hub/blob/main/authentication_examples/go/test.go
func (f *fireblocksClient) signJwt(path string, bodyJson interface{}, durationSeconds int64) (string, error) {
	nonce := uuid.New().String()
	now := time.Now().Unix()
	expiration := now + durationSeconds

	bodyBytes, err := json.Marshal(bodyJson)
	if err != nil {
		return "", fmt.Errorf("error marshaling JSON: %w", err)
	}

	h := sha256.New()
	h.Write(bodyBytes)
	hashed := h.Sum(nil)

	claims := jwt.MapClaims{
		"uri":      path,
		"nonce":    nonce,
		"iat":      now,
		"exp":      expiration,
		"sub":      f.apiKey,
		"bodyHash": hex.EncodeToString(hashed),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tokenString, err := token.SignedString(f.privateKey)
	if err != nil {
		return "", fmt.Errorf("error signing token: %w", err)
	}

	return tokenString, nil
}

// makeRequest makes a request to the Fireblocks API
// mostly copied from the Fireblocks example: https://github.com/fireblocks/developers-hub/blob/main/authentication_examples/go/test.go
func (f *fireblocksClient) makeRequest(ctx context.Context, method, path string, body interface{}) ([]byte, error) {
	url, err := url.JoinPath(f.baseURL, path)
	if err != nil {
		return nil, fmt.Errorf("error joining URL path with %s and %s: %w", f.baseURL, path, err)
	}
	f.logger.Debug("making request to Fireblocks", "method", method, "url", url)
	var reqBodyBytes []byte
	if body != nil {
		var err error
		reqBodyBytes, err = json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("error marshaling request body: %w", err)
		}
	}

	token, err := f.signJwt(path, body, int64(f.timeout.Seconds()))
	if err != nil {
		return nil, fmt.Errorf("error signing JWT: %w", err)
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(reqBodyBytes))
	if err != nil {
		return nil, fmt.Errorf("error creating HTTP request: %w", err)
	}

	if method == "POST" {
		req.Header.Set("Content-Type", "application/json")
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("X-API-KEY", f.apiKey)

	resp, err := f.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending HTTP request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		var errResp ErrorResponse
		err = json.Unmarshal(respBody, &errResp)
		if err != nil {
			return nil, fmt.Errorf("error parsing error response: %w", err)
		}
		return nil, fmt.Errorf("error response (%d) from Fireblocks with code %d: %s", resp.StatusCode, errResp.Code, errResp.Message)
	}

	return respBody, nil
}
