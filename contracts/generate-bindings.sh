#!/bin/bash

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

EIGENLAYER_CONTRACT_PATH=./lib/eigenlayer-contracts
cd $EIGENLAYER_CONTRACT_PATH

forge clean
forge build

el_contracts="DelegationManager Slasher StrategyManager IStrategy BLSPublicKeyCompendium EigenPodManager EigenPod"
for contract in $el_contracts; do
    create_binding . $contract ./../../bindings
done

avs_contracts="BLSRegistryCoordinatorWithIndices BLSOperatorStateRetriever StakeRegistry BLSPubkeyRegistry IBLSSignatureChecker"
for contract in $avs_contracts; do
    create_binding . $contract ./../../bindings
done

create_binding . IERC20 ./../../bindings