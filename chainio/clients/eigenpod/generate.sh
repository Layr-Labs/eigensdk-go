#!/bin/bash

# This script generates Go bindings for the EigenPods contracts
# The reason this script is separate from the bindings generation in the ~/contracts directory
# is because they are in a separate branch in eigenlayer-contracts and to unblock the development
# of the EigenPods client, we need to generate the bindings for the contracts in the feat/partial-withdrawal-batching branch.
# Once eigenlayer-contracts repo is stable and features are on single branch, we can move the bindings back to normal process

set -e

TMP_DIR=$(mktemp -d)
BINDINGS_DIR=bindings

cleanup_bindings_dir() {
  echo "Cleaning up the bindings directory"
  rm -rf ${BINDINGS_DIR}
}

clone() {
  echo "Cloning the EigenLayer contracts repository"
  git clone -b slashing-magnitudes --depth=1 https://github.com/Layr-Labs/eigenlayer-contracts.git "${TMP_DIR}"
}

generate_bindings() {
  echo "Generating bindings for the EigenPods contracts"
  current_dir=$(pwd)
  cd "${TMP_DIR}" && make bindings
  # shellcheck disable=SC2164
  cd "$current_dir"
  mkdir -p ${BINDINGS_DIR}
  generate_go "${TMP_DIR}"/out/IEigenPod.sol/IEigenPod.json IEigenPod
  generate_go "${TMP_DIR}"/out/IEigenPodManager.sol/IEigenPodManager.json IEigenPodManager
}

generate_go() {
  # Check if jq is installed
  if ! command -v jq &> /dev/null; then
      echo "jq is required but not installed. Please install jq (brew install jq)."
      exit 1
  fi

  temp_file=$(mktemp)
  echo "Generating Go bindings for the $1 contract"
  jq '.abi' "$1" > "$temp_file" && mv "$temp_file" "$1"

  abigen --abi "$1" --pkg bindings --type "$2" --out "$BINDINGS_DIR"/"$2".go
}

cleanup() {
  echo "Cleaning up the temporary directory"
  rm -rf "${TMP_DIR}"
}

main() {
  cleanup_bindings_dir
  clone
  generate_bindings
  cleanup
}

main
