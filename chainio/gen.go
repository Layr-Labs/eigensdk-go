package chainio

//go:generate mockgen -destination=./mocks/avsRegistryContractsReader.go -package=mocks github.com/Layr-Labs/eigensdk-go/chainio/clients/avsregistry AvsRegistryReader
//go:generate mockgen -destination=./mocks/avsRegistryContractsWriter.go -package=mocks github.com/Layr-Labs/eigensdk-go/chainio/clients/avsregistry AvsRegistryWriter
//go:generate mockgen -destination=./mocks/elContractsReader.go -package=mocks github.com/Layr-Labs/eigensdk-go/chainio/clients/elcontracts ELReader
//go:generate mockgen -destination=./mocks/elContractsWriter.go -package=mocks github.com/Layr-Labs/eigensdk-go/chainio/clients/elcontracts ELWriter
//go:generate mockgen -destination=./mocks/elContractsSubscriber.go -package=mocks github.com/Layr-Labs/eigensdk-go/chainio/clients/elcontracts ELSubscriber
//go:generate mockgen -destination=./mocks/ethclient.go -package=mocks github.com/Layr-Labs/eigensdk-go/chainio/clients/eth EthClient
//go:generate mockgen -destination=./mocks/eventSubscription.go -package=mocks github.com/ethereum/go-ethereum/event Subscription
