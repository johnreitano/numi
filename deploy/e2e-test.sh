#!/usr/bin/env bash

set -e
set -x

SCRIPT_DIR=$(dirname $(readlink -f $0))
cd ${SCRIPT_DIR}/..

rm -rf ~/.numi
numid init --chain-id numi-local myclient 1>/dev/null 2>/dev/null
deploy/add-test-keys.sh ~/.numi 1>/dev/null 2>/dev/null

echo "Users in numi blockchain before adding user:"
numid query numi list-user --chain-id numi-testnet-1 --node https://testnet-seed-0-rpc.numi.oktryme.com:26657 --output json | jq '.user'

yes | numid keys delete newbie --keyring-backend test --home ~/.numi 2>/dev/null || :
numid keys add newbie --keyring-backend test --home ~/.numi >/dev/null 2>/dev/null
NEWBIE_ADDR=$(numid keys show newbie -a --keyring-backend test --home ~/.numi)

BOB_ADDR=$(numid keys show bob -a --keyring-backend test --home ~/.numi)
OLIVER_ADDR=$(numid keys show oliver -a --keyring-backend test --home ~/.numi)

USER_ID=$(uuidgen)

BIO="I love Numi ($(date +"at %Y-%h-%d %H:%M:%S %p"))"

numid tx numi create-and-verify-user ${USER_ID} "Newbie" "Newman" "USA" "California" "San Diego" "${BIO}" ${BOB_ADDR} ${NEWBIE_ADDR} --keyring-backend test --from ${OLIVER_ADDR} --chain-id numi-testnet-1 --node https://testnet-seed-0-rpc.numi.oktryme.com:26657

sleep 7

echo "\nUsers in numi blockchain after adding user:"
numid query numi list-user --chain-id numi-testnet-1 --node https://testnet-seed-0-rpc.numi.oktryme.com:26657 --output json | jq '.user'
