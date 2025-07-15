# Integration Tests

We store an anvil state files in this directory, so that we can start an anvil chain with the correct state for integration tests.
```
anvil --load-state STATE_FILE.json
```

## Eigenlayer Deployment State File
`contracts-deployed-anvil-state.json` contains the eigenlayer Anvil deployment state.

It was created by running the `deploy-contracts-save-anvil-state.sh`.

```
./deploy-contracts-save-anvil-state.sh
```

See the main [README](../../README.md#dependencies) to understand why we deploy from the `experimental-reduce-strategy-manager-bytecode-size` branch of eigenlayer-contracts.

A handy way to get the Anvil state of a running Anvil server is using the `--dump-state` CLI parameter. Below is an example.

```
anvil --dump-state contracts-deployed-anvil-state.json
```
See [Anvil CLI](https://book.getfoundry.sh/reference/cli/anvil) for more CLI options.
