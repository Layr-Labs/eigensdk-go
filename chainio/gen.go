package chainio

//go:generate mockgen -destination=./mocks/avsRegistryContractsReader.go -package=mocks -mock_names=Reader=MockAVSReader github.com/Layr-Labs/eigensdk-go/chainio/clients/avsregistry Reader
//go:generate mockgen -destination=./mocks/avsRegistryContractsSubscriber.go -package=mocks -mock_names=Subscriber=MockAVSSubscriber github.com/Layr-Labs/eigensdk-go/chainio/clients/avsregistry Subscriber
//go:generate mockgen -destination=./mocks/avsRegistryContractsWriter.go -package=mocks -mock_names=Writer=MockAVSWriter github.com/Layr-Labs/eigensdk-go/chainio/clients/avsregistry Writer
//go:generate mockgen -destination=./mocks/elContractsReader.go -package=mocks github.com/Layr-Labs/eigensdk-go/chainio/clients/elcontracts ELReader
//go:generate mockgen -destination=./mocks/elContractsWriter.go -package=mocks github.com/Layr-Labs/eigensdk-go/chainio/clients/elcontracts ELWriter
//go:generate mockgen -destination=./mocks/ethclient.go -package=mocks -mock_names=Client=MockEthClient github.com/Layr-Labs/eigensdk-go/chainio/clients/eth Client
//go:generate mockgen -destination=./mocks/eventSubscription.go -package=mocks github.com/ethereum/go-ethereum/event Subscription
//go:generate mockgen -destination=./clients/mocks/fireblocks.go -package=mocks -mock_names=Client=MockFireblocksClient github.com/Layr-Labs/eigensdk-go/chainio/clients/fireblocks Client
//go:generate mockgen -destination=./clients/mocks/wallet.go -package=mocks github.com/Layr-Labs/eigensdk-go/chainio/clients/wallet Wallet
