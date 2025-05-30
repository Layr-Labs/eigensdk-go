package challenger

// Challenger configuration struct
type Config struct {
	// Ethereum WebSocket RPC URL to use for subscribing to on-chain events
	EthWsUrl string `toml:"eth_ws_url"`

	// Ethereum HTTP RPC URL to use for interacting with on-chain contracts
	EthHttpUrl string `toml:"eth_http_url"`
}
