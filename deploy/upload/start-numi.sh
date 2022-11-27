#!/usr/bin/env bash

# set -x # echo commands
set -e # exit on failure

CHAIN_ID=$(dasel -f ~/.numi/config/client.toml -p toml ".chain-id")
MONIKER=$(dasel -f ~/.numi/config/config.toml -p toml ".moniker")
NODE_ID=$(numid tendermint show-node-id)

echo "about to start node ${MONIKER} on chain ${CHAIN_ID} with node id ${NODE_ID}"
pkill numid || :
sleep 1
numid start --log_level warn
sleep 1
