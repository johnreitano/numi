#!/usr/bin/env bash

set -e
set -x

SCRIPT_DIR=$(dirname $(readlink -f $0))
cd ${SCRIPT_DIR}/..

rm -rf ~/.numi
numid init --chain-id numi-local myclient
deploy/add-test-keys.sh ~/.numi

yes | numid keys delete newbie --keyring-backend test --home ~/.numi 2>/dev/null || :
numid keys add newbie --keyring-backend test --home ~/.numi

BOB_ADDR=$(numid keys show bob -a --keyring-backend test --home ~/.numi)
OLIVER_ADDR=$(numid keys show oliver -a --keyring-backend test --home ~/.numi)
NEWBIE_ADDR=$(numid keys show newbie -a --keyring-backend test --home ~/.numi)

echo BOB_ADDR=${BOB_ADDR}
echo OLIVER_ADDR=${OLIVER_ADDR}
echo NEWBIE_ADDR=${NEWBIE_ADDR}

numid query numi list-user --chain-id numi-testnet-1 --node https://testnet-seed-0-rpc.numi.oktryme.com:26657

numid tx numi create-and-verify-user 5450612c-4665-4926-8037-e80ec1eb8544 Newbie Newman USA California "San Diego" "I just signed up on Numi ($(date +"at %Y-%h-%d %H:%M:%S %p"))" ${BOB_ADDR} ${NEWBIE_ADDR} --keyring-backend test --from ${OLIVER_ADDR} --chain-id numi-testnet-1 --node https://testnet-seed-0-rpc.numi.oktryme.com:26657

sleep 7
ssh ubuntu@testnet-seed-0-rpc.numi.oktryme.com sudo journalctl -xeu numi.service | tail -300 > t.log


# numid query numi list-user --chain-id numi-testnet-1 --node https://testnet-seed-0-rpc.numi.oktryme.com:26657
