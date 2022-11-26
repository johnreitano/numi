#!/usr/bin/env bash

set -e
# set -x

SCRIPT_DIR=$(dirname $(readlink -f $0))
cd ${SCRIPT_DIR}/..

VALIDATOR_IPS=$1
SEED_IPS=$2

# convert seed ips variable from a single string to an array
IFS=' ' read -r -a SEED_IPS <<< "${SEED_IPS}"

# create and initialize the the node config dirs
mkdir -p deploy/node-config
rm -rf deploy/node-config/seed*
for i in ${!SEED_IPS[@]}; do 
    MONIKER=seed${i}
    HOME_DIR=deploy/node-config/${MONIKER}
    numid init --chain-id numi-testnet-1 --home ${HOME_DIR} ${MONIKER}
    deploy/add-test-keys.sh ${HOME_DIR}
done

# update config files with external_address, persistent_peers and other values
P2P_PERSISTENT_PEERS=$(deploy/persistent-peers.sh "${VALIDATOR_IPS}")
for i in ${!SEED_IPS[@]}; do
  MONIKER=seed${i}
  HOME_DIR=deploy/node-config/${MONIKER}

  dasel put string -f ${HOME_DIR}/config/config.toml -p toml ".p2p.external_address" "tcp://${SEED_IPS[$i]}:26656"
  dasel put string -f ${HOME_DIR}/config/config.toml -p toml ".p2p.persistent_peers" "${P2P_PERSISTENT_PEERS}"
  dasel put string -f ${HOME_DIR}/config/config.toml -p toml ".rpc.laddr" "tcp://0.0.0.0:26657" # this allows extenal connections to the rpc port on the seed; should change this to 127.0.0.1 on prod
  dasel put string -f ${HOME_DIR}/config/config.toml -p toml ".rpc.tls_cert_file" "/home/ubuntu/cert/fullchain.pem"
  dasel put string -f ${HOME_DIR}/config/config.toml -p toml ".rpc.tls_key_file" "/home/ubuntu/cert/privkey.pem"
  dasel put bool -f ${HOME_DIR}/config/app.toml -p toml ".api.enable" true
  dasel put string -f ${HOME_DIR}/config/app.toml -p toml ".api.address" "tcp://localhost:1317"
done


# copy genesis file from primary validator node to seed nodes
PRIMARY_VALIDATOR_HOME_DIR=deploy/node-config/validator0
for i in ${!SEED_IPS[@]}; do
  MONIKER=seed${i}
  HOME_DIR=deploy/node-config/${MONIKER}
  cp ${PRIMARY_VALIDATOR_HOME_DIR}/config/genesis.json ${HOME_DIR}/config/
done

