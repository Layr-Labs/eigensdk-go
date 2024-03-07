#!/bin/bash

# cd to the directory of this script so that this can be run from anywhere
script_path=$(
    cd "$(dirname "${BASH_SOURCE[0]}")"
    pwd -P
)

# build abigen-with-interfaces docker image if it doesn't exist
if [[ "$(docker images -q abigen-with-interfaces 2> /dev/null)" == "" ]]; then
    docker build -t abigen-with-interfaces -f abigen-with-interfaces.Dockerfile "$script_path"
fi

function extract_abi_bytecode {
    local contract_json="$1"
    local solc_abi
    local solc_bin

    solc_abi=$(cat "${contract_json}" | jq -r '.abi')
    solc_bin=$(cat "${contract_json}" | jq -r '.bytecode.object')

    mkdir -p data
    echo "${solc_abi}" > data/tmp.abi
    echo "${solc_bin}" > data/tmp.bin
}

function create_directories {
    local binding_dir="$1"
    local contract="$2"
    
    mkdir -p "${binding_dir}/${contract}"
    rm -f "${binding_dir}/${contract}/binding.go"
}

function run_docker_command {
    local binding_dir="$1"
    local contract="$2"

    docker run -v "$(realpath "${binding_dir}"):/home/binding_dir" -v .:/home/repo abigen-with-interfaces --bin=/home/repo/data/tmp.bin --abi=/home/repo/data/tmp.abi --pkg=contract${contract} --out=/home/binding_dir/${contract}/binding.go
    rm -rf data/tmp.abi data/tmp.bin
}

function create_binding {
    local contract_dir="$1"
    local contract="$2"
    local binding_dir="$3"
    local contract_json="${contract_dir}/out/${contract}.sol/${contract}.json"

    echo "${contract}"
    extract_abi_bytecode "${contract_json}"
    create_directories "${binding_dir}" "${contract}"
    run_docker_command "${binding_dir}" "${contract}"
}

EIGENLAYER_MIDDLEWARE_PATH="${script_path}/lib/eigenlayer-middleware"
cd "${EIGENLAYER_MIDDLEWARE_PATH}"
# you might want to run forge clean if the contracts have changed
forge build

avs_contracts="RegistryCoordinator OperatorStateRetriever StakeRegistry BLSApkRegistry IBLSSignatureChecker ServiceManagerBase"
for contract in ${avs_contracts}; do
    create_binding . "${contract}" ../../bindings
done

EIGENLAYER_CONTRACT_PATH="${EIGENLAYER_MIDDLEWARE_PATH}/lib/eigenlayer-contracts"
cd "${EIGENLAYER_CONTRACT_PATH}"
forge build

el_contracts="DelegationManager ISlasher StrategyManager EigenPod EigenPodManager IStrategy IERC20 AVSDirectory"
for contract in ${el_contracts}; do
    create_binding . "${contract}" ../../../../bindings
done
