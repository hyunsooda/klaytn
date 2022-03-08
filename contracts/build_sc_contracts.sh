#!/bin/bash

SOLC_PATH=$1
ABIGEN_PATH=$2
OUTPUT=build

rm -rf $OUTPUT

if [[ -z $SOLC_PATH || -z $ABIGEN_PATH ]]; then
    echo "Usage $0 <SOLC_PATH> <ABIGEN_PATH>"
    exit 1
fi

export PATH=$PATH:$SOLC_PATH:$ABIGEN_PATH

# step1: Build contracts
function build_contract() {
    solc --allow-paths . --abi --bin sc_erc20/sc_token.sol -o $OUTPUT
    solc --allow-paths . --abi --bin sc_erc721/sc_nft.sol -o $OUTPUT
    solc --allow-paths . --abi --bin sc_kip7/sc_token.sol -o $OUTPUT
    solc --allow-paths . --abi --bin sc_kip17/sc_nft.sol -o $OUTPUT
    solc --allow-paths . --abi --bin bridge/Bridge.sol -o $OUTPUT
}


# step2: Build go-bulding handler
function build_binding() {
    set -e
    echo "Generate go-binding handler.."
    go generate
}

# step3: Replace
# TODO: What about push service chain related contract files into `sc` directory?
# For discussion, we leave it for now.

function do_replace_bin() {
    local TARGET_BRIDGE_CONTRACT_NAME=$1
    local TARGET_BINDING=$2
    local TARGET_BIN_FILE=$3
    local TARGET_BINDING_BIN=`cat $TARGET_BINDING | grep "var ${TARGET_BRIDGE_CONTRACT_NAME}Bin" | awk '{ print $4 }'`
    # Truncate 3 chars ("0x), 1 char(") at prefix and suffix, respectively.

    TARGET_BINDING_BIN=${TARGET_BINDING_BIN#\"0x}
    TARGET_BINDING_BIN=${TARGET_BINDING_BIN%\"}
    rm $TARGET_BIN_FILE; echo $TARGET_BINDING_BIN >> $TARGET_BIN_FILE
    truncate -s -1 $TARGET_BIN_FILE
}

function replace_bin() {
    local BRIDGE_CONTRACT_NAME=`cat bridge/Bridge.sol | grep -oh "contract [A-Za-z][A-Za-z0-9]*" | awk '{ print $2 }'`
    local ERC20_CONTRACT_NAME=`cat sc_erc20/sc_token.sol | grep -oh "contract [A-Za-z][A-Za-z0-9]*" | awk '{ print $2 }'`
    local ERC721_CONTRACT_NAME=`cat sc_erc721/sc_nft.sol | grep -oh "contract [A-Za-z][A-Za-z0-9]*" | awk '{ print $2 }'`
    local KIP7_CONTRACT_NAME=`cat sc_kip7/sc_token.sol | grep -oh "contract [A-Za-z][A-Za-z0-9]*" | awk '{ print $2 }'`
    local KIP17_CONTRACT_NAME=`cat sc_kip17/sc_nft.sol | grep -oh "contract [A-Za-z][A-Za-z0-9]*" | awk '{ print $2 }'`

    local BRIDGE_BINDING=bridge/Bridge.go
    local ERC20_BINDING=sc_erc20/sc_token.go
    local ERC721_BINDING=sc_erc721/sc_nft.go
    local KIP7_BINDING=sc_kip7/sc_token.go
    local KIP17_BINDING=sc_kip17/sc_nft.go

    do_replace_bin $BRIDGE_CONTRACT_NAME $BRIDGE_BINDING ${OUTPUT}/${BRIDGE_CONTRACT_NAME}.bin
    do_replace_bin $ERC20_CONTRACT_NAME $ERC20_BINDING ${OUTPUT}/${ERC20_CONTRACT_NAME}.bin
    do_replace_bin $ERC721_CONTRACT_NAME $ERC721_BINDING ${OUTPUT}/${ERC721_CONTRACT_NAME}.bin
    do_replace_bin $KIP7_CONTRACT_NAME $KIP7_BINDING ${OUTPUT}/${KIP7_CONTRACT_NAME}.bin
    do_replace_bin $KIP17_CONTRACT_NAME $KIP17_BINDING ${OUTPUT}/${KIP17_CONTRACT_NAME}.bin
}

build_contract
build_binding
replace_bin
