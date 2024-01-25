package chainio

//go:generate mockgen -destination=./mocks/avsEcdsaRegistryContractsReader.go -package=mocks github.com/Layr-Labs/eigensdk-go/chainio/clients/avsregistry/ecdsa AvsEcdsaRegistryReader
//go:generate mockgen -destination=./mocks/avsRegistryContractsReader.go -package=mocks github.com/Layr-Labs/eigensdk-go/chainio/clients/avsregistry/bls AvsRegistryReader
//go:generate mockgen -destination=./mocks/avsRegistryContractsSubscriber.go -package=mocks github.com/Layr-Labs/eigensdk-go/chainio/clients/avsregistry/bls AvsRegistrySubscriber
//go:generate mockgen -destination=./mocks/avsRegistryContractsWriter.go -package=mocks github.com/Layr-Labs/eigensdk-go/chainio/clients/avsregistry/bls AvsRegistryWriter
//go:generate mockgen -destination=./mocks/elContractsReader.go -package=mocks github.com/Layr-Labs/eigensdk-go/chainio/clients/elcontracts ELReader
//go:generate mockgen -destination=./mocks/elContractsWriter.go -package=mocks github.com/Layr-Labs/eigensdk-go/chainio/clients/elcontracts ELWriter
//go:generate mockgen -destination=./mocks/ethclient.go -package=mocks github.com/Layr-Labs/eigensdk-go/chainio/clients/eth EthClient
//go:generate mockgen -destination=./mocks/eventSubscription.go -package=mocks github.com/ethereum/go-ethereum/event Subscription
