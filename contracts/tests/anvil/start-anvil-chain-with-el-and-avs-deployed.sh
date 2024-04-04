#!/bin/bash

# cd to the directory of this script so that this can be run from anywhere
parent_path=$(
    cd "$(dirname "${BASH_SOURCE[0]}")"
    pwd -P
)
root_dir=$parent_path/../..

set -a
source $parent_path/utils.sh
source $root_dir/.env
# we overwrite some variables here because should always deploy to anvil (localhost)
ETH_HTTP_URL=http://localhost:8545
set +a

# start an anvil instance in the background that has eigenlayer contracts deployed
# we start anvil in the background so that we can run the below script
start_anvil_docker $parent_path/eigenlayer-eigencert-eigenda-strategies-deployed-operators-registered-with-eigenlayer-anvil-state.json ""
docker 

cd $root_dir/contracts
# we need to restart the anvil chain at the correct block, otherwise the indexRegistry has a quorumUpdate at the block number
# at which it was deployed (aka quorum was created/updated), but when we start anvil by loading state file it starts at block number 0
# so calling getOperatorListAtBlockNumber reverts because it thinks there are no quorums registered at block 0
# advancing chain manually like this is a current hack until https://github.com/foundry-rs/foundry/issues/6679 is merged
# also not that it doesn't really advance by the correct number of blocks.. not sure why, so we just forward by a bunch of blocks that should be enough
forge script script/utils/Utils.sol --sig "advanceChainByNBlocks(uint256)" 500 --rpc-url ${ETH_HTTP_URL:?} --private-key ${DEPLOYER_PRIVATE_KEY:?} --broadcast
echo "current block-number:" $(cast block-number)

docker attach anvil