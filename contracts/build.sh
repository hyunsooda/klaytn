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

SC_INTERFACEID_TEST_ALL=sc_interfaceid_test/bridges/support_all_bridge
SC_INTERFACEID_TEST_ERC20=sc_interfaceid_test/bridges/support_only_erc20_bridge
SC_INTERFACEID_TEST_ERC721=sc_interfaceid_test/bridges/support_only_erc721_bridge
SC_INTERFACEID_TEST_NO_ANY=sc_interfaceid_test/bridges/no_support_any_bridge
SC_INTERFACEID_TEST_INCORRECT_EVENT_SIG=sc_interfaceid_test/bridges/incorrect_event_signature


SC_INTERFACEID_TEST_ALL_SOL=${SC_INTERFACEID_TEST_ALL}/Bridge.sol
SC_INTERFACEID_TEST_ERC20_SOL=${SC_INTERFACEID_TEST_ERC20}/Bridge.sol
SC_INTERFACEID_TEST_ERC721_SOL=${SC_INTERFACEID_TEST_ERC721}/Bridge.sol
SC_INTERFACEID_TEST_NO_ANY_SOL=${SC_INTERFACEID_TEST_NO_ANY}/Bridge.sol
SC_INTERFACEID_TEST_INCORRECT_EVENT_SIG_SOL=${SC_INTERFACEID_TEST_INCORRECT_EVENT_SIG}/Bridge.sol

SC_INTERFACEID_TEST_ALL_GO=${SC_INTERFACEID_TEST_ALL}/Bridge.go
SC_INTERFACEID_TEST_ERC20_GO=${SC_INTERFACEID_TEST_ERC20}/Bridge.go
SC_INTERFACEID_TEST_ERC721_GO=${SC_INTERFACEID_TEST_ERC721}/Bridge.go
SC_INTERFACEID_TEST_NO_ANY_GO=${SC_INTERFACEID_TEST_NO_ANY}/Bridge.go
SC_INTERFACEID_TEST_INCORRECT_EVENT_SIG_GO=${SC_INTERFACEID_TEST_INCORRECT_EVENT_SIG}/Bridge.go

# step1: Build contracts
function build_contract() {
    solc --allow-paths . --abi --bin sc_erc20/sc_token.sol -o $OUTPUT
    solc --allow-paths . --abi --bin sc_erc721/sc_nft.sol -o $OUTPUT
    solc --allow-paths . --abi --bin bridge/Bridge.sol -o $OUTPUT
    # test contracts
    solc --allow-paths . --abi --bin $SC_INTERFACEID_TEST_ALL_SOL -o $OUTPUT
    solc --allow-paths . --abi --bin $SC_INTERFACEID_TEST_ERC20_SOL -o $OUTPUT
    solc --allow-paths . --abi --bin $SC_INTERFACEID_TEST_ERC721_SOL -o $OUTPUT
    solc --allow-paths . --abi --bin $SC_INTERFACEID_TEST_NO_ANY_SOL -o $OUTPUT
    solc --allow-paths . --abi --bin $SC_INTERFACEID_TEST_INCORRECT_EVENT_SIG_SOL -o $OUTPUT
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

    # test contracts
    local INTERFACEID_TEST_ALL_CONTRACT_NAME=`cat $SC_INTERFACEID_TEST_ALL_SOL | grep -oh "contract [A-Za-z][A-Za-z0-9]*" | awk '{ print $2 }'`
    local INTERFACEID_TEST_ERC20_CONTRACT_NAME=`cat $SC_INTERFACEID_TEST_ERC20_SOL | grep -oh "contract [A-Za-z][A-Za-z0-9]*" | awk '{ print $2 }'`
    local INTERFACEID_TEST_ERC721_CONTRACT_NAME=`cat $SC_INTERFACEID_TEST_ERC721_SOL | grep -oh "contract [A-Za-z][A-Za-z0-9]*" | awk '{ print $2 }'`
    local INTERFACEID_TEST_NO_ANY_CONTRACT_NAME=`cat $SC_INTERFACEID_TEST_NO_ANY_SOL | grep -oh "contract [A-Za-z][A-Za-z0-9]*" | awk '{ print $2 }'`
    local INTERFACEID_TEST_INCORRECT_EVENT_SIG_CONTRACT_NAME=`cat $SC_INTERFACEID_TEST_INCORRECT_EVENT_SIG_SOL | grep -oh "contract [A-Za-z][A-Za-z0-9]*" | awk '{ print $2 }'`

    local BRIDGE_BINDING=bridge/Bridge.go
    local ERC20_BINDING=sc_erc20/sc_token.go
    local ERC721_BINDING=sc_erc721/sc_nft.go

    local INTERFACEID_TEST_ALL_BINDING=$SC_INTERFACEID_TEST_ALL_GO
    local INTERFACEID_TEST_ERC20_BINDING=$SC_INTERFACEID_TEST_ERC20_GO
    local INTERFACEID_TEST_ERC721_BINDING=$SC_INTERFACEID_TEST_ERC721_GO
    local INTERFACEID_TEST_NO_ANY_BINDING=$SC_INTERFACEID_TEST_NO_ANY_GO
    local INTERFACEID_TEST_INCORRECT_EVENT_SIG_BINDING=$SC_INTERFACEID_TEST_INCORRECT_EVENT_SIG_GO

    do_replace_bin $BRIDGE_CONTRACT_NAME $BRIDGE_BINDING ${OUTPUT}/${BRIDGE_CONTRACT_NAME}.bin
    do_replace_bin $ERC20_CONTRACT_NAME $ERC20_BINDING ${OUTPUT}/${ERC20_CONTRACT_NAME}.bin
    do_replace_bin $ERC721_CONTRACT_NAME $ERC721_BINDING ${OUTPUT}/${ERC721_CONTRACT_NAME}.bin

    do_replace_bin $INTERFACEID_TEST_ALL_CONTRACT_NAME $INTERFACEID_TEST_ALL_BINDING ${OUTPUT}/${INTERFACEID_TEST_ALL_CONTRACT_NAME}.bin
    do_replace_bin $INTERFACEID_TEST_ERC20_CONTRACT_NAME $INTERFACEID_TEST_ERC20_BINDING ${OUTPUT}/${INTERFACEID_TEST_ERC20_CONTRACT_NAME}.bin
    do_replace_bin $INTERFACEID_TEST_ERC721_CONTRACT_NAME $INTERFACEID_TEST_ERC721_BINDING ${OUTPUT}/${INTERFACEID_TEST_ERC721_CONTRACT_NAME}.bin
    do_replace_bin $INTERFACEID_TEST_NO_ANY_CONTRACT_NAME $INTERFACEID_TEST_NO_ANY_BINDING ${OUTPUT}/${INTERFACEID_TEST_NO_ANY_CONTRACT_NAME}.bin
    do_replace_bin $INTERFACEID_TEST_INCORRECT_EVENT_SIG_CONTRACT_NAME $INTERFACEID_TEST_INCORRECT_EVENT_SIG_BINDING ${OUTPUT}/${INTERFACEID_TEST_INCORRECT_EVENT_SIG_CONTRACT_NAME}.bin
}

build_contract
build_binding
replace_bin
