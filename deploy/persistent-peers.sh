#!/usr/bin/env bash

set -e
set -x

# load ips into array - convert arg from a string to an array
VALIDATOR_IPS=$1
IFS=' ' read -r -a VALIDATOR_IPS <<< "${VALIDATOR_IPS}"

P2P_PERSISTENT_PEERS=""
for i in ${!VALIDATOR_IPS[@]}; do 
    HOME_DIR=deploy/node-config/validator${i}
    VALIDATOR_NODE_ID=$(numid tendermint show-node-id --home ${HOME_DIR})
    P2P_PERSISTENT_PEERS="${P2P_PERSISTENT_PEERS}${VALIDATOR_NODE_ID}@${VALIDATOR_IPS[$i]}:26656,"
done
echo ${P2P_PERSISTENT_PEERS}
