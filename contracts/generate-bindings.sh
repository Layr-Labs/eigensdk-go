#!/bin/bash

set -o errexit -o nounset -o pipefail

# cd to the directory of this script so that this can be run from anywhere
script_path=$(
    cd "$(dirname "${BASH_SOURCE[0]}")"
    pwd -P
)

# build abigen-with-interfaces docker image if it doesn't exist
if [[ "$(docker images -q abigen-with-interfaces 2>/dev/null)" == "" ]]; then
    docker build -t abigen-with-interfaces -f abigen-with-interfaces.Dockerfile $script_path
fi

# TODO: refactor this function.. it got really ugly with the dockerizing of abigen
function create_binding {
    contract_dir=$1
    contract=$2
    binding_dir=$3
    echo $contract
    mkdir -p $binding_dir/${contract}
    contract_json="$contract_dir/out/${contract}.sol/${contract}.json"
    solc_abi=$(cat ${contract_json} | jq -r '.abi')
    solc_bin=$(cat ${contract_json} | jq -r '.bytecode.object')

    mkdir -p data
    echo ${solc_abi} >data/tmp.abi
    echo ${solc_bin} >data/tmp.bin

    rm -f $binding_dir/${contract}/binding.go
    docker run -v $(realpath $binding_dir):/home/binding_dir -v .:/home/repo abigen-with-interfaces --bin=/home/repo/data/tmp.bin --abi=/home/repo/data/tmp.abi --pkg=contract${contract} --out=/home/binding_dir/${contract}/binding.go
    rm -rf data/tmp.abi data/tmp.bin
}

cd $script_path
forge build
sdk_contracts="ContractsRegistry"
for contract in $sdk_contracts; do
    create_binding . $contract ./bindings
done

EIGENLAYER_MIDDLEWARE_PATH=$script_path/lib/eigenlayer-middleware
cd $EIGENLAYER_MIDDLEWARE_PATH
# you might want to run forge clean if the contracts have changed
forge build

# No idea why but ordering of the contracts matters here... when I move them around sometimes bindings fail
avs_contracts="RegistryCoordinator IndexRegistry OperatorStateRetriever StakeRegistry BLSApkRegistry IBLSSignatureChecker ServiceManagerBase"
for contract in $avs_contracts; do
    create_binding . $contract ../../bindings
done

EIGENLAYER_CONTRACT_PATH=$EIGENLAYER_MIDDLEWARE_PATH/lib/eigenlayer-contracts
cd $EIGENLAYER_CONTRACT_PATH
forge build

# No idea why but the ordering of the contracts matters, and for some orderings abigen fails...
el_contracts="DelegationManager IRewardsCoordinator ISlasher StrategyManager EigenPod EigenPodManager IStrategy IERC20 AVSDirectory"
for contract in $el_contracts; do
    create_binding . $contract ../../../../bindings
done
