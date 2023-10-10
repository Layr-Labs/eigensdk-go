package services

//go:generate mockgen -destination=./mocks/pubkeycompendium.go -package=mocks github.com/Layr-Labs/eigensdk-go/services/pubkeycompendium PubkeyCompendiumService
//go:generate mockgen -destination=./mocks/avsregistry.go -package=mocks github.com/Layr-Labs/eigensdk-go/services/avsregistry AvsRegistryService

// We generate it in ./mocks/blsagg/ instead of ./mocks like the others because otherwise we get a circular dependency
// avsregistry -> mocks -> avsregistry
// because avsregistry_chaincaller_test -> for pubkeycompendium mock
// and blsaggregation mock -> avsregistry interface
// TODO: are there better ways to organize these dependencies? Maybe by using ben johnson
// and having teh avs registry interface be in the /avsregistry dir but the avsregistry_chaincaller
// and its test in a subdir?
//go:generate mockgen -destination=./mocks/blsagg/blsaggregation.go -package=mocks github.com/Layr-Labs/eigensdk-go/services/bls_aggregation BlsAggregationService
