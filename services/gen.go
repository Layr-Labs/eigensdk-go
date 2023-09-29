package services

//go:generate mockgen -destination=./mocks/pubkeycompendium.go -package=mocks github.com/Layr-Labs/eigensdk-go/services/pubkeycompendium PubkeyCompendiumService
//go:generate mockgen -destination=./mocks/avsregistry.go -package=mocks github.com/Layr-Labs/eigensdk-go/services/avsregistry AvsRegistryService
