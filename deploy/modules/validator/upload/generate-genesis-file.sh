#!/usr/bin/env bash

set -x
set -e

THIS_NODE_ID=$(~/upload/numid tendermint show-node-id)
for f in ~/.numi/config/gentx/gentx-*.json; do
    base=$(basename ${f})
    if [[ "${base}" != "gentx-${THIS_NODE_ID}.json" ]]; then
        ADDRESS=$(cat ${f} | jq -r '.body.messages[0].delegator_address')
        AMOUNT=$(cat ${f} | jq -r '.body.messages[0].value.amount')
        DENOM=$(cat ${f} | jq -r '.body.messages[0].value.denom')
        ~/upload/numid add-genesis-account ${ADDRESS} ${AMOUNT}${DENOM} || :
    fi
done
~/upload/numid collect-gentxs
