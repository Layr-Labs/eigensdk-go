package chainio

//go:generate mockgen -destination=./mocks/avsRegistryContractsReader.go -package=mocks github.com/Layr-Labs/eigensdk-go/chainio/avsregistry AvsRegistryReader
//go:generate mockgen -destination=./mocks/avsRegistryContractsWriter.go -package=mocks github.com/Layr-Labs/eigensdk-go/chainio/avsregistry AvsRegistryWriter
//go:generate mockgen -destination=./mocks/elContractsReader.go -package=mocks github.com/Layr-Labs/eigensdk-go/chainio/elcontracts ELReader
//go:generate mockgen -destination=./mocks/elContractsWriter.go -package=mocks github.com/Layr-Labs/eigensdk-go/chainio/elcontracts ELWriter
//go:generate mockgen -destination=./mocks/elContractsSubscriber.go -package=mocks github.com/Layr-Labs/eigensdk-go/chainio/elcontracts ELSubscriber
//go:generate mockgen -destination=./mocks/elContractsClient.go -package=mocks github.com/Layr-Labs/eigensdk-go/chainio/clients ELContractsClient
//go:generate mockgen -destination=./mocks/erc20ContractClient.go -package=mocks github.com/Layr-Labs/eigensdk-go/chainio/clients ERC20ContractClient
//go:generate mockgen -destination=./mocks/avsRegistryContractClient.go -package=mocks github.com/Layr-Labs/eigensdk-go/chainio/clients AVSRegistryContractsClient
//go:generate mockgen -destination=./mocks/ethclient.go -package=mocks github.com/Layr-Labs/eigensdk-go/chainio/clients/eth EthClient

//go:generate mockgen -destination=./mocks/eventSubscription.go -package=mocks github.com/ethereum/go-ethereum/event Subscription
