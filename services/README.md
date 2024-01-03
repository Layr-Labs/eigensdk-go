# Services

eigensdk services are high level APIs that provide backend like functionality for avs nodes and aggregators processes. We provide the following suite of services as soon as possible:

- [operator pubkeys service](./operatorpubkeys/)
  - this service simply indexes the [NewPubkeyRegistration](https://github.com/Layr-Labs/eigenlayer-middleware/blob/9aa6eb543fe38db6e41516f89f15b654ad4d6bf4/src/interfaces/IBLSApkRegistry.sol#L38) events and provides a single endpoint to query for the registered G1 and G2 pubkeys of a given operator address.
  - this service is needed for aggregators to get the pubkey of their registered operators so as to verify their signatures
- Registry service
  - this service will index the events on the [avs registry contracts](https://github.com/Layr-Labs/eigenlayer-middleware) and provide endpoints to query for the registered avs nodes, their stake at past blocks, operator id, etc.
  - this service is needed to be able to check that an eigenlayer operator was registered at a specific block and that the stake threshold had been met for an aggregate bls signature
- Signature aggregation service
  - this service will provide endpoints to aggregate operator signatures for various avs tasks
  - this service will aggregate signatures in the background and return an aggregated bls signature once it reached a threshold (this will require using the registry service to get operator stakes)
