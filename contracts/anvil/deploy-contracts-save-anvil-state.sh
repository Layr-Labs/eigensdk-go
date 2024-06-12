#!/bin/bash

# Enable the script to exit immediately if a command exits with a non-zero status
set -o errexit -o nounset -o pipefail

# Define your cleanup function
clean_up() {
    echo "Executing cleanup function..."
    set +e
    pkill -f anvil

    # Check if the exit status is non-zero
    exit_status=$?
    if [ $exit_status -ne 0 ]; then
        echo "Script exited due to set -e on line $1 with command '$2'. Exit status: $exit_status"
    fi
}
# Use trap to call the clean_up function when the script exits
trap 'clean_up $LINENO "$BASH_COMMAND"' EXIT

# cd to the directory of this script so that this can be run from anywhere
parent_path=$(
    cd "$(dirname "${BASH_SOURCE[0]}")"
    pwd -P
)
root_dir=$parent_path/../../..

set -a
source $parent_path/utils.sh
# we overwrite some variables here because should always deploy to anvil (localhost)
ETH_HTTP_URL=http://localhost:8545
set +a

# start an empty anvil chain in the background and dump its state to a json file upon exit
start_anvil_docker "" $parent_path/contracts-deployed-anvil-state.json
sleep 1

CHAIN_ID=$(cast chain-id)

# DEPLOY CONTRACT REGISTRY
# cd $parent_path/../../contracts
# forge create src/ContractsRegistry.sol:ContractsRegistry --rpc-url ${ETH_HTTP_URL:?} --private-key ${DEPLOYER_PRIVATE_KEY:?}

# DEPLOY EIGENLAYER
EIGEN_CONTRACTS_DIR=$parent_path/../../lib/eigenlayer-middleware/lib/eigenlayer-contracts
DEVNET_OUTPUT_DIR=$EIGEN_CONTRACTS_DIR/script/output/devnet
# deployment overwrites this file, so we save it as backup, because we want that output in our local files, and not in the eigenlayer-contracts submodule files
mv $DEVNET_OUTPUT_DIR/M2_from_scratch_deployment_data.json $DEVNET_OUTPUT_DIR/M2_from_scratch_deployment_data.json.bak
cd $EIGEN_CONTRACTS_DIR
forge script script/deploy/devnet/M2_Deploy_From_Scratch.s.sol --rpc-url http://localhost:8545 \
    --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 --broadcast \
    --sig "run(string memory configFileName)" -- M2_deploy_from_scratch.anvil.config.json
mv $DEVNET_OUTPUT_DIR/M2_from_scratch_deployment_data.json $root_dir/contracts/script/output/${CHAIN_ID:?}/eigenlayer_deployment_output.json
mv $DEVNET_OUTPUT_DIR/M2_from_scratch_deployment_data.json.bak $DEVNET_OUTPUT_DIR/M2_from_scratch_deployment_data.json
