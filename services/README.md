# Services

eigensdk services are high level APIs that provide backend like functionality for avs nodes and aggregators processes. Only the pubkeycompendium service is currently implemented, but the goal is to provide the following suite of services as soon as possible:

- [BLS pubkey compendium service](./pubkeycompendium/)
  - this service simply indexes the [NewPubkeyRegistration](https://github.com/Layr-Labs/eigenlayer-contracts/blob/ba22dbc18c72514f3a4a8b7b7b46cb59fef29034/src/contracts/middleware/BLSPublicKeyCompendium.sol#L23) events and provides a single endpoint to query for the registered G1 and G2 pubkeys of a given operator address.
  - this service is needed for aggregators to get the pubkey of their registered operators so as to verify their signatures
- Registry service
  - this service will index the events on the [avs registry contracts](https://github.com/Layr-Labs/eigenlayer-contracts/tree/11a8da30919268ef22828e7e39088110cbf64611/src/contracts/middleware) and provide endpoints to query for the registered avs nodes, their stake at past blocks, operator id, etc.
  - this service is needed to be able to check that an eigenlayer operator was registered at a specific block and that the stake threshold had been met for an aggregate bls signature
- Signature aggregation service
  - this service will provide endpoints to aggregate operator signatures for various avs tasks
  - this service will aggregate signatures in the background and return an aggregated bls signature once it reached a threshold (this will require using the registry service to get operator stakes)
