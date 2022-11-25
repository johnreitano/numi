#!/usr/bin/env bash
set -x
set -e

NODE_INDEX=$1

if [[ "${NODE_INDEX}" = "0" ]]; then
    MONIKER="red"
elif [[ "${NODE_INDEX}" = "1" ]]; then
    MONIKER="blue"
else
    MONIKER="green"
fi

if [[ "${NODE_INDEX}" != "0" ]]; then
    echo "sleeping for 5 seconds to give primary validator time to start up"
    sleep 5
fi

sleep 1
echo "about to start validator node ${MONIKER} with NODE_INDEX ${NODE_INDEX} and id $(~/upload/numid tendermint show-node-id)"
pkill numid || :
sleep 1
~/upload/numid start
sleep 1
