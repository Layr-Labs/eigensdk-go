#!/bin/bash

TMP_DIR=tmp
BINDINGS_DIR=bindings

cleanup_bindings_dir() {
  echo "Cleaning up the bindings directory"
  rm -rf ${BINDINGS_DIR}
}

create_tmp_dir() {
  echo "Creating a temporary directory"
  mkdir -p ${TMP_DIR}
}

clone() {
  echo "Cloning the EigenLayer contracts repository"
  git clone -b feat/partial-withdrawal-batching --depth=1 git@github.com:Layr-Labs/eigenlayer-contracts.git ${TMP_DIR}
}

generate_bindings() {
  echo "Generating bindings for the EigenPods contracts"
  cd ${TMP_DIR} && make bindings
  cd ..
  mkdir -p ${BINDINGS_DIR}
  generate_go ${TMP_DIR}/out/IEigenPod.sol/IEigenPod.json IEigenPod
  generate_go ${TMP_DIR}/out/IEigenPodManager.sol/IEigenPodManager.json IEigenPodManager
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
  rm -rf ${TMP_DIR}
}

main() {
  cleanup_bindings_dir
  create_tmp_dir
  clone
  generate_bindings
  cleanup
}

main