package fireblocks

type TransactionOperation string
type FeeLevel string

const (
	ContractCall TransactionOperation = "CONTRACT_CALL"
	Transfer     TransactionOperation = "TRANSFER"
	Mint         TransactionOperation = "MINT"
	Burn         TransactionOperation = "BURN"
	TypedMessage TransactionOperation = "TYPED_MESSAGE"
	Raw          TransactionOperation = "RAW"

	FeeLevelHigh   FeeLevel = "HIGH"
	FeeLevelMedium FeeLevel = "MEDIUM"
	FeeLevelLow    FeeLevel = "LOW"
)

type account struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

type extraParams struct {
	Calldata string `json:"contractCallData"`
}

type TransactionRequest struct {
	Operation TransactionOperation `json:"operation"`
	// ExternalTxID is an optional field that can be used as an idempotency key.
	ExternalTxID    string      `json:"externalTxId,omitempty"`
	AssetID         AssetID     `json:"assetId"`
	Source          account     `json:"source"`
	Destination     account     `json:"destination"`
	Amount          string      `json:"amount,omitempty"`
	ExtraParameters extraParams `json:"extraParameters"`
	// In case a transaction is stuck, specify the hash of the stuck transaction to replace it
	// by this transaction with a higher fee, or to replace it with this transaction with a zero fee and drop it from
	// the blockchain.
	ReplaceTxByHash string `json:"replaceTxByHash,omitempty"`
	// GasPrice and GasLimit are the gas price and gas limit for the transaction.
	// If GasPrice is specified (non-1559), MaxFee and PriorityFee are not required.
	GasPrice string `json:"gasPrice,omitempty"`
	GasLimit string `json:"gasLimit,omitempty"`
	// MaxFee and PriorityFee are the maximum and priority fees for the transaction.
	// If the transaction is stuck, the Fireblocks platform will replace the transaction with a new one with a higher
	// fee.
	// These fields are required if FeeLevel is not specified.
	MaxFee      string `json:"maxFee,omitempty"`
	PriorityFee string `json:"priorityFee,omitempty"`
	// FeeLevel is the fee level for the transaction which Fireblocks estimates based on the current network conditions.
	// The fee level can be HIGH, MEDIUM, or LOW.
	// If MaxFee and PriorityFee are not specified, the Fireblocks platform will use the default fee level MEDIUM.
	// Ref: https://developers.fireblocks.com/docs/gas-estimation#estimated-network-fee
	FeeLevel FeeLevel `json:"feeLevel,omitempty"`
}

type TransactionResponse struct {
	ID     string   `json:"id"`
	Status TxStatus `json:"status"`
}
