#!/usr/bin/env bash
set -x
set -e

NODE_INDEX=$1

if [[ "${NODE_INDEX}" = "0" ]]; then
    MONIKER="black"
elif [[ "${NODE_INDEX}" = "1" ]]; then
    MONIKER="white"
else
    MONIKER="gray"
fi

echo "About to start seed node ${MONIKER} with NODE_INDEX ${NODE_INDEX} and id $(~/upload/numid tendermint show-node-id)"
pkill numid || :
sleep 1
~/upload/numid start
sleep 1
