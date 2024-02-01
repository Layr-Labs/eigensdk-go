package services

//go:generate mockgen -destination=./mocks/operatorpubkeys.go -package=mocks github.com/Layr-Labs/eigensdk-go/services/bls/operatorpubkeys OperatorPubkeysService
//go:generate mockgen -destination=./mocks/avsregistry.go -package=mocks github.com/Layr-Labs/eigensdk-go/services/bls/avsregistry AvsRegistryService

// We generate it in ./mocks/blsagg/ instead of ./mocks like the others because otherwise we get a circular dependency
// avsregistry -> mocks -> avsregistry
// because avsregistry_chaincaller_test -> for blsApkRegistry mock
// and blsaggregation mock -> avsregistry interface
// TODO: are there better ways to organize these dependencies? Maybe by using ben johnson
// and having the avs registry interface be in the /avsregistry dir but the avsregistry_chaincaller
// and its test in a subdir?
//go:generate mockgen -destination=./mocks/blsagg/aggregation.go -package=mocks github.com/Layr-Labs/eigensdk-go/services/bls/aggregation BlsAggregationService
