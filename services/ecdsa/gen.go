package services

//go:generate mockgen -destination=./mocks/avsregistry.go -package=mocks github.com/Layr-Labs/eigensdk-go/services/ecdsa/avsregistry AvsRegistryService
//go:generate mockgen -destination=./mocks/aggregation.go -package=mocks github.com/Layr-Labs/eigensdk-go/services/ecdsa/aggregation EcdsaAggregationService
