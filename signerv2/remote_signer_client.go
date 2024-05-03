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
)

type RpcRequest struct {
	JsonRPC string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
	ID      int         `json:"id"`
}

type RemoteSignerClient interface {
	SignTransaction(from common.Address, tx *types.Transaction) (*types.Transaction, error)
}

// RemoteSigner is a client for a remote signer
// It implements Consensys Web3 Signer
// Reference: https://docs.web3signer.consensys.io/reference/api/json-rpc
type RemoteSigner struct {
	url    string
	client http.Client
}

func NewRemoteSignerClient(url string) RemoteSignerClient {
	client := http.Client{}
	return &RemoteSigner{client: client, url: url}
}

func (r RemoteSigner) SignTransaction(
	from common.Address,
	tx *types.Transaction,
) (*types.Transaction, error) {
	method := "eth_signTransaction"
	id := 1
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

	request := RpcRequest{
		JsonRPC: "2.0",
		Method:  method,
		Params:  params,
		ID:      id,
	}

	jsonData, err := json.Marshal(request)
	if err != nil {
		fmt.Println("Error marshalling request:", err)
		return nil, err
	}

	req, err := http.NewRequest("POST", r.url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := r.client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		fmt.Println("Error decoding response:", err)
		return nil, err
	}

	if result["error"] != nil {
		fmt.Println("Error in response:", result["error"])
		return nil, fmt.Errorf("error in response")
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
