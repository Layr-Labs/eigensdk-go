package chainio

//go:generate mockgen -destination=./mocks/ethclient.go -package=mocks -mock_names=Client=MockEthClient github.com/Layr-Labs/eigensdk-go/chainio/clients/eth Client
//go:generate mockgen -destination=./mocks/eventSubscription.go -package=mocks github.com/ethereum/go-ethereum/event Subscription
//go:generate mockgen -destination=./clients/mocks/fireblocks.go -package=mocks -mock_names=Client=MockFireblocksClient github.com/Layr-Labs/eigensdk-go/chainio/clients/fireblocks Client
//go:generate mockgen -destination=./clients/mocks/wallet.go -package=mocks github.com/Layr-Labs/eigensdk-go/chainio/clients/wallet Wallet
