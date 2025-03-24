package signerv2

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Layr-Labs/eigensdk-go/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/google/uuid"
)

// JsonRpcRequest is a struct for JSON RPC 2.0 request
// See: https://www.jsonrpc.org/specification
type JsonRpcRequest struct {
	JsonRPC string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
	ID      string      `json:"id"`
}

type Web3SignerClient interface {
	SignTransaction(from common.Address, tx *types.Transaction) (*types.Transaction, error)
}

// Web3Signer is a client for a remote signer
// It currently implements `eth_signTransaction` method of Consensys Web3 Signer
// Reference: https://docs.web3signer.consensys.io/reference/api/json-rpc#eth_signtransaction
type Web3Signer struct {
	url    string
	client http.Client
}

func NewWeb3SignerClient(url string) Web3SignerClient {
	client := http.Client{}
	return &Web3Signer{client: client, url: url}
}

func (r Web3Signer) SignTransaction(
	from common.Address,
	tx *types.Transaction,
) (*types.Transaction, error) {
	method := "eth_signTransaction"
	id := uuid.New().String()
	params := []map[string]string{
		{
			"from":     from.Hex(),
			"to":       tx.To().Hex(),
			"gas":      utils.Add0x(fmt.Sprintf("%x", tx.Gas())),
			"gasPrice": utils.Add0x(hex.EncodeToString(tx.GasPrice().Bytes())),
			"nonce":    utils.Add0x(fmt.Sprintf("%x", tx.Nonce())),
			"data":     utils.Add0x(hex.EncodeToString(tx.Data())),
		},
	}

	request := JsonRpcRequest{
		JsonRPC: "2.0",
		Method:  method,
		Params:  params,
		ID:      id,
	}

	jsonData, err := json.Marshal(request)
	if err != nil {
		return nil, utils.WrapError("error marshalling request", err)
	}

	resp, err := r.client.Post(r.url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, utils.WrapError("error decoding response", err)
	}

	if result["error"] != nil {
		return nil, utils.WrapError("error in response", fmt.Errorf("%v", result["error"]))
	}

	rlpEncodedSignedTx := result["result"].(string)
	rlpEncodedSignedTx = utils.Trim0x(rlpEncodedSignedTx)
	signedTxBytes, err := hex.DecodeString(rlpEncodedSignedTx)
	if err != nil {
		return nil, err
	}
	var signedTx types.Transaction
	err = rlp.DecodeBytes(signedTxBytes, &signedTx)
	if err != nil {
		return nil, err
	}
	return &signedTx, nil
}
