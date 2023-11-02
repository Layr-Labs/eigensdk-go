#!/bin/bash

# cd to the directory of this script so that this can be run from anywhere
script_path=$(
    cd "$(dirname "${BASH_SOURCE[0]}")"
    pwd -P
)

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
    abigen --bin=data/tmp.bin --abi=data/tmp.abi --pkg=contract${contract} --out=$binding_dir/${contract}/binding.go
    rm -rf ../data/tmp.abi ../data/tmp.bin
}

EIGENLAYER_MIDDLEWARE_PATH=$script_path/lib/eigenlayer-middleware
cd $EIGENLAYER_MIDDLEWARE_PATH
forge build

avs_contracts="BLSRegistryCoordinatorWithIndices BLSOperatorStateRetriever StakeRegistry BLSPubkeyRegistry IBLSSignatureChecker BLSPublicKeyCompendium"
for contract in $avs_contracts; do
    create_binding . $contract ../../bindings
done

EIGENLAYER_CONTRACT_PATH=$EIGENLAYER_MIDDLEWARE_PATH/lib/eigenlayer-contracts
cd $EIGENLAYER_CONTRACT_PATH
forge build

el_contracts="DelegationManager Slasher StrategyManager IStrategy EigenPodManager EigenPod IERC20"
for contract in $el_contracts; do
    create_binding . $contract ../../../../bindings
done
